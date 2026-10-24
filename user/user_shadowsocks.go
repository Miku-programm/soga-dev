package user

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/crypto/ssmethod"
)

// WtK9HQSy66 retains the original 80-byte Shadowsocks user layout.
type WtK9HQSy66 struct {
	*AhSFeSua
	aTswvh0     ssmethod.C6q3Upu
	bRxpsrh13ry string
	n5IKaq      []byte
	we1MNFya    cipher.Block
}

func MbpJoYH(p *types.Aga91y, node *relay.L337vk2G) (*WtK9HQSy66, error) { // 0x1820bc0
	c, ok := p.JAfXWwzYrW.(*types.YlfOV0DBBnx)
	if !ok {
		return nil, pNF7CSleWS(p.WZ89pTAOrHK4, "ss")
	}
	u := &WtK9HQSy66{AhSFeSua: cLLKJ_A31(p, node), aTswvh0: c.T1oDDIgYJHh}
	if _, err := u.UpdatePanelUser(p); err != nil {
		return nil, err
	}
	return u, nil
}
func (u *WtK9HQSy66) UpdatePanelUser(p *types.Aga91y) (bool, error) { // 0x1820ce0
	u.UpdateUser(p)
	c, ok := p.JAfXWwzYrW.(*types.YlfOV0DBBnx)
	if !ok {
		return false, pNF7CSleWS(p.WZ89pTAOrHK4, "ss")
	}
	if u.bRxpsrh13ry == c.A_wNUTkMh6pH {
		return false, nil
	}
	u.bRxpsrh13ry = c.A_wNUTkMh6pH
	if err := u.qNsBtysF(); err != nil {
		return false, err
	}
	u.ResetCtx()
	return true, nil
}
func (u *WtK9HQSy66) qNsBtysF() error { // 0x1820e20
	key, err := u.aTswvh0.PSK(u.bRxpsrh13ry)
	if err != nil {
		return err
	}
	if u.aTswvh0.Is2022() {
		block, err := aes.NewCipher(key)
		if err != nil {
			return err
		}
		u.we1MNFya = block
	}
	u.n5IKaq = key
	return nil
}
func (u *WtK9HQSy66) GetCipher() ssmethod.C6q3Upu { return u.aTswvh0 }  // 0x1820fa0
func (u *WtK9HQSy66) PSK() []byte                 { return u.n5IKaq }   // 0x1820fc0
func (u *WtK9HQSy66) GetBlock() cipher.Block      { return u.we1MNFya } // 0x1820fe0
func (u *WtK9HQSy66) NewAEAD(salt []byte) (cipher.AEAD, error) { // 0x1821000
	a, err := u.aTswvh0.NewAEAD(u.aTswvh0.SubKey(u.n5IKaq, salt))
	if err != nil {
		return nil, fmt.Errorf("user %v new aead err: %v", u.djdKrU, err)
	}
	return a, nil
}
func (u *WtK9HQSy66) NewAEADWith(s *ssmethod.MXXLkaPB, salt []byte) (cipher.AEAD, error) { // 0x18211e0
	key := s.KeyBuf(len(u.n5IKaq))
	u.aTswvh0.SubKeyTo(s, key, u.n5IKaq, salt)
	a, err := u.aTswvh0.NewAEAD(key)
	if err != nil {
		return nil, fmt.Errorf("user %v new aead err: %v", u.djdKrU, err)
	}
	return a, nil
}
func (u *WtK9HQSy66) String() string { // 0x1821440
	return fmt.Sprintf("id: %v, method: %v, password: %v", u.djdKrU, u.aTswvh0, u.bRxpsrh13ry)
}
