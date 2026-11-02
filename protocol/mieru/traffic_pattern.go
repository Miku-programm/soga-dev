package mieru

// 服务端 traffic pattern，按 mieru v3 apis/trafficpattern 与 pkg/protocol 移植。
// 面板下发的是 `mieru export traffic-pattern` 导出的 Base64(protobuf TrafficPattern)。
// 只影响服务端发出的数据：nonce 前缀、填充长度、会话帧 TCP 分片和低熵编码；
// 客户端发来的低熵数据帧无论服务端是否配置都能解码。

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/bits"
	mathrand "math/rand"
	"net"
	"strings"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
)

const (
	nonceTypeRandom        = 0
	nonceTypePrintable     = 1
	nonceTypePrintableSub  = 2
	nonceTypeFixed         = 3
	lowEntropyModeOff      = 0
	lowEntropyChunkLen     = 8
	lowEntropyRotationKeys = 31 // len(LowEntropyMaskRotation_name)
	maxPDU                 = 32 * 1024
)

// mieru common.Common64Set。
const common64Set = "WKYu29q87rl6ifoVgGGyiAN7JwYx8sgTSAebhHsySEeTfZKzWDOzhX6D1xneQlx0"

// trafficPattern 是生效后的配置；proto 中的 optional 字段用指针记录是否显式设置。
type trafficPattern struct {
	seed      *int32
	unlockAll bool

	fragmentEnable *bool
	fragmentSleep  *int32

	nonceType  *int32
	nonceMin   *int32
	nonceMax   *int32
	nonceUDP   *bool
	nonceFixed [][]byte

	maxMiddlePadding *int32
	maxEndPadding    *int32

	lowEntropyMode     *int32
	lowEntropyRotation *int32
}

var trafficPatternAt atomic.Pointer[trafficPattern]

// SetTrafficPattern 由 mieru inbound 启动时按节点配置设置；空字符串表示不启用，保持原有行为。
func SetTrafficPattern(encoded string) error {
	encoded = strings.TrimSpace(encoded)
	if encoded == "" {
		trafficPatternAt.Store(nil)
		return nil
	}
	p, err := parseTrafficPattern(encoded)
	if err != nil {
		return err
	}
	p.generateImplicit()
	trafficPatternAt.Store(p)
	return nil
}

func currentTrafficPattern() *trafficPattern { return trafficPatternAt.Load() }

func parseTrafficPattern(encoded string) (*trafficPattern, error) {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("decode traffic pattern base64 failed: %w", err)
	}
	p := new(trafficPattern)
	err = walkProto(b, func(num protowire.Number, v uint64, sub []byte) error {
		switch num {
		case 1:
			p.seed = i32p(v)
		case 2:
			p.unlockAll = v != 0
		case 3:
			return walkProto(sub, func(num protowire.Number, v uint64, _ []byte) error {
				switch num {
				case 1:
					p.fragmentEnable = boolp(v)
				case 2:
					p.fragmentSleep = i32p(v)
				}
				return nil
			})
		case 4:
			return walkProto(sub, func(num protowire.Number, v uint64, s []byte) error {
				switch num {
				case 1:
					p.nonceType = i32p(v)
				case 2:
					p.nonceUDP = boolp(v)
				case 3:
					p.nonceMin = i32p(v)
				case 4:
					p.nonceMax = i32p(v)
				case 5:
					prefix, err := hex.DecodeString(string(s))
					if err != nil {
						return fmt.Errorf("nonce customHexStrings %q is not a valid hex string: %w", s, err)
					}
					if len(prefix) > 12 {
						return fmt.Errorf("nonce customHexStrings %q exceeds 12 bytes", s)
					}
					p.nonceFixed = append(p.nonceFixed, prefix)
				}
				return nil
			})
		case 5:
			return walkProto(sub, func(num protowire.Number, v uint64, _ []byte) error {
				switch num {
				case 1:
					p.maxMiddlePadding = i32p(v)
				case 2:
					p.maxEndPadding = i32p(v)
				}
				return nil
			})
		case 6:
			return walkProto(sub, func(num protowire.Number, v uint64, _ []byte) error {
				switch num {
				case 1:
					p.lowEntropyMode = i32p(v)
				case 2:
					p.lowEntropyRotation = i32p(v)
				}
				return nil
			})
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse traffic pattern failed: %w", err)
	}
	if err := p.validate(); err != nil {
		return nil, fmt.Errorf("traffic pattern is invalid: %w", err)
	}
	return p, nil
}

// walkProto 遍历一层 protobuf 字段；varint 传 v，length-delimited 传 sub。
func walkProto(b []byte, fn func(protowire.Number, uint64, []byte) error) error {
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return protowire.ParseError(n)
		}
		b = b[n:]
		var v uint64
		var sub []byte
		switch typ {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(b)
		case protowire.BytesType:
			sub, n = protowire.ConsumeBytes(b)
		default:
			n = protowire.ConsumeFieldValue(num, typ, b)
		}
		if n < 0 {
			return protowire.ParseError(n)
		}
		b = b[n:]
		if typ == protowire.VarintType || typ == protowire.BytesType {
			if err := fn(num, v, sub); err != nil {
				return err
			}
		}
	}
	return nil
}

