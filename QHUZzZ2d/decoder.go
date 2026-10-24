// 已对照原样本 0x19358a0—0x1939f00 的十二个业务入口、RTTI 与字符串。
// 来源候选：https://github.com/hUYisnn/clash/blob/master/common/structure/structure.go
// 原协议配置库的此部分与该历史 Clash 依赖一致；这里不是替换转发内核。
// 原作者变量名不可恢复。版权及 GPLv3 许可见 licenses.go。
package QHUZzZ2d

// references: https://github.com/mitchellh/mapstructure

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// MBi603C is the configuration that is used to create a new decoder
type MBi603C struct {
	YONNn23   string
	RlQxpRzya bool
}

// LsMT6BB is the core of structure
type LsMT6BB struct {
	d01d_V594LQ *MBi603C
}

// K6YuYvsMD5Y return a LsMT6BB by MBi603C
func K6YuYvsMD5Y(d01d_V594LQ MBi603C) *LsMT6BB {
	if d01d_V594LQ.YONNn23 == "" {
		d01d_V594LQ.YONNn23 = "structure"
	}
	return &LsMT6BB{d01d_V594LQ: &d01d_V594LQ}
}

// Decode transform a map[string]any to a struct
func (d *LsMT6BB) Decode(src map[string]any, dst any) error {
	if reflect.TypeOf(dst).Kind() != reflect.Ptr {
		return fmt.Errorf("Decode must recive a ptr struct")
	}
	t := reflect.TypeOf(dst).Elem()
	v := reflect.ValueOf(dst).Elem()
	for idx := 0; idx < v.NumField(); idx++ {
		field := t.Field(idx)
		if field.Anonymous {
			if err := d.jUOTMOSzkd8K(field.Name, src, v.Field(idx)); err != nil {
				return err
			}
			continue
		}

		tag := field.Tag.Get(d.d01d_V594LQ.YONNn23)
		key, omitKey, found := strings.Cut(tag, ",")
		omitempty := found && omitKey == "omitempty"

		value, ok := src[key]
		if !ok || value == nil {
			if omitempty {
				continue
			}
			return fmt.Errorf("key '%s' missing", key)
		}

		err := d.kFuaz5EqcMhd(key, value, v.Field(idx))
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *LsMT6BB) kFuaz5EqcMhd(name string, data any, val reflect.Value) error {
	switch val.Kind() {
	case reflect.Int:
		return d.oFXosZkCznOb(name, data, val)
	case reflect.String:
		return d.h2HVgaIJ(name, data, val)
	case reflect.Bool:
		return d.lDEb3C(name, data, val)
	case reflect.Slice:
		return d.myOR7sbC(name, data, val)
	case reflect.Map:
		return d.ffScJWvztxAA(name, data, val)
	case reflect.Interface:
		return d.fLakdR88(name, data, val)
	case reflect.Struct:
		return d.jUOTMOSzkd8K(name, data, val)
	default:
		return fmt.Errorf("type %s not support", val.Kind().String())
	}
}

func (d *LsMT6BB) oFXosZkCznOb(name string, data any, val reflect.Value) (err error) {
	dataVal := reflect.ValueOf(data)
	kind := dataVal.Kind()
	switch {
	case kind == reflect.Int:
		val.SetInt(dataVal.Int())
	case kind == reflect.Float64 && d.d01d_V594LQ.RlQxpRzya:
		val.SetInt(int64(dataVal.Float()))
	case kind == reflect.String && d.d01d_V594LQ.RlQxpRzya:
		var i int64
		i, err = strconv.ParseInt(dataVal.String(), 0, val.Type().Bits())
		if err == nil {
			val.SetInt(i)
		} else {
			err = fmt.Errorf("cannot parse '%s' as int: %s", name, err)
		}
	default:
		err = fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type(),
		)
	}
	return err
}

