package MFFa9yT

import (
	"encoding/binary"
	"hash/crc32"
	"net"
	C "recovered.local/soga/CAodbhZf"
	G "recovered.local/soga/G6bHD5JA"
	V "recovered.local/soga/gk1z9v"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	"time"
)

// y80zV6 @0x1da2ae0 及三个闭包0x1e74060/0x1e74160/0x1e74240。
// 先遍历 IP 缓存，再扫描未尝试的主用户表；原时间偏差按有符号64位计算。
func (v *K4avFiYOPi) y80zV6(ip string, authID [16]byte) *V.WR4lMNfab {
	if v.l_ZoZQDxG7 == nil {
		return nil
	}
	if v.qgz9PfTS.IsForbidden(ip) {
		Log.V8GN5Stk(-4, "forbidden ip %v access", ip)
		return nil
	}
	started := time.Now()
	now := int64(C.C5Jff5vs0())
	tried := 0
	var decrypted [16]byte
	var found *V.WR4lMNfab
	check := func(base U.PgaNd3014U) *V.WR4lMNfab {
		tried++
		user := base.(*U.QKaIWV2)
		block := user.GetBlock()
		if block == nil {
			return nil
		}
		block.Decrypt(decrypted[:], authID[:])
		stamp := int64(binary.BigEndian.Uint64(decrypted[:8]))
		if stamp-now > 120 || now-stamp > 120 {
			return nil
		}
		if crc32.ChecksumIEEE(decrypted[:12]) != binary.BigEndian.Uint32(decrypted[12:]) {
			return nil
		}
		// 原 AEAD 分支不赋值 B2pr6gMaPO，维持零值；T_ul7Q_0j 指向收到的认证块。
		return &V.WR4lMNfab{Xwwpafq: user, QAkNeaNePGH: uint32(stamp), Ih7r8HyMvOa: true, T_ul7Q_0j: authID[:]}
	}
	seen := make(map[int64]struct{})
	v.l_ZoZQDxG7.RangeUser(ip, func(user U.PgaNd3014U) bool {
		seen[user.GetId()] = struct{}{}
		found = check(user)
		return found == nil
	})
	if found == nil {
		v.jc_QPN6.Range(func(_ int64, user U.PgaNd3014U) bool {
			if _, ok := seen[user.GetId()]; ok {
				return true
			}
			found = check(user)
			return found == nil
		})
	}
	if tried > 0 {
		if found != nil {
			v.qgz9PfTS.Valid(ip)
			Log.V8GN5Stk(-4, "VMessAEAD auth success, tried: %v, time: %v", tried, time.Since(started))
			v.l_ZoZQDxG7.AddUser(ip, found.Xwwpafq, 0)
		} else {
			v.qgz9PfTS.Invalid(ip)
			Log.V8GN5Stk(-4, "VMessAEAD auth failed, tried: %v, time: %v", tried, time.Since(started))
		}
	}
	return found
}
func (v *K4avFiYOPi) iUB_i9oBP_68(authID [16]byte) *V.WR4lMNfab { // 0x1da3100
	if v.dj_cTnaUp6f == nil {
		return nil
	}
	id, stamp, ok := v.dj_cTnaUp6f.Get(authID)
	if !ok {
		Log.TRTbrquD(-4, "VMessMD5 auth failed:", authID[:])
		return nil
	}
	base, ok := v.jc_QPN6.Get(int64(id))
	if !ok {
		return nil
	}
	Log.TRTbrquD(-4, "VMessMD5 auth success")
	return &V.WR4lMNfab{Xwwpafq: base.(*U.QKaIWV2), QAkNeaNePGH: stamp, B2pr6gMaPO: authID}
}

// s3mqrqd_ @0x1da3360 使用单次非空 Read，而非 ReadFull；保留短读作为 fallback 原数据。
func (v *K4avFiYOPi) s3mqrqd_(conn net.Conn) (*V.WR4lMNfab, []byte) {
	var authID [16]byte
	conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	var n int
	var err error
	for {
		n, err = conn.Read(authID[:])
		if n > 0 || err != nil {
			break
		}
	}
	conn.SetReadDeadline(time.Time{})
	if err != nil || n != 16 {
		return nil, authID[:n]
	}
	addr, err := G.JY6gu_Rqfz(conn.RemoteAddr())
	if err != nil {
		Log.TRTbrquD(4, err)
		return nil, nil
	}
	if auth := v.iUB_i9oBP_68(authID); auth != nil {
		return auth, nil
	}
	if auth := v.y80zV6(addr.IP.String(), authID); auth != nil {
		return auth, nil
	}
	return nil, authID[:]
}
