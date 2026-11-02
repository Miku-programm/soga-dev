package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOriginalTunnelConfigurationGateAndErrors(t *testing.T) {
	old, cert, key := configAt509E300, configAt509F9C0, configAt5096DC0.RM5SFRF
	flag, mirror, marker := keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800
	defer func() {
		configAt509E300, configAt509F9C0, configAt5096DC0.RM5SFRF = old, cert, key
		keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800 = flag, mirror, marker
	}()
	configAt509E300 = k_lu9bPK{GbyJkc: "invalid"}
	if v, e := R3DjPWSnt(); v != nil || e != nil {
		t.Fatal("disabled", v, e)
	}
	configAt509E300.JHXrVGNjg31 = true
	configAt5096DC0.RM5SFRF = ""
	T15Jifo(false)
	if v, e := R3DjPWSnt(); v != nil || e != nil {
		t.Fatal("original key gate", v, e)
	}
	configAt5096DC0.RM5SFRF = "test-key"
	if _, e := R3DjPWSnt(); e == nil || e.Error() != "not support tunnel type: invalid" {
		t.Fatal(e)
	}
	configAt509E300 = k_lu9bPK{JHXrVGNjg31: true, GbyJkc: "tcp-tunnel", ZjL7PiTY: "aes-128-gcm", D38NXN3uV: "p", RtHB6a: true}
	v, e := R3DjPWSnt()
	if e != nil || v == nil || !v.TunnelRecvProxy() || !v.IsMux() || v.IsProxyMode() {
		t.Fatal(v, e)
	}
	T15Jifo(true)
	if v, e := R3DjPWSnt(); e != nil || v == nil || !v.TunnelRecvProxy() || !v.IsMux() || v.IsProxyMode() {
		t.Fatal("tunnel must stay enabled regardless of state", v, e)
	}
	T15Jifo(false)
	configAt509E300.GbyJkc = "tls-tunnel"
	configAt509F9C0.C2BNS6_QJC = filepath.Join(t.TempDir(), "missing-cert")
	configAt509F9C0.JISdid54S5t = "missing-key"
	if _, e := R3DjPWSnt(); e == nil || !os.IsNotExist(e) {
		t.Fatal(e)
	}
	if e := os.WriteFile(configAt509F9C0.C2BNS6_QJC, []byte("bad-cert"), 0600); e != nil {
		t.Fatal(e)
	}
	if _, e := R3DjPWSnt(); e == nil || !os.IsNotExist(e) {
		t.Fatal("key read precedes certificate parse", e)
	}
}
