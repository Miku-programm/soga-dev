// Restored against IDA 0x1b28820 through 0x1b2a480.
// Source comparison: github.com/Ieooo/clash/transport/ssr/obfs (GPL-3.0).
// Original dependency revision is unproven; names/layout and deviations follow target evidence.
package hysteriaconn

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"encoding/binary"
	"io"
	mathRand "math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"recovered.local/soga/transport/tlsutil"
	"recovered.local/soga/buf/syncpool"
)

func init() {
	register("tls1.2_ticket_auth", aUNXJlk, 5)
	register("tls1.2_ticket_fastauth", aUNXJlk, 5)
}

type xIbCZlK struct {
	*Ja7FBino7HL
	*sFaLAK
}

func aUNXJlk(b *Ja7FBino7HL) AGNBYcIkk {
	r := &xIbCZlK{Ja7FBino7HL: b, sFaLAK: &sFaLAK{}}
	rand.Read(r.l57vr8H[:])
	return r
}

type cFCr3rhcmDX struct {
	net.Conn
	*xIbCZlK
	awCIY1      sync.Mutex
	vpScqu      int
	joNPsEaf    bytes.Buffer
	zJ742e_XsaS bytes.Buffer
	dhQYYCsc    bytes.Buffer
}

func (t *xIbCZlK) StreamConn(c net.Conn) net.Conn {
	return &cFCr3rhcmDX{Conn: c, xIbCZlK: t}
}

func (c *cFCr3rhcmDX) Read(b []byte) (int, error) {
	if c.joNPsEaf.Len() > 0 {
		return c.joNPsEaf.Read(b)
	}

	tbgQNW := syncpool.SharedPoolAt5096978.Get(20480)
	defer syncpool.SharedPoolAt5096978.Put(tbgQNW)
	n, err := c.Conn.Read(tbgQNW)
	if err != nil {
		return 0, err
	}

	if c.vpScqu == 8 {
		c.zJ742e_XsaS.Write(tbgQNW[:n])
		for c.zJ742e_XsaS.Len() > 5 {
			if !bytes.Equal(c.zJ742e_XsaS.Bytes()[:3], []byte{0x17, 3, 3}) {
				c.zJ742e_XsaS.Reset()
				return 0, errorAt5098560
			}
			size := int(binary.BigEndian.Uint16(c.zJ742e_XsaS.Bytes()[3:5]))
			if c.zJ742e_XsaS.Len() < 5+size {
				break
			}
			c.zJ742e_XsaS.Next(5)
			c.joNPsEaf.Write(c.zJ742e_XsaS.Next(size))
		}
		n, _ = c.joNPsEaf.Read(b)
		return n, nil
	}

	if n < 11+32+1+32 {
		return 0, errorAt5098570
	}

	if !hmac.Equal(tbgQNW[33:43], c.cC9t2Xktic(tbgQNW[11:33])[:10]) || !hmac.Equal(tbgQNW[n-10:n], c.cC9t2Xktic(tbgQNW[:n-10])[:10]) {
		return 0, errorAt5098580
	}

	c.Write(nil)
	return 0, nil
}

