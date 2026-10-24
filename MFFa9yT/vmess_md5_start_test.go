package MFFa9yT

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	U "recovered.local/soga/nzATvajt"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"testing"
	"time"
)

// Independent HMAC-MD5/AES-CFB client wire traverses the real native listener,
// authentication window, header parser, route and TCP/UDP target. This does not
// install an authentication replacement or change the sample's state limits.
func TestNativeVMessMD5StartRealDestinations(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	for _, mode := range []string{"md5", "auto"} {
		// Force the former collision preconditions instead of relying on random
		// ephemeral ports: one AuthID timestamp and the same TCP/UDP target port.
		// Only the client's fresh request IV/key distinguishes the replay input.
		stamp := time.Now().Unix()
		pairedPort := 0
		for _, kind := range []string{"tcp", "udp"} {
			t.Run(mode+"/"+kind, func(t *testing.T) {
				path := filepath.Join(t.TempDir(), "soga.conf")
				if e := os.WriteFile(path, []byte("listen=127.0.0.1\nvmess_auth_mode="+mode+"\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\nproxy_protocol=false\nudp_proxy_protocol=false\n"), 0600); e != nil {
					t.Fatal(e)
				}
				C.MkwNWciPHS(path)
				if e := C.TvBECnD9fM6(); e != nil {
					t.Fatal(e)
				}
				panel := &Q.D5gTIZ5aVg{H47AzZsYHB1N: &L.Wr5SVP{EMlXEc: 31, SwxJNgPpg: L.JW0CeHo3O{F2GXyZZ: "tcp"}}, VTme1SuHL6i_: L.Sxcf3wu3L{vmessPanelFixture(7, "11111111-1111-4111-8111-111111111111")}}
				core, e := ag90n_rZjJO(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: panel})
				if e != nil {
					t.Fatal(e)
				}
				v := core.(*K4avFiYOPi)
				defer v.Stop()
				if v.dj_cTnaUp6f == nil || (v.p54XvyV == nil) != (mode == "md5") || (v.qgz9PfTS == nil) != (mode == "md5") {
					t.Fatal("original authentication-mode constructor branches")
				}
				v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
				if e = v.Start(); e != nil {
					t.Fatal(e)
				}
				u0, _ := v.jc_QPN6.Get(7)
				u := u0.(*U.QKaIWV2)
				id := legacyVMessAuthID(u, stamp)
				until := time.Now().Add(2 * time.Second)
				for {
					if _, _, ok := v.dj_cTnaUp6f.Get(id); ok {
						break
					}
					if time.Now().After(until) {
						t.Fatal("original MD5 window did not populate")
					}
					time.Sleep(time.Millisecond)
				}
				done := make(chan error, 1)
				var target net.Addr
				if kind == "tcp" {
					l, e := net.Listen("tcp4", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					defer l.Close()
					target = l.Addr()
					pairedPort = l.Addr().(*net.TCPAddr).Port
					go func() {
						c, e := l.Accept()
						if e == nil {
							defer c.Close()
							c.SetDeadline(time.Now().Add(5 * time.Second))
							_, e = io.Copy(c, c)
						}
						done <- e
					}()
				} else {
					pc, e := net.ListenPacket("udp4", fmt.Sprintf("127.0.0.1:%d", pairedPort))
					if e != nil {
						t.Fatal(e)
					}
					defer pc.Close()
					target = pc.LocalAddr()
					go func() {
						pc.SetDeadline(time.Now().Add(5 * time.Second))
						buf := make([]byte, 4096)
						n, a, e := pc.ReadFrom(buf)
						if e == nil {
							_, e = pc.WriteTo(buf[:n], a)
						}
						done <- e
					}()
				}
				var port int
				switch a := target.(type) {
				case *net.TCPAddr:
					port = a.Port
				case *net.UDPAddr:
					port = a.Port
				default:
					t.Fatal(fmt.Sprintf("%T", a))
				}
				addr := v.l8Uq4M[0].Addr().String()
				c, e := net.DialTimeout("tcp", addr, time.Second)
				if e != nil {
					t.Fatal(e)
				}
				defer c.Close()
				c.SetDeadline(time.Now().Add(5 * time.Second))
				exerciseNativeVMessWireMode(t, c, u, port, kind, false, stamp)
				c.Close()
				v.Stop()
				if v.jc_QPN6.Len() != 0 || v.zb8f3cHWW8v.Err() != context.Canceled {
					t.Fatal("Stop retained users or context")
				}
				if c, e := net.DialTimeout("tcp", addr, 100*time.Millisecond); e == nil {
					c.Close()
					t.Fatal("listener remained open")
				}
				select {
				case e := <-done:
					if e != nil {
						t.Fatal(e)
					}
				case <-time.After(5 * time.Second):
					t.Fatal("target did not finish")
				}
			})
		}
	}
}
