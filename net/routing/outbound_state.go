package routing

// uymza6wOIV6 restores 0x1b7dbc0: initial enablement happens once.
func (out *GyyO51) uymza6wOIV6() {
	out.mJTqKai.Lock()
	defer out.mJTqKai.Unlock()
	if !out.vWopFdjg {
		out.fSEvC0BgQ = true
		out.vWopFdjg = true
	}
}

// IsEnabled restores 0x1b7dc80.
func (out *GyyO51) IsEnabled() bool {
	out.mJTqKai.RLock()
	defer out.mJTqKai.RUnlock()
	return out.fSEvC0BgQ
}

// ebDTwNX restores 0x1b7dd40. Results report disabled/enabled transitions.
// Three consecutive failures disable; two consecutive successes re-enable.
func (out *GyyO51) ebDTwNX(success bool) (bool, bool) {
	out.mJTqKai.Lock()
	defer out.mJTqKai.Unlock()
	if success {
		out.pJpTed = 0
		if out.fSEvC0BgQ {
			out.qoxZCsWf = 0
			return false, false
		}
		out.qoxZCsWf++
		if out.qoxZCsWf >= 2 {
			out.fSEvC0BgQ = true
			out.qoxZCsWf = 0
			return false, true
		}
		return false, false
	}
	out.qoxZCsWf = 0
	if !out.fSEvC0BgQ {
		return false, false
	}
	out.pJpTed++
	if out.pJpTed >= 3 {
		out.fSEvC0BgQ = false
		out.pJpTed = 0
		return true, false
	}
	return false, false
}

// IsDirect restores 0x1b804a0; comparison is case-sensitive.
func (out *GyyO51) IsDirect() bool { return out.NyLbiXb4OrlL == "direct" }

// Enabled restores 0x1b804e0, retaining order and a non-nil empty result.
func (outs J_enG0wE) Enabled() J_enG0wE {
	result := make(J_enG0wE, 0, len(outs))
	for _, out := range outs {
		if out.IsEnabled() {
			result = append(result, out)
		}
	}
	return result
}
