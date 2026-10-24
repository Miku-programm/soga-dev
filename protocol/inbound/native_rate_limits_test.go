package inbound

import (
	"bytes"
	"crypto/rand"
	"io"
	"net"
	"testing"
	"time"
)

// Exercise the original user and node byte accounting through authenticated
// native SS, rather than calling the limiter directly. The original limiter
// permits one second of burst and charges after copying a buffer. Consequently
// this checks a conservative lower bound, not an exact scheduling duration.
func TestNativeUserAndNodeDownloadRate(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	liveConfiguration(t, "127.0.0.1")
	for _, scope := range []string{"user", "node"} {
		t.Run(scope, func(t *testing.T) {
			node := liveNode()
			if scope == "node" {
				node.TtWKyh = 8
			}
			v, data := liveCore(t, node)
			if scope == "user" {
				data.VTme1SuHL6i_[0].LmOKwjha = 8
				liveUpdate(t, v, data)
			}
			target, err := net.Listen("tcp4", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			defer target.Close()
			payload := bytes.Repeat([]byte("native-rate-data"), 3*65536)
			// 16 bytes per phrase gives exactly 3 MiB, versus an 8 Mib/s
			// bucket with a 1 MiB initial burst.
			if len(payload) != 3<<20 {
				t.Fatal("fixture payload size", len(payload))
			}
			release := make(chan struct{})
			defer close(release)
			done := make(chan error, 1)
			go func() {
				conn, e := target.Accept()
				if e != nil {
					done <- e
					return
				}
				defer conn.Close()
				conn.SetDeadline(time.Now().Add(10 * time.Second))
				var request [4]byte
				if _, e = io.ReadFull(conn, request[:]); e == nil {
					_, e = io.Copy(conn, bytes.NewReader(payload))
				}
				done <- e
				<-release
			}()
			client := liveClient(t, v, false, "live-key", ssFixtureAddress(1, target.Addr().(*net.TCPAddr).Port))
			client.SetDeadline(time.Now().Add(10 * time.Second))
			salt := make([]byte, 16)
			if _, err = rand.Read(salt); err != nil {
				t.Fatal(err)
			}
			request := append(append([]byte{}, client.target...), []byte("send")...)
			wire := append(salt, ssClientFrame(liveAEAD("live-key", salt), make([]byte, 12), request)...)
			start := time.Now()
			if _, err = client.Write(wire); err != nil {
				t.Fatal(err)
			}
			if _, err = io.ReadFull(client, salt); err != nil {
				t.Fatal(err)
			}
			aead, nonce := liveAEAD("live-key", salt), make([]byte, 12)
			var got []byte
			for len(got) < len(payload) {
				chunk, e := ssReadClientFrame(client, aead, nonce)
				if e != nil {
					t.Fatalf("download %d/%d: %v", len(got), len(payload), e)
				}
				got = append(got, chunk...)
			}
			elapsed := time.Since(start)
			if !bytes.Equal(got, payload) {
				t.Fatal("decrypted target payload differs")
			}
			if elapsed < 1700*time.Millisecond {
				t.Fatalf("%s limit did not constrain native relay: 3 MiB in %v", scope, elapsed)
			}
			if e := <-done; e != nil {
				t.Fatal(e)
			}
			t.Logf("%s: authenticated 3 MiB download in %v", scope, elapsed)
		})
	}
}
