package muxerr

import (
	"bytes"
	"context"
	"crypto/rand"
	"net"
	"net/netip"
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
	"testing"
	"testing/synctest"
	"time"
	"unsafe"
)

var _ sess.KsKg3yls = (*SUfEmBl)(nil)

func TestNativeLegacyUDPWireAndAuthFailure(t *testing.T) {
	for _, method := range []ssmethod.C6q3Upu{"aes-128-gcm", "aes-256-gcm", "chacha20-ietf-poly1305"} {
		t.Run(string(method), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			u := legacyUser(t, method)
			users := &user.ZpNFqUVZp{GTFAWF: new(genmap.GTFAWF[int64, user.PgaNd3014U])}
			users.Add(u)
			server, e := net.ListenPacket("udp", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer server.Close()
			server.SetDeadline(time.Now().Add(3 * time.Second))
			client, e := net.ListenPacket("udp", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer client.Close()
			client.SetDeadline(time.Now().Add(3 * time.Second))
			c := OlQGbrpyFga(server, method, sess.YCcacra(ctx), users, limit.Azb3Zr(ctx, &config.Nl62Lnd_{}))
			wire := func(p []byte) []byte {
				salt := make([]byte, len(u.PSK()))
				rand.Read(salt)
				a, e := u.NewAEAD(salt)
				if e != nil {
					t.Fatal(e)
				}
				return append(salt, a.Seal(nil, make([]byte, 12), p, nil)...)
			}
			target := append([]byte{3, 11}, []byte("example.org")...)
			target = append(target, 0, 53)
			target = append(target, []byte("query")...)
			p := make([]byte, 2048)
			bad := wire(target)
			bad[len(bad)-1] ^= 1
			client.WriteTo(bad, server.LocalAddr())
			if n, u, a, d, e := c.ReadFromProxy(p); n != 0 || u != nil || a != nil || d != nil || e != nil {
				t.Fatal("original zero-result authentication failure", n, e)
			}
			valid := wire(target)
			client.WriteTo(valid, server.LocalAddr())
			n, base, addr, dest, e := c.ReadFromProxy(p)
			if e != nil || string(p[:n]) != "query" || base != u || dest.GetHost() != "example.org" || dest.GetPort() != 53 {
				t.Fatal("request", n, e, dest)
			}
			if n, e = c.WriteToProxy([]byte("reply"), u, addr, &net.UDPAddr{IP: net.IPv4(8, 8, 8, 8), Port: 53}); e != nil || n != 5 {
				t.Fatal("response", n, e)
			}
			n, _, e = client.ReadFrom(p)
			if e != nil {
				t.Fatal(e)
			}
			keyLen := len(u.PSK())
			a, e := u.NewAEAD(p[:keyLen])
			if e != nil {
				t.Fatal(e)
			}
			plain, e := a.Open(nil, make([]byte, 12), p[keyLen:n], nil)
			if e != nil || !bytes.Equal(plain, append([]byte{1, 8, 8, 8, 8, 0, 53}, []byte("reply")...)) {
				t.Fatal("response wire", e, plain)
			}
			// Address cache rejects a replaced pointer, but the independent IP candidate
			// cache still authenticates the old user until explicitly invalidated.
			replacement := legacyUser(t, method)
			users.Add(replacement)
			client.WriteTo(wire(target), server.LocalAddr())
			n, base, _, _, e = c.ReadFromProxy(p)
			if e != nil || n != 5 || base != u { t.Fatal("candidate cache remains independent", e, base) }
 c.fN2q7wJSa.RemoveUsers([]user.PgaNd3014U{u})
 client.WriteTo(wire(target), server.LocalAddr())
 n,base,_,_,e=c.ReadFromProxy(p)
 if e!=nil || n!=5 || base!=replacement {t.Fatal("explicit candidate invalidation",e,base)}
		})
	}
}
func TestNativeLegacyUDPOriginalCacheRotation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var layout SUfEmBl
		if unsafe.Sizeof(layout) != 120 || unsafe.Offsetof(layout.xxfcsG) != 80 || unsafe.Offsetof(layout.s3T3QxaU4Wf) != 96 {
			t.Fatal("RTTI")
		}
		u := legacyUser(t, "aes-128-gcm")
		users := &user.ZpNFqUVZp{GTFAWF: new(genmap.GTFAWF[int64, user.PgaNd3014U])}
		users.Add(u)
		c := OlQGbrpyFga(nil, "aes-128-gcm", nil, users, nil)
		addr := netip.MustParseAddrPort("127.0.0.1:7")
		c.cBywakP2a(addr, u)
		time.Sleep(2 * time.Minute)
		if c.hpzIczxIS(addr) != u || c.z7p8paU != nil {
			t.Fatal("rotation must be strict After")
		}
		time.Sleep(time.Nanosecond)
		if c.hpzIczxIS(addr) != u || c.z7p8paU[addr] != u || c.xxfcsG[addr] != u {
			t.Fatal("previous generation promotion")
		}
		users.Delete(u.GetId())
		if c.hpzIczxIS(addr) != nil {
			t.Fatal("deleted user cache")
		}
	})
}
