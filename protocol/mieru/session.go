package mieru

import (
	"context"
	"fmt"
	"io"
	"net"
	D "recovered.local/soga/buf/pool"
	N "recovered.local/soga/net/limit"
	U "recovered.local/soga/user"
	"sync"
)

// WMK1ddta3da retains the original 96-byte authenticated connection state.
type WMK1ddta3da struct {
	fFpIe7      context.Context
	bl3UaWZlq   context.CancelFunc
	jcTl_f8Fmn  net.Conn
	pnt0Bzyf    *N.O0UldNL4Eq
	eJsvyMJ     U.PgaNd3014U
	djqn5YPkLN  *rZDq_kinV4G
	pZaZ0uHV2KJ *qhFcsZC
	aAwbr_m35W  *zXJm01kG7F4E
	qPLI0s2     chan *TZ3YCqW
}
type TZ3YCqW struct{ j_Ra2Blepbh *xjTVP8 }
type zXJm01kG7F4E struct {
	ocfcBS      sync.RWMutex
	w4gwnrQrrXj map[uint32]*xjTVP8
}

func (m *zXJm01kG7F4E) mS0qzOa(id uint32) (*xjTVP8, bool) { //0x1c58120
	m.ocfcBS.RLock()
	defer m.ocfcBS.RUnlock()
	s, ok := m.w4gwnrQrrXj[id]
	return s, ok
}
func (m *zXJm01kG7F4E) iptPyaqX(id uint32, s *xjTVP8) { //0x1c58220
	m.ocfcBS.Lock()
	defer m.ocfcBS.Unlock()
	m.w4gwnrQrrXj[id] = s
}
func (m *zXJm01kG7F4E) jJn3FR6HBm5A(id uint32) (*xjTVP8, bool) { //0x1c58300
	m.ocfcBS.Lock()
	defer m.ocfcBS.Unlock()
	s, ok := m.w4gwnrQrrXj[id]
	if ok {
		delete(m.w4gwnrQrrXj, id)
	}
	return s, ok
}
func (m *zXJm01kG7F4E) ihh5hvgSsIi() { //0x1c58420
	m.ocfcBS.Lock()
	old := m.w4gwnrQrrXj
	m.w4gwnrQrrXj = make(map[uint32]*xjTVP8)
	m.ocfcBS.Unlock()
	for _, s := range old {
		s.Close()
	}
}
func M_03miDCRcqa(ctx context.Context, conn net.Conn, resolver *N.O0UldNL4Eq) (*WMK1ddta3da, error) { //0x1c54f20
	ctx, cancel := context.WithCancel(ctx)
	s := &WMK1ddta3da{fFpIe7: ctx, bl3UaWZlq: cancel, jcTl_f8Fmn: conn, pnt0Bzyf: resolver, aAwbr_m35W: &zXJm01kG7F4E{w4gwnrQrrXj: make(map[uint32]*xjTVP8)}, qPLI0s2: make(chan *TZ3YCqW, 8)}
	if err := s.init(); err != nil {
		cancel()
		return nil, err
	}
	return s, nil
}
func (s *WMK1ddta3da) init() error { //0x1c55100
	reader, user, meta, err := cMGKnW(s.jcTl_f8Fmn, func(callback func(*U.WIZQ28) bool) {
		s.pnt0Bzyf.RacA3he0TX.Range(func(_ int64, user U.PgaNd3014U) bool {
			if u, ok := user.(*U.WIZQ28); ok {
				return callback(u)
			}
			return true
		})
	})
	if err != nil {
		return err
	}
	writer, err := vfhZtLu(s.jcTl_f8Fmn, reader.n49jm3JWYo, user.(*U.WIZQ28).GetUsername())
	if err != nil {
		return err
	}
	s.eJsvyMJ = user
	s.djqn5YPkLN = reader
	s.pZaZ0uHV2KJ = writer
	stop := context.AfterFunc(user.GetCtx(), func() { s.Close() })
	go s.dTznYF0U(meta, stop)
	return nil
}
func (s *WMK1ddta3da) Next() (*TZ3YCqW, error) { //0x1c55360
	select {
	case <-s.fFpIe7.Done():
		return nil, io.EOF
	case v, ok := <-s.qPLI0s2:
		if !ok {
			return nil, io.EOF
		}
		return v, nil
	}
}
func (s *WMK1ddta3da) Close() error { //0x1c55440
	s.bl3UaWZlq()
	s.aAwbr_m35W.ihh5hvgSsIi()
	s.jcTl_f8Fmn.Close()
	return nil
}
func (s *WMK1ddta3da) dTznYF0U(first *IqB1foHWg, stop func() bool) { //0x1c554a0, cleanup0x1c5d220
	defer func() { stop(); close(s.qPLI0s2); s.aAwbr_m35W.ihh5hvgSsIi(); s.Close() }()
	seg, err := s.djqn5YPkLN.jQ8yjlkt(first)
	if err != nil {
		return
	}
	if s.d8yPV78aE(seg) != nil {
		return
	}
	for {
		select {
		case <-s.fFpIe7.Done():
			return
		default:
		}
		seg, err = s.djqn5YPkLN.Read()
		if err != nil {
			return
		}
		if s.d8yPV78aE(seg) != nil {
			return
		}
	}
}
func (s *WMK1ddta3da) d8yPV78aE(seg *jjyc26) error { //0x1c55600
	switch seg.jaPY7RU.Rf6N7kY6GVCs {
	case 2:
		return s.z33kfUKtTYUZ(seg)
	case 4, 5:
		return s.qzQPqE1Bn4fD(seg)
	case 6:
		return s.bH_Ryqwc5(seg)
	case 10:
		s.pZaZ0uHV2KJ.clientLowEntropy.Store(true)
		return s.bH_Ryqwc5(seg)
	case 8:
		return nil
	default:
		return fmt.Errorf("unexpected protocol %v from client", seg.jaPY7RU.Rf6N7kY6GVCs.String())
	}
}

