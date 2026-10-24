package MFFa9yT

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/miekg/dns"
	"golang.org/x/crypto/hkdf"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	E "recovered.local/soga/EP3M1uC"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	X "recovered.local/soga/xx5fRyrFw"
	"testing"
	"time"
)

// These fixtures call the original native constructor, Start and UpdateAPIData.
// Clients independently encode SS AES-128-GCM, and targets are real sockets.
// Neither the license clock nor the session gate is overridden.
// The child has only loopback: unchanged hourly authorization workers cannot
// contact a real endpoint even if their random delay expires during the test.
func liveProcess(t *testing.T) bool {
	t.Helper()
	if os.Getenv("SOGA_NATIVE_LIVE_NETNS") == "1" {
		interfaces, err := net.Interfaces()
		if err != nil || len(interfaces) != 1 || interfaces[0].Name != "lo" {
			t.Fatalf("network namespace not isolated: %v %v", interfaces, err)
		}
		return true
	}
	cmd := exec.Command("unshare", "-Urn", "sh", "-c", `/usr/sbin/ip link set lo up && exec "$@"`, "live-fixture", os.Args[0], "-test.run=^"+t.Name()+"$", "-test.timeout=30s", "-test.v")
	cmd.Env = append(os.Environ(), "SOGA_NATIVE_LIVE_NETNS=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("isolated fixture: %v\n%s", err, out)
	}
	t.Log(string(out))
	return false
}
func liveConfiguration(t *testing.T, listen string) {
	t.Helper()
	p := filepath.Join(t.TempDir(), "soga.conf")
	config := "listen=" + listen + "\naccess_log_enable=false\nlog_file_dir=\nban_private_ip=false\nforbidden_bit_torrent=false\nproxy_protocol=false\nudp_proxy_protocol=false\ndy_limit_enable=false\nsoga_key=\n"
	if err := os.WriteFile(p, []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(p)
	if err := C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
}
func liveCore(t *testing.T, node *L.Wr5SVP) (*Lp07FzBLx9V, *Q.D5gTIZ5aVg) {
	t.Helper()
	d := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{ssPanelFixture(7, "live-key")}}
	raw, err := aXSbxcDLdeEY(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: d})
	if err != nil {
		t.Fatal(err)
	}
	v := raw.(*Lp07FzBLx9V)
	t.Cleanup(v.Stop)
	if err = v.Start(); err != nil {
		t.Fatal(err)
	}
	return v, d
}
func liveNode() *L.Wr5SVP {
	return &L.Wr5SVP{EMlXEc: 32, FnMUny: L.AAMPrdcYTL6{HMkc4BgR: "aes-128-gcm"}}
}
func liveUpdate(t *testing.T, v *Lp07FzBLx9V, d *Q.D5gTIZ5aVg) {
	t.Helper()
	if v.UpdateAPIData(d) {
		t.Fatal("unexpected node fingerprint restart")
	}
}
func liveAEAD(password string, salt []byte) cipher.AEAD {
	key := md5.Sum([]byte(password))
	sub := make([]byte, 16)
	if _, err := io.ReadFull(hkdf.New(sha1.New, key[:], salt, []byte("ss-subkey")), sub); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(sub)
	if err != nil {
		panic(err)
	}
	a, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	return a
}

type liveWire struct {
	net.Conn
	password     string
	target       []byte
	tx, rx       cipher.AEAD
	tn, rn       []byte
	udp, started bool
}

