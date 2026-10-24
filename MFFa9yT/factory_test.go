package MFFa9yT

import (
	"context"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"reflect"
	"testing"
	"unsafe"
)

func TestNativeFactoryOriginalRegistrations(t *testing.T) {
	expected := map[C.Mpn2JltM]fH6caiUi{"anytls": aboatPBa, "hysteria": hpdxYbrDSPs, "mieru": tzgQP2vl, "ss": aXSbxcDLdeEY, "ssr": x4oEZV, "trojan": f1mPzegGt, "vless": oHr1LeATmS, "v2ray": ag90n_rZjJO}
	if len(inboundConstructorsAt5095E58) != 8 {
		t.Fatal("incomplete/extra registrations")
	}
	for key, want := range expected {
		got := inboundConstructorsAt5095E58[key]
		if got == nil || reflect.ValueOf(got).Pointer() != reflect.ValueOf(want).Pointer() {
			t.Fatalf("%s constructor mismatch", key)
		}
	}
	if unsafe.Sizeof(DV_XPJ4z0{}) != 40 || unsafe.Offsetof(DV_XPJ4z0{}.POpyjFWF8v) != 16 || unsafe.Offsetof(DV_XPJ4z0{}.HG5k0365) != 32 {
		t.Fatal("constructor argument RTTI")
	}
}
func TestNativeFactoryUnknownAndNilBehavior(t *testing.T) {
	for _, key := range []C.Mpn2JltM{"", "vmess", "V2RAY", "vless ", "hysteria2", "ss2022", "bad%v\x00"} {
		got, e := Pb5l5RLOHSi(&DV_XPJ4z0{POpyjFWF8v: key})
		if got != nil || e == nil || e.Error() != "unknown server type: "+string(key) {
			t.Fatalf("%q: %v %v", key, got, e)
		}
	}
	defer func() {
		if recover() == nil {
			t.Fatal("nil argument must panic before construction")
		}
	}()
	Pb5l5RLOHSi(nil)
}

// All eight entries invoke their actual recovered constructors, without fake
// registry replacements. The empty user list also avoids SSR constructor-time
// per-user listeners; no Start call is made.
func TestNativeFactoryConstructsAllEightNativeTypes(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	for _, tc := range []struct {
		key  C.Mpn2JltM
		want reflect.Type
	}{
		{"anytls", reflect.TypeFor[*R_n5ucb]()}, {"hysteria", reflect.TypeFor[*Aq5i0Bwez8]()}, {"mieru", reflect.TypeFor[*ONGT2EHM40r]()}, {"ss", reflect.TypeFor[*Lp07FzBLx9V]()}, {"ssr", reflect.TypeFor[*Ze0fOaAko]()}, {"trojan", reflect.TypeFor[*EfaIFpp]()}, {"vless", reflect.TypeFor[*KHonGaqbIy_6]()}, {"v2ray", reflect.TypeFor[*K4avFiYOPi]()},
	} {
		t.Run(string(tc.key), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			arg := &DV_XPJ4z0{Dmpl48sXM: ctx, POpyjFWF8v: tc.key, HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: &L.Wr5SVP{EMlXEc: 713}}}
			inbound, e := Pb5l5RLOHSi(arg)
			if e != nil {
				t.Fatal(e)
			}
			if reflect.TypeOf(inbound) != tc.want {
				t.Fatalf("%T, want %v", inbound, tc.want)
			}
			defer inbound.Stop()
			if inbound.UserExist(99) {
				t.Fatal("unexpected user")
			}
		})
	}
}
