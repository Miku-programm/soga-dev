package deadline

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
	"time"
	"unsafe"
)

func TestNativePipeLayoutAndPartialReads(t *testing.T) {
	if unsafe.Sizeof(P848NOz3su9{}) != 144 || unsafe.Sizeof(EjFsCHfka9Zc{}) != 144 || unsafe.Sizeof(TyIy_s{}) != 24 {
		t.Fatal("RTTI layout")
	}
	r, w := NewPipe()
	defer r.Close()
	defer w.Close()
	if unsafe.Pointer(r) != unsafe.Pointer(w) {
		t.Fatal("inline constructor alias")
	}
	payload := bytes.Repeat([]byte{0, 1, 2, 255}, 4096)
	done := make(chan error, 1)
	go func() {
		n, e := w.Write(payload)
		if e == nil && n != len(payload) {
			e = io.ErrShortWrite
		}
		w.Close()
		done <- e
	}()
	var got []byte
	b := make([]byte, 31)
	for {
		n, e := r.Read(b)
		got = append(got, b[:n]...)
		if e == io.EOF {
			break
		}
		if e != nil {
			t.Fatal(e)
		}
	}
	if e := <-done; e != nil {
		t.Fatal(e)
	}
	if !bytes.Equal(got, payload) {
		t.Fatal("partial transfer")
	}
}

func TestNativePipeCloseErrorPriority(t *testing.T) {
	first := errors.New("first")
	second := errors.New("second")
	r, w := NewPipe()
	w.CloseWithError(first)
	w.CloseWithError(second)
	if _, e := r.Read(make([]byte, 1)); e != first {
		t.Fatalf("first writer error: %v", e)
	}
	r.CloseWithError(second)
	if _, e := r.Read(make([]byte, 1)); e != io.ErrClosedPipe {
		t.Fatalf("both closed read: %v", e)
	}
	if _, e := w.Write(nil); e != io.ErrClosedPipe {
		t.Fatalf("both closed write: %v", e)
	}
	if e := r.SetReadDeadline(time.Time{}); e != io.ErrClosedPipe {
		t.Fatal(e)
	}
	r, w = NewPipe()
	r.CloseWithError(first)
	if _, e := w.Write([]byte{1}); e != first {
		t.Fatal(e)
	}
}

func TestNativePipeDeadlineResetAndPartialWrite(t *testing.T) {
	r, w := NewPipe()
	defer r.Close()
	defer w.Close()
	r.SetReadDeadline(time.Now().Add(-time.Second))
	if n, e := r.Read(nil); n != 0 || e != os.ErrDeadlineExceeded {
		t.Fatalf("expired read: %d %v", n, e)
	}
	r.SetReadDeadline(time.Time{})
	w.SetWriteDeadline(time.Now().Add(30 * time.Millisecond))
	result := make(chan struct {
		n int
		e error
	}, 1)
	go func() {
		n, e := w.Write([]byte("abcdef"))
		result <- struct {
			n int
			e error
		}{n, e}
	}()
	b := make([]byte, 2)
	if n, e := r.Read(b); n != 2 || e != nil || string(b) != "ab" {
		t.Fatalf("partial %d %v %q", n, e, b)
	}
	select {
	case x := <-result:
		if x.n != 2 || x.e != os.ErrDeadlineExceeded {
			t.Fatalf("partial timeout: %+v", x)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout failed")
	}
	w.SetWriteDeadline(time.Time{})
	go func() {
		n, e := w.Write(nil)
		result <- struct {
			n int
			e error
		}{n, e}
	}()
	// 原零长度写仍需与一个读操作同步，不能提前返回。
	select {
	case <-result:
		t.Fatal("empty write returned before read")
	case <-time.After(10 * time.Millisecond):
	}
	if n, e := r.Read(nil); n != 0 || e != nil {
		t.Fatal(n, e)
	}
	if x := <-result; x.n != 0 || x.e != nil {
		t.Fatal(x)
	}
}
