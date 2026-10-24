package FLhkKwJ8M

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"net"
	"time"
)

// mWynszCaKB retains the original 80-byte RTTI layout. The original has
// no additional locks around first Dial or its closed flag.
type mWynszCaKB struct {
	net.Conn
	xwg3XpzNo    net.Conn
	bCzUlUxzMr7  bool
	maRCahwMPX   chan bool
	dW68su       context.CancelFunc
	_BITXBt_ajlG context.Context
	mcbe7zks     *ELIQWabl31
}

// uTCy4V: 0x19bd3c0. Handshake is deferred until the first Write.
func uTCy4V(conn net.Conn, config *ELIQWabl31) (net.Conn, error) {
	ctx, cancel := context.WithCancel(context.Background())
	return &mWynszCaKB{xwg3XpzNo: conn, maRCahwMPX: make(chan bool, 1), dW68su: cancel, _BITXBt_ajlG: ctx, mcbe7zks: config}, nil
}

// Dial: 0x19bc9c0. RawURLEncoding is confirmed by base64.init@0x72a7c0's
// global 0x5095fd0 assignment (URL alphabet, padding -1).
func (c *mWynszCaKB) Dial(payload []byte) error {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.RawURLEncoding, &encoded)
	n := c.mcbe7zks.A3jvUyhB
	if len(payload) < n {
		n = len(payload)
	}
	if _, err := encoder.Write(payload[:n]); err != nil {
		return errors.New("failed to encode early data: " + err.Error())
	}
	if err := encoder.Close(); err != nil {
		return errors.New("failed to encode early data tail: " + err.Error())
	}
	var err error
	c.Conn, err = caPKDWZ1qaY9(c.xwg3XpzNo, c.mcbe7zks, &encoded)
	if err != nil {
		_ = c.Close()
		return errors.New("failed to dial WebSocket: " + err.Error())
	}
	c.maRCahwMPX <- true
	if n == len(payload) {
		return nil
	}
	_, err = c.Conn.Write(payload[n:])
	return err
}

// Write/Read: 0x19bcec0/0x19bcf80. Cancellation while waiting for Dial
// returns io.ErrUnexpectedEOF; an already-closed call returns ErrClosedPipe.
func (c *mWynszCaKB) Write(p []byte) (int, error) {
	if c.bCzUlUxzMr7 {
		return 0, io.ErrClosedPipe
	}
	if c.Conn != nil {
		return c.Conn.Write(p)
	}
	if err := c.Dial(p); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (c *mWynszCaKB) Read(p []byte) (int, error) {
	if c.bCzUlUxzMr7 {
		return 0, io.ErrClosedPipe
	}
	if c.Conn != nil {
		return c.Conn.Read(p)
	}
	select {
	case <-c.maRCahwMPX:
		return c.Conn.Read(p)
	case <-c._BITXBt_ajlG.Done():
		return 0, io.ErrUnexpectedEOF
	}
}

// Close: 0x19bd0c0. Before Dial, the retained raw connection is not closed.
func (c *mWynszCaKB) Close() error {
	c.bCzUlUxzMr7 = true
	c.dW68su()
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}

// Addresses: 0x19bd120/0x19bd180; deadlines: 0x19bd1e0..0x19bd340.
func (c *mWynszCaKB) LocalAddr() net.Addr {
	if c.Conn != nil {
		return c.Conn.LocalAddr()
	}
	return c.xwg3XpzNo.LocalAddr()
}
func (c *mWynszCaKB) RemoteAddr() net.Addr {
	if c.Conn != nil {
		return c.Conn.RemoteAddr()
	}
	return c.xwg3XpzNo.RemoteAddr()
}
func (c *mWynszCaKB) SetDeadline(t time.Time) error {
	if err := c.SetReadDeadline(t); err != nil {
		return err
	}
	return c.SetWriteDeadline(t)
}
func (c *mWynszCaKB) SetReadDeadline(t time.Time) error {
	if c.Conn != nil {
		return c.Conn.SetReadDeadline(t)
	}
	return nil
}
func (c *mWynszCaKB) SetWriteDeadline(t time.Time) error {
	if c.Conn != nil {
		return c.Conn.SetWriteDeadline(t)
	}
	return nil
}

var _ net.Conn = (*mWynszCaKB)(nil)
