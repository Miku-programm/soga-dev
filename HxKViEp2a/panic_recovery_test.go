package HxKViEp2a

import "testing"

func TestOriginalDeferredRecovery(t *testing.T) {
	for _, prefix := range []string{"", "test"} {
		finished := false
		func() { defer func() { finished = true }(); defer SlG_Y_mg(prefix); panic("fixture") }()
		if !finished {
			t.Fatal("deferred recovery")
		}
	}
	if SlG_Y_mg("outside defer") != nil {
		t.Fatal("recover outside panic")
	}
}
