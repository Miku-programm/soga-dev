package MFFa9yT

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
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
)

func rawAnyTLSFrame(cmd byte, id uint32, payload []byte) []byte {
	b := binary.BigEndian.AppendUint32([]byte{cmd}, id)
	b = binary.BigEndian.AppendUint16(b, uint16(len(payload)))
	return append(b, payload...)
}
func TestNativeAnyTLSConstructStartTCPUDPStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	path := filepath.Join(t.TempDir(), "soga.conf")
	os.WriteFile(path, []byte("listen=127.0.0.1\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\n"), 0600)
	C.MkwNWciPHS(path)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	for _, kind := range []string{"tcp", "udp"} {
		t.Run(kind, func(t *testing.T) {
			node := &L.Wr5SVP{EMlXEc: 32, QaZjDzw95: L.Tt_Moslrup{CgE0xkW: 0, J6bt8AOW8fL8: "none", UPnToXrM: "stop=0"}}
			userPanel := &L.Aga91y{WZ89pTAOrHK4: 9, JAfXWwzYrW: &L.ZtMFlBTZXyD{CBWeRfrO1ahs: "native-anytls-password"}}
			data := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{userPanel}}
			core, e := aboatPBa(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: data})
			if e != nil {
				t.Fatal(e)
			}
			v := core.(*R_n5ucb)
			defer v.Stop()
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			if e = v.Start(); e != nil {
				t.Fatal(e)
			}
			if len(v.qjukoB) != 1 {
				t.Fatal("listener count")
			}
			addr := v.qjukoB[0].Addr().String()
			var port int
			done := make(chan error, 1)
			if kind == "tcp" {
				echo, e := net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer echo.Close()
				port = echo.Addr().(*net.TCPAddr).Port
				go func() {
					c, e := echo.Accept()
					if e != nil {
						done <- e
						return
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(5 * time.Second))
					_, e = io.Copy(c, c)
					done <- e
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
						_, e = echo.WriteTo(b[:n], a)
					}
					done <- e
				}()
			}
			c, e := net.DialTimeout("tcp", addr, time.Second)
			if e != nil {
				t.Fatal(e)
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(5 * time.Second))
			auth := sha256.Sum256([]byte("native-anytls-password"))
			wire := append([]byte(nil), auth[:]...)
			wire = append(wire, 0, 3, 0xaa, 0xbb, 0xcc)
			settings := []byte(fmt.Sprintf("v=2\nclient=independent-test\npadding-md5=%x", md5.Sum([]byte("stop=0"))))
			wire = append(wire, rawAnyTLSFrame(4, 0, settings)...)
			wire = append(wire, rawAnyTLSFrame(1, 1, nil)...)
			target := binary.BigEndian.AppendUint16([]byte{1, 127, 0, 0, 1}, uint16(port))
			payload := bytes.Repeat([]byte("anytls-native\x00"), 200)
			request := append([]byte(nil), target...)
			expected := payload
			if kind == "udp" {
				magic := "sp.v2.udp-over-tcp.arpa"
				request = append([]byte{3, byte(len(magic))}, []byte(magic)...)
				request = append(request, 0, 0)
				request = append(request, 1)
				request = append(request, target...)
				expected = binary.BigEndian.AppendUint16(nil, uint16(len(payload)))
				expected = append(expected, payload...)
			}
			request = append(request, expected...)
			wire = append(wire, rawAnyTLSFrame(2, 1, request)...)
			if _, e = c.Write(wire); e != nil {
				t.Fatal(e)
			}
			var got []byte
			for len(got) < len(expected) {
				var h [7]byte
				if _, e = io.ReadFull(c, h[:]); e != nil {
					t.Fatalf("read %s native frame: %v", kind, e)
				}
				id := binary.BigEndian.Uint32(h[1:5])
				p := make([]byte, int(binary.BigEndian.Uint16(h[5:])))
				if _, e = io.ReadFull(c, p); e != nil {
					t.Fatal(e)
				}
				switch h[0] {
				case 2:
					if id != 1 {
						t.Fatal("wrong stream", id)
					}
					got = append(got, p...)
				case 0, 6, 7, 10:
				case 3, 5:
					t.Fatalf("unexpected frame %d %q", h[0], p)
				default:
					t.Fatalf("unknown frame %d", h[0])
				}
			}
			if !bytes.Equal(got, expected) {
				t.Fatal("payload corruption")
			}
			c.Close()
			user, _ := v.jc_QPN6.Get(9)
			previous := user.GetCtx()
			v.Stop()
			if v.jc_QPN6.Len() != 1 || previous.Err() != context.Canceled || user.GetCtx().Err() != nil || v.zb8f3cHWW8v.Err() != context.Canceled {
				t.Fatal("original stop state")
			}
			if c, e = net.DialTimeout("tcp", addr, 100*time.Millisecond); e == nil {
				c.Close()
				t.Fatal("listener remained open")
			}
			select {
			case e := <-done:
				if e != nil {
					t.Fatal(e)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("destination remained open")
			}
		})
	}
}
