package pcnXtbh8Z5p

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	S "recovered.local/soga/Bx4XCl"
	H "recovered.local/soga/Hx04IbKk"
	"sync"
)

// Mah1Kab retains the 48-byte RTTI layout. Embedded DrpXao maps to net.Conn.
type Mah1Kab struct {
	net.Conn
	iFfzJtf3G  int
	bsHi8N0rMT net.Addr
	i7IOAZNu   sync.Mutex
}

// sXUNG8z: 0x1b51e80. The write count includes address/length/CRLF bytes.
func sXUNG8z(w io.Writer, address, payload []byte) (int, error) {
	b := H.BaDekw7qe{}
	appendBuffer(&b, address)
	i := b.Grow(2)
	binary.BigEndian.PutUint16(b[i:], uint16(len(payload)))
	appendBuffer(&b, crlfAt5037D70)
	appendBuffer(&b, payload)
	return w.Write(b)
}

// AwkQMyI: 0x1b52100. Payloads over 8192 bytes are split into independent
// records. For that branch only, total success returns the payload length.
func AwkQMyI(w io.Writer, address, payload []byte) (int, error) {
	if len(payload) <= 8192 {
		return sXUNG8z(w, address, payload)
	}
	for offset := 0; offset < len(payload); {
		end := offset + 8192
		if end > len(payload) {
			end = len(payload)
		}
		n, err := sXUNG8z(w, address, payload[offset:end])
		if err != nil {
			return offset + n, err
		}
		offset = end
	}
	return len(payload), nil
}

// WroRKupgG3k7: 0x1b522a0. The caller's buffer is reused for the address,
// length, unchecked two-byte delimiter and payload. Underlying errors are
// replaced with the original stage-specific errors, not wrapped.
func WroRKupgG3k7(r io.Reader, b []byte) (net.Addr, int, int, error) {
	address, err := S.YhN2Ke2lC(r, b)
	if err != nil {
		return nil, 0, 0, errors.New("read addr error")
	}
	var addr *net.UDPAddr
	if len(address) > 0 {
		switch address[0] {
		case 1:
			ip := make(net.IP, 4)
			copy(ip, address[1:5])
			addr = &net.UDPAddr{IP: ip, Port: int(binary.BigEndian.Uint16(address[5:7]))}
		case 4:
			ip := make(net.IP, 16)
			copy(ip, address[1:17])
			addr = &net.UDPAddr{IP: ip, Port: int(binary.BigEndian.Uint16(address[17:19]))}
		}
	}
	if addr == nil {
		return nil, 0, 0, errors.New("parse addr error")
	}
	if _, err = io.ReadFull(r, b[:2]); err != nil {
		return nil, 0, 0, errors.New("read length error")
	}
	length := int(binary.BigEndian.Uint16(b[:2]))
	if length > 8192 {
		return nil, 0, 0, errors.New("packet invalid")
	}
	if _, err = io.ReadFull(r, b[:2]); err != nil {
		return nil, 0, 0, errors.New("read crlf error")
	}
	n := len(b)
	if n > length {
		n = length
	}
	if _, err = io.ReadFull(r, b[:n]); err != nil {
		return nil, 0, 0, errors.New("read packet error")
	}
	return addr, n, length - n, nil
}

// WriteTo: 0x1b529c0. Uses address.String(), not the net.Addr encoder.
func (c *Mah1Kab) WriteTo(b []byte, address net.Addr) (int, error) {
	return AwkQMyI(c, S.L4JRaRodkw(address.String()), b)
}

// ReadFrom: 0x1b52a80. A partial datagram is retained across calls, with a
// mutex covering both header and remainder reads. A read with an error returns
// zero and leaves the original remainder counter unchanged.
func (c *Mah1Kab) ReadFrom(b []byte) (int, net.Addr, error) {
	c.i7IOAZNu.Lock()
	defer c.i7IOAZNu.Unlock()
	if c.iFfzJtf3G != 0 {
		if len(b) > c.iFfzJtf3G {
			b = b[:c.iFfzJtf3G]
		}
		n, err := c.Conn.Read(b)
		if err != nil {
			return 0, nil, err
		}
		c.iFfzJtf3G -= n
		addr := c.bsHi8N0rMT
		if c.iFfzJtf3G == 0 {
			c.bsHi8N0rMT = nil
		}
		return n, addr, nil
	}
	addr, n, left, err := WroRKupgG3k7(c.Conn, b)
	if err != nil {
		return 0, nil, err
	}
	if left != 0 {
		c.iFfzJtf3G = left
		c.bsHi8N0rMT = addr
	}
	return n, addr, nil
}

var _ net.PacketConn = (*Mah1Kab)(nil)
