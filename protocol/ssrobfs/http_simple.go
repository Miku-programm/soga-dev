package ssrobfs

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	C "recovered.local/soga/transport/listener"
	P "recovered.local/soga/buf/pool"
	S "recovered.local/soga/protocol/ssrshared"
	"strings"
	"time"
)

// Original RTTI 0x36e2d58, 56 bytes. Both HTTP constructors use this type.
type nArN8vNMTaR struct {
	io.ReadWriter
	mI41hq      bool
	a1yGpXg6Y   bool
	q60_aN      bool
	z9a8K4aeGWS *S.L7E7H2aW8KVs
	C.JQi6241
}

func zbGjqkTbNkOk(rw io.ReadWriter) Jakq78yIQ             { return &nArN8vNMTaR{ReadWriter: rw} } // 0x1ca7ee0
func ujOdSji25S(rw io.ReadWriter) Jakq78yIQ               { return &nArN8vNMTaR{ReadWriter: rw} } // 0x1ca7fe0
func (c *nArN8vNMTaR) SetServerInfo(info *S.L7E7H2aW8KVs) { c.z9a8K4aeGWS = info }                //0x1ca8080
func (*nArN8vNMTaR) GetOverhead() int                     { return 0 }                            //0x1ca9280

// bDsaTMgl restores 0x1ca80e0. The original stops at the first percent
// component whose length differs from two; a one-digit component is padded.
// kJalfXv.VF9Lvv4m@6054a0 is mapped only for the empty separator Join domain.
func (*nArN8vNMTaR) bDsaTMgl(data []byte) []byte {
	parts := bytes.Split(data, []byte{'%'})
	pieces := make([]string, len(parts))
	for _, part := range parts[1:] {
		if len(part) < 2 {
			pieces = append(pieces, "0", string(part))
			break
		}
		pieces = append(pieces, string(part[:2]))
		if len(part) > 2 {
			break
		}
	}
	decoded, e := hex.DecodeString(strings.Join(pieces, ""))
	if e != nil {
		return nil
	}
	return decoded
}

// Read restores 0x1ca8460. Header framing requires a complete GET/POST header
// in the first Read. The parser never reads Request.Body in this call domain.
func (c *nArN8vNMTaR) Read(p []byte) (int, error) {
	if c.a1yGpXg6Y {
		if !c.q60_aN {
			return c.ReadWriter.Read(p)
		}
		b, e := c.GetReadBufferDirect()
		if e != nil {
			return 0, e
		}
		n := copy(p, b.Bytes())
		b.Skip(n)
		if b.IsEmpty() {
			c.FreeReadBuffer()
			c.q60_aN = false
		}
		return n, nil
	}
	scratch := P.PoolAt50961D8.Get(8192)
	defer P.PoolAt50961D8.Put(scratch)
	n, e := c.ReadWriter.Read(scratch)
	if e != nil {
		return 0, e
	}
	data := scratch[:n]
	if !bytes.HasPrefix(data, []byte("GET ")) && !bytes.HasPrefix(data, []byte("POST ")) {
		return 0, errors.New("http_simple: can not decode http header")
	}
	at := bytes.Index(data, []byte("\r\n\r\n"))
	if at < 0 {
		return 0, errors.New("http_simple: can not decode http header")
	}
	c.a1yGpXg6Y = true
	// Original C9@a246a0 uses MaxInt64 header limits. <=8192 source bytes and
	// unconsumed Body match the audited native-http-obfs-parser-mapping domain.
	req, e := http.ReadRequest(bufio.NewReader(bytes.NewReader(data[:at+4])))
	if e != nil && e != io.EOF {
		return 0, e
	}
	if req == nil {
		return 0, errors.New("parse http header failed")
	}
	host, _, _ := net.SplitHostPort(req.Host)
	if host == "" {
		host = req.Host
	}
	c.z9a8K4aeGWS.EFLMTH = host
	decoded := c.bDsaTMgl(data[:bytes.Index(data, []byte("\r\n"))])
	size := n - at - 4 + len(decoded)
	if size > 0 {
		c.q60_aN = true
		b, e := c.GetReadBuffer(size)
		if e != nil {
			return 0, e
		}
		b.Write(decoded)
		b.Write(data[at+4:])
	} else {
		c.q60_aN = false
	}
	return c.Read(p)
}

// Write restores 0x1ca8fc0, including setting the sent flag before the write
// and ignoring a successful short write for the first header+payload record.
func (c *nArN8vNMTaR) Write(p []byte) (int, error) {
	if c.mI41hq {
		return c.ReadWriter.Write(p)
	}
	header := fmt.Sprintf("HTTP/1.1 200 OK\r\nConnection: keep-alive\r\nContent-Encoding: gzip\r\nContent-Type: text/html\r\nDate: %v\r\nServer: nginx\r\nVary: Accept-Encoding\r\n\r\n", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	data := P.PoolAt50961D8.Get(len(header) + len(p))
	copy(data, header)
	copy(data[len(header):], p)
	c.mI41hq = true
	_, e := c.ReadWriter.Write(data)
	P.PoolAt50961D8.Put(data)
	if e != nil {
		return 0, e
	}
	return len(p), nil
}
