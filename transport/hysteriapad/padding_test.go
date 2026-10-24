package hysteriapad

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestNativePaddingRules(t *testing.T) {
	if unsafe.Sizeof(FZBbzkJPWbrk{}) != 56 {
		t.Fatal("RTTI")
	}
	for _, bad := range []string{"", "garbage", "0=1-2", "stop= 8", "stop=8\r", "stop=9223372036854775808"} {
		if R_GY_FGx([]byte(bad)) != nil {
			t.Fatalf("accepted %q", bad)
		}
	}
	b := []byte("stop=-1\n0=9-9,c,5,0-3,3-0,-1-2,2-3-4,x-y,+4-+4,7-7\n1=3-7,7-3")
	p := R_GY_FGx(b)
	if p == nil || p.ZJzf4Os != ^uint32(0) || p.GtWiNfVs7j != fmt.Sprintf("%x", md5.Sum(b)) {
		t.Fatal("parse/wrap/hash")
	}
	if got := p.GenerateRecordPayloadSizes(0); !reflect.DeepEqual(got, []int{9, -1, 4, 7}) {
		t.Fatal(got)
	}
	if p.GenerateRecordPayloadSizes(99) != nil {
		t.Fatal("missing rule")
	}
	for i := 0; i < 100; i++ {
		v := p.GenerateRecordPayloadSizes(1)
		if len(v) != 2 || v[0] < 3 || v[0] >= 7 || v[1] < 3 || v[1] >= 7 {
			t.Fatal(v)
		}
	}
	// 输入缓冲不复制；映射与摘要已计算，不随之后的输入修改重新解析。
	b[0] = 'X'
	if p.KruzWnr9zxA[0] != 'X' {
		t.Fatal("copied input")
	}
	q := DefaultAt5098340.Load()
	if q == nil || q.ZJzf4Os != 8 || !reflect.DeepEqual(q.GenerateRecordPayloadSizes(0), []int{30}) {
		t.Fatal("original decoded default")
	}
	if !reflect.DeepEqual(q.GenerateRecordPayloadSizes(3)[:1], []int{9}) {
		t.Fatal("record3")
	}
}
