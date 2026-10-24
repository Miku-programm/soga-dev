package GLaePRKYM

import (
	"bytes"
	"errors"
	"net"
	"reflect"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func TestNativeSalamanderUDP(t *testing.T) {
	a, e := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if e != nil {
		t.Fatal(e)
	}
	defer a.Close()
	b, e := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if e != nil {
		t.Fatal(e)
	}
	defer b.Close()
	x, _ := Ea9POnRF([]byte("native-password"))
	y, _ := Ea9POnRF([]byte("native-password"))
	left, right := WrapFromHysteriaInline(a, x), WrapFromHysteriaInline(b, y)
	left.SetDeadline(time.Now().Add(time.Second))
	right.SetDeadline(time.Now().Add(time.Second))
	if _, ok := left.(syscall.Conn); !ok {
		t.Fatal("UDP socket capability lost")
	}
	if _, e := a.WriteToUDP([]byte{1, 2}, b.LocalAddr().(*net.UDPAddr)); e != nil {
		t.Fatal(e)
	}
	payload := bytes.Repeat([]byte("native\x00"), 290)
	if n, e := left.WriteTo(payload, right.LocalAddr()); e != nil || n != len(payload) {
		t.Fatal(n, e)
	}
	buf := make([]byte, 2048)
	n, addr, e := right.ReadFrom(buf)
	if e != nil || !bytes.Equal(buf[:n], payload) || addr.String() != left.LocalAddr().String() {
		t.Fatal(n, addr, e)
	}
	if n, e = right.WriteTo(buf[:n], addr); e != nil || n != len(payload) {
		t.Fatal(n, e)
	}
	n, _, e = left.ReadFrom(buf)
	if e != nil || !bytes.Equal(buf[:n], payload) {
		t.Fatal(n, e)
	}
	for i := 0; i < 4; i++ {
		_, e := Ea9POnRF(make([]byte, i))
		if e != shortPasswordAt5098D90 {
			t.Fatal("short-key error identity")
		}
	}
	key := make([]byte, 4, 20)
	o, _ := Ea9POnRF(key)
	key[0] = 17
	if o.EDkKnfGR8w[0] != 17 {
		t.Fatal("key alias lost")
	}
}

type packetResult struct {
	b   []byte
	err error
}
type packetFake struct {
	net.PacketConn
	read     []packetResult
	written  []byte
	writeN   int
	writeErr error
}

func (f *packetFake) ReadFrom(b []byte) (int, net.Addr, error) {
	r := f.read[0]
	f.read = f.read[1:]
	return copy(b, r.b), &net.UDPAddr{Port: 5}, r.err
}
func (f *packetFake) WriteTo(b []byte, _ net.Addr) (int, error) {
	f.written = append([]byte{}, b...)
	return f.writeN, f.writeErr
}

func TestOriginalPacketErrorAndBufferBoundaries(t *testing.T) {
	o, _ := Ea9POnRF([]byte("test"))
	f := &packetFake{}
	v := WrapFromHysteriaInline(f, o)
	if _, ok := v.(syscall.Conn); ok {
		t.Fatal("invented socket capability")
	}
	// 原写入只按 error 判定成功，不检查底层短写；过大报文混淆返回零，仍提交空报文。
	for _, size := range []int{0, 1, 2040, 2041, 4096} {
		p := make([]byte, size)
		n, e := v.WriteTo(p, nil)
		if n != size || e != nil {
			t.Fatal(n, e)
		}
		want := size + 8
		if size > 2040 {
			want = 0
		}
		if len(f.written) != want {
			t.Fatal("wire size", len(f.written), want)
		}
	}
	sentinel := errors.New("transport failure")
	f.writeErr = sentinel
	if n, e := v.WriteTo([]byte{1}, nil); n != 0 || e != sentinel {
		t.Fatal(n, e)
	}
	wire := make([]byte, 9)
	o.Obfuscate([]byte{7}, wire)
	f.read = []packetResult{{[]byte{1}, nil}, {wire, sentinel}}
	buf := make([]byte, 1)
	n, _, e := v.ReadFrom(buf)
	if n != 1 || buf[0] != 7 || e != sentinel {
		t.Fatal(n, e)
	}
	f.read = []packetResult{{wire, sentinel}}
	n, _, e = v.ReadFrom(nil)
	if n != 0 || e != sentinel {
		t.Fatal(n, e)
	}
	var a DOqsrH8iC
	var b v2leTMr
	var c kIKvVu0qKkaB
	if !reflect.DeepEqual([]uintptr{unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Offsetof(b.g4ZgYMaCfW7), unsafe.Offsetof(b.qXrWg_XLkm)}, []uintptr{40, 96, 16, 56, 88}) {
		t.Fatal("RTTI layout")
	}
}
