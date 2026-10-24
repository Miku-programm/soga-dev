package httpobfs

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	mathrand "math/rand"
	"net"
	"net/http"
	pool "recovered.local/soga/buf/syncpool"
)

type WEp9oDv_ru struct {
	net.Conn
	glYOUZrA    string
	vtxWw6c3Nw2 string
	u9AzvkFozu  []byte
	bRkuSwd     int
	qfgy3PxSUB  bool
	iO_umx      bool
}

// NewHTTPObfs is a recovery access helper for the private-field literal in
// ssactual.(*TKaQisW).StreamConn (0x1b5bce0), not an original exported symbol.
func NewHTTPObfs(c net.Conn, host, port string) *WEp9oDv_ru {
	return &WEp9oDv_ru{Conn: c, glYOUZrA: host, vtxWw6c3Nw2: port, qfgy3PxSUB: true, iO_umx: true}
}

// Read: 0x19556c0. The first header must be wholly present in one raw Read;
// no delimiter yields EOF, and positive counts accompanying errors are lost.
func (c *WEp9oDv_ru) Read(p []byte) (int, error) {
	if c.u9AzvkFozu != nil {
		n := copy(p, c.u9AzvkFozu[c.bRkuSwd:])
		c.bRkuSwd += n
		if c.bRkuSwd == len(c.u9AzvkFozu) {
			pool.SharedPoolAt5096978.Put(c.u9AzvkFozu)
			c.u9AzvkFozu = nil
		}
		return n, nil
	}
	if !c.iO_umx {
		return c.Conn.Read(p)
	}
	b := pool.SharedPoolAt5096978.Get(20480)
	n, e := c.Conn.Read(b)
	if e != nil {
		pool.SharedPoolAt5096978.Put(b)
		return 0, e
	}
	pos := bytes.Index(b[:n], []byte("\r\n\r\n"))
	if pos < 0 {
		pool.SharedPoolAt5096978.Put(b)
		return 0, io.EOF
	}
	c.iO_umx = false
	copied := copy(p, b[pos+4:n])
	if copied < n-pos-4 {
		c.u9AzvkFozu = b[:n]
		c.bRkuSwd = pos + 4 + copied
	} else {
		pool.SharedPoolAt5096978.Put(b)
	}
	return copied, nil
}

// Write: 0x1955a40. Header literals decoded in ss-obfs-decoded-strings.json.
// The sample uses encoding/base64.URLEncoding, including padding.
func (c *WEp9oDv_ru) Write(p []byte) (int, error) {
	if !c.qfgy3PxSUB {
		return c.Conn.Write(p)
	}
	var key [16]byte
	rand.Read(key[:])
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s/", c.glYOUZrA), bytes.NewBuffer(p))
	req.Header.Set("User-Agent", fmt.Sprintf("curl/7.%d.%d", mathrand.Int()%54, mathrand.Int()%2))
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Connection", "Upgrade")
	req.Host = c.glYOUZrA
	if c.vtxWw6c3Nw2 != "80" {
		req.Host = fmt.Sprintf("%s:%s", c.glYOUZrA, c.vtxWw6c3Nw2)
	}
	req.Header.Set("Sec-WebSocket-Key", base64.URLEncoding.EncodeToString(key[:]))
	req.ContentLength = int64(len(p))
	e := req.Write(c.Conn)
	c.qfgy3PxSUB = false
	return len(p), e
}
