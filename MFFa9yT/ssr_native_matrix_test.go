//go:build linux

package MFFa9yT

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/cast5"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/salsa20"
	"io"
	"math/big"
	"net"
	"net/http"
	"os/exec"
	E "recovered.local/soga/E3F12urwswL"
	A "recovered.local/soga/LCe4Wkk2"
	L "recovered.local/soga/LFWTTXF7JR"
	F "recovered.local/soga/hfOuV8fOe4L"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"strings"
	"testing"
	"time"
)

// The client is the artifact's separate outbound implementation (LCe ->
// Gt/low/m3), not the native inbound Ta/tw/Roa codecs. Only intersecting
// registries are used: no fabricated auth_chain C-F outbound support.
func TestNativeSSRIndependentOutboundCipherProtocolMatrix(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, method := range []string{"none", "aes-128-cfb", "aes-192-cfb", "aes-256-cfb", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr", "rc4-md5", "chacha20-ietf"} {
		for _, protocol := range []string{"origin", "auth_aes128_md5", "auth_aes128_sha1", "auth_chain_a", "auth_chain_b"} {
			for _, udp := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%s/udp=%t", method, protocol, udp), func(t *testing.T) { ssrNativeMatrixExchange(t, method, protocol, "plain", udp) })
			}
		}
	}
}
func TestNativeSSRIndependentOutboundObfsProtocolMatrix(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, obfs := range []string{"http_simple", "http_post", "tls1.2_ticket_auth"} {
		for _, protocol := range []string{"origin", "auth_aes128_md5", "auth_aes128_sha1", "auth_chain_a", "auth_chain_b"} {
			t.Run(obfs+"/"+protocol, func(t *testing.T) { ssrNativeMatrixExchange(t, "aes-128-cfb", protocol, obfs, false) })
		}
	}
}
func ssrNativeMatrixCore(t *testing.T, method, protocol, obfs string) *Ze0fOaAko {
	t.Helper()
	data := &Q.D5gTIZ5aVg{H47AzZsYHB1N: &L.Wr5SVP{EMlXEc: 69}, VTme1SuHL6i_: L.Sxcf3wu3L{ssrPanelFixture(0, "none", "bootstrap-password")}}
	raw, err := x4oEZV(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: data})
	if err != nil {
		t.Fatal(err)
	}
	v := raw.(*Ze0fOaAko)
	t.Cleanup(v.Stop)
	if err = v.Start(); err != nil {
		t.Fatal(err)
	}
	old := v.mc0MXUU.sSKJ79jFt[0]
	user := ssrPanelFixture(0, method, "native-matrix-key")
	settings := user.JAfXWwzYrW.(*L.QcqpSL3i6ri)
	settings.H_y2xZht7S8 = F.Mu0vMIIs(protocol)
	settings.ZZIDsJ = F.PMorzD7u(obfs)
	data.VTme1SuHL6i_ = L.Sxcf3wu3L{user}
	if v.UpdateAPIData(data) {
		t.Fatal("unexpected node restart")
	}
	next := v.mc0MXUU.sSKJ79jFt[0]
	if next == nil || next == old || old.ot1kuknWf35.Err() != context.Canceled {
		t.Fatal("UpdateAPIData did not replace original port state")
	}
	if len(next.woqAbWjQH_A) != 1 || len(next.aaBpAQ44Zfs) != 1 {
		t.Fatal("updated listener pair absent")
	}
	return v
}
func ssrNativeMatrixExchange(t *testing.T, method, protocol, obfs string, udp bool) {
	t.Helper()
	v := ssrNativeMatrixCore(t, method, protocol, obfs)
	state := v.mc0MXUU.sSKJ79jFt[0]
	endpoint := state.woqAbWjQH_A[0].Addr().(*net.TCPAddr).Port
	if udp {
		endpoint = state.aaBpAQ44Zfs[0].LocalAddr().(*net.UDPAddr).Port
	}
	adapter, err := A.CR6bC5(A.ShadowsocksROptions{Name: "independent-ssr-wire", Server: "127.0.0.1", Port: endpoint, Password: "native-matrix-key", Cipher: method, Obfs: obfs, Protocol: protocol, UDP: true})
	if err != nil {
		t.Fatal(err)
	}
	targetPort, done := trojanMatrixTarget(t, udp)
	target := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: targetPort}
	metadata := &E.D0mLgQF{ZouEpmjw8: target.IP, JRXMVvg: E.AN18DDo(targetPort)}
	payload := bytes.Repeat([]byte("native-transport\x00"), 91)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if udp {
		client, err := adapter.ListenPacketContext(ctx, metadata)
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()
		client.SetDeadline(time.Now().Add(2 * time.Second))
		if _, err = client.WriteTo(payload, target); err != nil {
			t.Fatal(err)
		}
		got := make([]byte, 8192)
		n, from, err := client.ReadFrom(got)
		if err != nil || !bytes.Equal(got[:n], payload) || from == nil || from.String() != target.String() {
			t.Fatalf("UDP reply bytes=%d from=%v error=%v", n, from, err)
		}
	} else {
		client, err := adapter.DialContext(ctx, metadata)
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()
		client.SetDeadline(time.Now().Add(2 * time.Second))
		if _, err = client.Write(bytes.Clone(payload)); err != nil {
			t.Fatal(err)
		}
		got := make([]byte, len(payload))
		n, err := io.ReadFull(client, got)
		if err != nil || !bytes.Equal(got, payload) {
			first := 0
			for first < len(got) && got[first] == payload[first] {
				first++
			}
			t.Fatalf("TCP reply bytes=%d/%d error=%v first-diff=%d got=%x want=%x tail=%x", n, len(payload), err, first, got[:min(32, len(got))], payload[:min(32, len(payload))], got[max(0, len(got)-32):])
		}
	}
	select {
	case err := <-done:
		if err != nil {
			t.Fatal("target", err)
		}
	case <-ctx.Done():
		t.Fatal("target not reached")
	}
	user, ok := v.jc_QPN6.Get(42)
	if !ok {
		t.Fatal("native user missing")
	}
	up, down := user.GetAndClearTraffic()
	if down < int64(len(payload)) || (!udp && up < int64(len(payload))) {
		t.Fatalf("native user traffic %d/%d below payload %d", up, down, len(payload))
	}
	// Original per-port SSR UDP mode zero omits upload counting; do not impose
	// symmetry that is absent from the native packet callback.
	if udp && up != 0 {
		t.Fatalf("mode zero UDP upload changed: %d", up)
	}
}

