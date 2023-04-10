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

	return ""
}

func (v Result) Int() int64 {
	if v.raw.IsValid() && v.raw.Kind() == reflect.Int {
		return v.raw.Int()
	}

	return -1
}

func (v Result) Float() float64 {
	if v.raw.IsValid() && (v.raw.Kind() == reflect.Float64 || v.raw.Kind() == reflect.Float32) {
		return v.raw.Float()
	}

	return -1
}

func (v Result) Bool() bool {
	if v.raw.IsValid() && v.raw.Kind() == reflect.Bool {
		return v.raw.Bool()
	}

	return false
}

func Get(src interface{}, path string, def ...any) interface{} {
	val := get(reflect.ValueOf(src), strings.Split(path, "."), def...)
	if val.IsValid() && val.CanInterface() {
		return val.Interface()
	}

	return nil
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
