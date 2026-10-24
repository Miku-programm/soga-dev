package xx5fRyrFw

import "strings"

// DaCngxGKN: RTTI 0x36524e0 fields are three consecutive bools.
type DaCngxGKN struct{ zO7zz1DN, b2xC9d, mKTiEDp1QvQM bool }

// XSe2ktMsZnSQ restores 0xebe800; kJalfXv.Hej79BYe8@0x606740 is
// strings.ToLower. Comma-separated items are not whitespace-trimmed.
func XSe2ktMsZnSQ(pattern string) *DaCngxGKN {
	m := new(DaCngxGKN)
	for _, p := range strings.Split(strings.ToLower(pattern), ",") {
		switch p {
		case "tcp":
			m.zO7zz1DN = true
		case "udp":
			m.b2xC9d = true
		}
	}
	m.mKTiEDp1QvQM = !m.zO7zz1DN && !m.b2xC9d
	return m
}

// Match restores 0xebe8e0: nil/no recognized tokens allow all protocols.
func (m *DaCngxGKN) Match(protocol string) bool {
	if m == nil || m.mKTiEDp1QvQM {
		return true
	}
	switch protocol {
	case "tcp":
		return m.zO7zz1DN
	case "udp":
		return m.b2xC9d
	}
	return false
}
