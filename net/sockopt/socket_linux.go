package sockopt

import (
	"syscall"

	"golang.org/x/sys/unix"
	"recovered.local/soga/util/errutil"
)

// 原 tdNCv4fhOvkt 为命名函数类型；导出首字母以支持 Go 源码跨包引用。
type TdNCv4fhOvkt func(uintptr) error

// UA1QxZBYiPM @ 0x1585480。
func UA1QxZBYiPM(fd uintptr) error {
	if err := unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1); err != nil {
		return err
	}
	return unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
}

// ICQr8wkaYl @ 0x1585500。原函数忽略 IP_BIND_ADDRESS_NO_PORT 的设置错误。
func ICQr8wkaYl(fd uintptr) error {
	_ = unix.SetsockoptInt(int(fd), unix.IPPROTO_IP, unix.IP_BIND_ADDRESS_NO_PORT, 1)
	return nil
}

// CyQmxg 构造器在原二进制中内联；闭包 @ 0x1596020、0x1585560。
func CyQmxg(options ...TdNCv4fhOvkt) func(string, string, syscall.RawConn) error {
	return func(_, _ string, connection syscall.RawConn) error {
		var optionError error
		controlError := connection.Control(func(fd uintptr) {
			for _, option := range options {
				optionError = option(fd)
				if optionError != nil {
					return
				}
			}
		})
		return errutil.T2i1ItMaa(controlError, optionError)
	}
}
