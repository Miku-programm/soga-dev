package cNT24rG4F

import (
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func TestFileLimitRaisedOnlyInChild(t *testing.T) {
	if os.Getenv("SOGA_RECOVERED_RLIMIT_CHILD") == "1" {
		limit := syscall.Rlimit{Cur: 64, Max: 128}
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
			t.Fatal(err)
		}
		MeWz428lV2Z()
		if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
			t.Fatal(err)
		}
		if limit.Cur != 128 || limit.Max != 128 {
			t.Fatal(limit)
		}
		MeWz428lV2Z()
		return
	}
	var before, after syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &before); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestFileLimitRaisedOnlyInChild$")
	cmd.Env = append(os.Environ(), "SOGA_RECOVERED_RLIMIT_CHILD=1")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%v: %s", err, out)
	}
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &after); err != nil {
		t.Fatal(err)
	}
	if before != after {
		t.Fatal("parent limits changed")
	}
}
