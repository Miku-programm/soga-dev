// 根据原样本 RTTI 和 IDA 指令恢复，不是作者源码原文。
// 类型名称前缀 oyvvcZShtb 与 uncommon_package_path wd42zQ0FT 不同，按后者组织目录。
package wd42zQ0FT

import "sync"

// NG0hCZyRHK 对应 RTTI 0x35f5498；amd64 大小 48 字节。
type NG0hCZyRHK struct{ icFkNd sync.Map }

// LAQMf2tPf 对应 RTTI 0x363bae0；amd64 大小 32 字节。
type LAQMf2tPf struct {
	Z49q4d     string
	U6F6cw_aOk any
}

// AddETag 对应 0x16a8980；存储对象本身，不复制切片或其它对象。
func (c *NG0hCZyRHK) AddETag(key, etag string, obj any) {
	c.icFkNd.Store(key, &LAQMf2tPf{Z49q4d: etag, U6F6cw_aOk: obj})
}

// CheckETag 对应 0x16a8ac0；空标签可相等，匹配区分大小写。
func (c *NG0hCZyRHK) CheckETag(key, etag string) (any, bool) {
	value, ok := c.icFkNd.Load(key)
	if !ok {
		return nil, false
	}
	entry := value.(*LAQMf2tPf)
	if entry.Z49q4d == etag {
		return entry.U6F6cw_aOk, true
	}
	return nil, false
}

// llkc6c4 对应 0x16a8bc0；保留原始类型断言的 panic 行为。
func (c *NG0hCZyRHK) llkc6c4(key string) (*LAQMf2tPf, bool) {
	value, ok := c.icFkNd.Load(key)
	if !ok {
		return nil, false
	}
	return value.(*LAQMf2tPf), true
}

// GetETag 对应 0x16a8c60。
func (c *NG0hCZyRHK) GetETag(key string) string {
	entry, ok := c.llkc6c4(key)
	if !ok {
		return ""
	}
	return entry.Z49q4d
}

// GetObj 对应 0x16a8cc0；命中时即使对象是 nil 也返回 true。
func (c *NG0hCZyRHK) GetObj(key string) (any, bool) {
	entry, ok := c.llkc6c4(key)
	if !ok {
		return nil, false
	}
	return entry.U6F6cw_aOk, true
}