func i32p(v uint64) *int32 { x := int32(v); return &x }
func boolp(v uint64) *bool { x := v != 0; return &x }
func getI32(p *int32) int32 {
	if p == nil {
		return 0
	}
	return *p
}

func (p *trafficPattern) validate() error {
	if v := getI32(p.fragmentSleep); v < 0 || v > 100 {
		return fmt.Errorf("tcpFragment maxSleepMs %d out of range [0, 100]", v)
	}
	if v := getI32(p.nonceMin); v < 0 || v > 12 {
		return fmt.Errorf("nonce minLen %d out of range [0, 12]", v)
	}
	if v := getI32(p.nonceMax); v < 0 || v > 12 {
		return fmt.Errorf("nonce maxLen %d out of range [0, 12]", v)
	}
	if p.nonceMin != nil && p.nonceMax != nil && *p.nonceMin > *p.nonceMax {
		return fmt.Errorf("nonce minLen %d is greater than maxLen %d", *p.nonceMin, *p.nonceMax)
	}
	if v := getI32(p.maxMiddlePadding); v < 0 || v > 255 {
		return fmt.Errorf("padding maxMiddlePaddingLen %d out of range [0, 255]", v)
	}
	if v := getI32(p.maxEndPadding); v < 0 || v > 255 {
		return fmt.Errorf("padding maxEndPaddingLen %d out of range [0, 255]", v)
	}
	if v := getI32(p.lowEntropyMode); v < 0 || v > 4 {
		return fmt.Errorf("lowEntropy mode %d is invalid", v)
	}
	if p.lowEntropyRotation != nil && !validLowEntropyRotation(*p.lowEntropyRotation) {
		return fmt.Errorf("lowEntropy maskRotation %d is invalid", *p.lowEntropyRotation)
	}
	return nil
}

// fixedInt 对应 mieru rng.FixedInt。
func fixedInt(n int, hint string) int {
	if n <= 0 {
		return 0
	}
	b := sha256.Sum256([]byte(hint))
	b[0] &= 0x7f
	return int(binary.BigEndian.Uint32(b[:4])) % n
}

