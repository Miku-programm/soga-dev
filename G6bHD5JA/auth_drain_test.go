package G6bHD5JA

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
	"net"
	"os"
	"testing"
	"time"
)

type authDrainProbe struct {
	deadline time.Time
	readSize int
	reads    int
	closed   bool
}

func (c *authDrainProbe) Read(p []byte) (int, error) {
	c.readSize = len(p)
	c.reads++
	return 0, io.EOF
}
func (c *authDrainProbe) Write(p []byte) (int, error)       { panic("unexpected write") }
func (c *authDrainProbe) Close() error                      { c.closed = true; return nil }
func (c *authDrainProbe) LocalAddr() net.Addr               { return &net.TCPAddr{} }
func (c *authDrainProbe) RemoteAddr() net.Addr              { return &net.TCPAddr{} }
func (c *authDrainProbe) SetDeadline(time.Time) error       { panic("wrong deadline") }
func (c *authDrainProbe) SetWriteDeadline(time.Time) error  { panic("wrong deadline") }
func (c *authDrainProbe) SetReadDeadline(t time.Time) error { c.deadline = t; return nil }
func TestNativeAuthenticationFailureDrain(t *testing.T) {
	name, _ := os.Hostname()
	sum := sha256.Sum256([]byte(name))
	sum[0] &= 0x7f
	expected := binary.BigEndian.Uint32(sum[:4])
	if gSnxlJwdTs() != expected {
		t.Fatal("host fingerprint")
	}
	for _, bounds := range [][2]int{{4, 4}, {5, 2}, {-3, -4}} {
		if i3zxS3(bounds[0], bounds[1]) != bounds[0] {
			t.Fatal("degenerate range")
		}
	}
	for i := 0; i < 100; i++ {
		x := i3zxS3(3, 11)
		if x < 3 || x > 11 {
			t.Fatal(x)
		}
	}
	before := time.Now()
	c := new(authDrainProbe)
	Ad2VIxY(c)
	after := time.Now()
	offset := time.Duration(expected%20000) * time.Millisecond
	if c.deadline.Before(before.Add(2*time.Second+offset)) || c.deadline.After(after.Add(10*time.Second+offset)) {
		t.Fatal("deadline range", c.deadline)
	}
	if c.closed {
		t.Fatal("original helper must not close")
	}
	if c.reads > 1 || (c.reads == 1 && c.readSize != 1<<((expected&3)+12)) {
		t.Fatal("read buffer", c.reads, c.readSize)
	}
}
