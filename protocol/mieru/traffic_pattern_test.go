package mieru

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"math/bits"
	"strings"
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
)

// 用 protowire 独立构造 TrafficPattern 消息，字段号取自 mieru base.proto。
type pbField struct {
	num protowire.Number
	v   int64
	sub []pbField
	str string
}

func pbEncode(fields []pbField) []byte {
	var b []byte
	for _, f := range fields {
		switch {
		case f.sub != nil:
			b = protowire.AppendTag(b, f.num, protowire.BytesType)
			b = protowire.AppendBytes(b, pbEncode(f.sub))
		case f.str != "":
			b = protowire.AppendTag(b, f.num, protowire.BytesType)
			b = protowire.AppendString(b, f.str)
		default:
			b = protowire.AppendTag(b, f.num, protowire.VarintType)
			b = protowire.AppendVarint(b, uint64(f.v))
		}
	}
	return b
}

func pbPattern(fields ...pbField) string { return base64.StdEncoding.EncodeToString(pbEncode(fields)) }

func useTrafficPattern(t *testing.T, encoded string) *trafficPattern {
	t.Helper()
	if err := SetTrafficPattern(encoded); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { trafficPatternAt.Store(nil) })
	return currentTrafficPattern()
}

func TestTrafficPatternParseExplicit(t *testing.T) {
	encoded := pbPattern(
		pbField{num: 1, v: -7},
		pbField{num: 2, v: 1},
		pbField{num: 3, sub: []pbField{{num: 1, v: 1}, {num: 2, v: 30}}},
		pbField{num: 4, sub: []pbField{{num: 1, v: 3}, {num: 3, v: 2}, {num: 4, v: 9}, {num: 5, str: "00010203"}, {num: 5, str: "ff"}}},
		pbField{num: 5, sub: []pbField{{num: 1, v: 0}, {num: 2, v: 77}}},
		pbField{num: 6, sub: []pbField{{num: 1, v: 2}, {num: 2, v: 48}}},
		pbField{num: 99, v: 5}, // 未知字段忽略
	)
	p := useTrafficPattern(t, encoded)
	if *p.seed != -7 || !p.unlockAll || !*p.fragmentEnable || *p.fragmentSleep != 30 {
		t.Fatalf("%+v", p)
	}
	if *p.nonceType != nonceTypeFixed || *p.nonceMin != 2 || *p.nonceMax != 9 || len(p.nonceFixed) != 2 || !bytes.Equal(p.nonceFixed[0], []byte{0, 1, 2, 3}) {
		t.Fatalf("nonce %+v", p)
	}
	if p.maxMiddle() != 0 || p.maxEnd() != 77 {
		t.Fatal("padding", p.maxMiddle(), p.maxEnd())
	}
	if mode, rot := p.lowEntropy(); mode != 2 || rot != 48 {
		t.Fatal("low entropy", mode, rot)
	}
	// 显式设置的固定前缀写入 nonce 开头，后面的用户提示不受影响。
	nonce := bytes.Repeat([]byte{0xaa}, 24)
	p.applyNonce(nonce)
	if !bytes.HasPrefix(nonce, []byte{0, 1, 2, 3}) && !bytes.HasPrefix(nonce, []byte{0xff, 0xaa}) {
		t.Fatalf("fixed nonce %x", nonce)
	}
	if !bytes.Equal(nonce[12:], bytes.Repeat([]byte{0xaa}, 12)) {
		t.Fatalf("nonce tail changed %x", nonce)
	}
}

func TestTrafficPatternInvalid(t *testing.T) {
	t.Cleanup(func() { trafficPatternAt.Store(nil) })
	for _, tc := range []struct{ encoded, err string }{
		{"!!", "base64"},
		{pbPattern(pbField{num: 3, sub: []pbField{{num: 2, v: 101}}}), "maxSleepMs"},
		{pbPattern(pbField{num: 4, sub: []pbField{{num: 3, v: 9}, {num: 4, v: 3}}}), "greater than maxLen"},
		{pbPattern(pbField{num: 4, sub: []pbField{{num: 5, str: "zz"}}}), "hex"},
		{pbPattern(pbField{num: 5, sub: []pbField{{num: 2, v: 256}}}), "maxEndPaddingLen"},
		{pbPattern(pbField{num: 6, sub: []pbField{{num: 1, v: 5}}}), "mode"},
		{pbPattern(pbField{num: 6, sub: []pbField{{num: 2, v: 17}}}), "maskRotation"},
		{base64.StdEncoding.EncodeToString([]byte{0x0a, 0x05}), "parse"},
	} {
		if err := SetTrafficPattern(tc.encoded); err == nil || !strings.Contains(err.Error(), tc.err) {
			t.Fatalf("%q: %v", tc.encoded, err)
		}
	}
	if err := SetTrafficPattern("  "); err != nil || currentTrafficPattern() != nil {
		t.Fatal("empty pattern must disable", err)
	}
}

