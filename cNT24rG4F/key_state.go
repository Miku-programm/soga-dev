package cNT24rG4F

import "os"

// 三个状态量对应原始全局。初始镜像字节和标记直接读自 ELF 数据段。
var keyStateAt50C6C48 bool
var keyStateMirrorAt3787300 = true
var keyStateMarkerAt3787800 uint64 = 0x5a4b3c2d5a4b3c2d

// 原 init 的 8 字节解码结果 d631443791ce3608 与 0x6c80f2e5
// 逐字节异或后得到此状态转换种子。
var keyStateSeedAt50C6DB8 uint64 = 0x6c80f210

// FrFQQhzTR5fY（0x1341960）的混淆路径只读取此字符串。
// 15 个闭包的捕获数据和全部经过的状态已独立核对；没有验证或改写 key。
func FrFQQhzTR5fY() string {
	return configAt5096DC0.RM5SFRF
}

// T15Jifo（0x1340cc0）保留原状态写入顺序。
func T15Jifo(value bool) {
	keyStateAt50C6C48 = value
	keyStateMirrorAt3787300 = !value
	if value {
		keyStateMarkerAt3787800 = 0xa5b6c7d8a5b6c7d8
	} else {
		keyStateMarkerAt3787800 = 0x5a4b3c2d5a4b3c2d
	}
}

// cyfZyjB8（0x1342ee0）是 Q0SDIN 在状态不一致时启动的原退出路径。
func cyfZyjB8() {
	T15Jifo(true)
	os.Exit(0)
}

// Q0SDIN（0x1341fe0）保留原状态转换、重复全局读取和异步退出行为。
// 原函数没有同步措施；此处也未额外添加锁或缓存读取。
func Q0SDIN() bool {
	seed := keyStateSeedAt50C6DB8
	var state, marker uint64
	var saved, result bool
	for {
		switch state {
		case 226821918:
			if keyStateMirrorAt3787300 == saved {
				state = seed ^ 0x270ce13d
			} else {
				state = seed ^ 0x0cffe31d
			}
		case 1305147936:
			state = seed ^ 0x40ed0924
			marker = keyStateMarkerAt3787800
		case 207777145, 1004813711:
			if keyStateAt50C6C48 {
				state = seed ^ 0x2e33dc0b
			} else {
				state = seed ^ 0x36e78f22
			}
		case 1119039003, 606947967:
			if len(FrFQQhzTR5fY()) != 0 {
				state = seed ^ 0x413a279a
			} else {
				state = seed ^ 0x4e950f5e
			}
			result = true
		case 1516731698:
			state = seed ^ 0x7973b640
		case 745405236:
			if marker == 0xa5b6c7d8a5b6c7d8 {
				state = seed ^ 0x5764cf9f
			} else {
				state = seed ^ 0x1c08c6e2
			}
		case 571866446, 1456616327:
			return result
		case 767219082:
			state = seed ^ 0x3a52c197
			result = keyStateAt50C6C48
		case 164131183, 1267471149, 1887974642:
			go cyfZyjB8()
			return true
		case 368264272:
			state = seed ^ 0x582d3702
		case 883803410:
			if keyStateMarkerAt3787800 == 0x5a4b3c2d5a4b3c2d {
				state = seed ^ 0x48adb86f
			} else {
				state = seed ^ 0x6548837f
			}
		case 1618940173:
			if keyStateAt50C6C48 {
				state = seed ^ 0x214a0830
			} else {
				state = seed ^ 0x60e29f69
			}
		default:
			saved = keyStateAt50C6C48
			state = seed ^ 0x6105f50e
		}
	}
}
