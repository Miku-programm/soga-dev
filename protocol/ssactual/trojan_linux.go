//go:build linux

package ssactual

import (
	"context"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	E "recovered.local/soga/net/outbound"
	G "recovered.local/soga/transport/httpconn"
	P "recovered.local/soga/crypto/sstypes"
	O "recovered.local/soga/net/socket"
	"strconv"
)

// Original shared transport options, with retained RTTI names and proxy tags.
type THKLdNRQ struct {
	IDDJaTZ string `proxy:"grpc-service-name,omitempty"`
}
type HrHLa45bjuV struct {
	TTsHbH5qb  string            `proxy:"path,omitempty"`
	RuSpNqlE   map[string]string `proxy:"headers,omitempty"`
	Z87qja35Ut int               `proxy:"max-early-data,omitempty"`
	IsFhLn3aBa string            `proxy:"early-data-header-name,omitempty"`
}

type ZxrFcPHYpol struct {
	FFrN8gdu
	AKe9rX4WI  string      `proxy:"name"`
	HF6Vma     string      `proxy:"server"`
	GRc9GM8NFW int         `proxy:"port"`
	Anl_Yj     string      `proxy:"password"`
	DhlWunAqa  []string    `proxy:"alpn,omitempty"`
	IGRMFmQSwT string      `proxy:"sni,omitempty"`
	R1aqZuYO   bool        `proxy:"skip-cert-verify,omitempty"`
	TNtpmYwuy  bool        `proxy:"udp,omitempty"`
	XIgNO9dyP  string      `proxy:"network,omitempty"`
	ZYpNOKrxXx THKLdNRQ    `proxy:"grpc-opts,omitempty"`
	JXjd1Bs    HrHLa45bjuV `proxy:"ws-opts,omitempty"`
}

type NLN_l6 struct {
	*B8EHsVY7a
	kmQiVgXZZ1  *P.MrjKETeSu
	sWg8Yg      *ZxrFcPHYpol
	oVYKaxP     *tls.Config
	eXjtOT      *G.EgMrkYjAnvwk
	vuo5piSzgld *http2.Transport
}

// RFqiXSepf_9: 0x1b6a1a0; raw gRPC dial closure: 0x1b736e0.
func RFqiXSepf_9(o ZxrFcPHYpol) (*NLN_l6, error) {
	config := &P.LckgbWP0Vg4{HQxCy9slSH: o.Anl_Yj, MOWJ8N: o.DhlWunAqa, DAPCgJns: o.HF6Vma, KJRBYI3fW: o.R1aqZuYO}
	if o.IGRMFmQSwT != "" {
		config.DAPCgJns = o.IGRMFmQSwT
	}
	t := &NLN_l6{
		B8EHsVY7a:  &B8EHsVY7a{m50I26ZDL: o.AKe9rX4WI, x3HUaU8Ln1S1: net.JoinHostPort(o.HF6Vma, strconv.Itoa(o.GRc9GM8NFW)), xDan8hA5LV: o.PiTAGr6U5KDG, iOKxDl: 8, ntseIWz: o.TNtpmYwuy, bEbqkM8: o.VsBRTR},
		kmQiVgXZZ1: P.NewClient(config), sWg8Yg: &o,
	}
	if o.XIgNO9dyP == "grpc" {
		t.oVYKaxP = &tls.Config{NextProtos: o.DhlWunAqa, MinVersion: tls.VersionTLS12, InsecureSkipVerify: config.KJRBYI3fW, ServerName: config.DAPCgJns}
		t.vuo5piSzgld = &http2.Transport{TLSClientConfig: t.oVYKaxP, DisableCompression: true, AllowHTTP: false, PingTimeout: 0,
			DialTLSContext: G.AljPVof590m(func(string, string) (net.Conn, error) {
				conn, err := O.I_y9nlK1q8(context.Background(), "tcp", t.Addr(), t.DialOptions(nil)...)
				if err != nil {
					return nil, fmt.Errorf("%s connect error: %s", t.Addr(), err.Error())
				}
				mS0F0Ti(conn)
				return conn, nil
			}),
		}
		t.eXjtOT = &G.EgMrkYjAnvwk{JbYoxUV: o.ZYpNOKrxXx.IDDJaTZ, EbklgbD: config.DAPCgJns}
	}
	return t, nil
}

