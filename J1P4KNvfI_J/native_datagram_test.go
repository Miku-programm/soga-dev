package J1P4KNvfI_J

import (
	"context"
	"errors"
	"io"
	"net"
	G "recovered.local/soga/G6bHD5JA"
	H "recovered.local/soga/HxKViEp2a"
	C "recovered.local/soga/cNT24rG4F"
	N "recovered.local/soga/nzATvajt"
	"reflect"
	"testing"
	"time"
)

type nativePacketFixture struct {
	KsKg3yls
	address  net.Addr
	writes   []nativePacketWrite
	deadline time.Time
	writeErr error
}
type nativePacketWrite struct {
	data           string
	user           N.PgaNd3014U
	source, remote net.Addr
}

func (p *nativePacketFixture) LocalAddr() net.Addr { return p.address }
func (p *nativePacketFixture) SetWriteDeadline(d time.Time) error {
	p.deadline = d
	return errors.New("ignored deadline")
}
func (p *nativePacketFixture) WriteToProxy(b []byte, u N.PgaNd3014U, s, r net.Addr) (int, error) {
	p.writes = append(p.writes, nativePacketWrite{string(b), u, s, r})
	return 0, p.writeErr
}

type nativeResponseFixture struct {
	net.PacketConn
	reads        int
	readDeadline time.Time
	from         net.Addr
	cancel       context.CancelFunc
	final        error
}

func (p *nativeResponseFixture) SetReadDeadline(d time.Time) error {
	p.readDeadline = d
	return errors.New("ignored deadline")
}
func (p *nativeResponseFixture) ReadFrom(b []byte) (int, net.Addr, error) {
	p.reads++
	switch p.reads {
	case 1:
		return copy(b, "payload"), p.from, nil
	case 2:
		return 0, p.from, nil
	default:
		if p.cancel != nil {
			p.cancel()
		}
		return copy(b, "discard"), p.from, p.final
	}
}

func TestNativeDatagramSessionLayoutAndAddresses(t *testing.T) {
	typ := reflect.TypeFor[AFui66gdB]()
	offsets := []uintptr{0, 16, 32, 40, 56}
	if typ.Size() != 80 {
		t.Fatal(typ)
	}
	for i, o := range offsets {
		if typ.Field(i).Offset != o {
			t.Fatal(typ.Field(i))
		}
	}
	ip := net.IPv4(192, 0, 2, 3)
	local := &nativePacketFixture{address: &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 3000}}
	target := G.NewDomainAddress("example.test", 53)
	p := AFui66gdB{Y8YsJ7QnEI3h: local, AIO3Qb: &net.TCPAddr{IP: ip, Port: 4000}, PNliwRWkrU: target}
	s := TgFd4ag1(p)
	if s.YlxwftjCvz_ != 1 || s.Iayb4V != target || s.Y72kDV35 != 4000 || s.X1gBTWHQO99 != 3000 || s.Uqzvq94e != "" || s.C66t8X != "" || &s.KeHzUwuXWh0[0] != &ip[0] {
		t.Fatal(s)
	}
	p.Y8YsJ7QnEI3h = nil
	p.AIO3Qb = nil
	z := TgFd4ag1(p)
	if z.KeHzUwuXWh0 != nil || z.L4SIW3qbjS != nil || z.E3L0l8hZ <= s.E3L0l8hZ {
		t.Fatal(z)
	}
}

func TestNativeDatagramForwardZeroPayloadShortWriteAndError(t *testing.T) {
	source := &net.UDPAddr{IP: net.IPv4(192, 0, 2, 4), Port: 1000}
	from := &net.UDPAddr{IP: net.IPv4(198, 51, 100, 2), Port: 53}
	stop := errors.New("read end")
	r := &nativeResponseFixture{from: from, final: stop}
	p := &nativePacketFixture{}
	start := time.Now()
	err := Jv9wkYGl(context.Background(), r, source, p, nil)
	if err != stop || len(p.writes) != 2 || p.writes[0].data != "payload" || p.writes[1].data != "" || p.writes[0].source != source || p.writes[0].remote != from {
		t.Fatal(err, p.writes)
	}
	if p.deadline.Before(start.Add(time.Minute)) || p.deadline.After(time.Now().Add(time.Minute)) {
		t.Fatal(p.deadline)
	}
	delta := time.Duration(C.UDPTimeoutMinutes()) * time.Minute
	if r.readDeadline.Before(start.Add(delta)) || r.readDeadline.After(time.Now().Add(delta)) {
		t.Fatal(r.readDeadline)
	}
	p = &nativePacketFixture{writeErr: io.ErrClosedPipe}
	r = &nativeResponseFixture{from: from}
	if e := Jv9wkYGl(context.Background(), r, source, p, nil); e != io.ErrClosedPipe || r.reads != 1 {
		t.Fatal(e, r.reads)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if e := Jv9wkYGl(ctx, nil, nil, nil, nil); e != H.ErrorAt5097800 || errors.Is(e, context.Canceled) {
		t.Fatal(e)
	}
}
