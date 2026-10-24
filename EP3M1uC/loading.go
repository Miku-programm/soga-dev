package EP3M1uC

import (
	"context"
	"os"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"strings"
	"unsafe"
)

// Static pointer at 0x5029218 points to 0x503db20, whose only initialized
// byte is enable=true. Replacements are plain pointer stores in the original.
var routesAt5029218 = &UIfnUsim{CnK0a4: true}
var routesPathAt5096F90 = "/etc/soga/routes.toml"

// CurrentRoutes is an analysis bridge to the original package global.
func CurrentRoutes() *UIfnUsim { return routesAt5029218 }

// Zma0TSvL8 restores 0x1b88780: prefix relative paths without cleaning them.
func Zma0TSvL8(path string) {
	if !strings.HasPrefix(path, "/") {
		path = "/etc/soga/" + path
	}
	routesPathAt5096F90 = path
}

// AycgFFCBT5Zx restores 0x1b88860 and its URL-capturing closure.
func AycgFFCBT5Zx(ctx context.Context) {
	if url := cNT24rG4F.RoutesURL(); url != "" {
		go CAodbhZf.KJa7rJUaLPYe(ctx, url, func(data []byte) error {
			if err := yFE5sGa7LHpa(data); err != nil {
				return err
			}
			muc279Mpa6S4.V8GN5Stk(0, "load routes url <%v> success, routes num: %v, enable: %v, health_check: %v", url, len(routesAt5029218.MMKbp60np), routesAt5029218.CnK0a4, routesAt5029218.H839j9DYDD5k)
			return nil
		})
		return
	}
	cXm2XJR()
	if err := CAodbhZf.SharedFileWatcher().WatchChanged(ctx, func() bool { cXm2XJR(); return true }, routesPathAt5096F90); err != nil {
		muc279Mpa6S4.V8GN5Stk(8, "watch routes config file <%v> failed: %v", routesPathAt5096F90, err)
	}
}

// cXm2XJR restores 0x1b88ac0, including warning/error/info log levels.
func cXm2XJR() {
	data, err := os.ReadFile(routesPathAt5096F90)
	if err != nil {
		muc279Mpa6S4.V8GN5Stk(4, "read routes config file <%v> failed: %v", routesPathAt5096F90, err)
		return
	}
	if err = yFE5sGa7LHpa(data); err != nil {
		muc279Mpa6S4.V8GN5Stk(8, "load routes config file <%v> failed: %v", routesPathAt5096F90, err)
		return
	}
	muc279Mpa6S4.V8GN5Stk(0, "load routes config file <%v> success, routes num: %v, enable: %v, health_check: %v", routesPathAt5096F90, len(routesAt5029218.MMKbp60np), routesAt5029218.CnK0a4, routesAt5029218.H839j9DYDD5k)
}

// yFE5sGa7LHpa restores 0x1b88de0. Empty input preserves both routing and
// health manager; parsing/initialization failure does not replace either.
func yFE5sGa7LHpa(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	routes, err := KmKjIZi1(unsafe.String(unsafe.SliceData(data), len(data)))
	if err != nil {
		return err
	}
	routesAt5029218 = routes
	tCGadgsq(routes)
	return nil
}
