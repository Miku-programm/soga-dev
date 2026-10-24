//go:build linux

package ssactual

import (
	"context"
	"fmt"
	"net"
	plugin "recovered.local/soga/transport/tlssni"
	E "recovered.local/soga/net/outbound"
	decode "recovered.local/soga/transport/quicwrap"
	obfs "recovered.local/soga/transport/httpobfs"
	cipher "recovered.local/soga/crypto/ciphernames"
	O "recovered.local/soga/net/socket"
	"strconv"
)

type cJlUN29eP struct {
	VfjnRgEJIGGI string `obfs:"mode,omitempty"`
	OJdz_5FmLCKj string `obfs:"host,omitempty"`
}
type dXR4UapBiBa struct {
	GTwWMvxPP8   string            `obfs:"mode"`
	L3gscTlpiGj  string            `obfs:"host,omitempty"`
	RhRqPp       string            `obfs:"path,omitempty"`
	RZZqKB       bool              `obfs:"tls,omitempty"`
	E1FxinJdhCGe map[string]string `obfs:"headers,omitempty"`
	R8YiK2Smi    bool              `obfs:"skip-cert-verify,omitempty"`
	JFsSe6xg     bool              `obfs:"mux,omitempty"`
}
type TKaQisW struct {
	*B8EHsVY7a
	dx07jQfKo   cipher.HjDHFt_
	pChyaQS6    string
	ar09vYxbfPZ *cJlUN29eP
	x93TNqC4dB  *plugin.YgrMuLrWXz
}

// ShadowsocksOptions is an analysis name for the original 128-byte by-value
// argument; offsets follow 0x1b5cd00. Original field names lack retained RTTI.
type ShadowsocksOptions struct {
	FFrN8gdu
	Name          string
	Server        string
	Port          int
	Password      string
	Cipher        string
	UDP           bool
	Plugin        string
	PluginOptions map[string]any
}

// PQ9jOjtX: 0x1b5cd00. Unknown plugins are silently ignored in this sample.
func PQ9jOjtX(o ShadowsocksOptions) (*TKaQisW, error) {
	addr := net.JoinHostPort(o.Server, strconv.Itoa(o.Port))
	c, e := cipher.B93MVPtD(o.Cipher, nil, o.Password)
	if e != nil {
		return nil, fmt.Errorf("ss %s initialize error: %w", addr, e)
	}
	d := decode.K6YuYvsMD5Y(decode.MBi603C{YONNn23: "obfs", RlQxpRzya: true})
	var simple *cJlUN29eP
	var ws *plugin.YgrMuLrWXz
	mode := ""
	switch o.Plugin {
	case "obfs":
		simple = &cJlUN29eP{OJdz_5FmLCKj: "bing.com"}
		if e = d.Decode(o.PluginOptions, simple); e != nil {
			return nil, fmt.Errorf("ss %s initialize obfs error: %w", addr, e)
		}
		if simple.VfjnRgEJIGGI != "tls" && simple.VfjnRgEJIGGI != "http" {
			return nil, fmt.Errorf("ss %s obfs mode error: %s", addr, simple.VfjnRgEJIGGI)
		}
		mode = simple.VfjnRgEJIGGI
	case "v2ray-plugin":
		opt := dXR4UapBiBa{L3gscTlpiGj: "bing.com", JFsSe6xg: true}
		if e = d.Decode(o.PluginOptions, &opt); e != nil {
			return nil, fmt.Errorf("ss %s initialize v2ray-plugin error: %w", addr, e)
		}
		if opt.GTwWMvxPP8 != "websocket" {
			return nil, fmt.Errorf("ss %s obfs mode error: %s", addr, opt.GTwWMvxPP8)
		}
		mode = opt.GTwWMvxPP8
		ws = &plugin.YgrMuLrWXz{PRkTz1l: opt.L3gscTlpiGj, G3GjlvRQW: opt.RhRqPp, BwX3DgZaf: opt.E1FxinJdhCGe, AUavUP: opt.JFsSe6xg}
		if opt.RZZqKB {
			ws.PwG_Wda = true
			ws.HjmCr8 = opt.R8YiK2Smi
		}
	}
	return &TKaQisW{B8EHsVY7a: &B8EHsVY7a{m50I26ZDL: o.Name, x3HUaU8Ln1S1: addr, xDan8hA5LV: o.PiTAGr6U5KDG, iOKxDl: 2, ntseIWz: o.UDP, bEbqkM8: o.VsBRTR}, dx07jQfKo: c, pChyaQS6: mode, ar09vYxbfPZ: simple, x93TNqC4dB: ws}, nil
}

// StreamConn: 0x1b5bce0. A destination-header error retains the wrapped conn.
func (s *TKaQisW) StreamConn(c net.Conn, m *E.D0mLgQF) (net.Conn, error) {
	switch s.pChyaQS6 {
	case "tls":
		c = obfs.NewTLSObfs(c, s.ar09vYxbfPZ.OJdz_5FmLCKj)
	case "http":
		_, port, _ := net.SplitHostPort(s.Addr())
		c = obfs.NewHTTPObfs(c, s.ar09vYxbfPZ.OJdz_5FmLCKj, port)
	case "websocket":
		var e error
		c, e = plugin.DO4qmnxxv4_(c, s.x93TNqC4dB)
		if e != nil {
			return nil, fmt.Errorf("%s connect error: %w", s.Addr(), e)
		}
	}
	c = s.dx07jQfKo.StreamConn(c)
	_, e := c.Write(o2HGjDHIVO(m))
	return c, e
}

// DialContext: 0x1b5c120; deferred cleanup closes the raw TCP connection.
func (s *TKaQisW) DialContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.PnyceLR, err error) {
	c, err := O.I_y9nlK1q8(ctx, "tcp", s.Addr(), s.DialOptions(options)...)
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %w", s.Addr(), err)
	}
	mS0F0Ti(c)
	defer func() {
		if err != nil {
			c.Close()
		}
	}()
	wrapped, err := s.StreamConn(c, m)
	if err != nil {
		return nil, err
	}
	return &lGta_1QkvQ{Conn: wrapped, bnOtGuS: E.GYe3KD{s.Name()}}, nil
}

// ListenPacketContext: 0x1b5c7e0. Metadata is ignored; plugins apply only to TCP.
func (s *TKaQisW) ListenPacketContext(ctx context.Context, _ *E.D0mLgQF, options ...O.PNSpXRb) (E.C08JTZ2l, error) {
	c, e := O.OpEY1B7hjf(ctx, "udp", "", s.DialOptions(options)...)
	if e != nil {
		return nil, e
	}
	addr, e := g0YPF3("udp", s.Addr())
	if e != nil {
		c.Close()
		return nil, e
	}
	return &uMtOz6o{PacketConn: &dTAcwl_gA{PacketConn: s.dx07jQfKo.PacketConn(c), zbHzsqAUlM: addr}, f7i_T5n5JXyM: E.GYe3KD{s.Name()}}, nil
}

var _ E.C74CUtQ9FqH = (*TKaQisW)(nil)
