//go:build linux

package CbO_7wp3B

import (
	"time"

	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/q5bqCgTdxaTM"
)

// 四个独立状态机分别核对了立即成功、前两次失败、全部失败路径。
// 入口地址：dj 0x1fa7820、proxypanel 0x1fba580、
// v2raysocks 0x1fc7720、whmcs 0x1fc9400。
func u1nzch(count int)     { directStarterCallback(count, q5bqCgTdxaTM.Aod2QUfu) }
func yyCM1ej(count int)    { directStarterCallback(count, q5bqCgTdxaTM.WIkyLxzu) }
func ssDTthclJN(count int) { directStarterCallback(count, q5bqCgTdxaTM.CWcYaB_UjkSN) }
func oXJigpaEaR(count int) { directStarterCallback(count, q5bqCgTdxaTM.EcfkJGx1ZPwc) }

func directStarterCallback(count int, check func() error) {
	MFFa9yT.SetStarterUserCount(count)
	for attempt := 0; attempt < 3; attempt++ {
		err := check()
		if err == nil {
			cNT24rG4F.T15Jifo(false)
			return
		}
		if attempt == 2 {
			muc279Mpa6S4.FcL3BVMFuVAn(err)
			cNT24rG4F.T15Jifo(true)
			return
		}
		time.Sleep(time.Second)
	}
}
