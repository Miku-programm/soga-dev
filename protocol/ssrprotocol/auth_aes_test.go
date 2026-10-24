package ssrprotocol

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"hash"
	"io"
	"net"
	M "recovered.local/soga/util/genmap"
	H "recovered.local/soga/util/ssrhelper"
	E "recovered.local/soga/panel/types"
	T "recovered.local/soga/protocol/ssrcipher"
	S "recovered.local/soga/protocol/ssrshared"
	"testing"
	"time"
	"unsafe"
)

type aesWire struct {
	input  *bytes.Reader
	output bytes.Buffer
}

func (w *aesWire) Read(p []byte) (int, error)  { return w.input.Read(p) }
func (w *aesWire) Write(p []byte) (int, error) { return w.output.Write(p) }
func aesTag(algorithm func() hash.Hash, key []byte, parts ...[]byte) []byte {
	h := hmac.New(algorithm, key)
	for _, p := range parts {
		h.Write(p)
	}
	return h.Sum(nil)
}
func aesAuthFixture(algorithm func() hash.Hash, salt string, userKey, serverKey, iv, payload []byte) []byte {
	var decoded [16]byte
	binary.LittleEndian.PutUint32(decoded[:], uint32(time.Now().Unix()))
	binary.LittleEndian.PutUint32(decoded[4:], 234)
	binary.LittleEndian.PutUint32(decoded[8:], 567)
	binary.LittleEndian.PutUint16(decoded[12:], uint16(31+3+len(payload)+4))
	binary.LittleEndian.PutUint16(decoded[14:], 3)
	key := H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(userKey)+salt, 16)
	block, _ := aes.NewCipher(key)
	header := make([]byte, 31)
	header[0] = 42
	binary.LittleEndian.PutUint32(header[7:], 123)
	cipher.NewCBCEncrypter(block, make([]byte, 16)).CryptBlocks(header[11:27], decoded[:])
	authkey := append(append([]byte{}, iv...), serverKey...)
	copy(header[1:7], aesTag(algorithm, authkey, header[:1])[:6])
	copy(header[27:], aesTag(algorithm, authkey, header[7:27])[:4])
	data := append([]byte{0, 0, 0}, payload...)
	data = append(data, aesTag(algorithm, userKey, header, data)[:4]...)
	return append(header, data...)
}
func aesFrameFixture(algorithm func() hash.Hash, key, payload []byte, count uint32) []byte {
	authkey := append(append([]byte{}, key...), 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(authkey[len(key):], count)
	frame := make([]byte, 5+len(payload)+4)
	binary.LittleEndian.PutUint16(frame, uint16(len(frame)))
	copy(frame[2:4], aesTag(algorithm, authkey, frame[:2])[:2])
	frame[4] = 1
	copy(frame[5:], payload)
	copy(frame[len(frame)-4:], aesTag(algorithm, authkey, frame[:len(frame)-4])[:4])
	return frame
}
func TestAESNativeCipherAuthenticationAndFrames(t *testing.T) {
	if unsafe.Sizeof(lKlYpdomLrhv{}) != 216 {
		t.Fatal("RTTI")
	}
	for _, variant := range []struct {
		name      string
		algorithm func() hash.Hash
		ctor      ojt4CWtzUu
	}{{"auth_aes128_md5", md5.New, jKGRC86Lx9g}, {"auth_aes128_sha1", sha1.New, dffBLS}} {
		for _, mode := range []E.LaTpsgYQ7{0, 1, 2} {
			t.Run(variant.name+string(rune('0'+mode)), func(t *testing.T) {
				var input bytes.Buffer
				clientInfo := &S.L7E7H2aW8KVs{MY3lagqY: "server-password", L6BrlCXd: "aes-128-cfb"}
				client, e := T.VavyqXsqy9V(clientInfo, &input)
				if e != nil {
					t.Fatal(e)
				}
				defer client.Free()
				client.Write(nil)
				userKey := client.GetKey()
				if mode != 0 {
					h := variant.algorithm()
					h.Write([]byte("user-password"))
					userKey = h.Sum(nil)
				}
				auth := aesAuthFixture(variant.algorithm, variant.name, userKey, client.GetKey(), input.Bytes()[:16], []byte("initial"))
				frame := aesFrameFixture(variant.algorithm, userKey, []byte("second"), 1)
				client.Write(append(auth, frame...))
				raw := &aesWire{input: bytes.NewReader(input.Bytes())}
				serverInfo := &S.L7E7H2aW8KVs{MY3lagqY: "server-password", L6BrlCXd: "aes-128-cfb", NxYA16xMfF: mode, Uklg6401: 321, ZlsyKgz4EI: new(M.GTFAWF[int64, string])}
				serverInfo.ZlsyKgz4EI.Set(123, "user-password")
				server, e := T.VavyqXsqy9V(serverInfo, raw)
				if e != nil {
					t.Fatal(e)
				}
				defer server.Free()
				c := variant.ctor(server).(*lKlYpdomLrhv)
				c.SetServerInfo(serverInfo)
				defer c.Free()
				got := make([]byte, 13)
				for offset := 0; offset < len(got); {
					n, e := c.Read(got[offset:min(offset+2, len(got))])
					if e != nil {
						t.Fatal(e)
					}
					offset += n
				}
				if string(got) != "initialsecond" || c.CID() != 234 || c.eOnqFdRmAaYM != 2 {
					t.Fatal(string(got), c.CID(), c.eOnqFdRmAaYM)
				}
				wantUID := int64(321)
				if mode != 0 {
					wantUID = 123
				}
				if c.CurUID() != wantUID {
					t.Fatal(c.CurUID())
				}
				payload := bytes.Repeat([]byte("reply"), 2000)
				if n, e := c.Write(payload); n != len(payload) || e != nil {
					t.Fatal(n, e)
				}
				decrypt, e := T.VavyqXsqy9V(&S.L7E7H2aW8KVs{MY3lagqY: "server-password", L6BrlCXd: "aes-128-cfb"}, &raw.output)
				if e != nil {
					t.Fatal(e)
				}
				defer decrypt.Free()
				plain, e := io.ReadAll(decrypt)
				if e != nil {
					t.Fatal(e)
				}
				expected := append(aesFrameFixture(variant.algorithm, userKey, payload[:8000], 1), aesFrameFixture(variant.algorithm, userKey, payload[8000:], 2)...)
				if !bytes.Equal(plain, expected) {
					t.Fatal("encrypted response framing")
				}
			})
		}
	}
}
func TestAESFrameFailureCounterAndEmptyReply(t *testing.T) {
	key := []byte("test-key")
	cases := []struct {
		name  string
		frame []byte
		want  string
	}{{"hmac", []byte{9, 0, 0, 0}, "hmac1 error"}}
	for _, size := range []int{6, 32769} {
		header := make([]byte, 4)
		binary.LittleEndian.PutUint16(header, uint16(size))
		k := append(append([]byte{}, key...), 1, 0, 0, 0)
		copy(header[2:], aesTag(md5.New, k, header[:2])[:2])
		want := "packet length too small: 6"
		if size > 32768 {
			want = "packet length too large: 32769"
		}
		cases = append(cases, struct {
			name  string
			frame []byte
			want  string
		}{want, header, want})
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := jKGRC86Lx9g(nil).(*lKlYpdomLrhv)
			defer c.Free()
			c.h6LJlye = key
			c.ReadWriter = &aesWire{input: bytes.NewReader(tc.frame)}
			if _, e := c.ibH0hkF(); e == nil || e.Error() != tc.want || c.eOnqFdRmAaYM != 2 {
				t.Fatal(e, c.eOnqFdRmAaYM)
			}
		})
	}
	raw := &aesWire{input: bytes.NewReader(aesFrameFixture(md5.New, key, nil, 1))}
	c := jKGRC86Lx9g(nil).(*lKlYpdomLrhv)
	defer c.Free()
	c.h6LJlye = key
	c.ReadWriter = raw
	got, e := c.ibH0hkF()
	if got != nil || e != nil || !bytes.Equal(raw.output.Bytes(), aesFrameFixture(md5.New, key, nil, 1)) {
		t.Fatal(got, e)
	}
}

