package vlessvision

import (
	"crypto/tls"
	"net"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/transport/proxyproto"
	"recovered.local/soga/protocol/vless"
	"recovered.local/soga/reality"
)

// CocBnud6 @0x1d0a9e0。保留确切具体类型、分支顺序及只解一层的行为。
// 调用真实恢复的 Reality.NetConn 和 PROXY.GetRawConn getter。
func CocBnud6(conn *sess.Mdn09F6aMrpa) net.Conn {
	raw := conn.DrpXao
	wrapped, special := raw.(*vless.RfBfjVN4XE)
	if special {
		raw = wrapped.DrpXao
	}
	if value, ok := raw.(*vless.LUKUT5W0R6); ok {
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
	if value, ok := raw.(*proxyproto.MOTygWJHAwt5); ok {
		raw = value.GetRawConn()
	}
	return raw
}
