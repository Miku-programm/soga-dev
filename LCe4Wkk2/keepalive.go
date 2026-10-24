package LCe4Wkk2

import (
	"net"
	"time"
)

// mS0F0Ti restores 0x1b6c0c0. It affects only a concrete *net.TCPConn,
// enables keepalive and sets its period to 30 seconds; both errors are ignored.
func mS0F0Ti(c net.Conn) {
	if tcp, ok := c.(*net.TCPConn); ok {
		_ = tcp.SetKeepAlive(true)
		_ = tcp.SetKeepAlivePeriod(30 * time.Second)
	}
}

// fS848iPPRMrU is the original 16-byte embedded PacketConn wrapper.
type fS848iPPRMrU struct{ net.PacketConn }