func (d *LsMT6BB) h2HVgaIJ(name string, data any, val reflect.Value) (err error) {
	dataVal := reflect.ValueOf(data)
	kind := dataVal.Kind()
	switch {
	case kind == reflect.String:
		val.SetString(dataVal.String())
	case kind == reflect.Int && d.d01d_V594LQ.RlQxpRzya:
		val.SetString(strconv.FormatInt(dataVal.Int(), 10))
	default:
		err = fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type(),
		)
	}
	return err
}

func (d *LsMT6BB) lDEb3C(name string, data any, val reflect.Value) (err error) {
	dataVal := reflect.ValueOf(data)
	kind := dataVal.Kind()
	switch {
	case kind == reflect.Bool:
		val.SetBool(dataVal.Bool())
	case kind == reflect.Int && d.d01d_V594LQ.RlQxpRzya:
		val.SetBool(dataVal.Int() != 0)
	default:
		err = fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type(),
		)
	}
	return err
}

func (d *LsMT6BB) myOR7sbC(name string, data any, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	valType := val.Type()
	valElemType := valType.Elem()

	if dataVal.Kind() != reflect.Slice {
		return fmt.Errorf("'%s' is not a slice", name)
	}

	valSlice := val
	for i := 0; i < dataVal.Len(); i++ {
		currentData := dataVal.Index(i).Interface()
		for valSlice.Len() <= i {
			valSlice = reflect.Append(valSlice, reflect.Zero(valElemType))
		}
		fieldName := fmt.Sprintf("%s[%d]", name, i)
		if currentData == nil {
			// in weakly type mode, null will convert to zero value
			if d.d01d_V594LQ.RlQxpRzya {
				continue
			}
			// in non-weakly type mode, null will convert to nil if element's zero value is nil, otherwise return an error
			if elemKind := valElemType.Kind(); elemKind == reflect.Map || elemKind == reflect.Slice {
				continue
			}
			return fmt.Errorf("'%s' can not be null", fieldName)
		}
		currentField := valSlice.Index(i)
		if err := d.kFuaz5EqcMhd(fieldName, currentData, currentField); err != nil {
			return err
		}
	}

	val.Set(valSlice)
	return nil
}

func (d *LsMT6BB) ffScJWvztxAA(name string, data any, val reflect.Value) error {
	valType := val.Type()
	valKeyType := valType.Key()
	valElemType := valType.Elem()

	valMap := val

	if valMap.IsNil() {
		mapType := reflect.MapOf(valKeyType, valElemType)
		valMap = reflect.MakeMap(mapType)
	}

	dataVal := reflect.Indirect(reflect.ValueOf(data))
	if dataVal.Kind() != reflect.Map {
		return fmt.Errorf("'%s' expected a map, got '%s'", name, dataVal.Kind())
	}

	return d.s_nDeb70KRUg(name, dataVal, val, valMap)
}

func (d *LsMT6BB) s_nDeb70KRUg(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
	valType := val.Type()
	valKeyType := valType.Key()
	valElemType := valType.Elem()

	errors := make([]string, 0)

	if dataVal.Len() == 0 {
		if dataVal.IsNil() {
			if !val.IsNil() {
				val.Set(dataVal)
			}
		} else {
			val.Set(valMap)
		}

		return nil
	}

	for _, k := range dataVal.MapKeys() {
		fieldName := fmt.Sprintf("%s[%s]", name, k)

		currentKey := reflect.Indirect(reflect.New(valKeyType))
		if err := d.kFuaz5EqcMhd(fieldName, k.Interface(), currentKey); err != nil {
			errors = append(errors, err.Error())
			continue
		}

		v := dataVal.MapIndex(k).Interface()
		if v == nil {
			errors = append(errors, fmt.Sprintf("filed %s invalid", fieldName))
			continue
		}

		currentVal := reflect.Indirect(reflect.New(valElemType))
		if err := d.kFuaz5EqcMhd(fieldName, v, currentVal); err != nil {
			errors = append(errors, err.Error())
			continue
		}

		valMap.SetMapIndex(currentKey, currentVal)
	}

	val.Set(valMap)

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ","), []any{}...)
	}

	return nil
}

