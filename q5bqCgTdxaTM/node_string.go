package q5bqCgTdxaTM

import (
	"encoding/json"
	"fmt"
)

// T0j_Rco 对应 0x15cc3a0：原 json/v2 Marshal 使用 v1 兼容选项。
// hkCx9WWHy.CGwtjqBl@0x5d92c0 是 fmt.Sprint 路径。
func T0j_Rco(value any) string {
	data, err := json.Marshal(value)
	if err == nil {
		return string(data)
	}
	return fmt.Sprint(value)
}
