package FLhkKwJ8M

import (
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"recovered.local/soga/mzi1d3gDO7u"
	"sync"
)

type tojW0z34 struct{ io.Writer }
type vV27Q3k struct {
	io.Reader
	dphuptdYtW   []byte
	gOoPyboUld   []byte
	aauT0eoLd0b0 int
}
type sBT2tDxX3LZ struct {
	io.Writer
	cipher.AEAD
	qhhbVzCk   [32]byte
	hEJ1_Tabe1 uint16
	xNF1Gb47   []byte
	gLcBElmaW  sync.Mutex
}
type se83rHogAI7 struct {
	io.Reader
	cipher.AEAD
	gC5rR2aCAj   [32]byte
	aiE3W5V      []byte
	cz1edAAYy    int
	frnZkWRc     []byte
	jYxJbABejRtF []byte
	zIlPJ70bRN   uint16
}

// Write: 0x19ad820. Original ignores a short write with a nil error.
func (w *tojW0z34) Write(p []byte) (n int, err error) {
	b := mzi1d3gDO7u.SharedPoolAt5096978.Get(20480)
	defer mzi1d3gDO7u.SharedPoolAt5096978.Put(b)
	for n < len(p) {
		size := min(16384, len(p)-n)
		copy(b[2:16386], p[n:n+size])
		binary.BigEndian.PutUint16(b, uint16(size))
		if _, err = w.Writer.Write(b[:size+2]); err != nil {
			return
		}
		n += size
	}
	return
}

// Read: 0x19ad460. Partial body reads return zero and discard partial bytes.
func (r *vV27Q3k) Read(p []byte) (int, error) {
	if r.dphuptdYtW != nil {
		n := copy(p, r.dphuptdYtW[r.aauT0eoLd0b0:])
		r.aauT0eoLd0b0 += n
		if r.aauT0eoLd0b0 == len(r.dphuptdYtW) {
			mzi1d3gDO7u.SharedPoolAt5096978.Put(r.dphuptdYtW)
			r.dphuptdYtW = nil
		}
		return n, nil
	}
	if _, err := io.ReadFull(r.Reader, r.gOoPyboUld); err != nil {
		return 0, err
	}
	size := int(binary.BigEndian.Uint16(r.gOoPyboUld))
	if size > 17408 {
		return 0, errors.New("buffer is larger than standard")
	}
	if len(p) >= size {
		_, err := io.ReadFull(r.Reader, p[:size])
		if err != nil {
			return 0, err
		}
		return size, nil
	}
	b := mzi1d3gDO7u.SharedPoolAt5096978.Get(size)
	if _, err := io.ReadFull(r.Reader, b); err != nil {
		mzi1d3gDO7u.SharedPoolAt5096978.Put(b)
		return 0, err
	}
	n := copy(p, b)
	r.aauT0eoLd0b0 = n
	r.dphuptdYtW = b
	return n, nil
}

// Write: 0x19abbc0. Counter increments before the underlying write, including
// writes returning an error. Nonce bytes beyond offset 11 retain their values.
func (w *sBT2tDxX3LZ) Write(p []byte) (n int, err error) {
	w.gLcBElmaW.Lock()
	b := mzi1d3gDO7u.SharedPoolAt5096978.Get(20480)
	defer func() { w.gLcBElmaW.Unlock(); mzi1d3gDO7u.SharedPoolAt5096978.Put(b) }()
	for n < len(p) {
		size := min(16384-w.Overhead(), len(p)-n)
		copy(b[2:16386-w.Overhead()], p[n:n+size])
		binary.BigEndian.PutUint16(b, uint16(size+w.Overhead()))
		binary.BigEndian.PutUint16(w.qhhbVzCk[:], w.hEJ1_Tabe1)
		copy(w.qhhbVzCk[2:12], w.xNF1Gb47[2:12])
		w.Seal(b[2:2], w.qhhbVzCk[:w.NonceSize()], b[2:2+size], nil)
		w.hEJ1_Tabe1++
		if _, err = w.Writer.Write(b[:size+w.Overhead()+2]); err != nil {
			return
		}
		n += size
	}
	return
}

// Read: 0x19ac060. The original does not return the pooled allocation when
// authentication fails, and increments its nonce counter even on that error.
func (r *se83rHogAI7) Read(p []byte) (int, error) {
	if r.aiE3W5V != nil {
		n := copy(p, r.aiE3W5V[r.cz1edAAYy:])
		r.cz1edAAYy += n
		if r.cz1edAAYy == len(r.aiE3W5V) {
			mzi1d3gDO7u.SharedPoolAt5096978.Put(r.aiE3W5V)
			r.aiE3W5V = nil
		}
		return n, nil
	}
	if _, err := io.ReadFull(r.Reader, r.jYxJbABejRtF); err != nil {
		return 0, err
	}
	size := int(binary.BigEndian.Uint16(r.jYxJbABejRtF))
	if size > 17408 {
		return 0, errors.New("buffer is larger than standard")
	}
	b := mzi1d3gDO7u.SharedPoolAt5096978.Get(size)
	if _, err := io.ReadFull(r.Reader, b[:size]); err != nil {
		mzi1d3gDO7u.SharedPoolAt5096978.Put(b)
		return 0, err
	}
	binary.BigEndian.PutUint16(r.gC5rR2aCAj[:], r.zIlPJ70bRN)
	copy(r.gC5rR2aCAj[2:12], r.frnZkWRc[2:12])
	_, err := r.Open(b[:0], r.gC5rR2aCAj[:r.NonceSize()], b[:size], nil)
	r.zIlPJ70bRN++
	if err != nil {
		return 0, err
	}
	plain := b[:size-r.Overhead()]
	n := copy(p, plain)
	if len(p) < len(plain) {
		r.cz1edAAYy = n
		r.aiE3W5V = plain
	} else {
		mzi1d3gDO7u.SharedPoolAt5096978.Put(b)
	}
	return n, nil
}
