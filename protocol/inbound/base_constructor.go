package inbound

import (
	"context"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"recovered.local/soga/panel/registry"
	"recovered.local/soga/net/sniff"
)

// DV_XPJ4z0 为原RTTI的40字节入站构造参数；类型显示前缀oUi8Xtk3yv，
// uncommon_package_path为MFFa9yT。源码由二进制证据重建，不是原作者文件。
type DV_XPJ4z0 struct {
	Dmpl48sXM  context.Context
	POpyjFWF8v config.Mpn2JltM
	HG5k0365   *registry.D5gTIZ5aVg
}

// n2s6Gf9ey 对应0x1d3e800，包括真实动态限速管理器和公共状态初始化。
func n2s6Gf9ey(arg *DV_XPJ4z0) (*JWN154qJ4, error) {
	users := &user.ZpNFqUVZp{GTFAWF: &genmap.GTFAWF[int64, user.PgaNd3014U]{}}
	ctx, cancel := context.WithCancel(arg.Dmpl48sXM)
	limiter := &relay.L337vk2G{}
	sessions := &sess.TaQ7Prswa{}
	reports := &genmap.GTFAWF[int, *genmap.Fxs7SY[int64]]{}
	addr, port := config.G9UvxRacCj(), config.QnSmghY()
	manager := limit.DMIy6PEVB(ctx, users)
	b := &JWN154qJ4{zb8f3cHWW8v: ctx, yPUbKO: cancel, dwfzZM: limiter, jc_QPN6: users, hS8xXFfcBr: sessions, qE9aU9: reports, dZfoV0aXHW: addr, lAmttzTO8b: port, cEUo2gYKI: sniff.KZgBMVcpL{}, oSsZZT4EmeG: manager}
	if err := b.init(); err != nil {
		return nil, err
	}
	return b, nil
}
