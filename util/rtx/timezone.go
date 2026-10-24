package rtx

import "time"

// TimeLocationAt5095CB0 names the original shared global by its address.
// runtime.init @0xb41d2b decodes Asia/Shanghai then invokes time.LoadLocation
// @0x5a5d80; its error is discarded, preserving nil on a failed zone lookup.
var TimeLocationAt5095CB0 = func() *time.Location {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return location
}()
