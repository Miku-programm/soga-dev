package ssrprotocol

import (
	"bytes"
	"crypto/md5"
	S "recovered.local/soga/protocol/ssrshared"
	"testing"
	"testing/synctest"
	"time"
)

// Both original authentication routines truncate current time and both bounds
// to uint32. Virtual time covers the 2106 wrap without mutating clock globals.
func TestNativeAuthenticationUint32Timestamp(t *testing.T) {
	for _, tc := range []struct {
		unix  int64
		valid bool
	}{{1<<32 + 172800, true}, {1<<32 + 2, false}, {1<<32 - 2, false}} {
		synctest.Test(t, func(t *testing.T) {
			time.Sleep(time.Unix(tc.unix, 0).Sub(time.Now()))
			key := []byte("1234567890abcdef")
			aesConn := jKGRC86Lx9g(nil).(*lKlYpdomLrhv)
			defer aesConn.Free()
			aesConn.SetServerInfo(&S.L7E7H2aW8KVs{IawCjSkuTM4: key, SP2KC4lkkl3C: 16})
			aesConn.ReadWriter = &aesWire{input: bytes.NewReader(aesAuthFixture(md5.New, "auth_aes128_md5", key, key, nil, []byte("test")))}
			data, err := aesConn.rTiayU()
			if tc.valid {
				if err != nil || string(data) != "test" {
					t.Fatal(tc.unix, string(data), err)
				}
			} else if err == nil || err.Error() != "utc invalid" {
				t.Fatal(tc.unix, err)
			}
			chain := oPPzNk(nil).(*h2QD2cI1)
			defer chain.Free()
			chain.SetServerInfo(&S.L7E7H2aW8KVs{IawCjSkuTM4: key, SP2KC4lkkl3C: 16})
			header, _, _, _ := chainAuthFixture("auth_chain_a", key, key, nil)
			chain.ReadWriter = &aesWire{input: bytes.NewReader(header)}
			err = chain.hAt43dMn3W()
			if tc.valid {
				if err != nil {
					t.Fatal(tc.unix, err)
				}
			} else if err == nil || err.Error() != "utc invalid" {
				t.Fatal(tc.unix, err)
			}
		})
	}
}
