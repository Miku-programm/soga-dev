package rtx

import (
	"testing"
	"time"
)

func TestDatabaseDateTimeUsesLocalWallClock(t *testing.T) {
	before := time.Now().Truncate(time.Second)
	got := XMMIhT97lTE7()
	after := time.Now().Truncate(time.Second)
	parsed, err := time.ParseInLocation("2006-01-02 15:04:05", got, time.Local)
	if err != nil || parsed.Before(before) || parsed.After(after) {
		t.Fatalf("date %q outside [%v,%v]: %v", got, before, after, err)
	}
}
