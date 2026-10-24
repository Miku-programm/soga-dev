package yaml

import (
	"bytes"
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestOriginalLevelCodecs(t *testing.T) {
	if unsafe.Sizeof(GPDvVWMg{}) != 24 || unsafe.Sizeof(logrus.Logger{}) != 128 || unsafe.Sizeof(logrus.Entry{}) != 112 || unsafe.Sizeof(logrus.TextFormatter{}) != 88 {
		t.Fatalf("original/runtime dependency layout: event=%d logger=%d entry=%d formatter=%d", unsafe.Sizeof(GPDvVWMg{}), unsafe.Sizeof(logrus.Logger{}), unsafe.Sizeof(logrus.Entry{}), unsafe.Sizeof(logrus.TextFormatter{}))
	}
	for i, name := range []string{"debug", "info", "warning", "error", "silent"} {
		level := H2hbGW(i)
		if level.String() != name || (&GPDvVWMg{ZbnaAXx: level}).Type() != name {
			t.Fatal("level names")
		}
		data, e := level.MarshalJSON()
		if e != nil || string(data) != "\""+name+"\"" {
			t.Fatal("JSON marshal", string(data), e)
		}
		yaml, e := level.MarshalYAML()
		if e != nil || yaml != name {
			t.Fatal("YAML marshal")
		}
		var restored H2hbGW
		if e = restored.UnmarshalJSON(data); e != nil || restored != level {
			t.Fatal("roundtrip")
		}
		// Original ignores decode error if the callback did populate a known value.
		e = restored.UnmarshalYAML(func(v any) error { *(v.(*string)) = name; return errors.New("ignored decoder failure") })
		if e != nil || restored != level {
			t.Fatal("YAML error handling")
		}
	}
	for _, input := range []string{`"warn"`, `"INFO"`, `null`, `42`, `{`, `"unknown"`} {
		l := H2hbGW(3)
		e := l.UnmarshalJSON([]byte(input))
		if e == nil || e.Error() != "invalid mode" || l != 3 {
			t.Fatal(input, l, e)
		}
	}
	l := H2hbGW(2)
	e := l.UnmarshalYAML(func(any) error { return errors.New("decoder failure") })
	if e == nil || e.Error() != "invalid mode" || l != 2 {
		t.Fatal("YAML invalid")
	}
	if H2hbGW(-1).String() != "unknown" || H2hbGW(5).String() != "unknown" {
		t.Fatal("unknown level")
	}
}
func TestWarningsPublishBeforeConsoleFiltering(t *testing.T) {
	oldEvents, oldLevel, oldLogger := eventsAt5096988, levelAt3787AB0, loggerAt5029DF8
	defer func() { eventsAt5096988, levelAt3787AB0, loggerAt5029DF8 = oldEvents, oldLevel, oldLogger }()
	eventsAt5096988 = make(chan any)
	var output bytes.Buffer
	loggerAt5029DF8 = logrus.New()
	loggerAt5029DF8.SetOutput(&output)
	loggerAt5029DF8.SetLevel(logrus.DebugLevel)
	loggerAt5029DF8.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, DisableColors: true})
	levelAt3787AB0 = 4
	done := make(chan struct{})
	go func() { Dl2Gj6("cipher %s failed: %d", "test", 7); close(done) }()
	select {
	case <-done:
		t.Fatal("publication must block")
	case <-time.After(10 * time.Millisecond):
	}
	if output.Len() != 0 {
		t.Fatal("console output preceded publication")
	}
	select {
	case item := <-eventsAt5096988:
		if !reflect.DeepEqual(item, GPDvVWMg{2, "cipher test failed: 7"}) {
			t.Fatal("event", item)
		}
	case <-time.After(time.Second):
		t.Fatal("event timeout")
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("sender not released")
	}
	if output.Len() != 0 {
		t.Fatal("silent level did not filter console")
	}
	levelAt3787AB0 = 1
	uRTsMwuzLPTw(0, "hidden")
	uRTsMwuzLPTw(1, "info")
	uRTsMwuzLPTw(2, "warning")
	uRTsMwuzLPTw(3, "error")
	uRTsMwuzLPTw(4, "silent")
	uRTsMwuzLPTw(99, "unknown")
	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 3 || !strings.Contains(lines[0], "level=info") || !strings.Contains(lines[1], "level=warning") || !strings.Contains(lines[2], "level=error") {
		t.Fatal("level mapping", output.String())
	}
}
