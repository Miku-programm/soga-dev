// 按原 BOaPYrtdvWIg 机器码、RTTI、字符串解码恢复；
// 候选基线 Hysteria 9e90d7d 已逐方法比较。保留其授权说明：
// Copyright 2023 Toby
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package BOaPYrtdvWIg

import (
	"fmt"
	"os"
	"strconv"
	"time"

	common "recovered.local/soga/EPZGnxoa5J"

	"github.com/apernet/quic-go/congestion"
)

const (
	pktInfoSlotCount           = 5 // slot index is based on seconds, so this is basically how many seconds we sample
	minSampleCount             = 50
	minAckRate                 = 0.8
	congestionWindowMultiplier = 2

	debugEnv           = "HYSTERIA_BRUTAL_DEBUG"
	debugPrintInterval = 2
)

var _ congestion.CongestionControl = &ON7Lun{}

type ON7Lun struct {
	vbyI7ILUmS1 congestion.RTTStatsProvider
	_aNjasGyoX  congestion.ByteCount
	t9AYEa6hD   congestion.ByteCount
	iVGkY7      *common.OQpLDX

	bzbrOl       [pktInfoSlotCount]gHgCJl
	qB0G_BBq1s6u float64

	hxjEhjlUXU   bool
	mP5_CmSWJZaP int64
}

type gHgCJl struct {
	I_nxVko6a    int64
	EYG4t7Mh_5zN uint64
	ACD0yUo6K56C uint64
}

func GPauZ708C6(_aNjasGyoX uint64) *ON7Lun {
	hxjEhjlUXU, _ := strconv.ParseBool(os.Getenv(debugEnv))
	bs := &ON7Lun{
		_aNjasGyoX:   congestion.ByteCount(_aNjasGyoX),
		t9AYEa6hD:    congestion.InitialPacketSizeIPv4,
		qB0G_BBq1s6u: 1,
		hxjEhjlUXU:   hxjEhjlUXU,
	}
	bs.iVGkY7 = common.NewFromInline(func() congestion.ByteCount {
		return congestion.ByteCount(float64(bs._aNjasGyoX) / bs.qB0G_BBq1s6u)
	})
	return bs
}

func (b *ON7Lun) SetRTTStatsProvider(vbyI7ILUmS1 congestion.RTTStatsProvider) {
	b.vbyI7ILUmS1 = vbyI7ILUmS1
}

func (b *ON7Lun) TimeUntilSend(bytesInFlight congestion.ByteCount) time.Time {
	return b.iVGkY7.TimeUntilSend()
}

func (b *ON7Lun) HasPacingBudget(now time.Time) bool {
	return b.iVGkY7.Budget(now) >= b.t9AYEa6hD
}

func (b *ON7Lun) CanSend(bytesInFlight congestion.ByteCount) bool {
	return bytesInFlight <= b.GetCongestionWindow()
}

func (b *ON7Lun) GetCongestionWindow() congestion.ByteCount {
	rtt := b.vbyI7ILUmS1.SmoothedRTT()
	if rtt <= 0 {
		return 10240
	}
	cwnd := congestion.ByteCount(float64(b._aNjasGyoX) * rtt.Seconds() * congestionWindowMultiplier / b.qB0G_BBq1s6u)
	if cwnd < b.t9AYEa6hD {
		cwnd = b.t9AYEa6hD
	}
	return cwnd
}

func (b *ON7Lun) OnPacketSent(sentTime time.Time, bytesInFlight congestion.ByteCount,
	packetNumber congestion.PacketNumber, bytes congestion.ByteCount, isRetransmittable bool,
) {
	b.iVGkY7.SentPacket(sentTime, bytes)
}

func (b *ON7Lun) OnPacketAcked(number congestion.PacketNumber, ackedBytes congestion.ByteCount,
	priorInFlight congestion.ByteCount, eventTime time.Time,
) {
	// 原入口仅 RET；保留样本无操作行为。
}

func (b *ON7Lun) OnCongestionEvent(number congestion.PacketNumber, lostBytes congestion.ByteCount,
	priorInFlight congestion.ByteCount,
) {
	// 原入口仅 RET；保留样本无操作行为。
}

