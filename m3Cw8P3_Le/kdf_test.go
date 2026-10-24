package m3Cw8P3_Le

import (
	"encoding/hex"
	"testing"
)

func TestArNrGS_EVKnownExpansion(t *testing.T) {
	const vector = "5f4dcc3b5aa765d61d8327deb882cf992b95990a9151374abd8ff8c5a7a0fe08b7b4372cdfbcb3d16a2631b59b509e94"
	if got := hex.EncodeToString(ArNrGS_EV("password", 48)); got != vector {
		t.Fatal(got)
	}
	if got := hex.EncodeToString(ArNrGS_EV("password", 17)); got != vector[:34] {
		t.Fatal(got)
	}
	if ArNrGS_EV("password", 0) != nil {
		t.Fatal("zero must be nil")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("negative length did not panic")
		}
	}()
	ArNrGS_EV("password", -1)
}
