package VlSfV5

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"recovered.local/soga/A818JHb"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/NtxdAE"
	"recovered.local/soga/Ualx5UM_D6I"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/w2YaTd9_"
	"sync"
	"time"
)

type WCApOca6IQW7 = net.PacketConn

// SUfEmBl is the original 120-byte UDP state. The address-key RTTI (32 bytes)
// is net/netip.AddrPort, whose original package path is COgCoHHQaf.
type SUfEmBl struct {
	WCApOca6IQW7
	z2glYbdwcY7n w2YaTd9_.C6q3Upu
	fN2q7wJSa    *J1P4KNvfI_J.EafxyUZD
	dQxSInY      *nzATvajt.ZpNFqUVZp
	cGUPMa       *NtxdAE.M1Qn9BA
	ygT83ZmW     [16]byte
	pTtnsJ4pO    sync.Mutex
	xxfcsG       map[netip.AddrPort]*nzATvajt.WtK9HQSy66
	z7p8paU      map[netip.AddrPort]*nzATvajt.WtK9HQSy66
	s3T3QxaU4Wf  time.Time
}

func OlQGbrpyFga(conn net.PacketConn, method w2YaTd9_.C6q3Upu, candidates *J1P4KNvfI_J.EafxyUZD, users *nzATvajt.ZpNFqUVZp, invalid *NtxdAE.M1Qn9BA) *SUfEmBl { //0x1c60d00
	return &SUfEmBl{WCApOca6IQW7: conn, z2glYbdwcY7n: method, fN2q7wJSa: candidates, dQxSInY: users, cGUPMa: invalid, xxfcsG: make(map[netip.AddrPort]*nzATvajt.WtK9HQSy66), s3T3QxaU4Wf: time.Now().Add(2 * time.Minute)}
}
func (c *SUfEmBl) hpzIczxIS(addr netip.AddrPort) *nzATvajt.WtK9HQSy66 { //0x1c60ea0
	c.pTtnsJ4pO.Lock()
	now := time.Now()
	if now.After(c.s3T3QxaU4Wf) {
		c.z7p8paU = c.xxfcsG
		c.xxfcsG = make(map[netip.AddrPort]*nzATvajt.WtK9HQSy66)
		c.s3T3QxaU4Wf = now.Add(2 * time.Minute)
	}
	u, ok := c.xxfcsG[addr]
	if !ok {
		if old, found := c.z7p8paU[addr]; found {
			u = old
			c.xxfcsG[addr] = old
		}
	}
	c.pTtnsJ4pO.Unlock()
	if u == nil {
		return nil
	}
	current, ok := c.dQxSInY.Get(u.GetId())
	if !ok || current != u {
		return nil
	}
	return u
}
func (c *SUfEmBl) cBywakP2a(addr netip.AddrPort, u *nzATvajt.WtK9HQSy66) { //0x1c61120
	c.pTtnsJ4pO.Lock()
	c.xxfcsG[addr] = u
	c.pTtnsJ4pO.Unlock()
}

