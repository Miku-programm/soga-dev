package pcnXtbh8Z5p

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"io"
	"net"
	"net/http"
	H "recovered.local/soga/Hx04IbKk"
	"time"
)

// Original RTTI structures: sizes 64, 32 and 56 respectively.
type LckgbWP0Vg4 struct {
	HQxCy9slSH string
	MOWJ8N     []string
	DAPCgJns   string
	KJRBYI3fW  bool
}
type MrjKETeSu struct {
	rTtnkhQSU   *LckgbWP0Vg4
	_mr1q5jad7H []byte
}
type KTJbqZyG_iM struct {
	ZgOmB9r     string
	XOOe7a      string
	Cgk4_Jjow   string
	VA7lpXIPEuY http.Header
}

// These slices combine original ELF data with init@0x1b51560's decoded strings.
var alpnAt5037D30 = []string{"h2", "http/1.1"}
var alpnAt5037D50 = []string{"http/1.1"}
var crlfAt5037D70 = []byte{'\r', '\n'}

// NewClient is a named recovery of the inlined allocation in 0x1b6a1a0:
// the retained configuration pointer is not cloned.
func NewClient(config *LckgbWP0Vg4) *MrjKETeSu {
	return &MrjKETeSu{rTtnkhQSU: config, _mr1q5jad7H: g4hbN5HjkIIm([]byte(config.HQxCy9slSH))}
}

// g4hbN5HjkIIm: 0x1b52e00; SHA-224 followed by lowercase hexadecimal.
func g4hbN5HjkIIm(password []byte) []byte {
	digest := sha256.Sum224(password)
	result := make([]byte, 56)
	hex.Encode(result, digest[:])
	return result
}

// StreamConn: 0x1b51640. Handshake errors are returned without closing conn.
func (c *MrjKETeSu) StreamConn(conn net.Conn) (net.Conn, error) {
	alpn := alpnAt5037D30
	if len(c.rTtnkhQSU.MOWJ8N) != 0 {
		alpn = c.rTtnkhQSU.MOWJ8N
	}
	secured := tls.Client(conn, &tls.Config{NextProtos: alpn, MinVersion: tls.VersionTLS12, InsecureSkipVerify: c.rTtnkhQSU.KJRBYI3fW, ServerName: c.rTtnkhQSU.DAPCgJns})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := secured.HandshakeContext(ctx); err != nil {
		return nil, err
	}
	return secured, nil
}

// WriteHeader: 0x1b51b00; one Write, ignoring its returned byte count.
func (c *MrjKETeSu) WriteHeader(w io.Writer, command byte, address []byte) error {
	b := H.BaDekw7qe{}
	appendBuffer(&b, c._mr1q5jad7H)
	appendBuffer(&b, crlfAt5037D70)
	i := b.Grow(1)
	b[i] = command
	appendBuffer(&b, address)
	appendBuffer(&b, crlfAt5037D70)
	_, err := w.Write(b)
	return err
}

// Factored original inlined byte-buffer copy, preserving Grow allocation rules.
func appendBuffer(b *H.BaDekw7qe, p []byte) { i := b.Grow(len(p)); copy((*b)[i:], p) }

// PacketConn: 0x1b51de0.
func (c *MrjKETeSu) PacketConn(conn net.Conn) net.PacketConn { return &Mah1Kab{Conn: conn} }
