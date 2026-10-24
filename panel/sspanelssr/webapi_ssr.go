package sspanelssr

import (
	"encoding/json"
	"fmt"
	L "recovered.local/soga/panel/types"
	U "recovered.local/soga/panel/sspaneltypes"
	F "recovered.local/soga/panel/ssrnames"
	M "recovered.local/soga/util/logx"
	Q "recovered.local/soga/panel/registry"
)

func (p *MpE9pa4fA) viBFPHIjmS5() (string, error) {
	body, err := p.GetCommon("nodes")
	if err != nil {
		return "", fmt.Errorf("get ssr node failed: %v, result: %v", err, string(body))
	}
	var out U.XgjalGC
	if err = json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("parse ssr node failed: %v, result: %v", err, string(body))
	}
	for _, n := range out.JhtMOC {
		if n.YIrMyYIR == p.EA_5pkz.Yf1KqSqJpNI7 {
			return n.Jn_QCrBDcb, nil
		}
	}
	return "", fmt.Errorf("can not find ssr node id: %v", p.EA_5pkz.Yf1KqSqJpNI7)
}

// Both original API wrappers perform this loop only after all three tasks succeed.
func addSSRUsers(out *Q.D5gTIZ5aVg) {
	node := out.H47AzZsYHB1N
	for _, u := range out.VTme1SuHL6i_ {
		if v, ok := u.JAfXWwzYrW.(*L.QcqpSL3i6ri); ok && (v.BskzckkE7qI == 1 || v.BskzckkE7qI == 2) {
			if err := F.PsDhAx0aUZ(v.Z_taAb, v.H_y2xZht7S8, v.ZZIDsJ); err != nil {
				M.TRTbrquD(4, err, "user id:", u.WZ89pTAOrHK4)
			} else {
				node.ASJxBbFS.AddSSRUser(u)
			}
		}
	}
}
