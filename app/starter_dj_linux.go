//go:build linux

package app

import (
	"context"
	Q "recovered.local/soga/panel/registry"
	"time"
)

// AiB7DQuj is the original 8-byte dj-specific wrapper, not an alias for Wrdr.
type AiB7DQuj struct{ *WrdrARwuhuha }

// yo9Lzz2t @0x1f994e0 allocates both original objects.
func yo9Lzz2t(callback _Zpu8_d) G6miUObWbEE {
	return &AiB7DQuj{WrdrARwuhuha: &WrdrARwuhuha{zqoWpLRIH: callback}}
}

// avZC_rQmNC @0x1f995a0 starts the optional status task before delegating.
// A typed nil panel still satisfies its interface assertion, as in the original.
func (s *AiB7DQuj) avZC_rQmNC(ctx context.Context, panel Q.VyFk5uT) {
	if status, ok := panel.(Q.Nz4DV7cPH); ok {
		go djSubmitStatusLoop(ctx, status)
	}
	s.WrdrARwuhuha.avZC_rQmNC(ctx, panel)
}

// Start @0x1f99740, original goroutine wrapper @0x1ff3ce0.
func (s *AiB7DQuj) Start(ctx context.Context, panel Q.VyFk5uT) { go s.avZC_rQmNC(ctx, panel) }

// Original closure @0x1ff3b80 and defer @0x1ff3ca0: one interval read,
// first report only after the first tick, errors ignored, ticker stopped on exit.
// Nonpositive intervals retain time.NewTicker's panic, even with canceled ctx.
func djSubmitStatusLoop(ctx context.Context, panel Q.Nz4DV7cPH) {
	ticker := time.NewTicker(panel.GetSubmitInterval())
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			_ = panel.SubmitStatus()
		case <-ctx.Done():
			return
		}
	}
}
