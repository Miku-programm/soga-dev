// Package xIF4G3 是原 soga ELF 中 BBR 拥塞控制的源码表示。
// 它依据实际函数机器码、Go 类型元数据和字符串解码结果恢复，不是作者文件原文。
// pclntab 包名为 xIF4G3，RTTI 显示 ykyDTH6hodhi；本目录采用前者，保留已知混淆类型和字段名。
// 对照来源为 apernet/hysteria 9e90d7d 的 congestion/bbr，并逐函数核对原样本。
// 未单独保留符号但在调用者中已核准的内联片段，沿用对照源码的辅助函数名称；
// 这些辅助函数和 NewSenderFromInline 不额外计为原独立函数。
// bandwidthSample 未保留独立 RTTI，按 nNZK6dj0MNti 的实际返回值布局表达。
// 原样本中若干 SendAlgorithm 方法本来只有 RET；本包保留这些空操作，未以它们替代未知逻辑。
// 指令差异、栈槽映射、常量、直接调用和受限解码证据位于 work/ida-20261023/hysteria-wire。
package xIF4G3
