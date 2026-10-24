package G6bHD5JA

import (
	"fmt"
	"recovered.local/soga/Ualx5UM_D6I"
)

// VCQ8zrwYebS @0x15a25a0 writes UOT address tags 0/1/2.
// It selects the first resolved IP, otherwise retaining the domain.
// 返回原固定编码长度，即使异常 IP 长度导致 To16 返回 nil。
func VCQ8zrwYebS(b Ualx5UM_D6I.Mf3phSOVqgf, a *AHU2MRmjzQu) (int, error) {
	var n int
	if len(a.oKcMaQDwbH) > 0 {
		ip := a.oKcMaQDwbH[0]
		if ip.To4() != nil {
			b.WriteByte(0)
			b.Write(ip.To4())
			n = 7
		} else {
			b.WriteByte(1)
			b.Write(ip.To16())
			n = 19
		}
	} else {
		if len(a.dlcy_1) > 255 {
			return 0, fmt.Errorf("host too long: %v", a.dlcy_1)
		}
		b.WriteByte(2)
		b.WriteByte(byte(len(a.dlcy_1)))
		b.WriteString(a.dlcy_1)
		n = len(a.dlcy_1) + 4
	}
	b.WriteUint16(uint16(a.pDnAz_))
	return n, nil
}
