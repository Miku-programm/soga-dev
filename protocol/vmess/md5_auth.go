package vmess

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/binary"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	CA "recovered.local/soga/util/rtx"
	U "recovered.local/soga/user"
)

// 原 RTTI：认证时间桶 16 字节，管理器 88 字节，认证结果 56 字节。
type CWAyuvwV9 struct {
	iDoMs2     map[[16]byte][2]uint32
	sKKwaeHJ7J uint32
}
type MuywMbTwQ []*CWAyuvwV9

func (s MuywMbTwQ) Len() int           { return len(s) }                            // 0x1d13e60
func (s MuywMbTwQ) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }                  // 0x1d13d80
func (s MuywMbTwQ) Less(i, j int) bool { return s[i].sKKwaeHJ7J < s[j].sKKwaeHJ7J } // 0x1d13e20
type E3b13Z3FO struct {
	iukZFcM_4kl context.Context
	uuPtg2OHz   *U.ZpNFqUVZp
	gHO0Rwg     uint32
	dOa0Z1B     sync.RWMutex
	hz3qQl7H5   MuywMbTwQ
	nlX_TBEa    atomic.Pointer[[]*U.QKaIWV2]
}
type WR4lMNfab struct {
	Xwwpafq     *U.QKaIWV2
	QAkNeaNePGH uint32
	Ih7r8HyMvOa bool
	B2pr6gMaPO  [16]byte
	T_ul7Q_0j   []byte
}

func (a *WR4lMNfab) GetCmdKey() []byte { return a.Xwwpafq.GetCmdKey() } // 0x1d13c40
func (a *WR4lMNfab) GetIV() []byte { // 0x1d13c80
	var data [32]byte
	for i := 0; i < 4; i++ {
		binary.BigEndian.PutUint64(data[i*8:], uint64(a.QAkNeaNePGH))
	}
	sum := md5.Sum(data[:])
	return sum[:]
}

// NewMD5AuthManager 是分析者命名的跨包桥接；构造逻辑来自内联于
// inbound.ag90n_rZjJO @0x1da0d40 的原分支，不是原程序新增的工厂符号。
func NewMD5AuthManager(ctx context.Context, users *U.ZpNFqUVZp, shortWindow bool) *E3b13Z3FO {
	window := uint32(70)
	if shortWindow {
		window = 25
	}
	if window%10 != 0 {
		window = window/10*10 + 10
	}
	return &E3b13Z3FO{iukZFcM_4kl: ctx, uuPtg2OHz: users, gHO0Rwg: window, hz3qQl7H5: MuywMbTwQ{}}
}
func (m *E3b13Z3FO) StartTask() { go m.ofF3NP4P() } // 0x1d13e80
func (m *E3b13Z3FO) ofF3NP4P() { // 0x1d13f00
	defer func() { m.hzdGla95(); runtime.GC() }()
	start := uint32(CA.C5Jff5vs0()) - m.gHO0Rwg
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		now := uint32(CA.C5Jff5vs0())
		end := now + m.gHO0Rwg
		m.ajZZhr(now - m.gHO0Rwg)
		m.nSbchfUCNxuE(start, end)
		if users := m.nlX_TBEa.Swap(nil); users != nil {
			m.wP_5CanVePo(*users, now-m.gHO0Rwg, start)
		}
		start = end
		select {
		case <-ticker.C:
		case <-m.iukZFcM_4kl.Done():
			return
		}
	}
}
func (m *E3b13Z3FO) nSbchfUCNxuE(start, end uint32) { // 0x1d14100
	users := make([]*U.QKaIWV2, 0, m.uuPtg2OHz.Len()+100)
	m.uuPtg2OHz.Range(func(_ int64, u U.PgaNd3014U) bool { users = append(users, u.(*U.QKaIWV2)); return true })
	m.wP_5CanVePo(users, start, end)
}
func (m *E3b13Z3FO) wP_5CanVePo(users []*U.QKaIWV2, start, end uint32) { // 0x1d14220
	if len(users) == 0 {
		return
	}
	for t := start; t < end; t += 10 {
		go func(t uint32) {
			limit := t + 10
			if limit > end {
				limit = end
			}
			m.jxyBGY(m.zKRpQupzoQE(users, t, limit))
		}(t)
	}
}
func (m *E3b13Z3FO) zKRpQupzoQE(users []*U.QKaIWV2, start, end uint32) *CWAyuvwV9 { // 0x1d14360
	count := 0
	for _, u := range users {
		count += len(u.GetUserIds())
	}
	hashes := make(map[[16]byte][2]uint32, count*int(end-start))
	var stamp [8]byte
	var digest [16]byte
	for _, u := range users {
		for _, id := range u.GetUserIds() {
			h := hmac.New(md5.New, id.Yfd33z6YYdJa[:])
			for t := start; t < end; t++ {
				binary.BigEndian.PutUint64(stamp[:], uint64(t))
				if t != start {
					h.Reset()
				}
				h.Write(stamp[:])
				h.Sum(digest[:0])
				hashes[digest] = [2]uint32{uint32(u.GetId()), t}
			}
		}
	}
	return &CWAyuvwV9{hashes, start}
}
func (m *E3b13Z3FO) jxyBGY(bucket *CWAyuvwV9) { // 0x1d14640
	m.dOa0Z1B.Lock()
	m.hz3qQl7H5 = append(m.hz3qQl7H5, bucket)
	sort.Sort(m.hz3qQl7H5)
	m.dOa0Z1B.Unlock()
}
func (m *E3b13Z3FO) GenAuthForNewUsers(users []*U.QKaIWV2) { m.nlX_TBEa.Store(&users) } // 0x1d14760
func (m *E3b13Z3FO) Get(hash [16]byte) (uint32, uint32, bool) { // 0x1d14840
	m.dOa0Z1B.RLock()
	defer m.dOa0Z1B.RUnlock()
	for left, right := len(m.hz3qQl7H5)/2-1, len(m.hz3qQl7H5)/2; left >= 0 || right < len(m.hz3qQl7H5); left, right = left-1, right+1 {
		if left >= 0 {
			if v, ok := m.hz3qQl7H5[left].iDoMs2[hash]; ok {
				return v[0], v[1], true
			}
		}
		if right < len(m.hz3qQl7H5) {
			if v, ok := m.hz3qQl7H5[right].iDoMs2[hash]; ok {
				return v[0], v[1], true
			}
		}
	}
	return 0, 0, false
}
func (m *E3b13Z3FO) ajZZhr(before uint32) { // 0x1d14a80
	m.dOa0Z1B.Lock()
	defer m.dOa0Z1B.Unlock()
	n := 0
	for n < len(m.hz3qQl7H5) && m.hz3qQl7H5[n].sKKwaeHJ7J < before {
		m.hz3qQl7H5[n] = nil
		n++
	}
	if n > 0 {
		m.hz3qQl7H5 = m.hz3qQl7H5[n:]
	}
}
func (m *E3b13Z3FO) hzdGla95() { // 0x1d14c00
	m.dOa0Z1B.Lock()
	defer m.dOa0Z1B.Unlock()
	m.hz3qQl7H5 = MuywMbTwQ{}
}
