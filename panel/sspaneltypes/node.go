package sspaneltypes

import (
	"fmt"
	"net/url"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	M "recovered.local/soga/util/logx"
	Q "recovered.local/soga/panel/registry"
	W "recovered.local/soga/crypto/ssmethod"
	X "recovered.local/soga/audit/matcher"
	"strconv"
	"strings"
)

// jrW5oitRt @0x16e7d60 preserves Split rather than SplitN: extra '=' parts are discarded.
func jrW5oitRt(s string) url.Values {
	out := make(url.Values)
	for _, part := range strings.Split(s, "|") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		kv := strings.Split(part, "=")
		if len(kv) < 2 {
			M.TRTbrquD(4, "sspanel server key and value invalid:", part)
			continue
		}
		out.Add(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
	}
	return out
}

// caT62H0U8n @0x16ed800 modifies only the two port fields before parsing.
func caT62H0U8n(s string, n *L.Dv6QJyFW6bU) error {
	n.PU1Di18, n.C33jbIzI = 0, 0
	parts := strings.Split(s, ";")
	if len(parts) <= 1 {
		return nil
	}
	v := jrW5oitRt(parts[1])
	if x := v.Get("soga_port"); x != "" {
		var err error
		n.C33jbIzI, err = strconv.Atoi(x)
		if err != nil {
			return fmt.Errorf("%v: parse soga_port err: %v", s, err)
		}
	}
	if x := v.Get("port"); x != "" {
		var err error
		n.PU1Di18, err = strconv.Atoi(strings.Split(x, "#")[0])
		if err != nil {
			return fmt.Errorf("%v: parse port err: %v", s, err)
		}
	}
	return nil
}
func bmG8Nfaaa(s string, n *L.L9AiYC) error {
	n.GbD_2PU = 443
	parts := strings.Split(s, ";")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	if len(parts) == 1 || parts[1] == "" {
		return nil
	}
	x := jrW5oitRt(parts[1]).Get("port")
	if x == "" {
		return nil
	}
	ports := strings.Split(x, "#")
	x = ports[0]
	if len(ports) > 1 {
		x = ports[1]
	}
	v, err := strconv.Atoi(x)
	if err != nil {
		return fmt.Errorf("parse server failed: %v", s)
	}
	if v > 0 {
		n.GbD_2PU = v
	}
	return nil
}
func zDRJVbGhJ7v(n *L.JW0CeHo3O, s string) {
	v := jrW5oitRt(s)
	n.ChJnngzb = v.Get("path")
	n.KWffzRDF = v.Get("servicename")
	port, _ := strconv.Atoi(v.Get("inside_port"))
	if port > 0 {
		n.CKsOpayd = port
	}
}
func mAndhwjjX(s string, n *L.JW0CeHo3O) error {
	s = strings.TrimSpace(s)
	*n = L.JW0CeHo3O{CKsOpayd: 80, VUwf1n9g: 2, F2GXyZZ: "tcp", ChJnngzb: "/"}
	parts := strings.Split(s, ";")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	if len(parts) < 2 {
		return nil
	}
	var err error
	if parts[1] != "" {
		n.CKsOpayd, err = strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("parse server failed: %v, reason: port invalid: %v", s, parts[1])
		}
	} else {
		n.CKsOpayd = 0
	}
	if len(parts) < 3 {
		return nil
	}
	if parts[2] != "" {
		n.VUwf1n9g, err = strconv.Atoi(parts[2])
		if err != nil {
			return fmt.Errorf("parse server failed: %v, reason: alterId invalid: %v", s, parts[2])
		}
	} else {
		n.VUwf1n9g = 0
	}
	if len(parts) == 4 {
		return fmt.Errorf("parse server failed: %v", s)
	}
	if len(parts) < 5 {
		return nil
	}
	valid := func(s string) bool { return s == "h2" || s == "ws" || s == "tcp" || s == "grpc" }
	network, security := parts[3], parts[4]
	if !valid(network) {
		network, security = parts[4], parts[3]
		if !valid(network) {
			return fmt.Errorf("parse server failed: %v", s)
		}
	}
	n.F2GXyZZ = L.BoulTn31(network)
	n.UVeahqDBT = security == "tls" || security == "xtls"
	if len(parts) >= 6 {
		zDRJVbGhJ7v(n, parts[5])
	}
	return nil
}

