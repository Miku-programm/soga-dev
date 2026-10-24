// 根据原样本面板请求及缓存分支恢复，不是作者源码原文。
package xiaov2board

import (
	"fmt"
	"net/url"
	"recovered.local/soga/net/relay"
)

// GetCommon 对应 0x17a7100；缓存键是调用方原始 address，不含新增查询参数。
func (p *LRlqOaYmO) GetCommon(address string, extra url.Values) ([]byte, error) {
	query := p.fUwRFGkp()
	for key, values := range extra {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	status, data, headers, err := relay.IlWBbp1xdPw(address, query, p.iU3FqWNT(address))
	// 错误格式由 0x17a7601 解码；原格式是 %v，不提供错误解包。
	if err != nil {
		return nil, fmt.Errorf("%v, result: %v, status code: %v", err, string(data), status)
	}
	switch status {
	case 200:
		if etag := headers.Get("etag"); etag != "" {
			if cached, ok := p.eyq32s.CheckETag(address, etag); ok {
				return cached.([]byte), nil
			}
			p.eyq32s.AddETag(address, etag, data)
		}
		return data, nil
	case 304:
		// 原代码内联了 GetObj；保留 []byte 断言及 panic 行为。
		if cached, ok := p.eyq32s.GetObj(address); ok {
			return cached.([]byte), nil
		}
		return nil, fmt.Errorf("can not find eTag: %v", address)
	default:
		return nil, fmt.Errorf("request failed, result: %v, status code: %v", string(data), status)
	}
}

// PostCommon 对应 0x17a6d80，只有 HTTP 200 且读取无错才成功。
func (p *LRlqOaYmO) PostCommon(address string, value any) error {
	data, status, err := relay.ZBovtj1xky(address, p.fUwRFGkp(), value, nil)
	if err != nil {
		return fmt.Errorf("post %v err: %v, status code: %v, result: %v", address, err, status, string(data))
	}
	if status != 200 {
		return fmt.Errorf("post status code: %v, result: %v", status, string(data))
	}
	return nil
}
