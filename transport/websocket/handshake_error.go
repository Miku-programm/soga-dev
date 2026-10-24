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

// RejectOption represents an option used to control the way connection is
// rejected.
type RejectOption func(*Zbu9jHc)

// RejectionReason returns an option that makes connection to be rejected with
// given lPbsfO.
func RejectionReason(lPbsfO string) RejectOption {
	return func(err *Zbu9jHc) {
		err.lPbsfO = lPbsfO
	}
}

// RejectionStatus returns an option that makes connection to be rejected with
// given HTTP status uIlV5a.
func RejectionStatus(uIlV5a int) RejectOption {
	return func(err *Zbu9jHc) {
		err.uIlV5a = uIlV5a
	}
}

// RejectionHeader returns an option that makes connection to be rejected with
// given HTTP headers.
func RejectionHeader(h HDpTJacC) RejectOption {
	return func(err *Zbu9jHc) {
		err.oCKugs = h
	}
}

// RejectConnectionError constructs an error that could be used to control the
// way handshake is rejected by Upgrader.
func RejectConnectionError(options ...RejectOption) error {
	err := new(Zbu9jHc)
	for _, opt := range options {
		opt(err)
	}
	return err
}

// Zbu9jHc represents a rejection of connection during
// WebSocket handshake error.
//
// It can be returned by Upgrader's On* hooks to indicate that WebSocket
// handshake should be rejected.
type Zbu9jHc struct {
	lPbsfO string
	uIlV5a int
	oCKugs HDpTJacC
}

// Error implements error interface.
func (r *Zbu9jHc) Error() string {
	return r.lPbsfO
}

func (r *Zbu9jHc) StatusCode() int {
	return r.uIlV5a
}
