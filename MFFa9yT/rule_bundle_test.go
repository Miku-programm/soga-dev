package MFFa9yT

import (
	"net"
	"recovered.local/soga/xx5fRyrFw"
	"reflect"
	"testing"
)

func TestOriginalMixedRuleConstruction(t *testing.T) {
	typ := reflect.TypeFor[rHfs_oNXEJ]()
	if typ.Size() != 40 || typ.Field(1).Offset != 24 || typ.Field(2).Offset != 32 {
		t.Fatal(typ)
	}
	r, nd, ni, np := elCwduDdyGl(xx5fRyrFw.VHWq49iRat{" ip:192.0.2.0/24 ", "full:one.test", "port:443", "full:two.test", " "})
	if nd != 2 || ni != 1 || np != 1 {
		t.Fatal(nd, ni, np)
	}
	if len(r.a73gMUVz) != 2 || r.a73gMUVz[0].GetId() != 0 || r.a73gMUVz[1].GetId() != 1 {
		t.Fatal(r.a73gMUVz)
	}
	if _, ok := r.a73gMUVz.Test("one.test"); !ok {
		t.Fatal("domain rule missing")
	}
	if !r.fj_WUx4re.TestIP(net.ParseIP("192.0.2.30")) {
		t.Fatal("IP rule missing")
	}
	if !r.xDjyw2UC.Match(443) {
		t.Fatal("port rule missing")
	}
}
