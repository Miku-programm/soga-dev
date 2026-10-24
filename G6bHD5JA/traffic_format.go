package G6bHD5JA

import "fmt"

// LTaVAEf restores 0x15b29a0, including the original successive binary scaling.
func LTaVAEf(n uint64) string {
	switch {
	case n < 1024:
		return fmt.Sprintf("%dB", n)
	case n < 1<<20:
		return fmt.Sprintf("%.2fKB", float64(n)*(1.0/1024))
	case n < 1<<30:
		return fmt.Sprintf("%.2fMB", float64(n)*(1.0/1024)*(1.0/1024))
	case n < 1<<40:
		return fmt.Sprintf("%.2fGB", float64(n)*(1.0/1024)*(1.0/1024)*(1.0/1024))
	default:
		return fmt.Sprintf("%.2fTB", float64(n)*(1.0/1024)*(1.0/1024)*(1.0/1024)*(1.0/1024))
	}
}
