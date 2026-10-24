package vlessvision

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// 原镜像常量切片0x5036e50/70/90；cipher名称来自原init的5次decoder。
var bytesAt5036E50 = []byte{0, 0x2b, 0, 2, 3, 4}
var bytesAt5036E70 = []byte{0x16, 3}
var bytesAt5036E90 = []byte{0x16, 3, 3}
var ciphersAt5096330 = map[uint16]string{
	0x1301: "TLS_AES_128_GCM_SHA256",
	0x1302: "TLS_AES_256_GCM_SHA384",
	0x1303: "TLS_CHACHA20_POLY1305_SHA256",
	0x1304: "TLS_AES_128_CCM_SHA256",
	0x1305: "TLS_AES_128_CCM_8_SHA256",
}

// YByzFSgf @0x1d0a6e0。检测原TLS记录标记和supported_versions片段。
// 原逻辑会按底层cap读取session ID之后的cipher，而非额外限制在len以内。
func YByzFSgf(data []byte, state *MxS0kZ) {
	if len(data) == 0 {
		return
	}
	state.BUwqtS--
	if len(data) >= 6 {
		if bytes.Equal(bytesAt5036E90, data[:3]) && data[5] == 2 {
			state.TvZPBqb = int32(binary.BigEndian.Uint16(data[3:5])) + 5
			state.UFPiQl0bbHk = true
			state.VRlXaTaT = true
			if len(data) >= 79 && state.TvZPBqb >= 79 {
				offset := int(data[43])
				state.OqJe390R06y = binary.BigEndian.Uint16(data[44+offset : 46+offset])
			}
		} else if bytes.Equal(bytesAt5036E70, data[:2]) && data[5] == 1 {
			state.VRlXaTaT = true
		}
	}
	if state.TvZPBqb > 0 {
		remaining := int(state.TvZPBqb)
		state.TvZPBqb -= int32(len(data))
		size := min(len(data), remaining)
		if bytes.Index(data[:size], bytesAt5036E50) == -1 {
			if state.TvZPBqb <= 0 {
				state.BUwqtS = 0
			}
		} else {
			if cipher, ok := ciphersAt5096330[state.OqJe390R06y]; ok {
				if cipher != "TLS_AES_128_CCM_8_SHA256" {
					state.YdJioLGvuLW = true
				}
			} else {
				// 原指令构造该字符串后未使用；保留此可见的源码表达，不添加日志。
				_ = "Old cipher: " + strconv.FormatUint(uint64(state.OqJe390R06y), 16)
			}
			state.BUwqtS = 0
		}
	}
}
