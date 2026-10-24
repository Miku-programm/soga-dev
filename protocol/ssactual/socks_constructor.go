package ssactual

import (
	"crypto/tls"
	"net"
	"strconv"
)

// FFrN8gdu retains the original 24-byte option prefix and tags.
type FFrN8gdu struct {
	PiTAGr6U5KDG string `proxy:"interface-name,omitempty" group:"interface-name,omitempty"`
	VsBRTR       int    `proxy:"routing-mark,omitempty" group:"routing-mark,omitempty"`
}

// SocksOptions names the 104-byte by-value argument layout read at 0x1b66240.
// Its original source type/field names were not retained in RTTI. These descriptive
// names preserve the observed offsets; no unobserved validation/defaults are added.
type SocksOptions struct {
	FFrN8gdu
	Name           string
	Server         string
	Port           int
	Username       string
	Password       string
	TLS            bool
	UDP            bool
	SkipCertVerify bool
}

// AwDE24aRWC retains its original 56-byte RTTI layout.
type AwDE24aRWC struct {
	*B8EHsVY7a
	gmXALhCpyKF string
	eOL60dT1    string
	m0JWRLb2CH  bool
	p7gaqRDae   bool
	idTOPJOmPpB *tls.Config
}

// Auqn90u2NBQ: 0x1b66240. This constructor returns only a pointer, not an error.
func Auqn90u2NBQ(o SocksOptions) *AwDE24aRWC {
	var config *tls.Config
	if o.TLS {
		config = &tls.Config{InsecureSkipVerify: o.SkipCertVerify, ServerName: o.Server}
	}
	return &AwDE24aRWC{
		B8EHsVY7a:   &B8EHsVY7a{m50I26ZDL: o.Name, x3HUaU8Ln1S1: net.JoinHostPort(o.Server, strconv.Itoa(o.Port)), xDan8hA5LV: o.PiTAGr6U5KDG, iOKxDl: 5, ntseIWz: o.UDP, bEbqkM8: o.VsBRTR},
		gmXALhCpyKF: o.Username, eOL60dT1: o.Password, m0JWRLb2CH: o.TLS, p7gaqRDae: o.SkipCertVerify, idTOPJOmPpB: config,
	}
}
