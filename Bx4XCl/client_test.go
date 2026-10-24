package Bx4XCl

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

type scriptedConn struct {
	*bytes.Reader
	writes [][]byte
}

func (s *scriptedConn) Write(b []byte) (int, error) {
	s.writes = append(s.writes, append([]byte(nil), b...))
	return len(b), nil
}
func TestNegotiationPreservesOriginalHeaderBehavior(t *testing.T) {
	// The final reply has a deliberately invalid version/nonzero REP/reserved
	// byte. The original consumes these three bytes without inspecting them.
	s := &scriptedConn{Reader: bytes.NewReader([]byte{5, 0, 9, 7, 8, 1, 127, 0, 0, 1, 0, 80})}
	addr, err := AMwLSL6GewF(s, []byte{3, 1, 'a', 0, 80}, 1, nil)
	if err != nil || !bytes.Equal(addr, []byte{1, 127, 0, 0, 1, 0, 80}) {
		t.Fatal(addr, err)
	}
	if len(s.writes) != 2 || !bytes.Equal(s.writes[0], []byte{5, 1, 0}) || !bytes.Equal(s.writes[1], []byte{5, 1, 0, 3, 1, 'a', 0, 80}) {
		t.Fatal(s.writes)
	}
}
func TestUsernamePasswordAndFailurePaths(t *testing.T) {
	s := &scriptedConn{Reader: bytes.NewReader([]byte{5, 2, 99, 0, 5, 0, 0, 3, 1, 'a', 0, 1})}
	_, err := AMwLSL6GewF(s, []byte{1, 1, 1, 1, 1, 0, 1}, 3, &Auth{"u", "p"})
	if err != nil || !bytes.Equal(s.writes[0], []byte{5, 1, 2}) || !bytes.Equal(s.writes[1], []byte{1, 1, 'u', 1, 'p'}) {
		t.Fatal(err, s.writes)
	}
	for _, test := range []struct {
		reply []byte
		auth  *Auth
		want  string
	}{{[]byte{4, 0}, nil, "SOCKS version error"}, {[]byte{5, 2}, nil, "auth failed"}, {[]byte{5, 255}, nil, "SOCKS need auth"}, {[]byte{5, 2, 1, 1}, &Auth{}, "rejected username/password"}} {
		s := &scriptedConn{Reader: bytes.NewReader(test.reply)}
		_, err := AMwLSL6GewF(s, nil, 1, test.auth)
		if err == nil || err.Error() != test.want {
			t.Fatal(err, test.want)
		}
	}
}
func TestReadAddressPartialAndAliasing(t *testing.T) {
	b := make([]byte, 259)
	got, err := YhN2Ke2lC(bytes.NewReader([]byte{1, 10}), b)
	if len(got) != 7 || !errors.Is(err, io.ErrUnexpectedEOF) || &got[0] != &b[0] {
		t.Fatal(got, err)
	}
	if got, err = YhN2Ke2lC(bytes.NewReader([]byte{3}), b); got != nil || err != io.EOF {
		t.Fatal(got, err)
	}
	if _, err = YhN2Ke2lC(bytes.NewReader([]byte{9}), b); err != LT72UC(8) || err.Error() != "SOCKS error: 8" {
		t.Fatal(err)
	}
	if _, err = YhN2Ke2lC(bytes.NewReader(nil), b[:258]); err != io.ErrShortBuffer {
		t.Fatal(err)
	}
}
