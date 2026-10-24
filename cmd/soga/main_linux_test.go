package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	Cert "recovered.local/soga/cert"
	C "recovered.local/soga/config"
	"strings"
	"testing"
	"time"
)

func TestMainOriginalFlagRegistration(t *testing.T) {
	for _, name := range append(C.ConfigurationFlagNames(), Cert.FHSbLFuj2()...) {
		f := flag.Lookup(name)
		if f == nil || f.DefValue != "" {
			t.Fatalf("dynamic flag %q: %#v", name, f)
		}
	}
	defaults := map[string]string{"v": "false", "version": "false", "c": "/etc/soga/soga.conf", "b": "/etc/soga/blockList", "w": "/etc/soga/whiteList", "d": "/etc/soga/dns.yml", "r": "/etc/soga/routes.toml"}
	for name, want := range defaults {
		f := flag.Lookup(name)
		if f == nil || f.DefValue != want {
			t.Fatalf("flag %q: %#v, want %q", name, f, want)
		}
	}
}

// Invoke only main's version short circuit, never its application loop. The
// nonexistent path proves this branch precedes all configuration file effects.
func TestMainVersionShortCircuit(t *testing.T) {
	if arg := os.Getenv("SOGA_MAIN_VERSION_FIXTURE"); arg != "" {
		os.Args = []string{"soga", arg, "-c", os.Getenv("SOGA_MAIN_ABSENT_CONFIG")}
		main()
		return
	}
	for _, arg := range []string{"-v", "-version"} {
		t.Run(arg, func(t *testing.T) {
			path := t.TempDir() + "/never-created/soga.conf"
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestMainVersionShortCircuit$")
			cmd.Env = append(os.Environ(), "SOGA_MAIN_VERSION_FIXTURE="+arg, "SOGA_MAIN_ABSENT_CONFIG="+path)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("version branch: %v\n%s", err, out)
			}
			if !strings.HasPrefix(string(out), C.StartupVersion()+"\n") {
				t.Fatalf("version output: %s", out)
			}
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				t.Fatalf("version branch touched config: %v", err)
			}
		})
	}
}
