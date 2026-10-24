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
// 由原 ELF/IDA 与 gobwas/ws v1.4.0 对照恢复，非作者源码原文。
// 保留原 RTTI 名称；内联辅助表达式使用恢复名称，不计独立原函数。
package UYoO4Hz

import (
	"errors"
	"io"
	"io/ioutil"
	"strconv"

	ws "recovered.local/soga/FgcbiePIJLm"
	pbytes "recovered.local/soga/TIhkFUX1kZ"
)

// VGrYaNsp returned when peer has closed the connection with appropriate
// code and a textual reason.
type VGrYaNsp struct {
	EAGe4eFzGle ws.K9jDRT9Z
	Nd7HXs1GwPQ string
}

// Error implements error interface.
func (err VGrYaNsp) Error() string {
	return "ws closed: " + strconv.FormatUint(uint64(err.EAGe4eFzGle), 10) + " " + err.Nd7HXs1GwPQ
}

// GdvjKg5RudlP contains logic of handling control frames.
//
// The intentional way to use it is to read the next frame header from the
// connection, optionally check its validity via ws.IYI3RQDfj4() and if it is
// not a ws.OpText of ws.OpBinary (or ws.OpContinuation) – pass it to Handle()
// method.
//
// That is, passed header should be checked to get rid of unexpected errors.
//
// The Handle() method will read out all control frame payload (if any) and
// write necessary bytes as a rfc compatible response.
// 值类型的字段原名未保留；Src/Dst/State/DisableSrcCiphering 对应原偏移 0/16/32/33。
type GdvjKg5RudlP struct {
	Src   io.Reader
	Dst   io.Writer
	State ws.DwfKMV

	// DisableSrcCiphering disables unmasking payload data read from Src.
	// It is useful when UYoO4Hz.Reader is used or when frame payload already
	// pulled and ciphered out from the connection (and introduced by
	// bytes.Reader, for example).
	DisableSrcCiphering bool
}

// ErrNotControlFrame is returned by GdvjKg5RudlP to indicate that given
// header could not be handled.
var ErrNotControlFrame = errors.New("not a control frame")

// Handle handles control frames regarding to the c.State and writes responses
// to the c.Dst when needed.
//
// It returns ErrNotControlFrame when given header is not of ws.OpClose,
// ws.OpPing or ws.OpPong operation code.
func (c GdvjKg5RudlP) Handle(h ws.Wmyq5K0sQ0w) error {
	switch h.IsHe8C {
	case ws.OpPing:
		return c.HandlePing(h)
	case ws.OpPong:
		return c.HandlePong(h)
	case ws.OpClose:
		return c.HandleClose(h)
	}
	return ErrNotControlFrame
}

// HandlePing handles ping frame and writes specification compatible response
// to the c.Dst.
func (c GdvjKg5RudlP) HandlePing(h ws.Wmyq5K0sQ0w) error {
	if h.L4ixlgT8VpTo == 0 {
		// The most common case when ping is empty.
		// Note that when sending masked frame the mask for empty payload is
		// just four zero bytes.
		return ws.KzLwHzhWB90h(c.Dst, ws.Wmyq5K0sQ0w{
			HcIhwPLvD:    true,
			IsHe8C:       ws.OpPong,
			X4IS_ER04TVv: c.State.ClientSide(),
		})
	}

	// In other way reply with Pong frame with copied payload.
	p := pbytes.GetLen(int(h.L4ixlgT8VpTo) + ws.HeaderSize(ws.Wmyq5K0sQ0w{
		L4ixlgT8VpTo: h.L4ixlgT8VpTo,
		X4IS_ER04TVv: c.State.ClientSide(),
	}))
	defer pbytes.KqwWvOwoRkN(p)

	// Deal with ciphering i/o:
	// Masking key is used to mask the "Payload data" defined in the same
	// section as frame-payload-data, which includes "Extension data" and
	// "Application data".
	//
	// See https://tools.ietf.org/html/rfc6455#section-5.3
	//
	// NOTE: We prefer DH32d4Xf with preallocated buffer to
	// ws.KzLwHzhWB90h because it performs one syscall instead of two.
	w := ZB0fR0Aw(c.Dst, c.State, ws.OpPong, p)
	r := c.Src
	if c.State.ServerSide() && !c.DisableSrcCiphering {
		r = &T3e3P4wOxooC{zGmeMGooB: r, xVfRjn07o: h.BiCmXtyECQl}
	}

	_, err := io.Copy(w, r)
	if err == nil {
		err = w.Flush()
	}

	return err
}

