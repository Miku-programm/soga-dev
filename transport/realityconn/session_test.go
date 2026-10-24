package realityconn

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

func pair(t *testing.T, version int) (*DZOCfwWJsy, *DZOCfwWJsy) {
	t.Helper()
	a, b := net.Pipe()
	ca := defaultConfigRecovered()
	ca.RTNOROJC5 = version
	ca.DU0iTVAg = true
	cb := *ca
	left, err := CXFDjl_TYx8(a, ca)
	if err != nil {
		t.Fatal(err)
	}
	right := yfgupp(&cb, b, true)
	t.Cleanup(func() { left.Close(); right.Close() })
	return left, right
}
func TestOriginalNativeMuxDuplex(t *testing.T) {
	for _, version := range []int{1, 2} {
		t.Run(string(rune('0'+version)), func(t *testing.T) {
			server, client := pair(t, version)
			c, err := client.OpenStream()
			if err != nil {
				t.Fatal(err)
			}
			s, err := server.AcceptStream()
			if err != nil {
				t.Fatal(err)
			}
			if c.ID() != 3 || s.ID() != 3 || server.NumStreams() != 1 || client.NumStreams() != 1 {
				t.Fatal("stream lifecycle")
			}
			request := bytes.Repeat([]byte("request\x00"), 80000)
			response := bytes.Repeat([]byte("response\xff"), 60000)
			errCh := make(chan error, 2)
			go func() { _, e := c.Write(request); errCh <- e }()
			go func() { _, e := s.Write(response); errCh <- e }()
			c.SetReadDeadline(time.Now().Add(5 * time.Second))
			s.SetReadDeadline(time.Now().Add(5 * time.Second))
			readDone := make(chan error, 1)
			go func() {
				b := make([]byte, len(request))
				_, e := io.ReadFull(s, b)
				if e == nil && !bytes.Equal(b, request) {
					e = errors.New("request mismatch")
				}
				readDone <- e
			}()
			b := make([]byte, len(response))
			if _, e := io.ReadFull(c, b); e != nil || !bytes.Equal(b, response) {
				t.Fatalf("response: %v", e)
			}
			if e := <-readDone; e != nil {
				t.Fatal(e)
			}
			for range 2 {
				if e := <-errCh; e != nil {
					t.Fatal(e)
				}
			}
			if e := c.Close(); e != nil {
				t.Fatal(e)
			}
			b = make([]byte, 1)
			if n, e := s.Read(b); n != 0 || e != io.EOF {
				t.Fatalf("FIN: %d %v", n, e)
			}
			if e := c.Close(); e != io.ErrClosedPipe {
				t.Fatalf("second close %v", e)
			}
		})
	}
}
func TestOriginalVersionAdoptionAndConfigCopy(t *testing.T) {
	a, b := net.Pipe()
	defer b.Close()
	cfg := defaultConfigRecovered()
	cfg.DU0iTVAg = true
	s, e := CXFDjl_TYx8(a, cfg)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	hdr := []byte{2, 0, 0, 0, 7, 0, 0, 0}
	go b.Write(hdr)
	stream, e := s.AcceptStream()
	if e != nil {
		t.Fatal(e)
	}
	if stream.ID() != 7 || s.wpUkUgrJJHwN.RTNOROJC5 != 2 || cfg.RTNOROJC5 != 1 {
		t.Fatal("original version adoption or config copy lost")
	}
	go b.Write([]byte{0, 3, 0, 0, 0, 0, 0, 0})
	s.SetDeadline(time.Now().Add(time.Second))
	if _, e = s.AcceptStream(); e != invalidProtocolAt5097F40 {
		t.Fatalf("invalid protocol: %v", e)
	}
}
func TestOriginalConfigValidationAndDeadline(t *testing.T) {
	cfg := defaultConfigRecovered()
	if e := CaHL4gb(cfg); e != nil {
		t.Fatal(e)
	}
	for _, v := range []struct {
		mutate  func(*BWpFHH3x)
		message string
	}{
		{func(c *BWpFHH3x) { c.RTNOROJC5 = 3 }, "unsupported protocol version"},
		{func(c *BWpFHH3x) { c.PGgIC2 = 0 }, "keep-alive interval must be positive"},
		{func(c *BWpFHH3x) { c.PK9EpD = 1 }, "keep-alive timeout must be larger than keep-alive interval"},
		{func(c *BWpFHH3x) { c.XsqVEjw4 = 65536 }, "max frame size must not be larger than 65535"},
	} {
		c := *cfg
		v.mutate(&c)
		if e := CaHL4gb(&c); e == nil || e.Error() != v.message {
			t.Fatalf("%s: %v", v.message, e)
		}
	}
	s, _ := pair(t, 1)
	s.SetDeadline(time.Now().Add(-time.Second))
	if _, e := s.AcceptStream(); e != timeoutAt502B390 {
		t.Fatalf("deadline: %v", e)
	}
	s.SetDeadline(time.Time{})
	s.Close()
	if _, e := s.AcceptStream(); e != io.ErrClosedPipe {
		t.Fatalf("closed: %v", e)
	}
}

