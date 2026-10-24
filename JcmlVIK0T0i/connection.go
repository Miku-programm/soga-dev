package JcmlVIK0T0i

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/HxKViEp2a"
)

// MOTygWJHAwt5：原 104 字节惰性 PROXY 头连接。
type MOTygWJHAwt5 struct {
	fBUa21t     atomic.Value
	qq2AFg0aO   sync.Once
	d52sVi_x7P  error
	y7HtLBfne   net.Conn
	Nh0TcQyJd9  K_M6K7
	aVUuZorW_   *D_zCag1qTV.NADyavA
	eU1FKzIzXH  *URdA97qPFc
	H2qtkr2P    XTAkygU
	gQaLqCURuN1 time.Duration
}

func GXOJV4B0a4D(conn net.Conn, options ...func(*MOTygWJHAwt5)) *MOTygWJHAwt5 { // 0x1185400
	p := &MOTygWJHAwt5{aVUuZorW_: D_zCag1qTV.C8FauP3Auy(conn, 4096), y7HtLBfne: conn}
	for _, option := range options {
		option(p)
	}
	return p
}
func (p *MOTygWJHAwt5) Read(b []byte) (int, error) { // 0x1185560
	p.qq2AFg0aO.Do(func() { p.d52sVi_x7P = p.bkflJp() })
	if p.d52sVi_x7P != nil {
		return 0, p.d52sVi_x7P
	}
	if p.aVUuZorW_ != nil {
		if n := p.aVUuZorW_.Buffered(); n > 0 {
			if len(b) > n {
				b = b[:n]
			}
			return p.aVUuZorW_.Read(b)
		}
		p.aVUuZorW_.FreeBuf()
		p.aVUuZorW_ = nil
	}
	return p.y7HtLBfne.Read(b)
}
func (p *MOTygWJHAwt5) Write(b []byte) (int, error) { return p.y7HtLBfne.Write(b) } // 0x11856a0
func (p *MOTygWJHAwt5) Close() error                { return p.y7HtLBfne.Close() }  // 0x1185700
func (p *MOTygWJHAwt5) ProxyHeader() *URdA97qPFc { // 0x1185740
	p.qq2AFg0aO.Do(func() { p.d52sVi_x7P = p.bkflJp() })
	return p.eU1FKzIzXH
}
func (p *MOTygWJHAwt5) LocalAddr() net.Addr { return p.y7HtLBfne.LocalAddr() } // 0x11857a0
func (p *MOTygWJHAwt5) RemoteAddr() net.Addr { // 0x11857e0
	p.qq2AFg0aO.Do(func() { p.d52sVi_x7P = p.bkflJp() })
	if p.eU1FKzIzXH == nil || p.eU1FKzIzXH.LgI8Cr1feIYb.IsLocal() || p.d52sVi_x7P != nil {
		return p.y7HtLBfne.RemoteAddr()
	}
	return p.eU1FKzIzXH.Dafxa_BOdZdw
}
func (p *MOTygWJHAwt5) Raw() net.Conn { return p.y7HtLBfne } // 0x1185880
func (p *MOTygWJHAwt5) TCPConn() (*net.TCPConn, bool) {
	c, ok := p.y7HtLBfne.(*net.TCPConn)
	return c, ok
} // 0x11858a0
func (p *MOTygWJHAwt5) UnixConn() (*net.UnixConn, bool) {
	c, ok := p.y7HtLBfne.(*net.UnixConn)
	return c, ok
} // 0x11858c0
func (p *MOTygWJHAwt5) UDPConn() (*net.UDPConn, bool) {
	c, ok := p.y7HtLBfne.(*net.UDPConn)
	return c, ok
} // 0x11858e0
func (p *MOTygWJHAwt5) SetDeadline(t time.Time) error {
	p.fBUa21t.Store(t)
	return p.y7HtLBfne.SetDeadline(t)
} // 0x1185900
func (p *MOTygWJHAwt5) SetReadDeadline(t time.Time) error {
	p.fBUa21t.Store(t)
	return p.y7HtLBfne.SetReadDeadline(t)
}                                                          // 0x11859c0
func (p *MOTygWJHAwt5) SetWriteDeadline(t time.Time) error { return p.y7HtLBfne.SetWriteDeadline(t) } // 0x1185a80
func (p *MOTygWJHAwt5) GetRawConn() net.Conn               { return p.y7HtLBfne }                     // 0x1185de0
func (p *MOTygWJHAwt5) Free() { // 0x1185e00
	if p.aVUuZorW_ != nil {
		p.aVUuZorW_.FreeBuf()
	}
	HxKViEp2a.Li3c4CJ1_D(p.y7HtLBfne)
}
func (p *MOTygWJHAwt5) bkflJp() error { // 0x1185ae0
	if p.gQaLqCURuN1 > 0 {
		if err := p.y7HtLBfne.SetReadDeadline(time.Now().Add(p.gQaLqCURuN1)); err != nil {
			return err
		}
	}
	h, err := LavOMmUY2(p.aVUuZorW_)
	if p.gQaLqCURuN1 > 0 {
		var old time.Time
		if value := p.fBUa21t.Load(); value != nil {
			old = value.(time.Time)
		}
		if resetErr := p.y7HtLBfne.SetReadDeadline(old); resetErr != nil {
			return resetErr
		}
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			err = ErrorAt5097E80
		}
	}
	if err == ErrorAt5097E80 {
		if p.H2qtkr2P == 3 {
			return err
		}
		return nil
	}
	if err == nil && h != nil {
		switch p.H2qtkr2P {
		case 0, 3:
			if p.Nh0TcQyJd9 != nil {
				if err := p.Nh0TcQyJd9(h); err != nil {
					return err
				}
			}
			p.eU1FKzIzXH = h
		case 2:
			return ErrorAt5097EF0
		}
	}
	return err
}

