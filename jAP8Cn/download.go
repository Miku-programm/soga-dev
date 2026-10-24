package jAP8Cn

import (
	"context"
	"io"
	"net/http"
	A "recovered.local/soga/CAodbhZf"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/muc279Mpa6S4"
	"sync/atomic"
	"time"
)

// Original global0x5095C48 is written only after at least one download updates.
var updateTimeAt5095C48 atomic.Pointer[time.Time]

// MCOm37H:0xea6a60 returns the snapshot even when it is not newer.
func MCOm37H(previous time.Time) (time.Time, bool) {
	snapshot := updateTimeAt5095C48.Load()
	if snapshot == nil {
		return time.Time{}, false
	}
	value := *snapshot
	if value.IsZero() {
		return value, false
	}
	return value, value.After(previous)
}

// Yk2PlrnXo:0xea6b20, callback0xeaa9e0. The first download waits for a tick.
// Cancellation stops the timer; it is not passed into a running HTTP download.
func Yk2PlrnXo(ctx context.Context, interval time.Duration, geositeURL, geoipURL string) {
	L.V8GN5Stk(0, "start geo update, interval: %v", interval)
	A.SwUrlM3M(ctx, func() { rIljlGNSX7VY(geositeURL, geoipURL) }, interval)
}

// rIljlGNSX7VY:0xea6cc0 creates a shared ten-minute HTTP client for two jobs.
func rIljlGNSX7VY(geositeURL, geoipURL string) {
	updateGeoFiles(&http.Client{Timeout: 600 * time.Second}, geositeURL, geoipURL, HaC1u4AhOL2o, ZrUpfY5bpa)
}

// Analyst-named factoring of the original two closures0xeaaea0/0xeaaa20.
// Both always return nil to the parallel runner; failures are logged locally.
// There is deliberately no HTTP status validation before the original writer.
func updateGeoFiles(client *http.Client, geositeURL, geoipURL string, siteWriter, ipWriter func(io.Reader) (bool, error)) {
	var changed atomic.Bool
	download := func(url string, writer func(io.Reader) (bool, error), start, requestError, success, unchanged, writeError string) func() error {
		return func() error {
			L.V8GN5Stk(-4, start, url)
			response, err := client.Get(url)
			if err != nil {
				L.V8GN5Stk(4, requestError, url, err)
				return nil
			}
			defer response.Body.Close()
			updated, err := writer(response.Body)
			if err != nil {
				L.V8GN5Stk(4, writeError, err)
			} else if updated {
				L.V8GN5Stk(0, "%s", success)
				changed.Store(true)
			} else {
				L.V8GN5Stk(0, "%s", unchanged)
			}
			return nil
		}
	}
	_ = H.D5aDHlQDLp(
		download(geositeURL, siteWriter, "downloading geisite url <%v>", "download geosite url <%v> failed: %v", "update geosite data success", "geosite data not changed", "update geosite data failed: %v"),
		download(geoipURL, ipWriter, "downloading geoip url <%v>", "download geoip url <%v> failed: %v", "update geoip data success", "geoip data not changed", "update geoip data failed: %v"),
	)
	if changed.Load() {
		now := time.Now()
		updateTimeAt5095C48.Store(&now)
	}
}
