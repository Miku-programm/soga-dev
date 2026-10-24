package MFFa9yT

import (
	"bytes"
	"context"
	"crypto/ecdh"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"net"
	CA "recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/LFWTTXF7JR"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	R "recovered.local/soga/reality"
	"testing"
	"time"
)

// Prevent the independent client's outer TLS implementation from reading ahead
// into raw inner records after the Vision transition. The server's real buffer
// drain behavior is unchanged and exercised against a real TCP socket.
type visionRecordBoundaryConn struct {
	net.Conn
	header    []byte
	remaining int
}

func (c *visionRecordBoundaryConn) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(c.header) == 0 && c.remaining == 0 {
		h := make([]byte, 5)
		if _, e := io.ReadFull(c.Conn, h); e != nil {
			return 0, e
		}
		c.remaining = int(binary.BigEndian.Uint16(h[3:]))
		c.header = h
	}
	if len(c.header) > 0 {
		n := copy(p, c.header)
		c.header = c.header[n:]
		return n, nil
	}
	n, e := c.Conn.Read(p[:min(len(p), c.remaining)])
	c.remaining -= n
	return n, e
}

type independentVisionConn struct {
	net.Conn
	outer                     net.Conn
	id                        uuid.UUID
	firstWrite, firstRead     bool
	writeDirect, readDirect   bool
	pending                   []byte
	directWritten, directRead int
}

