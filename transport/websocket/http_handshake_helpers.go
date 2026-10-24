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
	"bytes"
	"io"
	"net/http"
	"strconv"

	httphead "recovered.local/soga/transport/httphead"
)

const (
	crlf          = "\r\n"
	colonAndSpace = ": "
	commaAndSpace = ", "
)

const (
	textHeadUpgrade = "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\n"
)

var (
	textHeadBadRequest          = bqGmr5c44XB(http.StatusBadRequest)
	textHeadInternalServerError = bqGmr5c44XB(http.StatusInternalServerError)
	textHeadUpgradeRequired     = bqGmr5c44XB(http.StatusUpgradeRequired)

	textTailErrHandshakeBadProtocol   = rmCwpJZyvsoh(ErrHandshakeBadProtocol)
	textTailErrHandshakeBadMethod     = rmCwpJZyvsoh(ErrHandshakeBadMethod)
	textTailErrHandshakeBadHost       = rmCwpJZyvsoh(ErrHandshakeBadHost)
	textTailErrHandshakeBadUpgrade    = rmCwpJZyvsoh(ErrHandshakeBadUpgrade)
	textTailErrHandshakeBadConnection = rmCwpJZyvsoh(ErrHandshakeBadConnection)
	textTailErrHandshakeBadSecAccept  = rmCwpJZyvsoh(ErrHandshakeBadSecAccept)
	textTailErrHandshakeBadSecKey     = rmCwpJZyvsoh(ErrHandshakeBadSecKey)
	textTailErrHandshakeBadSecVersion = rmCwpJZyvsoh(ErrHandshakeBadSecVersion)
	textTailErrUpgradeRequired        = rmCwpJZyvsoh(ErrHandshakeUpgradeRequired)
)

const (
	headerHost          = "Host"
	headerUpgrade       = "Upgrade"
	headerConnection    = "Connection"
	headerSecVersion    = "Sec-WebSocket-Version"
	headerSecProtocol   = "Sec-WebSocket-Protocol"
	headerSecExtensions = "Sec-WebSocket-Extensions"
	headerSecKey        = "Sec-WebSocket-Key"
	headerSecAccept     = "Sec-WebSocket-Accept"

	headerHostCanonical          = headerHost
	headerUpgradeCanonical       = headerUpgrade
	headerConnectionCanonical    = headerConnection
	headerSecVersionCanonical    = "Sec-Websocket-Version"
	headerSecProtocolCanonical   = "Sec-Websocket-Protocol"
	headerSecExtensionsCanonical = "Sec-Websocket-Extensions"
	headerSecKeyCanonical        = "Sec-Websocket-Key"
	headerSecAcceptCanonical     = "Sec-Websocket-Accept"
)

var (
	specHeaderValueUpgrade         = []byte("websocket")
	specHeaderValueConnection      = []byte("Upgrade")
	specHeaderValueConnectionLower = []byte("upgrade")
	specHeaderValueSecVersion      = []byte("13")
)

var (
	httpVersion1_0    = []byte("HTTP/1.0")
	httpVersion1_1    = []byte("HTTP/1.1")
	httpVersionPrefix = []byte("HTTP/")
)

