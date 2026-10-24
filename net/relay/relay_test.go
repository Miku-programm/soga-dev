package relay

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	D "recovered.local/soga/buf/pool"
	H "recovered.local/soga/util/errutil"
	U "recovered.local/soga/buf/bufiface"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type bufferedRelayFixture struct {
	net.Conn
	buf *D.X_aXX5
}

func (c *bufferedRelayFixture) Free() {
	if c.buf != nil {
		c.buf.Free()
	}
}
func (c *bufferedRelayFixture) ReadBuffer() (U.Mf3phSOVqgf, error) {
	if c.buf == nil {
		c.buf = D.S4OQXWIZ(4096)
	}
	c.buf.Reset()
	_, err := c.buf.ReadMaxFrom(c.Conn, 4096)
	return c.buf, err
}
func (c *bufferedRelayFixture) WriteBuffer(b U.Mf3phSOVqgf) (int, error) { return c.Write(b.Bytes()) }

func TestRelayBothDirectionsOverRealPipes(t *testing.T) {
	for _, buffered := range []bool{false, true} {
		a, left := net.Pipe()
		right, b := net.Pipe()
		for _, c := range []net.Conn{a, left, right, b} {
			c.SetDeadline(time.Now().Add(3 * time.Second))
			defer c.Close()
		}
		done := make(chan struct{})
		if buffered {
			c := &bufferedRelayFixture{Conn: left}
			defer c.Free()
			go func() { BKn06KsCTP(context.Background(), c, right); close(done) }()
		} else {
			go func() {
				FvtgLyQrWX(AsVABr{C8FYQVoBTa: context.Background(), GdPSmtGs_1: left, HEfD13ahR__b: right})
				close(done)
			}()
		}
		payloadA := bytes.Repeat([]byte("upstream"), 9000)
		payloadB := bytes.Repeat([]byte("downstream"), 8500)
		errs := make(chan error, 4)
		go func() { _, e := a.Write(payloadA); errs <- e }()
		go func() { _, e := b.Write(payloadB); errs <- e }()
		go func() {
			v := make([]byte, len(payloadA))
			_, e := io.ReadFull(b, v)
			if e == nil && !bytes.Equal(v, payloadA) {
				e = errors.New("upstream mismatch")
			}
			errs <- e
		}()
		go func() {
			v := make([]byte, len(payloadB))
			_, e := io.ReadFull(a, v)
			if e == nil && !bytes.Equal(v, payloadB) {
				e = errors.New("downstream mismatch")
			}
			errs <- e
		}()
		for i := 0; i < 4; i++ {
			if err := <-errs; err != nil {
				t.Fatal(buffered, err)
			}
		}
		a.Close()
		b.Close()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("relay failed to close both directions")
		}
	}
}

type relayErrorFixture struct {
	net.Conn
	readErr, writeErr error
	read              bool
	written           []byte
	deadlines         int
	short             bool
	panicRead         bool
}

func (c *relayErrorFixture) Read(p []byte) (int, error) {
	if c.panicRead {
		panic("read fixture")
	}
	if c.read {
		return 0, io.EOF
	}
	c.read = true
	return copy(p, "payload"), c.readErr
}
func (c *relayErrorFixture) Write(p []byte) (int, error) {
	c.written = append(c.written, p...)
	if c.short {
		return 1, nil
	}
	return len(p), c.writeErr
}
func (c *relayErrorFixture) SetWriteDeadline(time.Time) error {
	c.deadlines++
	return errors.New("ignored deadline failure")
}

