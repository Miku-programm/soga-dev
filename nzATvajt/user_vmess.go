package nzATvajt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/JafvCs5"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"sync/atomic"
)

// QKaIWV2 preserves the original 40-byte VMess user layout (RTTI 0x37221d8).
type QKaIWV2 struct {
	*AhSFeSua
	nmOHVD      *LFWTTXF7JR.Yn8BJoL
	fCmumUdZX6m atomic.Pointer[IC_1xll]
	gzCqh4BsIiD cipher.Block
}

func EOkO09L(p *LFWTTXF7JR.Aga91y, node *G6bHD5JA.L337vk2G) (*QKaIWV2, error) { // 0x18287e0
	if _, ok := p.JAfXWwzYrW.(*LFWTTXF7JR.Yn8BJoL); !ok {
		return nil, pNF7CSleWS(p.WZ89pTAOrHK4, "vmess")
	}
	u := &QKaIWV2{AhSFeSua: cLLKJ_A31(p, node)}
	empty := IC_1xll{}
	u.fCmumUdZX6m.Store(&empty)
	u.UpdatePanelUser(p)
	return u, nil
}

func (u *QKaIWV2) UpdatePanelUser(p *LFWTTXF7JR.Aga91y) (IC_1xll, bool) { // 0x1828900
	u.UpdateUser(p)
	c, ok := p.JAfXWwzYrW.(*LFWTTXF7JR.Yn8BJoL)
	if !ok {
		return nil, false
	}
	if cNT24rG4F.VmessAuthMode() == "aead" {
		c.FvQ2ltKOM7 = 0
	}
	if u.nmOHVD != nil && *u.nmOHVD == *c {
		return nil, false
	}
	old := *u.fCmumUdZX6m.Load()
	u.teCoJgj5(c.SfiEGpJo, c.FvQ2ltKOM7)
	credential := *c
	u.nmOHVD = &credential
	if c.FvQ2ltKOM7 != 0 {
		u.gzCqh4BsIiD = nil
	} else {
		key := JafvCs5.CVqKigbE26(u.GetCmdKey(), []byte("AES Auth ID Encryption"))
		u.gzCqh4BsIiD, _ = aes.NewCipher(key[:16])
	}
	u.ResetCtx()
	return old, true
}

func (u *QKaIWV2) teCoJgj5(text string, alterID int) { // 0x1828ec0
	if alterID < 0 {
		empty := IC_1xll{}
		u.fCmumUdZX6m.Store(&empty)
		return
	}
	ids := make(IC_1xll, 0, alterID+1)
	id, err := w6MaVR(text)
	if err != nil {
		muc279Mpa6S4.V8GN5Stk(8, "invalid uuid %s by user id: %d\n", text, u.djdKrU)
		return // Retain the old IDs; UpdatePanelUser still updates its other fields.
	}
	ids = append(ids, id)
	for remaining := alterID; remaining > 0; remaining-- {
		next := new(DALfaCvQ)
		xshVaRu(id.Yfd33z6YYdJa[:], next.Yfd33z6YYdJa[:])
		ids = append(ids, next)
		id = next
	}
	u.fCmumUdZX6m.Store(&ids)
}
func (u *QKaIWV2) GetCmdKey() []byte { // 0x18292a0
	ids := *u.fCmumUdZX6m.Load()
	if len(ids) > 0 {
		return ids[0].GetCmdKey()
	}
	return nil
}
func (u *QKaIWV2) GetBlock() cipher.Block { return u.gzCqh4BsIiD }         // 0x1829300
func (u *QKaIWV2) GetUserIds() IC_1xll    { return *u.fCmumUdZX6m.Load() } // 0x1829320
func (u *QKaIWV2) String() string { // 0x1829340
	return fmt.Sprintf("id: %v, alterId: %v, uuid: %v", u.djdKrU, u.nmOHVD.FvQ2ltKOM7, u.nmOHVD.SfiEGpJo)
}
