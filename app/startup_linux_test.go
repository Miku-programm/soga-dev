package app

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	C "recovered.local/soga/config"
	"strings"
	"testing"
	"time"
)

// These real early exits occur before NTP, panel I/O, listeners or licensing.
// Test them in child processes because the original fatal logger calls Exit(0).
func TestCentralStartupEarlyErrorsAndRecovery(t *testing.T) {
	if mode := os.Getenv("SOGA_CENTRAL_EARLY_EXIT"); mode != "" {
		var app *M0mmPxPJsa
		if mode == "nilreceiver" {
			app.InitLinux()
			os.Exit(91)
		}
		C.MkwNWciPHS(os.Getenv("SOGA_CENTRAL_EARLY_CONFIG"))
		if err := C.TvBECnD9fM6(); err != nil {
			t.Fatal(err)
		}
		app.so4m0_oRopC()
		os.Exit(92)
	}
	for _, tt := range []struct{ mode, config, want string }{
		{"unknown", "type=unknown-central-fixture\nserver_type=vless\n", "unknown panel: unknown-central-fixture"},
		{"no_nodes", "type=v2board\nserver_type=vless\nnode_id=\n", "缺少 node_id"},
		{"nilreceiver", "", "invalid memory address or nil pointer dereference"},
	} {
		t.Run(tt.mode, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "soga.conf")
			if err := os.WriteFile(path, []byte(tt.config+"log_file=false\naccess_log_enable=false\n"), 0600); err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestCentralStartupEarlyErrorsAndRecovery$")
			cmd.Env = append(os.Environ(), "SOGA_CENTRAL_EARLY_EXIT="+tt.mode, "SOGA_CENTRAL_EARLY_CONFIG="+path)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("original zero exit: %v\n%s", err, out)
			}
			if !strings.Contains(string(out), tt.want) {
				t.Fatalf("missing original diagnostic %q: %s", tt.want, out)
			}
		})
	}
}
