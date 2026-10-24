package FLhkKwJ8M

import (
	"golang.org/x/net/http2"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
)

type ALZXSf struct {
	Hmv1xGaNFZBN []string
	B41d7b       string
}
type hj_AshdDlG struct {
	DrpXao
	*http2.ClientConn
	yN3X0U     *io.PipeWriter
	i7bFsJpyAB *http.Response
	eTapnX6ZC  *ALZXSf
}

// P3TqaD0kr: 0x19b3a80.
func P3TqaD0kr(raw net.Conn, o *ALZXSf) (net.Conn, error) {
	transport := new(http2.Transport)
	cc, err := transport.NewClientConn(raw)
	if err != nil {
		return nil, err
	}
	return &hj_AshdDlG{DrpXao: raw, ClientConn: cc, eTapnX6ZC: o}, nil
}

// pgJB_Y1: 0x19b33e0. Request template at 0x2173f60 is PUT, HTTP/2,
// ProtoMajor=2 with all other fields zero before the assignments below.
func (c *hj_AshdDlG) pgJB_Y1() error {
	reader, writer := io.Pipe()
	host := c.eTapnX6ZC.Hmv1xGaNFZBN[rand.Intn(len(c.eTapnX6ZC.Hmv1xGaNFZBN))]
	request := &http.Request{Method: "PUT", Proto: "HTTP/2", ProtoMajor: 2, Host: host, URL: &url.URL{Scheme: "https", Host: host, Path: c.eTapnX6ZC.B41d7b}, Body: reader, Header: http.Header{"Accept-Encoding": []string{"identity"}}}
	response, err := c.ClientConn.RoundTrip(request)
	if err != nil {
		return err
	}
	c.yN3X0U = writer
	c.i7bFsJpyAB = response
	return nil
}

// Read: 0x19b3880. The response Close flag triggers a new request.
func (c *hj_AshdDlG) Read(p []byte) (int, error) {
	if c.i7bFsJpyAB == nil || c.i7bFsJpyAB.Close {
		if err := c.pgJB_Y1(); err != nil {
			return 0, err
		}
	}
	return c.i7bFsJpyAB.Body.Read(p)
}

// Write: 0x19b3940.
func (c *hj_AshdDlG) Write(p []byte) (int, error) {
	if c.yN3X0U == nil {
		if err := c.pgJB_Y1(); err != nil {
			return 0, err
		}
	}
	return c.yN3X0U.Write(p)
}

// Close: 0x19b3a00; both close results are ignored by the original.
func (c *hj_AshdDlG) Close() error {
	if c.yN3X0U != nil {
		c.yN3X0U.Close()
	}
	c.ClientConn.Close()
	return nil
}
