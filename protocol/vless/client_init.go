package vless

import (
	"crypto/ecdh"
	"crypto/mlkem"
	"errors"
	"lukechampine.com/blake3"
)

// Init recovers 0x1cf3320. Failed key/parameter parsing preserves the partially
// initialized fields, so a second Init also returns "already initialized".
func (s *GPFxUP0_) Init(keys [][]byte, mode, seconds uint32, padding string) error {
	if s.QUnOciyca6eb != nil {
		return errors.New("already initialized")
	}
	if len(keys) == 0 {
		return errors.New("empty nfsPKeysBytes")
	}
	s.QUnOciyca6eb = make([]any, len(keys))
	s.JTmPRiuw = keys
	s.J2battR3S = make([][32]byte, len(keys))
	for i, key := range keys {
		if len(key) == 32 {
			public, err := ecdh.X25519().NewPublicKey(key)
			s.QUnOciyca6eb[i] = public
			if err != nil {
				return err
			}
			s.JhDzI4ehWF += 64
		} else {
			public, err := mlkem.NewEncapsulationKey768(key)
			s.QUnOciyca6eb[i] = public
			if err != nil {
				return err
			}
			s.JhDzI4ehWF += 1120
		}
		digest := blake3.Sum512(key)
		copy(s.J2battR3S[i][:], digest[:32])
	}
	s.JhDzI4ehWF -= 32
	s.S799IXI = mode
	s.BsCV7uFJFW = seconds
	return TQNZfKT(padding, &s.Q8W_Oz, &s.PVif2PW)
}
