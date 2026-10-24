package ss2022

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"lukechampine.com/blake3"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/buf/bufiface"
	"recovered.local/soga/user"
	"time"
)

// XWewqTAtRxH2 对应 RTTI 0x36f4dc0，112 字节。
type XWewqTAtRxH2 struct {
	diohM7       *sess.Mdn09F6aMrpa
	v6K00RA5CJ2  *user.WtK9HQSy66
	gfbmydMandnH *relay.AHU2MRmjzQu
	aWzYHn7      []byte
	gc3drLRmucl  cipher.AEAD
	yr2dmViJJIP  []byte
	cWbjUh       bool
	jGMqO6       bufiface.Mf3phSOVqgf
}

// yGCzb2S 对应 0x1c6a0e0；失败路径原函数不自动 Free 缓冲区。
func yGCzb2S(conn *sess.Mdn09F6aMrpa, m *N07NaA) (*XWewqTAtRxH2, error) {
	r := &XWewqTAtRxH2{diohM7: conn, jGMqO6: pool.S4OQXWIZ(4096)}
	if err := r.init(m); err != nil {
		return nil, err
	}
	return r, nil
}

// incrementNonce 表达原各调用点内联的小端加一循环。
func incrementNonce(nonce []byte) {
	for i := range nonce {
		nonce[i]++
		if nonce[i] != 0 {
			break
		}
	}
}

// init 对应 0x1c6a1c0；错误文本见 ss2022-decoded-strings.json 原函数受限仿真证据。
func (r *XWewqTAtRxH2) init(m *N07NaA) error {
	method := m.kUidCbZ_CtGU
	keyLen, err := method.KeyLen()
	if err != nil {
		return err
	}
	r.diohM7.SetReadDeadline(time.Now().Add(time.Minute))
	n, err := r.jGMqO6.ReadMaxFrom(r.diohM7, keyLen+43)
	r.diohM7.SetReadDeadline(time.Time{})
	if err != nil {
		return err
	}
	if n < keyLen+43 {
		return fmt.Errorf("ss2022 bad header")
	}
	salt := make([]byte, keyLen)
	n, err = r.jGMqO6.Read(salt)
	if err != nil {
		return err
	}
	if n < keyLen {
		return fmt.Errorf("ss2022 impossible: n < salt length")
	}
	if m.CheckSalt(salt) {
		return fmt.Errorf("repeat salt")
	}
	keyLen, err = method.KeyLen()
	if err != nil {
		return err
	}
	material := make([]byte, keyLen*2)
	copy(material, m.gIdWgkT8p3)
	copy(material[keyLen:], salt)
	identityKey := make([]byte, keyLen)
	blake3.DeriveKey(identityKey, "shadowsocks 2022 identity subkey", material)
	block, err := aes.NewCipher(identityKey)
	if err != nil {
		return err
	}
	encrypted, err := r.jGMqO6.ReadBytes(16)
	if err != nil {
		return err
	}
	var hash [16]byte
	block.Decrypt(hash[:], encrypted)
	user, ok := m.sK7pRj.Get(hash)
	if !ok {
		return fmt.Errorf("user not exist")
	}
	aead, err := method.NewAEAD(method.SubKey(user.PSK(), salt))
	if err != nil {
		return err
	}
	nonce := make([]byte, 12)
	_, err = r.jGMqO6.Open(aead, nonce, nil)
	incrementNonce(nonce)
	if err != nil {
		return err
	}
	kind, err := r.jGMqO6.ReadUint8()
	if err != nil {
		return err
	}
	if kind != 0 {
		return fmt.Errorf("ss2022 header type wrong: %v", kind)
	}
	timestamp, err := r.jGMqO6.ReadUint64()
	if err != nil {
		return err
	}
	delta := int64(timestamp) - rtx.C5Jff5vs0()
	if delta < 0 {
		delta = -delta
	}
	if delta > 30 {
		return fmt.Errorf("ss2022 time error")
	}
	size, err := r.jGMqO6.ReadUint16()
	if err != nil {
		return err
	}
	if _, err = r.jGMqO6.ReadFullFrom(r.diohM7, int(size)+16); err != nil {
		return err
	}
	_, err = r.jGMqO6.Open(aead, nonce, nil)
	incrementNonce(nonce)
	if err != nil {
		return err
	}
	addr, err := relay.SzhW4tOzUZg(r.jGMqO6)
	if err != nil {
		return err
	}
	padding, err := r.jGMqO6.ReadUint16()
	if err != nil {
		return err
	}
	if _, err = r.jGMqO6.ReadBytes(int(padding)); err != nil {
		return err
	}
	r.diohM7.SetUser(user)
	r.v6K00RA5CJ2 = user
	r.gfbmydMandnH = addr
	r.aWzYHn7 = salt
	r.gc3drLRmucl = aead
	r.yr2dmViJJIP = nonce
	r.cWbjUh = !r.jGMqO6.IsEmpty()
	return nil
}

