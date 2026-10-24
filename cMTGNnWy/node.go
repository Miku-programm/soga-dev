package cMTGNnWy

import (
	"encoding/json"
	"fmt"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	W "recovered.local/soga/w2YaTd9_"
	"time"
)

func (n *Yq4D8wXbIQyY) GetPullInterval() time.Duration {
	return time.Duration(n.OWYIa8fDu.F2WCOEc5Eu6) * time.Second
}
func (n *Yq4D8wXbIQyY) GetPushInterval() time.Duration {
	return time.Duration(n.OWYIa8fDu.I8kigte) * time.Second
}

// ToPanelNode 对应0x16a8d80；三个传输分支的接受范围不同，TLS检查顺序依原指令。
func (n *Yq4D8wXbIQyY) ToPanelNode(id int) (*L.Wr5SVP, error) {
	out := &L.Wr5SVP{EMlXEc: int(id), Z1TIgHOQ: Q.T0j_Rco(n), HSA2b5V: C.PanelForbiddenBitTorrent()}
	switch n.B0XVtyvn0k31 {
	case "vmess":
		v := new(Fp4V6S84pbeq)
		if err := json.Unmarshal(n.OWRl87iaah, v); err != nil {
			return nil, err
		}
		transport := v.AnWvvG4Vva
		if transport == "websocket" {
			transport = "ws"
		}
		switch transport {
		case "ws", "h2", "tcp", "grpc":
		default:
			return nil, fmt.Errorf("unknown transport type: %v", v.AnWvvG4Vva)
		}
		out.SwxJNgPpg = L.JW0CeHo3O{CKsOpayd: v.ASTH4uK7, F2GXyZZ: L.BoulTn31(transport), UVeahqDBT: v.VsSEJhwCJT == "tls", ChJnngzb: v.ZgnJBd5rnF.NjZ_Np, KWffzRDF: v.ZgnJBd5rnF.Uce3uX}
	case "shadowsocks":
		v := new(JV07s34JhN6z)
		if err := json.Unmarshal(n.OWRl87iaah, v); err != nil {
			return nil, err
		}
		out.FnMUny = L.AAMPrdcYTL6{Am3QoeF: v.SMh2QKAfSzk, HMkc4BgR: W.C6q3Upu(v.Afx4JqDX6KD4), VuhWqou: v.DaDgO0Sh}
	case "vless":
		v := new(LuuJNqNq1YTl)
		if err := json.Unmarshal(n.OWRl87iaah, v); err != nil {
			return nil, err
		}
		transport := v.UePfBt4Gx
		if transport == "websocket" {
			transport = "ws"
		}
		switch transport {
		case "ws", "h2", "tcp", "grpc":
		default:
			return nil, fmt.Errorf("unknown transport type: %v", v.UePfBt4Gx)
		}
		out.Ck6DxjB = L.VkkFLH{GJtOdKOWBiFP: v.KzTz1UaVet, XmIaPj3CzOF: L.BoulTn31(transport), WAdJ4ftdZFk: L.D7FHIoaa8s4a(v.U7TxoE88w8), Nd3mBux1: v.V51ymS4.NjZ_Np, J1efyloFyhk: v.V51ymS4.Uce3uX, NMzMUQ2: v.WbrpN6f5U, BANVIZ72wh: []string{v.SYTfrb.HbXeOi732YO}, CrS6C1: v.SYTfrb.HX23KFdaHxy, DxKJ7lSe3m8F: []string{v.SYTfrb.KpKrDeAYv}, IqSfNzea: fmt.Sprintf("%s:%d", v.SYTfrb.StpFTNt5U, v.SYTfrb.T2S3gmIMZ)}
	case "trojan":
		v := new(PtTswARD)
		if err := json.Unmarshal(n.OWRl87iaah, v); err != nil {
			return nil, err
		}
		transport := v.AnWvvG4Vva
		if transport == "websocket" {
			transport = "ws"
		}
		switch transport {
		case "ws", "tcp", "grpc":
		default:
			return nil, fmt.Errorf("unknown transport type: %v", v.AnWvvG4Vva)
		}
		if v.VsSEJhwCJT != "tls" {
			return nil, fmt.Errorf("unknown security type: %v", v.VsSEJhwCJT)
		}
		out.JhIYwEukjd = L.L9AiYC{GbD_2PU: v.ASTH4uK7, T4cuhG7MTbG: L.BoulTn31(transport), PBDb3Wc: v.ZgnJBd5rnF.NjZ_Np, QVlaSCZ4RCK: v.ZgnJBd5rnF.Uce3uX}
	case "hysteria2":
		v := new(XnaEdZLZi_Lz)
		if err := json.Unmarshal(n.OWRl87iaah, v); err != nil {
			return nil, err
		}
		out.AYf5kQSQW = L.ZMYytVS{Li6nvK1: 2, PIvpxN6YK: v.FD3yBPArxu, B30Jt2kO: "plain", UjgtNuouuKVl: v.AMjK_y00H3Vf}
		if v.AMjK_y00H3Vf != "" {
			out.AYf5kQSQW.B30Jt2kO = "salamander"
		}
	default:
		return nil, fmt.Errorf("unknown node protocol: %s", n.B0XVtyvn0k31)
	}
	return out, nil
}
