package inbound

import (
	"os"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/config"
	"recovered.local/soga/crypto/rand"
	"time"
)

// bj10P7X@1de9fa0及sbbdX3aT6E@1dead60按各自12个init种子独立静态仿真恢复。
// 保留uint32回绕、3600秒严格大于、CAS不重试和共享非原子用户数写入。
func (v *KHonGaqbIy_6) bj10P7X() {
	now := uint32(rtx.C5Jff5vs0())
	userCountAt50C6DA0 = v.fUxO7UYAH
	if claimConnectionStateCheck(&v.jfLH8mvq5e03, now) {
		go v.checkVLESSConnectionState()
	}
}
func (v *KHonGaqbIy_6) sbbdX3aT6E() {
	now := uint32(rtx.C5Jff5vs0())
	userCountAt50C6DA0 = v.fUxO7UYAH
	if claimConnectionStateCheck(&v.jfLH8mvq5e03, now) {
		go v.checkVLESSConnectionState()
	}
}

// 四原闭包1ea4660/1ea6080/1ea7c80/1ea96a0独立验证后共用相同响应算法；
// 实际授权请求仍为原dcHP1aBmqOSa，保留全部原失败、删文件及退出条件。
func (v *KHonGaqbIy_6) checkVLESSConnectionState() {
	time.Sleep(time.Duration(rand.Co6D_zfls(60)) * time.Second)
	connectionStateResponse(&v.mKB02yhVfTfP, config.FrFQQhzTR5fY, dcHP1aBmqOSa, config.T15Jifo, os.RemoveAll, os.Remove, os.Exit)
}
