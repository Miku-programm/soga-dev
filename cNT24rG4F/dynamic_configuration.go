// 根据 IDA 构建函数、内联调用树及类型描述符恢复，非作者文件原文。
package cNT24rG4F

import (
	"log"
	"strconv"
	"time"

	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/tkbUwvW3ex"
)

// HwIE6xOrX1OH 对应类型描述符 0x36e96b8，amd64 大小 72 字节。
// 两个原 xspLdyHi.YV5V4fmUeHED 字段是 time.Duration 的混淆名字。
type HwIE6xOrX1OH struct {
	DiN7MxyB7  bool
	BgrTZk     tkbUwvW3ex.ITMfakhJ
	Vs46JphL   time.Duration
	Yt1TMC2    float64
	M6lj3M4NlG float64
	ZxdryEoK   time.Duration
	GK1kRuhBCq *G7A7csgrs62p.Fw_lgKgam_jg[int64]
}

var dynamicConfigurationAt5095AB0 *HwIE6xOrX1OH

// oatbDNUZrild 对应 0x1326360；时段解析错误在其余配置转换之前返回。
func oatbDNUZrild() (*HwIE6xOrX1OH, error) {
	c := &configAt509E2A0
	ranges, err := tkbUwvW3ex.Z4FKfucI_A5(c.CGGYGi0)
	if err != nil {
		return nil, err
	}
	triggerTime := time.Duration(c.PLAo2k09hxd) * time.Second
	if triggerTime <= 0 {
		triggerTime = 60 * time.Second
	} else if triggerTime <= 5*time.Second {
		triggerTime = 5 * time.Second
	}
	triggerSpeed := c.ILhRYq68NSUC
	if triggerSpeed <= 0 {
		triggerSpeed = 100
	}
	limitSpeed := c.LxiE3fm4BLS
	if limitSpeed <= 0 {
		limitSpeed = 30
	}
	limitTime := time.Duration(c.Zao5joI9JY) * time.Second
	whiteIDs := G7A7csgrs62p.PQQKui[int64]()
	if limitTime <= 0 {
		limitTime = 600 * time.Second
	}
	for _, text := range c.IrsBkXPlx {
		if value, err := strconv.ParseInt(text, 10, 64); err == nil {
			whiteIDs.Add(value)
		}
	}
	return &HwIE6xOrX1OH{
		DiN7MxyB7: c.EFqkTZF6Wdql, BgrTZk: ranges, Vs46JphL: triggerTime,
		Yt1TMC2: triggerSpeed, M6lj3M4NlG: limitSpeed, ZxdryEoK: limitTime, GK1kRuhBCq: whiteIDs,
	}, nil
}

// vlfjWv1b4Bbg 对应 0x13265c0；原错误分支调用标准 log.Fatal 并终止进程。
func vlfjWv1b4Bbg() {
	value, err := oatbDNUZrild()
	if err != nil {
		log.Fatal(err)
	}
	dynamicConfigurationAt5095AB0 = value
}
