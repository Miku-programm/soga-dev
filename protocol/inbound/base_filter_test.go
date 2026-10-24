package inbound

import (
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/user"
	"recovered.local/soga/net/sniff"
	"recovered.local/soga/audit/matcher"
	"testing"
)

type sniffFixture struct {
	kind sniff.MCe3PxlaMYM
	host string
}

func (s *sniffFixture) Type() sniff.MCe3PxlaMYM                 { return s.kind }
func (s *sniffFixture) Value() string                             { return s.host }
func (s *sniffFixture) Network() relay.CjhtWb                  { return 0 }
func (s *sniffFixture) Sniff([]byte) (sniff.FlCuFjj_V1u, error) { return s, nil }
func TestAddressRulesWhitelistPriorityAndReport(t *testing.T) {
	b := &JWN154qJ4{qE9aU9: &genmap.GTFAWF[int, *genmap.Fxs7SY[int64]]{}}
	white, _, _, _ := elCwduDdyGl(matcher.VHWq49iRat{"full:white.test", "ip:192.0.2.3", "port:999"})
	b._IoxtH.Store(white)
	black, _, _, _ := elCwduDdyGl(matcher.VHWq49iRat{"domain:test", "ip:198.51.100.0/24", "port:443"})
	b.oIgEuicmPL.Store(black)
	detectors := matcher.Ij9r21E0UOt3{matcher.NewPanelDetector(55, "full:panel.test")}
	b.e3JHP2MzU.Store(&detectors)
	u, err := user.Sa_NiNyuJ_ax(&types.Aga91y{WZ89pTAOrHK4: 7, JAfXWwzYrW: &types.IImIwacp{Dh5iiakU: "11111111-1111-4111-8111-111111111111"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		addr    string
		blocked bool
	}{{"white.test:443", false}, {"192.0.2.3:443", false}, {"blocked.test:999", false}, {"blocked.test:80", true}, {"panel.test:80", true}, {"198.51.100.4:80", true}, {"203.0.113.4:443", true}, {"safe.example:80", false}} {
		a, _ := relay.U80FDt(tc.addr)
		if got := b.TestAddr(a, u); got != tc.blocked {
			t.Fatal(tc, got)
		}
	}
	reports := b.iyyXcKcQH7k7()
	if len(reports) != 1 || reports[0].CAgUUdrv9 != 55 || reports[0].GbzQUT_4f != 7 {
		t.Fatal(reports)
	}
	for _, kind := range []sniff.MCe3PxlaMYM{"tls", "http", "quic"} {
		b.cEUo2gYKI = sniff.KZgBMVcpL{&sniffFixture{kind, "Example.TEST"}}
		if host, e := b.Sniff(nil, u, 0); host != "Example.TEST" || e != nil {
			t.Fatal(host, e)
		}
	}
	b.cEUo2gYKI = sniff.KZgBMVcpL{&sniffFixture{"bt", ""}}
	if _, e := b.Sniff(nil, u, 0); e == nil || e.Error() != "user 7 access bt protocol" {
		t.Fatal(e)
	}
	if host, e := b.Sniff(nil, u, 1); host != "" || e != nil {
		t.Fatal("network mismatch should be suppressed", host, e)
	}
}
