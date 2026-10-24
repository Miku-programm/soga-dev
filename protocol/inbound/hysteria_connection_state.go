package inbound

import (
	"os"
	"time"

	"recovered.local/soga/config"
	"recovered.local/soga/crypto/rand"
)

// sbbdX3aT6E (0x1dad3a0) and bj10P7X (0x1dae100) are independently
// decoded copies of the hourly check. The original subtraction is uint32;
// backwards time and wraparound retain their original arithmetic. A failed
// CAS is not retried. The four Hysteria closures and both methods have 168
// independent machine traces (hysteria-state-all-traces.json); their response
// effects agree with connectionStateResponse. The Hysteria request entry is vLhTYZafwY.
// Neither the user count nor the failure count is atomic.
func (v *Aq5i0Bwez8) sbbdX3aT6E() {
	userCountAt50C6DA0 = v.fUxO7UYAH
}

func (v *Aq5i0Bwez8) bj10P7X() {
	userCountAt50C6DA0 = v.fUxO7UYAH
}

func (v *Aq5i0Bwez8) checkHysteriaConnectionState() {
	time.Sleep(time.Duration(rand.Co6D_zfls(60)) * time.Second)
	connectionStateResponse(&v.mKB02yhVfTfP, config.FrFQQhzTR5fY,
		vLhTYZafwY, config.T15Jifo, os.RemoveAll, os.Remove, os.Exit)
}
