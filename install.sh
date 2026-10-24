#!/usr/bin/env bash
#
# soga 一键安装 / 更新脚本
#
#   bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh)
#   bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh) 2.17.0
#   bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh) type=v2board server_type=v2ray node_id=1 webapi_url=https://xxx.com webapi_key=xxx
#
# 环境变量:
#   SOGA_REPO     下载 release 的 GitHub 仓库，默认 Miku-programm/soga-dev
#   SOGA_TARBALL  使用本地 soga-linux-<arch>.tar.gz 安装，跳过下载

set -uo pipefail

SOGA_REPO="${SOGA_REPO:-Miku-programm/soga-dev}"
SOGA_TARBALL="${SOGA_TARBALL:-}"

INSTALL_DIR="/usr/local/soga"
CONF_DIR="/etc/soga"
SYSTEMD_DIR="/etc/systemd/system"

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

info() { echo -e "${green}$*${plain}"; }
warn() { echo -e "${yellow}$*${plain}"; }
die() { echo -e "${red}错误：${plain}$*" >&2; exit 1; }

[[ $EUID -ne 0 ]] && die "必须使用 root 用户运行此脚本"
command -v systemctl >/dev/null 2>&1 || die "systemctl 命令不存在，请使用较新版本的系统，例如 Ubuntu 18+、Debian 9+、CentOS 7+"

detect_arch() {
    case "$(uname -m)" in
        x86_64 | x64 | amd64) echo "amd64" ;;
        aarch64 | arm64) echo "arm64" ;;
        *) die "不支持的架构: $(uname -m)，仅支持 amd64 与 arm64" ;;
    esac
}

install_base() {
    local missing=()
    for cmd in curl tar; do
        command -v "$cmd" >/dev/null 2>&1 || missing+=("$cmd")
    done
    [[ ${#missing[@]} -eq 0 ]] && return
    info "安装依赖: ${missing[*]}"
    if command -v apt-get >/dev/null 2>&1; then
        apt-get update -y && apt-get install -y "${missing[@]}"
    elif command -v dnf >/dev/null 2>&1; then
        dnf install -y "${missing[@]}"
    elif command -v yum >/dev/null 2>&1; then
        yum install -y "${missing[@]}"
    elif command -v apk >/dev/null 2>&1; then
        apk add --no-cache "${missing[@]}"
    else
        die "无法自动安装 ${missing[*]}，请手动安装后重试"
    fi
}

# 0: running, 1: not running, 2: not installed
check_status() {
    [[ -f ${SYSTEMD_DIR}/soga.service ]] || return 2
    systemctl is-active --quiet soga && return 0
    return 1
}

download_tarball() {
    local version="$1" arch="$2" dest="$3" url
    if [[ -n "$SOGA_TARBALL" ]]; then
        info "使用本地安装包: ${SOGA_TARBALL}"
        cp -f "$SOGA_TARBALL" "$dest" || die "复制本地安装包失败"
        return
    fi
    if [[ -z "$version" ]]; then
        url="https://github.com/${SOGA_REPO}/releases/latest/download/soga-linux-${arch}.tar.gz"
        info "开始安装 soga 最新版"
    else
        url="https://github.com/${SOGA_REPO}/releases/download/v${version}/soga-linux-${arch}.tar.gz"
        info "开始安装 soga v${version}"
    fi
    curl -fL --retry 3 -o "$dest" "$url" \
        || die "下载 soga 失败: ${url}\n请确认该版本存在，且服务器能够访问 GitHub"
}

install_soga() {
    local version="${1#v}" arch tmp
    arch="$(detect_arch)"
    echo "架构: ${arch}"

    TMP_DIR="$(mktemp -d)"
    trap 'rm -rf "$TMP_DIR"' EXIT
    tmp="$TMP_DIR"
    download_tarball "$version" "$arch" "$tmp/soga.tar.gz"
    tar -xzf "$tmp/soga.tar.gz" -C "$tmp" || die "解压安装包失败"
    [[ -x "$tmp/soga/soga" ]] || die "安装包内缺少 soga/soga"

    # 二进制与附带文件整体替换；/etc/soga 下的配置不受影响
    rm -rf "$INSTALL_DIR"
    mv "$tmp/soga" "$INSTALL_DIR"
    chmod +x "$INSTALL_DIR/soga"

    mkdir -p "$CONF_DIR"
    local fresh=0
    if [[ ! -f "$CONF_DIR/soga.conf" ]]; then
        cp "$INSTALL_DIR/soga.conf" "$CONF_DIR/"
        fresh=1
    fi
    for f in blockList whiteList dns.yml routes.toml; do
        [[ -f "$CONF_DIR/$f" ]] || cp "$INSTALL_DIR/$f" "$CONF_DIR/"
    done

    cp -f "$INSTALL_DIR/soga.service" "$INSTALL_DIR/soga@.service" "$SYSTEMD_DIR/"
    install -m 0755 "$INSTALL_DIR/soga.sh" /usr/bin/soga
    systemctl daemon-reload
    systemctl enable soga >/dev/null 2>&1

    info "soga v$("$INSTALL_DIR/soga" -v) 安装完成，已设置开机自启"
    FRESH_INSTALL=$fresh
}

apply_configs() {
    local conf_file="$CONF_DIR/soga.conf" kv key value esc
    for kv in "$@"; do
        key="${kv%%=*}"
        value="${kv#*=}"
        if [[ -z "$key" || "$key" == "$kv" || ! "$key" =~ ^[A-Za-z0-9_]+$ ]]; then
            warn "忽略无效配置参数: ${kv}"
            continue
        fi
        esc="$(printf '%s' "$value" | sed -e 's/[\\|&]/\\&/g')"
        if grep -Eq "^[[:space:]]*${key}[[:space:]]*=" "$conf_file"; then
            sed -i -E "s|^[[:space:]]*${key}[[:space:]]*=.*|${key}=${esc}|" "$conf_file"
            info "已更新配置: ${key}=${value}"
        else
            echo "${key}=${value}" >> "$conf_file"
            info "已添加配置: ${key}=${value}"
        fi
    done
}

restart_and_report() {
    systemctl restart soga
    sleep 2
    if check_status; then
        info "soga 启动成功"
    else
        warn "soga 可能启动失败，请使用 soga log 查看日志"
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

main() {
    local version="" configs=() arg
    for arg in "$@"; do
        if [[ "$arg" == *=* ]]; then
            configs+=("$arg")
        elif [[ -z "$version" ]]; then
            version="$arg"
        fi
    done

    FRESH_INSTALL=0
    install_base
    install_soga "$version"

    if [[ ${#configs[@]} -gt 0 ]]; then
        info "开始写入配置"
        apply_configs "${configs[@]}"
    fi

    if [[ $FRESH_INSTALL -eq 1 && ${#configs[@]} -eq 0 ]]; then
        echo
        warn "全新安装，请先配置 ${CONF_DIR}/soga.conf，然后执行 soga start"
    else
        restart_and_report
    fi
    show_usage
}

main "$@"
