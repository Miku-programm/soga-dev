package CbO_7wp3B

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

type certificatePanelFixture struct {
	node  *L.Wr5SVP
	err   error
	calls int
}

func (*certificatePanelFixture) GetAPIConfig() *C.FCEAfQJwDA7e { return nil }
func (p *certificatePanelFixture) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	p.calls++
	return &Q.D5gTIZ5aVg{H47AzZsYHB1N: p.node}, p.err
}
func (*certificatePanelFixture) SubmitData(*Q.FGmNba4MhioY) error { return nil }

type certificateNodeFixture struct {
	certificatePanelFixture
	nodeCalls int
}

func (p *certificateNodeFixture) GetNode() (*L.Wr5SVP, error) { p.nodeCalls++; return p.node, p.err }

func certificateConfig(t *testing.T, content string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "soga.conf")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(path)
	if err := C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
}

func TestCertificateRequirementsByNodeAndProtocol(t *testing.T) {
	old := C.ConfigurationPath()
	defer C.MkwNWciPHS(old)
	for _, protocol := range []string{"v2ray", "vless", "trojan", "anytls"} {
		for _, security := range []string{"tls", "reality", "none", ""} {
			for _, optional := range []bool{false, true} {
				certificateConfig(t, "server_type="+protocol+"\n")
				node := &L.Wr5SVP{}
				node.SwxJNgPpg.UVeahqDBT = security == "tls"
				node.Ck6DxjB.WAdJ4ftdZFk = L.D7FHIoaa8s4a(security)
				node.JhIYwEukjd.AJjaqrBjj = L.D7FHIoaa8s4a(security)
				node.QaZjDzw95.J6bt8AOW8fL8 = L.D7FHIoaa8s4a(security)
				base := &certificatePanelFixture{node: node}
				fast := &certificateNodeFixture{certificatePanelFixture: *base}
				var panel Q.VyFk5uT = base
				if optional {
					panel = fast
				}
				got, err := d_7qy2u([]Q.VyFk5uT{panel})
				want := security == "tls" || ((protocol == "trojan" || protocol == "anytls") && security == "")
				if err != nil || got != want {
					t.Fatalf("%s/%s optional=%v: %v %v", protocol, security, optional, got, err)
				}
				if optional {
					if fast.nodeCalls != 1 || fast.calls != 0 {
						t.Fatal("optional GetNode path not preferred")
					}
				} else if base.calls != 1 {
					t.Fatal("fallback GetAPIData not used")
				}
			}
		}
	}
}

func TestCertificateFetchFailuresAndEarlyReturns(t *testing.T) {
	old := C.ConfigurationPath()
	defer C.MkwNWciPHS(old)
	sentinel := errors.New("node fetch failed")
	for _, protocol := range []string{"v2ray", "vless", "trojan", "anytls"} {
		certificateConfig(t, "server_type="+protocol+"\n")
		got, err := d_7qy2u([]Q.VyFk5uT{&certificatePanelFixture{err: sentinel}})
		fallback := protocol == "trojan" || protocol == "anytls"
		if got != fallback || (fallback && err != nil) || (!fallback && err != sentinel) {
			t.Fatalf("%s: %v %v", protocol, got, err)
		}
	}
	for _, tc := range []struct {
		config string
		want   bool
	}{
		{"server_type=hysteria\nforce_close_ssl=true\n", true},
		{"server_type=ss\n", false},
		{"server_type=ssr\n", false},
		{"server_type=trojan\nforce_close_ssl=true\n", false},
		{"server_type=hysteria\ncert_file=one\nkey_file=two\n", false},
		{"server_type=hysteria\ncert_file=one\n", true},
	} {
		certificateConfig(t, tc.config)
		p := &certificatePanelFixture{err: sentinel}
		got, err := d_7qy2u([]Q.VyFk5uT{p})
		if err != nil || got != tc.want || p.calls != 0 {
			t.Fatalf("%q: %v %v calls=%d", tc.config, got, err, p.calls)
		}
	}
}

func TestEncryptedTunnelRetainsOriginalAcceptedNames(t *testing.T) {
	old := C.ConfigurationPath()
	defer C.MkwNWciPHS(old)
	for _, mode := range []string{"tls", "tls-tunnel", "quic-tunnel", "ws-tls-tunnel", "tcp-tunnel", "ws-tunnel", "invalid"} {
		certificateConfig(t, "server_type=ss\ntunnel_enable=true\ntunnel_type="+mode+"\nsoga_key=test-nonempty\n")
		C.T15Jifo(false)
		want := mode == "tls-tunnel" || mode == "quic-tunnel" || mode == "ws-tls-tunnel"
		if got := gcTStJ6_zEzn(); got != want {
			t.Fatalf("%s: %v", mode, got)
		}
		C.T15Jifo(true)
		if gcTStJ6_zEzn() {
			t.Fatal("key state no longer gates tunnel branch")
		}
	}
	C.T15Jifo(false)
}
