// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package low0s_10x

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"math"
	mathRand "math/rand"
	"net"
	"strconv"
	"strings"

	"recovered.local/soga/H8_In_Vuy"
	"recovered.local/soga/mzi1d3gDO7u"
	"recovered.local/soga/x6HWbunbF"
)

type (
	jiISoa4JwDI func(key, data []byte) []byte
	z92zGaN     func([]byte) []byte
)

// init: 0x1b338e0.
func init() {
	register("auth_aes128_sha1", fGRPkFSzCrQ, 9)
}

type oNvVPHch3WB struct {
	bvFg1gLbd  string
	efwls6Jy   jiISoa4JwDI
	qaEqxSog1a z92zGaN
}

type sdA1Al25W struct {
	*PZfVxx4UGRb
	*t1vL_y55t
	*oNvVPHch3WB
	*bhRsfpsWw
	fq3rseqxscn []byte
	hv9KPm      bool
	tPJl5EtsP   bool
	kGLv0DUCiIh uint32
	pgem_p      uint32
}

// fGRPkFSzCrQ: 0x1b35260.
func fGRPkFSzCrQ(b *PZfVxx4UGRb) Shdj3DH1 {
	a := &sdA1Al25W{
		PZfVxx4UGRb: b,
		t1vL_y55t:   &t1vL_y55t{},
		oNvVPHch3WB: &oNvVPHch3WB{bvFg1gLbd: "auth_aes128_sha1", efwls6Jy: H8_In_Vuy.B4_4_BnW, qaEqxSog1a: H8_In_Vuy.YhCPGpz},
		bhRsfpsWw:   &bhRsfpsWw{},
	}
	a.eatUbiR0d()
	return a
}

// (*sdA1Al25W).eatUbiR0d: 0x1b35420.
func (a *sdA1Al25W) eatUbiR0d() {
	params := strings.Split(a.BOJR3xVshAb, ":")
	if len(params) > 1 {
		if qC81k5Mk, err := strconv.ParseUint(params[0], 10, 32); err == nil {
			binary.LittleEndian.PutUint32(a.qC81k5Mk[:], uint32(qC81k5Mk))
			a.oYNnYN0vN = a.qaEqxSog1a([]byte(params[1]))
		} else {
			x6HWbunbF.Dl2Gj6("Wrong protocol-param for %s, only digits are expected before ':'", a.bvFg1gLbd)
		}
	}
	if len(a.oYNnYN0vN) == 0 {
		a.oYNnYN0vN = a.Z2UHwhi2h4
		rand.Read(a.qC81k5Mk[:])
	}
}

// (*sdA1Al25W).StreamConn: 0x1b35600.
func (a *sdA1Al25W) StreamConn(c net.Conn, fq3rseqxscn []byte) net.Conn {
	p := &sdA1Al25W{
		PZfVxx4UGRb: a.PZfVxx4UGRb,
		t1vL_y55t:   a.uU1wpTCSA(),
		oNvVPHch3WB: a.oNvVPHch3WB,
		bhRsfpsWw:   a.bhRsfpsWw,
		kGLv0DUCiIh: 1,
		pgem_p:      1,
	}
	p.fq3rseqxscn = fq3rseqxscn
	return &A6KE0r{Conn: c, Shdj3DH1: p}
}

// (*sdA1Al25W).PacketConn: 0x1b357c0.
func (a *sdA1Al25W) PacketConn(c net.PacketConn) net.PacketConn {
	p := &sdA1Al25W{
		PZfVxx4UGRb: a.PZfVxx4UGRb,
		oNvVPHch3WB: a.oNvVPHch3WB,
		bhRsfpsWw:   a.bhRsfpsWw,
	}
	return &S_N0Ixh{WCApOca6IQW7: c, Shdj3DH1: p}
}

