package toml

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/BurntSushi/toml"
)

func TestMetadataAndNestedTables(t *testing.T) {
	if unsafe.Sizeof(XykYEQs{}) != 96 || unsafe.Sizeof(Q2Erylds4g{}) != 16 {
		t.Fatal("original amd64 decoder/metadata layout changed")
	}
	var got struct {
		Enabled bool `toml:"enable"`
		Routes  []struct {
			Rules []string `toml:"rules"`
			Outs  []struct {
				Server string `toml:"server"`
				Port   int    `toml:"port"`
			} `toml:"outs"`
		} `toml:"routes"`
	}
	md, err := Ikm5Xa("\ufeffenable=true\nunused=42\n[[routes]]\nrules=['*','ip:10.0.0.0/8']\n[[routes.outs]]\nserver='::1'\nport=443\n", &got)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Enabled || len(got.Routes) != 1 || len(got.Routes[0].Outs) != 1 || got.Routes[0].Outs[0].Server != "::1" || got.Routes[0].Outs[0].Port != 443 {
		t.Fatalf("nested tables: %+v", got)
	}
	// IsDefined descends maps; array-of-table elements are not map nodes.
	if !md.IsDefined("enable") || md.IsDefined("routes", "outs", "port") || !reflect.DeepEqual(md.Undecoded(), []toml.Key{{"unused"}}) {
		t.Fatalf("metadata: keys=%v undecoded=%v", md.Keys(), md.Undecoded())
	}
}

func TestInputAndSyntaxFailures(t *testing.T) {
	var nilMap *map[string]any
	for _, tc := range []struct {
		dst  any
		text string
	}{
		{nil, "toml: cannot decode to non-pointer <nil>"},
		{1, "toml: cannot decode to non-pointer \"int\""},
		{nilMap, "toml: cannot decode to nil value"},
	} {
		_, err := Ikm5Xa("a=1", tc.dst)
		if err == nil || !strings.HasPrefix(err.Error(), tc.text) {
			t.Fatalf("%T: %v", tc.dst, err)
		}
	}
	for _, data := range []string{"\x00a=1", "a=1\na=2", "a=[1,"} {
		var dst map[string]any
		_, err := Ikm5Xa(data, &dst)
		var pe toml.ParseError
		if !errors.As(err, &pe) {
			t.Fatalf("%q: %T %v", data, err, err)
		}
		if data == "a=[1," && pe.Position.Line != 0 {
			t.Fatalf("EOF retains the parser's zero position: %+v", pe.Position)
		}
	}
}

func TestReaderFailureIsReturnedUnchanged(t *testing.T) {
	want := errors.New("fixture read failure")
	decoder := toml.NewDecoder(failedReader{want})
	var dst map[string]any
	_, err := decoder.Decode(&dst)
	if err != want {
		t.Fatalf("got %v, want same error instance", err)
	}
}

type failedReader struct{ err error }

func (r failedReader) Read([]byte) (int, error) { return 0, r.err }