func TestTrafficPatternImplicit(t *testing.T) {
	// unlockAll=false 时 mieru 只生成保守的隐式值。
	for _, seed := range []int64{1, 12345, -3} {
		p := useTrafficPattern(t, pbPattern(pbField{num: 1, v: seed}))
		if *p.fragmentEnable || *p.fragmentSleep != 0 {
			t.Fatal("fragment", seed)
		}
		if typ := *p.nonceType; typ != nonceTypePrintable && typ != nonceTypePrintableSub {
			t.Fatal("nonce type", seed, typ)
		}
		if *p.nonceMin < 6 || *p.nonceMin > 12 || *p.nonceMax < *p.nonceMin || *p.nonceMax > 12 {
			t.Fatal("nonce len", seed, *p.nonceMin, *p.nonceMax)
		}
		if p.maxMiddle() < 0 || p.maxMiddle() > 127 || p.maxEnd() != 255 {
			t.Fatal("padding", seed, p.maxMiddle(), p.maxEnd())
		}
		if mode, rot := p.lowEntropy(); mode != 0 || !validLowEntropyRotation(int32(rot)) {
			t.Fatal("low entropy", seed, mode, rot)
		}
		// 同一 seed 的结果稳定。
		q, _ := parseTrafficPattern(pbPattern(pbField{num: 1, v: seed}))
		q.generateImplicit()
		if *q.nonceType != *p.nonceType || *q.nonceMin != *p.nonceMin || q.maxMiddle() != p.maxMiddle() {
			t.Fatal("unstable", seed)
		}
	}
	// 与 mieru rng.FixedInt 的定义逐步核对一项。
	p := useTrafficPattern(t, pbPattern(pbField{num: 1, v: 1}))
	if want := int32(fixedInt(2, "1:nonce.type") + 1); *p.nonceType != want {
		t.Fatal("nonce type formula", *p.nonceType, want)
	}
	// 显式字段不被隐式生成覆盖。
	p = useTrafficPattern(t, pbPattern(pbField{num: 2, v: 1}, pbField{num: 5, sub: []pbField{{num: 1, v: 9}}}))
	if p.maxMiddle() != 9 || p.fragmentSleep == nil || *p.fragmentSleep < 1 || *p.fragmentSleep > 100 {
		t.Fatal("explicit kept", p.maxMiddle(), *p.fragmentSleep)
	}
}

func TestTrafficPatternNonceRewrite(t *testing.T) {
	p := useTrafficPattern(t, pbPattern(pbField{num: 4, sub: []pbField{{num: 1, v: nonceTypePrintableSub}, {num: 3, v: 12}, {num: 4, v: 12}}}))
	for i := 0; i < 50; i++ {
		w, err := vfhZtLu(new(frameMemoryConn), testFrameAEAD(t), "user")
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range w.g9zKwmF[:12] {
			if !strings.ContainsRune(common64Set, rune(c)) {
				t.Fatalf("nonce %x", w.g9zKwmF)
			}
		}
		if !pyUlp642nl("user", w.g9zKwmF) {
			t.Fatal("user hint lost")
		}
	}
	p = useTrafficPattern(t, pbPattern(pbField{num: 4, sub: []pbField{{num: 1, v: nonceTypePrintable}, {num: 3, v: 4}, {num: 4, v: 8}}}))
	_ = p
	for i := 0; i < 50; i++ {
		w, _ := vfhZtLu(new(frameMemoryConn), testFrameAEAD(t), "")
		for _, c := range w.g9zKwmF[:4] {
			if c < 32 || c > 126 {
				t.Fatalf("printable nonce %x", w.g9zKwmF)
			}
		}
	}
}

