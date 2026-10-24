// 原样本 EPZGnxoa5J（RTTI 显示 siaEQHRJ）拥塞发送节拍器的源码表示。
// 机器码地址 0x1c0b860、0x1c0b940、0x1c0bae0、0x1c0bba0。
// 对照 Hysteria 9e90d7d 的 common/pacer.go 核对公式、调用顺序和溢出行为。
package hysteriaerr

import (
	"github.com/apernet/quic-go/congestion"
	"time"
)

type OQpLDX struct {
	gpqkhK     congestion.ByteCount
	jyPhMc     congestion.ByteCount
	rh7eeI     time.Time
	clqvXsSOvk func() congestion.ByteCount
}

// NewFromInline 是恢复工程桥接名；原构造被内联在 BBR/Brutal 构造中，
// 初始化值和字段赋值对应 0x1c0e080、0x1c1dfa0，不计独立原函数。
func NewFromInline(bandwidth func() congestion.ByteCount) *OQpLDX {
	return &OQpLDX{gpqkhK: 12800, jyPhMc: 1280, clqvXsSOvk: bandwidth}
}
func (p *OQpLDX) SentPacket(sendTime time.Time, size congestion.ByteCount) {
	budget := p.Budget(sendTime)
	if size > budget {
		p.gpqkhK = 0
	} else {
		p.gpqkhK = budget - size
	}
	p.rh7eeI = sendTime
}
func (p *OQpLDX) Budget(now time.Time) congestion.ByteCount {
	if p.rh7eeI.IsZero() {
		return p.maxBurstSizeFromInline()
	}
	budget := p.gpqkhK + p.clqvXsSOvk()*congestion.ByteCount(now.Sub(p.rh7eeI).Nanoseconds())/1e9
	if budget < 0 {
		budget = (1 << 62) - 1
	}
	return min(p.maxBurstSizeFromInline(), budget)
}

// 对应四处已内联的 max(4ms*带宽/1s,10*包大小)；保留每次真实带宽函数调用。
func (p *OQpLDX) maxBurstSizeFromInline() congestion.ByteCount {
	return max(4000000*p.clqvXsSOvk()/1000000000, 10*p.jyPhMc)
}
func (p *OQpLDX) TimeUntilSend() time.Time {
	if p.gpqkhK >= p.jyPhMc {
		return time.Time{}
	}
	diff := 1000000000 * uint64(p.jyPhMc-p.gpqkhK)
	bw := uint64(p.clqvXsSOvk())
	d := diff / bw
	if diff%bw > 0 {
		d++
	}
	return p.rh7eeI.Add(max(time.Millisecond, time.Duration(d)))
}
func (p *OQpLDX) SetMaxDatagramSize(size congestion.ByteCount) { p.jyPhMc = size }
