package l_5zgG8VLP

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func serverHello(cipher uint16) []byte {
	data := make([]byte, 100)
	copy(data, []byte{0x16, 3, 3, 0, 95, 2})
	data[43] = 0
	binary.BigEndian.PutUint16(data[44:], cipher)
	copy(data[70:], []byte{0, 0x2b, 0, 2, 3, 4})
	return data
}
func TestVisionTLSDetectionAndCipherExclusion(t *testing.T) {
	for _, cipher := range []uint16{0x1301, 0x1302, 0x1303, 0x1304, 0x1305, 0x9999} {
		state := &MxS0kZ{BUwqtS: 8, TvZPBqb: -1}
		YByzFSgf(serverHello(cipher), state)
		want := cipher >= 0x1301 && cipher <= 0x1304
		if state.YdJioLGvuLW != want || !state.UFPiQl0bbHk || !state.VRlXaTaT || state.OqJe390R06y != cipher || state.BUwqtS != 0 || state.TvZPBqb != 0 {
			t.Fatalf("cipher%x %+v", cipher, state)
		}
	}
	state := &MxS0kZ{BUwqtS: 8, TvZPBqb: -1}
	YByzFSgf(nil, state)
	if state.BUwqtS != 8 {
		t.Fatal("empty decremented")
	}
	YByzFSgf([]byte{0x16, 3, 1, 0, 5, 1}, state)
	if !state.VRlXaTaT || state.UFPiQl0bbHk || state.BUwqtS != 7 {
		t.Fatal(state)
	}
}
func TestVisionTLSSplitAndRecordBounds(t *testing.T) {
	data := serverHello(0x1301)
	for i := 70; i < 76; i++ {
		data[i] = 0
	}
	copy(data[90:], []byte{0, 0x2b, 0, 2, 3, 4})
	state := &MxS0kZ{BUwqtS: 8, TvZPBqb: -1}
	YByzFSgf(data[:80], state)
	if state.TvZPBqb != 20 || state.BUwqtS != 7 || state.YdJioLGvuLW {
		t.Fatal(state)
	}
	YByzFSgf(data[80:], state)
	if !state.YdJioLGvuLW || state.BUwqtS != 0 || state.TvZPBqb != 0 {
		t.Fatal(state)
	}
	// pattern跨两次调用边界不会被额外缓存匹配。
	state = &MxS0kZ{BUwqtS: 8, TvZPBqb: 6, OqJe390R06y: 0x1301}
	YByzFSgf([]byte{0, 0x2b, 0}, state)
	YByzFSgf([]byte{2, 3, 4}, state)
	if state.YdJioLGvuLW || state.BUwqtS != 0 {
		t.Fatal(state)
	}
	// 原cipher读取按cap允许扩展切片；这里len79、cap100，session offset40。
	data = bytes.Clone(serverHello(0))
	data[43] = 40
	binary.BigEndian.PutUint16(data[84:], 0x1302)
	state = &MxS0kZ{BUwqtS: 8, TvZPBqb: -1}
	YByzFSgf(data[:79], state)
	if state.OqJe390R06y != 0x1302 {
		t.Fatal(state)
	}
}
