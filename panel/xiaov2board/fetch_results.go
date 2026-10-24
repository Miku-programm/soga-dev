// 根据原样本 0x17a6420、0x17a6900 及其错误字符串 decoder 恢复。
package xiaov2board

import (
	"encoding/json"
	"fmt"
	"recovered.local/soga/config"
)

// GetUniNode 对应 0x17a6420，按原 JSON v1 兼容选项解析 config 返回值。
func (p *LRlqOaYmO) GetUniNode() (*BdT6k1HU8L, error) {
	data, err := p.GetCommon(p.kwoj3qch_(), nil)
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	result := new(BdT6k1HU8L)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	// 原非空 message 分支仍格式化刚刚成功的 Unmarshal 返回错误，即 <nil>。
	if result.Hb7ZhAG != "" {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	if config.XiaoV2NodeMode() {
		if err := result.ocVy8aVcKm(p.EA_5pkz.ZyAhkq5th8bf); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// GetUsersResult 对应 0x17a6900；这里只恢复原始响应，不替代用户限制或转换逻辑。
func (p *LRlqOaYmO) GetUsersResult() (*IMyQ5CxaikH, error) {
	data, err := p.GetCommon(p.GetUniUrl("user"), nil)
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	result := new(IMyQ5CxaikH)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	if result.EWFMfZX != "" {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	return result, nil
}
