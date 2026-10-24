package J1P4KNvfI_J

import (
	"net"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/NeGeSsM"
	"recovered.local/soga/cNT24rG4F"
)

// RReCtt 对应0x187f780；首字母导出以表达原跨包调用。
func (s *ZzkV7_5l) RReCtt(network string, dest *G6bHD5JA.AHU2MRmjzQu, ip net.IP) {
	if !cNT24rG4F.AccessLogEnabled() || dest == nil || s.PayAlbdv6 == nil {
		return
	}
	source := ""
	if s.KeHzUwuXWh0 != nil {
		source = s.KeHzUwuXWh0.String()
	}
	target := ""
	if ip != nil {
		target = ip.String()
	} else if first := dest.GetFirstIp(); first != nil {
		target = first.String()
	}
	NeGeSsM.IPblyWsS(network, s.IcoyFqRBU, s.PayAlbdv6.GetId(), source, dest.GetHost(), dest.GetPort(), target)
}

// WriteAccessLog 对应0x187f4c0。
func (s *ZzkV7_5l) WriteAccessLog(network string) { s.RReCtt(network, s.Iayb4V, s.IksfOg) }

// WriteAccessLogTo 对应0x187f520。
func (s *ZzkV7_5l) WriteAccessLogTo(network string, dest *G6bHD5JA.AHU2MRmjzQu) {
	s.RReCtt(network, dest, nil)
}

// WriteAccessLogUDPAddr 对应0x187f580，仍仅接受具体TCPAddr或UDPAddr。
func (s *ZzkV7_5l) WriteAccessLogUDPAddr(address net.Addr) {
	if !cNT24rG4F.AccessLogEnabled() || s.PayAlbdv6 == nil {
		return
	}
	ip, port, ok := sessionAddressParts(address)
	if !ok {
		return
	}
	source := ""
	if s.KeHzUwuXWh0 != nil {
		source = s.KeHzUwuXWh0.String()
	}
	NeGeSsM.IPblyWsS("udp", s.IcoyFqRBU, s.PayAlbdv6.GetId(), source, ip.String(), port, ip.String())
}
