package TaFdeUcnzQ

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"net"
	J "recovered.local/soga/JafvCs5"
	S "recovered.local/soga/ad5KEbw_EzWk"
	H "recovered.local/soga/hfOuV8fOe4L"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func infoFor(method H.ZvAxOt1l7Ar) *S.L7E7H2aW8KVs {
	return &S.L7E7H2aW8KVs{MY3lagqY: "original-native-SSR", L6BrlCXd: method}
}
func TestNativeCipherLayoutsAndAlgorithms(t *testing.T) {
	if unsafe.Sizeof(nfQqag72aa{}) != 32 || unsafe.Sizeof(KrOWYtzUYft{}) != 184 || unsafe.Sizeof(jSUbF0{}) != 216 || unsafe.Sizeof(jUB7SuavZpra{}) != 72 {
		t.Fatal("RTTI")
	}
	if len(cipherMapAt5096960) != 28 {
		t.Fatal("original algorithm map")
	}
	for name, spec := range cipherMapAt5096960 {
		t.Run(string(name), func(t *testing.T) {
			var wire bytes.Buffer
			writer, e := VavyqXsqy9V(infoFor(name), &wire)
			if e != nil {
				t.Fatal(e)
			}
			defer writer.Free()
			payload := make([]byte, 40001)
			for i := range payload {
				payload[i] = byte(i*17 + 9)
			}
			for _, p := range [][]byte{payload[:1], payload[1:126], payload[126:]} {
				if n, e := writer.Write(p); n != len(p) || e != nil {
					t.Fatal(n, e)
				}
			}
			reader, e := VavyqXsqy9V(infoFor(name), &wire)
			if e != nil {
				t.Fatal(e)
			}
			defer reader.Free()
			if reader.KeyLen() != spec.geqiu3e || !bytes.Equal(reader.GetKey(), writer.GetKey()) {
				t.Fatal("key")
			}
			got := make([]byte, len(payload))
			for n := 0; n < len(got); {
				end := min(n+117, len(got))
				m, e := reader.Read(got[n:end])
				if e != nil {
					t.Fatal(n, e)
				}
				if m == 0 {
					t.Fatal("stalled")
				}
				n += m
			}
			if !bytes.Equal(got, payload) {
				t.Fatal("stream roundtrip")
			}
		})
	}
}

type packetFixture struct {
	data    []byte
	err     error
	written []byte
}

func (p *packetFixture) ReadFrom(b []byte) (int, net.Addr, error) {
	n := copy(b, p.data)
	return n, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 325}, p.err
}
func (p *packetFixture) WriteTo(b []byte, _ net.Addr) (int, error) {
	p.written = bytes.Clone(b)
	return len(b), p.err
}
func (*packetFixture) Close() error                     { return nil }
func (*packetFixture) LocalAddr() net.Addr              { return &net.UDPAddr{} }
func (*packetFixture) SetDeadline(time.Time) error      { return nil }
func (*packetFixture) SetReadDeadline(time.Time) error  { return nil }
func (*packetFixture) SetWriteDeadline(time.Time) error { return nil }
func TestNativeCipherDatagramsAndReplay(t *testing.T) {
	for name, spec := range cipherMapAt5096960 {
		t.Run(string(name), func(t *testing.T) {
			pc := &packetFixture{}
			writer, e := DP5vq3(infoFor(name), pc)
			if e != nil {
				t.Fatal(e)
			}
			defer writer.Free()
			reader, e := DP5vq3(infoFor(name), pc)
			if e != nil {
				t.Fatal(e)
			}
			defer reader.Free()
			payload := bytes.Repeat([]byte("native UDP payload"), 39)
			for i := 0; i < 2; i++ {
				if n, e := writer.WriteTo(payload, &net.UDPAddr{}); n != len(payload) || e != nil {
					t.Fatal(n, e)
				}
				pc.data = pc.written
				dst := make([]byte, 4096)
				n, addr, e := reader.ReadFrom(dst)
				if e != nil || addr == nil || n != len(payload) || !bytes.Equal(dst[:n], payload) {
					t.Fatal(n, addr, e)
				}
				n, addr, e = reader.ReadFrom(dst)
				if spec.w5pQfi_Ptek != nil {
					if n != 0 || addr != nil || e == nil || e.Error() != "aead stream salt repeat" {
						t.Fatal("AEAD replay", n, addr, e)
					}
				} else if e != nil || n != len(payload) {
					t.Fatal("stream replay must only log", n, e)
				}
			}
		})
	}
}

