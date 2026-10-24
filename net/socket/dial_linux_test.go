//go:build linux

package socket

import (
	"context"
	"net"
	"strings"
	"testing"
	"unsafe"
)

func TestDialFailurePrecedenceAndNetwork(t *testing.T) {
	if unsafe.Sizeof(bW961oCfn3t{}) != 40 {
		t.Fatal("RTTI layout")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c, err := I_y9nlK1q8(ctx, "unix", "x")
	if c != nil || err == nil || err.Error() != "network invalid" {
		t.Fatal(c, err)
	}
	c, err = I_y9nlK1q8(ctx, "tcp", "127.0.0.1:1")
	if c != nil || err == nil {
		t.Fatal(c, err)
	}
	if _, ok := err.(*net.OpError); !ok {
		t.Fatalf("IPv4 dialing error must beat disabled IPv6: %T %v", err, err)
	}
	c, err = I_y9nlK1q8(ctx, "tcp", "[::1]:1")
	if c != nil || err == nil || err.Error() != "ip version error" {
		t.Fatal("both DNS failures must prefer IPv4", c, err)
	}
	c, err = I_y9nlK1q8(ctx, "tcp6", "[::1]:1")
	if c != nil || err == nil || !strings.Contains(err.Error(), "ipv6 disabled") {
		t.Fatal(c, err)
	}
}
