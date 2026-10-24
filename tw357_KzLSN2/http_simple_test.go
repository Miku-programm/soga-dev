package tw357_KzLSN2

import (
	"bytes"
	"errors"
	"io"
	S "recovered.local/soga/ad5KEbw_EzWk"
	"strings"
	"testing"
	"unsafe"
)

type wireFixture struct {
	read     *bytes.Reader
	output   bytes.Buffer
	writeErr error
	readErr  error
	short    bool
}

func (w *wireFixture) Read(p []byte) (int, error) {
	n, e := w.read.Read(p)
	if w.readErr != nil {
		return n, w.readErr
	}
	return n, e
}
func (w *wireFixture) Write(p []byte) (int, error) {
	w.output.Write(p)
	if w.short {
		return 0, w.writeErr
	}
	return len(p), w.writeErr
}
func TestHTTPObfsOriginalConstructorsAndPlain(t *testing.T) {
	if unsafe.Sizeof(rqkJYP_xJtF{}) != 16 || unsafe.Sizeof(nArN8vNMTaR{}) != 56 {
		t.Fatal("RTTI")
	}
	for _, ctor := range []q1_ASanae22{zbGjqkTbNkOk, ujOdSji25S} {
		c := ctor(&wireFixture{}).(*nArN8vNMTaR)
		if c.mI41hq || c.a1yGpXg6Y || c.q60_aN || c.z9a8K4aeGWS != nil || c.GetOverhead() != 0 {
			t.Fatal("original constructor")
		}
	}
	var b bytes.Buffer
	p := aR_Feln8A36(&b)
	p.SetServerInfo(nil)
	if p.GetOverhead() != 0 {
		t.Fatal("overhead")
	}
	p.Write([]byte("plain"))
	out := make([]byte, 5)
	if n, e := p.Read(out); n != 5 || e != nil || string(out) != "plain" {
		t.Fatal(n, e, string(out))
	}
}
func TestHTTPSimpleHeaderAndPayload(t *testing.T) {
	for _, method := range []string{"GET", "POST"} {
		for _, host := range []string{"example.com:8443", "[::1]:443", "example.com"} {
			frame := method + " /%41%42%43 HTTP/1.1\r\nHost: " + host + "\r\nContent-Length: 5\r\n\r\nTAIL!"
			w := &wireFixture{read: bytes.NewReader([]byte(frame))}
			c := ujOdSji25S(w).(*nArN8vNMTaR)
			info := &S.L7E7H2aW8KVs{}
			c.SetServerInfo(info)
			var got []byte
			for len(got) < 8 {
				out := make([]byte, 2)
				n, e := c.Read(out)
				if e != nil {
					t.Fatal(e)
				}
				got = append(got, out[:n]...)
			}
			if string(got) != "ABCTAIL!" || c.q60_aN {
				t.Fatal(got, c.q60_aN)
			}
			wantHost := "example.com"
			if strings.HasPrefix(host, "[") {
				wantHost = "::1"
			}
			if info.EFLMTH != wantHost {
				t.Fatal(info.EFLMTH)
			}
			if n, e := c.Read(make([]byte, 1)); n != 0 || e != io.EOF {
				t.Fatal(n, e)
			}
		}
	}
}
func TestHTTPSimpleFirstReadAndFlags(t *testing.T) {
	for _, wire := range []string{"GET / HTTP/1.1\r\nHost: h\r\n", "PUT / HTTP/1.1\r\nHost: h\r\n\r\n", "G", ""} {
		c := ujOdSji25S(&wireFixture{read: bytes.NewReader([]byte(wire))}).(*nArN8vNMTaR)
		c.SetServerInfo(&S.L7E7H2aW8KVs{})
		if n, e := c.Read(make([]byte, 8)); n != 0 || e == nil {
			t.Fatal(n, e)
		}
		if c.a1yGpXg6Y {
			t.Fatal("header flag before complete header")
		}
	}
	sentinel := errors.New("data with error")
	c := ujOdSji25S(&wireFixture{read: bytes.NewReader([]byte("GET / HTTP/1.1\r\n\r\n")), readErr: sentinel}).(*nArN8vNMTaR)
	if n, e := c.Read(make([]byte, 8)); n != 0 || e != sentinel {
		t.Fatal(n, e)
	}
	// The flag is set before parser validation, even when the request is bad.
	c = ujOdSji25S(&wireFixture{read: bytes.NewReader([]byte("GET / HTTP/x\r\n\r\n"))}).(*nArN8vNMTaR)
	if _, e := c.Read(make([]byte, 8)); e == nil || !c.a1yGpXg6Y {
		t.Fatal(e, c.a1yGpXg6Y)
	}
}
func TestHTTPSimpleWriteHeaderAndErrorState(t *testing.T) {
	for _, fail := range []bool{false, true} {
		w := &wireFixture{short: true}
		if fail {
			w.writeErr = errors.New("write")
		}
		c := ujOdSji25S(w).(*nArN8vNMTaR)
		n, e := c.Write([]byte("payload"))
		if !c.mI41hq || !strings.HasPrefix(w.output.String(), "HTTP/1.1 200 OK\r\nConnection: keep-alive\r\n") || !strings.HasSuffix(w.output.String(), "\r\n\r\npayload") {
			t.Fatal("response")
		}
		if fail {
			if n != 0 || e != w.writeErr {
				t.Fatal(n, e)
			}
		} else if n != 7 || e != nil {
			t.Fatal(n, e)
		}
		w.output.Reset()
		n, e = c.Write([]byte("raw"))
		if n != 0 || e != w.writeErr || w.output.String() != "raw" {
			t.Fatal("later writes propagate raw counts", n, e)
		}
	}
}
