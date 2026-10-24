package MFFa9yT

import (
	"bytes"
	"context"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"testing"
	"time"
	"unsafe"
)

// This client fixture builds raw metadata, key material and AEAD frames without
// calling any recovered writer/reader, parser, padding or nonce helper.
type inboundMieruClient struct {
	c             net.Conn
	a             cipher.AEAD
	send, receive []byte
	first         bool
}

func newInboundMieruClient(t *testing.T, c net.Conn, username, password string) *inboundMieruClient {
	t.Helper()
	secret := sha256.Sum256([]byte(password + "\x00" + username))
	var stamp [8]byte
	binary.BigEndian.PutUint64(stamp[:], uint64(time.Now().Round(2*time.Minute).Unix()))
	salt := sha256.Sum256(stamp[:])
	key := pbkdf2.Key(secret[:], salt[:], 64, 32, sha256.New)
	a, _ := chacha20poly1305.NewX(key)
	nonce := make([]byte, 24)
	rand.Read(nonce)
	sum := sha256.Sum256(append([]byte(username), nonce[:16]...))
	copy(nonce[20:], sum[:4])
	return &inboundMieruClient{c: c, a: a, send: nonce, first: true}
}
func inboundMieruNext(n []byte) {
	for i := len(n) - 1; i >= 0; i-- {
		n[i]++
		if n[i] != 0 {
			break
		}
	}
}
func (f *inboundMieruClient) write(kind byte, id, seq uint32, payload []byte) error {
	h := make([]byte, 32)
	h[0] = kind
	binary.BigEndian.PutUint32(h[2:6], uint32(time.Now().Unix()/60))
	binary.BigEndian.PutUint32(h[6:10], id)
	binary.BigEndian.PutUint32(h[10:14], seq)
	if kind < 6 {
		binary.BigEndian.PutUint16(h[15:17], uint16(len(payload)))
	} else {
		binary.BigEndian.PutUint16(h[22:24], uint16(len(payload)))
	}
	var b []byte
	if f.first {
		b = append(b, f.send...)
		f.first = false
	}
	b = append(b, f.a.Seal(nil, f.send, h, nil)...)
	inboundMieruNext(f.send)
	if len(payload) > 0 {
		b = append(b, f.a.Seal(nil, f.send, payload, nil)...)
		inboundMieruNext(f.send)
	}
	_, err := f.c.Write(b)
	return err
}
func (f *inboundMieruClient) read() (byte, uint32, []byte, error) {
	if f.receive == nil {
		f.receive = make([]byte, 24)
		if _, e := io.ReadFull(f.c, f.receive); e != nil {
			return 0, 0, nil, e
		}
	}
	b := make([]byte, 48)
	if _, e := io.ReadFull(f.c, b); e != nil {
		return 0, 0, nil, e
	}
	h, e := f.a.Open(nil, f.receive, b, nil)
	if e != nil {
		return 0, 0, nil, e
	}
	inboundMieruNext(f.receive)
	n, suffix, prefix := int(binary.BigEndian.Uint16(h[15:17])), int(h[17]), 0
	if h[0] >= 6 {
		n, suffix, prefix = int(binary.BigEndian.Uint16(h[22:24])), int(h[24]), int(h[21])
	}
	if _, e = io.CopyN(io.Discard, f.c, int64(prefix)); e != nil {
		return 0, 0, nil, e
	}
	var p []byte
	if n > 0 {
		b = make([]byte, n+16)
		if _, e = io.ReadFull(f.c, b); e != nil {
			return 0, 0, nil, e
		}
		p, e = f.a.Open(nil, f.receive, b, nil)
		if e != nil {
			return 0, 0, nil, e
		}
		inboundMieruNext(f.receive)
	}
	if _, e = io.CopyN(io.Discard, f.c, int64(suffix)); e != nil {
		return 0, 0, nil, e
	}
	return h[0], binary.BigEndian.Uint32(h[6:10]), p, nil
}
func TestNativeMieruConstructStartTCPUDPStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	if unsafe.Sizeof(ONGT2EHM40r{}) != 40 {
		t.Fatal("RTTI")
	}
	path := filepath.Join(t.TempDir(), "soga.conf")
	os.WriteFile(path, []byte("listen=127.0.0.1\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\n"), 0600)
	C.MkwNWciPHS(path)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	for _, network := range []string{"tcp", "udp"} {
		t.Run(network, func(t *testing.T) {
			node := &L.Wr5SVP{EMlXEc: 10, ABA4hq: L.Xh4IiOa0{GGSLzD: "TCP"}}
			panel := &L.Aga91y{WZ89pTAOrHK4: 111, JAfXWwzYrW: &L.DO9qqM9{Ejujyaj: "fixture", GltkJGvV: "password"}}
			core, e := tzgQP2vl(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{panel}}})
			if e != nil {
				t.Fatal(e)
			}
			v := core.(*ONGT2EHM40r)
			defer v.Stop()
			// A real recent-check timestamp selects the proved no-request branch.
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			if e = v.Start(); e != nil {
				t.Fatal(e)
			}
			if len(v.d7Fav9U0ykgg) != 1 {
				t.Fatal("listeners")
			}
			var port int
			if network == "tcp" {
				echo, e := net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer echo.Close()
				port = echo.Addr().(*net.TCPAddr).Port
				go func() {
					c, e := echo.Accept()
					if e != nil {
						return
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(5 * time.Second))
					io.Copy(c, c)
				}()
			} else {
				echo, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer echo.Close()
				port = echo.LocalAddr().(*net.UDPAddr).Port
				go func() {
					echo.SetDeadline(time.Now().Add(5 * time.Second))
					b := make([]byte, 4096)
					n, a, e := echo.ReadFrom(b)
					if e == nil {
						echo.WriteTo(b[:n], a)
					}
				}()
			}
			conn, e := net.DialTimeout("tcp", v.d7Fav9U0ykgg[0].Addr().String(), time.Second)
			if e != nil {
				t.Fatal(e)
			}
			defer conn.Close()
			conn.SetDeadline(time.Now().Add(5 * time.Second))
			client := newInboundMieruClient(t, conn, "fixture", "password")
			cmd := byte(1)
			if network == "udp" {
				cmd = 3
			}
			request := binary.BigEndian.AppendUint16([]byte{5, cmd, 0, 1, 127, 0, 0, 1}, uint16(port))
			if e = client.write(2, 99, 0, request); e != nil {
				t.Fatal(e)
			}
			kind, id, p, e := client.read()
			if e != nil || kind != 3 || id != 99 || len(p) != 0 {
				t.Fatal("open", kind, id, p, e)
			}
			kind, id, p, e = client.read()
			if e != nil || kind != 7 || id != 99 || !bytes.Equal(p, []byte{5, 0, 0, 1, 0, 0, 0, 0, 0, 0}) {
				t.Fatal("request", kind, id, p, e)
			}
			payload := bytes.Repeat([]byte("mieru-native\x00"), 100)
			send := payload
			if network == "udp" {
				inner := binary.BigEndian.AppendUint16([]byte{0, 0, 0, 1, 127, 0, 0, 1}, uint16(port))
				inner = append(inner, payload...)
				send = binary.BigEndian.AppendUint16([]byte{0}, uint16(len(inner)))
				send = append(send, inner...)
				send = append(send, 255)
			}
			if e = client.write(6, 99, 1, send); e != nil {
				t.Fatal(e)
			}
			var got []byte
			for len(got) < len(send) {
				kind, id, p, e = client.read()
				if e != nil || kind != 7 || id != 99 {
					t.Fatal("echo", kind, id, e)
				}
				got = append(got, p...)
			}
			if !bytes.Equal(got, send) {
				t.Fatalf("echo differs len %d/%d", len(got), len(send))
			}
			conn.Close()
			v.Stop()
			select {
			case <-v.zb8f3cHWW8v.Done():
			default:
				t.Fatal("not canceled")
			}
			if _, e = net.DialTimeout("tcp", v.d7Fav9U0ykgg[0].Addr().String(), 100*time.Millisecond); e == nil {
				t.Fatal("listener still open")
			}
		})
	}
}