func httpGetHeader(h http.Header, key string) string {
	if h == nil {
		return ""
	}
	v := h[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func z4doaSTR2E(h string, check func(string) bool) (ret string, ok bool) {
	ok = httphead.PgwXhS7T(strToBytes(h), func(v []byte) bool {
		if check(btsToString(v)) {
			ret = string(v)
			return false
		}
		return true
	})
	return ret, ok
}

func btsSelectExtensions(h []byte, selected []httphead.BIMZTmbYHb, check func(httphead.BIMZTmbYHb) bool) ([]httphead.BIMZTmbYHb, bool) {
	s := httphead.ExK8bKI2Xox{
		Flags: httphead.SelectCopy,
		Check: check,
	}
	return s.Select(h, selected)
}

func negotiateMaybe(in httphead.BIMZTmbYHb, dest []httphead.BIMZTmbYHb, f func(httphead.BIMZTmbYHb) (httphead.BIMZTmbYHb, error)) ([]httphead.BIMZTmbYHb, error) {
	if in.Size() == 0 {
		return dest, nil
	}
	opt, err := f(in)
	if err != nil {
		return nil, err
	}
	if opt.Size() > 0 {
		dest = append(dest, opt)
	}
	return dest, nil
}

func tHZ_mCh(
	h []byte, dest []httphead.BIMZTmbYHb,
	f func(httphead.BIMZTmbYHb) (httphead.BIMZTmbYHb, error),
) (_ []httphead.BIMZTmbYHb, err error) {
	index := -1
	var current httphead.BIMZTmbYHb
	ok := httphead.LU7plsn8k(h, func(i int, name, attr, val []byte) httphead.Control {
		if i != index {
			dest, err = negotiateMaybe(current, dest, f)
			if err != nil {
				return httphead.ControlBreak
			}
			index = i
			current = httphead.BIMZTmbYHb{KXa0Kt2I: name}
		}
		if attr != nil {
			current.Jd2j6aD7.Set(attr, val)
		}
		return httphead.ControlContinue
	})
	if !ok {
		return nil, ErrMalformedRequest
	}
	return negotiateMaybe(current, dest, f)
}

func pU61Wak(bw *bufio.Writer, key, value string) {
	pyi17ZDrONKV(bw, key)
	bw.WriteString(value)
	bw.WriteString(crlf)
}

func pyi17ZDrONKV(bw *bufio.Writer, key string) {
	bw.WriteString(key)
	bw.WriteString(colonAndSpace)
}

func g3pBIqz(bw *bufio.Writer, nonce []byte, hs Handshake, header HandshakeHeaderFunc) {
	bw.WriteString(textHeadUpgrade)

	pyi17ZDrONKV(bw, headerSecAccept)
	oExaw8(bw, nonce)
	bw.WriteString(crlf)

	if hs.Protocol != "" {
		pU61Wak(bw, headerSecProtocol, hs.Protocol)
	}
	if len(hs.Extensions) > 0 {
		pyi17ZDrONKV(bw, headerSecExtensions)
		httphead.KBgYhFImpuOF(bw, hs.Extensions)
		bw.WriteString(crlf)
	}
	if header != nil {
		header(bw)
	}

	bw.WriteString(crlf)
}

func dD1xM5DMpliN(bw *bufio.Writer, err error, code int, header HandshakeHeaderFunc) {
	switch code {
	case http.StatusBadRequest:
		bw.WriteString(textHeadBadRequest)
	case http.StatusInternalServerError:
		bw.WriteString(textHeadInternalServerError)
	case http.StatusUpgradeRequired:
		bw.WriteString(textHeadUpgradeRequired)
	default:
		e0cU5fypxOv(bw, code)
	}

	if header != nil {
		header(bw)
	}

	switch err {
	case ErrHandshakeBadProtocol:
		bw.WriteString(textTailErrHandshakeBadProtocol)
	case ErrHandshakeBadMethod:
		bw.WriteString(textTailErrHandshakeBadMethod)
	case ErrHandshakeBadHost:
		bw.WriteString(textTailErrHandshakeBadHost)
	case ErrHandshakeBadUpgrade:
		bw.WriteString(textTailErrHandshakeBadUpgrade)
	case ErrHandshakeBadConnection:
		bw.WriteString(textTailErrHandshakeBadConnection)
	case ErrHandshakeBadSecAccept:
		bw.WriteString(textTailErrHandshakeBadSecAccept)
	case ErrHandshakeBadSecKey:
		bw.WriteString(textTailErrHandshakeBadSecKey)
	case ErrHandshakeBadSecVersion:
		bw.WriteString(textTailErrHandshakeBadSecVersion)
	case ErrHandshakeUpgradeRequired:
		bw.WriteString(textTailErrUpgradeRequired)
	case nil:
		bw.WriteString(crlf)
	default:
		vksURXlT(bw, err)
	}
}

func e0cU5fypxOv(bw *bufio.Writer, code int) {
	bw.WriteString("HTTP/1.1 ")
	bw.WriteString(strconv.Itoa(code))
	bw.WriteByte(' ')
	bw.WriteString(http.StatusText(code))
	bw.WriteString(crlf)
	bw.WriteString("Content-Type: text/plain; charset=utf-8")
	bw.WriteString(crlf)
}

func vksURXlT(bw *bufio.Writer, err error) {
	body := err.Error()
	bw.WriteString("Content-Length: ")
	bw.WriteString(strconv.Itoa(len(body)))
	bw.WriteString(crlf)
	bw.WriteString(crlf)
	bw.WriteString(body)
}

func kwu2B80uWa(w http.ResponseWriter, body string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(code)
	w.Write([]byte(body))
}

func bqGmr5c44XB(code int) string {
	var buf bytes.Buffer
	bw := bufio.NewWriter(&buf)
	e0cU5fypxOv(bw, code)
	bw.Flush()
	return buf.String()
}

func rmCwpJZyvsoh(err error) string {
	var buf bytes.Buffer
	bw := bufio.NewWriter(&buf)
	vksURXlT(bw, err)
	bw.Flush()
	return buf.String()
}

type HDpTJacC interface {
	io.WriterTo
}

type BuqwQfCAXz string

func (s BuqwQfCAXz) WriteTo(w io.Writer) (int64, error) {
	n, err := io.WriteString(w, string(s))
	return int64(n), err
}

type HandshakeHeaderFunc func(io.Writer) (int64, error)

type NCYj6EtaAgQ http.Header

func (h NCYj6EtaAgQ) WriteTo(w io.Writer) (int64, error) {
	wr := jC7sxat{h7OgHq0GODs: w}
	err := http.Header(h).Write(&wr)
	return wr.zqxPTJEN, err
}

type jC7sxat struct {
	zqxPTJEN    int64
	h7OgHq0GODs io.Writer
}

func (w *jC7sxat) WriteString(s string) (int, error) {
	n, err := io.WriteString(w.h7OgHq0GODs, s)
	w.zqxPTJEN += int64(n)
	return n, err
}

func (w *jC7sxat) Write(p []byte) (int, error) {
	n, err := w.h7OgHq0GODs.Write(p)
	w.zqxPTJEN += int64(n)
	return n, err
}