func liveClient(t *testing.T, v *Lp07FzBLx9V, udp bool, password string, target []byte) *liveWire {
	t.Helper()
	network, addr := "tcp", v.vLRjbS7NDeZZ[0].Addr().String()
	if udp {
		network, addr = "udp", v.sSqGyaLW[0].LocalAddr().String()
	}
	c, err := net.Dial(network, addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Close() })
	return &liveWire{Conn: c, password: password, target: target, tn: make([]byte, 12), rn: make([]byte, 12), udp: udp}
}
func (c *liveWire) exchange() (string, error) {
	c.SetDeadline(time.Now().Add(350 * time.Millisecond))
	payload := []byte("ping")
	if c.udp || !c.started {
		payload = append(append([]byte(nil), c.target...), payload...)
	}
	var wire []byte
	if c.udp || !c.started {
		salt := make([]byte, 16)
		if _, e := rand.Read(salt); e != nil {
			return "", e
		}
		c.tx = liveAEAD(c.password, salt)
		wire = salt
	}
	if c.udp {
		wire = append(wire, c.tx.Seal(nil, make([]byte, 12), payload, nil)...)
	} else {
		wire = append(wire, ssClientFrame(c.tx, c.tn, payload)...)
	}
	c.started = true
	if _, err := c.Write(wire); err != nil {
		return "", err
	}
	if c.udp {
		buf := make([]byte, 4096)
		n, err := c.Read(buf)
		if err != nil {
			return "", err
		}
		if n < 32 {
			return "", io.ErrUnexpectedEOF
		}
		a := liveAEAD(c.password, buf[:16])
		p, err := a.Open(nil, make([]byte, 12), buf[16:n], nil)
		if err != nil {
			return "", err
		}
		if len(p) < 7 {
			return "", io.ErrUnexpectedEOF
		}
		return string(p[7:]), nil
	}
	if c.rx == nil {
		salt := make([]byte, 16)
		if _, err := io.ReadFull(c, salt); err != nil {
			return "", err
		}
		c.rx = liveAEAD(c.password, salt)
	}
	var reply []byte
	for len(reply) < 5 {
		p, err := ssReadClientFrame(c, c.rx, c.rn)
		if err != nil {
			return "", err
		}
		reply = append(reply, p...)
	}
	return string(reply), nil
}
func liveExpect(t *testing.T, c *liveWire, want string) {
	t.Helper()
	got, err := c.exchange()
	if err != nil || got != want {
		t.Fatalf("wire got %q, %v; want %q", got, err, want)
	}
}
func liveReject(t *testing.T, c *liveWire) {
	t.Helper()
	got, err := c.exchange()
	if err == nil {
		t.Fatalf("unexpected accepted packet %q", got)
	}
}
func liveTarget(t *testing.T, host string, port int, tag byte) int {
	t.Helper()
	l, err := net.Listen("tcp4", net.JoinHostPort(host, fmt.Sprint(port)))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { l.Close() })
	port = l.Addr().(*net.TCPAddr).Port
	pc, err := net.ListenPacket("udp4", net.JoinHostPort(host, fmt.Sprint(port)))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pc.Close() })
	go func() {
		for {
			c, e := l.Accept()
			if e != nil {
				return
			}
			go func() {
				defer c.Close()
				p := make([]byte, 4)
				for {
					c.SetDeadline(time.Now().Add(10 * time.Second))
					if _, e := io.ReadFull(c, p); e != nil {
						return
					}
					if _, e := c.Write(append([]byte{tag}, p...)); e != nil {
						return
					}
				}
			}()
		}
	}()
	go func() {
		p := make([]byte, 4096)
		for {
			n, a, e := pc.ReadFrom(p)
			if e != nil {
				return
			}
			pc.WriteTo(append([]byte{tag}, p[:n]...), a)
		}
	}()
	return port
}
func TestNativeLiveBlockWhiteAndExistingSessions(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	liveConfiguration(t, "127.0.0.1")
	v, d := liveCore(t, liveNode())
	port := liveTarget(t, "127.0.0.1", 0, 'A')
	target := ssFixtureAddress(1, port)
	tcp := liveClient(t, v, false, "live-key", target)
	udp := liveClient(t, v, true, "live-key", target)
	liveExpect(t, tcp, "Aping")
	liveExpect(t, udp, "Aping")
	d.ATGYGi0 = X.VHWq49iRat{"ip:127.0.0.1"}
	liveUpdate(t, v, d)
	liveExpect(t, tcp, "Aping")
	liveReject(t, liveClient(t, v, false, "live-key", target))
	liveReject(t, udp)
	d.NjjWbjH = X.VHWq49iRat{"ip:127.0.0.1"}
	liveUpdate(t, v, d)
	liveExpect(t, tcp, "Aping")
	liveExpect(t, udp, "Aping")
	liveExpect(t, liveClient(t, v, false, "live-key", target), "Aping")
	// Failed rule retrieval keeps the installed white rule rather than replacing it.
	d.NjjWbjH = nil
	d.IGP6TVQazVxp = errors.New("fixture retrieval failure")
	liveUpdate(t, v, d)
	liveExpect(t, udp, "Aping")
	liveExpect(t, liveClient(t, v, false, "live-key", target), "Aping")
	d.IGP6TVQazVxp = nil
	liveUpdate(t, v, d)
	liveReject(t, udp)
	liveExpect(t, tcp, "Aping")
	d.ATGYGi0 = nil
	d.BeqZvr6 = errors.New("fixture block retrieval failure")
	liveUpdate(t, v, d)
	liveReject(t, udp)
	d.BeqZvr6 = nil
	liveUpdate(t, v, d)
	liveExpect(t, udp, "Aping")
}
func TestNativeLiveTCPLimitAndCredentialCancellation(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	liveConfiguration(t, "127.0.0.1")
	v, d := liveCore(t, liveNode())
	target := ssFixtureAddress(1, liveTarget(t, "127.0.0.1", 0, 'A'))
	one := liveClient(t, v, false, "live-key", target)
	two := liveClient(t, v, false, "live-key", target)
	liveExpect(t, one, "Aping")
	liveExpect(t, two, "Aping")
	updated := ssPanelFixture(7, "live-key")
	updated.B4DYSxuA31za = 1
	d.VTme1SuHL6i_ = L.Sxcf3wu3L{updated}
	liveUpdate(t, v, d)
	liveExpect(t, one, "Aping")
	liveExpect(t, two, "Aping")
	liveReject(t, liveClient(t, v, false, "live-key", target))
	// Relaxing the cap permits a new connection without replacing credentials.
	updated = ssPanelFixture(7, "live-key")
	d.VTme1SuHL6i_ = L.Sxcf3wu3L{updated}
	liveUpdate(t, v, d)
	liveExpect(t, liveClient(t, v, false, "live-key", target), "Aping")
	updated = ssPanelFixture(7, "changed-live-key")
	d.VTme1SuHL6i_ = L.Sxcf3wu3L{updated}
	liveUpdate(t, v, d)
	liveReject(t, one)
	liveReject(t, two)
	liveExpect(t, liveClient(t, v, false, "changed-live-key", target), "Aping")
	liveReject(t, liveClient(t, v, false, "live-key", target))
}
func liveRoute(t *testing.T, port int) {
	t.Helper()
	r, err := E.KmKjIZi1(fmt.Sprintf("health_check=false\n[[routes]]\nrules=[\"node_id:32\"]\n[[routes.outs]]\nname=\"live-redirect\"\ntype=\"redirect\"\nserver=\"127.0.0.1\"\nport=%d\n", port))
	if err != nil {
		t.Fatal(err)
	}
	*E.CurrentRoutes() = *r
}
func TestNativeLiveRoutesKeepEstablishedOutbounds(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	liveConfiguration(t, "127.0.0.1")
	old := *E.CurrentRoutes()
	defer func() { *E.CurrentRoutes() = old }()
	a, b := liveTarget(t, "127.0.0.1", 0, 'A'), liveTarget(t, "127.0.0.1", 0, 'B')
	liveRoute(t, a)
	v, _ := liveCore(t, liveNode())
	target := ssFixtureAddress(1, 9)
	tcp, udp := liveClient(t, v, false, "live-key", target), liveClient(t, v, true, "live-key", target)
	liveExpect(t, tcp, "Aping")
	liveExpect(t, udp, "Aping")
	nat, ok := v.hS8xXFfcBr.Get(udp.LocalAddr().String())
	if !ok {
		t.Fatal("NAT absent after established reply")
	}
	liveRoute(t, b)
	liveExpect(t, tcp, "Aping")
	liveExpect(t, udp, "Aping")
	if after, ok := v.hS8xXFfcBr.Get(udp.LocalAddr().String()); !ok || after != nat {
		t.Fatal("route update replaced existing NAT")
	}
	liveExpect(t, liveClient(t, v, false, "live-key", target), "Bping")
	liveExpect(t, liveClient(t, v, true, "live-key", target), "Bping")
}
func liveDNSServer(t *testing.T, ip string) string {
	t.Helper()
	pc, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	s := &dns.Server{PacketConn: pc, Handler: dns.HandlerFunc(func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		for _, q := range r.Question {
			if q.Qtype == dns.TypeA {
				m.Answer = append(m.Answer, &dns.A{Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60}, A: net.ParseIP(ip)})
			}
		}
		w.WriteMsg(m)
	})}
	go s.ActivateAndServe()
	t.Cleanup(func() { s.Shutdown(); pc.Close() })
	return "udp://" + pc.LocalAddr().String()
}
func TestNativeLiveResolverChangesExistingUDPDestination(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	liveConfiguration(t, "127.0.0.1")
	port := liveTarget(t, "127.0.0.1", 0, 'A')
	liveTarget(t, "127.0.0.2", port, 'B')
	a, b := liveDNSServer(t, "127.0.0.1"), liveDNSServer(t, "127.0.0.2")
	node := liveNode()
	node.AjXcBfRcAXR = L.U863Mm01Nre{&L.Axa8uKIB4JL{K8BNfziIgb: []string{a}, CM0eMsqyP: "ipv4_only", CZ80ik: []string{"full:live.test"}}}
	v, d := liveCore(t, node)
	target := binary.BigEndian.AppendUint16(append([]byte{3, 9}, []byte("live.test")...), uint16(port))
	tcp, udp := liveClient(t, v, false, "live-key", target), liveClient(t, v, true, "live-key", target)
	liveExpect(t, tcp, "Aping")
	liveExpect(t, udp, "Aping")
	nat, ok := v.hS8xXFfcBr.Get(udp.LocalAddr().String())
	if !ok {
		t.Fatal("NAT absent after established reply")
	}
	prior := v.ffju0Enjo.Load()
	next := *node
	next.AjXcBfRcAXR = L.U863Mm01Nre{&L.Axa8uKIB4JL{K8BNfziIgb: []string{b}, CM0eMsqyP: "ipv4_only", CZ80ik: []string{"full:live.test"}}}
	d.H47AzZsYHB1N = &next
	liveUpdate(t, v, d)
	if v.ffju0Enjo.Load() == prior {
		t.Fatal("merged rules must replace resolver")
	}
	liveExpect(t, tcp, "Aping")
	liveExpect(t, udp, "Bping")
	if after, ok := v.hS8xXFfcBr.Get(udp.LocalAddr().String()); !ok || after != nat {
		t.Fatal("resolver update replaced existing NAT")
	}
	liveExpect(t, liveClient(t, v, false, "live-key", target), "Bping")
}
func TestNativeMultiListenFailureRollbackAndRetry(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	liveConfiguration(t, "127.0.0.1,127.0.0.2")
	for _, network := range []string{"tcp", "udp"} {
		t.Run(network, func(t *testing.T) {
			var blocker io.Closer
			var port int
			if network == "tcp" {
				l, e := net.Listen("tcp4", "127.0.0.2:0")
				if e != nil {
					t.Fatal(e)
				}
				blocker = l
				port = l.Addr().(*net.TCPAddr).Port
			} else {
				p, e := net.ListenPacket("udp4", "127.0.0.2:0")
				if e != nil {
					t.Fatal(e)
				}
				blocker = p
				port = p.LocalAddr().(*net.UDPAddr).Port
			}
			defer blocker.Close()
			v := ssStartFixture(t, port)
			if e := v.Start(); e == nil {
				t.Fatal("second address must fail")
			}
			if len(v.vLRjbS7NDeZZ) != 0 || len(v.sSqGyaLW) != 0 {
				t.Fatal("partial listeners published")
			}
			for _, host := range []string{"127.0.0.1", "127.0.0.2"} {
				if host == "127.0.0.1" || network != "tcp" {
					l, e := net.Listen("tcp4", net.JoinHostPort(host, fmt.Sprint(port)))
					if e != nil {
						t.Fatal("leaked TCP", e)
					}
					l.Close()
				}
				if host == "127.0.0.1" {
					p, e := net.ListenPacket("udp4", net.JoinHostPort(host, fmt.Sprint(port)))
					if e != nil {
						t.Fatal("leaked UDP", e)
					}
					p.Close()
				}
			}
			blocker.Close()
			if e := v.Start(); e != nil {
				t.Fatal("retry after rollback", e)
			}
			if len(v.vLRjbS7NDeZZ) != 2 || len(v.sSqGyaLW) != 2 {
				t.Fatal("retry not fully installed")
			}
			target := ssFixtureAddress(1, liveTarget(t, "127.0.0.1", 0, 'A'))
			liveExpect(t, liveClient(t, v, false, "udp-native-start", target), "Aping")
			liveExpect(t, liveClient(t, v, true, "udp-native-start", target), "Aping")
		})
	}
}
