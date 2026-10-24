// 根据原样本面板请求及缓存分支恢复，不是作者源码原文。
package BP3PTlx

import (
	"fmt"
	"net/url"
	"recovered.local/soga/G6bHD5JA"
)

// lVcAo38Q87cu 对应 0x1791be0；缓存键是调用方原始 address，不含新增查询参数。
func (p *KlO65BuwOt) lVcAo38Q87cu(address string, extra url.Values) ([]byte, error) {
	query := p.tgg_NuFN()
	for key, values := range extra {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	status, data, headers, err := G6bHD5JA.IlWBbp1xdPw(address, query, p.x7aWIuw06(address))
	// 错误格式由 0x17920c6 解码；原格式是 %v，不提供错误解包。
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

// iJrldZ0g 对应 0x1791860，只有 HTTP 200 且读取无错才成功。
func (p *KlO65BuwOt) iJrldZ0g(address string, value any) error {
	data, status, err := G6bHD5JA.ZBovtj1xky(address, p.tgg_NuFN(), value, nil)
	if err != nil {
		return fmt.Errorf("post %v err: %v, status code: %v, result: %v", address, err, status, string(data))
	}
	if status != 200 {
		return fmt.Errorf("post status code: %v, result: %v", status, string(data))
	}
	return nil
}
