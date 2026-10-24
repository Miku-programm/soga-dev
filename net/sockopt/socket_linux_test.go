package sockopt

import (
	"errors"
	"reflect"
	"testing"

	"golang.org/x/sys/unix"
)

type rawConnection struct {
	fd           uintptr
	controlError error
	call         bool
}

func (r rawConnection) Control(f func(uintptr)) error {
	if r.call {
		f(r.fd)
	}
	return r.controlError
}
func (r rawConnection) Read(func(uintptr) bool) error  { panic("unexpected Read") }
func (r rawConnection) Write(func(uintptr) bool) error { panic("unexpected Write") }

func TestSocketOptionsOnLocalDescriptor(t *testing.T) {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM|unix.SOCK_CLOEXEC, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer unix.Close(fd)
	if err := UA1QxZBYiPM(uintptr(fd)); err != nil {
		t.Fatal(err)
	}
	for _, option := range []int{unix.SO_REUSEADDR, unix.SO_REUSEPORT} {
		got, err := unix.GetsockoptInt(fd, unix.SOL_SOCKET, option)
		if err != nil || got != 1 {
			t.Fatalf("option %d: %d %v", option, got, err)
		}
	}
	if err := ICQr8wkaYl(uintptr(fd)); err != nil {
		t.Fatal(err)
	}
	got, err := unix.GetsockoptInt(fd, unix.IPPROTO_IP, unix.IP_BIND_ADDRESS_NO_PORT)
	if err != nil || got != 1 {
		t.Fatalf("bind option %d %v", got, err)
	}
	invalid := ^uintptr(0)
	if err := UA1QxZBYiPM(invalid); !errors.Is(err, unix.EBADF) {
		t.Fatalf("invalid fd: %v", err)
	}
	if err := ICQr8wkaYl(invalid); err != nil {
		t.Fatal("original bind helper ignores failure")
	}
}

func TestControlOptionsStopAndCombineErrors(t *testing.T) {
	var seen []uintptr
	opErr, ctlErr := errors.New("option"), errors.New("control")
	callback := CyQmxg(func(fd uintptr) error { seen = append(seen, fd); return nil }, func(fd uintptr) error { seen = append(seen, fd+1); return opErr }, func(uintptr) error { t.Fatal("ran after failed option"); return nil })
	err := callback("tcp", "ignored", rawConnection{fd: 17, controlError: ctlErr, call: true})
	if !reflect.DeepEqual(seen, []uintptr{17, 18}) || err.Error() != "control\noption" {
		t.Fatalf("%v %v", seen, err)
	}
	seen = nil
	err = callback("", "", rawConnection{controlError: ctlErr})
	if len(seen) != 0 || err.Error() != "control" || err == ctlErr {
		t.Fatal("control-only failure must be wrapped")
	}
	if err := CyQmxg()("", "", rawConnection{call: true}); err != nil {
		t.Fatal(err)
	}
	options := []TdNCv4fhOvkt{func(uintptr) error { return opErr }}
	captured := CyQmxg(options...)
	options[0] = func(uintptr) error { return nil }
	if captured("", "", rawConnection{call: true}) != nil {
		t.Fatal("original captures slice backing array")
	}
}
