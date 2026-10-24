// Package JX5NGDvCQwty restores the original native SS protocol factories.
package JX5NGDvCQwty

import (
	"net"
	"recovered.local/soga/GTiWvwS"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/NtxdAE"
	"recovered.local/soga/VlSfV5"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/w2YaTd9_"
)

// PacketOptions names the original 64-byte value argument reconstructed from
// CXui61CuSLa's ABI and its caller @0x1d6d5c0. The aggregate and field labels
// did not survive RTTI and are explicit analysis names, not author names.
type PacketOptions struct {
	Conn       net.PacketConn
	Cipher     w2YaTd9_.C6q3Upu
	Users      *nzATvajt.ZpNFqUVZp
	Candidates *J1P4KNvfI_J.EafxyUZD
	Invalid    *NtxdAE.M1Qn9BA
	Manager    *GTiWvwS.N07NaA
}

func CXui61CuSLa(o PacketOptions) J1P4KNvfI_J.KsKg3yls { //0x1c841c0
	if o.Cipher == "2022-blake3-aes-128-gcm" || o.Cipher == "2022-blake3-aes-256-gcm" {
		return GTiWvwS.NewPacketConnForFactory(o.Conn, o.Manager)
	}
	return VlSfV5.OlQGbrpyFga(o.Conn, o.Cipher, o.Candidates, o.Users, o.Invalid)
}
