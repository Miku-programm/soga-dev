package CBbCOBqme

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	quic "github.com/apernet/quic-go"
	"github.com/apernet/quic-go/http3"
	"net"
	B "recovered.local/soga/BaEWsPy"
	O "recovered.local/soga/GLaePRKYM"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	"sync"
)

type P3_qJL5v struct {
	fhdn5Z       context.Context
	yeJIrGaD12   net.PacketConn
	p2xMAL1hk    *quic.Listener
	fpLylxmI     *InfdqOmZK
	eIECnHeZd8   chan J.Ki8urSHTx44
	hcouJtJtvDTP func() *B.D23y2888a_
	sDLEnpCX     sync.Once
}

func VOhu4P9d2b(ctx context.Context, conn net.PacketConn, obfs, password string, config *InfdqOmZK, resolver func() *B.D23y2888a_) (*P3_qJL5v, error) { //0x1c31f20
	tlsConfig := http3.ConfigureTLSConfig(&tls.Config{Certificates: config.DGraqdoRv9ot.GinGbR, GetCertificate: config.DGraqdoRv9ot.EG8WQwgS, EncryptedClientHelloKeys: config.DGraqdoRv9ot.FQaEaYFoFF})
	q := config.I3PHCWIdg8EH
	quicConfig := &quic.Config{InitialStreamReceiveWindow: q.UIFPM249f2, MaxStreamReceiveWindow: q.ET3EbVcYR, InitialConnectionReceiveWindow: q.UCjF7o6elmz, MaxConnectionReceiveWindow: q.CtEnIL0gS, MaxIdleTimeout: q.YldcvpLEKua1, MaxIncomingStreams: q.VAZK_diUd, DisablePathMTUDiscovery: q.EJDv3cXV8, EnableDatagrams: true}
	switch obfs {
	case "", "plain":
	case "salamander":
		o, err := O.Ea9POnRF([]byte(password))
		if err != nil {
			return nil, err
		}
		conn = O.WrapFromHysteriaInline(conn, o)
	default:
		return nil, fmt.Errorf("unknown obfs: %v", obfs)
	}
	l, err := quic.Listen(conn, tlsConfig, quicConfig)
	if err != nil {
		return nil, err
	}
	v := &P3_qJL5v{fhdn5Z: ctx, yeJIrGaD12: conn, p2xMAL1hk: l, fpLylxmI: config, eIECnHeZd8: make(chan J.Ki8urSHTx44, 1024), hcouJtJtvDTP: resolver}
	go func() { //0x1c34ae0
		for {
			conn, err := l.Accept(ctx)
			if err != nil {
				return
			}
			go v.kIY21xyiE0Pa(conn)
		}
	}()
	return v, nil
}

func (v *P3_qJL5v) kIY21xyiE0Pa(conn quic.Connection) { //0x1c32620
	h := &cYzxTBfcV{szJp1HM: v.fhdn5Z, ff2_VHDMO_e: v.fpLylxmI, lsVoia1: conn, ehiG0t4uYJad: v.hcouJtJtvDTP().For(conn.LocalAddr()), dXJZalu06: v.eIECnHeZd8}
	server := &http3.Server{Handler: h, StreamHijacker: h.ProxyStreamHijacker}
	server.ServeQUICConn(conn)
	conn.CloseWithError(0x100, "")
}

func (v *P3_qJL5v) AcceptProxy() (J.Ki8urSHTx44, error) { //0x1c32900
	c := <-v.eIECnHeZd8
	if c == nil {
		return nil, errors.New("hysteria listener closed")
	}
	return c, nil
}
func (v *P3_qJL5v) Accept() (net.Conn, error) { return v.AcceptProxy() } //0x1c32a00
func (v *P3_qJL5v) Close() error { //0x1c32aa0
	v.sDLEnpCX.Do(func() { close(v.eIECnHeZd8) })
	return H.T2i1ItMaa(v.p2xMAL1hk.Close(), v.yeJIrGaD12.Close())
}
func (v *P3_qJL5v) Addr() net.Addr { return v.yeJIrGaD12.LocalAddr() } //0x1c32c40
