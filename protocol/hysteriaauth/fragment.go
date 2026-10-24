// 本文件按原二进制 0x1c2e460、0x1c2e880 恢复，不是作者文件原文。
// 原元数据包名 M14gyoUSscE，RTTI 显示包名 esxKUkjg。
package hysteriaauth

import "recovered.local/soga/crypto/misc"

// TOF1KM6hi 保留原函数的 uint8 分片计数与原边界行为。
// 原机器码没有后续上游版本添加的 maxPayload <= 0 检查。
func TOF1KM6hi(m *misc.ItLMumzBf, maxSize int) []misc.ItLMumzBf {
	if m.Size() <= maxSize {
		return []misc.ItLMumzBf{*m}
	}
	payload := m.HmoFPk
	maxPayload := maxSize - m.HeaderSize()
	count := uint8((len(payload) + maxPayload - 1) / maxPayload)
	fragments := make([]misc.ItLMumzBf, count)
	var id uint8
	for off := 0; off < len(payload); {
		size := min(len(payload)-off, maxPayload)
		fragment := *m
		fragment.QmDvxK_8Q, fragment.Kb9FLg7wuNE = id, count
		fragment.HmoFPk = payload[off : off+size]
		fragments[id] = fragment
		off += size
		id++
	}
	return fragments
}

// YWXI4jW3 是原 48 字节类型，字段名、顺序与 RTTI 一致。
type YWXI4jW3 struct {
	sQXA1D9aRa   uint16
	mk2YbtR      []*misc.ItLMumzBf
	uGzaQtaWh    uint8
	zLEI1e06OnKH int
}

func (d *YWXI4jW3) Feed(m *misc.ItLMumzBf) *misc.ItLMumzBf {
	if m.Kb9FLg7wuNE <= 1 {
		return m
	}
	if m.QmDvxK_8Q >= m.Kb9FLg7wuNE {
		return nil
	}
	if d.sQXA1D9aRa != m.CBHlJenOD || uint8(len(d.mk2YbtR)) != m.Kb9FLg7wuNE {
		d.sQXA1D9aRa = m.CBHlJenOD
		d.mk2YbtR = make([]*misc.ItLMumzBf, m.Kb9FLg7wuNE)
		d.mk2YbtR[m.QmDvxK_8Q] = m
		d.uGzaQtaWh = 1
		d.zLEI1e06OnKH = len(m.HmoFPk)
		return nil
	}
	if d.mk2YbtR[m.QmDvxK_8Q] != nil {
		return nil
	}
	d.mk2YbtR[m.QmDvxK_8Q] = m
	d.uGzaQtaWh++
	d.zLEI1e06OnKH += len(m.HmoFPk)
	if int(d.uGzaQtaWh) != len(d.mk2YbtR) {
		return nil
	}
	payload := make([]byte, d.zLEI1e06OnKH)
	off := 0
	for _, frag := range d.mk2YbtR {
		off += copy(payload[off:], frag.HmoFPk)
	}
	m.HmoFPk, m.QmDvxK_8Q, m.Kb9FLg7wuNE = payload, 0, 1
	return m
}
