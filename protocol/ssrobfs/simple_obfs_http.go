package ssrobfs

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	P "recovered.local/soga/buf/pool"
	S "recovered.local/soga/protocol/ssrshared"
	L "recovered.local/soga/util/logx"
	"time"
)

// Original k7yarZ7PgV6, 72-byte RTTI. The request retains the original body
// reader even after its one permitted body read; there is no Close/Free method.
type k7yarZ7PgV6 struct {
	io.ReadWriter
	xJOdz9bQ9L8  bool
	i00GQqMOFLio bool
	dNeNOip      []byte
	eEaGtYHWaeb  int
	d3aah5M      *http.Request
	mw4qCe0      *S.L7E7H2aW8KVs
}

func i0Z5v8x(rw io.ReadWriter) Jakq78yIQ               { return &k7yarZ7PgV6{ReadWriter: rw} } // 0x1cac120
func (c *k7yarZ7PgV6) SetServerInfo(i *S.L7E7H2aW8KVs) { c.mw4qCe0 = i }                       // 0x1cacbc0
func (*k7yarZ7PgV6) GetOverhead() int                  { return 0 }                            // 0x1cacc20

// Read restores 0x1cac1c0. The audited HTTP mapping is limited to a fresh
// ReadRequest over <=8192 memory bytes, one Body.Read, and no Body.Close.
// It deliberately reads the buffered tail raw after that single Body.Read.
func (c *k7yarZ7PgV6) Read(p []byte) (int, error) {
	if c.xJOdz9bQ9L8 {
		if c.dNeNOip == nil {
			return c.ReadWriter.Read(p)
		}
		n := copy(p, c.dNeNOip[c.eEaGtYHWaeb:])
		c.eEaGtYHWaeb += n
		if c.eEaGtYHWaeb >= len(c.dNeNOip) {
			P.PoolAt50961D8.Put(c.dNeNOip)
			c.dNeNOip = nil
		}
		return n, nil
	}
	c.dNeNOip = P.PoolAt50961D8.Get(8192)
	n, e := c.ReadWriter.Read(c.dNeNOip)
	if e != nil {
		return 0, e
	}
	r := bufio.NewReader(bytes.NewReader(c.dNeNOip[:n]))
	req, e := http.ReadRequest(r)
	if e != nil {
		return 0, e
	}
	c.xJOdz9bQ9L8 = true
	c.d3aah5M = req
	host, _, _ := net.SplitHostPort(req.Host)
	if host == "" {
		host = req.Host
	}
	c.mw4qCe0.EFLMTH = host
	body, e := req.Body.Read(c.dNeNOip)
	if e != nil && e != io.EOF {
		L.TRTbrquD(slog.LevelError, "ss http body too big")
		return 0, e
	}
	tail, e := r.Read(c.dNeNOip[body:])
	if e != nil && e != io.EOF {
		L.TRTbrquD(slog.LevelError, "ss http body 2 too big")
		return 0, e
	}
	if body <= 0 && tail <= 0 {
		P.PoolAt50961D8.Put(c.dNeNOip)
		c.dNeNOip = nil
	} else {
		c.dNeNOip = c.dNeNOip[:body+tail]
	}
	return c.Read(p)
}

// Write restores 0x1cac8e0: random bytes are not a WebSocket challenge hash.
// The first successful short write is reported as the complete plaintext size.
func (c *k7yarZ7PgV6) Write(p []byte) (int, error) {
	if c.i00GQqMOFLio {
		return c.ReadWriter.Write(p)
	}
	var random [20]byte
	rand.Read(random[:])
	header := fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\nServer: nginx/1.16.1\r\nDate: %s\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"), base64.StdEncoding.EncodeToString(random[:]))
	wire := make([]byte, len(header)+len(p))
	copy(wire, header)
	copy(wire[len(header):], p)
	c.i00GQqMOFLio = true
	_, e := c.ReadWriter.Write(wire)
	if e != nil {
		return 0, e
	}
	return len(p), nil
}
