package sspaneltypes

import (
	"encoding/json"
	"encoding/json/jsontext"
	L "recovered.local/soga/panel/types"
	F "recovered.local/soga/panel/ssrnames"
	"strconv"
)

// UnmarshalJSON 对应 0x16ecf00：外层失败不写字段，custom_config 错误被原样忽略。
func (n *NNrxxXeBt) UnmarshalJSON(b []byte) error {
	var v struct {
		VoQAT3uGq7  float64        `json:"node_speedlimit"`
		SwndEwba    float64        `json:"traffic_rate"`
		RzKa3NYa4   string         `json:"server"`
		U5aBm0NKS   string         `json:"type"`
		YLTkEdV7bE6 int            `json:"mu_only"`
		UAEOTcPzv   jsontext.Value `json:"custom_config"`
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	n.OFIQtbK5h = v.VoQAT3uGq7
	n.Z2pKW8byss = v.SwndEwba
	n.HIgDb6THe1A9 = v.RzKa3NYa4
	n.JCP2YA26 = v.U5aBm0NKS
	n.A0zacuR6Nzb = v.YLTkEdV7bE6
	_ = json.Unmarshal(v.UAEOTcPzv, &n.PzPwREkQp)
	return nil
}
func (c *Z3x174Dd4P3) IsVLess() bool { return c.DQ3N8yg6d == "1" || c.DQ3N8yg6d == "true" }
func (c *Z3x174Dd4P3) GetAlterId() int {
	n, err := strconv.Atoi(c.SFDGfYC)
	if err != nil {
		return -1
	}
	return n
}
func (c *Z3x174Dd4P3) GetStreamType() L.BoulTn31 { return L.BoulTn31(c.DwqtUCI) }
func (c *Z3x174Dd4P3) IsTls() bool               { return c.QR9CYdJCnh == "tls" || c.QR9CYdJCnh == "xtls" }
func (c *HS63ppHsdS) GetMethod() F.ZvAxOt1l7Ar   { return F.ZvAxOt1l7Ar(c.PyqeODiH) }
func (c *HS63ppHsdS) GetProtocol() F.Mu0vMIIs    { return F.Mu0vMIIs(c.BaCpMmr7) }
func (c *HS63ppHsdS) GetObfs() F.PMorzD7u        { return F.PMorzD7u(c.PDnKxsV8fttw) }
func (c *CEJsaaVtomQg) uX0Wxx() int {
	n, _ := strconv.Atoi(c.ZeH9JLmUs)
	if n <= 0 {
		n, _ = strconv.Atoi(c.ITA1ZyGMvV2)
	}
	return n
}
func (c *CEJsaaVtomQg) UpdateVMessNode(node *L.JW0CeHo3O) {
	v := c.Z3x174Dd4P3
	n, _ := strconv.Atoi(v.UPnr_Bjaq)
	if n <= 0 {
		n = c.uX0Wxx()
	}
	if n > 0 {
		node.CKsOpayd = n
	}
	if alter := v.GetAlterId(); alter >= 0 {
		node.VUwf1n9g = alter
	}
	if v.DwqtUCI != "" {
		node.F2GXyZZ = L.BoulTn31(v.DwqtUCI)
	}
	if v.IsTls() {
		node.UVeahqDBT = true
	}
	if v.JlkGYXM9 != "" {
		node.ChJnngzb = v.JlkGYXM9
	}
	if v.AMntbVC62jQS != "" {
		node.KWffzRDF = c.Z3x174Dd4P3.AMntbVC62jQS
	}
}
func (c *CEJsaaVtomQg) UpdateTrojanNode(node *L.L9AiYC) {
	n, _ := strconv.Atoi(c.VpDpbOJEJ)
	if n <= 0 {
		n = c.uX0Wxx()
	}
	if n > 0 {
		node.GbD_2PU = n
	}
	if c.DwqtUCI != "" {
		node.T4cuhG7MTbG = L.BoulTn31(c.DwqtUCI)
	}
	if c.JlkGYXM9 != "" {
		node.PBDb3Wc = c.JlkGYXM9
	}
	if c.AMntbVC62jQS != "" {
		node.QVlaSCZ4RCK = c.AMntbVC62jQS
	}
}
