package VasCJz

import (
	jsonv1 "encoding/json"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// iY_pOEhZS restores 0x1250a00. Splitting at the first hyphen is intentional:
// a leading negative sign is a separator, while the second bound may be negative.
// The original does not reorder bounds or check that from <= to.
func iY_pOEhZS(input string) (int32, int32, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, 0, fmt.Errorf("empty range string")
	}
	parts := strings.SplitN(input, "-", 2)
	from, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 32)
	if err != nil {
		return 0, 0, err
	}
	if len(parts) == 1 {
		return int32(from), int32(from), nil
	}
	to, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return int32(from), int32(to), nil
}

// awXBGVSDBEx7 is the method-free original object form (RTTI 0x36302f0).
// Its two int32 fields and JSON tags are identical to MUVoWl0.
type awXBGVSDBEx7 MUVoWl0

// UnmarshalJSON restores 0x1250c00, including numeric-string precedence,
// three distinct error classes, and assignment only after complete success.
func (r *MUVoWl0) UnmarshalJSON(data []byte) error {
	var raw jsontext.Value
	if err := json.Unmarshal(data, &raw, jsonv1.DefaultOptionsV1()); err != nil {
		return err
	}
	var number jsonv1.Number
	if err := json.Unmarshal(raw, &number, jsonv1.DefaultOptionsV1()); err == nil {
		value, err := number.Int64()
		if err != nil {
			return fmt.Errorf("RangeConfig: invalid number %q: %w", number, err)
		}
		if value < math.MinInt32 || value > math.MaxInt32 {
			return fmt.Errorf("RangeConfig: number out of int32 range: %d", value)
		}
		r.AS9zG362 = int32(value)
		r.HHTPwX0N = int32(value)
		return nil
	}
	var text string
	if err := json.Unmarshal(raw, &text, jsonv1.DefaultOptionsV1()); err == nil {
		from, to, err := iY_pOEhZS(text)
		if err != nil {
			return fmt.Errorf("RangeConfig: invalid string %q: %w", text, err)
		}
		r.AS9zG362 = from
		r.HHTPwX0N = to
		return nil
	}
	var object awXBGVSDBEx7
	if err := json.Unmarshal(raw, &object, jsonv1.DefaultOptionsV1()); err != nil {
		return err
	}
	*r = MUVoWl0(object)
	return nil
}
