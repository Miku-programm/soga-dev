// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package low0s_10x

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rc4"
	"encoding/base64"
	"encoding/binary"
	"io"
	"net"
	"strconv"
	"strings"

	"recovered.local/soga/H8_In_Vuy"
	"recovered.local/soga/m3Cw8P3_Le"
	"recovered.local/soga/mzi1d3gDO7u"
	"recovered.local/soga/x6HWbunbF"
)

// init: 0x1b338e0.
func init() {
	register("auth_chain_a", kB6Vhatt, 4)
}

type yBZDYfbTgTP3 func(int, []byte, *H8_In_Vuy.YE9p7Ja6vRfW) int

type vcMVzsjn5n struct {
	*PZfVxx4UGRb
	*t1vL_y55t
	*bhRsfpsWw
	fTzASkj      []byte
	cLIdrsul1DU  string
	htDwG7h1F    bool
	aeUq8k4O     bool
	knDyablm6A   []byte
	y4eh_DDXZsU  []byte
	s0XnF9FH     cipher.Stream
	_9Od97DLOEI  cipher.Stream
	f5YPnAoB     H8_In_Vuy.YE9p7Ja6vRfW
	nIbhQaw0D9r  H8_In_Vuy.YE9p7Ja6vRfW
	eXaIwIUeDfSi yBZDYfbTgTP3
	j4NIsaaxx    uint32
	nKJS2vXY6    uint32
}

// kB6Vhatt: 0x1b37780.
func kB6Vhatt(b *PZfVxx4UGRb) Shdj3DH1 {
	a := &vcMVzsjn5n{
		PZfVxx4UGRb: b,
		t1vL_y55t:   &t1vL_y55t{},
		bhRsfpsWw:   &bhRsfpsWw{},
		cLIdrsul1DU: "auth_chain_a",
	}
	a.eatUbiR0d()
	return a
}

// (*vcMVzsjn5n).eatUbiR0d: 0x1b378e0.
func (a *vcMVzsjn5n) eatUbiR0d() {
	params := strings.Split(a.BOJR3xVshAb, ":")
	if len(params) > 1 {
		if qC81k5Mk, err := strconv.ParseUint(params[0], 10, 32); err == nil {
			binary.LittleEndian.PutUint32(a.qC81k5Mk[:], uint32(qC81k5Mk))
			a.oYNnYN0vN = []byte(params[1])
		} else {
			x6HWbunbF.Dl2Gj6("Wrong protocol-param for %s, only digits are expected before ':'", a.cLIdrsul1DU)
		}
	}
	if len(a.oYNnYN0vN) == 0 {
		a.oYNnYN0vN = a.Z2UHwhi2h4
		rand.Read(a.qC81k5Mk[:])
	}
}

// (*vcMVzsjn5n).StreamConn: 0x1b37ac0.
func (a *vcMVzsjn5n) StreamConn(c net.Conn, fTzASkj []byte) net.Conn {
	p := &vcMVzsjn5n{
		PZfVxx4UGRb: a.PZfVxx4UGRb,
		t1vL_y55t:   a.uU1wpTCSA(),
		bhRsfpsWw:   a.bhRsfpsWw,
		cLIdrsul1DU: a.cLIdrsul1DU,
		j4NIsaaxx:   1,
		nKJS2vXY6:   1,
	}
	p.fTzASkj = fTzASkj
	p.eXaIwIUeDfSi = p.fv9XBEr7xo
	return &A6KE0r{Conn: c, Shdj3DH1: p}
}

// (*vcMVzsjn5n).PacketConn: 0x1b37ce0.
func (a *vcMVzsjn5n) PacketConn(c net.PacketConn) net.PacketConn {
	p := &vcMVzsjn5n{
		PZfVxx4UGRb: a.PZfVxx4UGRb,
		cLIdrsul1DU: a.cLIdrsul1DU,
		bhRsfpsWw:   a.bhRsfpsWw,
	}
	return &S_N0Ixh{WCApOca6IQW7: c, Shdj3DH1: p}
}

