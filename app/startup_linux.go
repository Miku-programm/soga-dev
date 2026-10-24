package app

import (
	"context"
	"net"
	CA "recovered.local/soga/util/rtx"
	"recovered.local/soga/net/udppacket"
	Cert "recovered.local/soga/cert"
	C "recovered.local/soga/config"
	Geo "recovered.local/soga/geo"
	Log "recovered.local/soga/util/logx"
	Q "recovered.local/soga/panel/registry"
	"strconv"
	"time"
)

// InitLinux restores 0x1fa8340 and its defer closure 0x200b600.
// The original fatal logger exits with status zero; cancellation is not deferred.
func (app *M0mmPxPJsa) InitLinux() {
	defer func() {
		if value := recover(); value != nil {
			Log.ROUJ82_bIQ(value)
		}
	}()
	app.spsw8VktR9z, app.un5PcsX7aN53 = context.WithCancel(context.Background())
	Log.TRTbrquD(0, "soga v"+C.StartupVersion())
	app.l24UlWt()
	url, key := C.StartupPanelService()
	if url != "" && key != "" {
		app.euCFnZuf()
	} else {
		app.so4m0_oRopC()
	}
}

// euCFnZuf:0x1fa01c0. This supplied binary logs the configured service URL,
// then directly invokes the ordinary startup function. No API request occurs
// inside this entry; this is not a claim about other releases.
func (app *M0mmPxPJsa) euCFnZuf() {
	url, _ := C.StartupPanelService()
	Log.V8GN5Stk(0, "soga panel mode, url: %v", url)
	app.so4m0_oRopC()
}

// so4m0_oRopC restores 0x1fa9220. The flattened machine's original seeds and
// closures were evaluated in a read-only bounded emulator; the recorded calls
// preserve ordering, including errors, certificates, and the license-key read.
func (app *M0mmPxPJsa) so4m0_oRopC() {
	starter, err := AewQql2U(C.FSOZSH7g(C.PanelType()))
	if err != nil {
		Log.ROUJ82_bIQ(err)
	}
	configs := C.SGx0HPPP()
	if len(configs) == 0 {
		Log.ROUJ82_bIQ("缺少 node_id")
	}
	if C.StartupAutoUpdate() {
		udppacket.ZAWxLZe5zfNN(app.spsw8VktR9z)
	}
	printTime := C.SIZIzCa() && !C.KvDEtVwdr1rT()
	CA.Kia4I9wKU(app.spsw8VktR9z, printTime)
	if enabled, hours, siteURL, ipURL := C.StartupGeoUpdate(); enabled {
		Geo.Yk2PlrnXo(app.spsw8VktR9z, time.Duration(hours)*time.Hour, siteURL, ipURL)
	}
	if enabled, addr, port := C.StartupPrometheus(); enabled {
		Wj71D8aYUiw6(net.JoinHostPort(addr, strconv.Itoa(port)))
	}
	if enabled, addr, port := C.StartupPprof(); enabled {
		D7Dl21U(net.JoinHostPort(addr, strconv.Itoa(port)))
	}
	panels := make([]Q.VyFk5uT, 0, len(configs))
	for _, config := range configs {
		panel, err := Q.SMkpVNb(config)
		if err != nil {
			Log.ROUJ82_bIQ(err)
		}
		panels = append(panels, panel)
	}
	cert, key := C.ManualCertificateFiles()
	manual := cert != "" && key != ""
	if !manual && gcTStJ6_zEzn() && gJwSzy(panels) {
		Cert.D18E0adGmOX(app.spsw8VktR9z)
	} else {
		automatic, err := d_7qy2u(panels)
		if err != nil {
			Log.ROUJ82_bIQ(err)
		}
		if automatic {
			if err := Cert.TFwUsh(app.spsw8VktR9z); err != nil {
				Log.ROUJ82_bIQ("apply cert failed:", err)
			}
		} else if manual {
			Cert.Ed4asPq5(app.spsw8VktR9z, cert, key)
		}
	}
	if err := app.vVeTh4DwS(); err != nil {
		Log.ROUJ82_bIQ(err)
	}
	// Both original empty/nonempty-key control-flow paths converge on the same
	// startup loop. Retain the real accessor call; protocol authorization remains
	// in each recovered starter/inbound and is not altered here.
	_ = C.FrFQQhzTR5fY()
	for _, panel := range panels {
		starter.Start(app.spsw8VktR9z, panel)
	}
}
