//go:build linux

package LCe4Wkk2

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/chacha20poly1305"
	"hash"
	"hash/crc32"
	"hash/fnv"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	E "recovered.local/soga/E3F12urwswL"
	O "recovered.local/soga/vP_e5OQAve"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func vmessPeerKDF(key []byte, labels ...string) []byte {
	var build func(int) hash.Hash
	build = func(n int) hash.Hash {
		if n == 0 {
			return hmac.New(sha256.New, []byte("VMess AEAD KDF"))
		}
		return hmac.New(func() hash.Hash { return build(n - 1) }, []byte(labels[n-1]))
	}
	h := build(len(labels))
	h.Write(key)
	return h.Sum(nil)
}
func vmessPeerGCM(key []byte) cipher.AEAD {
	b, e := aes.NewCipher(key)
	if e != nil {
		panic(e)
	}
	a, e := cipher.NewGCM(b)
	if e != nil {
		panic(e)
	}
	return a
}
func vmessPeerRead(r io.Reader, n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(r, b)
	return b, err
}

// This vmessPeer is separate test code which parses the wire, authenticates requests,
// and seals responses. It does not use recovered frames or handshake methods.
func vmessPeerReadHeader(c io.ReadWriter, commandKey, uuid []byte, aead bool) ([]byte, error) {
	auth, e := vmessPeerRead(c, 16)
	if e != nil {
		return nil, e
	}
	if aead {
		b, _ := aes.NewCipher(vmessPeerKDF(commandKey, "AES Auth ID Encryption")[:16])
		plain := make([]byte, 16)
		b.Decrypt(plain, auth)
		if binary.BigEndian.Uint32(plain[12:]) != crc32.ChecksumIEEE(plain[:12]) {
			return nil, fmt.Errorf("auth CRC")
		}
		encLen, e := vmessPeerRead(c, 18)
		if e != nil {
			return nil, e
		}
		nonce, e := vmessPeerRead(c, 8)
		if e != nil {
			return nil, e
		}
		key := vmessPeerKDF(commandKey, "VMess Header AEAD Key_Length", string(auth), string(nonce))
		iv := vmessPeerKDF(commandKey, "VMess Header AEAD Nonce_Length", string(auth), string(nonce))
		length, e := vmessPeerGCM(key[:16]).Open(nil, iv[:12], encLen, auth)
		if e != nil {
			return nil, e
		}
		enc, e := vmessPeerRead(c, int(binary.BigEndian.Uint16(length))+16)
		if e != nil {
			return nil, e
		}
		key = vmessPeerKDF(commandKey, "VMess Header AEAD Key", string(auth), string(nonce))
		iv = vmessPeerKDF(commandKey, "VMess Header AEAD Nonce", string(auth), string(nonce))
		return vmessPeerGCM(key[:16]).Open(nil, iv[:12], enc, auth)
	}
	var timestamp int64
	now := time.Now().Unix()
	for stamp := now - 10; stamp <= now+10; stamp++ {
		h := hmac.New(md5.New, uuid)
		binary.Write(h, binary.BigEndian, uint64(stamp))
		if hmac.Equal(h.Sum(nil), auth) {
			timestamp = stamp
			break
		}
	}
	if timestamp == 0 {
		return nil, fmt.Errorf("legacy auth")
	}
	var stamp [8]byte
	binary.BigEndian.PutUint64(stamp[:], uint64(timestamp))
	iv := md5.Sum(bytes.Repeat(stamp[:], 4))
	block, _ := aes.NewCipher(commandKey)
	stream := cipher.NewCFBDecrypter(block, iv[:])
	first, e := vmessPeerRead(c, 41)
	if e != nil {
		return nil, e
	}
	stream.XORKeyStream(first, first)
	if first[40] != 1 {
		return nil, fmt.Errorf("expected IPv4")
	}
	tail, e := vmessPeerRead(c, 4+int(first[35]>>4)+4)
	if e != nil {
		return nil, e
	}
	stream.XORKeyStream(tail, tail)
	return append(first, tail...), nil
}
func vmessPeerBodyAEAD(key []byte, security byte) cipher.AEAD {
	if security == 3 {
		return vmessPeerGCM(key)
	}
	a := md5.Sum(key)
	b := md5.Sum(a[:])
	joined := append(a[:], b[:]...)
	out, e := chacha20poly1305.New(joined)
	if e != nil {
		panic(e)
	}
	return out
}
func vmessPeerResponseHeader(key, iv []byte, auth byte, aead bool) []byte {
	plain := []byte{auth, 0, 0, 0}
	if !aead {
		block, _ := aes.NewCipher(key)
		out := make([]byte, 4)
		cipher.NewCFBEncrypter(block, iv).XORKeyStream(out, plain)
		return out
	}
	lk := vmessPeerKDF(key, "AEAD Resp Header Len Key")
	ln := vmessPeerKDF(iv, "AEAD Resp Header Len IV")
	out := vmessPeerGCM(lk[:16]).Seal(nil, ln[:12], []byte{0, 4}, nil)
	pk := vmessPeerKDF(key, "AEAD Resp Header Key")
	pn := vmessPeerKDF(iv, "AEAD Resp Header IV")
	return append(out, vmessPeerGCM(pk[:16]).Seal(nil, pn[:12], plain, nil)...)
}
func vmessPeerEcho(c io.ReadWriter, commandKey, uuid []byte, aead, udp bool, payload []byte) error {
	header, e := vmessPeerReadHeader(c, commandKey, uuid, aead)
	if e != nil {
		return e
	}
	if len(header) < 49 {
		return fmt.Errorf("short header")
	}
	f := fnv.New32a()
	f.Write(header[:len(header)-4])
	if f.Sum32() != binary.BigEndian.Uint32(header[len(header)-4:]) {
		return fmt.Errorf("request FNV")
	}
	command := byte(1)
	if udp {
		command = 2
	}
	if header[0] != 1 || header[36] != 0 || header[37] != command || binary.BigEndian.Uint16(header[38:40]) != 443 || !bytes.Equal(header[41:45], []byte{127, 0, 0, 1}) {
		return fmt.Errorf("request fields %x", header[:45])
	}
	iv, key := header[1:17], header[17:33]
	var responseKey, responseIV []byte
	if aead {
		k := sha256.Sum256(key)
		v := sha256.Sum256(iv)
		responseKey = k[:16]
		responseIV = v[:16]
	} else {
		k := md5.Sum(key)
		v := md5.Sum(iv)
		responseKey = k[:]
		responseIV = v[:]
	}
	options, security := header[34], header[35]&15
	var body []byte
	if options == 0 {
		body, e = vmessPeerRead(c, len(payload))
	} else {
		length, err := vmessPeerRead(c, 2)
		if err != nil {
			return err
		}
		body, e = vmessPeerRead(c, int(binary.BigEndian.Uint16(length)))
		if e == nil && security != 5 {
			nonce := append([]byte{0, 0}, iv[2:12]...)
			body, e = vmessPeerBodyAEAD(key, security).Open(nil, nonce, body, nil)
		}
	}
	if e != nil {
		return e
	}
	if !bytes.Equal(body, payload) {
		return fmt.Errorf("request body %x", body)
	}
	out := vmessPeerResponseHeader(responseKey, responseIV, header[33], aead)
	reply := append([]byte("echo:"), body...)
	if options == 0 {
		out = append(out, reply...)
	} else {
		if security != 5 {
			nonce := append([]byte{0, 0}, responseIV[2:12]...)
			reply = vmessPeerBodyAEAD(responseKey, security).Seal(nil, nonce, reply, nil)
		}
		out = append(out, byte(len(reply)>>8), byte(len(reply)))
		out = append(out, reply...)
	}
	_, e = c.Write(out)
	return e
}