// Independent primitives provide additional cipher wire coverage without using
// either native Ta or outbound m3 wrappers. Camellia and RC2 use system
// libcrypto; IDEA uses the separate mathematical reference below.
type ssrCipherFixtureSpec struct {
	name    string
	key, iv int
	aead    bool
}

var ssrIndependentCiphers = []ssrCipherFixtureSpec{
	{"none", 16, 0, false}, {"aes-128-cfb", 16, 16, false}, {"aes-192-cfb", 24, 16, false}, {"aes-256-cfb", 32, 16, false},
	{"aes-128-ctr", 16, 16, false}, {"aes-192-ctr", 24, 16, false}, {"aes-256-ctr", 32, 16, false},
	{"aes-128-ofb", 16, 16, false}, {"aes-192-ofb", 24, 16, false}, {"aes-256-ofb", 32, 16, false},
	{"des-cfb", 8, 8, false}, {"bf-cfb", 16, 8, false}, {"cast5-cfb", 16, 8, false},
	{"rc4", 16, 0, false}, {"rc4-md5", 16, 16, false}, {"rc4-md5-6", 16, 6, false},
	{"chacha20", 32, 8, false}, {"chacha20-ietf", 32, 12, false}, {"salsa20", 32, 8, false},
	{"camellia-128-cfb", 16, 16, false}, {"camellia-192-cfb", 24, 16, false}, {"camellia-256-cfb", 32, 16, false}, {"idea-cfb", 16, 8, false}, {"rc2-cfb", 16, 8, false},
	{"chacha20-ietf-poly1305", 32, 32, true}, {"aes-128-gcm", 16, 16, true}, {"aes-192-gcm", 24, 24, true}, {"aes-256-gcm", 32, 32, true},
}

