//go:build linux

package app

import (
	"time"

	"recovered.local/soga/protocol/inbound"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/panel/registry"
)

// 四个独立状态机分别核对了立即成功、前两次失败、全部失败路径。
// 入口地址：dj 0x1fa7820、proxypanel 0x1fba580、
// v2raysocks 0x1fc7720、whmcs 0x1fc9400。
func u1nzch(count int)     { directStarterCallback(count, registry.Aod2QUfu) }
func yyCM1ej(count int)    { directStarterCallback(count, registry.WIkyLxzu) }
func ssDTthclJN(count int) { directStarterCallback(count, registry.CWcYaB_UjkSN) }
func oXJigpaEaR(count int) { directStarterCallback(count, registry.EcfkJGx1ZPwc) }

func directStarterCallback(count int, check func() error) {
	inbound.SetStarterUserCount(count)
	for attempt := 0; attempt < 3; attempt++ {
		err := check()
		if err == nil {
			config.T15Jifo(false)
			return
		}
		if attempt == 2 {
			logx.FcL3BVMFuVAn(err)
			config.T15Jifo(true)
			return
		}
		time.Sleep(time.Second)
	}
}
