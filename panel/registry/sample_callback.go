package registry

// HlayUCRXc 对应所提供 ELF 的 0x16082a0。其入口五字节为
// 31 c0 31 db c3，即清零 error 的两个寄存器后直接返回；0x16082a5
// 后仍存在不可达状态机。这里忠实恢复该样本的可达行为，不表示官方
// 未修改版本的行为，也不改变其他检查路径。样本 SHA-256：
// 6660dcfbc7971188dcb751b440f4269aa718491b569d4c8629b751e6e339f92a。
func HlayUCRXc() error { return nil }

// ViXdEUtm@0x1624cc0 和 Bj7eIYADjFV@0x162ece0 已分别读取原 ELF：
// 两者入口同样为 31 c0 31 db c3；均保留样本实际早返回行为。
func ViXdEUtm() error    { return nil }
func Bj7eIYADjFV() error { return nil }

// 以下四个入口也分别通过原文件读取核对，五字节均为 31 c0 31 db c3。
// 地址分别为 0x15d7700、0x15eb960、0x1614220、0x161d500。
func Aod2QUfu() error     { return nil }
func WIkyLxzu() error     { return nil }
func CWcYaB_UjkSN() error { return nil }
func EcfkJGx1ZPwc() error { return nil }
