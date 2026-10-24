package hysteria

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"math"
	"testing"
	"unsafe"
)

// 固定输出来自原 ELF 构造函数受限执行；只模拟内存分配，不执行证书解析。
func TestOriginalConfigurationMachineGoldens(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x0), math.Float64frombits(0x42f0000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000000000000000000000800000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("1", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x8000000000000000), math.Float64frombits(0x42e0000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000000000000000000000800000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("2", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x3ff0000000000000), math.Float64frombits(0x42dfffffffffffff))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000002000000000000f8ffffffffffff0000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("3", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x4059000000000000), math.Float64frombits(0x3e10000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc06000000000400000000000000000000000000000000c8000000000000000000000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("4", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0xbff0000000000000), math.Float64frombits(0x7ff8000000000001))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000000000000000000000800000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("5", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x3fd0000000000000), math.Float64frombits(0xfff0000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000008000000000000000000000000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("6", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x7ff0000000000000), math.Float64frombits(0x7ff0000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000008000000000000000800000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("7", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0xfff0000000000000), math.Float64frombits(0x3fd0000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000000000800000000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("8", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x7ff8000000000001), math.Float64frombits(0xbff0000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000008000000000000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("9", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x3e10000000000000), math.Float64frombits(0x4059000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc060000000004000000000000000000000000000000000000000000000000c800000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("10", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x42dfffffffffffff), math.Float64frombits(0x3ff0000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc060000000004000000000000000000000000000000f8ffffffffffff00000200000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("11", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x42e0000000000000), math.Float64frombits(0x8000000000000000))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000008000000000000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("12", func(t *testing.T) {
		c, e := LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{Leaf: &x509.Certificate{}}}}, nil, math.Float64frombits(0x42f0000000000000), math.Float64frombits(0x0))
		if e != nil {
			t.Fatal(e)
		}
		got := hex.EncodeToString(unsafe.Slice((*byte)(unsafe.Pointer(&c.I3PHCWIdg8EH)), 88))
		if got != "000080000000000000008000000000000000400100000000000040010000000000ac23fc0600000000040000000000000000000000000000000000000000008000000000000000000000000000000000005847f80d000000" {
			t.Fatalf("got %s", got)
		}
	})
}
