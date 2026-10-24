package inbound

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	CA "recovered.local/soga/util/rtx"
	G "recovered.local/soga/transport/grpcproto"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

// 从原生入站构造和 Start 进入共享监听器，独立客户端发送真实协议字节。
// 下游是真实 TCP/UDP 目标，不以替代内核或模拟转发器代替受测代码。
func TestNativeTrojanTransportMatrix(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	// 原H2 Close的cancel→Flush存在并发关闭竞态；此隔离进程只验收串行wire。
	// 原顺序及panic另见AKP5mbedNKj/close_order_test.go，不宣称并发稳定性。
	runtime.GOMAXPROCS(1)
	path := filepath.Join(t.TempDir(), "soga.conf")
	if err := os.WriteFile(path, []byte("listen=127.0.0.1\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\n"), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(path)
	if err := C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	cert, key := hysteriaFixtureCertificate(t)
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM([]byte(cert)) {
		t.Fatal("fixture certificate")
	}
	for _, network := range []string{"tcp", "ws", "h2", "grpc", "xhttp"} {
		for _, encrypted := range []bool{false, true} {
			for _, udp := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/tls=%t/udp=%t", network, encrypted, udp), func(t *testing.T) {
					// XHTTP 原 Addr 有关闭接收队列的副作用，因此先选择端口再 Start。
					probe, err := net.Listen("tcp4", "127.0.0.1:0")
					if err != nil {
						t.Fatal(err)
					}
					port := probe.Addr().(*net.TCPAddr).Port
					addr := probe.Addr().String()
					probe.Close()
					security := "none"
					if encrypted {
						security = "tls"
					}
					node := &L.Wr5SVP{EMlXEc: 32, I3k4JnNp7nB: cert, XATiVun: key, JhIYwEukjd: L.L9AiYC{GbD_2PU: port, T4cuhG7MTbG: L.BoulTn31(network), AJjaqrBjj: L.D7FHIoaa8s4a(security), PBDb3Wc: "/native", QVlaSCZ4RCK: "native.Tunnel"}}
					panel := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{trojanPanelFixture(9, "transport-password")}}
					core, err := f1mPzegGt(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: panel})
					if err != nil {
						t.Fatal(err)
					}
					v := core.(*EfaIFpp)
					defer v.Stop()
					v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
					if err = v.Start(); err != nil {
						t.Fatal(err)
					}
					targetPort, targetDone := trojanMatrixTarget(t, udp)
					target := binary.BigEndian.AppendUint16([]byte{1, 127, 0, 0, 1}, uint16(targetPort))
					digest := sha256.Sum224([]byte("transport-password"))
					wire := []byte(hex.EncodeToString(digest[:]))
					command := byte(1)
					if udp {
						command = 3
					}
					wire = append(wire, 13, 10, command)
					wire = append(wire, target...)
					wire = append(wire, 13, 10)
					expected := bytes.Repeat([]byte("native-transport\x00"), 91)
					if udp {
						frame := append([]byte(nil), target...)
						frame = binary.BigEndian.AppendUint16(frame, uint16(len(expected)))
						frame = append(frame, 13, 10)
						expected = append(frame, expected...)
					}
					wire = append(wire, expected...)
					clientTLS := &tls.Config{RootCAs: roots, ServerName: "hysteria"}
					got := trojanMatrixExchange(t, network, encrypted, addr, clientTLS, wire, len(expected))
					if !bytes.Equal(got, expected) {
						t.Fatalf("target echo mismatch: %d / %d", len(got), len(expected))
					}
					select {
					case err = <-targetDone:
						if err != nil {
							t.Fatal(err)
						}
					case <-time.After(5 * time.Second):
						t.Fatal("target not reached")
					}
				})
			}
		}
	}
}

func trojanMatrixTarget(t *testing.T, udp bool) (int, <-chan error) {
	t.Helper()
	done := make(chan error, 1)
	if udp {
		l, err := net.ListenPacket("udp4", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { l.Close() })
		go func() {
			l.SetDeadline(time.Now().Add(5 * time.Second))
			p := make([]byte, 8192)
			n, a, e := l.ReadFrom(p)
			if e == nil {
				_, e = l.WriteTo(p[:n], a)
			}
			done <- e
		}()
		return l.LocalAddr().(*net.UDPAddr).Port, done
	}
	l, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { l.Close() })
	go func() {
		c, e := l.Accept()
		if e != nil {
			done <- e
			return
		}
		defer c.Close()
		c.SetDeadline(time.Now().Add(5 * time.Second))
		p := make([]byte, len("native-transport\x00")*91)
		_, e = io.ReadFull(c, p)
		if e == nil {
			_, e = c.Write(p)
		}
		done <- e
	}()
	return l.Addr().(*net.TCPAddr).Port, done
}

