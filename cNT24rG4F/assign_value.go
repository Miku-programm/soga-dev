// 本文件由原样本的 IDA 指令和静态解码常量恢复，非作者文件原文。
package cNT24rG4F

import (
	"reflect"
	"strconv"
	"time"

	"recovered.local/soga/vKl1IHr3"
)

// ieqMch_a 对应 0x13236c0，将配置文本写入反射字段。
// 保留原函数对不支持类型的 panic、窄整数截断和数字解析失败置零行为。
func ieqMch_a(field reflect.Value, text string) {
	if field.Type() == reflect.TypeOf(time.Duration(0)) {
		panic("conf tag does not support time.Duration")
	}
	switch field.Kind() {
	case reflect.Bool:
		field.SetBool(text == "true")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			value = 0
		}
		field.SetInt(value)
	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			value = 0
		}
		field.SetFloat(value)
	case reflect.String:
		field.SetString(text)
	case reflect.Slice:
		if field.Type().Elem().Kind() != reflect.String {
			panic("conf tag only supports []string slice, got " + field.Type().String())
		}
		field.Set(reflect.ValueOf(vKl1IHr3.DAaBs7_C(text, ",", -1)))
	default:
		panic("conf tag does not support " + field.Type().String())
	}
}