type batchWriter struct {
	mu     sync.Mutex
	writes [][]byte
	short  bool
	err    error
}

func (w *batchWriter) Read([]byte) (int, error) { return 0, io.EOF }
func (w *batchWriter) Close() error             { return nil }
func (w *batchWriter) Write(b []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.writes = append(w.writes, append([]byte(nil), b...))
	if w.err != nil {
		return 0, w.err
	}
	if w.short {
		return 0, nil
	}
	return len(b), nil
}
func TestOriginalMuxBatchingAndShortWriteContract(t *testing.T) {
	for _, short := range []bool{false, true} {
		w := &batchWriter{short: short}
		s := &DZOCfwWJsy{mlu21WY6: w, nHOcCGRza7V: make(chan struct{}), hHxKb51kQa: make(chan vwv8dqM, 16), nvN2Bu: make(chan struct{})}
		replies := []chan yoowfr{}
		for i := range 3 {
			ch := make(chan yoowfr, 1)
			replies = append(replies, ch)
			s.hHxKb51kQa <- vwv8dqM{fN47GQEn: TKUtOq{1, 2, uint32(i + 1), bytes.Repeat([]byte{byte(i)}, 10000)}, wbTHvwrgbD: ch}
		}
		done := make(chan struct{})
		go func() { s.hm9Q3Wo(); close(done) }()
		for _, ch := range replies {
			select {
			case r := <-ch:
				if r.b1C__L1 != 10000 || r.uRSOFlo != nil {
					t.Fatalf("original short-write contract: %+v", r)
				}
			case <-time.After(time.Second):
				t.Fatal("send blocked")
			}
		}
		close(s.nHOcCGRza7V)
		<-done
		if len(w.writes) != 2 || len(w.writes[0]) != 20016 || len(w.writes[1]) != 10008 {
			t.Fatalf("batch boundaries: %v", w.writes)
		}
		if binary.LittleEndian.Uint16(w.writes[0][2:]) != 10000 || binary.LittleEndian.Uint32(w.writes[0][4:]) != 1 {
			t.Fatal("wire header")
		}
	}
}
func TestOriginalWindowUpdateOnlyQueues(t *testing.T) {
	s := &DZOCfwWJsy{hHxKb51kQa: make(chan vwv8dqM, 16), wpUkUgrJJHwN: defaultConfigRecovered()}
	stream := newStream(9, 32768, s)
	// 无发送协程：原函数必须仅排队就返回，不能等待 result。
	done := make(chan error, 1)
	go func() { done <- stream.zH509K(42) }()
	select {
	case e := <-done:
		if e != nil {
			t.Fatal(e)
		}
	case <-time.After(time.Second):
		t.Fatal("unexpected synchronous ACK wait")
	}
	req := <-s.hHxKb51kQa
	if req.fN47GQEn.gUa4uFKR6 != 4 || binary.LittleEndian.Uint32(req.fN47GQEn.cAKGyUaj) != 42 {
		t.Fatal("ACK")
	}
}
