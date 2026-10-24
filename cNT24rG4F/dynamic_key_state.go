package cNT24rG4F

// Original init decoder 0x1344bc0 yields a274902e6d909df591d5c0;
// byte-wise XOR with 0x2c45a824 initializes this independent FSM seed.
var dynamicKeyStateSeedAt50C6DC8 uint64 = 0x2c45a85d

// HuQEoQ: 0x1342600. Preserve original state reads, independent seed, and
// asynchronous exit on inconsistent markers; do not collapse to a flag getter.
func HuQEoQ() bool {
	seed := dynamicKeyStateSeedAt50C6DC8
	var state, marker uint64
	var snapshot, result bool
	for {
		switch state {
		case 611538049:
			if keyStateMirrorAt3787300 == snapshot {
				state = seed ^ 0x5fb09f57
			} else {
				state = seed ^ 0x35a78353
			}
		case 1553490299:
			state = seed ^ 0x592394ca
			marker = keyStateMarkerAt3787800
		case 403619847:
			state = seed ^ 0x32f2f750
		case 1902515654, 900613051:
			if len(FrFQQhzTR5fY()) != 0 {
				state = seed ^ 0x44155f75
			} else {
				state = seed ^ 0x309365fd
			}
			result = true
		case 2025105290:
			if keyStateMarkerAt3787800 == 0x5a4b3c2d5a4b3c2d {
				state = seed ^ 0x19ebebe6
			} else {
				state = seed ^ 0xa2a341
			}
		case 1969634455:
			if marker == 0xa5b6c7d8a5b6c7d8 {
				state = seed ^ 0x6a4b54a0
			} else {
				state = seed ^ 0x51706af8
			}
		case 483839392, 1241293814:
			return result
		case 1750136616:
			result = keyStateAt50C6C48
			state = seed ^ 0x65b90bab
		case 753339164, 1945450250, 2100675237:
			go cyfZyjB8()
			return true
		case 434252558:
			if keyStateAt50C6C48 {
				state = seed ^ 0x70ddc926
			} else {
				state = seed ^ 0x344b685a
			}
		case 515333901:
			state = seed ^ 0x44e008ee
		case 1175387389, 1755685043:
			if keyStateAt50C6C48 {
				state = seed ^ 0x5d23bd9b
			} else {
				state = seed ^ 0x54f10fd7
			}
		default:
			snapshot = keyStateAt50C6C48
			state = seed ^ 0x836fcdc
		}
	}
}
