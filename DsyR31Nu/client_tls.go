package DsyR31Nu

import (
	"crypto/tls"
	utls "github.com/refraction-networking/utls"
	"net"
)

// Recovered address names below describe globals originally owned by
// GDNr5LruOqC. Only JZejsfshGB's dependencies are restored here, not that
// package's complete initializer. Original config@509D928 is 72 bytes;
// init@10A86A0 creates non-nil empty slices and a 256-entry TLS session cache.
// No public configuration API is inferred from these otherwise unwritten globals.
type clientTLSSettings struct {
	enabled, useUTLS       bool
	nextProtos             []string
	minVersion, maxVersion uint16
	cipherSuites           []uint16
	sessionCache           bool
}

var clientTLSSettingsAt509D928 = clientTLSSettings{
	nextProtos: make([]string, 0), cipherSuites: make([]uint16, 0),
}
var clientSessionCacheAt50975C0 = tls.NewLRUClientSessionCache(256)
var clientHelloAt5045F40 = utls.ClientHelloID{Client: "Randomized", Version: "0"}

// JZejsfshGB is the lazy client TLS wrapper at original VA 0x1171A60.
// TLS handshaking is deferred until the wrapped connection performs I/O.
func JZejsfshGB(raw net.Conn, skipVerify bool, serverName string) net.Conn {
	if clientTLSSettingsAt509D928.enabled {
		if clientTLSSettingsAt509D928.useUTLS {
			return utls.UClient(raw, &utls.Config{InsecureSkipVerify: skipVerify, ServerName: serverName}, clientHelloAt5045F40)
		}
		s := &clientTLSSettingsAt509D928
		c := &tls.Config{InsecureSkipVerify: skipVerify, ServerName: serverName}
		if len(s.nextProtos) != 0 {
			c.NextProtos = s.nextProtos
		}
		if s.minVersion != 0 {
			c.MinVersion = s.minVersion
		}
		if s.maxVersion != 0 {
			c.MaxVersion = s.maxVersion
		}
		if len(s.cipherSuites) != 0 {
			c.CipherSuites = s.cipherSuites
		}
		if s.sessionCache {
			c.ClientSessionCache = clientSessionCacheAt50975C0
		}
		return tls.Client(raw, c)
	}
	return tls.Client(raw, &tls.Config{InsecureSkipVerify: skipVerify, ServerName: serverName})
}
