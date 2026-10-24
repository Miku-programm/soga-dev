package inbound

import (
	"os"
	"sync/atomic"
	"time"

	"recovered.local/soga/util/rtx"
	"recovered.local/soga/config"
	"recovered.local/soga/crypto/rand"
)

// sbbdX3aT6E (0x1debb20) and bj10P7X (0x1dec8a0) are independently
// decoded copies of the hourly check. The original subtraction is uint32;
// backwards time and wraparound retain their original arithmetic. A failed
// CAS is not retried. Neither the user count nor the failure count is atomic.
func (v *K4avFiYOPi) sbbdX3aT6E() {
	now := uint32(rtx.C5Jff5vs0())
	userCountAt50C6DA0 = v.fUxO7UYAH
	if claimConnectionStateCheck(&v.jfLH8mvq5e03, now) {
		go v.checkVMessConnectionState()
	}
}

func (v *K4avFiYOPi) bj10P7X() {
	now := uint32(rtx.C5Jff5vs0())
	userCountAt50C6DA0 = v.fUxO7UYAH
	if claimConnectionStateCheck(&v.jfLH8mvq5e03, now) {
		go v.checkVMessConnectionState()
	}
}

func claimConnectionStateCheck(last *atomic.Uint32, now uint32) bool {
	previous := last.Load()
	return now-previous > 3600 && last.CompareAndSwap(previous, now)
}

func (v *K4avFiYOPi) checkVMessConnectionState() {
	time.Sleep(time.Duration(rand.Co6D_zfls(60)) * time.Second)
	connectionStateResponse(&v.mKB02yhVfTfP, config.FrFQQhzTR5fY,
		bKJv7R6xU, config.T15Jifo, os.RemoveAll, os.Remove, os.Exit)
}

// connectionStateResponse is the common, independently traced control flow
// of 0x1eab260, 0x1eacea0, 0x1eaeae0 and 0x1eb0860 after their random sleep.
// Explicit effect arguments let tests verify even the original uninstall and
// exit paths without performing them. Production supplies the real functions.
func connectionStateResponse(counter *int, key func() string,
	check func() (string, error), setState func(bool), removeAll, remove func(string) error,
	exit func(int)) {
	if key() == "" {
		setState(true)
	}
	code, err := check()
	if err != nil {
		if key() != "" {
			*counter++
		}
		return
	}
	switch code {
	case "0":
		*counter = 0
	case "100001":
		setState(true)
		*counter = 0
	case "100002", "100005":
		*counter++
		exit(0)
	case "100003", "100004", "100007":
		*counter++
		if *counter > 10 {
			exit(0)
		}
	case "100006", "100008", "200001":
		_ = removeAll("/usr/local/soga/")
		_ = remove("/etc/systemd/system/soga.service")
		_ = remove("/usr/local/bin/soga")
		_ = remove("/usr/bin/soga")
		*counter += 100
		exit(0)
	default:
		*counter += 100
		exit(0)
	}
}
