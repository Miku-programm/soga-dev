package rtx

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func TestMappedFileLayoutAndReload(t *testing.T) {
	mapped := &CCVciOvWaCEr{mnsdOl0CZUz: -1}
	if unsafe.Sizeof(*mapped) != 56 || unsafe.Offsetof(mapped.rVSeQKqr) != 24 || unsafe.Offsetof(mapped.mnsdOl0CZUz) != 48 {
		t.Fatal("RTTI layout differs")
	}
	if data, err := mapped.Data(); data != nil || err == nil || err.Error() != "data is not initialized" {
		t.Fatalf("uninitialized: %v %v", data, err)
	}
	path := filepath.Join(t.TempDir(), "database")
	if err := os.WriteFile(path, []byte("first database"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := mapped.g9iBYoc(path); err != nil {
		t.Fatal(err)
	}
	data, err := mapped.Data()
	if err != nil || string(data) != "first database" {
		t.Fatalf("mapped %q %v", data, err)
	}
	second, _ := mapped.Data()
	if &data[0] != &second[0] {
		t.Fatal("Data copied the mapped memory")
	}
	replacement := path + ".new"
	if err := os.WriteFile(replacement, []byte("replacement database"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(replacement, path); err != nil {
		t.Fatal(err)
	}
	if err := mapped.g9iBYoc(path); err != nil {
		t.Fatal(err)
	}
	data, err = mapped.Data()
	if err != nil || string(data) != "replacement database" {
		t.Fatalf("reload %q %v", data, err)
	}
	if err := mapped.g9iBYoc(path + ".missing"); err == nil {
		t.Fatal("missing file accepted")
	}
	if data, err := mapped.Data(); data != nil || err == nil {
		t.Fatal("failed reload retained old map")
	}
	if mapped.mnsdOl0CZUz != -1 {
		t.Fatal("failed reload retained descriptor")
	}
	if err := mapped.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestMappedFileZeroLengthAndCloseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty")
	if err := os.WriteFile(path, nil, 0600); err != nil {
		t.Fatal(err)
	}
	mapped := &CCVciOvWaCEr{mnsdOl0CZUz: -1}
	if err := mapped.g9iBYoc(path); err == nil {
		t.Fatal("empty mmap unexpectedly succeeded")
	}
	fd, err := syscall.Open(path, syscall.O_RDONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if err := syscall.Close(fd); err != nil {
		t.Fatal(err)
	}
	mapped.mnsdOl0CZUz = fd
	err = mapped.Close()
	if err == nil || !strings.Contains(err.Error(), "bad file descriptor") {
		t.Fatalf("Close %v", err)
	}
	if _, ok := err.(interface{ Unwrap() error }); ok {
		t.Fatal("original aggregate has no Unwrap")
	}
	if mapped.mnsdOl0CZUz != -1 || mapped.Close() != nil {
		t.Fatal("Close is not idempotent")
	}
}

func TestMappedFileConstructorDescriptorBehavior(t *testing.T) {
	if os.Getenv("SOGA_MMAP_CTOR_CHILD") == "1" {
		path := os.Getenv("SOGA_MMAP_CTOR_PATH")
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		mapped, err := LWzrRP(ctx, path)
		if err != nil {
			t.Fatal(err)
		}
		if mapped.mnsdOl0CZUz != 0 {
			t.Fatalf("original initial Close(0) not preserved: fd=%d", mapped.mnsdOl0CZUz)
		}
		data, err := mapped.Data()
		if err != nil || string(data) != "constructor data" {
			t.Fatalf("data %q, %v", data, err)
		}
		if err := mapped.Close(); err != nil {
			t.Fatal(err)
		}
		return
	}
	path := filepath.Join(t.TempDir(), "database")
	if err := os.WriteFile(path, []byte("constructor data"), 0600); err != nil {
		t.Fatal(err)
	}
	command := exec.Command(os.Args[0], "-test.run=^TestMappedFileConstructorDescriptorBehavior$")
	command.Env = append(os.Environ(), "SOGA_MMAP_CTOR_CHILD=1", "SOGA_MMAP_CTOR_PATH="+path)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("child: %v\n%s", err, output)
	}
}

func TestWatchChangedWriteAndCallbackStop(t *testing.T) {
	path := filepath.Join(t.TempDir(), "watched")
	if err := os.WriteFile(path, []byte("initial"), 0600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	called := make(chan struct{}, 10)
	if err := (&WaVcJUem{}).WatchChanged(ctx, func() bool { called <- struct{}{}; return false }, path); err != nil {
		t.Fatal(err)
	}
	deadline := time.NewTimer(3 * time.Second)
	defer deadline.Stop()
	tick := time.NewTicker(10 * time.Millisecond)
	defer tick.Stop()
	for {
		select {
		case <-called:
			if err := os.WriteFile(path, []byte("after stop"), 0600); err != nil {
				t.Fatal(err)
			}
			select {
			case <-called:
				t.Fatal("callback continued after false")
			case <-time.After(50 * time.Millisecond):
				return
			}
		case <-tick.C:
			if err := os.WriteFile(path, []byte("changed"), 0600); err != nil {
				t.Fatal(err)
			}
		case <-deadline.C:
			t.Fatal("file write produced no callback")
		}
	}
}
