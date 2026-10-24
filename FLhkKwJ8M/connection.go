package FLhkKwJ8M

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/chacha20poly1305"
	"io"
	"net"
)

type IxhMZ66CqoH1 struct {
	GV3P5rY     bool
	OyEjzi1v93  uint8
	TIBiNxg     []byte
	GHP6ivev2Db uint
}

// DrpXao retains the original embedded alias and net.Conn method promotion.
type DrpXao = net.Conn
type HRzHVG1yP5O struct {
	DrpXao
	xGnl3X       io.Reader
	_mBMg3       io.Writer
	mNrCvc       *IxhMZ66CqoH1
	fqAsdD3hjU   *OMOYhoLsaabd
	c646iDyBoW   []byte
	_7ezur1gq_ZX []byte
	f9_fTlq      []byte
	z1TtPlbnfnl  []byte
	kONmi_w_yL3D uint8
	iGH6ra       uint8
	kljQIi6N     uint8
	tu04N8lBj6I  bool
	ppPBszhDxm   bool
}

// Write: 0x19aeae0.
func (c *HRzHVG1yP5O) Write(p []byte) (int, error) { return c._mBMg3.Write(p) }

// Read: 0x19aeb40. A failed response check is retried on the next Read.
func (c *HRzHVG1yP5O) Read(p []byte) (int, error) {
	if !c.ppPBszhDxm {
		if err := c.dCFgAbCz(); err != nil {
			return 0, err
		}
		c.ppPBszhDxm = true
	}
	return c.xGnl3X.Read(p)
}

// wv3qc3: 0x19b0200. Unknown security codes deliberately retain nil data IO;
// the original still sends their request header before returning the connection.
func wv3qc3(conn net.Conn, id *OMOYhoLsaabd, dst *IxhMZ66CqoH1, security byte, aeadHeader bool) (*HRzHVG1yP5O, error) {
	var seed [33]byte
	rand.Read(seed[:])
	iv := make([]byte, 16)
	key := make([]byte, 16)
	copy(iv, seed[:16])
	copy(key, seed[16:32])
	var responseKey, responseIV []byte
	if aeadHeader {
		k := sha256.Sum256(key)
		v := sha256.Sum256(iv)
		responseKey = k[:16]
		responseIV = v[:16]
	} else {
		k := md5.Sum(key)
		v := md5.Sum(iv)
		responseKey = k[:]
		responseIV = v[:]
	}
	var reader io.Reader
	var writer io.Writer
	options := byte(1)
	switch security {
	case 3, 4:
		var enc, dec cipher.AEAD
		if security == 3 {
			b, _ := aes.NewCipher(key)
			enc, _ = cipher.NewGCM(b)
			b, _ = aes.NewCipher(responseKey)
			dec, _ = cipher.NewGCM(b)
		} else {
			k := md5.Sum(key)
			k2 := md5.Sum(k[:])
			var ck [32]byte
			copy(ck[:16], k[:])
			copy(ck[16:], k2[:])
			enc, _ = chacha20poly1305.New(ck[:])
			k = md5.Sum(responseKey)
			k2 = md5.Sum(k[:])
			copy(ck[:16], k[:])
			copy(ck[16:], k2[:])
			dec, _ = chacha20poly1305.New(ck[:])
		}
		writer = &sBT2tDxX3LZ{Writer: conn, AEAD: enc, xNF1Gb47: iv}
		reader = &se83rHogAI7{Reader: conn, AEAD: dec, frnZkWRc: responseIV, jYxJbABejRtF: make([]byte, 2)}
	case 5:
		writer = &tojW0z34{conn}
		reader = &vV27Q3k{Reader: conn, gOoPyboUld: make([]byte, 2)}
	case 6:
		security = 5
		if dst.GV3P5rY {
			writer = &tojW0z34{conn}
			reader = &vV27Q3k{Reader: conn, gOoPyboUld: make([]byte, 2)}
		} else {
			reader = conn
			writer = conn
			options = 0
		}
	}
	c := &HRzHVG1yP5O{DrpXao: conn, xGnl3X: reader, _mBMg3: writer, mNrCvc: dst, fqAsdD3hjU: id, c646iDyBoW: iv, _7ezur1gq_ZX: key, f9_fTlq: responseIV, z1TtPlbnfnl: responseKey, kONmi_w_yL3D: seed[32], iGH6ra: security, kljQIi6N: options, tu04N8lBj6I: aeadHeader}
	if err := c.nhGe6Mb_IeZ(); err != nil {
		return nil, err
	}
	return c, nil
}
