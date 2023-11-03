package tofu

import (
	"reflect"
	"strconv"
	"strings"
)

type Result struct {
	raw reflect.Value
}

func (v Result) Raw() interface{} {
	if v.raw.IsValid() && v.raw.CanInterface() {
		return v.raw.Interface()
	}
	return nil
}

func (v Result) String() string {
	if v.raw.IsValid() && v.raw.Kind() == reflect.String {
		return v.raw.String()
	}

	if v.raw.IsValid() && (v.raw.Kind() == reflect.Interface) {
		return CastToString(v.raw.Interface())
	}

	return ""
}

func (v Result) Int() int {
	if v.raw.IsValid() && (v.raw.Kind() == reflect.Int || v.raw.Kind() == reflect.Float64 || v.raw.Kind() == reflect.Float32) {
		return int(v.raw.Int())
	}

	if v.raw.IsValid() && (v.raw.Kind() == reflect.Interface) {
		return CastToInt(v.raw.Interface())
	}

	return -1
}

func (v Result) Float() float64 {
	if v.raw.IsValid() && (v.raw.Kind() == reflect.Float64 || v.raw.Kind() == reflect.Float32 || v.raw.Kind() == reflect.Int) {
		return float64(v.raw.Float())
	}

	if v.raw.IsValid() && (v.raw.Kind() == reflect.Interface) {
		return CastToFloat(v.raw.Interface())
	}

	return -1
}

func (v Result) Bool() bool {
	if v.raw.IsValid() && v.raw.Kind() == reflect.Bool {
		return v.raw.Bool()
	}

	if v.raw.IsValid() && (v.raw.Kind() == reflect.Interface) {
		return CastToBool(v.raw.Interface())
	}

	return false
}

func GetN(src interface{}, path string, def ...any) Result {
	val := get(reflect.ValueOf(src), strings.Split(path, "."), def...)
	if val.IsValid() && val.CanInterface() {
		return Result{
			raw: val,
		}
	}

	return Result{
		raw: reflect.ValueOf(nil),
	}
}

func Get(src interface{}, path string, def ...any) interface{} {
	val := get(reflect.ValueOf(src), strings.Split(path, "."), def...)
	if val.IsValid() && val.CanInterface() {
		return val.Interface()
	}

	return nil
}

func CastToInt(i any, def ...int) int {
	if v, ok := i.(int); ok {
		return v
	}
	if v, ok := i.(float32); ok {
		return int(v)
	}
	if v, ok := i.(float64); ok {
		return int(v)
	}
	if v, ok := i.(string); ok {
		res, err := strconv.Atoi(v)
		if err == nil {
			return res
		}
	}
	if v, ok := i.(bool); ok {
		if v {
			return 1
		} else {
			return 0
		}
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

func CastToFloat(i any, def ...float64) float64 {
	if v, ok := i.(float64); ok {
		return v
	}
	if v, ok := i.(int); ok {
		return float64(v)
	}
	if v, ok := i.(float32); ok {
		return float64(v)
	}
	if v, ok := i.(string); ok {
		res, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return res
		}
	}
	if v, ok := i.(bool); ok {
		if v {
			return 1
		} else {
			return 0
		}
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

func intToBool[T int | float64 | float32](v T) bool {
	if v <= 0 {
		return false
	}
	if v >= 1 {
		return true
	}

	return false
}

func CastToBool(i any, def ...bool) bool {
	if v, ok := i.(bool); ok {
		return v
	}
	if v, ok := i.(float64); ok {
		return intToBool(v)
	}
	if v, ok := i.(int); ok {
		return intToBool(v)
	}
	if v, ok := i.(float32); ok {
		return intToBool(v)
	}
	if v, ok := i.(string); ok {
		res, err := strconv.ParseBool(v)
		if err == nil {
			return res
		}
	}

	if len(def) > 0 {
		return def[0]
	}

	return false
}

func CastToString(i any, def ...string) string {
	if v, ok := i.(bool); ok {
		if v {
			return "true"
		} else {
			return "false"
		}
	}
	if v, ok := i.(int); ok {
		return strconv.Itoa(v)
	}
	if v, ok := i.(float64); ok {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}
	if v, ok := i.(float32); ok {
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	}
	if v, ok := i.(string); ok {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

func get(src reflect.Value, path []string, def ...any) reflect.Value {
	if len(path) > 0 && src.IsValid() {
		var next reflect.Value

		if src.Kind() == reflect.Array || src.Kind() == reflect.Slice {
			if idx, err := strconv.Atoi(path[0]); idx >= 0 && idx <= src.Len() && err == nil {
				next = src.Index(idx)
				if len(path) == 1 {
					return next
				}

				return get(next, path[1:], def...)
			}
		}

		if src.Kind() == reflect.Map || src.Kind() == reflect.Struct {
			if src.Kind() == reflect.Struct {
				next = src.FieldByName(path[0])
			} else {
				next = src.MapIndex(reflect.ValueOf(path[0]))
			}

			if next.IsValid() {
				if len(path) == 1 {
					return next
				}

				return get(next.Elem(), path[1:], def...)
			}
		}
	}

	if len(def) > 0 {
		return reflect.ValueOf(def[0])
	}

	return reflect.ValueOf(nil)
}
