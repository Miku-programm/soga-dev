package sogav1node

import (
	C "recovered.local/soga/config"
	"reflect"
	"testing"
)

func TestOriginalUserCountAndCredentialGates(t *testing.T) {
	old := stateAt378731D
	stateAt378731D = 1
	C.T15Jifo(true)
	defer func() { stateAt378731D = old; C.T15Jifo(false) }()
	in := make(RFBSFd3, 90)
	for i := range in {
		in[i] = &A4MyCjSpEhR{UZaO5D8ajd1: int64(i + 1), SvKI66: "secret"}
	}
	first := in.GetUsers()
	if len(first) != 90 || first[89].JAfXWwzYrW == nil || stateAt378731D != 0 {
		t.Fatal("initial bypass state must be consumed")
	}
	next := in.GetUsers()
	if len(next) != 90 || len(in) != 90 {
		t.Fatal("user list must not be truncated")
	}
	if in[88].GetUser(88).JAfXWwzYrW == nil || in[87].GetUser(87).JAfXWwzYrW == nil {
		t.Fatal("credentials must be present for all users")
	}
	if got := (RFBSFd3{}).GetUsers(); got == nil {
		t.Fatal("original nonnil empty")
	}
	if reflect.TypeOf(A4MyCjSpEhR{}).Size() != 56 {
		t.Fatal("RTTI user layout")
	}
}
