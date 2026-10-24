package vtQf23A9G

import (
	"bytes"
	"errors"
	"io"
	"net"
	C "recovered.local/soga/CAodbhZf"
	H "recovered.local/soga/HHGbXWqYaYt_"
	"testing"
	"time"
)

type recordConn struct {
	writes    [][]byte
	deadlines []time.Time
	closed    bool
	failAt    int
	short     bool
}

func (c *recordConn) Read([]byte) (int, error) { return 0, io.EOF }
func (c *recordConn) Write(p []byte) (int, error) {
	c.writes = append(c.writes, bytes.Clone(p))
	if c.failAt > 0 && len(c.writes) == c.failAt {
		return 1, io.ErrUnexpectedEOF
	}
	if c.short {
		return 1, nil
	}
	return len(p), nil
}
func (c *recordConn) Close() error                    { c.closed = true; return nil }
func (c *recordConn) LocalAddr() net.Addr             { return nil }
func (c *recordConn) RemoteAddr() net.Addr            { return nil }
func (c *recordConn) SetDeadline(time.Time) error     { return nil }
func (c *recordConn) SetReadDeadline(time.Time) error { return nil }
func (c *recordConn) SetWriteDeadline(d time.Time) error {
	c.deadlines = append(c.deadlines, d)
	return nil
}
func paddedFixture(spec string) (*BIkgBMz3M, *recordConn) {
	c := &recordConn{}
	var v C.EahpvU_iseQ[*H.FZBbzkJPWbrk]
	v.Store(H.R_GY_FGx([]byte(spec)))
	s := NewServerFromInline(c, nil, &v)
	s.rVguB66 = true
	return s, c
}
func TestPaddingExactWireRecordsAndOriginalShortWrites(t *testing.T) {
	s, c := paddedFixture("stop=3\n1=4-4,20-20,5-5\n2=11-11,c,20-20")
	n, e := s.ucYLRp([]byte("0123456789"), time.Second)
	if e != nil || n != 10 {
		t.Fatalf("%d %v", n, e)
	}
	want := [][]byte{[]byte("0123"), append([]byte("456789"), wireFrame(0, 0, make([]byte, 7))...), wireFrame(0, 0, make([]byte, 5))}
	if len(c.writes) != len(want) {
		t.Fatal(len(c.writes))
	}
	for i := range want {
		if !bytes.Equal(c.writes[i], want[i]) {
			t.Fatalf("record%d=%x want%x", i, c.writes[i], want[i])
		}
	}
	if len(c.deadlines) != 2 || c.deadlines[0].IsZero() || !c.deadlines[1].IsZero() {
		t.Fatal(c.deadlines)
	}
	n, e = s.ucYLRp([]byte("abcd"), 0)
	if e != nil || n != 4 || len(c.writes) != 4 || string(c.writes[3]) != "abcd" {
		t.Fatalf("gap=7 must not append pad: n%d e%v records%v", n, e, c.writes)
	}
	s.ucYLRp([]byte("stop"), 0)
	if s.rVguB66 || s.wLYsfs.Load() != 3 || string(c.writes[4]) != "stop" {
		t.Fatal("stop boundary")
	}
	s, c = paddedFixture("stop=2\n1=4-4,20-20")
	c.short = true
	n, e = s.ucYLRp([]byte("0123456789"), 0)
	if n != 10 || e != nil {
		t.Fatalf("original ignores intermediate short writes: %d %v", n, e)
	}
	s, c = paddedFixture("stop=2\n1=4-4,20-20")
	c.failAt = 2
	n, e = s.ucYLRp([]byte("0123456789"), 0)
	if n != 0 || !errors.Is(e, io.ErrUnexpectedEOF) {
		t.Fatalf("original error resets count: %d %v", n, e)
	}
}
func TestFrameLengthWrapAndControlErrorClose(t *testing.T) {
	s, c := paddedFixture("stop=1")
	p := bytes.Repeat([]byte{0xa3}, 65537)
	n, e := s.hOmlzcDoI(0x12345678, p)
	if n != len(p) || e != nil {
		t.Fatal(n, e)
	}
	if len(c.writes) != 1 || !bytes.Equal(c.writes[0], wireFrame(2, 0x12345678, p)) {
		t.Fatal("length truncation frame differs")
	}
	s, c = paddedFixture("stop=1")
	c.failAt = 1
	if _, e := s.hOmlzcDoI(1, []byte{1}); e == nil || c.closed {
		t.Fatal("data write error must not close session")
	}
	s, c = paddedFixture("stop=1")
	c.failAt = 1
	if _, e := s.ehs06TiU(8, 1, nil); e == nil || !c.closed || !s.IsClosed() {
		t.Fatal("control write error must close session")
	}
}

func TestBufferedConcatOriginalNilAndCopy(t *testing.T) {
	s, c := paddedFixture("stop=1")
	s.d0Ze8qja = true
	if n, e := s.ucYLRp(nil, 0); n != 0 || e != nil || s.vV_K36 != nil {
		t.Fatal("original zero-size concat must remain nil")
	}
	p := []byte("first")
	s.ucYLRp(p, 0)
	p[0] = 'X'
	if string(s.vV_K36) != "first" {
		t.Fatal("buffer retained caller data")
	}
	s.d0Ze8qja = false
	s.ucYLRp([]byte("next"), 0)
	if s.vV_K36 != nil || len(c.writes) != 1 || string(c.writes[0]) != "firstnext" {
		t.Fatal("flush/clear mismatch")
	}
}
