package config

import (
	"go.uber.org/dig"
	"log"
)

// Bnm1qf 对应 0x13354a0；依赖注册失败时原实现调用标准库 log.Fatal。
func Bnm1qf(constructor any, options ...dig.ProvideOption) {
	if err := containerAt5095AC8.Provide(constructor, options...); err != nil {
		log.Fatal(err)
	}
}

// DependencyContainer 是分析命名的访问桥，对应原跨包直接读取 0x5095ac8。
func DependencyContainer() *dig.Container { return containerAt5095AC8 }
