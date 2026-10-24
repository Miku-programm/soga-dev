package vtQf23A9G

import (
	"encoding/binary"
	D "recovered.local/soga/D_zCag1qTV"
	"slices"
	"time"
)

// hOmlzcDoI @0x18a4bc0；原长度字段直接转 uint16，不拆分超长数据。
func (s *BIkgBMz3M) hOmlzcDoI(id uint32, p []byte) (int, error) {
	b := D.LdHEmSs9(len(p) + 7)
	b.WriteByte(2)
	b.WriteUint32(id)
	b.WriteUint16(uint16(len(p)))
	b.Write(p)
	_, err := s.ucYLRp(b.Bytes(), 0)
	b.Free()
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
func (s *BIkgBMz3M) ehs06TiU(cmd byte, id uint32, p []byte) (int, error) { // 0x18a4e60
	b := D.LdHEmSs9(len(p) + 7)
	b.WriteByte(cmd)
	b.WriteUint32(id)
	b.WriteUint16(uint16(len(p)))
	b.Write(p)
	_, err := s.ucYLRp(b.Bytes(), 0)
	b.Free()
	if err != nil {
		s.Close()
		return 0, err
	}
	return len(p), nil
}

// ucYLRp @0x18a5120：保留原记录计数、填充范围、部分写入计数和延迟 flush。
// 拼接内联含溢出 panic("len out of range") 与 nil 起点 Grow，逐分支对应 slices.Concat。
func (s *BIkgBMz3M) ucYLRp(data []byte, timeout time.Duration) (int, error) {
	s.a5L1EeV4yvaw.Lock()
	defer s.a5L1EeV4yvaw.Unlock()
	if s.d0Ze8qja {
		s.vV_K36 = slices.Concat(s.vV_K36, data)
		return len(data), nil
	}
	if len(s.vV_K36) > 0 {
		data = slices.Concat(s.vV_K36, data)
		s.vV_K36 = nil
	}
	if timeout > 0 {
		s.c740FZf.SetWriteDeadline(time.Now().Add(timeout))
		defer s.c740FZf.SetWriteDeadline(time.Time{})
	}
	if s.rVguB66 {
		record := s.wLYsfs.Add(1)
		scheme := s.bmjW7a.Load()
		if record < scheme.ZJzf4Os {
			n := 0
			for _, size := range scheme.GenerateRecordPayloadSizes(record) {
				if size == -1 {
					if len(data) == 0 {
						break
					}
					continue
				}
				if len(data) > size {
					if _, err := s.c740FZf.Write(data[:size]); err != nil {
						return 0, err
					}
					n += size
					data = data[size:]
				} else {
					if len(data) > 0 {
						original := len(data)
						gap := size - len(data)
						if gap-7 > 0 {
							pad := make([]byte, gap)
							binary.BigEndian.PutUint16(pad[5:], uint16(gap-7))
							data = slices.Concat(data, pad)
						}
						if _, err := s.c740FZf.Write(data); err != nil {
							return 0, err
						}
						n += original
					} else {
						pad := make([]byte, size+7)
						binary.BigEndian.PutUint16(pad[5:], uint16(size))
						if _, err := s.c740FZf.Write(pad); err != nil {
							return 0, err
						}
					}
					data = nil
				}
			}
			if len(data) == 0 {
				return n, nil
			}
			written, err := s.c740FZf.Write(data)
			return n + written, err
		}
		s.rVguB66 = false
	}
	return s.c740FZf.Write(data)
}
