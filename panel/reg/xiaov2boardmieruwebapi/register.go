package xiaov2boardmieruwebapi

import (
	"recovered.local/soga/panel/xiaov2board"
	"recovered.local/soga/panel/registry"
)

// 新增：xiaov2b 的 mieru 节点复用 xiaov2board UniProxy 对接。
func init() { registry.PanelFactoriesAt5095E50["xiaov2board-mieru-webapi"] = xiaov2board.AnxiJP }
