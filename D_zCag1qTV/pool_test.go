package D_zCag1qTV

import (
	"sync"
	"testing"
	"unsafe"
)

func TestPoolOriginalLayoutAndAllSizeClasses(t *testing.T) {
	pool := ZCb00F()
	if unsafe.Sizeof(*pool) != 48 || unsafe.Offsetof(pool.bTG6g8JP7BLD) != 24 || len(pool.xF8TJyuJtUf) != 27 {
		t.Fatal("RTTI layout differs")
	}
	if pool.Get(-1) != nil || pool.Get(0) != nil {
		t.Fatal("nonpositive Get must return nil")
	}
	for size := 1; size <= 65537; size++ {
		want := size
		if size <= 16384 {
			want = 1
			for want < size {
				want *= 2
			}
		} else if size <= 65536 {
			want = ((size + 4095) / 4096) * 4096
		}
		data := pool.Get(size)
		if len(data) != size || cap(data) != want {
			t.Fatalf("Get(%d): len %d cap %d, want %d", size, len(data), cap(data), want)
		}
		pool.Put(data)
	}
	pool.Put(nil)
	pool.Put(make([]byte, 1001))
	pool.Put(make([]byte, 17001))
	pool.Put(make([]byte, 65537))
}

func TestPoolRejectsForeignCapacities(t *testing.T) {
	pool := ZCb00F()
	// 任意非档位容量若误入池，Get 的容量或切片边界会暴露该错误。
	for _, capacity := range []int{3, 7, 1001, 16385, 20000, 65535} {
		pool.Put(make([]byte, 1, capacity))
		data := pool.Get(capacity + 1)
		if cap(data) == capacity {
			t.Fatalf("accepted foreign cap %d", capacity)
		}
	}
}

func TestSharedPoolConcurrentUse(t *testing.T) {
	var wg sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				data := PoolAt50961D8.Get(8192 + worker)
				for j := range data {
					data[j] = byte(worker)
				}
				for _, b := range data {
					if b != byte(worker) {
						t.Error("buffer shared by concurrent consumers")
						break
					}
				}
				BE1pnyFd(data[:1])
			}
		}(worker)
	}
	wg.Wait()
}
