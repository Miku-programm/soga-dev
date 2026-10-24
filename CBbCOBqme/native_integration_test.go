package CBbCOBqme

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	quic "github.com/apernet/quic-go"
	"github.com/apernet/quic-go/http3"
	"github.com/apernet/quic-go/quicvarint"
	"io"
	"math/big"
	"net"
	"net/http"
	B "recovered.local/soga/BaEWsPy"
	O "recovered.local/soga/GLaePRKYM"
	F "recovered.local/soga/M14gyoUSscE"
	W "recovered.local/soga/mlbRsO44"
	U "recovered.local/soga/nzATvajt"
	"sync/atomic"
	"testing"
	"time"
)

type wireUser struct {
	U.PgaNd3014U
	ctx      context.Context
	up, down atomic.Int64
}

func (u *wireUser) GetCtx() context.Context { return u.ctx }
func (u *wireUser) AddUp(n int)             { u.up.Add(int64(n)) }
func (u *wireUser) AddDown(n int)           { u.down.Add(int64(n)) }
func (u *wireUser) IsTrafficOut() bool      { return false }

type wireAuth struct {
	user  *wireUser
	calls atomic.Int32
}

func (a *wireAuth) Authenticate(_ net.Addr, password string, _ uint64) (U.PgaNd3014U, bool) {
	a.calls.Add(1)
	return a.user, password == "test-password"
}
func wireCertificate(t *testing.T) tls.Certificate {
	t.Helper()
	pub, key, e := ed25519.GenerateKey(rand.Reader)
	if e != nil {
		t.Fatal(e)
	}
	cert := &x509.Certificate{SerialNumber: big.NewInt(77), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), DNSNames: []string{"hysteria"}, KeyUsage: x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}}
	der, e := x509.CreateCertificate(rand.Reader, cert, cert, pub, key)
	if e != nil {
		t.Fatal(e)
	}
	return tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}
}
func TestNativeHysteriaWire(t *testing.T) {
	for _, obfs := range []string{"plain", "salamander"} {
		for _, rx := range []string{"0", "10000000"} {
			t.Run(obfs+rx, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				userCtx, stopUser := context.WithCancel(ctx)
				defer stopUser()
				u := &wireUser{ctx: userCtx}
				auth := &wireAuth{user: u}
				cfg, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{wireCertificate(t)}}, auth, 100, 100)
				if e != nil {
					t.Fatal(e)
				}
				raw, e := net.ListenPacket("udp", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				resolver := &B.D23y2888a_{}
				listener, e := VOhu4P9d2b(ctx, raw, obfs, "obfs-test", cfg, func() *B.D23y2888a_ { return resolver })
				if e != nil {
					raw.Close()
					t.Fatal(e)
				}
				defer listener.Close()
				clientRaw, e := net.ListenPacket("udp", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer clientRaw.Close()
				clientPacket := clientRaw
				if obfs == "salamander" {
					o, e := O.Ea9POnRF([]byte("obfs-test"))
					if e != nil {
						t.Fatal(e)
					}
					clientPacket = O.WrapFromHysteriaInline(clientRaw, o)
				}
				conn, e := quic.Dial(ctx, clientPacket, listener.Addr(), &tls.Config{InsecureSkipVerify: true, NextProtos: []string{"h3"}}, &quic.Config{EnableDatagrams: true})
				if e != nil {
					t.Fatal(e)
				}
				defer conn.CloseWithError(0, "")
				transport := &http3.Transport{}
				client := transport.NewClientConn(conn)
				request := func(password string) int {
					t.Helper()
					r, e := http.NewRequestWithContext(ctx, "POST", "https://hysteria/auth", nil)
					if e != nil {
						t.Fatal(e)
					}
					r.Header.Set("Hysteria-Auth", password)
					r.Header.Set("Hysteria-CC-RX", rx)
					resp, e := client.RoundTrip(r)
					if e != nil {
						t.Fatal(e)
					}
					io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
					return resp.StatusCode
				}
				if s := request("wrong"); s != 404 {
					t.Fatal("bad auth", s)
				}
				if s := request("test-password"); s != 233 {
					t.Fatal("auth", s)
				}
				if s := request("wrong-after-success"); s != 233 || auth.calls.Load() != 2 {
					t.Fatal("cached auth", s, auth.calls.Load())
				}
				stream, e := conn.OpenStreamSync(ctx)
				if e != nil {
					t.Fatal(e)
				}
				stream.SetDeadline(time.Now().Add(3 * time.Second))
				defer stream.Close()
				dest := "127.0.0.1:2345"
				requestBytes := quicvarint.Append(nil, 0x401)
				requestBytes = quicvarint.Append(requestBytes, uint64(len(dest)))
				requestBytes = append(requestBytes, dest...)
				requestBytes = quicvarint.Append(requestBytes, 0)
				requestBytes = append(requestBytes, "hello"...)
				if _, e = stream.Write(requestBytes); e != nil {
					t.Fatal(e)
				}
				var status [1]byte
				if _, e = io.ReadFull(stream, status[:]); e != nil || status[0] != 0 {
					t.Fatal(status, e)
				}
				br := quicvarint.NewReader(stream)
				message, e := quicvarint.Read(br)
				if e != nil {
					t.Fatal(e)
				}
				if _, e = io.CopyN(io.Discard, stream, int64(message)); e != nil {
					t.Fatal(e)
				}
				pad, e := quicvarint.Read(br)
				if e != nil {
					t.Fatal(e)
				}
				if _, e = io.CopyN(io.Discard, stream, int64(pad)); e != nil {
					t.Fatal(e)
				}
				select {
				case proxy := <-listener.eIECnHeZd8:
					proxy.SetDeadline(time.Now().Add(3 * time.Second))
					defer proxy.Close()
					if proxy.NetworkType() != 0 || proxy.GetUser() != u || proxy.GetDestAddr().String() != dest {
						t.Fatal("TCP metadata")
					}
					buf := make([]byte, 5)
					if _, e = io.ReadFull(proxy, buf); e != nil || string(buf) != "hello" {
						t.Fatal(string(buf), e)
					}
					if _, e = proxy.Write([]byte("world")); e != nil {
						t.Fatal(e)
					}
					if _, e = io.ReadFull(stream, buf); e != nil || string(buf) != "world" {
						t.Fatal(string(buf), e)
					}
				case <-ctx.Done():
					t.Fatal("TCP accept timeout")
				}
				if u.up.Load() != 5 || u.down.Load() != 5 {
					t.Fatal("traffic", u.up.Load(), u.down.Load())
				}
				packet := W.ItLMumzBf{D35ycLc5aB6r: 7, Kb9FLg7wuNE: 1, XrDcBXcD: dest, HmoFPk: []byte("udp")}
				b := make([]byte, packet.Size())
				packet.Serialize(b)
				if e = conn.SendDatagram(b); e != nil {
					t.Fatal(e)
				}
				select {
				case proxy := <-listener.eIECnHeZd8:
					udp, ok := proxy.(*aK1tZAfWP2z)
					if !ok {
						t.Fatal("UDP type")
					}
					defer udp.Close()
					udp.SetReadDeadline(time.Now().Add(3 * time.Second))
					buf := make([]byte, 5000)
					n, addr, e := udp.ReadFromProxy(buf)
					if e != nil || string(buf[:n]) != "udp" || addr.String() != dest {
						t.Fatal(n, addr, e)
					}
					for _, payload := range [][]byte{[]byte("reply"), bytes.Repeat([]byte("fragment"), 450)} {
						if n, e = udp.WriteToProxy(payload, addr); e != nil || n != len(payload) {
							t.Fatal(n, e)
						}
						reassembly := &F.YWXI4jW3{}
						for {
							data, e := conn.ReceiveDatagram(ctx)
							if e != nil {
								t.Fatal(e)
							}
							part, e := W.CY9jLhcY1f5(data)
							if e != nil {
								t.Fatal(e)
							}
							full := reassembly.Feed(part)
							if full == nil {
								continue
							}
							if full.D35ycLc5aB6r != 7 || full.XrDcBXcD != dest || !bytes.Equal(full.HmoFPk, payload) {
								t.Fatal("UDP response")
							}
							break
						}
					}
				case <-ctx.Done():
					t.Fatal("UDP accept timeout")
				}
				stopUser()
				select {
				case <-conn.Context().Done():
				case <-ctx.Done():
					t.Fatal("user cancellation did not close QUIC")
				}
			})
		}
	}
}
