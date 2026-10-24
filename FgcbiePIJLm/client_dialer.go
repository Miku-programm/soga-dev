/*
The MIT License (MIT)

Copyright (c) 2017-2021 Sergey Kamardin <gobwas@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
// Selected client functions checked against original ELF C/ASM; reference gobwas/ws v1.4.0.
// Client field names are analysis names: original value-only Dialer RTTI did not survive.
package FgcbiePIJLm

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	httphead "recovered.local/soga/EsS8DL8PE9"
	pbufio "recovered.local/soga/gPJ6UJF"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultClientReadBufferSize  = 4096
	DefaultClientWriteBufferSize = 4096
)

var (
	ErrHandshakeBadStatus      = fmt.Errorf("unexpected http status")
	ErrHandshakeBadSubProtocol = fmt.Errorf("unexpected protocol in %q header", headerSecProtocol)
	ErrHandshakeBadExtensions  = fmt.Errorf("unexpected extensions in %q header", headerSecProtocol)
)

type Z4oWaKU5v struct {
	ReadBufferSize, WriteBufferSize int

	Timeout time.Duration

	Protocols []string

	Extensions []httphead.BIMZTmbYHb

	Header HDpTJacC

	Host string

	OnStatusError func(status int, reason []byte, resp io.Reader)

	OnHeader func(key, value []byte) (err error)

	NetDial func(ctx context.Context, network, addr string) (net.Conn, error)

	TLSClient func(conn net.Conn, hostname string) net.Conn

	TLSConfig *tls.Config

	WrapConn func(conn net.Conn) net.Conn
}

// Original VA 0x11a36e0.
func (d Z4oWaKU5v) Dial(ctx context.Context, urlstr string) (conn net.Conn, br *bufio.Reader, hs Handshake, err error) {
	u, err := url.ParseRequestURI(urlstr)
	if err != nil {
		return nil, nil, hs, err
	}

	dialctx := ctx

	var deadline time.Time
	if t := d.Timeout; t != 0 {
		deadline = time.Now().Add(t)
		if d, ok := ctx.Deadline(); !ok || deadline.Before(d) {
			var cancel context.CancelFunc
			dialctx, cancel = context.WithDeadline(ctx, deadline)
			defer cancel()
		}
	}
	if conn, err = d.xYiGvj(dialctx, u); err != nil {
		return conn, nil, hs, err
	}
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()
	if ctx == context.Background() {

		conn.SetDeadline(deadline)
		defer conn.SetDeadline(noDeadline)
	} else {

		done := cVc8h7D3zW(ctx, conn)
		defer func() {

			done(&err)
		}()
	}

	br, hs, err = d.Upgrade(conn, u)

	return conn, br, hs, err
}

var (
	netEmptyDialer net.Dialer

	tlsEmptyConfig tls.Config
)

// Inlined original access to the zero TLS config at 0x50A34E0.
func tlsDefaultConfig() *tls.Config {
	return &tlsEmptyConfig
}

// Original VA 0x11a3f60.
func clTie97(host, defaultPort string) (hostname, addr string) {
	var (
		colon   = strings.LastIndexByte(host, ':')
		bracket = strings.IndexByte(host, ']')
	)
	if colon > bracket {
		return host[:colon], host
	}
	return host, host + defaultPort
}

// Original VA 0x11a4080.
func (d Z4oWaKU5v) xYiGvj(ctx context.Context, u *url.URL) (conn net.Conn, err error) {
	xYiGvj := d.NetDial
	if xYiGvj == nil {
		xYiGvj = netEmptyDialer.DialContext
	}
	switch u.Scheme {
	case "ws":
		_, addr := clTie97(u.Host, ":80")
		conn, err = xYiGvj(ctx, "tcp", addr)
	case "wss":
		hostname, addr := clTie97(u.Host, ":443")
		conn, err = xYiGvj(ctx, "tcp", addr)
		if err != nil {
			return nil, err
		}
		fi5kco := d.TLSClient
		if fi5kco == nil {
			fi5kco = d.fi5kco
		}
		conn = fi5kco(conn, hostname)
	default:
		return nil, fmt.Errorf("unexpected websocket scheme: %q", u.Scheme)
	}
	if wrap := d.WrapConn; wrap != nil {
		conn = wrap(conn)
	}
	return conn, err
}

// Original VA 0x11a4440.
func (d Z4oWaKU5v) fi5kco(conn net.Conn, hostname string) net.Conn {
	config := d.TLSConfig
	if config == nil {
		config = tlsDefaultConfig()
	}
	if config.ServerName == "" {
		config = config.Clone()
		config.ServerName = hostname
	}

	return tls.Client(conn, config)
}

var (
	aLongTimeAgo = time.Unix(42, 0)
)

// Original VA 0x11a4600.
func (d Z4oWaKU5v) Upgrade(conn io.ReadWriter, u *url.URL) (br *bufio.Reader, hs Handshake, err error) {

	const (
		headerSeenUpgrade = 1 << iota
		headerSeenConnection
		headerSeenSecAccept

		headerSeenAll = 0 |
			headerSeenUpgrade |
			headerSeenConnection |
			headerSeenSecAccept
	)

	br = pbufio.GetReader(conn,
		nonZero(d.ReadBufferSize, DefaultClientReadBufferSize),
	)
	bw := pbufio.GetWriter(conn,
		nonZero(d.WriteBufferSize, DefaultClientWriteBufferSize),
	)
	defer func() {
		pbufio.PutWriter(bw)
		if br.Buffered() == 0 || err != nil {

			pbufio.PutReader(br)
			br = nil
		}
	}()

	nonce := make([]byte, nonceSize)
	rn71Bf(nonce)

	eLiGoGo1nj(bw, u, nonce, d.Protocols, d.Extensions, d.Header, d.Host)
	if err := bw.Flush(); err != nil {
		return br, hs, err
	}

	sl, err := aqxVoILI83(br)
	if err != nil {
		return br, hs, err
	}

	resp, err := l8NFoELj(sl)
	if err != nil {
		return br, hs, err
	}

	if resp.major != 1 || resp.minor < 1 {
		err = ErrHandshakeBadProtocol
		return br, hs, err
	}
	if resp.status != http.StatusSwitchingProtocols {
		err = XChC4Wxtj(resp.status)
		if onStatusError := d.OnStatusError; onStatusError != nil {

			onStatusError(resp.status, resp.reason,
				io.MultiReader(
					bytes.NewReader(sl),
					strings.NewReader(crlf),
					br,
				),
			)
		}
		return br, hs, err
	}

	var headerSeen byte
	for {
		line, e := aqxVoILI83(br)
		if e != nil {
			err = e
			return br, hs, err
		}
		if len(line) == 0 {

			break
		}

		k, v, ok := mdOcFVt(line)
		if !ok {
			err = ErrMalformedResponse
			return br, hs, err
		}

		switch btsToString(k) {
		case headerUpgradeCanonical:
			headerSeen |= headerSeenUpgrade
			if !bytes.Equal(v, specHeaderValueUpgrade) && !bytes.EqualFold(v, specHeaderValueUpgrade) {
				err = ErrHandshakeBadUpgrade
				return br, hs, err
			}

		case headerConnectionCanonical:
			headerSeen |= headerSeenConnection

			if !bytes.Equal(v, specHeaderValueConnection) && !bytes.EqualFold(v, specHeaderValueConnection) {
				err = ErrHandshakeBadConnection
				return br, hs, err
			}

		case headerSecAcceptCanonical:
			headerSeen |= headerSeenSecAccept
			if !tOuNh95S(v, nonce) {
				err = ErrHandshakeBadSecAccept
				return br, hs, err
			}

		case headerSecProtocolCanonical:

			for _, want := range d.Protocols {
				if string(v) == want {
					hs.Protocol = want
					break
				}
			}
			if hs.Protocol == "" {

				err = ErrHandshakeBadSubProtocol
				return br, hs, err
			}

		case headerSecExtensionsCanonical:
			hs.Extensions, err = n3LyWV(v, d.Extensions, hs.Extensions)
			if err != nil {
				return br, hs, err
			}

		default:
			if onHeader := d.OnHeader; onHeader != nil {
				if e := onHeader(k, v); e != nil {
					err = e
					return br, hs, err
				}
			}
		}
	}
	if err == nil && headerSeen != headerSeenAll {
		switch {
		case headerSeen&headerSeenUpgrade == 0:
			err = ErrHandshakeBadUpgrade
		case headerSeen&headerSeenConnection == 0:
			err = ErrHandshakeBadConnection
		case headerSeen&headerSeenSecAccept == 0:
			err = ErrHandshakeBadSecAccept
		default:
			panic("unknown headers state")
		}
	}
	return br, hs, err
}

type XChC4Wxtj int

// Original VA 0x11a57c0.
func (s XChC4Wxtj) Error() string {
	return "unexpected HTTP response status: " + strconv.Itoa(int(s))
}

// Inlined direct net.Error assertion from 0x11C8B60; no errors.As traversal.
func isTimeoutError(err error) bool {
	t, ok := err.(net.Error)
	return ok && t.Timeout()
}

// Original VA 0x11a5860.
func n3LyWV(selected []byte, wanted, received []httphead.BIMZTmbYHb) ([]httphead.BIMZTmbYHb, error) {
	if len(selected) == 0 {
		return received, nil
	}
	var (
		index  int
		option httphead.BIMZTmbYHb
		err    error
	)
	index = -1
	match := func() (ok bool) {
		for _, want := range wanted {

			if bytes.Equal(option.KXa0Kt2I, want.KXa0Kt2I) {

				want.Jd2j6aD7, _ = option.Jd2j6aD7.Copy(make([]byte, option.Jd2j6aD7.Size()))
				received = append(received, want)
				return true
			}
		}
		return false
	}
	ok := httphead.LU7plsn8k(selected, func(i int, name, attr, val []byte) httphead.Control {
		if i != index {

			index = i
			if i != 0 && !match() {

				err = ErrHandshakeBadExtensions
				return httphead.ControlBreak
			}
			option = httphead.BIMZTmbYHb{KXa0Kt2I: name}
		}
		if attr != nil {
			option.Jd2j6aD7.Set(attr, val)
		}
		return httphead.ControlContinue
	})
	if !ok {
		err = ErrMalformedResponse
		return received, err
	}
	if !match() {
		return received, ErrHandshakeBadExtensions
	}
	return received, err
}

// Original VA 0x11a5d80.
func cVc8h7D3zW(ctx context.Context, conn net.Conn) (done func(*error)) {
	var (
		quit      = make(chan struct{})
		interrupt = make(chan error, 1)
	)
	go func() {
		select {
		case <-quit:
			interrupt <- nil
		case <-ctx.Done():

			conn.SetDeadline(aLongTimeAgo)
			interrupt <- ctx.Err()
		}
	}()
	return func(err *error) {
		close(quit)

		if ctxErr := <-interrupt; ctxErr != nil && (*err == nil || isTimeoutError(*err)) {
			*err = ctxErr
		}
	}
}
