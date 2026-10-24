package cNT24rG4F

import "testing"

func TestVmessModeBridgeReadsOriginalGlobal(t *testing.T) {
	saved := configAt50A02C0.dZWc0A2uIe
	defer func() { configAt50A02C0.dZWc0A2uIe = saved }()
	for _, mode := range []F6DuWgn9N{"aead", "md5", "auto", ""} {
		configAt50A02C0.dZWc0A2uIe = mode
		if VmessAuthMode() != mode {
			t.Fatal("detached config state")
		}
	}
}
