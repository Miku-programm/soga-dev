package VasCJz

import (
	"bytes"
	"container/heap"
	"errors"
	"io"
	"testing"
	"unsafe"
)

func queueFixture(n int) *y0bchQKJ {
	return &y0bchQKJ{dL_ULZL: make(chan UmYbKhCpX, n), fAEWeMURGLm0: n}
}
func TestOriginalUploadQueueReordersAndSplits(t *testing.T) {
	if unsafe.Sizeof(y0bchQKJ{}) != 88 || unsafe.Offsetof(y0bchQKJ{}.fAEWeMURGLm0) != 80 || unsafe.Sizeof(UmYbKhCpX{}) != 72 {
		t.Fatal("RTTI")
	}
	q := queueFixture(30)
	defer q.Close()
	for _, seq := range []uint64{4, 2, 1, 3, 0} {
		if err := q.Push(UmYbKhCpX{JMZzwbXWi: seq, F6dEJnE3: bytes.Repeat([]byte{byte(seq)}, 7)}); err != nil {
			t.Fatal(err)
		}
	}
	var got []byte
	for len(got) < 35 {
		p := make([]byte, 3)
		n, e := q.Read(p)
		if e != nil {
			t.Fatal(e)
		}
		got = append(got, p[:n]...)
	}
	var want []byte
	for seq := 0; seq < 5; seq++ {
		want = append(want, bytes.Repeat([]byte{byte(seq)}, 7)...)
	}
	if !bytes.Equal(got, want) || q.bDCX9Qu != 5 {
		t.Fatalf("%x counter=%d", got, q.bDCX9Qu)
	}
	// 旧序号单独入队时，原函数清除该项后返回 (0,nil)。
	q.Push(UmYbKhCpX{JMZzwbXWi: 1, F6dEJnE3: []byte("old")})
	if n, e := q.Read(make([]byte, 8)); n != 0 || e != nil {
		t.Fatalf("old %d %v", n, e)
	}
	q.Push(UmYbKhCpX{JMZzwbXWi: 5})
	if n, e := q.Read(nil); n != 0 || e != nil || q.bDCX9Qu != 6 {
		t.Fatalf("empty %d %v", n, e)
	}
}

type queueReader struct {
	*bytes.Reader
	closed int
}

func (r *queueReader) Close() error { r.closed++; return io.ErrUnexpectedEOF }
func TestOriginalUploadQueueStreamAndClose(t *testing.T) {
	q := queueFixture(3)
	r := &queueReader{Reader: bytes.NewReader([]byte("stream"))}
	if e := q.Push(UmYbKhCpX{ACMt1h1: r}); e != nil {
		t.Fatal(e)
	}
	if e := q.Push(UmYbKhCpX{}); e == nil || e.Error() != "h.reader already exists" {
		t.Fatal(e)
	}
	p := make([]byte, 6)
	if n, e := q.Read(p); n != 6 || e != nil || string(p) != "stream" {
		t.Fatalf("%d %v %q", n, e, p)
	}
	if e := q.Close(); !errors.Is(e, io.ErrUnexpectedEOF) {
		t.Fatal(e)
	}
	if e := q.Close(); !errors.Is(e, io.ErrUnexpectedEOF) || r.closed != 2 {
		t.Fatalf("reclose %v %d", e, r.closed)
	}
	if e := q.Push(UmYbKhCpX{}); e == nil || e.Error() != "packet queue closed" {
		t.Fatal(e)
	}
	// Close 必须发现尚未被 Read 取出的流。
	q = queueFixture(3)
	r = &queueReader{Reader: bytes.NewReader(nil)}
	q.Push(UmYbKhCpX{ACMt1h1: r})
	q.Close()
	if r.closed != 1 || len(q.dL_ULZL) != 0 {
		t.Fatal("unread stream not closed")
	}
	if _, ok := <-q.dL_ULZL; ok {
		t.Fatal("channel remains open")
	}
}
func TestOriginalUploadQueueOverflowBoundaryAndClear(t *testing.T) {
	q := queueFixture(2)
	// 判断在 heap.Pop 之后且使用 >，所以需有 limit+2 个乱序包。
	for i := 1; i <= 4; i++ {
		heap.Push(&q.vPdgI_lx, UmYbKhCpX{JMZzwbXWi: uint64(i), F6dEJnE3: []byte{byte(i)}})
	}
	if n, e := q.Read(make([]byte, 1)); n != 0 || e == nil || e.Error() != "packet queue is too large" {
		t.Fatalf("%d %v", n, e)
	}
	backing := q.vPdgI_lx
	if e := q.Close(); e != nil {
		t.Fatal(e)
	}
	if len(q.vPdgI_lx) != 0 {
		t.Fatal("heap not cleared")
	}
	for _, p := range backing {
		if p.F6dEJnE3 != nil || p.Au22AAVkjZY != nil {
			t.Fatal("retained slices")
		}
	}
	if n, e := q.Read(make([]byte, 1)); n != 0 || e != io.EOF {
		t.Fatalf("closed %d %v", n, e)
	}
}
func TestOriginalXHTTPPathAndMetadataDefaults(t *testing.T) {
	for _, tc := range []struct{ raw, path, query string }{{"", "/", ""}, {"abc?x=%20?y", "/abc/", "x=%20?y"}, {"/a//?", "/a//", ""}, {"//x", "//x/", ""}} {
		c := &R0J19auWof{HoBz4V: tc.raw, XEKp9SbqKQ: "header", HqtZJaRZNw: "query"}
		if c.GetNormalizedPath() != tc.path || c.GetNormalizedQuery() != tc.query {
			t.Fatal(tc)
		}
	}
	for _, tc := range []struct{ place, session, seq string }{{"", "", ""}, {"header", "X-Session", "X-Seq"}, {"query", "x_session", "x_seq"}, {"cookie", "x_session", "x_seq"}, {"Header", "", ""}} {
		c := &R0J19auWof{XEKp9SbqKQ: tc.place, HqtZJaRZNw: tc.place}
		if c.GetNormalizedSessionKey() != tc.session || c.GetNormalizedSeqKey() != tc.seq {
			t.Fatal(tc)
		}
		c.G9oU9diBqr = "custom"
		c.NPLvCM = "sequence"
		if c.GetNormalizedSessionKey() != "custom" || c.GetNormalizedSeqKey() != "sequence" {
			t.Fatal("custom")
		}
	}
}
