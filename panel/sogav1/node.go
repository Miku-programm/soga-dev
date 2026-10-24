// 根据 0x16d5f60 的字段写入和调用恢复。
package sogav1

import (
	"recovered.local/soga/panel/types"
	"recovered.local/soga/net/addrtype"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/ssrnames"
	"recovered.local/soga/panel/registry"
	"recovered.local/soga/crypto/ssmethod"
	"strings"
	"time"
)

func (n *GX5_vLs) GetPullInterval() time.Duration {
	return time.Duration(n.XI5uZkt.H5wL4HQdqP) * time.Second
}
func (n *GX5_vLs) GetPushInterval() time.Duration {
	return time.Duration(n.XI5uZkt.ULhXjJaRde) * time.Second
}
func (n *GX5_vLs) ToPanelNode(id int) (*types.Wr5SVP, error) {
	c := &n.I8ECEmdqGb
	out := new(types.Wr5SVP)
	if c.HAei1i_ != nil {
		out.Ck6DxjB.FPhrzP645 = c.HAei1i_.ToXHttpConfig()
	}
	out.EMlXEc = id
	out.Z1TIgHOQ = registry.T0j_Rco(n)
	out.HSA2b5V = config.PanelForbiddenBitTorrent()
	out.TtWKyh = mergeSpeed(config.PanelNodeSpeedLimit(), n.XI5uZkt.ZZ6bu_Ki)
	out.SwxJNgPpg.CKsOpayd = c.GJd93ieNwt
	out.SwxJNgPpg.F2GXyZZ = c.BcHm_C3XOt
	out.SwxJNgPpg.ChJnngzb = c.X78zRUstxhd
	out.SwxJNgPpg.KWffzRDF = c.XZmYiPZisD8
	out.JhIYwEukjd.GbD_2PU = c.GJd93ieNwt
	out.JhIYwEukjd.T4cuhG7MTbG = c.BcHm_C3XOt
	out.JhIYwEukjd.AJjaqrBjj = c.Gpt4dj
	out.JhIYwEukjd.PBDb3Wc = c.X78zRUstxhd
	out.JhIYwEukjd.QVlaSCZ4RCK = c.XZmYiPZisD8
	out.JhIYwEukjd.HIlblJq1 = c.Tm4m67gYc
	out.JhIYwEukjd.SmInTc = c.RtH5luH9GF0t
	out.JhIYwEukjd.TJ8rFSEH = c.Mnuh5QkiZaeP
	out.JhIYwEukjd.ZxabzIRCNeV = c.DTjgP1
	out.FnMUny.Am3QoeF = c.GJd93ieNwt
	out.FnMUny.HMkc4BgR = c.YDJXeuot82
	out.FnMUny.VuhWqou = c.LXsSwyq7YZ
	out.FnMUny.PzwWMqEpPguJ = c.X78zRUstxhd
	out.FnMUny.DURh5VYO = c.CWNyQiAMJQ82
	out.ASJxBbFS.PU1Di18 = c.GJd93ieNwt
	out.ASJxBbFS.Qggyovo = c.MV3efn
	out.ASJxBbFS.D2rWMm0 = c.KRr2OnixG8
	out.ASJxBbFS.ND3AWUWg = ssrnames.PMorzD7u(c.NrtZd7)
	out.ASJxBbFS.JvlcaS = c.LXsSwyq7YZ
	out.Ck6DxjB.GJtOdKOWBiFP = c.GJd93ieNwt
	out.Ck6DxjB.XmIaPj3CzOF = c.BcHm_C3XOt
	out.Ck6DxjB.WAdJ4ftdZFk = c.Gpt4dj
	out.Ck6DxjB.Nd3mBux1 = c.X78zRUstxhd
	out.Ck6DxjB.J1efyloFyhk = c.XZmYiPZisD8
	out.Ck6DxjB.NMzMUQ2 = c.G9GmqrKxu9iJ
	out.Ck6DxjB.BANVIZ72wh = c.Tm4m67gYc
	out.Ck6DxjB.CrS6C1 = c.RtH5luH9GF0t
	out.Ck6DxjB.DxKJ7lSe3m8F = c.Mnuh5QkiZaeP
	out.Ck6DxjB.IqSfNzea = c.DTjgP1
	out.Ck6DxjB.DlDK5Y9AS = c.B1fm0zbCjL
	out.AYf5kQSQW.PIvpxN6YK = c.GJd93ieNwt
	out.AYf5kQSQW.B30Jt2kO = addrtype.B3YC1rotZ(c.NrtZd7)
	out.AYf5kQSQW.UjgtNuouuKVl = c.DDVa4fX
	out.AYf5kQSQW.HnHarW = c.RU9mVZaMebUI
	out.AYf5kQSQW.W4ggEz = c.Sj7k1l11MYq
	out.QaZjDzw95.CgE0xkW = c.GJd93ieNwt
	out.QaZjDzw95.J6bt8AOW8fL8 = c.Gpt4dj
	out.QaZjDzw95.Jq7YxRdb18Xx = c.Tm4m67gYc
	out.QaZjDzw95.CTa7WN = c.RtH5luH9GF0t
	out.QaZjDzw95.NWbT7wk = c.Mnuh5QkiZaeP
	out.QaZjDzw95.F9P9aeOu23U = c.DTjgP1
	out.ABA4hq.MQE3K4mq = c.GJd93ieNwt
	out.ABA4hq.CLWr9bQsb = c.Qz3ffZ1n4oYU
	out.SwxJNgPpg.UVeahqDBT = c.Gpt4dj == "tls"
	out.FnMUny.ErKysFf2I = ssmethod.GBXRYZ(c.NrtZd7)
	out.ASJxBbFS.QawbKS2x = true
	out.AYf5kQSQW.Li6nvK1 = 2
	out.QaZjDzw95.UPnToXrM = strings.Join(c.JcFiIM, "\n")
	out.ABA4hq.GGSLzD = strings.ToUpper(c.LtcfZGsn7n)
	mode := types.LaTpsgYQ7(2)
	if c.CG9XBFFHZw14 == "obfs" {
		mode = 1
	}
	out.ASJxBbFS.AddSSRUser(&types.Aga91y{JAfXWwzYrW: &types.QcqpSL3i6ri{YrjVl_2g_S: c.GJd93ieNwt, YyAgnMatmj: c.LXsSwyq7YZ, Z_taAb: c.MV3efn, H_y2xZht7S8: c.KRr2OnixG8, ZZIDsJ: ssrnames.PMorzD7u(c.NrtZd7), BskzckkE7qI: mode}})
	return out, nil
}
