package uazGa0Ap

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log/slog"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	Log "recovered.local/soga/muc279Mpa6S4"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

func (p *W3tZSymoah) kPwWuozWyN(in L.CWfcUv) error { //0x1751380
	out := make(XsPC72G_, 0, len(in))
	for _, v := range in {
		out = append(out, &Hjb4HGlua{v.Hy7Q97RROK, v.XJkGTQf_qkAt, v.I6LXObODav})
	}
	return p.cpDvcLHkAY("submit", out)
}
func (p *W3tZSymoah) boFGTijR(in L.CQhZrJ4dK) error { //0x1751520
	out := make(FbR_38UK4A4, 0, len(in))
	for _, v := range in {
		for _, ip := range v.PNtn9KAV {
			out = append(out, &FW_R_A{v.Vy9ybu, ip})
		}
	}
	return p.cpDvcLHkAY("onlineusers", out)
}
func (p *W3tZSymoah) mF8pKFA(in L.QmiwDVjV00lZ) error { //0x1751de0
	out := make(F97hByMs8, 0, len(in))
	for _, v := range in {
		out = append(out, &GFVixRJI2{v.GbzQUT_4f})
	}
	return p.cpDvcLHkAY("illegal", out)
}

// qDXjmtmuX restores17517c0. Read failures are warned, remaining fields are
// still collected and posted. The original never populates the net field.
func (p *W3tZSymoah) qDXjmtmuX() error {
	out := new(MKWGBV)
	if v, e := cpu.Percent(0, false); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get cpu percent failed:", e)
	} else {
		out.MwG7GAHIuy = fmt.Sprintf("%.0f%%", v[0])
	}
	if v, e := mem.VirtualMemory(); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get virtual memory failed:", e)
	} else {
		out.A6pOwZFaGw = fmt.Sprintf("%.0f%%", 100*(float64(v.Used)/float64(v.Total)))
	}
	if v, e := disk.Usage("/"); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get dist usage failed:", e)
	} else {
		out.AazLqAvDDU = fmt.Sprintf("%.0f%%", float64(v.Used)/float64(v.Total)*100)
	}
	if v, e := host.Uptime(); e != nil {
		Log.TRTbrquD(slog.LevelWarn, "get uptime failed:", e)
	} else {
		out.FmIi0I4dOmz0 = v
	}
	return p.cpDvcLHkAY("nodestatus", out)
}

// SubmitData restores1751f60 and its four actual concurrent task closures.
func (p *W3tZSymoah) SubmitData(in *Q.FGmNba4MhioY) error {
	return H.D5aDHlQDLp(func() error { return p.kPwWuozWyN(in.HTgPyE3VytZd) }, func() error { return p.boFGTijR(in.JNcHHwTZ7) }, p.qDXjmtmuX, func() error { return p.mF8pKFA(in.Y0ys4ma4Tno4) })
}
