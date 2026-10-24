package routing

import (
	"encoding/binary"
	"hash/fnv"
	"net"
	"testing"
)

func TestSelectorsHashRawBytes(t *testing.T) {
	routes := make(J_enG0wE, 13)
	for i := range routes {
		routes[i] = new(GyyO51)
	}
	ips := []net.IP{make(net.IP, 0), {127, 0, 0, 1}, net.ParseIP("127.0.0.1"), net.ParseIP("2001:db8::1")}
	for _, ip := range ips {
		h := fnv.New32a()
		h.Write(ip)
		want := routes[int(h.Sum32())%len(routes)]
		if (&twxaE4km_U{}).Select(routes, ip, 0) != want {
			t.Fatalf("raw IP hash mismatch for %x", []byte(ip))
		}
	}
	for _, id := range []int64{1, 256, 0x0102030405060708} {
		var data [8]byte
		binary.BigEndian.PutUint64(data[:], uint64(id))
		h := fnv.New32a()
		h.Write(data[:])
		if (&gMxUOGd{}).Select(routes, nil, id) != routes[int(h.Sum32())%len(routes)] {
			t.Fatal("user ID hash byte order mismatch")
		}
	}
}
func TestRoundRobinStartsAtOneAndWraps(t *testing.T) {
	routes := J_enG0wE{new(GyyO51), new(GyyO51), new(GyyO51)}
	s := new(bLJQqABD)
	for _, index := range []int{1, 2, 0, 1} {
		if s.Select(routes, nil, 0) != routes[index] {
			t.Fatal("round robin order mismatch")
		}
	}
	s.v1TcNPAI0.Store(^uint64(0))
	if s.Select(routes, nil, 0) != routes[0] {
		t.Fatal("atomic counter wrap must select index zero")
	}
}
func TestSingleRouteBypassesHashAndRandom(t *testing.T) {
	out := new(GyyO51)
	routes := J_enG0wE{out}
	for _, s := range []ZLdFKPeDrCJ{new(h1kE2OhgN), new(twxaE4km_U), new(gMxUOGd)} {
		if s.Select(routes, nil, -1) != out {
			t.Fatal("single route changed")
		}
	}
}
func TestEmptyRoutesPanicAndRoundRobinStillIncrements(t *testing.T) {
	s := new(bLJQqABD)
	func() {
		defer func() {
			if recover() == nil {
				t.Error("empty route selection must panic")
			}
		}()
		s.Select(nil, nil, 0)
	}()
	if s.v1TcNPAI0.Load() != 1 {
		t.Fatal("empty round robin must increment before divide panic")
	}
}
