package rtx

import (
	"fmt"
	"net"
)

// N4EF0n2 对应0xb46c60。仅枚举已启用接口，过滤回环及链路本地单播；
// 单个接口Addrs失败时跳过该接口，保留地址顺序、重复和原IP切片。
func N4EF0n2() ([]net.IP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("获取网络接口失败: %v", err)
	}
	var result []net.IP
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		addresses, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, address := range addresses {
			var ip net.IP
			switch a := address.(type) {
			case *net.IPNet:
				ip = a.IP
			case *net.IPAddr:
				ip = a.IP
			default:
				continue
			}
			if !ip.IsLoopback() && !ip.IsLinkLocalUnicast() {
				result = append(result, ip)
			}
		}
	}
	return result, nil
}
