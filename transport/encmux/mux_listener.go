package encmux

import (
	"context"
	"errors"
	"io"
	"net"
	"recovered.local/soga/transport/realityconn"
	"recovered.local/soga/util/logx"
)

// 原 RTTI bhXMvjeYm_m.PRSLFMd，实际包 g0MlUOhvO，amd64 大小 56。
// 以下是原二进制函数的源码表示，并非作者文件原文。
type PRSLFMd struct {
	O6B2qZ0p0
	tJhNCameSED  context.Context
	ih8lN8bRYVzI context.CancelFunc
	j6_Rmz52Z9   chan net.Conn
	vAYswb       *VTl94pzz2IkB
}

// Bn7A4E9vXH 对应 0x11d7240。
func Bn7A4E9vXH(listener net.Listener, config *VTl94pzz2IkB) net.Listener {
	ctx, cancel := context.WithCancel(context.Background())
	v := &PRSLFMd{listener, ctx, cancel, make(chan net.Conn, 256), config}
	v.aLZdoQfgn()
	return v
}

// aLZdoQfgn 对应 0x11d7380；协程/闭包为 0x11e59e0/0x11e5b20。
func (v *PRSLFMd) aLZdoQfgn() {
	go func() {
		defer func() { v.Close() }()
		for {
			conn, err := v.O6B2qZ0p0.Accept()
			if err != nil {
				return
			}
			go v.lHLHbAVjED(conn)
		}
	}()
}

// lHLHbAVjED 对应 0x11d7400。原汇编 0x11d74e8—0x11d7520
// 明确传入禁用 keepalive 的配置，不是 Server 的 nil 默认参数。
func (v *PRSLFMd) lHLHbAVjED(conn net.Conn) {
	defer func() { conn.Close(); logx.TRTbrquD(-4, "close mux server") }()
	logx.TRTbrquD(-4, "start mux server")
	session, err := realityconn.CXFDjl_TYx8(conn, &realityconn.BWpFHH3x{
		RTNOROJC5: 1, DU0iTVAg: true, XsqVEjw4: 32768, K3nkPQuOI: 4194304, CUuHLYfWnC8q: 65536,
	})
	if err != nil {
		logx.TRTbrquD(-4, "start mux server failed:", err)
		return
	}
	defer func() { session.Close() }()
	for {
		stream, err := session.AcceptStream()
		if err != nil {
			if err != io.EOF && !errors.Is(err, realityconn.TimeoutErrorAt502B390()) {
				logx.TRTbrquD(-4, "mux session accept err:", err)
			}
			return
		}
		logx.TRTbrquD(-4, "accept mux stream")
		// 原函数是直接 chansend，没有额外 ctx 分支；保持该阻塞语义。
		v.j6_Rmz52Z9 <- stream
	}
}

// Accept 对应 0x11d78e0；每次关闭错误均重新构造。
func (v *PRSLFMd) Accept() (net.Conn, error) {
	select {
	case c := <-v.j6_Rmz52Z9:
		if c != nil {
			return c, nil
		}
	case <-v.tJhNCameSED.Done():
	}
	return nil, errors.New("closed")
}

// Close 对应 0x11d7a20；原代码不关闭输出 channel，也不主动关闭已接收的 session。
func (v *PRSLFMd) Close() error { v.ih8lN8bRYVzI(); return v.O6B2qZ0p0.Close() }
