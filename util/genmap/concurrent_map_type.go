package genmap

import (
	"sync"
	"sync/atomic"
)

// GTFAWF 按 RTTI 0x363d160 的两个字段恢复。
type GTFAWF[K comparable, V any] struct {
	r5_Hdcu3GxXC sync.Map
	yaLIYjUJ     atomic.Int64
}
