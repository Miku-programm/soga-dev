// 根据原样本 0x1746480、0x1746900 及其错误字符串 decoder 恢复。
package webapibase

import (
	"encoding/json"
	"fmt"
)

// GetUniNode 对应 0x1746480，按原 JSON v1 兼容选项解析 config 返回值。
func (p *JZA3kKksUe) GetUniNode() (*U3nQ9iaQj, error) {
	data, err := p.GetCommon(p.GetUniUrl("config"), nil)
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	result := new(U3nQ9iaQj)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	// 原非空 message 分支仍格式化刚刚成功的 Unmarshal 返回错误，即 <nil>。
	if result.VKhFmhBGQ != "" {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	return result, nil
}

// GetUsersResult 对应 0x1746900；这里只恢复原始响应，不替代用户限制或转换逻辑。
func (p *JZA3kKksUe) GetUsersResult() (*D9MPLRTuA, error) {
	data, err := p.GetCommon(p.GetUniUrl("user"), nil)
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	result := new(D9MPLRTuA)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	if result.EWFMfZX != "" {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	return result, nil
}