// generateImplicit 对应 trafficpattern.Config.generateImplicitTrafficPattern，补齐未显式设置的字段。
func (p *trafficPattern) generateImplicit() {
	seed := int(getI32(p.seed))
	if p.seed == nil {
		seed = fsAE63SXO(math.MaxInt32)
	}
	key := func(name string) string { return fmt.Sprintf("%d:%s", seed, name) }
	set := func(v int) *int32 { x := int32(v); return &x }

	if p.fragmentEnable == nil {
		v := p.unlockAll && fixedInt(2, key("tcpFragment.enable")) == 1
		p.fragmentEnable = &v
	}
	if p.fragmentSleep == nil {
		if p.unlockAll {
			p.fragmentSleep = set(fixedInt(100, key("tcpFragment.maxSleepMs")) + 1)
		} else {
			p.fragmentSleep = set(0)
		}
	}

	if p.nonceType == nil {
		if p.unlockAll {
			p.nonceType = set(fixedInt(3, key("nonce.type")))
		} else {
			p.nonceType = set(fixedInt(2, key("nonce.type")) + 1)
		}
	}
	if p.nonceUDP == nil {
		v := fixedInt(2, key("nonce.applyToAllUDPPacket")) == 1
		p.nonceUDP = &v
	}
	if p.nonceMin == nil {
		if p.unlockAll {
			p.nonceMin = set(fixedInt(13, key("nonce.minLen")))
		} else {
			p.nonceMin = set(fixedInt(7, key("nonce.minLen")) + 6)
		}
	}
	if p.nonceMax == nil {
		minLen := int(*p.nonceMin)
		p.nonceMax = set(minLen + fixedInt(13-minLen, key("nonce.maxLen")))
	}

	if p.maxMiddlePadding == nil {
		p.maxMiddlePadding = set(max(fixedInt(256, key("padding.maxMiddlePaddingLen"))-128, 0))
	}
	if p.maxEndPadding == nil {
		if p.unlockAll {
			p.maxEndPadding = set(fixedInt(256, key("padding.maxEndPaddingLen")))
		} else {
			p.maxEndPadding = set(255)
		}
	}

	if p.lowEntropyMode == nil {
		if p.unlockAll {
			p.lowEntropyMode = set(fixedInt(5, key("lowEntropy.mode")))
		} else {
			p.lowEntropyMode = set(lowEntropyModeOff)
		}
	}
	if p.lowEntropyRotation == nil {
		idx := fixedInt(lowEntropyRotationKeys, key("lowEntropy.maskRotation"))
		if idx > 15 {
			idx = (idx - 15) * 16
		}
		p.lowEntropyRotation = set(idx)
	}
}

// applyNonce 对应 mieru cipher.newNonceTo，在用户提示写入之前改写 nonce 前缀。
func (p *trafficPattern) applyNonce(nonce []byte) {
	if p == nil {
		return
	}
	switch getI32(p.nonceType) {
	case nonceTypePrintable:
		goAOlbfD50OV(nonce, 0, p.nonceRewriteLen(len(nonce)))
	case nonceTypePrintableSub:
		for i := 0; i < p.nonceRewriteLen(len(nonce)); i++ {
			nonce[i] = common64Set[nonce[i]&0x3f]
		}
	case nonceTypeFixed:
		if len(p.nonceFixed) > 0 {
			copy(nonce, p.nonceFixed[mathrand.Intn(len(p.nonceFixed))])
		}
	}
}

func (p *trafficPattern) nonceRewriteLen(size int) int {
	minLen, maxLen := int(getI32(p.nonceMin)), min(int(getI32(p.nonceMax)), size)
	if minLen >= maxLen {
		return maxLen
	}
	return minLen + mathrand.Intn(maxLen-minLen+1)
}

// maxMiddle / maxEnd 返回填充上限；没有 pattern 时返回 -1 表示保持原逻辑。
func (p *trafficPattern) maxMiddle() int {
	if p == nil {
		return -1
	}
	return int(getI32(p.maxMiddlePadding))
}
func (p *trafficPattern) maxEnd() int {
	if p == nil {
		return -1
	}
	return int(getI32(p.maxEndPadding))
}

// writeFragmented 对应 StreamUnderlay.writeWithPossibleFragment，只用于会话帧。
func (p *trafficPattern) writeFragmented(conn net.Conn, data []byte) error {
	if p == nil || p.fragmentEnable == nil || !*p.fragmentEnable {
		_, err := conn.Write(data)
		return err
	}
	minLen := int(math.Sqrt(float64(len(data)))) + 1
	maxLen := max(minLen, len(data)/2)
	for remaining := data; len(remaining) > 0; {
		n := min(mathrand.Intn(maxLen-minLen+1)+minLen, len(remaining))
		if _, err := conn.Write(remaining[:n]); err != nil {
			return err
		}
		if sleep := getI32(p.fragmentSleep); sleep > 0 {
			time.Sleep(time.Duration(mathrand.Intn(int(sleep)+1)) * time.Millisecond)
		}
		remaining = remaining[n:]
	}
	return nil
}

