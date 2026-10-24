package MFFa9yT

import (
	"fmt"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	"time"
)

// _7le3Hqii @0x1da1200。原字符串直接拼接目录，不自动补路径分隔符。
func (v *K4avFiYOPi) _7le3Hqii() string {
	return fmt.Sprintf("%vip-user-cache-vmess-%v-%v", C.IPUserCacheDirectory(), C.AVktbLizjJl(), v.wiy6P3.Load().EMlXEc)
}

// saveVmessIPCache 为闭包0x1e73ac0的源级命名；原周期300秒，退出时停止计时器。
func (v *K4avFiYOPi) saveVmessIPCache() {
	ticker := time.NewTicker(300 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := v.l_ZoZQDxG7.Save(v._7le3Hqii()); err != nil {
				Log.V8GN5Stk(4, "save ip user cache err: %v", err)
			}
		case <-v.zb8f3cHWW8v.Done():
			return
		}
	}
}
