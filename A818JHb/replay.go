package A818JHb

import (
	"hash/maphash"
	"recovered.local/soga/bloom"
	"sync"
)

// AiTSgb7OJeoK 原RTTI waQhdPH.AiTSgb7OJeoK，80字节。
type AiTSgb7OJeoK struct {
	pg4Okax6g3a  int
	ii_apO       int
	cl8UHzwUk6z1 int
	uSo7A4DU     int
	d6aQwf8rDu   []bloom.Filter
	oKYp41ywPPbH sync.RWMutex
}

// init 对应 0x1c5dba0，原两次循环 runtime.rand，直到非零；MakeSeed 保留此语义。
var seedAt50C79F8 = maphash.MakeSeed()
var seedAt50C7A00 = maphash.MakeSeed()

// g1TDyioZI 对应 0x1c5dbe0，原调用 hash/maphash.Bytes 两次。
func g1TDyioZI(data []byte) (uint64, uint64) {
	return maphash.Bytes(seedAt50C79F8, data), maphash.Bytes(seedAt50C7A00, data)
}

// EJaHm2W 对应0x1c5dc80，除法、分配及零分层panic均保留。
func EJaHm2W(layers, capacity int, p float64) *AiTSgb7OJeoK {
	filters := make([]bloom.Filter, layers)
	cache := &AiTSgb7OJeoK{pg4Okax6g3a: capacity / layers, cl8UHzwUk6z1: layers, d6aQwf8rDu: filters}
	for i := range filters {
		filters[i] = bloom.New(cache.pg4Okax6g3a, p, g1TDyioZI)
	}
	return cache
}

// TestAndAdd 对应0x1c5de00。原严格count>limit时才轮换，命中任一层返回false。
func (c *AiTSgb7OJeoK) TestAndAdd(data []byte) bool {
	if c == nil {
		return true
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

var cacheAt5096958 *AiTSgb7OJeoK
var onceAt50C7D50 sync.Once

// mF1icpp 对应0x1c5e040，内部闭包0x1c5e160建立原十层百万容量。
func mF1icpp() *AiTSgb7OJeoK {
	onceAt50C7D50.Do(func() { cacheAt5096958 = EJaHm2W(10, 1000000, 1e-12) })
	return cacheAt5096958
}

// HFlnXWty 对应0x1c5e0a0，所有原调用共享同一反重放缓存。
func HFlnXWty(data []byte) bool { return mF1icpp().TestAndAdd(data) }
