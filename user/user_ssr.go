package user

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/ssrnames"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var ssrHostPatternAt5096318 = regexp.MustCompile(`%(\d+)m`) // init @ 0x181cea0.

// LJJkaD preserves the original 136-byte SSR user layout.
type LJJkaD struct {
	*AhSFeSua
	LINadiEOz2  int
	aVb5zSn     string
	pGs6LTCd    ssrnames.ZvAxOt1l7Ar
	qWkRlPbzPlq ssrnames.Mu0vMIIs
	dVoGxBq     ssrnames.PMorzD7u
	lK3OaWUYP   types.LaTpsgYQ7
	kHH6tGdN    map[int64]int
	jS2vrvcY    sync.Mutex
	wyY4xeau    int
	zBAP060J    int64
	ya8WV3ZP11A string
}

func PXHoe5(p *types.Aga91y, node *relay.L337vk2G) (*LJJkaD, error) { // 0x1821de0
	if _, ok := p.JAfXWwzYrW.(*types.QcqpSL3i6ri); !ok {
		return nil, pNF7CSleWS(p.WZ89pTAOrHK4, "ssr")
	}
	u := &LJJkaD{AhSFeSua: cLLKJ_A31(p, node), kHH6tGdN: make(map[int64]int)}
	u.UpdatePanelUser(p)
	return u, nil
}
func (u *LJJkaD) UpdatePanelUser(p *types.Aga91y) (bool, string) { // 0x1821ec0
	u.UpdateUser(p)
	old := u.aVb5zSn
	c, ok := p.JAfXWwzYrW.(*types.QcqpSL3i6ri)
	if !ok {
		return false, old
	}
	changed := u.LINadiEOz2 != c.YrjVl_2g_S || u.aVb5zSn != c.YyAgnMatmj || u.pGs6LTCd != c.Z_taAb || u.qWkRlPbzPlq != c.H_y2xZht7S8 || u.dVoGxBq != c.ZZIDsJ
	u.LINadiEOz2 = c.YrjVl_2g_S
	u.aVb5zSn = c.YyAgnMatmj
	u.pGs6LTCd = c.Z_taAb
	u.qWkRlPbzPlq = c.H_y2xZht7S8
	u.dVoGxBq = c.ZZIDsJ
	// Single-port mode is assigned unconditionally, but does not mark a change.
	u.lK3OaWUYP = c.BskzckkE7qI
	if changed {
		panel := config.SSRPanelType()
		switch panel {
		case "sspanel-uim", "soga-v1":
			u.zBAP060J = u.djdKrU
		case "proxypanel", "whmcs":
			u.zBAP060J = int64(u.LINadiEOz2)
		default:
			panic("can not get ssr uid, not support panel: " + panel)
		}
		u.ya8WV3ZP11A = ""
		u.ResetCtx()
	}
	return changed, old
}
func (u *LJJkaD) AddDevice(device int64) bool { // 0x18222e0
	u.jS2vrvcY.Lock()
	defer u.jS2vrvcY.Unlock()
	count, exists := u.kHH6tGdN[device]
	if u.i_fZvtng2iy > 0 && !exists && u.wyY4xeau >= u.i_fZvtng2iy {
		return false
	}
	if exists {
		u.kHH6tGdN[device] = count + 1
	} else {
		u.kHH6tGdN[device] = 1
		u.wyY4xeau++
	}
	return true
}
func (u *LJJkaD) DelDevice(device int64) { // 0x1822460
	u.jS2vrvcY.Lock()
	defer u.jS2vrvcY.Unlock()
	count, exists := u.kHH6tGdN[device]
	if !exists {
		return
	}
	if count > 1 {
		u.kHH6tGdN[device] = count - 1
	} else {
		delete(u.kHH6tGdN, device)
		u.wyY4xeau--
	}
}
func (u *LJJkaD) GenUserMd5() string { // 0x18225a0
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(u.djdKrU, 10)))
	h.Write([]byte(u.aVb5zSn))
	h.Write([]byte(u.pGs6LTCd))
	h.Write([]byte(u.dVoGxBq))
	h.Write([]byte(u.qWkRlPbzPlq))
	return hex.EncodeToString(h.Sum(nil))
}
func (u *LJJkaD) GetSSRHost() (string, error) { // 0x1822860
	if u.ya8WV3ZP11A != "" {
		return u.ya8WV3ZP11A, nil
	}
	suffix, template := config.SSRHostConfiguration()
	if suffix == "" || template == "" {
		return "", errors.New("must config mu_suffix and mu_regex")
	}
	matches := ssrHostPatternAt5096318.FindStringSubmatch(template)
	n, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", err
	}
	digest := u.GenUserMd5()
	if n > len(digest) {
		n = len(digest)
	}
	result := strings.Replace(template, matches[0], digest[:n], 1)
	result = strings.Replace(result, "%id", strconv.FormatInt(u.djdKrU, 10), 1)
	result = strings.Replace(result, "%suffix", suffix, 1)
	u.ya8WV3ZP11A = result
	return result, nil
}
func (u *LJJkaD) GetUID() int64                           { return u.zBAP060J }                                               // 0x1822ae0
func (u *LJJkaD) String() string                          { return fmt.Sprintf("id: %v, password: %v", u.djdKrU, u.aVb5zSn) } // 0x1822b00
func (u *LJJkaD) GetPassword() string                     { return u.aVb5zSn }                                                // 0x1822c00
func (u *LJJkaD) GetMethod() ssrnames.ZvAxOt1l7Ar      { return u.pGs6LTCd }                                               // 0x1822c20
func (u *LJJkaD) GetProtocol() ssrnames.Mu0vMIIs       { return u.qWkRlPbzPlq }                                            // 0x1822c40
func (u *LJJkaD) GetObfs() ssrnames.PMorzD7u           { return u.dVoGxBq }                                                // 0x1822c60
func (u *LJJkaD) GetSinglePortType() types.LaTpsgYQ7 { return u.lK3OaWUYP }                                              // 0x1822c80
