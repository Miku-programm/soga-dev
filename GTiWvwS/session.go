package GTiWvwS

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"net"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/w2YaTd9_"
	"sync/atomic"
)

// RGGcUu1R 保留原 RTTI 0x36eca48 的 1104 字节会话状态。
type RGGcUu1R struct {
	fl_h6QOtKast uint64
	e5Lk2aYUf    uint64
	c9dMiAM0t89w atomic.Uint64
	ob06bMoWqrSi net.Addr
	ul7vaxxk4q4  cipher.AEAD
	qXprYj6U9AK  cipher.AEAD
	paeyZKdn     KHKdG8c
}

// jr2_LZfPY0H 对应 0x1c68860。两个方向分别由各自的 session ID 派生密钥。
func jr2_LZfPY0H(id uint64, method w2YaTd9_.C6q3Upu, user *nzATvajt.WtK9HQSy66, addr net.Addr) (*RGGcUu1R, error) {
	var salt [8]byte
	binary.BigEndian.PutUint64(salt[:], id)
	incoming, err := method.NewAEAD(method.SubKey(user.PSK(), salt[:]))
	if err != nil {
		return nil, err
	}
	rand.Read(salt[:])
	server := binary.BigEndian.Uint64(salt[:])
	for server == id {
		rand.Read(salt[:])
		server = binary.BigEndian.Uint64(salt[:])
	}
	outgoing, err := method.NewAEAD(method.SubKey(user.PSK(), salt[:]))
	if err != nil {
		return nil, err
	}
	session := &RGGcUu1R{fl_h6QOtKast: server, e5Lk2aYUf: id, ob06bMoWqrSi: addr, ul7vaxxk4q4: incoming, qXprYj6U9AK: outgoing}
	session.c9dMiAM0t89w.Add(^uint64(0))
	return session, nil
}

// KHKdG8c 对应 0x3631cd0，滑动窗口是 128 个 64 位块。
type KHKdG8c struct {
	rsPnryPvPJ   uint64
	nypzxyc8qDxT [128]uint64
}

// Reset 对应 0x1c6f7a0；原函数只清最高序号及第一个块。
func (w *KHKdG8c) Reset() { w.rsPnryPvPJ = 0; w.nypzxyc8qDxT[0] = 0 }

// Check 对应 0x1c6f7c0。检查不消费序号；有效回溯距离包含 8128。
func (w *KHKdG8c) Check(id uint64) bool {
	if id > w.rsPnryPvPJ {
		return true
	}
	if w.rsPnryPvPJ-id > 8128 {
		return false
	}
	return w.nypzxyc8qDxT[(id>>6)&127]&(uint64(1)<<(id&63)) == 0
}

// Add 对应 0x1c6f800。不隐含调用 Check。
func (w *KHKdG8c) Add(id uint64) {
	if id > w.rsPnryPvPJ {
		block := w.rsPnryPvPJ >> 6
		count := (id >> 6) - block
		if count > 128 {
			count = 128
		}
		for count > 0 {
			block = (block + 1) & 127
			w.nypzxyc8qDxT[block] = 0
			count--
		}
		w.rsPnryPvPJ = id
	}
	w.nypzxyc8qDxT[(id>>6)&127] |= uint64(1) << (id & 63)
}
