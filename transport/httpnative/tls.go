package httpnative

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/util/errutil"
	"recovered.local/soga/crypto/ssmethod"
)

type DrpXao = net.Conn

// IkPl_ual retains the original 80-byte native TLS-obfuscation state.
type IkPl_ual struct {
	DrpXao
	w9I7sl       string
	naM5tdG      ssmethod.VE7a8rvh
	ubA_OeWCX    bool
	pX8an7       bool
	rfszTODt8c_O *pool.X_aXX5
	btPvDAOYtSO  int
	vGM3bgb8VDgL []byte
}

// NewTLSForFactory bridges the original cross-package struct literal in
// ssudp.FciVp4yFHL@0x1c83e20; this helper name is newly introduced.
func NewTLSForFactory(conn net.Conn, host string, mode ssmethod.VE7a8rvh) *IkPl_ual {
	return &IkPl_ual{DrpXao: conn, w9I7sl: host, naM5tdG: mode}
}

// Original static slices at 0x503b890/0x503b8b0 and init@0x1c79c40.
var versionAt503B890 = []byte{3, 3}
var changeAt503B8B0 = []byte{20, 3, 3, 0, 1, 1}
var helloTailAt509A750 = []byte{0xcc, 0xa8, 0, 0, 15, 0xff, 1, 0, 1, 0, 0, 23, 0, 0, 0, 11, 0, 2, 1, 0}