// No extra payload cleanup is added on duplicate/unknown-session paths.
func (s *WMK1ddta3da) z33kfUKtTYUZ(seg *jjyc26) error { //0x1c55780
	id := seg.jaPY7RU.HNaaDh
	if id == 0 {
		return fmt.Errorf("reserved session ID 0")
	}
	if _, exists := s.aAwbr_m35W.mS0qzOa(id); exists {
		return fmt.Errorf("duplicate session ID %d", id)
	}
	stream := vnqzugKQF1jX(id, s.eJsvyMJ, s.jcTl_f8Fmn, s.pZaZ0uHV2KJ)
	stream.auBuNh = s.fTCBPKK
	s.aAwbr_m35W.iptPyaqX(id, stream)
	if len(seg.mzrzt1uBj) > 0 {
		select {
		case stream.xfdnp9xA <- seg.mzrzt1uBj:
		case <-stream.zhIY4QEke:
			D.PoolAt50961D8.Put(seg.mzrzt1uBj)
		}
	}
	if err := s.pZaZ0uHV2KJ.WriteSessionSegment(&IqB1foHWg{Rf6N7kY6GVCs: 3, HNaaDh: id, Tt1OfazA: stream.qYZ2Uu.Add(1) - 1}); err != nil {
		return fmt.Errorf("write openSessionResponse: %w", err)
	}
	select {
	case s.qPLI0s2 <- &TZ3YCqW{stream}:
		return nil
	case <-s.fFpIe7.Done():
		return s.fFpIe7.Err()
	}
}
func (s *WMK1ddta3da) qzQPqE1Bn4fD(seg *jjyc26) error { //0x1c55d80
	if stream, ok := s.aAwbr_m35W.jJn3FR6HBm5A(seg.jaPY7RU.HNaaDh); ok {
		stream.Close()
		if seg.jaPY7RU.Rf6N7kY6GVCs == 4 {
			s.pZaZ0uHV2KJ.WriteSessionSegment(&IqB1foHWg{Rf6N7kY6GVCs: 5, HNaaDh: stream.xIwb7TaM, Tt1OfazA: stream.qYZ2Uu.Add(1) - 1})
		}
	}
	return nil
}
func (s *WMK1ddta3da) bH_Ryqwc5(seg *jjyc26) error { //0x1c55e60
	if stream, ok := s.aAwbr_m35W.mS0qzOa(seg.jaPY7RU.HNaaDh); ok {
		if len(seg.mzrzt1uBj) > 0 {
			select {
			case stream.xfdnp9xA <- seg.mzrzt1uBj:
			case <-stream.zhIY4QEke:
				D.PoolAt50961D8.Put(seg.mzrzt1uBj)
			}
		}
	} else {
		s.pZaZ0uHV2KJ.WriteSessionSegment(&IqB1foHWg{Rf6N7kY6GVCs: 4, HNaaDh: seg.jaPY7RU.HNaaDh, Tt1OfazA: seg.jaPY7RU.N01v7C})
	}
	return nil
}
func (s *WMK1ddta3da) fTCBPKK(stream *xjTVP8) { //0x1c56000
	if _, ok := s.aAwbr_m35W.jJn3FR6HBm5A(stream.xIwb7TaM); ok {
		s.pZaZ0uHV2KJ.WriteSessionSegment(&IqB1foHWg{Rf6N7kY6GVCs: 4, HNaaDh: stream.xIwb7TaM, Tt1OfazA: stream.qYZ2Uu.Add(1) - 1})
	}
}
