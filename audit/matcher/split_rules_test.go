package matcher

import (
	"reflect"
	"testing"
)

func TestOriginalRuleSplit(t *testing.T) {
	d, i, p := (VHWq49iRat{"\t", " ip:1.2.3.4 ", "geoip:cn", "IP:case", " port:80 ", "domain:x", "domain:x"}).Split()
	if !reflect.DeepEqual(d, []string{"IP:case", "domain:x", "domain:x"}) || !reflect.DeepEqual(i, []string{"ip:1.2.3.4", "geoip:cn"}) || !reflect.DeepEqual(p, []string{"port:80"}) {
		t.Fatal(d, i, p)
	}
	d, i, p = (VHWq49iRat(nil)).Split()
	if d == nil || i == nil || p == nil {
		t.Fatal("original empty slices are nonnil")
	}
}
