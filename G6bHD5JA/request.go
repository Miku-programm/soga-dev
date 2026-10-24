// 根据原样本 0x15a49e0 的指令及字符串解码结果恢复，不是作者源码原文。
package G6bHD5JA

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// qword_5097990 是原 init@0x159db20 写入的字符串全局变量。
// 变量原名未保留；此地址名不是推定的作者名称。值由 0x15a5e40 解出。
var qword_5097990 = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36"

// dfK5RMSe4h9B 对应 0x15a49e0。非 nil 查询参数完全替换已有查询字符串。
func dfK5RMSe4h9B(address, method string, query url.Values, body io.Reader, headers http.Header) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, address, body)
	if err != nil {
		return nil, err
	}
	// 键和值分别由 0x15ac340、0x15ac8e0、0x15acf40 解出。
	req.Header.Set("User-Agent", qword_5097990)
	req.Header.Set("Cache-Control", "no-cache")
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return req, nil
}
