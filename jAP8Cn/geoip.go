package jAP8Cn

import (
	"errors"
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"
	"recovered.local/soga/GYzZ5v"
	"recovered.local/soga/muc279Mpa6S4"
)

var geoIPMutexAt50C7FC0 sync.RWMutex

// YugPy6709RY（0xe9f260）返回匹配的 GeoIP 数据。protobuf 解码失败时
// 原函数同时返回已分配的对象和 error；扫描失败则对象为 nil。
func YugPy6709RY(code string) (*GYzZ5v.E4umdYA, error) {
	geoIPMutexAt50C7FC0.RLock()
	defer geoIPMutexAt50C7FC0.RUnlock()
	var data []byte
	if geoIPFileAt5095C40 != nil {
		data, _ = geoIPFileAt5095C40.Data()
	}
	if data == nil {
		data = embeddedGeoIPAt502DA40
	}
	message, err := kbUmw5F78(data, code)
	if err != nil {
		// 两个独立原字串解码器给出的内容相同，保留原错误分类分支。
		if errors.Is(err, codeNotFoundErrorAt5096F80) {
			return nil, fmt.Errorf("geoip:%v not exist", code)
		}
		return nil, fmt.Errorf("geoip:%v not exist", code)
	}
	result := new(GYzZ5v.E4umdYA)
	err = proto.Unmarshal(message, result)
	if err == nil {
		muc279Mpa6S4.V8GN5Stk(-4, "load geoip:%v success", code)
	}
	return result, err
}
