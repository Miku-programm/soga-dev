package J1P4KNvfI_J

import (
	"context"
	"net"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/cNT24rG4F"
	"strconv"
	"time"
)

// G3dXf_C86Ujq 对应0x1886c40。缓冲由调用方提供，不在此归还。
func G3dXf_C86Ujq(ctx context.Context, local NTe8CQ, remote net.PacketConn, buf []byte, session *ZzkV7_5l) (err error) {
	// 原func1是普通闭包调用recover helper，保留层级，不扩展其捕获能力。
	defer func() { HxKViEp2a.SlG_Y_mg("handle remote udp panic") }()
	logged := make(map[string]struct{})
	if session != nil && session.Iayb4V != nil {
		if ip := session.Iayb4V.GetFirstIp(); ip != nil {
			logged[net.JoinHostPort(ip.String(), strconv.Itoa(session.Iayb4V.GetPort()))] = struct{}{}
		}
	}
	for {
		select {
		case <-ctx.Done():
			return HxKViEp2a.ErrorAt5097800
		default:
		}
		// 原代码在读取local之前更新的是remote的读截止时间。
		_ = remote.SetReadDeadline(time.Now().Add(time.Duration(cNT24rG4F.UDPTimeoutMinutes()) * time.Minute))
		n, address, err := local.ReadFromProxy(buf)
		if err != nil {
			return err
		}
		if session != nil && cNT24rG4F.AccessLogEnabled() {
			key := address.String()
			if _, ok := logged[key]; !ok {
				if len(logged) >= 1024 {
					clear(logged)
				}
				logged[key] = struct{}{}
				session.WriteAccessLogUDPAddr(address)
			}
		}
		_ = remote.SetWriteDeadline(time.Now().Add(time.Minute))
		if _, err = remote.WriteTo(buf[:n], address); err != nil {
			return err
		}
	}
}
