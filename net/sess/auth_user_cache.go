package sess

import (
	"container/list"
	"context"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"recovered.local/soga/util/rtx"
	C "recovered.local/soga/config"
	U "recovered.local/soga/user"
)

// 原 RTTI 分别为 24、48、40 字节；字段次序保持原布局。
type HRRGg0nld struct {
	mHY8JyIJ6 U.PgaNd3014U
	jMdUpB    int64
}
type Elkpyuoj struct {
	nszHnbL5B sync.RWMutex
	ite4WB58  time.Duration
	gvgrVQca  *list.List
	zerMi8z   map[int64]*list.Element
}
type EafxyUZD struct {
	aRCMnhMrKi  sync.RWMutex
	dd_Dh_6Zne  time.Duration
	eGhvKTJIaDi *ttlcache.Cache[string, *Elkpyuoj]
}

// YCcacra @0x187d060：按原配置创建缓存，每小时清理两层过期项。
func YCcacra(ctx context.Context) *EafxyUZD {
	ttl := time.Duration(C.IPUserCacheHours()) * time.Hour
	c := &EafxyUZD{dd_Dh_6Zne: ttl, eGhvKTJIaDi: ttlcache.New(ttlcache.WithTTL[string, *Elkpyuoj](ttl))}
	rtx.SwUrlM3M(ctx, func() {
		c.aRCMnhMrKi.Lock()
		c.eGhvKTJIaDi.DeleteExpired()
		c.eGhvKTJIaDi.Range(func(item *ttlcache.Item[string, *Elkpyuoj]) bool { item.Value().DeleteExpired(); return true })
		c.aRCMnhMrKi.Unlock()
	}, time.Hour)
	return c
}

// AddUser @0x1884960：已有 ID 不替换用户，也不刷新过期时间。
func (l *Elkpyuoj) AddUser(u U.PgaNd3014U, expires int64) {
	l.nszHnbL5B.Lock()
	defer l.nszHnbL5B.Unlock()
	if _, ok := l.zerMi8z[u.GetId()]; ok {
		return
	}
	if expires <= 0 {
		expires = time.Now().Add(l.ite4WB58).UnixMilli()
	}
	e := l.gvgrVQca.PushFront(&HRRGg0nld{u, expires})
	l.zerMi8z[u.GetId()] = e
}

// RemoveUser @0x1884ca0。
func (l *Elkpyuoj) RemoveUser(u U.PgaNd3014U) {
	l.nszHnbL5B.Lock()
	defer l.nszHnbL5B.Unlock()
	if e, ok := l.zerMi8z[u.GetId()]; ok {
		l.gvgrVQca.Remove(e)
		delete(l.zerMi8z, u.GetId())
	}
}

// Range @0x1884e40：先回调，再清除过期项；回调 false 表示命中并停止。
// 命中时可续期并移到队首，即使该项在调用前已经过期。
func (l *Elkpyuoj) Range(fn func(*HRRGg0nld) bool, touch bool) {
	l.nszHnbL5B.Lock()
	defer l.nszHnbL5B.Unlock()
	for e := l.gvgrVQca.Front(); e != nil; {
		next := e.Next()
		v := e.Value.(*HRRGg0nld)
		now := time.Now()
		if !fn(v) {
			if touch {
				v.jMdUpB = now.Add(l.ite4WB58).UnixMilli()
				l.gvgrVQca.MoveToFront(e)
			}
			break
		}
		if v.jMdUpB <= now.UnixMilli() {
			l.gvgrVQca.Remove(e)
			delete(l.zerMi8z, v.mHY8JyIJ6.GetId())
		}
		e = next
	}
}
func (l *Elkpyuoj) Count() int { // 0x1885260
	l.nszHnbL5B.RLock()
	defer l.nszHnbL5B.RUnlock()
	return l.gvgrVQca.Len()
}
func (l *Elkpyuoj) DeleteExpired() { // 0x1885320
	l.nszHnbL5B.Lock()
	defer l.nszHnbL5B.Unlock()
	now := time.Now().UnixMilli()
	for e := l.gvgrVQca.Front(); e != nil; {
		next := e.Next()
		v := e.Value.(*HRRGg0nld)
		if v.jMdUpB <= now {
			l.gvgrVQca.Remove(e)
			delete(l.zerMi8z, v.mHY8JyIJ6.GetId())
		}
		e = next
	}
}
func (c *EafxyUZD) RangeUser(ip string, fn func(U.PgaNd3014U) bool) { // 0x187d1e0
	c.aRCMnhMrKi.RLock()
	defer c.aRCMnhMrKi.RUnlock()
	if item := c.eGhvKTJIaDi.Get(ip); item != nil {
		item.Value().Range(func(v *HRRGg0nld) bool { return fn(v.mHY8JyIJ6) }, true)
	}
}
func (c *EafxyUZD) AddUser(ip string, u U.PgaNd3014U, expires int64) { // 0x187d360
	c.aRCMnhMrKi.Lock()
	defer c.aRCMnhMrKi.Unlock()
	var l *Elkpyuoj
	if item := c.eGhvKTJIaDi.Get(ip); item != nil {
		l = item.Value()
	} else {
		l = &Elkpyuoj{ite4WB58: c.dd_Dh_6Zne, gvgrVQca: list.New(), zerMi8z: make(map[int64]*list.Element)}
		c.eGhvKTJIaDi.Set(ip, l, ttlcache.DefaultTTL)
	}
	l.AddUser(u, expires)
}
func (c *EafxyUZD) RemoveUsers(users []U.PgaNd3014U) { // 0x187d5a0
	if len(users) == 0 {
		return
	}
	c.aRCMnhMrKi.Lock()
	defer c.aRCMnhMrKi.Unlock()
	c.eGhvKTJIaDi.Range(func(item *ttlcache.Item[string, *Elkpyuoj]) bool {
		for _, u := range users {
			item.Value().RemoveUser(u)
		}
		return true
	})
}
