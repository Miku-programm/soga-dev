package JcmlVIK0T0i

import "errors"

// 错误文本由原 init 的纯解码器恢复；变量按原全局地址命名。
var (
	ErrorAt5097E00 = errors.New("proxyproto: can't read version 1 header")
	ErrorAt5097E10 = errors.New("proxyproto: version 1 header must be 107 bytes or less")
	ErrorAt5097E20 = errors.New("proxyproto: version 1 header is invalid, must end with \\r\\n")
	ErrorAt5097E30 = errors.New("proxyproto: can't read proxy protocol version and command")
	ErrorAt5097E40 = errors.New("proxyproto: can't read address family or protocol")
	ErrorAt5097E50 = errors.New("proxyproto: can't read length")
	ErrorAt5097E60 = errors.New("proxyproto: can't resolve source Unix address")
	ErrorAt5097E70 = errors.New("proxyproto: can't resolve destination Unix address")
	ErrorAt5097E80 = errors.New("proxyproto: proxy protocol signature not present")
	ErrorAt5097E90 = errors.New("proxyproto: unknown proxy protocol version")
	ErrorAt5097EA0 = errors.New("proxyproto: unsupported proxy protocol version and command")
	ErrorAt5097EB0 = errors.New("proxyproto: unsupported address family and protocol")
	ErrorAt5097EC0 = errors.New("proxyproto: invalid length")
	ErrorAt5097ED0 = errors.New("proxyproto: invalid address")
	ErrorAt5097EE0 = errors.New("proxyproto: invalid port number")
	ErrorAt5097EF0 = errors.New("proxyproto: upstream connection sent PROXY header but isn't allowed to send one")
	ErrorAt5097F00 = errors.New("proxyproto: truncated TLV")
	ErrorAt5097F10 = errors.New("proxyproto: malformed TLV Value")
	ErrorAt5097F20 = errors.New("proxyproto: incompatible TLV type")
	ErrorAt5097F30 = errors.New("proxyproto: uint16 overflow")
)
