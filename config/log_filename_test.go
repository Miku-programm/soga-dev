package config

import (
	"testing"
	"time"
)

func TestLogFilenameLocalTime(t *testing.T) {
	previous := time.Local
	t.Cleanup(func() { time.Local = previous })
	time.Local = time.FixedZone("test-local", 8*60*60)
	for _, tc := range []struct {
		name     string
		parse    func(string) (time.Time, bool)
		filename string
		hour     int
	}{
		{"access", CpXlPLFmKQcD, "soga-access-2024-02-29 09:00:00.log", 9},
		{"ordinary", YbaVop37K2p, "prefix/soga-anytls-2024-02-29.log.old", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := tc.parse(tc.filename)
			want := time.Date(2024, 2, 29, tc.hour, 0, 0, 0, time.Local)
			if !ok || got != want {
				t.Fatalf("got %v, %v; want %v", got, ok, want)
			}
		})
	}
}

func TestLogFilenameRejectsInvalidCalendar(t *testing.T) {
	for _, name := range []string{
		"soga-access-2025-02-29 09:00:00.log",
		"soga-access-2024-02-29 24:00:00.log",
		"soga-access-2024-00-29 09:00:00.log",
		"soga-access-2024-02-29 09:00:00.log.old",
		"",
	} {
		got, ok := CpXlPLFmKQcD(name)
		if ok || got != (time.Time{}) {
			t.Fatalf("%q: %v, %v", name, got, ok)
		}
	}
	for _, name := range []string{
		"soga-anytls-2025-02-29.log",
		"soga-anytls-2024-13-01.log",
		"soga-anytls-2024-02-00.log",
		"soga-节点-2024-02-29.log",
		"",
	} {
		got, ok := YbaVop37K2p(name)
		if ok || got != (time.Time{}) {
			t.Fatalf("%q: %v, %v", name, got, ok)
		}
	}
}
