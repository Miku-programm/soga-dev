package xiaov2boardhysteriawebapi

import (
	"recovered.local/soga/panel/xiaov2board"
	"recovered.local/soga/panel/registry"
)

// 原 xiaov2boardhysteriawebapi.init.0 对应 0x17b0100，func 值 0x377B700 → 0x17a5b00。
func init() { registry.PanelFactoriesAt5095E50["xiaov2board-hysteria-webapi"] = xiaov2board.AnxiJP }
