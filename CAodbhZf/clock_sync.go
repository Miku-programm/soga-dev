package CAodbhZf

import (
	"context"
	"errors"
	"net"
	"sync/atomic"
	"time"

	"github.com/beevik/ntp"
	"recovered.local/soga/muc279Mpa6S4"
)

// 原 init@0xb41a00 按此顺序建立七个服务地址。
var ntpServersAt5099630 = []string{
	"ntp.aliyun.com", "pool.ntp.org", "asia.pool.ntp.org", "time.asia.apple.com",
	"time.windows.com", "cn.ntp.org.cn", "edu.ntp.org.cn",
}

var timeOffsetAt50C7170 int64
var timeLogAt50C6C6A bool

// aep2YURd_Nd 对应 0xb48040。每个域名最多尝试一个 IPv4 和一个 IPv6 地址。
// 成功的 NTP 响应在调用者中验证；验证失败不会在这里换服务器。
func aep2YURd_Nd() (*ntp.Response, error) {
	for _, server := range ntpServersAt5099630 {
		addresses, err := net.LookupIP(server)
		if err != nil {
			continue
		}
		tried4, tried6 := false, false
		for _, address := range addresses {
			ip4 := address.To4()
			if !tried4 && ip4 != nil {
				response, err := ntp.QueryWithOptions(ip4.String(), ntp.QueryOptions{})
				if err == nil {
					return response, nil
				}
				tried4 = true
			}
			if !tried6 && ip4 == nil {
				response, err := ntp.QueryWithOptions(address.String(), ntp.QueryOptions{})
				if err == nil {
					return response, nil
				}
				tried6 = true
			}
		}
	}
	return nil, errors.New("udp issue")
}

// ejBA9tGpop 对应 0xb48380；网络或响应校验失败保留此前的偏移。
func ejBA9tGpop() {
	response, err := aep2YURd_Nd()
	if err != nil {
		if timeLogAt50C6C6A {
			muc279Mpa6S4.TRTbrquD(0, "sync time failed:", err)
		}
		return
	}
	if err := response.Validate(); err != nil {
		if timeLogAt50C6C6A {
			muc279Mpa6S4.TRTbrquD(0, "sync time validate failed:", err)
		}
		return
	}
	atomic.StoreInt64(&timeOffsetAt50C7170, int64(response.ClockOffset))
	if timeLogAt50C6C6A {
		muc279Mpa6S4.TRTbrquD(0, "time offset:", response.ClockOffset)
	}
}

// Kia4I9wKU 对应 0xb485c0；先同步一次，再每小时同步，取消时停止计时器。
func Kia4I9wKU(ctx context.Context, printLog bool) {
	timeLogAt50C6C6A = printLog
	ejBA9tGpop()
	go func() {
		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ejBA9tGpop()
			case <-ctx.Done():
				return
			}
		}
	}()
}

// Nj5UR_g2、C5Jff5vs0 分别对应 0xb486e0、0xb48680。
func Nj5UR_g2() time.Time {
	return time.Now().Add(time.Duration(atomic.LoadInt64(&timeOffsetAt50C7170)))
}
func C5Jff5vs0() int64 { return Nj5UR_g2().Unix() }
