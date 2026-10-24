// 原样本协议函数的 Go 源码表示，不是作者源码原文。
package vlessvision

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/transport/grpcproto"
)

// KYlGQuja 是原 RTTI 0x35849c0 的命名 uint8 类型。
type KYlGQuja uint8

// ZfmpaHtmu 对应 RTTI 0x36b68c0，24 字节，字段偏移 0/1/8/16。
type ZfmpaHtmu struct {
	h8FCLanKLDuj byte
	aj8eMgsHQRY  KYlGQuja
	if2alzq2rt6  *relay.AHU2MRmjzQu
	qssMeIUTR    *grpcproto.N_SQ5a
}

// a8D0ppAPaqu 对应 0x1d07580。长度为零仍分配并返回非 nil Addons。
func a8D0ppAPaqu(conn net.Conn) (*grpcproto.N_SQ5a, error) {
	size, err := relay.TBUcsYQLu25q(conn)
	if err != nil {
		return nil, err
	}
	addons := new(grpcproto.N_SQ5a)
	if size != 0 {
		data := make([]byte, int(size))
		if _, err := io.ReadFull(conn, data); err != nil {
			return nil, err
		}
		if err := proto.Unmarshal(data, addons); err != nil {
			return nil, err
		}
	}
	return addons, nil
}

// vF01yD4y 对应 0x1d071e0。版本字节由调用方传入；本函数从附加信息长度开始读。
// 命令 3 不读目标地址，其他命令均按原地址格式读取，不额外拒绝未知命令。
func vF01yD4y(conn net.Conn, version byte) (*ZfmpaHtmu, error) {
	if version != 0 {
		return nil, fmt.Errorf("unknown vless version: %v", version)
	}
	addons, err := a8D0ppAPaqu(conn)
	if err != nil {
		return nil, err
	}
	command, err := relay.TBUcsYQLu25q(conn)
	if err != nil {
		return nil, err
	}
	var destination *relay.AHU2MRmjzQu
	if command != 3 {
		destination, err = relay.CHYbTjleKL(conn)
		if err != nil {
			return nil, err
		}
	}
	return &ZfmpaHtmu{aj8eMgsHQRY: KYlGQuja(command), if2alzq2rt6: destination, qssMeIUTR: addons}, nil
}
func (*ZfmpaHtmu) ResponseHeaderLen() int { return 2 } // 0x1d074a0

// WriteResponseHeader @0x1d074c0 内联两次原 WriteByte；响应附加长度固定为零。
func (h *ZfmpaHtmu) WriteResponseHeader(b *pool.X_aXX5) {
	b.WriteByte(h.h8FCLanKLDuj)
	b.WriteByte(0)
}
