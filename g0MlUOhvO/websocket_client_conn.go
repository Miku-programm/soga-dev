package g0MlUOhvO

import (
	pkgerrors "github.com/pkg/errors"
	"io"
	"net"
	ws "recovered.local/soga/FgcbiePIJLm"
	util "recovered.local/soga/UYoO4Hz"
	"sync"
)

// RTTI bhXMvjeYm_m.F2Eb1b @0x36f0938: 56 bytes, original field order.
// This client wrapper has its own close-once semantics, unlike V72bzf8.
type F2Eb1b struct {
	DrpXao
	hESODkc8p  io.Reader
	jsKotLrECn util.KTXwG8eroaXF
	tqzpXGhPgw sync.Once
}

// fAbNA17x_Ire restores 0x11e2ee0 with client-side frame validation.
func (c *F2Eb1b) fAbNA17x_Ire() (io.Reader, error) {
	if c.hESODkc8p != nil {
		return c.hESODkc8p, nil
	}
	r := &util.HLKO3Cdqw{GX04GFKh1Gz: c.DrpXao, WXUh5psqY7: ws.StateClientSide, EbVEHBvY4lh: c.jsKotLrECn}
	for {
		h, e := r.NextFrame()
		if e != nil {
			return nil, e
		}
		if h.IsHe8C&8 == 0 {
			break
		}
		if e = c.jsKotLrECn(h, r); e != nil {
			return nil, e
		}
	}
	c.hESODkc8p = r
	return r, nil
}

// Read restores 0x11e30c0, including pkg/errors.Cause rather than errors.Is.
func (c *F2Eb1b) Read(p []byte) (int, error) {
	for {
		r, e := c.fAbNA17x_Ire()
		if e != nil {
			return 0, e
		}
		n, e := r.Read(p)
		if pkgerrors.Cause(e) != io.EOF {
			return n, e
		}
		c.hESODkc8p = nil
		if n > 0 {
			return n, nil
		}
	}
}

// Write restores 0x11e31e0: one masked binary message, zero count on error.
func (c *F2Eb1b) Write(p []byte) (int, error) {
	if e := util.WriteFrameRecovered(c.DrpXao, ws.StateClientSide, ws.OpBinary, true, p); e != nil {
		return 0, e
	}
	return len(p), nil
}

// Close restores 0x11e3320 and closure 0x11e6020. It sends an empty close
// frame once, ignores frame-write errors, and returns the underlying Close
// error only on the first invocation; later invocations return nil.
func (c *F2Eb1b) Close() (err error) {
	c.tqzpXGhPgw.Do(func() {
		_ = util.WriteFrameRecovered(c.DrpXao, ws.StateClientSide, ws.OpClose, true, []byte{})
		err = c.DrpXao.Close()
	})
	return
}
func (c *F2Eb1b) GetRawConn() net.Conn { return c.DrpXao } //0x11e33a0
