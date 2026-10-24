package jIIdlAO

import (
	"errors"
	"testing"
	"time"
	"unsafe"
)

func TestCacheSuccessErrorAndExpiration(t *testing.T) {
	g := NewCache(time.Hour)
	sentinel := errors.New("cached")
	v, err, shared := g.Do(func() (any, error) { return 17, sentinel })
	if v != 17 || err != sentinel || shared {
		t.Fatal(v, err, shared)
	}
	v, err, shared = g.Do(func() (any, error) { t.Fatal("cache miss"); return nil, nil })
	if v != 17 || err != sentinel || !shared {
		t.Fatal(v, err, shared)
	}
	g.mu.Lock()
	g.started = time.Now().Add(-2 * time.Hour)
	g.mu.Unlock()
	v, err, shared = g.Do(func() (any, error) { return 18, nil })
	if v != 18 || err != nil || shared {
		t.Fatal(v, err, shared)
	}
	if unsafe.Sizeof(*g) != 56 || unsafe.Sizeof(SRClJmN_{}) != 32 || unsafe.Sizeof(tdfGa3B_{}) != 48 {
		t.Fatal("layout")
	}
}
func TestConcurrentUsesSameResult(t *testing.T) {
	g := NewCache(0)
	started := make(chan struct{})
	release := make(chan struct{})
	done := make(chan struct{})
	go func() {
		defer close(done)
		g.Do(func() (any, error) { close(started); <-release; return "ready", nil })
	}()
	<-started
	// Keep a callback visibly in flight while starting the second caller.
	second := make(chan struct {
		v      any
		err    error
		shared bool
	}, 1)
	go func() {
		v, e, s := g.Do(func() (any, error) { return "unexpected", nil })
		second <- struct {
			v      any
			err    error
			shared bool
		}{v, e, s}
	}()
	// A positive TTL keeps both scheduling possibilities equivalent: the second
	// caller either waits for the in-flight result or reads the same cached result.
	g.mu.Lock()
	g.duration = time.Hour
	g.mu.Unlock()
	close(release)
	r := <-second
	<-done
	if r.v != "ready" || r.err != nil || !r.shared {
		t.Fatal(r)
	}
}
