package LFWTTXF7JR

import "testing"

func TestOriginalSSRCarrierReplacement(t *testing.T) {
	n := new(Dv6QJyFW6bU)
	ignored := &Aga91y{JAfXWwzYrW: &Yn8BJoL{SfiEGpJo: "uuid"}}
	n.AddSSRUser(ignored)
	a := &Aga91y{WZ89pTAOrHK4: 1, JAfXWwzYrW: &QcqpSL3i6ri{YrjVl_2g_S: 10, YyAgnMatmj: "one"}}
	b := &Aga91y{WZ89pTAOrHK4: 2, JAfXWwzYrW: &QcqpSL3i6ri{YrjVl_2g_S: 20, YyAgnMatmj: "two"}}
	replacement := &Aga91y{WZ89pTAOrHK4: 3, JAfXWwzYrW: &QcqpSL3i6ri{YrjVl_2g_S: 10, YyAgnMatmj: "changed"}}
	n.AddSSRUser(a)
	n.AddSSRUser(b)
	n.AddSSRUser(replacement)
	if len(n.kJbZdr) != 2 || n.kJbZdr[0] != replacement || n.kJbZdr[1] != b {
		t.Fatal(n.kJbZdr)
	}
	n.C33jbIzI = 20
	n.AddSSRUser(a)
	if a.JAfXWwzYrW.(*QcqpSL3i6ri).YrjVl_2g_S != 20 || len(n.kJbZdr) != 2 || n.kJbZdr[1] != a {
		t.Fatal("port override and replacement")
	}
}
