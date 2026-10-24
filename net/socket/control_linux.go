//go:build linux

package socket

import (
	"golang.org/x/sys/unix"
	"net"
	"syscall"
)

type socketControl = func(string, string, syscall.RawConn) error

// bindDeviceControl names the inlined closure construction. The closure body is
// c3ukcOJ40_.func1 at 0x1927140; deferral at 0x19270c0; syscall at 0x19273a0.
// Existing controls run last, only after this control succeeds.
func bindDeviceControl(previous socketControl, device string) socketControl {
	return func(network, address string, raw syscall.RawConn) (err error) {
		defer func() {
			if err == nil && previous != nil {
				err = previous(network, address, raw)
			}
		}()
		host, _, splitErr := net.SplitHostPort(address)
		if splitErr == nil {
			ip := net.ParseIP(host)
			if ip != nil && !ip.IsGlobalUnicast() {
				return nil
			}
		}
		var operationError error
		err = raw.Control(func(fd uintptr) { operationError = unix.BindToDevice(int(fd), device) })
		if operationError != nil {
			err = operationError
		}
		return err
	}
}

// markControl restores ewq0cKSDYv.func1 at 0x1927420, its deferred call at
// 0x19276e0 and syscall closure at 0x1927660 (SOL_SOCKET=1, SO_MARK=36).
func markControl(previous socketControl, mark int) socketControl {
	return func(network, address string, raw syscall.RawConn) (err error) {
		defer func() {
			if err == nil && previous != nil {
				err = previous(network, address, raw)
			}
		}()
		host, _, splitErr := net.SplitHostPort(address)
		if splitErr == nil {
			ip := net.ParseIP(host)
			if ip != nil && !ip.IsGlobalUnicast() {
				return nil
			}
		}
		var operationError error
		err = raw.Control(func(fd uintptr) { operationError = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, 36, mark) })
		if operationError != nil {
			err = operationError
		}
		return err
	}
}

// reuseControl restores jMWz9v.func1 at 0x1927760 and its closures
// 0x19278a0/0x1927d40. The two setsockopt errors are intentionally ignored.
func reuseControl(previous socketControl) socketControl {
	return func(network, address string, raw syscall.RawConn) (err error) {
		defer func() {
			if err == nil && previous != nil {
				err = previous(network, address, raw)
			}
		}()
		return raw.Control(func(fd uintptr) {
			_ = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
			_ = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		})
	}
}
