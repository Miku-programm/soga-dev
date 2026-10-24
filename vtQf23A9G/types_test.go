package vtQf23A9G

import (
	"net"
	B "recovered.local/soga/BVM7xtVwiJ"
	C "recovered.local/soga/CAodbhZf"
	H "recovered.local/soga/HHGbXWqYaYt_"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 原样本 all_types.json 的全部字段类型；标准库/已恢复依赖使用已验证的命名映射。
func TestOriginalRTTIFieldTypes(t *testing.T) {
	for _, tc := range []struct {
		v      any
		fields []reflect.Type
	}{
		{BIkgBMz3M{}, []reflect.Type{
			reflect.TypeFor[net.Conn](),
			reflect.TypeFor[sync.Mutex](),
			reflect.TypeFor[map[uint32]*CCXLCurYiT7](),
			reflect.TypeFor[atomic.Uint32](),
			reflect.TypeFor[sync.RWMutex](),
			reflect.TypeFor[sync.Once](),
			reflect.TypeFor[chan struct{}](),
			reflect.TypeFor[func()](),
			reflect.TypeFor[func()](),
			reflect.TypeFor[sync.Mutex](),
			reflect.TypeFor[uint64](),
			reflect.TypeFor[time.Time](),
			reflect.TypeFor[*C.EahpvU_iseQ[*H.FZBbzkJPWbrk]](),
			reflect.TypeFor[uint8](),
			reflect.TypeFor[bool](),
			reflect.TypeFor[bool](),
			reflect.TypeFor[bool](),
			reflect.TypeFor[[]uint8](),
			reflect.TypeFor[atomic.Uint32](),
			reflect.TypeFor[func(*CCXLCurYiT7)](),
		}},
		{CCXLCurYiT7{}, []reflect.Type{
			reflect.TypeFor[uint32](),
			reflect.TypeFor[*BIkgBMz3M](),
			reflect.TypeFor[*B.P848NOz3su9](),
			reflect.TypeFor[*B.EjFsCHfka9Zc](),
			reflect.TypeFor[B.TyIy_s](),
			reflect.TypeFor[sync.Once](),
			reflect.TypeFor[func()](),
			reflect.TypeFor[error](),
			reflect.TypeFor[sync.Once](),
		}},
	} {
		typ := reflect.TypeOf(tc.v)
		for i, want := range tc.fields {
			if f := typ.Field(i); f.Type != want {
				t.Fatalf("%s.%s=%s want %s", typ, f.Name, f.Type, want)
			}
		}
	}
}
