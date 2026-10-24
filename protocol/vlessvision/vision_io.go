package vlessvision

import (
	"bytes"
	"io"
	"net"

	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/sess"
)

// 以下类型及行为由原样本 RTTI、指令和调用关系还原，并非作者文件原文。
// LZ7AGw7 与 Qx58rV 原布局分别为80、72字节。
type LZ7AGw7 struct {
	io.Reader
	fB4afneh7k *sess.Mdn09F6aMrpa
	woJdOhbWO  *MxS0kZ
	sztBlp     *bytes.Reader
	k9rmTO     *bytes.Buffer
	lOw1FJn4V  bool
	ei1r2yb    bool
	dx5b6nPL   net.Conn
	oqHPkr     *pool.X_aXX5
}

type Qx58rV struct {
	io.Writer
	uYYP2p      *sess.Mdn09F6aMrpa
	taFXeH      *MxS0kZ
	oRVa2PobW_S []byte
	gljtGU3     net.Conn
}

// gT73UNf4 @0x1d08ee0。先耗尽TLS内存中的明文/原始输入，随后读取底层连接；
// 只有最后一种读取追加用户流量，不能给此前已计费的内存数据再次计费。
func (r *LZ7AGw7) gT73UNf4(p []byte) (int, error) {
	if !r.lOw1FJn4V && r.sztBlp.Len() > 0 {
		n, err := r.sztBlp.Read(p)
		if r.sztBlp.Len() == 0 {
			r.lOw1FJn4V = true
		}
		return n, err
	}
	if !r.ei1r2yb && r.k9rmTO.Len() > 0 {
		n, err := r.k9rmTO.Read(p)
		if r.k9rmTO.Len() == 0 {
			r.ei1r2yb = true
		}
		return n, err
	}
	if r.dx5b6nPL == nil {
		r.dx5b6nPL = CocBnud6(r.fB4afneh7k)
	}
	n, err := r.dx5b6nPL.Read(p)
	if n > 0 {
		r.fB4afneh7k.AddUp(n)
	}
	return n, err
}

// Read @0x1d091c0。底层一次返回数据和错误时，原函数先返回0/error，
// 缓冲中的数据仍在，下一次调用会直接取出；不改写为惯常的n/error处理。
func (r *LZ7AGw7) Read(p []byte) (int, error) {
	if r.oqHPkr != nil && r.oqHPkr.Len() > 0 {
		return r.oqHPkr.Read(p)
	}
	if r.woJdOhbWO.SDw6IIM1D51U {
		r.Free()
		return r.gT73UNf4(p)
	}
	r.oqHPkr.Reset()
	if _, err := r.oqHPkr.ReadMaxFrom(r.Reader, 8192); err != nil {
		return 0, err
	}
	out := r.oqHPkr
	state := r.woJdOhbWO
	if state.EO4y9oyDkzq || state.BUwqtS > 0 {
		out = DtcbafIY7SUd(out, state)
		if state.ZOEKPa7JPx5 <= 0 && state.PIsBMe <= 0 && state.EXbj1L5 != 0 {
			switch state.EXbj1L5 {
			case 1:
				state.EO4y9oyDkzq = false
			case 2:
				state.EO4y9oyDkzq = false
				state.SDw6IIM1D51U = true
			}
		} else {
			state.EO4y9oyDkzq = true
		}
	}
	if state.BUwqtS > 0 {
		YByzFSgf(out.Bytes(), state)
	}
	if out != r.oqHPkr {
		defer out.Free()
	}
	if out.IsEmpty() {
		r.oqHPkr.Reset()
		return 0, nil
	}
	n := copy(p, out.Bytes())
	out.Skip(n)
	if out != r.oqHPkr {
		r.oqHPkr.Reset()
		if !out.IsEmpty() {
			_, _ = r.oqHPkr.Write(out.Bytes())
		}
	}
	return n, nil
}

func (r *LZ7AGw7) Free() { // @0x1d096a0
	if r.oqHPkr != nil {
		r.oqHPkr.Free()
		r.oqHPkr = nil
	}
}

func (w *Qx58rV) oo4zzh8AE1b(p []byte) (int, error) { // @0x1d09720
	if w.gljtGU3 == nil {
		w.gljtGU3 = CocBnud6(w.uYYP2p)
	}
	n, err := w.gljtGU3.Write(p)
	if n > 0 {
		w.uYYP2p.AddDown(n)
	}
	return n, err
}

// Write @0x1d09840。原空写会返回线上padding帧的Write返回值；普通写在
// 任一底层写入错误时返回0，成功时返回原负载长度，忽略无错误的短写。
func (w *Qx58rV) Write(p []byte) (int, error) {
	state := w.taFXeH
	if state.Pd9MU7P6IX {
		return w.oo4zzh8AE1b(p)
	}
	if state.BUwqtS > 0 {
		YByzFSgf(p, state)
	}
	if !state.ZkTHd3 {
		return w.Writer.Write(p)
	}
	if len(p) == 0 {
		out := pool.S4OQXWIZ(2048)
		defer out.Free()
		IFq7RHaxl(out, nil, 0, &w.oRVa2PobW_S, true)
		return w.Writer.Write(out.Bytes())
	}
	parts := AaBEhR(nil, p)
	longPadding := state.VRlXaTaT
	out := pool.S4OQXWIZ(len(p))
	defer out.Free()
	var remainder [][]byte
	for i, part := range parts {
		if state.VRlXaTaT && len(part) >= 6 && bytes.Equal(part[:3], bytesAt5036EB0) {
			if state.YdJioLGvuLW {
				state.Pd9MU7P6IX = true
			}
			var command byte
			if i == len(parts)-1 {
				command = 1
				if state.YdJioLGvuLW {
					command = 2
				}
			}
			IFq7RHaxl(out, part, command, &w.oRVa2PobW_S, true)
			state.ZkTHd3 = false
			longPadding = false
		} else if state.UFPiQl0bbHk || state.BUwqtS > 1 {
			var command byte
			if i == len(parts)-1 && !state.ZkTHd3 {
				command = 1
				if state.YdJioLGvuLW {
					command = 2
				}
			}
			IFq7RHaxl(out, part, command, &w.oRVa2PobW_S, longPadding)
		} else {
			state.ZkTHd3 = false
			IFq7RHaxl(out, part, 1, &w.oRVa2PobW_S, longPadding)
			remainder = parts[i+1:]
			break
		}
	}
	if _, err := w.Writer.Write(out.Bytes()); err != nil {
		return 0, err
	}
	for _, part := range remainder {
		if _, err := w.Writer.Write(part); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}
