package MFFa9yT

import (
	"context"
	"fmt"
	"net"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/muc279Mpa6S4"
	"sync/atomic"
	"time"
)

// 原RTTI0x36b3440：184字节，显示名oUi8Xtk3yv，实际包MFFa9yT。
type rWW6Rmakjzl struct {
	af0cE6r5P string
	tkRa6u    string
	rqmdTVGqu net.Dialer
	sDhm_cyQ1 atomic.Pointer[[]net.IP]
}
type fsqq14xvr struct {
	u4oajSYOgl net.Conn
	ri0FDK4z   error
}

// gzasC2TjpYF 0x1d6a6e0。静态Dialer模板0x215d160仅Timeout=5s。
func gzasC2TjpYF(ctx context.Context, dest string) func(context.Context, string, string) (net.Conn, error) {
	d := net.Dialer{Timeout: 5 * time.Second}
	host, port, err := net.SplitHostPort(dest)
	if err != nil || net.ParseIP(host) != nil {
		return d.DialContext
	}
	v := &rWW6Rmakjzl{af0cE6r5P: host, tkRa6u: port, rqmdTVGqu: d}
	go func() { //0x1e6e080
		defer HxKViEp2a.SlG_Y_mg("init reality dest")
		for !v.dtw3TCUZjkQh() {
			select {
			case <-ctx.Done():
				return
			case <-time.NewTimer(5 * time.Second).C:
			}
		}
		CAodbhZf.SwUrlM3M(ctx, func() { v.dtw3TCUZjkQh() }, time.Minute)
	}()
	return v.DialContext
}
func (v *rWW6Rmakjzl) dtw3TCUZjkQh() bool { //0x1d6aa00
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ips, err := net.DefaultResolver.LookupIP(ctx, "ip", v.af0cE6r5P)
	if err != nil || len(ips) == 0 {
		muc279Mpa6S4.V8GN5Stk(-4, "resolve reality dest %v failed: %v", v.af0cE6r5P, err)
		return false
	}
	v.sDhm_cyQ1.Store(&ips)
	return true
}

// eF7oFHXp9Xe 0x1d6b660：首地址决定优先地址族，稳定分组；空输入原样panic。
func eF7oFHXp9Xe(ips []net.IP) (primary, secondary []net.IP) {
	firstV4 := ips[0].To4() != nil
	for _, ip := range ips {
		if (ip.To4() != nil) == firstV4 {
			primary = append(primary, ip)
		} else {
			secondary = append(secondary, ip)
		}
	}
	return
}
func (v *rWW6Rmakjzl) cUaTHGDAfd(ctx context.Context, network string, ips []net.IP) (net.Conn, error) { //0x1d6b400
	var last error
	for _, ip := range ips {
		c, err := v.rqmdTVGqu.DialContext(ctx, network, net.JoinHostPort(ip.String(), v.tkRa6u))
		if err == nil {
			return c, nil
		}
		last = err
		if ctx.Err() != nil {
			break
		}
	}
	return nil, last
}
func (v *rWW6Rmakjzl) DialContext(ctx context.Context, network, address string) (net.Conn, error) { //0x1d6ac40
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	ips := v.sDhm_cyQ1.Load()
	if ips == nil {
		return v.rqmdTVGqu.DialContext(ctx, network, address)
	}
	primary, secondary := eF7oFHXp9Xe(*ips)
	if len(secondary) == 0 {
		return v.cUaTHGDAfd(ctx, network, primary)
	}
	results := make(chan fsqq14xvr)
	done := make(chan struct{})
	defer close(done)
	dial := func(group []net.IP) { //func1@1e6e3c0 / deferred func1.1@1e6e4e0
		var r fsqq14xvr
		defer func() {
			if p := recover(); p != nil {
				muc279Mpa6S4.V8GN5Stk(8, "dial reality dest panic: %v", p)
				r = fsqq14xvr{ri0FDK4z: fmt.Errorf("dial reality dest panic: %v", p)}
			}
			select {
			case results <- r:
			case <-done:
				if r.u4oajSYOgl != nil {
					r.u4oajSYOgl.Close()
				}
			}
		}()
		r.u4oajSYOgl, r.ri0FDK4z = v.cUaTHGDAfd(ctx, network, group)
	}
	go dial(primary)
	timer := time.NewTimer(300 * time.Millisecond)
	defer timer.Stop()
	pending := 2
	var firstErr error
	select {
	case r := <-results:
		if r.ri0FDK4z == nil {
			return r.u4oajSYOgl, nil
		}
		firstErr = r.ri0FDK4z
		pending = 1
	case <-timer.C:
	}
	go dial(secondary)
	for ; pending > 0; pending-- {
		r := <-results
		if r.ri0FDK4z == nil {
			return r.u4oajSYOgl, nil
		}
		if firstErr == nil {
			firstErr = r.ri0FDK4z
		}
	}
	return nil, firstErr
}
