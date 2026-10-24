package vtQf23A9G

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	D "recovered.local/soga/D_zCag1qTV"
	H "recovered.local/soga/HHGbXWqYaYt_"
	S "recovered.local/soga/sn9h4GXYmaaA"
	"strconv"
)

// readPooledFromInline 仅合并原读取各分支重复的 Get/ReadFull/失败 Put。
func (s *BIkgBMz3M) readPooledFromInline(n int) ([]byte, error) {
	p := D.PoolAt50961D8.Get(n)
	if _, err := io.ReadFull(s.c740FZf, p); err != nil {
		D.PoolAt50961D8.Put(p)
		return nil, err
	}
	return p, nil
}

// wbBfTO632dea @0x18a2ec0：原 0..10 命令处理、设置协商、流生命周期。
// 原版对未知命令以及 SYN/FIN/PING/PONG 不消耗头部宣称的 payload，保持此行为。
func (s *BIkgBMz3M) wbBfTO632dea() error {
	defer s.Close()
	settingsSeen := false
	var header cNZriK
	for {
		if s.IsClosed() {
			return io.ErrClosedPipe
		}
		if _, err := io.ReadFull(s.c740FZf, header[:]); err != nil {
			return err
		}
		cmd := header[0]
		id := binary.BigEndian.Uint32(header[1:5])
		length := int(binary.BigEndian.Uint16(header[5:]))
		switch cmd {
		case 0:
			if length == 0 {
				continue
			}
			p, err := s.readPooledFromInline(length)
			if err != nil {
				return err
			}
			D.PoolAt50961D8.Put(p)
		case 1:
			if !s.pNuGhHxIQdo && !settingsSeen {
				s.ehs06TiU(5, 0, []byte("client did not send its settings"))
				return nil
			}
			s.ukg_2xDnH9s.Lock()
			if _, exists := s.mW8PvJYxYup9[id]; !exists {
				stream := newStreamFromInline(s, id)
				s.mW8PvJYxYup9[id] = stream
				go func() {
					if s.ulr87QdD != nil {
						s.ulr87QdD(stream)
					} else {
						stream.nolKeczqPhz(io.ErrClosedPipe)
					}
				}()
			}
			s.ukg_2xDnH9s.Unlock()
		case 2:
			if length == 0 {
				continue
			}
			p, err := s.readPooledFromInline(length)
			if err != nil {
				return err
			}
			s.ukg_2xDnH9s.RLock()
			stream, ok := s.mW8PvJYxYup9[id]
			s.ukg_2xDnH9s.RUnlock()
			if ok {
				stream.z8qCZQcE.Write(p)
			}
			D.PoolAt50961D8.Put(p)
		case 3:
			s.ukg_2xDnH9s.Lock()
			stream, ok := s.mW8PvJYxYup9[id]
			delete(s.mW8PvJYxYup9, id)
			s.ukg_2xDnH9s.Unlock()
			if ok {
				stream.mublR9bouj1Z()
			}
		case 4:
			if length == 0 {
				continue
			}
			p, err := s.readPooledFromInline(length)
			if err != nil {
				return err
			}
			if !s.pNuGhHxIQdo {
				settings := S.DsbUyzhj9(p)
				scheme := s.bmjW7a.Load()
				if settings["padding-md5"] != scheme.GtWiNfVs7j {
					if _, err := s.ehs06TiU(6, 0, scheme.KruzWnr9zxA); err != nil {
						D.PoolAt50961D8.Put(p)
						return err
					}
				}
				if v, err := strconv.Atoi(settings["v"]); err == nil && v >= 2 {
					s.gAJaebL = uint8(v)
					if _, err := s.ehs06TiU(10, 0, (S.L4jWYITSyXGq{"v": "2"}).ToBytes()); err != nil {
						D.PoolAt50961D8.Put(p)
						return err
					}
				}
				settingsSeen = true
			}
			D.PoolAt50961D8.Put(p)
		case 5:
			if length == 0 {
				continue
			}
			p, err := s.readPooledFromInline(length)
			if err != nil {
				return err
			}
			if s.pNuGhHxIQdo {
				logrus.StandardLogger().Logln(logrus.ErrorLevel, "[Alert from server]", string(p))
			}
			D.PoolAt50961D8.Put(p)
			return nil
		case 6:
			if length == 0 {
				continue
			}
			p := make([]byte, length)
			if _, err := io.ReadFull(s.c740FZf, p); err != nil {
				return err
			}
			if s.pNuGhHxIQdo && !debugPaddingAt50C6C92 {
				if H.TrBr1o4z(p) {
					logrus.StandardLogger().Logf(logrus.InfoLevel, "[Update padding succeed] %x\n", md5.Sum(p))
				} else {
					logrus.StandardLogger().Logf(logrus.WarnLevel, "[Update padding failed] %x\n", md5.Sum(p))
				}
			}
		case 7:
			s.wAiQ9IC2UZ.Lock()
			if s.c5NUXMoBI != nil {
				s.c5NUXMoBI()
				s.c5NUXMoBI = nil
			}
			s.wAiQ9IC2UZ.Unlock()
			if length == 0 {
				continue
			}
			p, err := s.readPooledFromInline(length)
			if err != nil {
				return err
			}
			s.ukg_2xDnH9s.RLock()
			stream, ok := s.mW8PvJYxYup9[id]
			s.ukg_2xDnH9s.RUnlock()
			if ok {
				stream.nolKeczqPhz(fmt.Errorf("remote: %s", string(p)))
			}
			D.PoolAt50961D8.Put(p)
		case 8:
			if _, err := s.ehs06TiU(9, id, nil); err != nil {
				return err
			}
		case 9:
		case 10:
			if length == 0 {
				continue
			}
			p, err := s.readPooledFromInline(length)
			if err != nil {
				return err
			}
			if s.pNuGhHxIQdo {
				if v, err := strconv.Atoi(S.DsbUyzhj9(p)["v"]); err == nil {
					s.gAJaebL = uint8(v)
				}
			}
			D.PoolAt50961D8.Put(p)
		}
	}
}
