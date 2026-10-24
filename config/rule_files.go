package config

import (
	"bufio"
	"context"
	"os"
	"strings"

	"recovered.local/soga/util/rtx"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/audit/matcher"
)

// Dgd6obkvt 对应 0x131a2e0。仅注册监视，首次加载由启动装配另行调用。
func Dgd6obkvt(path string) {
	if len(path) == 0 || path[0] != '/' {
		path = "/etc/soga/" + path
	}
	blockListPathAt5096DA0 = path
	err := rtx.SharedFileWatcher().WatchChanged(context.Background(), func() bool {
		LnLDMWbjk3() // 原闭包 0x1346e60。
		return true
	}, path)
	if err != nil {
		logx.V8GN5Stk(8, "watch block list file <%v> failed: %v", blockListPathAt5096DA0, err)
	}
}

// QJZTKBo 对应 0x1335660，不做路径清理、不检查文件是否存在。
func QJZTKBo(path string) {
	if len(path) == 0 || path[0] != '/' {
		path = "/etc/soga/" + path
	}
	dnsPathAt5096DE0 = path
}

// AcIcHDv 对应 0x133e5a0，回调使用调用时的全局路径。
func AcIcHDv(path string) {
	if len(path) == 0 || path[0] != '/' {
		path = "/etc/soga/" + path
	}
	whiteListPathAt5096DF0 = path
	err := rtx.SharedFileWatcher().WatchChanged(context.Background(), func() bool {
		I5PWiaSBwdar() // 原闭包 0x1346c40。
		return true
	}, path)
	if err != nil {
		logx.V8GN5Stk(8, "watch white list file <%v> failed: %v", whiteListPathAt5096DF0, err)
	}
}

// LnLDMWbjk3 对应 0x131a4a0。读取失败保留旧快照；配置远程 URL 时跳过本地文件。
func LnLDMWbjk3() {
	if configAt509B660.BooNOdXcSa3U != "" {
		return
	}
	file, err := os.Open(blockListPathAt5096DA0)
	if err != nil {
		return
	}
	defer file.Close()
	rules := readRuleFile(file)
	// 原日志先于 forbidden_ports 追加，因此计数只包含文件规则。
	logx.V8GN5Stk(0, "load block list <%v> success, rule num: %v", blockListPathAt5096DA0, len(rules))
	for _, port := range configAt50A2900.DUpePpa39q {
		rules = append(rules, "port:"+port)
	}
	blockListAt5095A90.Store(&rules)
}

// I5PWiaSBwdar 对应 0x133e760；原函数先发布快照，再输出日志。
func I5PWiaSBwdar() {
	if configAt5096E00.M4qDq3KsUi != "" {
		return
	}
	file, err := os.Open(whiteListPathAt5096DF0)
	if err != nil {
		return
	}
	defer file.Close()
	rules := readRuleFile(file)
	whiteListAt5095AD8.Store(&rules)
	logx.V8GN5Stk(0, "load white list <%v> success, rule num: %v", whiteListPathAt5096DF0, len(rules))
}

// readRuleFile 合并上述两个函数中相同的内联循环（分析命名）。
// 原 ReadLine 的 isPrefix 被忽略：超过 4096 字节的行按片段分别解析。
// 任意读取错误结束循环，并由调用方发布已读到的规则。
func readRuleFile(file *os.File) matcher.VHWq49iRat {
	rules := matcher.VHWq49iRat{}
	reader := bufio.NewReaderSize(file, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return rules
		}
		text := strings.TrimSpace(string(line))
		if len(text) == 0 || text[0] == '#' {
			continue
		}
		text = strings.TrimSpace(strings.Split(text, "#")[0])
		rules = append(rules, text)
	}
}
