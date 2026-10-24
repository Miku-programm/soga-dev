package ssrprotocol

import (
	CA "recovered.local/soga/util/rtx"
	S "recovered.local/soga/protocol/ssrshared"
	"slices"
)

// fhBiTh.GWIGI0SUzOOZ at 0x780280 is the original int-slice pdqsort
// wrapper: slices.pdqsortOrdered[int](s,0,len(s),bits.Len(uint(len(s)))).
// These call sites only need its ascending integer ordering.

// kEHh4m5 restores 0x1cd8f00. Repeated initialization appends to both lists.
func (c *yeGSuIuC) kEHh4m5(key []byte) {
	var r S.HelGdo2ETK
	r.InitFromBytes(key)
	for n := int(r.Next()%8) + 4; n > 0; n-- {
		c.fbupx9aL = append(c.fbupx9aL, int(r.Next()%2340%2040%1440))
	}
	slices.Sort(c.fbupx9aL)
	for n := int(r.Next()%16) + 8; n > 0; n-- {
		c.r24E2R = append(c.r24E2R, int(r.Next()%2340%2040%1440))
	}
	slices.Sort(c.r24E2R)
}

// SetServerInfo restores 0x1cd9400.
func (c *yeGSuIuC) SetServerInfo(info *S.L7E7H2aW8KVs) {
	c.oHHJQje = info
	c.kEHh4m5(info.IawCjSkuTM4)
}

// z0O7jemcj restores 0x1cd9480, including draws from the full list lengths
// rather than only from the suffix returned by the binary search.
func (c *yeGSuIuC) z0O7jemcj(length int, seed []byte, r *S.HelGdo2ETK) int {
	if length > 1440 {
		return 0
	}
	r.InitFromBytesLen(seed, length)
	target := length + c.oHHJQje.CJfUaAqL
	p, _ := slices.BinarySearch(c.fbupx9aL, target)
	i := p + int(r.Next()%uint64(len(c.fbupx9aL)))
	if i < len(c.fbupx9aL) {
		return c.fbupx9aL[i] - target
	}
	p, _ = slices.BinarySearch(c.r24E2R, target)
	i = p + int(r.Next()%uint64(len(c.r24E2R)))
	if i < len(c.r24E2R) {
		return c.r24E2R[i] - target
	}
	if i < p+len(c.r24E2R)-1 {
		return 0
	}
	n := r.Next()
	switch {
	case length > 1300:
		return int(n % 31)
	case length > 900:
		return int(n % 127)
	case length > 400:
		return int(n % 521)
	default:
		return int(n % 1021)
	}
}

// kEHh4m5 restores 0x1cd9c20. Its reset is a non-nil empty slice.
func (c *j8lPExkSUzBF) kEHh4m5(key []byte) {
	c.fbupx9aL = []int{}
	var r S.HelGdo2ETK
	r.InitFromBytes(key)
	for n := int(r.Next()%24) + 12; n > 0; n-- {
		c.fbupx9aL = append(c.fbupx9aL, int(r.Next()%2340%2040%1440))
	}
	slices.Sort(c.fbupx9aL)
}

// SetServerInfo restores 0x1cd9fa0.
func (c *j8lPExkSUzBF) SetServerInfo(info *S.L7E7H2aW8KVs) {
	c.oHHJQje = info
	c.kEHh4m5(info.IawCjSkuTM4)
}

// z0O7jemcj restores 0x1cda020.
func (c *j8lPExkSUzBF) z0O7jemcj(length int, seed []byte, r *S.HelGdo2ETK) int {
	r.InitFromBytesLen(seed, length)
	target := length + c.oHHJQje.CJfUaAqL
	if c.fbupx9aL[len(c.fbupx9aL)-1] <= target {
		if target >= 1440 {
			return 0
		}
		n := r.Next()
		switch {
		case target > 1300:
			return int(n % 31)
		case target > 900:
			return int(n % 127)
		case target > 400:
			return int(n % 521)
		default:
			return int(n % 1021)
		}
	}
	p, _ := slices.BinarySearch(c.fbupx9aL, target)
	i := p + int(r.Next()%uint64(len(c.fbupx9aL)-p))
	return c.fbupx9aL[i] - target
}

