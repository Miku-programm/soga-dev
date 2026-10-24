// 从原样本 0x169d140、0x169d1c0、0x169d280 恢复。
package connutil

import (
	"bytes"
	"os"
)

func a6N3H7(path string) (int, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return bytes.Count(b, []byte{'\n'}) - 1, nil
}

// 原实现保留 IPv4 读取错误，忽略 IPv6 读取错误。
func QMgUaHjx4G() (int, error) {
	v4, err := a6N3H7("/proc/net/tcp")
	if err != nil {
		return v4, err
	}
	v6, _ := a6N3H7("/proc/net/tcp6")
	return v4 + v6, nil
}
func NEfQI6r99p8() (int, error) {
	v4, err := a6N3H7("/proc/net/udp")
	if err != nil {
		return v4, err
	}
	v6, _ := a6N3H7("/proc/net/udp6")
	return v4 + v6, nil
}
