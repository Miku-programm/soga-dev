package CbO_7wp3B

import (
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/q5bqCgTdxaTM"
)

// encryptedTunnel is an analyst-extracted predicate from both original
// functions. Bare "tls" fails the outer supported-tunnel check in the binary.
func encryptedTunnel() bool {
	enabled, mode := cNT24rG4F.TunnelSettings()
	if !enabled || cNT24rG4F.Q0SDIN() {
		return false
	}
	return mode == "tls-tunnel" || mode == "quic-tunnel" || mode == "ws-tls-tunnel"
}

// gcTStJ6_zEzn restores 0x1f99880. The decompiler's apparent input is unused.
func gcTStJ6_zEzn() bool {
	if encryptedTunnel() {
		return true
	}
	switch cNT24rG4F.DnhHoT2aRvs() {
	case "ss", "ssr":
		return false
	case "hysteria":
		return true
	default:
		return !cNT24rG4F.ForceCloseSSL()
	}
}

// d_7qy2u restores 0x1f99ae0. Node fetch errors for Trojan/AnyTLS cause a
// warning and default certificate application; other protocols return errors.
func d_7qy2u(panels []q5bqCgTdxaTM.VyFk5uT) (bool, error) {
	cert, key := cNT24rG4F.ManualCertificateFiles()
	if cert != "" && key != "" {
		return false, nil
	}
	if !gcTStJ6_zEzn() {
		return false, nil
	}
	if encryptedTunnel() {
		return true, nil
	}
	if cNT24rG4F.DnhHoT2aRvs() == "hysteria" {
		return true, nil
	}
	switch cNT24rG4F.DnhHoT2aRvs() {
	case "v2ray", "vless", "trojan", "anytls":
	default:
		return false, nil
	}
	for _, panel := range panels {
		var node *LFWTTXF7JR.Wr5SVP
		var err error
		if provider, ok := panel.(q5bqCgTdxaTM.Chd8czf6); ok {
			node, err = provider.GetNode()
		} else {
			var data *q5bqCgTdxaTM.D5gTIZ5aVg
			data, err = panel.GetAPIData()
			if err == nil {
				node = data.H47AzZsYHB1N
			}
		}
		if err != nil {
			switch cNT24rG4F.DnhHoT2aRvs() {
			case "trojan", "anytls":
				muc279Mpa6S4.TRTbrquD(4, "get node failed, apply cert by default:", err)
				return true, nil
			default:
				return false, err
			}
		}
		switch cNT24rG4F.DnhHoT2aRvs() {
		case "v2ray":
			if node.SwxJNgPpg.UVeahqDBT {
				return true, nil
			}
		case "vless":
			if node.Ck6DxjB.WAdJ4ftdZFk == "tls" {
				return true, nil
			}
		case "anytls":
			if security := node.QaZjDzw95.J6bt8AOW8fL8; security != "reality" && security != "none" {
				return true, nil
			}
		case "trojan":
			if security := node.JhIYwEukjd.AJjaqrBjj; security != "reality" && security != "none" {
				return true, nil
			}
		}
	}
	return false, nil
}