func (d *LsMT6BB) jUOTMOSzkd8K(name string, data any, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))

	// If the type of the value to write to and the data match directly,
	// then we just set it directly instead of recursing into the structure.
	if dataVal.Type() == val.Type() {
		val.Set(dataVal)
		return nil
	}

	dataValKind := dataVal.Kind()
	switch dataValKind {
	case reflect.Map:
		return d.kqLYSRF8X(name, dataVal, val)
	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", name, dataVal.Kind())
	}
}

func (d *LsMT6BB) kqLYSRF8X(name string, dataVal, val reflect.Value) error {
	dataValType := dataVal.Type()
	if kind := dataValType.Key().Kind(); kind != reflect.String && kind != reflect.Interface {
		return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys",
			name, dataValType.Key().Kind())
	}

	dataValKeys := make(map[reflect.Value]struct{})
	dataValKeysUnused := make(map[any]struct{})
	for _, dataValKey := range dataVal.MapKeys() {
		dataValKeys[dataValKey] = struct{}{}
		dataValKeysUnused[dataValKey.Interface()] = struct{}{}
	}

	errors := make([]string, 0)

	// This slice will keep track of all the structs we'll be decoding.
	// There can be more than one struct if there are embedded structs
	// that are squashed.
	structs := make([]reflect.Value, 1, 5)
	structs[0] = val

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type field struct {
		field reflect.StructField
		val   reflect.Value
	}
	fields := []field{}
	for len(structs) > 0 {
		structVal := structs[0]
		structs = structs[1:]

		structType := structVal.Type()

		for i := 0; i < structType.NumField(); i++ {
			fieldType := structType.Field(i)
			fieldKind := fieldType.Type.Kind()

			// If "squash" is specified in the tag, we squash the field down.
			squash := false
			tagParts := strings.Split(fieldType.Tag.Get(d.d01d_V594LQ.YONNn23), ",")
			for _, tag := range tagParts[1:] {
				if tag == "squash" {
					squash = true
					break
				}
			}

			if squash {
				if fieldKind != reflect.Struct {
					errors = append(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", fieldType.Name, fieldKind).Error())
				} else {
					structs = append(structs, structVal.FieldByName(fieldType.Name))
				}
				continue
			}

			// Normal struct field, store it away
			fields = append(fields, field{fieldType, structVal.Field(i)})
		}
	}

	// for fieldType, field := range fields {
	for _, f := range fields {
		field, fieldValue := f.field, f.val
		fieldName := field.Name

		tagValue := field.Tag.Get(d.d01d_V594LQ.YONNn23)
		tagValue = strings.SplitN(tagValue, ",", 2)[0]
		if tagValue != "" {
			fieldName = tagValue
		}

		rawMapKey := reflect.ValueOf(fieldName)
		rawMapVal := dataVal.MapIndex(rawMapKey)
		if !rawMapVal.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for dataValKey := range dataValKeys {
				mK, ok := dataValKey.Interface().(string)
				if !ok {
					// Not a string key
					continue
				}

				if strings.EqualFold(mK, fieldName) {
					rawMapKey = dataValKey
					rawMapVal = dataVal.MapIndex(dataValKey)
					break
				}
			}

			if !rawMapVal.IsValid() {
				// There was no matching key in the map for the value in
				// the struct. Just ignore.
				continue
			}
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(dataValKeysUnused, rawMapKey.Interface())

		if !fieldValue.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !fieldValue.CanSet() {
			continue
		}

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if name != "" {
			fieldName = fmt.Sprintf("%s.%s", name, fieldName)
		}

		if err := d.kFuaz5EqcMhd(fieldName, rawMapVal.Interface(), fieldValue); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ","), []any{}...)
	}

	return nil
}

func (d *LsMT6BB) fLakdR88(name string, data any, val reflect.Value) (err error) {
	dataVal := reflect.ValueOf(data)
	val.Set(dataVal)
	return nil
}
