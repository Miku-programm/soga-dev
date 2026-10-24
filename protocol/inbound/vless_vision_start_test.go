package inbound

import (
	"bytes"
	"context"
	"crypto/ecdh"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"github.com/gofrs/uuid"
	"io"
	"net"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	Q "recovered.local/soga/panel/registry"
	R "recovered.local/soga/reality"
	"testing"
	"time"
)

// Independently encoded Vision padding and protobuf Addons cross the original
// MFF Start/authentication/flow checks and reach a real TCP destination. These
// cases exercise padding mode; they do not claim TLS-inside-TLS raw splicing.
func TestNativeVLESSVisionTLSRealityStart(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	cert, pemKey := hysteriaFixtureCertificate(t)
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM([]byte(cert))
	for _, security := range []string{"tls", "reality"} {
		t.Run(security, func(t *testing.T) {
			node := &L.Wr5SVP{EMlXEc: 32, I3k4JnNp7nB: cert, XATiVun: pemKey, Ck6DxjB: L.VkkFLH{XmIaPj3CzOF: "tcp", WAdJ4ftdZFk: L.D7FHIoaa8s4a(security), NMzMUQ2: "xtls-rprx-vision"}}
			var key *ecdh.PrivateKey
			var dest string
			if security == "reality" {
				var e error
				key, e = ecdh.X25519().GenerateKey(rand.Reader)
				if e != nil {
					t.Fatal(e)
				}
				dest = nativeRealityDestination(t)
				node.Ck6DxjB.BANVIZ72wh = []string{"hysteria"}
				node.Ck6DxjB.CrS6C1 = base64.RawURLEncoding.EncodeToString(key.Bytes())
				node.Ck6DxjB.IqSfNzea = dest
				node.Ck6DxjB.DxKJ7lSe3m8F = []string{"0102030405060708"}
			}
			core, e := oHr1LeATmS(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{panelVless(9, "11111111-1111-4111-8111-111111111111")}}})
			if e != nil {
				t.Fatal(e)
			}
			v := core.(*KHonGaqbIy_6)
			defer v.Stop()
			v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
			if e = v.Start(); e != nil {
				t.Fatal(e)
			}
			if security == "reality" {
				k := dest + " hysteria 2"
				defer R.GlobalPostHandshakeRecordsLens.Delete(k)
				until := time.Now().Add(4 * time.Second)
				for {
					if val, ok := R.GlobalPostHandshakeRecordsLens.Load(k); ok {
						if _, ok = val.([]int); ok {
							break
						}
					}
					if time.Now().After(until) {
						t.Fatal("probe not ready")
					}
					time.Sleep(time.Millisecond)
				}
			}
			addr := v.kX5i8U[0].Addr().String()
			var c net.Conn
			if security == "reality" {
				c = nativeRealityClient(t, addr, key)
			} else {
				c, e = tls.DialWithDialer(&net.Dialer{Timeout: time.Second}, "tcp", addr, &tls.Config{RootCAs: roots, ServerName: "hysteria", MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13})
				if e != nil {
					t.Fatal(e)
				}
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(5 * time.Second))
			port, done := trojanMatrixTarget(t, false)
			id := uuid.FromStringOrNil("11111111-1111-4111-8111-111111111111")
			wire := append([]byte{0}, id[:]...)
			addons := append([]byte{0x0a, 16}, []byte("xtls-rprx-vision")...)
			wire = append(wire, byte(len(addons)))
			wire = append(wire, addons...)
			wire = append(wire, 1)
			wire = binary.BigEndian.AppendUint16(wire, uint16(port))
			wire = append(wire, 1, 127, 0, 0, 1)
			payload := bytes.Repeat([]byte("native-transport\x00"), 91)
			wire = append(wire, id[:]...)
			wire = append(wire, 0)
			wire = binary.BigEndian.AppendUint16(wire, uint16(len(payload)))
			wire = append(wire, 0, 0)
			wire = append(wire, payload...)
			if _, e = c.Write(wire); e != nil {
				t.Fatal(e)
			}
			var header [18]byte
			if _, e = io.ReadFull(c, header[:]); e != nil {
				t.Fatal(e)
			}
			if !bytes.Equal(header[:2], []byte{0, 0}) || !bytes.Equal(header[2:], id[:]) {
				t.Fatal("response UUID/version")
			}
			var got []byte
			for len(got) < len(payload) {
				var frame [5]byte
				if _, e = io.ReadFull(c, frame[:]); e != nil {
					t.Fatal(e)
				}
				n, pad := int(binary.BigEndian.Uint16(frame[1:3])), int(binary.BigEndian.Uint16(frame[3:]))
				block := make([]byte, n+pad)
				if _, e = io.ReadFull(c, block); e != nil {
					t.Fatal(e)
				}
				if frame[0] > 1 {
					t.Fatal("unexpected direct-splice command")
				}
				got = append(got, block[:n]...)
			}
			if !bytes.Equal(got, payload) {
				t.Fatal("Vision target payload")
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
