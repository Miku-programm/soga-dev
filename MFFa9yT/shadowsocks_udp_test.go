package MFFa9yT

import (
	"context"
	"io"
	"net"
	D "recovered.local/soga/D_zCag1qTV"
	G "recovered.local/soga/G6bHD5JA"
	J "recovered.local/soga/J1P4KNvfI_J"
	N "recovered.local/soga/nzATvajt"
	"testing"
)

type ssDatagramReaderFixture struct {
	J.KsKg3yls
	calls, closed, freed int
	buffers              []*byte
}

func (p *ssDatagramReaderFixture) Close() error        { p.closed++; return nil }
func (p *ssDatagramReaderFixture) Free()               { p.freed++ }
func (p *ssDatagramReaderFixture) LocalAddr() net.Addr { return &net.UDPAddr{Port: 10000} }
func (p *ssDatagramReaderFixture) ReadFromProxy(buf []byte) (int, N.PgaNd3014U, net.Addr, *G.AHU2MRmjzQu, error) {
	p.calls++
	p.buffers = append(p.buffers, &buf[0])
	switch p.calls {
	case 1:
		return 0, nil, nil, nil, nil
	case 2:
		return copy(buf, "accepted"), nil, &net.UDPAddr{Port: 1}, G.NewDomainAddress("first.test", 53), nil
	case 3, 4:
		return copy(buf, "dropped"), nil, &net.UDPAddr{Port: 2}, G.NewDomainAddress("other.test", 53), nil
	default:
		return copy(buf, "error data ignored"), nil, nil, nil, io.EOF
	}
}
func TestShadowsocksNativeQueueOwnershipAndDrops(t *testing.T) {
	p := &ssDatagramReaderFixture{}
	q := make(chan J.AFui66gdB, 1)
	receiveShadowsocksDatagrams(context.Background(), p, q)
	if p.calls != 5 || p.closed != 1 || p.freed != 1 {
		t.Fatal(p)
	}
	got, ok := <-q
	if !ok || string(got.AGHsUHa) != "accepted" || got.PNliwRWkrU.GetHost() != "first.test" || got.Y8YsJ7QnEI3h != p {
		t.Fatal(got, ok)
	}
	defer D.PoolAt50961D8.Put(got.AGHsUHa)
	if _, ok = <-q; ok {
		t.Fatal("full queue failed to drop packets or close")
	}
	if p.buffers[0] != p.buffers[1] || p.buffers[1] == p.buffers[2] || p.buffers[2] != p.buffers[3] || p.buffers[3] != p.buffers[4] {
		t.Fatal("buffer ownership does not follow send success")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	p = &ssDatagramReaderFixture{}
	q = make(chan J.AFui66gdB, 1)
	receiveShadowsocksDatagrams(ctx, p, q)
	if p.calls != 0 || p.closed != 1 || p.freed != 1 {
		t.Fatal(p)
	}
	if _, ok = <-q; ok {
		t.Fatal("not closed")
	}
}
