package config

// 两个独立种子来自原 init 的字节解码和异或归约。
var panelKeySeedAt50C6DC0 uint64 = 0x1f0ea50f
var panelKeySeedAt50C6DD0 uint64 = 0x77ad37b6

// Gh2DnstC（0x1342300）保留每次全局读取和异步退出分支。
func Gh2DnstC() bool {
	seed := panelKeySeedAt50C6DC0
	var state, marker uint64
	var snapshot, result bool
	for {
		switch state {
		case 1678805806:
			if keyStateMirrorAt3787300 == snapshot {
				state = seed ^ 0x0559289c
			} else {
				state = seed ^ 0x53d5dad0
			}
		case 908297162:
			state = seed ^ 0x22339a4d
			marker = keyStateMarkerAt3787800
		case 0x44a2931, 1464426057:
			if keyStateAt50C6C48 {
				state = seed ^ 0x5f61667b
			} else {
				state = seed ^ 0x54ff4f03
			}
		case 1081066356, 1647270436:
			if len(FrFQQhzTR5fY()) != 0 {
				state = seed ^ 0x405e0a4a
			} else {
				state = seed ^ 0x57e710cf
			}
			result = true
		case 1274145292:
			if keyStateMarkerAt3787800 == 0x5a4b3c2d5a4b3c2d {
				state = seed ^ 0x7d21ff2b
			} else {
				state = seed ^ 0x1205ffd8
			}
		case 1027424066:
			if marker == 0xa5b6c7d8a5b6c7d8 {
				state = seed ^ 0x4847fb46
			} else {
				state = seed ^ 0x6871b40b
			}
		case 1223275968:
			state = seed ^ 0x236a239c
		case 1599123269:
			result = keyStateAt50C6C48
			state = seed ^ 0x702c6a73
		case 218847959, 441945491, 2004816132:
			go cyfZyjB8()
			return true
		case 1864552316, 1013221011:
			return result
		case 1289453535:
			state = seed ^ 0x081b61d6
		case 387302617:
			if keyStateAt50C6C48 {
				state = seed ^ 0x292d26c5
			} else {
				state = seed ^ 0x1b448c3e
			}
		default:
			state = seed ^ 0x7b1e2e21
			snapshot = keyStateAt50C6C48
		}
	}
}

// JatfUF38O（0x1342900）是另一个独立的面板检查状态机。
func JatfUF38O() bool {
	seed := panelKeySeedAt50C6DD0
	var state uint64
	var snapshot, mirror, result bool
	for {
		switch state {
		case 325968968:
			mirror = keyStateMirrorAt3787300
			state = seed ^ 0x12820114
		case 2098741048:
			if keyStateMarkerAt3787800 == 0xa5b6c7d8a5b6c7d8 {
				state = seed ^ 0x6e345c8e
			} else {
				state = seed ^ 0x2160600c
			}
		case 1412432126, 429484856:
			if keyStateAt50C6C48 {
				state = seed ^ 0x683874cd
			} else {
				state = seed ^ 0x0f8a3b48
			}
		case 529875835, 1183966196:
			if len(FrFQQhzTR5fY()) != 0 {
				state = seed ^ 0x68ef151d
			} else {
				state = seed ^ 0x0244f866
			}
			result = true
		case 2015825150:
			if keyStateMarkerAt3787800 == 0x5a4b3c2d5a4b3c2d {
				state = seed ^ 0x313cd442
			} else {
				state = seed ^ 0x5512e94a
			}
		case 1456297914, 582999804, 1588521645:
			go cyfZyjB8()
			return true
		case 1978257360, 152911492:
			return result
		case 524427947:
			state = seed ^ 0x26d6d4cd
		case 1508609296:
			state = seed ^ 0x2903dd1b
		case 1465286520:
			if keyStateAt50C6C48 {
				state = seed ^ 0x0ab5088e
			} else {
				state = seed ^ 0x239d3748
			}
		case 1697592994:
			if snapshot == mirror {
				state = seed ^ 0x2e46baa6
			} else {
				state = seed ^ 0x20fb48ce
			}
		case 1367073659:
			state = seed ^ 0x7eb00932
			result = keyStateAt50C6C48
		default:
			snapshot = keyStateAt50C6C48
			state = seed ^ 0x64c0d3fe
		}
	}
}
