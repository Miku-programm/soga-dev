package MFFa9yT

import (
	"fmt"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	"time"
)

// _7le3Hqii @0x1d6d460 directly concatenates the configured directory.
func (v *Lp07FzBLx9V) _7le3Hqii() string {
	return fmt.Sprintf("%vip-user-cache-ss-%v-%v", C.IPUserCacheDirectory(), C.AVktbLizjJl(), v.wiy6P3.Load().EMlXEc)
}

// saveShadowsocksIPCache is the source name for closure 0x1e6ea40.
func (v *Lp07FzBLx9V) saveShadowsocksIPCache() {
	ticker := time.NewTicker(300 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := v.wZhn2Y.Save(v._7le3Hqii()); err != nil {
				Log.V8GN5Stk(4, "save ip user cache err: %v", err)
			}
		case <-v.zb8f3cHWW8v.Done():
			return
		}
	}
}