type vmessDuplex struct {
	io.Reader
	io.Writer
}
type vmessHTTPWriter struct{ http.ResponseWriter }

func (w vmessHTTPWriter) Write(p []byte) (int, error) {
	n, e := w.ResponseWriter.Write(p)
	w.ResponseWriter.(http.Flusher).Flush()
	return n, e
}
func TestVMessAdapterAllTransports(t *testing.T) {
	for _, transport := range []string{"tcp", "tls", "http", "http-tls", "ws", "ws-tls", "h2", "grpc", "grpc-explicit"} {
		for _, udp := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/udp=%v", transport, udp), func(t *testing.T) {
				reports := make(chan error, 2)
				payload := []byte("independent transport peer")
				uuid := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
				commandKey := md5.Sum(append(append([]byte(nil), uuid...), []byte("c48619fe-8f02-49e0-b9e9-edf763e17e21")...))
				serve := func(s io.ReadWriter) { reports <- vmessPeerEcho(s, commandKey[:], uuid, true, udp, payload) }
				server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					switch transport {
					case "ws", "ws-tls":
						if r.URL.Path != "/tunnel" || r.Header.Get("X-Proof") != "yes" {
							reports <- fmt.Errorf("ws request %v", r.URL)
							return
						}
						c, e := (&websocket.Upgrader{}).Upgrade(w, r, nil)
						if e != nil {
							reports <- e
							return
						}
						defer c.Close()
						serve(&trojanWSTestStream{c: c})
					case "h2":
						if r.ProtoMajor != 2 || r.Method != "PUT" || r.URL.Path != "/h2" || r.Host != "h2.example" || r.Header.Get("Accept-Encoding") != "identity" {
							reports <- fmt.Errorf("h2 request %s %s %v %s %v", r.Proto, r.Method, r.URL, r.Host, r.Header)
							return
						}
						w.WriteHeader(200)
						w.(http.Flusher).Flush()
						serve(vmessDuplex{r.Body, vmessHTTPWriter{w}})
					default:
						if r.URL.Path != "/custom/Tun" {
							reports <- fmt.Errorf("grpc request %v", r.URL)
							return
						}
						w.WriteHeader(200)
						w.(http.Flusher).Flush()
						serve(&trojanGRPCTestStream{r: bufio.NewReader(r.Body), w: w})
					}
				}))
				server.EnableHTTP2 = transport == "h2" || strings.HasPrefix(transport, "grpc")
				server.StartTLS()
				defer server.Close()
				address := server.Listener.Addr().String()
				if transport == "tcp" || transport == "tls" || strings.HasPrefix(transport, "http") || transport == "ws" {
					listener, e := net.Listen("tcp", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					defer listener.Close()
					address = listener.Addr().String()
					if transport == "ws" {
						plain := &http.Server{Handler: server.Config.Handler}
						go plain.Serve(listener)
						defer plain.Close()
					} else {
						if transport == "tls" || transport == "http-tls" {
							listener = tls.NewListener(listener, &tls.Config{Certificates: server.TLS.Certificates})
						}
						go func() {
							c, e := listener.Accept()
							if e != nil {
								reports <- e
								return
							}
							defer c.Close()
							c.SetDeadline(time.Now().Add(5 * time.Second))
							if strings.HasPrefix(transport, "http") {
								r := bufio.NewReader(c)
								request, e := http.ReadRequest(r)
								if e != nil {
									reports <- e
									return
								}
								if request.Method != "POST" || request.URL.Path != "/cover" || request.Host != "cover.example" || request.Header.Get("X-Proof") != "yes" {
									reports <- fmt.Errorf("http request %v", request)
									return
								}
								c.Write([]byte("IGNORED STATUS LINE\r\nX-Test: yes\r\n\r\n"))
								serve(vmessDuplex{io.MultiReader(request.Body, r), c})
							} else {
								serve(c)
							}
						}()
					}
				}
				host, port, _ := net.SplitHostPort(address)
				number, _ := strconv.Atoi(port)
				network := transport
				switch transport {
				case "tcp", "tls":
					network = ""
				case "http-tls":
					network = "http"
				case "ws-tls":
					network = "ws"
				case "grpc-explicit":
					network = "grpc"
				}
				a, e := Fr9G0QQhtaac(YaiP1Z10VuOK{NwFIZN: "test", PaWTH6: host, KQ7hyM: number, MMeZdiHjQ: "00010203-0405-0607-0809-0a0b0c0d0e0f", YKnBs48P7XY: "AES-128-GCM", DS2NjE9s_: true, EtDA3oIPdZ7y: network, MLU1J_: transport != "tcp" && transport != "http" && transport != "ws", VPGpSFLRW: true, KphAInyo: EyDONQ3H9Jz{UmSoRT: "POST", Wt0Mo6: []string{"/cover"}, MpPwkigf3: map[string][]string{"Host": {"cover.example"}, "X-Proof": {"yes"}}}, GMSZn2d0QZm: OPYb56D{MrTNulEgJrCT: []string{"h2.example"}, WvC0eaw: "/h2"}, FN179GVH_: THKLdNRQ{IDDJaTZ: "custom"}, U7nq5HsQY: HrHLa45bjuV{TTsHbH5qb: "/tunnel", RuSpNqlE: map[string]string{"X-Proof": "yes"}}})
				if e != nil {
					t.Fatal(e)
				}
				if a.nmGdkI != nil {
					defer a.nmGdkI.CloseIdleConnections()
				}
				m := &E.D0mLgQF{ZouEpmjw8: net.IPv4(127, 0, 0, 1), JRXMVvg: 443}
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				var options []O.PNSpXRb
				if transport == "grpc-explicit" {
					options = []O.PNSpXRb{O.Fum1PEvMPct(0)}
				}
				if udp {
					m.AcXafZ1 = 1
					c, e := a.ListenPacketContext(ctx, m, options...)
					if e != nil {
						t.Fatal(e)
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(5 * time.Second))
					addr := &net.UDPAddr{IP: m.ZouEpmjw8, Port: 443}
					if n, e := c.WriteTo(payload, addr); n != len(payload) || e != nil {
						t.Fatal(n, e)
					}
					buf := make([]byte, 100)
					n, got, e := c.ReadFrom(buf)
					if e != nil || string(buf[:n]) != "echo:"+string(payload) || got.String() != addr.String() {
						t.Fatal(n, got, e, string(buf[:n]))
					}
				} else {
					c, e := a.DialContext(ctx, m, options...)
					if e != nil {
						t.Fatal(e)
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(5 * time.Second))
					if n, e := c.Write(payload); n != len(payload) || e != nil {
						t.Fatal(n, e)
					}
					buf := make([]byte, len(payload)+5)
					if _, e := io.ReadFull(c, buf); e != nil || string(buf) != "echo:"+string(payload) {
						t.Fatal(e, string(buf))
					}
				}
				select {
				case e := <-reports:
					if e != nil {
						t.Fatal(e)
					}
				case <-time.After(5 * time.Second):
					t.Fatal("peer timeout")
				}
			})
		}
	}
}
func TestVMessAdapterConstruction(t *testing.T) {
	if unsafe.Sizeof(YaiP1Z10VuOK{}) != 304 || unsafe.Sizeof(UpVCB2mE{}) != 48 || unsafe.Sizeof(keoydRNdlJ{}) != 32 || unsafe.Sizeof(EyDONQ3H9Jz{}) != 48 || unsafe.Sizeof(OPYb56D{}) != 40 {
		t.Fatal("RTTI")
	}
	o := YaiP1Z10VuOK{NwFIZN: "node", PaWTH6: "::1", KQ7hyM: 443, MMeZdiHjQ: "00010203-0405-0607-0809-0a0b0c0d0e0f", YKnBs48P7XY: "AUTO", EtDA3oIPdZ7y: "h2", MLU1J_: true, DS2NjE9s_: true}
	a, e := Fr9G0QQhtaac(o)
	if e != nil {
		t.Fatal(e)
	}
	if a.Addr() != "[::1]:443" || a.Type() != 7 || !a.SupportUDP() || a.vL1N8AU_5.GMSZn2d0QZm.MrTNulEgJrCT[0] != "www.example.com" || len(o.GMSZn2d0QZm.MrTNulEgJrCT) != 0 {
		t.Fatal(a)
	}
	o.MLU1J_ = false
	if _, e := Fr9G0QQhtaac(o); e == nil || e.Error() != "TLS must be true with h2/grpc network" {
		t.Fatal(e)
	}
	o.YKnBs48P7XY = "bad"
	if _, e := Fr9G0QQhtaac(o); e == nil || e.Error() != "unknown security type: bad" {
		t.Fatal(e)
	}
	o.YKnBs48P7XY = "none"
	o.MLU1J_ = true
	o.EtDA3oIPdZ7y = "grpc"
	a, e = Fr9G0QQhtaac(o)
	if e != nil {
		t.Fatal(e)
	}
	if a.o7IafTDLYDr.ServerName != "::1" || a.o7IafTDLYDr.MinVersion != 0 || a.o7IafTDLYDr.NextProtos != nil || a.nmGdkI.PingTimeout != 0 {
		t.Fatal("grpc defaults")
	}
}

