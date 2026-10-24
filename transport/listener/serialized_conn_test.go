package listener

import (
	"errors"
	"net"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type serializedFixture struct {
	net.Conn
	inside, max atomic.Int32
	panicWrite  bool
	err         error
}

func (c *serializedFixture) Write(p []byte) (int, error) {
	if c.panicWrite {
		panic("write panic")
	}
	n := c.inside.Add(1)
	for {
		old := c.max.Load()
		if old >= n || c.max.CompareAndSwap(old, n) {
			break
		}
	}
	time.Sleep(time.Millisecond)
	c.inside.Add(-1)
	return len(p), c.err
}
func TestOriginalSerializedConnection(t *testing.T) {
	typ := reflect.TypeFor[MhXtuhyPkR5]()
	if typ.Size() != 24 || typ.Field(1).Offset != 16 || !typ.Field(0).Anonymous || !typ.Field(1).Anonymous {
		t.Fatal(typ)
	}
	marker := errors.New("write failure")
	raw := &serializedFixture{err: marker}
	c := &MhXtuhyPkR5{DrpXao: raw}
	var wg sync.WaitGroup
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if n, err := c.Write([]byte{1, 2}); n != 2 || err != marker {
				t.Errorf("%v %v", n, err)
			}
		}()
	}
	wg.Wait()
	if raw.max.Load() != 1 || c.GetRawConn() != raw {
		t.Fatal(raw.max.Load())
	}
	raw.panicWrite = true
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected underlying panic")
			}
		}()
		c.Write(nil)
	}()
	if c.TryLock() {
		c.Unlock()
		t.Fatal("original panic path must leave mutex locked")
	}
	c.Unlock()
}
