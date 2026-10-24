package aRZNjEvm

import (
	"encoding/base64"
	"fmt"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/q5bqCgTdxaTM"
	"recovered.local/soga/wd42zQ0FT"
	"time"
)

// AnxiJP 对应 0x17a5b00，保留原接口返回值与独立缓存分配。
func AnxiJP(config *cNT24rG4F.FCEAfQJwDA7e) q5bqCgTdxaTM.VyFk5uT {
	return &LRlqOaYmO{EA_5pkz: config, eyq32s: new(wd42zQ0FT.NG0hCZyRHK)}
}

// GetNode 对应 0x17a78e0。
func (p *LRlqOaYmO) GetNode() (*LFWTTXF7JR.Wr5SVP, error) {
	node, err := p.GetUniNode()
	if err != nil {
		return nil, err
	}
	return node.ToPanelNode(p.EA_5pkz.Yf1KqSqJpNI7)
}

// GetAPIData 对应 0x17a7940；两个闭包分别按 0x17afd00.asm 与
// 0x17afc60.asm 恢复。保留并发执行、部分结果及 SS 密钥转换的原行为。
func (p *LRlqOaYmO) GetAPIData() (*q5bqCgTdxaTM.D5gTIZ5aVg, error) {
	out := new(q5bqCgTdxaTM.D5gTIZ5aVg)
	err := HxKViEp2a.D5aDHlQDLp(
		func() error {
			node, err := p.GetUniNode()
			if err != nil {
				return err
			}
			converted, err := node.ToPanelNode(p.EA_5pkz.Yf1KqSqJpNI7)
			if err != nil {
				return err
			}
			out.H47AzZsYHB1N = converted
			out.HYnfYWa = node.GetDetectors()
			out.Xdo1D9 = time.Duration(node.Pa8W2AZ6Iw.F2WCOEc5Eu6.Value()) * time.Second
			out.DDu2Nm_P = time.Duration(node.Pa8W2AZ6Iw.I8kigte.Value()) * time.Second
			return nil
		},
		func() error {
			users, err := p.GetUsersResult()
			if err != nil {
				return err
			}
			out.VTme1SuHL6i_ = users.ToPanelUsers(p.EA_5pkz.ZyAhkq5th8bf)
			return nil
		},
	)
	if err == nil && p.EA_5pkz.ZyAhkq5th8bf == "ss" {
		node := out.H47AzZsYHB1N
		size, keyErr := node.FnMUny.HMkc4BgR.KeyLen()
		if keyErr != nil {
			return nil, fmt.Errorf("ss node err: %v", keyErr)
		}
		for _, user := range out.VTme1SuHL6i_ {
			credential, ok := user.JAfXWwzYrW.(*LFWTTXF7JR.YlfOV0DBBnx)
			if !ok {
				continue
			}
			credential.T1oDDIgYJHh = node.FnMUny.HMkc4BgR
			if (node.FnMUny.HMkc4BgR == "2022-blake3-aes-128-gcm" || node.FnMUny.HMkc4BgR == "2022-blake3-aes-256-gcm") && len(credential.A_wNUTkMh6pH) >= size {
				credential.A_wNUTkMh6pH = base64.StdEncoding.EncodeToString([]byte(credential.A_wNUTkMh6pH[:size]))
			}
		}
	}
	return out, err
}
