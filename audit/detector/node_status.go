package detector

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	Log "recovered.local/soga/util/logx"
	"time"
)

// SubmitNodeInfo 对应 0x16ba720。保留原500ms CPU采样、逐项失败默认值及上报格式。
// detector.(*GMLWI3).SubmitNodeInfo：原样本入口 0x16ba720。
func (p *GMLWI3) SubmitNodeInfo(address string) error {
	uptime, err := host.Uptime()
	if err != nil {
		Log.TRTbrquD(4, "get uptime err:", err)
	}
	cpus, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		Log.TRTbrquD(4, "get cpu percent err:", err)
		cpus = []float64{0}
	}
	memory, err := mem.VirtualMemory()
	if err != nil {
		Log.TRTbrquD(4, "get mem percent err:", err)
		memory = &mem.VirtualMemoryStat{Total: 1}
	}
	counters, err := net.IOCounters(false)
	if err != nil {
		Log.TRTbrquD(4, "get io err:", err)
		counters = []net.IOCountersStat{{}}
	}
	usage, err := disk.Usage("/")
	if err != nil {
		Log.TRTbrquD(4, "get disk err:", err)
		usage = &disk.UsageStat{}
	}
	sent := float64(counters[0].BytesSent) / 1024 / 1024 / 1024
	received := float64(counters[0].BytesRecv) / 1024 / 1024 / 1024
	result := map[string]any{"cpu": fmt.Sprintf("%.2f%%", cpus[0]), "mem": fmt.Sprintf("%.2f%%", float64(memory.Used)/float64(memory.Total)*100), "net": fmt.Sprintf("%.2fGB↑ %.2fGB↓", sent, received), "disk": fmt.Sprintf("%.2f%%", usage.UsedPercent), "uptime": uptime}
	return p.PostCommon(address, result)
}