func TestAESNativeUDPMultiuser(t *testing.T) {
	for _, variant := range []struct {
		name      string
		algorithm func() hash.Hash
		ctor      ojt4CWtzUu
	}{{"md5", md5.New, jKGRC86Lx9g}, {"sha1", sha1.New, dffBLS}} {
		for _, mode := range []E.LaTpsgYQ7{0, 1, 2} {
			t.Run(variant.name+string(rune('0'+mode)), func(t *testing.T) {
				server, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer server.Close()
				client, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer client.Close()
				server.SetDeadline(time.Now().Add(3 * time.Second))
				client.SetDeadline(time.Now().Add(3 * time.Second))
				info := &S.L7E7H2aW8KVs{MY3lagqY: "password", L6BrlCXd: "none", NxYA16xMfF: mode, Uklg6401: 321, ZlsyKgz4EI: new(M.GTFAWF[int64, string])}
				info.ZlsyKgz4EI.Set(123, "user-password")
				cipher, e := T.DP5vq3(info, server)
				if e != nil {
					t.Fatal(e)
				}
				defer cipher.Free()
				c := variant.ctor(cipher).(*lKlYpdomLrhv)
				defer c.Free()
				c.SetServerInfo(info)
				key := cipher.GetKey()
				if mode != 0 {
					h := variant.algorithm()
					h.Write([]byte("user-password"))
					key = h.Sum(nil)
				}
				wire := append([]byte("request"), 123, 0, 0, 0)
				wire = append(wire, aesTag(variant.algorithm, key, wire)[:4]...)
				if _, e = client.WriteTo(wire, server.LocalAddr()); e != nil {
					t.Fatal(e)
				}
				buf := make([]byte, 512)
				n, uid, addr, e := c.ReadFromUID(buf)
				if e != nil || string(buf[:n]) != "request" {
					t.Fatal(n, uid, addr, e)
				}
				want := int64(321)
				if mode != 0 {
					want = 123
				}
				if uid != want {
					t.Fatal(uid)
				}
				if n, e = c.WriteToUID([]byte("response"), addr, 123); n != 8 || e != nil {
					t.Fatal(n, e)
				}
				n, _, e = client.ReadFrom(buf)
				expected := append([]byte("response"), aesTag(variant.algorithm, info.IawCjSkuTM4, []byte("response"))[:4]...)
				if e != nil || !bytes.Equal(buf[:n], expected) {
					t.Fatal(n, e, buf[:n])
				}
				wire[len(wire)-1] ^= 1
				client.WriteTo(wire, server.LocalAddr())
				n, uid, addr, e = c.ReadFromUID(buf)
				if n != 0 || uid != 0 || addr != nil || e == nil || e.Error() != "auth_aes128_* udp packet hmac err" {
					t.Fatal(n, uid, addr, e)
				}
			})
		}
	}
}
