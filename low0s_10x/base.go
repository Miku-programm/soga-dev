// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package low0s_10x

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	mathRand "math/rand"
	"sync"
	"time"

	"recovered.local/soga/m3Cw8P3_Le"
	"recovered.local/soga/mzi1d3gDO7u"
	"recovered.local/soga/x6HWbunbF"
)

type PZfVxx4UGRb struct {
	Z2UHwhi2h4  []byte
	UqLugVVWCKJ int
	BOJR3xVshAb string
}

type bhRsfpsWw struct {
	oYNnYN0vN []byte
	qC81k5Mk  [4]byte
}

type t1vL_y55t struct {
	kh9UY106p3Y6 [4]byte
	xOB1UU       uint32
	lDLI1IYh92   sync.Mutex
}

// (*t1vL_y55t).uU1wpTCSA: 0x1b3ed60.
func (a *t1vL_y55t) uU1wpTCSA() *t1vL_y55t {
	r := &t1vL_y55t{}
	a.lDLI1IYh92.Lock()
	defer a.lDLI1IYh92.Unlock()
	if a.xOB1UU > 0xff000000 || a.xOB1UU == 0 {
		rand.Read(a.kh9UY106p3Y6[:])
		a.xOB1UU = mathRand.Uint32() & 0xffffff
	}
	a.xOB1UU++
	copy(r.kh9UY106p3Y6[:], a.kh9UY106p3Y6[:])
	r.xOB1UU = a.xOB1UU
	return r
}

// (*t1vL_y55t).mHxvPk: 0x1b3eea0.
func (a *t1vL_y55t) mHxvPk(buf *bytes.Buffer) {
	binary.Write(buf, binary.LittleEndian, uint32(time.Now().Unix()))
	buf.Write(a.kh9UY106p3Y6[:])
	binary.Write(buf, binary.LittleEndian, a.xOB1UU)
}

// (*t1vL_y55t).vKsdajuSitet: 0x1b3efa0.
func (a *t1vL_y55t) vKsdajuSitet(b *bytes.Buffer, oYNnYN0vN []byte, paddings [2]int, salt string) error {
	encrypt := mzi1d3gDO7u.SharedPoolAt5096978.Get(16)
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(encrypt)
	binary.LittleEndian.PutUint32(encrypt, uint32(time.Now().Unix()))
	copy(encrypt[4:], a.kh9UY106p3Y6[:])
	binary.LittleEndian.PutUint32(encrypt[8:], a.xOB1UU)
	binary.LittleEndian.PutUint16(encrypt[12:], uint16(paddings[0]))
	binary.LittleEndian.PutUint16(encrypt[14:], uint16(paddings[1]))

	cipherKey := m3Cw8P3_Le.ArNrGS_EV(base64.StdEncoding.EncodeToString(oYNnYN0vN)+salt, 16)
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		x6HWbunbF.Dl2Gj6("New cipher error: %s", err.Error())
		return err
	}
	iv := bytes.Repeat([]byte{0}, 16)
	cbcCipher := cipher.NewCBCEncrypter(block, iv)

	cbcCipher.CryptBlocks(encrypt, encrypt)

	b.Write(encrypt)
	return nil
}