// (*vcMVzsjn5n).Decode: 0x1b37e20.
func (a *vcMVzsjn5n) Decode(dst, src *bytes.Buffer) error {
	if a.aeUq8k4O {
		dst.ReadFrom(src)
		return nil
	}
	for src.Len() > 4 {
		macKey := mzi1d3gDO7u.SharedPoolAt5096978.Get(len(a.oYNnYN0vN) + 4)
		defer mzi1d3gDO7u.SharedPoolAt5096978.Put(macKey)
		copy(macKey, a.oYNnYN0vN)
		binary.LittleEndian.PutUint32(macKey[len(a.oYNnYN0vN):], a.nKJS2vXY6)

		dataLength := int(binary.LittleEndian.Uint16(src.Bytes()[:2]) ^ binary.LittleEndian.Uint16(a.y4eh_DDXZsU[14:16]))
		eXaIwIUeDfSi := a.eXaIwIUeDfSi(dataLength, a.y4eh_DDXZsU, &a.nIbhQaw0D9r)
		length := dataLength + eXaIwIUeDfSi

		if length >= 4096 {
			a.aeUq8k4O = true
			src.Reset()
			return errorAt50985F0
		}

		if 4+length > src.Len() {
			break
		}

		serverHash := H8_In_Vuy.YPrsw38MN(macKey, src.Bytes()[:length+2])
		if !bytes.Equal(serverHash[:2], src.Bytes()[length+2:length+4]) {
			a.aeUq8k4O = true
			src.Reset()
			return errorAt5098600
		}
		a.y4eh_DDXZsU = serverHash

		pos := 2
		if dataLength > 0 && eXaIwIUeDfSi > 0 {
			pos += getRandStartPos(eXaIwIUeDfSi, &a.nIbhQaw0D9r)
		}
		wantedData := src.Bytes()[pos : pos+dataLength]
		a._9Od97DLOEI.XORKeyStream(wantedData, wantedData)
		if a.nKJS2vXY6 == 1 {
			dst.Write(wantedData[2:])
		} else {
			dst.Write(wantedData)
		}
		a.nKJS2vXY6++
		src.Next(length + 4)
	}
	return nil
}

// (*vcMVzsjn5n).Encode: 0x1b385a0.
func (a *vcMVzsjn5n) Encode(buf *bytes.Buffer, b []byte) error {
	if !a.htDwG7h1F {
		dataLength := t4pamzUU(b)
		a.upMb05UedR(buf, b[:dataLength])
		b = b[dataLength:]
		a.htDwG7h1F = true
	}
	for len(b) > 2800 {
		a.pZRlqlGM1E(buf, b[:2800])
		b = b[2800:]
	}
	if len(b) > 0 {
		a.pZRlqlGM1E(buf, b)
	}
	return nil
}

// (*vcMVzsjn5n).DecodePacket: 0x1b38720.
func (a *vcMVzsjn5n) DecodePacket(b []byte) ([]byte, error) {
	if len(b) < 9 {
		return nil, errorAt50985F0
	}
	if !bytes.Equal(H8_In_Vuy.YPrsw38MN(a.oYNnYN0vN, b[:len(b)-1])[:1], b[len(b)-1:]) {
		return nil, errorAt5098600
	}
	md5Data := H8_In_Vuy.YPrsw38MN(a.Z2UHwhi2h4, b[len(b)-8:len(b)-1])

	eXaIwIUeDfSi := zDiFyWALBK(md5Data, &a.nIbhQaw0D9r)

	key := m3Cw8P3_Le.ArNrGS_EV(base64.StdEncoding.EncodeToString(a.oYNnYN0vN)+base64.StdEncoding.EncodeToString(md5Data), 16)
	rc4Cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	wantedData := b[:len(b)-8-eXaIwIUeDfSi]
	rc4Cipher.XORKeyStream(wantedData, wantedData)
	return wantedData, nil
}

// (*vcMVzsjn5n).EncodePacket: 0x1b389c0.
func (a *vcMVzsjn5n) EncodePacket(buf *bytes.Buffer, b []byte) error {
	t1vL_y55t := mzi1d3gDO7u.SharedPoolAt5096978.Get(3)
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(t1vL_y55t)
	rand.Read(t1vL_y55t)

	md5Data := H8_In_Vuy.YPrsw38MN(a.Z2UHwhi2h4, t1vL_y55t)

	eXaIwIUeDfSi := zDiFyWALBK(md5Data, &a.f5YPnAoB)

	key := m3Cw8P3_Le.ArNrGS_EV(base64.StdEncoding.EncodeToString(a.oYNnYN0vN)+base64.StdEncoding.EncodeToString(md5Data), 16)
	rc4Cipher, err := rc4.NewCipher(key)
	if err != nil {
		return err
	}
	rc4Cipher.XORKeyStream(b, b)

	buf.Write(b)
	buf.ReadFrom(io.LimitReader(rand.Reader, int64(eXaIwIUeDfSi)))
	buf.Write(t1vL_y55t)
	binary.Write(buf, binary.LittleEndian, binary.LittleEndian.Uint32(a.qC81k5Mk[:])^binary.LittleEndian.Uint32(md5Data[:4]))
	buf.Write(H8_In_Vuy.YPrsw38MN(a.oYNnYN0vN, buf.Bytes())[:1])
	return nil
}