type vmessPacketFixture struct {
	net.Conn
	bytes.Buffer
}

func (c *vmessPacketFixture) Read(p []byte) (int, error)  { return c.Buffer.Read(p) }
func (c *vmessPacketFixture) Write(p []byte) (int, error) { return c.Buffer.Write(p) }
func TestVMessAddressAndPacketSemantics(t *testing.T) {
	domain := strings.Repeat("d", 300)
	m := &E.D0mLgQF{WJ95LyOL: domain, AcXafZ1: 1, JRXMVvg: 53}
	v := eEgDc4(m)
	if v.OyEjzi1v93 != 2 || len(v.TIBiNxg) != 301 || v.TIBiNxg[0] != 44 || !v.GV3P5rY || v.GHP6ivev2Db != 53 {
		t.Fatal(v)
	}
	m.WJ95LyOL = ""
	m.ZouEpmjw8 = net.IP{1, 2}
	v = eEgDc4(m)
	if v.OyEjzi1v93 != 3 || !bytes.Equal(v.TIBiNxg, make([]byte, 16)) {
		t.Fatal(v)
	}
	f := new(vmessPacketFixture)
	addr := &net.UDPAddr{IP: net.ParseIP("fe80::1"), Port: 53, Zone: "original"}
	c := &keoydRNdlJ{Conn: f, oDRJBoi: addr}
	other := *addr
	other.Zone = "different"
	if n, e := c.WriteTo([]byte("ok"), &other); n != 2 || e != nil {
		t.Fatal(n, e)
	}
	other.Port++
	if n, e := c.WriteTo([]byte("drop"), &other); n != 0 || e != udpDestinationErrorAt5097920 {
		t.Fatal(n, e)
	}
	b := make([]byte, 5)
	if n, a, e := c.ReadFrom(b); n != 2 || a != addr || e != nil {
		t.Fatal(n, a, e)
	}
	if n, a, e := c.ReadFrom(b); n != 0 || a != addr || e != io.EOF {
		t.Fatal(n, a, e)
	}
	for _, invalid := range []net.Addr{&net.TCPAddr{}, (*net.UDPAddr)(nil)} {
		func() {
			defer func() {
				if recover() == nil {
					t.Error("original assertion/dereference must panic")
				}
			}()
			c.WriteTo(nil, invalid)
		}()
	}
}
