package TaFdeUcnzQ

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	A "recovered.local/soga/A818JHb"
	J "recovered.local/soga/JafvCs5"
)

// bh0EfZohV restores 0x1c9d380. Replacing the per-packet AEAD retains an
// already allocated nonce, matching the original object's state transitions.
func (c *jSUbF0) bh0EfZohV(salt []byte) error {
	key := make([]byte, c.obzsFnGuS.geqiu3e)
	J.THj8k5(c.fkUrkngrG1x, salt, []byte("ss-subkey"), key)
	var e error
	c.et6j454EdcWG, e = c.obzsFnGuS.w5pQfi_Ptek(key)
	if e != nil {
		return e
	}
	if c.qaeiJJzh3hj == nil {
		c.qaeiJJzh3hj = make([]byte, c.et6j454EdcWG.NonceSize())
	}
	return nil
}

// n4dhtM restores 0x1c9dfa0.
func (c *jSUbF0) n4dhtM(salt []byte) error {
	key := make([]byte, c.obzsFnGuS.geqiu3e)
	J.THj8k5(c.fkUrkngrG1x, salt, []byte("ss-subkey"), key)
	var e error
	c.fsuDlw7T, e = c.obzsFnGuS.w5pQfi_Ptek(key)
	if e != nil {
		return e
	}
	if c.l8MMQ3mB9Vm == nil {
		c.l8MMQ3mB9Vm = make([]byte, c.fsuDlw7T.NonceSize())
	}
	return nil
}

// incrementNonce is the identical inline little-endian loop at the original
// Open/Seal call sites, including increment after failed Open.
func incrementNonce(n []byte) {
	for i := range n {
		n[i]++
		if n[i] != 0 {
			return
		}
	}
}

// nBjfJv restores 0x1c9d5c0: one authenticated length followed by one payload.
// The original accepts lengths through 32766; its writer splits at 16383.
func (c *jSUbF0) nBjfJv(p []byte) (int, error) {
	header, e := c.GetReadBuf(c.et6j454EdcWG.Overhead() + 2)
	if e != nil {
		return 0, e
	}
	if _, e = io.ReadFull(c.eMYicxBI, header); e != nil {
		return 0, e
	}
	_, e = c.et6j454EdcWG.Open(header[:0], c.qaeiJJzh3hj, header, nil)
	incrementNonce(c.qaeiJJzh3hj)
	if e != nil {
		return 0, e
	}
	size := int(binary.BigEndian.Uint16(header))
	if size > 32766 {
		return 0, fmt.Errorf("payload size too large: %v", size)
	}
	data, e := c.GetReadBuf(size + c.et6j454EdcWG.Overhead())
	if e != nil {
		return 0, e
	}
	if _, e = io.ReadFull(c.eMYicxBI, data); e != nil {
		return 0, e
	}
	if len(p) < size {
		_, e = c.et6j454EdcWG.Open(data[:0], c.qaeiJJzh3hj, data, nil)
		incrementNonce(c.qaeiJJzh3hj)
		if e != nil {
			return 0, e
		}
		copy(p, data[:size])
		c.jjwGY3_N = data
		c.g9tmWrw57CtG = len(p)
		c.foRWabfgA2FO = size
		return len(p), nil
	}
	_, e = c.et6j454EdcWG.Open(p[:0], c.qaeiJJzh3hj, data, nil)
	incrementNonce(c.qaeiJJzh3hj)
	if e != nil {
		return 0, e
	}
	return size, nil
}

