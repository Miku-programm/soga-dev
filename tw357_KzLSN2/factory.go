package tw357_KzLSN2

import (
	"fmt"
	"io"
	E "recovered.local/soga/hfOuV8fOe4L"
)

// Original init@1ca71a0 and init.0..5 register these seven exact keys.
// http_post and http_simple deliberately share the same reader type.
var obfsAt5096968 = map[E.PMorzD7u]q1_ASanae22{
	"http_post":          zbGjqkTbNkOk,
	"http_simple":        ujOdSji25S,
	"plain":              aR_Feln8A36,
	"none":               aR_Feln8A36,
	"simple_obfs_http":   i0Z5v8x,
	"simple_obfs_tls":    crpBB1_mZ,
	"tls1.2_ticket_auth": jRGJRCFN,
}

// Z_Mqu9of restores the native inbound factory at 0x1ca7260.
func Z_Mqu9of(name E.PMorzD7u, rw io.ReadWriter) (Jakq78yIQ, error) {
	if makeObfs, ok := obfsAt5096968[name]; ok {
		return makeObfs(rw), nil
	}
	return nil, fmt.Errorf("obfs %v not support", name)
}
