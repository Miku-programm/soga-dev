package gPJ6UJF

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestOriginalBufferPoolsResetAndSize(t *testing.T) {
	if reflect.TypeOf(QBjm6CMcar{}).Size() != 8 || reflect.TypeOf(RqHcmi{}).Size() != 8 {
		t.Fatal("original pointer layouts")
	}
	for _, n := range []int{0, 16, 255, 256, 257, 512, 65000, 65536, 70000} {
		r := GetReader(strings.NewReader("old"), n)
		size := r.Size()
		r.ReadByte()
		PutReader(r)
		r = GetReader(strings.NewReader("fresh"), n)
		b, e := io.ReadAll(r)
		if e != nil || string(b) != "fresh" || r.Size() != size {
			t.Fatal(n, string(b), e)
		}
		PutReader(r)
		var a, bw bytes.Buffer
		w := GetWriter(&a, n)
		w.WriteString("discard")
		PutWriter(w)
		w = GetWriter(&bw, n)
		w.WriteString("new")
		if e = w.Flush(); e != nil || bw.String() != "new" || a.Len() != 0 {
			t.Fatal(n, e, a.String(), bw.String())
		}
		PutWriter(w)
	}
}
