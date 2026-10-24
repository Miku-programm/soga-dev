// Package JX5NGDvCQwty restores the original native SS protocol factories.
package ssudp

import (
	"net"
	"recovered.local/soga/protocol/ss2022"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/transport/muxerr"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
)

// PacketOptions names the original 64-byte value argument reconstructed from
// CXui61CuSLa's ABI and its caller @0x1d6d5c0. The aggregate and field labels
// did not survive RTTI and are explicit analysis names, not author names.
type PacketOptions struct {
	Conn       net.PacketConn
	Cipher     ssmethod.C6q3Upu
	Users      *user.ZpNFqUVZp
	Candidates *sess.EafxyUZD
	Invalid    *limit.M1Qn9BA
	Manager    *ss2022.N07NaA
}

func CXui61CuSLa(o PacketOptions) sess.KsKg3yls { //0x1c841c0
	if o.Cipher == "2022-blake3-aes-128-gcm" || o.Cipher == "2022-blake3-aes-256-gcm" {
		return ss2022.NewPacketConnForFactory(o.Conn, o.Manager)
	}
	return muxerr.OlQGbrpyFga(o.Conn, o.Cipher, o.Candidates, o.Users, o.Invalid)
}
