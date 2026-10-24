package GTiWvwS

import (
	"fmt"
	"net"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/uLofn6j"
)

// JxxVAb 对应 RTTI 0x3631d60，24 字节。
type JxxVAb struct {
	i5SFpe     net.PacketConn
	aoeqs2GBaN *N07NaA
}

// WriteToProxy 对应 0x1c68ae0；传入客户端地址用于查会话，实际目的地取会话最新地址。
func (w *JxxVAb) WriteToProxy(payload []byte, user nzATvajt.PgaNd3014U, client, remote net.Addr) (int, error) {
	addr, err := G6bHD5JA.RNFJeaRY(remote)
	if err != nil {
		return 0, err
	}
	session, ok := w.aoeqs2GBaN.GetSessionByClientAddr(client)
	if !ok {
		return 0, fmt.Errorf("ss2022 udp session not exist")
	}
	padding := 0
	if len(payload) < 900 {
		padding = uLofn6j.Co6D_zfls(900-len(payload)) + 1
	}
	b := D_zCag1qTV.LdHEmSs9(0x2000)
	defer b.Free()
	b.WriteUint64(session.fl_h6QOtKast)
	b.WriteUint64(session.c9dMiAM0t89w.Add(1))
	b.WriteByte(1)
	b.WriteUint64(uint64(uint32(CAodbhZf.C5Jff5vs0())))
	b.WriteUint64(session.e5Lk2aYUf)
	b.WriteUint16(uint16(padding))
	if padding > 0 {
		b.WriteRandom(padding)
	}
	if _, err = G6bHD5JA.DdYroT1n(&b, addr); err != nil {
		return 0, err
	}
	header := b.Bytes()[:16]
	b.SealWriteWithStart(session.qXprYj6U9AK, 16, header[4:16], payload, nil)
	concrete := user.(*nzATvajt.WtK9HQSy66)
	concrete.GetBlock().Encrypt(header, header)
	if !concrete.AllowAndAddDown(b.Len()) {
		muc279Mpa6S4.V8GN5Stk(-4, "user %v udp speed limited", concrete.GetId())
		return 0, nil
	}
	if _, err = w.i5SFpe.WriteTo(b.Bytes(), session.ob06bMoWqrSi); err != nil {
		return 0, err
	}
	return len(payload), nil
}
func (w *JxxVAb) Free() {} // 0x1c694c0 原函数仅 RET。
