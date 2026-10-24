package GYzZ5v

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Y5md9Ywqjd4m 是样本 aXzEtV1OS 包的 grpc.ServerStream 类型别名。
// 方法集合由 RTTI 核对；别名保留原包装结构的嵌入字段名称。
type Y5md9Ywqjd4m = grpc.ServerStream

// SCv1eHEMZ 与 XZFdvGb2 的完整方法集合来自原接口 RTTI。
type SCv1eHEMZ interface{ Tun(XZFdvGb2) error }
type XZFdvGb2 interface {
	Send(*Z27tYEZ9hy1o) error
	Recv() (*Z27tYEZ9hy1o, error)
	grpc.ServerStream
}

type EmfCw_9j struct{}

// Tun @ 0xe89a20 是原程序生成的默认服务方法，不是待恢复业务的占位。
// 原错误码 12，错误字符串由 0xe89a4c 的纯解码调用恢复。
func (*EmfCw_9j) Tun(XZFdvGb2) error {
	return status.Errorf(codes.Unimplemented, "method Tun not implemented")
}

type utpvEgZY struct{ Y5md9Ywqjd4m }

// eehVxKUB @ 0xe89a80。原接口断言失败时直接 panic。
func eehVxKUB(server any, stream grpc.ServerStream) error {
	return server.(SCv1eHEMZ).Tun(&utpvEgZY{Y5md9Ywqjd4m: stream})
}

// Send @ 0xe89bc0。
func (s *utpvEgZY) Send(message *Z27tYEZ9hy1o) error { return s.SendMsg(message) }

// Recv @ 0xe89c20。读取失败丢弃已分配消息，原样返回错误。
func (s *utpvEgZY) Recv() (*Z27tYEZ9hy1o, error) {
	message := new(Z27tYEZ9hy1o)
	if err := s.RecvMsg(message); err != nil {
		return nil, err
	}
	return message, nil
}

// 静态 0x5057580：名称由 init 解码，流表位于 0x503e2a0。
// Methods 是原静态非 nil 空切片，双向流两个标记均为 true。
var serviceDescriptionAt5057580 = grpc.ServiceDesc{
	ServiceName: "protos.GunService",
	HandlerType: (*SCv1eHEMZ)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{{StreamName: "Tun", Handler: eehVxKUB, ServerStreams: true, ClientStreams: true}},
	Metadata:    "protos/grpc.proto",
}

// E9qSZvo37 @ 0xe87880。调用方提供服务名；metadata 原字面量为 gun.proto。
// 每次创建独立流表，不能复用可修改的全局 ServiceDesc。
func E9qSZvo37(serviceName string) grpc.ServiceDesc {
	return grpc.ServiceDesc{ServiceName: serviceName, HandlerType: (*SCv1eHEMZ)(nil), Methods: []grpc.MethodDesc{}, Streams: []grpc.StreamDesc{{StreamName: "Tun", Handler: eehVxKUB, ServerStreams: true, ClientStreams: true}}, Metadata: "gun.proto"}
}

// Qpv51XFhNT @ 0xe87a00。
func Qpv51XFhNT(server *grpc.Server, implementation SCv1eHEMZ, serviceName string) {
	description := E9qSZvo37(serviceName)
	server.RegisterService(&description, implementation)
}
