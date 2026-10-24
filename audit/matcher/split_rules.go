package matcher

import "strings"

// Split 对应 0xebadc0。去除两端空白并丢弃空行；前缀大小写敏感且保留，
// 其它规则全部放入第一组。原返回三个非 nil 切片，未排序、去重。
func (r VHWq49iRat) Split() (domains, ips, ports []string) {
	domains = []string{}
	ips = []string{}
	ports = []string{}
	for _, s := range r {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		switch {
		case strings.HasPrefix(s, "ip:"), strings.HasPrefix(s, "geoip:"):
			ips = append(ips, s)
		case strings.HasPrefix(s, "port:"):
			ports = append(ports, s)
		default:
			domains = append(domains, s)
		}
	}
	return
}
