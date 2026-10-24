package sess

import (
	"net"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/user"
	"time"
)

// KsKg3yls 为原入站多用户 UDP 接口，RTTI 0x364d1a0 的完整8方法集。
type KsKg3yls interface {
	Close() error
	Free()
	LocalAddr() net.Addr
	ReadFromProxy([]byte) (int, user.PgaNd3014U, net.Addr, *relay.AHU2MRmjzQu, error)
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	WriteToProxy([]byte, user.PgaNd3014U, net.Addr, net.Addr) (int, error)
}