func TestPipelinePreservesReadWriteErrorAndShortWriteBehavior(t *testing.T) {
	readErr, writeErr := errors.New("read failure"), errors.New("write failure")
	src, dst := &relayErrorFixture{readErr: readErr}, &relayErrorFixture{writeErr: writeErr}
	idle := QCkuttV(0, func() {})
	err := JWBWWZS(context.Background(), src, dst, nil, idle)
	if err == nil || err.Error() != "read failure\nwrite failure" || errors.Is(err, readErr) {
		t.Fatalf("original aggregate error differs: %v", err)
	}
	if string(dst.written) != "payload" || dst.deadlines != 1 || !idle.m_XMay1FFo7.Load() {
		t.Fatal("write path state differs")
	}
	src, dst = &relayErrorFixture{readErr: io.EOF}, &relayErrorFixture{short: true}
	if err = JWBWWZS(context.Background(), src, dst, make([]byte, 32), idle); err != io.EOF {
		t.Fatalf("short write synthesized an error: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err = JWBWWZS(ctx, src, dst, nil, idle); err != H.ErrorAt5097800 {
		t.Fatal("context error identity changed")
	}
}

func TestPipelineOriginalRecoveryClosureDoesNotSwallowPanic(t *testing.T) {
	defer func() {
		if recover() != "read fixture" {
			t.Fatal("panic propagation differs")
		}
	}()
	JWBWWZS(context.Background(), &relayErrorFixture{panicRead: true}, &relayErrorFixture{}, nil, QCkuttV(0, func() {}))
}

func TestAdaptiveRelayBuffersAndLayouts(t *testing.T) {
	if unsafe.Sizeof(AsVABr{}) != 96 || unsafe.Sizeof(JJEIU0hAF{}) != 40 {
		t.Fatal("original layouts differ")
	}
	s := eN9twxaAav{size: 4096}
	for _, pair := range [][2]int{{4096, 4096}, {4096, 8192}, {8192, 8192}, {8192, 16384}, {16384, 16384}} {
		if got := s.aIjQSD(pair[0]); got != pair[1] {
			t.Fatal(pair, got)
		}
	}
	for i := 0; i < 15; i++ {
		if s.aIjQSD(1) != 16384 {
			t.Fatal("shrunk early")
		}
	}
	if s.aIjQSD(1) != 8192 {
		t.Fatal("did not shrink after 16 small reads")
	}
	s.aIjQSD(1)
	s.aIjQSD(4096)
	if s.small != 0 || s.full != 0 {
		t.Fatal("half-full read must reset counters")
	}
	state := AsVABr{C93uQCY9TN: D.PoolAt50961D8.Get(4096), Mw4pNzmFlqA: D.PoolAt50961D8.Get(8192)}
	state.FreeBuf()
	state.FreeBuf()
	if state.C93uQCY9TN != nil || state.Mw4pNzmFlqA != nil {
		t.Fatal("buffers retained")
	}
}

func TestIdleTimerActivityAndStop(t *testing.T) {
	calls := 0
	timer := QCkuttV(4*time.Hour, func() { calls++ })
	defer timer.Stop()
	timer.xkqY3siS.Stop()
	for i := 0; i < 3; i++ {
		timer.a3rxvmzRzoE()
	}
	if calls != 0 {
		t.Fatal("early expiry")
	}
	timer.m_XMay1FFo7.Store(true)
	timer.a3rxvmzRzoE()
	if timer.rapKwb3eF != 0 {
		t.Fatal("activity failed to reset")
	}
	for i := 0; i < 4; i++ {
		timer.a3rxvmzRzoE()
	}
	if calls != 1 {
		t.Fatal("missing expiry")
	}
	timer.Stop()
	timer.a3rxvmzRzoE()
	if calls != 1 {
		t.Fatal("stopped callback fired")
	}
	zero := QCkuttV(0, func() {})
	zero.Stop()
	if zero.cmzS3I8m.Load() {
		t.Fatal("zero timer changed stopped bit")
	}
	fired := make(chan struct{}, 1)
	live := QCkuttV(20*time.Millisecond, func() { fired <- struct{}{} })
	defer live.Stop()
	select {
	case <-fired:
	case <-time.After(time.Second):
		t.Fatal("live timer did not expire")
	}
}

func TestBufferedReadErrorDropsAccompanyingData(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	c := &bufferedRelayFixture{Conn: &relayErrorFixture{Conn: a, readErr: io.EOF}}
	defer c.Free()
	dst := &relayErrorFixture{}
	eywOJGkx(context.Background(), c, dst, QCkuttV(0, func() {}))
	if len(dst.written) != 0 {
		t.Fatal("original ReadBuffer error no longer discards data")
	}
	if c.buf == nil || !strings.Contains(c.buf.String(), "payload") {
		t.Fatal("buffer unexpectedly freed")
	}
}