// j6QaqcoVDJ 对应 0x1c6b000。
func (r *XWewqTAtRxH2) j6QaqcoVDJ() (int, error) {
	if _, err := r.jGMqO6.ReadFullFrom(r.diohM7, r.gc3drLRmucl.Overhead()+2); err != nil {
		return 0, err
	}
	_, err := r.jGMqO6.Open(r.gc3drLRmucl, r.yr2dmViJJIP, nil)
	incrementNonce(r.yr2dmViJJIP)
	if err != nil {
		return 0, err
	}
	n, err := r.jGMqO6.ReadUint16()
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// pFvi5Xw 对应 0x1c6b120，原程序只有 size < len(p) 才采用直接解密分支。
func (r *XWewqTAtRxH2) pFvi5Xw(p []byte) (int, error) {
	n, err := r.j6QaqcoVDJ()
	if err != nil {
		return 0, err
	}
	if _, err = r.jGMqO6.ReadFullFrom(r.diohM7, n+r.gc3drLRmucl.Overhead()); err != nil {
		return 0, err
	}
	if n < len(p) {
		_, err = r.gc3drLRmucl.Open(p[:0], r.yr2dmViJJIP, r.jGMqO6.Bytes(), nil)
		incrementNonce(r.yr2dmViJJIP)
		if err != nil {
			return 0, err
		}
		r.jGMqO6.Reset()
		return n, nil
	}
	_, err = r.jGMqO6.Open(r.gc3drLRmucl, r.yr2dmViJJIP, nil)
	incrementNonce(r.yr2dmViJJIP)
	if err != nil {
		return 0, err
	}
	return r.jGMqO6.Read(p)
}
func (r *XWewqTAtRxH2) Read(p []byte) (int, error) {
	if r.jGMqO6.Len() > 0 {
		return r.jGMqO6.Read(p)
	}
	return r.pFvi5Xw(p)
}                                                          // 0x1c6b3a0
func (r *XWewqTAtRxH2) GetUser() user.PgaNd3014U       { return r.v6K00RA5CJ2 }  // 0x1c6b460
func (r *XWewqTAtRxH2) GetDestAddr() *relay.AHU2MRmjzQu { return r.gfbmydMandnH } // 0x1c6b480
func (r *XWewqTAtRxH2) NetworkType() sess.ZES6F5    { return 0 }              // 0x1c6b4a0
// cKgVQx 对应 0x1c6b4c0。
func (r *XWewqTAtRxH2) cKgVQx() error {
	r.jGMqO6.Reset()
	n, err := r.j6QaqcoVDJ()
	if err != nil {
		return err
	}
	r.jGMqO6.Reset()
	if _, err = r.jGMqO6.ReadFullFrom(r.diohM7, n+r.gc3drLRmucl.Overhead()); err != nil {
		return err
	}
	_, err = r.jGMqO6.Open(r.gc3drLRmucl, r.yr2dmViJJIP, nil)
	incrementNonce(r.yr2dmViJJIP)
	return err
}
func (r *XWewqTAtRxH2) ReadBuffer() (bufiface.Mf3phSOVqgf, error) { // 0x1c6b600
	if r.cWbjUh {
		r.cWbjUh = false
		return r.jGMqO6, nil
	}
	if err := r.cKgVQx(); err != nil {
		return nil, err
	}
	return r.jGMqO6, nil
}
func (r *XWewqTAtRxH2) Free() { r.jGMqO6.Free() } // 0x1c6b680
