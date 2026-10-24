package vtQf23A9G

import (
	"io"
	"net"
	"os"
	"time"
)

func (s *CCXLCurYiT7) Read(b []byte) (int, error) { return s.k_6JEbXwYJY.Read(b) } // 0x18a8920
func (s *CCXLCurYiT7) Write(b []byte) (int, error) { // 0x18a89a0
	select {
	case <-s.ko_ou22dHrhf.Wait():
		return 0, os.ErrDeadlineExceeded
	default:
	}
	if s.yKgxtGJV != nil {
		return 0, s.yKgxtGJV
	}
	return s.hnbi8qavj6Y.hOmlzcDoI(s.y7z3pbAl, b)
}
func (s *CCXLCurYiT7) Close() error { return s.nolKeczqPhz(io.ErrClosedPipe) } // 0x18a8a80
func (s *CCXLCurYiT7) mublR9bouj1Z() { // 0x18a8ac0 / 0x18a9520；原 502A6C0 为 poll.errNetClosing。
	won := false
	s.yTV1rdHC.Do(func() { s.yKgxtGJV = net.ErrClosed; s.k_6JEbXwYJY.Close(); won = true })
	if won && s.kCaaEFymiBDr != nil {
		s.kCaaEFymiBDr()
		s.kCaaEFymiBDr = nil
	}
}
func (s *CCXLCurYiT7) nolKeczqPhz(err error) error { // 0x18a8b80 / 0x18a95a0
	won := false
	s.yTV1rdHC.Do(func() { s.yKgxtGJV = err; s.k_6JEbXwYJY.Close(); won = true })
	if !won {
		return s.yKgxtGJV
	}
	if s.kCaaEFymiBDr != nil {
		s.kCaaEFymiBDr()
		s.kCaaEFymiBDr = nil
	}
	return s.hnbi8qavj6Y.ab0YmvFt(s.y7z3pbAl)
}
func (s *CCXLCurYiT7) SetReadDeadline(t time.Time) error  { return s.k_6JEbXwYJY.SetReadDeadline(t) } // 0x18a8c80
func (s *CCXLCurYiT7) SetWriteDeadline(t time.Time) error { s.ko_ou22dHrhf.Set(t); return nil }       // 0x18a8ce0
func (s *CCXLCurYiT7) SetDeadline(t time.Time) error {
	s.ko_ou22dHrhf.Set(t)
	return s.k_6JEbXwYJY.SetReadDeadline(t)
} // 0x18a8d40
func (s *CCXLCurYiT7) LocalAddr() net.Addr {
	if c, ok := s.hnbi8qavj6Y.c740FZf.(interface{ LocalAddr() net.Addr }); ok {
		return c.LocalAddr()
	}
	return nil
} // 0x18a8de0
func (s *CCXLCurYiT7) RemoteAddr() net.Addr {
	if c, ok := s.hnbi8qavj6Y.c740FZf.(interface{ RemoteAddr() net.Addr }); ok {
		return c.RemoteAddr()
	}
	return nil
} // 0x18a8ea0
func (s *CCXLCurYiT7) HandshakeFailure(err error) error { // 0x18a8f60
	won := false
	s.bo7peua2l.Do(func() { won = true })
	if won && s.hnbi8qavj6Y.gAJaebL >= 2 && err != nil {
		_, e := s.hnbi8qavj6Y.ehs06TiU(7, s.y7z3pbAl, []byte(err.Error()))
		return e
	}
	return nil
}
func (s *CCXLCurYiT7) HandshakeSuccess() error { // 0x18a9080
	won := false
	s.bo7peua2l.Do(func() { won = true })
	if won && s.hnbi8qavj6Y.gAJaebL >= 2 {
		_, e := s.hnbi8qavj6Y.ehs06TiU(7, s.y7z3pbAl, nil)
		return e
	}
	return nil
}
