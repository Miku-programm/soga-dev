package g0MlUOhvO

import (
	"fmt"
	"net"
	"recovered.local/soga/DsyR31Nu"
	"recovered.local/soga/JcmlVIK0T0i"
)

// WrapListener 对应 0x11db9c0。顺序来自原机器码：QUIC 返回原监听器，
// 其它模式依次包裹 TLS、WebSocket、加密记录、MUX。这里没有新增协议实现。
func (v *VTl94pzz2IkB) WrapListener(l net.Listener) (net.Listener, error) {
	if v.hZORuZstvS07 {
		return l, nil
	}
	if v.rj1DhSBHZ7O {
		if v._rKF9rrYaf8a == nil {
			return nil, fmt.Errorf("tunnel config certificate is nil")
		}
		l = DsyR31Nu.MYA8Jh(l, v._rKF9rrYaf8a, false)
	}
	if v.ufiMqR6x {
		l = GHrg5b6zAf(l, v.lyXfKPaz4m23)
	}
	l = &T4k2R7cJ{O6B2qZ0p0: l, tS5gsM72s: v}
	if v.e2TiiBKaBt {
		l = Bn7A4E9vXH(l, v)
	}
	return l, nil
}

// HNbZee047 对应 0x11df5c0：完整保留原外层验证与 PROXY 解析包裹顺序。
func HNbZee047(l net.Listener, v *VTl94pzz2IkB) (net.Listener, error) {
	wrapped, e := v.WrapListener(l)
	if e != nil {
		return nil, e
	}
	wrapped = &Ppx7OU{O6B2qZ0p0: wrapped, tS5gsM72s: v}
	if v.jrF7WDTK {
		wrapped = &JcmlVIK0T0i.KprDQw4aZEOg{I9Hb4SAu: wrapped}
	}
	return wrapped, nil
}
