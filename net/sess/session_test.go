package sess

import (
	"errors"
	"net"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/user"
	"reflect"
	"testing"
)

type sessionFixtureConn struct {
	net.Conn
	local, remote net.Addr
	user          user.PgaNd3014U
	dest          *relay.AHU2MRmjzQu
	kind          ZES6F5
	free, closed  int
	err           error
}

func (c *sessionFixtureConn) LocalAddr() net.Addr                { return c.local }
func (c *sessionFixtureConn) RemoteAddr() net.Addr               { return c.remote }
func (c *sessionFixtureConn) GetUser() user.PgaNd3014U       { return c.user }
func (c *sessionFixtureConn) GetDestAddr() *relay.AHU2MRmjzQu { return c.dest }
func (c *sessionFixtureConn) NetworkType() ZES6F5                { return c.kind }
func (c *sessionFixtureConn) Free()                              { c.free++ }
func (c *sessionFixtureConn) Close() error                       { c.closed++; return c.err }

type packetUnionFixture struct {
	NlU9CLU8MJPy
	free, closed int
	err          error
}

func (c *packetUnionFixture) Free()        { c.free++ }
func (c *packetUnionFixture) Close() error { c.closed++; return c.err }

func TestOriginalConnectionUnionDispatchAndLayout(t *testing.T) {
	typ := reflect.TypeFor[DqwVny0NcXb]()
	if typ.Size() != 40 || typ.Field(0).Offset != 0 || typ.Field(1).Offset != 8 || typ.Field(2).Offset != 24 {
		t.Fatal(typ)
	}
	var zero DqwVny0NcXb
	if c, ok := zero.AsStream(); !ok || c != nil {
		t.Fatal(c, ok)
	}
	if c, ok := zero.AsDatagram(); ok || c != nil {
		t.Fatal(c, ok)
	}
	err := errors.New("close sentinel")
	s := &sessionFixtureConn{err: err}
	d := &packetUnionFixture{err: err}
	for _, tag := range []HpOTlmdRdFv{0, 1, 2, 255} {
		c := DqwVny0NcXb{tag, s, d}
		c.Free()
		got := c.Close()
		if tag <= 1 && got != err || tag > 1 && got != nil {
			t.Fatal(tag, got)
		}
	}
	if s.free != 1 || s.closed != 1 || d.free != 1 || d.closed != 1 {
		t.Fatal(s, d)
	}
}
func TestOriginalSessionIdentityAddressAndShallowClone(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	ip := net.IPv4(198, 51, 100, 47)
	c := &sessionFixtureConn{Conn: a, local: &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}, remote: &net.UDPAddr{IP: ip, Port: 5678}, dest: new(relay.AHU2MRmjzQu), kind: 1}
	s := MHLL7QOie(c)
	if s.X1gBTWHQO99 != 1234 || s.Y72kDV35 != 5678 || s.Iayb4V != c.dest || s.YlxwftjCvz_ != 1 || s.C66t8X != s.Uqzvq94e {
		t.Fatal(s)
	}
	if &s.KeHzUwuXWh0[0] != &ip[0] {
		t.Fatal("remote IP copied")
	}
	clone := s.Clone()
	if clone.E3L0l8hZ <= s.E3L0l8hZ {
		t.Fatal("sequence did not advance")
	}
	clone.E3L0l8hZ = s.E3L0l8hZ
	if !reflect.DeepEqual(clone, s) || &clone.KeHzUwuXWh0[0] != &s.KeHzUwuXWh0[0] {
		t.Fatal("clone not shallow")
	}
	// net.Pipe 的 pipeAddr 不属于 TCPAddr/UDPAddr，原方法保留零值，不能解析其 String。
	c.local = a.LocalAddr()
	c.remote = a.RemoteAddr()
	u := MHLL7QOie(c)
	if u.L4SIW3qbjS != nil || u.KeHzUwuXWh0 != nil || u.C66t8X != "" || u.Uqzvq94e != "" {
		t.Fatal(u)
	}
}
func TestOriginalSessionLayout(t *testing.T) {
	typ := reflect.TypeFor[ZzkV7_5l]()
	offsets := []uintptr{0, 8, 16, 24, 48, 56, 80, 88, 104, 128, 136, 160, 168, 184, 200, 208, 216}
	if typ.Size() != 232 || typ.NumField() != len(offsets) {
		t.Fatal(typ)
	}
	for i, want := range offsets {
		if typ.Field(i).Offset != want {
			t.Fatal(typ.Field(i))
		}
	}
}
