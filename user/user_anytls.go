package user

import (
	"crypto/sha256"
	"fmt"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/panel/types"
)

// HaB1NEnqcph is the original 56-byte AnyTLS user record.
type HaB1NEnqcph struct {
	*AhSFeSua
	fMjUrMi   string
	w2oeIPBSz [32]byte
}

// NZl_wC1qYi recovers the constructor at 0x181cf00.
func NZl_wC1qYi(panel *types.Aga91y, nodeLimiter *relay.L337vk2G) (*HaB1NEnqcph, error) {
	u := &HaB1NEnqcph{AhSFeSua: cLLKJ_A31(panel, nodeLimiter)}
	if _, _, err := u.UpdatePanelUser(panel); err != nil {
		return nil, fmt.Errorf("update user %v failed: %v", u.djdKrU, err)
	}
	return u, nil
}

// UpdatePanelUser recovers 0x181d0e0. An unchanged password returns a zero
// previous-auth array. In particular, an initially empty password leaves
// the initial authentication array zero instead of hashing an empty string.
func (u *HaB1NEnqcph) UpdatePanelUser(panel *types.Aga91y) ([32]byte, bool, error) {
	u.AhSFeSua.UpdateUser(panel)
	var previous [32]byte
	credential, ok := panel.JAfXWwzYrW.(*types.ZtMFlBTZXyD)
	if !ok {
		return previous, false, pNF7CSleWS(panel.WZ89pTAOrHK4, "anytls")
	}
	changed := credential.CBWeRfrO1ahs != u.fMjUrMi
	if changed {
		previous = u.w2oeIPBSz
		u.fMjUrMi = credential.CBWeRfrO1ahs
		u.w2oeIPBSz = sha256.Sum256([]byte(u.fMjUrMi))
		u.AhSFeSua.ResetCtx()
	}
	return previous, changed, nil
}

// GetAuth recovers 0x181d2e0.
func (u *HaB1NEnqcph) GetAuth() [32]byte { return u.w2oeIPBSz }

// String recovers 0x181d320 and its decoded format string.
func (u *HaB1NEnqcph) String() string {
	return fmt.Sprintf("id: %v, password: %v", u.djdKrU, u.fMjUrMi)
}