// bzvPBLtdys99 restores 0x1cda820. The loop inspects the last appended
// value; it does not re-sort or inspect the maximum until it returns.
func (c *gdcp8ElhzO) bzvPBLtdys99(r *S.HelGdo2ETK) {
	for n := len(c.fbupx9aL); c.fbupx9aL[n-1] < 1300 && n < 64; n++ {
		c.fbupx9aL = append(c.fbupx9aL, int(r.Next()%2340%2040%1440))
	}
}

// kEHh4m5 restores 0x1cdaa00.
func (c *gdcp8ElhzO) kEHh4m5(key []byte) {
	c.fbupx9aL = []int{}
	var r S.HelGdo2ETK
	r.InitFromBytes(key)
	for n := int(r.Next()%24) + 12; n > 0; n-- {
		c.fbupx9aL = append(c.fbupx9aL, int(r.Next()%2340%2040%1440))
	}
	slices.Sort(c.fbupx9aL)
	n := len(c.fbupx9aL)
	c.bzvPBLtdys99(&r)
	if len(c.fbupx9aL) != n {
		slices.Sort(c.fbupx9aL)
	}
}

// SetServerInfo restores 0x1cdade0.
func (c *gdcp8ElhzO) SetServerInfo(info *S.L7E7H2aW8KVs) {
	c.oHHJQje = info
	c.kEHh4m5(info.IawCjSkuTM4)
}

// z0O7jemcj restores 0x1cdae60. Unlike C and E, this early return does
// not seed or advance the supplied PRNG.
func (c *gdcp8ElhzO) z0O7jemcj(length int, seed []byte, r *S.HelGdo2ETK) int {
	target := length + c.oHHJQje.CJfUaAqL
	if c.fbupx9aL[len(c.fbupx9aL)-1] <= target {
		return 0
	}
	r.InitFromBytesLen(seed, length)
	p, _ := slices.BinarySearch(c.fbupx9aL, target)
	i := p + int(r.Next()%uint64(len(c.fbupx9aL)-p))
	return c.fbupx9aL[i] - target
}

// z0O7jemcj restores 0x1cdb7c0. The seed is initialized but Next is not
// called; E selects the first qualifying entry deterministically.
func (c *yMywJL) z0O7jemcj(length int, seed []byte, r *S.HelGdo2ETK) int {
	r.InitFromBytesLen(seed, length)
	target := length + c.oHHJQje.CJfUaAqL
	if c.fbupx9aL[len(c.fbupx9aL)-1] <= target {
		return 0
	}
	p, _ := slices.BinarySearch(c.fbupx9aL, target)
	return c.fbupx9aL[p] - target
}

// SetServerInfo restores 0x1cdbdc0. The original truncates the synchronized
// Unix time to uint32 before division, and appends eight big-endian bytes.
func (c *pMrxs3_Z) SetServerInfo(info *S.L7E7H2aW8KVs) {
	c.oHHJQje = info
	period := (uint64(uint32(CA.C5Jff5vs0())) >> 1) * 3257812231 >> 47
	c.ztf_pH0rlj_E = []byte{}
	for i := 7; i >= 0; i-- {
		c.ztf_pH0rlj_E = append(c.ztf_pH0rlj_E, byte(period>>uint(i*8)))
	}
	c.kEHh4m5(info.IawCjSkuTM4)
}

// kEHh4m5 restores 0x1cdbf40. A private key copy is XORed with the eight
// period bytes. A key shorter than eight bytes panics before PRNG seeding.
func (c *pMrxs3_Z) kEHh4m5(key []byte) {
	c.fbupx9aL = []int{}
	var r S.HelGdo2ETK
	seed := make([]byte, len(key))
	copy(seed, key)
	for i := 0; i < 8; i++ {
		seed[i] ^= c.ztf_pH0rlj_E[i]
	}
	r.InitFromBytes(seed)
	for n := int(r.Next()%24) + 12; n > 0; n-- {
		c.fbupx9aL = append(c.fbupx9aL, int(r.Next()%2340%2040%1440))
	}
	slices.Sort(c.fbupx9aL)
	n := len(c.fbupx9aL)
	c.bzvPBLtdys99(&r)
	if len(c.fbupx9aL) != n {
		slices.Sort(c.fbupx9aL)
	}
}
