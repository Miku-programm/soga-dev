package cNT24rG4F

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

// 原 0x1342300/0x1342900 指令在 48 种状态下的受限执行结果。
// 原进程未运行；newproc 被记录，未在仿真器中启动。
var panelStateMachineGoldens = []struct {
	name         string
	flag, mirror bool
	marker       uint64
	key          string
	result, exit bool
}{
	{"Gh2DnstC", false, false, 0x0, "", true, true},
	{"Gh2DnstC", false, false, 0x0, "x", true, true},
	{"Gh2DnstC", false, false, 0xa5b6c7d8a5b6c7d8, "", true, true},
	{"Gh2DnstC", false, false, 0xa5b6c7d8a5b6c7d8, "x", true, true},
	{"Gh2DnstC", false, false, 0x5a4b3c2d5a4b3c2d, "", true, true},
	{"Gh2DnstC", false, false, 0x5a4b3c2d5a4b3c2d, "x", true, true},
	{"Gh2DnstC", false, true, 0x0, "", true, true},
	{"Gh2DnstC", false, true, 0x0, "x", true, true},
	{"Gh2DnstC", false, true, 0xa5b6c7d8a5b6c7d8, "", true, true},
	{"Gh2DnstC", false, true, 0xa5b6c7d8a5b6c7d8, "x", true, true},
	{"Gh2DnstC", false, true, 0x5a4b3c2d5a4b3c2d, "", true, false},
	{"Gh2DnstC", false, true, 0x5a4b3c2d5a4b3c2d, "x", false, false},
	{"Gh2DnstC", true, false, 0x0, "", true, true},
	{"Gh2DnstC", true, false, 0x0, "x", true, true},
	{"Gh2DnstC", true, false, 0xa5b6c7d8a5b6c7d8, "", true, false},
	{"Gh2DnstC", true, false, 0xa5b6c7d8a5b6c7d8, "x", true, false},
	{"Gh2DnstC", true, false, 0x5a4b3c2d5a4b3c2d, "", true, true},
	{"Gh2DnstC", true, false, 0x5a4b3c2d5a4b3c2d, "x", true, true},
	{"Gh2DnstC", true, true, 0x0, "", true, true},
	{"Gh2DnstC", true, true, 0x0, "x", true, true},
	{"Gh2DnstC", true, true, 0xa5b6c7d8a5b6c7d8, "", true, true},
	{"Gh2DnstC", true, true, 0xa5b6c7d8a5b6c7d8, "x", true, true},
	{"Gh2DnstC", true, true, 0x5a4b3c2d5a4b3c2d, "", true, true},
	{"Gh2DnstC", true, true, 0x5a4b3c2d5a4b3c2d, "x", true, true},
	{"JatfUF38O", false, false, 0x0, "", true, true},
	{"JatfUF38O", false, false, 0x0, "x", true, true},
	{"JatfUF38O", false, false, 0xa5b6c7d8a5b6c7d8, "", true, true},
	{"JatfUF38O", false, false, 0xa5b6c7d8a5b6c7d8, "x", true, true},
	{"JatfUF38O", false, false, 0x5a4b3c2d5a4b3c2d, "", true, true},
	{"JatfUF38O", false, false, 0x5a4b3c2d5a4b3c2d, "x", true, true},
	{"JatfUF38O", false, true, 0x0, "", true, true},
	{"JatfUF38O", false, true, 0x0, "x", true, true},
	{"JatfUF38O", false, true, 0xa5b6c7d8a5b6c7d8, "", true, true},
	{"JatfUF38O", false, true, 0xa5b6c7d8a5b6c7d8, "x", true, true},
	{"JatfUF38O", false, true, 0x5a4b3c2d5a4b3c2d, "", true, false},
	{"JatfUF38O", false, true, 0x5a4b3c2d5a4b3c2d, "x", false, false},
	{"JatfUF38O", true, false, 0x0, "", true, true},
	{"JatfUF38O", true, false, 0x0, "x", true, true},
	{"JatfUF38O", true, false, 0xa5b6c7d8a5b6c7d8, "", true, false},
	{"JatfUF38O", true, false, 0xa5b6c7d8a5b6c7d8, "x", true, false},
	{"JatfUF38O", true, false, 0x5a4b3c2d5a4b3c2d, "", true, true},
	{"JatfUF38O", true, false, 0x5a4b3c2d5a4b3c2d, "x", true, true},
	{"JatfUF38O", true, true, 0x0, "", true, true},
	{"JatfUF38O", true, true, 0x0, "x", true, true},
	{"JatfUF38O", true, true, 0xa5b6c7d8a5b6c7d8, "", true, true},
	{"JatfUF38O", true, true, 0xa5b6c7d8a5b6c7d8, "x", true, true},
	{"JatfUF38O", true, true, 0x5a4b3c2d5a4b3c2d, "", true, true},
	{"JatfUF38O", true, true, 0x5a4b3c2d5a4b3c2d, "x", true, true},
}

func TestPanelStateOriginalMachineGoldens(t *testing.T) {
	apply := func(index int) bool {
		tc := panelStateMachineGoldens[index]
		keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800 = tc.flag, tc.mirror, tc.marker
		configAt5096DC0.RM5SFRF = tc.key
		if tc.name == "Gh2DnstC" {
			return Gh2DnstC()
		}
		return JatfUF38O()
	}
	if child := os.Getenv("SOGA_PANEL_STATE_CASE"); child != "" {
		index, err := strconv.Atoi(child)
		if err != nil {
			os.Exit(97)
		}
		_, _ = os.Stdout.WriteString("ready\n")
		if apply(index) != panelStateMachineGoldens[index].result {
			os.Exit(98)
		}
		time.Sleep(time.Second)
		os.Exit(99)
	}
	flag, mirror, marker, key := keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800, configAt5096DC0.RM5SFRF
	t.Cleanup(func() {
		keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800 = flag, mirror, marker
		configAt5096DC0.RM5SFRF = key
	})
	for index, tc := range panelStateMachineGoldens {
		t.Run(fmt.Sprintf("%s/%d", tc.name, index), func(t *testing.T) {
			if !tc.exit {
				if got := apply(index); got != tc.result {
					t.Fatalf("got %v want %v", got, tc.result)
				}
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestPanelStateOriginalMachineGoldens$")
			cmd.Env = append(os.Environ(), "SOGA_PANEL_STATE_CASE="+strconv.Itoa(index))
			out, err := cmd.CombinedOutput()
			if err != nil || !bytes.Equal(out, []byte("ready\n")) {
				t.Fatalf("original async exit: %v %q", err, out)
			}
		})
	}
}
