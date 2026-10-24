// 依据原 ELF xIF4G3 包机器码及 RTTI 恢复的源码表示，不是作者文件原文。
// 对照基线：apernet/hysteria 9e90d7d，保留其原许可；证据在 hysteria-wire。
package hysteriabrutal

// A RingBuffer is a ring buffer.
// It acts as a heap that doesn't cause any allocations.
type YXV6CyZ5Ux[T any] struct {
	ckRjtgECJ               []T
	rxa2MP2pFk, gwZziNLjai8 int
	aY2XfPJecIWV            bool
}

// Init preallocs a buffer with a certain size.
func (r *YXV6CyZ5Ux[T]) Init(size int) {
	r.ckRjtgECJ = make([]T, size)
}

// Len returns the number of elements in the ring buffer.
func (r *YXV6CyZ5Ux[T]) Len() int {
	if r.aY2XfPJecIWV {
		return len(r.ckRjtgECJ)
	}
	if r.gwZziNLjai8 >= r.rxa2MP2pFk {
		return r.gwZziNLjai8 - r.rxa2MP2pFk
	}
	return r.gwZziNLjai8 - r.rxa2MP2pFk + len(r.ckRjtgECJ)
}

// Empty says if the ring buffer is empty.
func (r *YXV6CyZ5Ux[T]) Empty() bool {
	return !r.aY2XfPJecIWV && r.rxa2MP2pFk == r.gwZziNLjai8
}

// PushBack adds a new element.
// If the ring buffer is full, its capacity is increased first.
func (r *YXV6CyZ5Ux[T]) PushBack(t T) {
	if r.aY2XfPJecIWV || len(r.ckRjtgECJ) == 0 {
		r.grow()
	}
	r.ckRjtgECJ[r.gwZziNLjai8] = t
	r.gwZziNLjai8++
	if r.gwZziNLjai8 == len(r.ckRjtgECJ) {
		r.gwZziNLjai8 = 0
	}
	if r.gwZziNLjai8 == r.rxa2MP2pFk {
		r.aY2XfPJecIWV = true
	}
}

// PopFront returns the next element.
// It must not be called when the buffer is empty, that means that
// callers might need to check if there are elements in the buffer first.
func (r *YXV6CyZ5Ux[T]) PopFront() T {
	if r.Empty() {
		panic("github.com/quic-go/quic-go/internal/utils/ringbuffer: pop from an empty queue")
	}
	r.aY2XfPJecIWV = false
	t := r.ckRjtgECJ[r.rxa2MP2pFk]
	r.ckRjtgECJ[r.rxa2MP2pFk] = *new(T)
	r.rxa2MP2pFk++
	if r.rxa2MP2pFk == len(r.ckRjtgECJ) {
		r.rxa2MP2pFk = 0
	}
	return t
}

// Offset returns the offset element.
// It must not be called when the buffer is empty, that means that
// callers might need to check if there are elements in the buffer first
// and check if the index larger than buffer length.
func (r *YXV6CyZ5Ux[T]) Offset(index int) *T {
	if r.Empty() || index >= r.Len() {
		panic("github.com/quic-go/quic-go/internal/utils/ringbuffer: offset from invalid index")
	}
	offset := (r.rxa2MP2pFk + index) % len(r.ckRjtgECJ)
	return &r.ckRjtgECJ[offset]
}

// Front returns the front element.
// It must not be called when the buffer is empty, that means that
// callers might need to check if there are elements in the buffer first.
func (r *YXV6CyZ5Ux[T]) Front() *T {
	if r.Empty() {
		panic("github.com/quic-go/quic-go/internal/utils/ringbuffer: front from an empty queue")
	}
	return &r.ckRjtgECJ[r.rxa2MP2pFk]
}

// Back returns the back element.
// It must not be called when the buffer is empty, that means that
// callers might need to check if there are elements in the buffer first.
func (r *YXV6CyZ5Ux[T]) Back() *T {
	if r.Empty() {
		panic("github.com/quic-go/quic-go/internal/utils/ringbuffer: back from an empty queue")
	}
	return r.Offset(r.Len() - 1)
}

// Grow the maximum size of the queue.
// This method assume the queue is full.
func (r *YXV6CyZ5Ux[T]) grow() {
	oldRing := r.ckRjtgECJ
	newSize := len(oldRing) * 2
	if newSize == 0 {
		newSize = 1
	}
	r.ckRjtgECJ = make([]T, newSize)
	headLen := copy(r.ckRjtgECJ, oldRing[r.rxa2MP2pFk:])
	copy(r.ckRjtgECJ[headLen:], oldRing[:r.rxa2MP2pFk])
	r.rxa2MP2pFk, r.gwZziNLjai8, r.aY2XfPJecIWV = 0, len(oldRing), false
}

// Clear removes all elements.
func (r *YXV6CyZ5Ux[T]) Clear() {
	var zeroValue T
	for i := range r.ckRjtgECJ {
		r.ckRjtgECJ[i] = zeroValue
	}
	r.rxa2MP2pFk, r.gwZziNLjai8, r.aY2XfPJecIWV = 0, 0, false
}
