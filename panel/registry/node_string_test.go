package registry

import (
	"math"
	"testing"
)

func TestNodeStringJSONAndFallback(t *testing.T) {
	if got := T0j_Rco(map[string]any{"x": "<tag>", "a": 1}); got != `{"a":1,"x":"\u003ctag\u003e"}` {
		t.Fatal(got)
	}
	if got := T0j_Rco(math.Inf(1)); got != "+Inf" {
		t.Fatal(got)
	}
	if got := T0j_Rco(nil); got != "null" {
		t.Fatal(got)
	}
}
