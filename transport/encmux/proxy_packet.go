package encmux

import (
	"context"
	"github.com/jellydator/ttlcache/v3"
	"net"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/transport/utlsconfig"
	"recovered.local/soga/transport/proxyproto"
	"time"
)

type WCApOca6IQW7 = net.PacketConn

// ARZKOqmN5e is the original 40-byte state at RTTI 0x36f0a48.
type ARZKOqmN5e struct {
	WCApOca6IQW7
	akXo7g4lc bool
	sUKRhMwzm *ttlcache.Cache[string, net.Addr]
	aLeZsGA   *ttlcache.Cache[string, net.Addr]
}

// Original initialized duration at 0x5048d08 is 300000000000 ns.
var proxyTTLAt5048D08 = 300 * time.Second

func SimI1bq(ctx context.Context, conn net.PacketConn, force bool) net.PacketConn { //0x11d9c60
	forward := ttlcache.New[string, net.Addr](ttlcache.WithTTL[string, net.Addr](proxyTTLAt5048D08))
	reverse := ttlcache.New[string, net.Addr](ttlcache.WithTTL[string, net.Addr](proxyTTLAt5048D08))
	go func() { //0x11e5c40, defer0x11e5d80
		ticker := time.NewTicker(proxyTTLAt5048D08)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				forward.DeleteExpired()
				reverse.DeleteExpired()
			case <-ctx.Done():
				return
			}
		}
	}()
	return &ARZKOqmN5e{WCApOca6IQW7: conn, akXo7g4lc: force, sUKRhMwzm: forward, aLeZsGA: reverse}
}
func (p *ARZKOqmN5e) ssa9qlNCrrBM(addr net.Addr) (net.Addr, bool) { //0x11d9ec0
	// The original calls String twice before the lookup, even for a miss.
	key := addr.String()
	item := p.sUKRhMwzm.Get(addr.String())
	if item == nil {
		return nil, false
	}
	value := item.Value()
	p.sUKRhMwzm.Set(key, value, ttlcache.DefaultTTL)
	p.aLeZsGA.Set(value.String(), addr, ttlcache.DefaultTTL)
	return value, true
}
func (p *ARZKOqmN5e) h5MasPA(from, to net.Addr) { //0x11da020
	p.sUKRhMwzm.Set(from.String(), to, ttlcache.DefaultTTL)
	p.aLeZsGA.Set(to.String(), from, ttlcache.DefaultTTL)
}
func (p *ARZKOqmN5e) ReadFrom(b []byte) (int, net.Addr, error) { //0x11da120
	for {
		n, from, err := p.WCApOca6IQW7.ReadFrom(b)
		if err != nil {
			return n, from, err
		}
		r := pool.NewPrefilledForProxy(b[:n])
		header, _ := proxyproto.LavOMmUY2(r)
		if header != nil {
			if header.Dafxa_BOdZdw != nil {
				size := r.Buffered()
				addr, err := utlsconfig.Q_ZqnPkUlb9(header.Dafxa_BOdZdw)
				if err == nil {
					p.h5MasPA(from, addr)
					copy(b, b[n-size:n])
					return size, addr, nil
				}
			}
			// Headers with no source (including LOCAL) and invalid source
			// conversion are discarded even when force is false.
			continue
		}
		if addr, ok := p.ssa9qlNCrrBM(from); ok {
			return n, addr, nil
		}
		if !p.akXo7g4lc {
			return n, from, nil
		}
	}
}
func (p *ARZKOqmN5e) eQq7GgK(addr net.Addr) (net.Addr, bool) { //0x11da440
	key := addr.String()
	item := p.aLeZsGA.Get(key)
	if item == nil {
		return nil, false
	}
	value := item.Value()
	p.aLeZsGA.Set(key, value, ttlcache.DefaultTTL)
	p.sUKRhMwzm.Set(value.String(), addr, ttlcache.DefaultTTL)
	return value, true
}
func (p *ARZKOqmN5e) WriteTo(b []byte, addr net.Addr) (int, error) { //0x11da580
	if mapped, ok := p.eQq7GgK(addr); ok {
		return p.WCApOca6IQW7.WriteTo(b, mapped)
	}
	return p.WCApOca6IQW7.WriteTo(b, addr)
}
func (p *ARZKOqmN5e) GetRawPacketConn() net.PacketConn { return p.WCApOca6IQW7 } //0x11da680
