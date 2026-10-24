package sspanel

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	U "recovered.local/soga/panel/sspaneltypes"
	M "recovered.local/soga/util/logx"
	Q "recovered.local/soga/panel/registry"
)

// 与原 callsites 一致的宿主采样；变量仅使本地失败路径可确定测试。
var hostLoad = load.Avg
var hostUptime = host.Uptime

func (p *JYkPXk4ZdW) SubmitNodeInfo() error {
	avg, err := hostLoad()
	if avg == nil {
		M.TRTbrquD(4, "get load avg err:", err)
		avg = &load.AvgStat{}
	}
	uptime, err := hostUptime()
	if err != nil {
		M.TRTbrquD(4, "get uptime err:", err)
	}
	data := map[string]any{"load": fmt.Sprintf("%.2f %.2f %.2f", avg.Load1, avg.Load5, avg.Load15), "uptime": uptime}
	// 原 0x172d36e 丢弃 POST 错误，返回 nil。
	_ = p.PostCommon(fmt.Sprintf("nodes/%v/info", p.EA_5pkz.Yf1KqSqJpNI7), data)
	return nil
}
func (p *JYkPXk4ZdW) SubmitTraffic(data L.CWfcUv) error {
	if len(data) == 0 {
		return nil
	}
	return p.PostCommon("users/traffic", map[string]any{"data": data})
}
func (p *JYkPXk4ZdW) SubmitAliveIP(data L.CQhZrJ4dK) error {
	if len(data) == 0 {
		return nil
	}
	out := make([]*U.BQJu1p, 0, len(data))
	for _, u := range data {
		for _, ip := range u.PNtn9KAV {
			out = append(out, &U.BQJu1p{G1PPkdYRT: u.Vy9ybu, AWUmSo: ip})
		}
	}
	return p.PostCommon("users/aliveip", map[string]any{"data": out})
}
func (p *JYkPXk4ZdW) SubmitDetectLog(data L.QmiwDVjV00lZ) error {
	if len(data) == 0 {
		return nil
	}
	out := make([]*U.VkGzHpZbP7, 0, len(data))
	for _, d := range data {
		out = append(out, &U.VkGzHpZbP7{FRwqa_GCLS: d.GbzQUT_4f, AezaFx: d.CAgUUdrv9})
	}
	body := map[string]any{"data": out}
	return p.PostCommon("users/detectlog", &body)
}
func (p *JYkPXk4ZdW) SubmitData(data *Q.FGmNba4MhioY) error {
	return H.D5aDHlQDLp(p.SubmitNodeInfo, func() error { return p.SubmitTraffic(data.HTgPyE3VytZd) }, func() error { return p.SubmitDetectLog(data.Y0ys4ma4Tno4) }, func() error { return p.SubmitAliveIP(data.JNcHHwTZ7) })
}
