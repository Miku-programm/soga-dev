package utlsconfig

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
)

// Fqn4mE5 对应0x116eca0。保留原文本，ParseIP产生独立的16字节地址。
func Fqn4mE5(host string, port int) *FGMp1YT7 {
	return &FGMp1YT7{bt6_FOvs5: host, b6HasR_: net.ParseIP(host), nrxt1fRatq1Y: port}
}
func (a *FGMp1YT7) HasIP() bool     { return a.b6HasR_ != nil }                                            //0x116edc0
func (a *FGMp1YT7) GetIP() net.IP   { return a.b6HasR_ }                                                   //0x116ede0
func (a *FGMp1YT7) GetHost() string { return a.bt6_FOvs5 }                                                 //0x116ee00
func (a *FGMp1YT7) GetPort() int    { return a.nrxt1fRatq1Y }                                              //0x116ee20
func (a *FGMp1YT7) Join() string    { return net.JoinHostPort(a.bt6_FOvs5, strconv.Itoa(a.nrxt1fRatq1Y)) } //0x116ee40
func (a *FGMp1YT7) ToTCPAddr() (*net.TCPAddr, error) {
	if a.b6HasR_ == nil {
		return nil, fmt.Errorf("addr is not ip")
	}
	return &net.TCPAddr{IP: a.b6HasR_, Port: a.nrxt1fRatq1Y}, nil
} //0x116ef40
func (a *FGMp1YT7) ToUDPAddr() (*net.UDPAddr, error) {
	if a.b6HasR_ == nil {
		return nil, fmt.Errorf("addr is not ip")
	}
	return &net.UDPAddr{IP: a.b6HasR_, Port: a.nrxt1fRatq1Y}, nil
}                                  //0x116f080
func (a *FGMp1YT7) String() string { return fmt.Sprintf("%v:%v", a.bt6_FOvs5, a.nrxt1fRatq1Y) } //0x116f1c0
func (a *FGMp1YT7) EqualTo(b *FGMp1YT7) bool {
	return a.bt6_FOvs5 == b.bt6_FOvs5 && a.nrxt1fRatq1Y == b.nrxt1fRatq1Y
} //0x116f260
func (a *FGMp1YT7) AddrDataSize() int {
	if a.b6HasR_ == nil {
		return len(a.bt6_FOvs5) + 4
	}
	if a.b6HasR_.To4() != nil {
		return 7
	}
	return 19
} //0x116f2e0
// KayO9u 对应0x11701a0。
func KayO9u(reader io.Reader) (byte, error) {
	var b [1]byte
	if _, e := io.ReadFull(reader, b[:]); e != nil {
		return 0, e
	}
	return b[0], nil
}

// UXgN4RrAvUx 对应0x1170240；IP切片保留包含端口字节的底层容量。
// 域名分支也经过ParseIP，所以文本形式IP会保留文本并解析出IP。
func UXgN4RrAvUx(reader io.Reader) (*FGMp1YT7, error) {
	kind, err := KayO9u(reader)
	if err != nil {
		return nil, err
	}
	var size int
	switch kind {
	case 1:
		size = 6
	case 4:
		size = 18
	case 3:
		n, e := KayO9u(reader)
		if e != nil {
			return nil, e
		}
		size = int(n) + 2
	default:
		return nil, fmt.Errorf("unknown addr type: %v", kind)
	}
	data := make([]byte, size)
	if _, e := io.ReadFull(reader, data); e != nil {
		return nil, e
	}
	port := int(binary.BigEndian.Uint16(data[size-2:]))
	host := data[:size-2]
	if kind == 3 {
		return Fqn4mE5(string(host), port), nil
	}
	ip := net.IP(host)
	return &FGMp1YT7{bt6_FOvs5: ip.String(), b6HasR_: ip, nrxt1fRatq1Y: port}, nil
}
