package jAP8Cn

import (
	"fmt"
	"sync"

	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/GYzZ5v"
)

var geoSiteMutexAt50C7FA0 sync.RWMutex
var geoSiteFileAt5095C38 *CAodbhZf.CCVciOvWaCEr

// ZRQq51fC00J（0xe9ef40）优先使用已映射的数据；Data 返回 nil 时
// 回退到二进制内嵌数据。非 nil 的空切片不会触发回退。
func ZRQq51fC00J(code string) (*GYzZ5v.P3tOWWIHzX, bool, error) {
	geoSiteMutexAt50C7FA0.RLock()
	defer geoSiteMutexAt50C7FA0.RUnlock()
	var data []byte
	if geoSiteFileAt5095C38 != nil {
		data, _ = geoSiteFileAt5095C38.Data()
	}
	if data == nil {
		data = embeddedGeoSiteAt502DA20
	}
	site, err := mcrrOY(data, code)
	if err != nil {
		return nil, false, fmt.Errorf("find geosite failed: %v", err)
	}
	return site, site != nil, nil
}