func (b *ON7Lun) OnCongestionEventEx(priorInFlight congestion.ByteCount, eventTime time.Time, ackedPackets []congestion.AckedPacketInfo, lostPackets []congestion.LostPacketInfo) {
	currentTimestamp := eventTime.Unix()
	slot := currentTimestamp % pktInfoSlotCount
	if b.bzbrOl[slot].I_nxVko6a == currentTimestamp {
		b.bzbrOl[slot].ACD0yUo6K56C += uint64(len(lostPackets))
		b.bzbrOl[slot].EYG4t7Mh_5zN += uint64(len(ackedPackets))
	} else {
		// uninitialized slot or too old, reset
		b.bzbrOl[slot].I_nxVko6a = currentTimestamp
		b.bzbrOl[slot].EYG4t7Mh_5zN = uint64(len(ackedPackets))
		b.bzbrOl[slot].ACD0yUo6K56C = uint64(len(lostPackets))
	}
	b.gDjSGs_OvW3(currentTimestamp)
}

func (b *ON7Lun) SetMaxDatagramSize(size congestion.ByteCount) {
	b.t9AYEa6hD = size
	b.iVGkY7.SetMaxDatagramSize(size)
	if b.hxjEhjlUXU {
		b.pRMpAt85lTaU("SetMaxDatagramSize: %d", size)
	}
}

func (b *ON7Lun) gDjSGs_OvW3(currentTimestamp int64) {
	minTimestamp := currentTimestamp - pktInfoSlotCount
	var ackCount, lossCount uint64
	for _, info := range b.bzbrOl {
		if info.I_nxVko6a < minTimestamp {
			continue
		}
		ackCount += info.EYG4t7Mh_5zN
		lossCount += info.ACD0yUo6K56C
	}
	if ackCount+lossCount < minSampleCount {
		b.qB0G_BBq1s6u = 1
		if b.canPrintAckRateFromInline(currentTimestamp) {
			b.mP5_CmSWJZaP = currentTimestamp
			b.pRMpAt85lTaU("Not enough samples (total=%d, ack=%d, loss=%d, rtt=%d)",
				ackCount+lossCount, ackCount, lossCount, b.vbyI7ILUmS1.SmoothedRTT().Milliseconds())
		}
		return
	}
	rate := float64(ackCount) / float64(ackCount+lossCount)
	if rate < minAckRate {
		b.qB0G_BBq1s6u = minAckRate
		if b.canPrintAckRateFromInline(currentTimestamp) {
			b.mP5_CmSWJZaP = currentTimestamp
			b.pRMpAt85lTaU("ACK rate too low: %.2f, clamped to %.2f (total=%d, ack=%d, loss=%d, rtt=%d)",
				rate, minAckRate, ackCount+lossCount, ackCount, lossCount, b.vbyI7ILUmS1.SmoothedRTT().Milliseconds())
		}
		return
	}
	b.qB0G_BBq1s6u = rate
	if b.canPrintAckRateFromInline(currentTimestamp) {
		b.mP5_CmSWJZaP = currentTimestamp
		b.pRMpAt85lTaU("ACK rate: %.2f (total=%d, ack=%d, loss=%d, rtt=%d)",
			rate, ackCount+lossCount, ackCount, lossCount, b.vbyI7ILUmS1.SmoothedRTT().Milliseconds())
	}
}

func (b *ON7Lun) InSlowStart() bool {
	return false
}

func (b *ON7Lun) InRecovery() bool {
	return false
}

func (b *ON7Lun) MaybeExitSlowStart() {}

func (b *ON7Lun) OnRetransmissionTimeout(packetsRetransmitted bool) {}

func (b *ON7Lun) canPrintAckRateFromInline(currentTimestamp int64) bool {
	return b.hxjEhjlUXU && currentTimestamp-b.mP5_CmSWJZaP >= debugPrintInterval
}

func (b *ON7Lun) pRMpAt85lTaU(format string, a ...any) {
	fmt.Printf("[BrutalSender] [%s] %s\n",
		time.Now().Format("15:04:05"),
		fmt.Sprintf(format, a...))
}
