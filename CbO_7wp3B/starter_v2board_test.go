//go:build linux

package CbO_7wp3B

import (
	"crypto/rand"
	"errors"
	"testing"
	"time"

	"recovered.local/soga/cNT24rG4F"
)

// 原机码的两次比较路径分别消耗 1 或 2 次 h68 的 256 字节，
// 加上最初 32 字节及异步后继发送的 8 字节。队列接收建立完成同步。
func TestV2boardAttemptPreservesEntropyAndSignal(t *testing.T) {
	old := rand.Reader
	t.Cleanup(func() { rand.Reader = old; cNT24rG4F.T15Jifo(false) })
	for _, attempt := range []func(*int, chan int, chan struct{}){v2boardCallbackAttempt, xboardCallbackAttempt, xiaov2boardCallbackAttempt} {
		for _, tc := range []struct{ value, bytes int }{{0, 296}, {2, 296}, {3, 552}, {100, 552}} {
			r := &thresholdEntropy{}
			rand.Reader = r
			queue, done := make(chan int, 1), make(chan struct{}, 1)
			value := tc.value
			attempt(&value, queue, done)
			select {
			case got := <-queue:
				if got != 0 {
					t.Fatalf("successor %d", got)
				}
			case <-time.After(time.Second):
				t.Fatal("successor missing")
			}
			if len(done) != 1 || r.read != tc.bytes {
				t.Fatalf("value=%d: signal=%d entropy=%d, want %d", tc.value, len(done), r.read, tc.bytes)
			}
		}
	}
}

func TestV2boardReceiverConsumesOnlyOneValue(t *testing.T) {
	queue, done := make(chan int, 2), make(chan struct{}, 1)
	queue <- 100
	queue <- 101
	v2boardCallbackReceiver(queue, done)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("worker did not finish")
	}
	if got := <-queue; got != 101 {
		t.Fatalf("receiver consumed second value: %d", got)
	}
	select {
	case <-queue: // Drain the one successor goroutine.
	case <-time.After(time.Second):
		t.Fatal("successor missing")
	}
}

// 失败结果只作为外部调用夹具；不会更改样本的三个真实 nil 返回入口。
// 两条轨迹分别对应原 inner 的延迟重试与最后一次错误处理。
func TestStarterCallbackFailureBranches(t *testing.T) {
	t.Cleanup(func() { cNT24rG4F.T15Jifo(false) })
	for _, value := range []int{0, 2} {
		queue, done := make(chan int, 1), make(chan struct{}, 1)
		done <- struct{}{} // 原 selectnbsend 在通道已满时不能阻塞。
		calls, checks := 0, 0
		threshold := func([]byte) int { calls++; return 3 }
		check := func() error { checks++; return errors.New("fixture failure") }
		start := time.Now()
		starterCallbackAttempt(&value, queue, done, threshold, check)
		select {
		case got := <-queue:
			if value == 0 && (got != 1 || time.Since(start) < time.Second) {
				t.Fatalf("retry=%d elapsed=%s", got, time.Since(start))
			}
		case <-time.After(3 * time.Second):
			t.Fatal("successor missing")
		}
		if calls != 2 || checks != 1 || len(done) != 1 {
			t.Fatalf("threshold=%d checks=%d signals=%d", calls, checks, len(done))
		}
	}
}
