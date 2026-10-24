package anytlssettings

import (
	"sync"
	"time"
)

// LwNqe0FPzq @0x188ac60 与三个原闭包：独立 goroutine 等待定时或取消，
// 返回函数仅关闭取消通道一次；原版不保证取消与已就绪定时竞争时取消获胜。
func LwNqe0FPzq(d time.Duration, fn func()) func() {
	timer := time.NewTimer(d)
	cancel := make(chan struct{})
	go func() {
		defer timer.Stop()
		select {
		case <-timer.C:
			fn()
		case <-cancel:
		}
	}()
	var once sync.Once
	return func() { once.Do(func() { close(cancel) }) }
}

// ClientVersionAt5098D50 来自 init @0x188abe0 解码器，非作者原变量名。
var ClientVersionAt5098D50 = "anytls/0.0.11"
