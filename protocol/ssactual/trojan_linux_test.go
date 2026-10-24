//go:build linux

package ssactual

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	E "recovered.local/soga/net/outbound"
	O "recovered.local/soga/net/socket"
	"strconv"
	"testing"
	"time"
	"unsafe"
)

type trojanWSTestStream struct {
	c *websocket.Conn
	r io.Reader
}

func (s *trojanWSTestStream) Read(p []byte) (int, error) {
	for {
		if s.r == nil {
			_, r, e := s.c.NextReader()
			if e != nil {
				return 0, e
			}
			s.r = r
		}
		n, e := s.r.Read(p)
		if e == io.EOF {
			s.r = nil
			if n == 0 {
				continue
			}
			e = nil
		}
		return n, e
	}
}
func (s *trojanWSTestStream) Write(p []byte) (int, error) {
	e := s.c.WriteMessage(2, p)
	if e != nil {
		return 0, e
	}
	return len(p), nil
}

type trojanGRPCTestStream struct {
	r    *bufio.Reader
	w    http.ResponseWriter
	left uint64
}

func (s *trojanGRPCTestStream) Read(p []byte) (int, error) {
	if s.left == 0 {
		if _, e := s.r.Discard(6); e != nil {
			return 0, e
		}
		n, e := binary.ReadUvarint(s.r)
		if e != nil {
			return 0, e
		}
		s.left = n
	}
	n, e := io.ReadFull(s.r, p[:min(uint64(len(p)), s.left)])
	s.left -= uint64(n)
	return n, e
}
func (s *trojanGRPCTestStream) Write(p []byte) (int, error) {
	var pb [11]byte
	pb[0] = 10
	n := 1 + binary.PutUvarint(pb[1:], uint64(len(p)))
	frame := make([]byte, 5)
	binary.BigEndian.PutUint32(frame[1:], uint32(n+len(p)))
	frame = append(frame, pb[:n]...)
	frame = append(frame, p...)
	_, e := s.w.Write(frame)
	s.w.(http.Flusher).Flush()
	return len(p), e
}
func TestTrojanLayoutsAndConstructor(t *testing.T) {
	if unsafe.Sizeof(NLN_l6{}) != 48 || unsafe.Sizeof(ZxrFcPHYpol{}) != 208 || unsafe.Sizeof(HrHLa45bjuV{}) != 48 || unsafe.Sizeof(THKLdNRQ{}) != 16 {
		t.Fatal("RTTI layout")
	}
	o := ZxrFcPHYpol{AKe9rX4WI: "trojan", HF6Vma: "::1", GRc9GM8NFW: 443, Anl_Yj: "pw", IGRMFmQSwT: "sni.example", TNtpmYwuy: true, XIgNO9dyP: "grpc", R1aqZuYO: true}
	a, err := RFqiXSepf_9(o)
	if err != nil {
		t.Fatal(err)
	}
	if a.Addr() != "[::1]:443" || a.Type() != 8 || !a.SupportUDP() || a.oVYKaxP.ServerName != "sni.example" || a.oVYKaxP.MinVersion != tls.VersionTLS12 || a.oVYKaxP.NextProtos != nil || a.eXjtOT.EbklgbD != "sni.example" {
		t.Fatal(a)
	}
	o.Anl_Yj = "changed"
	if a.sWg8Yg.Anl_Yj != "pw" {
		t.Fatal("option struct not copied")
	}
}
func TestTrojanRealTransportTCPAndUDP(t *testing.T) {
	for _, transport := range []string{"tls", "ws", "grpc", "grpc-explicit"} {
		for _, udp := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/udp=%v", transport, udp), func(t *testing.T) {
				// The original grpc+explicit UDP path selects ordinary TLS.
				servedTransport := transport
				if transport == "grpc-explicit" {
					servedTransport = "grpc"
					if udp {
						servedTransport = "tls"
					}
				}
				reports := make(chan error, 2)
				expectedAddress := []byte{1, 192, 0, 2, 7, 0, 80}
				serve := func(stream io.ReadWriter) {
					header := make([]byte, 68)
					if _, e := io.ReadFull(stream, header); e != nil {
						reports <- e
						return
					}
					digest := sha256.Sum224([]byte("secret"))
					want := append([]byte(hex.EncodeToString(digest[:])+"\r\n"), byte(1))
					if udp {
						want[len(want)-1] = 3
					}
					want = append(want, expectedAddress...)
					want = append(want, '\r', '\n')
					if !bytes.Equal(header, want) {
						reports <- fmt.Errorf("header %x != %x", header, want)
						return
					}
					reports <- nil
					io.Copy(stream, stream)
				}
				server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if servedTransport == "ws" {
						if r.URL.Path != "/tunnel" || r.Header.Get("X-Proof") != "yes" {
							reports <- fmt.Errorf("websocket request %s", r.URL)
							return
						}
						c, e := (&websocket.Upgrader{}).Upgrade(w, r, nil)
						if e != nil {
							reports <- e
							return
						}
						defer c.Close()
						serve(&trojanWSTestStream{c: c})
					} else {
						if r.ProtoMajor != 2 || r.URL.Path != "/custom/Tun" {
							reports <- fmt.Errorf("grpc request %s %s", r.Proto, r.URL)
							return
						}
						w.WriteHeader(200)
						w.(http.Flusher).Flush()
						serve(&trojanGRPCTestStream{r: bufio.NewReader(r.Body), w: w})
					}
				}))
				server.EnableHTTP2 = servedTransport == "grpc"
				server.StartTLS()
				defer server.Close()
				address := server.Listener.Addr().String()
				if servedTransport == "tls" {
					listener, e := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: server.TLS.Certificates})
					if e != nil {
						t.Fatal(e)
					}
					defer listener.Close()
					address = listener.Addr().String()
					go func() {
						c, e := listener.Accept()
						if e != nil {
							reports <- e
							return
						}
						defer c.Close()
						c.SetDeadline(time.Now().Add(3 * time.Second))
						serve(c)
					}()
				}
				host, port, _ := net.SplitHostPort(address)
				number, _ := strconv.Atoi(port)
				network := transport
				if transport == "tls" {
					network = ""
				}
				if transport == "grpc-explicit" {
					network = "grpc"
				}
				adapter, e := RFqiXSepf_9(ZxrFcPHYpol{AKe9rX4WI: "test", HF6Vma: host, GRc9GM8NFW: number, Anl_Yj: "secret", R1aqZuYO: true, TNtpmYwuy: true, XIgNO9dyP: network, ZYpNOKrxXx: THKLdNRQ{IDDJaTZ: "custom"}, JXjd1Bs: HrHLa45bjuV{TTsHbH5qb: "/tunnel", RuSpNqlE: map[string]string{"X-Proof": "yes"}}})
				if e != nil {
					t.Fatal(e)
				}
				if adapter.vuo5piSzgld != nil {
					defer adapter.vuo5piSzgld.CloseIdleConnections()
				}
				m := &E.D0mLgQF{ZouEpmjw8: net.IPv4(192, 0, 2, 7), JRXMVvg: 80}
				var options []O.PNSpXRb
				if transport == "grpc-explicit" {
					options = []O.PNSpXRb{O.Fum1PEvMPct(0)}
				}
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()
				if udp {
					conn, e := adapter.ListenPacketContext(ctx, m, options...)
					if e != nil {
						t.Fatal(e)
					}
					defer conn.Close()
					conn.SetDeadline(time.Now().Add(3 * time.Second))
					payload := []byte("udp payload")
					n, e := conn.WriteTo(payload, &net.UDPAddr{IP: net.IPv4(192, 0, 2, 7), Port: 80})
					if e != nil || n != len(payload)+11 {
						t.Fatalf("UDP write %d %v", n, e)
					}
					b := make([]byte, 300)
					n, addr, e := conn.ReadFrom(b)
					if e != nil || string(b[:n]) != string(payload) || addr.String() != "192.0.2.7:80" {
						t.Fatalf("UDP read %d %v %v", n, addr, e)
					}
				} else {
					conn, e := adapter.DialContext(ctx, m, options...)
					if e != nil {
						t.Fatal(e)
					}
					defer conn.Close()
					conn.SetDeadline(time.Now().Add(3 * time.Second))
					payload := []byte("tcp payload")
					if _, e := conn.Write(payload); e != nil {
						t.Fatal(e)
					}
					reply := make([]byte, len(payload))
					if _, e := io.ReadFull(conn, reply); e != nil || !bytes.Equal(reply, payload) {
						t.Fatalf("TCP read %q %v", reply, e)
					}
				}
				select {
				case e := <-reports:
					if e != nil {
						t.Fatal(e)
					}
				case <-time.After(3 * time.Second):
					t.Fatal("missing header")
				}
			})
		}
	}
}
