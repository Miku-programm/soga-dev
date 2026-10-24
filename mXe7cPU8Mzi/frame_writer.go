package mXe7cPU8Mzi

import (
	"crypto/cipher"
	"fmt"
	"net"
	D "recovered.local/soga/D_zCag1qTV"
	U "recovered.local/soga/Ualx5UM_D6I"
	"sync"
	"time"
)

// qhFcsZC is the original 88-byte writer RTTI at 0x36df248.
type qhFcsZC struct {
	v_7InizyTC22 sync.Mutex
	_z_j84       net.Conn
	yn5oQQkf8    cipher.AEAD
	g9zKwmF      []byte
	oynKh2Fa     bool
	rAB3kZi8     string
}

func vfhZtLu(conn net.Conn, aead cipher.AEAD, password string) (*qhFcsZC, error) { //0x1c535c0
	nonce, err := o0fNBhk(password)
	if err != nil {
		return nil, fmt.Errorf("generate nonce: %w", err)
	}
	return &qhFcsZC{_z_j84: conn, yn5oQQkf8: aead, g9zKwmF: nonce, oynKh2Fa: true, rAB3kZi8: password}, nil
}

// incrementNonce is an analyst name for the repeated original big-endian loop.
func incrementNonce(nonce []byte) {
	for i := len(nonce) - 1; i >= 0; i-- {
		nonce[i]++
		if nonce[i] != 0 {
			break
		}
	}
}

func (w *qhFcsZC) bBLmwPhjPp54(b U.Mf3phSOVqgf, payload [][]byte) { //0x1c54100
	if len(payload) == 1 {
		b.SealWrite(w.yn5oQQkf8, w.g9zKwmF, payload[0], nil)
	} else {
		start := b.Len()
		for _, part := range payload[:len(payload)-1] {
			_, _ = b.Write(part)
		}
		b.SealWriteWithStart(w.yn5oQQkf8, start, w.g9zKwmF, payload[len(payload)-1], nil)
	}
	incrementNonce(w.g9zKwmF)
}

// WriteSessionSegment @0x1c53820. Header fields are updated before locking.
// Only the error from the single connection Write is returned; a short nil
// Write is not converted to io.ErrShortWrite by the original.
func (w *qhFcsZC) WriteSessionSegment(m *IqB1foHWg, payload ...[]byte) error {
	size := 0
	for _, p := range payload {
		size += len(p)
	}
	padding := a2dHepfMMQ(size, w.rAB3kZi8)
	m.HPUJVQ0S = uint16(size)
	m.FO7MEJ = byte(len(padding))
	m.WyK3yKmJKe1V = uint32(time.Now().Unix() / 60)
	w.v_7InizyTC22.Lock()
	defer w.v_7InizyTC22.Unlock()
	header := m.Marshal()
	capacity := len(padding) + 48
	if w.oynKh2Fa {
		capacity += 24
	}
	if size > 0 {
		capacity += size + 16
	}
	b := D.S4OQXWIZ(capacity)
	defer b.Free()
	if w.oynKh2Fa {
		_, _ = b.Write(w.g9zKwmF)
	}
	b.SealWrite(w.yn5oQQkf8, w.g9zKwmF, header[:], nil)
	incrementNonce(w.g9zKwmF)
	w.oynKh2Fa = false
	if size > 0 {
		w.bBLmwPhjPp54(b, payload)
	}
	_, _ = b.Write(padding)
	_, err := w._z_j84.Write(b.Bytes())
	return err
}

func (w *qhFcsZC) WriteDataSegment(m *IqB1foHWg, payload ...[]byte) error { //0x1c53c80
	prefix, suffix := byte(ofUaiB3MIo(256)), byte(ofUaiB3MIo(256))
	size := 0
	for _, p := range payload {
		size += len(p)
	}
	m.B8QsA0N2yo = prefix
	m.HPUJVQ0S = uint16(size)
	m.FO7MEJ = suffix
	m.WyK3yKmJKe1V = uint32(time.Now().Unix() / 60)
	w.v_7InizyTC22.Lock()
	defer w.v_7InizyTC22.Unlock()
	header := m.Marshal()
	capacity := int(prefix) + int(suffix) + 48
	if w.oynKh2Fa {
		capacity += 24
	}
	if size > 0 {
		capacity += size + 16
	}
	b := D.S4OQXWIZ(capacity)
	defer b.Free()
	if w.oynKh2Fa {
		_, _ = b.Write(w.g9zKwmF)
	}
	b.SealWrite(w.yn5oQQkf8, w.g9zKwmF, header[:], nil)
	incrementNonce(w.g9zKwmF)
	w.oynKh2Fa = false
	b.WriteRandom(int(prefix))
	if size > 0 {
		w.bBLmwPhjPp54(b, payload)
	}
	b.WriteRandom(int(suffix))
	_, err := w._z_j84.Write(b.Bytes())
	return err
}
