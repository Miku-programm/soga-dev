package BaEWsPy

import "net"

// ResolveDefault 暴露原全局解析器 0x5096320，供原 SQL 初始化路径使用。
func ResolveDefault(host string) ([]net.IP, error) {
	return defaultResolverAt5096320.Resolve(host)
}
