package relay

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
	"math"
	"math/rand"
	"net"
	"os"
	D "recovered.local/soga/buf/pool"
	"sync"
	"time"
)

var hostnameOnceAt50C7B60 sync.Once
var hostnameNumberAt50C6D50 uint32

func gSnxlJwdTs() uint32 { // 0x15a3de0；初始化闭包 0x15b5a60。
	hostnameOnceAt50C7B60.Do(func() {
		name, _ := os.Hostname()
		digest := sha256.Sum256([]byte(name))
		digest[0] &= 0x7f
		hostnameNumberAt50C6D50 = binary.BigEndian.Uint32(digest[:4])
	})
	return hostnameNumberAt50C6D50
}
func i3zxS3(low, high int) int { // 0x15a3e20
	if high <= low {
		return low
	}
	n := float64(rand.Intn(high - low + 1))
	f := rand.Float64()
	return int(n*math.Sqrt(f*f*f)) + low
}

// Ad2VIxY @0x15a3ec0：失败认证后按原主机摘要和分布设置读超时并消费数据。
// 既不关闭连接，也不向调用方返回读错误；关闭由原上层 defer 执行。
func Ad2VIxY(conn net.Conn) {
	delay := i3zxS3(2000, 10000) + int(gSnxlJwdTs()%20000)
	conn.SetReadDeadline(time.Now().Add(time.Duration(delay) * time.Millisecond))
	size := 1 << ((gSnxlJwdTs() & 3) + 12)
	buf := D.PoolAt50961D8.Get(int(size))
	defer D.BE1pnyFd(buf)
	minimum := i3zxS3(0, int(size)-254) + int(byte(gSnxlJwdTs()))
	if minimum > 0 {
		io.ReadAtLeast(conn, buf, minimum)
	}
}
