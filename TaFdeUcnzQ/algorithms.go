package TaFdeUcnzQ

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"encoding/binary"
	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/cast5"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/salsa20/salsa"
	P "recovered.local/soga/D_zCag1qTV"
	R "recovered.local/soga/J2SLI_g3TG"
	I "recovered.local/soga/JPLmG42"
	X "recovered.local/soga/_KBmXZd"
	C "recovered.local/soga/ubT9IZfVz0MQ"
)

// iuxGEQ restores 0x1ca13a0. The original accepts a block/error pair
// followed by the IV and direction, and propagates the block error first.
func iuxGEQ(block cipher.Block, err error, iv []byte, direction IUgK3nW) (cipher.Stream, error) {
	if err != nil {
		return nil, err
	}
	if direction == 1 {
		return cipher.NewCFBEncrypter(block, iv), nil
	}
	return cipher.NewCFBDecrypter(block, iv), nil
}
func rzaLikkN6bw_(key, iv []byte, _ IUgK3nW) (cipher.Stream, error) {
	b, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	return cipher.NewCTR(b, iv), nil
} // 0x1ca1220
func otGPQ5SQ(key, iv []byte, _ IUgK3nW) (cipher.Stream, error) {
	b, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	return cipher.NewOFB(b, iv), nil
}                                                               // 0x1ca12e0
func wl3p_IavOEJ(_, _ []byte, _ IUgK3nW) (cipher.Stream, error) { return &zclWYvDGjqa2{}, nil } // 0x1ca1480
func rMHlyA4axa(key, iv []byte, d IUgK3nW) (cipher.Stream, error) {
	b, e := aes.NewCipher(key)
	return iuxGEQ(b, e, iv, d)
} // 0x1ca14a0
func vw1gL73QDne(key, iv []byte, d IUgK3nW) (cipher.Stream, error) {
	b, e := des.NewCipher(key)
	return iuxGEQ(b, e, iv, d)
} // 0x1ca1580
func vhDmXBbbD(key, iv []byte, d IUgK3nW) (cipher.Stream, error) {
	b, e := blowfish.NewCipher(key)
	return iuxGEQ(b, e, iv, d)
} // 0x1ca1660
func iFGyatq(key, iv []byte, d IUgK3nW) (cipher.Stream, error) {
	b, e := cast5.NewCipher(key)
	return iuxGEQ(b, e, iv, d)
} // 0x1ca1760
// Returning the concrete pointer preserves the original typed-nil interface
// when these constructors fail (0x1ca18e0, 0x1ca1960, 0x1ca1b20, 0x1ca1bc0).
func g57NcXia7(key, _ []byte, _ IUgK3nW) (cipher.Stream, error) { return rc4.NewCipher(key) }
func qi9_aRSWW8(key, iv []byte, _ IUgK3nW) (cipher.Stream, error) {
	h := md5.New()
	h.Write(key)
	h.Write(iv)
	return rc4.NewCipher(h.Sum(nil))
}
func lPcdRtmHNa(key, iv []byte, _ IUgK3nW) (cipher.Stream, error) { return X.JpuCElV(key, iv) }
func dzRpwKaY(key, iv []byte, _ IUgK3nW) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(key, iv)
}
func uHwob9RZa(key, iv []byte, d IUgK3nW) (cipher.Stream, error) {
	b, e := C.RMycR7V(key)
	return iuxGEQ(b, e, iv, d)
} // 0x1ca1ca0
func gBu8RF4S2Et4(key, iv []byte, d IUgK3nW) (cipher.Stream, error) {
	b, e := I.L3EF5OwH(key)
	return iuxGEQ(b, e, iv, d)
} // 0x1ca1d80
func jYMM7TY(key, iv []byte, d IUgK3nW) (cipher.Stream, error) {
	b, e := R.JAdaDYP7aSa(key, 16)
	return iuxGEQ(b, e, iv, d)
}                                             // 0x1ca1e60: effective key size is 16 bits.
func mk62qv0(key []byte) (cipher.AEAD, error) { return chacha20poly1305.New(key) } // 0x1ca1f60
func jPLNP81(key []byte) (cipher.AEAD, error) {
	b, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	return cipher.NewGCM(b)
}                                                  // 0x1ca2040
func (*zclWYvDGjqa2) XORKeyStream(dst, src []byte) { copy(dst, src) } // 0x1ca2d20
func iuskSmap(key, iv []byte, _ IUgK3nW) (cipher.Stream, error) { // 0x1ca2dc0
	s := &jUB7SuavZpra{}
	copy(s.pD_OAaLAHa2H[:], iv[:8])
	copy(s.u9KzUjD[:], key[:32])
	return s, nil
}
func (s *jUB7SuavZpra) XORKeyStream(dst, src []byte) { // 0x1ca2ec0
	offset := s.cO2TNZUn7xB % 64
	size := len(src) + offset
	var work []byte
	if cap(dst) < size {
		if len(s.eWLprPSs16C) < size {
			if s.eWLprPSs16C != nil {
				P.PoolAt50961D8.Put(s.eWLprPSs16C)
			}
			s.eWLprPSs16C = P.PoolAt50961D8.Get(size)
		}
		work = s.eWLprPSs16C[:size]
	} else {
		work = dst[:size]
	}
	var counter [16]byte
	copy(counter[:8], s.pD_OAaLAHa2H[:])
	binary.LittleEndian.PutUint64(counter[8:], uint64(s.cO2TNZUn7xB/64))
	copy(work[offset:], src)
	salsa.XORKeyStream(work, work, &counter, &s.u9KzUjD)
	copy(dst, work[offset:])
	s.cO2TNZUn7xB += len(src)
}
