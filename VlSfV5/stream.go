// Package VlSfV5 reconstructs the target's native legacy Shadowsocks inbound.
package VlSfV5

import (
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"recovered.local/soga/A818JHb"
	"recovered.local/soga/CpHCdepV"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/NtxdAE"
	"recovered.local/soga/Ualx5UM_D6I"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/w2YaTd9_"
)

// Original init 0x1c5e1c0; strings decoded from the target, not inferred.
var errorAt50983F0 = errors.New("shadowsocks data length too long")
var errorAt5098400 = errors.New("salt repeat")
var errorAt5098410 = errors.New("not found user")

type DrpXao = net.Conn

// BgeR1vhZZG retains the original 200-byte RTTI 0x3724e30 layout.
type BgeR1vhZZG struct {
	DrpXao
	sY6Ak5rJDEPi w2YaTd9_.C6q3Upu
	fXCOXjGn5K   *J1P4KNvfI_J.EafxyUZD
	upTmVMTVIBAF *NtxdAE.M1Qn9BA
	kRHSVpkV     *nzATvajt.ZpNFqUVZp
	bu3p3B57gQjz *nzATvajt.WtK9HQSy66
	pDnKrW2r     *G6bHD5JA.AHU2MRmjzQu
	sFp8vo9jQ    cipher.AEAD
	xhW4VaUrBa4r []byte
	uYsj2V       cipher.AEAD
	jyca9mgLi    []byte
	kDc_uTWoO    int
	z9ea37s3b    bool
	f4ApgxW4hCCA Ualx5UM_D6I.Mf3phSOVqgf
	a24Pchp      Ualx5UM_D6I.Mf3phSOVqgf
}

// VLnoIxcGjx corresponds to 0x1c5e3e0.
func VLnoIxcGjx(conn net.Conn, method w2YaTd9_.C6q3Upu, candidates *J1P4KNvfI_J.EafxyUZD, invalid *NtxdAE.M1Qn9BA, users *nzATvajt.ZpNFqUVZp) (*BgeR1vhZZG, error) {
	c := &BgeR1vhZZG{sY6Ak5rJDEPi: method, fXCOXjGn5K: candidates, upTmVMTVIBAF: invalid, kRHSVpkV: users, kDc_uTWoO: -1, f4ApgxW4hCCA: D_zCag1qTV.S4OQXWIZ(4096), a24Pchp: D_zCag1qTV.S4OQXWIZ(16384)}
	if err := c.init(CpHCdepV.MlZo77Me(conn)); err != nil {
		HxKViEp2a.Li3c4CJ1_D(c)
		return nil, err
	}
	return c, nil
}
func (c *BgeR1vhZZG) init(conn net.Conn) error { //0x1c5e5e0
	if err := c.aHK6b5W(conn); err != nil {
		return err
	}
	counted := &J1P4KNvfI_J.Mdn09F6aMrpa{DrpXao: conn}
	counted.SetUser(c.bu3p3B57gQjz)
	c.DrpXao = counted
	dest, err := G6bHD5JA.SzhW4tOzUZg(c)
	if err != nil {
		return err
	}
	c.z9ea37s3b = !c.f4ApgxW4hCCA.IsEmpty()
	c.pDnKrW2r = dest
	return nil
}
func incrementNonce(n []byte) {
	for i := range n {
		n[i]++
		if n[i] != 0 {
			break
		}
	}
}

