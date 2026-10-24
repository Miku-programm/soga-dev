// 本包只恢复样本中的地址值类型，不引用完整 Xray 内核。
package net

import (
	"bytes"
	"net"
	Log "recovered.local/soga/util/logx"
	"strings"
)

type AddressFamily byte

const (
	AddressFamilyIPv4 AddressFamily = iota
	AddressFamilyIPv6
	AddressFamilyDomain
)

func (f AddressFamily) IsIPv4() bool   { return f == 0 }
func (f AddressFamily) IsIPv6() bool   { return f == 1 }
func (f AddressFamily) IsIP() bool     { return f|1 == 1 }
func (f AddressFamily) IsDomain() bool { return f == 2 }

type Address interface {
	IP() net.IP
	Domain() string
	Family() AddressFamily
	String() string
}
type ipv4Address [4]byte
type ipv6Address [16]byte
type domainAddress string

func (a ipv4Address) IP() net.IP              { return net.IP(a[:]) }
func (a ipv6Address) IP() net.IP              { return net.IP(a[:]) }
func (a domainAddress) IP() net.IP            { panic("Calling IP() on a DomainAddress.") }
func (a ipv4Address) Domain() string          { panic("Calling Domain() on an IPv4Address.") }
func (a ipv6Address) Domain() string          { panic("Calling Domain() on an IPv6Address.") }
func (a domainAddress) Domain() string        { return string(a) }
func (a ipv4Address) Family() AddressFamily   { return 0 }
func (a ipv6Address) Family() AddressFamily   { return 1 }
func (a domainAddress) Family() AddressFamily { return 2 }
func (a ipv4Address) String() string          { return net.IP(a[:]).String() }
func (a ipv6Address) String() string          { return "[" + net.IP(a[:]).String() + "]" }
func (a domainAddress) String() string        { return string(a) }

// 0x124fdc0：先去成对方括号，再按首尾字符决定是否 TrimSpace。
func ParseAddress(s string) Address {
	if len(s) != 0 && s[0] == '[' && s[len(s)-1] == ']' {
		s = s[1 : len(s)-1]
	}
	alnum := func(c byte) bool { return c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' }
	if len(s) > 0 && (!alnum(s[0]) || !alnum(s[len(s)-1])) {
		s = strings.TrimSpace(s)
	}
	if ip := net.ParseIP(s); ip != nil {
		return IPAddress(ip)
	}
	return domainAddress(s)
}

// 0x124ff20：复制成值类型；IPv4-mapped IPv6 会缩为四字节。
func IPAddress(ip []byte) Address {
	switch len(ip) {
	case 4:
		return ipv4Address([4]byte(ip))
	case 16:
		if bytes.Equal(ip[:10], make([]byte, 10)) && ip[10] == 255 && ip[11] == 255 {
			return IPAddress(ip[12:])
		}
		return ipv6Address([16]byte(ip))
	default:
		Log.TRTbrquD(8, "invalid IP format: ", ip)
		return nil
	}
}
