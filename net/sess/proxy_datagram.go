package sess

import (
	"context"
	"net"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/util/errutil"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"time"
)

// NTe8CQ 为原接口 RTTI0x364d100，实际包路径J1P4KNvfI_J。
type NTe8CQ interface {
	Close() error
	GetUser() user.PgaNd3014U
	LocalAddr() net.Addr
	ReadFromProxy([]byte) (int, net.Addr, error)
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	WriteToProxy([]byte, net.Addr) (int, error)
}

// Nw8wZD 对应0x18871c0和归还缓冲的func1@18888c0。
// 读写截止时间的设置错误、成功短写均按原逻辑忽略。
func Nw8wZD(ctx context.Context, remote net.PacketConn, local NTe8CQ) error {
	buf := pool.PoolAt50961D8.Get(4096)
	defer pool.PoolAt50961D8.Put(buf)
	for {
		select {
		case <-ctx.Done():
			return errutil.ErrorAt5097800
		default:
		}
		_ = remote.SetReadDeadline(time.Now().Add(time.Duration(config.UDPTimeoutMinutes()) * time.Minute))
		n, address, err := remote.ReadFrom(buf)
		if err != nil {
			return err
		}
		_ = local.SetWriteDeadline(time.Now().Add(time.Minute))
		if _, err = local.WriteToProxy(buf[:n], address); err != nil {
			return err
		}
	}
}
