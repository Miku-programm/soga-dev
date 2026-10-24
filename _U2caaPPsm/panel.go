// 从原样本 0x169e9a0—0x169f8c0、0x16a13a0 恢复；非作者源码原文。
package _U2caaPPsm

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/url"
	G "recovered.local/soga/G6bHD5JA"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	X "recovered.local/soga/xx5fRyrFw"
	"time"
)

type GYKBwK3t_aMN struct {
	JgRscM5NlW []string `json:"domain"`
}
type Biktr6aZofGI struct {
	BK3hpLmXsZ []A_JUhZ `json:"clients"`
}
type MyopJI struct {
	UwS4pcm []GYKBwK3t_aMN `json:"rules"`
}
type TdT9I0uMiAKm struct {
	W325LG  int          `json:"port"`
	BjjANcz Biktr6aZofGI `json:"settings"`
}
type E1rI_uHRR1 struct {
	Dps5ZWL  []TdT9I0uMiAKm `json:"inbounds"`
	YhC3piOh MyopJI         `json:"routing"`
}
type A_JUhZ struct {
	RRkv8aQBcG9 string `json:"id"`
	VZfvWQe6yy  string `json:"email"`
	E5QPI2L     int    `json:"alterId"`
}
type Iw4id8kfTTT struct {
	jamJA0Qs  *C.FCEAfQJwDA7e
	tdXPhll10 *DpeSn3T
}

// 原全局 0x5095e48；成功取得配置时替换，原访问未加锁。
var emailByIDAt5095E48 map[uint32]string

func lFFXD1M(email string) uint32 {
	sum := md5.Sum([]byte(email))
	id := binary.BigEndian.Uint32(sum[:4])
	emailByIDAt5095E48[id] = email
	return id
}
func (u *A_JUhZ) GetUser() *L.Aga91y {
	id := lFFXD1M(u.VZfvWQe6yy)
	return &L.Aga91y{WZ89pTAOrHK4: int64(id), LmOKwjha: C.PanelUserSpeedLimit(), Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit(), JAfXWwzYrW: &L.Yn8BJoL{SfiEGpJo: u.RRkv8aQBcG9, FvQ2ltKOM7: u.E5QPI2L}}
}
func (c *E1rI_uHRR1) ToPanelNode() *L.Wr5SVP {
	return &L.Wr5SVP{HSA2b5V: true, SwxJNgPpg: L.JW0CeHo3O{CKsOpayd: c.Dps5ZWL[1].W325LG, VUwf1n9g: -1, F2GXyZZ: "ws", UVeahqDBT: true, ChJnngzb: "/ws"}}
}
func (c *E1rI_uHRR1) GetUsers() L.Sxcf3wu3L {
	clients := c.Dps5ZWL[1].BjjANcz.BK3hpLmXsZ
	out := make(L.Sxcf3wu3L, 0, len(clients))
	for _, u := range clients {
		out = append(out, u.GetUser())
	}
	return out
}
func (c *E1rI_uHRR1) ToDetectors() X.Ij9r21E0UOt3 {
	out := make(X.Ij9r21E0UOt3, 0, len(c.YhC3piOh.UwS4pcm))
	for _, rule := range c.YhC3piOh.UwS4pcm {
		for _, domain := range rule.JgRscM5NlW {
			out = append(out, X.NewPanelDetector(0, domain))
		}
	}
	return out
}
func init()                                          { Q.PanelFactoriesAt5095E50["dj-v2ray-webapi"] = gIxmlk }
func gIxmlk(c *C.FCEAfQJwDA7e) Q.VyFk5uT             { return &Iw4id8kfTTT{jamJA0Qs: c} }
func (p *Iw4id8kfTTT) GetAPIConfig() *C.FCEAfQJwDA7e { return p.jamJA0Qs }
func (p *Iw4id8kfTTT) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	body, _, err := G.J1liFniHg(p.jamJA0Qs.W_ogGtF+"sub_server/get_config", url.Values{"secret": {p.jamJA0Qs.Bd08Li}})
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	config := new(E1rI_uHRR1)
	if err = json.Unmarshal(body, config); err != nil {
		return nil, fmt.Errorf("parse node failed: %v", err)
	}
	emailByIDAt5095E48 = make(map[uint32]string)
	node := config.ToPanelNode()
	users := config.GetUsers()
	detectors := config.ToDetectors()
	return &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: users, HYnfYWa: detectors}, nil
}
func (p *Iw4id8kfTTT) GetSubmitInterval() time.Duration { return 3 * time.Second }
func (p *Iw4id8kfTTT) SubmitData(data *Q.FGmNba4MhioY) error {
	traffic := make(map[string]map[string]int64)
	for _, entry := range data.HTgPyE3VytZd {
		if email, ok := emailByIDAt5095E48[uint32(entry.Hy7Q97RROK)]; ok {
			traffic[email] = map[string]int64{"upload": entry.XJkGTQf_qkAt, "download": entry.I6LXObODav}
		}
	}
	body, err := json.Marshal(traffic)
	if err != nil {
		return fmt.Errorf("submit traffic failed: %v", err)
	}
	_, _, err = G.J1liFniHg(p.jamJA0Qs.W_ogGtF+"sub_server/submit_traffic", url.Values{"secret": {p.jamJA0Qs.Bd08Li}, "email_traffic": {string(body)}})
	if err != nil {
		return fmt.Errorf("submit traffic failed: %v", err)
	}
	return nil
}
