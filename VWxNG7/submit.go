// 原始提交路径：0x16da0a0、0x16da260、0x16da460、0x16da660、0x16dab20。
package VWxNG7

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	log "recovered.local/soga/muc279Mpa6S4"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

func (p *CsKk5VzGb) kiKw99B(in L.CWfcUv) error {
	if len(in) == 0 {
		return nil
	}
	out := make([]*OTsrnZKa, 0, len(in))
	for _, u := range in {
		out = append(out, &OTsrnZKa{X7cdH6hK: u.Hy7Q97RROK, SszZxQ_BV: u.XJkGTQf_qkAt, CEKdEFoFV: u.I6LXObODav})
	}
	return p.olDPHF3j("traffic", out)
}
func (p *CsKk5VzGb) h1d1_Y(in L.CQhZrJ4dK) error {
	if len(in) == 0 {
		return nil
	}
	out := make([]*Tw_nfz, 0, len(in))
	for _, u := range in {
		out = append(out, &Tw_nfz{AdYIli: u.Vy9ybu, FmXlcnZSlSU: u.PNtn9KAV})
	}
	return p.olDPHF3j("alive_ip", out)
}
func (p *CsKk5VzGb) yoBIG34bwknA(in L.QmiwDVjV00lZ) error {
	if len(in) == 0 {
		return nil
	}
	out := make([]*V91695M, 0, len(in))
	for _, u := range in {
		out = append(out, &V91695M{VJB4pQGJ: u.GbzQUT_4f, JtAQyC57KMKb: u.CAgUUdrv9})
	}
	return p.olDPHF3j("audit_log", out)
}
func (p *CsKk5VzGb) x2xbc1KY() error {
	out := new(LygoL0kC5)
	if v, e := cpu.Percent(0, false); e != nil {
		log.TRTbrquD(4, "get cpu percent failed:", e)
	} else {
		out.BRhjXec = v[0]
	}
	if v, e := mem.VirtualMemory(); e != nil {
		log.TRTbrquD(4, "get virtual memory failed:", e)
	} else {
		out.ZmWD5iP.BcEb4PnfTS9 = v.Used
		out.ZmWD5iP.JvmfpVk = v.Total
	}
	if v, e := mem.SwapMemory(); e != nil {
		log.TRTbrquD(4, "get swap memory failed:", e)
	} else {
		out.J9mVD1.BcEb4PnfTS9 = v.Used
		out.J9mVD1.JvmfpVk = v.Total
	}
	if v, e := disk.Usage("/"); e != nil {
		log.TRTbrquD(4, "get disk usage failed:", e)
	} else {
		out.EpTvgGsJ5.BcEb4PnfTS9 = v.Used
		out.EpTvgGsJ5.JvmfpVk = v.Total
	}
	if v, e := host.Uptime(); e != nil {
		log.TRTbrquD(4, "get uptime failed:", e)
	} else {
		out.EUw7Qwoa4K0 = v
	}
	return p.olDPHF3j("status", out)
}
func (p *CsKk5VzGb) SubmitData(in *Q.FGmNba4MhioY) error {
	return HxKViEp2a.D5aDHlQDLp(
		func() error {
			if in == nil {
				return nil
			}
			return p.kiKw99B(in.HTgPyE3VytZd)
		},
		func() error {
			if in == nil {
				return nil
			}
			return p.h1d1_Y(in.JNcHHwTZ7)
		},
		func() error {
			if in == nil {
				return nil
			}
			return p.yoBIG34bwknA(in.Y0ys4ma4Tno4)
		},
		func() error { return p.x2xbc1KY() },
	)
}
