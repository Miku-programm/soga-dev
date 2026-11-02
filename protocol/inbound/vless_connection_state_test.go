package inbound

import (
	CA "recovered.local/soga/util/rtx"
	"testing"
)

func TestVLESSOriginalCurrentHourSharedCount(t *testing.T) {
	v := &KHonGaqbIy_6{JWN154qJ4: &JWN154qJ4{fUxO7UYAH: 73}}
	v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
	old := userCountAt50C6DA0
	defer func() { userCountAt50C6DA0 = old }()
	v.bj10P7X()
	if userCountAt50C6DA0 != 73 {
		t.Fatal("close-path count")
	}
	v.fUxO7UYAH = 81
	v.sbbdX3aT6E()
	if userCountAt50C6DA0 != 81 {
		t.Fatal("open-path count")
	}
}
