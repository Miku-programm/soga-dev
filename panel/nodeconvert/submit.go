package nodeconvert

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log/slog"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	Log "recovered.local/soga/util/logx"
	Q "recovered.local/soga/panel/registry"
)

func (p *DrEoaRCt2) ijuK35N2(in L.CWfcUv) error { //0x175cf80
	out := make(IvwLwWM, 0, len(in))
	for _, v := range in {
		out = append(out, &EMisRMdmr8d{v.Hy7Q97RROK, v.XJkGTQf_qkAt, v.I6LXObODav})
	}
	return p.rQrimKF088a("submit", out)
}
func (p *DrEoaRCt2) v2oTvqP(in L.CQhZrJ4dK) error { //0x175d120
	out := make(X22YvSZk, 0, len(in))
	for _, v := range in {
		for _, ip := range v.PNtn9KAV {
			out = append(out, &ZbwhxuO8Dq9a{v.Vy9ybu, ip})
		}
	}
	return p.rQrimKF088a("onlineusers", out)
}
func (p *DrEoaRCt2) bOqpKa(in L.QmiwDVjV00lZ) error { //0x175d9e0
	out := make(Ga4NII7c12I, 0, len(in))
	for _, v := range in {
		out = append(out, &EWcuZ1Kttay{v.GbzQUT_4f})
	}
	return p.rQrimKF088a("illegal", out)
}

// t9A_Zt2gOQQ restores175d3c0. Read failures are warned, remaining fields are
// still collected and posted. The original never populates the net field.
func (p *DrEoaRCt2) t9A_Zt2gOQQ() error {
	out := new(Vg6Nm38GB9)
	if v, e := cpu.Percent(0, false); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get cpu percent failed:", e)
	} else {
		out.EUjZVPy7nt = fmt.Sprintf("%.0f%%", v[0])
	}
	if v, e := mem.VirtualMemory(); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get virtual memory failed:", e)
	} else {
		out.Fpzi5hx = fmt.Sprintf("%.0f%%", 100*(float64(v.Used)/float64(v.Total)))
	}
	if v, e := disk.Usage("/"); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get dist usage failed:", e)
	} else {
		out.JzgG45gXrE = fmt.Sprintf("%.0f%%", float64(v.Used)/float64(v.Total)*100)
	}
	if v, e := host.Uptime(); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get uptime failed:", e)
	} else {
		out.PJJnV8 = v
	}
	return p.rQrimKF088a("nodestatus", out)
}

// SubmitData restores175db60 and its four actual concurrent task closures.
func (p *DrEoaRCt2) SubmitData(in *Q.FGmNba4MhioY) error {
	return H.D5aDHlQDLp(func() error { return p.ijuK35N2(in.HTgPyE3VytZd) }, func() error { return p.v2oTvqP(in.JNcHHwTZ7) }, p.t9A_Zt2gOQQ, func() error { return p.bOqpKa(in.Y0ys4ma4Tno4) })
}