func ssrIndependentKey(password string, size int) []byte {
	var key, last []byte
	for len(key) < size {
		h := md5.New()
		h.Write(last)
		h.Write([]byte(password))
		last = h.Sum(nil)
		key = append(key, last...)
	}
	return key[:size]
}
func ssrIndependentCrypt(t *testing.T, method string, key, iv, input []byte, decrypt bool) []byte {
	t.Helper()
	out := make([]byte, len(input))
	var stream cipher.Stream
	var block cipher.Block
	var err error
	switch {
	case strings.HasPrefix(method, "camellia-") || method == "rc2-cfb":
		return ssrIndependentLibcrypto(t, method, key, iv, input, decrypt)
	case method == "none":
		return bytes.Clone(input)
	case method == "salsa20":
		var k [32]byte
		copy(k[:], key)
		salsa20.XORKeyStream(out, input, iv, &k)
		return out
	case strings.HasPrefix(method, "chacha20"):
		nonce := iv
		if len(iv) == 8 {
			nonce = append(make([]byte, 4), iv...)
		}
		stream, err = chacha20.NewUnauthenticatedCipher(key, nonce)
	case strings.HasPrefix(method, "rc4"):
		if method != "rc4" {
			h := md5.New()
			h.Write(key)
			h.Write(iv)
			key = h.Sum(nil)
		}
		stream, err = rc4.NewCipher(key)
	default:
		switch {
		case strings.HasPrefix(method, "aes-"):
			block, err = aes.NewCipher(key)
		case method == "des-cfb":
			block, err = des.NewCipher(key)
		case method == "bf-cfb":
			block, err = blowfish.NewCipher(key)
		case method == "cast5-cfb":
			block, err = cast5.NewCipher(key)
		case method == "idea-cfb":
			block = newSSRIDEAReference(key)
		default:
			t.Fatal("unsupported independent primitive", method)
		}
		if err == nil {
			switch {
			case strings.HasSuffix(method, "-ctr"):
				stream = cipher.NewCTR(block, iv)
			case strings.HasSuffix(method, "-ofb"):
				stream = cipher.NewOFB(block, iv)
			case decrypt:
				stream = cipher.NewCFBDecrypter(block, iv)
			default:
				stream = cipher.NewCFBEncrypter(block, iv)
			}
		}
	}
	if err != nil {
		t.Fatal(err)
	}
	stream.XORKeyStream(out, input)
	return out
}
func ssrIndependentAEAD(t *testing.T, spec ssrCipherFixtureSpec, salt []byte) cipher.AEAD {
	t.Helper()
	sub := make([]byte, spec.key)
	if _, err := io.ReadFull(hkdf.New(sha1.New, ssrIndependentKey("native-matrix-key", spec.key), salt, []byte("ss-subkey")), sub); err != nil {
		t.Fatal(err)
	}
	var a cipher.AEAD
	var err error
	if spec.name == "chacha20-ietf-poly1305" {
		a, err = chacha20poly1305.New(sub)
	} else {
		var b cipher.Block
		b, err = aes.NewCipher(sub)
		if err == nil {
			a, err = cipher.NewGCM(b)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
	return a
}
func TestNativeSSRIndependentPrimitiveCipherMatrix(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, spec := range ssrIndependentCiphers {
		for _, udp := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/udp=%t", spec.name, udp), func(t *testing.T) {
				v := ssrNativeMatrixCore(t, spec.name, "origin", "plain")
				state := v.mc0MXUU.sSKJ79jFt[0]
				network, endpoint := "tcp", state.woqAbWjQH_A[0].Addr().String()
				if udp {
					network, endpoint = "udp", state.aaBpAQ44Zfs[0].LocalAddr().String()
				}
				targetPort, done := trojanMatrixTarget(t, udp)
				address := ssFixtureAddress(1, targetPort)
				payload := bytes.Repeat([]byte("native-transport\x00"), 91)
				plain := append(bytes.Clone(address), payload...)
				key := ssrIndependentKey("native-matrix-key", spec.key)
				iv := make([]byte, spec.iv)
				if _, err := rand.Read(iv); err != nil {
					t.Fatal(err)
				}
				wire := bytes.Clone(iv)
				if spec.aead {
					a := ssrIndependentAEAD(t, spec, iv)
					if udp {
						wire = append(wire, a.Seal(nil, make([]byte, 12), plain, nil)...)
					} else {
						wire = append(wire, ssClientFrame(a, make([]byte, 12), plain)...)
					}
				} else {
					wire = append(wire, ssrIndependentCrypt(t, spec.name, key, iv, plain, false)...)
				}
				client, err := net.Dial(network, endpoint)
				if err != nil {
					t.Fatal(err)
				}
				defer client.Close()
				client.SetDeadline(time.Now().Add(2 * time.Second))
				if _, err = client.Write(wire); err != nil {
					t.Fatal(err)
				}
				var got []byte
				if udp {
					buf := make([]byte, 8192)
					n, err := client.Read(buf)
					if err != nil {
						t.Fatal(err)
					}
					if n < spec.iv {
						t.Fatal("short UDP IV")
					}
					if spec.aead {
						got, err = ssrIndependentAEAD(t, spec, buf[:spec.iv]).Open(nil, make([]byte, 12), buf[spec.iv:n], nil)
						if err != nil {
							t.Fatal(err)
						}
					} else {
						got = ssrIndependentCrypt(t, spec.name, key, buf[:spec.iv], buf[spec.iv:n], true)
					}
					if !bytes.Equal(got, plain) {
						t.Fatalf("UDP plaintext mismatch: %d", len(got))
					}
				} else {
					readIV := make([]byte, spec.iv)
					if _, err = io.ReadFull(client, readIV); err != nil {
						t.Fatal(err)
					}
					if spec.aead {
						a := ssrIndependentAEAD(t, spec, readIV)
						nonce := make([]byte, 12)
						for len(got) < len(payload) {
							p, err := ssReadClientFrame(client, a, nonce)
							if err != nil {
								t.Fatal(err)
							}
							got = append(got, p...)
						}
					} else {
						encrypted := make([]byte, len(payload))
						if _, err = io.ReadFull(client, encrypted); err != nil {
							t.Fatal(err)
						}
						got = ssrIndependentCrypt(t, spec.name, key, readIV, encrypted, true)
					}
					if !bytes.Equal(got, payload) {
						t.Fatalf("TCP plaintext mismatch: %d", len(got))
					}
				}
				if err := <-done; err != nil {
					t.Fatal("actual target", err)
				}
			})
		}
	}
}
func TestNativeSSRIndependentSimpleObfsWire(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, obfs := range []string{"none", "simple_obfs_http", "simple_obfs_tls"} {
		t.Run(obfs, func(t *testing.T) {
			v := ssrNativeMatrixCore(t, "none", "origin", obfs)
			targetPort, done := trojanMatrixTarget(t, false)
			payload := bytes.Repeat([]byte("native-transport\x00"), 91)
			wire := append(ssFixtureAddress(1, targetPort), payload...)
			switch obfs {
			case "simple_obfs_http":
				wire = append([]byte(fmt.Sprintf("POST /native HTTP/1.1\r\nHost: example.test\r\nContent-Length: %d\r\n\r\n", len(wire))), wire...)
			case "simple_obfs_tls":
				hello := []byte{1, 0, 0, 0, 3, 3}
				r := make([]byte, 32)
				rand.Read(r)
				binary.BigEndian.PutUint32(r, uint32(time.Now().Unix()))
				hello = append(hello, r...)
				hello = append(hello, 32)
				sid := make([]byte, 32)
				rand.Read(sid)
				hello = append(hello, sid...)
				hello = append(hello, make([]byte, 62)...)
				hello = append(hello, 0, 35, byte(len(wire)>>8), byte(len(wire)))
				hello = append(hello, wire...)
				hello = append(hello, make([]byte, 7)...)
				host := "example.test"
				hello = append(hello, byte(len(host)>>8), byte(len(host)))
				hello = append(hello, host...)
				binary.BigEndian.PutUint16(hello[2:], uint16(len(hello)-4))
				wire = append([]byte{22, 3, 1, byte(len(hello) >> 8), byte(len(hello))}, hello...)
			}
			client, err := net.Dial("tcp", v.mc0MXUU.sSKJ79jFt[0].woqAbWjQH_A[0].Addr().String())
			if err != nil {
				t.Fatal(err)
			}
			defer client.Close()
			client.SetDeadline(time.Now().Add(2 * time.Second))
			if _, err = client.Write(wire); err != nil {
				t.Fatal(err)
			}
			var got []byte
			switch obfs {
			case "simple_obfs_http":
				r := bufio.NewReader(client)
				response, err := http.ReadResponse(r, nil)
				if err != nil || response.StatusCode != 101 {
					t.Fatal("HTTP obfs handshake", response, err)
				}
				got = make([]byte, len(payload))
				if _, err = io.ReadFull(r, got); err != nil {
					t.Fatal(err)
				}
			case "simple_obfs_tls":
				for i := 0; len(got) < len(payload); i++ {
					var h [5]byte
					if _, err = io.ReadFull(client, h[:]); err != nil {
						t.Fatal(err)
					}
					record := make([]byte, int(binary.BigEndian.Uint16(h[3:])))
					if _, err = io.ReadFull(client, record); err != nil {
						t.Fatal(err)
					}
					if i >= 2 {
						got = append(got, record...)
					} else if (i == 0 && h[0] != 22) || (i == 1 && h[0] != 20) {
						t.Fatal("unexpected fake TLS record", h)
					}
				}
			default:
				got = make([]byte, len(payload))
				if _, err = io.ReadFull(client, got); err != nil {
					t.Fatal(err)
				}
			}
			if !bytes.Equal(got, payload) {
				t.Fatalf("obfs plaintext mismatch: %d", len(got))
			}
			if err := <-done; err != nil {
				t.Fatal(err)
			}
		})
	}
}

