package MFFa9yT

import (
	"context"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/CAodbhZf"
	J "recovered.local/soga/J1P4KNvfI_J"
	L "recovered.local/soga/LFWTTXF7JR"
	N "recovered.local/soga/NtxdAE"
	C "recovered.local/soga/cNT24rG4F"
	U "recovered.local/soga/nzATvajt"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"strings"
	"testing"
	"time"
)

func ssClientNonce(p []byte) {
	for i := range p {
		p[i]++
		if p[i] != 0 {
			return
		}
	}
}
func ssClientFrame(a cipher.AEAD, nonce, p []byte) []byte {
	b := binary.BigEndian.AppendUint16(nil, uint16(len(p)))
	out := a.Seal(nil, nonce, b, nil)
	ssClientNonce(nonce)
	out = append(out, a.Seal(nil, nonce, p, nil)...)
	ssClientNonce(nonce)
	return out
}
func ssReadClientFrame(r io.Reader, a cipher.AEAD, nonce []byte) ([]byte, error) {
	b := make([]byte, 2+a.Overhead())
	if _, e := io.ReadFull(r, b); e != nil {
		return nil, e
	}
	b, e := a.Open(nil, nonce, b, nil)
	ssClientNonce(nonce)
	if e != nil {
		return nil, e
	}
	out := make([]byte, int(binary.BigEndian.Uint16(b))+a.Overhead())
	if _, e = io.ReadFull(r, out); e != nil {
		return nil, e
	}
	out, e = a.Open(nil, nonce, out, nil)
	ssClientNonce(nonce)
	return out, e
}
func ssFixtureAddress(tag byte, port int) []byte {
	return binary.BigEndian.AppendUint16([]byte{tag, 127, 0, 0, 1}, uint16(port))
}

// The client encodes real native SS frames; the remote socket is an actual TCP
// or UDP echo target. This tests the original authentication, policy and relay,
// including the three UDP-over-TCP framing variants, without a panel endpoint.
func TestNativeShadowsocksHandlerRealDestinations(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	path := filepath.Join(t.TempDir(), "soga.conf")
	if e := os.WriteFile(path, []byte("dy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\nlisten=127.0.0.1\n"), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(path)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	for _, mode := range []string{"tcp", "v1", "v2-connected", "v2-unconnected", "tcp-start", "v2-connected-start"} {
		t.Run(mode, func(t *testing.T) {
			useStart := strings.HasSuffix(mode, "-start")
			mode = strings.TrimSuffix(mode, "-start")
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			b, e := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: ctx})
			if e != nil {
				t.Fatal(e)
			}
			defer b.yPUbKO()
			v := &Lp07FzBLx9V{JWN154qJ4: b, wZhn2Y: J.YCcacra(ctx), gtKatlWm0: N.Azb3Zr(ctx, &C.Nl62Lnd_{})}
			v.wiy6P3.Store(&L.Wr5SVP{EMlXEc: 31, FnMUny: L.AAMPrdcYTL6{HMkc4BgR: "aes-128-gcm"}})
			if useStart {
				inbound, e := aXSbxcDLdeEY(&DV_XPJ4z0{Dmpl48sXM: ctx, HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: v.wiy6P3.Load(), VTme1SuHL6i_: L.Sxcf3wu3L{ssPanelFixture(7, "original-handler-wire")}}})
				if e != nil {
					t.Fatal(e)
				}
				v = inbound.(*Lp07FzBLx9V)
				defer v.Stop()
			} else {
				v.spcNlUL(L.Sxcf3wu3L{ssPanelFixture(7, "original-handler-wire")})
			}
			base, _ := v.jc_QPN6.Get(7)
			user := base.(*U.WtK9HQSy66)
			// Original hourly-check state, deliberately within its no-check interval.
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			var port int
			echo := make(chan error, 1)
			if mode == "tcp" {
				ln, e := net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer ln.Close()
				port = ln.Addr().(*net.TCPAddr).Port
				go func() {
					c, e := ln.Accept()
					if e != nil {
						echo <- e
						return
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(5 * time.Second))
					p := make([]byte, 17)
					_, e = io.ReadFull(c, p)
					if e == nil {
						_, e = c.Write(p)
					}
					echo <- e
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
					p := make([]byte, 4096)
					n, a, e := pc.ReadFrom(p)
					if e == nil {
						_, e = pc.WriteTo(p[:n], a)
					}
					echo <- e
				}()
			}
			var ln net.Listener
			var done chan struct{}
			if useStart {
				if e := v.Start(); e != nil {
					t.Fatal(e)
				}
				if len(v.vLRjbS7NDeZZ) != 1 || len(v.sSqGyaLW) != 1 {
					t.Fatal("native listener count")
				}
				ln = v.vLRjbS7NDeZZ[0]
			} else {
				var e error
				ln, e = net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer ln.Close()
				done = make(chan struct{})
				go func() {
					defer close(done)
					c, e := ln.Accept()
					if e != nil {
						return
					}
					v.nlJ8YjzjZ(ctx, c)
				}()
			}
			client, e := net.Dial("tcp", ln.Addr().String())
			if e != nil {
				t.Fatal(e)
			}
			defer client.Close()
			client.SetDeadline(time.Now().Add(5 * time.Second))
			payload := []byte("native ss payload")
			var plain []byte
			if mode == "tcp" {
				plain = append(ssFixtureAddress(1, port), payload...)
			} else {
				magic := "sp.v2.udp-over-tcp.arpa"
				if mode == "v1" {
					magic = "sp.udp-over-tcp.arpa"
				}
				plain = append([]byte{3, byte(len(magic))}, []byte(magic)...)
				plain = append(plain, 0, 0)
				if mode == "v1" {
					plain = append(plain, ssFixtureAddress(0, port)...)
				} else {
					flag := byte(0)
					if mode == "v2-connected" {
						flag = 1
					}
					plain = append(plain, flag)
					plain = append(plain, ssFixtureAddress(1, port)...)
					if mode == "v2-unconnected" {
						plain = append(plain, ssFixtureAddress(0, port)...)
					}
				}
				plain = binary.BigEndian.AppendUint16(plain, uint16(len(payload)))
				plain = append(plain, payload...)
			}
			salt := make([]byte, 16)
			rand.Read(salt)
			a, e := user.NewAEAD(salt)
			if e != nil {
				t.Fatal(e)
			}
			nonce := make([]byte, 12)
			if _, e = client.Write(append(salt, ssClientFrame(a, nonce, plain)...)); e != nil {
				t.Fatal(e)
			}
			serverSalt := make([]byte, 16)
			if _, e = io.ReadFull(client, serverSalt); e != nil {
				t.Fatal(e)
			}
			a, e = user.NewAEAD(serverSalt)
			if e != nil {
				t.Fatal(e)
			}
			clear(nonce)
			want := payload
			if mode != "tcp" {
				want = binary.BigEndian.AppendUint16(nil, uint16(len(payload)))
				want = append(want, payload...)
				if mode != "v2-connected" {
					want = append(ssFixtureAddress(0, port), want...)
				}
			}
			var got []byte
			for len(got) < len(want) {
				p, e := ssReadClientFrame(client, a, nonce)
				if e != nil {
					t.Fatal(e)
				}
				got = append(got, p...)
			}
			if string(got) != string(want) {
				t.Fatalf("wire %x want %x", got, want)
			}
			if e = <-echo; e != nil {
				t.Fatal(e)
			}
			client.Close()
			if done != nil {
				select {
				case <-done:
				case <-time.After(3 * time.Second):
					t.Fatal("handler did not close")
				}
			}
		})
	}
}
