package xx5fRyrFw

import (
	"recovered.local/soga/G7A7csgrs62p"
	"strconv"
	"strings"
)

type Z5Jdu9 struct {
	cI3Z5P7PZ *G7A7csgrs62p.Fw_lgKgam_jg[uint16]
}

// SekQr0GHk restores 0xebe3e0, removing an optional case-sensitive port: prefix.
func SekQr0GHk(patterns []string) *Z5Jdu9 {
	m := &Z5Jdu9{cI3Z5P7PZ: G7A7csgrs62p.PQQKui[uint16]()}
	for _, p := range patterns {
		m.iF3QxW1j2A(strings.TrimPrefix(p, "port:"))
	}
	return m
}

// iF3QxW1j2A restores 0xebe540. Reversed ranges are accepted, invalid tokens
// are silently skipped, and whitespace is not trimmed.
func (m *Z5Jdu9) iF3QxW1j2A(pattern string) {
	if strings.Contains(pattern, ",") {
		for _, p := range strings.Split(pattern, ",") {
			m.iF3QxW1j2A(p)
		}
		return
	}
	if strings.Contains(pattern, "-") {
		bounds := strings.SplitN(pattern, "-", 2)
		if len(bounds) != 2 {
			return
		}
		first, err := strconv.Atoi(bounds[0])
		if err != nil {
			return
		}
		last, err := strconv.Atoi(bounds[1])
		if err != nil {
			return
		}
		if last < first {
			first, last = last, first
		}
		for port := first; port <= last; port++ {
			if port >= 1 && port <= 65535 {
				m.cI3Z5P7PZ.Add(uint16(port))
			}
		}
		return
	}
	port, err := strconv.Atoi(pattern)
	if err == nil && port >= 1 && port <= 65535 {
		m.cI3Z5P7PZ.Add(uint16(port))
	}
}

// Match restores 0xebe780; range rejection precedes dereferencing the receiver.
func (m *Z5Jdu9) Match(port int) bool {
	if port < 1 || port > 65535 {
		return false
	}
	return m.cI3Z5P7PZ.Has(uint16(port))
}
