package inbound

import (
	"os"
	"time"

	"recovered.local/soga/util/rtx"
	"recovered.local/soga/config"
	"recovered.local/soga/crypto/rand"
)

// sbbdX3aT6E (0x1de5f20) and bj10P7X (0x1de5180) are independently
// decoded copies of the hourly check. The original subtraction is uint32;
// backwards time and wraparound retain their original arithmetic. A failed
// CAS is not retried. The four SS closures and both methods have 180
// independent machine traces (ss-state-all-traces.json); their response
// effects agree with connectionStateResponse. The SS request entry is fzRe58.
// Neither the user count nor the failure count is atomic.
func (v *Lp07FzBLx9V) sbbdX3aT6E() {
	now := uint32(rtx.C5Jff5vs0())
	userCountAt50C6DA0 = v.fUxO7UYAH
	if claimConnectionStateCheck(&v.jfLH8mvq5e03, now) {
		go v.checkShadowsocksConnectionState()
	}
}

func (v *Lp07FzBLx9V) bj10P7X() {
	now := uint32(rtx.C5Jff5vs0())
	userCountAt50C6DA0 = v.fUxO7UYAH
	if claimConnectionStateCheck(&v.jfLH8mvq5e03, now) {
		go v.checkShadowsocksConnectionState()
	}
}

func (v *Lp07FzBLx9V) checkShadowsocksConnectionState() {
	time.Sleep(time.Duration(rand.Co6D_zfls(60)) * time.Second)
	connectionStateResponse(&v.mKB02yhVfTfP, config.FrFQQhzTR5fY,
		fzRe58, config.T15Jifo, os.RemoveAll, os.Remove, os.Exit)
}
