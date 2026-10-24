package g0MlUOhvO

import (
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/sha3"
	"net"
	"recovered.local/soga/Y5PIQhBeRha"
	"time"
)

// dlEStwP4xJ 对应0x11d44e0。盐采用一次非空Read而非ReadFull，
// 设置60秒读取期限后清空，原反重放缓存命中时必须拒绝连接。
func (c *Cb7vj3KQA8) dlEStwP4xJ() error {
	salt := make([]byte, c.hp3cHCehD2pe.NonceLen())
	c.DrpXao.SetReadDeadline(time.Now().Add(60 * time.Second))
	var n int
	var err error
	for {
		n, err = c.DrpXao.Read(salt)
		if err != nil || n > 0 {
			break
		}
	}
	c.DrpXao.SetReadDeadline(time.Time{})
	if err != nil {
		return err
	}
	if n != len(salt) {
		return errors.New("read salt len wrong")
	}
	if !Y5PIQhBeRha.UJhJgiLn(salt) {
		return errors.New("repeat salt")
	}
	key, aead, err := c.hp3cHCehD2pe.NewAEAD(salt)
	if err != nil {
		return err
	}
	c.e8aahi_R3S = aead
	c.rOfYjOzCdA1L = salt
	shake := sha3.NewShake128()
	shake.Write(salt)
	shake.Write(key)
	c.w1vg8Dd = &ErrfcUUB{pzOWW_: shake}
	return nil
}

// hm8Y5JRNqwp 对应0x11d4a60，返回同一nonce底层切片作为首帧前缀。
func (c *Cb7vj3KQA8) hm8Y5JRNqwp() ([]byte, error) {
	salt := make([]byte, c.hp3cHCehD2pe.NonceLen())
	rand.Read(salt)
	key, aead, err := c.hp3cHCehD2pe.NewAEAD(salt)
	if err != nil {
		return nil, err
	}
	c.bUWvbqWzlH = aead
	c.cdwwdRLQG = salt
	shake := sha3.NewShake128()
	shake.Write(salt)
	shake.Write(key)
	c.r6A494 = &ErrfcUUB{pzOWW_: shake}
	return salt, nil
}

// Read 对应0x11d5dc0。
func (c *Cb7vj3KQA8) Read(p []byte) (int, error) {
	if c.e8aahi_R3S == nil {
		if err := c.dlEStwP4xJ(); err != nil {
			return 0, err
		}
	}
	return c.e21qHx(p)
}

// Write 对应0x11d5e60。
func (c *Cb7vj3KQA8) Write(p []byte) (int, error) {
	var prefix []byte
	if c.bUWvbqWzlH == nil {
		var err error
		prefix, err = c.hm8Y5JRNqwp()
		if err != nil {
			return 0, err
		}
	}
	return c.oH5wsFE(prefix, p)
}

// Accept 对应0x11d7180，接受时仅装配包装器，认证延迟到首次读取。
func (l *T4k2R7cJ) Accept() (net.Conn, error) {
	raw, err := l.O6B2qZ0p0.Accept()
	if err != nil {
		return nil, err
	}
	return &Cb7vj3KQA8{DrpXao: raw, hp3cHCehD2pe: l.tS5gsM72s}, nil
}
