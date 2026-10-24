package CpHCdepV

import (
	"errors"
	"net"
	"testing"
	"time"
	"unsafe"
)

type tunnelPacketFixture struct {
	net.Conn
	source    net.Addr
	closed    int
	writes    int
	readError error
	deadline  time.Time
}

func (c *tunnelPacketFixture) RemoteAddr() net.Addr { return c.source }
func (c *tunnelPacketFixture) Read(p []byte) (int, error) {
	copy(p, "abc")
	c.source = &net.TCPAddr{Port: 99}
	return 3, c.readError
}
func (c *tunnelPacketFixture) Write(p []byte) (int, error)        { c.writes++; return 1, c.readError }
func (c *tunnelPacketFixture) Close() error                       { c.closed++; return c.readError }
func (c *tunnelPacketFixture) LocalAddr() net.Addr                { return c.source }
func (c *tunnelPacketFixture) SetDeadline(d time.Time) error      { c.deadline = d; return c.readError }
func (c *tunnelPacketFixture) SetReadDeadline(d time.Time) error  { c.deadline = d; return c.readError }
func (c *tunnelPacketFixture) SetWriteDeadline(d time.Time) error { c.deadline = d; return c.readError }
func TestOriginalTunnelPacketWrappers(t *testing.T) {
	var a Hzw0dY3ZZL
	var b NSuBse
	if unsafe.Sizeof(a) != 16 || unsafe.Sizeof(b) != 32 || unsafe.Offsetof(b.lxRU7JIrLmL) != 16 {
		t.Fatal("RTTI")
	}
	before := &net.TCPAddr{Port: 3}
	raw := &tunnelPacketFixture{source: before, readError: errors.New("partial")}
	pc := NewTunnelPacketFromStream(raw)
	buf := make([]byte, 8)
	n, addr, e := pc.ReadFrom(buf)
	if n != 3 || addr != before || e != raw.readError || string(buf[:n]) != "abc" {
		t.Fatal(n, addr, e)
	}
	wrapper := pc.GetRawPacketConn()
	n, addr, e = wrapper.ReadFrom(buf)
	if n != 3 || addr != raw.source || addr == before || e != raw.readError {
		t.Fatal("dynamic remote after read")
	}
	if n, e = pc.WriteTo(buf, before); n != 1 || e != raw.readError || raw.writes != 1 {
		t.Fatal("short/error write forwarded")
	}
	for _, set := range []func(time.Time) error{pc.SetDeadline, pc.SetReadDeadline, pc.SetWriteDeadline} {
		d := time.Unix(77, 88)
		if set(d) != raw.readError || raw.deadline != d {
			t.Fatal("deadline")
		}
	}
	if pc.LocalAddr() != raw.source || pc.Close() != raw.readError || raw.closed != 1 {
		t.Fatal("delegation")
	}
}
