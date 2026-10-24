package ssmethod

import "strings"

// GBXRYZ 对应 0x10a4fc0；0x606740 为 strings.ToLower，原实现不去空白。
func GBXRYZ(value string) OdAX8q {
	switch strings.ToLower(value) {
	case "http", "simple_obfs_http":
		return "http"
	case "tls", "simple_obfs_tls":
		return "tls"
	default:
		return ""
	}
}
