package proxyproto

import "net"

type JvCeJK0s uint8
type ZYgyOe uint8
type XTAkygU int
type K_M6K7 func(*URdA97qPFc) error
type AXaIUgeevW2 func(net.Addr) (XTAkygU, error)

func (p JvCeJK0s) IsIPv4() bool     { return p&0xf0 == 0x10 }             // 0x1175700
func (p JvCeJK0s) IsIPv6() bool     { return p&0xf0 == 0x20 }             // 0x1175720
func (p JvCeJK0s) IsUnix() bool     { return p&0xf0 == 0x30 }             // 0x1175740
func (p JvCeJK0s) IsStream() bool   { return p&0x0f == 1 }                // 0x1175760
func (p JvCeJK0s) IsDatagram() bool { return p&0x0f == 2 }                // 0x1175780
func (p JvCeJK0s) IsUnspec() bool   { return p&0xf0 == 0 || p&0x0f == 0 } // 0x11757a0
func (p JvCeJK0s) uGOaaar() byte { // 0x11757c0
	if (p.IsIPv4() || p.IsIPv6() || p.IsUnix()) && (p.IsStream() || p.IsDatagram()) {
		return byte(p)
	}
	return 0
}
func (p ZYgyOe) IsLocal() bool  { return p == 0x20 }                    // 0x118b400
func (p ZYgyOe) IsProxy() bool  { return p == 0x21 }                    // 0x118b420
func (p ZYgyOe) IsUnspec() bool { return !p.IsLocal() && !p.IsProxy() } // 0x118b440

// URdA97qPFc：原 PROXY 协议头 64 字节布局。
type URdA97qPFc struct {
	TmZyJAm      uint8
	LgI8Cr1feIYb ZYgyOe
	UCYj4Sj65ui  JvCeJK0s
	Dafxa_BOdZdw net.Addr
	KhtRcTD89    net.Addr
	ttF0mSAdaYG  []byte
}
