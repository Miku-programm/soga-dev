package MFFa9yT

import "testing"

// 黄金值来自原机器指令在真实初始化 seed 下的有界 trace，非从 Go 实现生成。
func TestOriginalSessionGateHandlers(t *testing.T) {
	old := sessionBlockedAt50C6C7D
	defer func() { sessionBlockedAt50C6C7D = old }()
	expected := map[string]string{
		"100002": "soga key 到期", "100003": "soga key 对接地址错误",
		"100004": "soga key 后端类型错误", "100005": "soga key 用户数超出限制",
		"100006": "unknown error", "100007": "timestamp error", "100008": "soga key 错误",
		"100009": "soga key 未激活，请先绑定机器人激活", "200001": "unknown error",
	}
	handlers := euYBu_M()
	if len(handlers) != 11 {
		t.Fatal(len(handlers))
	}
	for _, before := range []bool{false, true} {
		for code, message := range expected {
			sessionBlockedAt50C6C7D = before
			err := handlers[code]()
			if err == nil || err.Error() != message || !SessionBlocked() {
				t.Fatalf("%s: error=%v blocked=%v", code, err, SessionBlocked())
			}
		}
		sessionBlockedAt50C6C7D = before
		if err := handlers["0"](); err != nil || SessionBlocked() != before {
			t.Fatal("success status must not reset the original global", err)
		}
	}
	delete(handlers, "0")
	if euYBu_M()["0"] == nil {
		t.Fatal("dispatcher map should be freshly constructed")
	}
}
