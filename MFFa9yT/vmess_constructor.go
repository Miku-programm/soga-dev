package MFFa9yT

import (
	G "recovered.local/soga/G7A7csgrs62p"
	J "recovered.local/soga/J1P4KNvfI_J"
	N "recovered.local/soga/NtxdAE"
	C "recovered.local/soga/cNT24rG4F"
	V "recovered.local/soga/gk1z9v"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
)

// ag90n_rZjJO @0x1da0d40，原生VMess构造，原模式分支分别创建认证依赖。
func ag90n_rZjJO(arg *DV_XPJ4z0) (Kz_8ajU7, error) {
	base, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &K4avFiYOPi{JWN154qJ4: base}
	if C.VmessAuthMode() != "aead" {
		v.dj_cTnaUp6f = V.NewMD5AuthManager(base.zb8f3cHWW8v, base.jc_QPN6, C.VmessReduceMemory())
	}
	if C.VmessAuthMode() != "md5" {
		v.p54XvyV = &U.ZpNFqUVZp{GTFAWF: &G.GTFAWF[int64, U.PgaNd3014U]{}}
		v.l_ZoZQDxG7 = J.YCcacra(base.zb8f3cHWW8v)
		v.qgz9PfTS = N.Azb3Zr(base.zb8f3cHWW8v, C.VmessInvalidAccessSettings())
	}
	v.UpdateAPIData(arg.HG5k0365)
	if v.dj_cTnaUp6f != nil {
		v.dj_cTnaUp6f.StartTask()
	}
	Log.TRTbrquD(0, "create vmess server")
	Log.V8GN5Stk(0, "VMess auth mode: %v", C.VmessAuthMode())
	return v, nil
}
