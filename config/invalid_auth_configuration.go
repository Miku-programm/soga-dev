package config

import "time"

// Nl62Lnd_ retains the original 32-byte RTTI at 0x36af780.
type Nl62Lnd_ struct {
	SEU_Z687ugp  bool
	GxuEkr       int
	QS255wlN7m2W time.Duration
	NjiRkFvidOaA time.Duration
}

// SSInvalidAccessSettings names the inline struct construction at 0x1d6d140.
// Only nonpositive values receive defaults, before nanosecond conversion.
func SSInvalidAccessSettings() *Nl62Lnd_ {
	c := &configAt50C8440
	count, duration, ban := c.UA4PZU, c.PWFghdU, c.SndqWD26h
	if count <= 0 {
		count = 30
	}
	if duration <= 0 {
		duration = 60
	}
	if ban <= 0 {
		ban = 600
	}
	return &Nl62Lnd_{SEU_Z687ugp: c.X4_AmA8j7, GxuEkr: count, QS255wlN7m2W: time.Duration(duration) * time.Second, NjiRkFvidOaA: time.Duration(ban) * time.Second}
}
