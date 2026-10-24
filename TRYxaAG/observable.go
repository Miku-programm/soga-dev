package TRYxaAG

import "sync"

// Original channel RTTI0x35a1b28/0x35a1ad8: both receive-only interface channels.
type F11bqIbX2 <-chan any
type ZutmWj26aQLd <-chan any

// SaTXuENDW: original24-byte subscription; no registration method survives in
// the linked binary. Close is idempotent, but does not remove any map entry.
type SaTXuENDW struct {
	szJyRBo2e9   chan any
	dhjdMxyR0jOd sync.Once
}

func (s *SaTXuENDW) Close() { s.dhjdMxyR0jOd.Do(func() { close(s.szJyRBo2e9) }) } //0x1b31d20

// EkVnc6a: original32-byte observable RTTI0x36bd700.
type EkVnc6a struct {
	dRVDIEcad    F11bqIbX2
	jHrXT_T2TvUz map[ZutmWj26aQLd]*SaTXuENDW
	nd3U58QVfEal sync.Mutex
	zGiaD7v      bool
}

// AkUZ2QiKCu3n:0x1b31c20. Retain the supplied channel and start one receiver.
func AkUZ2QiKCu3n(input F11bqIbX2) *EkVnc6a {
	o := &EkVnc6a{dRVDIEcad: input, jHrXT_T2TvUz: make(map[ZutmWj26aQLd]*SaTXuENDW)}
	go o.jIaNOx3wV8()
	return o
}

// jIaNOx3wV8:0x1b31960. Sending is blocking and holds the map mutex. Original
// code has no panic recovery or select/default branch for slow subscribers.
func (o *EkVnc6a) jIaNOx3wV8() {
	for item := range o.dRVDIEcad {
		o.nd3U58QVfEal.Lock()
		for _, subscriber := range o.jHrXT_T2TvUz {
			subscriber.szJyRBo2e9 <- item
		}
		o.nd3U58QVfEal.Unlock()
	}
	o.fmz_nvCHJDK()
}

// fmz_nvCHJDK:0x1b31ae0. Closed subscriptions remain in the original map.
func (o *EkVnc6a) fmz_nvCHJDK() {
	o.nd3U58QVfEal.Lock()
	defer o.nd3U58QVfEal.Unlock()
	o.zGiaD7v = true
	for _, subscriber := range o.jHrXT_T2TvUz {
		subscriber.Close()
	}
}
