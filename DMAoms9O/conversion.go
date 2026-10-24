// 依据原 ELF ToPanelNode/GetUser/GetUsers 的 C 与 ASM 逐分支恢复。
package DMAoms9O

import (
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

// 原地址 0x378731C 初始为 1；GetUsers 尾部写零，保留原一次性行为。
var firstUsersAt378731C = true

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

// DMAoms9O.(*Pn5vgsrEm5).GetUser：原样本入口 0x16cf0c0。
func (p *Pn5vgsrEm5) GetUser(index int) *L.Aga91y {
	out := &L.Aga91y{WZ89pTAOrHK4: p.EgeihzZFfj, LmOKwjha: speedLimit(p.Z_Y4R2Rbg, C.PanelUserSpeedLimit()), Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit()}
	out.JAfXWwzYrW = &L.Yn8BJoL{SfiEGpJo: p.Z6TxVE4XB}
	if !firstUsersAt378731C && C.Q0SDIN() && index >= 88 {
		out.JAfXWwzYrW = nil
	}
	return out
}

// DMAoms9O.(*FaNycS0m).GetUsers：原样本入口 0x16cf2a0。
func (p *FaNycS0m) GetUsers() L.Sxcf3wu3L {
	if !firstUsersAt378731C && C.JatfUF38O() && len(p.GPETCf2) > 88 {
		p.GPETCf2 = p.GPETCf2[:88]
	}
	out := make(L.Sxcf3wu3L, 0)
	for i, u := range p.GPETCf2 {
		out = append(out, u.GetUser(i))
	}
	firstUsersAt378731C = false
	return out
}

// DMAoms9O.(*GDzADaHU).ToPanelNode：原样本入口 0x16cee80。
func (p *GDzADaHU) ToPanelNode(id int) *L.Wr5SVP {
	if p.A05Abab == "http" {
		p.A05Abab = "h2"
	}
	out := &L.Wr5SVP{EMlXEc: id, Z1TIgHOQ: Q.T0j_Rco(p), HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: speedLimit(p.WO4euQQs, C.PanelNodeSpeedLimit()), W3Dw_B4wgbM: connLimit(p.AzPzYuvK9, C.PanelUserConnectionLimit())}
	out.SwxJNgPpg = L.JW0CeHo3O{CKsOpayd: p.I6bjCC2H4sJm, VUwf1n9g: p.QgIbz8hHD0, F2GXyZZ: L.BoulTn31(p.A05Abab), UVeahqDBT: p.HaC47go2d8Dz, ChJnngzb: p.LVgd5PJKI}
	return out
}
