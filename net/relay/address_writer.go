package relay

import (
	"fmt"
	"recovered.local/soga/buf/bufiface"
)

// HfFNot 对应 0x15a21c0：先端口，再地址类型与地址。错误发生前仍写入端口。
func HfFNot(b bufiface.Mf3phSOVqgf, a *AHU2MRmjzQu) (int, error) {
	b.WriteUint16(uint16(a.pDnAz_))
	if len(a.oKcMaQDwbH) != 0 {
		ip := a.oKcMaQDwbH[0]
		if v4 := ip.To4(); v4 != nil {
			b.WriteByte(1)
			b.Write(v4)
			return 7, nil
		}
		b.WriteByte(3)
		b.Write(ip.To16())
		return 19, nil
	}
	if len(a.dlcy_1) > 255 {
		return 0, fmt.Errorf("host too long: %v", a.dlcy_1)
	}
	b.WriteByte(2)
	b.WriteByte(uint8(len(a.dlcy_1)))
	b.WriteString(a.dlcy_1)
	return len(a.dlcy_1) + 4, nil
}
