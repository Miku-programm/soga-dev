//go:build linux

package CbO_7wp3B

import (
	"context"
	"sync"

	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/a3PdV_bmpRUK"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/q5bqCgTdxaTM"
)

// _Zpu8_d 为原 RTTI 0x359ece8 的 func(int)。
type _Zpu8_d func(int)

var starterLockAt50C6FA8 sync.Mutex
var starterOnceAt50C7AA0 sync.Once

// startOriginalInbound 是原 WrdrARwuhuha.vDpj7g@0x1fa1880 的源码提取。
// construct 参数对应原 MFFa9yT.Pb5l5RLOHSi 调用；不得接入替代协议实现。
// 原全局锁仅覆盖 Once，不覆盖 Start；回调以首次用户数在单独 goroutine 执行。
func startOriginalInbound(ctx context.Context, data *q5bqCgTdxaTM.D5gTIZ5aVg, callback _Zpu8_d,
	construct func(*MFFa9yT.DV_XPJ4z0) (MFFa9yT.Kz_8ajU7, error),
) (MFFa9yT.Kz_8ajU7, error) {
	return startInboundWithConstructor(&WrdrARwuhuha{zqoWpLRIH: callback}, ctx, data, construct)
}

// 生产入口与测试的边界夹具共用原控制流；接收者的回调字段只在 Once
// 真正执行时读取，不能在取消分支或构造错误分支前提前解引用接收者。
func startInboundWithConstructor(receiver *WrdrARwuhuha, ctx context.Context, data *q5bqCgTdxaTM.D5gTIZ5aVg,
	construct func(*MFFa9yT.DV_XPJ4z0) (MFFa9yT.Kz_8ajU7, error),
) (MFFa9yT.Kz_8ajU7, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	node := data.H47AzZsYHB1N
	if gcTStJ6_zEzn() && node.I3k4JnNp7nB != "" && node.XATiVun != "" {
		if _, _, err := a3PdV_bmpRUK.Qr4n7sWM5di(node.EMlXEc, node.I3k4JnNp7nB, node.XATiVun); err != nil {
			muc279Mpa6S4.TRTbrquD(4, "write panel cert failed:", err)
		}
	}
	inbound, err := construct(&MFFa9yT.DV_XPJ4z0{
		Dmpl48sXM: ctx, POpyjFWF8v: cNT24rG4F.Mpn2JltM(cNT24rG4F.DnhHoT2aRvs()), HG5k0365: data,
	})
	if err != nil {
		return nil, err
	}
	starterLockAt50C6FA8.Lock()
	starterOnceAt50C7AA0.Do(func() {
		go receiver.zqoWpLRIH(len(data.VTme1SuHL6i_))
	})
	starterLockAt50C6FA8.Unlock()
	if err := inbound.Start(); err != nil {
		inbound.Stop()
		return nil, err
	}
	return inbound, nil
}
