package assembler

import (
	"net/url"
	"reflect"
	"strconv"
)

const (
	blank              = ""
	namespaceSeparator = "."
)

// var bakedIn = map[reflect.Type]Converter{
// 	reflect.TypeOf(int8(0)): int8Converter,
// }

// // KeyValues is the conversions input of key and values
// type KeyValues struct {
// 	Key    []byte
// 	Values []string
// }

// // Converter is the type used for converting types of values
// type Converter func(kv []KeyValues) (reflect.Value, error)

// Decoder is the assembler decode instance
type Decoder struct {
	tagName string
}

// NewDecoder creates a new decoder instance
func NewDecoder() *Decoder {
	return &Decoder{
		tagName: "assembler",
	}
}

// Decode decodes the given values and set the cooresponding struct values on v
func (d *Decoder) Decode(v interface{}, values url.Values) (err error) {

	// val := reflect.Indirect(reflect.ValueOf(v))

	// fmt.Println(val)
	errs := map[string]error{}
	d.traverseStruct(reflect.ValueOf(v), values, "", errs)

	return
}

func (d *Decoder) traverseStruct(v reflect.Value, values url.Values, namespace string, errs map[string]error) {

	if v.Kind() == reflect.Ptr {
		if v.IsNil() && v.CanSet() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}
	// if v.Kind() == reflect.Ptr && !v.IsNil() {
	// 	v = v.Elem()
	// }

	if v.Kind() != reflect.Struct && v.Kind() != reflect.Interface {
		panic("value passed for validation is not a struct")
	}
	// if v.Kind() == reflect.Ptr {
	// 	if v.IsNil() {
	// 		// v.Elem().Set(reflect.New(v.Type().Elem()))
	// 		// fmt.Println(v)
	// 		// fmt.Println(v.Elem())
	// 		// fmt.Println(v.Type())
	// 		// fmt.Println(v.Type().Elem())
	// 		// fmt.Println(reflect.New(v.Type()))
	// 		// fmt.Println(reflect.New(v.Type()).Elem())
	// 	}

	// 	// v = v.Elem()
	// }

	// if v.Kind() != reflect.Struct && v.Kind() != reflect.Interface {
	// 	panic("value passed for validation is not a struct")
	// }

	// fmt.Println(v)
	typ := v.Type()
	numFields := v.NumField()
	var fld reflect.StructField
	var key string
	var nn string // new namespace

	for i := 0; i < numFields; i++ {

		fld = typ.Field(i)

		if fld.PkgPath != blank && !fld.Anonymous {
			continue
		}

		key = fld.Name

		if namespace == blank {
			nn = key
		} else {
			nn = namespace + namespaceSeparator + key
		}

		d.setFieldByType(v.Field(i), values, nn, 0, errs)
		// TODO check if pointer before switch

		// switch fld.Kind() {

		// }

		// 	v.traverseField(topStruct, currentStruct, current.Field(i), errPrefix, nsPrefix, errs, true, fld.Tag.Get(v.tagName), fld.Name, customName, partial, exclude, includeExclude, nil)
	}
}

func (d *Decoder) setFieldByType(current reflect.Value, values url.Values, namespace string, idx int, errs map[string]error) (set bool) {

	// var s string
	var arr []string
	var ok bool
	var err error
	// var val reflect.Value

	v, kind := d.ExtractType(current)

	switch kind {
	case reflect.Interface, reflect.Invalid:
		return
	case reflect.Ptr:
		// fmt.Println("I'm a Ptr", namespace, v.Type(), v.Type().Elem())
		// TODO check if field
		newVal := reflect.New(v.Type().Elem())
		if set = d.setFieldByType(newVal.Elem(), values, namespace, idx, errs); set {
			v.Set(newVal)
		}
		// TODO if newVal set then do this, otherwise no fileds here.
		// v.Set(newVal)
	case reflect.Int8:
		var i64 int64
		if arr, ok = values[namespace]; len(arr) == 0 {
			if !ok {
				return
			}

			set = true
			v.Set(reflect.ValueOf(int8(0)))
			return
		}

		if i64, err = strconv.ParseInt(arr[idx], 10, 8); err != nil {
			errs[namespace] = err
			return
		}

		v.Set(reflect.ValueOf(int8(i64)))
		set = true

	case reflect.Slice, reflect.Array:

		if arr, _ = values[namespace]; len(arr) == 0 {
			return
		}

		varr := reflect.MakeSlice(v.Type(), 0, len(arr))

		for i := 0; i < len(arr); i++ {
			newVal := reflect.New(v.Type().Elem()).Elem()

			if d.setFieldByType(newVal, values, namespace, i, errs) {
				varr = reflect.Append(varr, newVal)
			}
		}

		if varr.Len() == 0 {
			return
		}

		set = true
		v.Set(varr)
	}

	return
}

// func int8Converter(kv []KeyValues) (value reflect.Value, err error) {
// 	var v int64

// 	v, err = strconv.ParseInt(kv[0].Values[0], 10, 8)

// 	value = reflect.ValueOf(int8(v))
// 	return
// }
