package EP3M1uC

import (
	"encoding/binary"
	"net"
	"recovered.local/soga/uLofn6j"
)

// Select restores 0x1b794e0 (cryptographically random selection).
func (*h1kE2OhgN) Select(routes J_enG0wE, _ net.IP, _ int64) *GyyO51 {
	if len(routes) == 1 {
		return routes[0]
	}
	return routes[uLofn6j.Co6D_zfls(len(routes))]
}

// Select restores 0x1b795a0; the first index is 1, not 0.
func (s *bLJQqABD) Select(routes J_enG0wE, _ net.IP, _ int64) *GyyO51 {
	return routes[s.v1TcNPAI0.Add(1)%uint64(len(routes))]
}

// Select restores 0x1b795e0. A non-nil empty IP hashes the empty byte sequence.
func (*twxaE4km_U) Select(routes J_enG0wE, ip net.IP, _ int64) *GyyO51 {
	if len(routes) == 1 {
		return routes[0]
	}
	if ip == nil {
		return routes[uLofn6j.Co6D_zfls(len(routes))]
	}
	hash := uint32(2166136261)
	for _, b := range ip {
		hash = (hash ^ uint32(b)) * 16777619
	}
	return routes[hash%uint32(len(routes))]
}

// Select restores 0x1b79700, hashing the positive user ID in network byte order.
func (*gMxUOGd) Select(routes J_enG0wE, _ net.IP, userID int64) *GyyO51 {
	if len(routes) == 1 {
		return routes[0]
	}
	if userID <= 0 {
		return routes[uLofn6j.Co6D_zfls(len(routes))]
	}
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], uint64(userID))
	hash := uint32(2166136261)
	for _, b := range data {
		hash = (hash ^ uint32(b)) * 16777619
	}
	return routes[hash%uint32(len(routes))]
}