type readWriterFixture struct {
	input             []byte
	readErr, writeErr error
	shortWrite        bool
	maxRead           int
}

func (f *readWriterFixture) Read(p []byte) (int, error) {
	if f.maxRead > 0 && len(p) > f.maxRead {
		p = p[:f.maxRead]
	}
	n := copy(p, f.input)
	f.input = f.input[n:]
	return n, f.readErr
}
func (f *readWriterFixture) Write(p []byte) (int, error) {
	if f.shortWrite {
		return 0, f.writeErr
	}
	return len(p), f.writeErr
}
func TestNativeCipherErrorAndShortIO(t *testing.T) {
	if _, e := VavyqXsqy9V(&S.L7E7H2aW8KVs{}, nil); e != errorAt5098DB0 {
		t.Fatal(e)
	}
	if _, e := VavyqXsqy9V(&S.L7E7H2aW8KVs{MY3lagqY: "x"}, nil); e != errorAt5098DC0 {
		t.Fatal(e)
	}
	if _, e := VavyqXsqy9V(infoFor("AES-128-CFB"), nil); e == nil || e.Error() != "Unsupported encryption method: AES-128-CFB" {
		t.Fatal(e)
	}
	sentinel := errors.New("reader result")
	for _, name := range []H.ZvAxOt1l7Ar{"aes-128-cfb", "aes-128-gcm"} {
		f := &readWriterFixture{input: make([]byte, 128), maxRead: 7}
		c, e := VavyqXsqy9V(infoFor(name), f)
		if e != nil {
			t.Fatal(e)
		}
		_, e = c.Read(make([]byte, 12))
		want := "iv len not enough: 7"
		if name == "aes-128-gcm" {
			want = "salt len not enough: 7"
		}
		if e == nil || e.Error() != want {
			t.Fatal(e)
		}
		f = &readWriterFixture{input: make([]byte, 128), readErr: sentinel}
		c, _ = VavyqXsqy9V(infoFor(name), f)
		if n, e := c.Read(make([]byte, 12)); n != 0 || e != sentinel {
			t.Fatal(n, e)
		}
		f = &readWriterFixture{shortWrite: true}
		c, _ = VavyqXsqy9V(infoFor(name), f)
		if n, e := c.Write([]byte("successful short write")); n != 22 || e != nil {
			t.Fatal(n, e)
		}
	}
	for _, fn := range []func([]byte, []byte, IUgK3nW) (cipher.Stream, error){g57NcXia7, lPcdRtmHNa, dzRpwKaY} {
		s, e := fn(nil, nil, 0)
		if e == nil || s == nil || !reflect.ValueOf(s).IsNil() {
			t.Fatal("typed nil must be preserved", s, e)
		}
	}
	backing := make([]byte, 32)
	iv := make([]byte, 8)
	if _, e := iuskSmap(backing[:0], iv[:0], 1); e != nil {
		t.Fatal(e)
	}
	d := []byte{9}
	(&zclWYvDGjqa2{}).XORKeyStream(d, []byte{1, 2})
	if d[0] != 1 {
		t.Fatal(d)
	}
}

