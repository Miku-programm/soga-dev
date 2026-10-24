// 根据原样本 0x17a7de0、0x17a8080 恢复提交方法。
package xiaov2board

import (
	"recovered.local/soga/util/errutil"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/panel/registry"
	"strconv"
)

// zsD02JRnj 对应 0x17a7de0；同一用户后出现的记录覆盖前一条，并不累加。
func (p *LRlqOaYmO) zsD02JRnj(traffic types.CWfcUv) error {
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

// SubmitData 对应 0x17a8080，两项并发提交。
func (p *LRlqOaYmO) SubmitData(data *registry.FGmNba4MhioY) error {
	return errutil.D5aDHlQDLp(func() error { return p.zsD02JRnj(data.HTgPyE3VytZd) }, func() error { return p.lu3PfJZhV6t(data.JNcHHwTZ7) })
}

// lu3PfJZhV6t 对应 0x17a7f20，原实现明确丢弃在线 IP 上报错误。
func (p *LRlqOaYmO) lu3PfJZhV6t(alive types.CQhZrJ4dK) error {
	if len(alive) == 0 {
		return nil
	}
	address := p.GetUniUrl("alive")
	values := make(map[int64][]string)
	for _, entry := range alive {
		values[entry.Vy9ybu] = entry.PNtn9KAV
	}
	p.PostCommon(address, values)
	return nil
}
