package xiaov2board

import (
	"encoding/json"
	"fmt"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/cert"
	C "recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/crypto/hashutil"
	"sync"
)

// 原 init@0x179b920 的 map 0x5096920；混淆字符串限定解码已核准。
var protocolAt5096920 = map[string]C.Mpn2JltM{"vmess": "v2ray", "vless": "vless", "trojan": "trojan", "anytls": "anytls", "shadowsocks": "ss", "hysteria2": "hysteria"}
var certificateErrorsAt509C8C0 sync.Map

// ocVy8aVcKm 对应 0x179c420，先验证协议，再按原顺序修改响应。
func (n *BdT6k1HU8L) ocVy8aVcKm(serverType C.Mpn2JltM) error {
	protocol, ok := protocolAt5096920[n.DWMzz2]
	if !ok {
		return fmt.Errorf("unsupported v2node protocol: %v", n.DWMzz2)
	}
	if protocol != serverType {
		return fmt.Errorf("v2node protocol %v does not match server_type %v", n.DWMzz2, serverType)
	}
	n.edYRMY8 = true
	if n.Xg1j1ZA == 1 && n.PNpnXla.MtI3pIGhwl == "none" {
		n.Xg1j1ZA = 0
	}
	n.PabrtEwnkJ = n.LWnSaWm
	n.F7emlj = n.HnEAiVV
	if n.FoQe96xgtPg == "splithttp" {
		n.FoQe96xgtPg = "xhttp"
	}
	switch protocol {
	case "hysteria":
		n.LQQHdt4QjrX = 2
	case "ss":
		n.YHnnp66mM = ""
		if len(n.PabrtEwnkJ) != 0 {
			if err := json.Unmarshal(n.PabrtEwnkJ, &n.Enhxlh); err != nil {
				return fmt.Errorf("parse ss network settings failed: %v, result: %v", err, string(n.PabrtEwnkJ))
			}
		}
		if n.Enhxlh.XQkYg9e9wLB7 != "" || n.Enhxlh.UzanNits != "" {
			n.YHnnp66mM = "http"
		}
	}
	return nil
}

// kX3RSoCossI 对应 0x179d080；五个原字符串字段按原 trim/跳空 helper 拼接。
func (n *BdT6k1HU8L) kX3RSoCossI() (string, error) {
	if n.HDMF3C == "" {
		return "", nil
	}
	if n.HDMF3C != "mlkem768x25519plus" {
		return "", fmt.Errorf("unsupported vless decryption: %v", n.HDMF3C)
	}
	return hashutil.DP0s5dasePo([]string{n.HDMF3C, n.Qrd69gaE.WaRdy9XI06u, n.Qrd69gaE.OQ0yzunzWm, n.Qrd69gaE.FnvlwhZTa, n.Qrd69gaE.AWaMGzOhxPy}, "."), nil
}

// bybd31ST 对应 0x179cea0；Load/Store 独立，保留原并发时重复日志的可能。
func bybd31ST(id int, message string) {
	if old, ok := certificateErrorsAt509C8C0.Load(id); !ok || old != message {
		certificateErrorsAt509C8C0.Store(id, message)
		logx.V8GN5Stk(4, "node %v: %v", id, message)
	}
}

// wXQOQActA 对应 0x179ca40，非法面板证书只记日志并返回空证书。
func (n *BdT6k1HU8L) wXQOQActA(id int) (string, string) {
	if !n.edYRMY8 || n.Xg1j1ZA != 1 {
		certificateErrorsAt509C8C0.LoadAndDelete(id)
		return "", ""
	}
	switch n.PNpnXla.MtI3pIGhwl {
	case "self":
		return n.hMCNntSK(id)
	case "remote":
		cert, key := n.PNpnXla.Z6rZjvbHxvss, n.PNpnXla.UEzMLAOgTf
		if cert == "" || key == "" {
			bybd31ST(id, "cert_mode is remote but tls_cert or tls_key is empty, ignore panel cert")
			return "", ""
		}
		if err := rtx.ZvmhWx(cert, key); err != nil {
			bybd31ST(id, fmt.Sprintf("panel cert invalid, ignore panel cert: %v", err))
			return "", ""
		}
		certificateErrorsAt509C8C0.LoadAndDelete(id)
		return cert, key
	default:
		certificateErrorsAt509C8C0.LoadAndDelete(id)
		return "", ""
	}
}

// hMCNntSK 对应 0x179cce0，只考虑 server_names 首项，空首项回退单值。
func (n *BdT6k1HU8L) hMCNntSK(id int) (string, string) {
	domain := n.PNpnXla.VHVkiEr1NUE
	if len(n.PNpnXla.XK10gJ_O) > 0 && n.PNpnXla.XK10gJ_O[0] != "" {
		domain = n.PNpnXla.XK10gJ_O[0]
	}
	if domain == "" {
		domain = "example.com"
	}
	cert, key, err := cert.EEGj3Ra(domain)
	if err != nil {
		bybd31ST(id, fmt.Sprintf("create self-signed cert failed, ignore panel cert: %v", err))
		return "", ""
	}
	certificateErrorsAt509C8C0.LoadAndDelete(id)
	return cert, key
}
