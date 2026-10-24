package VasCJz

import (
	"container/heap"
	"errors"
	"io"
	B "recovered.local/soga/D_zCag1qTV"
	"runtime"
	"sync"
)

// 原 RTTI：72 字节的数据块；F6d 是未读部分，Au22 是归还池的原切片。
type UmYbKhCpX struct {
	ACMt1h1     io.ReadCloser
	F6dEJnE3    []byte
	Au22AAVkjZY []byte
	JMZzwbXWi   uint64
}

// 原 RTTI 88 字节。读路径按原实现由单个消费者调用。
type y0bchQKJ struct {
	lINtYg6Bh    io.ReadCloser
	gfws0DVq     bool
	dL_ULZL      chan UmYbKhCpX
	nqQf4VIYqWLJ sync.Mutex
	vPdgI_lx     eMSUo20X
	bDCX9Qu      uint64
	p5CcN5dU3    bool
	fAEWeMURGLm0 int
}
type eMSUo20X []UmYbKhCpX

// 原入口 0x126aa80；锁覆盖阻塞发送，不能改变成丢包或非阻塞发送。
func (q *y0bchQKJ) Push(p UmYbKhCpX) error {
	q.nqQf4VIYqWLJ.Lock()
	defer q.nqQf4VIYqWLJ.Unlock()
	if q.p5CcN5dU3 {
		return errors.New("packet queue closed")
	}
	if q.gfws0DVq {
		return errors.New("h.reader already exists")
	}
	if p.ACMt1h1 != nil {
		q.gfws0DVq = true
	}
	q.dL_ULZL <- p
	return nil
}

// 原入口 0x126acc0。首次关闭先 Gosched，再非阻塞排空通道及堆。
// 原反编译文本遗漏排空循环，以下顺序据 0x126add5—0x126af11 汇编。
func (q *y0bchQKJ) Close() error {
	q.nqQf4VIYqWLJ.Lock()
	defer q.nqQf4VIYqWLJ.Unlock()
	if !q.p5CcN5dU3 {
		q.p5CcN5dU3 = true
		runtime.Gosched()
	drain:
		for {
			select {
			case p, ok := <-q.dL_ULZL:
				if !ok {
					break drain
				}
				if p.ACMt1h1 != nil {
					q.lINtYg6Bh = p.ACMt1h1
				}
				if p.Au22AAVkjZY != nil {
					B.PoolAt50961D8.Put(p.Au22AAVkjZY)
				}
			default:
				break drain
			}
		}
		q.vPdgI_lx.l3NKvduJ6()
		close(q.dL_ULZL)
	}
	// 重复 Close 仍会再次调用 reader.Close，保留样本行为。
	if q.lINtYg6Bh != nil {
		return q.lINtYg6Bh.Close()
	}
	return nil
}

// 原入口 0x126af60。缺号时等待，旧号丢弃，溢出判断位于弹出最小项之后。
func (q *y0bchQKJ) Read(dst []byte) (int, error) {
	if q.lINtYg6Bh != nil {
		return q.lINtYg6Bh.Read(dst)
	}
	if q.p5CcN5dU3 {
		return 0, io.EOF
	}
	if len(q.vPdgI_lx) == 0 {
		p, ok := <-q.dL_ULZL
		if !ok {
			return 0, io.EOF
		}
		if p.ACMt1h1 != nil {
			q.lINtYg6Bh = p.ACMt1h1
			return q.lINtYg6Bh.Read(dst)
		}
		heap.Push(&q.vPdgI_lx, p)
	}
	for len(q.vPdgI_lx) > 0 {
		p := heap.Pop(&q.vPdgI_lx).(UmYbKhCpX)
		if p.JMZzwbXWi < q.bDCX9Qu {
			continue
		}
		if p.JMZzwbXWi > q.bDCX9Qu {
			if len(q.vPdgI_lx) > q.fAEWeMURGLm0 {
				return 0, errors.New("packet queue is too large")
			}
			heap.Push(&q.vPdgI_lx, p)
			next, ok := <-q.dL_ULZL
			if !ok {
				return 0, io.EOF
			}
			heap.Push(&q.vPdgI_lx, next)
			continue
		}
		n := copy(dst, p.F6dEJnE3)
		if n < len(p.F6dEJnE3) {
			p.F6dEJnE3 = p.F6dEJnE3[n:]
			heap.Push(&q.vPdgI_lx, p)
		} else {
			q.bDCX9Qu = p.JMZzwbXWi + 1
			if p.Au22AAVkjZY != nil {
				B.PoolAt50961D8.Put(p.Au22AAVkjZY)
			}
		}
		return n, nil
	}
	return 0, nil
}

func (h *eMSUo20X) Push(x any) { *h = append(*h, x.(UmYbKhCpX)) } // 0x126b700
func (h *eMSUo20X) Pop() any { // 0x126b8a0；保留容量及底层最后一项。
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}
func (h eMSUo20X) Len() int           { return len(h) }
func (h eMSUo20X) Less(i, j int) bool { return h[i].JMZzwbXWi < h[j].JMZzwbXWi }
func (h eMSUo20X) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *eMSUo20X) l3NKvduJ6() { // 0x126b980；归还池后仅清空两份切片。
	for i := range *h {
		p := &(*h)[i]
		if p.Au22AAVkjZY != nil {
			B.PoolAt50961D8.Put(p.Au22AAVkjZY)
			p.Au22AAVkjZY = nil
		}
		p.F6dEJnE3 = nil
	}
	*h = (*h)[:0]
}
