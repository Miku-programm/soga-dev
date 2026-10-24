package MFFa9yT

import (
	"bytes"
	"net"
	"testing"
	"time"

	L "recovered.local/soga/LFWTTXF7JR"
	V "recovered.local/soga/gk1z9v"
	U "recovered.local/soga/nzATvajt"
)

// Reproduce the old fixture's collision deterministically: ports differing by
// 256 generated identical request IV/key. MD5 auth @0x1da3100 includes the
// received AuthID in B2pr6gMaPO, so the same user/timestamp produces exactly the
// same 48-byte replay input at uaQJpenJx8c @0x1d179c0 despite TCP/UDP and port
// changes. The real auth manager and handshake must continue rejecting it.
func TestVMessMD5FixturePortCollisionAndOriginalReplay(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := newVmessAuthFixture(t)
	v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(7, "11111111-1111-4111-8111-111111111111")})
	base, _ := v.jc_QPN6.Get(7)
	u := base.(*U.QKaIWV2)
	v.dj_cTnaUp6f = V.NewMD5AuthManager(v.zb8f3cHWW8v, v.jc_QPN6, true)
	v.dj_cTnaUp6f.StartTask()
	stamp := time.Now().Unix()
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
	check := func(wire []byte, repeated bool) {
		t.Helper()
		auth := v.iUB_i9oBP_68([16]byte(wire[:16]))
		if auth == nil || auth.Ih7r8HyMvOa || auth.B2pr6gMaPO != id {
			t.Fatal("original MD5 authentication/replay identity changed")
		}
		server, client := net.Pipe()
		server.SetDeadline(time.Now().Add(time.Second))
		client.SetDeadline(time.Now().Add(time.Second))
		done := make(chan struct{})
		go func() { defer close(done); client.Write(wire[16:]); client.Close() }()
		framed, err := V.EAQVlR(server, auth)
		if framed != nil {
			framed.Free()
		}
		server.Close()
		client.Close()
		<-done
		if repeated {
			if err == nil || err.Error() != "vmess salt repeat" {
				t.Fatalf("same-second different-port TCP→UDP replay: %v", err)
			}
		} else if err != nil {
			t.Fatalf("fresh authenticated header: %v", err)
		}
	}
	ports := []int{35001, 35257}
	kinds := []string{"tcp", "udp"}
	var old [2][32]byte
	for i, port := range ports {
		for n := range old[i] {
			old[i][n] = byte(n+1) + byte(port)
		}
	}
	if old[0] != old[1] {
		t.Fatal("old fixture collision no longer reproduced")
	}
	for i, port := range ports {
		wire, _, _ := nativeVMessFixtureWireKeys(t, u, port, kinds[i], false, stamp, old[i])
		check(wire, i == 1)
	}
	a, ca, _ := nativeVMessFixtureWire(t, u, ports[0], kinds[0], false, stamp)
	b, cb, _ := nativeVMessFixtureWire(t, u, ports[1], kinds[1], false, stamp)
	if bytes.Equal(ca[1:33], cb[1:33]) || bytes.Equal(ca[1:33], old[0][:]) || bytes.Equal(cb[1:33], old[1][:]) {
		t.Fatal("normal fixture reused request IV/key")
	}
	check(a, false)
	check(b, false)
	t.Log("old port-derived pair: TCP accepted, same-second UDP rejected as vmess salt repeat; fresh independent pair: both accepted")
}
