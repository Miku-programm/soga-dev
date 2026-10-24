//go:build linux

package app

import "recovered.local/soga/panel/registry"

// Original main entries, independently traced for completion and timeout:
// ppanel 0x1fb95a0, soga-v1 0x1fbd300, sspanel-uim 0x1fc1b60.
// All nine receiving closures consume one value and spawn one attempt.
func aGvHzoD3fr(count int) { runPanelStarterCallback(count, ppanelCallbackAttempt) }
func jQGVEC1W(count int)   { runPanelStarterCallback(count, sogaV1CallbackAttempt) }
func joTwyrx(count int)    { runPanelStarterCallback(count, sspanelCallbackAttempt) }

// Each of the nine independent original attempt bodies was checked for low,
// high, retry, final error and second-comparison error paths (45 traces).
func ppanelCallbackAttempt(value *int, queue chan int, done chan struct{}) {
	starterCallbackAttempt(value, queue, done, fUu8qv6z, registry.AJlKiH)
}
func sogaV1CallbackAttempt(value *int, queue chan int, done chan struct{}) {
	starterCallbackAttempt(value, queue, done, mW8q8iAOtl, registry.IfXGMpuMHP)
}
func sspanelCallbackAttempt(value *int, queue chan int, done chan struct{}) {
	starterCallbackAttempt(value, queue, done, uAIwuL_g, registry.R4ZXq34Lnd)
}

// The three original thresholds 0x1fb8100, 0x1fbb520 and 0x1fbe060
// have distinct obfuscation states but the same arithmetic. All 32 successive
// dividend/divisor pairs were checked with zero, FF and ascending entropy,
// plus empty input and both clock parity branches. Random reads are retained.
func fUu8qv6z(input []byte) int   { return h68JK8g(input) }
func mW8q8iAOtl(input []byte) int { return h68JK8g(input) }
func uAIwuL_g(input []byte) int   { return h68JK8g(input) }
