package J1P4KNvfI_J

import (
	"context"
	"errors"
	"io"
	"net"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/nzATvajt"
	"reflect"
	"testing"
)

type measuredUser struct {
	nzATvajt.PgaNd3014U
	ctx      context.Context
	up, down int
	out      bool
	calls    []string
}

func (u *measuredUser) GetCtx() context.Context { u.calls = append(u.calls, "ctx"); return u.ctx }
func (u *measuredUser) AddUp(n int)             { u.calls = append(u.calls, "up"); u.up += n }
func (u *measuredUser) AddDown(n int)           { u.calls = append(u.calls, "down"); u.down += n }
func (u *measuredUser) IsTrafficOut() bool      { return u.out }
func (u *measuredUser) GetId() int64            { return 7 }

type measuredConn struct {
	net.Conn
	readData              []byte
	readErr, errorWrite   error
	writeN, reads, writes int
}

func (c *measuredConn) Read(p []byte) (int, error) {
	c.reads++
	n := copy(p, c.readData)
	return n, c.readErr
}
func (c *measuredConn) Write(p []byte) (int, error) { c.writes++; return c.writeN, c.errorWrite }
func TestUserConnectionLayoutAndLateBinding(t *testing.T) {
	typ := reflect.TypeFor[Mdn09F6aMrpa]()
	if typ.Size() != 64 {
		t.Fatal(typ.Size())
	}
	for i, offset := range []uintptr{0, 16, 24, 32, 48} {
		if typ.Field(i).Offset != offset {
			t.Fatal(typ.Field(i))
		}
	}
	raw := &measuredConn{readData: []byte("abc"), readErr: io.EOF, writeN: 1, errorWrite: io.ErrShortWrite}
	c := &Mdn09F6aMrpa{DrpXao: raw}
	if n, err := c.Read(make([]byte, 10)); n != 3 || err != io.EOF {
		t.Fatalf("%d %v", n, err)
	}
	if n, err := c.Write([]byte("hello")); n != 1 || err != io.ErrShortWrite {
		t.Fatalf("%d %v", n, err)
	}
	user := &measuredUser{ctx: context.Background()}
	c.SetUser(user)
	if user.up != 3 || user.down != 5 || c.dqmafe != 0 || c.e1uNaRZ != 0 || c.GetUser() != user || c.GetRawConn() != raw {
		t.Fatalf("%+v %+v", user, c)
	}
	if !reflect.DeepEqual(user.calls, []string{"ctx", "up", "down"}) {
		t.Fatal(user.calls)
	}
	c.AddUp(-2)
	c.AddDown(-3)
	if user.up != 1 || user.down != 2 {
		t.Fatal("signed counts changed")
	}
}
func TestUserConnectionOriginalCheckOrdering(t *testing.T) {
	raw := &measuredConn{readData: []byte("abc"), readErr: io.EOF, writeN: 1}
	user := &measuredUser{ctx: context.Background(), out: true}
	c := &Mdn09F6aMrpa{DrpXao: raw}
	c.SetUser(user)
	if n, err := c.Read(make([]byte, 3)); n != 0 || err == nil || err.Error() != "user traffic out" || raw.reads != 1 || user.up != 3 {
		t.Fatalf("%d %v %+v", n, err, user)
	}
	if n, err := c.Write([]byte("abcdef")); n != 0 || err == nil || raw.writes != 0 || user.down != 6 {
		t.Fatalf("%d %v %+v", n, err, user)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	user.ctx = ctx
	user.out = false
	c.SetUser(user)
	if n, err := c.Write([]byte("xx")); n != 0 || err != HxKViEp2a.ErrorAt5097800 || errors.Is(err, context.Canceled) || raw.writes != 0 {
		t.Fatalf("%d %v", n, err)
	}
}
