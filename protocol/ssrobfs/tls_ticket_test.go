package ssrobfs

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"io"
	"net"
	S "recovered.local/soga/protocol/ssrshared"
	"testing"
	"time"
	"unsafe"
)

func fixtureTag(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(data)
	return h.Sum(nil)[:10]
}
func ticketHello(key []byte) ([]byte, []byte) {
	sid := bytes.Repeat([]byte{0x64}, 32)
	authkey := append(append([]byte{}, key...), sid...)
	random := make([]byte, 32)
	rand.Read(random)
	binary.BigEndian.PutUint32(random, uint32(time.Now().Unix()))
	copy(random[22:], fixtureTag(authkey, random[:22]))
	b := []byte{1, 0, 0, 0, 3, 3}
	b = append(b, random...)
	b = append(b, 32)
	b = append(b, sid...)
	b = append(b, make([]byte, 48)...)
	b = append(b, []byte("host.example")...)
	b = append(b, 0, 23, 0, 0)
	binary.BigEndian.PutUint16(b[2:], uint16(len(b)-4))
	return append([]byte{22, 3, 1, byte(len(b) >> 8), byte(len(b))}, b...), authkey
}
func ticketFinish(key []byte) []byte {
	data := append([]byte{}, ticketPrefixAt509A770...)
	data = append(data, 32)
	random := make([]byte, 22)
	rand.Read(random)
	data = append(data, random...)
	return append(data, fixtureTag(key, data)...)
}
func TestTLSTicketActualConnection(t *testing.T) {
	if unsafe.Sizeof(g7X2glAwFd91{}) != 120 {
		t.Fatal("RTTI")
	}
	server, client := net.Pipe()
	defer client.Close()
	defer server.Close()
	client.SetDeadline(time.Now().Add(5 * time.Second))
	server.SetDeadline(time.Now().Add(5 * time.Second))
	key := []byte("original-key")
	hello, authkey := ticketHello(key)
	info := &S.L7E7H2aW8KVs{IawCjSkuTM4: key, SP2KC4lkkl3C: len(key)}
	c := jRGJRCFN(server).(*g7X2glAwFd91)
	c.SetServerInfo(info)
	defer c.Free()
	done := make(chan error, 1)
	go func() {
		b := make([]byte, 7)
		_, e := io.ReadFull(c, b)
		if e == nil && !bytes.Equal(b, []byte("message")) {
			e = io.ErrUnexpectedEOF
		}
		if e == nil {
			_, e = c.Write(b)
		}
		done <- e
	}()
	if _, e := client.Write(hello); e != nil {
		t.Fatal(e)
	}
	var handshake []byte
	for {
		header := make([]byte, 5)
		if _, e := io.ReadFull(client, header); e != nil {
			t.Fatal(e)
		}
		body := make([]byte, binary.BigEndian.Uint16(header[3:]))
		if _, e := io.ReadFull(client, body); e != nil {
			t.Fatal(e)
		}
		handshake = append(handshake, header...)
		handshake = append(handshake, body...)
		if header[0] == 20 {
			header = make([]byte, 5)
			io.ReadFull(client, header)
			body = make([]byte, binary.BigEndian.Uint16(header[3:]))
			io.ReadFull(client, body)
			handshake = append(handshake, header...)
			handshake = append(handshake, body...)
			break
		}
	}
	if !bytes.Equal(fixtureTag(authkey, handshake[:len(handshake)-10]), handshake[len(handshake)-10:]) {
		t.Fatal("server HMAC")
	}
	if !bytes.Equal(handshake[33:43], fixtureTag(authkey, handshake[11:33])) {
		t.Fatal("server random HMAC")
	}
	wire := append(ticketFinish(authkey), 23, 3, 3, 0, 7)
	wire = append(wire, []byte("message")...)
	if _, e := client.Write(wire); e != nil {
		t.Fatal(e)
	}
	output := make([]byte, 12)
	if _, e := io.ReadFull(client, output); e != nil {
		t.Fatal(e)
	}
	if string(output[5:]) != "message" {
		t.Fatal(output)
	}
	if e := <-done; e != nil {
		t.Fatal(e)
	}
	if info.EFLMTH != "host.example" || c.kr85QF != 5 || !c.g3yMrTl9L {
		t.Fatal(info, c.kr85QF, c.g3yMrTl9L)
	}
}
func TestTicketAuthenticationErrors(t *testing.T) {
	key := []byte("key")
	for _, mode := range []string{"bad-hmac", "bad-time", "replay", "finish-prefix", "finish-tag"} {
		t.Run(mode, func(t *testing.T) {
			hello, authkey := ticketHello(key)
			switch mode {
			case "bad-hmac":
				hello[27] ^= 1
			case "bad-time":
				binary.BigEndian.PutUint32(hello[11:], 1)
			}
			finish := ticketFinish(authkey)
			if mode == "finish-prefix" {
				finish[0] = 19
			}
			if mode == "finish-tag" {
				finish[len(finish)-1] ^= 1
			}
			w := &wireFixture{read: bytes.NewReader(append(hello, finish...))}
			c := jRGJRCFN(w).(*g7X2glAwFd91)
			defer c.Free()
			c.SetServerInfo(&S.L7E7H2aW8KVs{IawCjSkuTM4: key, SP2KC4lkkl3C: len(key)})
			e := c.edo12OH()
			if mode == "replay" {
				if e != nil {
					t.Fatal(e)
				}
				w.read = bytes.NewReader(hello)
				e = c.hS8VwDp3P()
			}
			want := map[string]string{"bad-hmac": "tls auth wrong sha1", "bad-time": "tls auth wrong time", "replay": "tls auth repeat attack detect", "finish-prefix": "handshake2 first data not match", "finish-tag": "handshake2 verify failed"}[mode]
			if e == nil || e.Error() != want {
				t.Fatal(e)
			}
		})
	}
}
