package inbound

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	F "recovered.local/soga/panel/ssrnames"
	U "recovered.local/soga/user"
	Q "recovered.local/soga/panel/registry"
	"testing"
	"time"
	"unsafe"
)

func ssrConfigFixture(t *testing.T) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "soga.conf")
	text := "type=sspanel-uim\nmu_suffix=example.test\nmu_regex=%0m%id.%suffix\nlisten=127.0.0.1\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\nproxy_protocol=false\nudp_proxy_protocol=false\ndy_limit_enable=false\n"
	if e := os.WriteFile(path, []byte(text), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(path)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
}
func ssrPanelFixture(port int, method, password string) *L.Aga91y {
	return &L.Aga91y{WZ89pTAOrHK4: 42, JAfXWwzYrW: &L.QcqpSL3i6ri{YrjVl_2g_S: port, YyAgnMatmj: password, Z_taAb: F.ZvAxOt1l7Ar(method), H_y2xZht7S8: "origin", ZZIDsJ: "plain"}}
}

// Tests deliberately implement wire encryption independently with crypto/aes.
func ssrSealFixture(method, password string, p []byte) []byte {
	if method == "none" {
		return append([]byte(nil), p...)
	}
	key := md5.Sum([]byte(password))
	block, _ := aes.NewCipher(key[:])
	iv := bytes.Repeat([]byte{0x73}, 16)
	out := append([]byte{}, iv...)
	out = append(out, make([]byte, len(p))...)
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(out[16:], p)
	return out
}
func ssrOpenFixture(t *testing.T, method, password string, p []byte) []byte {
	t.Helper()
	if method == "none" {
		return p
	}
	if len(p) < 16 {
		t.Fatal("short encrypted wire")
	}
	key := md5.Sum([]byte(password))
	block, _ := aes.NewCipher(key[:])
	out := make([]byte, len(p)-16)
	cipher.NewCFBDecrypter(block, p[:16]).XORKeyStream(out, p[16:])
	return out
}
func ssrNewFixture(t *testing.T, method, password string) *Ze0fOaAko {
	t.Helper()
	node := &L.Wr5SVP{EMlXEc: 69}
	any, e := x4oEZV(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{ssrPanelFixture(0, method, password)}}})
	if e != nil {
		t.Fatal(e)
	}
	v := any.(*Ze0fOaAko)
	v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0())) // Original hourly no-request path; no external authorization server accessed.
	t.Cleanup(v.Stop)
	return v
}
func TestNativeSSRConstructorTCPUDPAndLifecycle(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	if unsafe.Sizeof(Ze0fOaAko{}) != 48 || unsafe.Sizeof(xdWl88{}) != 40 || unsafe.Sizeof(pxPRmFCpW6SU{}) != 88 {
		t.Fatal("RTTI layout")
	}
	for _, method := range []string{"none", "aes-128-cfb"} {
		t.Run(method, func(t *testing.T) {
			v := ssrNewFixture(t, method, "native-ssr-password")
			state := v.mc0MXUU.sSKJ79jFt[0]
			if state == nil || len(state.woqAbWjQH_A) != 1 || len(state.aaBpAQ44Zfs) != 1 {
				t.Fatal("constructor must have already opened both listeners")
			}
			before := state
			if e := v.Start(); e != nil || v.mc0MXUU.sSKJ79jFt[0] != before {
				t.Fatal("Start unexpectedly replaced listeners", e)
			}
			for _, network := range []string{"tcp", "udp"} {
				t.Run(network, func(t *testing.T) {
					payload := []byte("native SSR actual " + network + " target")
					response := []byte("target reply " + network)
					done := make(chan error, 1)
					var target net.Addr
					if network == "tcp" {
						listener, e := net.Listen("tcp4", "127.0.0.1:0")
						if e != nil {
							t.Fatal(e)
						}
						defer listener.Close()
						target = listener.Addr()
						go func() {
							c, e := listener.Accept()
							if e == nil {
								defer c.Close()
								c.SetDeadline(time.Now().Add(3 * time.Second))
								buf := make([]byte, len(payload))
								_, e = io.ReadFull(c, buf)
								if e == nil && !bytes.Equal(buf, payload) {
									e = fmt.Errorf("target payload %q", buf)
								}
								if e == nil {
									_, e = c.Write(response)
								}
							}
							done <- e
						}()
					} else {
						p, e := net.ListenPacket("udp4", "127.0.0.1:0")
						if e != nil {
							t.Fatal(e)
						}
						defer p.Close()
						target = p.LocalAddr()
						go func() {
							p.SetDeadline(time.Now().Add(3 * time.Second))
							buf := make([]byte, 2048)
							n, a, e := p.ReadFrom(buf)
							if e == nil && !bytes.Equal(buf[:n], payload) {
								e = fmt.Errorf("target datagram %q", buf[:n])
							}
							if e == nil {
								_, e = p.WriteTo(response, a)
							}
							done <- e
						}()
					}
					_, port, _ := net.SplitHostPort(target.String())
					var portN int
					fmt.Sscan(port, &portN)
					address := ssFixtureAddress(1, portN)
					endpoint := state.woqAbWjQH_A[0].Addr()
					if network == "udp" {
						endpoint = state.aaBpAQ44Zfs[0].LocalAddr()
					}
					client, e := net.Dial(network, endpoint.String())
					if e != nil {
						t.Fatal(e)
					}
					defer client.Close()
					client.SetDeadline(time.Now().Add(4 * time.Second))
					if _, e = client.Write(ssrSealFixture(method, "native-ssr-password", append(append([]byte{}, address...), payload...))); e != nil {
						t.Fatal(e)
					}
					want := response
					if network == "udp" {
						want = append(append([]byte{}, address...), response...)
					}
					size := len(want)
					if method != "none" {
						size += 16
					}
					buf := make([]byte, 2048)
					var n int
					if network == "tcp" {
						n, e = io.ReadFull(client, buf[:size])
					} else {
						n, e = client.Read(buf)
					}
					if e != nil || !bytes.Equal(ssrOpenFixture(t, method, "native-ssr-password", buf[:n]), want) {
						t.Fatal("wire reply", n, e, buf[:n])
					}
					if e = <-done; e != nil {
						t.Fatal(e)
					}
				})
			}
			raw, _ := v.jc_QPN6.Get(42)
			u := raw.(*U.LJJkaD)
			ctx := u.GetCtx()
			oldPortCtx := state.ot1kuknWf35
			v.spcNlUL(L.Sxcf3wu3L{ssrPanelFixture(0, method, "changed")})
			next := v.mc0MXUU.sSKJ79jFt[0]
			if next == nil || next == state || next.gHBoVOaQY != u || u.GetPassword() != "changed" || oldPortCtx.Err() != context.Canceled || ctx.Err() != context.Canceled {
				t.Fatal("password update did not replace port in original order")
			}
			v.spcNlUL(nil)
			if _, ok := v.jc_QPN6.Get(42); ok || len(v.mc0MXUU.sSKJ79jFt) != 0 {
				t.Fatal("removed user retained")
			}
			if _, ok := v.jTyGfNkAw.Get(42); ok {
				t.Fatal("uid map retained")
			}
			if _, ok := v.tpQEL_qKogl.Get(42); ok {
				t.Fatal("password map retained")
			}
		})
	}
}