func ssrChainMAC(key []byte, parts ...[]byte) []byte {
	h := hmac.New(md5.New, key)
	for _, p := range parts {
		h.Write(p)
	}
	return h.Sum(nil)
}
func ssrChainUDPZeroPad(seed []byte) int {
	// One xorshift128+ draw, independently expressed from the wire specification.
	x, y := binary.LittleEndian.Uint64(seed[:8]), binary.LittleEndian.Uint64(seed[8:16])
	x ^= x << 23
	next := x ^ y ^ (x >> 17) ^ (y >> 26)
	return int((next + y) % 127)
}
func TestNativeSSRIndependentChainABCDEF(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, protocol := range []string{"auth_chain_a", "auth_chain_b", "auth_chain_c", "auth_chain_d", "auth_chain_e", "auth_chain_f"} {
		for _, udp := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/udp=%t", protocol, udp), func(t *testing.T) {
				v := ssrNativeMatrixCore(t, "none", protocol, "plain")
				state := v.mc0MXUU.sSKJ79jFt[0]
				payload := bytes.Repeat([]byte{0x35, 0x67, 0x89}, 486) //1458; first server frame includes two MSS bytes, total1460.
				var port int
				var done <-chan error
				if udp {
					port, done = trojanMatrixTarget(t, true)
				} else {
					l, err := net.Listen("tcp4", "127.0.0.1:0")
					if err != nil {
						t.Fatal(err)
					}
					defer l.Close()
					port = l.Addr().(*net.TCPAddr).Port
					finished := make(chan error, 1)
					done = finished
					go func() {
						c, err := l.Accept()
						if err == nil {
							defer c.Close()
							c.SetDeadline(time.Now().Add(2 * time.Second))
							p := make([]byte, len(payload))
							_, err = io.ReadFull(c, p)
							if err == nil && !bytes.Equal(p, payload) {
								err = fmt.Errorf("chain target payload differs")
							}
							if err == nil {
								_, err = c.Write(p)
							}
						}
						finished <- err
					}()
				}
				plain := append(ssFixtureAddress(1, port), payload...)
				key := ssrIndependentKey("native-matrix-key", 16)
				network, endpoint := "tcp", state.woqAbWjQH_A[0].Addr().String()
				if udp {
					network, endpoint = "udp", state.aaBpAQ44Zfs[0].LocalAddr().String()
				}
				client, err := net.Dial(network, endpoint)
				if err != nil {
					t.Fatal(err)
				}
				defer client.Close()
				client.SetDeadline(time.Now().Add(2 * time.Second))
				if udp {
					auth := make([]byte, 3)
					rand.Read(auth)
					mac := ssrChainMAC(key, auth)
					pad := ssrChainUDPZeroPad(mac)
					packet := make([]byte, len(plain)+pad+8)
					stream, _ := rc4.NewCipher(ssrIndependentKey(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(mac), 16))
					stream.XORKeyStream(packet, plain)
					copy(packet[len(packet)-8:], auth)
					binary.LittleEndian.PutUint32(packet[len(packet)-5:], 42^binary.LittleEndian.Uint32(mac))
					packet[len(packet)-1] = ssrChainMAC(key, packet[:len(packet)-1])[0]
					if _, err = client.Write(packet); err != nil {
						t.Fatal(err)
					}
					buf := make([]byte, 8192)
					n, err := client.Read(buf)
					if err != nil {
						t.Fatal(err)
					}
					reply := buf[:n]
					if n < 8 || reply[n-1] != ssrChainMAC(key, reply[:n-1])[0] {
						t.Fatal("UDP response HMAC")
					}
					mac = ssrChainMAC(key, reply[n-8:n-1])
					size := n - 8 - ssrChainUDPZeroPad(mac)
					if size < 0 {
						t.Fatal("UDP padding")
					}
					stream, _ = rc4.NewCipher(ssrIndependentKey(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(mac), 16))
					got := make([]byte, size)
					stream.XORKeyStream(got, reply[:size])
					if !bytes.Equal(got, plain) {
						t.Fatal("UDP decrypted address/payload mismatch")
					}
				} else {
					header := make([]byte, 36)
					rand.Read(header[:4])
					recvSeed := ssrChainMAC(key, header[:4])
					copy(header[4:12], recvSeed[:8])
					binary.LittleEndian.PutUint32(header[12:16], 42^binary.LittleEndian.Uint32(recvSeed[8:12]))
					auth := make([]byte, 16)
					binary.LittleEndian.PutUint32(auth, uint32(time.Now().Unix()))
					rand.Read(auth[4:12])
					binary.LittleEndian.PutUint16(auth[12:], 0)
					block, _ := aes.NewCipher(ssrIndependentKey(base64.StdEncoding.EncodeToString(key)+protocol, 16))
					cipher.NewCBCEncrypter(block, make([]byte, 16)).CryptBlocks(header[16:32], auth)
					sendSeed := ssrChainMAC(key, header[12:32])
					copy(header[32:], sendSeed[:4])
					sessionKey := ssrIndependentKey(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(recvSeed), 16)
					enc, _ := rc4.NewCipher(sessionKey)
					dec, _ := rc4.NewCipher(sessionKey)
					// 1465 request bytes and 1460 response bytes are above every native list
					// entry; original A-F choose zero padding here, without calling its strategy.
					frame := make([]byte, len(plain)+4)
					binary.LittleEndian.PutUint16(frame, uint16(len(plain))^binary.LittleEndian.Uint16(recvSeed[14:]))
					enc.XORKeyStream(frame[2:], plain)
					macKey := binary.LittleEndian.AppendUint32(bytes.Clone(key), 1)
					copy(frame[len(frame)-2:], ssrChainMAC(macKey, frame[:len(frame)-2])[:2])
					if _, err = client.Write(append(header, frame...)); err != nil {
						t.Fatal(err)
					}
					var length [2]byte
					if _, err = io.ReadFull(client, length[:]); err != nil {
						t.Fatal(err)
					}
					n := int(binary.LittleEndian.Uint16(length[:]) ^ binary.LittleEndian.Uint16(sendSeed[14:]))
					if n != 1460 {
						t.Fatalf("zero-padding fixture requires first response1460, got %d", n)
					}
					reply := make([]byte, n+2)
					if _, err = io.ReadFull(client, reply); err != nil {
						t.Fatal(err)
					}
					if !bytes.Equal(reply[n:], ssrChainMAC(macKey, length[:], reply[:n])[:2]) {
						t.Fatal("TCP response HMAC")
					}
					got := make([]byte, n)
					dec.XORKeyStream(got, reply[:n])
					if binary.LittleEndian.Uint16(got) != 1460 || !bytes.Equal(got[2:], payload) {
						t.Fatal("TCP MSS/payload mismatch")
					}
				}
				if err := <-done; err != nil {
					t.Fatal("target", err)
				}
			})
		}
	}
}

