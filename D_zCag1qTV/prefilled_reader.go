package D_zCag1qTV

import "io"

type vF878Xs struct{}

func (*vF878Xs) Read([]byte) (int, error) { return 0, io.EOF } //0xe90840

// NewPrefilledForProxy is an explicitly named bridge for the original inline
// NADyavA literal in g0MlUOhvO.ARZKOqmN5e.ReadFrom@0x11da120. The input slice
// remains caller-owned: it is not allocated from the pool and must not FreeBuf.
func NewPrefilledForProxy(p []byte) *NADyavA {
	return &NADyavA{bkCmHugyPy: p, c6HUuPfxb526: &vF878Xs{}, inKjzRUMN3d: len(p), lS7kcz: -1, zyE274D49tb: -1}
}
