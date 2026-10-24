//go:build linux

package a3PdV_bmpRUK

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

func TestOriginalACMETypeLayouts(t *testing.T) {
	for name, pair := range map[string][2]uintptr{
		"user":          {unsafe.Sizeof(s_CjqL_8{}), 24},
		"configuration": {unsafe.Sizeof(lego.Config{}), 80},
		"request":       {unsafe.Sizeof(certificate.ObtainRequest{}), 144},
		"resource":      {unsafe.Sizeof(certificate.Resource{}), 144},
		"registration":  {unsafe.Sizeof(registration.Resource{}), 112},
	} {
		if pair[0] != pair[1] {
			t.Fatalf("%s size %d != original %d", name, pair[0], pair[1])
		}
	}
	request := reflect.TypeOf(certificate.ObtainRequest{})
	bundle, _ := request.FieldByName("Bundle")
	if bundle.Offset != 96 {
		t.Fatal(bundle)
	}
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	reg := &registration.Resource{}
	user := &s_CjqL_8{key, reg}
	if user.GetEmail() != "" || user.GetPrivateKey() != key || user.GetRegistration() != reg {
		t.Fatal("user accessors")
	}
}

func TestDNSFieldListOriginalOrderAndFreshAllocation(t *testing.T) {
	fields := FHSbLFuj2()
	if len(fields) != 424 || fields[193] != "DNS_KEY" {
		t.Fatal(len(fields))
	}
	const expected = "6e3c1aefa39e2f8575dbb78ddc4df60312e1b039e7456fac7e71c691eade0143"
	if actual := fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(fields, "\x00")))); actual != expected {
		t.Fatal(actual)
	}
	for _, field := range fields {
		if !strings.HasPrefix(field, "DNS_") {
			t.Fatal(field)
		}
	}
	fields[0] = "mutated"
	if FHSbLFuj2()[0] != "DNS_CF_Token" {
		t.Fatal("array shared across calls")
	}
	args, env, count := cKh3OZ4rAaWw("dns_cf")
	if !reflect.DeepEqual(args, []string{"--dns", "dns_cf"}) || env == nil || len(env) != 0 || count != 0 {
		t.Fatal(args, env, count)
	}
}

type finalDataReader struct{ calls int }

func (r *finalDataReader) Read(p []byte) (int, error) { r.calls++; return copy(p, "lost"), io.EOF }

func TestACMEOutputReaderChunkAndTerminalError(t *testing.T) {
	oldOut, oldFlags, oldPrefix := log.Writer(), log.Flags(), log.Prefix()
	defer func() { log.SetOutput(oldOut); log.SetFlags(oldFlags); log.SetPrefix(oldPrefix) }()
	var out bytes.Buffer
	log.SetOutput(&out)
	log.SetFlags(0)
	log.SetPrefix("")
	u9iC_y2SBfunc1(strings.NewReader("a\nb"))
	if out.String() != "a\nb\n" {
		t.Fatal(out.String())
	}
	out.Reset()
	reader := &finalDataReader{}
	u9iC_y2SBfunc2(reader)
	if out.Len() != 0 || reader.calls != 1 {
		t.Fatal("original discards final bytes returned with EOF")
	}
}

func TestStartupWithoutDomainAndCanceledRenewal(t *testing.T) {
	if err := TFwUsh(context.Background()); err == nil || err.Error() != "please configure cert_domain" {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	hRYbCxp(ctx)
}
