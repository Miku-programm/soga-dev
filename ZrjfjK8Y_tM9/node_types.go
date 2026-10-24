// 原样本 RTTI 数据类型表示，不是作者源码原文。
// 使用原始标准库 jsontext.Value，构建需要 GOEXPERIMENT=jsonv2。
package ZrjfjK8Y_tM9

import (
	"encoding/json"
	"encoding/json/jsontext"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/hfOuV8fOe4L"
	"recovered.local/soga/vKl1IHr3"
	"recovered.local/soga/w2YaTd9_"
)

// DX8SNQb7vp 对应0x36bb600，大小64字节；原类型前缀wyihwaFO。
type DX8SNQb7vp struct {
	CEfy7Qj      int            `json:"id"`
	Fu2GAU2hpa1b jsontext.Value `json:"match"`
	ReTngXeu     string         `json:"action"`
	CoShU6j      string         `json:"action_value,omitempty"`
}

// IUXU5ZXqeiRC 对应0x35da2f0，元素描述符0x36bb600。
type IUXU5ZXqeiRC []DX8SNQb7vp

// U3nQ9iaQj 对应0x3711f70，大小192字节；保留原字段类型及标签。
type U3nQ9iaQj struct {
	VKhFmhBGQ string       `json:"message"`
	BDLnnhaR5 int          `json:"server_port"`
	RqE21w    IUXU5ZXqeiRC `json:"routes"`
	WNvzoiMPp struct {
		F2WCOEc5Eu6 CAodbhZf.JlKvxH[int] `json:"pull_interval"`
		I8kigte     CAodbhZf.JlKvxH[int] `json:"push_interval"`
	} `json:"base_config"`
	X6Jwaaf492C  LFWTTXF7JR.BoulTn31  `json:"network"`
	MDoXQFch8aWk jsontext.Value       `json:"networkSettings"`
	PqqmTY8L     int                  `json:"tls"`
	ELUl7_FnS    string               `json:"host"`
	BNKMfAf5     string               `json:"server_name"`
	FYeHum       w2YaTd9_.C6q3Upu     `json:"cipher"`
	D07VWhUsih   string               `json:"server_key"`
	L6f4wB51G0cC hfOuV8fOe4L.PMorzD7u `json:"obfs"`
}

// d09KQUdaP05h 对应0x17433a0，原调用经JSON v1兼容选项进入json/v2。
// 若不能解析为字符串数组，直接对原始JSON字节按逗号分割；不先解开字符串引号。
func (r *DX8SNQb7vp) d09KQUdaP05h() []string {
	var result []string
	if err := json.Unmarshal(r.Fu2GAU2hpa1b, &result); err != nil {
		return vKl1IHr3.DAaBs7_C(string(r.Fu2GAU2hpa1b), ",", -1)
	}
	return result
}
