package CAodbhZf

import (
	"encoding/json"
	"testing"
)

func TestRdPirOriginalParsing(t *testing.T) {
	var n RdPirvOZCnMz
	for _, tc := range []struct {
		in   string
		want bool
	}{{`true`, true}, {`1`, true}, {`"TRUE"`, true}, {`"1"`, true}, {`false`, false}, {`0`, false}, {`"yes"`, false}, {` true `, false}, {`" true "`, false}, {`invalid`, false}, {`""true""`, true}, {`null`, false}} {
		if err := n.UnmarshalJSON([]byte(tc.in)); err != nil {
			t.Fatal(err)
		}
		if n.Value() != tc.want {
			t.Fatalf("%q got %v", tc.in, n.Value())
		}
		got, err := json.Marshal(&n)
		if err != nil {
			t.Fatal(err)
		}
		want := "false"
		if tc.want {
			want = "true"
		}
		if string(got) != want {
			t.Fatalf("%q: %s", tc.in, got)
		}
	}
}
