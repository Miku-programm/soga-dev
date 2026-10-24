package MFFa9yT

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	CA "recovered.local/soga/CAodbhZf"
	E "recovered.local/soga/EP3M1uC"
	L "recovered.local/soga/LFWTTXF7JR"
	F "recovered.local/soga/hfOuV8fOe4L"
	U "recovered.local/soga/nzATvajt"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"strconv"
	"strings"
	"testing"
	"time"
)

// The fixture uses two actual native inbound cores: entrance node 32 and proxy
// peer node 33. A real TOML node_id rule sends only entrance sessions through
// the original LCe adapter; peer sessions use a separately configured direct node rule.
// No fake adapter or production test hook is involved.
func TestNativeSessionConfiguredSSAndSSRProxyRoutes(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, kind := range []string{"ss", "ssr"} {
		for _, network := range []string{"tcp", "udp"} {
			t.Run(kind+"/"+network, func(t *testing.T) { nativeProxyRouteFixture(t, kind, network, "origin", "plain") })
		}
	}
}

func nativeRouteSSCore(t *testing.T, nodeID int, password string) *Lp07FzBLx9V {
	t.Helper()
	node := &L.Wr5SVP{EMlXEc: nodeID, FnMUny: L.AAMPrdcYTL6{HMkc4BgR: "aes-128-gcm"}}
	core, err := aXSbxcDLdeEY(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{ssPanelFixture(7, password)}}})
	if err != nil {
		t.Fatal(err)
	}
	v := core.(*Lp07FzBLx9V)
	v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0())) // Original hourly no-request state; retain all checks.
	t.Cleanup(v.Stop)
	if err = v.Start(); err != nil {
		t.Fatal(err)
	}
	return v
}

