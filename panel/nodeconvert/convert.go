package nodeconvert

import (
	"errors"
	"fmt"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	W "recovered.local/soga/crypto/ssmethod"
	X "recovered.local/soga/audit/matcher"
)

// mCZQ7N3 restores0x1765f20, including default-network mutation before hashing.
func (in *PNf9Em87) mCZQ7N3() error {
	switch in.QxNm4UvLN {
	case "vmess", "vless", "trojan":
		switch in.Jedod2gPZ4UT.Id1HxAg5Oy {
		case "", "tls", "none":
		default:
			return fmt.Errorf("unknown security: %s", in.Jedod2gPZ4UT.Id1HxAg5Oy)
		}
	case "shadowsocks":
	default:
		return fmt.Errorf("unknown protocol: %s", in.QxNm4UvLN)
	}
	switch in.Jedod2gPZ4UT.WZVYFl {
	case "":
		in.Jedod2gPZ4UT.WZVYFl = "tcp"
	case "tcp", "ws", "h2", "grpc":
	default:
		return fmt.Errorf("unknown network: %s", in.Jedod2gPZ4UT.WZVYFl)
	}
	return nil
}
func (rules HidyXaJnC) cGVHFYeam0GS() bool { //0x17663a0
	for _, r := range rules {
		if r.Ab0SdSa == "block" {
			for _, p := range r.DaFBjr {
				if p == "bittorrent" {
					return true
				}
			}
		}
	}
	return false
}
func (r *AnlFZbNq) ToPanelNode() (*L.Wr5SVP, error) { //0x17664e0
	if len(r.Dps5ZWL) == 0 {
		return nil, errors.New("inbounds is empty")
	}
	in := r.Dps5ZWL[0]
	if err := in.mCZQ7N3(); err != nil {
		return nil, err
	}
	n := &L.Wr5SVP{Z1TIgHOQ: Q.T0j_Rco(in), HSA2b5V: C.PanelForbiddenBitTorrent() || r.YhC3piOh.UwS4pcm.cGVHFYeam0GS(), TtWKyh: C.PanelNodeSpeedLimit(), W3Dw_B4wgbM: C.PanelUserConnectionLimit()}
	n.SwxJNgPpg.CKsOpayd = in.SioIDq
	n.SwxJNgPpg.F2GXyZZ = L.BoulTn31(in.Jedod2gPZ4UT.WZVYFl)
	n.SwxJNgPpg.UVeahqDBT = in.Jedod2gPZ4UT.Id1HxAg5Oy == "tls"
	n.SwxJNgPpg.ChJnngzb = in.Jedod2gPZ4UT.D5bsRALb.JCRhkrluQQY
	n.SwxJNgPpg.KWffzRDF = in.Jedod2gPZ4UT.Ol5ZxI8vO.RVKq_aLvfWyF
	n.FnMUny.Am3QoeF = in.SioIDq
	n.FnMUny.HMkc4BgR = W.C6q3Upu(in.VhYAkS.P9ydbTayFu)
	n.FnMUny.VuhWqou = in.VhYAkS.XUfl9V
	n.JhIYwEukjd.GbD_2PU = in.SioIDq
	n.JhIYwEukjd.T4cuhG7MTbG = "tcp"
	return n, nil
}
func (r *AnlFZbNq) GetDetectors() X.Ij9r21E0UOt3 { //0x17667e0
	out := X.Ij9r21E0UOt3{}
	for _, v := range r.YhC3piOh.UwS4pcm {
		if v.Ab0SdSa == "block" {
			for _, d := range v.F3hpMd0 {
				out = append(out, X.NewPanelDetector(0, d))
			}
		}
	}
	return out
}
func (r *Oh1oNuTp) ToPanelUsers(kind C.Mpn2JltM) (L.Sxcf3wu3L, error) { //0x1768d60
	if r.DCbciSyJ_c != "ok" {
		return nil, errors.New(r.DCbciSyJ_c)
	}
	out := make(L.Sxcf3wu3L, 0, len(r.XH32wVY))
	for _, u := range r.XH32wVY {
		var credential L.ADVqOHSx
		switch kind {
		case "ss":
			credential = &L.YlfOV0DBBnx{A_wNUTkMh6pH: u.HZO6BjGa}
		case "v2ray":
			credential = &L.Yn8BJoL{SfiEGpJo: u.SDB5DM3E68Y}
		case "trojan":
			credential = &L.JxTGUar{ARARP1z0yA: u.NRLcM3Xm2zgj}
		}
		speed, global := u.WEkiWDf5si, C.PanelUserSpeedLimit()
		if speed < 0 {
			speed = 0
		} else if speed > 10000 {
			speed = 10000
		}
		if global < 0 {
			global = 0
		} else if global > 10000 {
			global = 10000
		}
		// UCOMISD/JA: unordered comparisons take the assignment path as in original.
		if speed > 0 && global > 0 {
			if !(global > speed) {
				speed = global
			}
		} else if !(speed > global) {
			speed = global
		}
		connections, g := u.C2TUmBU9gT, C.PanelUserConnectionLimit()
		if connections < 0 {
			connections = 0
		}
		if g < 0 {
			g = 0
		}
		if connections > 0 && g > 0 {
			if g > connections {
				g = connections
			}
		} else if g < connections {
			g = connections
		}
		out = append(out, &L.Aga91y{WZ89pTAOrHK4: u.PqTXyjlHNDD_, LmOKwjha: speed, Xmrna8: g, B4DYSxuA31za: C.PanelUserTCPLimit(), JAfXWwzYrW: credential})
	}
	return out, nil
}
