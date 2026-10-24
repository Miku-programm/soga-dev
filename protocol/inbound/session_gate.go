// 根据原 euYBu_M@0x1db0960 及十一处理器的 IDA 指令恢复。
// 原 init 区间 0x1d2429e..0x1d28e72 的实际 seed/capture 已受限仿真求值；
// 配套证据 session-initialization.json / session-handler-traces.json / session-dispatch-map.json。
package inbound

import (
	"errors"
	"recovered.local/soga/config"
)

// 原全局字节 0x50c6c7d，初值为零。保留原非原子读写，未增加重置行为。
var sessionBlockedAt50C6C7D bool

// SessionBlocked 是新增跨包访问桥接，对应中央会话处理器对 0x50c6c7d 的内联读取。
func SessionBlocked() bool { return sessionBlockedAt50C6C7D }

func euYBu_M() map[string]func() error {
	return map[string]func() error{
		"0": euYBu_Mfunc1, "100001": euYBu_Mfunc2, "100002": euYBu_Mfunc3,
		"100003": euYBu_Mfunc4, "100004": euYBu_Mfunc5, "100005": euYBu_Mfunc6,
		"100006": euYBu_Mfunc7, "100007": euYBu_Mfunc8, "100008": euYBu_Mfunc9,
		"100009": euYBu_Mfunc10, "200001": euYBu_Mfunc11,
	}
}

func euYBu_Mfunc1() error { return nil } // 0x1eb6ce0；不清除旧的禁止状态。

// 0x1e8bd20：先更新原 key 状态，然后读取 key。此处理器不写会话禁止字节。
func euYBu_Mfunc2() error {
	config.T15Jifo(true)
	if config.FrFQQhzTR5fY() != "" {
		return errors.New("soga key 错误")
	}
	return nil
}

func euYBu_Mfunc3() error { sessionBlockedAt50C6C7D = true; return errors.New("soga key 到期") } // 0x1eb6d00
func euYBu_Mfunc4() error {
	sessionBlockedAt50C6C7D = true
	return errors.New("soga key 对接地址错误")
} // 0x1eb7540
func euYBu_Mfunc5() error {
	sessionBlockedAt50C6C7D = true
	return errors.New("soga key 后端类型错误")
} // 0x1e8b4c0
func euYBu_Mfunc6() error {
	sessionBlockedAt50C6C7D = true
	return errors.New("soga key 用户数超出限制")
}                         // 0x1e8af00
func euYBu_Mfunc7() error { sessionBlockedAt50C6C7D = true; return errors.New("unknown error") }   // 0x1e8aa20
func euYBu_Mfunc8() error { sessionBlockedAt50C6C7D = true; return errors.New("timestamp error") } // 0x1e8a4a0
func euYBu_Mfunc9() error { sessionBlockedAt50C6C7D = true; return errors.New("soga key 错误") }     // 0x1e89f20
func euYBu_Mfunc10() error {
	sessionBlockedAt50C6C7D = true
	return errors.New("soga key 未激活，请先绑定机器人激活")
}                          // 0x1eb7da0
func euYBu_Mfunc11() error { sessionBlockedAt50C6C7D = true; return errors.New("unknown error") } // 0x1e896e0
