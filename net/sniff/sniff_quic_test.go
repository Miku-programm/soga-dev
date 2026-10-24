package sniff

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io"
	"reflect"
	"testing"
)

// 独立测试端按 wire 格式生成加密 Initial，不调用待测 HKDF 或 varint helper。
func testQUICLabel(secret []byte, label string, n int) []byte {
	info := []byte{byte(n >> 8), byte(n), byte(len(label) + 6)}
	info = append(info, []byte("tls13 "+label)...)
	info = append(info, 0)
	h := hmac.New(sha256.New, secret)
	h.Write(info)
	h.Write([]byte{1})
	return h.Sum(nil)[:n]
}
func testQUICVarint(v int) []byte {
	if v < 64 {
		return []byte{byte(v)}
	}
	return []byte{byte(v>>8) | 0x40, byte(v)}
}
func testQUICPacket(t *testing.T, version uint32, pn byte, pnLen int, frames []byte) ([]byte, int) {
	t.Helper()
	saltText := "38762cf7f55934b34d179ae6a4c80cadccbb7f0a"
	if version != 1 {
		saltText = "afbfec289993d24c9e9786f19c6111e04390a899"
	}
	salt, _ := hex.DecodeString(saltText)
	dcid := []byte{0x83, 0x94, 0xc8, 0xf0, 0x3e, 0x51, 0x57, 0x08}
	extract := hmac.New(sha256.New, salt)
	extract.Write(dcid)
	secret := testQUICLabel(extract.Sum(nil), "client in", 32)
	header := []byte{0xc0 | byte(pnLen-1), 0, 0, 0, 0, byte(len(dcid))}
	binary.BigEndian.PutUint32(header[1:], version)
	header = append(header, dcid...)
	header = append(header, 0, 0)
	header = append(header, testQUICVarint(pnLen+len(frames)+16)...)
	pnOffset := len(header)
	header = append(header, make([]byte, pnLen)...)
	header[len(header)-1] = pn
	block, _ := aes.NewCipher(testQUICLabel(secret, "quic key", 16))
	aead, _ := cipher.NewGCM(block)
	nonce := testQUICLabel(secret, "quic iv", 12)
	nonce[11] ^= pn
	packet := append([]byte{}, header...)
	packet = aead.Seal(packet, nonce, frames, header)
	block, _ = aes.NewCipher(testQUICLabel(secret, "quic hp", 16))
	var mask [16]byte
	block.Encrypt(mask[:], packet[pnOffset+4:pnOffset+20])
	packet[0] ^= mask[0] & 15
	for i := 0; i < pnLen; i++ {
		packet[pnOffset+i] ^= mask[i+1]
	}
	return packet, pnOffset
}
func testCryptoFrame(offset int, data []byte) []byte {
	b := []byte{6}
	b = append(b, testQUICVarint(offset)...)
	b = append(b, testQUICVarint(len(data))...)
	return append(b, data...)
}
func TestQUICOriginalInitialVersionsFramesAndInputMutation(t *testing.T) {
	hello := captureHello(t)[5:]
	for _, version := range []uint32{1, 0xff00001d} {
		for _, pn := range []byte{0, 1} {
			// PADDING/PING, ACK_ECN with a range, reordered contiguous CRYPTO, CONNECTION_CLOSE。
			frames := []byte{0, 0, 1, 3, 1, 0, 1, 0, 0, 0, 0, 0, 0}
			frames = append(frames, testCryptoFrame(30, hello[30:])...)
			frames = append(frames, testCryptoFrame(0, hello[:30])...)
			frames = append(frames, 28, 0, 6, 1, 'x', 0, 0)
			packet, off := testQUICPacket(t, version, pn, 1, frames)
			before := bytes.Clone(packet)
			result, err := (&naaHnaatZCex{}).Sniff(packet)
			if err != nil || result.Type() != "quic" || result.Value() != "Mixed.Example.test" {
				t.Fatalf("version %x pn %d: %v %v", version, pn, result, err)
			}
			if packet[0] != 0xc0 || packet[off] != pn || !bytes.Equal(packet[off+1:off+1+len(frames)], frames) || bytes.Equal(packet, before) {
				t.Fatal("input was not decrypted in place")
			}
		}
	}
	typ := reflect.TypeFor[jO2kbfaNQJ]()
	if typ.Size() != 16 || typ.Field(0).Name != "dunzUGo" {
		t.Fatal(typ)
	}
}
func TestQUICOriginalRejectionsAndFrameBoundaries(t *testing.T) {
	s := &naaHnaatZCex{}
	frames := testCryptoFrame(0, syntheticHello("quic.example")[5:])
	for _, tc := range []struct {
		pn     byte
		length int
	}{{2, 1}, {0, 2}, {1, 3}, {0, 4}} {
		packet, _ := testQUICPacket(t, 1, tc.pn, tc.length, frames)
		if _, err := s.Sniff(packet); err != errorAt5098390 {
			t.Fatal(tc, err)
		}
	}
	packet, off := testQUICPacket(t, 1, 0, 1, frames)
	for i := 0; i < len(packet); i++ {
		truncated := append([]byte{}, packet[:i]...)
		truncated = truncated[:len(truncated):len(truncated)]
		if r, e := s.Sniff(truncated); e == nil || r != nil {
			t.Fatalf("truncated %d unexpectedly accepted", i)
		}
	}
	tampered := bytes.Clone(packet)
	tampered[len(tampered)-1] ^= 1
	if _, e := s.Sniff(tampered); e == nil || e == errorAt5098380 || e == errorAt5098390 {
		t.Fatal("AEAD error was lost", e)
	}
	for _, bad := range [][]byte{{2}, {3, 0, 0, 0, 0, 0}, {6, 0, 2, 1}, {6, 0, 0}, {28, 0, 0, 4, 1}} {
		// 头保护样本还需至少三字节明文；前置 PADDING 不改变末端截断。
		bad = append([]byte{0, 0, 0}, bad...)
		p, _ := testQUICPacket(t, 1, 0, 1, bad)
		if _, e := s.Sniff(p); e != io.ErrUnexpectedEOF {
			t.Fatalf("frame %x: %v", bad, e)
		}
	}
	unknown, _ := testQUICPacket(t, 1, 0, 1, []byte{0, 0, 0, 0x1d})
	if _, e := s.Sniff(unknown); e != errorAt5098390 {
		t.Fatal(e)
	}
	packet[1] = 1
	if _, e := s.Sniff(packet); e != errorAt5098380 {
		t.Fatal(e)
	}
	_ = off
}
func TestOriginalSnifferFactory(t *testing.T) {
	for _, name := range []MCe3PxlaMYM{"bt", "http", "tls", "quic"} {
		s, e := NHzYZHDSd(name)
		if e != nil || s.Type() != name {
			t.Fatal(s, e)
		}
	}
	if s, e := NHzYZHDSd("QUIC"); s != nil || e == nil || e.Error() != "sniffer type <QUIC> not exist" {
		t.Fatal(s, e)
	}
}
