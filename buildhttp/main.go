// Command buildhttp builds or tests with the evidenced HTTP/1 server overlay.
// This is a reconstruction build tool, not a function from the supplied binary.
package main

import (
	"crypto/sha256"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// The source file retains Go's copyright notice and license. Its only edits are
// four recovered differences in chunkWriter.writeHeader at original 0xa4d040.
// This does not assert that unrelated private HTTP-fork functions match Go.
//
//go:embed testdata/server.go
var serverSource []byte

const referenceServerSHA256 = "0a4886872338cde6d0d4cf04ccd53a2b1debad33551d2e77b37116ae2b422709"

func run() error {
	goBinary := flag.String("go", "go", "Go 1.27.3 executable")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"build", "./..."}
	}
	if args[0] != "build" && args[0] != "test" {
		return fmt.Errorf("expected build or test, got %q", args[0])
	}
	for _, arg := range args[1:] {
		if arg == "-overlay" || strings.HasPrefix(arg, "-overlay=") {
			return errors.New("an additional overlay must be merged explicitly")
		}
	}
	rootBytes, err := exec.Command(*goBinary, "env", "GOROOT").Output()
	if err != nil {
		return fmt.Errorf("locate selected Go root: %w", err)
	}
	sourcePath := filepath.Join(strings.TrimSpace(string(rootBytes)), "src", "net", "http", "server.go")
	original, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	if got := fmt.Sprintf("%x", sha256.Sum256(original)); got != referenceServerSHA256 {
		return fmt.Errorf("unsupported net/http/server.go SHA256 %s; expected verified Go 1.27.3 source %s", got, referenceServerSHA256)
	}
	temporary, err := os.MkdirTemp("", "soga-http-overlay-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(temporary)
	replacement := filepath.Join(temporary, "server.go")
	if err := os.WriteFile(replacement, serverSource, 0600); err != nil {
		return err
	}
	overlay, err := json.Marshal(struct{ Replace map[string]string }{map[string]string{sourcePath: replacement}})
	if err != nil {
		return err
	}
	overlayPath := filepath.Join(temporary, "overlay.json")
	if err := os.WriteFile(overlayPath, overlay, 0600); err != nil {
		return err
	}
	childArgs := append([]string{args[0], "-overlay", overlayPath}, args[1:]...)
	child := exec.Command(*goBinary, childArgs...)
	child.Stdin = os.Stdin
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	return child.Run()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		var exit *exec.ExitError
		if errors.As(err, &exit) {
			os.Exit(exit.ExitCode())
		}
		os.Exit(1)
	}
}
