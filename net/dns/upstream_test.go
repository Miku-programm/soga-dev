package dns

import (
	"context"
	"github.com/miekg/dns"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRecoveredServerConstruction(t *testing.T) {
	for _, input := range []string{"local", "localhost", "local-custom"} {
		if _, ok := cDpHJRFxZ(input, nil).(*BkMuy3LHBq); !ok {
			t.Fatalf("local prefix %s", input)
		}
	}
	udp := cDpHJRFxZ("udp://192.0.2.1", nil).(*Zq5YXn)
	if udp.dqrvhAOMdU6 != "192.0.2.1:53" || udp.gRxcpc77sP.UDPSize != 1024 || udp.gRxcpc77sP.Timeout != 5*time.Second || udp.gRxcpc77sP.Net != "" {
		t.Fatal("udp client fields")
	}
	tcp := cDpHJRFxZ("tcp://[2001:db8::1]:5300", nil).(*KZFqTq8jd)
	if tcp.dqrvhAOMdU6 != "[2001:db8::1]:5300" || tcp.gRxcpc77sP.Net != "tcp" {
		t.Fatal("tcp client fields")
	}
	tls := cDpHJRFxZ("tls://2001:db8::1", nil).(*OPLAviioh9T)
	if tls.dqrvhAOMdU6 != "[2001:db8::1]:853" || tls.gRxcpc77sP.Net != "tcp-tls" {
		t.Fatal("tls client fields")
	}
	unknown := cDpHJRFxZ("unknown://192.0.2.2", nil).(*Zq5YXn)
	if unknown.dqrvhAOMdU6 != "192.0.2.2:53" {
		t.Fatal("unknown scheme fallback must use address suffix")
	}
	v4 := net.IPv4(192, 0, 2, 10)
	if dsVT1Buuy(v4, false, "2001:db8::1") != dialerAt5029CB0 {
		t.Fatal("mixed families must use shared dialer")
	}
	bound := dsVT1Buuy(v4, true, "resolver.test")
	if bound == dialerAt5029CB0 || !bound.LocalAddr.(*net.UDPAddr).IP.Equal(v4) {
		t.Fatal("bound udp dialer")
	}
}
func TestUDPAndTCPWireResolution(t *testing.T) {
	for _, network := range []string{"udp", "tcp"} {
		t.Run(network, func(t *testing.T) {
			server := &dns.Server{Handler: dns.HandlerFunc(func(w dns.ResponseWriter, q *dns.Msg) {
				reply := new(dns.Msg)
				reply.SetReply(q)
				reply.Answer = []dns.RR{&dns.A{Hdr: dns.RR_Header{Name: q.Question[0].Name, Rrtype: 1, Class: 1}, A: net.IPv4(192, 0, 2, 50)}, &dns.AAAA{Hdr: dns.RR_Header{Name: q.Question[0].Name, Rrtype: 28, Class: 1}, AAAA: net.ParseIP("2001:db8::50")}}
				if err := w.WriteMsg(reply); err != nil {
					t.Error(err)
				}
			})}
			var address string
			if network == "udp" {
				pc, err := net.ListenPacket("udp4", "127.0.0.1:0")
				if err != nil {
					t.Fatal(err)
				}
				server.PacketConn = pc
				address = pc.LocalAddr().String()
			} else {
				ln, err := net.Listen("tcp4", "127.0.0.1:0")
				if err != nil {
					t.Fatal(err)
				}
				server.Listener = ln
				address = ln.Addr().String()
			}
			started := make(chan struct{})
			server.NotifyStartedFunc = func() { close(started) }
			go func() {
				if err := server.ActivateAndServe(); err != nil {
					t.Error(err)
				}
			}()
			<-started
			t.Cleanup(func() { _ = server.Shutdown() })
			resolver := cDpHJRFxZ(network+"://"+address, nil)
			for _, typ := range []uint16{1, 28} {
				ips, err := resolver.LookupDNSType(context.Background(), "wire.test", typ)
				if err != nil || len(ips) != 1 {
					t.Fatalf("type %d: %v %v", typ, ips, err)
				}
				if (ips[0].To4() != nil) != (typ == 1) {
					t.Fatal("incorrect answer family")
				}
			}
		})
	}
}
func TestDoHHeadersStatusAndDNSCodeBehavior(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.Header.Get("Accept") != "application/dns-message" || r.Header.Get("Content-Type") != "application/dns-message" {
			t.Error("unexpected DNS HTTP request")
		}
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		var q dns.Msg
		if err = q.Unpack(payload); err != nil {
			t.Error(err)
			return
		}
		if q.Question[0].Name != "wire.test." {
			t.Error("domain not qualified")
		}
		response := new(dns.Msg)
		response.SetReply(&q)
		response.Rcode = dns.RcodeServerFailure
		response.Answer = []dns.RR{&dns.A{Hdr: dns.RR_Header{Name: "wire.test.", Rrtype: 1, Class: 1}, A: net.IPv4(192, 0, 2, 51)}}
		payload, err = response.Pack()
		if err != nil {
			t.Error(err)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write(payload)
	}))
	defer server.Close()
	resolver := eXmNUd9(server.URL, nil)
	ips, err := resolver.LookupDNSType(context.Background(), "wire.test", 1)
	if err != nil || len(ips) != 1 || ips[0].String() != "192.0.2.51" {
		t.Fatalf("original DoH accepts valid answer despite HTTP/DNS status: %v %v", ips, err)
	}
}
