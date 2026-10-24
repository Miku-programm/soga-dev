package app

import (
	"recovered.local/soga/net/routing"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"runtime"
)

// l24UlWt restores 0x1f9a9a0. Preserve the load, watcher, GC and logging order;
// the target's fatal logger exits with status zero on configuration failure.
func (app *M0mmPxPJsa) l24UlWt() {
	if err := config.TvBECnD9fM6(); err != nil {
		logx.Wa0FarWk("load config <%v> failed: %v", config.ConfigurationPath(), err)
	}
	logx.V8GN5Stk(0, "load config <%v> success", config.ConfigurationPath())
	config.LnLDMWbjk3()
	config.I5PWiaSBwdar()
	config.WreRtju(app.spsw8VktR9z)
	routing.AycgFFCBT5Zx(app.spsw8VktR9z)
	runtime.GC()
	config.W2uyDtOBN2()
	logx.TRTbrquD(0, "is docker:", config.VSHlrRuK7_())
	config.MeWz428lV2Z()
}
