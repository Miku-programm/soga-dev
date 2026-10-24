package connutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOriginalLineCounting(t *testing.T) {
	for _, tc := range []struct {
		s string
		n int
	}{{"", -1}, {"header", -1}, {"header\n", 0}, {"header\na\nb\n", 2}, {"header\na", 0}} {
		p := filepath.Join(t.TempDir(), "fixture")
		if err := os.WriteFile(p, []byte(tc.s), 0600); err != nil {
			t.Fatal(err)
		}
		n, e := a6N3H7(p)
		if e != nil || n != tc.n {
			t.Fatalf("%q: %d %v", tc.s, n, e)
		}
	}
	n, e := a6N3H7(filepath.Join(t.TempDir(), "missing"))
	if n != 0 || !os.IsNotExist(e) {
		t.Fatal(n, e)
	}
}
