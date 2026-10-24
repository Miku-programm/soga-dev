// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package low0s_10x

import (
	"net"
	"sort"

	"recovered.local/soga/H8_In_Vuy"
)

// init: 0x1b338e0.
func init() {
	register("auth_chain_b", fE2Ibsq, 4)
}

type cQb8tRiyK struct {
	*vcMVzsjn5n
	fbupx9aL []int
	r24E2R   []int
}

// fE2Ibsq: 0x1b3c120.
func fE2Ibsq(b *PZfVxx4UGRb) Shdj3DH1 {
	a := &cQb8tRiyK{
		vcMVzsjn5n: &vcMVzsjn5n{
			PZfVxx4UGRb: b,
			t1vL_y55t:   &t1vL_y55t{},
			bhRsfpsWw:   &bhRsfpsWw{},
			cLIdrsul1DU: "auth_chain_b",
		},
	}
	a.eatUbiR0d()
	return a
}

// (*cQb8tRiyK).StreamConn: 0x1b3c2c0.
func (a *cQb8tRiyK) StreamConn(c net.Conn, fTzASkj []byte) net.Conn {
	p := &cQb8tRiyK{
		vcMVzsjn5n: &vcMVzsjn5n{
			PZfVxx4UGRb: a.PZfVxx4UGRb,
			t1vL_y55t:   a.uU1wpTCSA(),
			bhRsfpsWw:   a.bhRsfpsWw,
			cLIdrsul1DU: a.cLIdrsul1DU,
			j4NIsaaxx:   1,
			nKJS2vXY6:   1,
		},
	}
	p.fTzASkj = fTzASkj
	p.eXaIwIUeDfSi = p.fv9XBEr7xo
	p.dKnDDY()
	return &A6KE0r{Conn: c, Shdj3DH1: p}
}

// (*cQb8tRiyK).dKnDDY: 0x1b3c520.
func (a *cQb8tRiyK) dKnDDY() {
	a.fbupx9aL = a.fbupx9aL[:0]
	a.r24E2R = a.r24E2R[:0]

	a.nIbhQaw0D9r.InitFromBin(a.Z2UHwhi2h4)
	length := a.nIbhQaw0D9r.Next()%8 + 4
	for ; length > 0; length-- {
		a.fbupx9aL = append(a.fbupx9aL, int(a.nIbhQaw0D9r.Next()%2340%2040%1440))
	}
	sort.Ints(a.fbupx9aL)

	length = a.nIbhQaw0D9r.Next()%16 + 8
	for ; length > 0; length-- {
		a.r24E2R = append(a.r24E2R, int(a.nIbhQaw0D9r.Next()%2340%2040%1440))
	}
	sort.Ints(a.r24E2R)
}

// (*cQb8tRiyK).fv9XBEr7xo: 0x1b3c8e0.
func (a *cQb8tRiyK) fv9XBEr7xo(length int, lashHash []byte, random *H8_In_Vuy.YE9p7Ja6vRfW) int {
	if length >= 1440 {
		return 0
	}
	random.InitFromBinAndLength(lashHash, length)
	pos := sort.Search(len(a.fbupx9aL), func(i int) bool { return a.fbupx9aL[i] >= length+a.UqLugVVWCKJ })
	finalPos := pos + int(random.Next()%uint64(len(a.fbupx9aL)))
	if finalPos < len(a.fbupx9aL) {
		return a.fbupx9aL[finalPos] - length - a.UqLugVVWCKJ
	}

	pos = sort.Search(len(a.r24E2R), func(i int) bool { return a.r24E2R[i] >= length+a.UqLugVVWCKJ })
	finalPos = pos + int(random.Next()%uint64(len(a.r24E2R)))
	if finalPos < len(a.r24E2R) {
		return a.r24E2R[finalPos] - length - a.UqLugVVWCKJ
	}
	if finalPos < pos+len(a.r24E2R)-1 {
		return 0
	}
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
