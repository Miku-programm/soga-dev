// 根据原样本 0x1747de0、0x1747f20 恢复提交方法。
package ZrjfjK8Y_tM9

import (
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/q5bqCgTdxaTM"
	"strconv"
)

// eALpAMr 对应 0x1747de0；同一用户后出现的记录覆盖前一条，并不累加。
func (p *JZA3kKksUe) eALpAMr(traffic LFWTTXF7JR.CWfcUv) error {
	if len(traffic) == 0 {
		return nil
	}
	address := p.GetUniUrl("push")
	values := make(map[string][2]int64)
	for _, entry := range traffic {
		values[strconv.FormatInt(entry.Hy7Q97RROK, 10)] = [2]int64{entry.XJkGTQf_qkAt, entry.I6LXObODav}
	}
	return p.PostCommon(address, values)
}

// SubmitData 对应 0x1747f20；原方法仅读取提交结构中的流量字段。
func (p *JZA3kKksUe) SubmitData(data *q5bqCgTdxaTM.FGmNba4MhioY) error {
	return p.eALpAMr(data.HTgPyE3VytZd)
}
