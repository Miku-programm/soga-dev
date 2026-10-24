package relay

import (
	"bytes"
	"io"
	"testing"
)

func TestOriginalAddressWireFormat(t *testing.T) {
	for _, tc := range []struct {
		wire []byte
		host string
		ip   bool
		port int
	}{
		{[]byte{1, 187, 1, 127, 0, 0, 1}, "127.0.0.1", true, 443},
		{[]byte{0, 53, 3, 0x20, 1, 0xd, 0xb8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, "2001:db8::1", true, 53},
		{append([]byte{0, 0, 2, 9}, []byte("127.0.0.1")...), "127.0.0.1", false, 0},
		{[]byte{255, 255, 2, 0}, "", false, 65535},
	} {
		input := bytes.NewReader(append(tc.wire, 0xab))
		got, err := CHYbTjleKL(input)
		if err != nil || got.GetHost() != tc.host || got.GetPort() != tc.port || got.IsIP() != tc.ip || input.Len() != 1 {
			t.Fatalf("%x => %+v %v remaining%d", tc.wire, got, err, input.Len())
		}
	}
}
func TestOriginalAddressTruncationAndTypeErrors(t *testing.T) {
	full := []byte{1, 187, 1, 127, 0, 0, 1}
	for length := 0; length < len(full); length++ {
		got, err := CHYbTjleKL(bytes.NewReader(full[:length]))
		if got != nil || err == nil {
			t.Fatalf("%d: %v %v", length, got, err)
		}
	}
	if got, err := CHYbTjleKL(bytes.NewReader([]byte{0, 80, 4})); got != nil || err == nil || err.Error() != "unknown addr type: 4" {
		t.Fatalf("%v %v", got, err)
	}
	if n, err := HoymNl_VBq(bytes.NewReader([]byte{255})); n != 0 || err != io.ErrUnexpectedEOF {
		t.Fatalf("%d %v", n, err)
	}
	if n, err := TBUcsYQLu25q(bytes.NewReader(nil)); n != 0 || err != io.EOF {
		t.Fatalf("%d %v", n, err)
	}
}
