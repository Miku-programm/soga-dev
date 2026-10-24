package xiaov2boardvlesswebapi

import (
	"recovered.local/soga/panel/xiaov2board"
	"recovered.local/soga/panel/registry"
)

// 原 xiaov2boardvlesswebapi.init.0 对应 0x17b0340，func 值 0x377B700 → 0x17a5b00。
func init() { registry.PanelFactoriesAt5095E50["xiaov2board-vless-webapi"] = xiaov2board.AnxiJP }
