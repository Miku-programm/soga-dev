package inbound

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"github.com/gofrs/uuid"
	"net"
	"os"
	"path/filepath"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"runtime"
	"testing"
	"time"
)

// 五种原监听装配×TLS/plain×TCP/UDP。复用独立传输客户端，而 VLESS
// 字节头由本测试独立构造；不调用原 VLESS 客户端使两端共享实现错误。
func TestNativeVLESSTransportMatrix(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	// The original h2 Close cancels its handler before Flush, retaining a known
	// concurrent teardown panic. This isolated wire fixture serializes execution;
	// it proves transport/auth/relay behavior, not concurrent-close stability.
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
		t.Fatal("certificate")
	}
	for _, network := range []string{"tcp", "ws", "h2", "grpc", "xhttp"} {
		for _, encrypted := range []bool{false, true} {
			for _, udp := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/tls=%t/udp=%t", network, encrypted, udp), func(t *testing.T) {
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
					node := &L.Wr5SVP{EMlXEc: 32, I3k4JnNp7nB: cert, XATiVun: key, Ck6DxjB: L.VkkFLH{GJtOdKOWBiFP: port, XmIaPj3CzOF: L.BoulTn31(network), WAdJ4ftdZFk: L.D7FHIoaa8s4a(security), Nd3mBux1: "/native", J1efyloFyhk: "native.Tunnel"}}
					core, err := oHr1LeATmS(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{panelVless(9, "11111111-1111-4111-8111-111111111111")}}})
					if err != nil {
						t.Fatal(err)
					}
					v := core.(*KHonGaqbIy_6)
					defer v.Stop()
					v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
					if err = v.Start(); err != nil {
						t.Fatal(err)
					}
					targetPort, done := trojanMatrixTarget(t, udp)
					wire, want := vlessMatrixWire(udp, targetPort)
					got := trojanMatrixExchange(t, network, encrypted, addr, &tls.Config{RootCAs: roots, ServerName: "hysteria"}, wire, len(want))
					if !bytes.Equal(got, want) {
						t.Fatalf("native target response %x != %x", got, want)
					}
					select {
					case err = <-done:
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
func vlessMatrixWire(udp bool, port int) (wire, want []byte) {
	id := uuid.FromStringOrNil("11111111-1111-4111-8111-111111111111")
	wire = append([]byte{0}, id[:]...)
	command := byte(1)
	if udp {
		command = 2
	}
	wire = append(wire, 0, command)
	wire = binary.BigEndian.AppendUint16(wire, uint16(port))
	wire = append(wire, 1, 127, 0, 0, 1)
	payload := bytes.Repeat([]byte("native-transport\x00"), 91)
	want = []byte{0, 0}
	if udp {
		wire = binary.BigEndian.AppendUint16(wire, uint16(len(payload)))
		want = binary.BigEndian.AppendUint16(want, uint16(len(payload)))
	}
	wire = append(wire, payload...)
	want = append(want, payload...)
	return
}
