package vmessaead

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"testing"
	"time"
)

type vmessMemoryConn struct {
	r        *bytes.Reader
	w        bytes.Buffer
	writeErr error
}

func (c *vmessMemoryConn) Read(p []byte) (int, error) { return c.r.Read(p) }
func (c *vmessMemoryConn) Write(p []byte) (int, error) {
	if c.writeErr != nil {
		return 0, c.writeErr
	}
	return c.w.Write(p)
}
func (c *vmessMemoryConn) Close() error                     { return nil }
func (c *vmessMemoryConn) LocalAddr() net.Addr              { return &net.TCPAddr{} }
func (c *vmessMemoryConn) RemoteAddr() net.Addr             { return &net.TCPAddr{} }
func (c *vmessMemoryConn) SetDeadline(time.Time) error      { return nil }
func (c *vmessMemoryConn) SetReadDeadline(time.Time) error  { return nil }
func (c *vmessMemoryConn) SetWriteDeadline(time.Time) error { return nil }
func customPeerResponse(key, iv, payload []byte, aead bool) []byte {
	if !aead {
		b, _ := aes.NewCipher(key)
		out := make([]byte, len(payload))
		cipher.NewCFBEncrypter(b, iv).XORKeyStream(out, payload)
		return out
	}
	var size [2]byte
	binary.BigEndian.PutUint16(size[:], uint16(len(payload)))
	lk := peerKDF(key, "AEAD Resp Header Len Key")
	ln := peerKDF(iv, "AEAD Resp Header Len IV")
	out := peerGCM(lk[:16]).Seal(nil, ln[:12], size[:], nil)
	pk := peerKDF(key, "AEAD Resp Header Key")
	pn := peerKDF(iv, "AEAD Resp Header IV")
	return append(out, peerGCM(pk[:16]).Seal(nil, pn[:12], payload, nil)...)
}
func TestResponseAuthenticationErrorBranches(t *testing.T) {
	key := make([]byte, 16)
	iv := bytes.Repeat([]byte{4}, 16)
	for _, aead := range []bool{false, true} {
		for _, v := range []struct {
			plain []byte
			want  string
		}{{[]byte{7, 0, 0, 0}, "unexpected response header"}, {[]byte{8, 0, 1, 0}, "dynamic port is not supported now"}, {[]byte{8, 99, 0, 255}, ""}} {
			raw := &vmessMemoryConn{r: bytes.NewReader(customPeerResponse(key, iv, v.plain, aead))}
			c := &HRzHVG1yP5O{DrpXao: raw, z1TtPlbnfnl: key, f9_fTlq: iv, kONmi_w_yL3D: 8, tu04N8lBj6I: aead}
			err := c.dCFgAbCz()
			if v.want == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || err.Error() != v.want {
				t.Fatalf("aead %v: %v", aead, err)
			}
		}
	}
	raw := &vmessMemoryConn{r: bytes.NewReader(customPeerResponse(key, iv, []byte{8, 0, 0}, true))}
	c := &HRzHVG1yP5O{DrpXao: raw, z1TtPlbnfnl: key, f9_fTlq: iv, kONmi_w_yL3D: 8, tu04N8lBj6I: true}
	if err := c.dCFgAbCz(); err == nil || err.Error() != "unexpected buffer length" {
		t.Fatal(err)
	}
}
func TestResponseLengthWrapAndRetry(t *testing.T) {
	key := make([]byte, 16)
	iv := make([]byte, 16)
	lk := peerKDF(key, "AEAD Resp Header Len Key")
	ln := peerKDF(iv, "AEAD Resp Header Len IV")
	wire := peerGCM(lk[:16]).Seal(nil, ln[:12], []byte{255, 255}, nil)
	wire = append(wire, make([]byte, 16)...)
	raw := &vmessMemoryConn{r: bytes.NewReader(wire)}
	c := &HRzHVG1yP5O{DrpXao: raw, z1TtPlbnfnl: key, f9_fTlq: iv, tu04N8lBj6I: true}
	if n, e := c.Read(nil); n != 0 || e == nil || c.ppPBszhDxm || raw.r.Len() != 1 {
		t.Fatal(n, e, raw.r.Len())
	}
	// A failed response read leaves the first-read flag unset; the next call retries.
	if n, e := c.Read(nil); n != 0 || e != io.ErrUnexpectedEOF || c.ppPBszhDxm {
		t.Fatal(n, e)
	}
}
func TestHandshakeFailureTypedNilAndUnknownSecurity(t *testing.T) {
	client, e := Iw0IEMM3ha(ClientConfiguration{UUID: "00000000-0000-0000-0000-000000000000", Security: "none", AEAD: true})
	if e != nil {
		t.Fatal(e)
	}
	sentinel := errors.New("write failed")
	raw := &vmessMemoryConn{r: bytes.NewReader(nil), writeErr: sentinel}
	dst := &IxhMZ66CqoH1{OyEjzi1v93: 1, TIBiNxg: []byte{1, 2, 3, 4}, GHP6ivev2Db: 80}
	conn, e := client.StreamConn(raw, dst)
	if e != sentinel || conn == nil || conn.(*HRzHVG1yP5O) != nil {
		t.Fatal(conn, e)
	}
	raw.writeErr = nil
	c, e := wv3qc3(raw, client.roEClw[0], dst, 2, true)
	if e != nil || c.xGnl3X != nil || c._mBMg3 != nil || c.iGH6ra != 2 || raw.w.Len() == 0 {
		t.Fatal(c, e)
	}
}