// lowEntropy 返回服务端发送低熵数据的模式与旋转方式；mode 为 0 表示关闭。
func (p *trafficPattern) lowEntropy() (mode, rotation uint8) {
	if p == nil {
		return 0, 0
	}
	return uint8(getI32(p.lowEntropyMode)), uint8(getI32(p.lowEntropyRotation))
}

// maxDataFragment 返回一个数据帧最多携带的明文长度，低熵编码后的长度不能超过 uint16。
func (p *trafficPattern) maxDataFragment() int {
	mode, _ := p.lowEntropy()
	if mode == lowEntropyModeOff {
		return maxPDU
	}
	src, _ := lowEntropyParams(mode)
	return min(maxPDU, math.MaxUint16/lowEntropyChunkLen*src)
}

// ---- 低熵编码，对应 mieru pkg/protocol/low_entropy.go ----

func lowEntropyParams(mode uint8) (sourceBytes, halfMaskOnes int) {
	switch mode {
	case 1:
		return 4, 16
	case 2:
		return 5, 20
	case 3:
		return 6, 24
	case 4:
		return 7, 28
	}
	return 0, 0
}

func validLowEntropyRotation(r int32) bool {
	return (r >= 0 && r <= 15) || (r >= 16 && r <= 240 && r%16 == 0)
}

func lowEntropyEncodedLen(extracted int, mode uint8) (int, error) {
	src, _ := lowEntropyParams(mode)
	if src == 0 {
		return 0, fmt.Errorf("invalid low entropy mode %d", mode)
	}
	n := (extracted + src - 1) / src * lowEntropyChunkLen
	if n > math.MaxUint16 {
		return 0, fmt.Errorf("encoded payload length for %d bytes exceeds %d", extracted, math.MaxUint16)
	}
	return n, nil
}

// validateLowEntropyMetadata 对应 validateLowEntropyDataAckMetadata。
func validateLowEntropyMetadata(m *IqB1foHWg) error {
	if int(m.ExtractedPayloadLen) > maxPDU {
		return fmt.Errorf("invalid extracted payload length %d", m.ExtractedPayloadLen)
	}
	if m.HPUJVQ0S%lowEntropyChunkLen != 0 {
		return fmt.Errorf("invalid low entropy payload length %d", m.HPUJVQ0S)
	}
	_, ones := lowEntropyParams(m.LowEntropyMode)
	if ones == 0 {
		return fmt.Errorf("invalid low entropy mode %d", m.LowEntropyMode)
	}
	if bits.OnesCount32(m.LowEntropyMask) != ones {
		return fmt.Errorf("low entropy mask has %d one-bits, want %d", bits.OnesCount32(m.LowEntropyMask), ones)
	}
	if !validLowEntropyRotation(int32(m.LowEntropyRotation)) {
		return fmt.Errorf("invalid low entropy mask rotation %d", m.LowEntropyRotation)
	}
	if m.ExtractedPayloadLen == 0 {
		if m.HPUJVQ0S != 0 {
			return fmt.Errorf("low entropy payload length is %d, want 0", m.HPUJVQ0S)
		}
		return nil
	}
	want, err := lowEntropyEncodedLen(int(m.ExtractedPayloadLen), m.LowEntropyMode)
	if err != nil {
		return err
	}
	if int(m.HPUJVQ0S) != want {
		return fmt.Errorf("low entropy payload length is %d, want %d", m.HPUJVQ0S, want)
	}
	return nil
}

func rotateLowEntropyMask(mask uint64, rotation uint8, chunk int) uint64 {
	if rotation == 0 || chunk == 0 {
		return mask
	}
	if rotation <= 15 {
		return bits.RotateLeft64(mask, -((chunk % 64) * int(rotation)))
	}
	return bits.RotateLeft64(mask, (chunk%64)*int(rotation/16))
}

