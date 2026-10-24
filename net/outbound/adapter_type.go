package outbound

// String restores 0x1930020. Literal table: 0x216c0b0; five other values
// were recovered with bounded pure-decoder emulation (adapter-decoded-strings.json).
func (t ASFyBDZD) String() string {
	switch t {
	case 0:
		return "Direct"
	case 1:
		return "Reject"
	case 2:
		return "Shadowsocks"
	case 3:
		return "ShadowsocksR"
	case 4:
		return "Snell"
	case 5:
		return "Socks5"
	case 6:
		return "Http"
	case 7:
		return "Vmess"
	case 8:
		return "Trojan"
	case 9:
		return "Relay"
	case 10:
		return "Selector"
	case 11:
		return "Fallback"
	case 12:
		return "URLTest"
	case 13:
		return "LoadBalance"
	default:
		return "Unknown"
	}
}