// Independent OpenSSL primitives are optional test-environment tools. Missing
// tools are visible skips, never reported as tested cipher coverage.
func ssrIndependentLibcrypto(t *testing.T, method string, key, iv, p []byte, decrypt bool) []byte {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cmd *exec.Cmd
	if method == "rc2-cfb" {
		python, err := exec.LookPath("python3")
		if err != nil {
			t.Skip("independent RC2 requires Python3 + system libcrypto")
		}
		// The native constructor explicitly supplies 16 effective key bits. OpenSSL
		// CLI's RC2 defaults differ, so call the independent legacy primitive with
		// that exact parameter. Only the fixed test key and loopback payload enter it.
		code := `import ctypes as c, ctypes.util, sys
lib=c.CDLL(ctypes.util.find_library("crypto"))
lib.RC2_set_key.argtypes=[c.c_void_p,c.c_int,c.c_void_p,c.c_int]
lib.RC2_cfb64_encrypt.argtypes=[c.c_void_p,c.c_void_p,c.c_long,c.c_void_p,c.c_void_p,c.POINTER(c.c_int),c.c_int]
key=bytes.fromhex(sys.argv[1]); iv=c.create_string_buffer(bytes.fromhex(sys.argv[2]),8)
schedule=(c.c_uint*64)(); lib.RC2_set_key(schedule,len(key),key,16)
data=sys.stdin.buffer.read(); out=c.create_string_buffer(len(data)); num=c.c_int(0)
lib.RC2_cfb64_encrypt(data,out,len(data),schedule,iv,c.byref(num),int(sys.argv[3]))
sys.stdout.buffer.write(out.raw)
`
		enc := "1"
		if decrypt {
			enc = "0"
		}
		cmd = exec.CommandContext(ctx, python, "-c", code, hex.EncodeToString(key), hex.EncodeToString(iv), enc)
	} else {
		openssl, err := exec.LookPath("openssl")
		if err != nil {
			t.Skip("independent Camellia requires OpenSSL")
		}
		args := []string{"enc", "-" + method, "-nosalt", "-nopad", "-K", hex.EncodeToString(key), "-iv", hex.EncodeToString(iv)}
		if decrypt {
			args = append(args, "-d")
		}
		cmd = exec.CommandContext(ctx, openssl, args...)
	}
	cmd.Stdin = bytes.NewReader(p)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("independent %s primitive: %v", method, err)
	}
	return out
}

