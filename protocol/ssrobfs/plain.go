package ssrobfs

import (
	"io"
	S "recovered.local/soga/protocol/ssrshared"
)

// Jakq78yIQ is the original four-method interface, RTTI 0x35fbde8.
type Jakq78yIQ interface {
	io.ReadWriter
	SetServerInfo(*S.L7E7H2aW8KVs)
	GetOverhead() int
}
type q1_ASanae22 func(io.ReadWriter) Jakq78yIQ

// rqkJYP_xJtF preserves the original 16-byte plain wrapper (RTTI 0x36426f8).
type rqkJYP_xJtF struct{ io.ReadWriter }

func aR_Feln8A36(rw io.ReadWriter) Jakq78yIQ { return &rqkJYP_xJtF{rw} } // 0x1cabfe0
// The original plain implementation intentionally ignores server information.
func (*rqkJYP_xJtF) SetServerInfo(*S.L7E7H2aW8KVs) {}           // 0x1cac080
func (*rqkJYP_xJtF) GetOverhead() int              { return 0 } // 0x1cac0a0
