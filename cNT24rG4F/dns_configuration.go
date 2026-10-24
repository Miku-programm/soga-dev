package cNT24rG4F

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"gopkg.in/yaml.v3"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/muc279Mpa6S4"
)

// AorTZHf4G9Yq：原 RTTI 0x3624f80，40 字节，字段偏移 0、16。
type AorTZHf4G9Yq struct {
	ZkuUXsveafC WbFtukHA9 `yaml:"strategy"`
	IxW5Yxj     []string  `yaml:"rules"`
}

var dnsRulesAt5095AD0 atomic.Pointer[LFWTTXF7JR.U863Mm01Nre]

// String 对应 0x1335560。
func (rule *AorTZHf4G9Yq) String() string {
	return fmt.Sprintf("{ strategy: %v, rules: %v }", rule.ZkuUXsveafC, rule.IxW5Yxj)
}

// WreRtju 对应 0x1335720；远程 URL 在启动 goroutine 时捕获。
func WreRtju(ctx context.Context) {
	if url := configAt509CB40.MFHBO42; url != "" {
		go CAodbhZf.KJa7rJUaLPYe(ctx, url, func(data []byte) error { return as3gRNOKx9KD(url, data) })
		return
	}
	ka8yZsm4z()
	err := CAodbhZf.SharedFileWatcher().WatchChanged(ctx, func() bool { ka8yZsm4z(); return true }, dnsPathAt5096DE0)
	if err != nil {
		muc279Mpa6S4.V8GN5Stk(8, "watch dns config file <%v> failed: %v", dnsPathAt5096DE0, err)
	}
}

// ka8yZsm4z 对应 0x1335980。文件读取失败不输出日志。
func ka8yZsm4z() {
	data, err := os.ReadFile(dnsPathAt5096DE0)
	if err != nil {
		return
	}
	if err := as3gRNOKx9KD(dnsPathAt5096DE0, data); err != nil {
		muc279Mpa6S4.V8GN5Stk(8, "load dns config <%v> failed: %v", dnsPathAt5096DE0, err)
	}
}

// as3gRNOKx9KD 对应 0x1335ac0。yaml.v3 的 Node/decoder 字段、112/152
// 字节布局和非 strict unmarshal 流程均与原 DUx29Y4U5a9 核对。
func as3gRNOKx9KD(name string, data []byte) error {
	rules := make(map[string]*AorTZHf4G9Yq)
	if firstErr := yaml.Unmarshal(data, &rules); firstErr != nil {
		legacy := make(map[string][]string)
		if secondErr := yaml.Unmarshal(data, &legacy); secondErr != nil {
			return fmt.Errorf("%v %v", firstErr, secondErr)
		}
		muc279Mpa6S4.TRTbrquD(4, "检测到 dns.yml 的格式是旧版本的(v2.7.5-)，目前自动兼容旧版本格式，未来会废弃旧格式，请尽快切换到新版本格式(v2.7.6+)")
		for key, values := range legacy {
			rules[key] = &AorTZHf4G9Yq{ZkuUXsveafC: "ipv4_first", IxW5Yxj: values}
		}
	}
	if len(rules) > 0 {
		muc279Mpa6S4.TRTbrquD(0, "--------------------------dns config start--------------------------")
		for key, rule := range rules {
			muc279Mpa6S4.V8GN5Stk(0, "%v: %v", key, rule)
		}
		muc279Mpa6S4.TRTbrquD(0, "---------------------------dns config end---------------------------")
	}
	muc279Mpa6S4.V8GN5Stk(0, "load dns config <%v> success, rule num: %v", name, len(rules))
	converted := D2YhfBI(rules)
	dnsRulesAt5095AD0.Store(&converted)
	return nil
}

// D2YhfBI 对应 0x1336200。逗号分割调用 0xb7d720：去空白、丢弃空项，
// 保留重复项。原 map 遍历顺序不固定，规则的服务器切片直接共享。
func D2YhfBI(rules map[string]*AorTZHf4G9Yq) LFWTTXF7JR.U863Mm01Nre {
	result := LFWTTXF7JR.U863Mm01Nre{}
	for key, rule := range rules {
		names := []string{}
		for _, part := range strings.Split(strings.TrimSpace(key), ",") {
			if part = strings.TrimSpace(part); part != "" {
				names = append(names, part)
			}
		}
		if len(names) != 0 {
			result = append(result, &LFWTTXF7JR.Axa8uKIB4JL{K8BNfziIgb: names, CM0eMsqyP: string(rule.ZkuUXsveafC), CZ80ik: rule.IxW5Yxj})
		}
	}
	return result
}
