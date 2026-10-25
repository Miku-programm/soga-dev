#!/usr/bin/env bash
#
# soga 多实例管理脚本，安装后位于 /usr/bin/soga
#
# 行为与官方 soga-cmd v0.0.6（vaxilu/soga-cmd）保持一致，下载源改为本仓库:
#   default 实例: /etc/soga/soga.conf，systemd 服务 soga.service
#   <name>  实例: /etc/soga/<name>/soga.conf，systemd 服务 soga@<name>.service（模板 soga@.service）
#   实例清单与更新检查缓存: $HOME/.soga/config.json
#
# 环境变量:
#   SOGA_REPO     下载 release 的 GitHub 仓库，默认 Miku-programm/soga-dev
#   SOGA_BRANCH   管理脚本所在分支，默认 main
#   SOGA_TARBALL  soga update 时使用本地 soga-linux-<arch>.tar.gz，跳过下载

set -o pipefail

SCRIPT_VERSION="0.0.7"

SOGA_REPO="${SOGA_REPO:-Miku-programm/soga-dev}"
SOGA_BRANCH="${SOGA_BRANCH:-main}"
SOGA_TARBALL="${SOGA_TARBALL:-}"
RAW_BASE="https://raw.githubusercontent.com/${SOGA_REPO}/${SOGA_BRANCH}"
API_BASE="https://api.github.com/repos/${SOGA_REPO}/releases"

INSTALL_DIR="/usr/local/soga"
SOGA_BIN="${INSTALL_DIR}/soga"
CONF_ROOT="/etc/soga"
UNIT_DIR="/etc/systemd/system"
WANTS_DIR="${UNIT_DIR}/multi-user.target.wants"
CMD_PATH="/usr/bin/soga"
ACME_SH="/root/.acme.sh/acme.sh"
STORE_DIR="${HOME:-}/.soga"
STORE_FILE="${STORE_DIR}/config.json"
ATTACHMENTS=(blockList whiteList routes.toml dns.yml)

RED=$'\e[0;31m'
GREEN=$'\e[0;32m'
YELLOW=$'\e[0;33m'
RESET=$'\e[0m'
SEP='————————————————'
ST_TEXT=("${GREEN}已运行${RESET}" "${YELLOW}未运行${RESET}" "${RED}未安装${RESET}")
YES_TEXT="${GREEN}是${RESET}"
NO_TEXT="${RED}否${RESET}"

ERR=""

# ---------- 基础工具 ----------

fail() {
    ERR="$*"
    return 1
}

println_color() {
    printf '%s%s%s\n' "$1" "$2" "$RESET"
}

printf_color() {
    local color="$1" fmt="$2" s
    shift 2
    # shellcheck disable=SC2059
    printf -v s -- "$fmt" "$@"
    printf '%s%s%s\n' "$color" "$s" "$RESET"
}

# trim_to <变量名> <字符串>：去掉首尾空白
trim_to() {
    local _trim_v="$2"
    _trim_v="${_trim_v#"${_trim_v%%[![:space:]]*}"}"
    _trim_v="${_trim_v%"${_trim_v##*[![:space:]]}"}"
    printf -v "$1" '%s' "$_trim_v"
}

# read_line：读取一行到 LINE（已去掉首尾空白），EOF 时返回失败
read_line() {
    local l
    if ! IFS= read -r l; then
        LINE=""
        fail "EOF"
        return 1
    fi
    trim_to LINE "$l"
}

# confirm <提示> <默认 y|n>
confirm() {
    local prompt="$1" def="$2" hint="y/N" ans
    [[ "$def" == "y" ]] && hint="Y/n"
    printf '%s [%s]: ' "$prompt" "$hint"
    if ! read_line || [[ -z "$LINE" ]]; then
        ans="$def"
    else
        ans="${LINE,,}"
    fi
    [[ "$ans" == "y" || "$ans" == "yes" ]]
}

wait_enter() {
    printf '%s按回车继续...%s' "$YELLOW" "$RESET"
    read_line || true
}

# 去掉开头的一个 v
strip_v() {
    local s
    trim_to s "$1"
    printf '%s' "${s#v}"
}

# 数字开头的版本号统一显示为 v 前缀
with_v() {
    local t rest
    trim_to t "$1"
    [[ -z "$t" ]] && return 0
    rest="$t"
    while [[ "$rest" == [vV]* ]]; do rest="${rest:1}"; done
    [[ -z "$rest" ]] && return 0
    if [[ "$rest" == [0-9]* ]]; then
        printf 'v%s' "$rest"
    else
        printf '%s' "$t"
    fi
}

# version_greater a b：按 主.次.修订 比较（每段取开头的数字），a > b 返回 0
version_greater() {
    local a b i x y
    a="$(strip_v "$1")"
    b="$(strip_v "$2")"
    local -a pa pb
    IFS='.' read -r -a pa <<<"$a"
    IFS='.' read -r -a pb <<<"$b"
    for i in 0 1 2; do
        x="${pa[i]:-}"
        y="${pb[i]:-}"
        x="${x%%[!0-9]*}"
        y="${y%%[!0-9]*}"
        x=$((10#${x:-0}))
        y=$((10#${y:-0}))
        ((x > y)) && return 0
        ((x < y)) && return 1
    done
    return 1
}

now_ms() {
    local t sec frac
    t="$(date +%s.%N)"
    sec="${t%%.*}"
    frac="${t#*.}"
    [[ "$frac" =~ ^[0-9]+$ ]] || frac=0
    frac="${frac}000"
    echo $((sec * 1000 + 10#${frac:0:3}))
}

detect_arch() {
    case "$(uname -m)" in
        x86_64 | x64 | amd64) ARCH="amd64" ;;
        aarch64 | arm64) ARCH="arm64" ;;
        *) fail "不支持的架构: $(uname -m)" ;;
    esac
}

# json_flatten：把 JSON（stdin）展开为 "路径<TAB>值" 行，例如 /instances/0/name<TAB>node1
json_flatten() {
    awk '
    function emit(v,    p, k) {
        p = ""
        for (k = 1; k <= d; k++) p = p "/" (t[k] == "o" ? key[k] : idx[k])
        print p "\t" v
    }
    { s = s $0 "\n" }
    END {
        n = length(s); i = 1; d = 0
        while (i <= n) {
            c = substr(s, i, 1)
            if (c == " " || c == "\t" || c == "\n" || c == "\r") { i++; continue }
            if (c == "{") { d++; t[d] = "o"; key[d] = ""; ek[d] = 1; i++; continue }
            if (c == "[") { d++; t[d] = "a"; idx[d] = 0; i++; continue }
            if (c == "}" || c == "]") { d--; i++; continue }
            if (c == ":") { ek[d] = 0; i++; continue }
            if (c == ",") { if (t[d] == "o") ek[d] = 1; else idx[d]++; i++; continue }
            if (c == "\"") {
                v = ""; i++
                while (i <= n) {
                    c = substr(s, i, 1)
                    if (c == "\\") {
                        e = substr(s, i + 1, 1)
                        if (e == "n" || e == "t" || e == "r") { v = v " "; i += 2 }
                        else if (e == "u") {
                            h = tolower(substr(s, i + 2, 4))
                            if (h == "003c") v = v "<"
                            else if (h == "003e") v = v ">"
                            else if (h == "0026") v = v "&"
                            else if (h == "0027") v = v "\047"
                            else v = v "?"
                            i += 6
                        } else { v = v e; i += 2 }
                        continue
                    }
                    if (c == "\"") { i++; break }
                    v = v c; i++
                }
                if (t[d] == "o" && ek[d]) key[d] = v; else emit(v)
                continue
            }
            j = i
            while (j <= n) {
                c = substr(s, j, 1)
                if (c == "," || c == "}" || c == "]" || c == " " || c == "\t" || c == "\n" || c == "\r") break
                j++
            }
            emit(substr(s, i, j - i)); i = j
        }
    }'
}

json_str() {
    local s="$1"
    s="${s//\\/\\\\}"
    s="${s//\"/\\\"}"
    s="${s//$'\n'/\\n}"
    s="${s//$'\t'/\\t}"
    s="${s//$'\r'/\\r}"
    s="${s//</\\u003c}"
    s="${s//>/\\u003e}"
    s="${s//&/\\u0026}"
    printf '"%s"' "$s"
}

# ---------- 路径 ----------

conf_dir() {
    if [[ "$1" == "default" ]]; then echo "$CONF_ROOT"; else echo "$CONF_ROOT/$1"; fi
}

conf_path() {
    echo "$(conf_dir "$1")/soga.conf"
}

unit_of() {
    if [[ "$1" == "default" ]]; then echo "soga.service"; else echo "soga@$1.service"; fi
}

# ---------- 实例清单 $HOME/.soga/config.json ----------

INST_NAMES=()
INST_CONFS=()
INST_DESCS=()
CHANNEL=""
UC_SET=0
UC_LAST=0
UC_SOGA=""
UC_CMD=""

store_normalize() {
    local -a names=() confs=() descs=()
    local i n c
    for i in "${!INST_NAMES[@]}"; do
        trim_to n "${INST_NAMES[i]}"
        [[ "$n" == "default" ]] && continue
        c="${INST_CONFS[i]:-}"
        if [[ -n "${INST_NAMES[i]}" ]]; then
            if [[ -z "$c" || "$c" == "$CONF_ROOT/${INST_NAMES[i]}.conf" ]]; then
                c="$(conf_path "${INST_NAMES[i]}")"
            fi
        fi
        names+=("${INST_NAMES[i]}")
        confs+=("$c")
        descs+=("${INST_DESCS[i]:-}")
    done
    INST_NAMES=("${names[@]}")
    INST_CONFS=("${confs[@]}")
    INST_DESCS=("${descs[@]}")
}

# 首次运行（尚无 config.json）时，把 /etc/soga 下已有的实例目录同步进清单
store_import_existing() {
    local d n
    [[ -d "$CONF_ROOT" ]] || return 0
    for d in "$CONF_ROOT"/*/; do
        [[ -f "${d}soga.conf" ]] || continue
        n="${d%/}"
        n="${n##*/}"
        [[ "$n" =~ ^[A-Za-z0-9]+$ && "$n" != "default" ]] || continue
        INST_NAMES+=("$n")
        INST_CONFS+=("$(conf_path "$n")")
        INST_DESCS+=("")
    done
}

store_load() {
    INST_NAMES=()
    INST_CONFS=()
    INST_DESCS=()
    CHANNEL=""
    UC_SET=0
    UC_LAST=0
    UC_SOGA=""
    UC_CMD=""
    [[ -n "${HOME:-}" ]] || fail "\$HOME is not defined" || return 1
    if [[ ! -e "$STORE_FILE" ]]; then
        store_import_existing
        store_normalize
        return 0
    fi
    [[ -r "$STORE_FILE" ]] || fail "open $STORE_FILE: permission denied" || return 1
    local path val i max=-1
    local -A n=() c=() s=()
    while IFS=$'\t' read -r path val; do
        case "$path" in
            /instances/*/name | /instances/*/conf | /instances/*/desc)
                i="${path#/instances/}"
                i="${i%%/*}"
                [[ "$i" =~ ^[0-9]+$ ]] || continue
                ((i > max)) && max=$i
                case "$path" in
                    */name) n[$i]="$val" ;;
                    */conf) c[$i]="$val" ;;
                    */desc) s[$i]="$val" ;;
                esac
                ;;
            /channel) CHANNEL="$val" ;;
            /update_check/last_check_at_ms)
                UC_SET=1
                [[ "$val" =~ ^[0-9]+$ ]] && UC_LAST="$val"
                ;;
            /update_check/latest_soga) UC_SET=1 && UC_SOGA="$val" ;;
            /update_check/latest_cmd) UC_SET=1 && UC_CMD="$val" ;;
        esac
    done < <(json_flatten <"$STORE_FILE")
    for ((i = 0; i <= max; i++)); do
        INST_NAMES+=("${n[$i]:-}")
        INST_CONFS+=("${c[$i]:-}")
        INST_DESCS+=("${s[$i]:-}")
    done
    store_normalize
}

