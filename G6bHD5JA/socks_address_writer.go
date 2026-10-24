package G6bHD5JA

import (
	"fmt"
	"recovered.local/soga/Ualx5UM_D6I"
)

// DdYroT1n 对应 0x15a1de0。选择第一个已解析 IP，未解析时保留域名；
// 返回原固定编码长度，即使异常 IP 长度导致 To16 返回 nil。
func DdYroT1n(b Ualx5UM_D6I.Mf3phSOVqgf, a *AHU2MRmjzQu) (int, error) {
	var n int
	if len(a.oKcMaQDwbH) > 0 {
		ip := a.oKcMaQDwbH[0]
		if ip.To4() != nil {
			b.WriteByte(1)
			b.Write(ip.To4())
			n = 7
		} else {
			b.WriteByte(4)
			b.Write(ip.To16())
			n = 19
		}
	} else {
		if len(a.dlcy_1) > 255 {
			return 0, fmt.Errorf("host too long: %v", a.dlcy_1)
		}
		b.WriteByte(3)
		b.WriteByte(byte(len(a.dlcy_1)))
		b.WriteString(a.dlcy_1)
		n = len(a.dlcy_1) + 4
	}
	b.WriteUint16(uint16(a.pDnAz_))
	return n, nil
}
