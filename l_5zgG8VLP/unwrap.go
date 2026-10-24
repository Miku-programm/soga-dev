package l_5zgG8VLP

import (
	"crypto/tls"
	"net"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/JcmlVIK0T0i"
	"recovered.local/soga/bQxeORT"
	"recovered.local/soga/reality"
)

// CocBnud6 @0x1d0a9e0。保留确切具体类型、分支顺序及只解一层的行为。
// 调用真实恢复的 Reality.NetConn 和 PROXY.GetRawConn getter。
func CocBnud6(conn *J1P4KNvfI_J.Mdn09F6aMrpa) net.Conn {
	raw := conn.DrpXao
	wrapped, special := raw.(*bQxeORT.RfBfjVN4XE)
	if special {
		raw = wrapped.DrpXao
	}
	if value, ok := raw.(*bQxeORT.LUKUT5W0R6); ok {
		return value
	}
	if !special {
		switch value := raw.(type) {
		case *tls.Conn:
			raw = value.NetConn()
		case *reality.Conn:
			raw = value.NetConn()
		}
	}
	if value, ok := raw.(*JcmlVIK0T0i.MOTygWJHAwt5); ok {
		raw = value.GetRawConn()
	}
	return raw
}
