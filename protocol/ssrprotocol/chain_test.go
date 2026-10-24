package ssrprotocol

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rc4"
	"encoding/base64"
	"encoding/binary"
	"fmt"
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

var chainVariants = []struct {
	name string
	ctor ojt4CWtzUu
}{
	{"auth_chain_a", oPPzNk}, {"auth_chain_b", gYQ1Wtm}, {"auth_chain_c", vHtqJ7xSQiSj}, {"auth_chain_d", daajdwCbv7}, {"auth_chain_e", xrwCp6}, {"auth_chain_f", hH3aFx},
}

func chainBase(v Vzuv_s9ZS5) *h2QD2cI1 {
	switch c := v.(type) {
	case *h2QD2cI1:
		return c
	case *yeGSuIuC:
		return c.h2QD2cI1
	case *j8lPExkSUzBF:
		return c.h2QD2cI1
	case *gdcp8ElhzO:
		return c.h2QD2cI1
	case *yMywJL:
		return c.h2QD2cI1
	case *pMrxs3_Z:
		return c.h2QD2cI1
	}
	panic("unknown test type")
}
func chainAuthFixture(salt string, userKey, serverKey, iv []byte) (header, recvSeed, sendSeed, key []byte) {
	header = make([]byte, 36)
	copy(header, []byte{1, 2, 3, 4})
	recvSeed = aesTag(md5.New, append(append([]byte{}, iv...), serverKey...), header[:4])
	copy(header[4:], recvSeed[:8])
	binary.LittleEndian.PutUint32(header[12:], 123^binary.LittleEndian.Uint32(recvSeed[8:]))
	var decoded [16]byte
	binary.LittleEndian.PutUint32(decoded[:], uint32(time.Now().Unix()))
	binary.LittleEndian.PutUint32(decoded[4:], 234)
	binary.LittleEndian.PutUint32(decoded[8:], 567)
	binary.LittleEndian.PutUint16(decoded[12:], 4)
	block, _ := aes.NewCipher(H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(userKey)+salt, 16))
	cipher.NewCBCEncrypter(block, make([]byte, 16)).CryptBlocks(header[16:32], decoded[:])
	sendSeed = aesTag(md5.New, userKey, header[12:32])
	copy(header[32:], sendSeed[:4])
	key = H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(userKey)+base64.StdEncoding.EncodeToString(recvSeed), 16)
	return
}

