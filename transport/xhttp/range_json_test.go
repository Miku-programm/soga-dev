package xhttp

import (
	"errors"
	"strconv"
	"testing"
)

// Golden results come from 35 bounded executions of the original 0x1250c00
// and 0x1250a00. JSON/strconv library calls use independently generated Go
// fixtures. Only the fixture object's Go type name is mapped to original RTTI.
func TestRangeJSONOriginalMachineFixtures(t *testing.T) {
	cases := []struct {
		input     string
		from, to  int32
		errorText string
	}{
		{"1", 1, 1, ""},
		{"-1", -1, -1, ""},
		{"2147483647", 2147483647, 2147483647, ""},
		{"-2147483648", -2147483648, -2147483648, ""},
		{"2147483648", 71, 93, "RangeConfig: number out of int32 range: 2147483648"},
		{"-2147483649", 71, 93, "RangeConfig: number out of int32 range: -2147483649"},
		{"9223372036854775808", 71, 93, "RangeConfig: invalid number \"9223372036854775808\": strconv.ParseInt: parsing \"9223372036854775808\": value out of range"},
		{"1.5", 71, 93, "RangeConfig: invalid number \"1.5\": strconv.ParseInt: parsing \"1.5\": invalid syntax"},
		{"1e3", 71, 93, "RangeConfig: invalid number \"1e3\": strconv.ParseInt: parsing \"1e3\": invalid syntax"},
		{"null", 71, 93, "RangeConfig: invalid number \"\": strconv.ParseInt: parsing \"\": invalid syntax"},
		{"\"123\"", 123, 123, ""},
		{"\"-1\"", -1, -1, ""},
		{"\"2147483648\"", 71, 93, "RangeConfig: number out of int32 range: 2147483648"},
		{"\"9223372036854775808\"", 71, 93, "RangeConfig: invalid number \"9223372036854775808\": strconv.ParseInt: parsing \"9223372036854775808\": value out of range"},
		{"\"1.5\"", 71, 93, "RangeConfig: invalid number \"1.5\": strconv.ParseInt: parsing \"1.5\": invalid syntax"},
		{"\"1-3\"", 1, 3, ""},
		{"\"  4 - 2  \"", 4, 2, ""},
		{"\"3--2\"", 3, -2, ""},
		{"\"-2-3\"", 71, 93, "RangeConfig: invalid string \"-2-3\": strconv.ParseInt: parsing \"\": invalid syntax"},
		{"\"\"", 71, 93, "RangeConfig: invalid string \"\": empty range string"},
		{"\"   \"", 71, 93, "RangeConfig: invalid string \"   \": empty range string"},
		{"\"x-4\"", 71, 93, "RangeConfig: invalid string \"x-4\": strconv.ParseInt: parsing \"x\": invalid syntax"},
		{"\"1-x\"", 71, 93, "RangeConfig: invalid string \"1-x\": strconv.ParseInt: parsing \"x\": invalid syntax"},
		{"\"1-2-3\"", 71, 93, "RangeConfig: invalid string \"1-2-3\": strconv.ParseInt: parsing \"2-3\": invalid syntax"},
		{"\"2147483648-1\"", 71, 93, "RangeConfig: invalid string \"2147483648-1\": strconv.ParseInt: parsing \"2147483648\": value out of range"},
		{"\"1-2147483648\"", 71, 93, "RangeConfig: invalid string \"1-2147483648\": strconv.ParseInt: parsing \"2147483648\": value out of range"},
		{"{\"from\":3,\"to\":2}", 3, 2, ""},
		{"{}", 0, 0, ""},
		{"{\"FROM\":-2,\"TO\":3}", -2, 3, ""},
		{"{\"from\":3,\"to\":\"2\"}", 71, 93, "json: cannot unmarshal string into Go struct field awXBGVSDBEx7.to of type int32"},
		{"{\"from\":null,\"to\":4}", 0, 4, ""},
		{"[]", 71, 93, "json: cannot unmarshal array into Go value of type xhttp.awXBGVSDBEx7"},
		{"true", 71, 93, "json: cannot unmarshal bool into Go value of type xhttp.awXBGVSDBEx7"},
		{"{", 71, 93, "unexpected end of JSON input"},
		{"1 2", 71, 93, "invalid character '2' after top-level value"},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got := MUVoWl0{AS9zG362: 71, HHTPwX0N: 93}
			err := got.UnmarshalJSON([]byte(tc.input))
			var errorText string
			if err != nil {
				errorText = err.Error()
			}
			if errorText != tc.errorText {
				t.Fatalf("error = %q, want %q", errorText, tc.errorText)
			}
			if got.AS9zG362 != tc.from || got.HHTPwX0N != tc.to {
				t.Fatalf("range = %+v, want [%d,%d]", got, tc.from, tc.to)
			}
		})
	}
}

func TestRangeJSONPreservesWrappedNumericErrors(t *testing.T) {
	for _, input := range []string{`1.5`, `"1-x"`, `"2147483648-1"`} {
		var value MUVoWl0
		err := value.UnmarshalJSON([]byte(input))
		var numberError *strconv.NumError
		if !errors.As(err, &numberError) {
			t.Fatalf("%s: lost wrapped ParseInt error: %v", input, err)
		}
	}
}

func TestRangeTextWhitespaceAndFirstDelimiter(t *testing.T) {
	for _, tc := range []struct {
		input    string
		from, to int32
	}{
		{"\u2003+12\u00a0", 12, 12}, {"5--7", 5, -7}, {"7-2", 7, 2},
	} {
		from, to, err := iY_pOEhZS(tc.input)
		if err != nil || from != tc.from || to != tc.to {
			t.Fatalf("%q: %d %d %v", tc.input, from, to, err)
		}
	}
	for _, input := range []string{"-1", "x-y", "1-x"} {
		from, to, err := iY_pOEhZS(input)
		if err == nil || from != 0 || to != 0 {
			t.Fatalf("%q: %d %d %v", input, from, to, err)
		}
	}
}
