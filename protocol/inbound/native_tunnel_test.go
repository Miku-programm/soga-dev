package inbound

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/sha3"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

// 独立测试客户端：只用标准密码原语和逐字节帧格式，不调用被测隧道编解码器。
type independentTunnelRecords struct {
	io.ReadWriter
	master           []byte
	tx, rx           cipher.AEAD
	txNonce, rxNonce []byte
	txMask, rxMask   sha3.ShakeHash
	pending          []byte
}

var tunnelFixtureInfo, _ = hex.DecodeString("d8b8c066c10e8b0d1b1ba3ad1f7de1525625d03f9d9f682f1bd6097b13a22b38")

func tunnelFixtureHKDF(secret, salt, info []byte, n int) []byte {
	p := make([]byte, n)
	if _, e := io.ReadFull(hkdf.New(sha3.New512, secret, salt, info), p); e != nil {
		panic(e)
	}
	return p
}
func tunnelFixtureNonce(p []byte) {
	for i := range p {
		p[i]++
		if p[i] != 0 {
			break
		}
	}
}
func tunnelFixtureMask(h sha3.ShakeHash) uint16 {
	var b [2]byte
	io.ReadFull(h, b[:])
	return binary.BigEndian.Uint16(b[:])
}
func (c *independentTunnelRecords) initDirection(salt []byte) (cipher.AEAD, sha3.ShakeHash, error) {
	key := tunnelFixtureHKDF(c.master, salt, tunnelFixtureInfo, 16)
	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, nil, e
	}
	a, e := cipher.NewGCM(block)
	mask := sha3.NewShake128()
	mask.Write(salt)
	mask.Write(key)
	return a, mask, e
}
func (c *independentTunnelRecords) Write(p []byte) (int, error) {
	var out []byte
	if c.tx == nil {
		c.txNonce = make([]byte, 12)
		if _, e := rand.Read(c.txNonce); e != nil {
			return 0, e
		}
		var e error
		c.tx, c.txMask, e = c.initDirection(c.txNonce)
		if e != nil {
			return 0, e
		}
		out = append(out, c.txNonce...)
	}
	pad := int(tunnelFixtureMask(c.txMask) & 127)
	header := binary.BigEndian.AppendUint16(nil, uint16(len(p)+pad)^tunnelFixtureMask(c.txMask))
	out = c.tx.Seal(out, c.txNonce, header, nil)
	tunnelFixtureNonce(c.txNonce)
	body := append(append([]byte(nil), p...), make([]byte, pad)...)
	out = c.tx.Seal(out, c.txNonce, body, nil)
	tunnelFixtureNonce(c.txNonce)
	for len(out) > 0 {
		n, e := c.ReadWriter.Write(out)
		if e != nil {
			return 0, e
		}
		if n == 0 {
			return 0, io.ErrNoProgress
		}
		out = out[n:]
	}
	return len(p), nil
}
func (c *independentTunnelRecords) Read(p []byte) (int, error) {
	for len(c.pending) == 0 {
		if c.rx == nil {
			c.rxNonce = make([]byte, 12)
			if _, e := io.ReadFull(c.ReadWriter, c.rxNonce); e != nil {
				return 0, e
			}
			var e error
			c.rx, c.rxMask, e = c.initDirection(c.rxNonce)
			if e != nil {
				return 0, e
			}
		}
		header := make([]byte, 18)
		if _, e := io.ReadFull(c.ReadWriter, header); e != nil {
			return 0, e
		}
		decoded, e := c.rx.Open(nil, c.rxNonce, header, nil)
		if e != nil {
			return 0, e
		}
		tunnelFixtureNonce(c.rxNonce)
		pad := int(tunnelFixtureMask(c.rxMask) & 127)
		n := int(binary.BigEndian.Uint16(decoded) ^ tunnelFixtureMask(c.rxMask))
		if n < pad {
			return 0, fmt.Errorf("invalid tunnel padding")
		}
		body := make([]byte, n+16)
		if _, e = io.ReadFull(c.ReadWriter, body); e != nil {
			return 0, e
		}
		decoded, e = c.rx.Open(nil, c.rxNonce, body, nil)
		if e != nil {
			return 0, e
		}
		tunnelFixtureNonce(c.rxNonce)
		c.pending = decoded[:n-pad]
	}
	n := copy(p, c.pending)
	c.pending = c.pending[n:]
	return n, nil
}

type independentTunnelWS struct {
	*websocket.Conn
	pending []byte
}

func (c *independentTunnelWS) Read(p []byte) (int, error) {
	for len(c.pending) == 0 {
		_, b, e := c.ReadMessage()
		if e != nil {
			return 0, e
		}
		c.pending = b
	}
	n := copy(p, c.pending)
	c.pending = c.pending[n:]
	return n, nil
}
func (c *independentTunnelWS) Write(p []byte) (int, error) {
	if e := c.WriteMessage(websocket.BinaryMessage, p); e != nil {
		return 0, e
	}
	return len(p), nil
}
func tunnelMuxFrame(cmd byte, id uint32, p []byte) []byte {
	b := []byte{1, cmd}
	b = binary.LittleEndian.AppendUint16(b, uint16(len(p)))
	b = binary.LittleEndian.AppendUint32(b, id)
	return append(b, p...)
}

