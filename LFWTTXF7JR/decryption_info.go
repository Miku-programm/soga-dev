package LFWTTXF7JR

import (
	"encoding/base64"
	"strconv"
	"strings"
)

// Bqm34MJ 按原 RTTI 0x36c9308 的56字节布局恢复；显示包名jcasg8c，
// uncommon_package_path 为LFWTTXF7JR。这是二进制源码表示，并非作者文件原文。
type Bqm34MJ struct {
	I_Wzoo_nf   string
	KAeFR260tbP uint32
	E3a03_Hf    int64
	R3tya30     int64
	LRGrsuCn    string
}

// GetDecryptionInfo 对应0x12763e0。保留原字符串切片及解码长度判断，
// 包括忽略base64错误而检验已解码长度、将所有短分段长度累计为前缀的行为。
func (v *VkkFLH) GetDecryptionInfo() (*Bqm34MJ, bool) {
	parts := strings.Split(v.DlDK5Y9AS, ".")
	if len(parts) < 4 || parts[0] != "mlkem768x25519plus" {
		return nil, false
	}
	out := &Bqm34MJ{I_Wzoo_nf: v.DlDK5Y9AS}
	switch parts[1] {
	case "native":
	case "xorpub":
		out.KAeFR260tbP = 1
	case "random":
		out.KAeFR260tbP = 2
	default:
		return nil, false
	}
	span := strings.SplitN(strings.TrimSuffix(parts[2], "s"), "-", 2)
	n, err := strconv.Atoi(span[0])
	if err != nil {
		return nil, false
	}
	out.E3a03_Hf = int64(n)
	if len(span) == 2 {
		n, err = strconv.Atoi(span[1])
		if err != nil {
			return nil, false
		}
		out.R3tya30 = int64(n)
	}
	prefix := 0
	for _, s := range parts[3:] {
		if len(s) >= 20 {
			b, _ := base64.RawURLEncoding.DecodeString(s)
			if len(b) != 32 && len(b) != 64 {
				return nil, false
			}
		} else {
			prefix += len(s) + 1
		}
	}
	out.I_Wzoo_nf = out.I_Wzoo_nf[len(parts[2])+27:]
	if prefix > 0 {
		out.LRGrsuCn = out.I_Wzoo_nf[:prefix-1]
		out.I_Wzoo_nf = out.I_Wzoo_nf[prefix:]
	}
	return out, true
}
