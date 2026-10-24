package inbound

import (
	"errors"
	"fmt"
	C "recovered.local/soga/config"
)

// Original global 0x5095e58, allocated in inbound.init at 0x1d233e5.
// This is a named-string-keyed map of the original constructor function type.
// It is initialized before the eight registration init functions.
var inboundConstructorsAt5095E58 = make(map[C.Mpn2JltM]fH6caiUi)

func init() { // Original init.0 / init.1 / init.2 / init.3 / init.4 / init.5 / init.6 / init.7.
	inboundConstructorsAt5095E58["anytls"] = aboatPBa      // 0x1d33420
	inboundConstructorsAt5095E58["hysteria"] = hpdxYbrDSPs // 0x1d569a0
	inboundConstructorsAt5095E58["mieru"] = tzgQP2vl       // 0x1d5f240
	inboundConstructorsAt5095E58["ss"] = aXSbxcDLdeEY      // 0x1d6d0e0
	inboundConstructorsAt5095E58["ssr"] = x4oEZV           // 0x1d7da60
	inboundConstructorsAt5095E58["trojan"] = f1mPzegGt     // 0x1d8e620
	inboundConstructorsAt5095E58["vless"] = oHr1LeATmS     // 0x1d965e0
	inboundConstructorsAt5095E58["v2ray"] = ag90n_rZjJO    // 0x1da0ce0
}

// Pb5l5RLOHSi (0x1d6cc00) uses DV_XPJ4z0.POpyjFWF8v at offset 16;
// it does not derive a protocol from the panel or normalize the supplied key.
// The constructor receives the same argument pointer and its results are
// returned unchanged. A nil argument retains the original dereference panic.
func Pb5l5RLOHSi(arg *DV_XPJ4z0) (Kz_8ajU7, error) {
	if constructor, ok := inboundConstructorsAt5095E58[arg.POpyjFWF8v]; ok {
		return constructor(arg)
	}
	const format = "unknown server type: %v" // Decoder 0x1d6cda0, callsite 0x1d6cc58.
	err := fmt.Errorf(format, arg.POpyjFWF8v)
	if err == nil {
		err = errors.New(format)
	} // Retain the original helper-result fallback.
	return nil, err
}
