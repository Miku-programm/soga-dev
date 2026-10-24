package nzATvajt

import (
	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/LFWTTXF7JR"
	"reflect"
	"testing"
)

func TestOriginalUserCollectionReplaceAndPromotedMethods(t *testing.T) {
	typ := reflect.TypeFor[ZpNFqUVZp]()
	if typ.Size() != 8 || typ.NumField() != 1 || !typ.Field(0).Anonymous || typ.Field(0).Name != "GTFAWF" {
		t.Fatal(typ)
	}
	m := &ZpNFqUVZp{new(G7A7csgrs62p.GTFAWF[int64, PgaNd3014U])}
	u := cLLKJ_A31(&LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 8}, nil)
	v := cLLKJ_A31(&LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 8}, nil)
	defer u.kI9WvUIev()
	defer v.kI9WvUIev()
	m.Add(u)
	m.Add(v)
	got, ok := m.Get(8)
	if !ok || got != v {
		t.Fatal(got, ok)
	}
	got, ok = m.GetAndDelete(8)
	if !ok || got != v {
		t.Fatal(got, ok)
	}
	if _, ok := m.Get(8); ok {
		t.Fatal("deleted user remains")
	}
}