// HandlePong handles pong frame by discarding it.
func (c GdvjKg5RudlP) HandlePong(h ws.Wmyq5K0sQ0w) error {
	if h.L4ixlgT8VpTo == 0 {
		return nil
	}

	buf := pbytes.GetLen(int(h.L4ixlgT8VpTo))
	defer pbytes.KqwWvOwoRkN(buf)

	// Discard pong message according to the RFC6455:
	// A Pong frame MAY be sent unsolicited. This serves as a
	// unidirectional heartbeat. A response to an unsolicited Pong frame
	// is not expected.
	_, err := io.CopyBuffer(ioutil.Discard, c.Src, buf)

	return err
}

// HandleClose handles close frame, makes protocol validity checks and writes
// specification compatible response to the c.Dst.
func (c GdvjKg5RudlP) HandleClose(h ws.Wmyq5K0sQ0w) error {
	if h.L4ixlgT8VpTo == 0 {
		err := ws.KzLwHzhWB90h(c.Dst, ws.Wmyq5K0sQ0w{
			HcIhwPLvD:    true,
			IsHe8C:       ws.OpClose,
			X4IS_ER04TVv: c.State.ClientSide(),
		})
		if err != nil {
			return err
		}

		// Due to RFC, we should interpret the code as no status code
		// received:
		//   If this Close control frame contains no status code, _The WebSocket
		//   Connection Close Code_ is considered to be 1005.
		//
		// See https://tools.ietf.org/html/rfc6455#section-7.1.5
		return VGrYaNsp{
			EAGe4eFzGle: ws.StatusNoStatusRcvd,
		}
	}

	// Prepare bytes both for reading reason and sending response.
	p := pbytes.GetLen(int(h.L4ixlgT8VpTo) + ws.HeaderSize(ws.Wmyq5K0sQ0w{
		L4ixlgT8VpTo: h.L4ixlgT8VpTo,
		X4IS_ER04TVv: c.State.ClientSide(),
	}))
	defer pbytes.KqwWvOwoRkN(p)

	// Get the subslice to read the frame payload out.
	subp := p[:h.L4ixlgT8VpTo]

	r := c.Src
	if c.State.ServerSide() && !c.DisableSrcCiphering {
		r = &T3e3P4wOxooC{zGmeMGooB: r, xVfRjn07o: h.BiCmXtyECQl}
	}
	if _, err := io.ReadFull(r, subp); err != nil {
		return err
	}

	code, reason := ws.ParseCloseFrameData(subp)
	if err := ws.QXKFKa(code, reason); err != nil {
		// Here we could not use the prepared bytes because there is no
		// guarantee that it may fit our protocol error closure code and a
		// reason.
		c.fVfEF8ExV2MD(err)
		return err
	}

	// Deal with ciphering i/o:
	// Masking key is used to mask the "Payload data" defined in the same
	// section as frame-payload-data, which includes "Extension data" and
	// "Application data".
	//
	// See https://tools.ietf.org/html/rfc6455#section-5.3
	//
	// NOTE: We prefer DH32d4Xf with preallocated buffer to
	// ws.KzLwHzhWB90h because it performs one syscall instead of two.
	w := ZB0fR0Aw(c.Dst, c.State, ws.OpClose, p)

	// RFC6455#5.5.1:
	// If an endpoint receives a Close frame and did not previously
	// send a Close frame, the endpoint MUST send a Close frame in
	// response. (When sending a Close frame in response, the endpoint
	// typically echoes the status code it received.)
	_, err := w.Write(p[:2])
	if err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return VGrYaNsp{
		EAGe4eFzGle: code,
		Nd7HXs1GwPQ: reason,
	}
}

func (c GdvjKg5RudlP) fVfEF8ExV2MD(reason error) error {
	f := ws.NewCloseFrame(ws.NewCloseFrameBody(
		ws.StatusProtocolError, reason.Error(),
	))
	if c.State.ClientSide() {
		ws.GU1uDCm70(f)
	}
	return ws.VzPOf1_5QN(c.Dst, f)
}
