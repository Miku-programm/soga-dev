package netiface

import (
	"errors"
	"net"
	"recovered.local/soga/net/dialctx"
	"time"
)

// Ah4poD retains the original 72-byte RTTI layout (_GT9LX.Ah4poD).
type Ah4poD struct {
	BdMyLmuSMQA  int
	Arz5udf      string
	JNv6km       []*net.IPNet
	LcX0_pZ9KNNi net.HardwareAddr
}

// Error strings decoded from init 0x1920920. Exported first-letter mappings
// expose the original shared error values to the restored socket package.
var InterfaceNotFoundAt5098650 = errors.New("interface not found")
var AddrNotFoundAt5098660 = errors.New("addr not found")

// Static pointer 0x5029cd8 -> object 0x504d028 has duration 0x4a817c800 = 20s.
var interfaceCacheAt5029CD8 = dialctx.NewCache(20 * time.Second)

// DINaz4: 0x1920a80.
func DINaz4(name string) (*Ah4poD, error) {
	value, err, _ := interfaceCacheAt5029CD8.Do(dinaz4func1)
	if err != nil {
		return nil, err
	}
	if iface, ok := value.(map[string]*Ah4poD)[name]; ok {
		return iface, nil
	}
	return nil, InterfaceNotFoundAt5098650
}

// dinaz4func1 restores DINaz4.func1 at 0x19220a0. Interfaces whose Addrs call
// fails are skipped; a non-*net.IPNet address triggers the original assertion.
func dinaz4func1() (any, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	result := make(map[string]*Ah4poD)
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		ips := make([]*net.IPNet, 0, len(addrs))
		for _, addr := range addrs {
			ip := addr.(*net.IPNet)
			if v4 := ip.IP.To4(); v4 != nil {
				ip.IP = v4
			}
			ips = append(ips, ip)
		}
		result[iface.Name] = &Ah4poD{iface.Index, iface.Name, ips, iface.HardwareAddr}
	}
	return result, nil
}

// PickIPv4Addr: 0x1920b40; predicate: 0x1922500.
func (a *Ah4poD) PickIPv4Addr(destination net.IP) (*net.IPNet, error) {
	return a.zpdC3VK3HIL(destination, func(ip *net.IPNet) bool { return ip.IP.To4() != nil })
}

// PickIPv6Addr: 0x1920ba0; predicate: 0x19224a0.
func (a *Ah4poD) PickIPv6Addr(destination net.IP) (*net.IPNet, error) {
	return a.zpdC3VK3HIL(destination, func(ip *net.IPNet) bool { return ip.IP.To4() == nil })
}

// zpdC3VK3HIL: 0x1920c00. Prefer a containing subnet, otherwise the first
// non-link-local address. Nil destination stops immediately at that fallback.
func (a *Ah4poD) zpdC3VK3HIL(destination net.IP, accept func(*net.IPNet) bool) (*net.IPNet, error) {
	var selected *net.IPNet
	for _, ip := range a.JNv6km {
		if !accept(ip) {
			continue
		}
		if selected == nil && !ip.IP.IsLinkLocalUnicast() {
			selected = ip
			if destination == nil {
				break
			}
		}
		if destination != nil && ip.Contains(destination) {
			return ip, nil
		}
	}
	if selected != nil {
		return selected, nil
	}
	return nil, AddrNotFoundAt5098660
}
