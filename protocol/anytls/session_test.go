package anytls

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	C "recovered.local/soga/util/rtx"
	H "recovered.local/soga/transport/hysteriapad"
	S "recovered.local/soga/protocol/anytlssettings"
	"sync/atomic"
	"testing"
	"time"
)

func wireFrame(cmd byte, id uint32, p []byte) []byte {
	b := make([]byte, 7+len(p))
	b[0] = cmd
	binary.BigEndian.PutUint32(b[1:], id)
	binary.BigEndian.PutUint16(b[5:], uint16(len(p)))
	copy(b[7:], p)
	return b
}
func sendFrame(t *testing.T, c net.Conn, cmd byte, id uint32, p []byte) {
	t.Helper()
	if _, e := c.Write(wireFrame(cmd, id, p)); e != nil {
		t.Fatal(e)
	}
}
func getFrame(t *testing.T, c net.Conn) (byte, uint32, []byte) {
	t.Helper()
	var h [7]byte
	if _, e := io.ReadFull(c, h[:]); e != nil {
		t.Fatal(e)
	}
	b := make([]byte, binary.BigEndian.Uint16(h[5:]))
	if _, e := io.ReadFull(c, b); e != nil {
		t.Fatal(e)
	}
	return h[0], binary.BigEndian.Uint32(h[1:]), b
}
func wantFrame(t *testing.T, c net.Conn, cmd byte, id uint32, p []byte) {
	t.Helper()
	a, b, q := getFrame(t, c)
	if a != cmd || b != id || !bytes.Equal(q, p) {
		t.Fatalf("frame=(%d,%d,%q), want (%d,%d,%q)", a, b, q, cmd, id, p)
	}
}

func serveFixture(t *testing.T, callback func(*CCXLCurYiT7)) (net.Conn, *BIkgBMz3M, <-chan struct{}) {
	t.Helper()
	ln, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	client, e := net.Dial("tcp", ln.Addr().String())
	if e != nil {
		ln.Close()
		t.Fatal(e)
	}
	conn, e := ln.Accept()
	ln.Close()
	if e != nil {
		client.Close()
		t.Fatal(e)
	}
	client.SetDeadline(time.Now().Add(5 * time.Second))
	var padding C.EahpvU_iseQ[*H.FZBbzkJPWbrk]
	padding.Store(H.R_GY_FGx([]byte("stop=1")))
	session := NewServerFromInline(conn, callback, &padding)
	done := make(chan struct{})
	go func() { session.Run(); close(done) }()
	t.Cleanup(func() {
		client.Close()
		session.Close()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("read loop leaked")
		}
	})
	return client, session, done
}

func TestServerActualTCPMultiplexSettingsFIN(t *testing.T) {
	for _, version := range []int{1, 2, 258} {
		t.Run(fmt.Sprint(version), func(t *testing.T) {
			var opened atomic.Int32
			c, s, _ := serveFixture(t, func(st *CCXLCurYiT7) { opened.Add(1); defer st.Close(); st.HandshakeSuccess(); io.Copy(st, st) })
			settings := []byte(fmt.Sprintf("v=%d\npadding-md5=%s", version, s.bmjW7a.Load().GtWiNfVs7j))
			sendFrame(t, c, 4, 0, settings)
			if version >= 2 {
				wantFrame(t, c, 10, 0, []byte("v=2"))
			}
			for id := uint32(1); id <= 3; id++ {
				sendFrame(t, c, 1, id, nil)
				if version >= 2 {
					wantFrame(t, c, 7, id, nil)
				}
				payload := bytes.Repeat([]byte{byte(id), 0, 255}, 2000)
				sendFrame(t, c, 2, id, payload)
				wantFrame(t, c, 2, id, payload)
			}
			// 重复 SYN 不重复创建流；后续 ping 回应确保已消费该头。
			sendFrame(t, c, 1, 1, nil)
			sendFrame(t, c, 8, 66, nil)
			wantFrame(t, c, 9, 66, nil)
			if opened.Load() != 3 {
				t.Fatalf("opened %d streams", opened.Load())
			}
			for id := uint32(1); id <= 3; id++ {
				sendFrame(t, c, 3, id, nil)
			}
			sendFrame(t, c, 8, 67, nil)
			wantFrame(t, c, 9, 67, nil)
			s.ukg_2xDnH9s.RLock()
			n := len(s.mW8PvJYxYup9)
			s.ukg_2xDnH9s.RUnlock()
			if n != 0 {
				t.Fatalf("remaining streams=%d", n)
			}
		})
	}
}

