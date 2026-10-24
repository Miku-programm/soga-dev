// 依据原样本RTTI及IDA恢复，不是原作者源码。
package encmux

import (
	"crypto/tls"
	"net"
	"recovered.local/soga/crypto/kdf"
	"recovered.local/soga/transport/utlsconfig"
	"recovered.local/soga/protocol/muxcmd"
)

// 保留原嵌入字段名，标准网络类型用别名保持真实方法集合。
type DrpXao = net.Conn
type O6B2qZ0p0 = net.Listener

// VTl94pzz2IkB 原144字节隧道状态，显示前缀bhXMvjeYm_m。
type VTl94pzz2IkB struct {
	ufiMqR6x     bool
	hZORuZstvS07 bool
	rj1DhSBHZ7O  bool
	e2TiiBKaBt   bool
	pyauaz       int
	p1Wtyln2GJ6  string
	psdXJ2Mi     bool
	pJLNnm       string
	_rKF9rrYaf8a *tls.Certificate
	lyXfKPaz4m23 string
	jrF7WDTK     bool
	hrXHbLocdr   bool
	dlAoUM       *utlsconfig.FGMp1YT7
	zkn35IN      *kdf.PAEsiXjhBp_C
	oWvGY4I9cA   []byte
	x5Lor8hDyn   [16]byte
}

// Ql1qrX 原56字节服务端认证连接，认证完成标志默认false。
type Ql1qrX struct {
	DrpXao
	fR9Xy3    *VTl94pzz2IkB
	cv7owVg   bool
	quUa_I    int
	mPP4038Qh muxcmd.JTfWJdJvwJ30
	vfKdlLJ   bool
	z5K6vhFUp *utlsconfig.FGMp1YT7
}

// Ppx7OU 原24字节监听器。
type Ppx7OU struct {
	O6B2qZ0p0
	tS5gsM72s *VTl94pzz2IkB
}
