package limit

import (
	"context"
	"recovered.local/soga/panel/types"
	"reflect"
	"testing"
)

func TestOriginalResolverReplacementAndLayout(t *testing.T) {
	typ := reflect.TypeFor[O0UldNL4Eq]()
	if typ.Size() != 64 {
		t.Fatal(typ.Size())
	}
	for i, offset := range []uintptr{0, 16, 24, 32, 40} {
		if typ.Field(i).Offset != offset {
			t.Fatal(typ.Field(i))
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s := &O0UldNL4Eq{RCaEGDlm: ctx}
	s.UpdateResolver(nil)
	first := s.GetResolver()
	if first == nil {
		t.Fatal("missing resolver")
	}
	s.UpdateResolver(types.U863Mm01Nre{})
	if s.GetResolver() != first {
		t.Fatal("nil/empty rules unequal")
	}
	rules := types.U863Mm01Nre{&types.Axa8uKIB4JL{K8BNfziIgb: []string{"local"}, CM0eMsqyP: "ipv4_only", CZ80ik: []string{"full:test.invalid"}}}
	s.UpdateResolver(rules)
	second := s.GetResolver()
	if second == first {
		t.Fatal("new rules ignored")
	}
	copied := types.U863Mm01Nre{&types.Axa8uKIB4JL{K8BNfziIgb: []string{"local"}, CM0eMsqyP: "ipv4_only", CZ80ik: []string{"full:test.invalid"}}}
	s.UpdateResolver(copied)
	if s.GetResolver() != second {
		t.Fatal("equal rules unexpectedly replaced")
	}
	copied[0].CM0eMsqyP = "ipv6_only"
	s.UpdateResolver(copied)
	if s.GetResolver() == second {
		t.Fatal("changed rule ignored")
	}
	if s.lNzIhzjHYz[0] != copied[0] {
		t.Fatal("rule pointers should be shared")
	}
}
