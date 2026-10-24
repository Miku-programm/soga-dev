package vless

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"sync"
	"time"
)

// GPFxUP0_ retains the complete original 232-byte client handshake state.
type GPFxUP0_ struct {
	QUnOciyca6eb []any
	JTmPRiuw     [][]byte
	J2battR3S    [][32]byte
	JhDzI4ehWF   int
	S799IXI      uint32
	BsCV7uFJFW   uint32
	Q8W_Oz       [][3]int
	PVif2PW      [][3]int
	QgCGo4Q6lLF  sync.RWMutex
	EQa3Gaz      time.Time
	Wliqugg3J89  []byte
	HMnsiUUhd    []byte
}

// RfBfjVN4XE retains all original fields, including the obfuscated private
// names used in the RTTI rather than the names of upstream implementations.
type RfBfjVN4XE struct {
	DrpXao
	HiUDG1Fl    bool
	R5MQKRg26Z  *GPFxUP0_
	ENEXKXNftX  []byte
	A9M03eHa    []byte
	AVDV5UPQ    *RaS2myUz
	EufjD7UrtI  *RaS2myUz
	WWjo2UGyM   []byte
	y6fLiKc     bytes.Buffer
	ponyGFmB4cO bytes.Reader
}

var writePoolAt5042620 = sync.Pool{New: func() any { return make([]byte, 8213) }}
var nonceMaximumAt509A1F0 = bytes.Repeat([]byte{255}, 12)

// Write recovers 0x1cf6d60. The original ignores short nil-error writes and
// returns zero on any write error even if earlier records were transmitted.
func (c *RfBfjVN4XE) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	buffer := writePoolAt5042620.Get().([]byte)
	defer writePoolAt5042620.Put(buffer)
	for offset := 0; offset < len(p); {
		n := min(8192, len(p)-offset)
		record := buffer[:n+21]
		record[0] = 23
		record[1] = 3
		record[2] = 3
		binary.BigEndian.PutUint16(record[3:5], uint16(n+16))
		rotate := bytes.Equal(c.AVDV5UPQ.NQGg9x[:], nonceMaximumAt509A1F0)
		c.AVDV5UPQ.Seal(record[:5], nil, p[offset:offset+n], record[:5])
		if rotate {
			c.AVDV5UPQ = WrnUSnC6(record, c.ENEXKXNftX, c.HiUDG1Fl)
		}
		if c.A9M03eHa != nil {
			record = append(c.A9M03eHa, record...)
			c.A9M03eHa = nil
		}
		if _, err := c.DrpXao.Write(record); err != nil {
			return 0, err
		}
		offset += n
	}
	return len(p), nil
}

// Read recovers 0x1cf72a0, including lazy salt, padding, session invalidation,
// ciphertext-driven key rotation and plaintext buffered across small reads.
func (c *RfBfjVN4XE) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if c.EufjD7UrtI == nil {
		salt := make([]byte, 16)
		if _, err := io.ReadFull(c.DrpXao, salt); err != nil {
			return 0, err
		}
		c.EufjD7UrtI = WrnUSnC6(salt, c.ENEXKXNftX, c.HiUDG1Fl)
		if stream, ok := c.DrpXao.(*LUKUT5W0R6); ok {
			stream.EO5NXI6ex6 = GdvNyLM(c.ENEXKXNftX, salt)
		}
	}
	if c.WWjo2UGyM != nil {
		if _, err := io.ReadFull(c.DrpXao, c.WWjo2UGyM); err != nil {
			return 0, err
		}
		if _, err := c.EufjD7UrtI.Open(c.WWjo2UGyM[:0], nil, c.WWjo2UGyM, nil); err != nil {
			return 0, err
		}
		c.WWjo2UGyM = nil
	}
	if c.ponyGFmB4cO.Len() > 0 {
		return c.ponyGFmB4cO.Read(p)
	}
	var header [5]byte
	if _, err := io.ReadFull(c.DrpXao, header[:]); err != nil {
		return 0, err
	}
	n, err := QkDUgv(header[:])
	if err != nil {
		if c.R5MQKRg26Z != nil && strings.Contains(err.Error(), "invalid header: ") {
			state := c.R5MQKRg26Z
			state.QgCGo4Q6lLF.Lock()
			if len(c.ENEXKXNftX) >= len(state.Wliqugg3J89) && bytes.Equal(c.ENEXKXNftX[:len(state.Wliqugg3J89)], state.Wliqugg3J89) {
				state.EQa3Gaz = time.Now()
			}
			state.QgCGo4Q6lLF.Unlock()
			return 0, errors.New("new handshake needed")
		}
		return 0, err
	}
	c.R5MQKRg26Z = nil
	if c.y6fLiKc.Cap() < n {
		c.y6fLiKc.Grow(n)
	}
	ciphertext := c.y6fLiKc.Bytes()[:n]
	if _, err := io.ReadFull(c.DrpXao, ciphertext); err != nil {
		return 0, err
	}
	plaintext := ciphertext[:n-16]
	if len(p) >= n-16 {
		plaintext = p[:n-16]
	}
	var next *RaS2myUz
	if bytes.Equal(c.EufjD7UrtI.NQGg9x[:], nonceMaximumAt509A1F0) {
		context := append(header[:], ciphertext...)
		next = WrnUSnC6(context, c.ENEXKXNftX, c.HiUDG1Fl)
	}
	_, err = c.EufjD7UrtI.Open(plaintext[:0], nil, ciphertext, header[:])
	// The original switches key state even when authentication fails.
	if next != nil {
		c.EufjD7UrtI = next
	}
	if err != nil {
		return 0, err
	}
	if len(p) >= n-16 {
		return n - 16, nil
	}
	copy(p, plaintext)
	c.ponyGFmB4cO.Reset(plaintext[len(p):])
	return len(p), nil
}
