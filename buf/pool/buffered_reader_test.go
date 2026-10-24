package pool

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"testing"
	"unsafe"
)

func TestRecoveredReaderAgainstBufio(t *testing.T) {
	data := strings.Repeat("abc\r\n中文\xffx\n", 40)
	for _, size := range []int{16, 32, 4096} {
		a := C8FauP3Auy(strings.NewReader(data), size)
		b := bufio.NewReaderSize(strings.NewReader(data), size)
		rng := rand.New(rand.NewSource(1825460))
		for i := 0; i < 500; i++ {
			var av, bv any
			var ae, be error
			switch rng.Intn(10) {
			case 0:
				n := rng.Intn(40) - 2
				x, e := a.Peek(n)
				y, f := b.Peek(n)
				av, bv = string(x), string(y)
				ae, be = e, f
			case 1:
				n := rng.Intn(12) - 2
				av, ae = a.Discard(n)
				bv, be = b.Discard(n)
			case 2:
				av, ae = a.ReadByte()
				bv, be = b.ReadByte()
			case 3:
				ae = a.UnreadByte()
				be = b.UnreadByte()
			case 4:
				x, n, e := a.ReadRune()
				y, m, f := b.ReadRune()
				av, bv = fmt.Sprint(x, n), fmt.Sprint(y, m)
				ae, be = e, f
			case 5:
				ae = a.UnreadRune()
				be = b.UnreadRune()
			case 6:
				x, e := a.ReadSlice('\n')
				y, f := b.ReadSlice('\n')
				av, bv = string(x), string(y)
				ae, be = e, f
			case 7:
				x, p, e := a.ReadLine()
				y, q, f := b.ReadLine()
				av, bv = fmt.Sprint(string(x), p), fmt.Sprint(string(y), q)
				ae, be = e, f
			case 8:
				av, ae = a.ReadString('\n')
				bv, be = b.ReadString('\n')
			case 9:
				n := rng.Intn(60)
				x, y := make([]byte, n), make([]byte, n)
				nx, e := a.Read(x)
				ny, f := b.Read(y)
				av, bv = string(x[:nx]), string(y[:ny])
				ae, be = e, f
			}
			if fmt.Sprint(av) != fmt.Sprint(bv) || fmt.Sprint(ae) != fmt.Sprint(be) || a.Buffered() != b.Buffered() {
				t.Fatalf("size=%d step=%d: %v/%v vs %v/%v", size, i, av, ae, bv, be)
			}
		}
		a.FreeBuf()
	}
}

type emptyReadSource struct{ calls int }

func (e *emptyReadSource) Read([]byte) (int, error) { e.calls++; return 0, nil }

type recordEmptyWriter struct{ sizes []int }

func (w *recordEmptyWriter) Write(p []byte) (int, error) {
	w.sizes = append(w.sizes, len(p))
	return len(p), nil
}
func TestOriginalReaderPoolLifecycleAndEmptyWrite(t *testing.T) {
	r := C8FauP3Auy(strings.NewReader("abc"), 16)
	if unsafe.Sizeof(*r) != 88 || unsafe.Offsetof(r.c6HUuPfxb526) != 24 || unsafe.Offsetof(r.lS7kcz) != 72 {
		t.Fatal("RTTI layout")
	}
	if C8FauP3Auy(r, 8) != r {
		t.Fatal("reuse")
	}
	if _, err := r.Peek(2); err != nil {
		t.Fatal(err)
	}
	underlying := r.c6HUuPfxb526
	r.FreeBuf()
	if r.bkCmHugyPy != nil || r.c6HUuPfxb526 != underlying || r.lS7kcz != -1 || r.zyE274D49tb != -1 {
		t.Fatal("FreeBuf reset")
	}
	r.Reset(strings.NewReader("abc"))
	if r.Size() != 4096 {
		t.Fatal("default pool buffer")
	}
	r.Reset(r)
	if r.c6HUuPfxb526 != r {
		t.Fatal("original Reset keeps self reference")
	}
	r.Reset(strings.NewReader("abc"))
	w := new(recordEmptyWriter)
	if n, err := r.WriteTo(w); n != 3 || err != nil {
		t.Fatal(n, err)
	}
	if len(w.sizes) < 2 || w.sizes[0] != 0 || w.sizes[1] != 3 {
		t.Fatal("original unconditional empty Write", w.sizes)
	}
	r.FreeBuf()
	empty := new(emptyReadSource)
	r = C8FauP3Auy(empty, 16)
	if _, err := r.Peek(1); err != io.ErrNoProgress || empty.calls != 100 {
		t.Fatal(err, empty.calls)
	}
	r.FreeBuf()
	r = C8FauP3Auy(bytes.NewBufferString(strings.Repeat("z", 40)+"\n"), 16)
	got, err := r.ReadBytes('\n')
	if err != nil || len(got) != 41 {
		t.Fatal(len(got), err)
	}
	r.FreeBuf()
}
