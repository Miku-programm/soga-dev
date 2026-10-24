package J1P4KNvfI_J

import (
	"fmt"
	"net"
	"recovered.local/soga/muc279Mpa6S4"
)

// String 对应0x187f9e0。用户为空时保留原直接调用 GetId 的 panic 行为。
func (s *ZzkV7_5l) String() string {
	return fmt.Sprintf("session id: %v, user: %v, dest addr: %v, client(%v:%v) -> local(%v:%v) -> remote_local(%v:%v) -> remote(%v:%v)", s.E3L0l8hZ, s.PayAlbdv6.GetId(), s.Iayb4V, s.KeHzUwuXWh0, s.Y72kDV35, s.L4SIW3qbjS, s.X1gBTWHQO99, s.R61M7ZTr, s.N15sqBj36, s.IksfOg, s.IqEsiK)
}

// SetOutboundLocal 对应0x187ef60。未知地址只记录警告，保留已有字段。
func (s *ZzkV7_5l) SetOutboundLocal(address net.Addr) {
	ip, port, ok := sessionAddressParts(address)
	if ok {
		s.R61M7ZTr = ip
		s.N15sqBj36 = port
		return
	}
	muc279Mpa6S4.V8GN5Stk(4, "session set outbound local failed: %v", address)
}

// SetOutboundRemote 对应0x187f0e0。IP切片保留原引用，不复制。
func (s *ZzkV7_5l) SetOutboundRemote(address net.Addr) {
	ip, port, ok := sessionAddressParts(address)
	if ok {
		s.IksfOg = ip
		s.IqEsiK = port
		return
	}
	muc279Mpa6S4.V8GN5Stk(4, "session set outbound remote failed: %v", address)
}
