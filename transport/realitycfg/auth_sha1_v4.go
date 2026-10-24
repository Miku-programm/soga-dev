// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package realitycfg

import (
	"bytes"
	cryptoRand "crypto/rand"
	"encoding/binary"
	"hash/adler32"
	"hash/crc32"
	"io"
	"math/rand"
	"net"

	"recovered.local/soga/transport/tlsutil"
	"recovered.local/soga/buf/syncpool"
)

// init: 0x1b338e0.
func init() {
	register("auth_sha1_v4", eQPSiGzN, 7)
}

type mh5mOR8HqIcG struct {
	*PZfVxx4UGRb
	*t1vL_y55t
	j4sXbBI9LwU []byte
	q5zf4h      bool
	enGr6iXKNAp bool
}

// eQPSiGzN: 0x1b3d2e0.
func eQPSiGzN(b *PZfVxx4UGRb) Shdj3DH1 {
	return &mh5mOR8HqIcG{PZfVxx4UGRb: b, t1vL_y55t: &t1vL_y55t{}}
}

// (*mh5mOR8HqIcG).StreamConn: 0x1b3d3a0.
func (a *mh5mOR8HqIcG) StreamConn(c net.Conn, j4sXbBI9LwU []byte) net.Conn {
	p := &mh5mOR8HqIcG{PZfVxx4UGRb: a.PZfVxx4UGRb, t1vL_y55t: a.uU1wpTCSA()}
	p.j4sXbBI9LwU = j4sXbBI9LwU
	return &A6KE0r{Conn: c, Shdj3DH1: p}
}

// (*mh5mOR8HqIcG).PacketConn: 0x1b3d500.
func (a *mh5mOR8HqIcG) PacketConn(c net.PacketConn) net.PacketConn {
	return c
}

// (*mh5mOR8HqIcG).Decode: 0x1b3d520.
func (a *mh5mOR8HqIcG) Decode(dst, src *bytes.Buffer) error {
	if a.enGr6iXKNAp {
		dst.ReadFrom(src)
		return nil
	}
	for src.Len() > 4 {
		if uint16(crc32.ChecksumIEEE(src.Bytes()[:2])&0xffff) != binary.LittleEndian.Uint16(src.Bytes()[2:4]) {
			src.Reset()
			return errorAt5098590
		}

		length := int(binary.BigEndian.Uint16(src.Bytes()[:2]))
		if length >= 8192 || length < 7 {
			a.enGr6iXKNAp = true
			src.Reset()
			return errorAt50985A0
		}
		if length > src.Len() {
			break
		}

		if adler32.Checksum(src.Bytes()[:length-4]) != binary.LittleEndian.Uint32(src.Bytes()[length-4:length]) {
			a.enGr6iXKNAp = true
			src.Reset()
			return errorAt50985B0
		}

		pos := int(src.Bytes()[4])
		if pos < 255 {
			pos += 4
		} else {
			pos = int(binary.BigEndian.Uint16(src.Bytes()[5:7])) + 4
		}
		dst.Write(src.Bytes()[pos : length-4])
		src.Next(length)
	}
	return nil
}

// (*mh5mOR8HqIcG).Encode: 0x1b3d880.
func (a *mh5mOR8HqIcG) Encode(buf *bytes.Buffer, b []byte) error {
	if !a.q5zf4h {
		dataLength := t4pamzUU(b)

		a.upMb05UedR(buf, b[:dataLength])
		b = b[dataLength:]

		a.q5zf4h = true
	}
	for len(b) > 8100 {
		a.pZRlqlGM1E(buf, b[:8100])
		b = b[8100:]
	}
	if len(b) > 0 {
		a.pZRlqlGM1E(buf, b)
	}

	return nil
}

// (*mh5mOR8HqIcG).DecodePacket: 0x1b3da00.
func (a *mh5mOR8HqIcG) DecodePacket(b []byte) ([]byte, error) { return b, nil }

// (*mh5mOR8HqIcG).EncodePacket: 0x1b3da20.
func (a *mh5mOR8HqIcG) EncodePacket(buf *bytes.Buffer, b []byte) error {
	buf.Write(b)
	return nil
}

