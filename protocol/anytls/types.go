// Package vtQf23A9G 依据原 ELF 的 AnyTLS 会话机器码和 RTTI 恢复。
// 这不是作者文件原文。RTTI 显示包名 _nhcdlvtIZN1，方法元数据包路径为 vtQf23A9G。
package anytls

import (
	"net"
	B "recovered.local/soga/util/deadline"
	C "recovered.local/soga/util/rtx"
	H "recovered.local/soga/transport/hysteriapad"
	"sync"
	"sync/atomic"
	"time"
)

type cNZriK [7]byte

type BIkgBMz3M struct {
	c740FZf      net.Conn
	a5L1EeV4yvaw sync.Mutex
	mW8PvJYxYup9 map[uint32]*CCXLCurYiT7
	iX5FOl       atomic.Uint32
	ukg_2xDnH9s  sync.RWMutex
	pGT2EZ20     sync.Once
	sF3ka3A2dYX  chan struct{}
	twJTEeJ0S    func()
	c5NUXMoBI    func()
	wAiQ9IC2UZ   sync.Mutex
	xZ8dRw2_KR   uint64
	re4cr8xRiVdH time.Time
	bmjW7a       *C.EahpvU_iseQ[*H.FZBbzkJPWbrk]
	gAJaebL      uint8
	pNuGhHxIQdo  bool
	rVguB66      bool
	d0Ze8qja     bool
	vV_K36       []byte
	wLYsfs       atomic.Uint32
	ulr87QdD     func(*CCXLCurYiT7)
}

type CCXLCurYiT7 struct {
	y7z3pbAl     uint32
	hnbi8qavj6Y  *BIkgBMz3M
	k_6JEbXwYJY  *B.P848NOz3su9
	z8qCZQcE     *B.EjFsCHfka9Zc
	ko_ou22dHrhf B.TyIy_s
	yTV1rdHC     sync.Once
	kCaaEFymiBDr func()
	yKgxtGJV     error
	bo7peua2l    sync.Once
}

// NewServerFromInline 是集成辅助名，对应 inbound.R_n5ucb 的连接处理器
// @0x1d34ee0 中 0x1d355.. 区域内联初始化；未增加原版没有的默认行为。
func NewServerFromInline(conn net.Conn, callback func(*CCXLCurYiT7), padding *C.EahpvU_iseQ[*H.FZBbzkJPWbrk]) *BIkgBMz3M {
	return &BIkgBMz3M{c740FZf: conn, ulr87QdD: callback, bmjW7a: padding, sF3ka3A2dYX: make(chan struct{}), mW8PvJYxYup9: make(map[uint32]*CCXLCurYiT7)}
}

// newStreamFromInline 合并 OpenStream 与读循环中相同的内联字段赋值。
func newStreamFromInline(s *BIkgBMz3M, id uint32) *CCXLCurYiT7 {
	r, w := B.NewPipe()
	return &CCXLCurYiT7{y7z3pbAl: id, hnbi8qavj6Y: s, k_6JEbXwYJY: r, z8qCZQcE: w, ko_ou22dHrhf: B.NewDeadline()}
}