// Original ordered comparisons retain the unordered (NaN) branches.
func mergeSpeed(a, b float64) float64 {
	if a < 0 {
		a = 0
	} else if a > 10000 {
		a = 10000
	}
	if b < 0 {
		b = 0
	} else if b > 10000 {
		b = 10000
	}
	if a > 0 && b > 0 {
		if !(b > a) {
			a = b
		}
	} else if !(a > b) {
		a = b
	}
	return a
}
func (n *H0izB2agvG) common(id int) *L.Wr5SVP {
	hash := Q.T0j_Rco(n.L8xOe2JrpGDq)
	return &L.Wr5SVP{EMlXEc: id, Z1TIgHOQ: hash, HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: mergeSpeed(n.L8xOe2JrpGDq.OFIQtbK5h, C.PanelNodeSpeedLimit()), W3Dw_B4wgbM: C.PanelUserConnectionLimit()}
}
func (n *H0izB2agvG) ToVMessNode(id int) (*L.Wr5SVP, error) {
	out := n.common(id)
	if err := mAndhwjjX(n.L8xOe2JrpGDq.HIgDb6THe1A9, &out.SwxJNgPpg); err != nil {
		return nil, err
	}
	n.L8xOe2JrpGDq.PzPwREkQp.UpdateVMessNode(&out.SwxJNgPpg)
	return out, nil
}
func (n *H0izB2agvG) ToTrojanNode(id int) (*L.Wr5SVP, error) {
	out := n.common(id)
	out.JhIYwEukjd.T4cuhG7MTbG = "tcp"
	if err := bmG8Nfaaa(n.L8xOe2JrpGDq.HIgDb6THe1A9, &out.JhIYwEukjd); err != nil {
		return nil, err
	}
	n.L8xOe2JrpGDq.PzPwREkQp.UpdateTrojanNode(&out.JhIYwEukjd)
	return out, nil
}
func (n *H0izB2agvG) ToSSRNode(id int) (*L.Wr5SVP, error) {
	out := n.common(id)
	out.ASJxBbFS.QawbKS2x = n.L8xOe2JrpGDq.A0zacuR6Nzb == 1
	if err := caT62H0U8n(n.L8xOe2JrpGDq.HIgDb6THe1A9, &out.ASJxBbFS); err != nil {
		return nil, err
	}
	n.L8xOe2JrpGDq.PzPwREkQp.UpdateSSRNode(&out.ASJxBbFS)
	return out, nil
}
func (n *H0izB2agvG) ToSSNode(id int) (*L.Wr5SVP, error) {
	out := n.common(id)
	n.L8xOe2JrpGDq.PzPwREkQp.UpdateSSNode(&out.FnMUny)
	return out, nil
}
func (c *CEJsaaVtomQg) UpdateSSNode(n *L.AAMPrdcYTL6) {
	n.Am3QoeF = c.uX0Wxx()
	if n.Am3QoeF <= 0 {
		return
	}
	if c.QUUtCN7gjjra != "" {
		n.HMkc4BgR = W.C6q3Upu(c.QUUtCN7gjjra)
		n.VuhWqou = c.IBgOeC
		return
	}
	cipher := W.C6q3Upu(c.PyqeODiH)
	switch cipher {
	case "aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "chacha20-ietf-poly1305", "2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm":
		n.HMkc4BgR = cipher
	default:
		M.TRTbrquD(4, "unknown custom_config ss cipher:", cipher)
	}
	n.VuhWqou = c.RNq55eY
}
func (d *E3s7VeiKCij) ToDetector() *X.IXZNQqqXDsV {
	rule := d.XEmS9xCt
	if !X.WUWGex(rule) {
		rule = "regexp:" + rule
	}
	return X.NewPanelDetector(d.Uyj0tudIZN, rule)
}
func (d *Nq0XWDbiQ) ToDetectors() X.Ij9r21E0UOt3 {
	out := make(X.Ij9r21E0UOt3, 0)
	for _, v := range d.L8xOe2JrpGDq {
		out = append(out, v.ToDetector())
	}
	return out
}