// (*vcMVzsjn5n).upMb05UedR: 0x1b38e40.
func (a *vcMVzsjn5n) upMb05UedR(poolBuf *bytes.Buffer, data []byte) {
	/*
		dataLength := len(data)
		12:	checkHead(4) and hmac of checkHead(8)
		4:	uint32 LittleEndian uid (uid = userID ^ last client hash)
		16:	encrypted data of authdata(12), uint16 LittleEndian overhead(2) and uint16 LittleEndian number zero(2)
		4:	last server hash(4)
		packedAuthDataLength := 12 + 4 + 16 + 4 + dataLength
	*/

	macKey := mzi1d3gDO7u.SharedPoolAt5096978.Get(len(a.fTzASkj) + len(a.Z2UHwhi2h4))
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(macKey)
	copy(macKey, a.fTzASkj)
	copy(macKey[len(a.fTzASkj):], a.Z2UHwhi2h4)

	// check head
	poolBuf.ReadFrom(io.LimitReader(rand.Reader, int64(4)))
	a.knDyablm6A = H8_In_Vuy.YPrsw38MN(macKey, poolBuf.Bytes())
	a.zWOwZlmS()
	poolBuf.Write(a.knDyablm6A[:8])
	// uid
	binary.Write(poolBuf, binary.LittleEndian, binary.LittleEndian.Uint32(a.qC81k5Mk[:])^binary.LittleEndian.Uint32(a.knDyablm6A[8:12]))
	// encrypted data
	err := a.vKsdajuSitet(poolBuf, a.oYNnYN0vN, [2]int{a.UqLugVVWCKJ, 0}, a.cLIdrsul1DU)
	if err != nil {
		poolBuf.Reset()
		return
	}
	// last server hash
	a.y4eh_DDXZsU = H8_In_Vuy.YPrsw38MN(a.oYNnYN0vN, poolBuf.Bytes()[12:])
	poolBuf.Write(a.y4eh_DDXZsU[:4])
	// packed data
	a.pZRlqlGM1E(poolBuf, data)
}

// (*vcMVzsjn5n).pZRlqlGM1E: 0x1b39320.
func (a *vcMVzsjn5n) pZRlqlGM1E(poolBuf *bytes.Buffer, data []byte) {
	a.s0XnF9FH.XORKeyStream(data, data)

	macKey := mzi1d3gDO7u.SharedPoolAt5096978.Get(len(a.oYNnYN0vN) + 4)
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(macKey)
	copy(macKey, a.oYNnYN0vN)
	binary.LittleEndian.PutUint32(macKey[len(a.oYNnYN0vN):], a.j4NIsaaxx)
	a.j4NIsaaxx++

	length := uint16(len(data)) ^ binary.LittleEndian.Uint16(a.knDyablm6A[14:16])

	originalLength := poolBuf.Len()
	binary.Write(poolBuf, binary.LittleEndian, length)
	a.xQ5i1_(poolBuf, data)
	a.knDyablm6A = H8_In_Vuy.YPrsw38MN(macKey, poolBuf.Bytes()[originalLength:])
	poolBuf.Write(a.knDyablm6A[:2])
}

// (*vcMVzsjn5n).xQ5i1_: 0x1b39660.
func (a *vcMVzsjn5n) xQ5i1_(poolBuf *bytes.Buffer, data []byte) {
	eXaIwIUeDfSi := a.eXaIwIUeDfSi(len(data), a.knDyablm6A, &a.f5YPnAoB)
	if len(data) == 0 {
		poolBuf.ReadFrom(io.LimitReader(rand.Reader, int64(eXaIwIUeDfSi)))
		return
	}
	if eXaIwIUeDfSi > 0 {
		startPos := getRandStartPos(eXaIwIUeDfSi, &a.f5YPnAoB)
		poolBuf.ReadFrom(io.LimitReader(rand.Reader, int64(startPos)))
		poolBuf.Write(data)
		poolBuf.ReadFrom(io.LimitReader(rand.Reader, int64(eXaIwIUeDfSi-startPos)))
		return
	}
	poolBuf.Write(data)
}

// getRandStartPos names the modulo arithmetic inlined in 0x1b37e20 and 0x1b39660.
func getRandStartPos(length int, random *H8_In_Vuy.YE9p7Ja6vRfW) int {
	if length == 0 {
		return 0
	}
	return int(int64(random.Next()%8589934609) % int64(length))
}

// (*vcMVzsjn5n).fv9XBEr7xo: 0x1b39980.
func (a *vcMVzsjn5n) fv9XBEr7xo(length int, lastHash []byte, random *H8_In_Vuy.YE9p7Ja6vRfW) int {
	if length > 1440 {
		return 0
	}
	random.InitFromBinAndLength(lastHash, length)
	if length > 1300 {
		return int(random.Next() % 31)
	}
	if length > 900 {
		return int(random.Next() % 127)
	}
	if length > 400 {
		return int(random.Next() % 521)
	}
	return int(random.Next() % 1021)
}

// (*vcMVzsjn5n).zWOwZlmS: 0x1b39c00.
func (a *vcMVzsjn5n) zWOwZlmS() {
	key := m3Cw8P3_Le.ArNrGS_EV(base64.StdEncoding.EncodeToString(a.oYNnYN0vN)+base64.StdEncoding.EncodeToString(a.knDyablm6A), 16)
	a.s0XnF9FH, _ = rc4.NewCipher(key)
	a._9Od97DLOEI, _ = rc4.NewCipher(key)
}

// zDiFyWALBK: 0x1b39d40.
func zDiFyWALBK(lastHash []byte, random *H8_In_Vuy.YE9p7Ja6vRfW) int {
	random.InitFromBin(lastHash)
	return int(random.Next() % 127)
}
