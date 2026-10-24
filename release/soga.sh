#!/usr/bin/env bash
#
# soga 管理脚本，安装后位于 /usr/bin/soga
#
# 实例约定:
#   default 实例: /etc/soga/soga.conf，systemd 服务 soga.service
#   <name>  实例: /etc/soga/<name>/soga.conf，systemd 服务 soga@<name>.service

set -uo pipefail

SOGA_REPO="${SOGA_REPO:-Miku-programm/soga-dev}"
SOGA_BRANCH="${SOGA_BRANCH:-main}"
RAW_BASE="https://raw.githubusercontent.com/${SOGA_REPO}/${SOGA_BRANCH}"

INSTALL_DIR="/usr/local/soga"
CONF_DIR="/etc/soga"
SYSTEMD_DIR="/etc/systemd/system"
RULE_FILES=(blockList whiteList dns.yml routes.toml)

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

info() { echo -e "${green}$*${plain}"; }
warn() { echo -e "${yellow}$*${plain}"; }
err() { echo -e "${red}错误：${plain}$*" >&2; }
die() { err "$*"; exit 1; }

require_root() {
    [[ $EUID -eq 0 ]] || die "必须使用 root 用户运行此命令"
}

# ---------- 实例 ----------

valid_name() {
    [[ "$1" =~ ^[A-Za-z0-9][A-Za-z0-9_.-]{0,63}$ ]]
}

instance_dir() {
    if [[ "$1" == "default" ]]; then echo "$CONF_DIR"; else echo "$CONF_DIR/$1"; fi
}

instance_conf() {
    echo "$(instance_dir "$1")/soga.conf"
}

instance_unit() {
    if [[ "$1" == "default" ]]; then echo "soga"; else echo "soga@$1"; fi
}

instance_exists() {
    [[ -f "$(instance_conf "$1")" ]]
}

require_instance() {
    instance_exists "$1" || die "实例 <$1> 不存在，使用 soga list 查看已有实例"
}

list_instances() {
    local d
    [[ -f "$CONF_DIR/soga.conf" ]] && echo "default"
    for d in "$CONF_DIR"/*/; do
        [[ -f "${d}soga.conf" ]] || continue
        d="${d%/}"
        echo "${d##*/}"
    done
}

instance_state() {
    local unit
    unit="$(instance_unit "$1")"
    systemctl is-active "$unit" 2>/dev/null || true
}

instance_enabled() {
    local unit
    unit="$(instance_unit "$1")"
    systemctl is-enabled "$unit" 2>/dev/null || true
}

# ---------- 配置 ----------

set_config() {
    local conf_file="$1" kv="$2" key value esc
    key="${kv%%=*}"
    value="${kv#*=}"
    if [[ -z "$key" || "$key" == "$kv" || ! "$key" =~ ^[A-Za-z0-9_]+$ ]]; then
        warn "忽略无效配置参数: ${kv}"
        return 1
    fi
    esc="$(printf '%s' "$value" | sed -e 's/[\\|&]/\\&/g')"
    if grep -Eq "^[[:space:]]*${key}[[:space:]]*=" "$conf_file"; then
        sed -i -E "s|^[[:space:]]*${key}[[:space:]]*=.*|${key}=${esc}|" "$conf_file"
        info "已更新配置: ${key}=${value}"
    else
        echo "${key}=${value}" >> "$conf_file"
        info "已添加配置: ${key}=${value}"
    fi
}

get_config() {
    local conf_file="$1" key="$2"
    grep -E "^[[:space:]]*${key}[[:space:]]*=" "$conf_file" | tail -n 1 | sed -E 's/^[^=]*=[[:space:]]*//'
}

# ---------- 命令 ----------

cmd_list() {
    local name state enabled
    printf '%-20s %-12s %-10s %s\n' "实例" "状态" "开机自启" "配置文件"
    while read -r name; do
        [[ -z "$name" ]] && continue
        state="$(instance_state "$name")"
        enabled="$(instance_enabled "$name")"
        printf '%-20s %-12s %-10s %s\n' "$name" "${state:-unknown}" "${enabled:-unknown}" "$(instance_conf "$name")"
    done < <(list_instances)
}

