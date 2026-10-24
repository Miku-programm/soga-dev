package cNT24rG4F

import (
	"strconv"
	"strings"
)

// 全局地址 0x5095ae0；cNT24rG4F.init 首先创建此集合。
var blockListWhiteIDsAt5095AE0 = make(map[int64]struct{})

// init0func1 对应 cNT24rG4F.init.0.func1，地址 0x1346e80。
// 再次调用时原集合不会被清空；无法解析的条目直接跳过。
func init0func1() {
	for _, text := range configAt509B660.U2hwBFC {
		if id, err := strconv.ParseInt(text, 10, 64); err == nil {
			blockListWhiteIDsAt5095AE0[id] = struct{}{}
		}
	}
}

// rHAcgy 对应地址 0x132dc40。两个混淆字符串已通过静态字节运算恢复。
func rHAcgy() {
	switch strings.ToLower(configAt5099340.DlaO1l) {
	case "strict":
		configAt5099340.zdRAxv6CNPMA = 1
	case "strict_path":
		configAt5099340.zdRAxv6CNPMA = 2
	case "strict_host":
		configAt5099340.zdRAxv6CNPMA = 3
	default:
		configAt5099340.zdRAxv6CNPMA = 0
	}
}

// mvo6TVLM 对应地址 0x1326740；默认目录由 0x13268a0、0x1326840 解码得到。
func mvo6TVLM() {
	if configAt509B380.TsRVDHP0SA <= 0 {
		configAt509B380.TsRVDHP0SA = 24
	}
	if configAt509B380.FpIPwenv == "" {
		configAt509B380.FpIPwenv = "/etc/soga/"
	}
	if !strings.HasSuffix(configAt509B380.FpIPwenv, "/") {
		configAt509B380.FpIPwenv += "/"
	}
}
