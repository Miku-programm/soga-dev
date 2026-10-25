#!/usr/bin/env bash
#
# soga 一键安装 / 更新脚本
#
#   bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh)
#   bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh) 2.17.0
#   bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh) type=v2board server_type=v2ray node_id=1 webapi_url=https://xxx.com webapi_key=xxx
#
# 先安装管理脚本 /usr/bin/soga，再由 soga update 安装 soga 程序与 systemd 服务。
#
# 环境变量:
#   SOGA_REPO     下载 release 的 GitHub 仓库，默认 Miku-programm/soga-dev
#   SOGA_BRANCH   管理脚本所在分支，默认 main
#   SOGA_TARBALL  使用本地 soga-linux-<arch>.tar.gz 安装，跳过下载

set -uo pipefail

SOGA_REPO="${SOGA_REPO:-Miku-programm/soga-dev}"
SOGA_BRANCH="${SOGA_BRANCH:-main}"
SOGA_TARBALL="${SOGA_TARBALL:-}"
export SOGA_REPO SOGA_BRANCH SOGA_TARBALL

CONF_DIR="/etc/soga"
CMD_PATH="/usr/bin/soga"

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

info() { echo -e "${green}$*${plain}"; }
warn() { echo -e "${yellow}$*${plain}"; }
die() { echo -e "${red}错误：${plain}$*" >&2; exit 1; }

[[ $EUID -ne 0 ]] && die "必须使用 root 用户运行此脚本"
command -v systemctl >/dev/null 2>&1 || die "systemctl 命令不存在，请使用较新版本的系统，例如 Ubuntu 18+、Debian 9+、CentOS 7+"

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
    else
        die "无法自动安装 ${missing[*]}，请手动安装后重试"
    fi
}

# 安装管理脚本：本地安装包中自带 soga.sh 时优先使用，否则从仓库下载
install_manager() {
    local tmp
    tmp="$(mktemp -d)"
    if [[ -n "$SOGA_TARBALL" ]] && tar --no-same-owner -xzf "$SOGA_TARBALL" -C "$tmp" soga/soga.sh 2>/dev/null; then
        install -m 0755 "$tmp/soga/soga.sh" "$CMD_PATH" || die "安装管理脚本失败"
    else
        curl -fsSL --retry 3 -o "$tmp/soga.sh" "https://raw.githubusercontent.com/${SOGA_REPO}/${SOGA_BRANCH}/release/soga.sh" \
            || die "下载 soga 管理脚本失败，请确保你的服务器能够下载 Github 的文件"
        bash -n "$tmp/soga.sh" || die "下载的管理脚本无效"
        install -m 0755 "$tmp/soga.sh" "$CMD_PATH" || die "安装管理脚本失败"
    fi
    rm -rf "$tmp"
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
    local version="" configs=() arg fresh=0
    for arg in "$@"; do
        if [[ "$arg" == *=* ]]; then
            configs+=("$arg")
        elif [[ -z "$version" ]]; then
            version="$arg"
        fi
    done

    install_base
    install_manager

    # soga update 只在 /etc/soga 不存在时创建 default 实例，残留的空目录先清掉
    if [[ ! -f "$CONF_DIR/soga.conf" ]]; then
        fresh=1
        rmdir "$CONF_DIR" 2>/dev/null
    fi

    if [[ -n "$version" ]]; then
        "$CMD_PATH" update "$version" || exit 1
    else
        "$CMD_PATH" update || exit 1
    fi
    [[ -f "$CONF_DIR/soga.conf" ]] || die "${CONF_DIR} 已存在但缺少 soga.conf，请检查后重试"

    if [[ ${#configs[@]} -gt 0 ]]; then
        info "开始写入配置"
        "$CMD_PATH" config "${configs[@]}" || exit 1
    fi

    systemctl enable soga >/dev/null 2>&1
    info "soga $("$CMD_PATH" version | sed -n 's/^soga 程序: //p') 安装完成，已设置开机自启"

    if [[ $fresh -eq 1 && ${#configs[@]} -eq 0 ]]; then
        echo
        warn "全新安装，请先配置 ${CONF_DIR}/soga.conf（或执行 soga 进入菜单使用「快速配置」），然后执行 soga start"
    else
        systemctl restart soga
        sleep 2
        if systemctl is-active --quiet soga; then
            info "soga 启动成功"
        else
            warn "soga 可能启动失败，请使用 soga log 查看日志"
        fi
    fi
    show_usage
}

main "$@"
