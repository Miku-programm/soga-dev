package anytls

import (
	"io"
	"os"
	S "recovered.local/soga/protocol/anytlssettings"
	"time"
)

// 原 init @0x18a2440，只有环境值完全等于 "1" 时禁止客户端更新填充策略。
var debugPaddingAt50C6C92 = os.Getenv("CLIENT_DEBUG_PADDING_SCHEME") == "1"

func (s *BIkgBMz3M) Run() { // 0x18a24c0
	if s.pNuGhHxIQdo {
		settings := S.L4jWYITSyXGq{"v": "2", "client": S.ClientVersionAt5098D50, "padding-md5": s.bmjW7a.Load().GtWiNfVs7j}
		s.d0Ze8qja = true
		s.ehs06TiU(4, 0, settings.ToBytes())
		go s.wbBfTO632dea()
	} else {
		s.wbBfTO632dea()
	}
}
func (s *BIkgBMz3M) IsClosed() bool { // 0x18a27a0
	select {
	case <-s.sF3ka3A2dYX:
		return true
	default:
		return false
	}
}
func (s *BIkgBMz3M) Close() error { // 0x18a2800 / 0x18a9180
	won := false
	s.pGT2EZ20.Do(func() { close(s.sF3ka3A2dYX); won = true })
	if !won {
		return io.ErrClosedPipe
	}
	if s.twJTEeJ0S != nil {
		s.twJTEeJ0S()
		s.twJTEeJ0S = nil
	}
	s.ukg_2xDnH9s.Lock()
	for _, stream := range s.mW8PvJYxYup9 {
		stream.mublR9bouj1Z()
	}
	s.mW8PvJYxYup9 = make(map[uint32]*CCXLCurYiT7)
	s.ukg_2xDnH9s.Unlock()
	return s.c740FZf.Close()
}
func (s *BIkgBMz3M) OpenStream() (*CCXLCurYiT7, error) { // 0x18a29a0
	if s.IsClosed() {
		return nil, io.ErrClosedPipe
	}
	id := s.iX5FOl.Add(1)
	stream := newStreamFromInline(s, id)
	if id >= 2 && s.gAJaebL >= 2 {
		s.wAiQ9IC2UZ.Lock()
		if s.c5NUXMoBI != nil {
			s.c5NUXMoBI()
		}
		s.c5NUXMoBI = S.LwNqe0FPzq(3*time.Second, func() { s.Close() })
		s.wAiQ9IC2UZ.Unlock()
	}
	if _, err := s.ehs06TiU(1, id, nil); err != nil {
		return nil, err
	}
	s.d0Ze8qja = false
	s.ukg_2xDnH9s.Lock()
	defer s.ukg_2xDnH9s.Unlock()
	if s.IsClosed() {
		return nil, io.ErrClosedPipe
	}
	s.mW8PvJYxYup9[id] = stream
	return stream, nil
}
func (s *BIkgBMz3M) ab0YmvFt(id uint32) error { // 0x18a4ae0
	if s.IsClosed() {
		return io.ErrClosedPipe
	}
	_, err := s.ehs06TiU(3, id, nil)
	s.ukg_2xDnH9s.Lock()
	delete(s.mW8PvJYxYup9, id)
	s.ukg_2xDnH9s.Unlock()
	return err
}
