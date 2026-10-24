package replaycache

import (
	"recovered.local/soga/bloom"
	"sync"
)

// At6yec9xcC 原RTTI duAfa938mUS.At6yec9xcC，80字节。
type At6yec9xcC struct {
	pg4Okax6g3a  int
	ii_apO       int
	cl8UHzwUk6z1 int
	uSo7A4DU     int
	d6aQwf8rDu   []bloom.Filter
	oKYp41ywPPbH sync.RWMutex
}

// mmCLGRNuEav 对应0x11741c0，依次计算FNV-1和FNV-1a。
func mmCLGRNuEav(data []byte) (uint64, uint64) {
	var a, b uint64 = 0xcbf29ce484222325, 0xcbf29ce484222325
	for _, v := range data {
		a = a*0x100000001b3 ^ uint64(v)
	}
	for _, v := range data {
		b = (b ^ uint64(v)) * 0x100000001b3
	}
	return a, b
}

// KprmqtO 对应0x1174240，除法、分配及零分层panic均保留。
func KprmqtO(layers, capacity int, p float64) *At6yec9xcC {
	filters := make([]bloom.Filter, layers)
	cache := &At6yec9xcC{pg4Okax6g3a: capacity / layers, cl8UHzwUk6z1: layers, d6aQwf8rDu: filters}
	for i := range filters {
		filters[i] = bloom.New(cache.pg4Okax6g3a, p, mmCLGRNuEav)
	}
	return cache
}

// TestAndAdd 对应0x11743c0。原严格count>limit时才轮换，命中任一层返回false。
func (c *At6yec9xcC) TestAndAdd(data []byte) bool {
	if c == nil {
		return false
	}
	c.oKYp41ywPPbH.Lock()
	defer c.oKYp41ywPPbH.Unlock()
	for _, f := range c.d6aQwf8rDu {
		if f.Test(data) {
			return false
		}
	}
	f := c.d6aQwf8rDu[c.ii_apO]
	if c.uSo7A4DU > c.pg4Okax6g3a {
		c.ii_apO = (c.ii_apO + 1) % c.cl8UHzwUk6z1
		f = c.d6aQwf8rDu[c.ii_apO]
		f.Reset()
		c.uSo7A4DU = 0
	}
	c.uSo7A4DU++
	f.Add(data)
	return true
}

var cacheAt50961B8 *At6yec9xcC
var onceAt50C7C20 sync.Once

// ulS5Ci6maP 对应0x1174600，内部闭包0x1174720建立原十层百万容量。
func ulS5Ci6maP() *At6yec9xcC {
	onceAt50C7C20.Do(func() { cacheAt50961B8 = KprmqtO(10, 1000000, 1e-12) })
	return cacheAt50961B8
}

// UJhJgiLn 对应0x1174660，所有原调用共享同一反重放缓存。
func UJhJgiLn(data []byte) bool { return ulS5Ci6maP().TestAndAdd(data) }
