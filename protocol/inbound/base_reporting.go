package inbound

import (
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"recovered.local/soga/panel/registry"
)

// wKvd7jacLJ 对应0x1d440a0/0x1e6c240；计数先被清空，仅任一正数时上报。
func (b *JWN154qJ4) wKvd7jacLJ() types.CWfcUv {
	out := make(types.CWfcUv, 0, 16)
	b.jc_QPN6.Range(func(_ int64, u user.PgaNd3014U) bool {
		up, down := u.GetAndClearTraffic()
		if up > 0 || down > 0 {
			out = append(out, &types.DLxslMwYJXQw{Hy7Q97RROK: u.GetId(), XJkGTQf_qkAt: up, I6LXObODav: down})
		}
		return true
	})
	return out
}

// iyyXcKcQH7k7 对应0x1d44180与两个闭包0x1e6c460/0x1e6c500。
func (b *JWN154qJ4) iyyXcKcQH7k7() types.QmiwDVjV00lZ {
	out := make(types.QmiwDVjV00lZ, 0)
	b.qE9aU9.Range(func(detector int, users *genmap.Fxs7SY[int64]) bool {
		users.Range(func(user int64) bool {
			out = append(out, &types.B9FS114R4F9h{GbzQUT_4f: user, CAgUUdrv9: detector})
			return true
		})
		b.qE9aU9.Delete(detector)
		return true
	})
	return out
}

// jdbxrQIN2 对应0x1d44240/0x1e6c600；保留每个用户返回的原指针。
func (b *JWN154qJ4) jdbxrQIN2() types.CQhZrJ4dK {
	out := make(types.CQhZrJ4dK, 0, 16)
	b.jc_QPN6.Range(func(_ int64, u user.PgaNd3014U) bool {
		if x := u.GetAliveIpData(); x != nil {
			out = append(out, x)
		}
		return true
	})
	return out
}
func (b *JWN154qJ4) UserExist(id int64) bool { _, ok := b.jc_QPN6.Get(id); return ok } //0x1d44320

// GetSubmitData 对应0x1d43be0。两种阈值均先乘1024并转换int64，
// 在线IP按完整流量快照筛选，再独立筛选上报流量；已清空的计数不会补回。
func (b *JWN154qJ4) GetSubmitData() *registry.FGmNba4MhioY {
	traffic := b.wKvd7jacLJ()
	alive := b.jdbxrQIN2()
	aliveKB, trafficKB := config.SubmitTrafficThresholds()
	if threshold := int64(aliveKB * 1024); threshold > 0 {
		sums := make(map[int64]int64, len(traffic))
		for _, x := range traffic {
			sums[x.Hy7Q97RROK] = x.XJkGTQf_qkAt + x.I6LXObODav
		}
		filtered := make(types.CQhZrJ4dK, 0, len(alive))
		for _, x := range alive {
			if sums[x.Vy9ybu] >= threshold {
				filtered = append(filtered, x)
			}
		}
		alive = filtered
	}
	if threshold := int64(trafficKB * 1024); threshold > 0 {
		filtered := make(types.CWfcUv, 0, len(traffic))
		for _, x := range traffic {
			if x.XJkGTQf_qkAt+x.I6LXObODav >= threshold {
				filtered = append(filtered, x)
			}
		}
		traffic = filtered
	}
	return &registry.FGmNba4MhioY{HTgPyE3VytZd: traffic, JNcHHwTZ7: alive, Y0ys4ma4Tno4: b.iyyXcKcQH7k7()}
}
