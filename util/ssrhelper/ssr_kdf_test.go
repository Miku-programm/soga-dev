package ssrhelper

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestNativeSSRKDFBoundaries(t *testing.T) {
	if Is8IqhjVE6("p", 0) != nil {
		t.Fatal("nil zero length")
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("negative key length must panic")
			}
		}()
		Is8IqhjVE6("p", -1)
	}()
	a := md5.Sum([]byte("password"))
	secondInput := append(append([]byte{}, a[:]...), []byte("password")...)
	b := md5.Sum(secondInput)
	want := append(append([]byte{}, a[:]...), b[:]...)
	for _, n := range []int{1, 15, 16, 17, 24, 32} {
		if got := Is8IqhjVE6("password", n); !bytes.Equal(got, want[:n]) || cap(got) < n {
			t.Fatal(n, got)
		}
	}
	// RFC 5869 test case 4 (SHA-1).
	key := bytes.Repeat([]byte{0x0b}, 11)
	salt, _ := hex.DecodeString("000102030405060708090a0b0c")
	info, _ := hex.DecodeString("f0f1f2f3f4f5f6f7f8f9")
	out := make([]byte, 42)
	THj8k5(key, salt, info, out)
	if hex.EncodeToString(out) != "085a01ea1b10f36933068b56efa5ad81a4f14b822f5b091568a9cdd4f155fda2c22e422478d305f3f896" {
		t.Fatal(out)
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("HKDF exhaustion must panic")
			}
		}()
		THj8k5(key, salt, info, make([]byte, 255*20+1))
	}()
}
