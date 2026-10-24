package vmessaead

import (
	"context"
	"crypto/tls"
	"net"
	"time"
)

// TLSConfiguration is an analyst-named 48-byte option value. The original type
// name is absent from RTTI; 0x19b7e60 reads fields at offsets 0,16,24.
type TLSConfiguration struct {
	ServerName         string
	InsecureSkipVerify bool
	NextProtos         []string
}

// Zr3X_01: 0x19b7e60. Always returns the TLS connection, even on handshake error.
func Zr3X_01(raw net.Conn, o *TLSConfiguration) (net.Conn, error) {
	conn := tls.Client(raw, &tls.Config{ServerName: o.ServerName, InsecureSkipVerify: o.InsecureSkipVerify, NextProtos: o.NextProtos})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err := conn.HandshakeContext(ctx)
	cancel()
	return conn, err
}
