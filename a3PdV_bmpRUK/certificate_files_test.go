package a3PdV_bmpRUK

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestOriginalCertificateLayoutAndTemplate(t *testing.T) {
	if unsafe.Sizeof(oWCp7YUAFId{}) != 64 {
		t.Fatal("snapshot layout")
	}
	typ := reflect.TypeOf(x509.Certificate{})
	for name, offset := range map[string]uintptr{"KeyUsage": 0x2f8, "BasicConstraintsValid": 0x378, "IsCA": 0x379} {
		f, _ := typ.FieldByName(name)
		if f.Offset != offset {
			t.Fatalf("%s offset %#x", name, f.Offset)
		}
	}
	for _, domain := range []string{"example.test", "127.0.0.1", "2001:db8::1"} {
		t.Run(domain, func(t *testing.T) {
			dir := t.TempDir()
			crt, key := filepath.Join(dir, "sub", "cert.pem"), filepath.Join(dir, "sub", "key.pem")
			before := time.Now()
			if err := lLTDzR(domain, crt, key, 1); err != nil {
				t.Fatal(err)
			}
			pair, err := tls.LoadX509KeyPair(crt, key)
			if err != nil {
				t.Fatal(err)
			}
			cert, err := x509.ParseCertificate(pair.Certificate[0])
			if err != nil {
				t.Fatal(err)
			}
			if cert.Subject.CommonName != domain || !cert.IsCA || !cert.BasicConstraintsValid || cert.KeyUsage != 37 {
				t.Fatalf("template %+v", cert)
			}
			if !reflect.DeepEqual(cert.ExtKeyUsage, []x509.ExtKeyUsage{1, 2}) {
				t.Fatal(cert.ExtKeyUsage)
			}
			if err = cert.VerifyHostname(domain); err != nil {
				t.Fatal(err)
			}
			if cert.NotBefore.Before(before.Add(-time.Hour-time.Second)) || cert.NotAfter.Before(before.AddDate(1, 0, 0).Add(-time.Second)) {
				t.Fatal("validity")
			}
			data, _ := os.ReadFile(key)
			block, _ := pem.Decode(data)
			if block.Type != "EC PRIVATE KEY" {
				t.Fatal(block.Type)
			}
			if err = RUKs8pSwE4(crt, key); err != nil {
				t.Fatal(err)
			}
			if err = lLTDzR(domain, crt, key, -1); err != nil {
				t.Fatal(err)
			}
			if err = RUKs8pSwE4(crt, key); err == nil || !strings.HasPrefix(err.Error(), "cached panel cert expired at ") {
				t.Fatal(err)
			}
		})
	}
}

func TestCertificateFileWritingAndSnapshot(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", "cert")
	if err := wLQ7ZW(path, "first", 0600); err != nil {
		t.Fatal(err)
	}
	old, _ := os.Stat(path)
	if err := wLQ7ZW(path, "first", 0644); err != nil {
		t.Fatal(err)
	}
	same, _ := os.Stat(path)
	if same.Mode() != old.Mode() || same.ModTime() != old.ModTime() {
		t.Fatal("unchanged data changed metadata")
	}
	if err := wLQ7ZW(path, "second", 0644); err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(path)
	if string(data) != "second" {
		t.Fatal(string(data))
	}
	snapshot, err := eG4MicXvEEf(path, path)
	if err != nil || snapshot.fErFXoSf7Sv != 6 || snapshot.eMARy7uR != 6 {
		t.Fatal(snapshot, err)
	}
	snapshot, err = eG4MicXvEEf(path, path+"missing")
	if err == nil || snapshot != (oWCp7YUAFId{}) {
		t.Fatal(snapshot, err)
	}
	targetDir := filepath.Join(dir, "directory")
	os.Mkdir(targetDir, 0755)
	if err = wLQ7ZW(targetDir, "bad", 0600); err == nil {
		t.Fatal("rename onto directory should fail")
	}
	leftovers, _ := filepath.Glob(targetDir + ".tmp*")
	if len(leftovers) != 0 {
		t.Fatal(leftovers)
	}
}

func TestSelfCertificateCacheAndFailures(t *testing.T) {
	old := storageDirectoryAt5096D90
	storageDirectoryAt5096D90 = t.TempDir()
	defer func() { storageDirectoryAt5096D90 = old }()
	for _, domain := range []string{"", ".", "..", "../bad", "dir/bad"} {
		if _, _, err := EEGj3Ra(domain); err == nil {
			t.Fatalf("accepted %q", domain)
		}
	}
	cert, key, err := EEGj3Ra("example.test")
	if err != nil {
		t.Fatal(err)
	}
	c2, k2, err := EEGj3Ra("example.test")
	if err != nil || cert != c2 || key != k2 {
		t.Fatal("cache not reused", err)
	}
	if _, err = tls.X509KeyPair([]byte(cert), []byte(key)); err != nil {
		t.Fatal(err)
	}
	missing := filepath.Join(t.TempDir(), "absent")
	if err = RUKs8pSwE4(missing, missing); err == nil || err.Error() != "cached panel cert not found" {
		t.Fatal(err)
	}
	if _, err = sur_b39JM6l(missing); err == nil || err.Error() != "cert file not exist: "+missing {
		t.Fatal(err)
	}
	os.WriteFile(missing, []byte("not pem"), 0600)
	if _, err = sur_b39JM6l(missing); err == nil || err.Error() != "cert file invalid: "+missing {
		t.Fatal(err)
	}
	a, b := J6c4NB_R(-3)
	if a != "/etc/soga/panel_-3.crt" || b != "/etc/soga/panel_-3.key" {
		t.Fatal(a, b)
	}
}
