package yaml

import (
	"encoding/json"
	"errors"
)

type H2hbGW int

var levelsAt5096980 = map[string]H2hbGW{"error": 3, "warning": 2, "info": 1, "debug": 0, "silent": 4}

// UnmarshalYAML:0x1b31fc0. The original ignores the callback error and still
// looks up the resulting string. Unknown input leaves the receiver unchanged.
func (l *H2hbGW) UnmarshalYAML(decode func(any) error) error {
	var value string
	_ = decode(&value)
	if level, ok := levelsAt5096980[value]; ok {
		*l = level
		return nil
	}
	return errors.New("invalid mode")
}

// UnmarshalJSON:0x1b32100 passes DefaultOptionsV1 to json/v2, equivalent here
// to encoding/json.Unmarshal, and likewise ignores the decoding error.
func (l *H2hbGW) UnmarshalJSON(data []byte) error {
	var value string
	_ = json.Unmarshal(data, &value)
	if level, ok := levelsAt5096980[value]; ok {
		*l = level
		return nil
	}
	return errors.New("invalid mode")
}
func (l H2hbGW) String() string { //0x1b32460
	switch l {
	case 0:
		return "debug"
	case 1:
		return "info"
	case 2:
		return "warning"
	case 3:
		return "error"
	case 4:
		return "silent"
	default:
		return "unknown"
	}
}
func (l H2hbGW) MarshalJSON() ([]byte, error) { return json.Marshal(l.String()) } //0x1b322a0
func (l H2hbGW) MarshalYAML() (any, error)    { return l.String(), nil }          //0x1b323a0

// GPDvVWMg: original24-byte log event, transmitted by value through the channel.
type GPDvVWMg struct {
	ZbnaAXx H2hbGW
	URokq7  string
}

func (e *GPDvVWMg) Type() string { return e.ZbnaAXx.String() } //0x1b333e0
