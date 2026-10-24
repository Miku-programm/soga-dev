package BaEWsPy

import (
	"context"
	"errors"
	"fmt"
	"net"
	"recovered.local/soga/HxKViEp2a"
	"time"
)

// DAlA5L: 0x158a2a0。
func DAlA5L(ctx context.Context, cfg *PQL1aLrtju) *D23y2888a_ {
	rules := make(ZrMh8EfBQar, 0, len(cfg.MLo0o543OAaK))
	for _, rule := range cfg.MLo0o543OAaK {
		rules = append(rules, o4RFaBT0MwY3(rule, cfg.DsLwiF3YkJiu))
	}
	servers := make(GG2P4UAyQqI, 0, len(cfg.Zfr5Q8ybETV))
	for _, server := range cfg.Zfr5Q8ybETV {
		servers = append(servers, cDpHJRFxZ(server, cfg.DsLwiF3YkJiu))
	}
	var resolver Fr3ugVA_Lh4_ = servers
	if len(servers) == 0 {
		resolver = &BkMuy3LHBq{}
	} else if cfg.DMa5Vg24rVR {
		resolver = &by9pe4XG{servers, &BkMuy3LHBq{}}
	}
	return &D23y2888a_{nPknfPMR8: ctx, bebyxp: cfg, p_zkWSwY4Y: u838o9bPnv(ctx, cfg.C_TYygjEa0), a06M4WZc: rules, yrk3eTuNjOam: resolver, bkaoo_vYuCup: dnsTypes(string(cfg.X6T7XOmp))}
}

// 两处原内联分支位于 0x1586af0 与 0x158a6ae；未归一化输入字符串。
func dnsTypes(strategy string) []uint16 {
	switch strategy {
	case "ipv4_only":
		return []uint16{1}
	case "ipv6_only":
		return []uint16{28}
	case "ipv6_first":
		return []uint16{28, 1}
	default:
		return []uint16{1, 28}
	}
}

// Resolve: 0x158ad40；singleflight 闭包 0x1595420。
func (r *D23y2888a_) Resolve(host string) ([]net.IP, error) {
	if ips, ok := r.p_zkWSwY4Y.Get(host); ok {
		return ips, nil
	}
	if ip := net.ParseIP(host); ip != nil {
		return []net.IP{ip}, nil
	}
	value, err, _ := r.wYE6QCP.Do(host, func() (any, error) {
		if ips, ok := r.p_zkWSwY4Y.Get(host); ok {
			return ips, nil
		}
		for _, rule := range r.a06M4WZc {
			if rule.Test(host) {
				return r.ForceResolve(rule, rule.qBHWeHSCLI, host)
			}
		}
		return r.ForceResolve(r.yrk3eTuNjOam, r.bkaoo_vYuCup, host)
	})
	if err != nil {
		return nil, err
	}
	return value.([]net.IP), nil
}

// ForceResolve: 0x158af00；任务闭包 0x15955a0。
func (r *D23y2888a_) ForceResolve(resolver Fr3ugVA_Lh4_, types []uint16, host string) ([]net.IP, error) {
	if len(types) != 1 {
		ipv4, ipv6 := ipv4AvailableAt50C6D60.Load(), ipv6AvailableAt50C6D64.Load()
		if ipv4 == 0 || ipv6 == 0 {
			enabled := make([]uint16, 0, len(types))
			for _, typ := range types {
				if typ == 1 && ipv4 != 0 || typ == 28 && ipv6 != 0 {
					enabled = append(enabled, typ)
				}
			}
			if len(enabled) > 0 {
				types = enabled
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	completed := make(chan error, len(types))
	results := make([][]net.IP, len(types))
	for i, typ := range types {
		go func(i int, typ uint16) {
			ips, err := resolver.LookupDNSType(ctx, host, typ)
			if err == nil {
				results[i] = ips
			}
			completed <- err
		}(i, typ)
	}
	errs := make([]error, 0, len(types))
	for range types {
		if err := <-completed; err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == len(types) {
		return nil, HxKViEp2a.T2i1ItMaa(errs...)
	}
	ips := make([]net.IP, 0, len(types))
	for _, result := range results {
		if len(result) > 0 {
			ips = append(ips, result[0])
		}
	}
	r.p_zkWSwY4Y.AddIP(host, ips)
	return ips, nil
}

// LookupDNSType: 0x1592a80。原本的 local 分支明确调用 net.DefaultResolver。
func (*BkMuy3LHBq) LookupDNSType(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	switch typ {
	case 1:
		return net.DefaultResolver.LookupIP(ctx, "ip4", host)
	case 28:
		return net.DefaultResolver.LookupIP(ctx, "ip6", host)
	}
	return nil, fmt.Errorf("unknown ip type: %v", typ)
}

// LookupDNSType: 0x1593680；首个成功结果立即返回并取消其余请求。
func (servers GG2P4UAyQqI) LookupDNSType(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	if len(servers) == 0 {
		return nil, errors.New("no resolver")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	completed := make(chan srGFtPJWg1, len(servers))
	for _, server := range servers {
		go func(server Fr3ugVA_Lh4_) {
			ips, err := server.LookupDNSType(ctx, host, typ)
			completed <- srGFtPJWg1{ips, err}
		}(server)
	}
	errs := make([]error, 0, len(servers))
	for range servers {
		result := <-completed
		if result.fQrbEreES7Lj == nil {
			return result.z5nt2pc7e, nil
		}
		errs = append(errs, result.fQrbEreES7Lj)
	}
	return nil, HxKViEp2a.T2i1ItMaa(errs...)
}

// LookupDNSType: 0x1593de0；备用解析同时启动，但主解析成功时优先主结果。
func (r *by9pe4XG) LookupDNSType(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	completed := make(chan srGFtPJWg1, 1)
	go func() { ips, err := r.bkuuYtpyiQS.LookupDNSType(ctx, host, typ); completed <- srGFtPJWg1{ips, err} }()
	ips, err := r.okBXre3jIo.LookupDNSType(ctx, host, typ)
	if err == nil {
		return ips, nil
	}
	fallback := <-completed
	if fallback.fQrbEreES7Lj != nil {
		return nil, HxKViEp2a.T2i1ItMaa(err, fallback.fQrbEreES7Lj)
	}
	return fallback.z5nt2pc7e, nil
}
