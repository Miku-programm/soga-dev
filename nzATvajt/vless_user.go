package nzATvajt

import (
	"fmt"
	"sync/atomic"

	"github.com/gofrs/uuid"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/LFWTTXF7JR"
)

// A6CvRbtca：原 RTTI 的 16 字节 VLESS 用户，嵌入用户基类及原子 UUID 指针。
type A6CvRbtca struct {
	*AhSFeSua
	eszXwjoQ atomic.Pointer[uuid.UUID]
}

// Sa_NiNyuJ_ax @ 0x1826e60。
func Sa_NiNyuJ_ax(p *LFWTTXF7JR.Aga91y, node *G6bHD5JA.L337vk2G) (*A6CvRbtca, error) {
	u := &A6CvRbtca{AhSFeSua: cLLKJ_A31(p, node)}
	if _, _, err := u.UpdatePanelUser(p); err != nil {
		return nil, fmt.Errorf("update user %v failed: %v", u.djdKrU, err)
	}
	return u, nil
}

// UpdatePanelUser @ 0x1827020。凭据不合法时公共字段依然已经更新。
func (u *A6CvRbtca) UpdatePanelUser(p *LFWTTXF7JR.Aga91y) (uuid.UUID, bool, error) {
	u.UpdateUser(p)
	credential, ok := p.JAfXWwzYrW.(*LFWTTXF7JR.IImIwacp)
	if !ok {
		return uuid.Nil, false, pNF7CSleWS(p.WZ89pTAOrHK4, "vless")
	}
	value := new(uuid.UUID)
	if err := value.Parse(credential.Dh5iiakU); err != nil {
		return uuid.Nil, false, fmt.Errorf("parse uuid <%v> failed: %v", credential.Dh5iiakU, err)
	}
	previous := u.eszXwjoQ.Load()
	var old uuid.UUID
	if previous != nil {
		old = *previous
	}
	changed := previous == nil || *value != old
	if changed {
		u.eszXwjoQ.Store(value)
		u.ResetCtx()
	}
	return old, changed, nil
}
func (u *A6CvRbtca) GetUuid() (uuid.UUID, bool) { // 0x1827320
	if value := u.eszXwjoQ.Load(); value != nil {
		return *value, true
	}
	return uuid.Nil, false
}
func (u *A6CvRbtca) String() string { // 0x1827360
	if value, ok := u.GetUuid(); ok {
		return fmt.Sprintf("id: %v, uuid: %v", u.djdKrU, value)
	}
	return fmt.Sprintf("id: %v, uuid: nil", u.djdKrU)
}
