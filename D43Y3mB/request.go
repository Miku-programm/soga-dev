package D43Y3mB

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recovered.local/soga/CAodbhZf"
	H "recovered.local/soga/G6bHD5JA"
	C "recovered.local/soga/cNT24rG4F"
	"strconv"
)

// GetAPIConfig 对应 0x16b99a0，返回原配置指针。
// D43Y3mB.(*GMLWI3).GetAPIConfig：原样本入口 0x16b99a0。
func (p *GMLWI3) GetAPIConfig() *C.FCEAfQJwDA7e { return p.TMfNUn }

// qjNarhJeaZH 对应 0x16b99c0；原时间戳先截为 uint32 后转十进制。
// D43Y3mB.(*GMLWI3).qjNarhJeaZH：原样本入口 0x16b99c0。
func (p *GMLWI3) qjNarhJeaZH() http.Header {
	h := make(http.Header)
	h.Add("key", p.TMfNUn.Bd08Li)
	h.Add("timestamp", strconv.FormatUint(uint64(uint32(CAodbhZf.C5Jff5vs0())), 10))
	return h
}

// GetCommon 对应 0x16b9c00；原读取路径不额外检查 HTTP 状态码。
// D43Y3mB.(*GMLWI3).GetCommon：原样本入口 0x16b9c00。
func (p *GMLWI3) GetCommon(address string) ([]byte, error) {
	_, body, _, err := H.IlWBbp1xdPw(address, nil, p.qjNarhJeaZH())
	if err != nil {
		return nil, err
	}
	return body, nil
}

// PostCommon 对应 0x16b9ca0，保留传输错误/HTTP状态/JSON/code四层判断及原错误文本。
// D43Y3mB.(*GMLWI3).PostCommon：原样本入口 0x16b9ca0。
func (p *GMLWI3) PostCommon(address string, value any) error {
	body, status, err := H.ZBovtj1xky(address, nil, value, p.qjNarhJeaZH())
	if err != nil {
		return fmt.Errorf("%v, result: %v, status code: %v", err, string(body), status)
	}
	if status != 200 {
		return fmt.Errorf("request failed, result: %v, status code: %v", string(body), status)
	}
	var response EIx7G7QCW
	if err = json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("%v, result: %v, status code: %v", err, string(body), status)
	}
	if response.NIg3b55LCOZ != 200 {
		return fmt.Errorf("result: %v", string(body))
	}
	return nil
}
