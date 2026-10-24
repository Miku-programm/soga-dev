package vmessaead

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/textproto"
)

type R5Z6tXgFroMW struct {
	YORtUpVbx    string
	V7kc7NuILllS string
	DHZvNLM2Sc0K []string
	DaWC9tFk2l4E map[string][]string
}
type ywaQaL struct {
	DrpXao
	jH8e5Hst  *R5Z6tXgFroMW
	ek1b272RQ *bufio.Reader
	xAi_PX    bool
}

// NewHTTPObfuscationConn exposes the original inline construction at 0x1b6c640.
func NewHTTPObfuscationConn(raw net.Conn, o *R5Z6tXgFroMW) net.Conn {
	return &ywaQaL{DrpXao: raw, jH8e5Hst: o}
}

// Read: 0x19b7480. Status line and MIME headers are consumed but not validated.
func (c *ywaQaL) Read(p []byte) (int, error) {
	if c.ek1b272RQ != nil {
		return c.ek1b272RQ.Read(p)
	}
	r := textproto.NewReader(bufio.NewReader(c.DrpXao))
	if _, err := r.ReadLine(); err != nil {
		return 0, err
	}
	if _, err := r.ReadMIMEHeader(); err != nil {
		return 0, err
	}
	c.ek1b272RQ = r.R
	return r.R.Read(p)
}

// Write: 0x19b7640. Empty path/header alternatives retain the original Intn panic.
// Request-construction errors are deliberately ignored before the dereference.
func (c *ywaQaL) Write(p []byte) (int, error) {
	if c.xAi_PX {
		return c.DrpXao.Write(p)
	}
	o := c.jH8e5Hst
	path := o.DHZvNLM2Sc0K[rand.Intn(len(o.DHZvNLM2Sc0K))]
	host := o.V7kc7NuILllS
	if hosts := o.DaWC9tFk2l4E["Host"]; len(hosts) != 0 {
		host = hosts[rand.Intn(len(hosts))]
	}
	method := o.YORtUpVbx
	if method == "" {
		method = "GET"
	}
	request, _ := http.NewRequest(method, fmt.Sprintf("http://%s%s", host, path), bytes.NewBuffer(p))
	for key, values := range o.DaWC9tFk2l4E {
		request.Header.Set(key, values[rand.Intn(len(values))])
	}
	request.ContentLength = int64(len(p))
	if err := request.Write(c.DrpXao); err != nil {
		return 0, err
	}
	c.xAi_PX = true
	return len(p), nil
}

// Close: 0x19b7c60.
func (c *ywaQaL) Close() error { return c.DrpXao.Close() }
