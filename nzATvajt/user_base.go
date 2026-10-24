package nzATvajt

import (
	"context"
	"fmt"
	"sync/atomic"

	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/tkbUwvW3ex"
)

// AhSFeSua 保留原用户基类的字段顺序及 amd64 上 152 字节布局。
type AhSFeSua struct {
	dcM4pRx3d0S context.Context
	kI9WvUIev   context.CancelFunc
	djdKrU      int64
	s08sSK0Icd  float64
	i_fZvtng2iy int
	ckKNgBSf    int
	hBZll21     int64
	uIawOCHn    int64
	pq4dsrgDKl9 int64
	aOJHFe7     *G6bHD5JA.L337vk2G
	ipvHjTY     *G6bHD5JA.L337vk2G
	tKYtJcjS    *G6bHD5JA.L337vk2G
	p1GKeQ9QSy  *tkbUwvW3ex.JW2SKH2ajwT6
	rQTIvZ      atomic.Int32
	pSde23G4D   atomic.Int64
	up4CJo      atomic.Int64
	alKKtPSK    atomic.Int64
	mVxQSdX     atomic.Int64
}

var _ PgaNd3014U = (*AhSFeSua)(nil)

// cLLKJ_A31 @ 0x1825460。原构造不初始化 TCP 限额，也不启用用户限速；
// 具体协议构造随后调用 UpdatePanelUser/UpdateUser 完成这两项。
func cLLKJ_A31(p *LFWTTXF7JR.Aga91y, node *G6bHD5JA.L337vk2G) *AhSFeSua {
	ctx, cancel := context.WithCancel(context.Background())
	return &AhSFeSua{dcM4pRx3d0S: ctx, kI9WvUIev: cancel, djdKrU: p.WZ89pTAOrHK4,
		s08sSK0Icd: p.LmOKwjha, i_fZvtng2iy: p.Xmrna8, hBZll21: p.XuBBcTThBW,
		uIawOCHn: p.UWsYsmytTi, pq4dsrgDKl9: p.IUlI_15Qf_, aOJHFe7: node,
		ipvHjTY: new(G6bHD5JA.L337vk2G), tKYtJcjS: new(G6bHD5JA.L337vk2G),
		p1GKeQ9QSy: tkbUwvW3ex.NewJW2SKH2ajwT6()}
}

