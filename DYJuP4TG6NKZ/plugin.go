package DYJuP4TG6NKZ

import (
	"crypto/tls"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"net/http"
	"net/netip"
	ws "recovered.local/soga/FLhkKwJ8M"
	buffer "recovered.local/soga/Hx04IbKk"
)

type YgrMuLrWXz struct {
	PRkTz1l     string
	D6eC4LPWJm0 string
	G3GjlvRQW   string
	BwX3DgZaf   map[string]string
	PwG_Wda     bool
	HjmCr8      bool
	AUavUP      bool
}
type ClSQMRRJ struct {
	net.Conn
	kmya1ow_2Zw  buffer.BaDekw7qe
	heaj4Yoh     [2]byte
	kYIntLPfH    [2]byte
	rrOTvcj      [2]byte
	w9i2ksWrHUTk []byte
	baWziQysg3CB int
}

// DO4qmnxxv4_: 0x19c3aa0. TLS ServerName follows the canonical Host header;
// no explicit MinVersion is set here. The MUX session targets 127.0.0.1:0.
func DO4qmnxxv4_(c net.Conn, options *YgrMuLrWXz) (net.Conn, error) {
	headers := make(http.Header)
	for k, v := range options.BwX3DgZaf {
		headers.Add(k, v)
	}
	config := &ws.ELIQWabl31{OhyeDUGTCGH: options.PRkTz1l, RTIB8Nkai: options.D6eC4LPWJm0, WOT5UbYYto: options.G3GjlvRQW, RlZC_oDxj: headers}
	if options.PwG_Wda {
		config.Xtp2zWRK = true
		config.VbB82ra = &tls.Config{ServerName: options.PRkTz1l, InsecureSkipVerify: options.HjmCr8, NextProtos: []string{"http/1.1"}}
		if host := headers.Get("Host"); host != "" {
			config.VbB82ra.ServerName = host
		}
	}
	conn, e := ws.E6YB0lYJp(c, config)
	if e != nil {
		return nil, e
	}
	if options.AUavUP {
		return VvqqszrNmTt(conn, [2]byte{}, 0, "127.0.0.1", ""), nil
	}
	return conn, nil
}

// VvqqszrNmTt: 0x19c2b00. Stack fields at +0/+2/+8/+24 map to the
// session bytes, destination port, host, and network. Parameter grouping and
// source spelling cannot be recovered; this signature exposes their exact values.
func VvqqszrNmTt(c net.Conn, id [2]byte, port uint16, host, network string) *ClSQMRRJ {
	b := buffer.BaDekw7qe(make([]byte, 0))
	b.PutSlice([]byte{0, 0})
	b.PutSlice(id[:])
	b.PutSlice([]byte{1, 0})
	kind := byte(1)
	if network == "udp" {
		kind = 2
	}
	b.PutSlice([]byte{kind, byte(port >> 8), byte(port)})
	ip, e := netip.ParseAddr(host)
	if e != nil {
		b.PutSlice([]byte{2})
		b.PutSlice([]byte(host))
	} else {
		if ip.Is4() {
			b.PutSlice([]byte{1})
		} else {
			b.PutSlice([]byte{3})
		}
		b.PutSlice(ip.AsSlice())
	}
	binary.BigEndian.PutUint16(b[:2], uint16(len(b)-2))
	return &ClSQMRRJ{Conn: c, heaj4Yoh: id, w9i2ksWrHUTk: b.Bytes()}
}

// Read: 0x19c2420. Metadata length is only bounded, not used to skip extra
// fields. Status uses a single Read; partial successful reads retain old bytes.
func (c *ClSQMRRJ) Read(p []byte) (int, error) {
	if c.baWziQysg3CB != 0 {
		n, e := c.Conn.Read(p[:min(len(p), c.baWziQysg3CB)])
		if e != nil {
			return 0, e
		}
		c.baWziQysg3CB -= n
		return n, nil
	}
	for {
		if _, e := io.ReadFull(c.Conn, c.kYIntLPfH[:]); e != nil {
			return 0, e
		}
		if binary.BigEndian.Uint16(c.kYIntLPfH[:]) > 512 {
			return 0, errors.New("invalid metalen")
		}
		if _, e := io.ReadFull(c.Conn, c.heaj4Yoh[:]); e != nil {
			return 0, e
		}
		if _, e := c.Conn.Read(c.rrOTvcj[:]); e != nil {
			return 0, e
		}
		if c.rrOTvcj[0] != 4 && c.rrOTvcj[1] == 1 {
			break
		}
	}
	if _, e := io.ReadFull(c.Conn, c.kYIntLPfH[:]); e != nil {
		return 0, e
	}
	c.baWziQysg3CB = int(binary.BigEndian.Uint16(c.kYIntLPfH[:]))
	n, e := c.Conn.Read(p[:min(len(p), c.baWziQysg3CB)])
	c.baWziQysg3CB -= n
	return n, e
}

// Write: 0x19c2800. Returns the complete framed write count (payload + 8),
// with no size splitting and uint16 payload-length truncation.
func (c *ClSQMRRJ) Write(p []byte) (int, error) {
	if c.w9i2ksWrHUTk != nil {
		if _, e := c.Conn.Write(c.w9i2ksWrHUTk); e != nil {
			return 0, e
		}
		c.w9i2ksWrHUTk = nil
	}
	c.kmya1ow_2Zw.Reset()
	c.kmya1ow_2Zw.PutSlice([]byte{0, 4})
	c.kmya1ow_2Zw.PutSlice(c.heaj4Yoh[:])
	c.kmya1ow_2Zw.PutSlice([]byte{2, 1, byte(len(p) >> 8), byte(len(p))})
	c.kmya1ow_2Zw.PutSlice(p)
	return c.Conn.Write(c.kmya1ow_2Zw.Bytes())
}

// Close: 0x19c2a40. Failure to write the close frame leaves the raw conn open.
func (c *ClSQMRRJ) Close() error {
	if _, e := c.Conn.Write([]byte{0, 4, c.heaj4Yoh[0], c.heaj4Yoh[1], 3, 0}); e != nil {
		return e
	}
	return c.Conn.Close()
}
