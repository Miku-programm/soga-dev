package vmess

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"fmt"
	"golang.org/x/crypto/sha3"
	"io"
	"log/slog"
	"net"
	B "recovered.local/soga/buf/pool"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/util/ssrhelper"
	W "recovered.local/soga/buf/bufiface"
	Log "recovered.local/soga/util/logx"
)

// DVEkfear 为原 VMess 帧连接，RTTI 192 字节；原 net.Conn 的混淆嵌入名为 DrpXao。
type DVEkfear struct {
	net.Conn
	aJeMzWMqaw  *Wpqfspy6zZ_B
	bL4KxcUfi   *URTFuf2wVWf
	onlB1SoVyd  *WR4lMNfab
	s6odlStcw3  sha3.ShakeHash
	vFAueMeKs9  sha3.ShakeHash
	xNm7NkjP    G1Wopa
	cC5h7oc7    BwCT6jfasz6
	ieASegJmYKY [2]byte
	eDSy4QE     [2]byte
	qvY6ucfXLA  bool
	bxlT0MtgU_Y bytes.Reader
	osKtcdtrQlz []byte
	dvRoTXVYrAl W.Mf3phSOVqgf
}

func W2tVTGQ(conn net.Conn, auth *WR4lMNfab, request *Wpqfspy6zZ_B) (*DVEkfear, error) { // 0x1d1c620
	if !request.OptS() && request.mTRhOfKViN != 5 {
		return nil, errors.New("opts == false is deprecated")
	}
	response := ELa2jh3VBB(request, auth)
	var readMask, writeMask sha3.ShakeHash
	if request.OptM() {
		readMask = sha3.NewShake128()
		readMask.Write(request.cEcPFEa7)
		writeMask = sha3.NewShake128()
		writeMask.Write(response.maCTf0DLoP)
	}
	dec, err := DtSw_eAbIQ(request.mTRhOfKViN, request.ezCHRvHCePBC, request.cEcPFEa7)
	if err != nil {
		return nil, err
	}
	enc, err := RsWvpjES(request.mTRhOfKViN, response.iqQ8HJ6OMB, response.maCTf0DLoP)
	if err != nil {
		return nil, err
	}
	return &DVEkfear{Conn: conn, aJeMzWMqaw: request, bL4KxcUfi: response, onlB1SoVyd: auth, s6odlStcw3: readMask, vFAueMeKs9: writeMask, xNm7NkjP: dec, cC5h7oc7: enc}, nil
}

func (c *DVEkfear) x4DVqa7CuTw() (int, int, error) { // 0x1d1cbc0
	if _, err := io.ReadFull(c.Conn, c.ieASegJmYKY[:]); err != nil {
		return 0, 0, err
	}
	size := binary.BigEndian.Uint16(c.ieASegJmYKY[:])
	padding := 0
	if c.aJeMzWMqaw.OptM() {
		if c.aJeMzWMqaw.OptPadding() {
			if _, err := c.s6odlStcw3.Read(c.ieASegJmYKY[:]); err != nil {
				return 0, 0, err
			}
			padding = int(binary.BigEndian.Uint16(c.ieASegJmYKY[:]) & 63)
		}
		if _, err := c.s6odlStcw3.Read(c.ieASegJmYKY[:]); err != nil {
			return 0, 0, err
		}
		size ^= binary.BigEndian.Uint16(c.ieASegJmYKY[:])
	}
	return int(size), padding, nil
}
func (c *DVEkfear) dPv5C6sn(size uint16) (uint16, int, error) { // 0x1d1dda0
	padding := uint16(0)
	if c.aJeMzWMqaw.OptM() {
		if c.aJeMzWMqaw.OptPadding() {
			if _, err := c.vFAueMeKs9.Read(c.eDSy4QE[:]); err != nil {
				return 0, 0, err
			}
			padding = binary.BigEndian.Uint16(c.eDSy4QE[:]) & 63
		}
		if _, err := c.vFAueMeKs9.Read(c.eDSy4QE[:]); err != nil {
			return 0, 0, err
		}
		size = (size + padding) ^ binary.BigEndian.Uint16(c.eDSy4QE[:])
	}
	binary.BigEndian.PutUint16(c.eDSy4QE[:], size)
	return size, int(padding), nil
}

func (c *DVEkfear) Read(p []byte) (int, error) { // 0x1d1cd80
	if c.bxlT0MtgU_Y.Len() > 0 {
		n, _ := c.bxlT0MtgU_Y.Read(p)
		if c.bxlT0MtgU_Y.Len() == 0 {
			B.PoolAt50961D8.Put(c.osKtcdtrQlz)
			c.osKtcdtrQlz = nil
		}
		return n, nil
	}
	if !c.aJeMzWMqaw.OptS() {
		return c.Conn.Read(p)
	}
	size, padding, err := c.x4DVqa7CuTw()
	if err != nil {
		return 0, err
	}
	if size > 32768 {
		return 0, fmt.Errorf("vmess packet size too big: %v", size)
	}
	var data []byte
	release := true
	if size > len(p) {
		data = B.PoolAt50961D8.Get(size)
		defer func() {
			if release {
				B.PoolAt50961D8.Put(data)
			}
		}()
	} else {
		data = p[:size]
	}
	if _, err = io.ReadFull(c.Conn, data); err != nil {
		return 0, err
	}
	// 原 Read 在读帧后检查 padding；ReadBuffer 在读帧前检查。
	if padding > size {
		return 0, fmt.Errorf("size(%v) < padding(%v)", size, padding)
	}
	data = data[:size-padding]
	n, err := c.xNm7NkjP.Decrypt(data, data)
	if err != nil {
		Log.TRTbrquD(slog.LevelDebug, "ignored vmess decrypt err:", err)
		return 0, nil
	}
	if n == 0 {
		return 0, io.EOF
	}
	copied := copy(p, data[:n])
	if n > copied {
		c.bxlT0MtgU_Y.Reset(data[copied:n])
		c.osKtcdtrQlz = data
		release = false
	}
	return copied, nil
}

