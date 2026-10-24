//go:build linux

package CbO_7wp3B

import "recovered.local/soga/q5bqCgTdxaTM"

// qu_iEPVMzK9e@0x1fce2c0 和 uaxpQZ@0x1fd4940 的主路径及超时分支
// 已分别模拟核对。六个工作闭包也逐个核对为一次接收后启动一次处理。
func qu_iEPVMzK9e(count int) { runPanelStarterCallback(count, xboardCallbackAttempt) }
func uaxpQZ(count int)       { runPanelStarterCallback(count, xiaov2boardCallbackAttempt) }

// 对应 xboard 三个处理函数 0x2014a80、0x20154e0、0x2015f20。
func xboardCallbackAttempt(value *int, queue chan int, done chan struct{}) {
	starterCallbackAttempt(value, queue, done, kCDpmsi, q5bqCgTdxaTM.ViXdEUtm)
}

// 对应 xiaov2board 三个处理函数 0x2016980、0x20173e0、0x2017e20。
func xiaov2boardCallbackAttempt(value *int, queue chan int, done chan struct{}) {
	starterCallbackAttempt(value, queue, done, zJOcPcQ33, q5bqCgTdxaTM.Bj7eIYADjFV)
}

// kCDpmsi@0x1fca220 与 zJOcPcQ33@0x1fd0060 的混淆状态不同，
// 但数据计算相同：三组随机数夹具均逐轮核对全部 32 次除法的被除数和
// 除数，并核对空输入、时间奇偶路径及随机读取次数，因此共享已恢复算术。
func kCDpmsi(input []byte) int   { return h68JK8g(input) }
func zJOcPcQ33(input []byte) int { return h68JK8g(input) }
