// Package EP3M1uC restores routing code from the supplied executable.
// Field names and tags follow RTTI; this is not the original source text.
package EP3M1uC

import (
	"net"
	"recovered.local/soga/E3F12urwswL"
	"sync"
	"sync/atomic"
)

type Fdaw35mqIz string
type U0f6rIw7p string
type Jw53oR7lC5 string
type WoOkPbuvC_h string
type J_enG0wE []*GyyO51
type ZLdFKPeDrCJ interface {
	Select(J_enG0wE, net.IP, int64) *GyyO51
}
type h1kE2OhgN struct{}
type twxaE4km_U struct{}
type gMxUOGd struct{}
type bLJQqABD struct{ v1TcNPAI0 atomic.Uint64 }

// GyyO51: RTTI 0x3733f08, size 448 on amd64.
type GyyO51 struct {
	HG8A97UIPf1A string     `toml:"name"`
	NyLbiXb4OrlL Fdaw35mqIz `toml:"type"`
	DM8KdAE6ZiZu string     `toml:"listen"`
	ZK61Na2nsqVQ string     `toml:"listen_ipv4"`
	LXpap7F      string     `toml:"listen_ipv6"`
	JaH5hv       U0f6rIw7p  `toml:"listen_strategy"`
	AHTlIEynXL   string     `toml:"server"`
	JO1YIpXk     int        `toml:"port"`
	PJkA7ZsW4FB  string     `toml:"uuid"`
	UkmhqYCQ9u   bool       `toml:"tls"`
	SRegcRbFQo   bool       `toml:"skip_cert_verify"`
	JqKHTUS0FUa  string     `toml:"sni"`
	FegVmoOB     int        `toml:"alter_id"`
	LUjxWBsS2Jl  Jw53oR7lC5 `toml:"network"`
	AJMYgKaXJz   string     `toml:"ws_path"`
	ZNfRIgJOM4uR string     `toml:"service_name"`
	CiOEvesmFF   string     `toml:"password"`
	ExA8yD       string     `toml:"cipher"`
	TmSpNP       string     `toml:"obfs"`
	Ca3YcWv688v  string     `toml:"obfs_param"`
	JuTg4y       string     `toml:"protocol"`
	P57BOZ       string     `toml:"protocol_param"`
	XFrkscEI     string     `toml:"username"`
	ucp4iC       E3F12urwswL.C74CUtQ9FqH
	iCySnjtd     net.IP
	xk9pAT       net.IP
	mJTqKai      sync.RWMutex
	fSEvC0BgQ    bool
	pJpTed       int
	qoxZCsWf     int
	vWopFdjg     bool
}
