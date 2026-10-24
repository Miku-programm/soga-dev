package geo

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/transport/grpcproto"
)

func siteFixture(t *testing.T, code, domain string) []byte {
	t.Helper()
	data, err := proto.Marshal(&grpcproto.P3tOWWIHzX{
		I0dBJNoqAhc: code,
		HGtbzYnVdt0: []*grpcproto.HTBZ_woRreQ{{AfECsc: 2, DSsRH80v92kO: domain}},
	})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func listFixture(sites ...[]byte) []byte {
	var result []byte
	for _, site := range sites {
		result = protowire.AppendTag(result, 1, protowire.BytesType)
		result = protowire.AppendBytes(result, site)
	}
	return result
}

func TestGeoSiteWireMatchingAndCountryReset(t *testing.T) {
	first := siteFixture(t, "CN", "first.example")
	second := siteFixture(t, "cn", "second.example")
	data := listFixture(siteFixture(t, "US", "other.example"), first, second)
	site, err := mcrrOY(data, "cN")
	if err != nil || site == nil || len(site.HGtbzYnVdt0) != 1 || site.HGtbzYnVdt0[0].DSsRH80v92kO != "first.example" {
		t.Fatalf("first match: %v %v", site, err)
	}
	if site.I0dBJNoqAhc != "" {
		t.Fatalf("Unmarshal reset semantics lost: %q", site.I0dBJNoqAhc)
	}
	if site, err := mcrrOY(data, "absent"); site != nil || err != nil {
		t.Fatalf("missing: %v %v", site, err)
	}
	if site, err := sUZtoMEpw_(siteFixture(t, "K", "unicode.example"), "k"); site == nil || err != nil {
		t.Fatalf("Unicode EqualFold differs: %v %v", site, err)
	}
	// 原 loader 只检查首字段，不搜索后面的 country_code。
	prefix := protowire.AppendString(protowire.AppendTag(nil, 3, protowire.BytesType), "unknown")
	if site, err := sUZtoMEpw_(append(prefix, first...), "cn"); site != nil || err != nil {
		t.Fatalf("accepted a non-leading country field: %v %v", site, err)
	}
	// 后续重复 country_code 仍由 protobuf 正常解码。
	duplicate := protowire.AppendString(protowire.AppendTag(first, 1, protowire.BytesType), "later")
	if site, err := sUZtoMEpw_(duplicate, "cn"); err != nil || site.I0dBJNoqAhc != "later" {
		t.Fatalf("duplicate: %v %v", site, err)
	}
}

func TestGeoSiteWireMalformedErrors(t *testing.T) {
	for _, tc := range []struct {
		name   string
		data   []byte
		prefix string
		inner  bool
	}{
		{"list tag", []byte{0x80}, "failed to decode tag:", false},
		{"list length", []byte{0x0a, 0x80}, "failed to decode length:", false},
		{"list unknown field order", []byte{0x10, 0x01}, "failed to consume field:", false},
		{"site tag", []byte{0x80}, "failed to decode geoSite tag:", true},
		{"site country bytes", []byte{0x0a, 5, 'x'}, "failed to decode country code:", true},
		{"site payload", []byte{0x0a, 2, 'c', 'n', 0x80}, "failed to unmarshal GeoSite:", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			if tc.inner {
				_, err = sUZtoMEpw_(tc.data, "cn")
			} else {
				_, err = mcrrOY(tc.data, "cn")
			}
			if err == nil || !strings.HasPrefix(err.Error(), tc.prefix) {
				t.Fatalf("error %v", err)
			}
		})
	}
	t.Run("out of bounds retains panic", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Fatal("missing original bounds panic")
			}
		}()
		_, _ = mcrrOY([]byte{0x0a, 0x7f}, "cn")
	})
}

func TestOriginalEmbeddedGeoSiteLookup(t *testing.T) {
	if len(embeddedGeoSiteAt502DA20) != 2144618 || fmt.Sprintf("%x", sha256.Sum256(embeddedGeoSiteAt502DA20)) != "441d689a2f129132e0dd4c51517fe370a2e66eb2cccb86e4f5cd010559d0b5ce" {
		t.Fatal("embedded GeoSite bytes differ")
	}
	site, found, err := ZRQq51fC00J("CN")
	if err != nil || !found || site == nil || len(site.HGtbzYnVdt0) == 0 {
		t.Fatalf("original CN lookup: %v %v %v", site, found, err)
	}
	if site, found, err := ZRQq51fC00J("__does_not_exist__"); site != nil || found || err != nil {
		t.Fatalf("absent: %v %v %v", site, found, err)
	}
}

