package encmux

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	D "recovered.local/soga/crypto/kdf"
	A "recovered.local/soga/transport/utlsconfig"
)

var derivationInfoAt5099910, derivationInfoAt5099930 []byte

// init 对应0x11da700。两串由原decoder受限仿真恢复；保留解码失败panic。
func init() {
	var err error
	derivationInfoAt5099910, err = hex.DecodeString("00f0fb99e48653e8ade12a596a954c9816116fd75d64be86b799e9b33464c670")
	if err != nil {
		panic(err)
	}
	derivationInfoAt5099930, err = hex.DecodeString("d8b8c066c10e8b0d1b1ba3ad1f7de1525625d03f9d9f682f1bd6097b13a22b38")
	if err != nil {
		panic(err)
	}
}

// Verify 对应0x11db580，原固定16字节普通相等比较。
func (v *VTl94pzz2IkB) Verify(token []byte) bool {
	return len(token) == 16 && bytes.Equal(v.x5Lor8hDyn[:], token)
}

// VerifyLen 对应0x11db600。
func (v *VTl94pzz2IkB) VerifyLen() int { return 16 }

// NonceLen 对应0x11db620。
func (v *VTl94pzz2IkB) NonceLen() int { return v.zkn35IN.NonceLen() }

// KeyLen 对应0x11db640。
func (v *VTl94pzz2IkB) KeyLen() int { return v.zkn35IN.KeyLen() }

// NewAEAD 对应0x11db660，原funcvalue376ec80明确指向x/crypto/sha3.New512。
func (v *VTl94pzz2IkB) NewAEAD(salt []byte) ([]byte, cipher.AEAD, error) {
	key := D.ENKzXYE1a(sha3.New512, v.oWvGY4I9cA, salt, derivationInfoAt5099930, v.zkn35IN.KeyLen())
	aead, err := v.zkn35IN.NewAEAD(key)
	return key, aead, err
}

// TunnelRecvProxy 对应0x11db780。
func (v *VTl94pzz2IkB) TunnelRecvProxy() bool { return v.jrF7WDTK }

// IsProxyMode 对应0x11db7a0。
func (v *VTl94pzz2IkB) IsProxyMode() bool { return v.hrXHbLocdr }

// ProxyAddr 对应0x11db7c0。
func (v *VTl94pzz2IkB) ProxyAddr() *A.FGMp1YT7 { return v.dlAoUM }

// IsMux 对应0x11db7e0。
func (v *VTl94pzz2IkB) IsMux() bool { return v.e2TiiBKaBt }

// IsQuic 对应0x11db800。
func (v *VTl94pzz2IkB) IsQuic() bool { return v.hZORuZstvS07 }

// MuxNum 对应0x11db820。
func (v *VTl94pzz2IkB) MuxNum() int { return v.pyauaz }