// Frames are constructed independently using the registered primitive and a
// manually encoded length, rather than exercising only our matching writer.
func makeFrame(t *testing.T, name H.ZvAxOt1l7Ar, size int, corrupt bool) (*KrOWYtzUYft, []byte) {
	t.Helper()
	spec := cipherMapAt5096960[name]
	salt := make([]byte, max(16, spec.geqiu3e))
	rand.Read(salt)
	key := make([]byte, spec.geqiu3e)
	J.THj8k5(J.Is8IqhjVE6("original-native-SSR", spec.geqiu3e), salt, []byte("ss-subkey"), key)
	a, e := spec.w5pQfi_Ptek(key)
	if e != nil {
		t.Fatal(e)
	}
	nonce := make([]byte, a.NonceSize())
	length := []byte{byte(size >> 8), byte(size)}
	frame := append([]byte{}, salt...)
	frame = a.Seal(frame, nonce, length, nil)
	nonce[0] = 1
	payload := bytes.Repeat([]byte{0xa7}, size)
	frame = a.Seal(frame, nonce, payload, nil)
	if corrupt {
		frame[len(salt)+2] ^= 1
	}
	c, e := VavyqXsqy9V(infoFor(name), bytes.NewBuffer(frame))
	if e != nil {
		t.Fatal(e)
	}
	return c, payload
}
func TestNativeAEADLengthBoundaryAndFailureNonce(t *testing.T) {
	for _, name := range []H.ZvAxOt1l7Ar{"aes-128-gcm", "chacha20-ietf-poly1305"} {
		for _, size := range []int{0, 1, 16383, 16384, 32766} {
			c, want := makeFrame(t, name, size, false)
			got := make([]byte, size)
			if size == 0 {
				if n, e := c.Read(got); n != 0 || e != nil {
					t.Fatal(n, e)
				}
			} else if _, e := io.ReadFull(c, got); e != nil {
				t.Fatal(size, e)
			}
			if !bytes.Equal(got, want) {
				t.Fatal(size)
			}
			if c.s_53ql.qaeiJJzh3hj[0] != 2 {
				t.Fatal("nonce")
			}
			c.s_53ql.Free()
		}
		c, _ := makeFrame(t, name, 32767, false)
		if n, e := c.Read(make([]byte, 32767)); n != 0 || e == nil || e.Error() != "payload size too large: 32767" {
			t.Fatal(n, e)
		}
		if c.s_53ql.qaeiJJzh3hj[0] != 1 {
			t.Fatal("oversized length nonce")
		}
		c.s_53ql.Free()
		c, _ = makeFrame(t, name, 10, true)
		if n, e := c.Read(make([]byte, 10)); n != 0 || e == nil {
			t.Fatal(n, e)
		}
		if c.s_53ql.qaeiJJzh3hj[0] != 1 {
			t.Fatal("failed Open nonce")
		}
		c.s_53ql.Free()
	}
	n := []byte{255, 255, 0}
	incrementNonce(n)
	if !bytes.Equal(n, []byte{0, 0, 1}) {
		t.Fatal(n)
	}
}
func TestNativeSalsaChunkingAndCapacity(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 8)
	rand.Read(key)
	rand.Read(nonce)
	a, _ := iuskSmap(key, nonce, 1)
	b, _ := iuskSmap(key, nonce, 0)
	plain := make([]byte, 279)
	rand.Read(plain)
	whole := make([]byte, len(plain))
	a.XORKeyStream(whole, plain)
	got := make([]byte, 0, len(plain))
	pos := 0
	for _, size := range []int{1, 63, 1, 65, 149} {
		p := make([]byte, size)
		b.XORKeyStream(p, plain[pos:pos+size])
		got = append(got, p...)
		pos += size
	}
	if !bytes.Equal(got, whole) {
		t.Fatal("salsa offset")
	}
	// A short destination consumes the full input; the original does not panic.
	c, _ := iuskSmap(key, nonce, 1)
	small := make([]byte, 3)
	c.XORKeyStream(small, plain[:19])
	if !bytes.Equal(small, whole[:3]) || c.(*jUB7SuavZpra).cO2TNZUn7xB != 19 {
		t.Fatal("short destination")
	}
}
