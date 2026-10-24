package CpHCdepV

import (
	E "recovered.local/soga/HxKViEp2a"
	"testing"
	"unsafe"
)

func TestNativeCipherBuffersOwnership(t *testing.T) {
	var b JQi6241
	if unsafe.Sizeof(b) != 24 || unsafe.Offsetof(b.csU3Cr) != 16 {
		t.Fatal("RTTI mismatch")
	}
	read, e := b.GetReadBufferDirect()
	if e != nil {
		t.Fatal(e)
	}
	read.Write([]byte("abcdef"))
	read.Skip(2)
	same, e := b.GetReadBufferDirect()
	if e != nil || same != read || string(same.Bytes()) != "cdef" {
		t.Fatal("direct access must preserve cursors")
	}
	p, e := b.GetReadBuf(2000)
	if e != nil || len(p) != 2000 || len(read.Bytes()) != 0 {
		t.Fatal("read reuse must reset and grow")
	}
	p[0] = 5
	w, e := b.GetWriteBuf(80)
	if e != nil || len(w) != 80 {
		t.Fatal("write allocation")
	}
	w[0] = 7
	if p[0] != 5 {
		t.Fatal("read and write must be independent")
	}
	b.FreeReadBuffer()
	if b.jqW714u != nil || b.tvnW2P == nil {
		t.Fatal("read release")
	}
	if _, e = b.GetReadBuf(5); e != nil {
		t.Fatal("partial release must not mark closed")
	}
	b.Free()
	b.Free()
	if b.jqW714u != nil || b.tvnW2P != nil {
		t.Fatal("released pointers")
	}
	if p, e = b.GetReadBuf(1); p != nil || e != E.ErrorAt5097810 {
		t.Fatal("read closed error identity")
	}
	if _, e = b.GetReadBufferDirect(); e != E.ErrorAt5097810 {
		t.Fatal("direct closed error identity")
	}
	if p, e = b.GetWriteBuf(1); p != nil || e != E.ErrorAt5097810 {
		t.Fatal("write closed error identity")
	}
}
