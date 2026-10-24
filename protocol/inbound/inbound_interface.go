package inbound

import "recovered.local/soga/panel/registry"

// Kz_8ajU7 原接口RTTI0x36214f0，方法表0x3621550共5项。
// 定义只表达原方法集合，不表示每种协议已完成恢复。
type Kz_8ajU7 interface {
	GetSubmitData() *registry.FGmNba4MhioY
	Start() error
	Stop()
	UpdateAPIData(*registry.D5gTIZ5aVg) bool
	UserExist(int64) bool
}

// fH6caiUi 为原工厂注册表的命名构造函数类型。
type fH6caiUi func(*DV_XPJ4z0) (Kz_8ajU7, error)
