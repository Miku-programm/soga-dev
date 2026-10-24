// 根据样本 RTTI 及 0x174dfe0、0x174e060、0x174e180 恢复，不是作者源码原文。
package CAodbhZf

import (
	"encoding/json"
	"strconv"
)

// JlKvxH 对应 noIJiSXMP.JlKvxH[int]，RTTI 0x35f7fb8，原包路径 CAodbhZf。
// 这里只表示已验证的 int 实例；二进制不能证明作者的完整泛型约束。
type JlKvxH[T int] struct{ _2L2dhl T }

// Value 对应 0x174e180。
func (n *JlKvxH[T]) Value() T { return n._2L2dhl }

// MarshalJSON 对应 0x174dfe0，经 JSON v1 语义编码内部数值。
func (n *JlKvxH[T]) MarshalJSON() ([]byte, error) { return json.Marshal(n._2L2dhl) }

// UnmarshalJSON 对应 0x174e060：先解析 json.Number，再以 float64 解析并转换到 int。
// 小数朝零截断；任何解析错误都不修改已有值。
func (n *JlKvxH[T]) UnmarshalJSON(data []byte) error {
	var number json.Number
	if err := json.Unmarshal(data, &number); err != nil {
		return err
	}
	value, err := strconv.ParseFloat(string(number), 64)
	if err != nil {
		return err
	}
	n._2L2dhl = T(value)
	return nil
}
