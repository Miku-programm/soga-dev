// 根据 IDA 的原始加载顺序及错误分支恢复；变量名由分析者指定。
package config

import (
	"flag"
	"maps"
	"os"
	"strings"

	"recovered.local/soga/util/rtx"
	"recovered.local/soga/crypto/hashutil"
)

// 初值由 0x13191f5 调用的字符串解码器 0x1324320 静态求值得到。
var configurationPathAt5096DB0 = "/etc/soga/soga.conf"

// ConfigurationPath exposes the original global to recovered startup callers.
func ConfigurationPath() string { return configurationPathAt5096DB0 }

// MkwNWciPHS 对应 0x1323b00。相对名称直接加前缀，不做路径清理。
func MkwNWciPHS(path string) {
	if path == "" || path[0] != '/' {
		path = "/etc/soga/" + path
	}
	configurationPathAt5096DB0 = path
}

// TvBECnD9fM6 对应 0x1323be0。保留原程序创建文件、合并来源和调用回调的顺序。
func TvBECnD9fM6() error {
	if !rtx.VSCJLy2No(configurationPathAt5096DB0) {
		_ = os.MkdirAll(rtx.EVZvjBCf(configurationPathAt5096DB0), 0755)
		file, err := os.OpenFile(configurationPathAt5096DB0, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		_ = file.Close()
	}
	configurationAt5095AA8 = maps.Clone(defaultsAt5095AA0)
	if err := zyxKcF82B(configurationPathAt5096DB0); err != nil {
		return err
	}
	for _, entry := range os.Environ() {
		parts := hashutil.DAaBs7_C(entry, "=", 2)
		if len(parts) != 2 {
			continue
		}
		name, value := parts[0], parts[1]
		if strings.HasPrefix(name, "DNS_") {
			if value != "" {
				configurationAt5095AA8[strings.ToLower(name)] = value
			}
		} else if _, exists := defaultsAt5095AA0[name]; exists {
			configurationAt5095AA8[strings.ToLower(name)] = value
		}
	}
	// 原匿名函数入口 0x1346da0；显式空值仍覆盖已有值。
	flag.CommandLine.Visit(func(f *flag.Flag) {
		value := f.Value.String()
		configurationAt5095AA8[strings.ToLower(f.Name)] = value
	})
	for name, value := range configurationAt5095AA8 {
		configurationAt5095AA8[name] = zuOmYdCNcr(value)
	}
	for _, configuration := range registeredAt5099320 {
		bFnBl5Ak9(configuration)
	}
	for _, callback := range afterLoadAt5099300 {
		callback()
	}
	return nil
}
