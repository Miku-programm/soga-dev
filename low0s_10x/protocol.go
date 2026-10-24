// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package low0s_10x

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net"
)

var (
	errorAt5098590 = errors.New("auth_sha1_v4 decode data wrong crc32")
	errorAt50985A0 = errors.New("auth_sha1_v4 decode data wrong length")
	errorAt50985B0 = errors.New("auth_sha1_v4 decode data wrong adler32")
	errorAt50985C0 = errors.New("auth_aes128 decode data wrong mac")
	errorAt50985D0 = errors.New("auth_aes128 decode data wrong length")
	errorAt50985E0 = errors.New("auth_aes128 decode data wrong checksum")
	errorAt50985F0 = errors.New("auth_chain decode data wrong length")
	errorAt5098600 = errors.New("auth_chain decode data wrong checksum")
)

type Shdj3DH1 interface {
	StreamConn(net.Conn, []byte) net.Conn
	PacketConn(net.PacketConn) net.PacketConn
	Decode(dst, src *bytes.Buffer) error
	Encode(buf *bytes.Buffer, b []byte) error
	DecodePacket([]byte) ([]byte, error)
	EncodePacket(buf *bytes.Buffer, b []byte) error
}

type jn3xSYNOk func(b *PZfVxx4UGRb) Shdj3DH1

var registryAt50964B0 = make(map[string]struct {
	overhead int
	new      jn3xSYNOk
})

// register restores the map writes in original init.0 through init.5.
func register(name string, c jn3xSYNOk, o int) {
	registryAt50964B0[name] = struct {
		overhead int
		new      jn3xSYNOk
	}{overhead: o, new: c}
}

// HBLvlfQr0d: 0x1b401c0.
func HBLvlfQr0d(name string, b *PZfVxx4UGRb) (Shdj3DH1, error) {
	if choice, ok := registryAt50964B0[name]; ok {
		b.UqLugVVWCKJ += choice.overhead
		return choice.new(b), nil
	}
	return nil, fmt.Errorf("protocol %s not supported", name)
}

// getHeadSize names the switch inlined into t4pamzUU at 0x1b40340.
func getHeadSize(b []byte, defaultValue int) int {
	if len(b) < 2 {
		return defaultValue
	}
	headType := b[0] & 7
	switch headType {
	case 1:
		return 7
	case 4:
		return 19
	case 3:
		return 4 + int(b[1])
	}
	return defaultValue
}

// t4pamzUU: 0x1b40340.
func t4pamzUU(b []byte) int {
	bLength := len(b)
	dataLength := getHeadSize(b, 30) + rand.Intn(32)
	if bLength < dataLength {
		return bLength
	}
	return dataLength
}
