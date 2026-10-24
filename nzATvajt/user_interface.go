package nzATvajt

import (
	"context"
	"recovered.local/soga/LFWTTXF7JR"
)

// PgaNd3014U：接口 RTTI0x36e5d80，PkgPath=nzATvajt，原类型前缀ocim0HtS。
// 19个方法名与签名直接取自原接口方法表；此处没有替代用户实现。
type PgaNd3014U interface {
	AddAliveIp(string, string) bool
	AddDown(int)
	AddTCP() bool
	AddUp(int)
	AllowAndAddDown(int) bool
	AllowAndAddUp(int) bool
	DelAliveIp(string, string)
	DelTCP()
	GetAliveIpData() *LFWTTXF7JR.LmNKoY94
	GetAndClearDyTraffic() (int64, int64)
	GetAndClearTraffic() (int64, int64)
	GetConnLimit() int
	GetCtx() context.Context
	GetId() int64
	GetSpeedLimit() float64
	IsTrafficOut() bool
	RecordAliveIp(string)
	ResetCtx()
	SetSpeedLimit(float64)
}
