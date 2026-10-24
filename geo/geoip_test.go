package geo

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"google.golang.org/protobuf/proto"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/transport/grpcproto"
)

func geoIPFixture(t *testing.T, code string) []byte {
	t.Helper()
	data, err := proto.Marshal(&grpcproto.E4umdYA{
		OhCky4mjC:  code,
		CnonHlPYKJ: []*grpcproto.ZlC9UehPDkp{{Agmm2h: []byte{192, 0, 2, 0}, HYmWBe: 24}},
	})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestGeoIPOriginalEmbeddedLookup(t *testing.T) {
	result, err := YugPy6709RY("cn")
	if err != nil || result == nil || result.OhCky4mjC != "CN" || len(result.CnonHlPYKJ) == 0 {
		t.Fatalf("original CN lookup: %v %v", result, err)
	}
	if result, err := YugPy6709RY("__does_not_exist__"); result != nil || err == nil || err.Error() != "geoip:__does_not_exist__ not exist" {
		t.Fatalf("absent code: %v %v", result, err)
	}
}

func TestGeoIPReturnObjectAndErrorBranches(t *testing.T) {
	savedData, savedFile := embeddedGeoIPAt502DA40, geoIPFileAt5095C40
	defer func() { embeddedGeoIPAt502DA40, geoIPFileAt5095C40 = savedData, savedFile }()
	geoIPFileAt5095C40 = nil
	for _, data := range [][]byte{nil, {}, {0x10}, {0x0a, 0x05, 0x0a, 0x7f}} {
		embeddedGeoIPAt502DA40 = data
		result, err := YugPy6709RY("CN")
		if result != nil || err == nil || err.Error() != "geoip:CN not exist" {
			t.Fatalf("scan error %x: %v %v", data, result, err)
		}
		if errors.Is(err, codeNotFoundErrorAt5096F80) {
			t.Fatal("original formatted error must not wrap sentinel")
		}
	}
	embeddedGeoIPAt502DA40 = listFixture([]byte{0x0a, 2, 'C', 'N', 0x80})
	result, err := YugPy6709RY("cn")
	if result == nil || err == nil {
		t.Fatalf("protobuf failure discarded object: %v %v", result, err)
	}
	embeddedGeoIPAt502DA40 = listFixture(geoIPFixture(t, "CN"))
	// 映射对象的 Data 失败返回 nil，仍使用内嵌数据。
	geoIPFileAt5095C40 = new(rtx.CCVciOvWaCEr)
	result, err = YugPy6709RY("cn")
	if err != nil || result == nil || len(result.CnonHlPYKJ) != 1 || !bytes.Equal(result.CnonHlPYKJ[0].Agmm2h, []byte{192, 0, 2, 0}) {
		t.Fatalf("nil map fallback: %v %v", result, err)
	}
}

func TestGeoIPMappedOverride(t *testing.T) {
	if os.Getenv("SOGA_GEOIP_CHILD") != "1" {
		command := exec.Command(os.Args[0], "-test.run=^TestGeoIPMappedOverride$")
		command.Env = append(os.Environ(), "SOGA_GEOIP_CHILD=1")
		if output, err := command.CombinedOutput(); err != nil {
			t.Fatalf("child: %v\n%s", err, output)
		}
		return
	}
	path := filepath.Join(t.TempDir(), "geoip.dat")
	if err := os.WriteFile(path, listFixture(geoIPFixture(t, "CUSTOM")), 0600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var err error
	geoIPFileAt5095C40, err = rtx.LWzrRP(ctx, path)
	if err != nil {
		t.Fatal(err)
	}
	result, err := YugPy6709RY("custom")
	if err != nil || result == nil || result.OhCky4mjC != "CUSTOM" {
		t.Fatalf("mapped lookup %v %v", result, err)
	}
	if result, err := YugPy6709RY("CN"); result != nil || err == nil {
		t.Fatalf("unexpected embedded fallback %v %v", result, err)
	}
	if err := geoIPFileAt5095C40.Close(); err != nil {
		t.Fatal(err)
	}
	if result, err := YugPy6709RY("CN"); result == nil || err != nil {
		t.Fatalf("closed map fallback %v %v", result, err)
	}
}
