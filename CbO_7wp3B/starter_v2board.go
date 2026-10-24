//go:build linux

package CbO_7wp3B

import (
	"time"

	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/q5bqCgTdxaTM"
	"recovered.local/soga/uLofn6j"
)

// lag2OyJwFxd 对应 0x1fc6980。三个接收者各处理一次队列元素，
// 而不是持续循环。初始元素为 0；后续元素由处理函数启动的 goroutine 发送。
func lag2OyJwFxd(count int) {
	runPanelStarterCallback(count, v2boardCallbackAttempt)
}

// 已分别核对 v2board、xboard、xiaov2board 主回调的两个 select 路径。
func runPanelStarterCallback(count int, attempt func(*int, chan int, chan struct{})) {
	MFFa9yT.SetStarterUserCount(count)
	done := make(chan struct{}, 1)
	queue := make(chan int, 1)
	queue <- 0
	go starterCallbackReceiver(queue, done, attempt)
	go starterCallbackReceiver(queue, done, attempt)
	go starterCallbackReceiver(queue, done, attempt)
	timer := time.NewTimer(time.Duration(600+uLofn6j.Co6D_zfls(600)) * time.Second)
	defer timer.Stop()
	select {
	case <-done:
	case <-timer.C:
		cNT24rG4F.T15Jifo(true)
	}
}

// 原三个接收闭包：0x2005b00、0x2005400、0x20046a0。
// 各自的混淆状态机经独立机码轨迹核对，均为一次接收后启动一个处理闭包。
func v2boardCallbackReceiver(queue chan int, done chan struct{}) {
	starterCallbackReceiver(queue, done, v2boardCallbackAttempt)
}

func starterCallbackReceiver(queue chan int, done chan struct{}, attempt func(*int, chan int, chan struct{})) {
	value := <-queue
	go attempt(&value, queue, done)
}

// 三个处理闭包 0x2012b80、0x20135e0、0x2014020 的业务路径相同。
// 重复阈值计算会读取随机数，不能合并；失败路径保留，尽管提供的样本中
// HlayUCRXc 入口直接返回 nil（见该函数的原字节注释）。
func v2boardCallbackAttempt(value *int, queue chan int, done chan struct{}) {
	starterCallbackAttempt(value, queue, done, h68JK8g, q5bqCgTdxaTM.HlayUCRXc)
}

func starterCallbackAttempt(value *int, queue chan int, done chan struct{}, threshold func([]byte) int, check func() error) {
	var random [32]byte
	uLofn6j.Lc_51rU7L(random[:])
	if *value < threshold(random[:]) {
		starterCallbackCheck(value, queue, random[:], threshold, check)
	} else if *value < threshold(random[:]) {
		starterCallbackCheck(value, queue, random[:], threshold, check)
	} else {
		go v2boardCallbackRandom(queue)
	}
	select {
	case done <- struct{}{}:
	default:
	}
}

func starterCallbackCheck(value *int, queue chan int, random []byte, threshold func([]byte) int, check func() error) {
	if err := check(); err != nil {
		if *value == threshold(random)-1 {
			muc279Mpa6S4.FcL3BVMFuVAn(err)
			cNT24rG4F.T15Jifo(true)
			cNT24rG4F.G_G1rBA("")
			go v2boardCallbackRandom(queue)
		} else {
			go func() {
				time.Sleep(time.Second)
				queue <- *value + 1
			}()
		}
	} else {
		cNT24rG4F.T15Jifo(false)
		go v2boardCallbackRandom(queue)
	}
}

func v2boardCallbackRandom(queue chan int) { queue <- int(uLofn6j.Ctna67jV()) }
