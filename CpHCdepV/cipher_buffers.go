package CpHCdepV

import (
	D "recovered.local/soga/D_zCag1qTV"
	E "recovered.local/soga/HxKViEp2a"
	"sync/atomic"
)

// JQi6241 follows original RTTI 0x3654ce8 (24 bytes). Its atomic flag guards
// released ownership; the read/write buffer pointers are not synchronized.
type JQi6241 struct {
	jqW714u *D.X_aXX5
	tvnW2P  *D.X_aXX5
	csU3Cr  atomic.Bool
}

// GetReadBuf restores 0x18bc920.
func (b *JQi6241) GetReadBuf(size int) ([]byte, error) {
	buf, err := b.GetReadBuffer(size)
	if err != nil {
		return nil, err
	}
	return buf.BytesTo(size), nil
}

// GetWriteBuf restores 0x18bc9c0.
func (b *JQi6241) GetWriteBuf(size int) ([]byte, error) {
	buf, err := b.GetWriteBuffer(size)
	if err != nil {
		return nil, err
	}
	return buf.BytesTo(size), nil
}

// GetReadBuffer restores 0x18bca60. Reuse always resets both cursors.
func (b *JQi6241) GetReadBuffer(size int) (*D.X_aXX5, error) {
	if b.csU3Cr.Load() {
		return nil, E.ErrorAt5097810
	}
	if b.jqW714u == nil {
		b.jqW714u = D.S4OQXWIZ(size)
	}
	b.jqW714u.Reset()
	b.jqW714u.GrowTo(size)
	return b.jqW714u, nil
}

// GetReadBufferDirect restores 0x18bcb20 and retains existing cursors.
func (b *JQi6241) GetReadBufferDirect() (*D.X_aXX5, error) {
	if b.csU3Cr.Load() {
		return nil, E.ErrorAt5097810
	}
	if b.jqW714u == nil {
		b.jqW714u = D.S4OQXWIZ(1024)
	}
	return b.jqW714u, nil
}

// GetWriteBuffer restores 0x18bcbc0.
func (b *JQi6241) GetWriteBuffer(size int) (*D.X_aXX5, error) {
	if b.csU3Cr.Load() {
		return nil, E.ErrorAt5097810
	}
	if b.tvnW2P == nil {
		b.tvnW2P = D.S4OQXWIZ(size)
	}
	b.tvnW2P.Reset()
	b.tvnW2P.GrowTo(size)
	return b.tvnW2P, nil
}

// FreeReadBuf/FreeReadBuffer restore 0x18bcc80/0x18bcd00.
func (b *JQi6241) FreeReadBuf() {
	if b.jqW714u != nil {
		b.jqW714u.Free()
		b.jqW714u = nil
	}
}
func (b *JQi6241) FreeReadBuffer() {
	if b.jqW714u != nil {
		b.jqW714u.Free()
		b.jqW714u = nil
	}
}

// FreeWriteBuf/FreeWriteBuffer restore 0x18bcd80/0x18bce00.
func (b *JQi6241) FreeWriteBuf() {
	if b.tvnW2P != nil {
		b.tvnW2P.Free()
		b.tvnW2P = nil
	}
}
func (b *JQi6241) FreeWriteBuffer() {
	if b.tvnW2P != nil {
		b.tvnW2P.Free()
		b.tvnW2P = nil
	}
}

// Free restores 0x18bce80. The original unconditional atomic Store does not
// make concurrent buffer use/free safe; sequential repeated Free is harmless.
func (b *JQi6241) Free() { b.csU3Cr.Store(true); b.FreeReadBuf(); b.FreeWriteBuf() }
