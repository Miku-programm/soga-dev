package MFFa9yT

import (
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/nzATvajt"
)

// TestAddr 对应0x1d44380：用户豁免、白名单优先，然后面板探测器、
// 本地域名/IP/私有地址/端口拦截。检查所有解析IP，不只首个。
func (b *JWN154qJ4) TestAddr(addr *G6bHD5JA.AHU2MRmjzQu, user nzATvajt.PgaNd3014U) bool {
	if cNT24rG4F.IsInspectionExempt(user.GetId()) {
		return false
	}
	host := ""
	if !addr.IsIP() {
		host = addr.GetHost()
	}
	ips, port := addr.GetIPs(), addr.GetPort()
	white := b._IoxtH.Load()
	if host != "" {
		for _, r := range white.a73gMUVz {
			if r.Test(host) {
				return false
			}
		}
	}
	for _, ip := range ips {
		if white.fj_WUx4re.TestIP(ip) {
			return false
		}
	}
	if white.xDjyw2UC.Match(port) {
		return false
	}
	black := b.oIgEuicmPL.Load()
	if host != "" {
		for _, r := range *b.e3JHP2MzU.Load() {
			if r.Test(host) {
				users, _ := b.qE9aU9.GetOrSet(r.GetId(), &G7A7csgrs62p.Fxs7SY[int64]{})
				users.Add(user.GetId())
				return true
			}
		}
		for _, r := range black.a73gMUVz {
			if r.Test(host) {
				return true
			}
		}
	}
	for _, ip := range ips {
		if black.fj_WUx4re.TestIP(ip) {
			return true
		}
		_, _, banPrivate := cNT24rG4F.InspectionConfiguration()
		if banPrivate && (ip.IsLoopback() || ip.IsPrivate()) {
			return true
		}
	}
	return black.xDjyw2UC.Match(port)
}

// TestPacket 对应0x1d44900；仅面板命中的规则写入上报集合，本地规则命中不写入。
func (b *JWN154qJ4) TestPacket(data []byte, user nzATvajt.PgaNd3014U) bool {
	enabled, maxLength, _ := cNT24rG4F.InspectionConfiguration()
	if !enabled {
		return false
	}
	if cNT24rG4F.IsInspectionExempt(user.GetId()) {
		return false
	}
	if len(data) > maxLength {
		data = data[:maxLength]
	}
	for _, r := range *b.e3JHP2MzU.Load() {
		if r.MatchPacket(data) {
			users, _ := b.qE9aU9.GetOrSet(r.GetId(), &G7A7csgrs62p.Fxs7SY[int64]{})
			users.Add(user.GetId())
			return true
		}
	}
	for _, r := range b.oIgEuicmPL.Load().a73gMUVz {
		if r.MatchPacket(data) {
			return true
		}
	}
	return false
}