func (c *DVEkfear) eaRC51(buf W.Mf3phSOVqgf) error { // 0x1d1d840
	response := c.bL4KxcUfi.ToBytes()
	if !c.onlB1SoVyd.Ih7r8HyMvOa {
		block, _ := aes.NewCipher(c.bL4KxcUfi.iqQ8HJ6OMB)
		buf.XORKeyStreamWrite(cipher.NewCFBEncrypter(block, c.bL4KxcUfi.maCTf0DLoP), response)
		return nil
	}
	key := J.CVqKigbE26(c.bL4KxcUfi.iqQ8HJ6OMB, []byte("AEAD Resp Header Len Key"))[:16]
	nonce := J.CVqKigbE26(c.bL4KxcUfi.maCTf0DLoP, []byte("AEAD Resp Header Len IV"))[:12]
	buf.WriteUint16(4)
	if err := J.Fc9Ososz7(key, nonce, buf.BytesTo(0), buf.BytesTo(2), nil); err != nil {
		return err
	}
	buf.SkipEnd(16)
	key = J.CVqKigbE26(c.bL4KxcUfi.iqQ8HJ6OMB, []byte("AEAD Resp Header Key"))[:16]
	nonce = J.CVqKigbE26(c.bL4KxcUfi.maCTf0DLoP, []byte("AEAD Resp Header IV"))[:12]
	if err := J.Fc9Ososz7(key, nonce, buf.BytesRange(18, 18), response, nil); err != nil {
		return err
	}
	buf.SkipEnd(20)
	return nil
}
func (c *DVEkfear) Write(p []byte) (int, error) { // 0x1d1dee0
	encryptedLen := c.cC5h7oc7.GetEncryptLen(p)
	if encryptedLen > 16384 {
		middle := len(p) / 2
		first, err := c.Write(p[:middle])
		if err != nil {
			return first, err
		}
		second, err := c.Write(p[middle:])
		return first + second, err
	}
	headerLen := 0
	if !c.qvY6ucfXLA {
		headerLen = 4
		if c.onlB1SoVyd.Ih7r8HyMvOa {
			headerLen = 38
		}
	}
	lengthLen, padding := 0, 0
	var size uint16
	var err error
	if c.aJeMzWMqaw.OptS() {
		size, padding, err = c.dPv5C6sn(uint16(encryptedLen))
		if err != nil {
			return 0, err
		}
		lengthLen = 2
	}
	buf := B.S4OQXWIZ(encryptedLen + headerLen + lengthLen + padding)
	defer buf.Free()
	if headerLen > 0 {
		if err = c.eaRC51(buf); err != nil {
			return 0, err
		}
	}
	if lengthLen > 0 {
		buf.WriteUint16(size)
	}
	n, err := c.cC5h7oc7.Encrypt(buf.BytesRange(buf.Len(), buf.Cap()), p)
	if err != nil {
		return 0, err
	}
	buf.SkipEnd(n)
	if padding > 0 {
		buf.WriteRandom(padding)
	}
	_, err = c.Conn.Write(buf.Bytes())
	// 原实现即使 Write 错误也置位；短写且无错误仍返回完整输入长度。
	c.qvY6ucfXLA = true
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (c *DVEkfear) ReadBuffer() (W.Mf3phSOVqgf, error) { // 0x1d1e460
	if c.dvRoTXVYrAl == nil {
		c.dvRoTXVYrAl = B.S4OQXWIZ(4096)
	}
	buf := c.dvRoTXVYrAl
	buf.Reset()
	if !c.aJeMzWMqaw.OptS() {
		if _, err := buf.ReadMaxFrom(c.Conn, 4096); err != nil {
			return nil, err
		}
		return buf, nil
	}
	size, padding, err := c.x4DVqa7CuTw()
	if err != nil {
		return nil, err
	}
	if size > 32768 {
		return nil, fmt.Errorf("vmess packet size too big: %v", size)
	}
	if padding > size {
		return nil, fmt.Errorf("size(%v) < padding(%v)", size, padding)
	}
	if _, err = buf.ReadFullFrom(c.Conn, size); err != nil {
		return nil, err
	}
	buf.ReverseEnd(padding)
	n, err := c.xNm7NkjP.Decrypt(buf.Bytes(), buf.Bytes())
	if err != nil {
		Log.TRTbrquD(slog.LevelDebug, "ignored vmess decrypt err:", err)
		buf.Reset()
		return buf, nil
	}
	if n == 0 {
		return nil, io.EOF
	}
	buf.ReverseEnd(buf.Len() - n)
	return buf, nil
}
func (c *DVEkfear) WriteBuffer(buf W.Mf3phSOVqgf) (int, error) { return c.Write(buf.Bytes()) } // 0x1d1eb40
func (c *DVEkfear) Free() { // 0x1d1ebc0
	B.PoolAt50961D8.Put(c.osKtcdtrQlz)
	c.osKtcdtrQlz = nil
	if c.dvRoTXVYrAl != nil {
		c.dvRoTXVYrAl.Free()
		c.dvRoTXVYrAl = nil
	}
	H.Li3c4CJ1_D(c.Conn)
}
