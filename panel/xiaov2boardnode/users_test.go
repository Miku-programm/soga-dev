package xiaov2boardnode

import (
	C "recovered.local/soga/config"
	"reflect"
	"testing"
)

func TestOriginalUserCountAndCredentialGates(t *testing.T) {
	old := stateAt3787320
	stateAt3787320 = 1
	C.T15Jifo(true)
	defer func() { stateAt3787320 = old; C.T15Jifo(false) }()
	in := make(QgFZpfpg8DH, 90)
	for i := range in {
		in[i] = &NzbHXtDtzF{LMCS6ltENGT: int64(i + 1), DhIHEmq: "secret"}
	}
	first := in.GetUsers()
	if len(first) != 90 || first[89].JAfXWwzYrW == nil || stateAt3787320 != 0 {
		t.Fatal("initial bypass state must be consumed")
	}
	next := in.GetUsers()
	if len(next) != 88 || len(in) != 90 {
		t.Fatal("original range cutoff, not input mutation")
	}
	if in[88].GetUser(88).JAfXWwzYrW != nil || in[87].GetUser(87).JAfXWwzYrW == nil {
		t.Fatal("original credential cutoff")
	}
	if got := (QgFZpfpg8DH{}).GetUsers(); got == nil {
		t.Fatal("original nonnil empty")
	}
	if reflect.TypeOf(NzbHXtDtzF{}).Size() != 48 {
		t.Fatal("RTTI user layout")
	}
}
