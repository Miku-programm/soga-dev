// 本文件根据原样本指令及字符串解码运算恢复，不是作者源码原文。
package tkbUwvW3ex

import (
	"fmt"
	"strings"

	"recovered.local/soga/vKl1IHr3"
)

// String 对应 0xec17e0；四个字段依次按两位十进制输出，不重新校验时段。
func (r *D1Teq5S1f) String() string {
	// 实参来自 0x5058e20 指针链、0x5058e50 及调用点的立即数。
	format := haZl2FW(0x9a7a3b4e, 0xfabd, 0x2e32, 0xb0f3, qxmp4I0)
	return fmt.Sprintf(format, r.LXcSWg26p, r.RBNUY7G, r.BZwEA15n, r.G1LUbU)
}

// Z4FKfucI_A5 对应 0xec1920。首个非法项使解析停止，但保留此前解析的结果。
// 空输入（包括只有逗号和空白的输入）产生一个覆盖全天的时段。
func Z4FKfucI_A5(text string) (ITMfakhJ, error) {
	result := make(ITMfakhJ, 0)
	for _, item := range vKl1IHr3.DAaBs7_C(text, ",", -1) {
		parsed, ok := DJr_rBzh(item)
		if !ok {
			// 实参来自 0x5058e6c、0x5058e28、0x5058e48 指针链及立即数 0x74。
			format := xybQauJZY_n(0xf2a99ecb, 0xa002, 0x65fb, 0x74, aWXs7MI)
			// 0x5d50c0 的无参数快路可返回 nil；此处始终传入一个参数，
			// 因而只走格式化错误路径，与 fmt.Errorf 的本次调用一致。
			return result, fmt.Errorf(format, item)
		}
		result = append(result, parsed)
	}
	if len(result) == 0 {
		result = append(result, &D1Teq5S1f{BZwEA15n: 24})
	}
	return result, nil
}

// String 对应 0xec1c00。保持原顺序，用英文逗号连接；不跳过空指针。
func (ranges ITMfakhJ) String() string {
	parts := make([]string, 0, len(ranges))
	for _, r := range ranges {
		if text := r.String(); text != "" {
			parts = append(parts, text)
		}
	}
	return strings.Join(parts, ",")
}

// qxmp4I0 对应 0xec2000；截取解码缓冲区的第 3 至 21 字节。
func qxmp4I0(data []byte) string { return string(data[3:22]) }

// haZl2FW 对应 0xec2060。种子逐字节取自原 ELF 的 0x20aecf6、0x20aed0e。
// 字节运算保留溢出回绕；回调收到完整 24 字节缓冲区。
func haZl2FW(a uint32, b, c, d uint16, callback func([]byte) string) string {
	left := [24]byte{0xf2, 0xfa, 0x55, 0x04, 0xa9, 0xcd, 0x54, 0x94, 0x36, 0xaa, 0xd3, 0xdd, 0x9f, 0xb5, 0x64, 0x3c, 0xd1, 0x95, 0x9f, 0x9b, 0x1c, 0xff, 0x66, 0x57}
	right := []byte{0x8b, 0x4a, 0x56, 0xaf, 0x93, 0xff, 0xb6, 0x5e, 0x13, 0xe8, 0xe1, 0x68, 0xb2, 0x0a, 0x54, 0x14, 0x05, 0x55, 0xba, 0xab, 0x2e, 0xb5, 0xe8, 0x24}
	a2, b0, b1, c0, c1, d0, d1 := byte(a>>16), byte(b), byte(b>>8), byte(c), byte(c>>8), byte(d), byte(d>>8)
	left[3] = 4 - a2
	left[23] += d0
	left[11] -= d1
	left[0] ^= b1
	left[9] += c1
	left[11] = (left[11] ^ b0) - d0 + c0
	left[17] ^= b1
	left[16] ^= d1
	right[15] = b1 + 20
	right[11] -= b0
	right[4] -= b1
	right[6] += 2 * b0
	right[21] ^= c1
	right[7] -= d1
	right[13] -= a2
	right[2] -= b0
	for i := range right {
		right[i] ^= left[i]
	}
	return callback(right)
}

// aWXs7MI 对应 0xec2240；仅返回前 31 字节，丢弃尾部混淆字节。
func aWXs7MI(data []byte) string { return string(data[:31]) }

// xybQauJZY_n 对应 0xec22a0。恢复原状态转移及全部字节运算，未执行原二进制。
func xybQauJZY_n(a uint32, b, c uint16, d byte, callback func([]byte) string) string {
	a0, a1, a2, a3 := byte(a), byte(a>>8), byte(a>>16), byte(a>>24)
	b0, b1, c0, c1 := byte(b), byte(b>>8), byte(c), byte(c>>8)
	data := make([]byte, 0, 33)
	state, count, accumulator := 4, 0, 157
	for state != 12 {
		accumulator ^= count * state
		switch state {
		case 0:
			data = append(data, b0+(a3^d^c1^0xaa)+a1, a1+(d^(c0+b1+4-a0))-d)
			state = 7
		case 1:
			data = append(data, c0+24)
			state = 14
		case 2:
			data = append(data, 0xb0-b1-b0, b0+2*d, d^((a2^0x84)-c0), 0xe8)
			state = 9
		case 3:
			data = append(data, c0+43, (d^(a2+54))-c0, 0x91, b1^0x35)
			state = 0
		case 4:
			data = append(data, a0+((c0^c1^(0x91-d))-c0), d^(a3+0x85), 0x13-a3-b1+c0+a2)
			state = 10
		case 5:
			data = append(data, c0^0xe6, b1+0x98, 0x64)
			state = 13
		case 6:
			data = append(data, 0x48)
			state = 5
		case 7:
			data = append(data, -(b1 + 28))
			state = 8
		case 8:
			for i := range data {
				data[i] -= byte(accumulator ^ i)
			}
			state = 12
		case 9:
			data = append(data, 0xfc)
			state = 6
		case 10:
			data = append(data, c1-14)
			state = 15
		case 11:
			data = append(data, a0^0xd3, b1+122, b1-82-b0)
			state = 2
		case 13:
			data = append(data, 0xda-b1, 8, 0x61, (b0^0x81)-d)
			state = 1
		case 14:
			data = append(data, c1^c0^(c0+b0+(c1^(65-a1))), a0+(c1^0xdd)+b1)
			state = 3
		case 15:
			data = append(data, c0+(c0^0xfe), c1+a3-59, (c1^0xa0)-d)
			state = 11
		}
		count++
	}
	return callback(data)
}
