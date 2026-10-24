package pool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"io"
	"strings"
	"testing"
	"unsafe"
)

func TestBufferLayoutCompactionAndAllocation(t *testing.T) {
	value := LdHEmSs9(3)
	if value.Cap() != 4 || value.Len() != 0 {
		t.Fatal("value constructor returned wrong state")
	}
	value.Free()
	b := S4OQXWIZ(5)
	defer b.Free()
	if unsafe.Sizeof(*b) != 40 || unsafe.Offsetof(b.cnahaWq6) != 24 || unsafe.Offsetof(b.__XpzgdOCbm) != 32 || b.Cap() != 8 || b.Len() != 0 {
		t.Fatal("original layout/initial capacity differs")
	}
	b.WriteString("abcdef")
	b.Skip(4)
	before := &b.kNItjuFiwnR[0]
	b.GrowTo(8)
	if b.cnahaWq6 != 4 {
		t.Fatal("GrowTo compacted without growing")
	}
	b.WriteString("1234")
	if &b.kNItjuFiwnR[0] != before || b.cnahaWq6 != 0 || b.String() != "ef1234" {
		t.Fatalf("compaction %q, start %d", b.String(), b.cnahaWq6)
	}
	b.Skip(2)
	b.GrowTo(17)
	if b.Cap() != 32 || b.cnahaWq6 != 0 || b.String() != "1234" {
		t.Fatalf("allocation %q cap%d", b.String(), b.Cap())
	}
	b.GrowBy(35)
	if b.Cap() != 128 {
		t.Fatalf("must request old capacity + delta: got%d", b.Cap())
	}
	b.Free()
	b.Free()
	if b.Cap() != 0 || b.Len() != 0 || b.Bytes() != nil {
		t.Fatal("Free retained data")
	}
}

func TestBufferNetworkEndianFixture(t *testing.T) {
	b := S4OQXWIZ(0)
	defer b.Free()
	b.WriteByte(0xaa)
	b.WriteUint16(0x1234)
	b.WriteUint32(0x12345678)
	b.WriteUint64(0x0123456789abcdef)
	b.WriteUint16Le(0x1234)
	b.WriteUint32Le(0x12345678)
	b.WriteUint64Le(0x0123456789abcdef)
	want, _ := hex.DecodeString("aa1234123456780123456789abcdef341278563412efcdab8967452301")
	if !bytes.Equal(b.Bytes(), want) {
		t.Fatalf("wire bytes %x", b.Bytes())
	}
	if v, e := b.ReadUint8(); e != nil || v != 0xaa {
		t.Fatalf("u8 %x %v", v, e)
	}
	if v, e := b.ReadUint16(); e != nil || v != 0x1234 {
		t.Fatalf("u16 %x %v", v, e)
	}
	if v, e := b.ReadUint32(); e != nil || v != 0x12345678 {
		t.Fatalf("u32 %x %v", v, e)
	}
	if v, e := b.ReadUint64(); e != nil || v != 0x0123456789abcdef {
		t.Fatalf("u64 %x %v", v, e)
	}
	if v, e := b.ReadUint16Le(); e != nil || v != 0x1234 {
		t.Fatalf("u16le %x %v", v, e)
	}
	if v, e := b.ReadUint32Le(); e != nil || v != 0x12345678 {
		t.Fatalf("u32le %x %v", v, e)
	}
	if v, e := b.ReadUint64Le(); e != nil || v != 0x0123456789abcdef {
		t.Fatalf("u64le %x %v", v, e)
	}
	if !b.IsEmpty() || b.cnahaWq6 != 0 || b.__XpzgdOCbm != 0 {
		t.Fatal("exhaustion did not reset indices")
	}
	if _, e := b.ReadUint8(); e != io.ErrUnexpectedEOF {
		t.Fatalf("numeric empty %v", e)
	}
	if _, e := b.Read(nil); e != io.EOF {
		t.Fatalf("stream empty %v", e)
	}
}

func TestBufferViewsAndBounds(t *testing.T) {
	b := S4OQXWIZ(8)
	defer b.Free()
	b.WriteString("abcdef")
	b.Skip(2)
	if string(b.BytesRange(1, 3)) != "de" || string(b.BytesFrom(2)) != "ef" || string(b.BytesTo(2)) != "cd" {
		t.Fatal("relative views differ")
	}
	b.SetByte(0, 'C')
	if b.Byte(0) != 'C' {
		t.Fatal("Byte/SetByte indexing")
	}
	if len(b.BytesTo(6)) != 6 {
		t.Fatal("view restricted to readable length instead of capacity")
	}
	data, err := b.ReadBytes(2)
	if err != nil || string(data) != "Cd" || cap(data) != 6 {
		t.Fatalf("view %q cap%d %v", data, cap(data), err)
	}
	if _, err := b.ReadBytes(3); err != io.ErrUnexpectedEOF || b.Len() != 2 {
		t.Fatal("short ReadBytes consumed data")
	}
	if n := b.ReverseEnd(10); n != 2 || !b.IsEmpty() || b.cnahaWq6 != 4 {
		t.Fatalf("ReverseEnd %d start%d", n, b.cnahaWq6)
	}
	b.Reset()
	b.SkipEnd(3)
	if b.Len() != 3 {
		t.Fatal("SkipEnd")
	}
	for name, action := range map[string]func(){"negative read": func() { b.ReadBytes(-1) }, "view bounds": func() { b.BytesTo(100) }, "byte bounds": func() { b.Byte(100) }} {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("original bounds panic missing")
				}
			}()
			action()
		})
	}
	copyValue := b.String()
	b.SetByte(0, '!')
	if len(copyValue) > 0 && copyValue[0] == '!' {
		t.Fatal("String returned aliased bytes")
	}
}

