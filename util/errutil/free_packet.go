package errutil

import "net"

// W81xCX4t 对应原 RTTI 0x35d7028，方法记录 0x35d7088。
type W81xCX4t interface{ GetRawPacketConn() net.PacketConn }

// PpnNLr_pf7He 对应 0xb26a40。原 switch 表 0x5040e80 依次为
// Free、net.Conn、GetRawPacketConn；遇到 stream 类型委托原 stream 释放链。
func PpnNLr_pf7He(conn net.PacketConn) {
	for conn != nil {
		switch v := conn.(type) {
		case JZgM1sMF:
			U9tk4Xosh9(v)
			return
		case net.Conn:
			Li3c4CJ1_D(v)
			return
		case W81xCX4t:
			conn = v.GetRawPacketConn()
		default:
			return
		}
	}
}
