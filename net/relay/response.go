// 根据原样本 HTTP 封装的指令恢复；保留 metadata 函数名，不是作者源码原文。
package relay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// vxLBn4qabL9L 对应 0x15a4f40；初始化发生在构造请求之前。
func vxLBn4qabL9L(address, method string, query url.Values, body io.Reader, headers http.Header, contentType string) (*http.Response, error) {
	clientOnceAt50C7B70.Do(WuvamjL5)
	req, err := dfK5RMSe4h9B(address, method, query, body, headers)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return clientAt5029428.Do(req)
}

// IlWBbp1xdPw 对应 0x15a51c0；读响应失败时丢弃部分内容，保留状态和头。
// Close 错误按原 deferwrap1 行为忽略，HTTP 非 2xx 状态本身不产生错误。
func IlWBbp1xdPw(address string, query url.Values, headers http.Header) (int, []byte, http.Header, error) {
	resp, err := vxLBn4qabL9L(address, "GET", query, nil, headers, "")
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, resp.Header, err
	}
	return resp.StatusCode, data, resp.Header, nil
}

// ZBovtj1xky 对应 0x15a57a0。原 json/v2.Marshal 使用 0x505b4e0 的
// DefaultOptionsV1，等价于此处 encoding/json.Marshal；不是默认 v2 编码行为。
func ZBovtj1xky(address string, query url.Values, value any, headers http.Header) ([]byte, int, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, 0, err
	}
	resp, err := vxLBn4qabL9L(address, "POST", query, bytes.NewBuffer(data), headers, "application/json")
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return data, resp.StatusCode, nil
}

// J1liFniHg 对应 0x15a5460，表单放请求体，URL 原查询字符串保留。
// Content-Type 由该调用点的字符串 decoder 解出。
func J1liFniHg(address string, values url.Values) ([]byte, int, error) {
	resp, err := vxLBn4qabL9L(address, "POST", nil, strings.NewReader(values.Encode()), nil, "application/x-www-form-urlencoded")
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return data, resp.StatusCode, nil
}
