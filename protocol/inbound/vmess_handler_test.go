package inbound

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"hash"
	"hash/fnv"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	U "recovered.local/soga/user"
	"testing"
	"time"
)

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
func clientResponseKeys(header []byte, aead bool) (key, iv []byte) {
	if aead {
		k, v := sha256.Sum256(header[17:33]), sha256.Sum256(header[1:17])
		return k[:16], v[:16]
	}
	k, v := md5.Sum(header[17:33]), md5.Sum(header[1:17])
	return k[:], v[:]
}

// 完整认证→原握手→公共策略→真实TCP/UDP目标，客户端独立编码wire。
func TestNativeVMessHandlerRealDestinations(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	p := filepath.Join(t.TempDir(), "soga.conf")
	os.WriteFile(p, []byte("dy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\n"), 0600)
	C.MkwNWciPHS(p)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	for _, kind := range []string{"tcp", "udp"} {
		t.Run(kind, func(t *testing.T) {
			v := newVmessAuthFixture(t)
			v.wiy6P3.Store(&L.Wr5SVP{EMlXEc: 31})
			v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(7, "11111111-1111-4111-8111-111111111111")})
			base, _ := v.jc_QPN6.Get(7)
			u := base.(*U.QKaIWV2)
			// 构造“本小时已完成状态检查”的原合法状态，只测试本地传输。
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			var port int
			echoDone := make(chan error, 1)
			if kind == "tcp" {
				ln, e := net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer ln.Close()
				port = ln.Addr().(*net.TCPAddr).Port
				go func() {
					c, e := ln.Accept()
					if e != nil {
						echoDone <- e
						return
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(5 * time.Second))
					_, e = io.Copy(c, c)
					echoDone <- e
				}()
			} else {
				pc, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer pc.Close()
				port = pc.LocalAddr().(*net.UDPAddr).Port
				go func() {
					pc.SetDeadline(time.Now().Add(5 * time.Second))
					b := make([]byte, 4096)
					n, a, e := pc.ReadFrom(b)
					if e == nil {
						_, e = pc.WriteTo(b[:n], a)
					}
					echoDone <- e
				}()
			}
			ln, e := net.Listen("tcp4", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer ln.Close()
			c, e := net.DialTimeout("tcp4", ln.Addr().String(), time.Second)
			if e != nil {
				t.Fatal(e)
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(5 * time.Second))
			raw, e := ln.Accept()
			if e != nil {
				t.Fatal(e)
			}
			defer raw.Close()
			done := make(chan struct{})
			go func() { v.tFdBaee6T(raw); close(done) }()
			exerciseNativeVMessWire(t, c, u, port, kind)
			c.Close()
			select {
			case <-done:
			case <-time.After(5 * time.Second):
				t.Fatal("handler did not close")
			}
			select {
			case e = <-echoDone:
				if e != nil {
					t.Fatal(e)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("echo did not close")
			}
		})
	}
}

func exerciseNativeVMessWire(t *testing.T, c io.ReadWriter, u *U.QKaIWV2, port int, kind string) {
	exerciseNativeVMessWireMode(t, c, u, port, kind, true, time.Now().Unix())
}
func legacyVMessAuthID(u *U.QKaIWV2, stamp int64) (id [16]byte) {
	key := u.GetUserIds()[0].Yfd33z6YYdJa
	h := hmac.New(md5.New, key[:])
	binary.Write(h, binary.BigEndian, uint64(stamp))
	copy(id[:], h.Sum(nil))
	return
}
func exerciseNativeVMessWireMode(t *testing.T, c io.ReadWriter, u *U.QKaIWV2, port int, kind string, aead bool, stamp int64) {
	t.Helper()
	wire, clear, payload := nativeVMessFixtureWire(t, u, port, kind, aead, stamp)
	var e error
	if _, e = c.Write(wire); e != nil {
		t.Fatal(e)
	}
	key, iv := clientResponseKeys(clear, aead)
	response, e := readClientResponse(c, key, iv, aead)
	if e != nil || len(response) != 4 || response[0] != 0x42 {
		t.Fatalf("response %x: kind=%s aead=%v target_port=%d auth_timestamp=%d request_iv_key=%x: %v", response, kind, aead, port, stamp, clear[1:33], e)
	}
	got := make([]byte, 0, len(payload))
	for len(got) < len(payload) {
		var size [2]byte
		if _, e = io.ReadFull(c, size[:]); e != nil {
			t.Fatal(e)
		}
		b := make([]byte, binary.BigEndian.Uint16(size[:]))
		if _, e = io.ReadFull(c, b); e != nil {
			t.Fatal(e)
		}
		got = append(got, b...)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("relay changed data: %d", len(got))
	}
}

// Independent clients need independent request entropy even when two targets
// share the low byte of their port. The original replay key ignores destination
// and transport; legacy MD5 clients in the same second also share the AuthID.
func nativeVMessFixtureWire(t *testing.T, u *U.QKaIWV2, port int, kind string, aead bool, stamp int64) (wire, clear, payload []byte) {
	t.Helper()
	var session [32]byte
	if _, err := rand.Read(session[:]); err != nil {
		t.Fatal(err)
	}
	return nativeVMessFixtureWireKeys(t, u, port, kind, aead, stamp, session)
}

// Explicit keys are only for demonstrating original replay rejection.
func nativeVMessFixtureWireKeys(t *testing.T, u *U.QKaIWV2, port int, kind string, aead bool, stamp int64, session [32]byte) (wire, clear, payload []byte) {
	t.Helper()
	clear = make([]byte, 38)
	clear[0] = 1
	copy(clear[1:33], session[:])
	clear[33] = 0x42
	clear[34] = 1
	clear[35] = 5
	clear[37] = 1
	if kind == "udp" {
		clear[37] = 2
	}
	clear = binary.BigEndian.AppendUint16(clear, uint16(port))
	clear = append(clear, 1, 127, 0, 0, 1)
	h := fnv.New32a()
	h.Write(clear)
	clear = binary.BigEndian.AppendUint32(clear, h.Sum32())
	cmd := u.GetCmdKey()
	if aead {
		id := vmessAuthID(u, stamp, true)
		nonce := []byte("nonce123")
		seal := func(label, ivlabel string, b []byte) []byte {
			return clientGCM(clientKDF(cmd, []byte(label), id[:], nonce)[:16], clientKDF(cmd, []byte(ivlabel), id[:], nonce)[:12], b, id[:])
		}
		wire = append(wire, id[:]...)
		wire = append(wire, seal("VMess Header AEAD Key_Length", "VMess Header AEAD Nonce_Length", binary.BigEndian.AppendUint16(nil, uint16(len(clear))))...)
		wire = append(wire, nonce...)
		wire = append(wire, seal("VMess Header AEAD Key", "VMess Header AEAD Nonce", clear)...)
	} else {
		id := legacyVMessAuthID(u, stamp)
		wire = append(wire, id[:]...)
		var timestamp [8]byte
		binary.BigEndian.PutUint64(timestamp[:], uint64(stamp))
		iv := md5.Sum(bytes.Repeat(timestamp[:], 4))
		block, err := aes.NewCipher(cmd)
		if err != nil {
			t.Fatal(err)
		}
		header := append([]byte(nil), clear...)
		cipher.NewCFBEncrypter(block, iv[:]).XORKeyStream(header, header)
		wire = append(wire, header...)
	}
	payload = bytes.Repeat([]byte("native vmess payload\x00"), 100)
	wire = binary.BigEndian.AppendUint16(wire, uint16(len(payload)))
	wire = append(wire, payload...)
	return
}
