package routing

import (
	"errors"
	"fmt"
	"net"
)

// initializeAddresses names the validation/address section of o7v2ON9D
// at 0x1b7c4c0, before its outbound-adapter switch. Empty replacement fields
// retain previously parsed addresses, as the original assignments do.
func (out *GyyO51) initializeAddresses() error {
	if out.NyLbiXb4OrlL != "direct" {
		if out.AHTlIEynXL == "" {
			return errors.New("please set out server")
		}
		if out.NyLbiXb4OrlL != "redirect" && out.JO1YIpXk == 0 {
			return errors.New("please set out port")
		}
		if uint(out.JO1YIpXk) > 65535 {
			return fmt.Errorf("out port invalid: %v", out.JO1YIpXk)
		}
	}
	if out.JaH5hv != "ipv4_first" && out.JaH5hv != "ipv6_first" {
		return fmt.Errorf("unknown listen_strategy: %v", out.JaH5hv)
	}
	if out.DM8KdAE6ZiZu != "" {
		ip := net.ParseIP(out.DM8KdAE6ZiZu)
		if ip == nil {
			return fmt.Errorf("listen is not ip: %v", out.DM8KdAE6ZiZu)
		}
		if ip.To4() != nil {
			if out.ZK61Na2nsqVQ == "" {
				out.ZK61Na2nsqVQ = out.DM8KdAE6ZiZu
			}
		} else if out.LXpap7F == "" {
			out.LXpap7F = out.DM8KdAE6ZiZu
		}
	}
	if out.ZK61Na2nsqVQ != "" {
		ip := net.ParseIP(out.ZK61Na2nsqVQ)
		if ip == nil || ip.To4() == nil {
			return fmt.Errorf("listen_ipv4 is not ipv4: %v", out.ZK61Na2nsqVQ)
		}
		out.iCySnjtd = ip
	}
	if out.LXpap7F != "" {
		ip := net.ParseIP(out.LXpap7F)
		if ip == nil || ip.To4() != nil {
			return fmt.Errorf("listen_ipv6 is not ipv6: %v", out.LXpap7F)
		}
		out.xk9pAT = ip
	}
	return nil
}
