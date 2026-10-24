package inbound

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	quic "github.com/apernet/quic-go"
	"github.com/apernet/quic-go/http3"
	"github.com/apernet/quic-go/quicvarint"
	"io"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	CA "recovered.local/soga/util/rtx"
	O "recovered.local/soga/transport/salamander"
	L "recovered.local/soga/panel/types"
	P "recovered.local/soga/net/addrtype"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"testing"
	"time"
	"unsafe"
)

func hysteriaFixtureCertificate(t *testing.T) (string, string) {
	t.Helper()
	pub, key, e := ed25519.GenerateKey(rand.Reader)
	if e != nil {
		t.Fatal(e)
	}
	cert := &x509.Certificate{SerialNumber: big.NewInt(91), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), DNSNames: []string{"hysteria"}, KeyUsage: x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}}
	der, e := x509.CreateCertificate(rand.Reader, cert, cert, pub, key)
	if e != nil {
		t.Fatal(e)
	}
	pk, e := x509.MarshalPKCS8PrivateKey(key)
	if e != nil {
		t.Fatal(e)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})), string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: pk}))
}

// This exercises the actual constructor, TLS/UDP listener, password lookup,
// shared routing and target sockets. TCP and UDP headers below are built and
// parsed independently of the recovered Hysteria wire helpers.
func TestNativeHysteriaConstructStartTCPUDPStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	if unsafe.Sizeof(Aq5i0Bwez8{}) != 88 {
		t.Fatal("RTTI")
	}
	path := filepath.Join(t.TempDir(), "soga.conf")
	if e := os.WriteFile(path, []byte("listen=127.0.0.1\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\n"), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(path)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	cert, key := hysteriaFixtureCertificate(t)
	for _, obfs := range []string{"plain", "salamander"} {
		for _, network := range []string{"tcp", "udp"} {
			t.Run(obfs+"/"+network, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				node := &L.Wr5SVP{EMlXEc: 39, I3k4JnNp7nB: cert, XATiVun: key, AYf5kQSQW: L.ZMYytVS{B30Jt2kO: P.B3YC1rotZ(obfs), UjgtNuouuKVl: "fixture-obfs", W4ggEz: 100, HnHarW: 100}}
				panel := &L.Aga91y{WZ89pTAOrHK4: 91, JAfXWwzYrW: &L.WQM2tKA{CBWeRfrO1ahs: "hysteria-native-password"}}
				core, e := hpdxYbrDSPs(&DV_XPJ4z0{Dmpl48sXM: ctx, HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{panel}}})
				if e != nil {
					t.Fatal(e)
				}
				v := core.(*Aq5i0Bwez8)
				defer v.Stop()
				v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0())) // Original hourly no-request branch.
				if e = v.Start(); e != nil {
					t.Fatal(e)
				}
				if len(v.uqbfp_xRjia) != 1 {
					t.Fatal("listeners")
				}
				raw, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer raw.Close()
				packet := raw
				if C.UDPProxyProtocolEnabled() {
					packet = &hysteriaProxyClient{PacketConn: raw}
				}
				if obfs == "salamander" {
					scrambler, e := O.Ea9POnRF([]byte("fixture-obfs"))
					if e != nil {
						t.Fatal(e)
					}
					packet = O.WrapFromHysteriaInline(packet, scrambler)
				}
				conn, e := quic.Dial(ctx, packet, v.uqbfp_xRjia[0].Addr(), &tls.Config{InsecureSkipVerify: true, NextProtos: []string{"h3"}}, &quic.Config{EnableDatagrams: true})
				if e != nil {
					t.Fatal(e)
				}
				defer conn.CloseWithError(0, "")
				// Hysteria owns QUIC datagrams; enabling HTTP/3 datagrams here would
				// start a competing HTTP/3 receiver in this historical client fork.
				transport := &http3.Transport{}
				client := transport.NewClientConn(conn)
				for _, password := range []string{"incorrect", "hysteria-native-password"} {
					r, e := http.NewRequestWithContext(ctx, "POST", "https://hysteria/auth", nil)
					if e != nil {
						t.Fatal(e)
					}
					r.Header.Set("Hysteria-Auth", password)
					r.Header.Set("Hysteria-CC-RX", "10000000")
					response, e := client.RoundTrip(r)
					if e != nil {
						t.Fatal(e)
					}
					io.Copy(io.Discard, response.Body)
					response.Body.Close()
					want := 233
					if password == "incorrect" {
						want = 404
					}
					if response.StatusCode != want {
						t.Fatal("auth", response.StatusCode, want)
					}
				}
				var dest string
				if network == "tcp" {
					echo, e := net.Listen("tcp4", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					defer echo.Close()
					dest = echo.Addr().String()
					go func() {
						c, e := echo.Accept()
						if e != nil {
							return
						}
						defer c.Close()
						c.SetDeadline(time.Now().Add(5 * time.Second))
						io.Copy(c, c)
					}()
				} else {
					echo, e := net.ListenPacket("udp4", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					defer echo.Close()
					dest = echo.LocalAddr().String()
					go func() {
						echo.SetDeadline(time.Now().Add(5 * time.Second))
						b := make([]byte, 4096)
						n, a, e := echo.ReadFrom(b)
						if e == nil {
							echo.WriteTo(b[:n], a)
						}
					}()
				}
				payload := bytes.Repeat([]byte("native-hysteria\x00"), 40)
				if network == "tcp" {
					stream, e := conn.OpenStreamSync(ctx)
					if e != nil {
						t.Fatal(e)
					}
					defer stream.Close()
					stream.SetDeadline(time.Now().Add(5 * time.Second))
					request := quicvarint.Append(nil, 0x401)
					request = quicvarint.Append(request, uint64(len(dest)))
					request = append(request, dest...)
					request = quicvarint.Append(request, 0)
					request = append(request, payload...)
					if _, e = stream.Write(request); e != nil {
						t.Fatal(e)
					}
					var status [1]byte
					if _, e = io.ReadFull(stream, status[:]); e != nil || status[0] != 0 {
						t.Fatal("TCP status", status, e)
					}
					for i := 0; i < 2; i++ {
						n, e := quicvarint.Read(quicvarint.NewReader(stream))
						if e != nil {
							t.Fatal(e)
						}
						if _, e = io.CopyN(io.Discard, stream, int64(n)); e != nil {
							t.Fatal(e)
						}
					}
					got := make([]byte, len(payload))
					if _, e = io.ReadFull(stream, got); e != nil || !bytes.Equal(got, payload) {
						t.Fatal("TCP echo", e)
					}
				} else {
					wire := binary.BigEndian.AppendUint32(nil, 73)
					wire = binary.BigEndian.AppendUint16(wire, 1)
					wire = append(wire, 0, 1)
					wire = quicvarint.Append(wire, uint64(len(dest)))
					wire = append(wire, dest...)
					wire = append(wire, payload...)
					if e = conn.SendDatagram(wire); e != nil {
						t.Fatal(e)
					}
					wire, e = conn.ReceiveDatagram(ctx)
					if e != nil {
						t.Fatal(e)
					}
					if len(wire) < 9 || binary.BigEndian.Uint32(wire) != 73 || wire[6] != 0 || wire[7] != 1 {
						t.Fatal("UDP header", wire)
					}
					br := bytes.NewReader(wire[8:])
					n, e := quicvarint.Read(br)
					if e != nil {
						t.Fatal(e)
					}
					addr := make([]byte, n)
					if _, e = io.ReadFull(br, addr); e != nil {
						t.Fatal(e)
					}
					got, _ := io.ReadAll(br)
					if string(addr) != dest || !bytes.Equal(got, payload) {
						t.Fatal("UDP echo", string(addr), len(got))
					}
				}
				u, ok := v.jc_QPN6.Get(91)
				if !ok {
					t.Fatal("missing user")
				}
				userCtx := u.GetCtx()
				bound := v.uqbfp_xRjia[0].Addr().String()
				v.Stop()
				select {
				case <-v.zb8f3cHWW8v.Done():
				default:
					t.Fatal("Stop context")
				}
				select {
				case <-userCtx.Done():
				default:
					t.Fatal("Stop user")
				}
				if _, ok = v.jc_QPN6.Get(91); !ok {
					t.Fatal("Stop unexpectedly cleared user map")
				}
				// Original Stop closes the UDP socket before resetting user
				// contexts; it cannot guarantee a close frame reaches clients.
				rebound, e := net.ListenPacket("udp", bound)
				if e != nil {
					t.Fatal("Stop did not release UDP listener", e)
				}
				rebound.Close()
			})
		}
	}
}