// aHK6b5W 0x1c5e760 and closures 0x1c63c60/1c63b40/1c63980.
func (c *BgeR1vhZZG) aHK6b5W(conn net.Conn) error {
	keyLen, err := c.sY6Ak5rJDEPi.KeyLen()
	if err != nil {
		return err
	}
	initial := make([]byte, keyLen+18)
	if _, err = io.ReadFull(conn, initial); err != nil {
		return err
	}
	salt := initial[:keyLen]
	if !A818JHb.HFlnXWty(salt) {
		return errorAt5098400
	}
	remote, err := G6bHD5JA.JY6gu_Rqfz(conn.RemoteAddr())
	if err != nil {
		return err
	}
	ip := remote.IP.String()
	nonce := make([]byte, 12)
	var length [2]byte
	state := w2YaTd9_.SaYhx5xog9()
	check := func(u *nzATvajt.WtK9HQSy66) (cipher.AEAD, bool) {
		a, e := u.NewAEADWith(state, salt)
		if e != nil {
			muc279Mpa6S4.TRTbrquD(4, e)
			return nil, false
		}
		_, e = a.Open(length[:0], nonce, initial[keyLen:], nil)
		if e != nil {
			return nil, false
		}
		return a, true
	}
	seen := make(map[int64]struct{})
	var selected *nzATvajt.WtK9HQSy66
	var aead cipher.AEAD
	c.fXCOXjGn5K.RangeUser(ip, func(base nzATvajt.PgaNd3014U) bool {
		u := base.(*nzATvajt.WtK9HQSy66)
		seen[u.GetId()] = struct{}{}
		if a, ok := check(u); ok {
			aead = a
			selected = u
			return false
		}
		return true
	})
	if selected == nil {
		count := 0
		c.kRHSVpkV.Range(func(_ int64, base nzATvajt.PgaNd3014U) bool {
			u := base.(*nzATvajt.WtK9HQSy66)
			if _, ok := seen[u.GetId()]; ok {
				return true
			}
			count++
			if count%50 == 0 && c.upTmVMTVIBAF.IsForbidden(ip) {
				return false
			}
			if a, ok := check(u); ok {
				aead = a
				selected = u
				return false
			}
			return true
		})
	}
	if selected == nil {
		return errorAt5098410
	}
	c.fXCOXjGn5K.AddUser(ip, selected, 0)
	c.bu3p3B57gQjz = selected
	c.sFp8vo9jQ = aead
	c.xhW4VaUrBa4r = nonce
	incrementNonce(nonce)
	c.kDc_uTWoO = int(binary.BigEndian.Uint16(length[:]))
	return nil
}
func (c *BgeR1vhZZG) gkew8UZXK9V() (int, error) { //0x1c5ed80
	if c.kDc_uTWoO >= 0 {
		n := c.kDc_uTWoO
		c.kDc_uTWoO = -1
		return n, nil
	}
	b := c.f4ApgxW4hCCA
	b.Reset()
	if _, err := b.ReadFullFrom(c.DrpXao, 2+c.sFp8vo9jQ.Overhead()); err != nil {
		return 0, err
	}
	_, err := b.Open(c.sFp8vo9jQ, c.xhW4VaUrBa4r, nil)
	incrementNonce(c.xhW4VaUrBa4r)
	if err != nil {
		return 0, err
	}
	n, err := b.ReadUint16()
	return int(n), err
}
func (c *BgeR1vhZZG) ob_9B3i6zZ() error { //0x1c5f500
	n, err := c.gkew8UZXK9V()
	if err != nil {
		return err
	}
	if n > 16383 {
		return errorAt50983F0
	}
	b := c.f4ApgxW4hCCA
	b.Reset()
	if _, err = b.ReadFullFrom(c.DrpXao, n+c.sFp8vo9jQ.Overhead()); err != nil {
		return err
	}
	_, err = b.Open(c.sFp8vo9jQ, c.xhW4VaUrBa4r, nil)
	incrementNonce(c.xhW4VaUrBa4r)
	return err
}
func (c *BgeR1vhZZG) Read(p []byte) (int, error) { //0x1c5ef60
	if c.f4ApgxW4hCCA.IsEmpty() {
		if err := c.ob_9B3i6zZ(); err != nil {
			return 0, err
		}
	}
	return c.f4ApgxW4hCCA.Read(p)
}
func (c *BgeR1vhZZG) dyrYP5aakv(salt []byte) error { //0x1c5f060
	a, err := c.bu3p3B57gQjz.NewAEAD(salt)
	if err != nil {
		return err
	}
	c.uYsj2V = a
	c.jyca9mgLi = make([]byte, a.NonceSize())
	return nil
}
func (c *BgeR1vhZZG) Write(p []byte) (int, error) { //0x1c5f160
	if len(p) > 16383 {
		half := len(p) >> 1
		n, err := c.Write(p[:half])
		if err != nil {
			return n, err
		}
		m, err := c.Write(p[half:])
		return n + m, err
	}
	b := c.a24Pchp
	b.Reset()
	if c.uYsj2V == nil {
		n, err := c.sY6Ak5rJDEPi.KeyLen()
		if err != nil {
			return 0, err
		}
		b.WriteRandom(n)
		if err = c.dyrYP5aakv(b.Bytes()); err != nil {
			return 0, err
		}
	}
	var length [2]byte
	binary.BigEndian.PutUint16(length[:], uint16(len(p)))
	b.SealWrite(c.uYsj2V, c.jyca9mgLi, length[:], nil)
	incrementNonce(c.jyca9mgLi)
	b.SealWrite(c.uYsj2V, c.jyca9mgLi, p, nil)
	incrementNonce(c.jyca9mgLi)
	if _, err := c.DrpXao.Write(b.Bytes()); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (c *BgeR1vhZZG) GetUser() nzATvajt.PgaNd3014U       { return c.bu3p3B57gQjz } //0x1c5f4a0
func (c *BgeR1vhZZG) GetDestAddr() *G6bHD5JA.AHU2MRmjzQu { return c.pDnKrW2r }     //0x1c5f4c0
func (c *BgeR1vhZZG) NetworkType() J1P4KNvfI_J.ZES6F5    { return 0 }              //0x1c5f4e0
func (c *BgeR1vhZZG) ReadBuffer() (Ualx5UM_D6I.Mf3phSOVqgf, error) { //0x1c5f700
	if c.z9ea37s3b {
		c.z9ea37s3b = false
		return c.f4ApgxW4hCCA, nil
	}
	if err := c.ob_9B3i6zZ(); err != nil {
		return nil, err
	}
	return c.f4ApgxW4hCCA, nil
}
func (c *BgeR1vhZZG) WriteBuffer(b Ualx5UM_D6I.Mf3phSOVqgf) (int, error) { return c.Write(b.Bytes()) }                                               //0x1c5f7a0
func (c *BgeR1vhZZG) Free()                                              { c.f4ApgxW4hCCA.Free(); c.a24Pchp.Free(); HxKViEp2a.Li3c4CJ1_D(c.DrpXao) } //0x1c5f820
