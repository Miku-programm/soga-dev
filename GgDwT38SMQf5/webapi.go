package GgDwT38SMQf5

import (
	"fmt"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	P "recovered.local/soga/PqJzgkx7s6D8"
	U "recovered.local/soga/U5ihx31fY"
	C "recovered.local/soga/cNT24rG4F"
	F "recovered.local/soga/hfOuV8fOe4L"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	W "recovered.local/soga/w2YaTd9_"
)

type OODwaPoXSI struct{ P.JYkPXk4ZdW }

func yMAu3n(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &OODwaPoXSI{P.JYkPXk4ZdW{EA_5pkz: c}} }
func init()                              { Q.PanelFactoriesAt5095E50["sspanel-uim-ss-webapi"] = yMAu3n }
func (p *OODwaPoXSI) GetNode() (*L.Wr5SVP, error) {
	r, err := p.GetNodeResult()
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	return r.ToSSNode(p.EA_5pkz.Yf1KqSqJpNI7)
}
func (p *OODwaPoXSI) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	out := new(Q.D5gTIZ5aVg)
	var users *U.O8T6nD3uPzb
	err := H.D5aDHlQDLp(func() error {
		v, e := p.GetNode()
		if e == nil {
			out.H47AzZsYHB1N = v
		}
		return e
	}, func() error {
		v, e := p.GetUsersResult()
		if e == nil {
			users = v
		}
		return e
	}, func() error {
		v, e := p.GetDetectors()
		if e == nil {
			out.HYnfYWa = v
		}
		return e
	})
	if err != nil {
		return out, err
	}
	if out.H47AzZsYHB1N.FnMUny.Am3QoeF > 0 {
		for _, u := range users.L8xOe2JrpGDq {
			u.BAyDRja4o = F.ZvAxOt1l7Ar(out.H47AzZsYHB1N.FnMUny.HMkc4BgR)
		}
		out.VTme1SuHL6i_ = users.GetUsers(p.EA_5pkz.ZyAhkq5th8bf)
		return out, nil
	}
	var carrier *U.QhOI5ti
	for _, u := range users.L8xOe2JrpGDq {
		if u.RiD2Bnl == 0 {
			continue
		}
		valid := false
		switch u.BAyDRja4o {
		case "aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "chacha20-ietf-poly1305", "2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm":
			valid = true
		}
		if valid && u.KzG7RsMW_4yx == "origin" && u.LEtQeEkSu == "plain" {
			carrier = u
			break
		}
	}
	if carrier == nil {
		return nil, fmt.Errorf("未找到 aead 单端口承载用户")
	}
	for _, u := range users.L8xOe2JrpGDq {
		u.BAyDRja4o = carrier.BAyDRja4o
	}
	out.VTme1SuHL6i_ = users.GetUsers(p.EA_5pkz.ZyAhkq5th8bf)
	out.H47AzZsYHB1N.FnMUny = L.AAMPrdcYTL6{Am3QoeF: carrier.TdPoGXzK, HMkc4BgR: W.C6q3Upu(carrier.BAyDRja4o), VuhWqou: carrier.ESm7aAEddE}
	return out, nil
}
