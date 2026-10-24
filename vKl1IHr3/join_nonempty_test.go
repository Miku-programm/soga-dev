package vKl1IHr3

import "testing"

func TestJoinNonemptyOriginalTrim(t *testing.T) {
	for _, tc := range []struct {
		in        []string
		sep, want string
	}{{nil, ".", ""}, {[]string{"  a ", "", "\u2003", " b\n"}, ".", "a.b"}, {[]string{"a", "b"}, "", "ab"}, {[]string{"a.", ".b"}, ".", "a...b"}} {
		if got := DP0s5dasePo(tc.in, tc.sep); got != tc.want {
			t.Fatal(got, tc)
		}
	}
}