type bufferReadFixture struct {
	chunks   []string
	terminal error
	calls    int
}

func (r *bufferReadFixture) Read(p []byte) (int, error) {
	r.calls++
	if len(r.chunks) == 0 {
		return 0, r.terminal
	}
	chunk := r.chunks[0]
	r.chunks = r.chunks[1:]
	n := copy(p, chunk)
	if len(r.chunks) == 0 {
		return n, r.terminal
	}
	return n, nil
}

func TestBufferReaderContracts(t *testing.T) {
	b := S4OQXWIZ(1)
	defer b.Free()
	if n, err := b.ReadFullFrom(strings.NewReader("ab"), 4); n != 2 || err != io.ErrUnexpectedEOF || b.String() != "ab" {
		t.Fatalf("ReadFull partial %d %v %q", n, err, b.String())
	}
	if n, err := b.ReadMaxFrom(strings.NewReader("cdef"), 2); n != 2 || err != nil || b.String() != "abcd" {
		t.Fatalf("ReadMax %d %v %q", n, err, b.String())
	}
	b.Reset()
	terminal := errors.New("reader terminal")
	r := &bufferReadFixture{chunks: []string{"", "xy"}, terminal: terminal}
	if n, err := b.ReadAllFrom(r, 2); n != 2 || err != terminal || b.String() != "xy" || r.calls != 2 {
		t.Fatalf("ReadAll %d %v %q calls%d", n, err, b.String(), r.calls)
	}
	b.Reset()
	r = &bufferReadFixture{chunks: []string{"done"}, terminal: io.EOF}
	if n, err := b.ReadFrom(r); n != 4 || err != nil || b.String() != "done" {
		t.Fatalf("ReadFrom %d %v %q", n, err, b.String())
	}
	if n, err := b.Read(nil); n != 0 || err != nil || b.Len() != 4 {
		t.Fatalf("zero read %d %v", n, err)
	}
}

func TestBufferAEADAndStreamFixtures(t *testing.T) {
	block, err := aes.NewCipher(make([]byte, 16))
	if err != nil {
		t.Fatal(err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal(err)
	}
	nonce := make([]byte, aead.NonceSize())
	plain := []byte("payload")
	additional := []byte("associated")
	expected := aead.Seal(nil, nonce, plain, additional)
	b := S4OQXWIZ(64)
	defer b.Free()
	b.WriteString("skip")
	b.Write(plain)
	b.Skip(4)
	b.Seal(aead, nonce, additional)
	if !bytes.Equal(b.Bytes(), expected) {
		t.Fatalf("Seal %x", b.Bytes())
	}
	if n, err := b.Open(aead, nonce, additional); n != len(plain) || err != nil || !bytes.Equal(b.Bytes(), plain) {
		t.Fatalf("Open %d %v %q", n, err, b.Bytes())
	}
	b.Reset()
	b.WriteString("prefix")
	if n := b.SealWrite(aead, nonce, plain, additional); n != len(expected) || !bytes.Equal(b.Bytes()[6:], expected) {
		t.Fatalf("SealWrite %d %x", n, b.Bytes())
	}
	b.Reset()
	b.WriteString("prefix")
	if n, err := b.OpenWrite(aead, nonce, expected, additional); n != len(plain) || err != nil || b.String() != "prefixpayload" {
		t.Fatalf("OpenWrite %d %v %q", n, err, b.String())
	}
	oldLength := b.Len()
	if _, err := b.OpenWrite(aead, nonce, []byte("invalid"), additional); err == nil || b.Len() != oldLength {
		t.Fatalf("OpenWrite auth failure %v len%d", err, b.Len())
	}
	b.Reset()
	b.WriteString("clearOLD")
	n := b.SealWriteWithStart(aead, 5, nonce, []byte("NEW"), additional)
	want := append([]byte("clear"), aead.Seal(nil, nonce, []byte("OLDNEW"), additional)...)
	if n != 3+aead.Overhead() || !bytes.Equal(b.Bytes(), want) {
		t.Fatalf("SealWriteWithStart %d %x", n, b.Bytes())
	}
	b.Reset()
	b.XORKeyStreamWrite(cipher.NewCTR(block, make([]byte, 16)), plain)
	out := make([]byte, 20)
	if n, err := b.XORKeyStreamRead(cipher.NewCTR(block, make([]byte, 16)), out); n != len(plain) || err != nil || !bytes.Equal(out[:n], plain) {
		t.Fatalf("stream %d %v %x", n, err, out)
	}
}