func lowBits(n int) uint64 {
	if n >= 64 {
		return math.MaxUint64
	}
	return 1<<n - 1
}

func pdep(x, mask uint64) uint64 {
	var out uint64
	for bit := uint64(1); mask != 0; bit <<= 1 {
		m := mask & -mask
		if x&bit != 0 {
			out |= m
		}
		mask &= mask - 1
	}
	return out
}

func pext(x, mask uint64) uint64 {
	var out uint64
	for bit := uint64(1); mask != 0; bit <<= 1 {
		m := mask & -mask
		if x&m != 0 {
			out |= bit
		}
		mask &= mask - 1
	}
	return out
}

// lowEntropyPaddingBit 对应 mieru 的 rng.FixedIntVH(2)，每台主机固定。
var lowEntropyPaddingBit = uint8(fsAE63SXO(2))

func encodeLowEntropy(src []byte, mode uint8, halfMask uint32, rotation uint8) []byte {
	srcPerChunk, _ := lowEntropyParams(mode)
	n, _ := lowEntropyEncodedLen(len(src), mode)
	out := make([]byte, n)
	mask := uint64(halfMask)<<32 | uint64(halfMask)
	for chunk, off := 0, 0; off < len(src); chunk, off = chunk+1, off+srcPerChunk {
		size := min(srcPerChunk, len(src)-off)
		var scratch [lowEntropyChunkLen]byte
		copy(scratch[lowEntropyChunkLen-size:], src[off:off+size])
		chunkMask := rotateLowEntropyMask(mask, rotation, chunk)
		v := pdep(binary.BigEndian.Uint64(scratch[:]), chunkMask)
		if lowEntropyPaddingBit == 1 {
			v |= ^pdep(lowBits(size*8), chunkMask)
		}
		binary.BigEndian.PutUint64(out[chunk*lowEntropyChunkLen:], v)
	}
	return out
}

func decodeLowEntropy(encoded []byte, extracted int, mode uint8, halfMask uint32, rotation uint8) ([]byte, error) {
	srcPerChunk, _ := lowEntropyParams(mode)
	if want, err := lowEntropyEncodedLen(extracted, mode); err != nil || len(encoded) != want {
		return nil, fmt.Errorf("encoded payload length is %d, want %d", len(encoded), want)
	}
	out := make([]byte, extracted)
	mask := uint64(halfMask)<<32 | uint64(halfMask)
	var paddingBit uint8
	for chunk, off := 0, 0; off < extracted; chunk, off = chunk+1, off+srcPerChunk {
		size := min(srcPerChunk, extracted-off)
		v := binary.BigEndian.Uint64(encoded[chunk*lowEntropyChunkLen:])
		chunkMask := rotateLowEntropyMask(mask, rotation, chunk)
		paddingMask := ^pdep(lowBits(size*8), chunkMask)
		padding := v & paddingMask
		if chunk == 0 {
			switch padding {
			case 0:
				paddingBit = 0
			case paddingMask:
				paddingBit = 1
			default:
				return nil, fmt.Errorf("low entropy chunk %d has mixed padding bits", chunk)
			}
		} else if (paddingBit == 0 && padding != 0) || (paddingBit == 1 && padding != paddingMask) {
			return nil, fmt.Errorf("low entropy chunk %d has non-uniform padding bits", chunk)
		}
		var scratch [lowEntropyChunkLen]byte
		binary.BigEndian.PutUint64(scratch[:], pext(v, chunkMask))
		copy(out[off:off+size], scratch[lowEntropyChunkLen-size:])
	}
	return out, nil
}

// newLowEntropyHalfMask 对应 rng.Uint32WithBits：随机选出 ones 个置位。
func newLowEntropyHalfMask(mode uint8) uint32 {
	_, ones := lowEntropyParams(mode)
	var positions [32]uint32
	for i := range positions {
		positions[i] = uint32(i)
	}
	var out uint32
	for i := 0; i < ones; i++ {
		j := i + mathrand.Intn(len(positions)-i)
		positions[i], positions[j] = positions[j], positions[i]
		out |= 1 << positions[i]
	}
	return out
}
