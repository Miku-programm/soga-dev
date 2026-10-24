package cNT24rG4F

import (
	"strings"
	"sync/atomic"

	"recovered.local/soga/xx5fRyrFw"
)

// 原静态字符串 global 0x5029ee0 指向 "2.17.0"（6 字节）。
var versionAt5029EE0 = "2.17.0"
var blockListAt5095A90 atomic.Pointer[xx5fRyrFw.VHWq49iRat]
var whiteListAt5095AD8 atomic.Pointer[xx5fRyrFw.VHWq49iRat]

// 恢复原 init.0—init.27 的实际行为与顺序。
// 默认 init@0x1319080 另有外部全局初始化尚待恢复；这里不代表整个配置包完成。
func init() { initializeRegisteredConfigurations() }

func initializeRegisteredConfigurations() {
	init0Recovered()
	init1Recovered()
	init2Recovered()
	init3Recovered()
	init4Recovered()
	init5Recovered()
	init6Recovered()
	init7Recovered()
	init8Recovered()
	init9Recovered()
	init10Recovered()
	init11Recovered()
	init12Recovered()
	init13Recovered()
	init14Recovered()
	init15Recovered()
	init16Recovered()
	init17Recovered()
	init18Recovered()
	init19Recovered()
	init20Recovered()
	init21Recovered()
	init22Recovered()
	init23Recovered()
	init24Recovered()
	init25Recovered()
	init26Recovered()
	init27Recovered()
}

// init0Recovered 对应 cNT24rG4F.init.0，地址 0x131a1e0。
func init0Recovered() {
	_NA1kft(&configAt509B660)
	afterLoadAt5099300 = append(afterLoadAt5099300, init0func1)
	rules := xx5fRyrFw.VHWq49iRat{}
	blockListAt5095A90.Store(&rules)
}

// init1Recovered 对应 cNT24rG4F.init.1，地址 0x131c8e0。
func init1Recovered() {
	_NA1kft(&configAt50A2260)
	afterLoadAt5099300 = append(afterLoadAt5099300, uQm8ubPKg)
}

// init2Recovered 对应 cNT24rG4F.init.2，地址 0x1320fa0。
func init2Recovered() {
	_NA1kft(&configAt50992E0)
}

// init3Recovered 对应 cNT24rG4F.init.3，地址 0x1321020。
func init3Recovered() {
	_NA1kft(&nodeConfigurationAt50A2AE0)
	afterLoadAt5099300 = append(afterLoadAt5099300, jkMcB1_)
}

// init4Recovered 对应 cNT24rG4F.init.4，地址 0x1322ee0。
func init4Recovered() {
	versionAt5029EE0 = strings.TrimSpace(versionAt5029EE0)
	_NA1kft(&configAt5096DC0)
}

// init5Recovered 对应 cNT24rG4F.init.5，地址 0x13260e0。
func init5Recovered() {
	_NA1kft(&configAt509F9C0)
}

// init6Recovered 对应 cNT24rG4F.init.6，地址 0x1326120。
func init6Recovered() {
	_NA1kft(&configAt509CB40)
	afterLoadAt5099300 = append(afterLoadAt5099300, init6func1)
}

// init7Recovered 对应 cNT24rG4F.init.7，地址 0x13261e0。
func init7Recovered() {
	_NA1kft(&configAt509E2A0)
	afterLoadAt5099300 = append(afterLoadAt5099300, vlfjWv1b4Bbg)
}

// init8Recovered 对应 cNT24rG4F.init.8，地址 0x1326640。
func init8Recovered() {
	_NA1kft(&configAt50C8440)
}

// init9Recovered 对应 cNT24rG4F.init.9，地址 0x1326680。
func init9Recovered() {
	_NA1kft(&configAt509B380)
	afterLoadAt5099300 = append(afterLoadAt5099300, mvo6TVLM)
}

// init10Recovered 对应 cNT24rG4F.init.10，地址 0x1326f80。
func init10Recovered() {
	_NA1kft(&configAt509B6A0)
	afterLoadAt5099300 = append(afterLoadAt5099300, init10func1)
}

// init11Recovered 对应 cNT24rG4F.init.11，地址 0x1329520。
func init11Recovered() {
	_NA1kft(&configAt50C6C49)
}

// init12Recovered 对应 cNT24rG4F.init.12，地址 0x1329560。
func init12Recovered() {
	_NA1kft(&configAt509BEA0)
	afterLoadAt5099300 = append(afterLoadAt5099300, qZO6G3A)
}

// init13Recovered 对应 cNT24rG4F.init.13，地址 0x132a160。
func init13Recovered() {
	_NA1kft(&configAt50A2900)
	afterLoadAt5099300 = append(afterLoadAt5099300, tzfa9ae)
}

// init14Recovered 对应 cNT24rG4F.init.14，地址 0x132a560。
func init14Recovered() {
	_NA1kft(&configAt509B3A0)
	afterLoadAt5099300 = append(afterLoadAt5099300, init14func1)
}

// init15Recovered 对应 cNT24rG4F.init.15，地址 0x132a620。
func init15Recovered() {
	_NA1kft(&configAt509B3C0)
	afterLoadAt5099300 = append(afterLoadAt5099300, init15func1)
}

// init16Recovered 对应 cNT24rG4F.init.16，地址 0x132a6e0。
func init16Recovered() {
	_NA1kft(&configAt509C940)
	afterLoadAt5099300 = append(afterLoadAt5099300, ekZtUL9pJS)
}

// init17Recovered 对应 cNT24rG4F.init.17，地址 0x132a7c0。
func init17Recovered() {
	_NA1kft(&configAt509CB80)
	afterLoadAt5099300 = append(afterLoadAt5099300, mQ5mMKnkLjl)
}

// init18Recovered 对应 cNT24rG4F.init.18，地址 0x132db00。
func init18Recovered() {
	_NA1kft(&configAt5096DD0)
}

// init19Recovered 对应 cNT24rG4F.init.19，地址 0x132db40。
func init19Recovered() {
	_NA1kft(&configAt509B3E0)
}

// init20Recovered 对应 cNT24rG4F.init.20，地址 0x132db80。
func init20Recovered() {
	_NA1kft(&configAt5099340)
	afterLoadAt5099300 = append(afterLoadAt5099300, rHAcgy)
}

// init21Recovered 对应 cNT24rG4F.init.21，地址 0x132ea00。
func init21Recovered() {
	_NA1kft(&configAt509B6E0)
}

// init22Recovered 对应 cNT24rG4F.init.22，地址 0x132ea40。
func init22Recovered() {
	_NA1kft(&configAt509EBE0)
}

// init23Recovered 对应 cNT24rG4F.init.23，地址 0x1330000。
func init23Recovered() {
	_NA1kft(&configAt5099360)
}

// init24Recovered 对应 cNT24rG4F.init.24，地址 0x1330040。
func init24Recovered() {
	_NA1kft(&configAt509E300)
}

// init25Recovered 对应 cNT24rG4F.init.25，地址 0x1333040。
func init25Recovered() {
	_NA1kft(&configAt50A02C0)
	afterLoadAt5099300 = append(afterLoadAt5099300, eifTpu4B)
}

// init26Recovered 对应 cNT24rG4F.init.26，地址 0x1335460。
func init26Recovered() {
	_NA1kft(&configAt50A2300)
}

// init27Recovered 对应 cNT24rG4F.init.27，地址 0x133e520。
func init27Recovered() {
	_NA1kft(&configAt5096E00)
	rules := xx5fRyrFw.VHWq49iRat{}
	whiteListAt5095AD8.Store(&rules)
}
