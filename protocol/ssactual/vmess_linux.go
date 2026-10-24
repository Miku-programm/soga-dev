//go:build linux

package ssactual

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	E "recovered.local/soga/net/outbound"
	V "recovered.local/soga/transport/vmessaead"
	G "recovered.local/soga/transport/httpconn"
	D "recovered.local/soga/net/dnssystem"
	O "recovered.local/soga/net/socket"
	"strconv"
	"strings"
)

type OPYb56D struct {
	MrTNulEgJrCT []string `proxy:"host,omitempty"`
	WvC0eaw      string   `proxy:"path,omitempty"`
}
type EyDONQ3H9Jz struct {
	UmSoRT    string              `proxy:"method,omitempty"`
	Wt0Mo6    []string            `proxy:"path,omitempty"`
	MpPwkigf3 map[string][]string `proxy:"headers,omitempty"`
}
type YaiP1Z10VuOK struct {
	FFrN8gdu
	NwFIZN       string      `proxy:"name"`
	PaWTH6       string      `proxy:"server"`
	KQ7hyM       int         `proxy:"port"`
	MMeZdiHjQ    string      `proxy:"uuid"`
	TJqENyK7     int         `proxy:"alterId"`
	YKnBs48P7XY  string      `proxy:"cipher"`
	DS2NjE9s_    bool        `proxy:"udp,omitempty"`
	EtDA3oIPdZ7y string      `proxy:"network,omitempty"`
	MLU1J_       bool        `proxy:"tls,omitempty"`
	VPGpSFLRW    bool        `proxy:"skip-cert-verify,omitempty"`
	JPPDxQ       string      `proxy:"servername,omitempty"`
	KphAInyo     EyDONQ3H9Jz `proxy:"http-opts,omitempty"`
	GMSZn2d0QZm  OPYb56D     `proxy:"h2-opts,omitempty"`
	FN179GVH_    THKLdNRQ    `proxy:"grpc-opts,omitempty"`
	U7nq5HsQY    HrHLa45bjuV `proxy:"ws-opts,omitempty"`
}
type UpVCB2mE struct {
	*B8EHsVY7a
	p61uGslT     *V.GC3qACg18wgq
	vL1N8AU_5    *YaiP1Z10VuOK
	o7IafTDLYDr  *tls.Config
	gw0sQtIBnRQp *G.EgMrkYjAnvwk
	nmGdkI       *http2.Transport
}

// Fr9G0QQhtaac: 0x1b6e7e0. The full int alterId selects AEAD before its
// uint16 truncation for the legacy-ID constructor. Only cipher is lowercased.
func Fr9G0QQhtaac(o YaiP1Z10VuOK) (*UpVCB2mE, error) {
	core, err := V.Iw0IEMM3ha(V.ClientConfiguration{UUID: o.MMeZdiHjQ, AlterID: uint16(o.TJqENyK7), Security: strings.ToLower(o.YKnBs48P7XY), Port: strconv.Itoa(o.KQ7hyM), Host: o.PaWTH6, AEAD: o.TJqENyK7 == 0})
	if err != nil {
		return nil, err
	}
	if (o.EtDA3oIPdZ7y == "h2" || o.EtDA3oIPdZ7y == "grpc") && !o.MLU1J_ {
		return nil, errors.New("TLS must be true with h2/grpc network")
	}
	v := &UpVCB2mE{B8EHsVY7a: &B8EHsVY7a{m50I26ZDL: o.NwFIZN, x3HUaU8Ln1S1: net.JoinHostPort(o.PaWTH6, strconv.Itoa(o.KQ7hyM)), xDan8hA5LV: o.PiTAGr6U5KDG, iOKxDl: 7, ntseIWz: o.DS2NjE9s_, bEbqkM8: o.VsBRTR}, p61uGslT: core, vL1N8AU_5: &o}
	if o.EtDA3oIPdZ7y == "h2" && len(o.GMSZn2d0QZm.MrTNulEgJrCT) == 0 {
		o.GMSZn2d0QZm.MrTNulEgJrCT = append(o.GMSZn2d0QZm.MrTNulEgJrCT, "www.example.com")
	}
	if o.EtDA3oIPdZ7y == "grpc" {
		name := o.JPPDxQ
		if name == "" {
			name, _, _ = net.SplitHostPort(v.Addr())
		}
		v.o7IafTDLYDr = &tls.Config{InsecureSkipVerify: o.VPGpSFLRW, ServerName: name}
		v.gw0sQtIBnRQp = &G.EgMrkYjAnvwk{JbYoxUV: o.FN179GVH_.IDDJaTZ, EbklgbD: name}
		v.nmGdkI = &http2.Transport{TLSClientConfig: v.o7IafTDLYDr, DisableCompression: true, AllowHTTP: false, PingTimeout: 0, DialTLSContext: G.AljPVof590m(func(string, string) (net.Conn, error) {
			conn, e := O.I_y9nlK1q8(context.Background(), "tcp", v.Addr(), v.DialOptions(nil)...)
			if e != nil {
				return nil, fmt.Errorf("%s connect error: %s", v.Addr(), e.Error())
			}
			mS0F0Ti(conn)
			return conn, nil
		})}
	}
	return v, nil
}

