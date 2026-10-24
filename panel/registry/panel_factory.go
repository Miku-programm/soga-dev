package registry

import (
	"fmt"
	C "recovered.local/soga/config"
)

// b1We2wTC9Hyw is the original named constructor function type.
type b1We2wTC9Hyw func(*C.FCEAfQJwDA7e) VyFk5uT

// PanelFactoriesAt5095E50 exports the original shared map under an analyst VA
// name, allowing recovered registration packages to reproduce inlined writes.
// Constructors are registered by their own original package init functions.
var PanelFactoriesAt5095E50 = make(map[string]b1We2wTC9Hyw)

// SMkpVNb:0x15cb0a0. Exact panel-protocol-api lookup without normalization or
// fallback, preserving the configuration pointer passed into the constructor.
func SMkpVNb(config *C.FCEAfQJwDA7e) (VyFk5uT, error) {
	key := string(config.OS91X026uML) + "-" + string(config.ZyAhkq5th8bf) + "-" + string(config.F3PCiaIra)
	if factory, ok := PanelFactoriesAt5095E50[key]; ok {
		return factory(config), nil
	}
	return nil, fmt.Errorf("not support server type <%v> for panel <%v> using <%v>", config.ZyAhkq5th8bf, config.OS91X026uML, config.F3PCiaIra)
}
