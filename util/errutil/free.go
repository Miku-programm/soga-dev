package errutil

import (
	"net"
	"reflect"
)

// 原RTTI接口xBpAmFj.JZgM1sMF和Gl5pJQIR；原动态类型表0x503e300
// 按Free、GetRawConn顺序选择分支。接口方法签名由原RTTI直接解码。
type JZgM1sMF interface{ Free() }
type Gl5pJQIR interface{ GetRawConn() net.Conn }

// U9tk4Xosh9 @0xb26020；原代码直接IsNil，非可空的值类型会panic。
func U9tk4Xosh9(value JZgM1sMF) {
	if value != nil && !reflect.ValueOf(value).IsNil() {
		value.Free()
	}
}

// Li3c4CJ1_D @0xb26980。先调用最外层Free；只有没有Free的包装层
// 才继续GetRawConn。原函数没有主动Close或循环引用检测。
func Li3c4CJ1_D(conn net.Conn) {
	for conn != nil {
		switch value := conn.(type) {
		case JZgM1sMF:
			U9tk4Xosh9(value)
			return
		case Gl5pJQIR:
			conn = value.GetRawConn()
		default:
			return
		}
	}
}
