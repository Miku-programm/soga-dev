package bQxeORT

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// IaLb4jaWtI @ 0x1cfb240 samples [min,max), returning an equal bound directly.
func IaLb4jaWtI(a, b int) int {
	if a == b {
		return a
	}
	if b < a {
		a, b = b, a
	}
	v, _ := rand.Int(rand.Reader, big.NewInt(int64(b-a)))
	return a + int(v.Int64())
}

// TQNZfKT @ 0x1cf84a0 appends alternating size and delay rules. Errors retain
// rules appended earlier in the parse; extra hyphen-separated fields are ignored.
func TQNZfKT(spec string, sizes, delays *[][3]int) error {
	if spec == "" {
		return nil
	}
	total := 0
	for index, segment := range strings.Split(spec, ".") {
		parts := strings.Split(segment, "-")
		if len(parts) < 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
			return errors.New("invalid padding lenth/gap parameter: " + segment)
		}
		var rule [3]int
		for i := range rule {
			v, err := strconv.Atoi(parts[i])
			if err != nil {
				return err
			}
			rule[i] = v
		}
		if index == 0 && (rule[0] < 100 || rule[1] < 35 || rule[2] < 35) {
			return errors.New("first padding length must not be smaller than 35")
		}
		if index&1 != 0 {
			*delays = append(*delays, rule)
		} else {
			*sizes = append(*sizes, rule)
			total += max(rule[1], rule[2])
		}
	}
	if total > 65553 {
		return errors.New("total padding length must not be larger than 65553")
	}
	return nil
}

// OODk_yfADN @ 0x1cf8a60 draws padding sizes and millisecond delays. The
// probability comparison is inclusive; a zero threshold still matches draw 0.
func OODk_yfADN(sizes, delays [][3]int) (int, []int, []time.Duration) {
	if len(sizes) == 0 {
		sizes = [][3]int{{100, 111, 1111}, {50, 0, 3333}}
		delays = [][3]int{{75, 0, 111}}
	}
	var lengths []int
	total := 0
	for _, r := range sizes {
		n := 0
		if r[0] >= IaLb4jaWtI(0, 100) {
			n = IaLb4jaWtI(r[1], r[2])
		}
		lengths = append(lengths, n)
		total += n
	}
	var times []time.Duration
	for _, r := range delays {
		n := 0
		if r[0] >= IaLb4jaWtI(0, 100) {
			n = IaLb4jaWtI(r[1], r[2])
		}
		times = append(times, time.Duration(n)*time.Millisecond)
	}
	return total, lengths, times
}
