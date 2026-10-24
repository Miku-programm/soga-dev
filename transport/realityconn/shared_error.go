package realityconn

// TimeoutErrorAt502B390 是恢复工程集成辅助，读取原共享 error 对象。
// g0MlUOhvO 的 errors.Is 分支与本包使用同一对象，不能重新创建同文本错误。
// 此桥接不计为已恢复的原始函数。
func TimeoutErrorAt502B390() error { return timeoutAt502B390 }
