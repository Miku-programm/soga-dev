package user

import (
	"strings"
	"testing"
	"unsafe"

	"github.com/gofrs/uuid"
	"recovered.local/soga/panel/types"
)

func TestVLESSOriginalUserLifecycle(t *testing.T) {
	first := uuid.Must(uuid.FromString("00112233-4455-6677-8899-aabbccddeeff"))
	second := uuid.Must(uuid.FromString("fedcba98-7654-3210-abcd-0123456789ab"))
	p := &types.Aga91y{WZ89pTAOrHK4: 23, LmOKwjha: 5, JAfXWwzYrW: &types.IImIwacp{Dh5iiakU: first.String()}}
	u, err := Sa_NiNyuJ_ax(p, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { u.kI9WvUIev() }()
	if unsafe.Sizeof(*u) != 16 || unsafe.Offsetof(u.eszXwjoQ) != 8 {
		t.Fatal("layout")
	}
	if id, ok := u.GetUuid(); !ok || id != first {
		t.Fatal(id, ok)
	}
	ctx := u.GetCtx()
	if old, changed, err := u.UpdatePanelUser(p); old != first || changed || err != nil || ctx != u.GetCtx() {
		t.Fatal(old, changed, err)
	}
	p.JAfXWwzYrW = &types.IImIwacp{Dh5iiakU: second.String()}
	if old, changed, err := u.UpdatePanelUser(p); old != first || !changed || err != nil || ctx.Err() == nil {
		t.Fatal(old, changed, err)
	}
	p.LmOKwjha = 2
	p.JAfXWwzYrW = nil
	ctx = u.GetCtx()
	if old, changed, err := u.UpdatePanelUser(p); old != uuid.Nil || changed || err == nil || err.Error() != "user 23 has no vless credential" {
		t.Fatal(old, changed, err)
	}
	if u.GetSpeedLimit() != 2 || u.GetCtx() != ctx || ctx.Err() != nil {
		t.Fatal("failed credential still updates base, retains context")
	}
	p.JAfXWwzYrW = &types.IImIwacp{Dh5iiakU: "bad"}
	if _, changed, err := u.UpdatePanelUser(p); changed || err == nil || !strings.HasPrefix(err.Error(), "parse uuid <bad> failed: ") {
		t.Fatal(changed, err)
	}
	if id, ok := u.GetUuid(); !ok || id != second {
		t.Fatal("error must preserve UUID")
	}
	if u.String() != "id: 23, uuid: "+second.String() {
		t.Fatal(u.String())
	}
	if got, err := Sa_NiNyuJ_ax(p, nil); got != nil || err == nil || !strings.HasPrefix(err.Error(), "update user 23 failed: parse uuid <bad> failed:") {
		t.Fatal(got, err)
	}
	zero := &A6CvRbtca{AhSFeSua: cLLKJ_A31(p, nil)}
	defer zero.kI9WvUIev()
	if id, ok := zero.GetUuid(); ok || id != uuid.Nil {
		t.Fatal(id, ok)
	}
	if zero.String() != "id: 23, uuid: nil" {
		t.Fatal(zero.String())
	}
}
