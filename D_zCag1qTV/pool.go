package D_zCag1qTV

import (
	"math/bits"
	"sync"
)

// RTTI dy6KhjOW1Tn.BPRhvK17，两个切片字段，amd64 大小 48 字节。
type BPRhvK17 struct {
	xF8TJyuJtUf  []sync.Pool
	bTG6g8JP7BLD []int
}

// 对应原 init.0（0xe98f60）的共享实例；导出名由分析者指定。
var PoolAt50961D8 = ZCb00F()

// ZCb00F（0xe98fa0）创建原 27 个容量档位。
func ZCb00F() *BPRhvK17 {
	pool := &BPRhvK17{xF8TJyuJtUf: make([]sync.Pool, 27), bTG6g8JP7BLD: make([]int, 27)}
	for i := range pool.xF8TJyuJtUf {
		size := 1 << i
		if size > 0x4000 {
			size = ((i - 14) << 12) + 0x4000
		}
		// 两个原闭包 0xe99700/0xe99760 均返回 make([]byte, size)。
		pool.xF8TJyuJtUf[i].New = func() any { return make([]byte, size) }
		pool.bTG6g8JP7BLD[i] = size
	}
	return pool
}

// 此辅助名由分析者指定，计算式来自 Get 和 Put 中相同的指令。
func sizeClass(size int) int {
	if size > 0x4000 {
		class := (size - 0x4000) >> 12
		if size&0xfff != 0 {
			return class + 15
		}
		return class + 14
	}
	high := bits.Len64(uint64(2*uint32(size)+1)) - 1
	class := high - 1
	if size != 1<<class {
		return high
	}
	return class
}

// Get（0xe99220）只重设长度，不清空池中已有内容。
func (pool *BPRhvK17) Get(size int) []byte {
	if size <= 0 {
		return nil
	}
	if size > 0x10000 {
		return make([]byte, size)
	}
	return pool.xF8TJyuJtUf[sizeClass(size)].Get().([]byte)[:size]
}

// Put（0xe99360）只有容量恰好等于档位大小时才归还。
func (pool *BPRhvK17) Put(data []byte) {
	capacity := cap(data)
	if capacity == 0 || capacity > 0x10000 {
		return
	}
	class := sizeClass(capacity)
	if pool.bTG6g8JP7BLD[class] == capacity {
		pool.xF8TJyuJtUf[class].Put(data)
	}
}

// BE1pnyFd（0xe994a0）将切片归还给原共享池。
func BE1pnyFd(data []byte) { PoolAt50961D8.Put(data) }