// Read restores 0x1c9dc60. Like the stream path it uses a single Read for salt.
func (c *jSUbF0) Read(p []byte) (int, error) {
	if c.et6j454EdcWG == nil {
		salt := make([]byte, max(16, c.obzsFnGuS.geqiu3e))
		n, e := c.eMYicxBI.Read(salt)
		if e != nil {
			return 0, e
		}
		if n != len(salt) {
			return 0, fmt.Errorf("salt len not enough: %v", n)
		}
		if !A.HFlnXWty(salt) {
			return 0, errorAt5098DA0
		}
		if e = c.bh0EfZohV(salt); e != nil {
			return 0, e
		}
	}
	if c.g9tmWrw57CtG < 0 {
		return c.nBjfJv(p)
	}
	n := copy(p, c.jjwGY3_N[c.g9tmWrw57CtG:c.foRWabfgA2FO])
	c.g9tmWrw57CtG += n
	if c.g9tmWrw57CtG >= c.foRWabfgA2FO {
		c.jjwGY3_N = nil
		c.g9tmWrw57CtG = -1
	}
	return n, nil
}

// Write restores 0x1c9e1a0, including recursive splitting and ignoring a
// successful short write reported by the wrapped writer.
func (c *jSUbF0) Write(p []byte) (int, error) {
	if len(p) > 16383 {
		n, e := c.Write(p[:16383])
		if e != nil {
			return n, e
		}
		m, e := c.Write(p[16383:])
		return n + m, e
	}
	var salt []byte
	if c.fsuDlw7T == nil {
		salt = make([]byte, max(16, c.obzsFnGuS.geqiu3e))
		rand.Read(salt)
		if e := c.n4dhtM(salt); e != nil {
			return 0, e
		}
	}
	overhead := c.fsuDlw7T.Overhead()
	out, e := c.GetWriteBuf(len(salt) + 2 + overhead + len(p) + overhead)
	if e != nil {
		return 0, e
	}
	copy(out, salt)
	offset := len(salt)
	binary.BigEndian.PutUint16(out[offset:], uint16(len(p)))
	c.fsuDlw7T.Seal(out[offset:offset], c.l8MMQ3mB9Vm, out[offset:offset+2], nil)
	incrementNonce(c.l8MMQ3mB9Vm)
	offset += 2 + overhead
	c.fsuDlw7T.Seal(out[offset:offset], c.l8MMQ3mB9Vm, p, nil)
	incrementNonce(c.l8MMQ3mB9Vm)
	if _, e = c.eMYicxBI.Write(out); e != nil {
		return 0, e
	}
	return len(p), nil
}

// ReadFrom restores 0x1c9e640, including distinct per-call replay errors.
func (c *jSUbF0) ReadFrom(p []byte) (int, net.Addr, error) {
	n, addr, e := c.WCApOca6IQW7.ReadFrom(p)
	if e != nil {
		return 0, nil, e
	}
	size := max(16, c.obzsFnGuS.geqiu3e)
	if n < size {
		return 0, nil, fmt.Errorf("salt len not enough: %v", n)
	}
	if !A.HFlnXWty(p[:size]) {
		return 0, nil, errors.New("aead stream salt repeat")
	}
	if e = c.bh0EfZohV(p[:size]); e != nil {
		return 0, nil, e
	}
	copy(p, p[size:n])
	_, e = c.et6j454EdcWG.Open(p[:0], c.qaeiJJzh3hj, p[:n-size], nil)
	if e != nil {
		return 0, nil, e
	}
	return n - size - c.et6j454EdcWG.Overhead(), addr, nil
}

// WriteTo restores 0x1c9ea40. Its nonce does not advance on datagrams.
func (c *jSUbF0) WriteTo(p []byte, addr net.Addr) (int, error) {
	salt := make([]byte, max(16, c.obzsFnGuS.geqiu3e))
	rand.Read(salt)
	if e := c.n4dhtM(salt); e != nil {
		return 0, e
	}
	out, e := c.GetWriteBuf(len(salt) + len(p) + c.fsuDlw7T.Overhead())
	if e != nil {
		return 0, e
	}
	copy(out, salt)
	c.fsuDlw7T.Seal(out[len(salt):len(salt)], c.l8MMQ3mB9Vm, p, nil)
	if _, e = c.WCApOca6IQW7.WriteTo(out, addr); e != nil {
		return 0, e
	}
	return len(p), nil
}
