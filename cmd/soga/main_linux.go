// Command soga restores the supplied ELF's main package. It is recovered source,
// not the author's original source tree. Original machine entry: 0x2019c00.
package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	mathrand "math/rand"
	"os"
	"os/signal"
	App "recovered.local/soga/app"
	Routes "recovered.local/soga/net/routing"
	Cert "recovered.local/soga/cert"
	C "recovered.local/soga/config"
	Geo "recovered.local/soga/geo"
	Log "recovered.local/soga/util/logx"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"syscall"
	"time"
)

var (
	versionAt509EB80    bool
	configPathAt509EB88 string
	blockPathAt509EB98  string
	whitePathAt509EBA8  string
	dnsPathAt509EBB8    string
	routesPathAt509EBC8 string
)

// main.init.0:0x20187c0. The dynamically registered configuration and DNS
// environment flags intentionally default to empty strings.
func init() {
	for _, name := range C.ConfigurationFlagNames() {
		flag.String(name, "", "")
	}
	for _, name := range Cert.FHSbLFuj2() {
		flag.String(name, "", "")
	}
	flag.BoolVar(&versionAt509EB80, "v", false, "Show version")
	flag.BoolVar(&versionAt509EB80, "version", false, "Show version")
	flag.StringVar(&configPathAt509EB88, "c", "/etc/soga/soga.conf", "Set config file")
	flag.StringVar(&blockPathAt509EB98, "b", "/etc/soga/blockList", "Set blockList file")
	flag.StringVar(&whitePathAt509EBA8, "w", "/etc/soga/whiteList", "Set whiteList file")
	flag.StringVar(&dnsPathAt509EBB8, "d", "/etc/soga/dns.yml", "Set dns config file")
	flag.StringVar(&routesPathAt509EBC8, "r", "/etc/soga/routes.toml", "Set routes config file")
}

// main.init.1:0x2019b20 and its actual goroutine 0x2019f40. Returning from a
// goroutine while it owns an OS thread causes the Go runtime to retire it.
func init() { go monitorThreads() }
func monitorThreads() {
	ticker := time.NewTicker(120 * time.Second)
	defer ticker.Stop()
	profile := pprof.Lookup("threadcreate")
	baseline := max(8, 4*runtime.NumCPU())
	for range ticker.C {
		n := (profile.Count() - baseline) / 3
		for ; n > 0; n-- {
			killOne()
		}
	}
}

// main.killOne:0x2019b60 and closure 0x2019f20. No UnlockOSThread call exists.
func killOne() { go func() { runtime.LockOSThread() }() }

// main.init.2:0x2019ba0. The original ignores crypto/rand.Read's result.
func init() {
	var seed [8]byte
	rand.Read(seed[:])
	mathrand.Seed(int64(binary.BigEndian.Uint64(seed[:])))
	debug.SetGCPercent(50)
}

func main() {
	flag.CommandLine.Parse(os.Args[1:])
	if versionAt509EB80 {
		fmt.Fprintln(os.Stdout, C.StartupVersion())
		return
	}
	C.MkwNWciPHS(configPathAt509EB88)
	C.Dgd6obkvt(blockPathAt509EB98)
	C.AcIcHDv(whitePathAt509EBA8)
	C.QJZTKBo(dnsPathAt509EBB8)
	Routes.Zma0TSvL8(routesPathAt509EBC8)
	Geo.Tac58sht()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	// main.main.func1:0x2019e80. Only SIGUSR1 reloads. The original does not
	// unregister signal delivery and does not include Invoke's error in its log.
	if err := C.DependencyContainer().Invoke(func(app *App.M0mmPxPJsa) {
		for {
			app.InitLinux()
			sig := <-signals
			app.InitUnix()
			if sig != syscall.SIGUSR1 {
				return
			}
		}
	}); err != nil {
		Log.TRTbrquD(8, "unknown error")
	}
}
