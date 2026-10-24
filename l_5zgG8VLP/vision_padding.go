package l_5zgG8VLP

import (
	"bytes"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/uLofn6j"
)

// MxS0kZ 对应原 RTTI0x371aa38，64字节；原状态字段名称全部保留。
type MxS0kZ struct {
	RW3eHGaIH0ya []byte
	BUwqtS       int
	YdJioLGvuLW  bool
	UFPiQl0bbHk  bool
	VRlXaTaT     bool
	OqJe390R06y  uint16
	TvZPBqb      int32
	EO4y9oyDkzq  bool
	SDw6IIM1D51U bool
	Z43zq7       int32
	ZOEKPa7JPx5  int32
	PIsBMe       int32
	EXbj1L5      byte
	ZkTHd3       bool
	Pd9MU7P6IX   bool
}

// 原切片全局0x5036eb0直接指向镜像中的17 03 03。
var bytesAt5036EB0 = []byte{0x17, 3, 3}

// AaBEhR @0x1d09f80。只切片不复制；优先在21..8171范围内最后的TLS标记拆分。
func AaBEhR(out [][]byte, data []byte) [][]byte {
	if len(data) < 8171 {
		return append(out, data)
	}
	at := bytes.LastIndex(data, bytesAt5036EB0)
	if uint64(at-21) > 0x1fd6 {
		at = 4096
	}
	out = AaBEhR(out, data[:at])
	return AaBEhR(out, data[at:])
}

// IFq7RHaxl @0x1d0a120。随机仅决定padding长度；原padding字节直接推进
// 池缓冲写位置，不额外随机填充或清零。prefix指针存在时写入后置nil。
func IFq7RHaxl(out *D_zCag1qTV.X_aXX5, data []byte, command byte, prefix *[]byte, longPadding bool) {
	var padding int
	if len(data) < 900 && longPadding {
		padding = uLofn6j.Co6D_zfls(500) - len(data) + 900
	} else {
		padding = uLofn6j.Co6D_zfls(256)
	}
	if padding > 8171-len(data) {
		padding = 8171 - len(data)
	}
	prefixLen := 0
	if prefix != nil {
		prefixLen = len(*prefix)
	}
	out.GrowBy(padding + prefixLen + len(data) + 5)
	if prefix != nil {
		_, _ = out.Write(*prefix)
		*prefix = nil
	}
	out.WriteByte(command)
	out.WriteUint16(uint16(len(data)))
	out.WriteUint16(uint16(padding))
	if data != nil {
		_, _ = out.Write(data)
	}
	out.SkipEnd(padding)
}

// DtcbafIY7SUd @0x1d0a3a0。原状态可跨输入缓冲；未识别UUID时返回输入
// 指针，已识别时返回新缓冲。非零command结束后剩余字节复制但不推进源位置。
func DtcbafIY7SUd(input *D_zCag1qTV.X_aXX5, state *MxS0kZ) *D_zCag1qTV.X_aXX5 {
	if state.Z43zq7 == -1 && state.ZOEKPa7JPx5 == -1 && state.PIsBMe == -1 {
		if input.Len() < 21 || !bytes.Equal(state.RW3eHGaIH0ya, input.BytesTo(16)) {
			return input
		}
		input.Skip(16)
		state.Z43zq7 = 5
	}
	out := D_zCag1qTV.S4OQXWIZ(input.Len())
	for input.Len() > 0 {
		if state.Z43zq7 > 0 {
			b, err := input.ReadUint8()
			if err != nil {
				return out
			}
			switch state.Z43zq7 {
			case 5:
				state.EXbj1L5 = b
			case 4:
				state.ZOEKPa7JPx5 = int32(b) << 8
			case 3:
				state.ZOEKPa7JPx5 |= int32(b)
			case 2:
				state.PIsBMe = int32(b) << 8
			case 1:
				state.PIsBMe |= int32(b)
			}
			state.Z43zq7--
		} else if state.ZOEKPa7JPx5 > 0 {
			n := min(int(state.ZOEKPa7JPx5), input.Len())
			part := input.BytesTo(n)
			input.Skip(n)
			_, _ = out.Write(part)
			state.ZOEKPa7JPx5 -= int32(n)
		} else {
			n := min(int(state.PIsBMe), input.Len())
			input.Skip(n)
			state.PIsBMe -= int32(n)
		}
		if state.Z43zq7 <= 0 && state.ZOEKPa7JPx5 <= 0 && state.PIsBMe <= 0 {
			if state.EXbj1L5 != 0 {
				state.Z43zq7 = -1
				state.ZOEKPa7JPx5 = -1
				state.PIsBMe = -1
				if input.Len() > 0 {
					_, _ = out.Write(input.Bytes())
				}
				return out
			}
			state.Z43zq7 = 5
		}
	}
	return out
}
