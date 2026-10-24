package CAodbhZf

import "sync/atomic"

// EahpvU_iseQ 对应 16 字节原子值包装；原包装结构允许存储 nil 指针。
// Load/Store/Swap/CompareAndSwap：0x188f000—0x188f1e0。
type EahpvU_iseQ[T any] struct{ value atomic.Value }
type ezkt4tvlO[T any] struct{ value T }

func (v *EahpvU_iseQ[T]) Load() (zero T) {
	if x := v.value.Load(); x != nil {
		return x.(ezkt4tvlO[T]).value
	}
	return
}
func (v *EahpvU_iseQ[T]) Store(x T) { v.value.Store(ezkt4tvlO[T]{x}) }
func (v *EahpvU_iseQ[T]) Swap(x T) (zero T) {
	if previous := v.value.Swap(ezkt4tvlO[T]{x}); previous != nil {
		return previous.(ezkt4tvlO[T]).value
	}
	return
}
func (v *EahpvU_iseQ[T]) CompareAndSwap(old, next T) bool {
	return v.value.CompareAndSwap(ezkt4tvlO[T]{old}, ezkt4tvlO[T]{next})
}
