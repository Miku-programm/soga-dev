package D43Y3mB

import (
	L "recovered.local/soga/LFWTTXF7JR"
	"strings"
)

// SubmitTraffic 对应 0x16bb1e0。空列表不发请求。
// D43Y3mB.(*GMLWI3).SubmitTraffic：原样本入口 0x16bb1e0。
func (p *GMLWI3) SubmitTraffic(address string, traffic L.CWfcUv) error {
	if len(traffic) == 0 {
		return nil
	}
	out := make([]*ymId_GfJFS, 0)
	for _, v := range traffic {
		out = append(out, &ymId_GfJFS{W0KlV7rEtxgz: v.Hy7Q97RROK, QavUxYiaO: v.XJkGTQf_qkAt, Vdf1KaaiPt: v.I6LXObODav})
	}
	return p.PostCommon(address, out)
}

// SubmitAliveIP 对应 0x16bb3c0；同用户IP以逗号连接，保留输入顺序。
// D43Y3mB.(*GMLWI3).SubmitAliveIP：原样本入口 0x16bb3c0。
func (p *GMLWI3) SubmitAliveIP(address string, ips L.CQhZrJ4dK) error {
	if len(ips) == 0 {
		return nil
	}
	out := make([]*ucFGrGVJeia, 0)
	for _, v := range ips {
		out = append(out, &ucFGrGVJeia{M9gh81OOM4ca: v.Vy9ybu, A6Zb5j5: strings.Join(v.PNtn9KAV, ",")})
	}
	return p.PostCommon(address, out)
}

// SubmitDetectLog 对应 0x16bb5e0/0x16c0380；逐项异步提交，原实现忽略请求错误。
// D43Y3mB.(*GMLWI3).SubmitDetectLog：原样本入口 0x16bb5e0。
func (p *GMLWI3) SubmitDetectLog(address string, logs L.QmiwDVjV00lZ) error {
	for _, v := range logs {
		r := &ksuCP0Hi{CYhv312VDo: v.GbzQUT_4f, F5beiROWvEE: v.CAgUUdrv9, L6tA4h_: "无"}
		go p.PostCommon(address, r)
	}
	return nil
}