// StreamConn: 0x1b6c640. Transport failures pass through unchanged.
func (v *UpVCB2mE) StreamConn(conn net.Conn, m *E.D0mLgQF) (net.Conn, error) {
	o := v.vL1N8AU_5
	host, port, _ := net.SplitHostPort(v.Addr())
	name := host
	if o.JPPDxQ != "" {
		name = o.JPPDxQ
	}
	tlsOptions := &V.TLSConfiguration{ServerName: name, InsecureSkipVerify: o.VPGpSFLRW}
	var err error
	switch o.EtDA3oIPdZ7y {
	case "h2":
		tlsOptions.NextProtos = []string{"h2"}
		conn, err = V.Zr3X_01(conn, tlsOptions)
		if err != nil {
			return nil, err
		}
		conn, err = V.P3TqaD0kr(conn, &V.ALZXSf{Hmv1xGaNFZBN: o.GMSZn2d0QZm.MrTNulEgJrCT, B41d7b: o.GMSZn2d0QZm.WvC0eaw})
	case "ws":
		options := &V.ELIQWabl31{OhyeDUGTCGH: host, RTIB8Nkai: port, WOT5UbYYto: o.U7nq5HsQY.TTsHbH5qb, A3jvUyhB: o.U7nq5HsQY.Z87qja35Ut, I_3sggK: o.U7nq5HsQY.IsFhLn3aBa}
		if len(o.U7nq5HsQY.RuSpNqlE) != 0 {
			options.RlZC_oDxj = make(http.Header)
			for k, value := range o.U7nq5HsQY.RuSpNqlE {
				options.RlZC_oDxj.Add(k, value)
			}
		}
		if o.MLU1J_ {
			options.Xtp2zWRK = true
			options.VbB82ra = &tls.Config{ServerName: host, InsecureSkipVerify: o.VPGpSFLRW, NextProtos: []string{"http/1.1"}}
			if o.JPPDxQ != "" {
				options.VbB82ra.ServerName = o.JPPDxQ
			} else if h := options.RlZC_oDxj.Get("Host"); h != "" {
				options.VbB82ra.ServerName = h
			}
		}
		conn, err = V.E6YB0lYJp(conn, options)
	case "grpc":
		conn, err = G.H6iElYbPlX4(conn, v.o7IafTDLYDr, v.gw0sQtIBnRQp)
	case "http":
		if o.MLU1J_ {
			conn, err = V.Zr3X_01(conn, tlsOptions)
			if err != nil {
				return nil, err
			}
		}
		conn = V.NewHTTPObfuscationConn(conn, &V.R5Z6tXgFroMW{YORtUpVbx: o.KphAInyo.UmSoRT, V7kc7NuILllS: host, DHZvNLM2Sc0K: o.KphAInyo.Wt0Mo6, DaWC9tFk2l4E: o.KphAInyo.MpPwkigf3})
	default:
		if o.MLU1J_ {
			conn, err = V.Zr3X_01(conn, tlsOptions)
		}
	}
	if err != nil {
		return nil, err
	}
	return v.p61uGslT.StreamConn(conn, eEgDc4(m))
}

// DialContext: 0x1b6d1e0. The direct path returns a nonnil chain wrapper even
// when StreamConn fails, whereas the shared-gRPC path returns nil on that error.
func (v *UpVCB2mE) DialContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.PnyceLR, err error) {
	if v.nmGdkI != nil && len(options) == 0 {
		conn, e := G.TwFKkx(v.nmGdkI, v.gw0sQtIBnRQp)
		if e != nil {
			return nil, e
		}
		defer func() {
			if err != nil {
				conn.Close()
			}
		}()
		secured, e := v.p61uGslT.StreamConn(conn, eEgDc4(m))
		if e != nil {
			return nil, e
		}
		return &lGta_1QkvQ{Conn: secured, bnOtGuS: E.GYe3KD{v.Name()}}, nil
	}
	conn, err := O.I_y9nlK1q8(ctx, "tcp", v.Addr(), v.DialOptions(options)...)
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %s", v.Addr(), err.Error())
	}
	mS0F0Ti(conn)
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()
	secured, err := v.StreamConn(conn, m)
	return &lGta_1QkvQ{Conn: secured, bnOtGuS: E.GYe3KD{v.Name()}}, err
}

// ListenPacketContext: 0x1b6dbe0. DNS mutates destination IP but retains Host.
func (v *UpVCB2mE) ListenPacketContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.C08JTZ2l, err error) {
	if m.ZouEpmjw8 == nil {
		ip, e := D.TEx0aVr56JU(m.WJ95LyOL)
		if e != nil {
			return nil, errors.New("can't resolve ip")
		}
		m.ZouEpmjw8 = ip
	}
	var conn, secured net.Conn
	if v.nmGdkI != nil && len(options) == 0 {
		conn, err = G.TwFKkx(v.nmGdkI, v.gw0sQtIBnRQp)
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				conn.Close()
			}
		}()
		secured, err = v.p61uGslT.StreamConn(conn, eEgDc4(m))
	} else {
		conn, err = O.I_y9nlK1q8(ctx, "tcp", v.Addr(), v.DialOptions(options)...)
		if err != nil {
			return nil, fmt.Errorf("%s connect error: %s", v.Addr(), err.Error())
		}
		mS0F0Ti(conn)
		defer func() {
			if err != nil {
				conn.Close()
			}
		}()
		secured, err = v.StreamConn(conn, m)
	}
	if err != nil {
		return nil, fmt.Errorf("new vmess client error: %v", err)
	}
	var addr *net.UDPAddr
	if m.AcXafZ1 == 1 && m.ZouEpmjw8 != nil {
		addr = &net.UDPAddr{IP: m.ZouEpmjw8, Port: int(m.JRXMVvg)}
	}
	return &uMtOz6o{PacketConn: &keoydRNdlJ{Conn: secured, oDRJBoi: addr}, f7i_T5n5JXyM: E.GYe3KD{v.Name()}}, nil
}

var _ E.C74CUtQ9FqH = (*UpVCB2mE)(nil)
