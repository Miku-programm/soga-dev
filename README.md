# soga-dev

soga 节点后端（v2.17.0）的源码，支持 linux/amd64 与 linux/arm64。

## 一键安装

```bash
# 安装 / 更新最新版
bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh)

# 安装指定版本
bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh) 2.17.0

# 安装并写入配置
bash <(curl -Ls https://raw.githubusercontent.com/Miku-programm/soga-dev/main/install.sh) type=xboard server_type=v2ray node_id=1 webapi_url=https://xxx.com/ webapi_key=xxx
```

需要 root 用户和 systemd。安装内容：

| 路径 | 说明 |
| --- | --- |
| `/usr/local/soga/` | soga 程序 |
| `/etc/soga/soga.conf` | default 实例配置，重装时保留 |
| `/etc/soga/{blockList,whiteList,dns.yml,routes.toml}` | 规则文件，重装时保留 |
| `/etc/systemd/system/soga.service`、`soga@.service` | systemd 服务 |
| `/usr/bin/soga` | 管理脚本 |

## 管理命令

```
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
```

### 多实例

- 内置 `default` 实例：`/etc/soga/soga.conf`，服务名 `soga`，不可移除
- 自定义实例 `<name>`：`/etc/soga/<name>/soga.conf`，服务名 `soga@<name>`

`soga new` 以 default 实例的配置和规则文件为模板，并把 `log_file_dir`、`ip_user_cache_save_dir`、`access_log_dir` 指向实例自己的目录，再叠加命令行传入的 `k=v`。

```bash
soga new node2 node_id=2
soga start node2
soga enable node2
soga log node2 -f
```

## 从源码构建

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o soga ./cmd/soga
```

推送 `v*` tag 后，GitHub Actions 会构建 `soga-linux-{amd64,arm64}.tar.gz` 并发布到 Release，一键安装脚本从这里下载。

## 注意

配置项 `auto_update=true` 时，程序内置的自动更新会从 `vaxilu/soga` 的官方 Release 下载并覆盖 `/usr/local/soga/`，而不是本仓库的构建。如需使用本仓库的版本，请保持 `auto_update=false`，用 `soga update` 更新。

## License

[MIT](LICENSE)