// This test-only IDEA reference uses a 128-bit rotation for key expansion and
// explicit modular arithmetic, independently of native JPLmG42's word schedule.
// Only Encrypt is needed by standard-library CFB in either direction.
type ssrIDEAReference struct{ keys [52]uint16 }

func newSSRIDEAReference(key []byte) *ssrIDEAReference {
	r := new(ssrIDEAReference)
	state := new(big.Int).SetBytes(key)
	mask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))
	for i := range r.keys {
		if i > 0 && i%8 == 0 {
			left := new(big.Int).Lsh(state, 25)
			right := new(big.Int).Rsh(state, 103)
			state = new(big.Int).And(new(big.Int).Or(left, right), mask)
		}
		word := new(big.Int).Rsh(state, uint(112-16*(i%8)))
		r.keys[i] = uint16(word.Uint64())
	}
	return r
}
func (*ssrIDEAReference) BlockSize() int { return 8 }
func (*ssrIDEAReference) Decrypt([]byte, []byte) {
	panic("CFB reference uses Encrypt in both directions")
}
func (r *ssrIDEAReference) Encrypt(dst, src []byte) {
	mul := func(x, y uint16) uint16 {
		a, b := uint64(x), uint64(y)
		if a == 0 {
			a = 65536
		}
		if b == 0 {
			b = 65536
		}
		return uint16((a * b) % 65537)
	}
	a, b, c, d := binary.BigEndian.Uint16(src), binary.BigEndian.Uint16(src[2:]), binary.BigEndian.Uint16(src[4:]), binary.BigEndian.Uint16(src[6:])
	for round := 0; round < 8; round++ {
		k := r.keys[round*6:]
		a = mul(a, k[0])
		b += k[1]
		c += k[2]
		d = mul(d, k[3])
		e := mul(a^c, k[4])
		f := mul((b^d)+e, k[5])
		e += f
		a, b, c, d = a^f, c^f, b^e, d^e
	}
	binary.BigEndian.PutUint16(dst, mul(a, r.keys[48]))
	binary.BigEndian.PutUint16(dst[2:], c+r.keys[49])
	binary.BigEndian.PutUint16(dst[4:], b+r.keys[50])
	binary.BigEndian.PutUint16(dst[6:], mul(d, r.keys[51]))
}
func TestSSRIndependentIDEAReferenceOriginalTrace(t *testing.T) {
	// Original-machine trace recorded earlier in native-ssr-block-original-traces.json.
	key, _ := hex.DecodeString("00254a6f94b9de03284d7297bce1062b")
	plain, _ := hex.DecodeString("000b16212c37424d")
	got := make([]byte, 8)
	newSSRIDEAReference(key).Encrypt(got, plain)
	if hex.EncodeToString(got) != "43379fe8227eebd2" {
		t.Fatalf("independent IDEA reference: %x", got)
	}
}
