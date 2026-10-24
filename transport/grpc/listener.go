package grpc

import (
	"context"
	"crypto/tls"
	"errors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"recovered.local/soga/transport/grpcproto"
	"recovered.local/soga/util/logx"
)

type O6B2qZ0p0 = net.Listener

// Q1SOcdpqfGG：原 56 字节结构，空默认服务嵌入与 listener 同在偏移 0。
type Q1SOcdpqfGG struct {
	grpcproto.EmfCw_9j
	O6B2qZ0p0
	m8StEa_m     *grpc.Server
	bvW4Be       context.Context
	jUNp5VibZl9x context.CancelFunc
	w6Om7JaO_bc  chan net.Conn
}

// SNiVo8 @ 0x18b12c0。原构造修改调用方 TLS NextProtos，再创建服务。
func SNiVo8(listener net.Listener, serviceName string, tlsConfig *tls.Config) *Q1SOcdpqfGG {
	var server *grpc.Server
	if tlsConfig != nil {
		tlsConfig.NextProtos = []string{"h2"}
		server = grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))
	} else {
		server = grpc.NewServer()
	}
	ctx, cancel := context.WithCancel(context.Background())
	l := &Q1SOcdpqfGG{O6B2qZ0p0: listener, m8StEa_m: server, bvW4Be: ctx, jUNp5VibZl9x: cancel, w6Om7JaO_bc: make(chan net.Conn, 0x4000)}
	grpcproto.Qpv51XFhNT(server, l, serviceName)
	go func() {
		if err := server.Serve(listener); err != nil {
			logx.TRTbrquD(0, "grpc serve err:", err)
		}
	}()
	return l
}

// Tun @ 0x18b1600。原实现无条件入队，再等待 listener 的子 context 取消。
func (l *Q1SOcdpqfGG) Tun(stream grpcproto.XZFdvGb2) error {
	ctx, cancel := context.WithCancel(l.bvW4Be)
	conn := HpSkqj(stream, l.O6B2qZ0p0.Addr(), cancel)
	l.w6Om7JaO_bc <- conn
	<-ctx.Done()
	return nil
}

// Accept @ 0x18b1760。拿到连接后再次检查取消；关闭刚取出的连接并丢弃关闭错误。
func (l *Q1SOcdpqfGG) Accept() (net.Conn, error) {
	select {
	case conn := <-l.w6Om7JaO_bc:
		select {
		case <-l.bvW4Be.Done():
			if conn != nil {
				_ = conn.Close()
			}
			return nil, errors.New("grpc listener closed")
		default:
			return conn, nil
		}
	case <-l.bvW4Be.Done():
		return nil, errors.New("grpc listener closed")
	}
}

// Close @ 0x18b19a0。取消→非优雅停止 grpc→关闭底层 listener，返回最后一次关闭错误。
func (l *Q1SOcdpqfGG) Close() error { l.jUNp5VibZl9x(); l.m8StEa_m.Stop(); return l.O6B2qZ0p0.Close() }
