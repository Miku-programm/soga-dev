package k_qeiAYoxa

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"recovered.local/soga/cNT24rG4F"
)

func TestRetentionClosuresPropagateErrorsAndTruncate(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file")
	if err := os.WriteFile(path, nil, 0600); err != nil {
		t.Fatal(err)
	}
	entry, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 10, 23, 12, 0, 0, 0, time.UTC)
	when := now.Add(-3*time.Hour + time.Nanosecond)
	parse := func(string) (time.Time, bool) { return when, true }
	var removed []string
	fn := retentionWalker(now, 2, time.Hour, parse, &removed)
	if err := fn("file", fs.FileInfoToDirEntry(entry), nil); err != nil {
		t.Fatal(err)
	}
	if len(removed) != 0 {
		t.Fatal("fractional age must truncate before comparison")
	}
	when = now.Add(-3 * time.Hour)
	if err := fn("file", fs.FileInfoToDirEntry(entry), nil); err != nil {
		t.Fatal(err)
	}
	if len(removed) != 1 {
		t.Fatal("strictly older file not selected")
	}
	sentinel := errors.New("walk failure")
	if err := fn("file", nil, sentinel); err != sentinel {
		t.Fatal("walk error not preserved")
	}
}

func TestRetentionJobsOnlyTouchFixtureFiles(t *testing.T) {
	if os.Getenv("SOGA_RETENTION_HELPER") != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestRetentionJobsOnlyTouchFixtureFiles$")
		cmd.Env = append(os.Environ(), "SOGA_RETENTION_HELPER=1")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%v: %s", err, out)
		}
		return
	}
	time.Local = time.UTC
	root := t.TempDir()
	access := filepath.Join(root, "access")
	logs := filepath.Join(root, "logs")
	for _, dir := range []string{access, logs, filepath.Join(access, "nested"), filepath.Join(logs, "nested")} {
		if err := os.MkdirAll(dir, 0700); err != nil {
			t.Fatal(err)
		}
	}
	now := time.Now()
	hour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.Local)
	day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	accessName := func(hours int) string {
		return "soga-access-" + hour.Add(-time.Duration(hours)*time.Hour).Format("2006-01-02 15") + ":00:00.log"
	}
	logName := func(days int) string { return "soga-vless-" + day.AddDate(0, 0, -days).Format("2006-01-02") + ".log" }
	remove := []string{filepath.Join(access, accessName(3)), filepath.Join(logs, logName(3)), filepath.Join(logs, "nested", logName(3))}
	keep := []string{filepath.Join(access, accessName(2)), filepath.Join(access, "nested", accessName(3)), filepath.Join(logs, logName(2)), filepath.Join(access, "unrelated"), filepath.Join(logs, "unrelated")}
	for _, p := range append(append([]string{}, remove...), keep...) {
		if err := os.WriteFile(p, []byte("fixture"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	config := filepath.Join(root, "soga.conf")
	text := fmt.Sprintf("access_log_enable=false\naccess_log_mode=local\naccess_log_dir=%s\naccess_log_retention_hours=2\nlog_file_dir=%s\nlog_file_retention_days=2\nserver_type=vless\n", access, logs)
	if err := os.WriteFile(config, []byte(text), 0600); err != nil {
		t.Fatal(err)
	}
	cNT24rG4F.MkwNWciPHS(config)
	if err := cNT24rG4F.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	a, b := rjabCj7Pjm(), nrTbX9QD()
	if a.Cron() != "0 1 * * * ?" || b.Cron() != "0 0 * * * ?" {
		t.Fatal("wrong original schedules")
	}
	a.Run()
	b.Run()
	for _, p := range remove {
		if _, err := os.Stat(p); !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected removal %q: %v", p, err)
		}
	}
	for _, p := range keep {
		if _, err := os.Stat(p); err != nil {
			t.Fatalf("unexpected removal %q: %v", p, err)
		}
	}
}
