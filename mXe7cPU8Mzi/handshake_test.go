package mXe7cPU8Mzi

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/pbkdf2"
	"hash/fnv"
	"io"
	L "recovered.local/soga/LFWTTXF7JR"
	U "recovered.local/soga/nzATvajt"
	"testing"
	"time"
	"unsafe"
)

func testMieruUser(t *testing.T, id int64, name, password string) *U.WIZQ28 {
	t.Helper()
	u, err := U.U1KSvhrlJFSB(&L.Aga91y{WZ89pTAOrHK4: id, JAfXWwzYrW: &L.DO9qqM9{Ejujyaj: name, GltkJGvV: password}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	return u
}
func usersIterator(users ...*U.WIZQ28) func(func(*U.WIZQ28) bool) {
	return func(f func(*U.WIZQ28) bool) {
		for _, u := range users {
			if !f(u) {
				break
			}
		}
	}
}

// Client fixture uses raw wire offsets and independent key construction.
func firstClientFrame(t *testing.T, user *U.WIZQ28, slot int, hint bool) ([]byte, []byte) {
	t.Helper()
	nonce := make([]byte, 24)
	for i := range nonce {
		nonce[i] = byte(i + slot + 8)
	}
	if hint {
		sum := sha256.Sum256(append([]byte(user.GetUsername()), nonce[:16]...))
		copy(nonce[20:], sum[:4])
	}
	raw := make([]byte, 32)
	raw[0] = 2
	binary.BigEndian.PutUint32(raw[2:6], uint32(time.Now().Unix()/60))
	binary.BigEndian.PutUint32(raw[6:10], 0x12345678)
	material := append([]byte(user.GetPassword()), 0)
	material = append(material, []byte(user.GetUsername())...)
	secret := sha256.Sum256(material)
	var timestamp [8]byte
	binary.BigEndian.PutUint64(timestamp[:], uint64(time.Now().Round(2*time.Minute).Add(time.Duration(slot)*2*time.Minute).Unix()))
	salt := sha256.Sum256(timestamp[:])
	key := pbkdf2.Key(secret[:], salt[:], 64, 32, sha256.New)
	a, err := chacha20poly1305.NewX(key)
	if err != nil {
		t.Fatal(err)
	}
	return append(nonce, a.Seal(nil, nonce, raw, nil)...), raw
}

func TestHandshakeMatchingAndFallback(t *testing.T) {
	user := testMieruUser(t, 91, "alice", "secret")
	wrong := testMieruUser(t, 92, "alice", "wrong")
	other := testMieruUser(t, 93, "bob", "secret")
	for _, slot := range []int{-1, 0, 1} {
		for _, hint := range []bool{false, true} {
			wire, want := firstClientFrame(t, user, slot, hint)
			got, a, plain, err := ofDJn56sP0(wire[:24], wire[24:], usersIterator(other, wrong, user))
			if err != nil || got != user || a == nil || !bytes.Equal(plain, want) {
				t.Fatalf("slot %d hint %v: %v %x", slot, hint, err, plain)
			}
		}
	}
	wire, _ := firstClientFrame(t, user, 0, true)
	got, a, plain, err := ofDJn56sP0(wire[:24], wire[24:], usersIterator(wrong, other))
	if got != nil || a != nil || plain != nil || err == nil || err.Error() != "decrypt failed: no matching user" {
		t.Fatalf("wrong password %v", err)
	}
}

type deadlineFrameConn struct {
	frameMemoryConn
	deadlines []time.Time
}

func (c *deadlineFrameConn) SetReadDeadline(v time.Time) error {
	c.deadlines = append(c.deadlines, v)
	return nil
}
func TestHandshakeDeadlineReplayAndFailureOrder(t *testing.T) {
	old := yCMyVg()
	replayAt5096328 = b3L9YozU4h(1024, time.Hour)
	t.Cleanup(func() { replayAt5096328 = old })
	user := testMieruUser(t, 94, "user", "pw")
	wire, _ := firstClientFrame(t, user, 0, true)
	for repeat := 0; repeat < 2; repeat++ {
		conn := new(deadlineFrameConn)
		conn.Buffer.Write(wire)
		r, got, m, err := cMGKnW(conn, usersIterator(user))
		if repeat == 0 {
			if err != nil || r == nil || got != user || m.HNaaDh != 0x12345678 || !bytes.Equal(r.xUtUMCF3grV, wire[:24]) {
				t.Fatalf("first handshake %#v %v", m, err)
			}
		} else if r != nil || got != nil || m != nil || err == nil || err.Error() != "replay attack detected (decrypt succeeded)" {
			t.Fatalf("replay %v", err)
		}
		wantCalls := 2
		if repeat == 1 {
			wantCalls = 3
		}
		if len(conn.deadlines) != wantCalls || !conn.deadlines[len(conn.deadlines)-1].IsZero() || time.Until(conn.deadlines[0]) < 28*time.Second {
			t.Fatalf("deadlines %v", conn.deadlines)
		}
	}
	corrupt := append([]byte(nil), wire...)
	corrupt[50] ^= 1
	conn := new(deadlineFrameConn)
	conn.Buffer.Write(corrupt)
	_, _, _, err := cMGKnW(conn, usersIterator(user))
	if err == nil || err.Error() != "replay attack detected (decrypt failed)" {
		t.Fatalf("replay priority %v", err)
	}
	short := new(deadlineFrameConn)
	short.Buffer.Write(wire[:71])
	_, _, _, err = cMGKnW(short, usersIterator(user))
	if !errors.Is(err, io.ErrUnexpectedEOF) || err.Error() != "read first segment: unexpected EOF" || !short.deadlines[1].IsZero() {
		t.Fatalf("short frame %v", err)
	}
}
func TestReplayGenerationsAndExpiry(t *testing.T) {
	if unsafe.Sizeof(fA9rSx{}) != 64 {
		t.Fatal("RTTI layout")
	}
	data := []byte("1234567890123456suffix")
	h := fnv.New64a()
	h.Write(data)
	if ijqr7F(data) != h.Sum64() {
		t.Fatal("FNV mismatch")
	}
	r := b3L9YozU4h(2, time.Hour)
	if r.n7hmYhVTVA(data[:15]) || len(r.cgoZbeQ) != 0 {
		t.Fatal("short stored")
	}
	if r.n7hmYhVTVA(data) || !r.n7hmYhVTVA(append(append([]byte(nil), data[:16]...), "different"...)) {
		t.Fatal("fingerprint ignored suffix")
	}
	second := bytes.Repeat([]byte{2}, 16)
	third := bytes.Repeat([]byte{3}, 16)
	if r.n7hmYhVTVA(second) || r.n7hmYhVTVA(third) || !r.n7hmYhVTVA(data) {
		t.Fatal("generation preservation")
	}
	r.ljZIb5e9OR = time.Now().Add(-2 * time.Hour)
	if r.n7hmYhVTVA(data) || len(r.vksDKih) != 0 {
		t.Fatal("old generations retained")
	}
	r.ljZIb5e9OR = time.Now().Add(-time.Minute)
	if !r.n7hmYhVTVA(data) {
		t.Fatal("single expiry lost previous generation")
	}
}
