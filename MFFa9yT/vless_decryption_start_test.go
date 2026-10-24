package MFFa9yT

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/mlkem"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
	"io"
	"lukechampine.com/blake3"
	"net"
	CA "recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/LFWTTXF7JR"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"strings"
	"testing"
	"time"
)

// These finite wire fixtures use only cryptographic primitives, never the
// recovered bQxeORT client, server, nonce or framing helpers.
type decryptionFixtureAEAD struct {
	cipher.AEAD
	nonce [12]byte
}

func decryptionFixtureCipher(context, material []byte, useAES bool) *decryptionFixtureAEAD {
	var key [32]byte
	blake3.DeriveKey(key[:], string(context), material)
	var a cipher.AEAD
	if useAES {
		block, _ := aes.NewCipher(key[:])
		a, _ = cipher.NewGCM(block)
	} else {
		a, _ = chacha20poly1305.New(key[:])
	}
	return &decryptionFixtureAEAD{AEAD: a}
}
func (a *decryptionFixtureAEAD) next() []byte {
	for i := 11; i >= 0; i-- {
		a.nonce[i]++
		if a.nonce[i] != 0 {
			break
		}
	}
	return a.nonce[:]
}
func (a *decryptionFixtureAEAD) seal(p, aad []byte) []byte { return a.AEAD.Seal(nil, a.next(), p, aad) }
func (a *decryptionFixtureAEAD) open(p, aad []byte) ([]byte, error) {
	return a.AEAD.Open(nil, a.next(), p, aad)
}
func decryptionFixtureCTR(material, iv []byte) cipher.Stream {
	var key [32]byte
	blake3.DeriveKey(key[:], "VLESS", material)
	block, _ := aes.NewCipher(key[:])
	return cipher.NewCTR(block, iv)
}

type decryptionFixtureSession struct{ ticket, forward []byte }
type decryptionFixtureWire struct {
	raw                       net.Conn
	send, receive             *decryptionFixtureAEAD
	sendHeader, receiveHeader cipher.Stream
	material                  []byte
	session                   *decryptionFixtureSession
	resume                    bool
}