func TestOriginalEmbeddedGeoIPBytes(t *testing.T) {
	if len(embeddedGeoIPAt502DA40) != 23058591 || fmt.Sprintf("%x", sha256.Sum256(embeddedGeoIPAt502DA40)) != "e9002979e0df72bce1c8751ff70725386594c551db684b7a232935b8b2bb8aa2" {
		t.Fatal("embedded GeoIP bytes differ from ELF")
	}
}

func TestGeoFileInitialization(t *testing.T) {
	if os.Getenv("SOGA_GEO_INIT_CHILD") != "1" {
		command := exec.Command(os.Args[0], "-test.run=^TestGeoFileInitialization$")
		command.Env = append(os.Environ(), "SOGA_GEO_INIT_CHILD=1")
		if output, err := command.CombinedOutput(); err != nil {
			t.Fatalf("child: %v\n%s", err, output)
		}
		return
	}
	path := filepath.Join(t.TempDir(), "test.dat")
	var mapped *rtx.CCVciOvWaCEr
	initializeGeoFile(path, []byte("embedded bytes"), &mapped, "write %v failed: %v", "map %v failed: %v")
	if mapped == nil {
		t.Fatal("missing file not created and mapped")
	}
	data, err := mapped.Data()
	if err != nil || string(data) != "embedded bytes" {
		t.Fatalf("created %q %v", data, err)
	}
	if err := mapped.Close(); err != nil {
		t.Fatal(err)
	}
	initializeGeoFile(path, []byte("must not overwrite"), &mapped, "write %v failed: %v", "map %v failed: %v")
	data, err = mapped.Data()
	if err != nil || string(data) != "embedded bytes" {
		t.Fatalf("existing %q %v", data, err)
	}
	if err := mapped.Close(); err != nil {
		t.Fatal(err)
	}
	previous := mapped
	initializeGeoFile(path+"/missing-parent/file", []byte("data"), &mapped, "write %v failed: %v", "map %v failed: %v")
	if mapped != previous {
		t.Fatal("write failure overwrote previous mapping")
	}
	initializeGeoFile(filepath.Dir(path), nil, &mapped, "write %v failed: %v", "map %v failed: %v")
	if mapped != nil {
		t.Fatal("mapping failure did not clear global")
	}
}

func TestGeoSiteMappedOverrideAndFallback(t *testing.T) {
	if os.Getenv("SOGA_GEOSITE_CHILD") != "1" {
		command := exec.Command(os.Args[0], "-test.run=^TestGeoSiteMappedOverrideAndFallback$")
		command.Env = append(os.Environ(), "SOGA_GEOSITE_CHILD=1")
		if output, err := command.CombinedOutput(); err != nil {
			t.Fatalf("child: %v\n%s", err, output)
		}
		return
	}
	path := filepath.Join(t.TempDir(), "geosite.dat")
	data := listFixture(siteFixture(t, "custom", "mapped.example"))
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var err error
	geoSiteFileAt5095C38, err = rtx.LWzrRP(ctx, path)
	if err != nil {
		t.Fatal(err)
	}
	site, found, err := ZRQq51fC00J("custom")
	if err != nil || !found || site.HGtbzYnVdt0[0].DSsRH80v92kO != "mapped.example" {
		t.Fatalf("override: %v %v %v", site, found, err)
	}
	if _, found, err := ZRQq51fC00J("CN"); found || err != nil {
		t.Fatalf("embedded used while map available: %v %v", found, err)
	}
	if err := geoSiteFileAt5095C38.Close(); err != nil {
		t.Fatal(err)
	}
	if _, found, err := ZRQq51fC00J("CN"); !found || err != nil {
		t.Fatalf("nil map did not fall back: %v %v", found, err)
	}
	geoSiteFileAt5095C38 = nil
	embeddedGeoSiteAt502DA20 = []byte{0x80}
	if _, found, err := ZRQq51fC00J("cn"); found || err == nil || !strings.HasPrefix(err.Error(), "find geosite failed: failed to decode tag:") {
		t.Fatalf("wrapper: %v %v", found, err)
	}
}
