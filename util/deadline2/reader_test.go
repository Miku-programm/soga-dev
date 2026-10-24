package deadline2

import (
	"bytes"
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

type gatedReader struct {
	started chan struct{}
	release chan struct{}
	data    []byte
	err     error
}

func (g *gatedReader) Read(p []byte) (int, error) {
	close(g.started)
	<-g.release
	return copy(p, g.data), g.err
}

func TestDeadlineTransitions(t *testing.T) {
	d := NewDeadline()
	initial := d.Wait()
	if d.Active() {
		t.Fatal("initially active")
	}
	d.Set(time.Now().Add(-time.Second))
	if !d.Active() {
		t.Fatal("not active")
	}
	select {
	case <-initial:
	default:
		t.Fatal("past deadline did not close")
	}
	d.Set(time.Time{})
	if d.Active() || d.Wait() == initial {
		t.Fatal("reset failed")
	}
	next := d.Wait()
	d.Set(time.Now().Add(10 * time.Millisecond))
	select {
	case <-next:
	case <-time.After(time.Second):
		t.Fatal("timer did not close")
	}
	d.Set(time.Time{})
	if d.Wait() == next {
		t.Fatal("expired timer channel reused")
	}
	d.Set(time.Now().Add(time.Hour))
	d.Set(time.Time{})
	select {
	case <-d.Wait():
		t.Fatal("stopped timer closed replacement")
	default:
	}
	if reflect.TypeOf(*d).Size() != 32 || reflect.TypeOf(Q5KYzU{}).Size() != 80 || reflect.TypeOf(ypZm4YFaA{}).Size() != 40 {
		t.Fatal("RTTI size mismatch")
	}
}

func TestReadPreservesPendingAfterTimeoutAndLosesSplitError(t *testing.T) {
	sentinel := errors.New("read partial")
	source := &gatedReader{started: make(chan struct{}), release: make(chan struct{}), data: []byte("abcdef"), err: sentinel}
	r := NewReader(source)
	r.SetReadDeadline(time.Now().Add(20 * time.Millisecond))
	result := make(chan error, 1)
	go func() { _, err := r.Read(make([]byte, 8)); result <- err }()
	<-source.started
	select {
	case err := <-result:
		if err != os.ErrDeadlineExceeded {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("read did not time out")
	}
	if r.g0cpWB == nil {
		t.Fatal("pending read lost")
	}
	r.SetReadDeadline(time.Time{})
	close(source.release)
	p := make([]byte, 2)
	n, err := r.Read(p)
	if n != 2 || err != nil || string(p) != "ab" {
		t.Fatal(n, err, string(p))
	}
	r.SetReadDeadline(time.Now().Add(-time.Second))
	for _, want := range []string{"cd", "ef"} {
		n, err = r.Read(p)
		if n != 2 || err != nil || string(p) != want {
			t.Fatal(n, err, string(p))
		}
	}
	if r.fZxJxjHPB2Ny != nil {
		t.Fatal("retained exhausted buffer")
	}
	if _, err = r.Read(p); err != os.ErrDeadlineExceeded {
		t.Fatal(err)
	}
}

func TestReadDirectAndWholeResultError(t *testing.T) {
	r := NewReader(bytes.NewBufferString("data"))
	p := make([]byte, 8)
	if n, err := r.Read(p); n != 4 || err != nil {
		t.Fatal(n, err)
	}
	source := &gatedReader{started: make(chan struct{}), release: make(chan struct{}), data: []byte("x"), err: io.EOF}
	close(source.release)
	r = NewReader(source)
	r.SetReadDeadline(time.Now().Add(time.Hour))
	defer r.SetReadDeadline(time.Time{})
	if n, err := r.Read(p); n != 1 || err != io.EOF {
		t.Fatal(n, err)
	}
}
