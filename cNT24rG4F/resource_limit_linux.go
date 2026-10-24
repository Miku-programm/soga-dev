package cNT24rG4F

import (
	"syscall"

	"recovered.local/soga/muc279Mpa6S4"
)

// MeWz428lV2Z 对应 0x133df60，将当前进程 RLIMIT_NOFILE 软限制提升到硬限制。
func MeWz428lV2Z() {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		muc279Mpa6S4.TRTbrquD(-4, "getting Rlimit failed:", err)
		return
	}
	if limit.Cur < limit.Max {
		limit.Cur = limit.Max
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
			muc279Mpa6S4.TRTbrquD(-4, "setting Rlimit failed", err)
		}
	}
}
