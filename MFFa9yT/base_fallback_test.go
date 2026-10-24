package MFFa9yT

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"io"
	"net"
	L "recovered.local/soga/LFWTTXF7JR"
	"testing"
	"time"
)

func TestFallbackActualPrefixAndBidirectionalRelay(t *testing.T) {
	for _, prefix := range [][]byte{nil, []byte("prefix:")} {
		t.Run(string(prefix), func(t *testing.T) {
			server, e := net.Listen("tcp4", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer server.Close()
			ready := make(chan struct{})
			done := make(chan struct{})
			go func() {
				defer close(done)
				c, e := server.Accept()
				if e != nil {
					return
				}
				defer c.Close()
				close(ready)
				c.SetDeadline(time.Now().Add(5 * time.Second))
				io.Copy(c, c)
			}()
			client, inbound := net.Pipe()
			defer client.Close()
			defer inbound.Close()
			client.SetDeadline(time.Now().Add(5 * time.Second))
			addr := server.Addr().(*net.TCPAddr)
			b := &JWN154qJ4{zb8f3cHWW8v: context.Background(), dZfoV0aXHW: addr.IP.String(), lAmttzTO8b: addr.Port}
			finished := make(chan struct{})
			go func() { b.xx3H26VL(inbound, prefix); close(finished) }()
			<-ready
			payload := bytes.Repeat([]byte("forward"), 3000)
			written := make(chan error, 1)
			go func() { _, e := client.Write(payload); written <- e }()
			got := make([]byte, len(prefix)+len(payload))
			if _, e := io.ReadFull(client, got); e != nil {
				t.Fatal(e)
			}
			if !bytes.Equal(got, append(append([]byte{}, prefix...), payload...)) {
				t.Fatal("prefix/relay changed")
			}
			if e := <-written; e != nil {
				t.Fatal(e)
			}
			client.Close()
			select {
			case <-finished:
			case <-time.After(5 * time.Second):
				t.Fatal("relay did not end")
			}
			<-done
		})
	}
}
func TestECHConfigurationPreservationAndErrors(t *testing.T) {
	b := &JWN154qJ4{}
	if e := b.dINraX_uOx5T(nil); e != nil {
		t.Fatal(e)
	}
	node := &L.Wr5SVP{}
	b.wiy6P3.Store(node)
	config := &tls.Config{EncryptedClientHelloKeys: []tls.EncryptedClientHelloKey{{PrivateKey: []byte("old")}}}
	if e := b.dINraX_uOx5T(config); e != nil || string(config.EncryptedClientHelloKeys[0].PrivateKey) != "old" {
		t.Fatal("empty preserved", e)
	}
	node.LKMSis7zkX = "!bad!"
	if e := b.dINraX_uOx5T(config); e == nil || string(config.EncryptedClientHelloKeys[0].PrivateKey) != "old" {
		t.Fatal("invalid preserved", e)
	}
	node.LKMSis7zkX = base64.StdEncoding.EncodeToString([]byte{0, 2, 1, 2, 0, 3, 3, 4, 5})
	if e := b.dINraX_uOx5T(config); e != nil || len(config.EncryptedClientHelloKeys) != 1 || !bytes.Equal(config.EncryptedClientHelloKeys[0].PrivateKey, []byte{1, 2}) {
		t.Fatal("key replacement", e)
	}
}
