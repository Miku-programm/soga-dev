package inbound

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"recovered.local/soga/transport/realitykdf"
)

func TestHysteriaStateResponseStatusExtraction(t *testing.T) {
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
		if got := rWeAdhA96(wire, request); got != tc.want {
			t.Fatalf("%s: %q want %q", tc.plain, got, tc.want)
		}
	}
	for _, wire := range []string{`invalid json`, `{"data":"!","nonce":"!"}`, `{"data":"AA==","nonce":"AAAAAAAAAAAAAAAA"}`} {
		if got := rWeAdhA96([]byte(wire), request); got != "100005" {
			t.Fatalf("malformed response %s: %s", wire, got)
		}
	}
}

func TestHysteriaStateRequestLocalHTTP(t *testing.T) {
	var paths []string
	var bodies []map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		if r.Method != "POST" || r.Header.Get("Content-Type") != "application/json" {
			t.Error("original request method/content type")
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		bodies = append(bodies, body)
		if r.URL.Path != "/accepted" {
			w.WriteHeader(503)
			return
		}
		w.Write([]byte("invalid encrypted response"))
	}))
	defer server.Close()
	got, err := requestHysteriaConnectionState([]string{server.URL + "/rejected", server.URL + "/accepted", server.URL + "/unreached"})
	if err != nil || got != "100005" {
		t.Fatalf("status %q, error %v", got, err)
	}
	if !reflect.DeepEqual(paths, []string{"/rejected", "/accepted"}) {
		t.Fatalf("attempts %v", paths)
	}
	if reflect.DeepEqual(bodies[0], bodies[1]) {
		t.Fatal("original loop must generate a fresh encrypted request per attempt")
	}
	for _, body := range bodies {
		if len(body) != 3 {
			t.Fatalf("wire fields %v", body)
		}
	}
	paths = nil
	got, err = requestHysteriaConnectionState([]string{server.URL + "/a", server.URL + "/b", server.URL + "/c", server.URL + "/d"})
	if got != "" || err == nil || err.Error() != "check soga failed" || len(paths) != 4 {
		t.Fatalf("all failed: %q %v %v", got, err, paths)
	}
}
