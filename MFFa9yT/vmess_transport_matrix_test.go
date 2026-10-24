package MFFa9yT

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"hash/fnv"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	CA "recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	V "recovered.local/soga/gk1z9v"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

// 实际原生 VMess 构造→Start→独立客户端 AEAD 认证→真实目标。
// 只复用传输 fixture 与独立密码学 helper，不调用服务端用户的派生密钥或客户端内核。
func TestNativeVMessTransportMatrix(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	// 原 H2 callback cancel→Flush 的关闭竞态仍保留；隔离进程串行 wire
	// 验收不代表并发关闭稳定性，原问题另有 AKP close_order_test.go 证据。
	runtime.GOMAXPROCS(1)
	config := filepath.Join(t.TempDir(), "soga.conf")
	if e := os.WriteFile(config, []byte("listen=127.0.0.1\nvmess_auth_mode=aead\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\n"), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(config)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
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
					probe, e := net.Listen("tcp4", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					port, addr := probe.Addr().(*net.TCPAddr).Port, probe.Addr().String()
					probe.Close()
					node := &L.Wr5SVP{EMlXEc: 31, I3k4JnNp7nB: cert, XATiVun: key, SwxJNgPpg: L.JW0CeHo3O{CKsOpayd: port, F2GXyZZ: L.BoulTn31(network), UVeahqDBT: encrypted, ChJnngzb: "/native", KWffzRDF: "native.Tunnel"}}
					core, e := ag90n_rZjJO(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{vmessPanelFixture(7, "11111111-1111-4111-8111-111111111111")}}})
					if e != nil {
						t.Fatal(e)
					}
					v := core.(*K4avFiYOPi)
					defer v.Stop()
					v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
					if e = v.Start(); e != nil {
						t.Fatal(e)
					}
					targetPort, done := trojanMatrixTarget(t, udp)
					wire, header, payload := independentVMessTransportWire(t, targetPort, udp)

					got := trojanMatrixExchange(t, network, encrypted, addr, &tls.Config{RootCAs: roots, ServerName: "hysteria"}, wire, 38+2+len(payload))
					reader := bytes.NewReader(got)
					rk, iv := clientResponseKeys(header, true)
					response, e := readClientResponse(reader, rk, iv, true)
					if e != nil || len(response) != 4 || response[0] != 0x42 {
						t.Fatalf("response=%x err=%v", response, e)
					}
					var reply []byte
					for len(reply) < len(payload) {
						var size [2]byte
						if _, e = io.ReadFull(reader, size[:]); e != nil {
							t.Fatal(e)
						}
						frame := make([]byte, int(binary.BigEndian.Uint16(size[:])))
						if _, e = io.ReadFull(reader, frame); e != nil {
							t.Fatal(e)
						}
						reply = append(reply, frame...)
					}
					if !bytes.Equal(reply, payload) {
						t.Fatalf("target echo %d != %d", len(reply), len(payload))
					}
					select {
					case e = <-done:
						if e != nil {
							t.Fatal(e)
						}
					case <-time.After(5 * time.Second):
						t.Fatal("target not reached")
					}
				})
			}
		}
	}
}
func independentVMessTransportWire(t *testing.T, port int, udp bool) (wire, clear, payload []byte) {
	t.Helper()
	// Port-derived bytes used only 256 possible key/IV pairs and triggered the
	// original process-wide replay filter across independent transport cases.
	var session [32]byte
	if _, err := rand.Read(session[:]); err != nil {
		t.Fatal(err)
	}
	return independentVMessTransportWireKeys(t, port, udp, session)
}
func independentVMessTransportWireKeys(t *testing.T, port int, udp bool, session [32]byte) (wire, clear, payload []byte) {
	t.Helper()
	id := uuid.FromStringOrNil("11111111-1111-4111-8111-111111111111")
	commandKey := md5.Sum(append(append([]byte(nil), id[:]...), []byte("c48619fe-8f02-49e0-b9e9-edf763e17e21")...))
	var authPlain, authID [16]byte
	binary.BigEndian.PutUint64(authPlain[:8], uint64(time.Now().Unix()))
	if _, err := rand.Read(authPlain[8:12]); err != nil {
		t.Fatal(err)
	}
	binary.BigEndian.PutUint32(authPlain[12:], crc32.ChecksumIEEE(authPlain[:12]))
	authCipher, e := aes.NewCipher(clientKDF(commandKey[:], []byte("AES Auth ID Encryption"))[:16])
	if e != nil {
		t.Fatal(e)
	}
	authCipher.Encrypt(authID[:], authPlain[:])
	clear = make([]byte, 38)
	clear[0] = 1
	copy(clear[1:33], session[:])
	clear[33] = 0x42
	clear[34] = 1
	clear[35] = 5
	clear[37] = 1
	if udp {
		clear[37] = 2
	}
	clear = binary.BigEndian.AppendUint16(clear, uint16(port))
	clear = append(clear, 1, 127, 0, 0, 1)
	h := fnv.New32a()
	h.Write(clear)
	clear = binary.BigEndian.AppendUint32(clear, h.Sum32())
	nonce := make([]byte, 8)
	if _, err := rand.Read(nonce); err != nil {
		t.Fatal(err)
	}
	seal := func(key, iv string, p []byte) []byte {
		return clientGCM(clientKDF(commandKey[:], []byte(key), authID[:], nonce)[:16], clientKDF(commandKey[:], []byte(iv), authID[:], nonce)[:12], p, authID[:])
	}
	wire = append(wire, authID[:]...)
	wire = append(wire, seal("VMess Header AEAD Key_Length", "VMess Header AEAD Nonce_Length", binary.BigEndian.AppendUint16(nil, uint16(len(clear))))...)
	wire = append(wire, nonce...)
	wire = append(wire, seal("VMess Header AEAD Key", "VMess Header AEAD Nonce", clear)...)
	payload = bytes.Repeat([]byte("native-transport\x00"), 91)
	wire = binary.BigEndian.AppendUint16(wire, uint16(len(payload)))
	wire = append(wire, payload...)
	return
}

