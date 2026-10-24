package bQxeORT

import (
	"crypto/ecdh"
	"crypto/mlkem"
	"errors"
	"sync"
	"time"

	"lukechampine.com/blake3"
)

// 原样本RTTI的服务端224字节状态和72字节会话；字段及布局保持原名。
// 这是依据原指令重建的Go源码表达，不是作者文件原文。
type LgF3OB1QX struct {
	EiKprMhvVvc  []any
	SpB4gP6aCP   [][]byte
	CPktYs15Wt   [][32]byte
	DptoPRgY     int
	F9GbshEXgaiY uint32
	MOy_iE1vq    int64
	OIHUERfuz    int64
	EA3t4sK      [][3]int
	PBHCqS       [][3]int
	VgRY0_5S9F37 sync.RWMutex
	JNEGgibSOG   bool
	GPXD9SMy     map[int64][16]byte
	CpNeohe      [][16]byte
	GIHdzaFGhi8  map[[16]byte]*MZKKvEA81HV
}
type MZKKvEA81HV struct {
	MwUsryYRwlpa []byte
	VaNnCL       sync.Map
}

// Init @0x1cfb320；失败也保留已写入的初始化状态，不能自动回滚重试。
func (s *LgF3OB1QX) Init(keys [][]byte, mode uint32, minimum, maximum int64, padding string) error {
	if s.EiKprMhvVvc != nil {
		return errors.New("already initialized")
	}
	if len(keys) == 0 {
		return errors.New("empty nfsSKeysBytes")
	}
	s.EiKprMhvVvc = make([]any, len(keys))
	s.SpB4gP6aCP = make([][]byte, len(keys))
	s.CPktYs15Wt = make([][32]byte, len(keys))
	for i, key := range keys {
		if len(key) == 32 {
			private, err := ecdh.X25519().NewPrivateKey(key)
			s.EiKprMhvVvc[i] = private
			if err != nil {
				return err
			}
			s.SpB4gP6aCP[i] = private.PublicKey().Bytes()
			s.DptoPRgY += 64
		} else {
			private, err := mlkem.NewDecapsulationKey768(key)
			s.EiKprMhvVvc[i] = private
			if err != nil {
				return err
			}
			s.SpB4gP6aCP[i] = private.EncapsulationKey().Bytes()
			s.DptoPRgY += 1120
		}
		digest := blake3.Sum512(s.SpB4gP6aCP[i])
		copy(s.CPktYs15Wt[i][:], digest[:32])
	}
	s.DptoPRgY -= 32
	s.F9GbshEXgaiY = mode
	s.MOy_iE1vq, s.OIHUERfuz = minimum, maximum
	if err := TQNZfKT(padding, &s.EA3t4sK, &s.PBHCqS); err != nil {
		return err
	}
	if minimum > 0 || maximum > 0 {
		s.GPXD9SMy = make(map[int64][16]byte)
		s.CpNeohe = make([][16]byte, 0, 1024)
		s.GIHdzaFGhi8 = make(map[[16]byte]*MZKKvEA81HV)
		go func() { // 原Init.func1 @0x1d01a80，完整ASM恢复。
			for {
				time.Sleep(time.Minute)
				s.VgRY0_5S9F37.Lock()
				if s.JNEGgibSOG {
					s.VgRY0_5S9F37.Unlock()
					return
				}
				minute := time.Now().Unix() / 60
				last := s.GPXD9SMy[minute]
				delete(s.GPXD9SMy, minute)
				delete(s.GPXD9SMy, minute-1)
				if last != ([16]byte{}) {
					for i, id := range s.CpNeohe {
						delete(s.GIHdzaFGhi8, id)
						if id == last {
							s.CpNeohe = s.CpNeohe[i+1:]
							break
						}
					}
				}
				s.VgRY0_5S9F37.Unlock()
			}
		}()
	}
	return nil
}

func (s *LgF3OB1QX) Close() error { // @0x1cfbc80
	s.VgRY0_5S9F37.Lock()
	s.JNEGgibSOG = true
	s.VgRY0_5S9F37.Unlock()
	return nil
}
