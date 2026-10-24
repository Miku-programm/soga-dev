package u3UOJmirww7

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"io"
	"net"
	pool "recovered.local/soga/mzi1d3gDO7u"
)

type DqUAnLs6a struct {
	io.Writer
	cipher.AEAD
	pk3MztlNaO7 [32]byte
}
type Gh18ibyrqaX struct {
	io.Reader
	cipher.AEAD
	fOP3Pl9EegxZ [32]byte
	sZSEFoeE     []byte
	rx6vUmUYPW   int
}
type Zh37aRKMK struct {
	net.PacketConn
	HugVbq
}
type AZVSCxZmZQ struct {
	net.Conn
	HugVbq
	i24UCJDW   *Gh18ibyrqaX
	dApbxu3E0Q *DqUAnLs6a
}

// zeroNonceAt50C8760 is the original 128-byte package global.
var zeroNonceAt50C8760 [128]byte

// SJwXfg414: 0x1945500. Salt generation precedes the destination-length check.
func SJwXfg414(dst, plaintext []byte, c HugVbq) ([]byte, error) {
	salt := dst[:c.SaltSize()]
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	a, err := c.Encrypter(salt)
	if err != nil {
		return nil, err
	}
	if len(dst) < len(salt)+len(plaintext)+a.Overhead() {
		return nil, io.ErrShortBuffer
	}
	out := a.Seal(dst[len(salt):len(salt)], zeroNonceAt50C8760[:a.NonceSize()], plaintext, nil)
	return dst[:len(salt)+len(out)], nil
}

// ED88t50zm: 0x19457c0.
func ED88t50zm(dst, packet []byte, c HugVbq) ([]byte, error) {
	size := c.SaltSize()
	if len(packet) < size {
		return nil, errorAt5098530
	}
	a, err := c.Decrypter(packet[:size])
	if err != nil {
		return nil, err
	}
	if len(packet) < size+a.Overhead() {
		return nil, errorAt5098530
	}
	if len(packet) > size+a.Overhead()+len(dst) {
		return nil, io.ErrShortBuffer
	}
	return a.Open(dst[:0], zeroNonceAt50C8760[:a.NonceSize()], packet[size:], nil)
}

// WriteTo: 0x1945a80. Count is plaintext length even on the underlying error.
func (c *Zh37aRKMK) WriteTo(p []byte, addr net.Addr) (int, error) {
	b := pool.SharedPoolAt5096978.Get(65536)
	defer pool.SharedPoolAt5096978.Put(b)
	b, err := SJwXfg414(b, p, c)
	if err != nil {
		return 0, err
	}
	_, err = c.PacketConn.WriteTo(b, addr)
	return len(p), err
}

// ReadFrom: 0x1945c80. On decryption failure retain the encrypted count/address.
func (c *Zh37aRKMK) ReadFrom(p []byte) (int, net.Addr, error) {
	n, addr, err := c.PacketConn.ReadFrom(p)
	if err != nil {
		return n, addr, err
	}
	out, err := ED88t50zm(p[c.SaltSize():], p[:n], c)
	if err != nil {
		return n, addr, err
	}
	copy(p, out)
	return len(out), addr, nil
}

// incrementNonce is the identical little-endian loop in each original call site.
func incrementNonce(n []byte) {
	for i := range n {
		n[i]++
		if n[i] != 0 {
			break
		}
	}
}

// Write: 0x1946440. A zero input emits only an encrypted length field.
func (w *DqUAnLs6a) Write(p []byte) (n int, err error) {
	b := pool.SharedPoolAt5096978.Get(17408)
	defer pool.SharedPoolAt5096978.Put(b)
	nonce := w.pk3MztlNaO7[:w.NonceSize()]
	overhead := w.Overhead()
	if len(p) == 0 {
		b = b[:2+overhead]
		b[0] = 0
		b[1] = 0
		w.Seal(b[:0], nonce, b[:2], nil)
		incrementNonce(nonce)
		_, err = w.Writer.Write(b)
		return
	}
	for n < len(p) && err == nil {
		size := min(16383, len(p)-n)
		b = b[:size+2*overhead+2]
		binary.BigEndian.PutUint16(b, uint16(size))
		w.Seal(b[:0], nonce, b[:2], nil)
		incrementNonce(nonce)
		w.Seal(b[:2+overhead], nonce, p[n:n+size], nil)
		incrementNonce(nonce)
		_, err = w.Writer.Write(b)
		n += size
	}
	return
}

// ReadFrom: 0x19469a0. Even a terminal zero-byte read emits two AEAD records.
func (w *DqUAnLs6a) ReadFrom(r io.Reader) (total int64, err error) {
	b := pool.SharedPoolAt5096978.Get(17408)
	defer pool.SharedPoolAt5096978.Put(b)
	nonce := w.pk3MztlNaO7[:w.NonceSize()]
	overhead := w.Overhead()
	offset := 2 + overhead
	for {
		n, re := r.Read(b[offset : offset+16383])
		total += int64(n)
		binary.BigEndian.PutUint16(b, uint16(n))
		w.Seal(b[:0], nonce, b[:2], nil)
		incrementNonce(nonce)
		w.Seal(b[:offset], nonce, b[offset:offset+n], nil)
		incrementNonce(nonce)
		_, err = w.Writer.Write(b[:n+2*overhead+2])
		if err != nil {
			return
		}
		if re != nil {
			if re != io.EOF {
				err = re
			}
			return
		}
	}
}

