package xboard

import (
	"fmt"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/crypto/hashutil"
	"sync"
)

var certificateErrorsAt509C880 sync.Map

// oN0fey 对应 0x1788ee0；跳过没有等号的段，后来的重复选项覆盖前者。
func (n *OGz0Bb0) oN0fey() map[string]string {
	out := make(map[string]string)
	for _, option := range hashutil.DAaBs7_C(n.XXzsbzH3U7, ";", -1) {
		parts := hashutil.DAaBs7_C(option, "=", 2)
		if len(parts) == 2 {
			out[parts[0]] = parts[1]
		}
	}
	return out
}

// bGlB2e_hEh1 对应 0x1789220，保留独立 Load/Store 及原日志级别。
func bGlB2e_hEh1(id int, message string) {
	if old, ok := certificateErrorsAt509C880.Load(id); !ok || old != message {
		certificateErrorsAt509C880.Store(id, message)
		logx.V8GN5Stk(4, "node %v: %v", id, message)
	}
}

// yU6F1oN6a 对应 0x1789000，无效证书只记录警告，不返回错误。
func (n *OGz0Bb0) yU6F1oN6a(id int) (string, string) {
	if n.DiDLlInAz.H_HT1L != "content" {
		certificateErrorsAt509C880.LoadAndDelete(id)
		return "", ""
	}
	cert, key := n.DiDLlInAz.HrieyL3JKL5, n.DiDLlInAz.FjIridNPs19h
	if cert == "" || key == "" {
		bGlB2e_hEh1(id, "cert_mode is content but cert_content or key_content is empty, ignore panel cert")
		return "", ""
	}
	if err := rtx.ZvmhWx(cert, key); err != nil {
		bGlB2e_hEh1(id, fmt.Sprintf("panel cert content invalid, ignore panel cert: %v", err))
		return "", ""
	}
	certificateErrorsAt509C880.LoadAndDelete(id)
	return cert, key
}

// pC0PsvYElMrC 对应 0x1789400。
func (n *OGz0Bb0) pC0PsvYElMrC() (string, error) {
	if !n.Biccak.IMskp1_c.Orib8hhz.Value() {
		return "", nil
	}
	if n.Biccak.IMskp1_c.HQ5AFB2w4c == "" {
		return "", fmt.Errorf("ech enabled but key is empty")
	}
	return n.Biccak.IMskp1_c.HQ5AFB2w4c, nil
}
