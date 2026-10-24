package CpHCdepV

import (
	"crypto/tls"
	"recovered.local/soga/reality"
	"testing"
	"unsafe"
)

func TestOriginalListenerArgumentAndString(t *testing.T) {
	o := Ap9uD2{Host: "::", Port: 443, Network: "tcp", ProxyProtocol: true}
	if unsafe.Sizeof(o) != 112 || unsafe.Offsetof(o.TLS) != 40 || unsafe.Offsetof(o.Reality) != 48 || unsafe.Offsetof(o.Path) != 56 || unsafe.Offsetof(o.ServiceName) != 72 || unsafe.Offsetof(o.ProxyProtocol) != 88 || unsafe.Offsetof(o.Tunnel) != 96 || unsafe.Offsetof(o.XHTTP) != 104 {
		t.Fatal("original caller layout")
	}
	if o.String() != "(:::443) (tcp + no tls) (proxy protocol: true) (force proxy protocol: false)" {
		t.Fatal(o.String())
	}
	o.Reality = &reality.Config{}
	if o.String() != "(:::443) (tcp + reality) (proxy protocol: true) (force proxy protocol: false)" {
		t.Fatal(o.String())
	}
	o.TLS = &tls.Config{}
	if o.String() != "(:::443) (tcp + tls) (proxy protocol: true) (force proxy protocol: false)" {
		t.Fatal(o.String())
	}
}
