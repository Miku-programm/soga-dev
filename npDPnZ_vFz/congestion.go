package npDPnZ_vFz

import (
	quic "github.com/apernet/quic-go"
	B "recovered.local/soga/BOaPYrtdvWIg"
	R "recovered.local/soga/xIF4G3"
)

// XpQZo1W7 对应原入口 0x1c25440。
func XpQZo1W7(conn quic.Connection, bps uint64) { conn.SetCongestionControl(B.GPauZ708C6(bps)) }

// AW6N_Jc 对应原入口 0x1c253a0，初始窗口为32包，最大窗口为20000包。
func AW6N_Jc(conn quic.Connection) {
	conn.SetCongestionControl(R.NewSenderFromInline(R.MDeCv6qOB_J{}, R.CMryyxB(conn.RemoteAddr())))
}
