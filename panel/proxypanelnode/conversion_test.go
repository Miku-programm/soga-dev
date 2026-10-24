package proxypanelnode

import (
	C "recovered.local/soga/config"
	"strings"
	"testing"
)

func TestOriginalSSRSinglePortAndParseError(t *testing.T) {
	p := QVdbsbqVT3{AZlE6hE7M7: "8443"}
	n, e := p.ToPanelNode(1)
	if n != nil || e == nil || e.Error() != "请在 ssr 节点中启用单端口" {
		t.Fatal(n, e)
	}
	p.C8gP2P = 1
	p.AZlE6hE7M7 = "bad"
	n, e = p.ToPanelNode(1)
	if n != nil || e == nil || e.Error() != "port invalid: 0" {
		t.Fatal(n, e)
	}
	p.AZlE6hE7M7 = "-1"
	n, e = p.ToPanelNode(1)
	if e != nil || n.ASJxBbFS.PU1Di18 != -1 {
		t.Fatal(n, e)
	}
	if !strings.Contains(n.Z1TIgHOQ, `"port":"-1"`) {
		t.Fatal(n.Z1TIgHOQ)
	}
}
func TestOriginalSSRFirstFetchAndRestrictedCount(t *testing.T) {
	oldFirst := firstUsersAt378731A
	oldState := C.Q0SDIN()
	defer func() { firstUsersAt378731A = oldFirst; C.T15Jifo(oldState) }()
	C.T15Jifo(true)
	firstUsersAt378731A = true
	response := JpYAk64U2{}
	for i := range 100 {
		response.GPETCf2 = append(response.GPETCf2, &CXFPuaTkSm{QZuKzrg: int64(i)})
	}
	first := response.GetUsers()
	if len(first) != 100 || first[99].JAfXWwzYrW == nil || firstUsersAt378731A {
		t.Fatal("first-fetch behavior")
	}
	next := response.GetUsers()
	if len(next) != 88 || len(response.GPETCf2) != 88 {
		t.Fatal(len(next), len(response.GPETCf2))
	}
	if (&CXFPuaTkSm{}).GetUser(88).JAfXWwzYrW != nil || (&CXFPuaTkSm{}).GetUser(87).JAfXWwzYrW == nil {
		t.Fatal("credential restriction boundary")
	}
}
