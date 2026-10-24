package errutil

import L "recovered.local/soga/util/logx"

// SlG_Y_mg:0xb25d20. Must be called directly by defer for recover semantics.
// A blank prefix still consumes the panic, but suppresses the original log.
func SlG_Y_mg(prefix string) any {
	value := recover()
	if value != nil && len(prefix) != 0 {
		L.TRTbrquD(8, prefix, "panic:", value)
	}
	return value
}
