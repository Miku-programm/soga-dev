package CpHCdepV

import (
	"crypto/tls"
	"fmt"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/VasCJz"
	"recovered.local/soga/g0MlUOhvO"
	"recovered.local/soga/reality"
)

// Ap9uD2 is the original named 112-byte listener argument. Its method names
// survive pclntab, but no struct RTTI descriptor was found. The field names
// below are explicit analysis labels from NewListener and its callers' exact
// offsets, not recovered author identifiers.
type Ap9uD2 struct {
	Host               string                  // +0
	Port               int                     // +16
	Network            LFWTTXF7JR.BoulTn31     // +24
	TLS                *tls.Config             // +40
	Reality            *reality.Config         // +48
	Path               string                  // +56: HTTP/2 and WebSocket path
	ServiceName        string                  // +72: gRPC service
	ProxyProtocol      bool                    // +88
	ForceProxyProtocol bool                    // +89
	Tunnel             *g0MlUOhvO.VTl94pzz2IkB // +96
	XHTTP              *VasCJz.R0J19auWof      // +104
}

func (o *Ap9uD2) String() string { //0x18b9ea0
	security := "no tls"
	if o.TLS != nil {
		security = "tls"
	} else if o.Reality != nil {
		security = "reality"
	}
	return fmt.Sprintf("(%v:%v) (%v + %v) (proxy protocol: %v) (force proxy protocol: %v)", o.Host, o.Port, o.Network, security, o.ProxyProtocol, o.ForceProxyProtocol)
}
