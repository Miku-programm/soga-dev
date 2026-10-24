// 本文件依据所提供二进制的 IDA 指令恢复为可编译 Go 源码，非作者文件原文。
// 样本 SHA256：6660dcfbc7971188dcb751b440f4269aa718491b569d4c8629b751e6e339f92a。
// 函数名保留 pclntab 中的名称；文件名和下面的全局变量名由分析者指定。
// 配套分析证据保存在工作目录 ida-20261023/evidence/focused-bounded。
package config

import (
	"os"
	"strings"
)

// configurationAt5095AA8 对应原程序地址 0x5095aa8 的字符串映射。
// 初始化由已核对的包初始化函数 0x1319080 中的 makemap_small 对应。
var configurationAt5095AA8 = make(map[string]string)

// zyxKcF82B 对应原函数 0x1324020。
// 先完整读取文件，再逐行处理；读取失败不改动映射。
func zyxKcF82B(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		configurationAt5095AA8[strings.ToLower(key)] = value
	}
	return nil
}

// zuOmYdCNcr 对应原函数 0x13241a0。
// 只有“空格加井号”开始行内注释；该处理与文件读取是两个独立阶段。
func zuOmYdCNcr(value string) string {
	if index := strings.Index(value, " #"); index >= 0 {
		value = value[:index]
	}
	return strings.TrimSpace(value)
}
