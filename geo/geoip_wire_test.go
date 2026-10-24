package geo

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

// 下列期望值来自原 kbUmw5F78 机器码的受限仿真；没有执行原完整程序。
func TestGeoIPWireOriginalMachineCodeVectors(t *testing.T) {
	cases := []struct {
		data, code, result, errorName string
		panicExpected                 bool
		resultCap                     int
	}{
		{"", "CN", "", "not_found", false, 0},
		{"00", "CN", "", "file", false, 0},
		{"0a", "CN", "", "not_found", false, 0},
		{"0a01", "CN", "", "not_found", false, 0},
		{"0a0510", "CN", "", "file", false, 0},
		{"0a050a7f", "CN", "", "not_found", false, 0},
		{"0a050a0243", "CN", "", "length", false, 0},
		{"0a8080808080808080808001", "CN", "", "varint", false, 0},
		{"0a80808080808080808080", "CN", "", "not_found", false, 0},
		{"0affffffffffffffffff010a02434e", "CN", "", "", true, 0},
		{"0a090a02434e12030a0178", "CN", "0a02434e12030a0178", "", false, 9},
		{"0a090a02434e12030a0178", "cn", "0a02434e12030a0178", "", false, 9},
		{"0a090a02434e12030a0178", "US", "", "not_found", false, 0},
		{"0a090a02434e12030a01780a040a025553", "US", "0a025553", "", false, 4},
		{"0a85010a820141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "0a820141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141", "", false, 133},
		{"0a85010a8201414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141410a090a02434e12030a0178", "cn", "0a02434e12030a0178", "", false, 9},
		{"0a020a00", "", "", "not_found", false, 0},
		{"0a020a000a090a02434e12030a0178", "cn", "0a02434e12030a0178", "", false, 9},
		{"0a000a00", "", "", "not_found", false, 0},
		{"0a010a02434e", "ZZ", "", "file", false, 0},
		{"", "cn", "", "not_found", false, 0},
		{"0a", "cn", "", "not_found", false, 0},
		{"0a09", "cn", "", "not_found", false, 0},
		{"0a090a", "cn", "", "not_found", false, 0},
		{"0a090a02", "cn", "", "not_found", false, 0},
		{"0a090a0243", "cn", "", "length", false, 0},
		{"0a090a02434e", "cn", "", "length", false, 0},
		{"0a090a02434e12", "cn", "", "length", false, 0},
		{"0a090a02434e1203", "cn", "", "length", false, 0},
		{"0a090a02434e12030a", "cn", "", "length", false, 0},
		{"0a090a02434e12030a01", "cn", "", "length", false, 0},
	}
	errorsByName := map[string]error{"length": readLengthErrorAt5096F50, "file": invalidFileErrorAt5096F60, "varint": invalidVarintErrorAt5096F70, "not_found": codeNotFoundErrorAt5096F80}
	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			data, _ := hex.DecodeString(tc.data)
			// 原仿真输入 cap=len；保持相同切片容量以核对输出视图。
			data = data[:len(data):len(data)]
			defer func() {
				r := recover()
				if (r != nil) != tc.panicExpected {
					t.Fatalf("panic=%v expected=%v", r, tc.panicExpected)
				}
			}()
			got, err := kbUmw5F78(data, tc.code)
			if tc.panicExpected {
				t.Fatal("missing original panic")
			}
			if tc.errorName != "" {
				if got != nil || err != errorsByName[tc.errorName] {
					t.Fatalf("got %x %v; expected %s", got, err, tc.errorName)
				}
				return
			}
			expected, _ := hex.DecodeString(tc.result)
			if err != nil || !bytes.Equal(got, expected) || cap(got) != tc.resultCap {
				t.Fatalf("got %x cap%d %v; expected %x cap%d", got, cap(got), err, expected, tc.resultCap)
			}
		})
	}
}
