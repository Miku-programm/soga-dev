package xiaov2boardanytlswebapi

import (
	"recovered.local/soga/panel/xiaov2board"
	"recovered.local/soga/panel/registry"
)

// 原 xiaov2boardanytlswebapi.init.0 对应 0x17b0040，func 值 0x377B700 → 0x17a5b00。
func init() { registry.PanelFactoriesAt5095E50["xiaov2board-anytls-webapi"] = xiaov2board.AnxiJP }
