package GLaePRKYM

import (
	"net"
	"sync"
	"syscall"
	"time"
)

type v2leTMr struct {
	CljgB3      net.PacketConn
	YC2F3s1     PSIAOZXXoq
	chfYpt5C9   []byte
	g4ZgYMaCfW7 sync.Mutex
	a9nyha3     []byte
	qXrWg_XLkm  sync.Mutex
}

type kIKvVu0qKkaB struct {
	*v2leTMr
	HGmsg9Ffs0Gc *net.UDPConn
}

// WrapFromHysteriaInline 对应 0x1c32170 后的原内联构造。
// 原实现固定使用 2048 字节读写缓冲，且仅为实际 *net.UDPConn 暴露套接字接口。
func WrapFromHysteriaInline(conn net.PacketConn, obfs PSIAOZXXoq) net.PacketConn {
	v := &v2leTMr{CljgB3: conn, YC2F3s1: obfs, chfYpt5C9: make([]byte, 2048), a9nyha3: make([]byte, 2048)}
	if udp, ok := conn.(*net.UDPConn); ok {
		return &kIKvVu0qKkaB{v, udp}
	}
	return v
}

func (v *v2leTMr) ReadFrom(b []byte) (int, net.Addr, error) { // 0x1c2c380
	for {
		v.g4ZgYMaCfW7.Lock()
		n, addr, err := v.CljgB3.ReadFrom(v.chfYpt5C9)
		if n <= 0 {
			v.g4ZgYMaCfW7.Unlock()
			return n, addr, err
		}
		n = v.YC2F3s1.Deobfuscate(v.chfYpt5C9[:n], b)
		v.g4ZgYMaCfW7.Unlock()
		if n > 0 || err != nil {
			return n, addr, err
		}
	}
}

func (v *v2leTMr) WriteTo(b []byte, addr net.Addr) (int, error) { // 0x1c2c580
	v.qXrWg_XLkm.Lock()
	n := v.YC2F3s1.Obfuscate(b, v.a9nyha3)
	_, err := v.CljgB3.WriteTo(v.a9nyha3[:n], addr)
	v.qXrWg_XLkm.Unlock()
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (v *v2leTMr) Close() error                               { return v.CljgB3.Close() }                 // 0x1c2c700
func (v *v2leTMr) LocalAddr() net.Addr                        { return v.CljgB3.LocalAddr() }             // 0x1c2c740
func (v *v2leTMr) SetDeadline(t time.Time) error              { return v.CljgB3.SetDeadline(t) }          // 0x1c2c780
func (v *v2leTMr) SetReadDeadline(t time.Time) error          { return v.CljgB3.SetReadDeadline(t) }      // 0x1c2c7e0
func (v *v2leTMr) SetWriteDeadline(t time.Time) error         { return v.CljgB3.SetWriteDeadline(t) }     // 0x1c2c840
func (v *kIKvVu0qKkaB) SetReadBuffer(n int) error             { return v.HGmsg9Ffs0Gc.SetReadBuffer(n) }  // 0x1c2c8a0
func (v *kIKvVu0qKkaB) SetWriteBuffer(n int) error            { return v.HGmsg9Ffs0Gc.SetWriteBuffer(n) } // 0x1c2c8e0
func (v *kIKvVu0qKkaB) SyscallConn() (syscall.RawConn, error) { return v.HGmsg9Ffs0Gc.SyscallConn() }     // 0x1c2c920
