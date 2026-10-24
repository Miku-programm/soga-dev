package J1P4KNvfI_J

import (
	"net"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
)

type trackedPacketConn struct {
	net.PacketConn
	closes atomic.Int32
}

func (c *trackedPacketConn) Close() error { c.closes.Add(1); return nil }
func TestPacketCacheOriginalLifecycleAndConcurrentCreation(t *testing.T) {
	if reflect.TypeFor[TaQ7Prswa]().Size() != 48 || reflect.TypeFor[HEzs3hFmtN]().Size() != 24 || reflect.TypeFor[HEzs3hFmtN]().Field(1).Offset != 8 {
		t.Fatal("RTTI layout")
	}
	var m TaQ7Prswa
	var wg sync.WaitGroup
	var fresh atomic.Int32
	got := make(chan *HEzs3hFmtN, 32)
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, loaded := m.GetOrCreate("same")
			if !loaded {
				fresh.Add(1)
			}
			got <- v
		}()
	}
	wg.Wait()
	close(got)
	r, _ := m.Get("same")
	for v := range got {
		if v != r {
			t.Fatal("different pending result")
		}
	}
	if fresh.Load() != 1 {
		t.Fatal(fresh.Load())
	}
	select {
	case <-r.Ready():
		t.Fatal("ready too soon")
	default:
	}
	conn := new(trackedPacketConn)
	r.Complete(conn)
	<-r.Ready()
	if r.Conn() != conn {
		t.Fatal("wrong connection")
	}
	m.CloseAll()
	m.CloseAll()
	if _, ok := m.Get("same"); !ok || conn.closes.Load() != 2 {
		t.Fatal("CloseAll unexpectedly deleted or skipped")
	}
	m.Delete("same")
	if _, ok := m.Get("same"); ok || conn.closes.Load() != 3 {
		t.Fatal("Delete")
	}
	pending, _ := m.GetOrCreate("pending")
	m.Delete("pending")
	late := new(trackedPacketConn)
	pending.Complete(late)
	m.CloseAll()
	if late.closes.Load() != 0 {
		t.Fatal("removed pending entry was closed")
	}
	empty, _ := m.GetOrCreate("nil")
	empty.Complete(nil)
	m.Delete("nil")
}
