package config

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestAssignObservedConversions(t *testing.T) {
	var b bool
	for _, text := range []string{"true", "TRUE", "1", " true", "false", ""} {
		ieqMch_a(reflect.ValueOf(&b).Elem(), text)
		if b != (text == "true") {
			t.Errorf("bool conversion for %q: %v", text, b)
		}
	}
	var small int8
	ieqMch_a(reflect.ValueOf(&small).Elem(), "128")
	if small != -128 {
		t.Fatal("narrow signed integer must retain reflect truncation")
	}
	var integer int64
	for _, text := range []string{"9223372036854775808", "0x10", " 1", "bad"} {
		integer = 99
		ieqMch_a(reflect.ValueOf(&integer).Elem(), text)
		if integer != 0 {
			t.Errorf("invalid integer %q must become zero, got %d", text, integer)
		}
	}
	var f float32
	ieqMch_a(reflect.ValueOf(&f).Elem(), "1e40")
	if !math.IsInf(float64(f), 1) {
		t.Fatal("float32 overflow after successful float64 parsing must become infinity")
	}
	ieqMch_a(reflect.ValueOf(&f).Elem(), "1e400")
	if f != 0 {
		t.Fatal("float64 parse overflow must become zero")
	}
	type textValue string
	var text textValue
	ieqMch_a(reflect.ValueOf(&text).Elem(), " untrimmed ")
	if text != " untrimmed " {
		t.Fatal("string assignment must preserve its input")
	}
	type textList []string
	var list textList
	ieqMch_a(reflect.ValueOf(&list).Elem(), " a, ,b ")
	if !reflect.DeepEqual(list, textList{"a", "b"}) {
		t.Fatalf("slice assignment: %v", list)
	}
	type durationValue time.Duration
	var duration durationValue
	ieqMch_a(reflect.ValueOf(&duration).Elem(), "123")
	if duration != 123 {
		t.Fatal("only the exact time.Duration type is prohibited")
	}
}

func TestAssignPanicBoundaries(t *testing.T) {
	type namedText string
	var duration time.Duration
	var integers []int
	var unsigned uint
	var namedStrings []namedText
	for _, tc := range []struct {
		name    string
		value   reflect.Value
		message string
	}{
		{"duration", reflect.ValueOf(&duration).Elem(), "conf tag does not support time.Duration"},
		{"integer slice", reflect.ValueOf(&integers).Elem(), "conf tag only supports []string slice, got []int"},
		{"unsigned", reflect.ValueOf(&unsigned).Elem(), "conf tag does not support uint"},
		{"named slice element", reflect.ValueOf(&namedStrings).Elem(), ""},
		{"invalid value", reflect.Value{}, ""},
		{"unaddressable value", reflect.ValueOf(7), ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				got := recover()
				if got == nil {
					t.Fatal("expected panic")
				}
				if tc.message != "" && got != tc.message {
					t.Errorf("panic %v, want %q", got, tc.message)
				}
			}()
			ieqMch_a(tc.value, "1")
		})
	}
}
