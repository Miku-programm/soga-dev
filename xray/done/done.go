// 此小包依据样本中的 soga-go/xray/done 符号恢复，不引入外部完整内核。
package done

import "sync"

type Instance struct {
	access sync.Mutex
	c      chan struct{}
	closed bool
}

// New 表达原调用点内联的初始化；样本没有独立的这个构造符号。
func New() *Instance { return &Instance{c: make(chan struct{})} }
func (d *Instance) Done() bool { // 0x124f8e0
	select {
	case <-d.c:
		return true
	default:
		return false
	}
}
func (d *Instance) Wait() <-chan struct{} { return d.c } // 0x124f940
func (d *Instance) Close() error { // 0x124f960
	d.access.Lock()
	defer d.access.Unlock()
	if !d.closed {
		d.closed = true
		close(d.c)
	}
	return nil
}
