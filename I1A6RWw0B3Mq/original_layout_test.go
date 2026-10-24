package I1A6RWw0B3Mq

import (
	"reflect"
	"testing"
)

func TestRecoveredOriginalRTTI(t *testing.T) {
	{
		typ := reflect.TypeOf(ZK5THfX{})
		if typ.Size() != 120 || typ.NumField() != 9 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "QvNk7x1dJQ" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "json:\"id\" gorm:\"column:id;primaryKey;autoIncrement\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "Pgpq5sNc" || f.Offset != 8 || f.Type.Size() != 16 || string(f.Tag) != "json:\"domain\" gorm:\"column:domain;uniqueIndex:idx_cert_domain\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(2)
			if f.Name != "IFa4vG" || f.Offset != 24 || f.Type.Size() != 1 || string(f.Tag) != "json:\"auto_renew\" gorm:\"column:auto_renew\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(3)
			if f.Name != "FHaYRa" || f.Offset != 32 || f.Type.Size() != 16 || string(f.Tag) != "json:\"email\" gorm:\"column:email\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(4)
			if f.Name != "VemC_yd" || f.Offset != 48 || f.Type.Size() != 16 || string(f.Tag) != "json:\"provider_type\" gorm:\"column:provider_type\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(5)
			if f.Name != "JdrlDaD1k" || f.Offset != 64 || f.Type.Size() != 16 || string(f.Tag) != "json:\"provider_config\" gorm:\"column:provider_config\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(6)
			if f.Name != "Y1mdfY" || f.Offset != 80 || f.Type.Size() != 16 || string(f.Tag) != "json:\"certContent\" gorm:\"column:cert_content\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(7)
			if f.Name != "SzI2MbNsk" || f.Offset != 96 || f.Type.Size() != 16 || string(f.Tag) != "json:\"keyContent\" gorm:\"column:key_content\"" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(8)
			if f.Name != "ICiRoQoM1K" || f.Offset != 112 || f.Type.Size() != 8 || string(f.Tag) != "json:\"expiryTime\" gorm:\"-\"" {
				t.Fatal(typ, f)
			}
		}
	}
}
