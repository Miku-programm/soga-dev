package G6bHD5JA

import (
	"fmt"
	"net"
	"recovered.local/soga/BaEWsPy"
	"strconv"
)

// AHU2MRmjzQu：原 RTTI 56 字节，字段偏移 0/16/40/48。
type AHU2MRmjzQu struct {
	dlcy_1     string
	oKcMaQDwbH []net.IP
	pDnAz_     int
	r_VdNaPA   bool
}

// RNFJeaRY @ 0x159de00。原函数仅接受 TCPAddr/UDPAddr，保留底层 IP 切片。
func RNFJeaRY(addr net.Addr) (*AHU2MRmjzQu, error) {
	var ip net.IP
	var port int
	switch a := addr.(type) {
	case *net.TCPAddr:
		ip, port = a.IP, a.Port
	case *net.UDPAddr:
		ip, port = a.IP, a.Port
	default:
		return nil, fmt.Errorf("unknown addr: %v", addr)
	}
	return &AHU2MRmjzQu{dlcy_1: ip.String(), oKcMaQDwbH: []net.IP{ip}, pDnAz_: port, r_VdNaPA: true}, nil
}

// U80FDt @ 0x159e1c0。端口使用 Atoi，原实现没有额外范围检查。
func U80FDt(address string) (*AHU2MRmjzQu, error) {
	host, portText, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(portText)
	if err != nil {
		return nil, err
	}
	if ip := net.ParseIP(host); ip != nil {
		return &AHU2MRmjzQu{dlcy_1: ip.String(), oKcMaQDwbH: []net.IP{ip}, pDnAz_: port, r_VdNaPA: true}, nil
	}
	return &AHU2MRmjzQu{dlcy_1: host, pDnAz_: port}, nil
}

func (a *AHU2MRmjzQu) HasIP() bool { return len(a.oKcMaQDwbH) != 0 } // 0x159e3e0
func (a *AHU2MRmjzQu) IsIP() bool  { return a.r_VdNaPA }             // 0x159e400
func (a *AHU2MRmjzQu) GetFirstIp() net.IP {
	if len(a.oKcMaQDwbH) == 0 {
		return nil
	}
	return a.oKcMaQDwbH[0]
}                                       // 0x159e420
func (a *AHU2MRmjzQu) GetIPs() []net.IP { return a.oKcMaQDwbH } // 0x159e440
func (a *AHU2MRmjzQu) GetHost() string  { return a.dlcy_1 }     // 0x159e460
// SetHost @ 0x159e480 不清除已解析 IP；只设置主机和 literal 标记。
func (a *AHU2MRmjzQu) SetHost(host string) { a.dlcy_1 = host; a.r_VdNaPA = false }
func (a *AHU2MRmjzQu) GetPort() int        { return a.pDnAz_ }                                           // 0x159e4e0
func (a *AHU2MRmjzQu) Join() string        { return net.JoinHostPort(a.dlcy_1, strconv.Itoa(a.pDnAz_)) } // 0x159e500
func (a *AHU2MRmjzQu) HostPort() string    { return net.JoinHostPort(a.dlcy_1, strconv.Itoa(a.pDnAz_)) } // 0x159ed20

// ToTCPAddrs @ 0x159e600。
func (a *AHU2MRmjzQu) ToTCPAddrs() ([]*net.TCPAddr, error) {
	if len(a.oKcMaQDwbH) == 0 {
		return nil, fmt.Errorf("addr is not ip")
	}
	out := make([]*net.TCPAddr, 0, len(a.oKcMaQDwbH))
	for _, ip := range a.oKcMaQDwbH {
		out = append(out, &net.TCPAddr{IP: ip, Port: a.pDnAz_})
	}
	return out, nil
}
func (a *AHU2MRmjzQu) ToUDPAddr() (*net.UDPAddr, error) { // 0x159e880
	if len(a.oKcMaQDwbH) == 0 {
		return nil, fmt.Errorf("addr is not ip")
	}
	return &net.UDPAddr{IP: a.oKcMaQDwbH[0], Port: a.pDnAz_}, nil
}
func (a *AHU2MRmjzQu) ResolveHost(resolver *BaEWsPy.D23y2888a_) error { // 0x159e9e0
	if len(a.oKcMaQDwbH) != 0 || a.r_VdNaPA {
		return nil
	}
	ips, err := resolver.Resolve(a.dlcy_1)
	if err != nil {
		return err
	}
	a.oKcMaQDwbH = ips
	return nil
}

// String @ 0x159eaa0 和 HostPort 不同；原展示字符串不为 IPv6 加括号。
func (a *AHU2MRmjzQu) String() string {
	if a.r_VdNaPA {
		return fmt.Sprintf("%v:%v", a.GetFirstIp(), a.pDnAz_)
	}
	if len(a.oKcMaQDwbH) != 0 {
		return fmt.Sprintf("%v(%v):%v", a.dlcy_1, a.oKcMaQDwbH, a.pDnAz_)
	}
	return fmt.Sprintf("%v:%v", a.dlcy_1, a.pDnAz_)
}
func (a *AHU2MRmjzQu) EqualTo(other *AHU2MRmjzQu) bool {
	return a.dlcy_1 == other.dlcy_1 && a.pDnAz_ == other.pDnAz_
} // 0x159ee20
func (a *AHU2MRmjzQu) Size() int { // 0x159eea0
	if len(a.oKcMaQDwbH) == 0 {
		return len(a.dlcy_1) + 4
	}
	if a.oKcMaQDwbH[0].To4() != nil {
		return 7
	}
	return 19
}
