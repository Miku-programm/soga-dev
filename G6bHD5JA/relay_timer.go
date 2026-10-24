package G6bHD5JA

import (
	"sync/atomic"
	"time"
)

// JJEIU0hAF retains the original 40-byte layout at RTTI 0x36dcea8.
type JJEIU0hAF struct {
	m_XMay1FFo7 atomic.Bool
	cmzS3I8m    atomic.Bool
	xkqY3siS    *time.Timer
	qbjO8iq     time.Duration
	rapKwb3eF   int
	lgyY48      func()
}

// QCkuttV: 0x159dbc0. Idle duration is divided into four checks.
func QCkuttV(duration time.Duration, callback func()) *JJEIU0hAF {
	t := &JJEIU0hAF{qbjO8iq: duration / 4, lgyY48: callback}
	if duration > 0 {
		t.xkqY3siS = time.AfterFunc(time.Duration(1<<63-1), t.a3rxvmzRzoE)
		t.xkqY3siS.Reset(t.qbjO8iq)
	}
	return t
}

// Stop: 0x159dd00. A timer with no underlying Timer leaves stopped false.
func (t *JJEIU0hAF) Stop() {
	if t.xkqY3siS != nil {
		t.cmzS3I8m.Store(true)
		t.xkqY3siS.Stop()
	}
}

// a3rxvmzRzoE: 0x159dd60. Activity resets the consecutive idle-check count.
func (t *JJEIU0hAF) a3rxvmzRzoE() {
	if t.cmzS3I8m.Load() {
		return
	}
	if t.m_XMay1FFo7.Swap(false) {
		t.rapKwb3eF = 0
	} else {
		t.rapKwb3eF++
		if t.rapKwb3eF >= 4 {
			t.lgyY48()
			return
		}
	}
	t.xkqY3siS.Reset(t.qbjO8iq)
	if t.cmzS3I8m.Load() {
		t.xkqY3siS.Stop()
	}
}
