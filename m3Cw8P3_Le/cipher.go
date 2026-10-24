package m3Cw8P3_Le

import (
	"errors"
	"net"
	stream "recovered.local/soga/VwwneWVFFv"
	aead "recovered.local/soga/u3UOJmirww7"
	"strings"
)

type HjDHFt_ interface {
	StreamConn(net.Conn) net.Conn
	PacketConn(net.PacketConn) net.PacketConn
}
type Dc5Zt0 struct {
	aead.HugVbq
	Ez2IY3fv []byte
}
type E3HaL6A8 struct {
	stream.ZDHlm7vHo
	Ez2IY3fv []byte
}
type cKUfPpsN2WU struct{}

// init 0x194d8c0. All nonliteral algorithm names decoded from original thunks;
// funcval 0x377d1c0/1d8/1d0 resolve to CTR/CFB/RC4-MD5 constructors.
var errorAt5098510 = errors.New("cipher not supported")
var aeadAt5096498 = map[string]struct {
	Iqtgx53      int
	NHaagjNZeaDf func([]byte) (aead.HugVbq, error)
}{
	"AEAD_AES_128_GCM": {16, aead.Rk8ppUT2aLh}, "AEAD_AES_192_GCM": {24, aead.Rk8ppUT2aLh}, "AEAD_AES_256_GCM": {32, aead.Rk8ppUT2aLh}, "AEAD_CHACHA20_POLY1305": {32, aead.IEFKpuUGsJ}, "AEAD_XCHACHA20_POLY1305": {32, aead.BeqNJFSF},
}
var streamAt50964A0 = map[string]struct {
	Iqtgx53      int
	NHaagjNZeaDf func([]byte) (stream.ZDHlm7vHo, error)
}{
	"RC4-MD5": {16, stream.F3Yl5rl}, "AES-128-CTR": {16, stream.E6vQRRYI}, "AES-192-CTR": {24, stream.E6vQRRYI}, "AES-256-CTR": {32, stream.E6vQRRYI}, "AES-128-CFB": {16, stream.HQZtzHEUA3}, "AES-192-CFB": {24, stream.HQZtzHEUA3}, "AES-256-CFB": {32, stream.HQZtzHEUA3}, "CHACHA20-IETF": {32, stream.E9kDmJ55P}, "XCHACHA20": {32, stream.AfbgTLVHrju},
}

// B93MVPtD: 0x194e000. Names are uppercased without trimming. Explicit empty
// keys derive from password; nonempty keys retain their original backing slice.
func B93MVPtD(name string, key []byte, password string) (HjDHFt_, error) {
	name = strings.ToUpper(name)
	switch name {
	case "NONE", "DUMMY", "PLAIN":
		return &cKUfPpsN2WU{}, nil
	case "CHACHA20-IETF-POLY1305":
		name = "AEAD_CHACHA20_POLY1305"
	case "XCHACHA20-IETF-POLY1305":
		name = "AEAD_XCHACHA20_POLY1305"
	case "AES-128-GCM":
		name = "AEAD_AES_128_GCM"
	case "AES-192-GCM":
		name = "AEAD_AES_192_GCM"
	case "AES-256-GCM":
		name = "AEAD_AES_256_GCM"
	}
	if entry, ok := aeadAt5096498[name]; ok {
		if len(key) == 0 {
			key = ArNrGS_EV(password, entry.Iqtgx53)
		}
		if len(key) != entry.Iqtgx53 {
			return nil, aead.ZPe9opSsSr9k(entry.Iqtgx53)
		}
		c, e := entry.NHaagjNZeaDf(key)
		return &Dc5Zt0{c, key}, e
	}
	if entry, ok := streamAt50964A0[name]; ok {
		if len(key) == 0 {
			key = ArNrGS_EV(password, entry.Iqtgx53)
		}
		if len(key) != entry.Iqtgx53 {
			return nil, stream.XJzap3(entry.Iqtgx53)
		}
		c, e := entry.NHaagjNZeaDf(key)
		return &E3HaL6A8{c, key}, e
	}
	return nil, errorAt5098510
}

// Original wrappers 0x194e680/194e740/194e800/194e8c0 preserve the outer
// cipher object as the dynamically embedded interface value.
func (c *Dc5Zt0) StreamConn(conn net.Conn) net.Conn { return &aead.AZVSCxZmZQ{Conn: conn, HugVbq: c} }
func (c *Dc5Zt0) PacketConn(conn net.PacketConn) net.PacketConn {
	return &aead.Zh37aRKMK{PacketConn: conn, HugVbq: c}
}
func (c *E3HaL6A8) StreamConn(conn net.Conn) net.Conn {
	return &stream.AriwPpia{Conn: conn, ZDHlm7vHo: c}
}
func (c *E3HaL6A8) PacketConn(conn net.PacketConn) net.PacketConn {
	return &stream.HtnJMv4Lsng{PacketConn: conn, ZDHlm7vHo: c}
}

// 0x194e980/194e9a0.
func (*cKUfPpsN2WU) StreamConn(c net.Conn) net.Conn             { return c }
func (*cKUfPpsN2WU) PacketConn(c net.PacketConn) net.PacketConn { return c }
