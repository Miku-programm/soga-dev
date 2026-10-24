package httpnative

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/util/errutil"
	"recovered.local/soga/crypto/ssmethod"
	"time"
)

// Foi2jG retains the original 80-byte native HTTP-obfuscation state.
type Foi2jG struct {
	DrpXao
	vxx3ZE5S     bool
	bYfXkyD      bool
	eAvkcvV0sihC *bufio.Reader
	d8LZbJ0      *pool.X_aXX5
	ceaIAG       string
	aTUGtv8      string
	ewxQcQsYiLi  ssmethod.VE7a8rvh
}

// NewHTTPForFactory is an explicitly named reconstruction bridge for the
// cross-package struct literal in original ssudp.FciVp4yFHL@0x1c83e20.
func NewHTTPForFactory(conn net.Conn, path, host string, mode ssmethod.VE7a8rvh) *Foi2jG {
	return &Foi2jG{DrpXao: conn, ceaIAG: path, aTUGtv8: host, ewxQcQsYiLi: mode}
}
func (c *Foi2jG) cdA93sBqKlxa(p []byte) (int, error) { //0x1c79cc0
	c.vxx3ZE5S = true
	c.d8LZbJ0 = pool.S4OQXWIZ(8192)
	if _, err := c.d8LZbJ0.ReadMaxFrom(c.DrpXao, 8192); err != nil {
		return 0, err
	}
	r := bufio.NewReader(c.d8LZbJ0)
	// Original C9Zu2eNw8mPJ@0xa246a0: ReadRequest with MaxInt64 header
	// count, over the already bounded 8192-byte buffer, and never reading Body.
	// The matched parser scope and unreachable resource-limit differences are
	// recorded in native-http-obfs-parser-mapping.json; this is not a claim
	// that the entire obfuscated HTTP fork equals the standard library.
	request, err := http.ReadRequest(r)
	if err != nil {
		return 0, err
	}
	// The original falls back according to host length, not SplitHostPort's error.
	host, _, _ := net.SplitHostPort(request.Host)
	if len(host) == 0 {
		host = request.Host
	}
	mode := c.ewxQcQsYiLi
	if (mode == 1 || mode == 3) && host != c.aTUGtv8 {
		return 0, fmt.Errorf("client host (%v) not match server host (%v)", host, c.aTUGtv8)
	}
	if (mode == 1 || mode == 2) && request.URL.Path != c.ceaIAG {
		return 0, fmt.Errorf("client path (%v) not match server path (%v)", request.URL.Path, c.ceaIAG)
	}
	if r.Buffered() > 0 {
		c.eAvkcvV0sihC = r
	}
	return c.Read(p)
}
func (c *Foi2jG) Read(p []byte) (int, error) { //0x1c7a560
	if !c.vxx3ZE5S {
		return c.cdA93sBqKlxa(p)
	}
	if c.eAvkcvV0sihC != nil {
		if c.eAvkcvV0sihC.Buffered() > 0 {
			return c.eAvkcvV0sihC.Read(p)
		}
		c.eAvkcvV0sihC = nil
	}
	if c.d8LZbJ0 != nil {
		if c.d8LZbJ0.Len() > 0 {
			return c.d8LZbJ0.Read(p)
		}
		c.d8LZbJ0.Free()
		c.d8LZbJ0 = nil
	}
	return c.DrpXao.Read(p)
}
func (c *Foi2jG) Write(p []byte) (int, error) { //0x1c7a740
	if c.bYfXkyD {
		return c.DrpXao.Write(p)
	}
	var random [20]byte
	rand.Read(random[:])
	accept := base64.StdEncoding.EncodeToString(random[:])
	header := fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\nServer: nginx/1.27.0\r\nDate: %s\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"), accept)
	b := pool.S4OQXWIZ(len(p) + len(header))
	defer b.Free()
	b.WriteString(header)
	b.Write(p)
	c.bYfXkyD = true
	if _, err := c.DrpXao.Write(b.Bytes()); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (c *Foi2jG) Free() { //0x1c7ab80
	if c.d8LZbJ0 != nil {
		c.d8LZbJ0.Free()
		c.d8LZbJ0 = nil
	}
	errutil.Li3c4CJ1_D(c.DrpXao)
}
