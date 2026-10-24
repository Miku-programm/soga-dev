// 根据原样本 RTTI 0x35e9468 与 0xb45700、0xb457c0、0xb45840 恢复。
package CAodbhZf

import (
	"encoding/json"
	"strings"
)

type RdPirvOZCnMz struct{ _2L2dhl bool }

func (n *RdPirvOZCnMz) Value() bool                  { return n._2L2dhl }
func (n *RdPirvOZCnMz) MarshalJSON() ([]byte, error) { return json.Marshal(n._2L2dhl) }
func (n *RdPirvOZCnMz) UnmarshalJSON(data []byte) error {
	s := strings.ToLower(strings.Trim(string(data), "\""))
	n._2L2dhl = s == "1" || s == "true"
	return nil
}
