package gk1z9v

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"io"
	"net"
	"testing"
	"time"
	"unsafe"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/sha3"
	B "recovered.local/soga/D_zCag1qTV"
	L "recovered.local/soga/LFWTTXF7JR"
	U "recovered.local/soga/nzATvajt"
)

// 测试客户端独立构造标准层叠 HMAC，不调用恢复工程 KDF。
func clientKDF(key []byte, path ...[]byte) []byte {
	var create func() hash.Hash = func() hash.Hash { return hmac.New(sha256.New, []byte("VMess AEAD KDF")) }
	for _, piece := range path {
		previous := create
		p := append([]byte(nil), piece...)
		create = func() hash.Hash { return hmac.New(previous, p) }
	}
	h := create()
	h.Write(key)
	return h.Sum(nil)
}
func clientGCM(key, nonce, plain, aad []byte) []byte {
	block, _ := aes.NewCipher(key)
	a, _ := cipher.NewGCM(block)
	return a.Seal(nil, nonce, plain, aad)
}
func clientOpen(key, nonce, encrypted []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	a, _ := cipher.NewGCM(block)
	return a.Open(nil, nonce, encrypted, nil)
}
func nativeAuthFixture(t *testing.T, aead bool, id byte) *WR4lMNfab {
	t.Helper()
	user, e := U.EOkO09L(&L.Aga91y{WZ89pTAOrHK4: 31, JAfXWwzYrW: &L.Yn8BJoL{SfiEGpJo: "01234567-89ab-cdef-0123-456789abcdef"}}, nil)
	if e != nil {
		t.Fatal(e)
	}
	a := &WR4lMNfab{Xwwpafq: user, QAkNeaNePGH: 1700000000, Ih7r8HyMvOa: aead, T_ul7Q_0j: bytes.Repeat([]byte{id}, 16)}
	copy(a.B2pr6gMaPO[:], a.T_ul7Q_0j)
	return a
}
func clientHandshake(auth *WR4lMNfab, clear []byte) []byte {
	if !auth.Ih7r8HyMvOa {
		block, _ := aes.NewCipher(auth.GetCmdKey())
		out := make([]byte, len(clear))
		cipher.NewCFBEncrypter(block, auth.GetIV()).XORKeyStream(out, clear)
		return out
	}
	nonce := []byte("nonce123")
	length := binary.BigEndian.AppendUint16(nil, uint16(len(clear)))
	k := clientKDF(auth.GetCmdKey(), []byte("VMess Header AEAD Key_Length"), auth.T_ul7Q_0j, nonce)[:16]
	iv := clientKDF(auth.GetCmdKey(), []byte("VMess Header AEAD Nonce_Length"), auth.T_ul7Q_0j, nonce)[:12]
	out := clientGCM(k, iv, length, auth.T_ul7Q_0j)
	out = append(out, nonce...)
	k = clientKDF(auth.GetCmdKey(), []byte("VMess Header AEAD Key"), auth.T_ul7Q_0j, nonce)[:16]
	iv = clientKDF(auth.GetCmdKey(), []byte("VMess Header AEAD Nonce"), auth.T_ul7Q_0j, nonce)[:12]
	return append(out, clientGCM(k, iv, clear, auth.T_ul7Q_0j)...)
}
func clientAEAD(sec IPas09, key []byte) cipher.AEAD {
	if sec == 3 {
		b, _ := aes.NewCipher(key)
		a, _ := cipher.NewGCM(b)
		return a
	}
	a := md5.Sum(key)
	b := md5.Sum(a[:])
	c, _ := chacha20poly1305.New(append(a[:], b[:]...))
	return c
}
func clientFrame(sec IPas09, key, iv, p []byte, mask sha3.ShakeHash, padded bool, counter uint16) []byte {
	payload := append([]byte(nil), p...)
	if sec != 5 {
		nonce := append([]byte(nil), iv[:12]...)
		binary.BigEndian.PutUint16(nonce, counter)
		payload = clientAEAD(sec, key).Seal(nil, nonce, p, nil)
	}
	size := uint16(len(payload))
	padding := 0
	var scratch [2]byte
	if mask != nil {
		if padded {
			mask.Read(scratch[:])
			padding = int(binary.BigEndian.Uint16(scratch[:]) & 63)
		}
		mask.Read(scratch[:])
		size = (size + uint16(padding)) ^ binary.BigEndian.Uint16(scratch[:])
	}
	out := binary.BigEndian.AppendUint16(nil, size)
	out = append(out, payload...)
	return append(out, bytes.Repeat([]byte{0x57}, padding)...)
}
func readClientFrame(r io.Reader, sec IPas09, key, iv []byte, mask sha3.ShakeHash, padded bool, counter uint16) ([]byte, error) {
	var length [2]byte
	if _, e := io.ReadFull(r, length[:]); e != nil {
		return nil, e
	}
	size := binary.BigEndian.Uint16(length[:])
	padding := 0
	if mask != nil {
		if padded {
			mask.Read(length[:])
			padding = int(binary.BigEndian.Uint16(length[:]) & 63)
		}
		mask.Read(length[:])
		size ^= binary.BigEndian.Uint16(length[:])
	}
	data := make([]byte, size)
	if _, e := io.ReadFull(r, data); e != nil {
		return nil, e
	}
	data = data[:len(data)-padding]
	if sec == 5 {
		return data, nil
	}
	nonce := append([]byte(nil), iv[:12]...)
	binary.BigEndian.PutUint16(nonce, counter)
	return clientAEAD(sec, key).Open(nil, nonce, data, nil)
}
func clientResponseKeys(header []byte, aead bool) (key, iv []byte) {
	if aead {
		k, v := sha256.Sum256(header[17:33]), sha256.Sum256(header[1:17])
		return k[:16], v[:16]
	}
	k, v := md5.Sum(header[17:33]), md5.Sum(header[1:17])
	return k[:], v[:]
}
func readClientResponse(r io.Reader, key, iv []byte, aead bool) ([]byte, error) {
	if !aead {
		p := make([]byte, 4)
		if _, e := io.ReadFull(r, p); e != nil {
			return nil, e
		}
		b, _ := aes.NewCipher(key)
		cipher.NewCFBDecrypter(b, iv).XORKeyStream(p, p)
		return p, nil
	}
	length := make([]byte, 18)
	if _, e := io.ReadFull(r, length); e != nil {
		return nil, e
	}
	p, e := clientOpen(clientKDF(key, []byte("AEAD Resp Header Len Key"))[:16], clientKDF(iv, []byte("AEAD Resp Header Len IV"))[:12], length)
	if e != nil {
		return nil, e
	}
	body := make([]byte, int(binary.BigEndian.Uint16(p))+16)
	if _, e = io.ReadFull(r, body); e != nil {
		return nil, e
	}
	return clientOpen(clientKDF(key, []byte("AEAD Resp Header Key"))[:16], clientKDF(iv, []byte("AEAD Resp Header IV"))[:12], body)
}
func TestNativeVMessHandshakeAndRealDuplexFrames(t *testing.T) {
	if unsafe.Sizeof(DVEkfear{}) != 192 || unsafe.Offsetof(DVEkfear{}.dvRoTXVYrAl) != 176 || unsafe.Sizeof(KUJQ1W{}) != 32 {
		t.Fatal("RTTI")
	}
	for _, aead := range []bool{false, true} {
		for _, sec := range []IPas09{3, 4, 5} {
			for _, options := range []byte{1, 5, 13} {
				t.Run(fmt.Sprintf("aead=%v/sec=%v/opts=%v", aead, sec, options), func(t *testing.T) {
					auth := nativeAuthFixture(t, aead, byte(sec)*20+options)
					if aead {
						auth.B2pr6gMaPO[1] = 1
					}
					clear := headerFixture(1, 1, options, []byte{192, 0, 2, 7})
					clear[35] = 0x20 | byte(sec)
					clear = finishHeaderChecksum(clear[:len(clear)-4])
					inbound := bytes.Repeat([]byte("upload"), 1900)
					outbound := bytes.Repeat([]byte("download"), 6000)
					listener, e := net.Listen("tcp4", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					client, e := net.DialTimeout("tcp4", listener.Addr().String(), time.Second)
					if e != nil {
						listener.Close()
						t.Fatal(e)
					}
					server, e := listener.Accept()
					listener.Close()
					if e != nil {
						client.Close()
						t.Fatal(e)
					}
					defer server.Close()
					defer client.Close()
					server.SetDeadline(time.Now().Add(5 * time.Second))
					client.SetDeadline(time.Now().Add(5 * time.Second))
					done := make(chan error, 1)
					go func() {
						c, e := EAQVlR(server, auth)
						if e != nil {
							done <- e
							return
						}
						defer c.Free()
						if c.GetDestAddr().HostPort() != "192.0.2.7:443" || c.GetUser() != auth.Xwwpafq || c.NetworkType() != 0 {
							done <- errors.New("metadata")
							return
						}
						got := make([]byte, len(inbound))
						for at := 0; at < len(got); {
							end := at + 79
							if end > len(got) {
								end = len(got)
							}
							n, e := c.Read(got[at:end])
							if e != nil {
								done <- e
								return
							}
							at += n
						}
						if !bytes.Equal(got, inbound) {
							done <- errors.New("upload mismatch")
							return
						}
						b := B.S4OQXWIZ(len(outbound))
						defer b.Free()
						b.Write(outbound)
						n, e := c.WriteBuffer(b)
						if e == nil && n != len(outbound) {
							e = errors.New("write count")
						}
						done <- e
					}()
					if _, e := client.Write(clientHandshake(auth, clear)); e != nil {
						t.Fatal(e)
					}
					var inMask, outMask sha3.ShakeHash
					key, iv := clientResponseKeys(clear, aead)
					if options&4 != 0 {
						inMask = sha3.NewShake128()
						inMask.Write(clear[1:17])
						outMask = sha3.NewShake128()
						outMask.Write(iv)
					}
					if _, e := client.Write(clientFrame(sec, clear[17:33], clear[1:17], inbound, inMask, options&8 != 0, 0)); e != nil {
						t.Fatal(e)
					}
					response, e := readClientResponse(client, key, iv, aead)
					if e != nil || !bytes.Equal(response, []byte{0x42, 0, 0, 0}) {
						t.Fatal(response, e)
					}
					var got []byte
					for counter := uint16(0); len(got) < len(outbound); counter++ {
						p, e := readClientFrame(client, sec, key, iv, outMask, options&8 != 0, counter)
						if e != nil {
							t.Fatal(e)
						}
						got = append(got, p...)
					}
					if !bytes.Equal(got, outbound) {
						t.Fatal("download mismatch")
					}
					if e := <-done; e != nil {
						t.Fatal(e)
					}
				})
			}
		}
	}
}

type memoryConn struct {
	bytes.Buffer
	writeErr error
	short    bool
	freed    int
}

func (c *memoryConn) Read(p []byte) (int, error) { return c.Buffer.Read(p) }
func (c *memoryConn) Write(p []byte) (int, error) {
	if c.writeErr != nil {
		return 0, c.writeErr
	}
	if c.short {
		return 0, nil
	}
	return c.Buffer.Write(p)
}
func (c *memoryConn) Close() error                     { return nil }
func (c *memoryConn) LocalAddr() net.Addr              { return &net.TCPAddr{} }
func (c *memoryConn) RemoteAddr() net.Addr             { return &net.TCPAddr{} }
func (c *memoryConn) SetDeadline(time.Time) error      { return nil }
func (c *memoryConn) SetReadDeadline(time.Time) error  { return nil }
func (c *memoryConn) SetWriteDeadline(time.Time) error { return nil }
func (c *memoryConn) Free()                            { c.freed++ }
func TestNativeVMessReplayAndTamperedHandshake(t *testing.T) {
	for _, aead := range []bool{false, true} {
		auth := nativeAuthFixture(t, aead, 241)
		if aead {
			auth.B2pr6gMaPO[1] = 9
		}
		clear := headerFixture(1, 1, 1, []byte{192, 0, 2, 8})
		wire := clientHandshake(auth, clear)
		raw := new(memoryConn)
		raw.Buffer.Write(wire)
		c, e := EAQVlR(raw, auth)
		if e != nil {
			t.Fatal(e)
		}
		c.Free()
		raw = new(memoryConn)
		raw.Buffer.Write(wire)
		if c, e = EAQVlR(raw, auth); c != nil || e != replayErrorAt5098420 || raw.freed != 1 {
			t.Fatal(c, e, raw.freed)
		}
		if aead {
			for _, offset := range []int{0, 17, 26, len(wire) - 1} {
				bad := append([]byte(nil), wire...)
				bad[offset] ^= 1
				raw = new(memoryConn)
				raw.Buffer.Write(bad)
				if c, e = EAQVlR(raw, auth); c != nil || e == nil || e == replayErrorAt5098420 {
					t.Fatal(offset, c, e)
				}
			}
		}
	}
}
func TestNativeVMessReadBufferFailuresAndWriteSemantics(t *testing.T) {
	auth := nativeAuthFixture(t, true, 242)
	h, e := Gk1af7wDMnD(bytes.NewReader(headerFixture(1, 1, 1, []byte{192, 0, 2, 9})))
	if e != nil {
		t.Fatal(e)
	}
	raw := new(memoryConn)
	c, e := W2tVTGQ(raw, auth, h)
	if e != nil {
		t.Fatal(e)
	}
	defer c.Free()
	frame := clientFrame(3, h.ezCHRvHCePBC, h.cEcPFEa7, []byte("buffer payload"), nil, false, 0)
	raw.Buffer.Write(frame)
	b, e := c.ReadBuffer()
	if e != nil || string(b.Bytes()) != "buffer payload" {
		t.Fatal(b, e)
	}
	bad := clientFrame(3, h.ezCHRvHCePBC, h.cEcPFEa7, []byte("bad frame"), nil, false, 1)
	bad[len(bad)-1] ^= 1
	raw.Buffer.Write(bad)
	b, e = c.ReadBuffer()
	if e != nil || !b.IsEmpty() {
		t.Fatal("original suppresses decrypt errors", b, e)
	}
	raw.Buffer.Write(clientFrame(3, h.ezCHRvHCePBC, h.cEcPFEa7, nil, nil, false, 2))
	if b, e = c.ReadBuffer(); b != nil || e != io.EOF {
		t.Fatal(b, e)
	}
	raw.Buffer.Write([]byte{0x80, 1})
	if b, e = c.ReadBuffer(); b != nil || e == nil || e.Error() != "vmess packet size too big: 32769" {
		t.Fatal(b, e)
	}
	raw.writeErr = io.ErrClosedPipe
	n, e := c.Write([]byte("x"))
	if n != 0 || e != io.ErrClosedPipe || !c.qvY6ucfXLA {
		t.Fatal(n, e, c.qvY6ucfXLA)
	}
	raw.writeErr = nil
	raw.short = true
	n, e = c.Write([]byte("x"))
	if n != 1 || e != nil {
		t.Fatal("original ignores short write count", n, e)
	}
}