func TestNativeTrojanEncryptedTunnelMatrix(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	cert, key := hysteriaFixtureCertificate(t)
	dir := t.TempDir()
	certPath, keyPath := filepath.Join(dir, "cert.pem"), filepath.Join(dir, "key.pem")
	for p, data := range map[string]string{certPath: cert, keyPath: key} {
		if e := os.WriteFile(p, []byte(data), 0600); e != nil {
			t.Fatal(e)
		}
	}
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM([]byte(cert))
	for _, mode := range []string{"tcp-tunnel", "tls-tunnel", "ws-tunnel", "ws-tls-tunnel"} {
		for _, mux := range []bool{false, true} {
			for _, udp := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/mux=%t/udp=%t", mode, mux, udp), func(t *testing.T) {
					// 非空配置键只是原始本地状态分支的测试输入；未模拟远端授权成功。
					conf := fmt.Sprintf("listen=127.0.0.1\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=1\nudp_timeout=1\nsoga_key=fixture-only-no-authority-contact\ntunnel_enable=true\ntunnel_type=%s\ntunnel_mux=%t\ntunnel_password=tunnel-fixture\ntunnel_method=aes-128-gcm\ntunnel_ws_path=/native\ncert_file=%s\nkey_file=%s\n", mode, mux, certPath, keyPath)
					path := filepath.Join(dir, "soga.conf")
					if e := os.WriteFile(path, []byte(conf), 0600); e != nil {
						t.Fatal(e)
					}
					C.MkwNWciPHS(path)
					if e := C.TvBECnD9fM6(); e != nil {
						t.Fatal(e)
					}
					cfg, e := C.R3DjPWSnt()
					if e != nil || cfg == nil {
						t.Fatalf("tunnel constructor: %v", e)
					}
					node := &L.Wr5SVP{EMlXEc: 32, JhIYwEukjd: L.L9AiYC{T4cuhG7MTbG: "tcp", AJjaqrBjj: "none"}}
					panel := &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{trojanPanelFixture(9, "tunnel-inner")}}
					core, e := f1mPzegGt(&DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: panel})
					if e != nil {
						t.Fatal(e)
					}
					v := core.(*EfaIFpp)
					defer v.Stop()
					v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
					if e = v.Start(); e != nil {
						t.Fatal(e)
					}
					addr := v.cDyEMW2[0].Addr().String()
					tc := &tls.Config{RootCAs: roots, ServerName: "hysteria"}
					var conn io.ReadWriter
					if strings.HasPrefix(mode, "ws-") {
						scheme := "ws"
						if mode == "ws-tls-tunnel" {
							scheme = "wss"
						}
						d := websocket.Dialer{TLSClientConfig: tc, HandshakeTimeout: 3 * time.Second}
						w, _, e := d.Dial(scheme+"://"+addr+"/native", nil)
						if e != nil {
							t.Fatal(e)
						}
						defer w.Close()
						w.SetReadDeadline(time.Now().Add(5 * time.Second))
						conn = &independentTunnelWS{Conn: w}
					} else {
						raw, e := net.DialTimeout("tcp", addr, time.Second)
						if e != nil {
							t.Fatal(e)
						}
						defer raw.Close()
						raw.SetDeadline(time.Now().Add(5 * time.Second))
						conn = raw
						if mode == "tls-tunnel" {
							conn = tls.Client(raw, tc)
						}
					}
					sum := sha3.Sum512([]byte("tunnel-fixture"))
					master := sum[:16]
					token := tunnelFixtureHKDF(master, tunnelFixtureInfo, []byte("tunnel-fixture"), 16)
					encrypted := &independentTunnelRecords{ReadWriter: conn, master: master}
					targetPort, done := trojanMatrixTarget(t, udp)
					target := binary.BigEndian.AppendUint16([]byte{1, 127, 0, 0, 1}, uint16(targetPort))
					digest := sha256.Sum224([]byte("tunnel-inner"))
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
					payload := append(append([]byte(nil), token...), 1)
					payload = append(payload, []byte("PROXY TCP4 192.0.2.44 198.51.100.9 12345 443\r\n")...)
					payload = append(payload, wire...)
					if mux {
						payload = append(tunnelMuxFrame(0, 3, nil), tunnelMuxFrame(2, 3, payload)...)
					}
					if _, e = encrypted.Write(payload); e != nil {
						t.Fatal(e)
					}
					var got []byte
					if mux {
						for len(got) < len(expected) {
							var header [8]byte
							if _, e = io.ReadFull(encrypted, header[:]); e != nil {
								t.Fatal(e)
							}
							body := make([]byte, int(binary.LittleEndian.Uint16(header[2:4])))
							if _, e = io.ReadFull(encrypted, body); e != nil {
								t.Fatal(e)
							}
							if header[0] != 1 || binary.LittleEndian.Uint32(header[4:]) != 3 {
								t.Fatal("mux response header", header)
							}
							if header[1] == 2 {
								got = append(got, body...)
							} else if header[1] == 1 {
								t.Fatal("premature stream FIN")
							}
						}
					} else {
						got = make([]byte, len(expected))
						if _, e = io.ReadFull(encrypted, got); e != nil {
							t.Fatal(e)
						}
					}
					if !bytes.Equal(got, expected) {
						t.Fatal("tunnel target response differs")
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