// 下面两个工厂为原 WTqa9yT / O8oMdhQ 的内联主体及已导出的闭包。
func WTqa9yT(policy XTAkygU) func(*MOTygWJHAwt5) {
	return func(p *MOTygWJHAwt5) { p.H2qtkr2P = policy }
}
func O8oMdhQ(validate K_M6K7) func(*MOTygWJHAwt5) {
	return func(p *MOTygWJHAwt5) {
		if validate != nil {
			p.Nh0TcQyJd9 = validate
		}
	}
}

type KprDQw4aZEOg struct {
	I9Hb4SAu net.Listener
	Uw20KUu1 AXaIUgeevW2
	GlUfdaIP K_M6K7
	MI_6xtup time.Duration
}

var DefaultHeaderTimeoutAt37879D8 = time.Duration(10000000000)

func (l *KprDQw4aZEOg) Accept() (net.Conn, error) { // 0x1185200
	c, err := l.I9Hb4SAu.Accept()
	if err != nil {
		return nil, err
	}
	var policy XTAkygU
	if l.Uw20KUu1 != nil {
		policy, err = l.Uw20KUu1(c.RemoteAddr())
		if err != nil {
			c.Close()
			return nil, err
		}
		if policy == 4 {
			return c, nil
		}
	}
	p := GXOJV4B0a4D(c, WTqa9yT(policy), O8oMdhQ(l.GlUfdaIP))
	if l.MI_6xtup == 0 {
		l.MI_6xtup = DefaultHeaderTimeoutAt37879D8
	}
	p.gQaLqCURuN1 = l.MI_6xtup
	return p, nil
}
func (l *KprDQw4aZEOg) Close() error   { return l.I9Hb4SAu.Close() } // 0x1185380
func (l *KprDQw4aZEOg) Addr() net.Addr { return l.I9Hb4SAu.Addr() }  // 0x11853c0
