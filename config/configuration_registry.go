// 本文件依据 IDA 函数体恢复；全局变量名由分析者按原地址指定。
package config

import (
	"reflect"
	"strings"
)

// 0x1319080 初始化默认值映射；另外两个列表保持原程序的零值。
var defaultsAt5095AA0 = make(map[string]string)
var registeredAt5099320 []any
var afterLoadAt5099300 []func()

// uku6j8A 对应 0x1323380。conf 标签存在就有效，包括空名称；不去除空格。
func uku6j8A(field reflect.StructField) (name, defaultText string, present bool) {
	name, present = field.Tag.Lookup("conf")
	if !present {
		return "", "", false
	}
	name = strings.ToLower(name)
	defaultText, _ = field.Tag.Lookup("default")
	return name, defaultText, true
}

// _NA1kft 对应 0x1322f60。检查默认值的可转换性，再登记原结构指针。
// 转换发生在临时字段上；此阶段不把默认值写进传入的结构实例。
func _NA1kft(configuration any) {
	typ := reflect.TypeOf(configuration).Elem()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		name, defaultText, present := uku6j8A(field)
		if !present {
			continue
		}
		ieqMch_a(reflect.New(field.Type).Elem(), defaultText)
		defaultsAt5095AA0[name] = defaultText
	}
	registeredAt5099320 = append(registeredAt5099320, configuration)
}

// bFnBl5Ak9 对应 0x1323420。直接使用最终映射里的文本，缺键读为空字符串。
func bFnBl5Ak9(configuration any) {
	value := reflect.ValueOf(configuration).Elem()
	for i := 0; i < value.NumField(); i++ {
		name, _, present := uku6j8A(value.Type().Field(i))
		if present {
			ieqMch_a(value.Field(i), configurationAt5095AA8[name])
		}
	}
}