func nativeProxyRouteFixture(t *testing.T, kind, network, protocol, obfs string) {
	t.Helper()
	entry := nativeRouteSSCore(t, 32, "route-entrance-key")
	var peerAddr net.Addr
	var peerUser U.PgaNd3014U
	var peerChecked <-chan error
	cipher := "aes-128-gcm"
	if kind == "ss" {
		peer := nativeRouteSSCore(t, 33, "route-peer-key")
		peerAddr = peer.vLRjbS7NDeZZ[0].Addr()
		if network == "udp" {
			peerAddr = peer.sSqGyaLW[0].LocalAddr()
		}
		peerUser, _ = peer.jc_QPN6.Get(7)
	} else if kind == "ssr" {
		cipher = "aes-128-cfb"
		node := &L.Wr5SVP{EMlXEc: 33}
		panelUser := ssrPanelFixture(0, cipher, "route-peer-key")
		panelUser.JAfXWwzYrW.(*L.QcqpSL3i6ri).H_y2xZht7S8 = F.Mu0vMIIs(protocol)
		panelUser.JAfXWwzYrW.(*L.QcqpSL3i6ri).ZZIDsJ = F.PMorzD7u(obfs)
		raw, err := x4oEZV(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{panelUser}}})
		if err != nil {
			t.Fatal(err)
		}
		peer := raw.(*Ze0fOaAko)
		peer.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
		t.Cleanup(peer.Stop)
		if err = peer.Start(); err != nil {
			t.Fatal(err)
		}
		state := peer.mc0MXUU.sSKJ79jFt[0]
		peerAddr = state.woqAbWjQH_A[0].Addr()
		if network == "udp" {
			peerAddr = state.aaBpAQ44Zfs[0].LocalAddr()
		}
		peerUser, _ = peer.jc_QPN6.Get(42)
	} else if kind == "vmess" {
		cipher = "aes-128-gcm"
		node := &L.Wr5SVP{EMlXEc: 33, SwxJNgPpg: L.JW0CeHo3O{F2GXyZZ: "tcp"}}
		core, e := ag90n_rZjJO(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{vmessPanelFixture(7, "11111111-1111-4111-8111-111111111111")}}})
		if e != nil {
			t.Fatal(e)
		}
		peer := core.(*K4avFiYOPi)
		t.Cleanup(peer.Stop)
		peer.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
		if e = peer.Start(); e != nil {
			t.Fatal(e)
		}
		peerAddr = peer.l8Uq4M[0].Addr()
		peerUser, _ = peer.jc_QPN6.Get(7)
	} else if kind == "trojan" {
		cert, key := hysteriaFixtureCertificate(t)
		node := &L.Wr5SVP{EMlXEc: 33, I3k4JnNp7nB: cert, XATiVun: key, JhIYwEukjd: L.L9AiYC{T4cuhG7MTbG: "tcp", AJjaqrBjj: "tls"}}
		core, e := f1mPzegGt(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{trojanPanelFixture(9, "route-peer-key")}}})
		if e != nil {
			t.Fatal(e)
		}
		peer := core.(*EfaIFpp)
		t.Cleanup(peer.Stop)
		peer.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
		if e = peer.Start(); e != nil {
			t.Fatal(e)
		}
		peerAddr = peer.cDyEMW2[0].Addr()
		peerUser, _ = peer.jc_QPN6.Get(9)
	} else {
		peerAddr, peerChecked = nativeIndependentProxyPeer(t, kind, network)
	}
	_, peerPort, _ := net.SplitHostPort(peerAddr.String())
	routeText := fmt.Sprintf(`health_check=false
[[routes]]
rules=["node_id:32"]
[[routes.outs]]
name="native-proxy-peer"
type=%q
server="127.0.0.1"
port=%s
password="route-peer-key"
cipher=%q
uuid="11111111-1111-4111-8111-111111111111"
username="route-user"
skip_cert_verify=true
obfs=%q
protocol=%q
[[routes]]
rules=["node_id:33"]
[[routes.outs]]
type="direct"
`, kind, peerPort, cipher, obfs, protocol)
	routes, err := E.KmKjIZi1(routeText)
	if err != nil {
		t.Fatal(err)
	}
	active := E.CurrentRoutes()
	old := *active
	*active = *routes
	defer func() { *active = old }()
	payload := bytes.Repeat([]byte("two-hop original route\x00"), 97)
	response := bytes.Repeat([]byte("final target response\x00"), 83)
	done := make(chan error, 1)
	var targetPort int
	if network == "tcp" {
		target, e := net.Listen("tcp4", "127.0.0.1:0")
		if e != nil {
			t.Fatal(e)
		}
		defer target.Close()
		targetPort = target.Addr().(*net.TCPAddr).Port
		go func() {
			conn, e := target.Accept()
			if e == nil {
				defer conn.Close()
				conn.SetDeadline(time.Now().Add(5 * time.Second))
				p := make([]byte, len(payload))
				_, e = io.ReadFull(conn, p)
				if e == nil && !bytes.Equal(p, payload) {
					e = fmt.Errorf("TCP target payload %x", p)
				}
				if e == nil {
					_, e = conn.Write(response)
				}
			}
			done <- e
		}()
	} else {
		target, e := net.ListenPacket("udp4", "127.0.0.1:0")
		if e != nil {
			t.Fatal(e)
		}
		defer target.Close()
		target.SetDeadline(time.Now().Add(5 * time.Second))
		targetPort = target.LocalAddr().(*net.UDPAddr).Port
		go func() {
			p := make([]byte, 4096)
			n, a, e := target.ReadFrom(p)
			if e == nil && !bytes.Equal(p[:n], payload) {
				e = fmt.Errorf("UDP target payload %x", p[:n])
			}
			if e == nil {
				_, e = target.WriteTo(response, a)
			}
			done <- e
		}()
	}
	entryAddr := entry.vLRjbS7NDeZZ[0].Addr()
	if network == "udp" {
		entryAddr = entry.sSqGyaLW[0].LocalAddr()
	}
	client, err := net.Dial(network, entryAddr.String())
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	client.SetDeadline(time.Now().Add(6 * time.Second))
	rawUser, _ := entry.jc_QPN6.Get(7)
	user := rawUser.(*U.WtK9HQSy66)
	salt := make([]byte, 16)
	rand.Read(salt)
	a, err := user.NewAEAD(salt)
	if err != nil {
		t.Fatal(err)
	}
	address := ssFixtureAddress(1, targetPort)
	plain := append(bytes.Clone(address), payload...)
	wire := bytes.Clone(salt)
	if network == "tcp" {
		wire = append(wire, ssClientFrame(a, make([]byte, 12), plain)...)
	} else {
		wire = append(wire, a.Seal(nil, make([]byte, 12), plain, nil)...)
	}
	if _, err = client.Write(wire); err != nil {
		t.Fatal(err)
	}
	var received []byte
	if network == "tcp" {
		if _, err = io.ReadFull(client, salt); err != nil {
			t.Fatal(err)
		}
		a, err = user.NewAEAD(salt)
		if err != nil {
			t.Fatal(err)
		}
		nonce := make([]byte, 12)
		for len(received) < len(response) {
			p, e := ssReadClientFrame(client, a, nonce)
			if e != nil {
				t.Fatal(e)
			}
			received = append(received, p...)
		}
	} else {
		p := make([]byte, 4096)
		n, e := client.Read(p)
		if e != nil {
			t.Fatal(e)
		}
		if n < 32 {
			t.Fatal("short reply", n)
		}
		a, err = user.NewAEAD(p[:16])
		if err != nil {
			t.Fatal(err)
		}
		received, err = a.Open(nil, make([]byte, 12), p[16:n], nil)
		if err != nil {
			t.Fatal(err)
		}
		response = append(bytes.Clone(address), response...)
	}
	if !bytes.Equal(received, response) {
		t.Fatalf("reply=%x want=%x", received, response)
	}
	if err = <-done; err != nil {
		t.Fatal(err)
	}
	// A direct bypass would reach the target but leave the real peer user's
	// accounting empty. Positive peer traffic proves the configured second hop.
	// Original SSR mode0 UDP skips upload accounting in ReadFromProxy; its
	// producer only substitutes the port user, so require download traffic there.
	if peerUser == nil {
		select {
		case e := <-peerChecked:
			if e != nil {
				t.Fatal(e)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("independent proxy did not confirm actual destination dial")
		}
		return
	}
	up, down := peerUser.GetAndClearTraffic()
	if down <= 0 || (up <= 0 && !(kind == "ssr" && network == "udp")) {
		t.Fatalf("proxy peer bypassed or unaccounted: up=%d down=%d", up, down)
	}
}

