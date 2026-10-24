package tkbUwvW3ex

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"unsafe"
)

func TestLicenseCryptoLayoutsAndKDF(t *testing.T) {
	if unsafe.Sizeof(KUD0DUCag{}) != 72 || unsafe.Sizeof(D4bEdSC7On{}) != 32 {
		t.Fatal("RTTI size mismatch")
	}
	key, err := aStQje([]byte("example shared secret"))
	if err != nil {
		t.Fatal(err)
	}
	// Independently computed with Python hmac SHA-256 extract/one-block expand.
	if hex.EncodeToString(key) != "a11da64f1d9969bc44fa3cfdf6d05d326859f997f8be82a377641f878c998b26" {
		t.Fatalf("derived %x", key)
	}
	if _, err = kVljP2(nil, nil, "", 255*32+1); err == nil {
		t.Fatal("oversized expand accepted")
	}
}

func TestLicenseRequestLocalCrypto(t *testing.T) {
	for _, data := range [][]byte{nil, []byte(`{"timestamp":123,"soga_key":"fixture"}`)} {
		req, err := LBab24XA(data)
		if err != nil {
			t.Fatal(err)
		}
		key, err := base64.StdEncoding.DecodeString(req.Nc_QRM4fWhLJ)
		if err != nil {
			t.Fatal(err)
		}
		if _, err = ecdh.P256().NewPublicKey(key); err != nil {
			t.Fatal(err)
		}
		nonce, _ := base64.StdEncoding.DecodeString(req.EK9X3prNY)
		encrypted, _ := base64.StdEncoding.DecodeString(req.YQ02ZfnPFDqE)
		if len(nonce) != 12 || len(req.tbNcCO9yF) != 32 {
			t.Fatal("wrong nonce/session key length")
		}
		block, _ := aes.NewCipher(req.tbNcCO9yF)
		gcm, _ := cipher.NewGCM(block)
		plain, err := gcm.Open(nil, nonce, encrypted, []byte(licenseKeyContext))
		if err != nil || !bytes.Equal(plain, data) {
			t.Fatalf("request decrypt: %q %v", plain, err)
		}
		if _, err = gcm.Open(nil, nonce, encrypted, nil); err == nil {
			t.Fatal("request omitted AAD")
		}
		encoded, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]any
		if err = json.Unmarshal(encoded, &fields); err != nil {
			t.Fatal(err)
		}
		if len(fields) != 3 || fields["data"] != req.YQ02ZfnPFDqE || fields["nonce"] != req.EK9X3prNY || fields["key"] != req.Nc_QRM4fWhLJ {
			t.Fatalf("wire fields %s", encoded)
		}
	}
}

func TestLicenseResponseAuthenticationAndErrors(t *testing.T) {
	req := &KUD0DUCag{tbNcCO9yF: bytes.Repeat([]byte{7}, 32)}
	nonce := bytes.Repeat([]byte{3}, 12)
	block, _ := aes.NewCipher(req.tbNcCO9yF)
	gcm, _ := cipher.NewGCM(block)
	plain := []byte(`{"code":100002}`)
	enc := gcm.Seal(nil, nonce, plain, nil)
	response := D4bEdSC7On{LGzvNG: base64.StdEncoding.EncodeToString(enc), C4YQM8: base64.StdEncoding.EncodeToString(nonce)}
	decoded, err := response.Decode(req)
	if err != nil || !bytes.Equal(decoded, plain) {
		t.Fatalf("response %q %v", decoded, err)
	}
	enc[0] ^= 1
	response.LGzvNG = base64.StdEncoding.EncodeToString(enc)
	if _, err = response.Decode(req); err == nil || !strings.HasPrefix(err.Error(), "decrypt response: ") {
		t.Fatalf("tampered response: %v", err)
	}
	for _, tc := range []struct {
		data, nonce string
		key         []byte
		prefix      string
	}{
		{"!", "!", nil, "decode response data: "}, {"AA==", "!", nil, "decode nonce: "}, {"AA==", "AA==", nil, "create cipher: "},
	} {
		_, err := (&D4bEdSC7On{LGzvNG: tc.data, C4YQM8: tc.nonce}).Decode(&KUD0DUCag{tbNcCO9yF: tc.key})
		if err == nil || !strings.HasPrefix(err.Error(), tc.prefix) || errors.Unwrap(err) == nil {
			t.Fatalf("%s: %v", tc.prefix, err)
		}
	}
}

func TestLicenseResponseWrongNoncePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("original GCM nonce panic lost")
		}
	}()
	_, _ = (&D4bEdSC7On{LGzvNG: "AA==", C4YQM8: "AA=="}).Decode(&KUD0DUCag{tbNcCO9yF: make([]byte, 32)})
}
