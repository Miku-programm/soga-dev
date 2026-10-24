package geo

import (
	"context"
	"os"

	"recovered.local/soga/util/rtx"
	"recovered.local/soga/util/logx"
)

var geoIPFileAt5095C40 *rtx.CCVciOvWaCEr

// idPXJg8（0xe9e6e0）的返回值由原解码指令恢复。
func idPXJg8() string { return "/etc/soga/geosite.dat" }

// Tac58sht（0xe9dea0）依次加载两个数据库；不存在时写入原嵌入内容。
// 原函数不创建父目录，写入失败只记录日志，映射失败会把对应全局置 nil。
func Tac58sht() {
	initializeGeoFile(idPXJg8(), embeddedGeoSiteAt502DA20, &geoSiteFileAt5095C38,
		"write geosite data to %v failed: %v", "load geosite %v failed: %v")
	initializeGeoFile("/etc/soga/geoip.dat", embeddedGeoIPAt502DA40, &geoIPFileAt5095C40,
		"write geoip data to %v failed: %v", "load geoip %v failed: %v")
}

// 此辅助名由分析者指定；两组分支在原函数中的执行顺序保持不变。
func initializeGeoFile(path string, data []byte, destination **rtx.CCVciOvWaCEr, writeError, mapError string) {
	if !rtx.VSCJLy2No(path) {
		if err := os.WriteFile(path, data, 0644); err != nil {
			logx.V8GN5Stk(8, writeError, path, err)
			return
		}
	}
	var err error
	*destination, err = rtx.LWzrRP(context.Background(), path)
	if err != nil {
		logx.V8GN5Stk(8, mapError, path, err)
	}
}