// Client fixtures use the padding strategies already verified against 185
// independent original machine-code cases, but independently build/authenticate
// frames and use separate standard RC4 instances (never server Write or Read).
func chainClientFrame(strategy d_JgwRxIbD4o, stream *rc4.Cipher, userKey, seed, payload []byte, counter uint32) ([]byte, []byte) {
	var rng S.HelGdo2ETK
	pad := strategy.z0O7jemcj(len(payload), seed, &rng)
	b := make([]byte, len(payload)+pad+4)
	binary.LittleEndian.PutUint16(b, uint16(len(payload))^binary.LittleEndian.Uint16(seed[14:]))
	offset := 0
	if len(payload) > 0 && pad > 0 {
		offset = int(rng.Next() % 0x200000011 % uint64(pad))
	}
	stream.XORKeyStream(b[2+offset:], payload)
	key := append(append([]byte{}, userKey...), 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(key[len(userKey):], counter)
	next := aesTag(md5.New, key, b[:len(b)-2])
	copy(b[len(b)-2:], next[:2])
	return b, next
}
func chainClientRead(t *testing.T, r io.Reader, strategy d_JgwRxIbD4o, stream *rc4.Cipher, userKey, seed []byte, count uint32) ([]byte, []byte) {
	t.Helper()
	header := make([]byte, 2)
	if _, e := io.ReadFull(r, header); e != nil {
		t.Fatal(e)
	}
	length := int(binary.LittleEndian.Uint16(header) ^ binary.LittleEndian.Uint16(seed[14:]))
	var rng S.HelGdo2ETK
	pad := strategy.z0O7jemcj(length, seed, &rng)
	b := make([]byte, length+pad+2)
	if _, e := io.ReadFull(r, b); e != nil {
		t.Fatal(e)
	}
	key := append(append([]byte{}, userKey...), 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(key[len(userKey):], count)
	next := aesTag(md5.New, key, header, b[:len(b)-2])
	if !bytes.Equal(next[:2], b[len(b)-2:]) {
		t.Fatal("response frame HMAC")
	}
	offset := 0
	if length > 0 && pad > 0 {
		offset = int(rng.Next() % 0x200000011 % uint64(pad))
	}
	plain := b[offset : offset+length]
	stream.XORKeyStream(plain, plain)
	return plain, next
}
func TestNativeChainActualTCP(t *testing.T) {
	sizes := []uintptr{unsafe.Sizeof(h2QD2cI1{}), unsafe.Sizeof(yeGSuIuC{}), unsafe.Sizeof(j8lPExkSUzBF{}), unsafe.Sizeof(gdcp8ElhzO{}), unsafe.Sizeof(yMywJL{}), unsafe.Sizeof(pMrxs3_Z{})}
	want := []uintptr{328, 56, 8, 8, 8, 32}
	for i := range sizes {
		if sizes[i] != want[i] {
			t.Fatal(sizes)
		}
	}
	for _, variant := range chainVariants {
		for _, mode := range []E.LaTpsgYQ7{0, 1, 2} {
			t.Run(fmt.Sprintf("%s/%d", variant.name, mode), func(t *testing.T) {
				listener, e := net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer listener.Close()
				client, e := net.Dial("tcp4", listener.Addr().String())
				if e != nil {
					t.Fatal(e)
				}
				defer client.Close()
				client.SetDeadline(time.Now().Add(5 * time.Second))
				server, e := listener.Accept()
				if e != nil {
					t.Fatal(e)
				}
				defer server.Close()
				server.SetDeadline(time.Now().Add(5 * time.Second))
				info := &S.L7E7H2aW8KVs{MY3lagqY: "server-password", L6BrlCXd: "aes-128-cfb", NxYA16xMfF: mode, Uklg6401: 321, CJfUaAqL: 4, ZlsyKgz4EI: new(M.GTFAWF[int64, string])}
				info.ZlsyKgz4EI.Set(123, "user-password")
				outer, e := T.VavyqXsqy9V(info, server)
				if e != nil {
					t.Fatal(e)
				}
				defer outer.Free()
				v := variant.ctor(outer)
				v.SetServerInfo(info)
				defer chainBase(v).Free()
				var wire bytes.Buffer
				clientOuter, e := T.VavyqXsqy9V(&S.L7E7H2aW8KVs{MY3lagqY: "server-password", L6BrlCXd: "aes-128-cfb"}, &wire)
				if e != nil {
					t.Fatal(e)
				}
				defer clientOuter.Free()
				clientOuter.Write(nil)
				userKey := clientOuter.GetKey()
				if mode != 0 {
					userKey = []byte("user-password")
				}
				header, recvSeed, sendSeed, key := chainAuthFixture(variant.name, userKey, clientOuter.GetKey(), wire.Bytes()[:16])
				strategy := variant.ctor(nil)
				strategy.SetServerInfo(&S.L7E7H2aW8KVs{IawCjSkuTM4: info.IawCjSkuTM4, CJfUaAqL: 4})
				enc, _ := rc4.NewCipher(key)
				dec, _ := rc4.NewCipher(key)
				frame, next := chainClientFrame(chainBase(strategy).d_JgwRxIbD4o, enc, userKey, recvSeed, []byte("initial"), 1)
				second, _ := chainClientFrame(chainBase(strategy).d_JgwRxIbD4o, enc, userKey, next, []byte("second"), 2)
				clientOuter.Write(append(append(header, frame...), second...))
				if _, e = client.Write(wire.Bytes()); e != nil {
					t.Fatal(e)
				}
				got := make([]byte, 13)
				for off := 0; off < len(got); {
					n, e := v.Read(got[off:min(off+2, len(got))])
					if e != nil {
						t.Fatal(e)
					}
					off += n
				}
				uid := int64(321)
				if mode != 0 {
					uid = 123
				}
				if string(got) != "initialsecond" || v.CID() != 234 || v.CurUID() != uid {
					t.Fatal(string(got), v.CID(), v.CurUID())
				}
				payload := bytes.Repeat([]byte("response"), 2000)
				done := make(chan error, 1)
				go func() {
					n, e := v.Write(payload)
					if n != len(payload) && e == nil {
						e = fmt.Errorf("write count %d", n)
					}
					done <- e
				}()
				decrypt, e := T.VavyqXsqy9V(&S.L7E7H2aW8KVs{MY3lagqY: "server-password", L6BrlCXd: "aes-128-cfb"}, client)
				if e != nil {
					t.Fatal(e)
				}
				defer decrypt.Free()
				var response []byte
				for count := uint32(1); len(response) < len(payload)+2; count++ {
					part, next := chainClientRead(t, decrypt, chainBase(strategy).d_JgwRxIbD4o, dec, userKey, sendSeed, count)
					sendSeed = next
					response = append(response, part...)
				}
				if e = <-done; e != nil {
					t.Fatal(e)
				}
				if binary.LittleEndian.Uint16(response) != 1460 || !bytes.Equal(response[2:], payload) {
					t.Fatal("response plaintext/first MSS")
				}
			})
		}
	}
}
func TestNativeChainActualUDP(t *testing.T) {
	for _, variant := range chainVariants {
		for _, mode := range []E.LaTpsgYQ7{0, 1, 2} {
			t.Run(fmt.Sprintf("%s/%d", variant.name, mode), func(t *testing.T) {
				server, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer server.Close()
				server.SetDeadline(time.Now().Add(3 * time.Second))
				client, e := net.ListenPacket("udp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer client.Close()
				client.SetDeadline(time.Now().Add(3 * time.Second))
				info := &S.L7E7H2aW8KVs{MY3lagqY: "password", L6BrlCXd: "none", NxYA16xMfF: mode, Uklg6401: 321, ZlsyKgz4EI: new(M.GTFAWF[int64, string])}
				info.ZlsyKgz4EI.Set(123, "user-password")
				outer, e := T.DP5vq3(info, server)
				if e != nil {
					t.Fatal(e)
				}
				defer outer.Free()
				v := variant.ctor(outer)
				v.SetServerInfo(info)
				defer chainBase(v).Free()
				key := info.IawCjSkuTM4
				if mode != 0 {
					key = []byte("user-password")
				}
				auth := []byte{3, 4, 5}
				mac := aesTag(md5.New, info.IawCjSkuTM4, auth)
				var rng S.HelGdo2ETK
				rng.InitFromBytes(mac)
				pad := int(rng.Next() % 127)
				request := []byte("request")
				packet := make([]byte, len(request)+pad+8)
				stream, _ := rc4.NewCipher(H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(mac), 16))
				stream.XORKeyStream(packet, request)
				copy(packet[len(packet)-8:], auth)
				binary.LittleEndian.PutUint32(packet[len(packet)-5:], 123^binary.LittleEndian.Uint32(mac))
				packet[len(packet)-1] = aesTag(md5.New, key, packet[:len(packet)-1])[0]
				client.WriteTo(packet, server.LocalAddr())
				buf := make([]byte, 2048)
				n, uid, addr, e := v.ReadFromUID(buf)
				wantUID := int64(321)
				if mode != 0 {
					wantUID = 123
				}
				if e != nil || n != len(request) || string(buf[:n]) != "request" || uid != wantUID {
					t.Fatal(n, uid, addr, e)
				}
				if n, e = v.WriteToUID([]byte("response"), addr, 123); n != 8 || e != nil {
					t.Fatal(n, e)
				}
				n, _, e = client.ReadFrom(buf)
				if e != nil {
					t.Fatal(e)
				}
				reply := buf[:n]
				if reply[n-1] != aesTag(md5.New, key, reply[:n-1])[0] {
					t.Fatal("reply HMAC")
				}
				mac = aesTag(md5.New, info.IawCjSkuTM4, reply[n-8:n-1])
				rng.InitFromBytes(mac)
				pad = int(rng.Next() % 127)
				stream, _ = rc4.NewCipher(H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(mac), 16))
				stream.XORKeyStream(reply, reply[:n-pad-8])
				if string(reply[:n-pad-8]) != "response" {
					t.Fatal("response")
				}
				packet[len(packet)-1] ^= 1
				client.WriteTo(packet, server.LocalAddr())
				n, uid, addr, e = v.ReadFromUID(buf)
				if n != 0 || uid != 0 || addr != nil || e == nil || e.Error() != "auth_chain_* udp packet hmac err" {
					t.Fatal(n, uid, addr, e)
				}
			})
		}
	}
}
