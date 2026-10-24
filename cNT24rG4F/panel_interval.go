package cNT24rG4F

import (
	"recovered.local/soga/uLofn6j"
	"time"
)

// EVzHyyNDe 对应 0x132a3a0。原区间是 [max(0,d-r),d+r)，保留溢出和非法范围 panic。
func EVzHyyNDe(d time.Duration) time.Duration {
	spread := time.Duration(configAt50A2900.KS63SD) * time.Second
	if spread == 0 {
		return d
	}
	low, high := d-spread, d+spread
	if low < 0 {
		low = 0
	}
	return low + time.Duration(uLofn6j.TfHpoOcv(int64(high-low)))
}

// PanelPollingIntervals 只读桥：对应原两个 atomic 初值读取 0x50A2900/0x50A2908。
func PanelPollingIntervals() (check, submit time.Duration) {
	return time.Duration(configAt50A2900.Qxaw2_S) * time.Second, time.Duration(configAt50A2900.Fak_7LgVV) * time.Second
}
