package HxKViEp2a

import (
	"net"
	"testing"
)

type freeProbe struct {
	net.Conn
	frees, unwraps int
	next           net.Conn
}

func (p *freeProbe) Free()                { p.frees++ }
func (p *freeProbe) GetRawConn() net.Conn { p.unwraps++; return p.next }

type unwrapProbe struct {
	net.Conn
	next    net.Conn
	unwraps int
}

func (p *unwrapProbe) GetRawConn() net.Conn { p.unwraps++; return p.next }

type valueFreeProbe struct{}

func (valueFreeProbe) Free() {}

func TestOriginalFreeDispatch(t *testing.T) {
	var absent *freeProbe
	U9tk4Xosh9(absent)
	U9tk4Xosh9(nil)
	leaf := &freeProbe{}
	inner := &freeProbe{next: leaf}
	outer := &unwrapProbe{next: inner}
	Li3c4CJ1_D(outer)
	if outer.unwraps != 1 || inner.frees != 1 || inner.unwraps != 0 || leaf.frees != 0 {
		t.Fatalf("wrong Free precedence: %+v %+v %+v", outer, inner, leaf)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("original non-nilable value must panic in IsNil")
		}
	}()
	U9tk4Xosh9(valueFreeProbe{})
}