# 本进程未读到更新检查结果时，沿用磁盘上（可能由后台检查刚写入）的结果
store_keep_update_check() {
    local path val
    ((UC_SET)) && return 0
    [[ -r "$STORE_FILE" ]] || return 0
    while IFS=$'\t' read -r path val; do
        case "$path" in
            /update_check/last_check_at_ms) [[ "$val" =~ ^[0-9]+$ ]] && UC_SET=1 && UC_LAST="$val" ;;
            /update_check/latest_soga) UC_SET=1 && UC_SOGA="$val" ;;
            /update_check/latest_cmd) UC_SET=1 && UC_CMD="$val" ;;
        esac
    done < <(json_flatten <"$STORE_FILE")
}

store_save() {
    store_normalize
    store_keep_update_check
    local out i
    out=$'{\n  "instances": '
    if ((${#INST_NAMES[@]} == 0)); then
        out+="null"
    else
        out+="["
        for i in "${!INST_NAMES[@]}"; do
            ((i > 0)) && out+=","
            out+=$'\n    {\n      "name": '"$(json_str "${INST_NAMES[i]}")"
            out+=$',\n      "conf": '"$(json_str "${INST_CONFS[i]}")"
            out+=$',\n      "desc": '"$(json_str "${INST_DESCS[i]}")"
            out+=$'\n    }'
        done
        out+=$'\n  ]'
    fi
    [[ -n "$CHANNEL" ]] && out+=$',\n  "channel": '"$(json_str "$CHANNEL")"
    if ((UC_SET)); then
        out+=$',\n  "update_check": {\n    "last_check_at_ms": '"$UC_LAST"
        [[ -n "$UC_SOGA" ]] && out+=$',\n    "latest_soga": '"$(json_str "$UC_SOGA")"
        [[ -n "$UC_CMD" ]] && out+=$',\n    "latest_cmd": '"$(json_str "$UC_CMD")"
        out+=$'\n  }'
    fi
    out+=$'\n}'
    if ! mkdir -p "$STORE_DIR"; then
        fail "mkdir $STORE_DIR 失败"
        return 1
    fi
    chmod 700 "$STORE_DIR" 2>/dev/null
    if ! (umask 077 && printf '%s' "$out" >"$STORE_FILE"); then
        fail "open $STORE_FILE: 写入失败"
        return 1
    fi
    chmod 600 "$STORE_FILE" 2>/dev/null
    return 0
}

# find_instance <name>：找到返回 0，FOUND_CONF 为清单中记录的配置路径
find_instance() {
    local n i
    trim_to n "$1"
    if [[ "$n" == "default" ]]; then
        FOUND_CONF="$CONF_ROOT/soga.conf"
        return 0
    fi
    for i in "${!INST_NAMES[@]}"; do
        if [[ "${INST_NAMES[i]}" == "$n" ]]; then
            FOUND_CONF="${INST_CONFS[i]}"
            return 0
        fi
    done
    return 1
}

require_instance() {
    local n
    trim_to n "$1"
    find_instance "$n" || fail "无此实例 \"$n\""
}

store_add_instance() {
    local n
    trim_to n "$1"
    [[ -z "$n" || "$n" == "default" ]] && return 1
    find_instance "$n" && return 1
    INST_NAMES+=("$n")
    INST_CONFS+=("$(conf_path "$n")")
    INST_DESCS+=("")
    return 0
}

store_remove_instance() {
    local n i
    trim_to n "$1"
    [[ "$n" == "default" ]] && return 0
    for i in "${!INST_NAMES[@]}"; do
        if [[ "${INST_NAMES[i]}" == "$n" ]]; then
            unset 'INST_NAMES[i]' 'INST_CONFS[i]' 'INST_DESCS[i]'
        fi
    done
    INST_NAMES=("${INST_NAMES[@]}")
    INST_CONFS=("${INST_CONFS[@]}")
    INST_DESCS=("${INST_DESCS[@]}")
}

store_rename_instance() {
    local old new i
    trim_to old "$1"
    trim_to new "$2"
    [[ -z "$old" || -z "$new" || "$new" == "default" ]] && return 1
    find_instance "$new" && return 1
    for i in "${!INST_NAMES[@]}"; do
        if [[ "${INST_NAMES[i]}" == "$old" ]]; then
            INST_NAMES[i]="$new"
            INST_CONFS[i]="$(conf_path "$new")"
            return 0
        fi
    done
    return 1
}

# 菜单中显示的配置路径（清单中记录的优先）
menu_conf_path() {
    if [[ "$1" != "default" ]] && find_instance "$1" && [[ -n "$FOUND_CONF" ]]; then
        echo "$FOUND_CONF"
    else
        conf_path "$1"
    fi
}

# ---------- soga.conf 读写（保留原有格式与注释） ----------

# conf_get <file> <key>：输出第一个同名配置的值（去掉首尾空白与引号）
conf_get() {
    [[ -f "$1" ]] || return 1
    K="$2" awk '
    BEGIN { k = ENVIRON["K"]; found = 0 }
    found { next }
    {
        line = $0; t = line; sub(/^[ \t]+/, "", t)
        p = index(line, "=")
        if (t == "" || substr(t, 1, 1) == "#" || substr(t, 1, 1) == ";" || p == 0) next
        kk = substr(line, 1, p - 1); gsub(/^[ \t\r\v\f]+|[ \t\r\v\f]+$/, "", kk)
        if (kk != k) next
        v = substr(line, p + 1)
        gsub(/^[ \t\r\v\f]+|[ \t\r\v\f]+$/, "", v)
        gsub(/^["\047]+|["\047]+$/, "", v)
        print v; found = 1
    }
    END { exit found ? 0 : 1 }' "$1"
}

# conf_set <file> <key> <value>：修改所有同名配置；不存在时追加（保持文件结尾换行）
conf_set() {
    local f="$1" endnl=0 tmp rc
    [[ -s "$f" && -z "$(tail -c 1 "$f")" ]] && endnl=1
    tmp="$(mktemp)" || return 1
    if ! K="$2" V="$3" ENDNL="$endnl" awk '
        BEGIN { k = ENVIRON["K"]; v = ENVIRON["V"]; found = 0; n = 0 }
        {
            line = $0; t = line; sub(/^[ \t]+/, "", t)
            p = index(line, "=")
            if (t != "" && substr(t, 1, 1) != "#" && substr(t, 1, 1) != ";" && p > 0) {
                kk = substr(line, 1, p - 1); gsub(/^[ \t\r\v\f]+|[ \t\r\v\f]+$/, "", kk)
                if (kk == k) {
                    rest = substr(line, p + 1); match(rest, /^[ \t]*/)
                    line = substr(line, 1, p) substr(rest, 1, RLENGTH) v
                    found = 1
                }
            }
            L[++n] = line
        }
        END {
            if (!found) L[++n] = k "=" v
            out = ""
            for (i = 1; i <= n; i++) out = out (i > 1 ? "\n" : "") L[i]
            if (ENVIRON["ENDNL"] == "1") out = out "\n"
            printf "%s", out
        }' "$f" >"$tmp"; then
        rm -f "$tmp"
        return 1
    fi
    cat "$tmp" >"$f"
    rc=$?
    rm -f "$tmp"
    return $rc
}

# set_instance_dir <file> <dir>：日志与 IP 缓存目录指向实例目录
set_instance_dir() {
    local d="${2%/}/" k
    for k in log_file_dir ip_user_cache_save_dir; do
        conf_set "$1" "$k" "$d" || return 1
    done
}

# relocate_instance_dir <file> <new_dir> <old_dir...>
relocate_instance_dir() {
    local f="$1" new="${2%/}/" k v o
    shift 2
    for k in log_file_dir ip_user_cache_save_dir; do
        v="$(conf_get "$f" "$k")" || continue
        trim_to v "$v"
        v="${v%/}"
        for o in "$@"; do
            if [[ "$v" == "${o%/}" ]]; then
                conf_set "$f" "$k" "$new" || return 1
                break
            fi
        done
    done
}

SC_KEYS=(type server_type node_id soga_key api webapi_url webapi_key db_host db_port db_name db_user db_password)
declare -A SC=()

# format_conf <dir>：用 SC 中的值生成完整 soga.conf
format_conf() {
    local d="$1"
    [[ "$d" == */ ]] || d="$d/"
    cat <<EOF
# 基础配置
type=${SC[type]:-}
server_type=${SC[server_type]:-}
node_id=${SC[node_id]:-}
soga_key=${SC[soga_key]:-}

# webapi 或 db 对接任选一个
api=${SC[api]:-}

# webapi 对接信息
webapi_url=${SC[webapi_url]:-}
webapi_key=${SC[webapi_key]:-}

# db 对接信息
db_host=${SC[db_host]:-}
db_port=${SC[db_port]:-}
db_name=${SC[db_name]:-}
db_user=${SC[db_user]:-}
db_password=${SC[db_password]:-}

# 日志文件夹，建议不要修改
log_file_dir=${d}

# 用户 ip 缓存文件夹，建议不要修改
ip_user_cache_save_dir=${d}

# 更多配置项请看文档根据需求自行添加
EOF
}

# load_parsed <name>：把实例配置的基础项读入 SC，文件存在返回 0
load_parsed() {
    local f k v
    SC=()
    f="$(conf_path "$1")"
    [[ -f "$f" ]] || return 1
    for k in "${SC_KEYS[@]}"; do
        if v="$(conf_get "$f" "$k")"; then SC[$k]="$v"; fi
    done
    return 0
}

default_conf_content() {
    SC=([type]=xboard [server_type]=ss [api]=webapi)
    format_conf "$(conf_dir "$1")"
}

ensure_attachments() {
    local d f
    d="$(conf_dir "$1")"
    mkdir -p "$d" || fail "mkdir $d 失败" || return 1
    for f in "${ATTACHMENTS[@]}"; do
        [[ -e "$d/$f" ]] && continue
        : >"$d/$f" || fail "open $d/$f: 创建失败" || return 1
    done
}

# ---------- systemd ----------

run_cmd() {
    "$@"
    local rc=$?
    ((rc == 0)) || fail "exit status $rc"
}

svc_run() {
    run_cmd systemctl "$1" "$(unit_of "$2")"
}

svc_reset_failed() {
    systemctl reset-failed "$(unit_of "$1")" >/dev/null 2>&1
}

svc_daemon_reload() {
    run_cmd systemctl daemon-reload
}

svc_show_status() {
    run_cmd systemctl status "$(unit_of "$1")" --no-pager -l
}

# svc_show_log <name> <lines> <follow 0|1>
svc_show_log() {
    if [[ "$3" == "1" ]]; then
        run_cmd journalctl -u "$(unit_of "$1")" -e --no-pager -f -n "$2"
    else
        run_cmd journalctl -u "$(unit_of "$1")" -e --no-pager -n "$2" </dev/null
    fi
}

# svc_status <name>：输出 0 运行中，1 未运行，2 未安装
svc_status() {
    local out state
    if ! out="$(systemctl show "$(unit_of "$1")" -p ActiveState 2>/dev/null)"; then
        echo 2
        return
    fi
    state="$(printf '%s\n' "$out" | sed -n 's/^[[:space:]]*ActiveState=//p' | head -n 1)"
    trim_to state "$state"
    if [[ -z "$state" ]]; then
        echo 2
    elif [[ "$state" == "active" ]]; then
        echo 0
    else
        echo 1
    fi
}

svc_enabled() {
    local out
    out="$(systemctl is-enabled "$(unit_of "$1")" 2>/dev/null)" || return 1
    trim_to out "$out"
    [[ "$out" == "enabled" || "$out" == "enabled-runtime" ]]
}

exec_start() {
    local d="$1"
    printf '%s -c %s/soga.conf -b %s/blockList -w %s/whiteList -r %s/routes.toml -d %s/dns.yml' \
        "$SOGA_BIN" "$d" "$d" "$d" "$d" "$d"
}

unit_content() {
    cat <<EOF
[Unit]
Description=soga Service
After=network.target nss-lookup.target
Wants=network.target

[Service]
User=root
Group=root
Type=simple
LimitAS=infinity
LimitRSS=infinity
LimitCORE=infinity
LimitNOFILE=999999
WorkingDirectory=/usr/local/soga/
ExecStart=$(exec_start "$1")
Restart=on-failure
RestartSec=5
StartLimitInterval=100s
StartLimitBurst=3

[Install]
WantedBy=multi-user.target
EOF
}

UNIT_CHANGED=0

# ensure_unit_file <path> <conf_dir>：内容不同时才写入
ensure_unit_file() {
    local content
    content="$(unit_content "$2")"$'\n'
    if [[ -f "$1" ]] && [[ "$(
        cat "$1"
        printf x
    )" == "${content}x" ]]; then
        return 0
    fi
    if ! mkdir -p "${1%/*}" || ! printf '%s' "$content" >"$1"; then
        fail "写入 $1 失败"
        return 1
    fi
    UNIT_CHANGED=1
}

ensure_default_unit() {
    ensure_unit_file "$UNIT_DIR/soga.service" "$CONF_ROOT"
}

ensure_template_unit() {
    ensure_unit_file "$UNIT_DIR/soga@.service" "$CONF_ROOT/%i"
}

# 旧版为单个实例单独生成的 soga@<name>.service 普通文件才会被改写
ensure_instance_unit() {
    local f="$UNIT_DIR/soga@$1.service"
    [[ -z "$1" || "$1" == "default" ]] && return 0
    [[ -L "$f" || ! -f "$f" ]] && return 0
    ensure_unit_file "$f" "$CONF_ROOT/$1"
}

remove_instance_unit() {
    local u
    u="$(unit_of "$1")"
    rm -f "$WANTS_DIR/$u"
    rm -f "$UNIT_DIR/$u"
}

# scan_units：从 systemd 中找出所有 soga 实例
scan_units() {
    local f b u
    for f in "$UNIT_DIR"/soga@*.service "$WANTS_DIR"/soga@*.service; do
        [[ -e "$f" || -L "$f" ]] || continue
        b="${f##*/}"
        [[ "$b" == "soga@.service" ]] && continue
        b="${b#soga@}"
        b="${b%.service}"
        [[ -n "$b" ]] && echo "$b"
    done
    if [[ -e "$UNIT_DIR/soga.service" || -e "$WANTS_DIR/soga.service" ]]; then
        echo "default"
    fi
    while read -r u _; do
        case "$u" in
            soga@.service) ;;
            soga.service) echo "default" ;;
            soga@*.service)
                u="${u#soga@}"
                echo "${u%.service}"
                ;;
        esac
    done < <(systemctl list-unit-files --type=service --no-legend 2>/dev/null)
}

# scan_instances：/etc/soga 与 systemd 中存在的全部实例（排序去重）
scan_instances() {
    local d
    {
        [[ -f "$CONF_ROOT/soga.conf" ]] && echo "default"
        for d in "$CONF_ROOT"/*/; do
            [[ -d "$d" ]] || continue
            d="${d%/}"
            echo "${d##*/}"
        done
        scan_units
    } | while read -r d; do [[ -n "$d" ]] && echo "$d"; done | LC_ALL=C sort -u
}

