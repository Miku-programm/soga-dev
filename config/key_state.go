package config

// 授权检查已移除：原 Q0SDIN（0x1341fe0）状态机及其不一致时的
// 异步退出路径一并去除，现在始终返回已授权状态（false）。
// 原状态变量与 T15Jifo 保留，用于记录配置中 soga_key 的载入。
var keyStateAt50C6C48 bool
var keyStateMirrorAt3787300 = true
var keyStateMarkerAt3787800 uint64 = 0x5a4b3c2d5a4b3c2d

// FrFQQhzTR5fY（0x1341960）的混淆路径只读取此字符串。
func FrFQQhzTR5fY() string {
	return configAt5096DC0.RM5SFRF
}

// T15Jifo（0x1340cc0）保留原状态写入顺序。
func T15Jifo(value bool) {
	keyStateAt50C6C48 = value
	keyStateMirrorAt3787300 = !value
	if value {
		keyStateMarkerAt3787800 = 0xa5b6c7d8a5b6c7d8
	} else {
		keyStateMarkerAt3787800 = 0x5a4b3c2d5a4b3c2d
	}
}

// Q0SDIN 原为授权状态机（0x1341fe0）；授权检查移除后始终返回已授权（false）。
func Q0SDIN() bool {
	return false
}
