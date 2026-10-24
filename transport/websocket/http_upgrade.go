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
// 原 ELF 的 0x11ab0a0、0x11b78e0 及调用链与 gobwas/ws v1.4.0 对照恢复。
// HTTP 仅在已核对的服务/劫持调用域使用 Go 标准类型；不宣称整包 fork 完全相同。
package websocket

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	httphead "recovered.local/soga/transport/httphead"
)

// Constants used by ConnUpgrader.
const (
	DefaultServerReadBufferSize  = 4096
	DefaultServerWriteBufferSize = 512
)

// Errors used by both client and server when preparing WebSocket handshake.
var (
	ErrHandshakeBadProtocol = RejectConnectionError(
		RejectionStatus(http.StatusHTTPVersionNotSupported),
		RejectionReason("handshake error: bad HTTP protocol version"),
	)
	ErrHandshakeBadMethod = RejectConnectionError(
		RejectionStatus(http.StatusMethodNotAllowed),
		RejectionReason("handshake error: bad HTTP request method"),
	)
	ErrHandshakeBadHost = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerHost)),
	)
	ErrHandshakeBadUpgrade = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerUpgrade)),
	)
	ErrHandshakeBadConnection = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerConnection)),
	)
	ErrHandshakeBadSecAccept = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecAccept)),
	)
	ErrHandshakeBadSecKey = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecKey)),
	)
	ErrHandshakeBadSecVersion = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecVersion)),
	)
)

// ErrMalformedResponse is returned by Dialer to indicate that server response
// can not be parsed.
var ErrMalformedResponse = fmt.Errorf("malformed HTTP response")

// ErrMalformedRequest is returned when HTTP request can not be parsed.
var ErrMalformedRequest = RejectConnectionError(
	RejectionStatus(http.StatusBadRequest),
	RejectionReason("malformed HTTP request"),
)

// ErrHandshakeUpgradeRequired is returned by Upgrader to indicate that
// connection is rejected because given WebSocket version is malformed.
//
// According to RFC6455:
// If this version does not match a version understood by the server, the
// server MUST abort the WebSocket handshake described in this section and
// instead send an appropriate HTTP error code (such as 426 Upgrade Required)
// and a |Sec-WebSocket-Version| header field indicating the version(s) the
// server is capable of understanding.
var ErrHandshakeUpgradeRequired = RejectConnectionError(
	RejectionStatus(http.StatusUpgradeRequired),
	RejectionHeader(BuqwQfCAXz(headerSecVersion+": 13\r\n")),
	RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecVersion)),
)

// ErrNotHijacker is an error returned when http.ResponseWriter does not
// implement http.Hijacker interface.
var ErrNotHijacker = RejectConnectionError(
	RejectionStatus(http.StatusInternalServerError),
	RejectionReason("given http.ResponseWriter is not a http.Hijacker"),
)

// DefaultHTTPUpgrader is an OKnCFm3eQ that holds no options and is used by
// UpgradeHTTP function.
var DefaultHTTPUpgrader OKnCFm3eQ

// UpgradeHTTP is like OKnCFm3eQ{}.Upgrade().
func UpgradeHTTP(r *http.Request, w http.ResponseWriter) (net.Conn, *bufio.ReadWriter, Handshake, error) {
	return DefaultHTTPUpgrader.Upgrade(r, w)
}

// DefaultUpgrader is an Upgrader that holds no options and is used by Upgrade
// function.

// Upgrade is like Upgrader{}.Upgrade().

// OKnCFm3eQ contains options for upgrading connection to websocket from
// net/http Handler arguments.
// 字段原名没有存活 RTTI；以下名称对应原 40B 值类型偏移 0/8/16/24/32。
type OKnCFm3eQ struct {
	// Timeout is the maximum amount of time an Upgrade() will spent while
	// writing handshake response.
	//
	// The default is no timeout.
	Timeout time.Duration

	// Header is an optional http.Header mapping that could be used to
	// write additional headers to the handshake response.
	//
	// Note that if present, it will be written in any result of handshake.
	Header http.Header

	// Protocol is the select function that is used to select subprotocol from
	// list requested by client. If this field is set, then the first matched
	// protocol is sent to a client as negotiated.
	Protocol func(string) bool

	// Extension is the select function that is used to select extensions from
	// list requested by client. If this field is set, then the all matched
	// extensions are sent to a client as negotiated.
	//
	// Deprecated: use Negotiate instead.
	Extension func(httphead.BIMZTmbYHb) bool

	// Negotiate is the callback that is used to negotiate extensions from
	// the client's offer. If this field is set, then the returned non-zero
	// extensions are sent to the client as accepted extensions in the
	// response.
	//
	// The argument is only valid until the Negotiate callback returns.
	//
	// If returned error is non-nil then connection is rejected and response is
	// sent with appropriate HTTP error code and body set to error message.
	//
	// RejectConnectionError could be used to get more control on response.
	Negotiate func(httphead.BIMZTmbYHb) (httphead.BIMZTmbYHb, error)
}

