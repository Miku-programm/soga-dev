package JPLmG42

import (
	"encoding/hex"
	"fmt"
	"testing"
	"unsafe"
)

var originalBlockTraces = []struct {
	Key, Plain, State, Cipher string
	Effective                 int
}{
	{"00254a6f94b9de03284d7297bce1062b", "000b16212c37424d", "25006f4ab99403de4d289772e1bc2b0629dfbc735006e59a792f0cc20056944a0c7835a15eca84f3ac189500be29e752946be7bd31080158532aa57cf0ce421910ceb0625402f9a69d4b32e0d7847b2904604da997f2c03b096552ae9cf7c520e59b772eca805c132dd989d1367fe1679cf7c5209bfbf79a40c4955d04604da9d76b297bce1f9e055402f9a6b2b5f031bee60f3f532aa57c123acff719428861be29e752fae554e77c0c56160c7835a1be8500aaf43da73f5006e59af6f6d720d5f92f9e4d2897721c4c91b5476b7980", "43379fe8227eebd2", 0},
	{"13385d82a7ccf1163b6085aacff4193e", "010c17222d38434e", "3813825dcca716f1603baa85f4cf3e194f05e299762c0bc19f5532e8267cbb7058c482edab16d03ff864e14c0a76339f2d047f56c9a099f0ecc23e158867dbb141ffe19385332ad8cf7c631108b6ac5a67c2b00bf954229e6cc7b510fe592783a9603cf38e4521d85c86c40c72ba1d1cfe592783297e9438de61162767c2b00bfae3f8499deec6e485332ad8811cbf00254ee9f8ecc23e15eeea375f81a900850a76339f1593089b30c0982658c482ede5e9da83ce172f15762c0bc1bf03b1fac2e60648603baa85a9d47ea2345837c5", "9803cccc3276fdf2", 0},
}

func TestOriginalBlockMachineCode(t *testing.T) {
	for i, row := range originalBlockTraces {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			key, _ := hex.DecodeString(row.Key)
			plain, _ := hex.DecodeString(row.Plain)
			want, _ := hex.DecodeString(row.Cipher)
			block, e := L3EF5OwH(key)
			if e != nil {
				t.Fatal(e)
			}
			state := block.(*pP3B3b6zEfB)
			if unsafe.Sizeof(*state) != 208 {
				t.Fatal("RTTI layout")
			}
			if hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(state)), 208)) != row.State {
				t.Fatal("key schedule does not match original machine code")
			}
			out := make([]byte, len(plain))
			block.Encrypt(out, plain)
			if hex.EncodeToString(out) != row.Cipher {
				t.Fatalf("cipher=%x want%x", out, want)
			}
			block.Decrypt(out, out)
			if hex.EncodeToString(out) != row.Plain {
				t.Fatal("decrypt")
			}
		})
	}
}
