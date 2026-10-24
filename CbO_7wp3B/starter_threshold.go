package CbO_7wp3B

import (
	"time"

	"recovered.local/soga/uLofn6j"
)

// h68JK8g 对应 0x1fc2920。保留原随机读取次数、混合运算和时间分支，
// 没有把该控制链替换为常量。原机码的三组 32 字节输入轨迹逐轮核对了累积值，
// 并另行检查空输入与时间奇偶两条路径。
func h68JK8g(input []byte) int {
	var accumulator int64 = 0xdeadbeef
	for index, value := range input {
		var mixed int64
		if index%7 == 0 {
			mixed = accumulator ^ (int64(value) << uint(index%15))
		} else {
			mixed = accumulator + int64(value^0xcc)
		}
		accumulator = mixed % (uLofn6j.Ctna67jV()%0x7fffffff + 1)
	}
	masked := accumulator & 0xffffff
	factor := int64(1)
	if accumulator > 0 {
		factor = accumulator%3 + 1
	}
	base := ((masked + 100) - (factor*masked)/factor) / 10
	var result int64
	if time.Now().UnixNano()%2 == 0 {
		result = base + 7
		result -= 14
	} else {
		result = base + 9
		result -= 16
	}
	return int((result << 2) >> 2)
}
