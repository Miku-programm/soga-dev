package ssudp

import (
	"net"
	"recovered.local/soga/protocol/ss2022"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/transport/muxerr"
	"recovered.local/soga/transport/httpnative"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
)

// StreamOptions names the original 120-byte value argument, reconstructed from
// FciVp4yFHL's stack ABI and inbound.Lp07FzBLx9V's TCP-session caller. These
// aggregate field labels are explicit analysis names; no aggregate RTTI survives.
type StreamOptions struct {
	Conn       net.Conn
	Cipher     ssmethod.C6q3Upu
	Obfs       string
	Mode       ssmethod.VE7a8rvh
	Path       string
	Host       string
	Candidates *sess.EafxyUZD
	Invalid    *limit.M1Qn9BA
	Users      *user.ZpNFqUVZp
	Manager    *ss2022.N07NaA
}

func FciVp4yFHL(o StreamOptions) (sess.Ki8urSHTx44, error) { //0x1c83e20
	conn := o.Conn
	switch o.Obfs {
	case "tls":
		conn = httpnative.NewTLSForFactory(conn, o.Host, o.Mode)
	case "http":
		path := o.Path
		if len(path) == 0 || path[0] != '/' {
			path = "/" + path
		}
		conn = httpnative.NewHTTPForFactory(conn, path, o.Host, o.Mode)
	}
	if o.Cipher == "2022-blake3-aes-128-gcm" || o.Cipher == "2022-blake3-aes-256-gcm" {
		return ss2022.YZbzDDrJX(conn, o.Manager)
	}
	// The original always sets this concrete interface itab, including when
	// VLnoIxcGjx returns a nil pointer with an error. Preserve the typed nil.
	c, err := muxerr.VLnoIxcGjx(conn, o.Cipher, o.Candidates, o.Invalid, o.Users)
	return c, err
}
