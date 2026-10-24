package routing

import (
	"fmt"
	"strings"
	"testing"
)

func TestTOMLConstructsEveryOutbound(t *testing.T) {
	for _, kind := range []string{"direct", "redirect", "ss", "ssr", "http", "socks", "vmess", "trojan"} {
		t.Run(kind, func(t *testing.T) {
			config := fmt.Sprintf(`health_check = false
[[routes]]
rules = ["*"]
balance = "round_robin"
[[routes.outs]]
type = %q
server = "127.0.0.1"
port = 443
password = "test-secret"
cipher = "aes-128-cfb"
uuid = "00000000-0000-4000-8000-000000000001"
obfs = "plain"
protocol = "origin"
`, kind)
			routes, err := KmKjIZi1(config)
			if err != nil {
				t.Fatal(err)
			}
			if !routes.CnK0a4 || routes.H839j9DYDD5k || len(routes.MMKbp60np) != 1 {
				t.Fatalf("bad configuration: %+v", routes)
			}
			rule := routes.MMKbp60np[0]
			out := rule.UwkgrlFf8[0]
			if out.LUjxWBsS2Jl != "tcp" || out.JaH5hv != "ipv4_first" || out.JqKHTUS0FUa != "127.0.0.1" {
				t.Fatalf("defaults lost: %+v", out)
			}
			if !rule.__ev4XIbr || len(*rule.jf_lZ8wS2GNO.Load()) != 1 {
				t.Fatal("matcher or enabled snapshot missing")
			}
			if _, ok := rule.yvcsQOnq.(*bLJQqABD); !ok {
				t.Fatal("balance selection lost")
			}
			if kind == "redirect" {
				if out.ucp4iC != nil {
					t.Fatal("redirect acquired adapter")
				}
				return
			}
			if out.ucp4iC == nil {
				t.Fatal("adapter absent")
			}
			if kind != "direct" && out.ucp4iC.Addr() != "127.0.0.1:443" {
				t.Fatal(out.ucp4iC.Addr())
			}
			if kind != "http" && !out.ucp4iC.SupportUDP() {
				t.Fatal("UDP template lost")
			}
		})
	}
}

func TestDisabledRoutesKeepDefaultsButSkipValidation(t *testing.T) {
	routes, err := KmKjIZi1("enable=false\n[[routes]]\n[[routes.outs]]\ntype='unknown'\n")
	if err != nil {
		t.Fatal(err)
	}
	out := routes.MMKbp60np[0].UwkgrlFf8[0]
	if routes.CnK0a4 || !routes.H839j9DYDD5k || out.LUjxWBsS2Jl != "tcp" || out.ucp4iC != nil {
		t.Fatal("disabled initialization order changed")
	}
	_, err = KmKjIZi1("[[routes]]\n[[routes.outs]]\ntype='ss'\n")
	if err == nil || !strings.Contains(err.Error(), "please set out server") {
		t.Fatalf("outbound failure must precede empty matcher failure: %v", err)
	}
}

func TestRouteHotReloadPreservesPreviousStateOnFailure(t *testing.T) {
	oldRoutes, oldManager := routesAt5029218, healthManagerAt5095C60
	routesAt5029218, healthManagerAt5095C60 = &UIfnUsim{}, nil
	t.Cleanup(func() { tCGadgsq(nil); routesAt5029218, healthManagerAt5095C60 = oldRoutes, oldManager })
	if err := yFE5sGa7LHpa([]byte("enable=true\nhealth_check=true\n")); err != nil {
		t.Fatal(err)
	}
	active, manager := CurrentRoutes(), healthManagerAt5095C60
	if manager == nil {
		t.Fatal("health manager absent")
	}
	for _, bad := range []string{"[[", "[[routes]]\n"} {
		if err := yFE5sGa7LHpa([]byte(bad)); err == nil {
			t.Fatal("invalid configuration accepted")
		}
		if CurrentRoutes() != active || healthManagerAt5095C60 != manager {
			t.Fatal("failure replaced active state")
		}
	}
	if err := yFE5sGa7LHpa(nil); err != nil {
		t.Fatal(err)
	}
	if CurrentRoutes() != active || healthManagerAt5095C60 != manager {
		t.Fatal("empty input replaced active state")
	}
	if err := yFE5sGa7LHpa([]byte("enable=false\n")); err != nil {
		t.Fatal(err)
	}
	if CurrentRoutes() == active || healthManagerAt5095C60 != nil {
		t.Fatal("disabled replacement did not stop health manager")
	}
	select {
	case <-manager.bXx0sIP.Done():
	default:
		t.Fatal("previous manager context not canceled")
	}
}

func TestRoutesPathPrefixPreservesUncleanSegments(t *testing.T) {
	old := routesPathAt5096F90
	defer func() { routesPathAt5096F90 = old }()
	for input, want := range map[string]string{"": "/etc/soga/", "../routes.toml": "/etc/soga/../routes.toml", "/tmp/../routes.toml": "/tmp/../routes.toml"} {
		Zma0TSvL8(input)
		if routesPathAt5096F90 != want {
			t.Fatalf("%q: %q", input, routesPathAt5096F90)
		}
	}
}
