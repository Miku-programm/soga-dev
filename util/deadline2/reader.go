package deadline2

import (
	"io"
	"os"
	"recovered.local/soga/buf/pool"
	"time"
)

type ypZm4YFaA struct {
	pwylHDJMbOdZ []byte
	kVjk1jW      error
}

// Q5KYzU：原 RTTI 80 字节。
type Q5KYzU struct {
	ncAEio       io.Reader
	e2ItPDQzgp   *TErFW7l9B2S
	g0cpWB       chan ypZm4YFaA
	da86H9jr     []byte
	fZxJxjHPB2Ny []byte
}

// NewReader 表达 HpSkqj 中原内联构造，恢复时新增名称。
func NewReader(reader io.Reader) *Q5KYzU                   { return &Q5KYzU{ncAEio: reader, e2ItPDQzgp: NewDeadline()} }
func (r *Q5KYzU) SetReadDeadline(deadline time.Time) error { r.e2ItPDQzgp.Set(deadline); return nil } // 0x124f260

// Read @ 0x124f2c0。已缓冲数据优先于过期检查；超时后保留正在运行的读取。
// 原函数不保证多个 Read 并发安全，也没有在后续分片保留此前的读取错误。
func (r *Q5KYzU) Read(p []byte) (int, error) {
	if len(r.da86H9jr) != 0 {
		n := copy(p, r.da86H9jr)
		r.da86H9jr = r.da86H9jr[n:]
		if len(r.da86H9jr) == 0 {
			pool.PoolAt50961D8.Put(r.fZxJxjHPB2Ny)
			r.fZxJxjHPB2Ny = nil
		}
		return n, nil
	}
	if r.g0cpWB == nil {
		if !r.e2ItPDQzgp.Active() {
			return r.ncAEio.Read(p)
		}
		select {
		case <-r.e2ItPDQzgp.Wait():
			return 0, os.ErrDeadlineExceeded
		default:
		}
		channel := make(chan ypZm4YFaA, 1)
		buffer := pool.PoolAt50961D8.Get(len(p))
		go func() { // 原 Read.func1 @ 0x124f840；切片边界按原值检查。
			n, err := r.ncAEio.Read(buffer)
			channel <- ypZm4YFaA{pwylHDJMbOdZ: buffer[:n], kVjk1jW: err}
		}()
		r.g0cpWB = channel
	}
	select {
	case <-r.e2ItPDQzgp.Wait():
		return 0, os.ErrDeadlineExceeded
	case result := <-r.g0cpWB:
		r.g0cpWB = nil
		n := copy(p, result.pwylHDJMbOdZ)
		if len(result.pwylHDJMbOdZ) > n {
			r.da86H9jr = result.pwylHDJMbOdZ[n:]
			r.fZxJxjHPB2Ny = result.pwylHDJMbOdZ
			return n, nil
		}
		pool.PoolAt50961D8.Put(result.pwylHDJMbOdZ)
		return n, result.kVjk1jW
	}
}
