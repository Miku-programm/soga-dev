package nzATvajt

import (
	"bytes"
	"crypto/md5"
	"github.com/gofrs/uuid"
)

// DALfaCvQ is the original 33-byte VMess identity record (RTTI 0x3659080).
type DALfaCvQ struct {
	Yfd33z6YYdJa [16]byte
	dp2tOOgp     [16]byte
	kSqxkHOPkJ   bool
}
type IC_1xll []*DALfaCvQ

func w6MaVR(text string) (*DALfaCvQ, error) { // 0x1828540
	var id uuid.UUID
	if err := id.Parse(text); err != nil {
		return nil, err
	}
	return &DALfaCvQ{Yfd33z6YYdJa: [16]byte(id)}, nil
}

func (id *DALfaCvQ) GetCmdKey() []byte { // 0x1828600
	if id.kSqxkHOPkJ {
		return id.dp2tOOgp[:]
	}
	h := md5.New()
	h.Write(id.Yfd33z6YYdJa[:])
	h.Write([]byte("c48619fe-8f02-49e0-b9e9-edf763e17e21"))
	result := h.Sum(nil)
	copy(id.dp2tOOgp[:], result)
	id.kSqxkHOPkJ = true
	// The first call returns the allocated digest; later calls expose the cache.
	return result
}

func xshVaRu(previous, destination []byte) { // 0x1828c20
	h := md5.New()
	h.Write(previous)
	h.Write([]byte("16167dc8-16b6-4e6d-b8bb-65dd68113a81"))
	for {
		h.Sum(destination[:0])
		// Compare the original destination slice, including its original length.
		if !bytes.Equal(destination, previous) {
			return
		}
		h.Write([]byte("533eff8a-4113-4b10-b5ce-0f5d76b98cd2"))
	}
}
