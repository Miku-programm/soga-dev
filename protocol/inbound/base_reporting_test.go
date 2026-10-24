package inbound

import (
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/user"
	"reflect"
	"testing"
)

func TestBaseReportingLayoutAndDrain(t *testing.T) {
	typ := reflect.TypeFor[JWN154qJ4]()
	if typ.Size() != 336 {
		t.Fatal(typ.Size())
	}
	for i, off := range []uintptr{0, 16, 24, 32, 40, 48, 56, 64, 88, 96, 104, 112, 120, 128, 152, 176, 192, 200, 208, 216, 232, 256, 264, 288, 312} {
		if typ.Field(i).Offset != off {
			t.Fatal(typ.Field(i))
		}
	}
	b := &JWN154qJ4{jc_QPN6: &user.ZpNFqUVZp{GTFAWF: &genmap.GTFAWF[int64, user.PgaNd3014U]{}}, qE9aU9: &genmap.GTFAWF[int, *genmap.Fxs7SY[int64]]{}}
	u, err := user.Sa_NiNyuJ_ax(&types.Aga91y{WZ89pTAOrHK4: 7, JAfXWwzYrW: &types.IImIwacp{Dh5iiakU: "11111111-1111-4111-8111-111111111111"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b.jc_QPN6.Add(u)
	if !b.UserExist(7) || b.UserExist(8) {
		t.Fatal("identity")
	}
	u.AddUp(11)
	u.AddDown(-2)
	out := b.wKvd7jacLJ()
	if len(out) != 1 || out[0].Hy7Q97RROK != 7 || out[0].XJkGTQf_qkAt != 11 || out[0].I6LXObODav != -2 {
		t.Fatal(out)
	}
	if len(b.wKvd7jacLJ()) != 0 {
		t.Fatal("not drained")
	}
	u.AddUp(-3)
	if len(b.wKvd7jacLJ()) != 0 {
		t.Fatal("negative-only traffic must be omitted")
	}
	if up, down := u.GetAndClearTraffic(); up != 0 || down != 0 {
		t.Fatal(up, down)
	}
	s := &genmap.Fxs7SY[int64]{}
	s.Add(7)
	s.Add(7)
	s.Add(8)
	b.qE9aU9.Set(22, s)
	reports := b.iyyXcKcQH7k7()
	if len(reports) != 2 || b.qE9aU9.Len() != 0 {
		t.Fatal(reports)
	}
	seen := map[int64]bool{}
	for _, r := range reports {
		if r.CAgUUdrv9 != 22 {
			t.Fatal(r)
		}
		seen[r.GbzQUT_4f] = true
	}
	if !seen[7] || !seen[8] {
		t.Fatal(seen)
	}
	if empty := b.iyyXcKcQH7k7(); empty == nil || len(empty) != 0 {
		t.Fatal(empty)
	}
	if empty := b.jdbxrQIN2(); empty == nil || len(empty) != 0 || cap(empty) != 16 {
		t.Fatal(empty)
	}
}