func (c *IkPl_ual) wwoi7Usrwb() error { //0x1c7dc40
	var header [5]byte
	if _, err := io.ReadFull(c.DrpXao, header[:]); err != nil {
		return err
	}
	if header[0] != 22 || header[1] != 3 {
		return fmt.Errorf("not a tls handshake record: %x", header[:3])
	}
	size := int(binary.BigEndian.Uint16(header[3:]))
	if size == 0 {
		return fmt.Errorf("empty tls handshake record")
	}
	b := pool.S4OQXWIZ(size)
	defer b.Free()
	if _, err := b.ReadFullFrom(c.DrpXao, size); err != nil {
		return err
	}
	session, host, payload, err := vYzlYln7glEA(b.Bytes())
	if err != nil {
		return err
	}
	if (c.naM5tdG == 1 || c.naM5tdG == 3) && len(c.w9I7sl) != 0 && host != c.w9I7sl {
		return fmt.Errorf("client host (%v) not match server host (%v)", host, c.w9I7sl)
	}
	if len(session) == 32 {
		c.vGM3bgb8VDgL = append([]byte(nil), session...)
	}
	if len(payload) != 0 {
		c.rfszTODt8c_O = pool.S4OQXWIZ(len(payload))
		c.rfszTODt8c_O.Write(payload)
	}
	return nil
}
func (c *IkPl_ual) oK9c_TP8a8Aq() (int, error) { //0x1c7e480
	var h [5]byte
	if _, err := io.ReadFull(c.DrpXao, h[:]); err != nil {
		return 0, err
	}
	if h[0] != 23 || h[1] != 3 {
		return 0, fmt.Errorf("not a tls application data record: %x", h[:3])
	}
	return int(binary.BigEndian.Uint16(h[3:])), nil
}
func (c *IkPl_ual) Read(p []byte) (int, error) { //0x1c7e680
	if !c.ubA_OeWCX {
		if err := c.wwoi7Usrwb(); err != nil {
			return 0, err
		}
		c.ubA_OeWCX = true
	}
	if c.rfszTODt8c_O != nil {
		if c.rfszTODt8c_O.Len() > 0 {
			return c.rfszTODt8c_O.Read(p)
		}
		c.rfszTODt8c_O.Free()
		c.rfszTODt8c_O = nil
	}
	if len(p) == 0 {
		return 0, nil
	}
	for c.btPvDAOYtSO == 0 {
		n, err := c.oK9c_TP8a8Aq()
		if err != nil {
			return 0, err
		}
		c.btPvDAOYtSO = n
	}
	if len(p) > c.btPvDAOYtSO {
		p = p[:c.btPvDAOYtSO]
	}
	n, err := c.DrpXao.Read(p)
	c.btPvDAOYtSO -= n
	return n, err
}
func (c *IkPl_ual) Write(p []byte) (int, error) { //0x1c7e8a0
	if !c.pX8an7 {
		c.pX8an7 = true
		return c.yG3B8yJAs(p)
	}
	if len(p) == 0 {
		return 0, nil
	}
	b := pool.S4OQXWIZ(5*(len(p)>>14) + len(p) + 5)
	defer b.Free()
	jwLEcxzeXJH(b, p, 23)
	if _, err := c.DrpXao.Write(b.Bytes()); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (c *IkPl_ual) yG3B8yJAs(p []byte) (int, error) { //0x1c7eae0
	b := pool.S4OQXWIZ(len(changeAt503B8B0) + 5*(len(p)>>14) + len(p) + 101)
	defer b.Free()
	b.WriteByte(22)
	b.Write(versionAt503B890)
	b.WriteUint16(91)
	b.WriteByte(2)
	b.WriteByte(0)
	b.WriteUint16(87)
	b.Write(versionAt503B890)
	b.WriteUint32(uint32(rtx.C5Jff5vs0()))
	b.WriteRandom(28)
	b.WriteByte(32)
	if len(c.vGM3bgb8VDgL) == 32 {
		b.Write(c.vGM3bgb8VDgL)
	} else {
		b.WriteRandom(32)
	}
	b.Write(helloTailAt509A750)
	b.Write(changeAt503B8B0)
	jwLEcxzeXJH(b, p, 22)
	if _, err := c.DrpXao.Write(b.Bytes()); err != nil {
		return 0, err
	}
	return len(p), nil
}

// jwLEcxzeXJH@0x1c7efe0 writes one record even for an empty payload. Subsequent
// chunks always use application-data type 23, including when the first is 22.
func jwLEcxzeXJH(b *pool.X_aXX5, p []byte, kind byte) int {
	for {
		n := min(len(p), 16384)
		b.WriteByte(kind)
		b.Write(versionAt503B890)
		b.WriteUint16(uint16(n))
		written, _ := b.Write(p[:n])
		p = p[n:]
		if len(p) == 0 {
			return written
		}
		kind = 23
	}
}
func (c *IkPl_ual) Free() {
	if c.rfszTODt8c_O != nil {
		c.rfszTODt8c_O.Free()
		c.rfszTODt8c_O = nil
	}
	errutil.Li3c4CJ1_D(c.DrpXao)
} //0x1c7ef60

// s_lsSqY46j is the original stack-only slice cursor. Field labels below are
// analysis names; no descriptor for this stack-only type survived in RTTI.
type s_lsSqY46j struct {
	data   []byte
	offset int
}

func (r *s_lsSqY46j) take(n int) ([]byte, error) {
	if len(r.data)-r.offset < n {
		return nil, io.ErrUnexpectedEOF
	}
	p := r.data[r.offset : r.offset+n]
	r.offset += n
	return p, nil
}
func (r *s_lsSqY46j) oU_t8xNC1A() ([]byte, error) {
	p, e := r.take(1)
	if e != nil {
		return nil, e
	}
	return r.take(int(p[0]))
} //0x1c7fa60
func (r *s_lsSqY46j) yscpqMaLQMs() ([]byte, error) {
	p, e := r.take(2)
	if e != nil {
		return nil, e
	}
	return r.take(int(binary.BigEndian.Uint16(p)))
} //0x1c7fb60
// jHOgG9GS@0x1c7f940 ignores malformed SNI, returning an empty host.
func jHOgG9GS(p []byte) string {
	r := s_lsSqY46j{data: p}
	list, e := r.yscpqMaLQMs()
	if e != nil {
		return ""
	}
	r = s_lsSqY46j{data: list}
	for r.offset < len(r.data) {
		kind, e := r.take(1)
		if e != nil {
			return ""
		}
		name, e := r.yscpqMaLQMs()
		if e != nil {
			return ""
		}
		if kind[0] == 0 {
			return string(name)
		}
	}
	return ""
}

// vYzlYln7glEA@0x1c7f1a0 reads the original ClientHello framing. It deliberately
// ignores handshake length/version, and keeps the last SNI and ticket extension.
func vYzlYln7glEA(p []byte) (session []byte, host string, ticket []byte, err error) {
	r := s_lsSqY46j{data: p}
	kind, e := r.take(1)
	if e != nil {
		return nil, "", nil, e
	}
	if kind[0] != 1 {
		return nil, "", nil, fmt.Errorf("not a tls client hello: %x", kind[0])
	}
	for _, n := range []int{3, 2, 32} {
		if _, e = r.take(n); e != nil {
			return nil, "", nil, e
		}
	}
	session, e = r.oU_t8xNC1A()
	if e != nil {
		return session, "", nil, e
	}
	if _, e = r.yscpqMaLQMs(); e != nil {
		return session, "", nil, e
	}
	if _, e = r.oU_t8xNC1A(); e != nil {
		return session, "", nil, e
	}
	extensions, e := r.yscpqMaLQMs()
	if e != nil {
		return session, "", nil, e
	}
	r = s_lsSqY46j{data: extensions}
	for r.offset < len(r.data) {
		kind, e := r.take(2)
		if e != nil {
			return session, host, ticket, e
		}
		data, e := r.yscpqMaLQMs()
		if e != nil {
			return session, host, ticket, e
		}
		switch binary.BigEndian.Uint16(kind) {
		case 0:
			host = jHOgG9GS(data)
		case 35:
			ticket = data
		}
	}
	return session, host, ticket, nil
}
