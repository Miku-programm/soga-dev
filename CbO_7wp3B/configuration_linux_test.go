package CbO_7wp3B

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"recovered.local/soga/cNT24rG4F"
)

func TestConfigurationFailureRetainsOriginalZeroExit(t *testing.T) {
	if path := os.Getenv("SOGA_TEST_STARTUP_BAD_PATH"); path != "" {
		cNT24rG4F.MkwNWciPHS(path)
		app := &M0mmPxPJsa{spsw8VktR9z: context.Background()}
		app.l24UlWt()
		os.Exit(93)
	}
	dir := t.TempDir()
	file := filepath.Join(dir, "regular-file")
	if err := os.WriteFile(file, []byte("x"), 0600); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestConfigurationFailureRetainsOriginalZeroExit$")
	cmd.Env = append(os.Environ(), "SOGA_TEST_STARTUP_BAD_PATH="+filepath.Join(file, "soga.conf"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("original zero exit changed: %v\n%s", err, output)
	}
	if !strings.Contains(string(output), "load config <") || !strings.Contains(string(output), "failed:") {
		t.Fatalf("missing diagnostic: %s", output)
	}
	if strings.Contains(string(output), "success") {
		t.Fatalf("startup continued after failure: %s", output)
	}
}
