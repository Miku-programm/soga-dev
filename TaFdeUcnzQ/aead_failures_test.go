package TaFdeUcnzQ

import (
	"bytes"
	"errors"
	"io"
	H "recovered.local/soga/hfOuV8fOe4L"
	"testing"
)

// Original 0x1c9d5c0 increments the nonce after Open, including failure,
// but not after a failed ReadFull. Both direct and buffered payload branches
// have the same ordering. The independent frame builder supplies the wire.
func TestNativeAEADPayloadFailureState(t *testing.T) {
	for _, name := range []H.ZvAxOt1l7Ar{"aes-128-gcm", "chacha20-ietf-poly1305"} {
		for _, dstSize := range []int{3, 10} {
			for _, truncated := range []bool{false, true} {
				c, _ := makeFrame(t, name, 10, false)
				a := c.s_53ql
				raw := a.eMYicxBI.(*bytes.Buffer)
				wire := bytes.Clone(raw.Bytes())
				if truncated {
					wire = wire[:len(wire)-1]
				} else {
					wire[len(wire)-1] ^= 1
				}
				a.eMYicxBI = bytes.NewBuffer(wire)
				n, err := c.Read(make([]byte, dstSize))
				if n != 0 || err == nil {
					t.Fatalf("%s/%d/%v: %d %v", name, dstSize, truncated, n, err)
				}
				wantNonce := byte(2)
				if truncated {
					wantNonce = 1
					if !errors.Is(err, io.ErrUnexpectedEOF) {
						t.Fatal(err)
					}
				}
				if a.qaeiJJzh3hj[0] != wantNonce || a.g9tmWrw57CtG != -1 || a.jjwGY3_N != nil {
					t.Fatalf("%s: nonce=%x offset=%d pending=%x", name, a.qaeiJJzh3hj, a.g9tmWrw57CtG, a.jjwGY3_N)
				}
				a.Free()
			}
		}
	}
}

// Original UDP 0x1c9e640 discards source/count on read, size, or authentication
// errors and never advances the nonce. The packet fixture models data+error.
func TestNativeAEADDatagramFailureState(t *testing.T) {
	for _, name := range []H.ZvAxOt1l7Ar{"aes-128-gcm", "chacha20-ietf-poly1305"} {
		for _, failure := range []string{"read", "salt", "tag"} {
			pc := &packetFixture{}
			c, err := DP5vq3(infoFor(name), pc)
			if err != nil {
				t.Fatal(err)
			}
			if _, err = c.WriteTo([]byte("payload"), nil); err != nil {
				t.Fatal(err)
			}
			pc.data = bytes.Clone(pc.written)
			sentinel := errors.New("data with error")
			switch failure {
			case "read":
				pc.err = sentinel
			case "salt":
				pc.data = pc.data[:7]
			case "tag":
				pc.data[len(pc.data)-1] ^= 1
			}
			n, addr, err := c.ReadFrom(make([]byte, 128))
			if n != 0 || addr != nil || err == nil {
				t.Fatalf("%s/%s: %d %v %v", name, failure, n, addr, err)
			}
			if failure == "read" && err != sentinel {
				t.Fatal(err)
			}
			if failure == "salt" && err.Error() != "salt len not enough: 7" {
				t.Fatal(err)
			}
			for _, b := range c.s_53ql.qaeiJJzh3hj {
				if b != 0 {
					t.Fatalf("UDP nonce advanced: %x", c.s_53ql.qaeiJJzh3hj)
				}
			}
			c.s_53ql.Free()
		}
	}
}