# ---------- 系统检查与依赖 ----------

require_root() {
    [[ $EUID -eq 0 ]] || fail "错误: 必须使用 root 用户运行" || return 1
    command -v systemctl >/dev/null 2>&1 || fail "错误: 未检测到 systemctl，请确认运行在 systemd 系发行版"
}

check_os() {
    local s f
    if [[ -r /etc/os-release ]]; then
        s="$(tr '[:upper:]' '[:lower:]' </etc/os-release)"
        if [[ "$s" == *"id="* ]] && [[ "$s" =~ ubuntu|debian|centos|rhel|fedora|rocky|almalinux ]]; then
            return 0
        fi
        fail "未检测到受支持的系统发行版（请使用 Ubuntu/Debian/CentOS 系）"
        return 1
    fi
    for f in /etc/issue /proc/version; do
        [[ -r "$f" ]] || continue
        s="$(tr '[:upper:]' '[:lower:]' <"$f")"
        [[ "$s" =~ ubuntu|debian|centos|"red hat"|rhel ]] && return 0
    done
    [[ -e /etc/redhat-release ]] && return 0
    fail "未检测到受支持的系统发行版"
}

distro_family() {
    local s=""
    [[ -r /etc/os-release ]] && s="$(tr '[:upper:]' '[:lower:]' </etc/os-release)"
    if [[ "$s" =~ centos|rhel|rocky|almalinux|fedora ]]; then
        echo centos
    elif [[ "$s" =~ ubuntu|debian ]]; then
        echo debian
    elif [[ -e /etc/redhat-release ]]; then
        echo centos
    elif command -v apt >/dev/null 2>&1; then
        echo debian
    elif command -v yum >/dev/null 2>&1; then
        echo centos
    fi
}

ensure_base() {
    local -a missing=()
    local c
    for c in wget curl tar socat crontab; do
        command -v "$c" >/dev/null 2>&1 || missing+=("$c")
    done
    ((${#missing[@]} == 0)) && return 0
    echo "  安装基础依赖: ${missing[*]}" >&2
    case "$(distro_family)" in
        centos)
            run_cmd yum install -y epel-release || fail "安装 epel-release 失败: $ERR" || return 1
            run_cmd yum install -y wget curl tar crontabs socat tzdata || fail "yum 安装基础依赖失败: $ERR"
            ;;
        debian)
            run_cmd apt install -y wget curl tar cron socat tzdata || fail "apt 安装基础依赖失败: $ERR"
            ;;
        *)
            fail "无法识别的发行版，无法自动安装基础依赖，请手动安装: ${missing[*]}"
            ;;
    esac
}

ensure_acme() {
    [[ -e "$ACME_SH" ]] && return 0
    command -v curl >/dev/null 2>&1 || fail "curl 不存在，无法安装 acme.sh" || return 1
    echo "  安装 acme.sh ..." >&2
    run_cmd sh -c "curl https://get.acme.sh | sh" || fail "安装 acme.sh 失败: $ERR" || return 1
    [[ -e "$ACME_SH" ]] || fail "acme.sh 安装后未找到 $ACME_SH" || return 1
    run_cmd "$ACME_SH" --set-default-ca --server letsencrypt || fail "acme.sh 设置默认 CA 失败: $ERR"
}

# ---------- GitHub release ----------

# github_get <url>：成功时 HTTP_CODE 为状态码，HTTP_BODY 为响应内容
github_get() {
    local out rc
    out="$(curl -sSL -H 'Accept: application/vnd.github+json' -H 'User-Agent: soga-manager' \
        -w $'\n%{http_code}' "$1" 2>&1)"
    rc=$?
    if ((rc != 0)); then
        fail "${out##*$'\n'}"
        return 1
    fi
    HTTP_CODE="${out##*$'\n'}"
    HTTP_BODY="${out%$'\n'*}"
}

# release_pick_from_list <flat>：开发版取预发布与正式版中较新的一个，结果 PICK_IDX
release_pick_from_list() {
    local flat="$1" i tag pre draft p_idx="" s_idx="" p_tag="" s_tag=""
    for ((i = 0; i < 100; i++)); do
        grep -q "^/$i/" <<<"$flat" || break
        tag="$(sed -n "s#^/$i/tag_name\t##p" <<<"$flat" | head -n 1)"
        pre="$(sed -n "s#^/$i/prerelease\t##p" <<<"$flat" | head -n 1)"
        draft="$(sed -n "s#^/$i/draft\t##p" <<<"$flat" | head -n 1)"
        trim_to tag "$tag"
        [[ "$draft" == "true" || -z "$tag" ]] && continue
        if [[ "$pre" == "true" ]]; then
            [[ -z "$p_idx" ]] && p_idx="$i" && p_tag="$tag"
        else
            [[ -z "$s_idx" ]] && s_idx="$i" && s_tag="$tag"
        fi
        [[ -n "$p_idx" && -n "$s_idx" ]] && break
    done
    if [[ -n "$p_idx" && -n "$s_idx" ]]; then
        if version_greater "$p_tag" "$s_tag"; then PICK_IDX="$p_idx"; else PICK_IDX="$s_idx"; fi
    elif [[ -n "$p_idx" ]]; then
        PICK_IDX="$p_idx"
    elif [[ -n "$s_idx" ]]; then
        PICK_IDX="$s_idx"
    else
        return 1
    fi
}

