// 依据原 ELF xIF4G3 包机器码及 RTTI 恢复的源码表示，不是作者文件原文。
// 对照基线：apernet/hysteria 9e90d7d，保留其原许可；证据在 hysteria-wire。
package xIF4G3

import (
	"math"
	"time"

	"github.com/apernet/quic-go/congestion"
)

const (
	infBandwidth = JaadYOR_Vaf6(math.MaxUint64)
)

// Bandwidth of a connection
type JaadYOR_Vaf6 uint64

const (
	// BitsPerSecond is 1 bit per second
	BitsPerSecond JaadYOR_Vaf6 = 1
	// BytesPerSecond is 1 byte per second
	BytesPerSecond = 8 * BitsPerSecond
)

// BandwidthFromDelta calculates the bandwidth from a number of bytes and a time delta
func BandwidthFromDelta(bytes congestion.ByteCount, delta time.Duration) JaadYOR_Vaf6 {
	return JaadYOR_Vaf6(bytes) * JaadYOR_Vaf6(time.Second) / JaadYOR_Vaf6(delta) * BytesPerSecond
}
