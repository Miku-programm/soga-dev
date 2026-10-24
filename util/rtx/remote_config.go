package rtx

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"recovered.local/soga/util/logx"
)

// 原指针全局 0x5029248 指向 0x50449f0 的 http.Client；前五个机器字
// 为零，Timeout 为 60000000000。原 init 不改写该客户端。
var remoteClientAt5029248 = &http.Client{Timeout: time.Minute}

// oDPndU2e 对应 0xb4d160。不检查 HTTP 状态码；任何响应体都交由调用方解析。
func oDPndU2e(url string) ([]byte, error) {
	response, err := remoteClientAt5029248.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

// KJa7rJUaLPYe 对应 0xb4d340：立即取一次，之后每分钟重取。
// ctx 只控制两次请求之间的等待；原 HTTP 请求没有绑定该 context。
func KJa7rJUaLPYe(ctx context.Context, url string, load func([]byte) error) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	var previous []byte
	for {
		data, err := oDPndU2e(url)
		if err != nil {
			logx.V8GN5Stk(8, "load url <%v> failed: %v", url, err)
		} else if !bytes.Equal(previous, data) {
			if err := load(data); err != nil {
				logx.V8GN5Stk(8, "load url <%v> failed: %v", url, err)
			} else {
				previous = data
			}
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return
		}
	}
}