// ux1mIly: 0x1946e80. Length flags are masked; nonces advance after Open even
// when authentication fails, but not when ReadFull fails.
func (r *Gh18ibyrqaX) ux1mIly(b []byte) (int, error) {
	nonce := r.fOP3Pl9EegxZ[:r.NonceSize()]
	overhead := r.Overhead()
	header := b[:2+overhead]
	if _, err := io.ReadFull(r.Reader, header); err != nil {
		return 0, err
	}
	_, err := r.Open(header[:0], nonce, header, nil)
	incrementNonce(nonce)
	if err != nil {
		return 0, err
	}
	n := int(binary.BigEndian.Uint16(header) & 16383)
	if n == 0 {
		return 0, errorAt5098540
	}
	data := b[:n+overhead]
	if _, err = io.ReadFull(r.Reader, data); err != nil {
		return 0, err
	}
	_, err = r.Open(data[:0], nonce, data, nil)
	incrementNonce(nonce)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Read: 0x19471c0. The error path deliberately does not return its allocated
// scratch slice to the original pool; zero-length calls still consume a chunk.
func (r *Gh18ibyrqaX) Read(p []byte) (int, error) {
	if r.sZSEFoeE == nil {
		if len(p) >= 16383+r.Overhead() {
			return r.ux1mIly(p)
		}
		b := pool.SharedPoolAt5096978.Get(17408)
		n, err := r.ux1mIly(b)
		if err != nil {
			return 0, err
		}
		r.sZSEFoeE = b[:n]
		r.rx6vUmUYPW = 0
	}
	n := copy(p, r.sZSEFoeE[r.rx6vUmUYPW:])
	r.rx6vUmUYPW += n
	if r.rx6vUmUYPW == len(r.sZSEFoeE) {
		pool.SharedPoolAt5096978.Put(r.sZSEFoeE[:cap(r.sZSEFoeE)])
		r.sZSEFoeE = nil
	}
	return n, nil
}

// WriteTo: 0x19473e0. Keep the scratch buffer on input EOF/error; preserve
// unwritten data on writer errors. There is no synthetic ErrShortWrite check.
func (r *Gh18ibyrqaX) WriteTo(w io.Writer) (total int64, err error) {
	if r.sZSEFoeE == nil {
		r.sZSEFoeE = pool.SharedPoolAt5096978.Get(17408)
		r.rx6vUmUYPW = len(r.sZSEFoeE)
	}
	for {
		if r.rx6vUmUYPW < len(r.sZSEFoeE) {
			n, e := w.Write(r.sZSEFoeE[r.rx6vUmUYPW:])
			r.rx6vUmUYPW += n
			total += int64(n)
			if e != nil {
				if r.rx6vUmUYPW == len(r.sZSEFoeE) {
					pool.SharedPoolAt5096978.Put(r.sZSEFoeE[:cap(r.sZSEFoeE)])
					r.sZSEFoeE = nil
				}
				return total, e
			}
			continue
		}
		n, e := r.ux1mIly(r.sZSEFoeE)
		if e != nil {
			if e == io.EOF {
				e = nil
			}
			return total, e
		}
		r.sZSEFoeE = r.sZSEFoeE[:n]
		r.rx6vUmUYPW = 0
	}
}

// eXuhbtvdbc: 0x1947600; dERlZM7: 0x1947960.
func (c *AZVSCxZmZQ) eXuhbtvdbc() error {
	salt := make([]byte, c.SaltSize())
	if _, err := io.ReadFull(c.Conn, salt); err != nil {
		return err
	}
	a, err := c.Decrypter(salt)
	if err != nil {
		return err
	}
	c.i24UCJDW = &Gh18ibyrqaX{Reader: c.Conn, AEAD: a}
	return nil
}
func (c *AZVSCxZmZQ) dERlZM7() error {
	salt := make([]byte, c.SaltSize())
	if _, err := rand.Read(salt); err != nil {
		return err
	}
	a, err := c.Encrypter(salt)
	if err != nil {
		return err
	}
	if _, err = c.Conn.Write(salt); err != nil {
		return err
	}
	c.dApbxu3E0Q = &DqUAnLs6a{Writer: c.Conn, AEAD: a}
	return nil
}

// Read/WriteTo/Write/ReadFrom: 0x1947820/19478c0/1947b40/1947be0.
func (c *AZVSCxZmZQ) Read(p []byte) (int, error) {
	if c.i24UCJDW == nil {
		if e := c.eXuhbtvdbc(); e != nil {
			return 0, e
		}
	}
	return c.i24UCJDW.Read(p)
}
func (c *AZVSCxZmZQ) WriteTo(w io.Writer) (int64, error) {
	if c.i24UCJDW == nil {
		if e := c.eXuhbtvdbc(); e != nil {
			return 0, e
		}
	}
	return c.i24UCJDW.WriteTo(w)
}
func (c *AZVSCxZmZQ) Write(p []byte) (int, error) {
	if c.dApbxu3E0Q == nil {
		if e := c.dERlZM7(); e != nil {
			return 0, e
		}
	}
	return c.dApbxu3E0Q.Write(p)
}
func (c *AZVSCxZmZQ) ReadFrom(r io.Reader) (int64, error) {
	if c.dApbxu3E0Q == nil {
		if e := c.dERlZM7(); e != nil {
			return 0, e
		}
	}
	return c.dApbxu3E0Q.ReadFrom(r)
}
