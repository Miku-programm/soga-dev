package Hx04IbKk

import (
	"bytes"
	"io"
	"testing"
)

func TestReadFullRollbackKeepsPartialData(t *testing.T) {
	b := BaDekw7qe(make([]byte, 1, 8))
	b[0] = 9
	if err := b.ReadFull(bytes.NewReader([]byte{1, 2}), 4); err != io.ErrUnexpectedEOF {
		t.Fatal(err)
	}
	if len(b) != 1 || !bytes.Equal(b[:3], []byte{9, 1, 2}) {
		t.Fatal(b[:3])
	}
	if err := b.ReadFull(bytes.NewReader([]byte{7, 8}), 2); err != nil || !bytes.Equal(b, []byte{9, 7, 8}) {
		t.Fatal(b, err)
	}
}
