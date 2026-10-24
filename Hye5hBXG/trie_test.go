package Hye5hBXG

import (
	"testing"
	"unsafe"
)

func TestTriePrecedence(t *testing.T) {
	exact := &YhrM38Rda{HJKQN58GQYL: "exact"}
	wild := &YhrM38Rda{HJKQN58GQYL: "wild"}
	suffix := &YhrM38Rda{HJKQN58GQYL: "suffix"}
	t1 := NewEmptyTrie()
	t1.root.xMLHHG1nh["com"] = &YhrM38Rda{xMLHHG1nh: map[string]*YhrM38Rda{"example": {xMLHHG1nh: map[string]*YhrM38Rda{"www": exact, "*": wild, "": suffix}}}}
	for _, test := range []struct {
		s    string
		want *YhrM38Rda
	}{{"www.example.com", exact}, {"x.example.com", wild}, {"a.b.example.com", suffix}, {"example.com", nil}, {"x.example.com.", nil}, {"x.EXAMPLE.com", nil}, {".example.com", nil}} {
		if got := t1.Search(test.s); got != test.want {
			t.Fatal(test.s, got, test.want)
		}
	}
	if unsafe.Sizeof(*exact) != 24 || unsafe.Sizeof(*t1) != 8 {
		t.Fatal("layout")
	}
}
func TestDomainSplitQuirks(t *testing.T) {
	for _, test := range []struct {
		s  string
		ok bool
	}{{"", false}, {".", false}, {"a..b", false}, {".a", true}, {"a", true}, {"a.", false}} {
		_, ok := Pbem1Ia(test.s)
		if ok != test.ok {
			t.Fatal(test.s, ok)
		}
	}
}
