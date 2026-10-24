package JX5NGDvCQwty

import (
	"net"
	"recovered.local/soga/GTiWvwS"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/NtxdAE"
	"recovered.local/soga/VlSfV5"
	"recovered.local/soga/digGeaHQxeZ"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/w2YaTd9_"
)

// StreamOptions names the original 120-byte value argument, reconstructed from
// FciVp4yFHL's stack ABI and MFFa9yT.Lp07FzBLx9V's TCP-session caller. These
// aggregate field labels are explicit analysis names; no aggregate RTTI survives.
type StreamOptions struct {
	Conn       net.Conn
	Cipher     w2YaTd9_.C6q3Upu
	Obfs       string
	Mode       w2YaTd9_.VE7a8rvh
	Path       string
	Host       string
	Candidates *J1P4KNvfI_J.EafxyUZD
	Invalid    *NtxdAE.M1Qn9BA
	Users      *nzATvajt.ZpNFqUVZp
	Manager    *GTiWvwS.N07NaA
}

func FciVp4yFHL(o StreamOptions) (J1P4KNvfI_J.Ki8urSHTx44, error) { //0x1c83e20
	conn := o.Conn
	switch o.Obfs {
	case "tls":
		conn = digGeaHQxeZ.NewTLSForFactory(conn, o.Host, o.Mode)
	case "http":
		path := o.Path
		if len(path) == 0 || path[0] != '/' {
			path = "/" + path
		}
		conn = digGeaHQxeZ.NewHTTPForFactory(conn, path, o.Host, o.Mode)
	}
	if o.Cipher == "2022-blake3-aes-128-gcm" || o.Cipher == "2022-blake3-aes-256-gcm" {
		return GTiWvwS.YZbzDDrJX(conn, o.Manager)
	}
	// The original always sets this concrete interface itab, including when
	// VLnoIxcGjx returns a nil pointer with an error. Preserve the typed nil.
	c, err := VlSfV5.VLnoIxcGjx(conn, o.Cipher, o.Candidates, o.Invalid, o.Users)
	return c, err
}