func TestNativeSSRSinglePortSyncAndBoundaries(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	carrier := ssrPanelFixture(0, "none", "carrier")
	carrier.WZ89pTAOrHK4 = 999
	carrier.JAfXWwzYrW.(*L.QcqpSL3i6ri).BskzckkE7qI = 2
	node := &L.Wr5SVP{EMlXEc: 69, ASJxBbFS: L.Dv6QJyFW6bU{QawbKS2x: true}}
	// Fixture writes the original private []user field at node offset432. This is
	// test-only access to real RTTI, not a fabricated production state interface.
	*(*[]*L.Aga91y)(unsafe.Add(unsafe.Pointer(node), 432)) = []*L.Aga91y{carrier}
	input := ssrPanelFixture(5001, "none", "first")
	any, e := x4oEZV(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{input}}})
	if e != nil {
		t.Fatal(e)
	}
	v := any.(*Ze0fOaAko)
	v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
	defer v.Stop()
	if len(v.mc0MXUU.sSKJ79jFt) != 1 || v.mc0MXUU.sSKJ79jFt[0] == nil {
		t.Fatal("carrier port not created")
	}
	if _, ok := v.jc_QPN6.Get(999); ok {
		t.Fatal("carrier incorrectly inserted into ordinary user set")
	}
	raw, ok := v.jc_QPN6.Get(42)
	if !ok {
		t.Fatal("ordinary user missing")
	}
	u := raw.(*U.LJJkaD)
	oldctx := u.GetCtx()
	host, _ := u.GetSSRHost()
	if x, ok := v.ktET5DFha.Get(host); !ok || x != u {
		t.Fatal("host map")
	}
	update := ssrPanelFixture(5001, "none", "second")
	v.spcNlUL(L.Sxcf3wu3L{update})
	if password, ok := v.tpQEL_qKogl.Get(42); !ok || password != "second" || oldctx.Err() != context.Canceled {
		t.Fatal("singleport credential update")
	}
	if got, ok := v.jTyGfNkAw.Get(42); !ok || got != u {
		t.Fatal("singleport UID map identity")
	}
	users := make(L.Sxcf3wu3L, 90)
	for i := range users {
		p := ssrPanelFixture(5001+i, "none", fmt.Sprint(i))
		p.WZ89pTAOrHK4 = int64(i + 1)
		users[i] = p
	}
	limited := C.Q0SDIN()
	v.spcNlUL(users)
	want := 90
	if limited {
		want = 88
	}
	count := 0
	v.jc_QPN6.Range(func(_ int64, _ U.PgaNd3014U) bool { count++; return true })
	if count != want || v.fUxO7UYAH != want || userCountAt50C6DA0 != want {
		t.Fatal("original user bound", count, v.fUxO7UYAH, userCountAt50C6DA0, want)
	}
	v.spcNlUL(nil)
	count = 0
	v.jc_QPN6.Range(func(_ int64, _ U.PgaNd3014U) bool { count++; return true })
	if count != 0 || len(v.mc0MXUU.sSKJ79jFt) != 1 {
		t.Fatal("singleport removal must leave carrier listener")
	}
	state := v.mc0MXUU.sSKJ79jFt[0]
	v.Stop()
	if state.ot1kuknWf35.Err() != context.Canceled || v.mc0MXUU.sSKJ79jFt[0] != state {
		t.Fatal("StopAll must cancel but retain map entries")
	}
}
func TestNativeSSRUDPBindFailureRollsBackTCP(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	occupied, e := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if e != nil {
		t.Fatal(e)
	}
	defer occupied.Close()
	port := occupied.LocalAddr().(*net.UDPAddr).Port
	node := &L.Wr5SVP{EMlXEc: 69}
	any, e := x4oEZV(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{ssrPanelFixture(port, "none", "bind")}}})
	if e != nil {
		t.Fatal("original constructor logs per-user failures rather than returning them", e)
	}
	v := any.(*Ze0fOaAko)
	defer v.Stop()
	if len(v.mc0MXUU.sSKJ79jFt) != 0 {
		t.Fatal("published partial listener")
	}
	if _, ok := v.jc_QPN6.Get(42); ok {
		t.Fatal("failed user not removed")
	}
	listener, e := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: port})
	if e != nil {
		t.Fatal("leaked TCP listener", e)
	}
	listener.Close()
	if e = v.mc0MXUU.hmycHk6ebPwS(port); e == nil || e.Error() != fmt.Sprintf("can not find ssr port: %v", port) {
		t.Fatal(e)
	}
	if e = v.mc0MXUU.Update(&L.Aga91y{WZ89pTAOrHK4: 73}); e == nil || e.Error() != "user 73 has no ssr credential" {
		t.Fatal(e)
	}
}
