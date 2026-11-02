package mieru

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/chacha20poly1305"
	"io"
	"net"
	D "recovered.local/soga/buf/pool"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type frameMemoryConn struct {
	bytes.Buffer
	writeErr error
	short    bool
}

func (c *frameMemoryConn) Close() error                     { return nil }
func (c *frameMemoryConn) LocalAddr() net.Addr              { return &net.TCPAddr{} }
func (c *frameMemoryConn) RemoteAddr() net.Addr             { return &net.TCPAddr{} }
func (c *frameMemoryConn) SetDeadline(time.Time) error      { return nil }
func (c *frameMemoryConn) SetReadDeadline(time.Time) error  { return nil }
func (c *frameMemoryConn) SetWriteDeadline(time.Time) error { return nil }
func (c *frameMemoryConn) Write(p []byte) (int, error) {
	if c.writeErr != nil {
		return 0, c.writeErr
	}
	if c.short {
		return 0, nil
	}
	return c.Buffer.Write(p)
}
func testFrameAEAD(t *testing.T) cipher.AEAD {
	t.Helper()
	key := bytes.Repeat([]byte{42}, 32)
	a, err := chacha20poly1305.NewX(key)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

// This decoder is deliberately independent of rZDq_kinV4G: split the wire
// using raw offsets, authenticate with x/crypto, and check every padding span.
func TestFrameWriterWire(t *testing.T) {
	// qhFcsZC 原 88 字节，新增 traffic pattern 字段后为 104。
	if unsafe.Sizeof(qhFcsZC{}) != 104 || unsafe.Sizeof(rZDq_kinV4G{}) != 56 || unsafe.Sizeof(jjyc26{}) != 32 {
		t.Fatal("RTTI layout mismatch")
	}
	for _, kind := range []byte{2, 3, 4, 5, 6, 7, 8, 9} {
		for _, parts := range [][][]byte{nil, {[]byte("one")}, {[]byte("a"), nil, []byte("bc")}} {
			conn := new(frameMemoryConn)
			a := testFrameAEAD(t)
			w, err := vfhZtLu(conn, a, "secret")
			if err != nil {
				t.Fatal(err)
			}
			initial := append([]byte(nil), w.g9zKwmF...)
			for repeat := 0; repeat < 2; repeat++ {
				m := &IqB1foHWg{Rf6N7kY6GVCs: xofvu_nD97c(kind), HNaaDh: 0x12345678, Tt1OfazA: 0xfedcba98, N01v7C: 0x99887766, UlMT6MeaV2: 0xabcd, VGN7gjQ4p: 9}
				if kind < 6 {
					err = w.WriteSessionSegment(m, parts...)
				} else {
					err = w.WriteDataSegment(m, parts...)
				}
				if err != nil {
					t.Fatal(err)
				}
				wire := append([]byte(nil), conn.Bytes()...)
				conn.Reset()
				if repeat == 0 {
					if !bytes.Equal(wire[:24], initial) {
						t.Fatal("nonce prefix")
					}
					wire = wire[24:]
				}
				plain, err := a.Open(nil, initial, wire[:48], nil)
				if err != nil {
					t.Fatal(err)
				}
				incrementNonce(initial)
				if expected := m.Marshal(); !bytes.Equal(plain, expected[:]) {
					t.Fatalf("header %x != %x", plain, expected)
				}
				wire = wire[48:]
				if kind >= 6 {
					wire = wire[int(m.B8QsA0N2yo):]
				}
				want := bytes.Join(parts, nil)
				if len(want) > 0 {
					got, e := a.Open(nil, initial, wire[:len(want)+16], nil)
					if e != nil || !bytes.Equal(got, want) {
						t.Fatalf("payload %x %v", got, e)
					}
					incrementNonce(initial)
					wire = wire[len(want)+16:]
				}
				// Session padding can theoretically be 256 bytes (encoded byte wraps),
				// preserving the original length selection rather than normalizing it.
				if byte(len(wire)) != m.FO7MEJ {
					t.Fatalf("suffix %d != %d", len(wire), m.FO7MEJ)
				}
				if !bytes.Equal(initial, w.g9zKwmF) || w.oynKh2Fa {
					t.Fatal("nonce progression")
				}
			}
		}
	}
}
func makeReadFixture(t *testing.T, kind byte, payload []byte) (*rZDq_kinV4G, *frameMemoryConn, []byte) {
	t.Helper()
	a := testFrameAEAD(t)
	nonce := bytes.Repeat([]byte{255}, 24)
	conn := new(frameMemoryConn)
	m := &IqB1foHWg{Rf6N7kY6GVCs: xofvu_nD97c(kind), WyK3yKmJKe1V: uint32(time.Now().Unix() / 60), HPUJVQ0S: uint16(len(payload)), FO7MEJ: 3}
	if kind >= 6 {
		m.B8QsA0N2yo = 2
	}
	header := m.Marshal()
	counter := append([]byte(nil), nonce...)
	incrementNonce(counter)
	conn.Buffer.Write(a.Seal(nil, counter, header[:], nil))
	if kind >= 6 {
		conn.Buffer.Write([]byte{1, 2})
	}
	if len(payload) > 0 {
		incrementNonce(counter)
		conn.Buffer.Write(a.Seal(nil, counter, payload, nil))
	}
	conn.Buffer.Write([]byte{3, 4, 5})
	return &rZDq_kinV4G{aMKXGpHBfg: conn, n49jm3JWYo: a, xUtUMCF3grV: nonce}, conn, counter
}
func TestFrameReaderIndependentWire(t *testing.T) {
	for _, kind := range []byte{2, 4, 6, 8} {
		for _, payload := range [][]byte{nil, []byte("authenticated payload")} {
			r, c, nonce := makeReadFixture(t, kind, payload)
			got, err := r.Read()
			if err != nil || !bytes.Equal(got.mzrzt1uBj, payload) || !bytes.Equal(nonce, r.xUtUMCF3grV) || c.Len() != 0 {
				t.Fatalf("%d: %#v %v remaining %d", kind, got, err, c.Len())
			}
			D.PoolAt50961D8.Put(got.mzrzt1uBj)
		}
	}
}
func TestFrameFailuresPreserveNonceAndWrapping(t *testing.T) {
	for _, where := range []string{"metadata-short", "metadata-auth", "payload-short", "payload-auth", "prefix-short", "suffix-short"} {
		r, c, _ := makeReadFixture(t, 6, []byte("abc"))
		wire := append([]byte(nil), c.Bytes()...)
		switch where {
		case "metadata-short":
			wire = wire[:47]
		case "metadata-auth":
			wire[0] ^= 1
		case "payload-short":
			wire = wire[:51]
		case "payload-auth":
			wire[50] ^= 1
		case "prefix-short":
			wire = wire[:49]
		case "suffix-short":
			wire = wire[:len(wire)-1]
		}
		c.Reset()
		c.Buffer.Write(wire)
		got, err := r.Read()
		if got != nil || err == nil {
			t.Fatalf("%s: %#v %v", where, got, err)
		}
		prefix := map[string]string{"metadata-short": "read metadata:", "metadata-auth": "decrypt metadata:", "payload-short": "read payload:", "payload-auth": "decrypt payload:", "prefix-short": "discard prefix padding:", "suffix-short": "discard suffix padding:"}[where]
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("%s: %v", where, err)
		}
		if strings.HasSuffix(where, "short") && !errors.Is(err, io.ErrUnexpectedEOF) && !errors.Is(err, io.EOF) {
			t.Fatalf("error identity %v", err)
		}
		want := strings.Repeat("ff", 24)
		if where != "metadata-short" {
			want = strings.Repeat("00", 24)
		}
		if where == "payload-auth" || where == "suffix-short" {
			want = strings.Repeat("00", 23) + "01"
		}
		if hex.EncodeToString(r.xUtUMCF3grV) != want {
			t.Fatalf("%s nonce %x", where, r.xUtUMCF3grV)
		}
	}
	sentinel := errors.New("write failed")
	for _, short := range []bool{false, true} {
		c := &frameMemoryConn{short: short}
		if !short {
			c.writeErr = sentinel
		}
		w, err := vfhZtLu(c, testFrameAEAD(t), "")
		if err != nil {
			t.Fatal(err)
		}
		before := append([]byte(nil), w.g9zKwmF...)
		incrementNonce(before)
		err = w.WriteDataSegment(&IqB1foHWg{Rf6N7kY6GVCs: 6})
		if (short && err != nil) || (!short && err != sentinel) || !bytes.Equal(before, w.g9zKwmF) || w.oynKh2Fa {
			t.Fatalf("write side effects %v", err)
		}
	}
}
