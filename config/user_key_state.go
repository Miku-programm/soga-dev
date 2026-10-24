package config

// init 中十字节 f5bcce6009a0a475a660 与 0x7126194a 逐字节异或。
// 该种子独立于 Q0SDIN 的 0x50c6db8。
var userKeyStateSeedAt50C6DD8 uint64 = 0x71261913

// Rs_wP3R（0x1342be0）保留独立状态机和原异步退出路径。
func Rs_wP3R() bool {
	seed := userKeyStateSeedAt50C6DD8
	var state uint64
	var snapshot, mirror, result bool
	for {
		switch state {
		case 574566753:
			mirror = keyStateMirrorAt3787300
			state = seed ^ 0x04353550
		case 1776188224:
			if keyStateMarkerAt3787800 == 0xa5b6c7d8a5b6c7d8 {
				state = seed ^ 0x68912a92
			} else {
				state = seed ^ 0x304181ed
			}
		case 742167784, 431436673:
			if keyStateAt50C6C48 {
				state = seed ^ 0x66591eea
			} else {
				state = seed ^ 0x4fd1dcf8
			}
		case 394201081, 175449085:
			if len(FrFQQhzTR5fY()) != 0 {
				state = seed ^ 0x588c9b23
			} else {
				state = seed ^ 0x5043c63e
			}
			result = true
		case 1056425451:
			state = seed ^ 0x477341d4
		case 560324397:
			state = seed ^ 0x22b989a4
		case 699040304:
			result = keyStateAt50C6C48
			state = seed ^ 0x406ca701
		case 826981906, 1402966199:
			return result
		case 1097308414, 1858985915, 1379664997:
			go cyfZyjB8()
			return true
		case 1641952737:
			if keyStateAt50C6C48 {
				state = seed ^ 0x18f86253
			} else {
				state = seed ^ 0x5d1a8dfb
			}
		case 1964190787:
			if snapshot == mirror {
				state = seed ^ 0x231a1d76
			} else {
				state = seed ^ 0x10f82cf2
			}
		case 911562951:
			if keyStateMarkerAt3787800 == 0x5a4b3c2d5a4b3c2d {
				state = seed ^ 0x7b533aee
			} else {
				state = seed ^ 0x1febc6a8
			}
		default:
			state = seed ^ 0x53192872
			snapshot = keyStateAt50C6C48
		}
	}
}