func (c *cFCr3rhcmDX) Write(b []byte) (int, error) {
	c.awCIY1.Lock()
	defer c.awCIY1.Unlock()
	length := len(b)
	if c.vpScqu == 8 {
		tbgQNW := syncpool.BytesBufferPoolAt5044220.Get().(*bytes.Buffer)
		defer syncpool.LjZLCz9(tbgQNW)
		for len(b) > 2048 {
			size := mathRand.Intn(4096) + 100
			if len(b) < size {
				size = len(b)
			}
			f75BSHe(tbgQNW, b[:size])
			b = b[size:]
		}
		if len(b) > 0 {
			f75BSHe(tbgQNW, b)
		}
		_, err := c.Conn.Write(tbgQNW.Bytes())
		if err != nil {
			return 0, err
		}
		return length, nil
	}

	if len(b) > 0 {
		f75BSHe(&c.dhQYYCsc, b)
	}

	if c.vpScqu == 0 {
		c.vpScqu = 1

		data := syncpool.BytesBufferPoolAt5044220.Get().(*bytes.Buffer)
		defer syncpool.LjZLCz9(data)

		data.Write([]byte{3, 3})
		c.fxMTvTqho7Lr(data)
		data.WriteByte(0x20)
		data.Write(c.l57vr8H[:])
		data.Write([]byte{0x00, 0x1c, 0xc0, 0x2b, 0xc0, 0x2f, 0xcc, 0xa9, 0xcc, 0xa8, 0xcc, 0x14, 0xcc, 0x13, 0xc0, 0x0a, 0xc0, 0x14, 0xc0, 0x09, 0xc0, 0x13, 0x00, 0x9c, 0x00, 0x35, 0x00, 0x2f, 0x00, 0x0a})
		data.Write([]byte{0x1, 0x0})

		ext := syncpool.BytesBufferPoolAt5044220.Get().(*bytes.Buffer)
		defer syncpool.LjZLCz9(ext)

		host := c.kGVO20Z()
		ext.Write([]byte{0xff, 0x01, 0x00, 0x01, 0x00})
		hU4pfL4l(ext, host)
		ext.Write([]byte{0, 0x17, 0, 0})
		c.h0EaidQ7(ext, host)
		ext.Write([]byte{0x00, 0x0d, 0x00, 0x16, 0x00, 0x14, 0x06, 0x01, 0x06, 0x03, 0x05, 0x01, 0x05, 0x03, 0x04, 0x01, 0x04, 0x03, 0x03, 0x01, 0x03, 0x03, 0x02, 0x01, 0x02, 0x03})
		ext.Write([]byte{0x00, 0x05, 0x00, 0x05, 0x01, 0x00, 0x00, 0x00, 0x00})
		ext.Write([]byte{0x00, 0x12, 0x00, 0x00})
		ext.Write([]byte{0x75, 0x50, 0x00, 0x00})
		ext.Write([]byte{0x00, 0x0b, 0x00, 0x02, 0x01, 0x00})
		ext.Write([]byte{0x00, 0x0a, 0x00, 0x06, 0x00, 0x04, 0x00, 0x17, 0x00, 0x18})

		binary.Write(data, binary.BigEndian, uint16(ext.Len()))
		data.ReadFrom(ext)

		ret := syncpool.BytesBufferPoolAt5044220.Get().(*bytes.Buffer)
		defer syncpool.LjZLCz9(ret)

		ret.Write([]byte{0x16, 3, 1})
		binary.Write(ret, binary.BigEndian, uint16(data.Len()+4))
		ret.Write([]byte{1, 0})
		binary.Write(ret, binary.BigEndian, uint16(data.Len()))
		ret.ReadFrom(data)

		_, err := c.Conn.Write(ret.Bytes())
		if err != nil {
			return 0, err
		}
		return length, nil
	} else if c.vpScqu == 1 && len(b) == 0 {
		tbgQNW := syncpool.BytesBufferPoolAt5044220.Get().(*bytes.Buffer)
		defer syncpool.LjZLCz9(tbgQNW)

		tbgQNW.Write([]byte{0x14, 3, 3, 0, 1, 1, 0x16, 3, 3, 0, 0x20})
		tbgQNW.ReadFrom(io.LimitReader(rand.Reader, int64(22)))
		tbgQNW.Write(c.cC9t2Xktic(tbgQNW.Bytes())[:10])
		tbgQNW.ReadFrom(&c.dhQYYCsc)

		c.vpScqu = 8

		_, err := c.Conn.Write(tbgQNW.Bytes())
		return 0, err
	}
	return length, nil
}

func f75BSHe(tbgQNW *bytes.Buffer, data []byte) {
	tbgQNW.Write([]byte{0x17, 3, 3})
	binary.Write(tbgQNW, binary.BigEndian, uint16(len(data)))
	tbgQNW.Write(data)
}

func (t *xIbCZlK) fxMTvTqho7Lr(tbgQNW *bytes.Buffer) {
	binary.Write(tbgQNW, binary.BigEndian, uint32(time.Now().Unix()))
	tbgQNW.ReadFrom(io.LimitReader(rand.Reader, int64(18)))
	tbgQNW.Write(t.cC9t2Xktic(tbgQNW.Bytes()[tbgQNW.Len()-22:])[:10])
}

func hU4pfL4l(tbgQNW *bytes.Buffer, u string) {
	len := uint16(len(u))
	tbgQNW.Write([]byte{0, 0})
	binary.Write(tbgQNW, binary.BigEndian, len+5)
	binary.Write(tbgQNW, binary.BigEndian, len+3)
	tbgQNW.WriteByte(0)
	binary.Write(tbgQNW, binary.BigEndian, len)
	tbgQNW.WriteString(u)
}

func (c *cFCr3rhcmDX) h0EaidQ7(tbgQNW *bytes.Buffer, u string) {
	length := 16 * (mathRand.Intn(17) + 8)
	tbgQNW.Write([]byte{0, 0x23})
	binary.Write(tbgQNW, binary.BigEndian, uint16(length))
	tbgQNW.ReadFrom(io.LimitReader(rand.Reader, int64(length)))
}

func (t *xIbCZlK) cC9t2Xktic(data []byte) []byte {
	key := syncpool.SharedPoolAt5096978.Get(len(t.JtkCoK71gtjm) + 32)
	defer syncpool.SharedPoolAt5096978.Put(key)
	copy(key, t.JtkCoK71gtjm)
	copy(key[len(t.JtkCoK71gtjm):], t.l57vr8H[:])

	sha1Data := tlsutil.B4_4_BnW(key, data)
	return sha1Data[:10]
}

func (t *xIbCZlK) kGVO20Z() string {
	host := t.NhWlHt412H8S
	if len(host) == 0 {
		host = t.O1k0EzteK
	}
	if len(host) > 0 && host[len(host)-1] >= '0' && host[len(host)-1] <= '9' {
		host = ""
	}
	hosts := strings.Split(host, ",")
	host = hosts[mathRand.Intn(len(hosts))]
	return host
}