// (*sdA1Al25W).Decode: 0x1b35900.
func (a *sdA1Al25W) Decode(dst, src *bytes.Buffer) error {
	if a.tPJl5EtsP {
		dst.ReadFrom(src)
		return nil
	}
	for src.Len() > 4 {
		macKey := mzi1d3gDO7u.SharedPoolAt5096978.Get(len(a.oYNnYN0vN) + 4)
		defer mzi1d3gDO7u.SharedPoolAt5096978.Put(macKey)
		copy(macKey, a.oYNnYN0vN)
		binary.LittleEndian.PutUint32(macKey[len(a.oYNnYN0vN):], a.pgem_p)
		if !bytes.Equal(a.efwls6Jy(macKey, src.Bytes()[:2])[:2], src.Bytes()[2:4]) {
			src.Reset()
			return errorAt50985C0
		}

		length := int(binary.LittleEndian.Uint16(src.Bytes()[:2]))
		if length >= 8192 || length < 7 {
			a.tPJl5EtsP = true
			src.Reset()
			return errorAt50985D0
		}
		if length > src.Len() {
			break
		}

		if !bytes.Equal(a.efwls6Jy(macKey, src.Bytes()[:length-4])[:4], src.Bytes()[length-4:length]) {
			a.tPJl5EtsP = true
			src.Reset()
			return errorAt50985E0
		}

		a.pgem_p++

		pos := int(src.Bytes()[4])
		if pos < 255 {
			pos += 4
		} else {
			pos = int(binary.LittleEndian.Uint16(src.Bytes()[5:7])) + 4
		}
		dst.Write(src.Bytes()[pos : length-4])
		src.Next(length)
	}
	return nil
}

// (*sdA1Al25W).Encode: 0x1b35f40.
func (a *sdA1Al25W) Encode(buf *bytes.Buffer, b []byte) error {
	fullDataLength := len(b)
	if !a.hv9KPm {
		dataLength := t4pamzUU(b)
		a.upMb05UedR(buf, b[:dataLength])
		b = b[dataLength:]
		a.hv9KPm = true
	}
	for len(b) > 8100 {
		a.pZRlqlGM1E(buf, b[:8100], fullDataLength)
		b = b[8100:]
	}
	if len(b) > 0 {
		a.pZRlqlGM1E(buf, b, fullDataLength)
	}
	return nil
}

// (*sdA1Al25W).DecodePacket: 0x1b360e0.
func (a *sdA1Al25W) DecodePacket(b []byte) ([]byte, error) {
	if len(b) < 4 {
		return nil, errorAt50985D0
	}
	if !bytes.Equal(a.efwls6Jy(a.Z2UHwhi2h4, b[:len(b)-4])[:4], b[len(b)-4:]) {
		return nil, errorAt50985E0
	}
	return b[:len(b)-4], nil
}

// (*sdA1Al25W).EncodePacket: 0x1b36200.
func (a *sdA1Al25W) EncodePacket(buf *bytes.Buffer, b []byte) error {
	buf.Write(b)
	buf.Write(a.qC81k5Mk[:])
	buf.Write(a.efwls6Jy(a.oYNnYN0vN, buf.Bytes())[:4])
	return nil
}

// (*sdA1Al25W).pZRlqlGM1E: 0x1b36320.
func (a *sdA1Al25W) pZRlqlGM1E(poolBuf *bytes.Buffer, data []byte, fullDataLength int) {
	dataLength := len(data)
	randDataLength := a.sErAgcxCJNy(dataLength, fullDataLength)
	/*
		2:	uint16 LittleEndian packedDataLength
		2:	hmac of packedDataLength
		3:	maxRandDataLengthPrefix (min:1)
		4:	hmac of packedData except the last 4 bytes
	*/
	packedDataLength := 2 + 2 + 3 + randDataLength + dataLength + 4
	if randDataLength < 128 {
		packedDataLength -= 2
	}

	macKey := mzi1d3gDO7u.SharedPoolAt5096978.Get(len(a.oYNnYN0vN) + 4)
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(macKey)
	copy(macKey, a.oYNnYN0vN)
	binary.LittleEndian.PutUint32(macKey[len(a.oYNnYN0vN):], a.kGLv0DUCiIh)
	a.kGLv0DUCiIh++

	binary.Write(poolBuf, binary.LittleEndian, uint16(packedDataLength))
	poolBuf.Write(a.efwls6Jy(macKey, poolBuf.Bytes()[poolBuf.Len()-2:])[:2])
	a.bQmucT1jtu(poolBuf, randDataLength)
	poolBuf.Write(data)
	poolBuf.Write(a.efwls6Jy(macKey, poolBuf.Bytes()[poolBuf.Len()-packedDataLength+4:])[:4])
}

