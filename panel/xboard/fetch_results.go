// 根据原样本 0x1790f60、0x17913e0 及其错误字符串 decoder 恢复。
package xboard

import (
	"encoding/json"
	"fmt"
)

// iEok3SmDhK 对应 0x1790f60，按原 JSON v1 兼容选项解析 config 返回值。
func (p *KlO65BuwOt) iEok3SmDhK() (*OGz0Bb0, error) {
	data, err := p.lVcAo38Q87cu(p.rTN7wPdD("config"), nil)
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	result := new(OGz0Bb0)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	// 原非空 message 分支仍格式化刚刚成功的 Unmarshal 返回错误，即 <nil>。
	if result.QAc4G6IUW != "" {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	return result, nil
}

// kDrbLuahIFA6 对应 0x17913e0；这里只恢复原始响应，不替代用户限制或转换逻辑。
func (p *KlO65BuwOt) kDrbLuahIFA6() (*Sf0Wjail3Z, error) {
	data, err := p.lVcAo38Q87cu(p.rTN7wPdD("user"), nil)
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	result := new(Sf0Wjail3Z)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	if result.EWFMfZX != "" {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	return result, nil
}
