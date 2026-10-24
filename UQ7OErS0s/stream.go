package UQ7OErS0s

import (
	"fmt"
	"net"
	G "recovered.local/soga/G6bHD5JA"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	R "recovered.local/soga/RoaZP1oJy"
	T "recovered.local/soga/TaFdeUcnzQ"
	S "recovered.local/soga/ad5KEbw_EzWk"
	U "recovered.local/soga/nzATvajt"
	O "recovered.local/soga/tw357_KzLSN2"
)

type DrpXao = net.Conn

// Original 80-byte RTTI c4opfNpyW.Z7PK9WJdYeE.
type Z7PK9WJdYeE struct {
	DrpXao
	rdpZQIeqaxDX *T.KrOWYtzUYft
	xkvBaI       R.Vzuv_s9ZS5
	ai6QQ5       O.Jakq78yIQ
	bIfs9Fg      *S.L7E7H2aW8KVs
	nbQRr7HlusRl *U.LJJkaD
	aWi7sedI43   *G.AHU2MRmjzQu
}

// RfxaTYLqLUJ restores 0x1ce3d40: accounting -> obfuscation -> cipher ->
// authentication, then destination parsing, user lookup, and deferred traffic
// accounting. Construction errors before init do not add cleanup absent in ELF.
func RfxaTYLqLUJ(raw net.Conn, info *S.L7E7H2aW8KVs) (*Z7PK9WJdYeE, error) {
	accounting := &J.Mdn09F6aMrpa{DrpXao: raw}
	obfs, err := O.Z_Mqu9of(info.UNRudV1377x9, accounting)
	if err != nil {
		return nil, err
	}
	encryption, err := T.VavyqXsqy9V(info, obfs)
	if err != nil {
		return nil, err
	}
	protocol, err := R.HrQWlmA(info.C8Qsa8G8XMw0, encryption)
	if err != nil {
		return nil, err
	}
	obfs.SetServerInfo(info)
	protocol.SetServerInfo(info)
	info.CJfUaAqL = obfs.GetOverhead() + protocol.GetOverhead()
	c := &Z7PK9WJdYeE{DrpXao: raw, rdpZQIeqaxDX: encryption, xkvBaI: protocol, ai6QQ5: obfs, bIfs9Fg: info}
	if err = c.init(); err != nil {
		H.Li3c4CJ1_D(c)
		return nil, err
	}
	accounting.SetUser(c.nbQRr7HlusRl)
	return c, nil
}

// init restores 0x1ce4080; obfs mode 1 looks up the original shared EFLMTH
// field, which obfs parsing populates, rather than inventing a separate lookup.
func (c *Z7PK9WJdYeE) init() error {
	dest, err := G.SzhW4tOzUZg(c.xkvBaI)
	if err != nil {
		return err
	}
	var user *U.LJJkaD
	var ok bool
	switch c.bIfs9Fg.NxYA16xMfF {
	case 0, 2:
		uid := c.bIfs9Fg.Uklg6401
		if uid <= 0 {
			uid = c.xkvBaI.CurUID()
		}
		user, ok = c.bIfs9Fg.N3pMXuwX85c.Get(uid)
		if !ok {
			return fmt.Errorf("uid <%v> not exist", uid)
		}
	case 1:
		user, ok = c.bIfs9Fg.RLXBf_luGj.Get(c.bIfs9Fg.EFLMTH)
		if !ok {
			return fmt.Errorf("host <%v> not exist", c.bIfs9Fg.EFLMTH)
		}
	default:
		return fmt.Errorf("not support single port type: %v", c.bIfs9Fg.NxYA16xMfF)
	}
	c.nbQRr7HlusRl = user
	c.aWi7sedI43 = dest
	return nil
}
func (c *Z7PK9WJdYeE) Read(p []byte) (int, error)  { return c.xkvBaI.Read(p) }  //0x1ce44c0
func (c *Z7PK9WJdYeE) Write(p []byte) (int, error) { return c.xkvBaI.Write(p) } //0x1ce4520
func (c *Z7PK9WJdYeE) Close() error                { return c.DrpXao.Close() }  //0x1ce4580
func (c *Z7PK9WJdYeE) CurUID() int64 {
	if c.bIfs9Fg.Uklg6401 > 0 {
		return c.bIfs9Fg.Uklg6401
	}
	return c.xkvBaI.CurUID()
}                                                     //0x1ce45c0
func (c *Z7PK9WJdYeE) CID() int64                     { return c.xkvBaI.CID() }   //0x1ce4620
func (c *Z7PK9WJdYeE) Host() string                   { return c.bIfs9Fg.EFLMTH } //0x1ce4660
func (c *Z7PK9WJdYeE) GetServerInfo() *S.L7E7H2aW8KVs { return c.bIfs9Fg }        //0x1ce4680
func (c *Z7PK9WJdYeE) GetUser() U.PgaNd3014U          { return c.nbQRr7HlusRl }   //0x1ce46a0, includes typed nil
func (c *Z7PK9WJdYeE) GetDestAddr() *G.AHU2MRmjzQu    { return c.aWi7sedI43 }     //0x1ce46c0
func (c *Z7PK9WJdYeE) NetworkType() J.ZES6F5          { return 0 }                //0x1ce46e0
func (c *Z7PK9WJdYeE) Free() { //0x1ce4700
	H.U9tk4Xosh9(c.rdpZQIeqaxDX)
	if v, ok := c.xkvBaI.(H.JZgM1sMF); ok {
		H.U9tk4Xosh9(v)
	}
	if v, ok := c.ai6QQ5.(H.JZgM1sMF); ok {
		H.U9tk4Xosh9(v)
	}
	H.Li3c4CJ1_D(c.DrpXao)
}