// Original uaQJpenJx8c @0x1d179c0 checks zero AEAD B2pr6gMaPO + IV + key,
// not the target port or the encrypted AuthID. Different destinations therefore
// must not reuse request entropy. Keep this rejection as a real backend check.
func TestVMessTransportFixtureFreshKeysAndOriginalReplay(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	_, a, _ := independentVMessTransportWire(t, 35001, false)
	_, b, _ := independentVMessTransportWire(t, 35257, false)
	if bytes.Equal(a[1:33], b[1:33]) {
		t.Fatal("fixture reused session key/IV for ports with identical low byte")
	}
	v := newVmessAuthFixture(t)
	v.spcNlUL(L.Sxcf3wu3L{vmessPanelFixture(7, "11111111-1111-4111-8111-111111111111")})
	var session [32]byte
	if _, err := rand.Read(session[:]); err != nil {
		t.Fatal(err)
	}
	for i, port := range []int{35001, 35257} {
		wire, _, _ := independentVMessTransportWireKeys(t, port, false, session)
		auth := v.y80zV6("127.0.0.1", [16]byte(wire[:16]))
		if auth == nil {
			t.Fatal("independent AEAD authentication failed")
		}
		server, client := net.Pipe()
		server.SetDeadline(time.Now().Add(time.Second))
		client.SetDeadline(time.Now().Add(time.Second))
		writeDone := make(chan struct{})
		go func() { defer close(writeDone); client.Write(wire[16:]); client.Close() }()
		framed, err := V.EAQVlR(server, auth)
		if framed != nil {
			framed.Free()
		}
		server.Close()
		client.Close()
		<-writeDone
		if i == 0 && err != nil {
			t.Fatalf("fresh request: %v", err)
		}
		if i == 1 && (err == nil || err.Error() != "vmess salt repeat") {
			t.Fatalf("different-port repeated key/IV: %v", err)
		}
	}
}
