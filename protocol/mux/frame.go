package mux

import (
	"fmt"
	"io"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/buf/bufiface"
)

// 类型名称及字段来自原 n7_GPlUYmWO RTTI；方法包实际为 nXZ99wahk。
type FEGjbz2y uint8
type OaPaSragw1 uint8
type CUwik6VP2WU0 struct {
	ljtkJw       uint16
	xvTw9VQDQdY  FEGjbz2y
	m3vf0n8      uint8
	ePmcTrBC     OaPaSragw1
	hUHpeGfV     *relay.AHU2MRmjzQu
	gQBBY59xyosW bool
	c8AtRj7hb    bufiface.Mf3phSOVqgf
}

// String 对应 0x1cebb60；KeepAlive 由原解码器还原。
func (c FEGjbz2y) String() string {
	switch c {
	case 1:
		return "New"
	case 2:
		return "Keep"
	case 3:
		return "End"
	case 4:
		return "KeepAlive"
	default:
		return "Unknown"
	}
}
func (f *CUwik6VP2WU0) HasData() bool  { return f.m3vf0n8&1 != 0 } //0x1cec6c0
func (f *CUwik6VP2WU0) HasError() bool { return f.m3vf0n8&2 != 0 } //0x1cec6e0
// FreeData 对应 0x1cec700；只释放存在的缓冲，并清空原接口。
func (f *CUwik6VP2WU0) FreeData() {
	if f.c8AtRj7hb != nil {
		f.c8AtRj7hb.Free()
		f.c8AtRj7hb = nil
	}
}

// HlVrAnC111 对应 0x1cebc20；完整读取声明的元数据（4..512 字节），
// New 必须包含目的地址，Keep 仅在元数据长度>4 时包含地址，其余额外字节忽略。
func HlVrAnC111(r io.Reader) (*CUwik6VP2WU0, error) {
	n, err := relay.HoymNl_VBq(r)
	if err != nil {
		return nil, err
	}
	if n < 4 {
		return nil, fmt.Errorf("mux metadata size < 4: %v", int(n))
	}
	if n > 512 {
		return nil, fmt.Errorf("mux metadata size is too big: %v", int(n))
	}
	metadata := pool.S4OQXWIZ(int(n))
	defer metadata.Free()
	if _, err := metadata.ReadFullFrom(r, int(n)); err != nil {
		return nil, err
	}
	id, err := metadata.ReadUint16()
	if err != nil {
		return nil, err
	}
	cmd, err := metadata.ReadUint8()
	if err != nil {
		return nil, err
	}
	flags, err := metadata.ReadUint8()
	if err != nil {
		return nil, err
	}
	frame := &CUwik6VP2WU0{ljtkJw: id, xvTw9VQDQdY: FEGjbz2y(cmd), m3vf0n8: flags}
	switch cmd {
	case 1, 2:
		if cmd == 1 || n > 4 {
			network, err := metadata.ReadUint8()
			if err != nil {
				return nil, err
			}
			address, err := relay.CHYbTjleKL(metadata)
			if err != nil {
				return nil, err
			}
			frame.ePmcTrBC = OaPaSragw1(network)
			frame.hUHpeGfV = address
			frame.gQBBY59xyosW = cmd == 2
		}
	case 3, 4:
	default:
		return nil, fmt.Errorf("unknown mux cmd: %v", FEGjbz2y(cmd))
	}
	if frame.HasData() {
		data, err := frame.xq8JKB3(r)
		if err != nil {
			return nil, err
		}
		frame.c8AtRj7hb = data
	}
	return frame, nil
}

// xq8JKB3 对应 0x1cec780；先在同一个原缓冲中读两字节长度，消费后再读数据，
// 出错时归还缓冲，成功时转移所有权给调用者。零长度仍返回非 nil 的空缓冲。
func (f *CUwik6VP2WU0) xq8JKB3(r io.Reader) (bufiface.Mf3phSOVqgf, error) {
	b := pool.S4OQXWIZ(2)
	if _, err := b.ReadFullFrom(r, 2); err != nil {
		b.Free()
		return nil, err
	}
	n, err := b.ReadUint16()
	if err != nil {
		b.Free()
		return nil, err
	}
	if _, err := b.ReadFullFrom(r, int(n)); err != nil {
		b.Free()
		return nil, err
	}
	return b, nil
}
