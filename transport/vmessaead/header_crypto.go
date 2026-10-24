package vmessaead

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"hash/crc32"
	"recovered.local/soga/util/bytes"
	"time"
)

// iAK7_ZD8: 0x19b0020.
func iAK7_ZD8(t time.Time) []byte {
	var timestamp [8]byte
	binary.BigEndian.PutUint64(timestamp[:], uint64(t.Unix()))
	h := md5.New()
	for range 4 {
		h.Write(timestamp[:])
	}
	return h.Sum(nil)
}

// otLnwGxNYh7: 0x19b4a60. ReadFull failure is deliberately ignored as in target.
func otLnwGxNYh7(key []byte, timestamp uint64) (result [16]byte) {
	b := bytes.BaDekw7qe{}
	b.Grow(8)
	binary.BigEndian.PutUint64(b, timestamp)
	b.ReadFull(rand.Reader, 4)
	checksum := crc32.ChecksumIEEE(b)
	pos := b.Grow(4)
	binary.BigEndian.PutUint32(b[pos:pos+4], checksum)
	block, _ := aes.NewCipher(xXVvCU5cjuB(key, "AES Auth ID Encryption")[:16])
	block.Encrypt(result[:], b)
	return
}

// s9ADKgL: 0x19b4cc0. Wire order: auth ID, encrypted length, nonce, encrypted payload.
func s9ADKgL(key [16]byte, payload []byte, t time.Time) []byte {
	auth := otLnwGxNYh7(key[:], uint64(t.Unix()))
	var nonce [8]byte
	rand.Read(nonce[:])
	var length [2]byte
	binary.BigEndian.PutUint16(length[:], uint16(len(payload)))
	lk := xXVvCU5cjuB(key[:], "VMess Header AEAD Key_Length", string(auth[:]), string(nonce[:]))
	ln := xXVvCU5cjuB(key[:], "VMess Header AEAD Nonce_Length", string(auth[:]), string(nonce[:]))
	block, _ := aes.NewCipher(lk[:16])
	aead, _ := cipher.NewGCM(block)
	encryptedLength := aead.Seal(nil, ln[:12], length[:], auth[:])
	pk := xXVvCU5cjuB(key[:], "VMess Header AEAD Key", string(auth[:]), string(nonce[:]))
	pn := xXVvCU5cjuB(key[:], "VMess Header AEAD Nonce", string(auth[:]), string(nonce[:]))
	block, _ = aes.NewCipher(pk[:16])
	aead, _ = cipher.NewGCM(block)
	encryptedPayload := aead.Seal(nil, pn[:12], payload, auth[:])
	b := bytes.BaDekw7qe{}
	for _, part := range [][]byte{auth[:], encryptedLength, nonce[:], encryptedPayload} {
		pos := b.Grow(len(part))
		copy(b[pos:pos+len(part)], part)
	}
	return b
}
