package logx

import (
	"log/slog"
	"math"
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type attrTestLogValuer struct {
	remaining int
	calls     *int
}

func (v attrTestLogValuer) LogValue() slog.Value {
	*v.calls++
	if v.remaining > 0 {
		return slog.AnyValue(attrTestLogValuer{v.remaining - 1, v.calls})
	}
	return slog.StringValue("resolved")
}

type attrTestPanickingLogValuer struct{}

func (attrTestPanickingLogValuer) LogValue() slog.Value { panic("log-value-panic") }

type attrTestGroupLogValuer struct{}

func (attrTestGroupLogValuer) LogValue() slog.Value {
	return slog.GroupValue(slog.String("child", "nested"))
}

func TestRecoveredLogValueResolution(t *testing.T) {
	calls := 0
	got := jh9bDVK_MFF(slog.AnyValue(attrTestLogValuer{120, &calls}))
	if got.Kind() != slog.KindString || got.String() != "resolved" || calls != 121 {
		t.Fatalf("resolution %v, calls %d", got, calls)
	}
	defer func() {
		if got := recover(); got != "log-value-panic" {
			t.Fatalf("panic not preserved: %v", got)
		}
	}()
	jh9bDVK_MFF(slog.AnyValue(attrTestPanickingLogValuer{}))
}

func TestRecoveredAttributeGroupsAndEmptyValues(t *testing.T) {
	initial := []zapcore.Field{zap.String("existing", "first")}
	group := slog.Group("outer",
		slog.String("name", "value"),
		slog.Group("inner", slog.Int("count", 7)),
		slog.Group("", slog.Bool("inline", true)),
		slog.Group("empty"),
		slog.Any("lazy", attrTestGroupLogValuer{}),
	)
	got := wFqiIG4V(initial, group, "root.")
	want := []zapcore.Field{zap.String("existing", "first"), zap.String("root.outer.name", "value"), zap.Int64("root.outer.inner.count", 7), zap.Bool("root.outer.inline", true), zap.String("root.outer.lazy.child", "nested")}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
	if got := wFqiIG4V(nil, slog.Attr{}, "prefix."); got != nil {
		t.Fatalf("empty attr allocated fields: %#v", got)
	}
	if got := wFqiIG4V(nil, slog.Group("empty"), ""); got != nil {
		t.Fatalf("empty group allocated fields: %#v", got)
	}
	emptyKey := wFqiIG4V(nil, slog.String("", "value"), "root.")
	if len(emptyKey) != 1 || emptyKey[0].Key != "root." || emptyKey[0].String != "value" {
		t.Fatalf("nonzero empty-key attr lost: %#v", emptyKey)
	}
}

func TestRecoveredAttributeKinds(t *testing.T) {
	instant := time.Date(2026, 10, 23, 12, 34, 56, 123, time.FixedZone("test", 3600))
	negativeZero := math.Copysign(0, -1)
	object := struct{ Count int }{42}
	cases := []struct {
		attr slog.Attr
		want zapcore.Field
	}{
		{slog.Bool("k", true), zap.Bool("p.k", true)},
		{slog.Duration("k", -3*time.Second), zap.Duration("p.k", -3*time.Second)},
		{slog.Float64("k", negativeZero), zap.Float64("p.k", negativeZero)},
		{slog.Float64("k", math.Float64frombits(0x7ff8000000000123)), zap.Float64("p.k", math.Float64frombits(0x7ff8000000000123))},
		{slog.Int64("k", math.MinInt64), zap.Int64("p.k", math.MinInt64)},
		{slog.String("k", ""), zap.String("p.k", "")},
		{slog.Time("k", instant), zap.Time("p.k", instant)},
		{slog.Uint64("k", math.MaxUint64), zap.Uint64("p.k", math.MaxUint64)},
		{slog.Any("k", object), zap.Any("p.k", object)},
		{slog.Any("k", nil), zap.Any("p.k", nil)},
	}
	for _, tc := range cases {
		got := wFqiIG4V(nil, tc.attr, "p.")
		if len(got) != 1 || !reflect.DeepEqual(got[0], tc.want) {
			t.Errorf("%v: got %#v, want %#v", tc.attr, got, tc.want)
		}
	}
}
