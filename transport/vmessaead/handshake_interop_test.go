package vmessaead

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
	"hash/crc32"
	"hash/fnv"
	"io"
	"net"
	"testing"
	"time"
)

func peerRead(r io.Reader, n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(r, b)
	return b, err
}

// This peer is separate test code which parses the wire, authenticates requests,
// and seals responses. It does not use recovered frames or handshake methods.
func peerReadHeader(c net.Conn, commandKey, uuid []byte, aead bool) ([]byte, error) {
	auth, e := peerRead(c, 16)
	if e != nil {
		return nil, e
	}
	if aead {
		b, _ := aes.NewCipher(peerKDF(commandKey, "AES Auth ID Encryption")[:16])
		plain := make([]byte, 16)
		b.Decrypt(plain, auth)
		if binary.BigEndian.Uint32(plain[12:]) != crc32.ChecksumIEEE(plain[:12]) {
			return nil, fmt.Errorf("auth CRC")
		}
		encLen, e := peerRead(c, 18)
		if e != nil {
			return nil, e
		}
		nonce, e := peerRead(c, 8)
		if e != nil {
			return nil, e
		}
		key := peerKDF(commandKey, "VMess Header AEAD Key_Length", string(auth), string(nonce))
		iv := peerKDF(commandKey, "VMess Header AEAD Nonce_Length", string(auth), string(nonce))
		length, e := peerGCM(key[:16]).Open(nil, iv[:12], encLen, auth)
		if e != nil {
			return nil, e
		}
		enc, e := peerRead(c, int(binary.BigEndian.Uint16(length))+16)
		if e != nil {
			return nil, e
		}
		key = peerKDF(commandKey, "VMess Header AEAD Key", string(auth), string(nonce))
		iv = peerKDF(commandKey, "VMess Header AEAD Nonce", string(auth), string(nonce))
		return peerGCM(key[:16]).Open(nil, iv[:12], enc, auth)
	}
	var timestamp int64
	now := time.Now().Unix()
	for stamp := now - 10; stamp <= now+10; stamp++ {
		h := hmac.New(md5.New, uuid)
		binary.Write(h, binary.BigEndian, uint64(stamp))
		if hmac.Equal(h.Sum(nil), auth) {
			timestamp = stamp
			break
		}
	}
	if timestamp == 0 {
		return nil, fmt.Errorf("legacy auth")
	}
	var stamp [8]byte
	binary.BigEndian.PutUint64(stamp[:], uint64(timestamp))
	iv := md5.Sum(bytes.Repeat(stamp[:], 4))
	block, _ := aes.NewCipher(commandKey)
	stream := cipher.NewCFBDecrypter(block, iv[:])
	first, e := peerRead(c, 41)
	if e != nil {
		return nil, e
	}
	stream.XORKeyStream(first, first)
	if first[40] != 1 {
		return nil, fmt.Errorf("expected IPv4")
	}
	tail, e := peerRead(c, 4+int(first[35]>>4)+4)
	if e != nil {
		return nil, e
	}
	stream.XORKeyStream(tail, tail)
	return append(first, tail...), nil
}
func peerBodyAEAD(key []byte, security byte) cipher.AEAD {
	if security == 3 {
		return peerGCM(key)
	}
	a := md5.Sum(key)
	b := md5.Sum(a[:])
	joined := append(a[:], b[:]...)
	out, e := chacha20poly1305.New(joined)
	if e != nil {
		panic(e)
	}
	return out
}
func peerResponseHeader(key, iv []byte, auth byte, aead bool) []byte {
	plain := []byte{auth, 0, 0, 0}
	if !aead {
		block, _ := aes.NewCipher(key)
		out := make([]byte, 4)
		cipher.NewCFBEncrypter(block, iv).XORKeyStream(out, plain)
		return out
	}
	lk := peerKDF(key, "AEAD Resp Header Len Key")
	ln := peerKDF(iv, "AEAD Resp Header Len IV")
	out := peerGCM(lk[:16]).Seal(nil, ln[:12], []byte{0, 4}, nil)
	pk := peerKDF(key, "AEAD Resp Header Key")
	pn := peerKDF(iv, "AEAD Resp Header IV")
	return append(out, peerGCM(pk[:16]).Seal(nil, pn[:12], plain, nil)...)
}
func peerEcho(c net.Conn, commandKey, uuid []byte, aead, udp bool, payload []byte) error {
	defer c.Close()
	c.SetDeadline(time.Now().Add(5 * time.Second))
	header, e := peerReadHeader(c, commandKey, uuid, aead)
	if e != nil {
		return e
	}
	if len(header) < 49 {
		return fmt.Errorf("short header")
	}
	f := fnv.New32a()
	f.Write(header[:len(header)-4])
	if f.Sum32() != binary.BigEndian.Uint32(header[len(header)-4:]) {
		return fmt.Errorf("request FNV")
	}
	command := byte(1)
	if udp {
		command = 2
	}
	if header[0] != 1 || header[36] != 0 || header[37] != command || binary.BigEndian.Uint16(header[38:40]) != 443 || !bytes.Equal(header[41:45], []byte{127, 0, 0, 1}) {
		return fmt.Errorf("request fields %x", header[:45])
	}
	iv, key := header[1:17], header[17:33]
	var responseKey, responseIV []byte
	if aead {
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
	options, security := header[34], header[35]&15
	var body []byte
	if options == 0 {
		body, e = peerRead(c, len(payload))
	} else {
		length, err := peerRead(c, 2)
		if err != nil {
			return err
		}
		body, e = peerRead(c, int(binary.BigEndian.Uint16(length)))
		if e == nil && security != 5 {
			nonce := append([]byte{0, 0}, iv[2:12]...)
			body, e = peerBodyAEAD(key, security).Open(nil, nonce, body, nil)
		}
	}
	if e != nil {
		return e
	}
	if !bytes.Equal(body, payload) {
		return fmt.Errorf("request body %x", body)
	}
	out := peerResponseHeader(responseKey, responseIV, header[33], aead)
	reply := append([]byte("echo:"), body...)
	if options == 0 {
		out = append(out, reply...)
	} else {
		if security != 5 {
			nonce := append([]byte{0, 0}, responseIV[2:12]...)
			reply = peerBodyAEAD(responseKey, security).Seal(nil, nonce, reply, nil)
		}
		out = append(out, byte(len(reply)>>8), byte(len(reply)))
		out = append(out, reply...)
	}
	_, e = c.Write(out)
	return e
}
func TestVMessTCPInteroperability(t *testing.T) {
	for _, aead := range []bool{false, true} {
		for _, security := range []string{"aes-128-gcm", "chacha20-poly1305", "none", "zero"} {
			for _, udp := range []bool{false, true} {
				t.Run(fmt.Sprintf("aead=%t/%s/udp=%t", aead, security, udp), func(t *testing.T) {
					client, e := Iw0IEMM3ha(ClientConfiguration{UUID: "00010203-0405-0607-0809-0a0b0c0d0e0f", Security: security, AEAD: aead})
					if e != nil {
						t.Fatal(e)
					}
					listener, e := net.Listen("tcp", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					defer listener.Close()
					done := make(chan error, 1)
					payload := []byte("actual TCP VMess payload")
					go func() {
						conn, e := listener.Accept()
						if e != nil {
							done <- e
							return
						}
						done <- peerEcho(conn, client.roEClw[0].P3zlkY2W1n, client.gyJC3dk_8[:], aead, udp, payload)
					}()
					raw, e := net.Dial("tcp", listener.Addr().String())
					if e != nil {
						t.Fatal(e)
					}
					defer raw.Close()
					raw.SetDeadline(time.Now().Add(5 * time.Second))
					conn, e := client.StreamConn(raw, &IxhMZ66CqoH1{GV3P5rY: udp, OyEjzi1v93: 1, TIBiNxg: []byte{127, 0, 0, 1}, GHP6ivev2Db: 443})
					if e != nil {
						t.Fatal(e)
					}
					if n, e := conn.Write(payload); e != nil || n != len(payload) {
						t.Fatal(n, e)
					}
					reply := make([]byte, len(payload)+5)
					_, e = io.ReadFull(conn, reply)
					if e != nil {
						t.Fatal(e)
					}
					if string(reply) != "echo:"+string(payload) {
						t.Fatal(string(reply))
					}
					if e = <-done; e != nil {
						t.Fatal(e)
					}
				})
			}
		}
	}
}
