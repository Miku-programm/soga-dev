package dns

import (
	"net"
	"recovered.local/soga/config"
)

// bQkmjIa: 0x158ab60。配置检查与原指令顺序一致。
func bQkmjIa(address net.Addr) net.IP {
	if address == nil || !config.AutoOutIPEnabled() {
		return nil
	}
	var ip net.IP
	switch address := address.(type) {
	case *net.UDPAddr:
		ip = address.IP
	case *net.TCPAddr:
		ip = address.IP
	}
	if ip == nil || ip.IsLoopback() || ip.IsUnspecified() {
		return nil
	}
	if v4 := ip.To4(); v4 != nil {
		return v4
	}
	return ip
}

// For: 0x158a860。按来源地址共享解析器，锁内再次查询以避免重复构造。
func (r *D23y2888a_) For(address net.Addr) *D23y2888a_ {
	ip := bQkmjIa(address)
	if ip == nil {
		return r
	}
	key := ip.String()
	if cached, ok := r.zCEngBnOF7.Get(key); ok {
		return cached
	}
	r.iAybdd2Iwc.Lock()
	defer r.iAybdd2Iwc.Unlock()
	if cached, ok := r.zCEngBnOF7.Get(key); ok {
		return cached
	}
	cfg := *r.bebyxp
	cfg.DsLwiF3YkJiu = ip
	child := DAlA5L(r.nPknfPMR8, &cfg)
	r.zCEngBnOF7.Set(key, child)
	return child
}