cmd_new() {
    require_root
    local name="${1:-}"
    [[ -n "$name" ]] || die "用法: soga new <name> [k1=v1 ...]"
    shift
    valid_name "$name" || die "实例名只能包含字母、数字、. _ -，且以字母或数字开头"
    [[ "$name" != "default" ]] || die "default 为内置实例名"
    instance_exists "$name" && die "实例 <$name> 已存在"
    [[ -e "$CONF_DIR/$name" ]] && die "$CONF_DIR/$name 已存在且不是 soga 实例"

    local dir conf f
    dir="$(instance_dir "$name")"
    conf="$(instance_conf "$name")"
    mkdir -p "$dir"
    if [[ -f "$CONF_DIR/soga.conf" ]]; then
        cp "$CONF_DIR/soga.conf" "$conf"
    else
        cp "$INSTALL_DIR/soga.conf" "$conf"
    fi
    for f in "${RULE_FILES[@]}"; do
        if [[ -f "$CONF_DIR/$f" ]]; then
            cp "$CONF_DIR/$f" "$dir/"
        elif [[ -f "$INSTALL_DIR/$f" ]]; then
            cp "$INSTALL_DIR/$f" "$dir/"
        fi
    done
    # 每个实例的日志、IP 缓存、访问日志放在自己的目录，避免多实例互相覆盖
    set_config "$conf" "log_file_dir=$dir/" >/dev/null
    set_config "$conf" "ip_user_cache_save_dir=$dir/" >/dev/null
    set_config "$conf" "access_log_dir=$dir/access_log" >/dev/null

    local kv
    for kv in "$@"; do
        set_config "$conf" "$kv"
    done
    systemctl daemon-reload
    info "实例 <$name> 创建成功，配置文件: $conf"
    echo "使用 soga start $name 启动，soga enable $name 设置开机自启"
}

cmd_remove() {
    require_root
    [[ $# -gt 0 ]] || die "用法: soga remove <name> [name...]"
    local name unit
    for name in "$@"; do
        if [[ "$name" == "default" ]]; then
            err "default 实例不可移除"
            continue
        fi
        if ! instance_exists "$name"; then
            err "实例 <$name> 不存在"
            continue
        fi
        unit="$(instance_unit "$name")"
        systemctl stop "$unit" >/dev/null 2>&1
        systemctl disable "$unit" >/dev/null 2>&1
        rm -rf "$(instance_dir "$name")"
        info "实例 <$name> 已移除"
    done
    systemctl daemon-reload
}

cmd_rename() {
    require_root
    local old="${1:-}" new="${2:-}"
    [[ -n "$old" && -n "$new" ]] || die "用法: soga rename <old_name> <new_name>"
    [[ "$old" != "default" && "$new" != "default" ]] || die "default 实例不可重命名"
    valid_name "$new" || die "实例名只能包含字母、数字、. _ -，且以字母或数字开头"
    require_instance "$old"
    instance_exists "$new" && die "实例 <$new> 已存在"
    [[ -e "$CONF_DIR/$new" ]] && die "$CONF_DIR/$new 已存在"

    local old_unit new_unit was_active was_enabled old_dir new_dir conf key value
    old_unit="$(instance_unit "$old")"
    new_unit="$(instance_unit "$new")"
    old_dir="$(instance_dir "$old")"
    new_dir="$(instance_dir "$new")"
    was_active="$(instance_state "$old")"
    was_enabled="$(instance_enabled "$old")"

    systemctl stop "$old_unit" >/dev/null 2>&1
    systemctl disable "$old_unit" >/dev/null 2>&1
    mv "$old_dir" "$new_dir"

    # 指向旧实例目录的路径配置一并改到新目录
    conf="$(instance_conf "$new")"
    for key in log_file_dir ip_user_cache_save_dir access_log_dir; do
        value="$(get_config "$conf" "$key")"
        if [[ "$value" == "$old_dir" || "$value" == "$old_dir/"* ]]; then
            set_config "$conf" "${key}=${new_dir}${value#"$old_dir"}" >/dev/null
        fi
    done

    systemctl daemon-reload
    [[ "$was_enabled" == "enabled" ]] && systemctl enable "$new_unit" >/dev/null 2>&1
    [[ "$was_active" == "active" ]] && systemctl start "$new_unit"
    info "实例 <$old> 已重命名为 <$new>"
}

cmd_service() {
    require_root
    local action="$1" name="${2:-default}" unit
    require_instance "$name"
    unit="$(instance_unit "$name")"
    case "$action" in
        start | restart)
            systemctl "$action" "$unit"
            sleep 2
            if [[ "$(instance_state "$name")" == "active" ]]; then
                info "实例 <$name> ${action} 成功"
            else
                err "实例 <$name> 可能启动失败，请使用 soga log $name 查看日志"
                return 1
            fi
            ;;
        stop)
            systemctl stop "$unit" && info "实例 <$name> 已停止"
            ;;
        enable)
            systemctl enable "$unit" >/dev/null 2>&1 && info "实例 <$name> 已设置开机自启"
            ;;
        disable)
            systemctl disable "$unit" >/dev/null 2>&1 && info "实例 <$name> 已取消开机自启"
            ;;
    esac
}

