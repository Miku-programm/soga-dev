package realitykdf

import "sync"

// JW2SKH2ajwT6 retains the original 32-byte layout (type descriptor 0x36b2300).
type JW2SKH2ajwT6 struct {
	aM9TY50Pn   map[string]int
	x8InevskM   map[string]int
	rl2GEL      map[string]struct{}
	gRUHgk0Uufx sync.Mutex
}

// NewJW2SKH2ajwT6 names the constructor inlined in user.cLLKJ_A31 at
// 0x1825460. The binary contains no separate constructor symbol.
func NewJW2SKH2ajwT6() *JW2SKH2ajwT6 {
	return &JW2SKH2ajwT6{
		aM9TY50Pn: make(map[string]int),
		x8InevskM: make(map[string]int),
		rl2GEL:    make(map[string]struct{}),
	}
}

// AddCount recovers 0xec0960. The limit applies to distinct first keys;
// lowering it below the current count rejects even an existing first key.
func (a *JW2SKH2ajwT6) AddCount(first, second string, limit int) bool {
	a.gRUHgk0Uufx.Lock()
	defer a.gRUHgk0Uufx.Unlock()
	if limit > 0 {
		count := len(a.aM9TY50Pn)
		if limit == count {
			if _, exists := a.aM9TY50Pn[first]; !exists {
				return false
			}
		} else if limit < count {
			return false
		}
	}
	a.aM9TY50Pn[first]++
	a.x8InevskM[second]++
	return true
}

// SubCount recovers 0xec0b60. Each key is decremented independently.
func (a *JW2SKH2ajwT6) SubCount(first, second string) {
	a.gRUHgk0Uufx.Lock()
	defer a.gRUHgk0Uufx.Unlock()
	if count, exists := a.aM9TY50Pn[first]; exists {
		if count <= 1 {
			delete(a.aM9TY50Pn, first)
		} else {
			a.aM9TY50Pn[first]--
		}
	}
	if count, exists := a.x8InevskM[second]; exists {
		if count <= 1 {
			delete(a.x8InevskM, second)
		} else {
			a.x8InevskM[second]--
		}
	}
}

// GetAndClearAliveIps recovers 0xec0d20. A disconnected recorded IP is
// included one final time; connected recorded IPs survive subsequent reads.
func (a *JW2SKH2ajwT6) GetAndClearAliveIps() []string {
	a.gRUHgk0Uufx.Lock()
	defer a.gRUHgk0Uufx.Unlock()
	result := make([]string, 0, len(a.rl2GEL))
	var remove []string
	for ip := range a.rl2GEL {
		result = append(result, ip)
		if _, exists := a.x8InevskM[ip]; !exists {
			remove = append(remove, ip)
		}
	}
	for _, ip := range remove {
		delete(a.rl2GEL, ip)
	}
	return result
}

// Record recovers 0xec1160. Counting a connection does not itself record it.
func (a *JW2SKH2ajwT6) Record(ip string) {
	a.gRUHgk0Uufx.Lock()
	defer a.gRUHgk0Uufx.Unlock()
	a.rl2GEL[ip] = struct{}{}
}