func TestSettingsRequiredAndPaddingNegotiation(t *testing.T) {
	t.Run("missing", func(t *testing.T) {
		c, _, done := serveFixture(t, nil)
		sendFrame(t, c, 4, 0, nil)
		sendFrame(t, c, 1, 9, nil)
		wantFrame(t, c, 5, 0, []byte("client did not send its settings"))
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("missing settings did not close")
		}
	})
	t.Run("different-hash", func(t *testing.T) {
		c, s, _ := serveFixture(t, nil)
		sendFrame(t, c, 4, 0, []byte("v=2\npadding-md5=invalid"))
		wantFrame(t, c, 6, 0, s.bmjW7a.Load().KruzWnr9zxA)
		wantFrame(t, c, 10, 0, []byte("v=2"))
		sendFrame(t, c, 1, 4, nil)
		wantFrame(t, c, 3, 4, nil)
	})
	t.Run("negative-version-still-settings", func(t *testing.T) {
		c, s, _ := serveFixture(t, nil)
		sendFrame(t, c, 4, 0, []byte("v=-1\npadding-md5="+s.bmjW7a.Load().GtWiNfVs7j))
		sendFrame(t, c, 1, 9, nil)
		wantFrame(t, c, 3, 9, nil)
	})
}

func TestOriginalIgnoredLengthsAndUnknownCommand(t *testing.T) {
	c, _, _ := serveFixture(t, nil)
	// 原 0x18a2ec0 只消耗未知命令头；如果误加通用 payload 跳过，以下 ping 将不能解析。
	b := wireFrame(200, 1, nil)
	binary.BigEndian.PutUint16(b[5:], 50)
	b = append(b, wireFrame(8, 33, nil)...)
	if _, e := c.Write(b); e != nil {
		t.Fatal(e)
	}
	wantFrame(t, c, 9, 33, nil)
	sendFrame(t, c, 0, 0, []byte("padding bytes"))
	sendFrame(t, c, 2, 999, []byte("unknown stream"))
	sendFrame(t, c, 5, 0, nil)
	sendFrame(t, c, 8, 34, nil)
	wantFrame(t, c, 9, 34, nil)
}

func TestStreamDeadlineErrorsAndHandshakeOnce(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	b.SetDeadline(time.Now().Add(3 * time.Second))
	s := NewServerFromInline(a, nil, &H.DefaultAt5098340)
	s.gAJaebL = 2
	st := newStreamFromInline(s, 9)
	s.mW8PvJYxYup9[9] = st
	st.SetWriteDeadline(time.Now().Add(-time.Second))
	if n, e := st.Write([]byte{1}); n != 0 || !errors.Is(e, os.ErrDeadlineExceeded) {
		t.Fatalf("write timeout %d %v", n, e)
	}
	st.SetDeadline(time.Now().Add(-time.Second))
	if _, e := st.Read(make([]byte, 1)); !errors.Is(e, os.ErrDeadlineExceeded) {
		t.Fatal(e)
	}
	st.SetDeadline(time.Time{})
	if st.LocalAddr().String() != a.LocalAddr().String() || st.RemoteAddr().String() != a.RemoteAddr().String() {
		t.Fatal("stream addresses changed")
	}
	done := make(chan error, 1)
	go func() { done <- st.HandshakeFailure(errors.New("denied")) }()
	wantFrame(t, b, 7, 9, []byte("denied"))
	if e := <-done; e != nil {
		t.Fatal(e)
	}
	if e := st.HandshakeSuccess(); e != nil {
		t.Fatal(e)
	} // 无第二个 ACK。
	calls := 0
	st.kCaaEFymiBDr = func() { calls++ }
	st.mublR9bouj1Z()
	st.mublR9bouj1Z()
	if calls != 1 {
		t.Fatal(calls)
	}
	if _, e := st.Write(nil); !errors.Is(e, net.ErrClosed) {
		t.Fatalf("remote close: %v", e)
	}
	if e := st.Close(); !errors.Is(e, net.ErrClosed) {
		t.Fatal(e)
	}
	s.Close()
	if e := s.Close(); e != io.ErrClosedPipe {
		t.Fatal(e)
	}
	if _, e := s.OpenStream(); e != io.ErrClosedPipe {
		t.Fatal(e)
	}
}

