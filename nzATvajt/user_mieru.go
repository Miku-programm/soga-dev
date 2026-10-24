package nzATvajt

import (
	"fmt"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/LFWTTXF7JR"
)

type WIZQ28 struct {
	*AhSFeSua
	vW0QDzs_DIhU string
	birIwyB_DM   string
}

// U1KSvhrlJFSB recovers 0x181f1c0.
func U1KSvhrlJFSB(panel *LFWTTXF7JR.Aga91y, nodeLimiter *G6bHD5JA.L337vk2G) (*WIZQ28, error) {
	u := &WIZQ28{AhSFeSua: cLLKJ_A31(panel, nodeLimiter)}
	if _, err := u.UpdatePanelUser(panel); err != nil {
		return nil, fmt.Errorf("update user %v failed: %v", u.djdKrU, err)
	}
	return u, nil
}

// UpdatePanelUser recovers 0x181f380.
func (u *WIZQ28) UpdatePanelUser(panel *LFWTTXF7JR.Aga91y) (bool, error) {
	u.AhSFeSua.UpdateUser(panel)
	credential, ok := panel.JAfXWwzYrW.(*LFWTTXF7JR.DO9qqM9)
	if !ok {
		return false, pNF7CSleWS(panel.WZ89pTAOrHK4, "mieru")
	}
	changed := u.vW0QDzs_DIhU != credential.Ejujyaj || u.birIwyB_DM != credential.GltkJGvV
	if changed {
		u.vW0QDzs_DIhU = credential.Ejujyaj
		u.birIwyB_DM = credential.GltkJGvV
		u.AhSFeSua.ResetCtx()
	}
	return changed, nil
}

func (u *WIZQ28) GetUsername() string { return u.vW0QDzs_DIhU }
func (u *WIZQ28) GetPassword() string { return u.birIwyB_DM }
func (u *WIZQ28) String() string {
	return fmt.Sprintf("id: %v, username: %v, password: %v", u.djdKrU, u.vW0QDzs_DIhU, u.birIwyB_DM)
}
