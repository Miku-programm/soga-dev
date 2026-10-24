package ssrprotocol

import (
	"bytes"
	T "recovered.local/soga/protocol/ssrcipher"
	S "recovered.local/soga/protocol/ssrshared"
	"testing"
	"unsafe"
)

func TestOriginCipherRoundTrip(t *testing.T) {
	if unsafe.Sizeof(f_E7c41{}) != 32 {
		t.Fatal("RTTI")
	}
	var wire bytes.Buffer
	out, e := T.VavyqXsqy9V(&S.L7E7H2aW8KVs{MY3lagqY: "test", L6BrlCXd: "aes-128-cfb"}, &wire)
	if e != nil {
		t.Fatal(e)
	}
	defer out.Free()
	c := jyvuea4xsHn(out)
	c.SetServerInfo(nil)
	if c.CurUID() != 0 || c.CID() != 0 || c.GetOverhead() != 0 {
		t.Fatal("origin metadata")
	}
	if _, e = c.Write([]byte("native origin")); e != nil {
		t.Fatal(e)
	}
	in, e := T.VavyqXsqy9V(&S.L7E7H2aW8KVs{MY3lagqY: "test", L6BrlCXd: "aes-128-cfb"}, &wire)
	if e != nil {
		t.Fatal(e)
	}
	defer in.Free()
	b := make([]byte, 64)
	n, e := jyvuea4xsHn(in).Read(b)
	if e != nil || string(b[:n]) != "native origin" {
		t.Fatal(n, e, string(b[:n]))
	}
}
