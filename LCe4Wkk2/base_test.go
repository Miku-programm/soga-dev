package LCe4Wkk2

import (
	"context"
	"net"
	E "recovered.local/soga/E3F12urwswL"
	O "recovered.local/soga/vP_e5OQAve"
	"reflect"
	"testing"
	"unsafe"
)

type chainAdapter struct{ *B8EHsVY7a }

func (*chainAdapter) DialContext(context.Context, *E.D0mLgQF, ...O.PNSpXRb) (E.PnyceLR, error) {
	panic("not exercised")
}
func TestOriginalBaseAndChains(t *testing.T) {
	b := &B8EHsVY7a{m50I26ZDL: "named", x3HUaU8Ln1S1: "addr", iOKxDl: 5, ntseIWz: true}
	if b.Name() != "named" || b.Addr() != "addr" || b.Type() != 5 || !b.SupportUDP() || b.Unwrap(nil) != nil {
		t.Fatal("base fields")
	}
	data, err := b.MarshalJSON()
	if err != nil || string(data) != `{"type":"Socks5"}` {
		t.Fatal(string(data), err)
	}
	a, z := net.Pipe()
	defer a.Close()
	defer z.Close()
	c, err := b.StreamConn(a, nil)
	if c != a || err == nil || err.Error() != "no support" {
		t.Fatal(c, err)
	}
	p, err := b.ListenPacketContext(context.Background(), nil)
	if p != nil || err == nil || err.Error() != "no support" {
		t.Fatal(p, err)
	}
	wrapped := &lGta_1QkvQ{Conn: a, bnOtGuS: E.GYe3KD{"first"}}
	wrapped.AppendToChains(&chainAdapter{b})
	if !reflect.DeepEqual(wrapped.Chains(), E.GYe3KD{"first", "named"}) {
		t.Fatal(wrapped.Chains())
	}
	wrapped.Chains()[0] = "mutated"
	if wrapped.bnOtGuS[0] != "mutated" {
		t.Fatal("copied chains")
	}
	packet := new(uMtOz6o)
	packet.AppendToChains(&chainAdapter{b})
	if packet.Chains()[0] != "named" {
		t.Fatal(packet.Chains())
	}
	if unsafe.Sizeof(*b) != 72 || unsafe.Sizeof(*wrapped) != 40 || unsafe.Sizeof(*packet) != 40 {
		t.Fatal("RTTI layout")
	}
}
func TestSocksConstructorLayoutAndTLS(t *testing.T) {
	o := SocksOptions{FFrN8gdu: FFrN8gdu{"eth9", 123}, Name: "node", Server: "2001:db8::1", Port: -1, Username: "u", Password: "p", TLS: true, UDP: true, SkipCertVerify: true}
	a := Auqn90u2NBQ(o)
	if a.Addr() != "[2001:db8::1]:-1" || a.Name() != "node" || a.Type() != 5 || !a.SupportUDP() {
		t.Fatal(a)
	}
	if a.gmXALhCpyKF != "u" || a.eOL60dT1 != "p" || a.xDan8hA5LV != "eth9" || a.bEbqkM8 != 123 {
		t.Fatal(a)
	}
	if a.idTOPJOmPpB == nil || a.idTOPJOmPpB.ServerName != o.Server || !a.idTOPJOmPpB.InsecureSkipVerify {
		t.Fatal(a.idTOPJOmPpB)
	}
	if a.idTOPJOmPpB.MinVersion != 0 || a.idTOPJOmPpB.RootCAs != nil {
		t.Fatal("unexpected TLS defaults")
	}
	o.TLS = false
	a = Auqn90u2NBQ(o)
	if a.idTOPJOmPpB != nil || a.m0JWRLb2CH || !a.p7gaqRDae {
		t.Fatal(a)
	}
	if unsafe.Sizeof(o) != 104 || unsafe.Offsetof(o.TLS) != 96 || unsafe.Offsetof(o.SkipCertVerify) != 98 || unsafe.Sizeof(*a) != 56 {
		t.Fatal("RTTI/ABI layout")
	}
}
func TestAdapterOptionAppend(t *testing.T) {
	b := new(B8EHsVY7a)
	if b.DialOptions(nil) != nil {
		t.Fatal("nil changed")
	}
	o := []O.PNSpXRb{O.VcC742("caller")}
	b.xDan8hA5LV = "adapter"
	b.bEbqkM8 = -1
	got := b.DialOptions(o)
	if len(got) != 3 || len(o) != 1 {
		t.Fatal(len(got), len(o))
	}
}
