package relay

import (
	"io"
	"net/http"
	"net/url"
)

// RequestFromRecoveredCall 是恢复工程跨包可编译性桥接，不新增原业务函数。
// 原调用点可直接访问私有 vxLBn4qabL9L，此桥完整转发原六个参数和返回值。
func RequestFromRecoveredCall(address, method string, query url.Values, body io.Reader, headers http.Header, contentType string) (*http.Response, error) {
	return vxLBn4qabL9L(address, method, query, body, headers, contentType)
}
