package types

import (
	"bytes"
	"encoding/base64"
	"encoding/pem"
	"strings"
	"testing"
)

func TestECHWireInputs(t *testing.T) {
	wire := []byte{0, 2, 1, 2, 0, 3, 3, 4, 5, 0, 0, 0, 0}
	enc := base64.StdEncoding.EncodeToString(wire)
	for _, input := range []string{enc, " \n" + enc + "\t", string(pem.EncodeToMemory(&pem.Block{Type: "ARBITRARY", Bytes: wire})) + "ignored trailing text"} {
		keys, err := YUH4nAWc(input)
		if err != nil {
			t.Fatal(err)
		}
		if len(keys) != 2 || !bytes.Equal(keys[0].PrivateKey, []byte{1, 2}) || !bytes.Equal(keys[0].Config, []byte{3, 4, 5}) || !keys[0].SendAsRetry || !keys[1].SendAsRetry || len(keys[1].Config) != 0 || len(keys[1].PrivateKey) != 0 {
			t.Fatalf("%+v", keys)
		}
	}
	for _, wire := range [][]byte{{0}, {0, 2, 1}, {0, 0}, {0, 0, 0}, {0, 0, 0, 1}, {0, 0, 0, 0, 0}} {
		keys, err := YUH4nAWc(base64.StdEncoding.EncodeToString(wire))
		if keys != nil || err == nil || err.Error() != "ech key 解析失败" {
			t.Fatalf("%x: %v %v", wire, keys, err)
		}
	}
	if keys, err := YUH4nAWc(" \t"); keys != nil || err == nil || err.Error() != "ech key 为空" {
		t.Fatalf("%v %v", keys, err)
	}
	if _, err := YUH4nAWc("!!"); err == nil || !strings.HasPrefix(err.Error(), "ech key 格式错误: ") {
		t.Fatal(err)
	}
}
