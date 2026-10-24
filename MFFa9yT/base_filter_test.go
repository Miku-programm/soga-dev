package MFFa9yT

import (
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/qwn0vjA"
	"recovered.local/soga/xx5fRyrFw"
	"testing"
)

type sniffFixture struct {
	kind qwn0vjA.MCe3PxlaMYM
	host string
}

func (s *sniffFixture) Type() qwn0vjA.MCe3PxlaMYM                 { return s.kind }
func (s *sniffFixture) Value() string                             { return s.host }
func (s *sniffFixture) Network() G6bHD5JA.CjhtWb                  { return 0 }
func (s *sniffFixture) Sniff([]byte) (qwn0vjA.FlCuFjj_V1u, error) { return s, nil }
func TestAddressRulesWhitelistPriorityAndReport(t *testing.T) {
	b := &JWN154qJ4{qE9aU9: &G7A7csgrs62p.GTFAWF[int, *G7A7csgrs62p.Fxs7SY[int64]]{}}
	white, _, _, _ := elCwduDdyGl(xx5fRyrFw.VHWq49iRat{"full:white.test", "ip:192.0.2.3", "port:999"})
	b._IoxtH.Store(white)
	black, _, _, _ := elCwduDdyGl(xx5fRyrFw.VHWq49iRat{"domain:test", "ip:198.51.100.0/24", "port:443"})
	b.oIgEuicmPL.Store(black)
	detectors := xx5fRyrFw.Ij9r21E0UOt3{xx5fRyrFw.NewPanelDetector(55, "full:panel.test")}
	b.e3JHP2MzU.Store(&detectors)
	u, err := nzATvajt.Sa_NiNyuJ_ax(&LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 7, JAfXWwzYrW: &LFWTTXF7JR.IImIwacp{Dh5iiakU: "11111111-1111-4111-8111-111111111111"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		addr    string
		blocked bool
	}{{"white.test:443", false}, {"192.0.2.3:443", false}, {"blocked.test:999", false}, {"blocked.test:80", true}, {"panel.test:80", true}, {"198.51.100.4:80", true}, {"203.0.113.4:443", true}, {"safe.example:80", false}} {
		a, _ := G6bHD5JA.U80FDt(tc.addr)
		if got := b.TestAddr(a, u); got != tc.blocked {
			t.Fatal(tc, got)
		}
	}
	reports := b.iyyXcKcQH7k7()
	if len(reports) != 1 || reports[0].CAgUUdrv9 != 55 || reports[0].GbzQUT_4f != 7 {
		t.Fatal(reports)
	}
	for _, kind := range []qwn0vjA.MCe3PxlaMYM{"tls", "http", "quic"} {
		b.cEUo2gYKI = qwn0vjA.KZgBMVcpL{&sniffFixture{kind, "Example.TEST"}}
		if host, e := b.Sniff(nil, u, 0); host != "Example.TEST" || e != nil {
			t.Fatal(host, e)
		}
	}
	b.cEUo2gYKI = qwn0vjA.KZgBMVcpL{&sniffFixture{"bt", ""}}
	if _, e := b.Sniff(nil, u, 0); e == nil || e.Error() != "user 7 access bt protocol" {
		t.Fatal(e)
	}
	if host, e := b.Sniff(nil, u, 1); host != "" || e != nil {
		t.Fatal("network mismatch should be suppressed", host, e)
	}
}
