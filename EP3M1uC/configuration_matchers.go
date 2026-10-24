package EP3M1uC

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/xx5fRyrFw"
)

// initializeMatchers 为分析者提取的源码函数，完整对应原规则初始化
// 0x1b86b12 之后的逻辑；出站初始化成功后才会执行这一段。
func (rule *DgAjL9UN) initializeMatchers() error {
	if len(rule.UwkgrlFf8) == 0 {
		return errors.New("please configure at least one outs")
	}
	if len(rule.OZmvCwn) == 0 && rule.Ehbw9wa == "" {
		return errors.New("please configure at least one rules or protocol")
	}
	for _, protocol := range splitConfigurationList(strings.ToLower(rule.Ehbw9wa)) {
		if protocol != "tcp" && protocol != "udp" {
			return fmt.Errorf("invalid route protocol: %v, only support tcp/udp", rule.Ehbw9wa)
		}
	}
	rule.yr2tuiHRVF = xx5fRyrFw.XSe2ktMsZnSQ(rule.Ehbw9wa)
	if len(rule.OZmvCwn) == 0 {
		rule.__ev4XIbr = true
	}
	var ips, ports []string
	var nodes []int
	for _, pattern := range rule.OZmvCwn {
		switch {
		case pattern == "" || pattern == "*":
			rule.__ev4XIbr = true
		case strings.HasPrefix(pattern, "ip:"), strings.HasPrefix(pattern, "geoip:"):
			ips = append(ips, pattern)
		case strings.HasPrefix(pattern, "port:"):
			ports = append(ports, pattern)
		case strings.HasPrefix(pattern, "node_id:"):
			for _, value := range splitConfigurationList(pattern[8:]) {
				if node, err := strconv.Atoi(value); err == nil {
					nodes = append(nodes, node)
				}
			}
		default:
			rule.hSORasT5h = append(rule.hSORasT5h, xx5fRyrFw.NewRouteMatcher(pattern))
		}
	}
	rule.cPaGes = xx5fRyrFw.HcZBFktYT(ips)
	rule.aru24wr = xx5fRyrFw.SekQr0GHk(ports)
	rule.ov3c4tIGkqkb = G7A7csgrs62p.PQQKui[int]()
	for _, node := range nodes {
		rule.ov3c4tIGkqkb.Add(node)
	}
	return nil
}

// vKl1IHr3.DAaBs7_C@0xb7d720：两侧去空白并丢弃空项，保留重复项。
func splitConfigurationList(value string) []string {
	var result []string
	for _, item := range strings.Split(strings.TrimSpace(value), ",") {
		if item = strings.TrimSpace(item); item != "" {
			result = append(result, item)
		}
	}
	return result
}
