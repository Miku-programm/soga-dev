package xboard

import (
	"encoding/json"
	"fmt"
	L "recovered.local/soga/panel/types"
	"recovered.local/soga/net/addrtype"
	"recovered.local/soga/transport/xhttp"
	C "recovered.local/soga/config"
	"recovered.local/soga/panel/registry"
	"recovered.local/soga/crypto/hashutil"
	"recovered.local/soga/crypto/ssmethod"
	"recovered.local/soga/audit/matcher"
	"strings"
)

// v4QYwh661u5 / de9oPsKx / ctgEywElYK 对应 0x1788620/0x1788740/0x17889a0。
func (r *CYKo3knGCrtI) v4QYwh661u5() []string {
	var out []string
	if err := json.Unmarshal(r.Fu2GAU2hpa1b, &out); err != nil {
		return hashutil.DAaBs7_C(string(r.Fu2GAU2hpa1b), ",", -1)
	}
	return out
}
func (r *CYKo3knGCrtI) de9oPsKx(rules []*L.Axa8uKIB4JL) []*L.Axa8uKIB4JL {
	if r.ReTngXeu != "dns" || r.CoShU6j == "" {
		return rules
	}
	patterns := r.v4QYwh661u5()
	if len(patterns) == 0 {
		return rules
	}
	return append(rules, &L.Axa8uKIB4JL{K8BNfziIgb: []string{r.CoShU6j}, CM0eMsqyP: "ipv4_first", CZ80ik: patterns})
}
func (r *CYKo3knGCrtI) ctgEywElYK(rules []*matcher.IXZNQqqXDsV) []*matcher.IXZNQqqXDsV {
	if r.ReTngXeu != "block" {
		return rules
	}
	for _, pattern := range r.v4QYwh661u5() {
		rules = append(rules, matcher.NewPanelDetector(r.CEfy7Qj, pattern))
	}
	return rules
}

// GetDetectors 对应 0x178ac60。
func (n *OGz0Bb0) GetDetectors() matcher.Ij9r21E0UOt3 {
	out := make(matcher.Ij9r21E0UOt3, 0)
	for _, route := range n.UgPWEe7mnbcK {
		out = route.ctgEywElYK(out)
	}
	return out
}

