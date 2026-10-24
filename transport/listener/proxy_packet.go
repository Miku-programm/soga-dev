package listener

import (
	"context"
	"github.com/jellydator/ttlcache/v3"
	"net"
	B "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	P "recovered.local/soga/transport/proxyproto"
	C "recovered.local/soga/config"
	"time"
)

// EuI1uJ2zx 保留样本 RTTI 0x36f0f98 的 40 字节布局。
type EuI1uJ2zx struct {
	WCApOca6IQW7
	akXo7g4lc bool
	sUKRhMwzm *ttlcache.Cache[string, net.Addr]
	aLeZsGA   *ttlcache.Cache[string, net.Addr]
}

func KlRWVpX39J(ctx context.Context, conn net.PacketConn, force bool) net.PacketConn { // 0x18bbee0
	ttl := time.Duration(C.UDPTimeoutMinutes()) * time.Minute
	forward := ttlcache.New[string, net.Addr](ttlcache.WithTTL[string, net.Addr](ttl))
	reverse := ttlcache.New[string, net.Addr](ttlcache.WithTTL[string, net.Addr](ttl))
	go func() { // 0x18c2920
		ticker := time.NewTicker(ttl)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				forward.DeleteExpired()
				reverse.DeleteExpired()
			}
		}
	}()
	return &EuI1uJ2zx{conn, force, forward, reverse}
}

func (c *EuI1uJ2zx) x2OW30(addr net.Addr) (net.Addr, bool) { // 0x18bc140
	key := addr.String()
	item := c.sUKRhMwzm.Get(addr.String())
	if item == nil {
		return nil, false
	}
	peer := item.Value()
	c.sUKRhMwzm.Set(key, peer, 0)
	c.aLeZsGA.Set(peer.String(), addr, 0)
	return peer, true
}

func (c *EuI1uJ2zx) pEL3YD3aSM2r(src, dst net.Addr) { // 0x18bc2a0
	c.sUKRhMwzm.Set(src.String(), dst, 0)
	c.aLeZsGA.Set(dst.String(), src, 0)
}

func (c *EuI1uJ2zx) ReadFrom(p []byte) (int, net.Addr, error) { // 0x18bc3a0
	for {
		n, addr, err := c.WCApOca6IQW7.ReadFrom(p)
		if err != nil {
			return n, addr, err
		}
		r := B.NewPrefilledForProxy(p[:n])
		header, _ := P.LavOMmUY2(r)
		if header != nil {
			if header.Dafxa_BOdZdw == nil {
				continue
			}
			remaining := r.Buffered()
			peer, err := G.LQYEqaxN(header.Dafxa_BOdZdw)
			if err != nil {
				continue
			}
			c.pEL3YD3aSM2r(addr, peer)
			copy(p, p[n-remaining:n])
			return remaining, peer, nil
		}
		if peer, ok := c.x2OW30(addr); ok {
			return n, peer, nil
		}
		if !c.akXo7g4lc {
			return n, addr, nil
		}
	}
}

func (c *EuI1uJ2zx) b_LUR8(addr net.Addr) (net.Addr, bool) { // 0x18bc6c0
	key := addr.String()
	item := c.aLeZsGA.Get(key)
	if item == nil {
		return nil, false
	}
	peer := item.Value()
	c.aLeZsGA.Set(key, peer, 0)
	c.sUKRhMwzm.Set(peer.String(), addr, 0)
	return peer, true
}

func (c *EuI1uJ2zx) WriteTo(p []byte, addr net.Addr) (int, error) { // 0x18bc800
	if peer, ok := c.b_LUR8(addr); ok {
		return c.WCApOca6IQW7.WriteTo(p, peer)
	}
	return c.WCApOca6IQW7.WriteTo(p, addr)
}

func (c *EuI1uJ2zx) GetRawPacketConn() net.PacketConn { return c.WCApOca6IQW7 } // 0x18bc900
