package config

import "os"

// VSHlrRuK7_ 对应 0x1324240。环境变量名由原解码器恢复；
// 只接受精确的 "true"，不做大小写转换或空白处理。
func VSHlrRuK7_() bool {
	return os.Getenv("IS_DOCKER") == "true"
}
