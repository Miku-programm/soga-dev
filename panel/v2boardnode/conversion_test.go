package v2boardnode

import (
	C "recovered.local/soga/config"
	"testing"
)

func TestOriginalVMessFirstFetchAndRestrictedCount(t *testing.T) {
	oldFirst := firstUsersAt378731C
	oldState := C.Q0SDIN()
	defer func() { firstUsersAt378731C = oldFirst; C.T15Jifo(oldState) }()
	C.T15Jifo(true)
	firstUsersAt378731C = true
	response := FaNycS0m{}
	for i := range 100 {
		response.GPETCf2 = append(response.GPETCf2, &Pn5vgsrEm5{EgeihzZFfj: int64(i)})
	}
	first := response.GetUsers()
	if len(first) != 100 || first[99].JAfXWwzYrW == nil || firstUsersAt378731C {
		t.Fatal("first-fetch behavior")
	}
	next := response.GetUsers()
	if len(next) != 88 || len(response.GPETCf2) != 88 {
		t.Fatal(len(next), len(response.GPETCf2))
	}
	if (&Pn5vgsrEm5{}).GetUser(88).JAfXWwzYrW != nil || (&Pn5vgsrEm5{}).GetUser(87).JAfXWwzYrW == nil {
		t.Fatal("credential restriction boundary")
	}
}