func (c *independentVisionConn) Write(p []byte) (int, error) {
	if c.writeDirect {
		n, e := c.Conn.Write(p)
		c.directWritten += n
		return n, e
	}
	command := byte(0)
	for at := 0; at+5 <= len(p); {
		n := int(binary.BigEndian.Uint16(p[at+3 : at+5]))
		if p[at] == 23 && p[at+1] == 3 && p[at+2] == 3 {
			command = 2
			break
		}
		if at+5+n > len(p) {
			break
		}
		at += 5 + n
	}
	var frame []byte
	if !c.firstWrite {
		frame = append(frame, c.id[:]...)
		c.firstWrite = true
	}
	frame = append(frame, command)
	frame = binary.BigEndian.AppendUint16(frame, uint16(len(p)))
	frame = append(frame, 0, 0)
	frame = append(frame, p...)
	if _, e := c.outer.Write(frame); e != nil {
		return 0, e
	}
	if command == 2 {
		c.writeDirect = true
	}
	return len(p), nil
}
func (c *independentVisionConn) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(c.pending) > 0 {
		n := copy(p, c.pending)
		c.pending = c.pending[n:]
		return n, nil
	}
	if c.readDirect {
		n, e := c.Conn.Read(p)
		c.directRead += n
		return n, e
	}
	if !c.firstRead {
		head := make([]byte, 18)
		if _, e := io.ReadFull(c.outer, head); e != nil {
			return 0, e
		}
		if !bytes.Equal(head[:2], []byte{0, 0}) || !bytes.Equal(head[2:], c.id[:]) {
			return 0, fmt.Errorf("bad VLESS/Vision prefix")
		}
		c.firstRead = true
	}
	for {
		var h [5]byte
		if _, e := io.ReadFull(c.outer, h[:]); e != nil {
			return 0, e
		}
		n, pad := int(binary.BigEndian.Uint16(h[1:3])), int(binary.BigEndian.Uint16(h[3:]))
		payload := make([]byte, n+pad)
		if _, e := io.ReadFull(c.outer, payload); e != nil {
			return 0, e
		}
		if h[0] == 1 {
			return 0, fmt.Errorf("server ended padding without direct switch")
		}
		if h[0] > 2 {
			return 0, fmt.Errorf("bad Vision command %d", h[0])
		}
		if h[0] == 2 {
			c.readDirect = true
		}
		c.pending = payload[:n]
		if n == 0 {
			if c.readDirect {
				return c.Read(p)
			}
			continue
		}
		n = copy(p, c.pending)
		c.pending = c.pending[n:]
		return n, nil
	}
}
func TestNativeVLESSVisionInnerTLSRawSplice(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	cert, pemKey := hysteriaFixtureCertificate(t)
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM([]byte(cert))
	pair, e := tls.X509KeyPair([]byte(cert), []byte(pemKey))
	if e != nil {
		t.Fatal(e)
	}
	for _, security := range []string{"tls", "reality"} {
		t.Run(security, func(t *testing.T) {
			node := &L.Wr5SVP{EMlXEc: 32, I3k4JnNp7nB: cert, XATiVun: pemKey, Ck6DxjB: L.VkkFLH{XmIaPj3CzOF: "tcp", WAdJ4ftdZFk: L.D7FHIoaa8s4a(security), NMzMUQ2: "xtls-rprx-vision"}}
			var private *ecdh.PrivateKey
			var dest string
			if security == "reality" {
				private, _ = ecdh.X25519().GenerateKey(rand.Reader)
				dest = nativeRealityDestination(t)
				node.Ck6DxjB.BANVIZ72wh = []string{"hysteria"}
				node.Ck6DxjB.CrS6C1 = base64.RawURLEncoding.EncodeToString(private.Bytes())
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
				key := dest + " hysteria 2"
				defer R.GlobalPostHandshakeRecordsLens.Delete(key)
				until := time.Now().Add(4 * time.Second)
				for {
					if val, ok := R.GlobalPostHandshakeRecordsLens.Load(key); ok {
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
			target, e := net.Listen("tcp4", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer target.Close()
			done := make(chan error, 1)
			payloads := [][]byte{bytes.Repeat([]byte("first TLS encrypted target packet/"), 300), bytes.Repeat([]byte("second packet must be raw/"), 1900)}
			go func() {
				raw, e := target.Accept()
				if e != nil {
					done <- e
					return
				}
				defer raw.Close()
				raw.SetDeadline(time.Now().Add(8 * time.Second))
				s := tls.Server(raw, &tls.Config{Certificates: []tls.Certificate{pair}, MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13, SessionTicketsDisabled: true, CurvePreferences: []tls.CurveID{tls.X25519}})
				if e = s.Handshake(); e != nil {
					done <- e
					return
				}
				for _, want := range payloads {
					got := make([]byte, len(want))
					if _, e = io.ReadFull(s, got); e != nil {
						done <- e
						return
					}
					if !bytes.Equal(got, want) {
						done <- fmt.Errorf("inner TLS target content differs")
						return
					}
					if _, e = s.Write(got); e != nil {
						done <- e
						return
					}
				}
				done <- nil
			}()
			raw, e := net.DialTimeout("tcp", v.kX5i8U[0].Addr().String(), time.Second)
			if e != nil {
				t.Fatal(e)
			}
			defer raw.Close()
			raw.SetDeadline(time.Now().Add(8 * time.Second))
			boundary := &visionRecordBoundaryConn{Conn: raw}
			var outer net.Conn
			if security == "reality" {
				outer = nativeRealityOnConn(t, boundary, private)
			} else {
				tlsConn := tls.Client(boundary, &tls.Config{RootCAs: roots, ServerName: "hysteria", MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13})
				if e = tlsConn.Handshake(); e != nil {
					t.Fatal(e)
				}
				outer = tlsConn
			}
			id := uuid.FromStringOrNil("11111111-1111-4111-8111-111111111111")
			header := append([]byte{0}, id[:]...)
			addons := append([]byte{0x0a, 16}, []byte("xtls-rprx-vision")...)
			header = append(header, byte(len(addons)))
			header = append(header, addons...)
			header = append(header, 1)
			header = binary.BigEndian.AppendUint16(header, uint16(target.Addr().(*net.TCPAddr).Port))
			header = append(header, 1, 127, 0, 0, 1)
			if _, e = outer.Write(header); e != nil {
				t.Fatal(e)
			}
			transport := &independentVisionConn{Conn: raw, outer: outer, id: id}
			inner := tls.Client(transport, &tls.Config{RootCAs: roots, ServerName: "hysteria", MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13, CurvePreferences: []tls.CurveID{tls.X25519}})
			if e = inner.Handshake(); e != nil {
				t.Fatal("inner TLS handshake", e)
			}
			for _, payload := range payloads {
				if _, e = inner.Write(payload); e != nil {
					t.Fatal(e)
				}
				got := make([]byte, len(payload))
				if _, e = io.ReadFull(inner, got); e != nil {
					t.Fatal(e)
				}
				if !bytes.Equal(got, payload) {
					t.Fatal("inner target echo differs")
				}
			}
			if !transport.writeDirect || !transport.readDirect || transport.directWritten < len(payloads[1]) || transport.directRead < len(payloads[1]) {
				t.Fatalf("raw splice not proven write=%t/%d read=%t/%d", transport.writeDirect, transport.directWritten, transport.readDirect, transport.directRead)
			}
			if e = <-done; e != nil {
				t.Fatal(e)
			}
		})
	}
}
