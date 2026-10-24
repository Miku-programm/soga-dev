package cNT24rG4F

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestKeyStateTransitions(t *testing.T) {
	flag, mirror, marker, key := keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800, configAt5096DC0.RM5SFRF
	t.Cleanup(func() {
		keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800 = flag, mirror, marker
		configAt5096DC0.RM5SFRF = key
	})
	for _, value := range []bool{false, true, false} {
		T15Jifo(value)
		if keyStateAt50C6C48 != value || keyStateMirrorAt3787300 == value {
			t.Fatal("state and inverse state disagree")
		}
		wantMarker := uint64(0x5a4b3c2d5a4b3c2d)
		if value {
			wantMarker = 0xa5b6c7d8a5b6c7d8
		}
		if keyStateMarkerAt3787800 != wantMarker {
			t.Fatalf("marker: %#x", keyStateMarkerAt3787800)
		}
		for _, key := range []string{"", "plain-key", " \x00\n"} {
			configAt5096DC0.RM5SFRF = key
			if got := FrFQQhzTR5fY(); got != key {
				t.Fatalf("getter changed %q to %q", key, got)
			}
			for name, check := range map[string]func() bool{"Q0SDIN": Q0SDIN, "Rs_wP3R": Rs_wP3R, "HuQEoQ": HuQEoQ, "Gh2DnstC": Gh2DnstC, "JatfUF38O": JatfUF38O} {
				if got, want := check(), key == "" || value; got != want {
					t.Fatalf("%s flag %v key %q: got %v, want %v", name, value, key, got, want)
				}
			}
		}
	}
}

func TestKeyStateInconsistencyExits(t *testing.T) {
	if scenario := os.Getenv("SOGA_KEY_STATE_TEST_CHILD"); scenario != "" {
		configAt5096DC0.RM5SFRF = "test-key"
		switch scenario {
		case "mirror":
			T15Jifo(false)
			keyStateMirrorAt3787300 = false
		case "false-marker":
			T15Jifo(false)
			keyStateMarkerAt3787800 = 0
		case "true-marker":
			T15Jifo(true)
			keyStateMarkerAt3787800 = 0
		}
		_, _ = os.Stdout.WriteString("child-ready\n")
		check := Q0SDIN
		if os.Getenv("SOGA_KEY_STATE_CHECK") == "Rs_wP3R" {
			check = Rs_wP3R
		}
		if os.Getenv("SOGA_KEY_STATE_CHECK") == "HuQEoQ" {
			check = HuQEoQ
		}
		if os.Getenv("SOGA_KEY_STATE_CHECK") == "Gh2DnstC" {
			check = Gh2DnstC
		}
		if os.Getenv("SOGA_KEY_STATE_CHECK") == "JatfUF38O" {
			check = JatfUF38O
		}
		if !check() {
			os.Exit(98)
		}
		time.Sleep(time.Second)
		os.Exit(99)
	}
	for _, check := range []string{"Q0SDIN", "Rs_wP3R", "HuQEoQ", "Gh2DnstC", "JatfUF38O"} {
		for _, scenario := range []string{"mirror", "false-marker", "true-marker"} {
			t.Run(check+"/"+scenario, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestKeyStateInconsistencyExits$")
				cmd.Env = append(os.Environ(), "SOGA_KEY_STATE_TEST_CHILD="+scenario, "SOGA_KEY_STATE_CHECK="+check)
				out, err := cmd.CombinedOutput()
				if err != nil || !bytes.Equal(out, []byte("child-ready\n")) {
					t.Fatalf("exit path: %v, output %q", err, out)
				}
			})
		}
	}
}
