package webapi

import (
	"context"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	M "recovered.local/soga/util/logx"
	Q "recovered.local/soga/panel/registry"
	"time"
)

func (p *Zkt5w_2o) m0ZwxicZQFM3(alive L.CQhZrJ4dK) error {
	if len(alive) == 0 {
		return nil
	}
	users := make(QFeT8QtEgxS, 0, len(alive))
	for _, u := range alive {
		for _, ip := range u.PNtn9KAV {
			users = append(users, &LC4dqy{M9gh81OOM4ca: u.Vy9ybu, A6Zb5j5: ip})
		}
	}
	return p.wVyldOdHZ3gi("online", &struct {
		Gd8mcBP80d QFeT8QtEgxS `json:"users"`
	}{users})
}
func (p *Zkt5w_2o) vTdzQJ(traffic L.CWfcUv) error {
	if len(traffic) == 0 {
		return nil
	}
	values := make(IodE2Y, 0, len(traffic))
	for _, v := range traffic {
		values = append(values, &DjMeiycBVU{I4dYSFy5bqX: v.Hy7Q97RROK, QdZ1mt63MOLf: v.XJkGTQf_qkAt, BJNQsH4BpW2: v.I6LXObODav})
	}
	return p.wVyldOdHZ3gi("push", &struct {
		YDc9Teuc IodE2Y `json:"traffic"`
	}{values})
}
func (p *Zkt5w_2o) djhDIRaqzWaO() error {
	var status struct {
		VyfxnLbeHB float64 `json:"cpu"`
		ZsaTFsOg   float64 `json:"mem"`
		GjA4BtU    float64 `json:"disk"`
		IpyjSPcdt  int64   `json:"updated_at"`
	}
	status.IpyjSPcdt = time.Now().UnixMilli()
	ctx := context.Background()
	if v, e := cpu.PercentWithContext(ctx, 0, false); e != nil {
		M.TRTbrquD(4, "get cpu percent failed:", e)
	} else {
		status.VyfxnLbeHB = v[0]
	}
	if v, e := mem.VirtualMemory(); e != nil {
		M.TRTbrquD(4, "get virtual memory failed:", e)
	} else {
		status.ZsaTFsOg = v.UsedPercent
	}
	if v, e := disk.UsageWithContext(ctx, "/"); e != nil {
		M.TRTbrquD(4, "get disk usage failed:", e)
	} else {
		status.GjA4BtU = v.UsedPercent
	}
	return p.wVyldOdHZ3gi("status", status)
}

// 原 nil data 直接返回；非 nil 数据同时执行在线、流量和状态三路任务。
func (p *Zkt5w_2o) SubmitData(data *Q.FGmNba4MhioY) error {
	if data == nil {
		return nil
	}
	return H.D5aDHlQDLp(func() error { return p.m0ZwxicZQFM3(data.JNcHHwTZ7) }, func() error { return p.vTdzQJ(data.HTgPyE3VytZd) }, func() error { return p.djhDIRaqzWaO() })
}
