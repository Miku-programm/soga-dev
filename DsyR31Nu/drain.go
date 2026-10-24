package DsyR31Nu

import (
	"net"
	D "recovered.local/soga/D_zCag1qTV"
	"time"
)

// Q6wDan 对应0x1171760；size<=0是持续丢弃直到出错，而不是不读。
func Q6wDan(conn net.Conn, size int, timeout time.Duration) error {
	buf := D.PoolAt50961D8.Get(4096)
	defer D.PoolAt50961D8.Put(buf)
	return TFaQ3h6zUaKf(conn, buf, size, timeout)
}

// TFaQ3h6zUaKf 对应0x1171880。有限长度模式只设置一次读期限，
// 无限模式每次读前更新；读取正好足够字节同时返回错误时仍传播错误。
func TFaQ3h6zUaKf(conn net.Conn, buf []byte, size int, timeout time.Duration) error {
	if size <= 0 {
		for {
			if timeout > 0 {
				_ = conn.SetReadDeadline(time.Now().Add(timeout))
			}
			if _, err := conn.Read(buf); err != nil {
				return err
			}
		}
	}
	if timeout > 0 {
		_ = conn.SetReadDeadline(time.Now().Add(timeout))
	}
	for size > 0 {
		n, err := conn.Read(buf[:min(len(buf), size)])
		size -= n
		if err != nil {
			return err
		}
	}
	return nil
}
