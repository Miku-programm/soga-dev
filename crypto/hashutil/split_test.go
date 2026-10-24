package hashutil

import (
	"reflect"
	"testing"
)

func TestSplitBoundaries(t *testing.T) {
	for _, tc := range []struct {
		text, sep string
		n         int
		want      []string
	}{
		{" a : : b ", ":", -1, []string{"a", "b"}},
		{" a : b : c ", ":", 2, []string{"a", "b : c"}},
		{"  a  ", ":", 0, []string{}},
		{"\u2003 \t", ":", -1, []string{}},
		{"中 文", "", -1, []string{"中", "文"}},
		{" a:b ", ":", 1, []string{"a:b"}},
	} {
		got := DAaBs7_C(tc.text, tc.sep, tc.n)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("split(%q, %q, %d) = %#v; want %#v", tc.text, tc.sep, tc.n, got, tc.want)
		}
	}
}
