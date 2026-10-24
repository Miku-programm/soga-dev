package FLhkKwJ8M

import (
	"crypto/md5"
	"github.com/gofrs/uuid/v5"
)

// OMOYhoLsaabd: original 32-byte UUID/command-key pair.
type OMOYhoLsaabd struct {
	AJxI2iBn   *uuid.UUID
	P3zlkY2W1n []byte
}

// h7OyM4GsmHif: 0x19b80a0; keeps the supplied UUID pointer.
func h7OyM4GsmHif(id *uuid.UUID) *OMOYhoLsaabd {
	h := md5.New()
	h.Write(id[:])
	h.Write([]byte("c48619fe-8f02-49e0-b9e9-edf763e17e21"))
	p := &OMOYhoLsaabd{AJxI2iBn: id, P3zlkY2W1n: make([]byte, 16)}
	h.Sum(p.P3zlkY2W1n[:0])
	return p
}

// oYNYqO: 0x19b82e0.
func oYNYqO(previous *uuid.UUID) *uuid.UUID {
	h := md5.New()
	h.Write(previous[:])
	h.Write([]byte("16167dc8-16b6-4e6d-b8bb-65dd68113a81"))
	next := new(uuid.UUID)
	for {
		h.Sum(next[:0])
		if *next != *previous {
			return next
		}
		h.Write([]byte("533eff8a-4113-4b10-b5ce-0f5d76b98cd2"))
	}
}

// o_nBA2: 0x19b8500. Alter IDs precede the primary ID and share its command key.
func o_nBA2(primary *OMOYhoLsaabd, count uint16) []*OMOYhoLsaabd {
	ids := make([]*OMOYhoLsaabd, count)
	previous := primary.AJxI2iBn
	for i := range ids {
		next := oYNYqO(previous)
		ids[i] = &OMOYhoLsaabd{next, primary.P3zlkY2W1n}
		previous = next
	}
	return append(ids, primary)
}
