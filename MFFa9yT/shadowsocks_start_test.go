package MFFa9yT

import (
	"context"
	"crypto/rand"
	"net"
	"os"
	"path/filepath"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	U "recovered.local/soga/nzATvajt"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"testing"
	"time"
)

func ssStartFixture(t *testing.T, port int) *Lp07FzBLx9V {
	t.Helper()
	node := &L.Wr5SVP{EMlXEc: 39, FnMUny: L.AAMPrdcYTL6{Am3QoeF: port, HMkc4BgR: "aes-128-gcm"}}
	i, e := aXSbxcDLdeEY(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{ssPanelFixture(7, "udp-native-start")}}})
	if e != nil {
		t.Fatal(e)
	}
	v := i.(*Lp07FzBLx9V)
	t.Cleanup(v.Stop)
	return v
}
func TestNativeShadowsocksStartUDPAndFailureCleanup(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	p := filepath.Join(t.TempDir(), "soga.conf")
	os.WriteFile(p, []byte("listen=127.0.0.1\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\nproxy_protocol=false\nudp_proxy_protocol=false\ndy_limit_enable=false\n"), 0600)
	C.MkwNWciPHS(p)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	t.Run("native-udp", func(t *testing.T) {
		v := ssStartFixture(t, 0)
		if e := v.Start(); e != nil {
			t.Fatal(e)
		}
		if len(v.vLRjbS7NDeZZ) != 1 || len(v.sSqGyaLW) != 1 {
			t.Fatal("native transport setup")
		}
		user, _ := v.jc_QPN6.Get(7)
		u := user.(*U.WtK9HQSy66)
		target, e := net.ListenPacket("udp4", "127.0.0.1:0")
		if e != nil {
			t.Fatal(e)
		}
		defer target.Close()
		echoed := make(chan error, 1)
		go func() {
			target.SetDeadline(time.Now().Add(3 * time.Second))
			p := make([]byte, 4096)
			n, a, e := target.ReadFrom(p)
			if e == nil {
				_, e = target.WriteTo(p[:n], a)
			}
			echoed <- e
		}()
		client, e := net.Dial("udp", v.sSqGyaLW[0].LocalAddr().String())
		if e != nil {
			t.Fatal(e)
		}
		defer client.Close()
		client.SetDeadline(time.Now().Add(3 * time.Second))
		plain := append(ssFixtureAddress(1, target.LocalAddr().(*net.UDPAddr).Port), []byte("original UDP wire")...)
		salt := make([]byte, 16)
		rand.Read(salt)
		a, e := u.NewAEAD(salt)
		if e != nil {
			t.Fatal(e)
		}
		if _, e = client.Write(append(salt, a.Seal(nil, make([]byte, 12), plain, nil)...)); e != nil {
			t.Fatal(e)
		}
		p := make([]byte, 4096)
		n, e := client.Read(p)
		if e != nil {
			t.Fatal(e)
		}
		if n < 32 {
			t.Fatal("short reply")
		}
		a, e = u.NewAEAD(p[:16])
		if e != nil {
			t.Fatal(e)
		}
		decoded, e := a.Open(nil, make([]byte, 12), p[16:n], nil)
		if e != nil || string(decoded) != string(plain) {
			t.Fatal(e, string(decoded))
		}
		if e = <-echoed; e != nil {
			t.Fatal(e)
		}
	})
	t.Run("udp-bind-failure-closes-tcp", func(t *testing.T) {
		occupied, e := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
		if e != nil {
			t.Fatal(e)
		}
		defer occupied.Close()
		port := occupied.LocalAddr().(*net.UDPAddr).Port
		v := ssStartFixture(t, port)
		if e = v.Start(); e == nil {
			t.Fatal("expected UDP bind failure")
		}
		if len(v.vLRjbS7NDeZZ) != 0 || len(v.sSqGyaLW) != 0 {
			t.Fatal("partial state published")
		}
		listener, e := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: port})
		if e != nil {
			t.Fatal("TCP leaked after UDP failure:", e)
		}
		listener.Close()
	})
}