// (*mh5mOR8HqIcG).pZRlqlGM1E: 0x1b3daa0.
func (a *mh5mOR8HqIcG) pZRlqlGM1E(poolBuf *bytes.Buffer, data []byte) {
	dataLength := len(data)
	randDataLength := a.lT0bGN2F(dataLength)
	/*
		2:	uint16 BigEndian packedDataLength
		2:	uint16 LittleEndian crc32Data & 0xffff
		3:	maxRandDataLengthPrefix (min:1)
		4:	adler32Data
	*/
	packedDataLength := 2 + 2 + 3 + randDataLength + dataLength + 4
	if randDataLength < 128 {
		packedDataLength -= 2
	}

	binary.Write(poolBuf, binary.BigEndian, uint16(packedDataLength))
	binary.Write(poolBuf, binary.LittleEndian, uint16(crc32.ChecksumIEEE(poolBuf.Bytes()[poolBuf.Len()-2:])&0xffff))
	a.bQmucT1jtu(poolBuf, randDataLength)
	poolBuf.Write(data)
	binary.Write(poolBuf, binary.LittleEndian, adler32.Checksum(poolBuf.Bytes()[poolBuf.Len()-packedDataLength+4:]))
}

// (*mh5mOR8HqIcG).upMb05UedR: 0x1b3dd00.
func (a *mh5mOR8HqIcG) upMb05UedR(poolBuf *bytes.Buffer, data []byte) {
	dataLength := len(data)
	randDataLength := a.lT0bGN2F(12 + dataLength)
	/*
		2:	uint16 BigEndian packedAuthDataLength
		4:	uint32 LittleEndian crc32Data
		3:	maxRandDataLengthPrefix (min: 1)
		12:	authDataLength
		10:	hmacSHA1DataLength
	*/
	packedAuthDataLength := 2 + 4 + 3 + randDataLength + 12 + dataLength + 10
	if randDataLength < 128 {
		packedAuthDataLength -= 2
	}

	salt := []byte("auth_sha1_v4")
	crcData := syncpool.SharedPoolAt5096978.Get(len(salt) + len(a.Z2UHwhi2h4) + 2)
	defer syncpool.SharedPoolAt5096978.Put(crcData)
	binary.BigEndian.PutUint16(crcData, uint16(packedAuthDataLength))
	copy(crcData[2:], salt)
	copy(crcData[2+len(salt):], a.Z2UHwhi2h4)

	key := syncpool.SharedPoolAt5096978.Get(len(a.j4sXbBI9LwU) + len(a.Z2UHwhi2h4))
	defer syncpool.SharedPoolAt5096978.Put(key)
	copy(key, a.j4sXbBI9LwU)
	copy(key[len(a.j4sXbBI9LwU):], a.Z2UHwhi2h4)

	poolBuf.Write(crcData[:2])
	binary.Write(poolBuf, binary.LittleEndian, crc32.ChecksumIEEE(crcData))
	a.bQmucT1jtu(poolBuf, randDataLength)
	a.mHxvPk(poolBuf)
	poolBuf.Write(data)
	poolBuf.Write(tlsutil.B4_4_BnW(key, poolBuf.Bytes()[poolBuf.Len()-packedAuthDataLength+10:])[:10])
}

// (*mh5mOR8HqIcG).bQmucT1jtu: 0x1b3e200.
func (a *mh5mOR8HqIcG) bQmucT1jtu(poolBuf *bytes.Buffer, size int) {
	if size < 128 {
		poolBuf.WriteByte(byte(size + 1))
		poolBuf.ReadFrom(io.LimitReader(cryptoRand.Reader, int64(size)))
		return
	}
	poolBuf.WriteByte(255)
	binary.Write(poolBuf, binary.BigEndian, uint16(size+3))
	poolBuf.ReadFrom(io.LimitReader(cryptoRand.Reader, int64(size)))
}

// (*mh5mOR8HqIcG).lT0bGN2F: 0x1b3e3c0.
func (a *mh5mOR8HqIcG) lT0bGN2F(size int) int {
	if size > 1200 {
		return 0
	}
	if size > 400 {
		return rand.Intn(256)
	}
	return rand.Intn(512)
}
