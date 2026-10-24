package CAodbhZf

import (
	"context"
	"encoding/binary"
	"net"
	"reflect"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/beevik/ntp"
)

func TestNTPDependencyLayouts(t *testing.T) {
	for _, tc := range []struct {
		typ     reflect.Type
		size    uintptr
		offsets []uintptr
	}{
		{reflect.TypeOf(ntp.QueryOptions{}), 128, []uintptr{0, 8, 16, 32, 40, 72, 96, 104, 112, 120}},
		{reflect.TypeOf(ntp.Response{}), 168, []uintptr{0, 8, 32, 40, 48, 56, 60, 64, 88, 96, 104, 112, 120, 128, 144, 152}},
		{reflect.TypeOf(ntp.AuthOptions{}), 32, []uintptr{0, 8, 24}},
	} {
		if tc.typ.Size() != tc.size || tc.typ.NumField() != len(tc.offsets) {
			t.Fatalf("%v layout", tc.typ)
		}
		for i, want := range tc.offsets {
			if tc.typ.Field(i).Offset != want {
				t.Fatalf("%v field %d", tc.typ, i)
			}
		}
	}
	now := time.Now()
	r := ntp.Response{Time: now, ReferenceTime: now.Add(-time.Hour), Stratum: 1}
	if err := r.Validate(); err != nil {
		t.Fatal(err)
	}
	r.RootDispersion = 16*time.Second + 1
	if err := r.Validate(); err != ntp.ErrInvalidDispersion {
		t.Fatalf("dispersion boundary: %v", err)
	}
}

func TestClockOffsetFailureAndCancellation(t *testing.T) {
	oldServers, oldOffset, oldLog := ntpServersAt5099630, atomic.LoadInt64(&timeOffsetAt50C7170), timeLogAt50C6C6A
	t.Cleanup(func() {
		ntpServersAt5099630 = oldServers
		atomic.StoreInt64(&timeOffsetAt50C7170, oldOffset)
		timeLogAt50C6C6A = oldLog
	})
	ntpServersAt5099630 = nil
	atomic.StoreInt64(&timeOffsetAt50C7170, int64(7*time.Second))
	if r, err := aep2YURd_Nd(); r != nil || err == nil || err.Error() != "udp issue" {
		t.Fatalf("empty servers: %v %v", r, err)
	}
	ejBA9tGpop()
	if atomic.LoadInt64(&timeOffsetAt50C7170) != int64(7*time.Second) {
		t.Fatal("failed query cleared offset")
	}
	synctest.Test(t, func(t *testing.T) {
		if got := Nj5UR_g2(); !got.Equal(time.Now().Add(7 * time.Second)) {
			t.Fatal(got)
		}
		if C5Jff5vs0() != time.Now().Unix()+7 {
			t.Fatal("adjusted Unix timestamp")
		}
		ctx, cancel := context.WithCancel(context.Background())
		Kia4I9wKU(ctx, false)
		time.Sleep(2 * time.Hour)
		cancel()
		synctest.Wait()
	})
}

func TestClockSyncWithActualLocalNTP(t *testing.T) {
	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 123})
	if err != nil {
		t.Skipf("local NTP port unavailable: %v", err)
	}
	defer conn.Close()
	oldServers, oldOffset, oldLog := ntpServersAt5099630, atomic.LoadInt64(&timeOffsetAt50C7170), timeLogAt50C6C6A
	t.Cleanup(func() {
		ntpServersAt5099630 = oldServers
		atomic.StoreInt64(&timeOffsetAt50C7170, oldOffset)
		timeLogAt50C6C6A = oldLog
	})
	ntpServersAt5099630 = []string{"127.0.0.1"}
	timeLogAt50C6C6A = false
	done := make(chan error, 1)
	go func() {
		for i := 0; i < 2; i++ {
			_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
			request := make([]byte, 512)
			n, peer, err := conn.ReadFromUDP(request)
			if err != nil {
				done <- err
				return
			}
			if n < 48 {
				done <- net.InvalidAddrError("short NTP request")
				return
			}
			response := make([]byte, 48)
			response[0] = 0x24
			response[1] = 1
			response[3] = 236
			if i == 1 {
				response[1] = 0
			}
			stamp := func(dst []byte, instant time.Time) {
				sec := uint64(instant.Unix() + 2208988800)
				frac := (uint64(instant.Nanosecond()) << 32) / 1e9
				binary.BigEndian.PutUint64(dst, sec<<32|frac)
			}
			now := time.Now().Add(2 * time.Second)
			stamp(response[16:24], now.Add(-time.Hour))
			copy(response[24:32], request[40:48])
			stamp(response[32:40], now)
			stamp(response[40:48], now)
			if _, err = conn.WriteToUDP(response, peer); err != nil {
				done <- err
				return
			}
		}
		done <- nil
	}()
	ejBA9tGpop()
	offset := time.Duration(atomic.LoadInt64(&timeOffsetAt50C7170))
	if offset < 1500*time.Millisecond || offset > 2500*time.Millisecond {
		t.Fatalf("local clock offset: %v", offset)
	}
	ejBA9tGpop()
	if time.Duration(atomic.LoadInt64(&timeOffsetAt50C7170)) != offset {
		t.Fatal("invalid stratum changed cached offset")
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

func TestNTPDependencyRealUDP(t *testing.T) {
	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	done := make(chan error, 1)
	go func() {
		for i := 0; i < 2; i++ {
			_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
			request := make([]byte, 512)
			n, peer, err := conn.ReadFromUDP(request)
			if err != nil {
				done <- err
				return
			}
			if n < 48 {
				done <- net.InvalidAddrError("short NTP request")
				return
			}
			response := make([]byte, 48)
			response[0] = 0x24
			response[1] = 1
			response[3] = 236
			if i == 1 {
				response[1] = 0
			}
			stamp := func(dst []byte, instant time.Time) {
				sec := uint64(instant.Unix() + 2208988800)
				frac := (uint64(instant.Nanosecond()) << 32) / 1e9
				binary.BigEndian.PutUint64(dst, sec<<32|frac)
			}
			now := time.Now().Add(2 * time.Second)
			stamp(response[16:24], now.Add(-time.Hour))
			copy(response[24:32], request[40:48])
			stamp(response[32:40], now)
			stamp(response[40:48], now)
			if _, err = conn.WriteToUDP(response, peer); err != nil {
				done <- err
				return
			}
		}
		done <- nil
	}()
	response, err := ntp.QueryWithOptions(conn.LocalAddr().String(), ntp.QueryOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if err = response.Validate(); err != nil {
		t.Fatal(err)
	}
	if response.ClockOffset < 1500*time.Millisecond || response.ClockOffset > 2500*time.Millisecond {
		t.Fatalf("local clock offset: %v", response.ClockOffset)
	}
	response, err = ntp.QueryWithOptions(conn.LocalAddr().String(), ntp.QueryOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if response.Validate() != ntp.ErrKissOfDeath {
		t.Fatalf("invalid stratum: %v", response.Validate())
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}
