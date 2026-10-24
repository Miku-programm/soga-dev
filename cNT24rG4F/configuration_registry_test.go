package cNT24rG4F

import (
	"reflect"
	"testing"
)

func isolateRegistry(t *testing.T) {
	t.Helper()
	defaults, values, registered, callbacks := defaultsAt5095AA0, configurationAt5095AA8, registeredAt5099320, afterLoadAt5099300
	t.Cleanup(func() {
		defaultsAt5095AA0, configurationAt5095AA8 = defaults, values
		registeredAt5099320, afterLoadAt5099300 = registered, callbacks
	})
	defaultsAt5095AA0 = make(map[string]string)
	configurationAt5095AA8 = make(map[string]string)
	registeredAt5099320, afterLoadAt5099300 = nil, nil
}

func TestRegistrationAndAssignmentAreSeparate(t *testing.T) {
	isolateRegistry(t)
	type config struct {
		Count    int    `conf:"COUNT" default:"12"`
		Blank    string `conf:"" default:"empty-key"`
		Untagged string
		Spaced   string `conf:" x " default:" whitespace "`
	}
	c := config{Count: 99, Untagged: "keep"}
	_NA1kft(&c)
	if c.Count != 99 || c.Blank != "" || len(registeredAt5099320) != 1 || registeredAt5099320[0] != &c {
		t.Fatal("registration must retain pointer without assigning defaults")
	}
	want := map[string]string{"count": "12", "": "empty-key", " x ": " whitespace "}
	if !reflect.DeepEqual(defaultsAt5095AA0, want) {
		t.Fatalf("registered defaults: %#v", defaultsAt5095AA0)
	}
	configurationAt5095AA8 = map[string]string{"count": "-8", "": "assigned"}
	bFnBl5Ak9(&c)
	if c.Count != -8 || c.Blank != "assigned" || c.Untagged != "keep" || c.Spaced != "" {
		t.Fatalf("assignment differs: %#v", c)
	}
}

func TestDuplicateDefaultsAndRegistrationFailure(t *testing.T) {
	isolateRegistry(t)
	a := struct {
		X int `conf:"x" default:"1"`
	}{}
	b := struct {
		X int `conf:"x" default:"2"`
	}{}
	_NA1kft(&a)
	_NA1kft(&b)
	if defaultsAt5095AA0["x"] != "2" || len(registeredAt5099320) != 2 {
		t.Fatal("later registrations must overwrite default text and retain both pointers")
	}
	bad := struct {
		Before      string `conf:"before" default:"retained"`
		Unsupported uint   `conf:"unsupported"`
	}{}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("unsupported default type must panic")
			}
		}()
		_NA1kft(&bad)
	}()
	if defaultsAt5095AA0["before"] != "retained" || len(registeredAt5099320) != 2 {
		t.Fatal("registration failure must preserve earlier map writes but not append instance")
	}
}
