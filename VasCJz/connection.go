package VasCJz

import (
	"io"
	"net"
	Q "recovered.local/soga/qcqMSUWwIx40"
	"time"
)

// 原 RTTI 80 字节，0x125a0e0—0x125a360 的 XHTTP 流连接。
type _0_9jH struct {
	scAGARMF0XsA io.WriteCloser
	xxabacttFOn  io.ReadCloser
	pcYQQ0       *Q.Q5KYzU
	n6EztFBWOT   net.Addr
	gSBX2NrxZ_8y net.Addr
	gvvJIQwQyA   func()
}

func (c *_0_9jH) Write(p []byte) (int, error) { return c.scAGARMF0XsA.Write(p) }
func (c *_0_9jH) Read(p []byte) (int, error)  { return c.pcYQQ0.Read(p) }
func (c *_0_9jH) Close() error {
	if c.gvvJIQwQyA != nil {
		c.gvvJIQwQyA()
	}
	writeErr := c.scAGARMF0XsA.Close()
	readErr := c.xxabacttFOn.Close()
	if writeErr != nil {
		return writeErr
	}
	return readErr
}
func (c *_0_9jH) LocalAddr() net.Addr  { return c.gSBX2NrxZ_8y }
func (c *_0_9jH) RemoteAddr() net.Addr { return c.n6EztFBWOT }

// 样本的 SetDeadline 只设置读超时；写超时接口直接成功返回。
func (c *_0_9jH) SetDeadline(t time.Time) error     { return c.pcYQQ0.SetReadDeadline(t) }
func (c *_0_9jH) SetReadDeadline(t time.Time) error { return c.pcYQQ0.SetReadDeadline(t) }
func (c *_0_9jH) SetWriteDeadline(time.Time) error  { return nil }

var _ net.Conn = (*_0_9jH)(nil)
