// Package Izrs0A_HS restores the original HTTP/2 Gun transport.
package httpconn

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"go.uber.org/atomic"
	"golang.org/x/net/http2"
	"recovered.local/soga/util/bytes"
	"recovered.local/soga/buf/syncpool"
)

type EgMrkYjAnvwk struct {
	JbYoxUV string
	EbklgbD string
}

// Nma1KEcWah preserves the original 96-byte RTTI layout.
type Nma1KEcWah struct {
	gy7Tt5cwR2i_ *http.Response
	fZ6viLKTqLal *http.Request
	jE67IuTLu    *http2.Transport
	h7EQWALPWw   *io.PipeWriter
	tYx02bvA     sync.Once
	gFnUW3RHY    *atomic.Bool
	cN_qy4ypoit  error
	qaeg__ni6JWF int
	ivecq3       *bufio.Reader
	mWaHrPjpuc   *time.Timer
}

// init: 0x1b4cfa0; strings are decoded in grpc-decoded-strings.json.
var errorAt5098610 = errors.New("invalid length")
var errorAt5098620 = errors.New("buffer too small")
var headerAt50964B8 = http.Header{
	"content-type": {"application/grpc"},
	"user-agent":   {"grpc-go/1.36.0"},
}

// gUVbPH: 0x1b4d3a0. Status and response headers are not validated.
func (c *Nma1KEcWah) gUVbPH() {
	response, err := c.jE67IuTLu.RoundTripOpt(c.fZ6viLKTqLal, http2.RoundTripOpt{})
	if err != nil {
		c.cN_qy4ypoit = err
		c.h7EQWALPWw.Close()
	} else if c.gFnUW3RHY.Load() {
		response.Body.Close()
	} else {
		c.gy7Tt5cwR2i_ = response
		c.ivecq3 = bufio.NewReader(response.Body)
	}
}

// Read: 0x1b4d640. The original skips six framing bytes without validating them.
func (c *Nma1KEcWah) Read(p []byte) (int, error) {
	c.tYx02bvA.Do(c.gUVbPH)
	if c.cN_qy4ypoit != nil {
		return 0, c.cN_qy4ypoit
	}
	if c.qaeg__ni6JWF > 0 {
		n, err := io.ReadFull(c.ivecq3, p[:min(len(p), c.qaeg__ni6JWF)])
		c.qaeg__ni6JWF -= n
		return n, err
	}
	if c.gy7Tt5cwR2i_ == nil {
		return 0, net.ErrClosed
	}
	if _, err := c.ivecq3.Discard(6); err != nil {
		return 0, err
	}
	length, err := binary.ReadUvarint(c.ivecq3)
	if err != nil {
		return 0, errorAt5098610
	}
	n, err := io.ReadFull(c.ivecq3, p[:min(uint64(len(p)), length)])
	if err == nil && int(length-uint64(n)) > 0 {
		c.qaeg__ni6JWF = int(length - uint64(n))
	}
	return n, err
}

// Write: 0x1b4d840; pool return closure: 0x1b50fe0. Even a failed write
// returns the input payload length; the pipe's encoded byte count is discarded.
func (c *Nma1KEcWah) Write(p []byte) (int, error) {
	var protobuf [11]byte
	protobuf[0] = 10
	n := 1 + binary.PutUvarint(protobuf[1:], uint64(len(p)))
	var frame [5]byte
	binary.BigEndian.PutUint32(frame[1:], uint32(n+len(p)))
	b := syncpool.AppendPoolAt5044260.Get().(*bytes.BaDekw7qe)
	defer syncpool.Wr41CoUmtL0(b)
	b.PutSlice(frame[:])
	b.PutSlice(protobuf[:n])
	b.PutSlice(p)
	_, err := c.h7EQWALPWw.Write(b.Bytes())
	if err == io.ErrClosedPipe && c.cN_qy4ypoit != nil {
		err = c.cN_qy4ypoit
	}
	return len(p), err
}

