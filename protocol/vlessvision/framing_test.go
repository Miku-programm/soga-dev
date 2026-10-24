package vlessvision

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"
)

type capturingWriter struct {
	calls [][]byte
	n     int
	err   error
}

func (w *capturingWriter) Write(p []byte) (int, error) {
	w.calls = append(w.calls, bytes.Clone(p))
	return w.n, w.err
}
func TestFramingOriginalLayoutsAndCommands(t *testing.T) {
	for _, tc := range []struct {
		typ  reflect.Type
		size uintptr
	}{{reflect.TypeFor[KhhcaT7](), 24}, {reflect.TypeFor[CZs766affqOH](), 32}, {reflect.TypeFor[OuMHpcpymd1](), 32}} {
		if tc.typ.Size() != tc.size {
			t.Fatal(tc.typ)
		}
	}
	reader := bytes.NewReader(nil)
	writer := new(bytes.Buffer)
	for _, command := range []KYlGQuja{1, 2, 3} {
		h := &ZfmpaHtmu{aj8eMgsHQRY: command}
		r, err := iP6eRhiZ(reader, h)
		if err != nil {
			t.Fatal(err)
		}
		if command == 2 {
			if _, ok := r.(*KhhcaT7); !ok {
				t.Fatal(r)
			}
		} else if r != reader {
			t.Fatal("identity changed")
		}
		w, err := yPiHcQrp4Vs(writer, h)
		if err != nil {
			t.Fatal(err)
		}
		if command == 2 {
			if _, ok := w.(*OuMHpcpymd1); !ok {
				t.Fatal(w)
			}
		} else if _, ok := w.(*CZs766affqOH); !ok {
			t.Fatal(w)
		}
	}
	if r, err := iP6eRhiZ(reader, &ZfmpaHtmu{aj8eMgsHQRY: 7}); r != nil || err == nil || err.Error() != "unknown command: 7" {
		t.Fatalf("%v %v", r, err)
	}
	if w, err := yPiHcQrp4Vs(writer, &ZfmpaHtmu{aj8eMgsHQRY: 7}); w != nil || err == nil || err.Error() != "unknown command: 7" {
		t.Fatalf("%v %v", w, err)
	}
}
func TestStreamHeaderOnceAndOriginalShortWrite(t *testing.T) {
	sink := &capturingWriter{n: 1}
	w := &CZs766affqOH{Writer: sink, g7F22C: &ZfmpaHtmu{h8FCLanKLDuj: 4}}
	if n, err := w.Write([]byte("abc")); n != 3 || err != nil || !bytes.Equal(sink.calls[0], []byte{4, 0, 'a', 'b', 'c'}) {
		t.Fatalf("%d %v %v", n, err, sink.calls)
	}
	if n, err := w.Write([]byte("xy")); n != 1 || err != nil || string(sink.calls[1]) != "xy" {
		t.Fatalf("%d %v", n, err)
	}
	sink = &capturingWriter{err: io.ErrClosedPipe}
	w = &CZs766affqOH{Writer: sink, g7F22C: &ZfmpaHtmu{}}
	if n, err := w.Write(nil); n != 0 || !errors.Is(err, io.ErrClosedPipe) {
		t.Fatalf("%d %v", n, err)
	}
	sink.err = nil
	_, _ = w.Write([]byte("x"))
	if string(sink.calls[1]) != "x" {
		t.Fatal("failed first header retried")
	}
}
func TestDatagramFramesAndOversizeOriginalLength(t *testing.T) {
	sink := new(capturingWriter)
	w := &OuMHpcpymd1{Writer: sink, g7F22C: &ZfmpaHtmu{h8FCLanKLDuj: 6}}
	if n, err := w.Write([]byte("abc")); n != 3 || err != nil || !bytes.Equal(sink.calls[0], []byte{6, 0, 0, 3, 'a', 'b', 'c'}) {
		t.Fatalf("%d %v %v", n, err, sink.calls)
	}
	if n, err := w.Write(nil); n != 0 || err != nil || !bytes.Equal(sink.calls[1], []byte{0, 0}) {
		t.Fatal(sink.calls[1])
	}
	data := make([]byte, 65537)
	if n, err := w.Write(data); n != len(data) || err != nil || len(sink.calls[2]) != 65539 || sink.calls[2][0] != 0 || sink.calls[2][1] != 1 {
		t.Fatalf("%d %v", n, err)
	}
	sink.err = io.ErrClosedPipe
	if n, err := w.Write([]byte("x")); n != 0 || err != io.ErrClosedPipe {
		t.Fatalf("%d %v", n, err)
	}
}
func TestDatagramShortReadDiscardsRemainder(t *testing.T) {
	wire := []byte{0, 5, 'a', 'b', 'c', 'd', 'e', 0, 2, 'x', 'y', 0, 0}
	r := &KhhcaT7{Reader: bytes.NewReader(wire)}
	out := make([]byte, 2)
	if n, err := r.Read(out); n != 2 || err != nil || string(out) != "ab" {
		t.Fatalf("%d %v %q", n, err, out)
	}
	if n, err := r.Read(out); n != 2 || err != nil || string(out) != "xy" {
		t.Fatalf("%d %v %q", n, err, out)
	}
	if n, err := r.Read(out); n != 0 || err != nil {
		t.Fatalf("empty frame %d %v", n, err)
	}
	if n, err := r.Read(out); n != 0 || err != io.EOF {
		t.Fatalf("end %d %v", n, err)
	}
	r = &KhhcaT7{Reader: bytes.NewReader([]byte{0, 4, 'a', 'b', 'c'})}
	if n, err := r.Read(out); n != 2 || err != io.EOF {
		t.Fatalf("discard EOF %d %v", n, err)
	}
	r = &KhhcaT7{Reader: bytes.NewReader([]byte{0, 4, 'a'})}
	if n, err := r.Read(out); n != 1 || err != io.ErrUnexpectedEOF {
		t.Fatalf("payload short %d %v", n, err)
	}
	r = &KhhcaT7{Reader: bytes.NewReader([]byte{0, 2, 'a', 'b', 0, 1, 'c'})}
	if n, err := r.Read(nil); n != 0 || err != nil {
		t.Fatalf("zero destination %d %v", n, err)
	}
	if n, err := r.Read(out); n != 1 || err != nil || out[0] != 'c' {
		t.Fatalf("next frame %d %v", n, err)
	}
}
