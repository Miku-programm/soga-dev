// 依据原二进制 0x17434c0、0x1743720、0x1743920、0x1744120 恢复。
// 解码字符串来自原 decoder 的限定执行；不是作者原始源码文件。
package ZrjfjK8Y_tM9

import (
	"encoding/json"
	"fmt"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/q5bqCgTdxaTM"
	"recovered.local/soga/xx5fRyrFw"
)

func (r *DX8SNQb7vp) pLKk1MatbA(rules []*LFWTTXF7JR.Axa8uKIB4JL) []*LFWTTXF7JR.Axa8uKIB4JL {
	if r.ReTngXeu != "dns" || r.CoShU6j == "" {
		return rules
	}
	patterns := r.d09KQUdaP05h()
	if len(patterns) == 0 {
		return rules
	}
	return append(rules, &LFWTTXF7JR.Axa8uKIB4JL{K8BNfziIgb: []string{r.CoShU6j}, CM0eMsqyP: "ipv4_first", CZ80ik: patterns})
}
func (r *DX8SNQb7vp) t2MGGp8qf(rules []*xx5fRyrFw.IXZNQqqXDsV) []*xx5fRyrFw.IXZNQqqXDsV {
	if r.ReTngXeu != "block" {
		return rules
	}
	for _, pattern := range r.d09KQUdaP05h() {
		rules = append(rules, xx5fRyrFw.NewPanelDetector(r.CEfy7Qj, pattern))
	}
	return rules
}
func (n *U3nQ9iaQj) GetDetectors() xx5fRyrFw.Ij9r21E0UOt3 {
	out := make(xx5fRyrFw.Ij9r21E0UOt3, 0)
	for _, route := range n.RqE21w {
		out = route.t2MGGp8qf(out)
	}
	return out
}
func (n *U3nQ9iaQj) ToPanelNode(id int) (*LFWTTXF7JR.Wr5SVP, error) {
	if n.X6Jwaaf492C == "" {
		n.X6Jwaaf492C = "tcp"
	}
	var path, service string
	switch n.X6Jwaaf492C {
	case "ws":
		var settings Re7aaf
		if err := json.Unmarshal(n.MDoXQFch8aWk, &settings); err != nil {
			return nil, fmt.Errorf("parse ws settings failed: %v, result: %v", err, string(n.MDoXQFch8aWk))
		}
		path = settings.JCRhkrluQQY
	case "grpc":
		var settings GIAIBDeJa25
		if err := json.Unmarshal(n.MDoXQFch8aWk, &settings); err != nil {
			return nil, fmt.Errorf("parse grpc settings failed: %v, result: %v", err, string(n.MDoXQFch8aWk))
		}
		service = settings.RVKq_aLvfWyF
	}
	raw := q5bqCgTdxaTM.T0j_Rco(n)
	forbidden := cNT24rG4F.PanelForbiddenBitTorrent()
	speed := cNT24rG4F.PanelNodeSpeedLimit()
	connections := cNT24rG4F.PanelUserConnectionLimit()
	rules := make([]*LFWTTXF7JR.Axa8uKIB4JL, 0)
	for _, route := range n.RqE21w {
		rules = route.pLKk1MatbA(rules)
	}
	return &LFWTTXF7JR.Wr5SVP{
		EMlXEc: id, Z1TIgHOQ: raw, HSA2b5V: forbidden, TtWKyh: speed, W3Dw_B4wgbM: connections, AjXcBfRcAXR: rules,
		SwxJNgPpg:  LFWTTXF7JR.JW0CeHo3O{CKsOpayd: n.BDLnnhaR5, F2GXyZZ: n.X6Jwaaf492C, UVeahqDBT: n.PqqmTY8L != 0, ChJnngzb: path, KWffzRDF: service},
		JhIYwEukjd: LFWTTXF7JR.L9AiYC{GbD_2PU: n.BDLnnhaR5, T4cuhG7MTbG: n.X6Jwaaf492C, PBDb3Wc: path, QVlaSCZ4RCK: service},
		FnMUny:     LFWTTXF7JR.AAMPrdcYTL6{Am3QoeF: n.BDLnnhaR5, HMkc4BgR: n.FYeHum, VuhWqou: n.D07VWhUsih},
	}, nil
}
