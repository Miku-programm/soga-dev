package l_5zgG8VLP

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/bQxeORT"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/reality"
)

// 原 RTTI 0x3584a00 的命名 string 类型。
type TIJsDKdn1j string

// V7ZRKYAK @0x1d027a0。输入已经消费了UUID和version；认证仍须由上游
// 完成，原入口强断言实际用户类型，保留失败条件及原来的检查顺序。
func V7ZRKYAK(conn *J1P4KNvfI_J.Mdn09F6aMrpa, allowedFlow string, version byte) (*J3LOqQMzo, error) {
	header, err := vF01yD4y(conn, version)
	if err != nil {
		return nil, err
	}
	flow := TIJsDKdn1j(header.qssMeIUTR.AfJaYuHm3H)
	var input *bytes.Reader
	var rawInput *bytes.Buffer
	switch flow {
	case "":
		if allowedFlow == "xtls-rprx-vision" && header.aj8eMgsHQRY == 1 {
			return nil, errors.New("can not use empty flow")
		}
	case "xtls-rprx-vision":
		if allowedFlow != string(flow) {
			return nil, fmt.Errorf("can not able to use flow: %v", flow)
		}
		var typ reflect.Type
		var pointer unsafe.Pointer
		switch raw := conn.DrpXao.(type) {
		case *bQxeORT.RfBfjVN4XE:
			typ, pointer = reflect.TypeOf(raw).Elem(), unsafe.Pointer(raw)
		case *tls.Conn:
			if raw.ConnectionState().Version != tls.VersionTLS13 {
				return nil, fmt.Errorf("failed to use %v, found outer tls version %v", flow, raw.ConnectionState().Version)
			}
			typ, pointer = reflect.TypeOf(raw).Elem(), unsafe.Pointer(raw)
		case *reality.Conn:
			typ, pointer = reflect.TypeOf(raw).Elem(), unsafe.Pointer(raw)
		default:
			return nil, errors.New("XTLS only supports TLS and REALITY directly for now")
		}
		// 原指令确实忽略FieldByName的ok值，且随后用Offset直接定位。
		// RfBfjVN4XE的混淆RTTI字段名不匹配时Offset为0；这里不修复原行为。
		inputField, _ := typ.FieldByName("input")
		rawField, _ := typ.FieldByName("rawInput")
		input = (*bytes.Reader)(unsafe.Add(pointer, inputField.Offset))
		rawInput = (*bytes.Buffer)(unsafe.Add(pointer, rawField.Offset))
	default:
		return nil, fmt.Errorf("unknown vless flow %v", flow)
	}
	user := conn.GetUser().(*nzATvajt.A6CvRbtca)
	id, exists := user.GetUuid()
	if !exists {
		return nil, fmt.Errorf("failed to get user %v uuid", user.GetId())
	}
	state := &MxS0kZ{
		RW3eHGaIH0ya: id[:], BUwqtS: 8, TvZPBqb: -1,
		EO4y9oyDkzq: true, Z43zq7: -1, ZOEKPa7JPx5: -1, PIsBMe: -1, ZkTHd3: true,
	}
	reader, err := iP6eRhiZ(conn, header)
	if err != nil {
		return nil, err
	}
	if flow == "xtls-rprx-vision" {
		reader = &LZ7AGw7{Reader: reader, fB4afneh7k: conn, woJdOhbWO: state,
			sztBlp: input, k9rmTO: rawInput, oqHPkr: D_zCag1qTV.S4OQXWIZ(8192)}
	}
	writer, err := yPiHcQrp4Vs(conn, header)
	if err != nil {
		return nil, err
	}
	if header.aj8eMgsHQRY != 2 && flow == "xtls-rprx-vision" {
		prefix := make([]byte, len(state.RW3eHGaIH0ya))
		copy(prefix, state.RW3eHGaIH0ya)
		writer = &Qx58rV{Writer: writer, uYYP2p: conn, taFXeH: state, oRVa2PobW_S: prefix}
	}
	result := &J3LOqQMzo{DrpXao: conn, g9LlwZz4Rdy: reader, oWbWTrr2MUiU: writer,
		d7otQpiN: conn.GetUser(), hC8tTwO: header}
	if flow == "xtls-rprx-vision" {
		result.eJeaO7LW = state
	}
	return result, nil
}