// ys3P2cu 0x1c61240, closures 0x1c64180/0x1c64080/0x1c63ee0.
func (c *SUfEmBl) ys3P2cu(packet []byte, addr netip.AddrPort, ip string, dst Ualx5UM_D6I.Mf3phSOVqgf) (*nzATvajt.WtK9HQSy66, error) {
	keyLen, err := c.z2glYbdwcY7n.KeyLen()
	if err != nil {
		return nil, err
	}
	if len(packet) < keyLen+16 {
		return nil, fmt.Errorf("ss udp packet len not enough: %v", len(packet))
	}
	salt := packet[:keyLen]
	if !A818JHb.HFlnXWty(salt) {
		return nil, errorAt5098400
	}
	nonce := make([]byte, 12)
	dst.Reset()
	state := w2YaTd9_.SaYhx5xog9()
	check := func(u *nzATvajt.WtK9HQSy66) bool {
		a, err := u.NewAEADWith(state, salt)
		if err != nil {
			muc279Mpa6S4.TRTbrquD(4, err)
			return false
		}
		_, err = dst.OpenWrite(a, nonce[:a.NonceSize()], packet[keyLen:], nil)
		return err == nil
	}
	if cached := c.hpzIczxIS(addr); cached != nil && check(cached) {
		return cached, nil
	}
	seen := make(map[int64]struct{})
	var selected *nzATvajt.WtK9HQSy66
	c.fN2q7wJSa.RangeUser(ip, func(base nzATvajt.PgaNd3014U) bool {
		u := base.(*nzATvajt.WtK9HQSy66)
		seen[u.GetId()] = struct{}{}
		if check(u) {
			selected = u
			return false
		}
		return true
	})
	if selected == nil {
		count := 0
		c.dQxSInY.Range(func(_ int64, base nzATvajt.PgaNd3014U) bool {
			u := base.(*nzATvajt.WtK9HQSy66)
			if _, ok := seen[u.GetId()]; ok {
				return true
			}
			count++
			if count%50 == 0 && c.cGUPMa.IsForbidden(ip) {
				return false
			}
			if check(u) {
				selected = u
				return false
			}
			return true
		})
	}
	if selected == nil {
		return nil, errorAt5098410
	}
	c.fN2q7wJSa.AddUser(ip, selected, 0)
	c.cBywakP2a(addr, selected)
	return selected, nil
}

// ReadFromProxy 0x1c619c0 consumes one datagram per invocation. Invalid/auth-
// limited packets produce the original zero result with nil error, not a retry.
func (c *SUfEmBl) ReadFromProxy(p []byte) (int, nzATvajt.PgaNd3014U, net.Addr, *G6bHD5JA.AHU2MRmjzQu, error) {
	n, addr, err := c.WCApOca6IQW7.ReadFrom(p)
	if n <= 0 {
		return 0, nil, addr, nil, err
	}
	udp, err := G6bHD5JA.LQYEqaxN(addr)
	if err != nil {
		return 0, nil, nil, nil, err
	}
	ip := udp.IP.String()
	if c.cGUPMa.IsForbidden(ip) {
		return 0, nil, nil, nil, nil
	}
	b := D_zCag1qTV.S4OQXWIZ(8192)
	defer b.Free()
	u, err := c.ys3P2cu(p[:n], udp.AddrPort(), ip, b)
	if err != nil {
		if errors.Is(err, errorAt5098410) {
			c.cGUPMa.Invalid(ip)
		}
		return 0, nil, nil, nil, nil
	}
	c.cGUPMa.Valid(ip)
	if !u.AllowAndAddUp(n) {
		muc279Mpa6S4.V8GN5Stk(-4, "user %v speed limited", u.GetId())
		return 0, u, addr, nil, nil
	}
	dest, err := G6bHD5JA.SzhW4tOzUZg(b)
	if err != nil {
		return 0, nil, nil, nil, nil
	}
	n, _ = b.Read(p)
	return n, u, addr, dest, nil
}
func (c *SUfEmBl) WriteToProxy(p []byte, base nzATvajt.PgaNd3014U, client, remote net.Addr) (int, error) { //0x1c62260
	dest, err := G6bHD5JA.RNFJeaRY(remote)
	if err != nil {
		return 0, err
	}
	u := base.(*nzATvajt.WtK9HQSy66)
	keyLen, err := c.z2glYbdwcY7n.KeyLen()
	if err != nil {
		return 0, err
	}
	b := D_zCag1qTV.S4OQXWIZ(8192)
	defer b.Free()
	b.WriteRandom(keyLen)
	if _, err = G6bHD5JA.DdYroT1n(b, dest); err != nil {
		return 0, err
	}
	a, err := u.NewAEAD(b.BytesTo(keyLen))
	if err != nil {
		return 0, err
	}
	b.SealWriteWithStart(a, keyLen, c.ygT83ZmW[:a.NonceSize()], p, nil)
	if !u.AllowAndAddDown(b.Len()) {
		muc279Mpa6S4.V8GN5Stk(-4, "user %v speed limited", u.GetId())
		return 0, nil
	}
	if _, err = c.WCApOca6IQW7.WriteTo(b.Bytes(), client); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (c *SUfEmBl) Free() { HxKViEp2a.PpnNLr_pf7He(c.WCApOca6IQW7) } //0x1c62860