func decryptionFixtureHandshake(t *testing.T, raw net.Conn, publics [][]byte, mode int, useAES bool, session *decryptionFixtureSession) *decryptionFixtureWire {
	t.Helper()
	salt := make([]byte, 16)
	rand.Read(salt)
	request := append([]byte{}, salt...)
	var shared []byte
	var previous cipher.Stream
	for i, p := range publics {
		var encoded []byte
		if len(p) == 32 {
			pub, e := ecdh.X25519().NewPublicKey(p)
			if e != nil {
				t.Fatal(e)
			}
			k, _ := ecdh.X25519().GenerateKey(rand.Reader)
			encoded = k.PublicKey().Bytes()
			shared, e = k.ECDH(pub)
			if e != nil {
				t.Fatal(e)
			}
		} else {
			pub, e := mlkem.NewEncapsulationKey768(p)
			if e != nil {
				t.Fatal(e)
			}
			shared, encoded = pub.Encapsulate()
		}
		if mode != 0 {
			decryptionFixtureCTR(p, salt).XORKeyStream(encoded, encoded)
		}
		if previous != nil {
			previous.XORKeyStream(encoded[:32], encoded[:32])
		}
		request = append(request, encoded...)
		if i+1 < len(publics) {
			previous = decryptionFixtureCTR(shared, salt)
			hash := blake3.Sum512(publics[i+1])
			enc := append([]byte{}, hash[:32]...)
			previous.XORKeyStream(enc, enc)
			request = append(request, enc...)
		}
	}
	hs := decryptionFixtureCipher(salt, shared, useAES)
	out := &decryptionFixtureWire{raw: raw, session: session}
	if len(session.ticket) > 0 {
		out.resume = true
		out.material = append(append([]byte{}, session.forward...), shared...)
		request = append(request, hs.seal([]byte{0, 32}, nil)...)
		sealed := hs.seal(session.ticket, nil)
		request = append(request, sealed...)
		out.send = decryptionFixtureCipher(sealed, out.material, useAES)
		if mode == 2 {
			out.sendHeader = decryptionFixtureCTR(out.material, salt)
		}
		if _, e := raw.Write(request); e != nil {
			t.Fatal(e)
		}
		return out
	}
	kem, _ := mlkem.GenerateKey768()
	x, _ := ecdh.X25519().GenerateKey(rand.Reader)
	pair := append(kem.EncapsulationKey().Bytes(), x.PublicKey().Bytes()...)
	request = append(request, hs.seal([]byte{4, 208}, nil)...)
	request = append(request, hs.seal(pair, nil)...)
	request = append(request, hs.seal([]byte{0, 17}, nil)...)
	request = append(request, hs.seal([]byte{0x72}, nil)...)
	if _, e := raw.Write(request); e != nil {
		t.Fatal(e)
	}
	encrypted := make([]byte, 1136)
	if _, e := io.ReadFull(raw, encrypted); e != nil {
		t.Fatal(e)
	}
	server, e := hs.AEAD.Open(nil, bytes.Repeat([]byte{255}, 12), encrypted, nil)
	if e != nil {
		t.Fatal(e)
	}
	kemSecret, e := kem.Decapsulate(server[:1088])
	if e != nil {
		t.Fatal(e)
	}
	xp, e := ecdh.X25519().NewPublicKey(server[1088:1120])
	if e != nil {
		t.Fatal(e)
	}
	xSecret, e := x.ECDH(xp)
	if e != nil {
		t.Fatal(e)
	}
	session.forward = append(append([]byte{}, kemSecret...), xSecret...)
	out.material = append(append([]byte{}, session.forward...), shared...)
	out.send = decryptionFixtureCipher(pair, out.material, useAES)
	out.receive = decryptionFixtureCipher(server, out.material, useAES)
	token := make([]byte, 32)
	if _, e = io.ReadFull(raw, token); e != nil {
		t.Fatal(e)
	}
	session.ticket, e = out.receive.open(token, nil)
	if e != nil {
		t.Fatal(e)
	}
	if binary.BigEndian.Uint16(session.ticket) == 0 {
		t.Fatal("server did not issue reusable ticket")
	}
	length := make([]byte, 18)
	if _, e = io.ReadFull(raw, length); e != nil {
		t.Fatal(e)
	}
	length, e = out.receive.open(length, nil)
	if e != nil {
		t.Fatal(e)
	}
	padding := make([]byte, int(binary.BigEndian.Uint16(length)))
	if _, e = io.ReadFull(raw, padding); e != nil {
		t.Fatal(e)
	}
	if _, e = out.receive.open(padding, nil); e != nil {
		t.Fatal(e)
	}
	if mode == 2 {
		out.sendHeader = decryptionFixtureCTR(out.material, salt)
		out.receiveHeader = decryptionFixtureCTR(out.material, session.ticket)
	}
	return out
}
func (c *decryptionFixtureWire) exchange(t *testing.T, wire []byte, size int, useAES bool, mode int) []byte {
	t.Helper()
	header := []byte{23, 3, 3, 0, 0}
	binary.BigEndian.PutUint16(header[3:], uint16(len(wire)+16))
	body := c.send.seal(wire, header)
	if c.sendHeader != nil {
		c.sendHeader.XORKeyStream(header, header)
	}
	if _, e := c.raw.Write(append(header, body...)); e != nil {
		t.Fatal(e)
	}
	if c.resume {
		salt := make([]byte, 16)
		if _, e := io.ReadFull(c.raw, salt); e != nil {
			t.Fatal(e)
		}
		c.receive = decryptionFixtureCipher(salt, c.material, useAES)
		if mode == 2 {
			c.receiveHeader = decryptionFixtureCTR(c.material, salt)
		}
	}
	var out []byte
	for len(out) < size {
		h := make([]byte, 5)
		if _, e := io.ReadFull(c.raw, h); e != nil {
			t.Fatal(e)
		}
		if c.receiveHeader != nil {
			c.receiveHeader.XORKeyStream(h, h)
		}
		if !bytes.Equal(h[:3], []byte{23, 3, 3}) {
			t.Fatal("record header", h)
		}
		n := int(binary.BigEndian.Uint16(h[3:]))
		if n < 16 || n > 16640 {
			t.Fatal("record length", n)
		}
		enc := make([]byte, n)
		if _, e := io.ReadFull(c.raw, enc); e != nil {
			t.Fatal(e)
		}
		plain, e := c.receive.open(enc, h)
		if e != nil {
			t.Fatal(e)
		}
		out = append(out, plain...)
	}
	return out
}
func TestNativeVLESSDecryptionStartFullResumeMatrix(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for mode, name := range []string{"native", "xorpub", "random"} {
		for _, mixed := range []bool{false, true} {
			for _, useAES := range []bool{false, true} {
				for _, udp := range []bool{false, true} {
					t.Run(fmt.Sprintf("%s/mixed=%t/aes=%t/udp=%t", name, mixed, useAES, udp), func(t *testing.T) {
						x, _ := ecdh.X25519().GenerateKey(rand.Reader)
						secrets := [][]byte{x.Bytes()}
						publics := [][]byte{x.PublicKey().Bytes()}
						if mixed {
							k, _ := mlkem.GenerateKey768()
							x2, _ := ecdh.X25519().GenerateKey(rand.Reader)
							secrets = append(secrets, k.Bytes(), x2.Bytes())
							publics = append(publics, k.EncapsulationKey().Bytes(), x2.PublicKey().Bytes())
						}
						var encoded []string
						for _, s := range secrets {
							encoded = append(encoded, base64.RawURLEncoding.EncodeToString(s))
						}
						config := "mlkem768x25519plus." + name + ".600-600s.100-35-35." + strings.Join(encoded, ".")
						node := &L.Wr5SVP{EMlXEc: 32, Ck6DxjB: L.VkkFLH{XmIaPj3CzOF: "tcp", DlDK5Y9AS: config}}
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
						session := &decryptionFixtureSession{}
						for pass := range 2 {
							target, done := trojanMatrixTarget(t, udp)
							wire, want := vlessMatrixWire(udp, target)
							raw, e := net.DialTimeout("tcp", v.kX5i8U[0].Addr().String(), time.Second)
							if e != nil {
								t.Fatal(e)
							}
							raw.SetDeadline(time.Now().Add(5 * time.Second))
							defer raw.Close()
							client := decryptionFixtureHandshake(t, raw, publics, mode, useAES, session)
							if client.resume != (pass == 1) {
								t.Fatal("wrong handshake path")
							}
							got := client.exchange(t, wire, len(want), useAES, mode)
							if !bytes.Equal(got, want) {
								t.Fatal("decrypted target response mismatch")
							}
							raw.Close()
							select {
							case e = <-done:
								if e != nil {
									t.Fatal(e)
								}
							case <-time.After(5 * time.Second):
								t.Fatal("target not reached")
							}
						}
					})
				}
			}
		}
	}
}
