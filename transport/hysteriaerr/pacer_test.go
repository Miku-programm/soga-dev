package hysteriaerr

import (
	"github.com/apernet/quic-go/congestion"
	"reflect"
	"testing"
	"time"
)

func TestPacerOriginalLayout(t *testing.T) {
	ty := reflect.TypeFor[OQpLDX]()
	if ty.Size() != 48 {
		t.Fatal(ty.Size())
	}
	names := []string{"gpqkhK", "jyPhMc", "rh7eeI", "clqvXsSOvk"}
	off := []uintptr{0, 8, 16, 40}
	for i, n := range names {
		f := ty.Field(i)
		if f.Name != n || f.Offset != off[i] || f.Tag != "" {
			t.Fatal(f)
		}
	}
}
func TestPacerBudgetAndSend(t *testing.T) {
	calls := 0
	p := NewFromInline(func() congestion.ByteCount { calls++; return 1000000 })
	now := time.Unix(1, 0)
	if p.Budget(now) != 12800 || calls != 1 || !p.TimeUntilSend().IsZero() {
		t.Fatal(p, calls)
	}
	p.SentPacket(now, 12800)
	if p.gpqkhK != 0 {
		t.Fatal(p)
	}
	if got := p.TimeUntilSend(); !got.Equal(now.Add(1280 * time.Microsecond)) {
		t.Fatal(got)
	}
	if got := p.Budget(now.Add(500 * time.Microsecond)); got != 500 {
		t.Fatal(got)
	}
	p.SentPacket(now, 100000)
	if p.gpqkhK != 0 {
		t.Fatal(p)
	}
	p.SetMaxDatagramSize(1500)
	if p.Budget(now.Add(time.Second)) != 15000 {
		t.Fatal(p)
	}
}
func TestPacerOriginalOverflowAndZeroBandwidth(t *testing.T) {
	now := time.Unix(1, 0)
	p := NewFromInline(func() congestion.ByteCount { return 1 << 40 })
	p.rh7eeI = now
	p.gpqkhK = 0
	if got := p.Budget(now.Add(8388608 * time.Nanosecond)); got != 4000000*(1<<40)/1000000000 {
		t.Fatal(got)
	}
	p.clqvXsSOvk = func() congestion.ByteCount { return 0 }
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("original integer division by zero")
			}
		}()
		p.TimeUntilSend()
	}()
}