# resolve_release <channel> <version>：结果 REL_TAG 与 REL_FLAT（单个 release 的展开结果）
resolve_release() {
    local channel="$1" version="$2" v flat
    if [[ -n "${version// /}" ]]; then
        v="$(strip_v "$version")"
        github_get "$API_BASE/tags/$v" || fail "查询版本 $v 失败: $ERR" || return 1
        [[ "$HTTP_CODE" == "404" ]] && { fail "版本 $v 不存在"; return 1; }
        [[ "$HTTP_CODE" == "200" ]] || fail "GitHub API 返回 $HTTP_CODE" || return 1
        REL_FLAT="$(json_flatten <<<"$HTTP_BODY")"
    elif [[ "$channel" == "dev" ]]; then
        github_get "$API_BASE?per_page=20" || fail "请求 GitHub API 失败: $ERR" || return 1
        [[ "$HTTP_CODE" == "200" ]] || fail "GitHub API 返回 $HTTP_CODE" || return 1
        flat="$(json_flatten <<<"$HTTP_BODY")" || fail "解析 release 列表失败" || return 1
        if ! release_pick_from_list "$flat"; then
            resolve_release stable ""
            return
        fi
        REL_FLAT="$(sed -n "s#^/$PICK_IDX/#/#p" <<<"$flat")"
    else
        github_get "$API_BASE/latest" || fail "请求 GitHub API 失败: $ERR" || return 1
        [[ "$HTTP_CODE" == "200" ]] || fail "GitHub API 返回 $HTTP_CODE" || return 1
        REL_FLAT="$(json_flatten <<<"$HTTP_BODY")"
    fi
    REL_TAG="$(sed -n 's#^/tag_name\t##p' <<<"$REL_FLAT" | head -n 1)"
}

# release_asset_url <asset_name>
release_asset_url() {
    local i
    i="$(awk -F'\t' -v n="$1" '$1 ~ /^\/assets\/[0-9]+\/name$/ && $2 == n { split($1, a, "/"); print a[3]; exit }' <<<"$REL_FLAT")"
    [[ -n "$i" ]] || return 1
    sed -n "s#^/assets/$i/browser_download_url\t##p" <<<"$REL_FLAT" | head -n 1
}

# latest_soga_tag <channel>
latest_soga_tag() {
    local flat
    if [[ "$1" == "dev" ]]; then
        github_get "$API_BASE?per_page=20" || return 1
        [[ "$HTTP_CODE" == "200" ]] || fail "GitHub API 返回 $HTTP_CODE" || return 1
        flat="$(json_flatten <<<"$HTTP_BODY")"
        release_pick_from_list "$flat" || fail "未找到可用 release" || return 1
        sed -n "s#^/$PICK_IDX/tag_name\t##p" <<<"$flat" | head -n 1
    else
        github_get "$API_BASE/latest" || return 1
        [[ "$HTTP_CODE" == "200" ]] || fail "GitHub API 返回 $HTTP_CODE" || return 1
        json_flatten <<<"$HTTP_BODY" | sed -n 's#^/tag_name\t##p' | head -n 1
    fi
}

