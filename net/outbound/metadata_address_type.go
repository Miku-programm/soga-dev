package outbound

// AddrType restores 0x19335c0. A nonnil malformed IP is classified as IPv6;
// a nil IP or nonempty host is classified as a domain name.
func (m *D0mLgQF) AddrType() int {
	if m.WJ95LyOL != "" || m.ZouEpmjw8 == nil {
		return 3
	}
	if m.ZouEpmjw8.To4() != nil {
		return 1
	}
	return 4
}
