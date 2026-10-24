package wd42zQ0FT

import (
	"reflect"
	"sync"
	"testing"
)

func TestCacheRTTI(t *testing.T) {
	rt := reflect.TypeOf(NG0hCZyRHK{})
	if rt.Size() != 48 || rt.NumField() != 1 || rt.Field(0).Name != "icFkNd" || rt.Field(0).Offset != 0 || rt.Field(0).Type != reflect.TypeOf(sync.Map{}) {
		t.Fatalf("cache layout: %v/%d", rt, rt.Size())
	}
	rt = reflect.TypeOf(LAQMf2tPf{})
	if rt.Size() != 32 || rt.NumField() != 2 || rt.Field(0).Name != "Z49q4d" || rt.Field(0).Type.Kind() != reflect.String || rt.Field(1).Name != "U6F6cw_aOk" || rt.Field(1).Type != reflect.TypeFor[any]() || rt.Field(1).Offset != 16 {
		t.Fatalf("entry layout: %v/%d", rt, rt.Size())
	}
}
func TestETagStateAndObjectIdentity(t *testing.T) {
	var c NG0hCZyRHK
	if v, ok := c.GetObj("missing"); ok || v != nil {
		t.Fatal(v, ok)
	}
	if v, ok := c.CheckETag("missing", ""); ok || v != nil {
		t.Fatal(v, ok)
	}
	if c.GetETag("missing") != "" {
		t.Fatal("miss tag")
	}
	body := []byte("one")
	c.AddETag("url", `"A"`, body)
	body[0] = 'O'
	if v, ok := c.CheckETag("url", `"A"`); !ok || string(v.([]byte)) != "One" {
		t.Fatal(v, ok)
	}
	if v, ok := c.CheckETag("url", `"a"`); ok || v != nil {
		t.Fatal(v, ok)
	}
	if c.GetETag("url") != `"A"` {
		t.Fatal("tag mismatch")
	}
	c.AddETag("url", "", nil)
	if v, ok := c.GetObj("url"); !ok || v != nil {
		t.Fatal(v, ok)
	}
	if v, ok := c.CheckETag("url", ""); !ok || v != nil {
		t.Fatal(v, ok)
	}
}
func TestCacheRejectsCorruptStoredType(t *testing.T) {
	var c NG0hCZyRHK
	c.icFkNd.Store("bad", 123)
	for _, f := range []func(){func() { c.GetObj("bad") }, func() { c.CheckETag("bad", "") }} {
		func() {
			defer func() {
				if recover() == nil {
					t.Fatal("missing type-assertion panic")
				}
			}()
			f()
		}()
	}
}
func TestConcurrentIndependentCacheEntries(t *testing.T) {
	var c NG0hCZyRHK
	var wg sync.WaitGroup
	for _, key := range []string{"a", "b", "c", "d"} {
		wg.Go(func() {
			for range 100 {
				c.AddETag(key, key, key)
				v, ok := c.CheckETag(key, key)
				if !ok || v != key {
					t.Errorf("cache %q: %v %v", key, v, ok)
				}
			}
		})
	}
	wg.Wait()
}
