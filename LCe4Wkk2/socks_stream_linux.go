//go:build linux

package LCe4Wkk2

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	S "recovered.local/soga/Bx4XCl"
	E "recovered.local/soga/E3F12urwswL"
	O "recovered.local/soga/vP_e5OQAve"
	"time"
)

// StreamConn: 0x1b64340. TLS gets its own five-second background deadline;
// the following SOCKS negotiation has no additional deadline in this method.
func (s *AwDE24aRWC) StreamConn(conn net.Conn, m *E.D0mLgQF) (net.Conn, error) {
	if s.m0JWRLb2CH {
		secured := tls.Client(conn, s.idTOPJOmPpB)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := secured.HandshakeContext(ctx); err != nil {
			return nil, fmt.Errorf("%s connect error: %w", s.Addr(), err)
		}
		conn = secured
	}
	var auth *S.Auth
	if s.gmXALhCpyKF != "" {
		auth = &S.Auth{Username: s.gmXALhCpyKF, Password: s.eOL60dT1}
	}
	if _, err := S.AMwLSL6GewF(conn, o2HGjDHIVO(m), 1, auth); err != nil {
		return nil, err
	}
	return conn, nil
}

// DialContext: 0x1b648a0; failure-close closure: 0x1b73220. The original
// connection is closed only when StreamConn returns an error.
func (s *AwDE24aRWC) DialContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.PnyceLR, err error) {
	conn, err := O.I_y9nlK1q8(ctx, "tcp", s.Addr(), s.DialOptions(options)...)
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %w", s.Addr(), err)
	}
	mS0F0Ti(conn)
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()
	secured, err := s.StreamConn(conn, m)
	if err != nil {
		return nil, err
	}
	return &lGta_1QkvQ{Conn: secured, bnOtGuS: E.GYe3KD{s.Name()}}, nil
}
