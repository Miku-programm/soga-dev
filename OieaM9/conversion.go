// 依据原 ELF ToPanelNode/GetUser/GetUsers 的 C 与 ASM 逐分支恢复。
package OieaM9

import (
	"errors"
	"fmt"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"strconv"
)

// 原地址 0x378731A 初始为 1；GetUsers 尾部写零，保留原一次性行为。
var firstUsersAt378731A = true

// 下列局部辅助表达原内联限速/连接数选择，包含负数与 NaN 分支；不是额外原函数。
func speedLimit(a, b float64) float64 {
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
		if b <= a {
			a = b
		}
	} else if !(a > b) {
		a = b
	}
	return a
}

func connLimit(a, b int) int {
	if a < 0 {
		a = 0
	}
	if b < 0 {
		b = 0
	}
	if a <= 0 || b <= 0 {
		if b < a {
			b = a
		}
	} else if b > a {
		b = a
	}
	return b
}

// OieaM9.(*CXFPuaTkSm).GetUser：原样本入口 0x16c21e0。
func (p *CXFPuaTkSm) GetUser(index int) *L.Aga91y {
	out := &L.Aga91y{WZ89pTAOrHK4: p.QZuKzrg, LmOKwjha: speedLimit(p.T1hUUd2, C.PanelUserSpeedLimit()), Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit()}
	out.JAfXWwzYrW = &L.QcqpSL3i6ri{YrjVl_2g_S: p.SXYRZj, YyAgnMatmj: p.JaWN6wQMY_8k}
	if !firstUsersAt378731A && C.JatfUF38O() && index >= 88 {
		out.JAfXWwzYrW = nil
	}
	return out
}

// OieaM9.(*JpYAk64U2).GetUsers：原样本入口 0x16c23c0。
func (p *JpYAk64U2) GetUsers() L.Sxcf3wu3L {
	if !firstUsersAt378731A && C.Q0SDIN() && len(p.GPETCf2) > 88 {
		p.GPETCf2 = p.GPETCf2[:88]
	}
	out := make(L.Sxcf3wu3L, 0)
	for i, u := range p.GPETCf2 {
		out = append(out, u.GetUser(i))
	}
	firstUsersAt378731A = false
	return out
}

// OieaM9.(*QVdbsbqVT3).ToPanelNode：原样本入口 0x16c0580。
func (p *QVdbsbqVT3) ToPanelNode(id int) (*L.Wr5SVP, error) {
	if p.C8gP2P == 0 {
		return nil, errors.New("请在 ssr 节点中启用单端口")
	}
	port, err := strconv.Atoi(p.AZlE6hE7M7)
	if err != nil {
		return nil, fmt.Errorf("port invalid: %v", port)
	}
	out := &L.Wr5SVP{EMlXEc: id, Z1TIgHOQ: Q.T0j_Rco(p), HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: speedLimit(p.DR80QsI, C.PanelNodeSpeedLimit()), W3Dw_B4wgbM: connLimit(p.K9N3GWPMdM3, C.PanelUserConnectionLimit())}
	out.ASJxBbFS = L.Dv6QJyFW6bU{PU1Di18: port, Qggyovo: p.G4HJNQ1, D2rWMm0: p.DXDUf20, ND3AWUWg: p.SCImeBpi43, JvlcaS: p.UGE6Z0a, QawbKS2x: true}
	return out, nil
}
