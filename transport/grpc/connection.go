package grpc

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"recovered.local/soga/transport/grpcproto"
	"recovered.local/soga/util/deadline2"
)

// QAUsR_ub1j：原接口 RTTI 的三个方法。
type QAUsR_ub1j interface {
	Context() context.Context
	Recv() (*grpcproto.Z27tYEZ9hy1o, error)
	Send(*grpcproto.Z27tYEZ9hy1o) error
}

type wtSXw52c8Wt4 struct {
	pmucGnkB5N6I QAUsR_ub1j
	dHRknat6     *bytes.Reader
}
type JOPzd1Giq struct {
	hbenTdWXvxzx QAUsR_ub1j
	wvmTP8N9hd   *deadline2.Q5KYzU
	btyTfM6xlk   context.CancelFunc
	bUtNnscX5XA5 net.Addr
	cXB7IzhpTQG  net.Addr
}

// HpSkqj @ 0x18ae4e0。保留原 peer 与转发头的优先级和端口继承。
func HpSkqj(stream QAUsR_ub1j, local net.Addr, cancel context.CancelFunc) *JOPzd1Giq {
	conn := &JOPzd1Giq{hbenTdWXvxzx: stream, wvmTP8N9hd: deadline2.NewReader(&wtSXw52c8Wt4{pmucGnkB5N6I: stream}), btyTfM6xlk: cancel, bUtNnscX5XA5: local}
	if p, ok := peer.FromContext(stream.Context()); ok {
		conn.cXB7IzhpTQG = p.Addr
		if headers, ok := metadata.FromIncomingContext(stream.Context()); ok {
			if ip, ok := bcVDUGx(headers); ok {
				port := 0
				switch a := p.Addr.(type) {
				case *net.TCPAddr:
					port = a.Port
				case *net.UDPAddr:
					port = a.Port
				}
				conn.cXB7IzhpTQG = &net.TCPAddr{IP: ip, Port: port}
			}
		}
	} else {
		conn.cXB7IzhpTQG = &net.TCPAddr{IP: net.IP{0, 0, 0, 0}, Port: 0}
	}
	return conn
}

// bcVDUGx @ 0x18ae920：原解码常量 x-real-ip 优先于 x-forwarded-for。
func bcVDUGx(headers metadata.MD) (net.IP, bool) {
	if ip, ok := sskxMZZytSNW(headers[strings.ToLower("x-real-ip")]); ok {
		return ip, true
	}
	return sskxMZZytSNW(headers[strings.ToLower("x-forwarded-for")])
}

// sskxMZZytSNW @ 0x18aea40：每个值先整体解析，再从逗号分段中取首个可解析 IP。
func sskxMZZytSNW(values []string) (net.IP, bool) {
	for _, value := range values {
		if len(value) == 0 {
			continue
		}
		if ip := vqwoT_Ao(strings.TrimSpace(value)); ip != nil {
			return ip, true
		}
		for _, part := range strings.Split(value, ",") {
			if ip := vqwoT_Ao(strings.TrimSpace(part)); ip != nil {
				return ip, true
			}
		}
	}
	return nil, false
}

// vqwoT_Ao @ 0x18aeb80。端口只参与 SplitHostPort，不进行数值校验。
func vqwoT_Ao(value string) net.IP {
	if len(value) == 0 {
		return nil
	}
	if ip := net.ParseIP(value); ip != nil {
		return ip
	}
	host, _, err := net.SplitHostPort(value)
	if err != nil {
		return nil
	}
	return net.ParseIP(strings.Trim(host, "[]"))
}

// Read @ 0x18aece0：空 frame 返回 0,nil；Recverror 转成原固定错误，不包装原因。
func (r *wtSXw52c8Wt4) Read(p []byte) (int, error) {
	if r.dHRknat6 == nil {
		message, err := r.pmucGnkB5N6I.Recv()
		if err != nil {
			return 0, errors.New("unable to read from gun tunnel")
		}
		r.dHRknat6 = bytes.NewReader(message.Q2bCNbJ12i)
	}
	n, err := r.dHRknat6.Read(p)
	if r.dHRknat6.Len() == 0 || err == io.EOF {
		r.dHRknat6 = nil
		err = nil
	}
	return n, err
}
func (c *JOPzd1Giq) Read(p []byte) (int, error) { return c.wvmTP8N9hd.Read(p) } // 0x18aefc0
func (c *JOPzd1Giq) Write(p []byte) (int, error) { // 0x18af020
	if err := c.hbenTdWXvxzx.Send(&grpcproto.Z27tYEZ9hy1o{Q2bCNbJ12i: p}); err != nil {
		return 0, errors.New("unable to send data over gun")
	}
	return len(p), nil
}
func (c *JOPzd1Giq) Close() error {
	if c.btyTfM6xlk != nil {
		c.btyTfM6xlk()
	}
	return nil
}                                                      // 0x18af1a0
func (c *JOPzd1Giq) LocalAddr() net.Addr               { return c.bUtNnscX5XA5 }                  // 0x18af1e0
func (c *JOPzd1Giq) RemoteAddr() net.Addr              { return c.cXB7IzhpTQG }                   // 0x18af200
func (c *JOPzd1Giq) SetDeadline(t time.Time) error     { return c.wvmTP8N9hd.SetReadDeadline(t) } // 0x18af220
func (c *JOPzd1Giq) SetReadDeadline(t time.Time) error { return c.wvmTP8N9hd.SetReadDeadline(t) } // 0x18af280
func (c *JOPzd1Giq) SetWriteDeadline(time.Time) error  { return nil }                             // 0x18af2e0，原明确不处理。
