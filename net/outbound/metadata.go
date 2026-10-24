package outbound

import (
	"net"
	"strconv"
)

// RemoteAddress restores 0x1933380. Host takes precedence over destination IP.
func (metadata *D0mLgQF) RemoteAddress() string {
	host := metadata.WJ95LyOL
	if host == "" {
		host = metadata.ZouEpmjw8.String()
	}
	return net.JoinHostPort(host, strconv.FormatUint(uint64(metadata.JRXMVvg), 10))
}
