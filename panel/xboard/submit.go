// 根据 0x17928c0、0x1792a20、0x1792bc0、0x1792fe0 及其闭包恢复。
package xboard

import (
	"context"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"recovered.local/soga/util/errutil"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/panel/registry"
	"strconv"
)

func (p *KlO65BuwOt) iyqZ1OyNhjmI(traffic types.CWfcUv) error {
	if len(traffic) == 0 {
		return nil
	}
	address := p.rTN7wPdD("push")
	values := make(map[string][2]int64, len(traffic))
	for _, entry := range traffic {
		values[strconv.FormatInt(entry.Hy7Q97RROK, 10)] = [2]int64{entry.XJkGTQf_qkAt, entry.I6LXObODav}
	}
	return p.iJrldZ0g(address, values)
}
func (p *KlO65BuwOt) cCw6i_L9ERk(alive types.CQhZrJ4dK) error {
	if len(alive) == 0 {
		return nil
	}
	address := p.rTN7wPdD("alive")
	values := make(map[string][]string, len(alive))
	for _, entry := range alive {
		values[strconv.FormatInt(entry.Vy9ybu, 10)] = entry.PNtn9KAV
	}
	return p.iJrldZ0g(address, values)
}

// gKXqCyN 保留原依次采样、错误时记录并继续及 CPU 空切片越界行为。
// gopsutil v2.20.9 为已核对所用 API/字段布局的兼容版本，原补丁版本未证明。
func (p *KlO65BuwOt) gKXqCyN() error {
	out := new(ISasilPp)
	ctx := context.Background()
	if value, err := cpu.PercentWithContext(ctx, 0, false); err != nil {
		logx.TRTbrquD(4, "get cpu percent failed:", err)
	} else {
		out.ArsVcdZv7Va = value[0]
	}
	if value, err := mem.VirtualMemory(); err != nil {
		logx.TRTbrquD(4, "get virtual memory failed:", err)
	} else {
		out.FntmTKM6E.BcEb4PnfTS9 = value.Used
		out.FntmTKM6E.JvmfpVk = value.Total
	}
	if value, err := mem.SwapMemoryWithContext(ctx); err != nil {
		logx.TRTbrquD(4, "get swap memory failed:", err)
	} else {
		out.AqVaAjc1cqg.BcEb4PnfTS9 = value.Used
		out.AqVaAjc1cqg.JvmfpVk = value.Total
	}
	if value, err := disk.UsageWithContext(ctx, "/"); err != nil {
		logx.TRTbrquD(4, "get disk usage failed:", err)
	} else {
		out.WtY6pUI_1.BcEb4PnfTS9 = value.Used
		out.WtY6pUI_1.JvmfpVk = value.Total
	}
	return p.iJrldZ0g(p.rTN7wPdD("status"), out)
}

// SubmitData 始终上报状态，nil data 只跳过流量及在线 IP 两项。
func (p *KlO65BuwOt) SubmitData(data *registry.FGmNba4MhioY) error {
	return errutil.D5aDHlQDLp(
		func() error {
			if data != nil {
				return p.iyqZ1OyNhjmI(data.HTgPyE3VytZd)
			}
			return nil
		},
		func() error {
			if data != nil {
				return p.cCw6i_L9ERk(data.JNcHHwTZ7)
			}
			return nil
		},
		func() error { return p.gKXqCyN() },
	)
}
