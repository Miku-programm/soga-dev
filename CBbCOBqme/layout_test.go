package CBbCOBqme

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestNativeHysteriaOriginalLayouts(t *testing.T) {
	var listener P3_qJL5v
	var handler cYzxTBfcV
	var stream StUVueP6a
	var udp aK1tZAfWP2z
	var sessions chFbgN
	got := []uintptr{
		unsafe.Sizeof(listener), unsafe.Offsetof(listener.yeJIrGaD12), unsafe.Offsetof(listener.p2xMAL1hk), unsafe.Offsetof(listener.fpLylxmI), unsafe.Offsetof(listener.eIECnHeZd8), unsafe.Offsetof(listener.hcouJtJtvDTP), unsafe.Offsetof(listener.sDLEnpCX),
		unsafe.Sizeof(handler), unsafe.Offsetof(handler.ff2_VHDMO_e), unsafe.Offsetof(handler.lsVoia1), unsafe.Offsetof(handler.chDy7ZwdT), unsafe.Offsetof(handler.gjUa1J3Pmhm), unsafe.Offsetof(handler.tZIs5a3Jb), unsafe.Offsetof(handler.ehiG0t4uYJad), unsafe.Offsetof(handler.dXJZalu06),
		unsafe.Sizeof(BqTbYn4{}), unsafe.Sizeof(stream), unsafe.Offsetof(stream.bQt7p94rWfG), unsafe.Offsetof(stream.z8ptDc4),
		unsafe.Sizeof(udp), unsafe.Offsetof(udp.oPadewGA), unsafe.Offsetof(udp.bibMnwQ), unsafe.Offsetof(udp.v7t5F5tZ), unsafe.Offsetof(udp.oi0B93oah), unsafe.Offsetof(udp.iXQ9zjKttm0F), unsafe.Offsetof(udp.g54qtKk8h), unsafe.Offsetof(udp.viQjY4UoFT6o), unsafe.Offsetof(udp.hlcbMoqH4), unsafe.Offsetof(udp.iIraKZ), unsafe.Offsetof(udp.smLqRCJL3Msj), unsafe.Offsetof(udp.a8Q2aY),
		unsafe.Sizeof(sessions), unsafe.Offsetof(sessions.eAEJVc), unsafe.Offsetof(sessions.ykxy0e), unsafe.Offsetof(sessions.ndlUQyGetN), unsafe.Offsetof(sessions.iuT4giLUja), unsafe.Offsetof(sessions.w7USyrd7), unsafe.Sizeof(kTVWUcbD5rU{}),
	}
	want := []uintptr{80, 16, 32, 40, 48, 56, 64, 88, 16, 24, 40, 44, 56, 72, 80, 48, 40, 16, 32, 128, 16, 24, 40, 48, 56, 64, 72, 80, 96, 104, 120, 120, 16, 24, 40, 48, 64, 8}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("RTTI offsets/sizes got %v want %v", got, want)
	}
}
