package inbound

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/gofrs/uuid"
	"net"
	P "recovered.local/soga/transport/listener"
	G "recovered.local/soga/util/genmap"
	E "recovered.local/soga/protocol/vless"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
	"recovered.local/soga/reality"
	"strings"
)

// oHr1LeATmS@1d96640，原生VLESS入站构造。
func oHr1LeATmS(arg *DV_XPJ4z0) (Kz_8ajU7, error) {
	base, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &KHonGaqbIy_6{JWN154qJ4: base, vU51MPl_at_: &G.GTFAWF[uuid.UUID, *U.A6CvRbtca]{}}
	v.UpdateAPIData(arg.HG5k0365)
	Log.TRTbrquD(0, "create vless server")
	return v, nil
}

// Start@1d967a0与defer1e726c0；失败只关闭此次已创建监听器。
func (v *KHonGaqbIy_6) Start() (err error) {
	node := v.wiy6P3.Load()
	cfg := &node.Ck6DxjB
	if cfg.DlDK5Y9AS != "" && cfg.DlDK5Y9AS != "none" {
		info, ok := cfg.GetDecryptionInfo()
		if !ok {
			return fmt.Errorf("invalid vless decryption: %v", cfg.DlDK5Y9AS)
		}
		var keys [][]byte
		for _, s := range strings.Split(info.I_Wzoo_nf, ".") {
			k, _ := base64.RawURLEncoding.DecodeString(s)
			keys = append(keys, k)
		}
		crypt := &E.LgF3OB1QX{}
		if err = crypt.Init(keys, info.KAeFR260tbP, info.E3a03_Hf, info.R3tya30, info.LRGrsuCn); err != nil {
			return fmt.Errorf("failed to use vless decryption: %v", err)
		}
		v.i420mh4Yr = crypt
	}
	var tlsConfig *tls.Config
	var realityConfig *reality.Config
	switch cfg.WAdJ4ftdZFk {
	case "tls":
		if !C.ForceCloseSSL() {
			tlsConfig, err = v.kaVZkL()
			if err != nil {
				return err
			}
			if err = v.dINraX_uOx5T(tlsConfig); err != nil {
				return err
			}
		}
	case "reality":
		realityConfig, err = fwKvAvIIe(v.zb8f3cHWW8v, cfg.BANVIZ72wh, cfg.CrS6C1, cfg.IqSfNzea, cfg.DxKJ7lSe3m8F)
		if err != nil {
			return err
		}
	}
	if cfg.NMzMUQ2 != "" && cfg.NMzMUQ2 != "xtls-rprx-vision" {
		return fmt.Errorf("not support vless flow: %v", cfg.NMzMUQ2)
	}
	tunnel, err := C.R3DjPWSnt()
	if err != nil {
		return err
	}
	if tunnel != nil {
		C.QitK3yi()
	}
	var listeners []net.Listener
	defer func() {
		if err != nil {
			for _, l := range listeners {
				l.Close()
			}
		}
	}()
	for _, host := range C.EdlWFAt() {
		proxy, force := C.ListenerProxyProtocolSettings()
		options := &P.Ap9uD2{Host: host, Port: cfg.GJtOdKOWBiFP, Network: cfg.XmIaPj3CzOF, TLS: tlsConfig, Reality: realityConfig, Path: cfg.Nd3mBux1, ServiceName: cfg.J1efyloFyhk, ProxyProtocol: proxy, ForceProxyProtocol: force, Tunnel: tunnel, XHTTP: cfg.FPhrzP645}
		var l net.Listener
		l, err = options.NewListener()
		if err != nil {
			return err
		}
		if v.i420mh4Yr != nil {
			Log.V8GN5Stk(0, "VLESS listen: %v, decryption: %v", options.String(), cfg.DlDK5Y9AS)
		} else {
			Log.V8GN5Stk(0, "VLESS listen: %v, decryption: none", options.String())
		}
		listeners = append(listeners, l)
	}
	v.kX5i8U = listeners
	for _, l := range listeners {
		v.xAm4LgsEB8j(v.zb8f3cHWW8v, l)
	}
	return nil
}
