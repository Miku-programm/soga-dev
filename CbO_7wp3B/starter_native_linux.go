//go:build linux

package CbO_7wp3B

import (
	"context"
	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/q5bqCgTdxaTM"
)

// WrdrARwuhuha 保留原 RTTI 的八字节结构和原回调函数类型。
// 此处不保存协议工厂替代项；vDpj7g 总是调用原 Pb5l5RLOHSi。
type WrdrARwuhuha struct{ zqoWpLRIH _Zpu8_d }

// vDpj7g @0x1fa1880。传递接收者本身，保持原延迟字段读取和 nil 行为。
func (s *WrdrARwuhuha) vDpj7g(ctx context.Context, data *q5bqCgTdxaTM.D5gTIZ5aVg) (MFFa9yT.Kz_8ajU7, error) {
	return startInboundWithConstructor(s, ctx, data, MFFa9yT.Pb5l5RLOHSi)
}

// avZC_rQmNC @0x1fa5a20。方法值仅捕获指针，不提前读取接收者字段。
func (s *WrdrARwuhuha) avZC_rQmNC(ctx context.Context, panel q5bqCgTdxaTM.VyFk5uT) {
	runOriginalStarter(ctx, panel, s.vDpj7g)
}

// Start @0x1fa7720、gowrap @0x1ffd280：异步发起，不等待拉取配置或启动。
func (s *WrdrARwuhuha) Start(ctx context.Context, panel q5bqCgTdxaTM.VyFk5uT) {
	go s.avZC_rQmNC(ctx, panel)
}

// G6miUObWbEE 对应原 starter interface，只有 Start(context.Context, panel)。
type G6miUObWbEE interface {
	Start(context.Context, q5bqCgTdxaTM.VyFk5uT)
}
type bRy8oqwuq6cY func(_Zpu8_d) G6miUObWbEE

// eXC63toqu @0x1f98c40；原构造只分配八字节对象并存回调，无启动副作用。
func eXC63toqu(callback _Zpu8_d) G6miUObWbEE { return &WrdrARwuhuha{zqoWpLRIH: callback} }
