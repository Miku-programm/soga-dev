package CAodbhZf

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLiteralPathSlicing(t *testing.T) {
	for _, tc := range []struct{ path, want string }{
		{"", ""}, {"file", ""}, {"/", ""}, {"/a/b", "/a/"},
		{"/a/b/", "/a/"}, {"/a/b//", "/a/b/"}, {"a/../b", "a/../"},
		{`C:\dir\file`, ""}, {"a//b", "a//"},
	} {
		if got := EVZvjBCf(tc.path); got != tc.want {
			t.Errorf("directory(%q) = %q, want %q", tc.path, got, tc.want)
		}
	}
}

func TestStatExistence(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "file")
	if VSCJLy2No(file) || !VSCJLy2No(dir) {
		t.Fatal("missing path or directory handled incorrectly")
	}
	if err := os.WriteFile(file, nil, 0600); err != nil {
		t.Fatal(err)
	}
	if !VSCJLy2No(file) || VSCJLy2No(filepath.Join(file, "child")) {
		t.Fatal("file or ENOTDIR handled incorrectly")
	}
}
