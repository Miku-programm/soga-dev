package udpnat

import (
	"testing"
	"time"
	"unsafe"
)

func TestOriginalObservableBlockingAndClose(t *testing.T) {
	if unsafe.Sizeof(EkVnc6a{}) != 32 || unsafe.Sizeof(SaTXuENDW{}) != 24 {
		t.Fatal("RTTI layout")
	}
	input := make(chan any)
	o := AkUZ2QiKCu3n(input)
	output := make(chan any)
	s := &SaTXuENDW{szJyRBo2e9: output}
	o.nd3U58QVfEal.Lock()
	o.jHrXT_T2TvUz[ZutmWj26aQLd(output)] = s
	o.nd3U58QVfEal.Unlock()
	firstAccepted := make(chan struct{})
	bothAccepted := make(chan struct{})
	go func() { input <- 1; close(firstAccepted); input <- 2; close(bothAccepted); close(input) }()
	select {
	case <-firstAccepted:
	case <-time.After(time.Second):
		t.Fatal("receiver missing")
	}
	select {
	case <-bothAccepted:
		t.Fatal("must block while first subscriber send waits")
	case <-time.After(10 * time.Millisecond):
	}
	for _, want := range []int{1, 2} {
		select {
		case got := <-output:
			if got != want {
				t.Fatal(got)
			}
		case <-time.After(time.Second):
			t.Fatal("lost event")
		}
	}
	select {
	case _, open := <-output:
		if open {
			t.Fatal("subscriber not closed")
		}
	case <-time.After(time.Second):
		t.Fatal("close timeout")
	}
	s.Close()
	s.Close()
	o.nd3U58QVfEal.Lock()
	closed, count := o.zGiaD7v, len(o.jHrXT_T2TvUz)
	o.nd3U58QVfEal.Unlock()
	if !closed || count != 1 {
		t.Fatal("closed flag/map retention")
	}
}
func TestObservableNoSubscribersStillDrains(t *testing.T) {
	input := make(chan any)
	o := AkUZ2QiKCu3n(input)
	done := make(chan struct{})
	go func() {
		for i := 0; i < 100; i++ {
			input <- i
		}
		close(input)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("undrained channel")
	}
	deadline := time.Now().Add(time.Second)
	for {
		o.nd3U58QVfEal.Lock()
		closed := o.zGiaD7v
		o.nd3U58QVfEal.Unlock()
		if closed {
			return
		}
		if time.Now().After(deadline) {
			t.Fatal("never closed")
		}
		time.Sleep(time.Millisecond)
	}
}