// The configured outbound and original inbound have separate implementations;
// this matrix catches shared-reconstruction assumptions that a same-code
// roundtrip cannot. UDP ignores obfs in the original, so auth variants run
// both transports, while obfs-specific cases use TCP.
func TestNativeSSRProxyAuthenticationAndObfsRoutes(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, protocol := range []string{"auth_aes128_md5", "auth_aes128_sha1", "auth_chain_a", "auth_chain_b"} {
		for _, network := range []string{"tcp", "udp"} {
			t.Run(protocol+"/"+network, func(t *testing.T) { nativeProxyRouteFixture(t, "ssr", network, protocol, "plain") })
		}
	}
	for _, obfs := range []string{"http_simple", "http_post", "tls1.2_ticket_auth"} {
		t.Run(obfs, func(t *testing.T) { nativeProxyRouteFixture(t, "ssr", "tcp", "origin", obfs) })
	}
}

// Native SSR server protocols are not the same registry as SSR outbound
// protocols. Original low0s init registers only chain A/B, whereas Roa accepts
// A-F. Preserve the route factory rejection instead of expanding the kernel.
func TestNativeSSRRouteRejectsServerOnlyChainVariants(t *testing.T) {
	for _, protocol := range []string{"auth_chain_c", "auth_chain_d", "auth_chain_e", "auth_chain_f"} {
		_, err := E.KmKjIZi1(fmt.Sprintf(`health_check=false
[[routes]]
rules=["node_id:32"]
[[routes.outs]]
type="ssr"
server="127.0.0.1"
port=443
password="fixture"
cipher="aes-128-cfb"
obfs="plain"
protocol=%q
`, protocol))
		if err == nil || !strings.Contains(err.Error(), "protocol "+protocol+" not supported") {
			t.Fatal(protocol, err)
		}
	}
}

