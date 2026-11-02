// 依据原 ELF ToPanelNode/GetUser/GetUsers 的 C 与 ASM 逐分支恢复。
package whmcsnode

import (
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

// 原地址 0x378731B 初始为 1；GetUsers 尾部写零，保留原一次性行为。
var firstUsersAt378731B = true

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

// whmcsnode.(*PqISXuLk).GetUser：原样本入口 0x16c8f60。
func (p *PqISXuLk) GetUser(index int) *L.Aga91y {
	out := &L.Aga91y{WZ89pTAOrHK4: p.Jrq3niigv, LmOKwjha: speedLimit(p.Bexa45zkL, C.PanelUserSpeedLimit()), Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit()}
	out.JAfXWwzYrW = &L.JxTGUar{ARARP1z0yA: p.Jm6azqIw}
	return out
}

// whmcsnode.(*BoM67db).GetUsers：原样本入口 0x16c9140。
func (p *BoM67db) GetUsers() L.Sxcf3wu3L {
	out := make(L.Sxcf3wu3L, 0)
	for i, u := range p.GPETCf2 {
		out = append(out, u.GetUser(i))
	}
	firstUsersAt378731B = false
	return out
}

// whmcsnode.(*XhDo0vyZH).ToPanelNode：原样本入口 0x16c8da0。
func (p *XhDo0vyZH) ToPanelNode(id int) *L.Wr5SVP {
	out := &L.Wr5SVP{EMlXEc: id, Z1TIgHOQ: Q.T0j_Rco(p), HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: speedLimit(p.JLAJhSUZPpQv, C.PanelNodeSpeedLimit()), W3Dw_B4wgbM: connLimit(p.JQsJVj, C.PanelUserConnectionLimit())}
	out.JhIYwEukjd = L.L9AiYC{GbD_2PU: p.Jbn82w0HZp, T4cuhG7MTbG: "tcp"}
	return out
}
