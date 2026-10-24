// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package low0s_10x

import (
	"bytes"
	"net"

	"recovered.local/soga/mzi1d3gDO7u"
)

type A6KE0r struct {
	net.Conn
	Shdj3DH1
	wt_fjza     bytes.Buffer
	kq485GyMshc bytes.Buffer
}

// (*A6KE0r).Read: 0x1b44f80.
func (c *A6KE0r) Read(b []byte) (int, error) {
	if c.wt_fjza.Len() > 0 {
		return c.wt_fjza.Read(b)
	}

	buf := mzi1d3gDO7u.SharedPoolAt5096978.Get(20480)
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(buf)
	n, err := c.Conn.Read(buf)
	if err != nil {
		return 0, err
	}
	c.kq485GyMshc.Write(buf[:n])
	err = c.Decode(&c.wt_fjza, &c.kq485GyMshc)
	if err != nil {
		return 0, err
	}
	n, _ = c.wt_fjza.Read(b)
	return n, nil
}

// (*A6KE0r).Write: 0x1b45320.
func (c *A6KE0r) Write(b []byte) (int, error) {
	bLength := len(b)
	buf := mzi1d3gDO7u.BytesBufferPoolAt5044220.Get().(*bytes.Buffer)
	defer mzi1d3gDO7u.LjZLCz9(buf)
	err := c.Encode(buf, b)
	if err != nil {
		return 0, err
	}
	_, err = c.Conn.Write(buf.Bytes())
	if err != nil {
		return 0, err
	}
	return bLength, nil
}
