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
	"net/url"
	httphead "recovered.local/soga/EsS8DL8PE9"
)

type httpResponseLine struct {
	major, minor int
	status       int
	reason       []byte
}

// Original VA 0x11ab280.
func l8NFoELj(line []byte) (resp httpResponseLine, err error) {
	var (
		proto  []byte
		status []byte
	)
	proto, status, resp.reason = cU_pjSUO(line, ' ')

	var ok bool
	resp.major, resp.minor, ok = lJiBRXY(proto)
	if !ok {
		return resp, ErrMalformedResponse
	}

	var convErr error
	resp.status, convErr = fm9r93qPp(status)
	if convErr != nil {
		return resp, ErrMalformedResponse
	}

	return resp, nil
}

// Original VA 0x11abf60.
func bGww_wdToRn9(bw *bufio.Writer, key string, value []byte) {
	pyi17ZDrONKV(bw, key)
	bw.Write(value)
	bw.WriteString(crlf)
}

// Original VA 0x11ac080.
func eLiGoGo1nj(
	bw *bufio.Writer,
	u *url.URL,
	nonce []byte,
	protocols []string,
	extensions []httphead.BIMZTmbYHb, header HDpTJacC,
	host string,
) {
	bw.WriteString("GET ")
	bw.WriteString(u.RequestURI())
	bw.WriteString(" HTTP/1.1\r\n")

	if host == "" {
		host = u.Host
	}
	pU61Wak(bw, headerHost, host)

	bGww_wdToRn9(bw, headerUpgrade, specHeaderValueUpgrade)
	bGww_wdToRn9(bw, headerConnection, specHeaderValueConnection)
	bGww_wdToRn9(bw, headerSecVersion, specHeaderValueSecVersion)

	pU61Wak(bw, headerSecKey, btsToString(nonce))

	if len(protocols) > 0 {
		pyi17ZDrONKV(bw, headerSecProtocol)
		for i, p := range protocols {
			if i > 0 {
				bw.WriteString(commaAndSpace)
			}
			bw.WriteString(p)
		}
		bw.WriteString(crlf)
	}

	if len(extensions) > 0 {
		pyi17ZDrONKV(bw, headerSecExtensions)
		httphead.KBgYhFImpuOF(bw, extensions)
		bw.WriteString(crlf)
	}

	if header != nil {
		header.WriteTo(bw)
	}

	bw.WriteString(crlf)
}