func (u *AhSFeSua) UpdateUser(p *LFWTTXF7JR.Aga91y) { // 0x1825680
	u.s08sSK0Icd = p.LmOKwjha
	u.i_fZvtng2iy = p.Xmrna8
	u.ckKNgBSf = p.B4DYSxuA31za
	u.hBZll21 = p.XuBBcTThBW
	u.uIawOCHn = p.UWsYsmytTi
	u.pq4dsrgDKl9 = p.IUlI_15Qf_
	u.ibGRzBUe()
}
func (u *AhSFeSua) ibGRzBUe() { // 0x1825700
	var n int64
	// 原 ucomisd/jb 在 NaN 时亦进入转换分支，不能改成 speed > 0。
	if !(u.s08sSK0Icd <= 0) {
		n = int64(1024.0 * (0.125 * (u.s08sSK0Icd * 1024.0)))
	}
	u.ipvHjTY.UpdateLimit(n)
	u.tKYtJcjS.UpdateLimit(n)
}
func (u *AhSFeSua) GetId() int64      { return u.djdKrU }      // 0x18257c0
func (u *AhSFeSua) GetConnLimit() int { return u.i_fZvtng2iy } // 0x18257e0
func (u *AhSFeSua) AddUp(n int) { // 0x1825800
	u.pSde23G4D.Add(int64(n))
	u.alKKtPSK.Add(int64(n))
	if u.aOJHFe7 != nil {
		u.aOJHFe7.Add(n)
	}
	u.ipvHjTY.Add(n)
}
func (u *AhSFeSua) AllowAndAddUp(n int) bool { // 0x1825880
	u.pSde23G4D.Add(int64(n))
	u.alKKtPSK.Add(int64(n))
	return u.tKYtJcjS.AllowAndAdd(n)
}
func (u *AhSFeSua) AddDown(n int) { // 0x18258e0
	u.up4CJo.Add(int64(n))
	u.mVxQSdX.Add(int64(n))
	if u.aOJHFe7 != nil {
		u.aOJHFe7.Add(n)
	}
	u.ipvHjTY.Add(n)
}
func (u *AhSFeSua) AllowAndAddDown(n int) bool { // 0x1825960
	u.up4CJo.Add(int64(n))
	u.mVxQSdX.Add(int64(n))
	return u.tKYtJcjS.AllowAndAdd(n)
}
func (u *AhSFeSua) GetAndClearTraffic() (int64, int64) { // 0x18259e0
	up := u.pSde23G4D.Swap(0)
	down := u.up4CJo.Swap(0)
	return up, down
}
func (u *AhSFeSua) GetAndClearDyTraffic() (int64, int64) { // 0x1825a00
	up := u.alKKtPSK.Swap(0)
	down := u.mVxQSdX.Swap(0)
	return up, down
}
func (u *AhSFeSua) IsTrafficOut() bool { // 0x1825a20
	return u.pq4dsrgDKl9 > 0 && u.pq4dsrgDKl9 <= u.hBZll21+u.uIawOCHn+u.pSde23G4D.Load()+u.up4CJo.Load()
}
func (u *AhSFeSua) AddAliveIp(ip, kind string) bool {
	return u.p1GKeQ9QSy.AddCount(ip, kind, u.i_fZvtng2iy)
}                                              // 0x1825a60
func (u *AhSFeSua) DelAliveIp(ip, kind string) { u.p1GKeQ9QSy.SubCount(ip, kind) } // 0x1825ae0
func (u *AhSFeSua) RecordAliveIp(ip string)    { u.p1GKeQ9QSy.Record(ip) }         // 0x1825b60
func (u *AhSFeSua) GetAliveIpData() *LFWTTXF7JR.LmNKoY94 { // 0x1825bc0
	ips := u.p1GKeQ9QSy.GetAndClearAliveIps()
	if len(ips) == 0 {
		return nil
	}
	return &LFWTTXF7JR.LmNKoY94{Vy9ybu: u.djdKrU, PNtn9KAV: ips}
}
func (u *AhSFeSua) AddTCP() bool { // 0x1825c80，原检查与加计数不是单个 CAS。
	if u.ckKNgBSf > 0 && u.ckKNgBSf <= int(u.rQTIvZ.Load()) {
		return false
	}
	u.rQTIvZ.Add(1)
	return true
}
func (u *AhSFeSua) DelTCP() { u.rQTIvZ.Add(-1) } // 0x1825cc0
func (u *AhSFeSua) ResetCtx() { // 0x1825ce0
	if u.kI9WvUIev != nil {
		u.kI9WvUIev()
	}
	u.dcM4pRx3d0S, u.kI9WvUIev = context.WithCancel(context.Background())
}
func (u *AhSFeSua) GetCtx() context.Context { return u.dcM4pRx3d0S }             // 0x1825d80
func (u *AhSFeSua) GetSpeedLimit() float64  { return u.ipvHjTY.GetSpeedLimit() } // 0x1825da0
func (u *AhSFeSua) SetSpeedLimit(speed float64) { // 0x1825de0
	if speed <= 0 {
		u.ibGRzBUe()
		return
	}
	n := int64(1024.0 * (0.125 * (speed * 1024.0)))
	u.ipvHjTY.UpdateLimit(n)
	u.tKYtJcjS.UpdateLimit(n)
}
func pNF7CSleWS(id int64, protocol string) error { // 0x1825e80
	return fmt.Errorf("user %v has no %v credential", id, protocol)
}
