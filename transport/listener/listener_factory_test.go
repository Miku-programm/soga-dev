package listener

import (
	"crypto/tls"
	"io"
	"net"
	"os"
	"os/exec"
	P "recovered.local/soga/transport/proxyproto"
	L "recovered.local/soga/panel/types"
	X "recovered.local/soga/transport/xhttp"
	"recovered.local/soga/reality"
	"testing"
	"time"
)

func TestListenerFactoryDefaultAndProxy(t *testing.T) {
	for _, force := range []bool{false, true} {
		options := Ap9uD2{Host: "127.0.0.1", ProxyProtocol: true, ForceProxyProtocol: force}
		listener, e := options.NewListener()
		if e != nil {
			t.Fatal(e)
		}
		p, ok := listener.(*P.KprDQw4aZEOg)
		if !ok || options.Network != "tcp" {
			t.Fatal("default/proxy type")
		}
		if force {
			policy, e := p.Uw20KUu1(nil)
			if policy != 3 || e != nil {
				t.Fatal("original force callback")
			}
		} else if p.Uw20KUu1 != nil {
			t.Fatal("default policy")
		}
		done := make(chan error, 1)
		go func() {
			c, e := listener.Accept()
			if e != nil {
				done <- e
				return
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(3 * time.Second))
			var b [1]byte
			_, e = io.ReadFull(c, b[:])
			if e == nil && (b[0] != 'x' || c.RemoteAddr().String() != "192.0.2.3:1234") {
				e = io.ErrUnexpectedEOF
			}
			done <- e
		}()
		c, e := net.Dial("tcp", listener.Addr().String())
		if e != nil {
			t.Fatal(e)
		}
		_, e = c.Write([]byte("PROXY TCP4 192.0.2.3 192.0.2.4 1234 4321\r\nx"))
		if e != nil {
			t.Fatal(e)
		}
		if e = <-done; e != nil {
			t.Fatal(e)
		}
		c.Close()
		listener.Close()
	}
}
func TestListenerFactoryTransportConstruction(t *testing.T) {
	for _, network := range []string{"tcp", "ws", "h2", "grpc", "xhttp"} {
		t.Run(network, func(t *testing.T) {
			o := Ap9uD2{Host: "127.0.0.1", Network: L.BoulTn31(network), Path: "/fixture", ServiceName: "fixture", XHTTP: &X.R0J19auWof{}}
			listener, e := o.NewListener()
			if e != nil {
				t.Fatal(e)
			}
			if listener == nil {
				t.Fatal("nil listener")
			}
			listener.Close()
		})
	}
	// ws chooses TLS before checking Reality; h2/grpc reject Reality even with TLS.
	o := Ap9uD2{Host: "127.0.0.1", Network: "ws", Path: "/fixture", TLS: &tls.Config{}, Reality: &reality.Config{}}
	l, e := o.NewListener()
	if e != nil {
		t.Fatal(e)
	}
	l.Close()
}

// Keep the original error-branch socket leak inside a subprocess so the test
// process can reclaim its sockets on exit. This checks original behavior,
// rather than changing resource ownership to hide a recovered difference.
func TestListenerFactoryOriginalExplicitErrorsKeepSocket(t *testing.T) {
	if os.Getenv("SS_LISTENER_ERROR_FIXTURE") != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestListenerFactoryOriginalExplicitErrorsKeepSocket$")
		cmd.Env = append(os.Environ(), "SS_LISTENER_ERROR_FIXTURE=1")
		if b, e := cmd.CombinedOutput(); e != nil {
			t.Fatalf("%v %s", e, b)
		}
		return
	}
	for _, tc := range []struct {
		network, want string
		reality       bool
	}{{"invalid", "unknown stream type: <invalid>", false}, {"ws", "reality not support websocket", true}, {"h2", "reality not support h2 or grpc", true}, {"grpc", "reality not support h2 or grpc", true}} {
		probe, e := net.Listen("tcp4", "127.0.0.1:0")
		if e != nil {
			t.Fatal(e)
		}
		port := probe.Addr().(*net.TCPAddr).Port
		addr := probe.Addr().String()
		probe.Close()
		o := Ap9uD2{Host: "127.0.0.1", Port: port, Network: L.BoulTn31(tc.network)}
		if tc.reality {
			o.Reality = &reality.Config{}
		}
		l, e := o.NewListener()
		if l != nil || e == nil || e.Error() != tc.want {
			t.Fatal(l, e)
		}
		c, e := net.DialTimeout("tcp", addr, time.Second)
		if e != nil {
			t.Fatal("original setupError cleanup changed:", e)
		}
		c.Close()
	}
}
