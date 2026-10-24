package user

import (
	"fmt"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/panel/types"
)

type V0wEAoUPFk struct {
	*AhSFeSua
	xWfXJaXWsM1k string
}

// PB4jVTrQHa recovers 0x181e3e0. The initial password is installed before
// UpdatePanelUser, so construction does not reset the newly created context.
func PB4jVTrQHa(panel *types.Aga91y, nodeLimiter *relay.L337vk2G) (*V0wEAoUPFk, error) {
	credential, ok := panel.JAfXWwzYrW.(*types.WQM2tKA)
	if !ok {
		return nil, pNF7CSleWS(panel.WZ89pTAOrHK4, "hysteria")
	}
	u := &V0wEAoUPFk{AhSFeSua: cLLKJ_A31(panel, nodeLimiter), xWfXJaXWsM1k: credential.CBWeRfrO1ahs}
	if _, err := u.UpdatePanelUser(panel); err != nil {
		return nil, err
	}
	return u, nil
}

// UpdatePanelUser recovers 0x181e520.
func (u *V0wEAoUPFk) UpdatePanelUser(panel *types.Aga91y) (bool, error) {
	u.AhSFeSua.UpdateUser(panel)
	credential, ok := panel.JAfXWwzYrW.(*types.WQM2tKA)
	if !ok {
		return false, pNF7CSleWS(panel.WZ89pTAOrHK4, "hysteria")
	}
	changed := credential.CBWeRfrO1ahs != u.xWfXJaXWsM1k
	if changed {
		u.xWfXJaXWsM1k = credential.CBWeRfrO1ahs
		u.AhSFeSua.ResetCtx()
	}
	return changed, nil
}

func (u *V0wEAoUPFk) GetPassword() string { return u.xWfXJaXWsM1k }
func (u *V0wEAoUPFk) String() string {
	return fmt.Sprintf("id: %v, password: %v", u.djdKrU, u.xWfXJaXWsM1k)
}
