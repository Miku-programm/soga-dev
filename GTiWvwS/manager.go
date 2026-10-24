package GTiWvwS

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"github.com/jellydator/ttlcache/v3"
	"lukechampine.com/blake3"
	"net"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/w2YaTd9_"
	"time"
	"unsafe"
)

// N07NaA 对应原 RTTI 0x36f4ca0，96 字节。C6q3Upu 是命名 string，非 interface。
type N07NaA struct {
	kUidCbZ_CtGU w2YaTd9_.C6q3Upu
	gIdWgkT8p3   []byte
	sK7pRj       *G7A7csgrs62p.GTFAWF[[16]byte, *nzATvajt.WtK9HQSy66]
	naPQvgu      *G7A7csgrs62p.GTFAWF[int64, [16]byte]
	dnTM_jovq    *ttlcache.Cache[string, struct{}]
	su5Rzwte     cipher.Block
	cKL3iGa54    *ttlcache.Cache[uint64, *RGGcUu1R]
	p_kZ2IwdCBhC *ttlcache.Cache[string, *RGGcUu1R]
}

// DkAuKfGU 对应 0x1c6ebe0，以及两个清理闭包 0x1c702e0 / 0x1c70340。
func DkAuKfGU(ctx context.Context, conf *LFWTTXF7JR.AAMPrdcYTL6) (*N07NaA, error) {
	key, err := conf.HMkc4BgR.PSK(conf.VuhWqou)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	salts := ttlcache.New[string, struct{}](ttlcache.WithTTL[string, struct{}](time.Minute), ttlcache.WithDisableTouchOnHit[string, struct{}]())
	duration := time.Duration(cNT24rG4F.UDPTimeoutMinutes()) * time.Minute
	sessions := ttlcache.New[uint64, *RGGcUu1R](ttlcache.WithTTL[uint64, *RGGcUu1R](duration))
	addresses := ttlcache.New[string, *RGGcUu1R](ttlcache.WithTTL[string, *RGGcUu1R](duration))
	CAodbhZf.SwUrlM3M(ctx, func() { salts.DeleteExpired() }, time.Minute)
	CAodbhZf.SwUrlM3M(ctx, func() { sessions.DeleteExpired(); addresses.DeleteExpired() }, duration)
	return &N07NaA{kUidCbZ_CtGU: conf.HMkc4BgR, gIdWgkT8p3: key, sK7pRj: new(G7A7csgrs62p.GTFAWF[[16]byte, *nzATvajt.WtK9HQSy66]), naPQvgu: new(G7A7csgrs62p.GTFAWF[int64, [16]byte]), dnTM_jovq: salts, su5Rzwte: block, cKL3iGa54: sessions, p_kZ2IwdCBhC: addresses}, nil
}
func (m *N07NaA) GetCipher() w2YaTd9_.C6q3Upu { return m.kUidCbZ_CtGU } // 0x1c6f0a0
func (m *N07NaA) GetPSK() []byte              { return m.gIdWgkT8p3 }   // 0x1c6f0c0
// UpdateUser 对应 0x1c6f0e0，原类型断言及旧 hash 删除发生于新 hash 写入之前。
func (m *N07NaA) UpdateUser(value nzATvajt.PgaNd3014U) {
	user := value.(*nzATvajt.WtK9HQSy66)
	id := user.GetId()
	if hash, ok := m.naPQvgu.Get(id); ok {
		m.sK7pRj.Delete(hash)
	}
	full := blake3.Sum512(user.PSK())
	var hash [16]byte
	copy(hash[:], full[:16])
	m.sK7pRj.Set(hash, user)
	m.naPQvgu.Set(id, hash)
}
func (m *N07NaA) RemoveUser(id int64) {
	if hash, ok := m.naPQvgu.GetAndDelete(id); ok {
		m.sK7pRj.Delete(hash)
	}
}                                                                          // 0x1c6f2c0
func (m *N07NaA) GetUserByHash(hash [16]byte) (*nzATvajt.WtK9HQSy66, bool) { return m.sK7pRj.Get(hash) } // 0x1c6f340
func (m *N07NaA) GetHashById(id int64) ([16]byte, bool)                    { return m.naPQvgu.Get(id) }  // 0x1c6f380
// CheckSalt 对应 0x1c6f3e0；空 slice 在原 &salt[0] 表达式处 panic。
func (m *N07NaA) CheckSalt(salt []byte) bool {
	return m.dnTM_jovq.Get(unsafe.String(&salt[0], len(salt))) != nil
}
func (m *N07NaA) AddSalt(salt []byte) { m.dnTM_jovq.Set(string(salt), struct{}{}, ttlcache.DefaultTTL) } // 0x1c6f480
// GetOrNewSession 对应 0x1c6f500。已有 ID 可更换源地址，不重新派生密钥。
func (m *N07NaA) GetOrNewSession(id uint64, user *nzATvajt.WtK9HQSy66, addr net.Addr) (*RGGcUu1R, error) {
	var session *RGGcUu1R
	if item := m.cKL3iGa54.Get(id); item != nil {
		session = item.Value()
		session.ob06bMoWqrSi = addr
	} else {
		var err error
		session, err = jr2_LZfPY0H(id, m.kUidCbZ_CtGU, user, addr)
		if err != nil {
			return nil, err
		}
	}
	m.cKL3iGa54.Set(id, session, ttlcache.DefaultTTL)
	m.p_kZ2IwdCBhC.Set(addr.String(), session, ttlcache.DefaultTTL)
	return session, nil
}

// GetSessionByClientAddr 对应 0x1c6f680，同时刷新两种索引的有效期。
func (m *N07NaA) GetSessionByClientAddr(addr net.Addr) (*RGGcUu1R, bool) {
	key := addr.String()
	item := m.p_kZ2IwdCBhC.Get(key)
	if item == nil {
		return nil, false
	}
	session := item.Value()
	m.cKL3iGa54.Set(session.e5Lk2aYUf, session, ttlcache.DefaultTTL)
	m.p_kZ2IwdCBhC.Set(key, session, ttlcache.DefaultTTL)
	return session, true
}
