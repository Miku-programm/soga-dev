package BP3PTlx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	C "recovered.local/soga/cNT24rG4F"
	"testing"
)

func TestStatusUsesMemoryCountersAndContinuesAfterSampleFailure(t *testing.T) {
	proc := t.TempDir()
	t.Setenv("HOST_PROC", proc)
	// stat 和 meminfo 来自独立固定样本，重复采样保持 CPU 计数维度。
	if err := os.WriteFile(filepath.Join(proc, "stat"), []byte("cpu 0 0 0 0 0 0 0 0 0 0\n"), 0600); err != nil {
		t.Fatal(err)
	}
	const meminfo = "MemTotal: 1000 kB\nMemFree: 100 kB\nMemAvailable: 700 kB\nBuffers: 20 kB\nCached: 30 kB\nSReclaimable: 50 kB\n"
	if err := os.WriteFile(filepath.Join(proc, "meminfo"), []byte(meminfo), 0600); err != nil {
		t.Fatal(err)
	}
	records := make(chan map[string]json.RawMessage, 2)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]json.RawMessage
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		records <- body
	}))
	defer server.Close()
	p := NiF7OSBK(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}).(*KlO65BuwOt)
	if err := p.gKXqCyN(); err != nil {
		t.Fatal(err)
	}
	body := <-records
	if string(body["cpu"]) != "0" || string(body["mem"]) != `{"used":819200,"total":1024000}` {
		t.Fatalf("%s / %s", body["cpu"], body["mem"])
	}
	if err := os.WriteFile(filepath.Join(proc, "meminfo"), []byte("MemTotal: bad kB\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := p.gKXqCyN(); err != nil {
		t.Fatal(err)
	}
	body = <-records
	if string(body["mem"]) != `{"used":0,"total":0}` || body["swap"] == nil || body["disk"] == nil {
		t.Fatal(body)
	}
}