// Close: 0x1b4dac0. No transport shutdown or timer cancellation is performed.
func (c *Nma1KEcWah) Close() error {
	c.gFnUW3RHY.Swap(true)
	if c.gy7Tt5cwR2i_ != nil {
		c.gy7Tt5cwR2i_.Body.Close()
	}
	return c.h7EQWALPWw.Close()
}

// SetReadDeadline: 0x1b4dc40; SetWriteDeadline: 0x1b4dca0.
func (c *Nma1KEcWah) SetReadDeadline(t time.Time) error  { return c.SetDeadline(t) }
func (c *Nma1KEcWah) SetWriteDeadline(t time.Time) error { return c.SetDeadline(t) }

// SetDeadline: 0x1b4dd00; closure: 0x1b51020. A zero deadline closes
// immediately, rather than disabling deadlines, matching the original.
func (c *Nma1KEcWah) SetDeadline(t time.Time) error {
	d := time.Until(t)
	if c.mWaHrPjpuc != nil {
		c.mWaHrPjpuc.Reset(d)
	} else {
		c.mWaHrPjpuc = time.AfterFunc(d, func() { c.Close() })
	}
	return nil
}

// TwFKkx: 0x1b4de20. The URL's +16 field is Opaque, not RawPath.
func TwFKkx(transport *http2.Transport, options *EgMrkYjAnvwk) (net.Conn, error) {
	service := options.JbYoxUV
	if service == "" {
		service = "GunService"
	}
	reader, writer := io.Pipe()
	request := &http.Request{
		Method: "POST", Body: reader,
		URL:   &url.URL{Scheme: "https", Host: options.EbklgbD, Path: fmt.Sprintf("/%s/Tun", service), Opaque: fmt.Sprintf("//%s/%s/Tun", options.EbklgbD, service)},
		Proto: "HTTP/2", ProtoMajor: 2, ProtoMinor: 0, Header: headerAt50964B8,
	}
	c := &Nma1KEcWah{fZ6viLKTqLal: request, jE67IuTLu: transport, h7EQWALPWw: writer, gFnUW3RHY: atomic.NewBool(false)}
	go c.tYx02bvA.Do(c.gUVbPH)
	return c, nil
}

// AljPVof590m restores the inlined closure factory, whose original closure
// body is 0x1b511a0. Its raw dial callback intentionally receives no context.
func AljPVof590m(dial func(string, string) (net.Conn, error)) func(context.Context, string, string, *tls.Config) (net.Conn, error) {
	return func(ctx context.Context, network, address string, config *tls.Config) (net.Conn, error) {
		conn, err := dial(network, address)
		if err != nil {
			return nil, err
		}
		secure := tls.Client(conn, config)
		if err = secure.HandshakeContext(ctx); err != nil {
			conn.Close()
			return nil, err
		}
		protocol := secure.ConnectionState().NegotiatedProtocol
		if protocol != "h2" {
			secure.Close()
			return nil, fmt.Errorf("http2: unexpected ALPN protocol %s, want %s", protocol, "h2")
		}
		return secure, nil
	}
}

// H6iElYbPlX4: 0x1b4e360; existing-connection callback: 0x1b510a0.
func H6iElYbPlX4(conn net.Conn, config *tls.Config, options *EgMrkYjAnvwk) (net.Conn, error) {
	transport := &http2.Transport{DialTLSContext: AljPVof590m(func(string, string) (net.Conn, error) { return conn, nil }), TLSClientConfig: config, DisableCompression: true, AllowHTTP: false, PingTimeout: 0}
	return TwFKkx(transport, options)
}

// LocalAddr: 0x1b4db40; RemoteAddr: 0x1b4dbc0. Both expose the
// original net.IPv4zero global (net.init 0x65a60f), not socket addresses.
func (c *Nma1KEcWah) LocalAddr() net.Addr  { return &net.TCPAddr{IP: net.IPv4zero} }
func (c *Nma1KEcWah) RemoteAddr() net.Addr { return &net.TCPAddr{IP: net.IPv4zero} }

var _ net.Conn = (*Nma1KEcWah)(nil)
