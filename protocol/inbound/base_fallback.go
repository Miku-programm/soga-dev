package inbound

import (
	"net"
	D "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	L "recovered.local/soga/util/logx"
	"strconv"
)

// xx3H26VL 对应0x1d44b80。原回落在已有前缀非空时仍先读取一次补充数据，
// 不把此前缀立即单独发出；第一次读取/写入的任何错误均结束回落。
func (b *JWN154qJ4) xx3H26VL(conn net.Conn, prefix []byte) {
	if b.dZfoV0aXHW == "" || b.lAmttzTO8b <= 0 {
		return
	}
	address := net.JoinHostPort(b.dZfoV0aXHW, strconv.Itoa(b.lAmttzTO8b))
	L.TRTbrquD(-4, "handle fallback:", address)
	remote, err := G.DE9CHPNLbK(nil, address)
	if err != nil {
		return
	}
	defer remote.Close()
	if len(prefix) != 0 {
		buf := D.PoolAt50961D8.Get(1024)
		defer func() { D.PoolAt50961D8.Put(buf) }()
		copied := copy(buf, prefix)
		n, err := conn.Read(buf[copied:])
		if err != nil {
			return
		}
		if _, err = remote.Write(buf[:copied+n]); err != nil {
			return
		}
	}
	relay := G.AsVABr{C8FYQVoBTa: b.zb8f3cHWW8v, GdPSmtGs_1: conn, HEfD13ahR__b: remote, C93uQCY9TN: D.PoolAt50961D8.Get(1024), Mw4pNzmFlqA: D.PoolAt50961D8.Get(1024)}
	defer func() { relay.FreeBuf() }()
	G.FvtgLyQrWX(relay)
	L.TRTbrquD(-4, "handle fallback end")
}
