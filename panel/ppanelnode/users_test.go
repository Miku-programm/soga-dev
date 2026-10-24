package ppanelnode

import (
	C "recovered.local/soga/config"
	"reflect"
	"testing"
)

func TestOriginalUserCountAndCredentialGates(t *testing.T) {
	old := stateAt378731E
	stateAt378731E = 1
	C.T15Jifo(true)
	defer func() { stateAt378731E = old; C.T15Jifo(false) }()
	in := make(CgPle0Z, 90)
	for i := range in {
		in[i] = &Ya7brDb{SlNaxp: int64(i + 1), GUDL30zd: "secret"}
	}
	first := in.GetUsers()
	if len(first) != 90 || first[89].JAfXWwzYrW == nil || stateAt378731E != 0 {
		t.Fatal("initial bypass state must be consumed")
	}
	next := in.GetUsers()
	if len(next) != 88 || len(in) != 90 {
		t.Fatal("original range cutoff, not input mutation")
	}
	if in[88].GetUser(88).JAfXWwzYrW != nil || in[87].GetUser(87).JAfXWwzYrW == nil {
		t.Fatal("original credential cutoff")
	}
	if got := (CgPle0Z{}).GetUsers(); got == nil {
		t.Fatal("original nonnil empty")
	}
	if reflect.TypeOf(Ya7brDb{}).Size() != 56 {
		t.Fatal("RTTI user layout")
	}
}