// Upgrade upgrades http connection to the websocket connection.
//
// It hijacks net.Conn from w and returns received net.Conn and
// bufio.ReadWriter. On successful handshake it returns Handshake struct
// describing handshake info.
func (u OKnCFm3eQ) Upgrade(r *http.Request, w http.ResponseWriter) (conn net.Conn, rw *bufio.ReadWriter, hs Handshake, err error) {
	// Hijack connection first to get the ability to write rejection errors the
	// same way as in Upgrader.
	conn, rw, err = gpl0cZhBY5OA(w)
	if err != nil {
		kwu2B80uWa(w, err.Error(), http.StatusInternalServerError)
		return conn, rw, hs, err
	}

	// See https://tools.ietf.org/html/rfc6455#section-4.1
	// The method of the request MUST be GET, and the HTTP version MUST be at least 1.1.
	var nonce string
	if r.Method != http.MethodGet {
		err = ErrHandshakeBadMethod
	} else if r.ProtoMajor < 1 || (r.ProtoMajor == 1 && r.ProtoMinor < 1) {
		err = ErrHandshakeBadProtocol
	} else if r.Host == "" {
		err = ErrHandshakeBadHost
	} else if u := httpGetHeader(r.Header, headerUpgradeCanonical); u != "websocket" && !strings.EqualFold(u, "websocket") {
		err = ErrHandshakeBadUpgrade
	} else if c := httpGetHeader(r.Header, headerConnectionCanonical); c != "Upgrade" && !uEdadz_EwJwD(c, "upgrade") {
		err = ErrHandshakeBadConnection
	} else if nonce = httpGetHeader(r.Header, headerSecKeyCanonical); len(nonce) != nonceSize {
		err = ErrHandshakeBadSecKey
	} else if v := httpGetHeader(r.Header, headerSecVersionCanonical); v != "13" {
		// According to RFC6455:
		//
		// If this version does not match a version understood by the server,
		// the server MUST abort the WebSocket handshake described in this
		// section and instead send an appropriate HTTP error code (such as 426
		// Upgrade Required) and a |Sec-WebSocket-Version| header field
		// indicating the version(s) the server is capable of understanding.
		//
		// So we branching here cause empty or not present version does not
		// meet the ABNF rules of RFC6455:
		//
		// version = DIGIT | (NZDIGIT DIGIT) |
		// ("1" DIGIT DIGIT) | ("2" DIGIT DIGIT)
		// ; Limited to 0-255 range, with no leading zeros
		//
		// That is, if version is really invalid – we sent 426 status, if it
		// not present or empty – it is 400.
		if v != "" {
			err = ErrHandshakeUpgradeRequired
		} else {
			err = ErrHandshakeBadSecVersion
		}
	}
	if check := u.Protocol; err == nil && check != nil {
		ps := r.Header[headerSecProtocolCanonical]
		for i := 0; i < len(ps) && err == nil && hs.Protocol == ""; i++ {
			var ok bool
			hs.Protocol, ok = z4doaSTR2E(ps[i], check)
			if !ok {
				err = ErrMalformedRequest
			}
		}
	}
	if f := u.Negotiate; err == nil && f != nil {
		for _, h := range r.Header[headerSecExtensionsCanonical] {
			hs.Extensions, err = tHZ_mCh(strToBytes(h), hs.Extensions, f)
			if err != nil {
				break
			}
		}
	}
	// DEPRECATED path.
	if check := u.Extension; err == nil && check != nil && u.Negotiate == nil {
		xs := r.Header[headerSecExtensionsCanonical]
		for i := 0; i < len(xs) && err == nil; i++ {
			var ok bool
			hs.Extensions, ok = btsSelectExtensions(strToBytes(xs[i]), hs.Extensions, check)
			if !ok {
				err = ErrMalformedRequest
			}
		}
	}

	// Clear deadlines set by server.
	conn.SetDeadline(noDeadline)
	if t := u.Timeout; t != 0 {
		conn.SetWriteDeadline(time.Now().Add(t))
		defer conn.SetWriteDeadline(noDeadline)
	}

	var header u5HVCPFx
	if h := u.Header; h != nil {
		header[0] = NCYj6EtaAgQ(h)
	}
	if err == nil {
		g3pBIqz(rw.Writer, strToBytes(nonce), hs, header.WriteTo)
		err = rw.Writer.Flush()
	} else {
		var code int
		if rej, ok := err.(*Zbu9jHc); ok {
			code = rej.uIlV5a
			header[1] = rej.oCKugs
		}
		if code == 0 {
			code = http.StatusInternalServerError
		}
		dD1xM5DMpliN(rw.Writer, err, code, header.WriteTo)
		// Do not store Flush() error to not override already existing one.
		_ = rw.Writer.Flush()
	}
	return conn, rw, hs, err
}

type u5HVCPFx [2]HDpTJacC

func (hs u5HVCPFx) WriteTo(w io.Writer) (n int64, err error) {
	for i := 0; i < len(hs) && err == nil; i++ {
		if h := hs[i]; h != nil {
			var m int64
			m, err = h.WriteTo(w)
			n += m
		}
	}
	return n, err
}

// Handshake 是原返回值的 40B 源码表达，值类型名称未存活。
type Handshake struct {
	Protocol   string
	Extensions []httphead.BIMZTmbYHb
}

var noDeadline time.Time
