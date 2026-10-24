// Restored against IDA 0x1b279e0, 0x1b27a00.
// Source comparison: github.com/Ieooo/clash/transport/ssr/obfs (GPL-3.0).
// Original dependency revision is unproven; names/layout and deviations follow target evidence.
package GtTssd5z

import "net"

type gVNW1qyLB struct{}

func init() {
	register("plain", c2yO0k0, 0)
}

func c2yO0k0(b *Ja7FBino7HL) AGNBYcIkk {
	return &gVNW1qyLB{}
}

func (p *gVNW1qyLB) StreamConn(c net.Conn) net.Conn { return c }
