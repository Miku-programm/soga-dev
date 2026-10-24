package internet

import (
	"net"
	"strings"
)

// 原 RTTI 0x35f2bd0，24 字节；此样本确实只保留这一个字段。
type SocketConfig struct {
	TrustedXForwardedFor []string `protobuf:"bytes,23,rep,name=trusted_x_forwarded_for,json=trustedXForwardedFor,proto3" json:"trusted_x_forwarded_for,omitempty"`
}

// IsValidHTTPHost 0x1250900。含冒号时忽略 SplitHostPort 的错误。
func IsValidHTTPHost(request, expected string) bool {
	request = strings.ToLower(request)
	expected = strings.ToLower(expected)
	if strings.Contains(request, ":") {
		request, _, _ = net.SplitHostPort(request)
	}
	return request == expected
}
