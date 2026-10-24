package _U2caaPPsm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	netstat "github.com/shirou/gopsutil/net"
	"net/url"
	G "recovered.local/soga/G6bHD5JA"
	L "recovered.local/soga/lm_zcVAb"
	M "recovered.local/soga/muc279Mpa6S4"
	"time"
)

// 原 RTTI 160 字节；累计计数和采样间差值分别保留。
type DpeSn3T struct {
	ABamh2lGc                                                              time.Time
	Ext49lvSw                                                              uint64
	CUcRQ2                                                                 float64
	Z9XlEaJL, Iqf1PJb9gAuJ, ND2QesObp, ZoKBTLsJbb, MyXxXp12C, BMQhQ6It6ocB uint64
	SOaNHKQg3a, ZiFazv7, XR83N6vHA                                         float64
	DrY4mmyOP, IMkIJFzv, Rl5I6SPhu50u, XixSqS4zM6                          uint64
	OEcQalt, PaNuEBV208                                                    int
}

// aYdnSsv7yZ 对应 0x169f8c0。采样失败仍继续；原 CPU 空结果会越界。
func (p *Iw4id8kfTTT) aYdnSsv7yZ() *DpeSn3T {
	out := &DpeSn3T{ABamh2lGc: time.Now()}
	ctx := context.Background()
	if v, e := cpu.PercentWithContext(ctx, 0, false); e != nil {
		M.TRTbrquD(4, "get cpu percent failed:", e)
	} else {
		out.CUcRQ2 = v[0]
	}
	if v, e := host.UptimeWithContext(ctx); e != nil {
		M.TRTbrquD(4, "get uptime failed:", e)
	} else {
		out.Ext49lvSw = v
	}
	if v, e := mem.VirtualMemory(); e != nil {
		M.TRTbrquD(4, "get virtual memory failed:", e)
	} else {
		out.Z9XlEaJL = v.Used
		out.Iqf1PJb9gAuJ = v.Total
	}
	if v, e := mem.SwapMemoryWithContext(ctx); e != nil {
		M.TRTbrquD(4, "get swap memory failed:", e)
	} else {
		out.ND2QesObp = v.Used
		out.ZoKBTLsJbb = v.Total
	}
	if v, e := disk.UsageWithContext(ctx, "/"); e != nil {
		M.TRTbrquD(4, "get dist usage failed:", e)
	} else {
		out.MyXxXp12C = v.Used
		out.BMQhQ6It6ocB = v.Total
	}
	if v, e := load.AvgWithContext(ctx); e != nil {
		M.TRTbrquD(4, "get load avg failed:", e)
	} else {
		out.SOaNHKQg3a = v.Load1
		out.ZiFazv7 = v.Load5
		out.XR83N6vHA = v.Load15
	}
	if v, e := netstat.IOCountersWithContext(ctx, false); e != nil {
		M.TRTbrquD(4, "get io counters failed:", e)
	} else if len(v) == 0 {
		M.TRTbrquD(4, "can not find io counters")
	} else {
		out.Rl5I6SPhu50u = v[0].BytesSent
		out.XixSqS4zM6 = v[0].BytesRecv
		if previous := p.tdXPhll10; previous != nil {
			seconds := out.ABamh2lGc.Sub(previous.ABamh2lGc).Seconds()
			out.DrY4mmyOP = uint64(float64(out.Rl5I6SPhu50u-previous.Rl5I6SPhu50u) / seconds)
			out.IMkIJFzv = uint64(float64(out.XixSqS4zM6-previous.XixSqS4zM6) / seconds)
		}
	}
	if v, e := L.QMgUaHjx4G(); e != nil {
		M.TRTbrquD(4, "get connections failed:", e)
	} else {
		out.OEcQalt = v
	}
	if v, e := L.NEfQI6r99p8(); e != nil {
		M.TRTbrquD(4, "get connections failed:", e)
	} else {
		out.PaNuEBV208 = v
	}
	p.tdXPhll10 = out
	return out
}

// SubmitInfo 对应 0x16a0360；只有状态上报要求 HTTP 200。
func (p *Iw4id8kfTTT) SubmitInfo(s *DpeSn3T) error {
	data := map[string]any{
		"v2": map[string]any{"code": 0}, "uptime": s.Ext49lvSw, "cpu": map[string]any{"percent": s.CUcRQ2},
		"memory":      map[string]any{"used": s.Z9XlEaJL, "total": s.Iqf1PJb9gAuJ},
		"swap":        map[string]any{"used": s.ND2QesObp, "total": s.ZoKBTLsJbb},
		"disk":        map[string]any{"used": s.MyXxXp12C, "total": s.BMQhQ6It6ocB},
		"loads":       []float64{s.SOaNHKQg3a, s.ZiFazv7, s.XR83N6vHA},
		"net_io":      map[string]any{"up": s.DrY4mmyOP, "down": s.IMkIJFzv},
		"net_traffic": map[string]any{"sent": s.Rl5I6SPhu50u, "recv": s.XixSqS4zM6},
		"tcp_count":   s.OEcQalt, "udp_count": s.PaNuEBV208,
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, status, err := G.J1liFniHg(p.jamJA0Qs.W_ogGtF+"sub_server/submit_info", url.Values{"secret": {p.jamJA0Qs.Bd08Li}, "status": {string(body)}})
	if err != nil || status != 200 {
		return fmt.Errorf("submit status failed: %v, status code: %v", err, status)
	}
	return nil
}
func (p *Iw4id8kfTTT) SubmitStatus() error { return p.SubmitInfo(p.aYdnSsv7yZ()) }