func TestNativeConfiguredHTTPAndSOCKSAndVMessAndTrojanRoutes(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, kind := range []string{"http", "socks", "vmess", "trojan"} {
		networks := []string{"tcp", "udp"}
		if kind == "http" {
			networks = networks[:1]
		}
		for _, network := range networks {
			t.Run(kind+"/"+network, func(t *testing.T) { nativeProxyRouteFixture(t, kind, network, "origin", "plain") })
		}
	}
}

// This peer is test-only and parses the public wire independently. It actually
// dials the requested loopback destination; it never returns a canned echo.
func nativeIndependentProxyPeer(t *testing.T, kind, network string) (net.Addr, <-chan error) {
	t.Helper()
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { listener.Close() })
	checked := make(chan error, 1)
	go func() {
		conn, e := listener.Accept()
		if e != nil {
			checked <- e
			return
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		reader := bufio.NewReader(conn)
		var target string
		if kind == "http" {
			request, err := http.ReadRequest(reader)
			if err != nil {
				checked <- err
				return
			}
			if request.Method != "CONNECT" || request.Header.Get("Proxy-Authorization") != "Basic "+base64.StdEncoding.EncodeToString([]byte("route-user:route-peer-key")) {
				checked <- fmt.Errorf("unexpected CONNECT %s auth=%q", request.Method, request.Header.Get("Proxy-Authorization"))
				return
			}
			target = request.Host
		} else {
			header := make([]byte, 2)
			if _, e = io.ReadFull(reader, header); e != nil {
				checked <- e
				return
			}
			if header[0] != 5 || header[1] == 0 {
				checked <- fmt.Errorf("bad SOCKS greeting %x", header)
				return
			}
			methods := make([]byte, int(header[1]))
			if _, e = io.ReadFull(reader, methods); e != nil {
				checked <- e
				return
			}
			if !bytes.Contains(methods, []byte{2}) {
				checked <- fmt.Errorf("authentication method missing %x", methods)
				return
			}
			if _, e = conn.Write([]byte{5, 2}); e != nil {
				checked <- e
				return
			}
			if _, e = io.ReadFull(reader, header); e != nil {
				checked <- e
				return
			}
			if header[0] != 1 {
				checked <- fmt.Errorf("auth version %d", header[0])
				return
			}
			username := make([]byte, int(header[1]))
			if _, e = io.ReadFull(reader, username); e != nil {
				checked <- e
				return
			}
			n, e := reader.ReadByte()
			if e != nil {
				checked <- e
				return
			}
			password := make([]byte, int(n))
			if _, e = io.ReadFull(reader, password); e != nil {
				checked <- e
				return
			}
			if string(username) != "route-user" || string(password) != "route-peer-key" {
				checked <- fmt.Errorf("SOCKS credentials mismatch")
				return
			}
			if _, e = conn.Write([]byte{1, 0}); e != nil {
				checked <- e
				return
			}
			command := make([]byte, 3)
			if _, e = io.ReadFull(reader, command); e != nil {
				checked <- e
				return
			}
			expected := byte(1)
			if network == "udp" {
				expected = 3
			}
			if !bytes.Equal(command, []byte{5, expected, 0}) {
				checked <- fmt.Errorf("SOCKS request %x", command)
				return
			}
			target, _, e = readIndependentSocksAddress(reader)
			if e != nil {
				checked <- e
				return
			}
			if network == "udp" {
				pc, err := net.ListenPacket("udp4", "127.0.0.1:0")
				if err != nil {
					checked <- err
					return
				}
				defer pc.Close()
				pc.SetDeadline(time.Now().Add(5 * time.Second))
				bound := binary.BigEndian.AppendUint16([]byte{5, 0, 0, 1, 127, 0, 0, 1}, uint16(pc.LocalAddr().(*net.UDPAddr).Port))
				if _, e = conn.Write(bound); e != nil {
					checked <- e
					return
				}
				buf := make([]byte, 4096)
				n, client, e := pc.ReadFrom(buf)
				if e != nil {
					checked <- e
					return
				}
				if n < 3 || !bytes.Equal(buf[:3], []byte{0, 0, 0}) {
					checked <- fmt.Errorf("SOCKS packet prefix")
					return
				}
				packet := bytes.NewReader(buf[3:n])
				host, header, e := readIndependentSocksAddress(packet)
				if e != nil {
					checked <- e
					return
				}
				if e = checkIndependentLoopback(host); e != nil {
					checked <- e
					return
				}
				remote, e := net.DialTimeout("udp4", host, time.Second)
				if e != nil {
					checked <- e
					return
				}
				defer remote.Close()
				remote.SetDeadline(time.Now().Add(5 * time.Second))
				if _, e = remote.Write(buf[n-packet.Len() : n]); e != nil {
					checked <- e
					return
				}
				response := make([]byte, 4096)
				m, e := remote.Read(response)
				if e != nil {
					checked <- e
					return
				}
				wire := append([]byte{0, 0, 0}, header...)
				wire = append(wire, response[:m]...)
				_, e = pc.WriteTo(wire, client)
				checked <- e
				// UDP association control remains open until the actual adapter closes it.
				io.Copy(io.Discard, reader)
				return
			}
		}
		if e = checkIndependentLoopback(target); e != nil {
			checked <- e
			return
		}
		remote, e := net.DialTimeout("tcp4", target, time.Second)
		if e != nil {
			checked <- e
			return
		}
		defer remote.Close()
		remote.SetDeadline(time.Now().Add(5 * time.Second))
		reply := []byte("HTTP/1.1 200 Connection established\r\n\r\n")
		if kind == "socks" {
			reply = []byte{5, 0, 0, 1, 127, 0, 0, 1, 0, 0}
		}
		if _, e = conn.Write(reply); e != nil {
			checked <- e
			return
		}
		go func() { io.Copy(remote, reader) }()
		_, e = io.Copy(conn, remote)
		checked <- e
	}()
	return listener.Addr(), checked
}
func checkIndependentLoopback(address string) error {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return err
	}
	if host != "127.0.0.1" {
		return fmt.Errorf("nonfixture destination %q", address)
	}
	return nil
}
func readIndependentSocksAddress(r io.Reader) (string, []byte, error) {
	var tag [1]byte
	if _, e := io.ReadFull(r, tag[:]); e != nil {
		return "", nil, e
	}
	wire := []byte{tag[0]}
	var host string
	switch tag[0] {
	case 1, 4:
		size := 4
		if tag[0] == 4 {
			size = 16
		}
		p := make([]byte, size)
		if _, e := io.ReadFull(r, p); e != nil {
			return "", nil, e
		}
		wire = append(wire, p...)
		host = net.IP(p).String()
	case 3:
		if _, e := io.ReadFull(r, tag[:]); e != nil {
			return "", nil, e
		}
		wire = append(wire, tag[0])
		p := make([]byte, int(tag[0]))
		if _, e := io.ReadFull(r, p); e != nil {
			return "", nil, e
		}
		wire = append(wire, p...)
		host = string(p)
	default:
		return "", nil, fmt.Errorf("invalid SOCKS address type %d", tag[0])
	}
	var port [2]byte
	if _, e := io.ReadFull(r, port[:]); e != nil {
		return "", nil, e
	}
	wire = append(wire, port[:]...)
	return net.JoinHostPort(host, strconv.Itoa(int(binary.BigEndian.Uint16(port[:])))), wire, nil
}
