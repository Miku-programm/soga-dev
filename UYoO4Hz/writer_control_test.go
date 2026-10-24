package UYoO4Hz

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	ws "recovered.local/soga/FgcbiePIJLm"
	"testing"
	"time"
)

func TestNativeWriterRealConnection(t *testing.T) {
	for _, state := range []ws.DwfKMV{ws.StateServerSide, ws.StateClientSide} {
		t.Run(map[ws.DwfKMV]string{ws.StateServerSide: "server", ws.StateClientSide: "client"}[state], func(t *testing.T) {
			a, b := net.Pipe()
			defer a.Close()
			defer b.Close()
			a.SetDeadline(time.Now().Add(5 * time.Second))
			b.SetDeadline(time.Now().Add(5 * time.Second))
			payload := bytes.Repeat([]byte("native-buffer-websocket"), 4100)
			result := make(chan error, 1)
			go func() {
				w := NewWriterBuffer(a, state, ws.OpBinary, make([]byte, 130))
				_, e := w.Write(payload[:17])
				if e == nil {
					_, e = w.Write(payload[17:])
				}
				if e == nil {
					e = w.Flush()
				}
				result <- e
			}()
			peer := ws.StateClientSide
			if state == ws.StateClientSide {
				peer = ws.StateServerSide
			}
			r := &HLKO3Cdqw{GX04GFKh1Gz: b, WXUh5psqY7: peer}
			h, e := r.NextFrame()
			if e != nil || h.IsHe8C != ws.OpBinary {
				t.Fatal(h, e)
			}
			got, e := io.ReadAll(r)
			if e != nil || !bytes.Equal(got, payload) {
				t.Fatal(len(got), e)
			}
			if e = <-result; e != nil {
				t.Fatal(e)
			}
		})
	}
}
func TestNativeWriterBoundariesAndPreservedState(t *testing.T) {
	var b bytes.Buffer
	w := NewWriterBuffer(&b, ws.StateServerSide, ws.OpBinary, make([]byte, 8))
	if e := w.Flush(); e != nil || b.Len() != 0 {
		t.Fatal(e, b.Bytes())
	}
	w.Write(nil)
	w.Flush()
	if !bytes.Equal(b.Bytes(), []byte{0x82, 0}) {
		t.Fatal(b.Bytes())
	}
	b.Reset()
	w.DisableFlush()
	p := bytes.Repeat([]byte{7}, 65536)
	if n, e := w.Write(p); e != nil || n != len(p) {
		t.Fatal(n, e)
	}
	w.Flush()
	if !bytes.Equal(b.Bytes()[:10], []byte{0x82, 127, 0, 0, 0, 0, 0, 1, 0, 0}) || !bytes.Equal(b.Bytes()[10:], p) {
		t.Fatal("growth or header")
	}
	err := errors.New("persistent original write error")
	w.vQT2iWLEILFo = err
	w.Reset(&b, ws.StateServerSide, ws.OpText)
	if _, e := w.Write([]byte("x")); e != err {
		t.Fatal("original Reset preserves err", e)
	}
}

type emptyNativeReader struct{}

func (emptyNativeReader) Read([]byte) (int, error) { return 0, nil }
func TestNativeWriterReadFromStall(t *testing.T) {
	var b bytes.Buffer
	w := NewWriterBuffer(&b, ws.StateServerSide, ws.OpBinary, make([]byte, 8))
	n, e := w.ReadFrom(emptyNativeReader{})
	if n != 0 || e != io.ErrNoProgress {
		t.Fatal(n, e)
	}
}
func TestNativeControlWireReplies(t *testing.T) {
	for _, mask := range []bool{false, true} {
		p := []byte("ping data")
		wire := append([]byte(nil), p...)
		key := [4]byte{1, 7, 9, 13}
		if mask {
			for i := range wire {
				wire[i] ^= key[i%4]
			}
		}
		var b bytes.Buffer
		c := GdvjKg5RudlP{Src: bytes.NewReader(wire), Dst: &b, State: ws.StateServerSide, DisableSrcCiphering: !mask}
		e := c.Handle(ws.Wmyq5K0sQ0w{IsHe8C: ws.OpPing, L4ixlgT8VpTo: int64(len(p)), BiCmXtyECQl: key})
		if e != nil || !bytes.Equal(b.Bytes(), append([]byte{0x8a, byte(len(p))}, p...)) {
			t.Fatal(b.Bytes(), e)
		}
	}
	for _, data := range [][]byte{nil, append(binary.BigEndian.AppendUint16(nil, 1000), []byte("bye")...), {0x03, 0xed}, {0x03}} {
		var b bytes.Buffer
		c := GdvjKg5RudlP{Src: bytes.NewReader(data), Dst: &b, State: ws.StateServerSide, DisableSrcCiphering: true}
		e := c.Handle(ws.Wmyq5K0sQ0w{IsHe8C: ws.OpClose, L4ixlgT8VpTo: int64(len(data))})
		switch len(data) {
		case 0:
			if e.Error() != "ws closed: 1005 " || !bytes.Equal(b.Bytes(), []byte{0x88, 0}) {
				t.Fatal(b.Bytes(), e)
			}
		case 5:
			if e.Error() != "ws closed: 1000 bye" || !bytes.Equal(b.Bytes(), []byte{0x88, 2, 3, 232}) {
				t.Fatal(b.Bytes(), e)
			}
		default:
			if e == nil || b.Len() < 4 || !bytes.Equal(b.Bytes()[:4], []byte{0x88, byte(b.Len() - 2), 3, 234}) {
				t.Fatal(b.Bytes(), e)
			}
		}
	}
	var b bytes.Buffer
	c := GdvjKg5RudlP{Src: bytes.NewReader([]byte("pong")), Dst: &b, State: ws.StateServerSide}
	if e := c.Handle(ws.Wmyq5K0sQ0w{IsHe8C: ws.OpPong, L4ixlgT8VpTo: 4}); e != nil || b.Len() != 0 {
		t.Fatal(e, b.Bytes())
	}
	if e := c.Handle(ws.Wmyq5K0sQ0w{IsHe8C: 2}); e != ErrNotControlFrame {
		t.Fatal(e)
	}
}
func TestNativeControlCallbackOnFragment(t *testing.T) {
	wire := wsWire(2, false, []byte("one"))
	wire = append(wire, wsWire(9, true, []byte("check"))...)
	wire = append(wire, wsWire(0, true, []byte("two"))...)
	var reply bytes.Buffer
	r := &HLKO3Cdqw{GX04GFKh1Gz: bytes.NewReader(wire), WXUh5psqY7: ws.StateServerSide, EbVEHBvY4lh: L9DiQX(&reply, ws.StateServerSide)}
	if _, e := r.NextFrame(); e != nil {
		t.Fatal(e)
	}
	p, e := io.ReadAll(r)
	if e != nil || string(p) != "onetwo" || !bytes.Equal(reply.Bytes(), []byte{138, 5, 'c', 'h', 'e', 'c', 'k'}) {
		t.Fatal(string(p), reply.Bytes(), e)
	}
}
