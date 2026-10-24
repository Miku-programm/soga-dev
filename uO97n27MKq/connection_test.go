package uO97n27MKq

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	J "recovered.local/soga/J1P4KNvfI_J"
	"testing"
	"time"
	"unsafe"
)

// 客户端夹具直接按协议编码，未调用恢复后的地址/帧编码函数。
func fixtureAddress(ip net.IP, port uint16) []byte {
	if v4 := ip.To4(); v4 != nil {
		return append(append([]byte{1}, v4...), byte(port>>8), byte(port))
	}
	return append(append([]byte{4}, ip.To16()...), byte(port>>8), byte(port))
}
func fixtureHeader(command byte, address []byte) []byte {
	p := append([]byte{13, 10, command}, address...)
	return append(p, 13, 10)
}
func fixturePacket(address, payload []byte) []byte {
	p := append([]byte(nil), address...)
	p = append(p, byte(len(payload)>>8), byte(len(payload)), 13, 10)
	return append(p, payload...)
}
func localPair(t *testing.T) (net.Conn, net.Conn) {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	client, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	server, err := l.Accept()
	if err != nil {
		client.Close()
		t.Fatal(err)
	}
	t.Cleanup(func() { client.Close(); server.Close() })
	client.SetDeadline(time.Now().Add(5 * time.Second))
	server.SetDeadline(time.Now().Add(5 * time.Second))
	return client, server
}
func TestNativeTrojanTCPDuplex(t *testing.T) {
	var layout DaFYwI
	if unsafe.Sizeof(layout) != 64 || unsafe.Offsetof(layout.uszzyXZ5k) != 24 || unsafe.Offsetof(layout.dXeEnBG) != 40 || unsafe.Offsetof(layout.nZSKiASba) != 56 {
		t.Fatal("RTTI")
	}
	for _, ip := range []net.IP{net.IPv4(192, 0, 2, 1), net.ParseIP("2001:db8::1")} {
		t.Run(ip.String(), func(t *testing.T) {
			client, server := localPair(t)
			done := make(chan error, 1)
			upload := bytes.Repeat([]byte("native-trojan-upload"), 2100)
			download := bytes.Repeat([]byte("native-trojan-download"), 1900)
			go func() {
				c, err := Lzwh7y6zTaq(&J.Mdn09F6aMrpa{DrpXao: server}, nil)
				if err != nil {
					done <- err
					return
				}
				if c.NetworkType() != 0 || c.IsMux() || c.GetDestAddr().GetPort() != 443 || !c.GetDestAddr().GetFirstIp().Equal(ip) {
					done <- errors.New("header state")
					return
				}
				received := make([]byte, len(upload))
				_, err = io.ReadFull(c, received)
				if err == nil && !bytes.Equal(received, upload) {
					err = errors.New("upload data")
				}
				if err == nil {
					_, err = c.Write(download)
				}
				done <- err
			}()
			if _, err := client.Write(append(fixtureHeader(1, fixtureAddress(ip, 443)), upload...)); err != nil {
				t.Fatal(err)
			}
			received := make([]byte, len(download))
			if _, err := io.ReadFull(client, received); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(received, download) {
				t.Fatal("download data")
			}
			if err := <-done; err != nil {
				t.Fatal(err)
			}
		})
	}
}
func TestNativeTrojanUDPFramingAndTruncation(t *testing.T) {
	for _, ip := range []net.IP{net.IPv4(192, 0, 2, 2), net.ParseIP("2001:db8::2")} {
		t.Run(ip.String(), func(t *testing.T) {
			client, server := localPair(t)
			done := make(chan error, 1)
			addr := fixtureAddress(ip, 5353)
			go func() {
				c, err := Lzwh7y6zTaq(&J.Mdn09F6aMrpa{DrpXao: server}, nil)
				if err != nil {
					done <- err
					return
				}
				if c.NetworkType() != 1 {
					done <- errors.New("UDP state")
					return
				}
				small := make([]byte, 3)
				n, a, err := c.ReadFromProxy(small)
				if err != nil || n != 3 || string(small) != "abc" {
					done <- errors.New("truncation")
					return
				}
				n, a, err = c.ReadFromProxy(small)
				if err != nil || n != 2 || string(small[:n]) != "XY" {
					done <- errors.New("frame boundary after discard")
					return
				}
				if a.String() != (&net.UDPAddr{IP: ip, Port: 5353}).String() {
					done <- errors.New("datagram destination")
					return
				}
				_, err = c.WriteToProxy([]byte("reply"), a)
				done <- err
			}()
			wire := fixtureHeader(3, addr)
			wire = append(wire, fixturePacket(addr, []byte("abcdefghi"))...)
			wire = append(wire, fixturePacket(addr, []byte("XY"))...)
			if _, err := client.Write(wire); err != nil {
				t.Fatal(err)
			}
			want := fixturePacket(addr, []byte("reply"))
			got := make([]byte, len(want))
			if _, err := io.ReadFull(client, got); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("wire reply %x", got)
			}
			if err := <-done; err != nil {
				t.Fatal(err)
			}
		})
	}
}