func TestLowEntropyCodecVector(t *testing.T) {
	// 手算：mode 32、half-mask 0x0000ffff、不旋转。源 0x01020304 的低 16 位放入
	// 掩码第 0-15 位，高 16 位放入第 32-47 位。
	enc := encodeLowEntropy([]byte{1, 2, 3, 4}, 1, 0x0000ffff, 0)
	want := uint64(0x0000010200000304)
	if lowEntropyPaddingBit == 1 {
		want |= ^uint64(0x0000ffff0000ffff)
	}
	if len(enc) != 8 || binary.BigEndian.Uint64(enc) != want {
		t.Fatalf("%x", enc)
	}
	src := []byte("hello low entropy world, 0123456789!")
	for mode := uint8(1); mode <= 4; mode++ {
		_, ones := lowEntropyParams(mode)
		for _, rot := range []uint8{0, 1, 7, 15, 16, 48, 240} {
			for trial := 0; trial < 5; trial++ {
				mask := newLowEntropyHalfMask(mode)
				if bits.OnesCount32(mask) != ones {
					t.Fatal("mask bits")
				}
				for n := 0; n <= len(src); n++ {
					enc := encodeLowEntropy(src[:n], mode, mask, rot)
					got, err := decodeLowEntropy(enc, n, mode, mask, rot)
					if err != nil || !bytes.Equal(got, src[:n]) {
						t.Fatalf("mode %d rot %d n %d: %v", mode, rot, n, err)
					}
				}
			}
		}
	}
	// 填充位不一致必须被拒绝。
	enc = encodeLowEntropy(src, 1, 0x0000ffff, 0)
	enc[9] ^= 0x80
	if _, err := decodeLowEntropy(enc, len(src), 1, 0x0000ffff, 0); err == nil {
		t.Fatal("tampered padding accepted")
	}
	if _, err := decodeLowEntropy(enc[:8], len(src), 1, 0x0000ffff, 0); err == nil {
		t.Fatal("bad length accepted")
	}
}

func TestLowEntropyMetadata(t *testing.T) {
	now := time.Unix(1700000000, 0)
	m := &IqB1foHWg{Rf6N7kY6GVCs: 10, WyK3yKmJKe1V: uint32(now.Unix() / 60), HNaaDh: 5, Tt1OfazA: 6, N01v7C: 7, UlMT6MeaV2: 8, B8QsA0N2yo: 1, FO7MEJ: 2,
		LowEntropyMode: 1, LowEntropyMask: 0x0000ffff, ExtractedPayloadLen: 9, HPUJVQ0S: 24, LowEntropyRotation: 3}
	raw := m.Marshal()
	if raw[1] != 1 || binary.BigEndian.Uint32(raw[25:]) != 0x0000ffff || binary.BigEndian.Uint16(raw[29:]) != 9 || raw[31] != 3 {
		t.Fatalf("%x", raw)
	}
	got, err := decodeMetadataAt(raw[:], now)
	if err != nil || *got != *m {
		t.Fatalf("%+v %v", got, err)
	}
	for _, mutate := range []func(*IqB1foHWg){
		func(m *IqB1foHWg) { m.LowEntropyMask = 0xffffffff },
		func(m *IqB1foHWg) { m.HPUJVQ0S = 16 },
		func(m *IqB1foHWg) { m.LowEntropyMode = 0 },
		func(m *IqB1foHWg) { m.LowEntropyRotation = 17 },
	} {
		bad := *m
		mutate(&bad)
		raw := bad.Marshal()
		if _, err := decodeMetadataAt(raw[:], now); err == nil {
			t.Fatalf("accepted %+v", bad)
		}
	}
	// 非低熵帧不写入扩展字节。
	plain := *m
	plain.Rf6N7kY6GVCs = 7
	if raw := plain.Marshal(); raw[1] != 0 || !bytes.Equal(raw[25:], make([]byte, 7)) {
		t.Fatalf("%x", raw)
	}
}

