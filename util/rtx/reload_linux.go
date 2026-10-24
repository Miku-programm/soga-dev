//go:build linux

package rtx

import "syscall"

// Jx2ev8z 对应 0xb4cf40：向当前进程发送 SIGUSR1，调用方的信号处理器执行重载。
func Jx2ev8z() error { return syscall.Kill(syscall.Getpid(), syscall.SIGUSR1) }
