package registry

import "time"

// Nz4DV7cPH is the optional panel interface at original RTTI 0x35e06a0,
// selected by the dj starter's type assertion descriptor at 0x502db00.
type Nz4DV7cPH interface {
	GetSubmitInterval() time.Duration
	SubmitStatus() error
}