// writeAndRead 用写端产生数据，再用读端按握手后的 nonce 状态解析。
func writeAndRead(t *testing.T, w *qhFcsZC, conn *frameMemoryConn, m *IqB1foHWg, payload []byte) (*IqB1foHWg, []byte) {
	t.Helper()
	nonce := append([]byte(nil), w.g9zKwmF...)
	if err := w.WriteDataSegment(m, payload); err != nil {
		t.Fatal(err)
	}
	wire := conn.Bytes()[24:]
	plain, err := w.yn5oQQkf8.Open(nil, nonce, wire[:48], nil)
	if err != nil {
		t.Fatal(err)
	}
	header, err := Ox0UfI5kIfcR(plain)
	if err != nil {
		t.Fatal(err)
	}
	conn.Next(24 + 48)
	r := &rZDq_kinV4G{aMKXGpHBfg: conn, n49jm3JWYo: w.yn5oQQkf8, xUtUMCF3grV: nonce}
	seg, err := r.jQ8yjlkt(header)
	if err != nil {
		t.Fatal(err)
	}
	if conn.Len() != 0 {
		t.Fatal("trailing bytes", conn.Len())
	}
	return header, seg.mzrzt1uBj
}

func TestLowEntropyFrameRoundTrip(t *testing.T) {
	useTrafficPattern(t, pbPattern(
		pbField{num: 5, sub: []pbField{{num: 1, v: 0}, {num: 2, v: 0}}},
		pbField{num: 6, sub: []pbField{{num: 1, v: 2}, {num: 2, v: 3}}},
	))
	payload := bytes.Repeat([]byte("mieru low entropy "), 50)

	// 客户端还没用过低熵时，服务端照常发送普通数据帧。
	conn := new(frameMemoryConn)
	w, _ := vfhZtLu(conn, testFrameAEAD(t), "u")
	header, got := writeAndRead(t, w, conn, &IqB1foHWg{Rf6N7kY6GVCs: 7, HNaaDh: 1}, payload)
	if header.Rf6N7kY6GVCs != 7 || header.B8QsA0N2yo != 0 || header.FO7MEJ != 0 || !bytes.Equal(got, payload) {
		t.Fatalf("plain %+v", header)
	}

	conn = new(frameMemoryConn)
	w, _ = vfhZtLu(conn, testFrameAEAD(t), "u")
	w.clientLowEntropy.Store(true)
	header, got = writeAndRead(t, w, conn, &IqB1foHWg{Rf6N7kY6GVCs: 7, HNaaDh: 1}, payload)
	if header.Rf6N7kY6GVCs != 11 || header.LowEntropyMode != 2 || header.LowEntropyRotation != 3 || int(header.ExtractedPayloadLen) != len(payload) || !bytes.Equal(got, payload) {
		t.Fatalf("low entropy %+v", header)
	}
	if want, _ := lowEntropyEncodedLen(len(payload), 2); int(header.HPUJVQ0S) != want {
		t.Fatal("wire length", header.HPUJVQ0S, want)
	}
	// 低熵模式下单帧明文上限保证编码长度不溢出 uint16。
	if n := w.pattern.maxDataFragment(); n > maxPDU {
		t.Fatal(n)
	} else if enc, err := lowEntropyEncodedLen(n, 2); err != nil || enc > 65535 {
		t.Fatal(enc, err)
	}
}

type countingConn struct {
	frameMemoryConn
	writes int
}

func (c *countingConn) Write(p []byte) (int, error) { c.writes++; return c.frameMemoryConn.Write(p) }

func TestTrafficPatternSessionFragmentAndPadding(t *testing.T) {
	useTrafficPattern(t, pbPattern(
		pbField{num: 3, sub: []pbField{{num: 1, v: 1}, {num: 2, v: 0}}},
		pbField{num: 5, sub: []pbField{{num: 2, v: 0}}},
	))
	conn := new(countingConn)
	w, _ := vfhZtLu(conn, testFrameAEAD(t), "u")
	m := &IqB1foHWg{Rf6N7kY6GVCs: 3, HNaaDh: 1}
	if err := w.WriteSessionSegment(m, bytes.Repeat([]byte{1}, 200)); err != nil {
		t.Fatal(err)
	}
	if m.FO7MEJ != 0 || conn.Len() != 24+48+200+16 {
		t.Fatal("session padding", m.FO7MEJ, conn.Len())
	}
	if conn.writes < 2 {
		t.Fatal("not fragmented", conn.writes)
	}
}
