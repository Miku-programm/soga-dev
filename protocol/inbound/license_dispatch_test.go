package inbound

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"

	"recovered.local/soga/config"
	"recovered.local/soga/transport/realitykdf"
)

func TestLicenseResponseStatusExtraction(t *testing.T) {
	request, err := realitykdf.LBab24XA([]byte("offline test"))
	if err != nil {
		t.Fatal(err)
	}
	// Reading the private session bytes is test-only and does not add an export
	// to the original wire type. No licensing endpoint is contacted.
	key := reflect.ValueOf(request).Elem().FieldByName("tbNcCO9yF").Bytes()
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, 12)
	for _, tc := range []struct{ plain, want string }{
		{`{"code":100002,"expiry_time":-1,"max_user_num":-2}`, "100002"},
		{`{"code":-987}`, "-987"}, {`{}`, "0"}, {`null`, "0"}, {`{"CODE":100003}`, "100003"},
		{`{"code":"0"}`, "100005"}, {`not json`, "100005"},
	} {
		sealed := gcm.Seal(nil, nonce, []byte(tc.plain), nil)
		wire, _ := json.Marshal(realitykdf.D4bEdSC7On{LGzvNG: base64.StdEncoding.EncodeToString(sealed), C4YQM8: base64.StdEncoding.EncodeToString(nonce)})
		if got := aNU8Vqc_q(wire, request); got != tc.want {
			t.Fatalf("%s: %q want %q", tc.plain, got, tc.want)
		}
	}
	for _, wire := range []string{`invalid json`, `{"data":"!","nonce":"!"}`, `{"data":"AA==","nonce":"AAAAAAAAAAAAAAAA"}`} {
		if got := aNU8Vqc_q([]byte(wire), request); got != "100005" {
			t.Fatalf("malformed response %s: %s", wire, got)
		}
	}
}

func TestLicensePayloadWireValues(t *testing.T) {
	oldCount := userCountAt50C6DA0
	userCountAt50C6DA0 = 123
	t.Cleanup(func() { userCountAt50C6DA0 = oldCount })
	fields := licensePayload()
	if len(fields) != 17 {
		t.Fatalf("field count %d", len(fields))
	}
	panel, api, address, host, port, name, version := config.LicensePanelFields()
	if !strings.HasSuffix(address, "/") {
		address += "/"
	}
	apiInfo := address + ":" + config.H7eU0uGXU()
	dbInfo := host + ":" + strconv.Itoa(port) + ":" + name
	apiMD5 := md5.Sum([]byte(apiInfo))
	dbMD5 := md5.Sum([]byte(dbInfo))
	domainMD5 := md5.Sum([]byte(gS8fWl()))
	want := map[string]any{
		"user_num": 123, "panel": panel, "api": api, "version": version,
		"api_info": apiInfo, "db_info": dbInfo, "domain_origin": gS8fWl(), "domain": hex.EncodeToString(domainMD5[:]),
		"api_md5": base64.StdEncoding.EncodeToString(apiMD5[:]), "db_md5": base64.StdEncoding.EncodeToString(dbMD5[:]),
	}
	for name, value := range want {
		if fields[name] != value {
			t.Fatalf("%s = %#v want %#v", name, fields[name], value)
		}
	}
	if _, ok := fields["timestamp"].(uint32); !ok {
		t.Fatal("timestamp must retain uint32 truncation")
	}
	if _, ok := fields["rss_total"].(uint64); !ok {
		t.Fatal("total memory type")
	}
	if _, ok := fields["rss"].(uint64); !ok {
		t.Fatal("Go Sys memory type")
	}
	if _, err := bj4dWcdriU(); err != nil {
		t.Fatalf("request construction: %v", err)
	}
}

func TestLicenseMetricsDependencyRTTI(t *testing.T) {
	if unsafe.Sizeof(load.AvgStat{}) != 24 || unsafe.Sizeof(mem.VirtualMemoryStat{}) != 280 {
		t.Fatal("gopsutil struct sizes differ from original RTTI")
	}
	if unsafe.Offsetof(mem.VirtualMemoryStat{}.Total) != 0 {
		t.Fatal("total field offset")
	}
}
