package nzATvajt

import (
	"fmt"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/q5bqCgTdxaTM"
)

// DWaKIV1T is the original 72-byte Trojan user (RTTI 0x371fb60).
type DWaKIV1T struct {
	*AhSFeSua
	drXNEIJ6 *LFWTTXF7JR.JxTGUar
	IdeY3L   [56]byte
}

func KUa7qViw(p *LFWTTXF7JR.Aga91y, node *G6bHD5JA.L337vk2G) (*DWaKIV1T, error) { // 0x18245a0
	if _, ok := p.JAfXWwzYrW.(*LFWTTXF7JR.JxTGUar); !ok {
		return nil, pNF7CSleWS(p.WZ89pTAOrHK4, "trojan")
	}
	u := &DWaKIV1T{AhSFeSua: cLLKJ_A31(p, node)}
	u.UpdatePanelUser(p)
	return u, nil
}

func (u *DWaKIV1T) UpdatePanelUser(p *LFWTTXF7JR.Aga91y) *[56]byte { // 0x1824680
	u.UpdateUser(p)
	c, ok := p.JAfXWwzYrW.(*LFWTTXF7JR.JxTGUar)
	if !ok {
		return nil
	}
	if u.drXNEIJ6 != nil && *u.drXNEIJ6 == *c {
		return nil
	}
	previous := u.IdeY3L
	if c.ARARP1z0yA != "" {
		copy(u.IdeY3L[:], q5bqCgTdxaTM.J26YsI2DwxPQ(c.ARARP1z0yA))
	} else {
		copy(u.IdeY3L[:], c.QiUSI27)
	}
	// A short precomputed value intentionally leaves the old tail in place.
	credential := *c
	u.drXNEIJ6 = &credential
	u.ResetCtx()
	return &previous
}
func (u *DWaKIV1T) String() string { // 0x18248c0
	return fmt.Sprintf("id: %v, password: %v", u.djdKrU, u.drXNEIJ6.ARARP1z0yA)
}