func TestClientRunBufferedFirstFlightAndRemoteRejection(t *testing.T) {
	a, b := net.Pipe()
	b.SetDeadline(time.Now().Add(4 * time.Second))
	defer b.Close()
	s := NewServerFromInline(a, nil, &H.DefaultAt5098340)
	s.pNuGhHxIQdo = true
	defer s.Close()
	s.Run()
	st, e := s.OpenStream()
	if e != nil {
		t.Fatal(e)
	}
	if len(s.vV_K36) == 0 || s.d0Ze8qja {
		t.Fatal("Run/SYN buffering mismatch")
	}
	sent := make(chan error, 1)
	go func() { _, e := st.Write([]byte("hello")); sent <- e }()
	cmd, id, p := getFrame(t, b)
	if cmd != 4 || id != 0 {
		t.Fatalf("settings %d %d", cmd, id)
	}
	settings := S.DsbUyzhj9(p)
	if settings["client"] != "anytls/0.0.11" || settings["v"] != "2" || settings["padding-md5"] != H.DefaultAt5098340.Load().GtWiNfVs7j {
		t.Fatal(settings)
	}
	wantFrame(t, b, 1, 1, nil)
	wantFrame(t, b, 2, 1, []byte("hello"))
	if e := <-sent; e != nil {
		t.Fatal(e)
	}
	sendFrame(t, b, 10, 0, []byte("v=2"))
	sendFrame(t, b, 7, 1, []byte("refused"))
	wantFrame(t, b, 3, 1, nil)
	// FIN writer完成后会从map移除；用ping作为readloop继续执行的同步点。
	sendFrame(t, b, 8, 8, nil)
	wantFrame(t, b, 9, 8, nil)
	if _, e := st.Write(nil); e == nil || e.Error() != "remote: refused" {
		t.Fatalf("remote rejection %v", e)
	}
}

func TestClientSecondStreamACKCancelsOriginalTimeout(t *testing.T) {
	a, b := net.Pipe()
	defer b.Close()
	b.SetDeadline(time.Now().Add(4 * time.Second))
	s := NewServerFromInline(a, nil, &H.DefaultAt5098340)
	s.pNuGhHxIQdo = true
	defer s.Close()
	s.Run()
	first, e := s.OpenStream()
	if e != nil {
		t.Fatal(e)
	}
	sent := make(chan error, 1)
	go func() { _, e := first.Write(nil); sent <- e }()
	getFrame(t, b)
	wantFrame(t, b, 1, 1, nil)
	wantFrame(t, b, 2, 1, nil)
	if e := <-sent; e != nil {
		t.Fatal(e)
	}
	sendFrame(t, b, 10, 0, []byte("v=2"))
	sendFrame(t, b, 8, 71, nil)
	wantFrame(t, b, 9, 71, nil)
	opened := make(chan *CCXLCurYiT7, 1)
	go func() { st, e := s.OpenStream(); sent <- e; opened <- st }()
	wantFrame(t, b, 1, 2, nil)
	if e := <-sent; e != nil {
		t.Fatal(e)
	}
	second := <-opened
	s.wAiQ9IC2UZ.Lock()
	armed := s.c5NUXMoBI != nil
	s.wAiQ9IC2UZ.Unlock()
	if !armed {
		t.Fatal("second stream timer not armed")
	}
	sendFrame(t, b, 7, 2, nil)
	sendFrame(t, b, 8, 72, nil)
	wantFrame(t, b, 9, 72, nil)
	s.wAiQ9IC2UZ.Lock()
	cancelled := s.c5NUXMoBI == nil
	s.wAiQ9IC2UZ.Unlock()
	if !cancelled {
		t.Fatal("ACK did not cancel timer")
	}
	if s.IsClosed() || second.yKgxtGJV != nil {
		t.Fatal("successful ACK closed stream")
	}
}

