package sTqn2l9

import (
	C "recovered.local/soga/cNT24rG4F"
	"testing"
)

func TestOriginalTrojanFirstFetchAndRestrictedCount(t *testing.T) {
	oldFirst := firstUsersAt378731B
	oldState := C.Q0SDIN()
	defer func() { firstUsersAt378731B = oldFirst; C.T15Jifo(oldState) }()
	C.T15Jifo(true)
	firstUsersAt378731B = true
	response := BoM67db{}
	for i := range 100 {
		response.GPETCf2 = append(response.GPETCf2, &PqISXuLk{Jrq3niigv: int64(i)})
	}
	first := response.GetUsers()
	if len(first) != 100 || first[99].JAfXWwzYrW == nil || firstUsersAt378731B {
		t.Fatal("first-fetch behavior")
	}
	next := response.GetUsers()
	if len(next) != 88 || len(response.GPETCf2) != 88 {
		t.Fatal(len(next), len(response.GPETCf2))
	}
	if (&PqISXuLk{}).GetUser(88).JAfXWwzYrW != nil || (&PqISXuLk{}).GetUser(87).JAfXWwzYrW == nil {
		t.Fatal("credential restriction boundary")
	}
}
