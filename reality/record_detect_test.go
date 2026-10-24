package reality

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestForkDetectReadRejectsTruncatedRecord(t *testing.T) {
	for _, tt := range []struct {
		name string
		data []byte
		want []int
	}{
		{"complete then truncated", []byte{23, 3, 3, 0, 2, 8, 9, 23, 3, 3, 0, 20, 1, 2}, []int{7}},
		{"truncated only", []byte{23, 3, 3, 0, 20, 1}, nil},
		{"two records", []byte{23, 3, 3, 0, 1, 8, 23, 3, 3, 0, 0}, []int{6, 5}},
		{"wrong type", []byte{22, 3, 3, 0, 0}, nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			a, b := net.Pipe()
			defer a.Close()
			done := make(chan struct{})
			go func() { defer close(done); b.Write(tt.data); b.Close() }()
			key := t.Name()
			defer GlobalPostHandshakeRecordsLens.Delete(key)
			c := &DetectConn{DrpXao: a, Key: key, CcsSent: true}
			n, e := c.Read(make([]byte, 10))
			if n != 0 || e != io.EOF {
				t.Fatalf("read = %d %v", n, e)
			}
			v, ok := GlobalPostHandshakeRecordsLens.Load(key)
			if !ok || !reflect.DeepEqual(v, tt.want) {
				t.Fatalf("lengths = %#v", v)
			}
			<-done
		})
	}
}

func TestDetectPassthroughAndCCS(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	c := &DetectConn{DrpXao: a}
	done := make(chan error, 1)
	go func() {
		_, e := b.Write([]byte("plain"))
		if e != nil {
			done <- e
			return
		}
		v := make([]byte, 3)
		_, e = io.ReadFull(b, v)
		if e == nil && !bytes.Equal(v, []byte{20, 3, 3}) {
			e = fmt.Errorf("wrong wire bytes")
		}
		done <- e
	}()
	p := make([]byte, 5)
	if _, e := io.ReadFull(c, p); e != nil || string(p) != "plain" {
		t.Fatalf("passthrough %q %v", p, e)
	}
	if _, e := c.Write([]byte{20, 3, 3}); e != nil || !c.CcsSent {
		t.Fatalf("CCS %v %v", c.CcsSent, e)
	}
	if e := <-done; e != nil {
		t.Fatal(e)
	}
}

func waitDetectResults(t *testing.T, cfg *Config) {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		complete := true
		for alpn := 0; alpn < 3; alpn++ {
			v, ok := GlobalPostHandshakeRecordsLens.Load(cfg.Dest + " probe.invalid " + strconv.Itoa(alpn))
			_, typed := v.([]int)
			complete = complete && ok && typed
		}
		if complete {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatal("probe results not published")
}
func clearDetectResults(cfg *Config) {
	for alpn := 0; alpn < 3; alpn++ {
		GlobalPostHandshakeRecordsLens.Delete(cfg.Dest + " probe.invalid " + strconv.Itoa(alpn))
	}
}

type closeObservedConn struct {
	net.Conn
	closed *atomic.Int32
}

func (c *closeObservedConn) Close() error { c.closed.Add(1); return c.Conn.Close() }

func TestForkDetectCustomDialAndDeferredClose(t *testing.T) {
	var calls, closes atomic.Int32
	errs := make(chan error, 3)
	cfg := &Config{Type: "original-custom-network", Dest: t.Name(), ServerNames: map[string]bool{"probe.invalid": true}}
	cfg.DialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
		calls.Add(1)
		if ctx == nil || ctx.Err() != nil || network != cfg.Type || address != cfg.Dest {
			errs <- errors.New("dial arguments differ")
		} else {
			errs <- nil
		}
		a, b := net.Pipe()
		b.Close()
		return &closeObservedConn{a, &closes}, nil
	}
	defer clearDetectResults(cfg)
	DetectPostHandshakeRecordsLens(cfg)
	waitDetectResults(t, cfg)
	if calls.Load() != 3 || closes.Load() != 3 {
		t.Fatalf("calls=%d closes=%d", calls.Load(), closes.Load())
	}
	for i := 0; i < 3; i++ {
		if e := <-errs; e != nil {
			t.Fatal(e)
		}
	}
	DetectPostHandshakeRecordsLens(cfg)
	if calls.Load() != 3 {
		t.Fatal("cached targets probed again")
	}
}

func TestForkDetectNilDialUsesTCP(t *testing.T) {
	listener, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer listener.Close()
	done := make(chan error, 1)
	go func() {
		for i := 0; i < 3; i++ {
			c, e := listener.Accept()
			if e != nil {
				done <- e
				return
			}
			c.Close()
		}
		done <- nil
	}()
	cfg := &Config{Type: "not-a-valid-network", Dest: listener.Addr().String(), ServerNames: map[string]bool{"probe.invalid": true}}
	defer clearDetectResults(cfg)
	DetectPostHandshakeRecordsLens(cfg)
	waitDetectResults(t, cfg)
	select {
	case e := <-done:
		if e != nil {
			t.Fatal(e)
		}
	case <-time.After(time.Second):
		t.Fatal("nil dial path did not use TCP")
	}
}
