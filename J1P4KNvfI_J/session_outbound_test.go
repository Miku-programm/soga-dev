package J1P4KNvfI_J

import (
	"net"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/nzATvajt"
	"testing"
)

func TestOriginalOutboundAddressTypesAndSharedSlices(t *testing.T) {
	s := &ZzkV7_5l{}
	local := &net.UDPAddr{IP: net.IP{192, 0, 2, 1}, Port: 100}
	remote := &net.TCPAddr{IP: net.IP{198, 51, 100, 1}, Port: 200}
	s.SetOutboundLocal(local)
	s.SetOutboundRemote(remote)
	local.IP[3] = 9
	remote.IP[3] = 8
	if s.R61M7ZTr.String() != "192.0.2.9" || s.IksfOg.String() != "198.51.100.8" || s.N15sqBj36 != 100 || s.IqEsiK != 200 {
		t.Fatal(s)
	}
	s.SetOutboundLocal(nil)
	s.SetOutboundRemote(&net.UnixAddr{Name: "/tmp/example", Net: "unix"})
	if s.R61M7ZTr.String() != "192.0.2.9" || s.IksfOg.String() != "198.51.100.8" {
		t.Fatal("unknown addresses changed state")
	}
}

func TestOriginalSessionString(t *testing.T) {
	u, e := nzATvajt.Sa_NiNyuJ_ax(&LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 42, JAfXWwzYrW: &LFWTTXF7JR.IImIwacp{Dh5iiakU: "11111111-1111-4111-8111-111111111111"}}, nil)
	if e != nil {
		t.Fatal(e)
	}
	dest, _ := G6bHD5JA.U80FDt("example.test:443")
	s := &ZzkV7_5l{E3L0l8hZ: 9, PayAlbdv6: u, Iayb4V: dest, KeHzUwuXWh0: net.IP{192, 0, 2, 1}, Y72kDV35: 123}
	want := "session id: 9, user: 42, dest addr: example.test:443, client(192.0.2.1:123) -> local(<nil>:0) -> remote_local(<nil>:0) -> remote(<nil>:0)"
	if got := s.String(); got != want {
		t.Fatalf("%s", got)
	}
}
