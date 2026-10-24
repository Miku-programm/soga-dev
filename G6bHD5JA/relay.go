package G6bHD5JA

import (
	"context"
	"net"
	D "recovered.local/soga/D_zCag1qTV"
	H "recovered.local/soga/HxKViEp2a"
	U "recovered.local/soga/Ualx5UM_D6I"
	C "recovered.local/soga/cNT24rG4F"
	"sync"
	"time"
)

type EHIIYP interface {
	net.Conn
	Free()
	ReadBuffer() (U.Mf3phSOVqgf, error)
	WriteBuffer(U.Mf3phSOVqgf) (int, error)
}

// AsVABr is the original 96-byte transfer argument.
type AsVABr struct {
	C8FYQVoBTa   context.Context
	GdPSmtGs_1   net.Conn
	HEfD13ahR__b net.Conn
	C93uQCY9TN   []byte
	Mw4pNzmFlqA  []byte
}

// FreeBuf: 0x15aed60. Caller-provided buffers are returned to the shared pool.
func (s *AsVABr) FreeBuf() {
	D.PoolAt50961D8.Put(s.C93uQCY9TN)
	D.PoolAt50961D8.Put(s.Mw4pNzmFlqA)
	s.C93uQCY9TN = nil
	s.Mw4pNzmFlqA = nil
}

// Fields are analyst-named; 0x15aeca0 uses three consecutive machine integers.
type eN9twxaAav struct{ size, full, small int }

func (s *eN9twxaAav) aIjQSD(n int) int {
	if n == s.size && s.size < 16384 {
		s.small = 0
		s.full++
		if s.full >= 2 {
			s.full = 0
			s.size = min(s.size*2, 16384)
		}
	} else if n >= s.size/2 || s.size <= 4096 {
		s.full = 0
		s.small = 0
	} else {
		s.full = 0
		s.small++
		if s.small >= 16 {
			s.small = 0
			s.size = max(s.size/2, 4096)
		}
	}
	return s.size
}

// j_HzamVxRg: 0x15aee20. Deadline errors are deliberately ignored.
func j_HzamVxRg(conn net.Conn, next *time.Time) {
	now := time.Now()
	if now.After(*next) {
		duration := C.WdJLtcGXr()
		_ = conn.SetWriteDeadline(now.Add(duration))
		*next = now.Add(duration / 2)
	}
}

// JWBWWZS: 0x15af260. Partial writes without errors remain accepted by the
// original. Only internally allocated buffers use adaptive resizing.
func JWBWWZS(ctx context.Context, src, dst net.Conn, buf []byte, idle *JJEIU0hAF) (err error) {
	owned := buf == nil
	sizing := eN9twxaAav{size: 4096}
	if owned {
		buf = D.PoolAt50961D8.Get(4096)
		defer func() { D.PoolAt50961D8.Put(buf) }()
	}
	defer func() { H.SlG_Y_mg("pipeline") }()
	var next time.Time
	for {
		select {
		case <-ctx.Done():
			return H.ErrorAt5097800
		default:
		}
		n, readErr := src.Read(buf)
		if n > 0 {
			idle.m_XMay1FFo7.Store(true)
			j_HzamVxRg(dst, &next)
			_, writeErr := dst.Write(buf[:n])
			if writeErr != nil {
				return H.T2i1ItMaa(readErr, writeErr)
			}
			if owned {
				if size := sizing.aIjQSD(n); size != len(buf) {
					D.PoolAt50961D8.Put(buf)
					buf = D.PoolAt50961D8.Get(size)
				}
			}
		}
		if readErr != nil {
			return readErr
		}
	}
}

// FvtgLyQrWX: 0x15aef40. Close ordering and the non-deferred WaitGroup.Done
// follow the target, including its behavior if a Close call panics.
func FvtgLyQrWX(s AsVABr) {
	idle := QCkuttV(time.Duration(C.TCPTimeoutMinutes())*time.Minute, func() {
		defer H.SlG_Y_mg("relay idle")
		s.HEfD13ahR__b.Close()
		s.GdPSmtGs_1.Close()
	})
	defer idle.Stop()
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		defer H.SlG_Y_mg("relay")
		JWBWWZS(s.C8FYQVoBTa, s.GdPSmtGs_1, s.HEfD13ahR__b, s.C93uQCY9TN, idle)
		s.GdPSmtGs_1.Close()
		s.HEfD13ahR__b.Close()
		wait.Done()
	}()
	defer H.SlG_Y_mg("relay")
	JWBWWZS(s.C8FYQVoBTa, s.HEfD13ahR__b, s.GdPSmtGs_1, s.Mw4pNzmFlqA, idle)
	s.HEfD13ahR__b.Close()
	s.GdPSmtGs_1.Close()
	wait.Wait()
}

// eywOJGkx: 0x15b14e0. A buffer returned together with an error is ignored;
// the target does not Free buffers returned by ReadBuffer here.
func eywOJGkx(ctx context.Context, src EHIIYP, dst net.Conn, idle *JJEIU0hAF) {
	var next time.Time
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		buf, err := src.ReadBuffer()
		if err != nil {
			return
		}
		if buf.Len() > 0 {
			idle.m_XMay1FFo7.Store(true)
			j_HzamVxRg(dst, &next)
			if _, err = dst.Write(buf.Bytes()); err != nil {
				return
			}
		}
	}
}

// kMn_Sz0s: 0x15b1660. The reusable buffer may be replaced after each write.
func kMn_Sz0s(ctx context.Context, src net.Conn, dst EHIIYP, idle *JJEIU0hAF) {
	sizing := eN9twxaAav{size: 4096}
	size := 4096
	buf := D.S4OQXWIZ(size)
	defer func() { buf.Free() }()
	var next time.Time
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		buf.Reset()
		n, readErr := buf.ReadMaxFrom(src, size)
		if n > 0 {
			idle.m_XMay1FFo7.Store(true)
			j_HzamVxRg(dst, &next)
			if _, err := dst.WriteBuffer(buf); err != nil {
				return
			}
			if newSize := sizing.aIjQSD(n); newSize != size {
				buf.Free()
				buf = D.S4OQXWIZ(newSize)
				size = newSize
			}
		}
		if readErr != nil {
			return
		}
	}
}

// BKn06KsCTP: 0x15b11e0, the buffer-interface transfer path.
func BKn06KsCTP(ctx context.Context, src EHIIYP, dst net.Conn) {
	idle := QCkuttV(time.Duration(C.TCPTimeoutMinutes())*time.Minute, func() {
		defer H.SlG_Y_mg("relay idle")
		dst.Close()
		src.Close()
	})
	defer idle.Stop()
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		defer func() { H.SlG_Y_mg("relay") }()
		eywOJGkx(ctx, src, dst, idle)
		src.Close()
		dst.Close()
		wait.Done()
	}()
	defer func() { H.SlG_Y_mg("relay") }()
	kMn_Sz0s(ctx, dst, src, idle)
	dst.Close()
	src.Close()
	wait.Wait()
}
