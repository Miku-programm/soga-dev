package geo

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type failingGeoReader struct {
	read bool
	err  error
}

func (r *failingGeoReader) Read(data []byte) (int, error) {
	if r.read {
		return 0, r.err
	}
	r.read = true
	return copy(data, "partial"), r.err
}

// 只暴露 Reader，确保成功路径也覆盖 CopyBuffer 的普通读取分支。
type plainGeoReader struct{ io.Reader }

func TestGeoUpdateFileReplacementAndErrors(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "geosite.dat")
	if err := os.WriteFile(path, []byte("previous"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := kBh3sTD66VJD(plainGeoReader{strings.NewReader("updated data")}, path); err != nil {
		t.Fatal(err)
	}
	if data, err := os.ReadFile(path); err != nil || string(data) != "updated data" {
		t.Fatalf("replacement %q %v", data, err)
	}
	if _, err := os.Stat(path + ".tmp"); !os.IsNotExist(err) {
		t.Fatalf("temporary rename %v", err)
	}
	if err := kBh3sTD66VJD(strings.NewReader(""), path); err == nil || err.Error() != "download geosite data failed: empty file" {
		t.Fatalf("empty error %v", err)
	}
	if stat, err := os.Stat(path + ".tmp"); err != nil || stat.Size() != 0 {
		t.Fatalf("empty temporary removed: %v %v", stat, err)
	}
	readError := errors.New("original reader error")
	if err := kBh3sTD66VJD(&failingGeoReader{err: readError}, path); err != readError {
		t.Fatalf("reader error wrapped: %v", err)
	}
	if data, err := os.ReadFile(path + ".tmp"); err != nil || string(data) != "partial" {
		t.Fatalf("partial file %q %v", data, err)
	}
	if data, err := os.ReadFile(path); err != nil || string(data) != "updated data" {
		t.Fatalf("failed update replaced destination: %q %v", data, err)
	}
	if err := kBh3sTD66VJD(strings.NewReader("data"), filepath.Join(root, "missing", "geoip.dat")); err == nil || !strings.HasPrefix(err.Error(), "open geosite file failed:") {
		t.Fatalf("open error %v", err)
	}
	directory := filepath.Join(root, "directory")
	if err := os.Mkdir(directory, 0700); err != nil {
		t.Fatal(err)
	}
	if err := kBh3sTD66VJD(strings.NewReader("data"), directory); err == nil || !strings.HasPrefix(err.Error(), "rename geosite file failed:") {
		t.Fatalf("rename error %v", err)
	}
	if _, err := os.Stat(directory + ".tmp"); !os.IsNotExist(err) {
		t.Fatalf("rename failure did not remove temporary %v", err)
	}
}