func trojanMatrixExchange(t *testing.T, network string, encrypted bool, addr string, tc *tls.Config, wire []byte, size int) []byte {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	read := func(r io.Reader) []byte {
		t.Helper()
		p := make([]byte, size)
		if n, e := io.ReadFull(r, p); e != nil {
			t.Fatalf("%s tls=%t response read %d/%d bytes: %v; prefix=%x", network, encrypted, n, size, e, p[:min(n, 64)])
		}
		return p
	}
	scheme := "http"
	if encrypted {
		scheme = "https"
	}
	switch network {
	case "tcp":
		raw, e := (&net.Dialer{}).DialContext(ctx, "tcp", addr)
		if e != nil {
			t.Fatal(e)
		}
		defer raw.Close()
		raw.SetDeadline(time.Now().Add(5 * time.Second))
		var c net.Conn = raw
		if encrypted {
			c = tls.Client(raw, tc)
		}
		if _, e = c.Write(wire); e != nil {
			t.Fatal(e)
		}
		return read(c)
	case "ws":
		scheme = "ws"
		if encrypted {
			scheme = "wss"
		}
		d := websocket.Dialer{TLSClientConfig: tc, HandshakeTimeout: 3 * time.Second}
		c, resp, e := d.DialContext(ctx, scheme+"://"+addr+"/native", nil)
		if e != nil {
			t.Fatalf("websocket: %v (%v)", e, resp)
		}
		defer c.Close()
		c.SetReadDeadline(time.Now().Add(5 * time.Second))
		if e = c.WriteMessage(websocket.BinaryMessage, wire); e != nil {
			t.Fatal(e)
		}
		var out []byte
		for len(out) < size {
			_, p, e := c.ReadMessage()
			if e != nil {
				t.Fatal(e)
			}
			out = append(out, p...)
		}
		return out
	case "grpc":
		var creds credentials.TransportCredentials = insecure.NewCredentials()
		if encrypted {
			creds = credentials.NewTLS(tc)
		}
		c, e := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
		if e != nil {
			t.Fatal(e)
		}
		defer c.Close()
		stream, e := c.NewStream(ctx, &grpc.StreamDesc{ServerStreams: true, ClientStreams: true}, "/native.Tunnel/Tun")
		if e != nil {
			t.Fatal(e)
		}
		if e = stream.SendMsg(&G.Z27tYEZ9hy1o{Q2bCNbJ12i: wire}); e != nil {
			t.Fatal(e)
		}
		var out []byte
		for len(out) < size {
			var p G.Z27tYEZ9hy1o
			if e = stream.RecvMsg(&p); e != nil {
				t.Fatal(e)
			}
			out = append(out, p.Q2bCNbJ12i...)
		}
		return out
	case "h2":
		tr := &http2.Transport{AllowHTTP: !encrypted, TLSClientConfig: tc}
		if !encrypted {
			tr.DialTLSContext = func(ctx context.Context, n, a string, _ *tls.Config) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, n, a)
			}
		}
		defer tr.CloseIdleConnections()
		upload, writer := io.Pipe()
		defer upload.Close()
		defer writer.Close()
		go func() {
			_, e := writer.Write(wire)
			if e != nil {
				writer.CloseWithError(e)
			}
		}()
		req, e := http.NewRequestWithContext(ctx, "POST", scheme+"://"+addr+"/native", upload)
		if e != nil {
			t.Fatal(e)
		}
		resp, e := tr.RoundTrip(req)
		if e != nil {
			t.Fatal(e)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 || resp.ProtoMajor != 2 {
			t.Fatal(resp.Status, resp.Proto)
		}
		return read(resp.Body)
	case "xhttp":
		tr := &http.Transport{TLSClientConfig: tc}
		defer tr.CloseIdleConnections()
		client := &http.Client{Transport: tr}
		req, e := http.NewRequestWithContext(ctx, "POST", scheme+"://"+addr+"/fixture/0", bytes.NewReader(wire))
		if e != nil {
			t.Fatal(e)
		}
		req.Header.Set("Referer", "http://fixture/?x_padding="+strings.Repeat("X", 100))
		resp, e := client.Do(req)
		if e != nil {
			t.Fatal(e)
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal(resp.Status)
		}
		req, e = http.NewRequestWithContext(ctx, "GET", scheme+"://"+addr+"/fixture", nil)
		if e != nil {
			t.Fatal(e)
		}
		req.Header.Set("Referer", "http://fixture/?x_padding="+strings.Repeat("X", 100))
		resp, e = client.Do(req)
		if e != nil {
			t.Fatal(e)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatal(resp.Status)
		}
		return read(resp.Body)
	}
	t.Fatalf("unknown test transport %q", network)
	return nil
}
