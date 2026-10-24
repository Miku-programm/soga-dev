// 依据原 ELF xIF4G3 包机器码及 RTTI 恢复的源码表示，不是作者文件原文。
// 对照基线：apernet/hysteria 9e90d7d，保留其原许可；证据在 hysteria-wire。
package xIF4G3

import (
	"github.com/apernet/quic-go/congestion"
)

// packetNumberIndexedQueue is a queue of mostly continuous numbered entries
// which supports the following operations:
// - adding elements to the end of the queue, or at some point past the end
// - removing elements in any order
// - retrieving elements
// If all elements are inserted in order, all of the operations above are
// amortized O(1) time.
//
// Internally, the data structure is a deque where each element is marked as
// present or not.  The deque starts at the lowest present index.  Whenever an
// element is removed, it's marked as not present, and the front of the deque is
// cleared of elements that are not present.
//
// The tail of the queue is not cleared due to the assumption of entries being
// inserted in order, though removing all elements of the queue will return it
// to its initial state.
//
// Note that this data structure is inherently hazardous, since an addition of
// just two entries will cause it to consume all of the memory available.
// Because of that, it is not a general-purpose container and should not be used
// as one.

type tWtTcf[T any] struct {
	cPIy2wvomr bool
	ySYKCoJI   T
}

type h8SYlleQ[T any] struct {
	fSrleCn8X1  YXV6CyZ5Ux[tWtTcf[T]]
	bDBgyW      int
	f8sQndcVreq congestion.PacketNumber
}

func newPacketNumberIndexedQueue[T any](size int) *h8SYlleQ[T] {
	q := &h8SYlleQ[T]{
		f8sQndcVreq: invalidPacketNumber,
	}

	q.fSrleCn8X1.Init(size)

	return q
}

// Emplace inserts data associated |packet_number| into (or past) the end of the
// queue, filling up the missing intermediate entries as necessary.  Returns
// true if the element has been inserted successfully, false if it was already
// in the queue or inserted out of order.
func (p *h8SYlleQ[T]) Emplace(packetNumber congestion.PacketNumber, entry *T) bool {
	if packetNumber == invalidPacketNumber || entry == nil {
		return false
	}

	if p.IsEmpty() {
		p.fSrleCn8X1.PushBack(tWtTcf[T]{
			cPIy2wvomr: true,
			ySYKCoJI:   *entry,
		})
		p.bDBgyW = 1
		p.f8sQndcVreq = packetNumber
		return true
	}

	// Do not allow insertion out-of-order.
	if packetNumber <= p.LastPacket() {
		return false
	}

	// Handle potentially missing elements.
	offset := int(packetNumber - p.FirstPacket())
	if gap := offset - p.fSrleCn8X1.Len(); gap > 0 {
		for i := 0; i < gap; i++ {
			p.fSrleCn8X1.PushBack(tWtTcf[T]{})
		}
	}

	p.fSrleCn8X1.PushBack(tWtTcf[T]{
		cPIy2wvomr: true,
		ySYKCoJI:   *entry,
	})
	p.bDBgyW++
	return true
}

// GetEntry Retrieve the entry associated with the packet number.  Returns the pointer
// to the entry in case of success, or nullptr if the entry does not exist.
func (p *h8SYlleQ[T]) GetEntry(packetNumber congestion.PacketNumber) *T {
	ew := p.hNtkAeGD(packetNumber)
	if ew == nil {
		return nil
	}

	return &ew.ySYKCoJI
}

// Remove, Same as above, but if an entry is present in the queue, also call f(entry)
// before removing it.
func (p *h8SYlleQ[T]) Remove(packetNumber congestion.PacketNumber, f func(T)) bool {
	ew := p.hNtkAeGD(packetNumber)
	if ew == nil {
		return false
	}
	if f != nil {
		f(ew.ySYKCoJI)
	}
	ew.cPIy2wvomr = false
	p.bDBgyW--

	if packetNumber == p.FirstPacket() {
		p.iloj4aku()
	}

	return true
}

// RemoveUpTo, but not including |packet_number|.
// Unused slots in the front are also removed, which means when the function
// returns, |first_packet()| can be larger than |packet_number|.
func (p *h8SYlleQ[T]) RemoveUpTo(packetNumber congestion.PacketNumber) {
	for !p.fSrleCn8X1.Empty() &&
		p.f8sQndcVreq != invalidPacketNumber &&
		p.f8sQndcVreq < packetNumber {
		if p.fSrleCn8X1.Front().cPIy2wvomr {
			p.bDBgyW--
		}
		p.fSrleCn8X1.PopFront()
		p.f8sQndcVreq++
	}
	p.iloj4aku()

	return
}

// IsEmpty return if queue is empty.
func (p *h8SYlleQ[T]) IsEmpty() bool {
	return p.bDBgyW == 0
}

// NumberOfPresentEntries returns the number of entries in the queue.
func (p *h8SYlleQ[T]) NumberOfPresentEntries() int {
	return p.bDBgyW
}

// EntrySlotsUsed returns the number of entries allocated in the underlying deque.  This is
// proportional to the memory usage of the queue.
func (p *h8SYlleQ[T]) EntrySlotsUsed() int {
	return p.fSrleCn8X1.Len()
}

// LastPacket returns packet number of the first entry in the queue.
func (p *h8SYlleQ[T]) FirstPacket() (packetNumber congestion.PacketNumber) {
	return p.f8sQndcVreq
}

// LastPacket returns packet number of the last entry ever inserted in the queue.  Note that the
// entry in question may have already been removed.  Zero if the queue is
// empty.
func (p *h8SYlleQ[T]) LastPacket() (packetNumber congestion.PacketNumber) {
	if p.IsEmpty() {
		return invalidPacketNumber
	}

	return p.f8sQndcVreq + congestion.PacketNumber(p.fSrleCn8X1.Len()-1)
}

func (p *h8SYlleQ[T]) iloj4aku() {
	for !p.fSrleCn8X1.Empty() && !p.fSrleCn8X1.Front().cPIy2wvomr {
		p.fSrleCn8X1.PopFront()
		p.f8sQndcVreq++
	}
	if p.fSrleCn8X1.Empty() {
		p.f8sQndcVreq = invalidPacketNumber
	}
}

func (p *h8SYlleQ[T]) hNtkAeGD(packetNumber congestion.PacketNumber) *tWtTcf[T] {
	if packetNumber == invalidPacketNumber ||
		p.IsEmpty() ||
		packetNumber < p.f8sQndcVreq {
		return nil
	}

	offset := int(packetNumber - p.f8sQndcVreq)
	if offset >= p.fSrleCn8X1.Len() {
		return nil
	}

	ew := p.fSrleCn8X1.Offset(offset)
	if ew == nil || !ew.cPIy2wvomr {
		return nil
	}

	return ew
}
