package AKP5mbedNKj

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

type closeOrderWriter struct {
	events   *[]string
	finished bool
}

func (*closeOrderWriter) Header() http.Header         { return make(http.Header) }
func (*closeOrderWriter) Write(p []byte) (int, error) { return len(p), nil }
func (*closeOrderWriter) WriteHeader(int)             {}
func (w *closeOrderWriter) Flush() {
	*w.events = append(*w.events, "flush")
	if w.finished {
		panic("Header called after Handler finished")
	}
}

type closeOrderBody struct{ events *[]string }

func (*closeOrderBody) Read([]byte) (int, error) { return 0, io.EOF }
func (b *closeOrderBody) Close() error           { *b.events = append(*b.events, "body.close"); return nil }

// 原 0x18b528b 先调用取消回调，0x18b52e0 再 Flush，0x18b52f7 关闭 Body。
// 原代码不捕获 Flush panic；此测试明确保留这一行为，不伪造安全关闭语义。
func TestOriginalCloseCallbackFlushBodyOrder(t *testing.T) {
	for _, finishBeforeFlush := range []bool{false, true} {
		var events []string
		w := &closeOrderWriter{events: &events}
		c := &IwDQb32fnlH9{oC5aEPzrec9w: w, px3ua9JUg: &http.Request{Body: &closeOrderBody{&events}}}
		c.g2MwZxLFm_YD = func() { events = append(events, "cancel"); w.finished = finishBeforeFlush }
		var panicked any
		func() { defer func() { panicked = recover() }(); _ = c.Close() }()
		want := []string{"cancel", "flush", "body.close"}
		if finishBeforeFlush {
			want = want[:2]
			if panicked != "Header called after Handler finished" {
				t.Fatalf("original propagated panic: %v", panicked)
			}
		} else if panicked != nil {
			t.Fatal(panicked)
		}
		if !reflect.DeepEqual(events, want) || !c.fMRTh247ym2l {
			t.Fatal(events, c.fMRTh247ym2l)
		}
		if !c.TryLock() {
			t.Fatal("deferred unlock lost")
		}
		c.Unlock()
		if err := c.Close(); err != nil || !reflect.DeepEqual(events, want) {
			t.Fatal("second close", events, err)
		}
	}
}