type byteConnection struct {
	net.Conn
	r        *bytes.Reader
	written  []byte
	writeErr error
	closed   bool
}

func (c *byteConnection) Read(p []byte) (int, error) { return c.r.Read(p) }
func (c *byteConnection) Write(p []byte) (int, error) {
	c.written = append(c.written, p...)
	return 1, c.writeErr
}
func (c *byteConnection) Close() error { c.closed = true; return nil }
func TestNativeTrojanHeaderFailures(t *testing.T) {
	header := fixtureHeader(1, fixtureAddress(net.IPv4(192, 0, 2, 1), 443))
	for n := 0; n < len(header); n++ {
		raw := &byteConnection{r: bytes.NewReader(header[:n])}
		if c, err := Lzwh7y6zTaq(&J.Mdn09F6aMrpa{DrpXao: raw}, nil); c != nil || err == nil {
			t.Fatalf("truncated header %d", n)
		}
		if raw.closed {
			t.Fatal("original constructor Free must not Close")
		}
	}
	for _, index := range []int{0, 1, len(header) - 2, len(header) - 1} {
		bad := bytes.Clone(header)
		bad[index] ^= 1
		_, err := Lzwh7y6zTaq(&J.Mdn09F6aMrpa{DrpXao: &byteConnection{r: bytes.NewReader(bad)}}, nil)
		if err == nil || !bytes.Contains([]byte(err.Error()), []byte("not crlf:")) {
			t.Fatal("bad delimiter", index, err)
		}
	}
	for _, command := range []byte{0, 2, 4, 255} {
		bad := bytes.Clone(header)
		bad[2] = command
		_, err := Lzwh7y6zTaq(&J.Mdn09F6aMrpa{DrpXao: &byteConnection{r: bytes.NewReader(bad)}}, nil)
		if err == nil || !bytes.Contains([]byte(err.Error()), []byte("unknown trojan cmd:")) {
			t.Fatal("bad command", command, err)
		}
	}
}
func TestNativeTrojanDatagramErrorSemantics(t *testing.T) {
	addr := fixtureAddress(net.IPv4(192, 0, 2, 1), 53)
	// 正文不足，原实现不向调用者返回部分长度或地址。
	truncated := fixturePacket(addr, []byte("abc"))
	truncated = truncated[:len(truncated)-1]
	c := &DaFYwI{DrpXao: &byteConnection{r: bytes.NewReader(truncated)}}
	if n, a, err := c.ReadFromProxy(make([]byte, 3)); n != 0 || a != nil || err != io.ErrUnexpectedEOF {
		t.Fatal("partial payload result", n, a, err)
	}
	// 已填满短缓冲后，丢弃尾部失败则保留n和地址。
	c = &DaFYwI{DrpXao: &byteConnection{r: bytes.NewReader(truncated)}}
	if n, a, err := c.ReadFromProxy(make([]byte, 1)); n != 1 || a == nil || err == nil {
		t.Fatal("discard failure result", n, a, err)
	}
	raw := &byteConnection{r: bytes.NewReader(nil)}
	c = &DaFYwI{DrpXao: raw}
	payload := make([]byte, 65536)
	n, err := c.WriteToProxy(payload, &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 53})
	if err != nil || n != len(payload) || binary.BigEndian.Uint16(raw.written[7:9]) != 0 || len(raw.written) != len(payload)+11 {
		t.Fatal("uint16 length wrap and ignored short write")
	}
	raw.writeErr = io.ErrClosedPipe
	if n, err = c.WriteToProxy([]byte("x"), &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 53}); n != 0 || err != io.ErrClosedPipe {
		t.Fatal("write failure")
	}
}
