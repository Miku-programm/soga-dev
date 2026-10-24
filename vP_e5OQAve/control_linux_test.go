//go:build linux

package vP_e5OQAve

import (
	"errors"
	"syscall"
	"testing"
)

type testRawConn struct {
	called int
	run    bool
	err    error
}

func (r *testRawConn) Control(f func(uintptr)) error {
	r.called++
	if r.run {
		f(^uintptr(0))
	}
	return r.err
}
func (*testRawConn) Read(func(uintptr) bool) error  { panic("not used") }
func (*testRawConn) Write(func(uintptr) bool) error { panic("not used") }
func TestControlsLoopbackAndPreviousOrder(t *testing.T) {
	for _, makeControl := range []func(socketControl) socketControl{
		func(p socketControl) socketControl { return bindDeviceControl(p, "no-device") },
		func(p socketControl) socketControl { return markControl(p, 7) },
	} {
		raw := new(testRawConn)
		previousCalls := 0
		sentinel := errors.New("previous")
		control := makeControl(func(network, address string, r syscall.RawConn) error {
			previousCalls++
			if raw.called != 0 || network != "tcp4" || address != "127.0.0.1:7" || r != raw {
				t.Fatal("unexpected args/order")
			}
			return sentinel
		})
		if err := control("tcp4", "127.0.0.1:7", raw); err != sentinel || previousCalls != 1 {
			t.Fatal(err, previousCalls)
		}
	}
}
func TestControlErrorPrecedence(t *testing.T) {
	rawError := errors.New("raw")
	for _, factory := range []func(socketControl) socketControl{
		func(p socketControl) socketControl { return bindDeviceControl(p, "invalid-device") },
		func(p socketControl) socketControl { return markControl(p, 7) },
	} {
		called := 0
		previous := func(string, string, syscall.RawConn) error { called++; return nil }
		raw := &testRawConn{err: rawError}
		if err := factory(previous)("tcp4", "invalid-address", raw); err != rawError || called != 0 {
			t.Fatal(err, called)
		}
		raw = &testRawConn{run: true, err: rawError}
		err := factory(previous)("tcp4", "8.8.8.8:53", raw)
		if err == nil || err == rawError || called != 0 {
			t.Fatal("setsockopt error must override Control error", err, called)
		}
	}
}
func TestReuseIgnoresSetsockoptFailure(t *testing.T) {
	called := 0
	raw := &testRawConn{run: true}
	fn := reuseControl(func(string, string, syscall.RawConn) error {
		called++
		if raw.called != 1 {
			t.Fatal("previous ran first")
		}
		return nil
	})
	if err := fn("udp", "", raw); err != nil || called != 1 {
		t.Fatal(err, called)
	}
	sentinel := errors.New("raw")
	raw.err = sentinel
	if err := fn("udp", "", raw); err != sentinel || called != 1 {
		t.Fatal(err, called)
	}
}