# latest_cmd_version：仓库中最新管理脚本的版本号
latest_cmd_version() {
    local v
    v="$(curl -fsSL "$RAW_BASE/release/soga.sh" 2>/dev/null | sed -n 's/^SCRIPT_VERSION="\([^"]*\)".*/\1/p' | head -n 1)"
    [[ -n "$v" ]] || return 1
    echo "$v"
}

# download_file <url> <dest>：下载并显示进度
download_file() {
    local url="$1" dest="$2" total code pid size rc codefile
    total="$(curl -sIL "$url" 2>/dev/null | tr -d '\r' | awk 'tolower($1) == "content-length:" { n = $2 } END { print n + 0 }')"
    codefile="$(mktemp)" || return 1
    curl -sSL -o "$dest" -w '%{http_code}' "$url" >"$codefile" 2>"$codefile.err" &
    pid=$!
    while kill -0 "$pid" 2>/dev/null; do
        if ((total > 0)); then
            size="$(stat -c %s "$dest" 2>/dev/null || echo 0)"
            printf '\r  下载中 %d%% (%d/%d KB)' $((100 * size / total)) $((size / 1024)) $((total >> 10))
        fi
        sleep 0.2
    done
    wait "$pid"
    rc=$?
    code="$(cat "$codefile")"
    if ((rc == 0 && total > 0)); then
        size="$(stat -c %s "$dest" 2>/dev/null || echo 0)"
        printf '\r  下载中 %d%% (%d/%d KB)' $((100 * size / total)) $((size / 1024)) $((total >> 10))
    fi
    printf '\n'
    if ((rc != 0)); then
        fail "下载失败: $(tail -n 1 "$codefile.err")"
        rm -f "$codefile" "$codefile.err"
        return 1
    fi
    rm -f "$codefile" "$codefile.err"
    [[ "$code" == "200" ]] || fail "下载返回 $code"
}

# extract_binary <tgz> <dest_dir>：从安装包中取出名为 soga 的程序
extract_binary() {
    local tmpd f rc
    tmpd="$(mktemp -d)" || return 1
    if ! tar --no-same-owner -xzf "$1" -C "$tmpd" 2>/dev/null; then
        rm -rf "$tmpd"
        fail "解压 gzip 失败: 安装包已损坏"
        return 1
    fi
    f="$(find "$tmpd" -type f -name soga | head -n 1)"
    if [[ -z "$f" ]]; then
        rm -rf "$tmpd"
        fail "tar 中未找到 soga 二进制"
        return 1
    fi
    mkdir -p "$2" && install -m 0755 "$f" "$2/soga"
    rc=$?
    rm -rf "$tmpd"
    ((rc == 0)) || fail "写入 $2/soga 失败"
}

# install_binary <version>：下载并安装 soga 程序
install_binary() {
    local version="$1" channel="${CHANNEL:-stable}" label="稳定版" url tmp tag
    detect_arch || return 1
    ensure_base || return 1
    ensure_acme || return 1
    tmp="$(mktemp --suffix=.tar.gz)" || return 1
    if [[ -n "$SOGA_TARBALL" ]]; then
        echo "  使用本地安装包: $SOGA_TARBALL"
        if ! cp -f "$SOGA_TARBALL" "$tmp"; then
            rm -f "$tmp"
            fail "读取本地安装包失败: $SOGA_TARBALL"
            return 1
        fi
        tag=""
    else
        [[ "$channel" == "dev" ]] && label="开发版"
        echo "  查询 GitHub release 信息（分支: $label）..."
        if ! resolve_release "$channel" "$version"; then
            rm -f "$tmp"
            return 1
        fi
        tag="$REL_TAG"
        if ! url="$(release_asset_url "soga-linux-$ARCH.tar.gz")"; then
            rm -f "$tmp"
            fail "release $tag 中未找到 soga-linux-$ARCH.tar.gz"
            return 1
        fi
        println_color "$GREEN" "  安装 soga $(with_v "$tag") ($ARCH)"
        if ! download_file "$url" "$tmp"; then
            rm -f "$tmp"
            return 1
        fi
    fi
    [[ -e "$SOGA_BIN" ]] && mv -f "$SOGA_BIN" "$SOGA_BIN.bak"
    if ! extract_binary "$tmp" "$INSTALL_DIR"; then
        [[ -e "$SOGA_BIN.bak" ]] && mv -f "$SOGA_BIN.bak" "$SOGA_BIN"
        rm -f "$tmp"
        fail "解压失败: $ERR"
        return 1
    fi
    rm -f "$tmp" "$SOGA_BIN.bak"
    [[ -z "$tag" ]] && tag="$("$SOGA_BIN" -v 2>/dev/null)"
    println_color "$GREEN" "  soga $(with_v "$tag") 安装完成"
}

ensure_service_units() {
    local n
    UNIT_CHANGED=0
    ensure_default_unit || return 1
    ensure_template_unit || return 1
    while read -r n; do
        [[ -z "$n" || "$n" == "default" ]] && continue
        ensure_instance_unit "$n" || return 1
    done < <({
        printf '%s\n' "${INST_NAMES[@]}"
        scan_instances
    } | LC_ALL=C sort -u)
    if ((UNIT_CHANGED)); then
        svc_daemon_reload || return 1
    fi
    return 0
}

# install_soga <version>：安装/更新 soga 程序并刷新 systemd 服务
install_soga() {
    install_binary "$1" || return 1
    ensure_service_units
}

binary_version() {
    local out
    if [[ ! -e "$SOGA_BIN" ]]; then
        echo "(未安装)"
        return
    fi
    if out="$("$SOGA_BIN" -v 2>/dev/null)"; then
        trim_to out "$out"
        echo "$out"
    else
        echo "(未知)"
    fi
}

# self_update：SELF_UPDATED=1 表示已替换管理脚本
self_update() {
    local latest tmp
    SELF_UPDATED=0
    tmp="$(mktemp)" || return 1
    if ! curl -fsSL -o "$tmp" "$RAW_BASE/release/soga.sh"; then
        rm -f "$tmp"
        fail "查询管理脚本最新版本失败: $RAW_BASE/release/soga.sh"
        return 1
    fi
    latest="$(sed -n 's/^SCRIPT_VERSION="\([^"]*\)".*/\1/p' "$tmp" | head -n 1)"
    if [[ -z "$latest" ]] || ! bash -n "$tmp" 2>/dev/null; then
        rm -f "$tmp"
        fail "下载的管理脚本无效"
        return 1
    fi
    if ! version_greater "$latest" "$SCRIPT_VERSION"; then
        rm -f "$tmp"
        printf_color "$GREEN" "管理脚本已是最新版本 %s" "$(with_v "$SCRIPT_VERSION")"
        return 0
    fi
    printf_color "$YELLOW" "当前版本 %s，最新版本 %s" "$(with_v "$SCRIPT_VERSION")" "$(with_v "$latest")"
    printf_color "$GREEN" "  正在更新到 %s ..." "$(with_v "$latest")"
    if ! install -m 0755 "$tmp" "$CMD_PATH.tmp" || ! mv -f "$CMD_PATH.tmp" "$CMD_PATH"; then
        rm -f "$tmp" "$CMD_PATH.tmp"
        fail "写入 $CMD_PATH 失败"
        return 1
    fi
    rm -f "$tmp"
    printf_color "$GREEN" "管理脚本已更新到 %s，请重新运行" "$(with_v "$latest")"
    SELF_UPDATED=1
}

# ---------- 更新检查（每 12 小时一次，后台进行） ----------

refresh_cached_async() {
    local now
    now="$(now_ms)"
    if ((UC_SET && UC_LAST > 0 && now - UC_LAST < 43200000)); then
        return 0
    fi
    refresh_cached_job </dev/null >/dev/null 2>&1 &
    disown 2>/dev/null || true
}

# 后台任务：查询最新版本并写回 config.json（重新读取磁盘上的内容，只更新检查结果）
refresh_cached_job() {
    local soga_tag cmd_tag s_ok c_ok
    soga_tag="$(latest_soga_tag "${CHANNEL:-stable}")"
    s_ok=$?
    cmd_tag="$(latest_cmd_version)"
    c_ok=$?
    ((s_ok != 0 && c_ok != 0)) && return 0
    store_load || return 0
    UC_SET=1
    ((s_ok == 0)) && UC_SOGA="$soga_tag"
    ((c_ok == 0)) && UC_CMD="$cmd_tag"
    UC_LAST="$(now_ms)"
    store_save
}

update_notice() {
    local cur="$1" items=""
    ((UC_SET)) || return 0
    if [[ -n "$UC_CMD" ]] && version_greater "$UC_CMD" "$SCRIPT_VERSION"; then
        items="管理脚本 $(with_v "$UC_CMD")"
    fi
    if [[ -n "$UC_SOGA" ]]; then
        case "$cur" in
            "" | "(未安装)" | "(未知)" | "(模拟)") items="${items:+$items, }soga $(with_v "$UC_SOGA")" ;;
            *) version_greater "$UC_SOGA" "$cur" && items="${items:+$items, }soga $(with_v "$UC_SOGA")" ;;
        esac
    fi
    [[ -n "$items" ]] && echo "发现新版本: $items"
    return 0
}

# ---------- 实例操作 ----------

validate_name() {
    local n
    trim_to n "$1"
    [[ -n "$n" ]] || fail "实例名不能为空" || return 1
    [[ "$n" =~ ^[A-Za-z0-9]+$ ]] || fail "实例名仅允许字母与数字: $n"
}

# list_items <是否包含 default>：结果 LI_NAMES / LI_STATUS / LI_ENABLED
list_items() {
    local n
    LI_NAMES=()
    LI_STATUS=()
    LI_ENABLED=()
    [[ "$1" == "1" ]] && LI_NAMES+=("default")
    for n in "${INST_NAMES[@]}"; do
        trim_to n "$n"
        [[ -z "$n" || "$n" == "default" ]] && continue
        LI_NAMES+=("$n")
    done
    for n in "${LI_NAMES[@]}"; do
        LI_STATUS+=("$(svc_status "$n")")
        if svc_enabled "$n"; then LI_ENABLED+=(1); else LI_ENABLED+=(0); fi
    done
}

bin_exists() {
    [[ -e "$SOGA_BIN" ]]
}

# instance_create <name> [配置内容文件]：写入服务、配置与规则文件
instance_create() {
    local n="$1" src="${2:-}" f
    if [[ "$n" == "default" ]]; then
        ensure_default_unit || return 1
    else
        ensure_template_unit || return 1
    fi
    f="$(conf_path "$n")"
    if [[ ! -e "$f" ]]; then
        mkdir -p "$(conf_dir "$n")" || fail "mkdir $(conf_dir "$n") 失败" || return 1
        if [[ -n "$src" && -s "$src" ]]; then
            cat "$src" >"$f" || fail "写入 $f 失败" || return 1
        else
            default_conf_content "$n" >"$f" || fail "写入 $f 失败" || return 1
        fi
        chmod 644 "$f" 2>/dev/null
    fi
    ensure_attachments "$n" || return 1
    svc_daemon_reload
}

instance_delete() {
    svc_run stop "$1"
    svc_run disable "$1"
    remove_instance_unit "$1"
    svc_daemon_reload
}

# instance_new <name> [k=v ...]：结果 NEW_CONF_PATH
instance_new() {
    local n tmp="" kv k
    trim_to n "$1"
    shift
    [[ "$n" == *=* ]] && { fail "实例名 \"$n\" 包含 '='，疑似漏写名称"; return 1; }
    [[ "$n" == "default" ]] && { fail "default 为内置实例，不可新增；安装 soga 时自动创建"; return 1; }
    validate_name "$n" || return 1
    find_instance "$n" && { fail "实例 $n 已存在"; return 1; }
    if [[ -f "$CONF_ROOT/soga.conf" ]]; then
        tmp="$(mktemp)" || return 1
        cat "$CONF_ROOT/soga.conf" >"$tmp"
        set_instance_dir "$tmp" "$(conf_dir "$n")"
        for kv in "$@"; do
            trim_to k "${kv%%=*}"
            conf_set "$tmp" "$k" "${kv#*=}"
        done
    fi
    if ! store_add_instance "$n"; then
        rm -f "$tmp"
        fail "实例 $n 已存在"
        return 1
    fi
    if ! store_save || ! instance_create "$n" "$tmp"; then
        rm -f "$tmp"
        return 1
    fi
    if [[ -z "$tmp" ]] && (($# > 0)); then
        apply_kv "$n" "$@" || return 1
    fi
    rm -f "$tmp"
    NEW_CONF_PATH="$(conf_path "$n")"
}

# apply_kv <name> k=v ...
apply_kv() {
    local n="$1" f kv k
    shift
    f="$(conf_path "$n")"
    [[ -f "$f" ]] || fail "实例 $n 的 soga.conf 不存在" || return 1
    [[ -r "$f" ]] || fail "读取实例 $n 配置失败: permission denied" || return 1
    for kv in "$@"; do
        trim_to k "${kv%%=*}"
        conf_set "$f" "$k" "${kv#*=}" || fail "写入实例 $n 配置失败: $f" || return 1
    done
}

# instance_remove <name...>：结果 REMOVED 为成功移除的实例
instance_remove() {
    local -a names=()
    local n x dup first=""
    REMOVED=()
    for n in "$@"; do
        trim_to n "$n"
        [[ -n "$n" ]] || fail "实例名为空" || return 1
        [[ "$n" != "default" ]] || fail "default 为内置实例，不可移除" || return 1
        dup=0
        for x in "${names[@]}"; do [[ "$x" == "$n" ]] && dup=1; done
        ((dup)) && continue
        require_instance "$n" || return 1
        names+=("$n")
    done
    for n in "${names[@]}"; do
        if ! instance_delete "$n"; then
            [[ -z "$first" ]] && first="删除实例 $n 失败: $ERR"
            continue
        fi
        store_remove_instance "$n"
        rm -rf "${CONF_ROOT:?}/$n"
        REMOVED+=("$n")
    done
    store_save || return 1
    [[ -z "$first" ]] || fail "$first"
}

instance_rename() {
    local old new was_status was_enabled=0
    trim_to old "$1"
    trim_to new "$2"
    [[ -n "$old" ]] || fail "原实例名为空" || return 1
    [[ "$old" != "default" && "$new" != "default" ]] || fail "default 为内置实例，不可参与改名" || return 1
    [[ "$new" != "$old" ]] || fail "新旧名称相同" || return 1
    validate_name "$new" || return 1
    require_instance "$old" || return 1
    find_instance "$new" && { fail "实例 $new 已存在"; return 1; }
    [[ -e "$(conf_path "$new")" ]] && { fail "目标配置目录已存在: $(conf_dir "$new")"; return 1; }

    was_status="$(svc_status "$old")"
    svc_enabled "$old" && was_enabled=1
    svc_run stop "$old"
    if ! mv "$(conf_dir "$old")" "$(conf_dir "$new")" 2>/dev/null; then
        [[ "$was_status" == "0" ]] && svc_run start "$old"
        fail "移动配置目录失败: rename $(conf_dir "$old") $(conf_dir "$new")"
        return 1
    fi
    if [[ -f "$(conf_path "$new")" ]]; then
        relocate_instance_dir "$(conf_path "$new")" "$(conf_dir "$new")" "$(conf_dir "$old")" "$CONF_ROOT" ||
            fail "更新实例配置路径失败: $(conf_path "$new")" || return 1
    fi
    svc_run disable "$old"
    remove_instance_unit "$old"
    store_rename_instance "$old" "$new" || fail "更新实例清单失败: $old → $new" || return 1
    store_save || return 1
    svc_daemon_reload || return 1
    if ((was_enabled)); then
        svc_run enable "$new" || fail "启用新实例失败: $ERR" || return 1
    fi
    if [[ "$was_status" == "0" ]]; then
        svc_reset_failed "$new"
        svc_run start "$new" || fail "启动新实例失败: $ERR" || return 1
    fi
    return 0
}

uninstall_all() {
    local n
    while read -r n; do
        [[ -n "$n" ]] && instance_delete "$n"
    done < <(scan_instances)
    rm -f "$UNIT_DIR/soga@.service"
    systemctl daemon-reload
    systemctl reset-failed >/dev/null 2>&1
    rm -rf "$CONF_ROOT" "$INSTALL_DIR"
    return 0
}

restart_running() {
    local -a names=()
    local i inc=0 joined
    bin_exists && inc=1
    list_items "$inc"
    for i in "${!LI_NAMES[@]}"; do
        [[ "${LI_STATUS[i]}" == "0" ]] && names+=("${LI_NAMES[i]}")
    done
    ((${#names[@]} == 0)) && return 0
    joined="$(printf '%s, ' "${names[@]}")"
    printf '  正在重启运行中的实例: %s\n' "${joined%, }"
    for i in "${names[@]}"; do
        svc_reset_failed "$i"
        if svc_run restart "$i"; then
            printf_color "$GREEN" "已重启 %s" "$i"
        else
            printf_color "$RED" "重启 %s 失败: %s" "$i" "$ERR"
        fi
    done
}

# ---------- 快速配置向导 ----------

F_LABELS=(
    "① 面板上报类型 type（面板类型）"
    "② 节点协议 server_type"
    "③ 节点 ID (node_id)，必填，多个节点 ID 用英文逗号分隔（例: 1 或 1,2,3）"
    "④ soga_key（授权码，选填；留空则为社区版，最多 88 名用户）"
    "⑤ 面板对接方式"
    "⑥ WebAPI 根地址 webapi_url（如 https://your-panel.com/）"
    "⑦ WebAPI 密钥 webapi_key"
    "⑥ 数据库主机 db_host"
    "⑦ 数据库端口 db_port"
    "⑧ 数据库名 db_name"
    "⑨ 数据库用户名 db_user"
    "⑩ 数据库密码 db_password"
)
F_REQUIRED=(1 1 1 0 1 1 1 1 1 1 1 1)
F_DEFAULTS=("" "" "" "" "" "" "" "" "3306" "" "root" "")
TYPE_OPTIONS=(
    "xboard|xboard（cedar2025/Xboard）"
    "xiaov2board|xiaov2board（wyx2685/v2board）"
    "sspanel-uim|sspanel-uim（SSPanel-UIM）"
    "soga-v1|soga-v1"
    "v2board|v2board（原版v2board）"
    "v2raysocks|v2raysocks"
    "whmcs|whmcs"
    "proxypanel|proxypanel"
    "ppanel|ppanel"
)
SERVER_TYPE_OPTIONS=(
    "ss|ss（Shadowsocks）"
    "v2ray|v2ray（VMess）"
    "trojan|trojan"
    "vless|vless"
    "ssr|ssr"
    "hysteria|hysteria"
    "anytls|anytls"
    "mieru|mieru"
)

validate_node_id() {
    local s p
    local -a parts
    trim_to s "$1"
    [[ -n "$s" ]] || fail "节点 ID 不能为空" || return 1
    IFS=',' read -r -a parts <<<"$s"
    for p in "${parts[@]}"; do
        trim_to p "$p"
        [[ -n "$p" ]] && return 0
    done
    fail "请至少填写一个有效的节点 ID"
}

# prompt_text <label> <default> <required> <validator>：结果 WZ_VALUE
prompt_text() {
    local label="$1" def="$2" req="$3" validator="$4" hint v
    while true; do
        if ((req)); then
            if [[ -n "$def" ]]; then hint="不可留空，回车默认 $def"; else hint="不可留空"; fi
        else
            if [[ -n "$def" ]]; then hint="回车默认 $def"; else hint="可留空"; fi
        fi
        printf '%s (%s): ' "$label" "$hint"
        read_line || return 1
        v="$LINE"
        [[ -z "$v" ]] && trim_to v "$def"
        if ((req)) && [[ -z "$v" ]]; then
            println_color "$YELLOW" "不能为空，请重新输入"
            continue
        fi
        if [[ -n "$validator" ]] && ! "$validator" "$v"; then
            println_color "$YELLOW" "$ERR"
            continue
        fi
        WZ_VALUE="$v"
        return 0
    done
}

# pick_option <label> <默认序号> <value|label>...：结果 WZ_VALUE
pick_option() {
    local label="$1" def="$2" i n mark
    shift 2
    local -a opts=("$@")
    n=${#opts[@]}
    echo
    println_color "$GREEN" "$label"
    for i in "${!opts[@]}"; do
        mark=""
        ((i == def)) && mark="${YELLOW} [默认]${RESET}"
        printf ' %s%2d.%s %s%s\n' "$GREEN" $((i + 1)) "$RESET" "${opts[i]#*|}" "$mark"
    done
    printf '请选择 [1-%d，回车=默认]: ' "$n"
    read_line || return 1
    if [[ -z "$LINE" ]]; then
        WZ_VALUE="${opts[def]%%|*}"
        return 0
    fi
    i=0
    if [[ "$LINE" =~ ^[+-]?[0-9]+ ]]; then
        i="${BASH_REMATCH[0]#+}"
        if [[ "$i" == -* ]]; then i=0; else i=$((10#$i)); fi
    fi
    if ((i <= 0 || i > n)); then
        println_color "$YELLOW" "无效选择，已使用默认项"
        i=$((def + 1))
    fi
    WZ_VALUE="${opts[i - 1]%%|*}"
}

# wizard_run：以 SC 中的现有值作为默认值，结果写回 SC（未出现的项为空）
wizard_run() {
    local -A old_vals=()
    local i k def j di validator
    local -a opts
    for k in "${SC_KEYS[@]}"; do old_vals[$k]="${SC[$k]:-}"; done
    SC=()
    for i in "${!SC_KEYS[@]}"; do
        k="${SC_KEYS[i]}"
        case "$k" in
            webapi_url | webapi_key) [[ "${SC[api]:-}" == "webapi" ]] || continue ;;
            db_*) [[ "${SC[api]:-}" == "db" ]] || continue ;;
        esac
        def="${F_DEFAULTS[i]}"
        [[ -n "${old_vals[$k]}" ]] && def="${old_vals[$k]}"
        opts=()
        case "$k" in
            type) opts=("${TYPE_OPTIONS[@]}") ;;
            server_type) opts=("${SERVER_TYPE_OPTIONS[@]}") ;;
            api)
                case "${SC[type]:-}" in
                    whmcs) opts=("db|db") ;;
                    sspanel-uim) opts=("webapi|webapi" "db|db") ;;
                    *) opts=("webapi|webapi") ;;
                esac
                ;;
        esac
        if ((${#opts[@]} == 1)); then
            WZ_VALUE="${opts[0]%%|*}"
        elif ((${#opts[@]} == 0)); then
            validator=""
            [[ "$k" == "node_id" ]] && validator=validate_node_id
            prompt_text "${F_LABELS[i]}" "$def" "${F_REQUIRED[i]}" "$validator" || return 1
        else
            di=0
            trim_to def "$def"
            for j in "${!opts[@]}"; do
                if [[ "${opts[j]%%|*}" == "$def" ]]; then
                    di=$j
                    break
                fi
            done
            pick_option "${F_LABELS[i]}" "$di" "${opts[@]}" || return 1
        fi
        SC[$k]="$WZ_VALUE"
    done
}

# ---------- 交互菜单 ----------

# 菜单项: "SEP" 为分隔线，其余为 "动作函数|显示文字"
MENU_ITEMS=()
MENU_PRE_PROMPT=""
MENU_QUIT=0

# prompt_menu <返回项文字> <是否允许输入实例名>：结果 PICK_ACTION / PICK_QUIT / PICK_INSTANCE
prompt_menu() {
    local exit_label="$1" resolver="$2" count=0 item n target i
    PICK_ACTION=""
    PICK_QUIT=0
    PICK_INSTANCE=""
    printf ' %s%2d.%s %s\n' "$GREEN" 0 "$RESET" "$exit_label"
    for item in "${MENU_ITEMS[@]}"; do
        if [[ "$item" == "SEP" ]]; then
            echo "$SEP"
        else
            count=$((count + 1))
            printf ' %s%2d.%s %s\n' "$GREEN" "$count" "$RESET" "${item#*|}"
        fi
    done
    echo
    if [[ -n "$MENU_PRE_PROMPT" ]]; then
        "$MENU_PRE_PROMPT"
        echo
    fi
    if ((resolver)); then
        printf '请输入选择 [0-%d] 或 [实例名]: ' "$count"
    else
        printf '请输入选择 [0-%d]: ' "$count"
    fi
    read_line || return 1
    if [[ "$LINE" == "0" ]]; then
        PICK_QUIT=1
        return 0
    fi
    target=0
    if [[ "$LINE" =~ ^[+-]?[0-9]+ ]]; then
        target="${BASH_REMATCH[0]#+}"
        if [[ "$target" == -* ]]; then target=0; else target=$((10#$target)); fi
    fi
    if ((target > 0 && target <= count)); then
        n=0
        for item in "${MENU_ITEMS[@]}"; do
            [[ "$item" == "SEP" ]] && continue
            n=$((n + 1))
            if ((n == target)); then
                PICK_ACTION="${item%%|*}"
                return 0
            fi
        done
    fi
    if ((resolver)) && [[ -n "$LINE" ]]; then
        for i in "${LI_NAMES[@]}"; do
            if [[ "$i" == "$LINE" ]]; then
                PICK_INSTANCE="$i"
                return 0
            fi
        done
    fi
    println_color "$RED" "无效选择"
    return 0
}

# run_action <函数>：动作可设置 ACT_QUIT（返回上级/退出）或 ACT_SUB（进入实例菜单）
run_action() {
    ACT_QUIT=0
    ACT_SUB=""
    if ! "$1"; then
        println_color "$RED" "$ERR"
        return 0
    fi
    if [[ -n "$ACT_SUB" ]]; then
        instance_menu "$ACT_SUB" || return 1
        return 0
    fi
    if ((ACT_QUIT)); then
        MENU_QUIT=1
        return 0
    fi
    wait_enter
    return 0
}

# ----- 实例菜单 -----

CUR_INSTANCE=""

im_require_binary() {
    bin_exists || fail "soga 程序未安装"
}

im_start() {
    im_require_binary || return 1
    if [[ "$(svc_status "$CUR_INSTANCE")" == "0" ]]; then
        println_color "$YELLOW" "已在运行"
        return 0
    fi
    svc_reset_failed "$CUR_INSTANCE"
    svc_run start "$CUR_INSTANCE" || return 1
    println_color "$GREEN" "已启动"
}

im_stop() {
    svc_run stop "$CUR_INSTANCE" || return 1
    println_color "$GREEN" "已停止"
}

im_restart() {
    im_require_binary || return 1
    svc_reset_failed "$CUR_INSTANCE"
    svc_run restart "$CUR_INSTANCE" || return 1
    println_color "$GREEN" "已重启"
}

im_reload() {
    im_require_binary || return 1
    svc_reset_failed "$CUR_INSTANCE"
    svc_run reload-or-restart "$CUR_INSTANCE" || return 1
    println_color "$GREEN" "已重载"
}

im_status() {
    svc_show_status "$CUR_INSTANCE"
}

im_log() {
    svc_show_log "$CUR_INSTANCE" 1000 0
}

im_log_follow() {
    svc_show_log "$CUR_INSTANCE" 1000 1
}

im_enable() {
    svc_run enable "$CUR_INSTANCE" || return 1
    println_color "$GREEN" "已设置开机自启"
}

im_disable() {
    svc_run disable "$CUR_INSTANCE" || return 1
    println_color "$GREEN" "已取消开机自启"
}

im_quick_config() {
    local f tmp k
    printf_color "$YELLOW" "快速配置将只更新 %s 中的基础对接配置项" "$(menu_conf_path "$CUR_INSTANCE")"
    confirm "是否继续？" n || return 0
    load_parsed "$CUR_INSTANCE" || load_parsed default
    wizard_run || return 1
    f="$(conf_path "$CUR_INSTANCE")"
    if [[ -f "$f" ]]; then
        tmp="$(mktemp)" || return 1
        cat "$f" >"$tmp"
        for k in "${SC_KEYS[@]}"; do
            conf_set "$tmp" "$k" "${SC[$k]:-}"
        done
        if ! cat "$tmp" >"$f"; then
            rm -f "$tmp"
            fail "写入 $f 失败"
            return 1
        fi
        rm -f "$tmp"
    else
        if ! mkdir -p "$(conf_dir "$CUR_INSTANCE")" || ! format_conf "$(conf_dir "$CUR_INSTANCE")" >"$f"; then
            fail "写入 $f 失败"
            return 1
        fi
    fi
    println_color "$GREEN" "快速配置已保存"
}

default_editor() {
    local e
    trim_to e "${EDITOR:-}"
    if [[ -n "$e" ]]; then
        EDITOR_CMD="$e"
        return 0
    fi
    for e in vim nano vi; do
        if command -v "$e" >/dev/null 2>&1; then
            EDITOR_CMD="$e"
            return 0
        fi
    done
    fail "未找到可用编辑器（请设置 EDITOR，或安装 vim / nano / vi）"
}

im_edit() {
    default_editor || return 1
    run_cmd "$EDITOR_CMD" "$(conf_dir "$CUR_INSTANCE")/$1"
}

im_edit_conf() { im_edit soga.conf; }
im_edit_blocklist() { im_edit blockList; }
im_edit_whitelist() { im_edit whiteList; }
im_edit_routes() { im_edit routes.toml; }
im_edit_dns() { im_edit dns.yml; }

im_delete() {
    local n rc
    confirm "确定删除实例 $CUR_INSTANCE ?" n || return 0
    instance_remove "$CUR_INSTANCE"
    rc=$?
    for n in "${REMOVED[@]}"; do
        printf_color "$GREEN" "已删除 %s" "$n"
    done
    ((rc == 0)) || return 1
    ACT_QUIT=1
}

im_rename() {
    local new
    printf '请输入新名称 (原名: %s): ' "$CUR_INSTANCE"
    if ! read_line; then
        println_color "$RED" "$ERR"
        return 0
    fi
    if [[ -z "$LINE" || "$LINE" == "$CUR_INSTANCE" ]]; then
        println_color "$YELLOW" "已取消"
        return 0
    fi
    new="$LINE"
    confirm "确定将实例 $CUR_INSTANCE 重命名为 $new ?" n || return 0
    if ! instance_rename "$CUR_INSTANCE" "$new"; then
        println_color "$RED" "$ERR"
        return 0
    fi
    printf_color "$GREEN" "已重命名: %s → %s" "$CUR_INSTANCE" "$new"
    ACT_QUIT=1
}

im_print_status() {
    local st en=0
    st="$(svc_status "$CUR_INSTANCE")"
    svc_enabled "$CUR_INSTANCE" && en=1
    printf '  状态: %s    自启: ' "${ST_TEXT[st]}"
    if ((en)); then echo "$YES_TEXT"; else echo "$NO_TEXT"; fi
    printf '  配置: %s%s%s\n' "$YELLOW" "$(menu_conf_path "$CUR_INSTANCE")" "$RESET"
}

instance_menu() {
    local -a saved_items=("${MENU_ITEMS[@]}")
    local saved_quit="$MENU_QUIT" saved_cur="$CUR_INSTANCE" saved_pre="$MENU_PRE_PROMPT" rc=0
    CUR_INSTANCE="$1"
    MENU_QUIT=0
    while true; do
        echo
        MENU_ITEMS=(
            SEP
            "im_start|启动"
            "im_stop|停止"
            "im_restart|重启"
            "im_reload|重载配置"
            SEP
            "im_status|查看服务状态详情"
            "im_log|查看最近日志 (1000行)"
            "im_log_follow|持续输出日志"
            SEP
            "im_enable|设置开机自启"
            "im_disable|取消开机自启"
            SEP
            "im_quick_config|快速配置 ${YELLOW}soga.conf${RESET}"
            "im_edit_conf|编辑 ${YELLOW}soga.conf${RESET}"
            "im_edit_blocklist|编辑 ${YELLOW}blockList${RESET}"
            "im_edit_whitelist|编辑 ${YELLOW}whiteList${RESET}"
            "im_edit_routes|编辑 ${YELLOW}routes.toml${RESET}"
            "im_edit_dns|编辑 ${YELLOW}dns.yml${RESET}"
        )
        if [[ "$CUR_INSTANCE" != "default" ]]; then
            MENU_ITEMS+=(SEP "im_delete|删除此实例" "im_rename|修改名称")
        fi
        println_color "$GREEN" "  soga [$CUR_INSTANCE] 管理"
        echo
        im_print_status
        echo
        MENU_PRE_PROMPT=""
        if ! prompt_menu "返回上级" 0; then
            rc=1
            break
        fi
        ((PICK_QUIT)) && break
        if [[ -n "$PICK_ACTION" ]]; then
            if ! run_action "$PICK_ACTION"; then
                rc=1
                break
            fi
            ((MENU_QUIT)) && break
        fi
    done
    MENU_ITEMS=("${saved_items[@]}")
    MENU_QUIT="$saved_quit"
    MENU_PRE_PROMPT="$saved_pre"
    CUR_INSTANCE="$saved_cur"
    return $rc
}

# ----- 主菜单 -----

# 菜单「安装/更新 soga」与切换分支后的更新：与 soga update 一样，装完后重启运行中的实例
mm_install_with() {
    local root_missing=0
    [[ -e "$CONF_ROOT" ]] || root_missing=1
    install_soga "$1" || return 1
    if ((root_missing)); then
        instance_create default || return 1
    fi
    store_save || return 1
    restart_running
    return 0
}

mm_install() {
    printf '输入版本号留空则安装最新: '
    read_line
    mm_install_with "$LINE"
}

mm_uninstall() {
    confirm "确定完全卸载 soga？" n || return 0
    uninstall_all || return 1
    INST_NAMES=()
    INST_CONFS=()
    INST_DESCS=()
    store_save || return 1
    println_color "$GREEN" "卸载完成"
    ACT_QUIT=1
}

mm_switch_channel() {
    local label
    if [[ "$CHANNEL" == "dev" ]]; then
        CHANNEL="stable"
        label="稳定版"
    else
        CHANNEL="dev"
        label="开发版"
    fi
    store_save || return 1
    printf_color "$GREEN" "已切换到%s" "$label"
    confirm "是否立即按新分支更新 soga？" y || return 0
    mm_install_with ""
}

mm_self_update() {
    self_update || return 1
    ((SELF_UPDATED)) && ACT_QUIT=1
    return 0
}

mm_add_instance() {
    local n i tmp
    if ! bin_exists; then
        println_color "$YELLOW" "soga 程序未安装，请先安装"
        confirm "是否安装 soga？" y || fail "需要先安装 soga 才能新增实例" || return 1
        install_soga "" || return 1
    fi
    printf '请输入新实例名称 (如 node1): '
    read_line || return 1
    n="$LINE"
    [[ "$n" == "default" ]] && { fail "default 为内置实例，不可新增；安装 soga 后自动创建"; return 1; }
    validate_name "$n" || return 1
    for i in "${INST_NAMES[@]}"; do
        [[ "$i" == "$n" ]] && { fail "实例已存在"; return 1; }
    done
    load_parsed default
    wizard_run || return 1
    tmp="$(mktemp)" || return 1
    format_conf "$(conf_dir "$n")" >"$tmp"
    if ! store_add_instance "$n"; then
        rm -f "$tmp"
        fail "实例已存在"
        return 1
    fi
    if ! store_save || ! instance_create "$n" "$tmp"; then
        rm -f "$tmp"
        return 1
    fi
    rm -f "$tmp"
    println_color "$GREEN" "已添加实例 $n"
}

mm_open_instance() {
    ACT_SUB="$MM_TARGET"
}

mm_print_instances() {
    local i st en
    echo "  实例列表:"
    if ((${#LI_NAMES[@]} == 0)); then
        echo "    (无) 请先新增实例或从系统同步"
        return
    fi
    for i in "${!LI_NAMES[@]}"; do
        st="${LI_STATUS[i]}"
        [[ "$st" == "0" || "$st" == "1" ]] || st=2
        if ((LI_ENABLED[i])); then en="$YES_TEXT"; else en="$NO_TEXT"; fi
        printf '    %-12s [%s] [自启: %s]  %s%s%s\n' "${LI_NAMES[i]}" "${ST_TEXT[st]}" "$en" \
            "$YELLOW" "$(menu_conf_path "${LI_NAMES[i]}")" "$RESET"
    done
}

main_menu() {
    local inc i cur notice channel_label resolver
    store_load || return 1
    refresh_cached_async
    MENU_QUIT=0
    while true; do
        echo
        inc=0
        bin_exists && inc=1
        list_items "$inc"
        channel_label="稳定版"
        [[ "$CHANNEL" == "dev" ]] && channel_label="开发版"
        MENU_ITEMS=(
            SEP
            "mm_install|安装/更新 soga"
            "mm_uninstall|完全卸载 soga"
            "mm_switch_channel|切换 soga 安装分支  [当前: ${YELLOW}${channel_label}${RESET}]"
            SEP
            "mm_self_update|更新管理脚本"
            "mm_add_instance|新增 soga 实例"
            SEP
        )
        for i in "${!LI_NAMES[@]}"; do
            MENU_ITEMS+=("mm_open_instance:${LI_NAMES[i]}|管理实例 [${LI_NAMES[i]}]")
        done
        cur="$(binary_version)"
        println_color "$GREEN" "  soga 管理脚本 $(with_v "$SCRIPT_VERSION")  [soga: $(with_v "$cur")]"
        notice="$(update_notice "$cur")"
        [[ -n "$notice" ]] && println_color "$YELLOW" "  $notice"
        echo
        MENU_PRE_PROMPT=mm_print_instances
        resolver=0
        ((${#LI_NAMES[@]} > 0)) && resolver=1
        prompt_menu "退出" "$resolver" || return 1
        MENU_PRE_PROMPT=""
        ((PICK_QUIT)) && return 0
        if [[ -n "$PICK_INSTANCE" ]]; then
            instance_menu "$PICK_INSTANCE" || return 1
            continue
        fi
        if [[ "$PICK_ACTION" == mm_open_instance:* ]]; then
            MM_TARGET="${PICK_ACTION#mm_open_instance:}"
            PICK_ACTION=mm_open_instance
        fi
        if [[ -n "$PICK_ACTION" ]]; then
            run_action "$PICK_ACTION" || return 1
            ((MENU_QUIT)) && return 0
        fi
    done
}

# ---------- 命令行 ----------

declare -A CMD_SHORT=(
    [config]="查看/修改实例配置文件：无 kv 参数时输出配置内容，带 kv 参数时修改或新增对应配置"
    [disable]="取消实例开机自启"
    [enable]="设置实例开机自启"
    [list]="列出已有的实例及状态"
    [log]="查看日志（默认 1000 行；加 -f 持续输出）"
    [new]="新建实例：以 default 实例配置为模板，可自定义实例配置"
    [remove]="移除实例（支持同时移除多个；default 不可移除）"
    [rename]="重命名实例（迁移配置目录与 systemd 状态）"
    [restart]="重启实例"
    [self-update]="更新管理脚本"
    [start]="启动实例"
    [stop]="停止实例"
    [uninstall]="完全卸载 soga 与实例"
    [update]="安装/更新 soga"
    [version]="输出版本"
)
declare -A CMD_USE=(
    [config]="config [name] [k1=v1 k2=v2 ...]"
    [disable]="disable [name]"
    [enable]="enable [name]"
    [list]="list"
    [log]="log [name]"
    [new]="new <name> [k1=v1 k2=v2 ...]"
    [remove]="remove <name> [name...]"
    [rename]="rename <old_name> <new_name>"
    [restart]="restart [name]"
    [self-update]="self-update"
    [start]="start [name]"
    [stop]="stop [name]"
    [uninstall]="uninstall"
    [update]="update [version]"
    [version]="version"
)
CMD_ORDER=(config disable enable help list log new remove rename restart self-update start stop uninstall update version)

root_help() {
    local c
    echo "soga 多实例管理脚本"
    echo
    echo "Usage:"
    echo "  soga [flags]"
    echo "  soga [command]"
    echo
    echo "Available Commands:"
    for c in "${CMD_ORDER[@]}"; do
        printf '  %-11s %s\n' "$c" "${CMD_SHORT[$c]:-Help about any command}"
    done
    echo
    echo "Flags:"
    echo "  -h, --help      help for soga"
    echo "  -v, --version   输出版本"
    echo
    echo 'Use "soga [command] --help" for more information about a command.'
}

cmd_help() {
    local c="$1"
    echo "${CMD_SHORT[$c]}"
    echo
    echo "Usage:"
    echo "  soga ${CMD_USE[$c]} [flags]"
    echo
    echo "Flags:"
    if [[ "$c" == "log" ]]; then
        echo "  -f, --follow      持续日志输出"
        echo "  -h, --help        help for log"
        echo "  -n, --lines int   尾部行数 (default 1000)"
    else
        echo "  -h, --help   help for $c"
    fi
}

print_versions() {
    echo "管理脚本: $(with_v "$SCRIPT_VERSION")"
    echo "soga 程序: $(with_v "$(binary_version)")"
}

require_binary() {
    bin_exists || fail "soga 程序未安装，请先执行: soga update"
}

# parse_kv k=v...：校验 key=value 形式
parse_kv() {
    local kv k
    for kv in "$@"; do
        trim_to k "${kv%%=*}"
        if [[ "$kv" != *=* || -z "$k" ]]; then
            fail "无效参数 \"$kv\"，需为 key=value 形式"
            return 1
        fi
    done
}

name_arg() {
    if [[ -n "${1:-}" ]]; then echo "$1"; else echo "default"; fi
}

c_list() {
    local i inc=0 st en
    store_load || return 1
    bin_exists && inc=1
    list_items "$inc"
    if ((${#LI_NAMES[@]} == 0)); then
        echo "(无实例)"
        return 0
    fi
    for i in "${!LI_NAMES[@]}"; do
        st="${LI_STATUS[i]}"
        [[ "$st" == "0" || "$st" == "1" ]] || st=2
        if ((LI_ENABLED[i])); then en="$YES_TEXT"; else en="$NO_TEXT"; fi
        printf '%s\t%s\t自启:%s\n' "${LI_NAMES[i]}" "${ST_TEXT[st]}" "$en"
    done
}

c_new() {
    require_binary || return 1
    parse_kv "${@:2}" || return 1
    store_load || return 1
    instance_new "$@" || return 1
    printf_color "$GREEN" "已添加实例 %s（配置: %s）" "$1" "$NEW_CONF_PATH"
}

c_remove() {
    local n rc
    store_load || return 1
    instance_remove "$@"
    rc=$?
    for n in "${REMOVED[@]}"; do
        printf_color "$GREEN" "已移除 %s" "$n"
    done
    return $rc
}

c_rename() {
    store_load || return 1
    instance_rename "$1" "$2" || return 1
    printf_color "$GREEN" "已重命名: %s → %s" "$1" "$2"
}

c_start() {
    local n
    n="$(name_arg "${1:-}")"
    store_load || return 1
    require_binary || return 1
    require_instance "$n" || return 1
    svc_reset_failed "$n"
    svc_run start "$n"
}

c_stop() {
    local n
    n="$(name_arg "${1:-}")"
    store_load || return 1
    require_instance "$n" || return 1
    svc_run stop "$n"
}

c_restart() {
    local n
    n="$(name_arg "${1:-}")"
    store_load || return 1
    require_binary || return 1
    require_instance "$n" || return 1
    svc_reset_failed "$n"
    svc_run restart "$n"
}

c_enable() {
    local n
    n="$(name_arg "${1:-}")"
    store_load || return 1
    require_instance "$n" || return 1
    svc_run enable "$n"
}

c_disable() {
    local n
    n="$(name_arg "${1:-}")"
    store_load || return 1
    require_instance "$n" || return 1
    svc_run disable "$n"
}

c_log() {
    local n lines="$LOG_LINES"
    n="$(name_arg "${1:-}")"
    store_load || return 1
    require_instance "$n" || return 1
    if ((lines <= 0)); then
        if ((LOG_FOLLOW)); then lines=1000; else lines=100; fi
    fi
    svc_show_log "$n" "$lines" "$LOG_FOLLOW"
}

c_config() {
    local n="default" path kv k
    if (($# > 0)) && [[ "$1" != *=* ]]; then
        n="$1"
        shift
    fi
    parse_kv "$@" || return 1
    store_load || return 1
    require_instance "$n" || return 1
    if (($# == 0)); then
        path="$(conf_path "$n")"
        [[ -r "$path" ]] || fail "读取配置 $path 失败: open $path: no such file or directory" || return 1
        cat "$path"
        [[ -s "$path" && -n "$(tail -c 1 "$path")" ]] && echo
        return 0
    fi
    apply_kv "$n" "$@" || return 1
    for kv in "$@"; do
        trim_to k "${kv%%=*}"
        printf '已更新 %s=%s\n' "$k" "${kv#*=}"
    done
}

c_update() {
    local root_missing=0
    store_load || return 1
    [[ -e "$CONF_ROOT" ]] || root_missing=1
    install_soga "${1:-}" || return 1
    if ((root_missing)); then
        instance_create default || return 1
    fi
    store_save || return 1
    restart_running
    return 0
}

c_uninstall() {
    local ans=""
    printf '%s [y/N] ' "确定完全卸载 soga？"
    read -r ans _ || true
    ans="${ans,,}"
    [[ "$ans" == "y" || "$ans" == "yes" ]] || return 0
    store_load || return 1
    uninstall_all || return 1
    INST_NAMES=()
    INST_CONFS=()
    INST_DESCS=()
    store_save
}

c_self_update() {
    self_update
}

c_version() {
    print_versions
}

levenshtein() {
    local a="${1,,}" b="${2,,}" i j cost la lb v
    local -a prev cur
    la=${#a}
    lb=${#b}
    for ((j = 0; j <= lb; j++)); do prev[j]=$j; done
    for ((i = 1; i <= la; i++)); do
        cur=("$i")
        for ((j = 1; j <= lb; j++)); do
            cost=1
            [[ "${a:i-1:1}" == "${b:j-1:1}" ]] && cost=0
            v=$((prev[j] + 1))
            ((cur[j - 1] + 1 < v)) && v=$((cur[j - 1] + 1))
            ((prev[j - 1] + cost < v)) && v=$((prev[j - 1] + cost))
            cur[j]=$v
        done
        prev=("${cur[@]}")
    done
    echo "${prev[lb]}"
}

# 未知命令的建议（编辑距离 ≤ 2 或前缀匹配）
suggest_cmd() {
    local typed="$1" c
    for c in "${CMD_ORDER[@]}"; do
        [[ "$c" == "help" ]] && continue
        if [[ "$c" == "${typed,,}"* ]] || (($(levenshtein "$typed" "$c") <= 2)); then
            echo "$c"
        fi
    done
}

die_red() {
    println_color "$RED" "$1"
    exit 1
}

main() {
    local cmd="${1:-}" arg v n only_args=0 msg s
    local -a args=()

    if [[ $# -eq 0 ]]; then
        require_root || die_red "$ERR"
        check_os || die_red "$ERR"
        main_menu || die_red "$ERR"
        exit 0
    fi

    case "$cmd" in
        -h | --help)
            root_help
            exit 0
            ;;
        -v | --version)
            print_versions
            exit 0
            ;;
        help)
            if [[ -n "${2:-}" && -n "${CMD_SHORT[$2]:-}" ]]; then cmd_help "$2"; else root_help; fi
            exit 0
            ;;
        --*) die_red "unknown flag: ${cmd%%=*}" ;;
        -*) die_red "unknown shorthand flag: '${cmd:1:1}' in $cmd" ;;
    esac

    if [[ -z "${CMD_SHORT[$cmd]:-}" ]]; then
        msg="unknown command \"$cmd\" for \"soga\""
        s="$(suggest_cmd "$cmd")"
        if [[ -n "$s" ]]; then
            msg+=$'\n\nDid you mean this?'
            while read -r arg; do msg+=$'\n\t'"$arg"; done <<<"$s"
            msg+=$'\n'
        fi
        die_red "$msg"
    fi
    shift

    # 解析参数与选项
    LOG_LINES=1000
    LOG_FOLLOW=0
    while (($# > 0)); do
        arg="$1"
        shift
        if ((only_args)) || [[ "$arg" != -* || "$arg" == "-" ]]; then
            args+=("$arg")
            continue
        fi
        case "$arg" in
            --) only_args=1 ;;
            -h | --help)
                cmd_help "$cmd"
                exit 0
                ;;
            -f | --follow)
                [[ "$cmd" == "log" ]] || die_red "unknown shorthand flag: 'f' in -f"
                LOG_FOLLOW=1
                ;;
            -n | --lines | -n* | --lines=*)
                if [[ "$cmd" != "log" ]]; then
                    [[ "$arg" == --* ]] && die_red "unknown flag: ${arg%%=*}"
                    die_red "unknown shorthand flag: 'n' in $arg"
                fi
                case "$arg" in
                    -n | --lines)
                        (($# > 0)) || die_red "flag needs an argument: 'n' in -n"
                        v="$1"
                        shift
                        ;;
                    --lines=*) v="${arg#--lines=}" ;;
                    *) v="${arg#-n}" ;;
                esac
                [[ "$v" =~ ^[+-]?[0-9]+$ ]] ||
                    die_red "invalid argument \"$v\" for \"-n, --lines\" flag: strconv.ParseInt: parsing \"$v\": invalid syntax"
                v="${v#+}"
                if [[ "$v" == -* ]]; then LOG_LINES=0; else LOG_LINES=$((10#$v)); fi
                ;;
            --*) die_red "unknown flag: ${arg%%=*}" ;;
            *) die_red "unknown shorthand flag: '${arg:1:1}' in $arg" ;;
        esac
    done

    # 参数个数校验
    n=${#args[@]}
    case "$cmd" in
        new | remove)
            ((n >= 1)) || die_red "requires at least 1 arg(s), only received $n"
            ;;
        rename)
            ((n == 2)) || die_red "accepts 2 arg(s), received $n"
            ;;
        start | stop | restart | enable | disable | log | update)
            ((n <= 1)) || die_red "accepts at most 1 arg(s), received $n"
            ;;
        uninstall | self-update | version)
            ((n == 0)) || die_red "unknown command \"${args[0]}\" for \"soga $cmd\""
            ;;
    esac

    require_root || die_red "$ERR"
    check_os || die_red "$ERR"

    "c_${cmd//-/_}" "${args[@]}" || die_red "$ERR"
    exit 0
}

if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
    main "$@"
fi
