package LpKMLQuaAsk

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"recovered.local/soga/Hye5hBXG"
	"strings"
	"time"
)

// ResolverSlots names the three observed resolver-interface calls at itab
// offsets +32/+40/+48. Its original type and method names were not retained;
// these descriptive source names preserve the recovered signatures and slots.
type ResolverSlots interface {
	LookupAll(context.Context, string) ([]net.IP, error)
	LookupIPv4(context.Context, string) ([]net.IP, error)
	LookupIPv6(context.Context, string) ([]net.IP, error)
}

var customResolverAt5097930 ResolverSlots
var hostsAt5029418 = Hye5hBXG.NewEmptyTrie()

// Static state: disable byte 0x3787310=1; duration 0x3787948=5e9;
// resolver pointer 0x5029330 targets a zero-valued net.Resolver at 0x509b480.
var ipv6DisabledAt3787310 = true
var timeoutAt3787948 = 5 * time.Second
var resolverAt5029330 = new(net.Resolver)
var noIPAt5097940 = errors.New("couldn't find ip")
var versionErrorAt5097950 = errors.New("ip version error")
var ipv6DisabledErrorAt5097960 = errors.New("ipv6 disabled")

// init.0 (0x191cc20) initializes this original two-string list. Its additional
// fallback DNS functions are outside this restored lookup path.
var serversAt503D730 = []string{"114.114.114.114:53", "8.8.8.8:53"}

// IYMzyr8f: 0x191dcc0. Host-map values use To4, while literal IPv4 strings
// return ParseIP's 16-byte representation. The system lookup creates a new
// Background timeout and deliberately does not inherit the supplied context.
func IYMzyr8f(ctx context.Context, host string) ([]net.IP, error) {
	if node := hostsAt5029418.Search(host); node != nil {
		if ip := node.HJKQN58GQYL.(net.IP).To4(); ip != nil {
			return []net.IP{ip}, nil
		}
	}
	if ip := net.ParseIP(host); ip != nil {
		if strings.Contains(host, ":") {
			return nil, versionErrorAt5097950
		}
		return []net.IP{ip}, nil
	}
	if customResolverAt5097930 != nil {
		return customResolverAt5097930.LookupIPv4(ctx, host)
	}
	timed, cancel := context.WithTimeout(context.Background(), timeoutAt3787948)
	defer cancel()
	ips, err := resolverAt5029330.LookupIP(timed, "ip4", host)
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, noIPAt5097940
	}
	return ips, nil
}

// Jqw3ag: 0x191e400. The disable check precedes even literal and hosts lookup.
func Jqw3ag(ctx context.Context, host string) ([]net.IP, error) {
	if ipv6DisabledAt3787310 {
		return nil, ipv6DisabledErrorAt5097960
	}
	if node := hostsAt5029418.Search(host); node != nil {
		if ip := node.HJKQN58GQYL.(net.IP).To16(); ip != nil {
			return []net.IP{ip}, nil
		}
	}
	if ip := net.ParseIP(host); ip != nil {
		if !strings.Contains(host, ":") {
			return nil, versionErrorAt5097950
		}
		return []net.IP{ip}, nil
	}
	if customResolverAt5097930 != nil {
		return customResolverAt5097930.LookupIPv6(ctx, host)
	}
	timed, cancel := context.WithTimeout(context.Background(), timeoutAt3787948)
	defer cancel()
	ips, err := resolverAt5029330.LookupIP(timed, "ip6", host)
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, noIPAt5097940
	}
	return ips, nil
}

// Adyk1_7N: 0x191e280; K4HgV0sXi: 0x191ea20. These use math/rand,
// unlike the separately recovered cryptographic route selector.
func Adyk1_7N(host string) (net.IP, error) {
	ips, err := IYMzyr8f(context.Background(), host)
	return chooseIP(host, ips, err)
}
func K4HgV0sXi(host string) (net.IP, error) {
	ips, err := Jqw3ag(context.Background(), host)
	return chooseIP(host, ips, err)
}

// chooseIP factors the identical inlined selection/error paths at those VAs.
func chooseIP(host string, ips []net.IP, err error) (net.IP, error) {
	if err != nil {
		return nil, err
	}
	if len(ips) != 0 {
		return ips[rand.Intn(len(ips))], nil
	}
	return nil, fmt.Errorf("%w: %s", noIPAt5097940, host)
}

// EXJDpRK0: 0x191eba0. Unlike family-specific lookups, the all-family system
// branch uses the supplied context directly without adding a timeout.
func EXJDpRK0(ctx context.Context, host string, resolver ResolverSlots) ([]net.IP, error) {
	if node := hostsAt5029418.Search(host); node != nil {
		return []net.IP{node.HJKQN58GQYL.(net.IP)}, nil
	}
	if resolver != nil {
		if ipv6DisabledAt3787310 {
			return resolver.LookupIPv4(ctx, host)
		}
		return resolver.LookupAll(ctx, host)
	}
	if ipv6DisabledAt3787310 {
		return IYMzyr8f(ctx, host)
	}
	if ip := net.ParseIP(host); ip != nil {
		return []net.IP{ip}, nil
	}
	ips, err := resolverAt5029330.LookupIP(ctx, "ip", host)
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, noIPAt5097940
	}
	return ips, nil
}

// TEx0aVr56JU: 0x191eee0.
func TEx0aVr56JU(host string) (net.IP, error) {
	ips, err := EXJDpRK0(context.Background(), host, customResolverAt5097930)
	return chooseIP(host, ips, err)
}
