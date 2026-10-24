package G6bHD5JA

import "net"

// HYeSYzPzN 对应原函数 0x15b3740。前缀为完整长度时直接返回原切片，
// 因而 IPv4 映射地址不会被无条件缩成 4 字节；非法前缀保留 net.IP.Mask 的 nil 结果。
func HYeSYzPzN(ip net.IP, ipv4Prefix, ipv6Prefix int) net.IP {
	if v4 := ip.To4(); v4 != nil {
		if ipv4Prefix != 32 {
			return v4.Mask(net.CIDRMask(ipv4Prefix, 32))
		}
	} else if ipv6Prefix != 128 {
		return ip.Mask(net.CIDRMask(ipv6Prefix, 128))
	}
	return ip
}
