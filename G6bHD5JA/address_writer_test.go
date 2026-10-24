package G6bHD5JA

import (
	"bytes"
	"net"
	"recovered.local/soga/D_zCag1qTV"
	"strings"
	"testing"
)

func TestPortFirstAddressEncoding(t *testing.T) {
	for _, s := range []string{"1.2.3.4:53", "[::1]:443", "example.org:80"} {
		a, _ := U80FDt(s)
		b := D_zCag1qTV.S4OQXWIZ(32)
		n, err := HfFNot(b, a)
		if err != nil || n != b.Len() {
			t.Fatal(n, err)
		}
		out, err := CHYbTjleKL(bytes.NewReader(b.Bytes()))
		if err != nil || out.Join() != s {
			t.Fatal(out, err)
		}
		b.Free()
	}
	b := D_zCag1qTV.S4OQXWIZ(2)
	defer b.Free()
	a := &AHU2MRmjzQu{dlcy_1: strings.Repeat("a", 256), pDnAz_: 123}
	if n, e := HfFNot(b, a); n != 0 || e == nil || e.Error() != "host too long: "+a.dlcy_1 || !bytes.Equal(b.Bytes(), []byte{0, 123}) {
		t.Fatal(n, e, b.Bytes())
	}
	b.Reset()
	a.oKcMaQDwbH = []net.IP{nil}
	n, e := HfFNot(b, a)
	if n != 19 || e != nil || !bytes.Equal(b.Bytes(), []byte{0, 123, 3}) {
		t.Fatal("original invalid IP accounting", n, e, b.Bytes())
	}
}