// d6u7Ig5KSAc: 0x1b36740.
func d6u7Ig5KSAc(max int, d float64) int {
	base := mathRand.Float64()
	if d-0 > 1e-6 {
		a := 1 - d
		base = (math.Sqrt(a*a+4*d*base) - a) / (2 * d)
	}
	return int(base * float64(max))
}

// (*sdA1Al25W).sErAgcxCJNy: 0x1b36800.
func (a *sdA1Al25W) sErAgcxCJNy(dataLength, fullDataLength int) int {
	if fullDataLength >= 32*1024-a.UqLugVVWCKJ {
		return 0
	}
	// 1460: tcp_mss
	revLength := 1460 - dataLength - 9
	if revLength == 0 {
		return 0
	}
	if revLength < 0 {
		if revLength > -1460 {
			return d6u7Ig5KSAc(revLength+1460, -0.3)
		}
		return mathRand.Intn(32)
	}
	if dataLength > 900 {
		return mathRand.Intn(revLength)
	}
	return d6u7Ig5KSAc(revLength, -0.3)
}

// (*sdA1Al25W).upMb05UedR: 0x1b368e0.
func (a *sdA1Al25W) upMb05UedR(poolBuf *bytes.Buffer, data []byte) {
	if len(data) == 0 {
		return
	}
	dataLength := len(data)
	randDataLength := a.kLUMSZc(dataLength)
	/*
		7:	checkHead(1) and hmac of checkHead(6)
		4:	userID
		16:	encrypted data of authdata(12), uint16 BigEndian packedDataLength(2) and uint16 BigEndian randDataLength(2)
		4:	hmac of userID and encrypted data
		4:	hmac of packedAuthData except the last 4 bytes
	*/
	packedAuthDataLength := 7 + 4 + 16 + 4 + randDataLength + dataLength + 4

	macKey := mzi1d3gDO7u.SharedPoolAt5096978.Get(len(a.fq3rseqxscn) + len(a.Z2UHwhi2h4))
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(macKey)
	copy(macKey, a.fq3rseqxscn)
	copy(macKey[len(a.fq3rseqxscn):], a.Z2UHwhi2h4)

	poolBuf.WriteByte(byte(mathRand.Intn(256)))
	poolBuf.Write(a.efwls6Jy(macKey, poolBuf.Bytes())[:6])
	poolBuf.Write(a.qC81k5Mk[:])
	err := a.t1vL_y55t.vKsdajuSitet(poolBuf, a.oYNnYN0vN, [2]int{packedAuthDataLength, randDataLength}, a.bvFg1gLbd)
	if err != nil {
		poolBuf.Reset()
		return
	}
	poolBuf.Write(a.efwls6Jy(macKey, poolBuf.Bytes()[7:])[:4])
	poolBuf.ReadFrom(io.LimitReader(rand.Reader, int64(randDataLength)))
	poolBuf.Write(data)
	poolBuf.Write(a.efwls6Jy(a.oYNnYN0vN, poolBuf.Bytes())[:4])
}

// (*sdA1Al25W).kLUMSZc: 0x1b36e40.
func (a *sdA1Al25W) kLUMSZc(size int) int {
	if size > 400 {
		return mathRand.Intn(512)
	}
	return mathRand.Intn(1024)
}

// (*sdA1Al25W).bQmucT1jtu: 0x1b36ea0.
func (a *sdA1Al25W) bQmucT1jtu(poolBuf *bytes.Buffer, size int) {
	if size < 128 {
		poolBuf.WriteByte(byte(size + 1))
		poolBuf.ReadFrom(io.LimitReader(rand.Reader, int64(size)))
		return
	}
	poolBuf.WriteByte(255)
	binary.Write(poolBuf, binary.LittleEndian, uint16(size+3))
	poolBuf.ReadFrom(io.LimitReader(rand.Reader, int64(size)))
}