func TestClientPaddingUpdateAndVersionCast(t *testing.T) {
	old := H.DefaultAt5098340.Load()
	defer H.DefaultAt5098340.Store(old)
	oldDebug := debugPaddingAt50C6C92
	defer func() { debugPaddingAt50C6C92 = oldDebug }()
	for _, debug := range []bool{false, true} {
		t.Run(fmt.Sprint(debug), func(t *testing.T) {
			H.DefaultAt5098340.Store(old)
			debugPaddingAt50C6C92 = debug
			a, b := net.Pipe()
			defer b.Close()
			b.SetDeadline(time.Now().Add(4 * time.Second))
			var separate C.EahpvU_iseQ[*H.FZBbzkJPWbrk]
			separate.Store(old)
			s := NewServerFromInline(a, nil, &separate)
			s.pNuGhHxIQdo = true
			done := make(chan error, 1)
			go func() { done <- s.wbBfTO632dea() }()
			defer func() { s.Close(); <-done }()
			update := []byte("stop=4\n1=15-15")
			sendFrame(t, b, 6, 0, update)
			sendFrame(t, b, 10, 0, []byte("v=-1"))
			sendFrame(t, b, 8, 20, nil)
			wantFrame(t, b, 9, 20, nil)
			if s.gAJaebL != 255 {
				t.Fatal("client version did not cast signed integer")
			}
			if separate.Load() != old {
				t.Fatal("original update must target shared default, not session pointer")
			}
			if debug {
				if H.DefaultAt5098340.Load() != old {
					t.Fatal("debug unexpectedly updated scheme")
				}
			} else {
				if !bytes.Equal(H.DefaultAt5098340.Load().KruzWnr9zxA, update) {
					t.Fatal("global scheme not updated")
				}
			}
		})
	}
}

func TestReadLoopTruncatedFramesClose(t *testing.T) {
	for _, cmd := range []byte{0, 2, 4, 5, 6, 7, 10} {
		t.Run(fmt.Sprint(cmd), func(t *testing.T) {
			a, b := net.Pipe()
			s := NewServerFromInline(a, nil, &H.DefaultAt5098340)
			done := make(chan error, 1)
			go func() { done <- s.wbBfTO632dea() }()
			b.Write(wireFrame(cmd, 1, []byte{1, 2})[:8])
			b.Close()
			select {
			case err := <-done:
				if !errors.Is(err, io.ErrUnexpectedEOF) {
					t.Fatal(err)
				}
				if !s.IsClosed() {
					t.Fatal("truncation left session open")
				}
			case <-time.After(time.Second):
				t.Fatal("truncated read hung")
			}
		})
	}
	a, b := net.Pipe()
	s := NewServerFromInline(a, nil, &H.DefaultAt5098340)
	done := make(chan error, 1)
	go func() { done <- s.wbBfTO632dea() }()
	b.Write([]byte{2, 0})
	b.Close()
	if e := <-done; !errors.Is(e, io.ErrUnexpectedEOF) {
		t.Fatal(e)
	}
}