// ToPanelNode 对应 0x1789520；缺失的 network 保持原空值。
func (n *OGz0Bb0) ToPanelNode(id int) (*L.Wr5SVP, error) {
	var path, service string
	var xhttp *xhttp.R0J19auWof
	switch n.T7xGio {
	case "ws":
		var settings OFybGF5iIp
		if err := json.Unmarshal(n.MiIwSabwwMec, &settings); err != nil {
			return nil, fmt.Errorf("parse ws settings failed: %v, result: %v", err, string(n.MiIwSabwwMec))
		}
		path = settings.JCRhkrluQQY
	case "grpc":
		var settings YMnnZhxYFt
		if err := json.Unmarshal(n.MiIwSabwwMec, &settings); err != nil {
			return nil, fmt.Errorf("parse grpc settings failed: %v, result: %v", err, string(n.MiIwSabwwMec))
		}
		service = settings.RVKq_aLvfWyF
	case "xhttp":
		var settings SNlDfbf3j
		if err := json.Unmarshal(n.MiIwSabwwMec, &settings); err != nil {
			return nil, fmt.Errorf("parse xhttp settings failed: %v, result: %v", err, string(n.MiIwSabwwMec))
		}
		xhttp = settings.ToXHttpConfig()
	}
	security := L.D7FHIoaa8s4a("none")
	if n.GCZiNr1s == 1 {
		security = "tls"
	} else if n.GCZiNr1s == 2 {
		security = "reality"
	}
	obfs, password := "plain", ""
	if n.Zd6n_x4ye != "" {
		obfs = "salamander"
		if n.Zd6n_x4ye == "salamander" {
			password = n.FBsOqCUmDeVc
		} else {
			password = n.Zd6n_x4ye
		}
	}
	ech, err := n.pC0PsvYElMrC()
	if err != nil {
		return nil, err
	}
	if ech != "" {
		if _, err := L.YUH4nAWc(ech); err != nil {
			return nil, err
		}
	}
	cert, key := n.yU6F1oN6a(id)
	var plugin map[string]string
	switch n.KVZ5bT {
	case "":
	case "obfs":
		plugin = n.oN0fey()
	default:
		return nil, fmt.Errorf("unsupported ss plugin: %v", n.KVZ5bT)
	}
	raw := registry.T0j_Rco(n)
	forbidden, speed, connections := C.PanelForbiddenBitTorrent(), C.PanelNodeSpeedLimit(), C.PanelUserConnectionLimit()
	rules := make([]*L.Axa8uKIB4JL, 0)
	for _, route := range n.UgPWEe7mnbcK {
		rules = route.de9oPsKx(rules)
	}
	return &L.Wr5SVP{
		EMlXEc: id, LKMSis7zkX: ech, I3k4JnNp7nB: cert, XATiVun: key, Z1TIgHOQ: raw, HSA2b5V: forbidden, TtWKyh: speed, W3Dw_B4wgbM: connections, AjXcBfRcAXR: rules,
		SwxJNgPpg:  L.JW0CeHo3O{CKsOpayd: n.SYdzk2V2lQ, F2GXyZZ: n.T7xGio, UVeahqDBT: n.GCZiNr1s != 0, ChJnngzb: path, KWffzRDF: service},
		JhIYwEukjd: L.L9AiYC{GbD_2PU: n.SYdzk2V2lQ, T4cuhG7MTbG: n.T7xGio, AJjaqrBjj: security, PBDb3Wc: path, QVlaSCZ4RCK: service, HIlblJq1: []string{n.Biccak.NaZM2e}, SmInTc: n.Biccak.BK1oMAx32PV, TJ8rFSEH: []string{n.Biccak.Fx6PX5j40}, ZxabzIRCNeV: fmt.Sprintf("%v:%v", n.Biccak.NaZM2e, n.Biccak.IRzKwqKa.Value())},
		FnMUny:     L.AAMPrdcYTL6{Am3QoeF: n.SYdzk2V2lQ, HMkc4BgR: n.SAOC5LCS9Ot, VuhWqou: n.AUEY_z, ErKysFf2I: ssmethod.GBXRYZ(plugin["obfs"]), PzwWMqEpPguJ: plugin["path"], DURh5VYO: plugin["obfs-host"]},
		Ck6DxjB:    L.VkkFLH{GJtOdKOWBiFP: n.SYdzk2V2lQ, XmIaPj3CzOF: n.T7xGio, WAdJ4ftdZFk: security, Nd3mBux1: path, J1efyloFyhk: service, NMzMUQ2: n.JkDMoRe3f, BANVIZ72wh: []string{n.Biccak.NaZM2e}, CrS6C1: n.Biccak.BK1oMAx32PV, DxKJ7lSe3m8F: []string{n.Biccak.Fx6PX5j40}, IqSfNzea: fmt.Sprintf("%v:%v", n.Biccak.NaZM2e, n.Biccak.IRzKwqKa.Value()), DlDK5Y9AS: n.PdJsmVeR8so, FPhrzP645: xhttp},
		AYf5kQSQW:  L.ZMYytVS{Li6nvK1: n.H_GY_Mo, PIvpxN6YK: n.SYdzk2V2lQ, B30Jt2kO: addrtype.B3YC1rotZ(obfs), UjgtNuouuKVl: password, HnHarW: n.CkxHfcd, W4ggEz: n.W8TM3dWc},
		QaZjDzw95:  L.Tt_Moslrup{CgE0xkW: n.SYdzk2V2lQ, UPnToXrM: strings.Join(n.Y8Ije40Uo, "\n"), J6bt8AOW8fL8: "tls"},
		ABA4hq:     L.Xh4IiOa0{MQE3K4mq: n.SYdzk2V2lQ, GGSLzD: strings.ToUpper(n.ZOry8Au3q), CLWr9bQsb: n.AiySCX6nT},
	}, nil
}
