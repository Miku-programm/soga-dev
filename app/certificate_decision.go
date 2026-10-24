package app

import (
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/panel/registry"
)

// encryptedTunnel is an analyst-extracted predicate from both original
// functions. Bare "tls" fails the outer supported-tunnel check in the binary.
func encryptedTunnel() bool {
	enabled, mode := config.TunnelSettings()
	if !enabled || config.Q0SDIN() {
		return false
	}
	return mode == "tls-tunnel" || mode == "quic-tunnel" || mode == "ws-tls-tunnel"
}

// gcTStJ6_zEzn restores 0x1f99880. The decompiler's apparent input is unused.
func gcTStJ6_zEzn() bool {
	if encryptedTunnel() {
		return true
	}
	switch config.DnhHoT2aRvs() {
	case "ss", "ssr":
		return false
	case "hysteria":
		return true
	default:
		return !config.ForceCloseSSL()
	}
}

// d_7qy2u restores 0x1f99ae0. Node fetch errors for Trojan/AnyTLS cause a
// warning and default certificate application; other protocols return errors.
func d_7qy2u(panels []registry.VyFk5uT) (bool, error) {
	cert, key := config.ManualCertificateFiles()
	if cert != "" && key != "" {
		return false, nil
	}
	if !gcTStJ6_zEzn() {
		return false, nil
	}
	if encryptedTunnel() {
		return true, nil
	}
	if config.DnhHoT2aRvs() == "hysteria" {
		return true, nil
	}
	switch config.DnhHoT2aRvs() {
	case "v2ray", "vless", "trojan", "anytls":
	default:
		return false, nil
	}
	for _, panel := range panels {
		var node *types.Wr5SVP
		var err error
		if provider, ok := panel.(registry.Chd8czf6); ok {
			node, err = provider.GetNode()
		} else {
			var data *registry.D5gTIZ5aVg
			data, err = panel.GetAPIData()
			if err == nil {
				node = data.H47AzZsYHB1N
			}
		}
		if err != nil {
			switch config.DnhHoT2aRvs() {
			case "trojan", "anytls":
				logx.TRTbrquD(4, "get node failed, apply cert by default:", err)
				return true, nil
			default:
				return false, err
			}
		}
		switch config.DnhHoT2aRvs() {
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
