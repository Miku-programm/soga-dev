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
| `/usr/local/soga/soga` | soga 程序 |
| `/etc/soga/soga.conf` | default 实例配置，重装时保留 |
| `/etc/soga/{blockList,whiteList,dns.yml,routes.toml}` | 规则文件，重装时保留 |
| `/etc/systemd/system/soga.service`、`soga@.service` | systemd 服务 |
| `/usr/bin/soga` | 管理脚本（行为与官方 soga-cmd v0.0.6 一致） |
| `~/.soga/config.json` | 实例清单、安装分支与更新检查缓存 |

安装时会像官方脚本一样补齐基础依赖（wget、curl、tar、socat、cron）并安装 acme.sh。

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

直接执行 `soga` 进入交互菜单：安装/更新、完全卸载、切换安装分支（稳定版/开发版）、更新管理脚本、新增实例（带快速配置向导），以及每个实例的启动/停止/重启/重载、状态与日志、开机自启、快速配置和编辑配置文件。菜单提示处也可以直接输入实例名进入该实例的管理菜单。

### 多实例

- 内置 `default` 实例：`/etc/soga/soga.conf`，服务名 `soga`，不可移除
- 自定义实例 `<name>`：`/etc/soga/<name>/soga.conf`，服务名 `soga@<name>`，实例名只能由字母和数字组成

`soga new` 以 default 实例的配置为模板，把 `log_file_dir`、`ip_user_cache_save_dir` 指向实例自己的目录，再叠加命令行传入的 `k=v`；规则文件在实例目录下新建为空文件。

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

## 访问日志

`access_log_enable=true` 时每个连接记录一条，`access_host` 是客户端原始请求的目标（sniff 改写前），`host` 是 sniff 识别出的域名：

```
时间|协议|node_id|user_id|user_ip|access_host|host|dest_ip|端口
2026-08-22 09:51:44|tcp|12|1001|1.2.3.4|31.13.92.37|www.facebook.com|31.13.92.37|443
2026-08-22 09:52:10|tcp|12|1002|5.6.7.8|www.facebook.com|www.facebook.com|31.13.92.37|443
```

第一条 `access_host` 是 IP、`host` 是域名，即直连 IP 并伪装 SNI；第二条是正常按域名访问。ClickHouse 模式对应新增 `access_host` 列，开启 `access_log_clickhouse_auto_create` 时会自动给旧表补上该列。

## 注意

配置项 `auto_update=true` 时，程序每 24 小时检查一次本仓库的最新 Release，发布超过 72 小时的新版本会自动下载、覆盖 `/usr/local/soga/` 并执行 `soga restart`。

## License

[MIT](LICENSE)