cmd_log() {
    local name="default" follow=0 arg
    for arg in "$@"; do
        if [[ "$arg" == "-f" ]]; then follow=1; else name="$arg"; fi
    done
    require_instance "$name"
    if [[ $follow -eq 1 ]]; then
        journalctl -u "$(instance_unit "$name")" -n 1000 -f --no-pager
    else
        journalctl -u "$(instance_unit "$name")" -n 1000 --no-pager
    fi
}

cmd_config() {
    local name="default"
    if [[ $# -gt 0 && "$1" != *=* ]]; then
        name="$1"
        shift
    fi
    require_instance "$name"
    local conf
    conf="$(instance_conf "$name")"
    if [[ $# -eq 0 ]]; then
        cat "$conf"
        return
    fi
    require_root
    local kv
    for kv in "$@"; do
        set_config "$conf" "$kv"
    done
    if [[ "$name" == "default" ]]; then
        echo "配置已修改，使用 soga restart 使其生效"
    else
        echo "配置已修改，使用 soga restart $name 使其生效"
    fi
}

cmd_update() {
    require_root
    local running=() name
    while read -r name; do
        [[ "$name" != "default" && "$(instance_state "$name")" == "active" ]] && running+=("$name")
    done < <(list_instances)

    bash <(curl -fLs "$RAW_BASE/install.sh") "$@" || die "更新失败"

    for name in "${running[@]}"; do
        systemctl restart "$(instance_unit "$name")" && info "实例 <$name> 已重启"
    done
}

cmd_uninstall() {
    require_root
    if [[ "${1:-}" != "-y" ]]; then
        read -rp "确定要卸载 soga 吗? 所有实例都将被停止 [y/N]: " ans
        [[ "$ans" == "y" || "$ans" == "Y" ]] || { echo "已取消"; return; }
    fi
    local name
    while read -r name; do
        systemctl stop "$(instance_unit "$name")" >/dev/null 2>&1
        systemctl disable "$(instance_unit "$name")" >/dev/null 2>&1
    done < <(list_instances)
    rm -f "$SYSTEMD_DIR/soga.service" "$SYSTEMD_DIR/soga@.service"
    systemctl daemon-reload
    systemctl reset-failed >/dev/null 2>&1
    rm -rf "$INSTALL_DIR"

    local purge="n"
    if [[ "${1:-}" == "-y" ]]; then
        purge="${2:-n}"
    else
        read -rp "是否同时删除配置目录 ${CONF_DIR}? [y/N]: " purge
    fi
    if [[ "$purge" == "y" || "$purge" == "Y" || "$purge" == "--purge" ]]; then
        rm -rf "$CONF_DIR"
        info "已删除 ${CONF_DIR}"
    else
        echo "配置目录 ${CONF_DIR} 已保留"
    fi
    rm -f /usr/bin/soga
    info "soga 已卸载"
}

cmd_self_update() {
    require_root
    local tmp
    tmp="$(mktemp)"
    if curl -fLs -o "$tmp" "$RAW_BASE/release/soga.sh" && bash -n "$tmp"; then
        install -m 0755 "$tmp" /usr/bin/soga
        rm -f "$tmp"
        info "管理脚本已更新"
    else
        rm -f "$tmp"
        die "下载管理脚本失败"
    fi
}

cmd_version() {
    if [[ -x "$INSTALL_DIR/soga" ]]; then
        echo "soga v$("$INSTALL_DIR/soga" -v)"
    else
        die "soga 未安装"
    fi
}

show_usage() {
    cat <<'EOF'
soga 管理脚本使用方法:
----------------------------------------------------------
soga                              - 显示管理菜单 (功能更多)
soga list                         - 列出已有的实例及状态
soga new <name> [k1=v1 ...]       - 新建实例
soga remove <name> [name...]      - 移除实例
soga rename <old_name> <new_name> - 重命名实例
soga start [name]                 - 启动实例
soga stop [name]                  - 停止实例
soga restart [name]               - 重启实例
soga enable [name]                - 设置实例开机自启
soga disable [name]               - 取消实例开机自启
soga log [name]                   - 查看实例日志(默认 1000 行)
soga log [name] -f                - 持续输出实例日志
soga config [name]                - 查看实例配置文件
soga config [name] k1=v1 k2=v2    - 修改或新增配置
soga update [version]             - 安装/更新 soga
soga uninstall                    - 完全卸载 soga
soga self-update                  - 更新管理脚本
soga version                      - 查看版本
----------------------------------------------------------
EOF
}

ask_name() {
    local name
    read -rp "实例名 (留空为 default): " name
    echo "${name:-default}"
}

show_menu() {
    local choice name args
    while true; do
        echo
        echo -e "${green}soga 管理菜单${plain}"
        echo "----------------------------------------------"
        echo "  1. 查看实例列表"
        echo "  2. 启动实例"
        echo "  3. 停止实例"
        echo "  4. 重启实例"
        echo "  5. 查看实例日志"
        echo "  6. 查看实例配置"
        echo "  7. 修改实例配置"
        echo "  8. 新建实例"
        echo "  9. 移除实例"
        echo " 10. 设置开机自启"
        echo " 11. 取消开机自启"
        echo " 12. 更新 soga"
        echo " 13. 更新管理脚本"
        echo " 14. 卸载 soga"
        echo " 15. 查看版本"
        echo "  0. 退出"
        echo "----------------------------------------------"
        read -rp "请输入选择: " choice
        case "$choice" in
            1) cmd_list ;;
            2) cmd_service start "$(ask_name)" ;;
            3) cmd_service stop "$(ask_name)" ;;
            4) cmd_service restart "$(ask_name)" ;;
            5) cmd_log "$(ask_name)" ;;
            6) cmd_config "$(ask_name)" ;;
            7)
                name="$(ask_name)"
                read -rp "输入配置 (k1=v1 k2=v2): " -a args
                cmd_config "$name" "${args[@]}"
                ;;
            8)
                read -rp "新实例名: " name
                read -rp "额外配置 (k1=v1 k2=v2，可留空): " -a args
                cmd_new "$name" "${args[@]}"
                ;;
            9)
                read -rp "要移除的实例名: " name
                cmd_remove "$name"
                ;;
            10) cmd_service enable "$(ask_name)" ;;
            11) cmd_service disable "$(ask_name)" ;;
            12)
                read -rp "版本号 (留空为最新版): " name
                cmd_update ${name:+"$name"}
                ;;
            13) cmd_self_update ;;
            14) cmd_uninstall ;;
            15) cmd_version ;;
            0) exit 0 ;;
            *) err "无效选择" ;;
        esac
    done
}

main() {
    if [[ $# -eq 0 ]]; then
        show_menu
        return
    fi
    local cmd="$1"
    shift
    case "$cmd" in
        list) cmd_list ;;
        new) cmd_new "$@" ;;
        remove) cmd_remove "$@" ;;
        rename) cmd_rename "$@" ;;
        start | stop | restart | enable | disable) cmd_service "$cmd" "$@" ;;
        log) cmd_log "$@" ;;
        config) cmd_config "$@" ;;
        update) cmd_update "$@" ;;
        uninstall) cmd_uninstall "$@" ;;
        self-update) cmd_self_update ;;
        version | -v | --version) cmd_version ;;
        help | -h | --help) show_usage ;;
        *)
            err "未知命令: $cmd"
            show_usage
            exit 1
            ;;
    esac
}

main "$@"
