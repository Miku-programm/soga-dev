//go:build linux

package app

import (
	"errors"
	"fmt"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
)

// Original named-key maps at 0x5095c50 and 0x5095c58. Every original entry is
// present; dj retains its distinct constructor and optional status-report loop.
var starterConstructorsAt5095C50 = map[C.FSOZSH7g]bRy8oqwuq6cY{
	"dj": yo9Lzz2t, "ppanel": eXC63toqu, "proxypanel": eXC63toqu,
	"soga-v1": eXC63toqu, "sspanel-uim": eXC63toqu, "v2board": eXC63toqu,
	"v2raysocks": eXC63toqu, "whmcs": eXC63toqu, "xboard": eXC63toqu, "xiaov2board": eXC63toqu,
}
var starterCallbacksAt5095C58 = map[C.FSOZSH7g]_Zpu8_d{
	"dj": u1nzch, "ppanel": aGvHzoD3fr, "proxypanel": yyCM1ej,
	"soga-v1": jQGVEC1W, "sspanel-uim": joTwyrx, "v2board": lag2OyJwFxd,
	"v2raysocks": ssDTthclJN, "whmcs": oXJigpaEaR, "xboard": qu_iEPVMzK9e, "xiaov2board": uaxpQZ,
}

// AewQql2U @0x1fa0f20. No spelling normalization. The original missing-callback
// branch logs "impossible" and exits through ROUJ82_bIQ; it does not silently
// substitute another callback. Literal decoders are recorded separately.
func AewQql2U(kind C.FSOZSH7g) (G6miUObWbEE, error) {
	if construct, ok := starterConstructorsAt5095C50[kind]; ok {
		callback, exists := starterCallbacksAt5095C58[kind]
		if !exists {
			Log.ROUJ82_bIQ("impossible")
		}
		return construct(callback), nil
	}
	const format = "unknown panel: %v"
	err := fmt.Errorf(format, kind)
	if err == nil {
		err = errors.New(format)
	}
	return nil, err
}
