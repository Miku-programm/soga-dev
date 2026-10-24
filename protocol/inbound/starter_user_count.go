package inbound

// SetStarterUserCount 是对 CbO 启动回调写入原共享地址 0x50c6da0 的桥接。
// 与入站限制及请求载荷使用同一全局；原调用未使用原子操作。
func SetStarterUserCount(count int) { userCountAt50C6DA0 = count }
