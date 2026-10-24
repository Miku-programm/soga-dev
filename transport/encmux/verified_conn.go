package encmux

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	D "recovered.local/soga/buf/pool"
	A "recovered.local/soga/transport/utlsconfig"
	H "recovered.local/soga/protocol/muxcmd"
	X "recovered.local/soga/util/errutil"
)

// aa_M9apoPgl 对应0x11df7e0。
func (c *Ql1qrX) aa_M9apoPgl() (int, error) {
	var p [2]byte
	if _, e := io.ReadFull(c.DrpXao, p[:]); e != nil {
		return 0, e
	}
	return int(binary.BigEndian.Uint16(p[:])), nil
}

// yZuSuMvna 对应0x11df8e0；超出调用方缓冲区的整帧剩余数据被实际丢弃。
func (c *Ql1qrX) yZuSuMvna(p []byte) (int, error) {
	size, e := c.aa_M9apoPgl()
	if e != nil {
		return 0, e
	}
	n, e := io.ReadFull(c.DrpXao, p[:min(size, len(p))])
	if e == nil && n < size {
		e = A.Q6wDan(c.DrpXao, size-n, 0)
	}
	return n, e
}

// icZPfTN 对应0x11dfa60。只重试(0,nil)，16字节令牌的短读或附带错误都失败。
func (c *Ql1qrX) icZPfTN() (bool, error) {
	p := make([]byte, 16)
	for {
		n, e := c.DrpXao.Read(p)
		if e != nil {
			return false, e
		}
		if n > 0 {
			if n != 16 {
				return false, nil
			}
			return [16]byte(p) == c.fR9Xy3.x5Lor8hDyn, nil
		}
	}
}

// oEVIXGXloU 对应0x11dfbc0。
func (c *Ql1qrX) oEVIXGXloU() error {
	var p [1]byte
	if _, e := io.ReadFull(c.DrpXao, p[:]); e != nil {
		return e
	}
	switch p[0] {
	case 1, 2:
		c.mPP4038Qh = H.JTfWJdJvwJ30(p[0])
		return nil
	case 3:
		return c.zMYWqz8I40()
	default:
		return fmt.Errorf("tunnel wrong cmd: %v", p[0])
	}
}

// zMYWqz8I40 对应0x11dfdc0。原proxy标志在后续读取失败前已经设置。
func (c *Ql1qrX) zMYWqz8I40() error {
	if !c.fR9Xy3.hrXHbLocdr {
		return errors.New("not proxy mode")
	}
	c.vfKdlLJ = true
	var p [1]byte
	if _, e := io.ReadFull(c.DrpXao, p[:]); e != nil {
		return e
	}
	switch p[0] {
	case 1, 2, 4:
		c.mPP4038Qh = H.JTfWJdJvwJ30(p[0])
	default:
		return fmt.Errorf("tunnel wrong cmd: %v", p[0])
	}
	address, e := A.UXgN4RrAvUx(c.DrpXao)
	if e != nil {
		return e
	}
	c.z5K6vhFUp = address
	return nil
}

// b1DOoDi4qQI 对应0x11e0120；只有令牌及命令全部成功才记为认证完成。
func (c *Ql1qrX) b1DOoDi4qQI() error {
	if c.cv7owVg {
		return nil
	}
	ok, e := c.icZPfTN()
	if e != nil {
		return e
	}
	if !ok {
		return errors.New("tunnel verify failed")
	}
	if e = c.oEVIXGXloU(); e != nil {
		return e
	}
	c.cv7owVg = true
	return nil
}
func (c *Ql1qrX) Read(p []byte) (int, error) { //0x11e0240
	if e := c.b1DOoDi4qQI(); e != nil {
		return 0, e
	}
	if len(p) == 0 {
		return 0, nil
	}
	switch c.mPP4038Qh {
	case 1, 4:
		return c.DrpXao.Read(p)
	case 2:
		return c.yZuSuMvna(p)
	default:
		return 0, fmt.Errorf("unknown proxy type: %v", c.mPP4038Qh)
	}
}

// tvQhWOesTs 对应0x11e0400。原长度写入uint16不拒绝大帧，短写无错误仍回报全部长度。
func (c *Ql1qrX) tvQhWOesTs(p []byte) (int, error) {
	b := D.S4OQXWIZ(len(p) + 2)
	defer b.Free()
	b.SkipEnd(2)
	binary.BigEndian.PutUint16(b.Bytes(), uint16(len(p)))
	_, _ = b.Write(p)
	if _, e := c.DrpXao.Write(b.Bytes()); e != nil {
		return 0, e
	}
	return len(p), nil
}
func (c *Ql1qrX) Write(p []byte) (int, error) { //0x11e0640
	if e := c.b1DOoDi4qQI(); e != nil {
		return 0, e
	}
	switch c.mPP4038Qh {
	case 1, 4:
		return c.DrpXao.Write(p)
	case 2:
		return c.tvQhWOesTs(p)
	default:
		return 0, fmt.Errorf("unknown proxy type: %v", c.mPP4038Qh)
	}
}
func (c *Ql1qrX) ReadFrom(p []byte) (int, net.Addr, error) { //0x11e0800
	if e := c.b1DOoDi4qQI(); e != nil {
		return 0, nil, e
	}
	if len(p) == 0 {
		return 0, c.DrpXao.RemoteAddr(), nil
	}
	n, e := c.yZuSuMvna(p)
	if e != nil {
		return 0, nil, e
	}
	return n, c.DrpXao.RemoteAddr(), nil
}
func (c *Ql1qrX) WriteTo(p []byte, _ net.Addr) (int, error) {
	if e := c.b1DOoDi4qQI(); e != nil {
		return 0, e
	}
	return c.tvQhWOesTs(p)
} //0x11e0900
func (c *Ql1qrX) GetProxyType() (H.JTfWJdJvwJ30, error) {
	if e := c.b1DOoDi4qQI(); e != nil {
		return 0, e
	}
	return c.mPP4038Qh, nil
} //0x11e09a0
func (c *Ql1qrX) IsProxy() (bool, error) {
	if e := c.b1DOoDi4qQI(); e != nil {
		return false, e
	}
	return c.vfKdlLJ, nil
}                                           //0x11e0a00
func (c *Ql1qrX) GetProxyAddr() *A.FGMp1YT7 { return c.z5K6vhFUp }     //0x11e0a60
func (c *Ql1qrX) GetRawConn() net.Conn      { return c.DrpXao }        //0x11e0a80
func (c *Ql1qrX) Free()                     { X.Li3c4CJ1_D(c.DrpXao) } //0x11e0aa0
// Accept 对应0x11df700；只包装连接，认证由第一次Read/Write/GetProxyType驱动。
func (l *Ppx7OU) Accept() (net.Conn, error) {
	conn, e := l.O6B2qZ0p0.Accept()
	if e != nil {
		return nil, e
	}
	return &Ql1qrX{DrpXao: conn, fR9Xy3: l.tS5gsM72s}, nil
}
