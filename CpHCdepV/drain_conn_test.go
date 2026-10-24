package CpHCdepV

import (
	"bytes"
	"io"
	"net"
	"testing"
)

type drainFixture struct {
	net.Conn
	reader *bytes.Reader
	err    error
	freed  bool
}

func (c *drainFixture) Read(p []byte) (int, error) {
	n, e := c.reader.Read(p)
	if c.err != nil {
		return n, c.err
	}
	return n, e
}
func (c *drainFixture) Free() { c.freed = true }
func TestOriginalRandomDrainWrapper(t *testing.T) {
	for i := 0; i < 100; i++ {
		c := MlZo77Me(nil).(*QBM9_LSBqsu)
		if c.ndn9Wlx < 54 || c.ndn9Wlx > 3382 {
			t.Fatal(c.ndn9Wlx)
		}
	}
	raw := &drainFixture{reader: bytes.NewReader([]byte("abcdefgh"))}
	c := &QBM9_LSBqsu{DrpXao: raw, ndn9Wlx: 5}
	b := make([]byte, 2)
	c.Read(b)
	if c.ndn9Wlx != 3 {
		t.Fatal(c.ndn9Wlx)
	}
	c.Free()
	if !raw.freed || raw.reader.Len() != 3 {
		t.Fatal("exact drain then Free")
	}
	raw = &drainFixture{reader: bytes.NewReader([]byte("abcdefgh")), err: io.EOF}
	c = &QBM9_LSBqsu{DrpXao: raw, ndn9Wlx: 1}
	n, e := c.Read(b)
	if n != 2 || e != io.EOF || c.ndn9Wlx != 1 {
		t.Fatal("n with error must not decrement")
	}
	raw.err = nil
	c.Read(b)
	if c.ndn9Wlx != -1 {
		t.Fatal("counter may cross zero")
	}
	c.Free()
	if raw.reader.Len() != 4 {
		t.Fatal("negative counter does not drain")
	}
}
