package FLhkKwJ8M

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

// pd89U7 is the original recursive HMAC factory, 32 bytes in the target.
// It differs from the separate JafvCs5 factory, which shares hash instances.
type pd89U7 struct {
	qe_y3FIf    *pd89U7
	zAtsk7OMEP1 []byte
}

// Create: 0x19b49e0.
func (p *pd89U7) Create() hash.Hash {
	if p.qe_y3FIf == nil {
		return hmac.New(sha256.New, p.zAtsk7OMEP1)
	}
	return hmac.New(p.qe_y3FIf.Create, p.zAtsk7OMEP1)
}

// xXVvCU5cjuB: 0x19b47c0. Labels are strings, including binary label bytes.
func xXVvCU5cjuB(key []byte, path ...string) []byte {
	p := &pd89U7{zAtsk7OMEP1: []byte("VMess AEAD KDF")}
	for _, label := range path {
		p = &pd89U7{qe_y3FIf: p, zAtsk7OMEP1: []byte(label)}
	}
	h := p.Create()
	h.Write(key)
	return h.Sum(nil)
}
