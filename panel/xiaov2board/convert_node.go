package xiaov2board

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

// aJAAq8pQ / aiB8JLHaT / gTcS5lgmN8 对应 0x179bb60/0x179bc80/0x179bee0。
func (r *FPZsKqad5Ox) aJAAq8pQ() []string {
	var out []string
	if err := json.Unmarshal(r.Fu2GAU2hpa1b, &out); err != nil {
		return hashutil.DAaBs7_C(string(r.Fu2GAU2hpa1b), ",", -1)
	}
	return out
}
func (r *FPZsKqad5Ox) aiB8JLHaT(rules []*L.Axa8uKIB4JL) []*L.Axa8uKIB4JL {
	if r.ReTngXeu != "dns" || r.CoShU6j == "" {
		return rules
	}
	patterns := r.aJAAq8pQ()
	if len(patterns) == 0 {
		return rules
	}
	return append(rules, &L.Axa8uKIB4JL{K8BNfziIgb: []string{r.CoShU6j}, CM0eMsqyP: "ipv4_first", CZ80ik: patterns})
}
func (r *FPZsKqad5Ox) gTcS5lgmN8(rules []*matcher.IXZNQqqXDsV) []*matcher.IXZNQqqXDsV {
	if r.ReTngXeu != "block" {
		return rules
	}
	for _, pattern := range r.aJAAq8pQ() {
		rules = append(rules, matcher.NewPanelDetector(r.CEfy7Qj, pattern))
	}
	return rules
}

// GetDetectors 对应 0x179e820。
func (n *BdT6k1HU8L) GetDetectors() matcher.Ij9r21E0UOt3 {
	out := make(matcher.Ij9r21E0UOt3, 0)
	for _, route := range n.JVdfiT {
		out = route.gTcS5lgmN8(out)
	}
	return out
}

// ToPanelNode 对应 0x179d320。原函数不为缺失的 network 补默认值。
func (n *BdT6k1HU8L) ToPanelNode(id int) (*L.Wr5SVP, error) {
	var path, service string
	var xhttp *xhttp.R0J19auWof
	switch n.FoQe96xgtPg {
	case "ws":
		var settings Xd0IG2Z
		if err := json.Unmarshal(n.PabrtEwnkJ, &settings); err != nil {
			return nil, fmt.Errorf("parse ws settings failed: %v, result: %v", err, string(n.PabrtEwnkJ))
		}
		path = settings.JCRhkrluQQY
	case "grpc":
		var settings GJOCQqa
		if err := json.Unmarshal(n.PabrtEwnkJ, &settings); err != nil {
			return nil, fmt.Errorf("parse grpc settings failed: %v, result: %v", err, string(n.PabrtEwnkJ))
		}
		service = settings.RVKq_aLvfWyF
	case "xhttp":
		var settings DUnaqZqqbnWW
		if err := json.Unmarshal(n.PabrtEwnkJ, &settings); err != nil {
			return nil, fmt.Errorf("parse xhttp settings failed: %v, result: %v", err, string(n.PabrtEwnkJ))
		}
		xhttp = settings.ToXHttpConfig()
	}
	security := L.D7FHIoaa8s4a("none")
	if n.Xg1j1ZA == 1 {
		security = "tls"
	} else if n.Xg1j1ZA == 2 {
		security = "reality"
	}
	dest := n.PNpnXla.KP60FJlF4w
	if dest == "" {
		dest = n.PNpnXla.VHVkiEr1NUE
	}
	dest = fmt.Sprintf("%v:%v", dest, n.PNpnXla.Nf0Qo4R5MqM.Value())
	names := n.PNpnXla.XK10gJ_O
	if len(names) == 0 {
		names = []string{n.PNpnXla.VHVkiEr1NUE}
	}
	shorts := n.PNpnXla.PU3pcO6LIV1
	if len(shorts) == 0 {
		shorts = []string{n.PNpnXla.IwPy4K9c}
	}
	decryption, err := n.kX3RSoCossI()
	legacySecurity := L.D7FHIoaa8s4a("tls")
	if n.edYRMY8 {
		legacySecurity = security
	}
	if err != nil {
		return nil, fmt.Errorf("unknown vless decryption: %v, err: %v", n.HDMF3C, err)
	}
	cert, key := n.wXQOQActA(id)
	raw := registry.T0j_Rco(n)
	forbidden, speed, connections := C.PanelForbiddenBitTorrent(), C.PanelNodeSpeedLimit(), C.PanelUserConnectionLimit()
	rules := make([]*L.Axa8uKIB4JL, 0)
	for _, route := range n.JVdfiT {
		rules = route.aiB8JLHaT(rules)
	}
	obfs := ssmethod.GBXRYZ(n.YHnnp66mM)
	padding := strings.Join(n.WJ6BZoDVKsX4, "\n")
	return &L.Wr5SVP{
		EMlXEc: id, I3k4JnNp7nB: cert, XATiVun: key, Z1TIgHOQ: raw, HSA2b5V: forbidden, TtWKyh: speed, W3Dw_B4wgbM: connections, AjXcBfRcAXR: rules,
		SwxJNgPpg:  L.JW0CeHo3O{CKsOpayd: n.IhUtwVuCp, F2GXyZZ: n.FoQe96xgtPg, UVeahqDBT: n.Xg1j1ZA != 0, ChJnngzb: path, KWffzRDF: service},
		JhIYwEukjd: L.L9AiYC{GbD_2PU: n.IhUtwVuCp, T4cuhG7MTbG: n.FoQe96xgtPg, AJjaqrBjj: legacySecurity, PBDb3Wc: path, QVlaSCZ4RCK: service, HIlblJq1: names, SmInTc: n.PNpnXla.K2YRRA, TJ8rFSEH: shorts, ZxabzIRCNeV: dest},
		FnMUny:     L.AAMPrdcYTL6{Am3QoeF: n.IhUtwVuCp, HMkc4BgR: n.QIdmzg42XK, VuhWqou: n.S7aO_SWe, ErKysFf2I: obfs, PzwWMqEpPguJ: n.Enhxlh.XQkYg9e9wLB7, DURh5VYO: n.Enhxlh.UzanNits},
		Ck6DxjB:    L.VkkFLH{GJtOdKOWBiFP: n.IhUtwVuCp, XmIaPj3CzOF: n.FoQe96xgtPg, WAdJ4ftdZFk: security, Nd3mBux1: path, J1efyloFyhk: service, NMzMUQ2: n.HfRMkQmY7, BANVIZ72wh: names, CrS6C1: n.PNpnXla.K2YRRA, DxKJ7lSe3m8F: shorts, IqSfNzea: dest, DlDK5Y9AS: decryption, FPhrzP645: xhttp},
		AYf5kQSQW:  L.ZMYytVS{Li6nvK1: n.LQQHdt4QjrX, PIvpxN6YK: n.IhUtwVuCp, B30Jt2kO: addrtype.B3YC1rotZ(n.YHnnp66mM), UjgtNuouuKVl: n.F7emlj, HnHarW: n.Ada6jz_, W4ggEz: n.MAKypb},
		QaZjDzw95:  L.Tt_Moslrup{CgE0xkW: n.IhUtwVuCp, UPnToXrM: padding, J6bt8AOW8fL8: legacySecurity, Jq7YxRdb18Xx: names, CTa7WN: n.PNpnXla.K2YRRA, NWbT7wk: shorts, F9P9aeOu23U: dest},
	}, nil
}
