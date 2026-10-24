package UQ7OErS0s

import (
	"bytes"
	"log/slog"
	"net"
	B "recovered.local/soga/D_zCag1qTV"
	G "recovered.local/soga/G6bHD5JA"
	H "recovered.local/soga/HxKViEp2a"
	R "recovered.local/soga/RoaZP1oJy"
	T "recovered.local/soga/TaFdeUcnzQ"
	S "recovered.local/soga/ad5KEbw_EzWk"
	L "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	"sync"
)

type WCApOca6IQW7 = net.PacketConn

// Original 48-byte RTTI c4opfNpyW.AXWGpkYq.
type AXWGpkYq struct {
	WCApOca6IQW7
	sync.Mutex
	j6haDfFi R.Vzuv_s9ZS5
	adiN53   *S.L7E7H2aW8KVs
}

func INOdx8(raw net.PacketConn, info *S.L7E7H2aW8KVs) (*AXWGpkYq, error) { //0x1ce5da0
	encryption, err := T.DP5vq3(info, raw)
	if err != nil {
		return nil, err
	}
	protocol, err := R.HrQWlmA(info.C8Qsa8G8XMw0, encryption)
	if err != nil {
		return nil, err
	}
	protocol.SetServerInfo(info)
	return &AXWGpkYq{WCApOca6IQW7: raw, j6haDfFi: protocol, adiN53: info}, nil
}

// ReadFromProxy restores 0x1ce5ee0 and its original deferred closure. The
// recovery helper remains nested, matching the original call depth. An address
// parse error is dropped in favor of the underlying packet error, as shown by
// 0x1ce641d leaving the named error slot unchanged.
func (c *AXWGpkYq) ReadFromProxy(p []byte) (n int, user U.PgaNd3014U, src net.Addr, dest *G.AHU2MRmjzQu, err error) {
	defer func() {
		if H.SlG_Y_mg("read from proxy") != nil {
			n = 0
			err = nil
		}
	}()
	var uid int64
	n, uid, src, err = c.j6haDfFi.ReadFromUID(p)
	if n <= 0 {
		return 0, nil, src, nil, err
	}
	reader := bytes.NewReader(p[:n])
	target, parseErr := G.SzhW4tOzUZg(reader)
	if parseErr != nil {
		return 0, nil, src, nil, err
	}
	payloadLen := copy(p, p[n-reader.Len():n])
	if c.adiN53.NxYA16xMfF != 2 {
		return payloadLen, nil, src, target, nil
	}
	u, ok := c.adiN53.N3pMXuwX85c.Get(uid)
	if !ok {
		L.V8GN5Stk(slog.LevelDebug, "uid %v not exist", uid)
		return 0, nil, src, nil, err
	}
	if !u.AhSFeSua.AllowAndAddUp(n) {
		L.V8GN5Stk(slog.LevelDebug, "user %v speed limited", u.GetId())
		return 0, u, src, target, nil
	}
	return payloadLen, u, src, target, nil
}

// WriteToProxy restores 0x1ce65a0. It accounts the address header, asserts the
// original concrete SSR user, and returns the full protocol count. The lock
// uses explicit Unlock; a protocol panic retains the original locked state.
func (c *AXWGpkYq) WriteToProxy(p []byte, user U.PgaNd3014U, client, target net.Addr) (int, error) {
	address, err := G.RNFJeaRY(target)
	if err != nil {
		return 0, err
	}
	u := user.(*U.LJJkaD)
	buffer := B.LdHEmSs9(8192)
	defer buffer.Free()
	if _, err = G.DdYroT1n(&buffer, address); err != nil {
		return 0, err
	}
	buffer.Write(p)
	if !user.AllowAndAddDown(buffer.Len()) {
		L.V8GN5Stk(slog.LevelDebug, "user %v speed limited", user.GetId())
		return 0, nil
	}
	c.Lock()
	n, err := c.j6haDfFi.WriteToUID(buffer.Bytes(), client, u.GetUID())
	c.Unlock()
	return n, err
}
func (c *AXWGpkYq) WriteTo(p []byte, addr net.Addr) (int, error) { //0x1ce6aa0
	c.Lock()
	n, err := c.j6haDfFi.WriteTo(p, addr)
	c.Unlock()
	return n, err
}
func (c *AXWGpkYq) ReadFrom(p []byte) (int, net.Addr, error) { return c.j6haDfFi.ReadFrom(p) } //0x1ce6be0
func (c *AXWGpkYq) GetServerInfo() *S.L7E7H2aW8KVs           { return c.adiN53 }               //0x1ce6c40
func (c *AXWGpkYq) Free() { //0x1ce6c60
	if v, ok := c.j6haDfFi.(H.JZgM1sMF); ok {
		H.U9tk4Xosh9(v)
	}
	H.PpnNLr_pf7He(c.WCApOca6IQW7)
}