// zGQVMayYr: 0x1b685e0. Trojan's websocket path forwards neither of the
// shared max-early-data/header-name fields; only an ed query can enable it.
func (t *NLN_l6) zGQVMayYr(conn net.Conn) (net.Conn, error) {
	if t.sWg8Yg.XIgNO9dyP != "ws" {
		return t.kmQiVgXZZ1.StreamConn(conn)
	}
	host, port, _ := net.SplitHostPort(t.Addr())
	options := &P.KTJbqZyG_iM{ZgOmB9r: host, XOOe7a: port, Cgk4_Jjow: t.sWg8Yg.JXjd1Bs.TTsHbH5qb}
	if t.sWg8Yg.IGRMFmQSwT != "" {
		options.ZgOmB9r = t.sWg8Yg.IGRMFmQSwT
	}
	if len(t.sWg8Yg.JXjd1Bs.RuSpNqlE) != 0 {
		options.VA7lpXIPEuY = make(http.Header)
		for key, value := range t.sWg8Yg.JXjd1Bs.RuSpNqlE {
			options.VA7lpXIPEuY.Add(key, value)
		}
	}
	return t.kmQiVgXZZ1.StreamWebsocketConn(conn, options)
}

// StreamConn: 0x1b688c0. On header failure the secured connection is returned.
func (t *NLN_l6) StreamConn(conn net.Conn, m *E.D0mLgQF) (net.Conn, error) {
	var err error
	if t.vuo5piSzgld != nil {
		conn, err = G.H6iElYbPlX4(conn, t.oVYKaxP, t.eXjtOT)
	} else {
		conn, err = t.zGQVMayYr(conn)
	}
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %w", t.Addr(), err)
	}
	return conn, t.kmQiVgXZZ1.WriteHeader(conn, 1, o2HGjDHIVO(m))
}

// DialContext: 0x1b68b80. Explicit options disable shared gRPC transport reuse.
func (t *NLN_l6) DialContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.PnyceLR, err error) {
	if t.vuo5piSzgld != nil && len(options) == 0 {
		conn, e := G.TwFKkx(t.vuo5piSzgld, t.eXjtOT)
		if e != nil {
			return nil, e
		}
		if e = t.kmQiVgXZZ1.WriteHeader(conn, 1, o2HGjDHIVO(m)); e != nil {
			conn.Close()
			return nil, e
		}
		return &lGta_1QkvQ{Conn: conn, bnOtGuS: E.GYe3KD{t.Name()}}, nil
	}
	conn, err := O.I_y9nlK1q8(ctx, "tcp", t.Addr(), t.DialOptions(options)...)
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %w", t.Addr(), err)
	}
	mS0F0Ti(conn)
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()
	secured, err := t.StreamConn(conn, m)
	if err != nil {
		return nil, err
	}
	return &lGta_1QkvQ{Conn: secured, bnOtGuS: E.GYe3KD{t.Name()}}, nil
}

// ListenPacketContext: 0x1b69580. The explicit-options branch uses zGQVMayYr
// directly, even when the configured network is grpc; this asymmetry is original.
func (t *NLN_l6) ListenPacketContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.C08JTZ2l, err error) {
	var conn net.Conn
	if t.vuo5piSzgld != nil && len(options) == 0 {
		conn, err = G.TwFKkx(t.vuo5piSzgld, t.eXjtOT)
		if err != nil {
			return nil, fmt.Errorf("%s connect error: %w", t.Addr(), err)
		}
		defer func() {
			if err != nil {
				conn.Close()
			}
		}()
	} else {
		raw, e := O.I_y9nlK1q8(ctx, "tcp", t.Addr(), t.DialOptions(options)...)
		if e != nil {
			return nil, fmt.Errorf("%s connect error: %w", t.Addr(), e)
		}
		defer func() {
			if err != nil {
				raw.Close()
			}
		}()
		mS0F0Ti(raw)
		conn, err = t.zGQVMayYr(raw)
		if err != nil {
			return nil, fmt.Errorf("%s connect error: %w", t.Addr(), err)
		}
	}
	if err = t.kmQiVgXZZ1.WriteHeader(conn, 3, o2HGjDHIVO(m)); err != nil {
		return nil, err
	}
	return &uMtOz6o{PacketConn: t.kmQiVgXZZ1.PacketConn(conn), f7i_T5n5JXyM: E.GYe3KD{t.Name()}}, nil
}

var _ E.C74CUtQ9FqH = (*NLN_l6)(nil)
