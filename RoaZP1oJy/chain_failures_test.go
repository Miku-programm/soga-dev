package RoaZP1oJy

import (
	"bytes"
	"crypto/md5"
	"crypto/rc4"
	"encoding/binary"
	S "recovered.local/soga/ad5KEbw_EzWk"
	E "recovered.local/soga/hfOuV8fOe4L"
	"testing"
)

func TestProtocolFactoryOriginalRegistry(t *testing.T) {
	if len(protocolsAt5096970) != 9 {
		t.Fatal(len(protocolsAt5096970))
	}
	for name := range protocolsAt5096970 {
		c, e := HrQWlmA(name, nil)
		if e != nil || c == nil {
			t.Fatal(name, c, e)
		}
	}
	for _, name := range []E.Mu0vMIIs{"none", "AUTH_CHAIN_A", "auth_sha1_v4", ""} {
		c, e := HrQWlmA(name, nil)
		if c != nil || e == nil || e.Error() != "protocol "+string(name)+" not support" {
			t.Fatal(c, e)
		}
	}
}
func TestChainFailuresPreserveOrdering(t *testing.T) {
	key := []byte("1234567890abcdef")
	for _, tc := range []struct{ where, want string }{{"first", "hmac1 error"}, {"second", "hmac2 error"}, {"frame", "hmac wrong"}} {
		t.Run(tc.where, func(t *testing.T) {
			c := oPPzNk(nil).(*h2QD2cI1)
			defer c.Free()
			c.SetServerInfo(&S.L7E7H2aW8KVs{IawCjSkuTM4: key, SP2KC4lkkl3C: 16, Uklg6401: 123})
			header, seed, _, streamKey := chainAuthFixture("auth_chain_a", key, key, nil)
			enc, _ := rc4.NewCipher(streamKey)
			frame, _ := chainClientFrame(c, enc, key, seed, []byte("abc"), 1)
			switch tc.where {
			case "first":
				header[4] ^= 1
			case "second":
				header[32] ^= 1
			case "frame":
				frame[len(frame)-1] ^= 1
			}
			c.ReadWriter = &aesWire{input: bytes.NewReader(append(header, frame...))}
			n, e := c.Read(make([]byte, 10))
			if n != 0 || e == nil || e.Error() != tc.want {
				t.Fatal(n, e)
			}
			if c.zOQ_6vA == nil {
				t.Fatal("first HMAC state omitted")
			}
			if tc.where == "frame" {
				if !c.rmErDQPkS31O || c.aBfz5qx5HJ7 != 2 {
					t.Fatal("authenticated/counter side effects", c.rmErDQPkS31O, c.aBfz5qx5HJ7)
				}
			} else if c.rmErDQPkS31O || c.aBfz5qx5HJ7 != 1 {
				t.Fatal("early error mutated frame state")
			}
		})
	}
	// Oversize is rejected before buffer read or sequence-key advancement.
	c := oPPzNk(nil).(*h2QD2cI1)
	c.zOQ_6vA = make([]byte, 16)
	c.oHHJQje = &S.L7E7H2aW8KVs{}
	wire := make([]byte, 2)
	binary.LittleEndian.PutUint16(wire, 4096)
	c.ReadWriter = &aesWire{input: bytes.NewReader(wire)}
	if _, e := c.ibH0hkF(make([]byte, 1)); e == nil || e.Error() != "auth_chain_* over size" || c.aBfz5qx5HJ7 != 1 {
		t.Fatal(e, c.aBfz5qx5HJ7)
	}
}
func TestChainFirstWriteMutatesSpareCapacity(t *testing.T) {
	c := oPPzNk(nil).(*h2QD2cI1)
	defer c.Free()
	c.oHHJQje = &S.L7E7H2aW8KVs{}
	c.emo_yYrLZ0DL = make([]byte, 16)
	c.fOtNm3lE9H5 = make([]byte, 20)
	copy(c.fOtNm3lE9H5, []byte("1234567890abcdef"))
	c.pAvqa5e1EyuH, _ = rc4.NewCipher([]byte("1234567890abcdef"))
	c.ReadWriter = &aesWire{input: bytes.NewReader(nil)}
	backing := make([]byte, 8)
	copy(backing, []byte("abc"))
	if n, e := c.Write(backing[:3]); n != 3 || e != nil {
		t.Fatal(n, e)
	}
	if !bytes.Equal(backing[:5], []byte{180, 5, 'a', 'b', 'c'}) {
		t.Fatal(backing)
	}
	if c.jLK6jCJ00azL != 1460 || c.aL42JWaXIu != 2 {
		t.Fatal(c.jLK6jCJ00azL, c.aL42JWaXIu)
	}
	// Complete HMAC helper must not truncate to the separate ten-byte helper.
	if !bytes.Equal(c.hmcSfmiGu([]byte("key"), []byte("a"), []byte("b")), aesTag(md5.New, []byte("key"), []byte("ab"))) {
		t.Fatal("HMAC")
	}
	defer func() {
		if recover() != "please use WriteToUID()" {
			t.Error("WriteTo original panic")
		}
	}()
	c.WriteTo(nil, nil)
}
