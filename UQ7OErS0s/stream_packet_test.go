package UQ7OErS0s

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	M "recovered.local/soga/G7A7csgrs62p"
	E "recovered.local/soga/LFWTTXF7JR"
	R "recovered.local/soga/RoaZP1oJy"
	S "recovered.local/soga/ad5KEbw_EzWk"
	U "recovered.local/soga/nzATvajt"
	"testing"
	"time"
	"unsafe"
)

func wrapperInfo(t *testing.T, mode E.LaTpsgYQ7) (*S.L7E7H2aW8KVs, *U.LJJkaD) {
	t.Helper()
	// All-zero SSR credential fields avoid altering package configuration, while
	// the real constructor still initializes the accounting and limiter state.
	user, err := U.PXHoe5(&E.Aga91y{WZ89pTAOrHK4: 42, JAfXWwzYrW: &E.QcqpSL3i6ri{}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	info := &S.L7E7H2aW8KVs{MY3lagqY: "password", L6BrlCXd: "none", C8Qsa8G8XMw0: "origin", UNRudV1377x9: "plain", NxYA16xMfF: mode, Uklg6401: 42, EFLMTH: "user.example", N3pMXuwX85c: new(M.GTFAWF[int64, *U.LJJkaD]), RLXBf_luGj: new(M.GTFAWF[string, *U.LJJkaD])}
	info.N3pMXuwX85c.Set(42, user)
	info.N3pMXuwX85c.Set(0, user)
	info.RLXBf_luGj.Set("user.example", user)
	return info, user
}
func tcpPair(t *testing.T) (net.Conn, net.Conn) {
	t.Helper()
	l, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer l.Close()
	a, e := net.Dial("tcp4", l.Addr().String())
	if e != nil {
		t.Fatal(e)
	}
	b, e := l.Accept()
	if e != nil {
		a.Close()
		t.Fatal(e)
	}
	a.SetDeadline(time.Now().Add(3 * time.Second))
	b.SetDeadline(time.Now().Add(3 * time.Second))
	t.Cleanup(func() { a.Close(); b.Close() })
	return a, b
}

var addressFixture = []byte{3, 12, 'e', 'x', 'a', 'm', 'p', 'l', 'e', '.', 't', 'e', 's', 't', 0x01, 0xbb}

func TestNativeSSRStreamWrapperTCP(t *testing.T) {
	if unsafe.Sizeof(Z7PK9WJdYeE{}) != 80 || unsafe.Sizeof(AXWGpkYq{}) != 48 {
		t.Fatal("RTTI")
	}
	for _, mode := range []E.LaTpsgYQ7{0, 1, 2} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			client, server := tcpPair(t)
			info, user := wrapperInfo(t, mode)
			client.Write(append(append([]byte{}, addressFixture...), []byte("payload")...))
			c, err := RfxaTYLqLUJ(server, info)
			if err != nil {
				t.Fatal(err)
			}
			defer c.Free()
			if c.GetUser() != user || c.GetServerInfo() != info || c.NetworkType() != 0 || c.CurUID() != 42 || c.CID() != 0 || c.Host() != "user.example" || c.GetDestAddr().GetHost() != "example.test" || c.GetDestAddr().GetPort() != 443 {
				t.Fatal("constructed state", c.GetDestAddr())
			}
			got := make([]byte, 7)
			if _, err = io.ReadFull(c, got); err != nil || string(got) != "payload" {
				t.Fatal(string(got), err)
			}
			if n, e := c.Write([]byte("reply")); n != 5 || e != nil {
				t.Fatal(n, e)
			}
			reply := make([]byte, 5)
			io.ReadFull(client, reply)
			if string(reply) != "reply" {
				t.Fatal(string(reply))
			}
			up, down := user.GetAndClearTraffic()
			if up != int64(len(addressFixture)+7) || down != 5 {
				t.Fatal("raw traffic accounting", up, down)
			}
		})
	}
}
func TestNativeSSRConstructorLookupErrors(t *testing.T) {
	for _, mode := range []E.LaTpsgYQ7{0, 1, 3} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			client, server := tcpPair(t)
			info, _ := wrapperInfo(t, mode)
			info.N3pMXuwX85c = new(M.GTFAWF[int64, *U.LJJkaD])
			info.RLXBf_luGj = new(M.GTFAWF[string, *U.LJJkaD])
			client.Write(addressFixture)
			c, err := RfxaTYLqLUJ(server, info)
			want := map[E.LaTpsgYQ7]string{0: "uid <42> not exist", 1: "host <user.example> not exist", 3: "not support single port type: 3"}[mode]
			if c != nil || err == nil || err.Error() != want {
				t.Fatal(c, err)
			}
		})
	}
}
func packetPair(t *testing.T) (net.PacketConn, net.PacketConn) {
	t.Helper()
	a, e := net.ListenPacket("udp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	b, e := net.ListenPacket("udp4", "127.0.0.1:0")
	if e != nil {
		a.Close()
		t.Fatal(e)
	}
	a.SetDeadline(time.Now().Add(3 * time.Second))
	b.SetDeadline(time.Now().Add(3 * time.Second))
	t.Cleanup(func() { a.Close(); b.Close() })
	return a, b
}
func TestNativeSSRPacketWrapperUDP(t *testing.T) {
	for _, mode := range []E.LaTpsgYQ7{0, 1, 2} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			client, server := packetPair(t)
			info, user := wrapperInfo(t, mode)
			c, e := INOdx8(server, info)
			if e != nil {
				t.Fatal(e)
			}
			defer c.Free()
			wire := append(append([]byte{}, addressFixture...), []byte("payload")...)
			client.WriteTo(wire, server.LocalAddr())
			buf := make([]byte, 2048)
			n, u, src, dest, e := c.ReadFromProxy(buf)
			if n != 7 || string(buf[:n]) != "payload" || src.String() != client.LocalAddr().String() || dest.GetHost() != "example.test" || e != nil {
				t.Fatal(n, u, src, dest, e)
			}
			if (mode == 2 && u != user) || (mode != 2 && u != nil) {
				t.Fatal("mode-specific user", u)
			}
			up, _ := user.GetAndClearTraffic()
			wantUp := int64(0)
			if mode == 2 {
				wantUp = int64(len(wire))
			}
			if up != wantUp {
				t.Fatal(up, wantUp)
			}
			target := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 2), Port: 444}
			n, e = c.WriteToProxy([]byte("reply"), user, src, target)
			if n != 12 || e != nil {
				t.Fatal("count includes SOCKS address", n, e)
			}
			n, _, e = client.ReadFrom(buf)
			if e != nil || !bytes.Equal(buf[:n], []byte{1, 127, 0, 0, 2, 1, 188, 'r', 'e', 'p', 'l', 'y'}) {
				t.Fatal(n, e, buf[:n])
			}
			_, down := user.GetAndClearTraffic()
			if down != 12 {
				t.Fatal(down)
			}
			// The original drops SOCKS address parsing errors when packet reading itself
			// succeeded; retaining that quirk prevents an invented error path.
			client.WriteTo([]byte{255}, server.LocalAddr())
			n, u, src, dest, e = c.ReadFromProxy(buf)
			if n != 0 || u != nil || dest != nil || e != nil || src == nil {
				t.Fatal(n, u, src, dest, e)
			}
		})
	}
}

// The following oracle checks the precise error/result ordering using a test
// packet fixture. Production still calls the actual restored protocol.
type resultPacket struct {
	R.Vzuv_s9ZS5
	n    int
	err  error
	data []byte
}

func (p *resultPacket) ReadFromUID(b []byte) (int, int64, net.Addr, error) {
	copy(b, p.data)
	return p.n, 0, &net.UDPAddr{}, p.err
}
func TestNativeSSRPacketErrorRetention(t *testing.T) {
	info, _ := wrapperInfo(t, 0)
	sentinel := errors.New("underlying")
	for _, tt := range []struct {
		data []byte
		n    int
		want error
	}{{[]byte{255}, 1, sentinel}, {nil, 0, sentinel}, {addressFixture, len(addressFixture), nil}} {
		raw := &resultPacket{n: tt.n, err: sentinel, data: tt.data}
		c := &AXWGpkYq{j6haDfFi: raw, adiN53: info}
		n, _, _, _, err := c.ReadFromProxy(make([]byte, 128))
		if n != 0 || err != tt.want {
			t.Fatal(n, err, tt.want)
		}
	}
}
