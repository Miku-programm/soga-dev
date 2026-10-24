package GTiWvwS

import (
	"encoding/binary"
	"net"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
)

// FcUA_X7yd 对应 RTTI 0x3659320，32 字节。
type FcUA_X7yd struct {
	kaT8OKQ    net.PacketConn
	gXGi5QPTn6 *N07NaA
	bDkiHh_    int
}

// ReadFromProxy 对应 0x1c64f20；坏包、未知用户及重放均继续读取下一报文。
func (r *FcUA_X7yd) ReadFromProxy(dst []byte) (int, nzATvajt.PgaNd3014U, net.Addr, *G6bHD5JA.AHU2MRmjzQu, error) {
	for {
		n, client, err := r.kaT8OKQ.ReadFrom(dst)
		if err != nil {
			return 0, nil, nil, nil, err
		}
		if n < r.bDkiHh_ {
			continue
		}
		packet := dst[:n]
		r.gXGi5QPTn6.su5Rzwte.Decrypt(dst[:16], dst[:16])
		var hash [16]byte
		r.gXGi5QPTn6.su5Rzwte.Decrypt(hash[:], dst[16:32])
		for i := range hash {
			hash[i] ^= dst[i]
		}
		user, ok := r.gXGi5QPTn6.GetUserByHash(hash)
		if !ok {
			continue
		}
		if !user.AllowAndAddUp(n) {
			continue
		}
		id := binary.BigEndian.Uint64(dst[:8])
		sequence := binary.BigEndian.Uint64(dst[8:16])
		session, err := r.gXGi5QPTn6.GetOrNewSession(id, user, client)
		if err != nil {
			continue
		}
		if !session.paeyZKdn.Check(sequence) {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 repeat packet id:", sequence)
			continue
		}
		if _, err = session.ul7vaxxk4q4.Open(dst[32:32], dst[4:16], packet[32:], nil); err != nil {
			continue
		}
		session.paeyZKdn.Add(sequence)
		p := CAodbhZf.NewDCeuEa2b(dst[32 : n-session.ul7vaxxk4q4.Overhead()])
		kind, err := p.ReadUInt8()
		if err != nil {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp read header type failed")
			continue
		}
		if kind != 0 {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp header type wrong:", kind)
			continue
		}
		stamp, err := p.ReadUInt64()
		if err != nil {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp read timestamp failed")
			continue
		}
		delta := int64(stamp) - CAodbhZf.C5Jff5vs0()
		if delta < 0 {
			delta = -delta
		}
		if delta > 30 {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp time error")
			continue
		}
		pad, err := p.ReadUInt16()
		if err != nil {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp read padding length failed")
			continue
		}
		if err = p.Skip(int(pad)); err != nil {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp padding length not enough:", pad)
			continue
		}
		addr, err := G6bHD5JA.SzhW4tOzUZg(p)
		if err != nil {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp read addr failed")
			continue
		}
		n, err = p.Read(dst)
		if err != nil {
			muc279Mpa6S4.TRTbrquD(-4, "ss2022 udp read payload failed")
			continue
		}
		return n, user, client, addr, nil
	}
}
func (r *FcUA_X7yd) Free() {} // 0x1c65be0 原机器码仅 RET，没有资源释放。
