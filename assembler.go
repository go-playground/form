package assembler

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	blank              = ""
	namespaceSeparator = "."
)

var (
	timeType = reflect.TypeOf(time.Time{})
)

// DecodeErrors is a map of errors encountered during form decoding
type DecodeErrors map[string]error

func (d DecodeErrors) Error() string {
	buff := bytes.NewBufferString(blank)

	for _, err := range d {
		buff.WriteString(err.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

type key struct {
	value       string
	searchValue string
}

type index struct {
	value       int
	searchValue string
}

type recursiveData struct {
	isSlice  bool
	isMap    bool
	sliceLen int
	keys     []*key
	indicies []*index
}

type dataMap map[string]*recursiveData

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
	errs := make(DecodeErrors)
	dm := d.parseMapData(values)

	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() && val.CanSet() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}
	// if v.Kind() == reflect.Ptr && !v.IsNil() {
	// 	v = v.Elem()
	// }

	if val.Kind() != reflect.Struct && val.Kind() != reflect.Interface {
		panic("value passed for validation is not a struct")
	}

	d.traverseStruct(val, values, "", dm, errs)

	if len(errs) == 0 {
		return nil
	}

	err = errs

	return
}

// find a way to parse this once and only when needed?
func (d *Decoder) parseMapData(values url.Values) (dm dataMap) {

	dm = make(dataMap)
	var idx int
	var idx2 int
	var cum int

	for k := range values {

		idx, idx2, cum = 0, 0, 0

		for {
			if idx = strings.Index(k[cum:], "["); idx == -1 {
				break
			}

			if idx2 = strings.Index(k[cum:], "]"); idx2 == -1 {
				log.Panicf("Invalid formatting for key '%s' missing bracket", k)
			}

			var rd *recursiveData
			var ok bool

			if rd, ok = dm[k[:idx+cum]]; !ok {
				rd = new(recursiveData)
				dm[k[:idx+cum]] = rd
			}

			// fmt.Println(k, cum+idx+1, cum+idx2)
			j, err := strconv.Atoi(k[cum+idx+1 : cum+idx2])
			if err != nil {
				// is map + key
				rd.isMap = true
				k := &key{
					value:       k[cum+idx+1 : cum+idx2],
					searchValue: k[cum+idx : cum+idx2+1],
				}
				rd.keys = append(rd.keys, k)
			} else {
				// is slice + indicies
				rd.isSlice = true

				if j > rd.sliceLen {
					rd.sliceLen = j
				}

				ind := &index{
					value:       j,
					searchValue: k[cum+idx : cum+idx2+1],
				}
				rd.indicies = append(rd.indicies, ind)
			}

			cum += idx2 + 1
		}
	}

	// fmt.Println(len(dm))
	// for k, v := range dm {
	// 	fmt.Println(k, v)
	// 	if v.isSlice {
	// 		fmt.Println(len(v.indicies), v.sliceLen)
	// 		for _, sv := range v.indicies {
	// 			fmt.Println(sv.value, sv.searchValue)
	// 		}
	// 	}
	// }

	return
}

func (d *Decoder) traverseStruct(v reflect.Value, values url.Values, namespace string, dm dataMap, errs DecodeErrors) (set bool) {

	// if v.Kind() == reflect.Ptr {
	// 	// if v.IsNil() && v.CanSet() {
	// 	// 	v.Set(reflect.New(v.Type().Elem()))
	// 	// }
	// 	v = v.Elem()
	// }
	// if v.Kind() == reflect.Ptr && !v.IsNil() {
	// 	v = v.Elem()
	// }

	// if v.Kind() != reflect.Struct && v.Kind() != reflect.Interface {
	// 	panic("value passed for validation is not a struct")
	// }
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

		if d.setFieldByType(v.Field(i), values, nn, 0, dm, errs) {
			set = true
		}
		// TODO check if pointer before switch

		// switch fld.Kind() {

		// }

		// 	v.traverseField(topStruct, currentStruct, current.Field(i), errPrefix, nsPrefix, errs, true, fld.Tag.Get(v.tagName), fld.Name, customName, partial, exclude, includeExclude, nil)
	}

	return
}

func (d *Decoder) setFieldByType(current reflect.Value, values url.Values, namespace string, idx int, dm dataMap, errs DecodeErrors) (set bool) {

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

		if arr = values[namespace]; len(arr) == 0 {
			return
		}

		// if v.IsNil() {
		newVal := reflect.New(v.Type().Elem())
		if set = d.setFieldByType(newVal.Elem(), values, namespace, idx, dm, errs); set {
			v.Set(newVal)
		}
		// }

		// set = d.setFieldByType(v.Elem(), values, namespace, idx, dm, errs)

	case reflect.String:
		if arr = values[namespace]; len(arr) == 0 {
			return
		}

		v.SetString(arr[0])
		set = true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if arr, ok = values[namespace]; len(arr) == 0 {
			if !ok {
				return
			}

			v.SetInt(0)
			// v.Set(reflect.ValueOf(int8(0)))
			set = true
			return
		}

		var u64 uint64

		if u64, err = strconv.ParseUint(arr[idx], 10, 64); err != nil || v.OverflowUint(u64) {
			errs[namespace] = fmt.Errorf("Invalid Unsigned Integer Value '%s', Type '%v'", arr[idx], v.Type())
			return
		}

		v.SetUint(u64)
		// v.Set(reflect.ValueOf(int8(i64)))
		set = true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if arr, ok = values[namespace]; len(arr) == 0 {
			if !ok {
				return
			}

			v.SetInt(0)
			// v.Set(reflect.ValueOf(int8(0)))
			set = true
			return
		}

		var i64 int64

		if i64, err = strconv.ParseInt(arr[idx], 10, 64); err != nil || v.OverflowInt(i64) {
			errs[namespace] = fmt.Errorf("Invalid Integer Value '%s', Type '%v'", arr[idx], v.Type())
			return
		}

		v.SetInt(i64)
		// v.Set(reflect.ValueOf(int8(i64)))
		set = true

	case reflect.Float32, reflect.Float64:
		if arr = values[namespace]; len(arr) == 0 {
			return
		}

		var f float64

		if f, err = strconv.ParseFloat(arr[0], 64); err != nil || v.OverflowFloat(f) {
			errs[namespace] = fmt.Errorf("Invalid Float Value '%s', Type '%v'", arr[0], v.Type())
			return
		}

		v.SetFloat(f)
		set = true

	case reflect.Bool:

		if arr = values[namespace]; len(arr) == 0 {
			return
		}

		var b bool

		if b, err = strconv.ParseBool(arr[0]); err != nil {
			errs[namespace] = fmt.Errorf("Invalid Boolean Value '%s', Type '%v'", arr[idx], v.Type())
			return
		}

		v.SetBool(b)
		set = true

	case reflect.Slice, reflect.Array:

		if arr, _ = values[namespace]; len(arr) == 0 {

			// maybe it's an numbered array i.e. Pnone[0].Number
			if rd := dm[namespace]; rd != nil {

				var varr reflect.Value

				sl := rd.sliceLen + 1

				if v.IsNil() {
					varr = reflect.MakeSlice(v.Type(), sl, sl)
				} else if v.Len() < sl {
					varr = reflect.MakeSlice(v.Type(), sl, sl)
					reflect.Copy(varr, v)
				} else {
					varr = v
				}

				for i := 0; i < len(rd.indicies); i++ {
					newVal := reflect.New(v.Type().Elem()).Elem()

					if d.setFieldByType(newVal, values, namespace+rd.indicies[i].searchValue, 0, dm, errs) {
						set = true
						varr.Index(rd.indicies[i].value).Set(newVal)
					}
				}

				if !set {
					return
				}

				v.Set(varr)
			}

			return
		}

		var varr reflect.Value
		var existing bool

		if v.IsNil() {
			varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
		} else if v.Len() < len(arr) {
			varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
			reflect.Copy(varr, v)
		} else {
			existing = true
			varr = v
		}

		for i := 0; i < len(arr); i++ {
			newVal := reflect.New(v.Type().Elem()).Elem()

			if d.setFieldByType(newVal, values, namespace, i, dm, errs) {
				set = true
				varr.Index(i).Set(newVal)
			}
		}

		if !set || existing {
			return
		}

		v.Set(varr)

	case reflect.Map:

		var rd *recursiveData

		// no natural map support so skip directly to dm lookup
		if rd = dm[namespace]; rd == nil {
			return
		}

		var existing bool
		var mp reflect.Value
		typ := v.Type()

		if v.IsNil() {
			mp = reflect.MakeMap(typ)
		} else {
			existing = true
			mp = v
		}

		for i := 0; i < len(rd.keys); i++ {
			newVal := reflect.New(typ.Elem()).Elem()
			kv := reflect.New(typ.Key()).Elem()

			if err := d.getMapKey(rd.keys[i].value, kv); err != nil {
				errs[namespace] = err
				continue
			}

			if d.setFieldByType(newVal, values, namespace+rd.keys[i].searchValue, 0, dm, errs) {
				set = true
				mp.SetMapIndex(kv, newVal)
			}
		}

		if !set || existing {
			return
		}

		v.Set(mp)

	// set here
	case reflect.Struct:

		// TODO: time is struct
		if v.Type() == timeType {

			// look for custom type here

			// Parse time... but how?
			// think maybe this should be left out and must specify a custom type function
			// because everyone may have different requirements...
			return
		}

		set = d.traverseStruct(v, values, namespace, dm, errs)
	default:
		// look for custom type here
		fmt.Println("Currently unknown!")
	}

	return
}

func (d *Decoder) getMapKey(key string, current reflect.Value) (err error) {

	v, kind := d.ExtractType(current)

	switch kind {
	case reflect.Interface, reflect.Invalid:
		return
	case reflect.Ptr:

		// if v.IsNil() {
		newVal := reflect.New(v.Type().Elem())
		if err = d.getMapKey(key, newVal.Elem()); err == nil {
			v.Set(newVal)
		}
		// }

		// set = d.setFieldByType(v.Elem(), values, namespace, idx, dm, errs)

	case reflect.String:
		v.SetString(key)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		u64, e := strconv.ParseUint(key, 10, 64)
		if e != nil || v.OverflowUint(u64) {
			err = fmt.Errorf("Invalid Unsigned Integer Value '%s', Type '%v'", key, v.Type())
			return
		}

		v.SetUint(u64)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		i64, e := strconv.ParseInt(key, 10, 64)
		if e != nil || v.OverflowInt(i64) {
			err = fmt.Errorf("Invalid Integer Value '%s', Type '%v'", key, v.Type())
			return
		}

		v.SetInt(i64)

	case reflect.Float32, reflect.Float64:

		f, e := strconv.ParseFloat(key, 64)
		if e != nil || v.OverflowFloat(f) {
			err = fmt.Errorf("Invalid Float Value '%s', Type '%v'", key, v.Type())
			return
		}

		v.SetFloat(f)

	case reflect.Bool:

		b, e := strconv.ParseBool(key)
		if e != nil {
			err = fmt.Errorf("Invalid Boolean Value '%s', Type '%v'", key, v.Type())
			return
		}

		v.SetBool(b)

	default:
		// look for custom type? or should it be done before this switch, must check out bson.ObjectId because is of typee
		// string but requires a specific method to ensure that it's valid
		err = fmt.Errorf("Unsupported Map Key '%s', Type '%v'", key, v.Type())
	}

	return
}

// maybe just return error instead of passing namespace + errs map...
// func (d *Decoder) getMapKey(key string, current reflect.Value) (v reflect.Value, err error) {

// 	v, kind := d.ExtractType(current)

// 	switch kind {
// 	case reflect.Interface, reflect.Invalid:
// 		return
// 	case reflect.Ptr:

// 		// if v.IsNil() {
// 		v = reflect.New(v.Type().Elem())
// 		myVal, e := d.getMapKey(key, v.Elem()) // TODO: must test this!!!!
// 		err = e
// 		// newVal.Set(myVal)
// 		// }

// 		// set = d.setFieldByType(v.Elem(), values, namespace, idx, dm, errs)

// 	case reflect.String:
// 		return reflect.ValueOf(key), nil

// 	case reflect.Uint:

// 		if u64, e := strconv.ParseUint(key, 10, 64); e != nil || v.OverflowUint(u64) {
// 			err = fmt.Errorf("Invalid Uint Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(uint(u64))
// 		}

// 		return

// 	case reflect.Uint8:

// 		if u64, e := strconv.ParseUint(key, 10, 8); e != nil || v.OverflowUint(u64) {
// 			err = fmt.Errorf("Invalid Uint8 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(uint8(u64))
// 		}

// 		return

// 	case reflect.Uint16:

// 		if u64, e := strconv.ParseUint(key, 10, 16); e != nil || v.OverflowUint(u64) {
// 			err = fmt.Errorf("Invalid Uint16 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(uint16(u64))
// 		}

// 		return

// 	case reflect.Uint32:

// 		if u64, e := strconv.ParseUint(key, 10, 32); e != nil || v.OverflowUint(u64) {
// 			err = fmt.Errorf("Invalid Uint32 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(uint32(u64))
// 		}

// 		return

// 	case reflect.Uint64:

// 		if u64, e := strconv.ParseUint(key, 10, 64); e != nil || v.OverflowUint(u64) {
// 			err = fmt.Errorf("Invalid Uint64 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(uint64(u64))
// 		}

// 		return

// 	case reflect.Int:

// 		if i64, e := strconv.ParseInt(key, 10, 64); e != nil || v.OverflowInt(i64) {
// 			err = fmt.Errorf("Invalid Int Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(int(i64))
// 		}

// 		return

// 	case reflect.Int8:

// 		if i64, e := strconv.ParseInt(key, 10, 8); e != nil || v.OverflowInt(i64) {
// 			err = fmt.Errorf("Invalid Int8 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(int8(i64))
// 		}

// 		return

// 	case reflect.Int16:

// 		if i64, e := strconv.ParseInt(key, 10, 16); e != nil || v.OverflowInt(i64) {
// 			err = fmt.Errorf("Invalid Int16 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(int16(i64))
// 		}

// 		return

// 	case reflect.Int32:

// 		if i64, e := strconv.ParseInt(key, 10, 32); e != nil || v.OverflowInt(i64) {
// 			err = fmt.Errorf("Invalid Int32 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(int32(i64))
// 		}

// 		return

// 	case reflect.Int64:

// 		if i64, e := strconv.ParseInt(key, 10, 64); e != nil || v.OverflowInt(i64) {
// 			err = fmt.Errorf("Invalid Int64 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(int64(i64))
// 		}

// 		return

// 	case reflect.Float32:

// 		if f, e := strconv.ParseFloat(key, 32); e != nil || v.OverflowFloat(f) {
// 			err = fmt.Errorf("Invalid Float32 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(float32(f))
// 		}

// 		return

// 	case reflect.Float64:

// 		if f, e := strconv.ParseFloat(key, 64); e != nil || v.OverflowFloat(f) {
// 			err = fmt.Errorf("Invalid Float64 Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(float64(f))
// 		}

// 		return

// 	case reflect.Bool:

// 		if b, e := strconv.ParseBool(key); e != nil {
// 			err = fmt.Errorf("Invalid Boolean Value '%s', Type '%v'", key, v.Type())
// 		} else {
// 			v = reflect.ValueOf(b)
// 		}

// 		return

// 	case reflect.Slice, reflect.Array:

// 		if arr, _ = values[namespace]; len(arr) == 0 {

// 			// maybe it's an numbered array i.e. Pnone[0].Number
// 			if rd := dm[namespace]; rd != nil {

// 				var varr reflect.Value

// 				sl := rd.sliceLen + 1

// 				if v.IsNil() {
// 					varr = reflect.MakeSlice(v.Type(), sl, sl)
// 				} else if v.Len() < sl {
// 					varr = reflect.MakeSlice(v.Type(), sl, sl)
// 					reflect.Copy(varr, v)
// 				} else {
// 					varr = v
// 				}

// 				for i := 0; i < len(rd.indicies); i++ {
// 					newVal := reflect.New(v.Type().Elem()).Elem()

// 					if d.setFieldByType(newVal, values, namespace+rd.indicies[i].searchValue, 0, dm, errs) {
// 						set = true
// 						varr.Index(rd.indicies[i].value).Set(newVal)
// 					}
// 				}

// 				if !set {
// 					return
// 				}

// 				v.Set(varr)
// 			}

// 			return
// 		}

// 		var varr reflect.Value

// 		if v.IsNil() {
// 			varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
// 		} else if v.Len() < len(arr) {
// 			varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
// 			reflect.Copy(varr, v)
// 		} else {
// 			varr = v
// 		}

// 		for i := 0; i < len(arr); i++ {
// 			newVal := reflect.New(v.Type().Elem()).Elem()

// 			if d.setFieldByType(newVal, values, namespace, i, dm, errs) {
// 				set = true
// 				varr.Index(i).Set(newVal)
// 			}
// 		}

// 		if !set {
// 			return
// 		}

// 		v.Set(varr)

// 	// case reflect.Map:
// 	// 	if arr = values[namespace]; len(arr) == 0 {
// 	// 		return
// 	// 	}

// 	// 	v.Set(reflect.ValueOf(arr[0]))
// 	// 	set = true

// 	case reflect.Map:

// 		var rd *recursiveData

// 		if rd = dm[namespace]; rd == nil {
// 			return
// 		}

// 		var mp reflect.Value
// 		typ := v.Type()

// 		if v.IsNil() {
// 			mp = reflect.MakeMap(typ)
// 		} else {
// 			mp = v
// 		}

// 		for i := 0; i < len(rd.keys); i++ {
// 			newVal := reflect.New(typ.Elem())

// 			if d.setFieldByType(newVal, values, namespace+rd.keys[i].searchValue, 0, dm, errs) {
// 				set = true
// 				mp.SetMapIndex(rd.keys[i].value, newVal)
// 			}
// 		}

// 		if !set {
// 			return
// 		}

// 		// set here
// 	case reflect.Struct:

// 		// TODO: time is struct
// 		if v.Type() == timeType {
// 			// Parse time... but how?
// 			// think maybe this should be left out and must specify a custom type function
// 			// because everyone may have different requirements...
// 			return
// 		}

// 		set = d.traverseStruct(v, values, namespace, dm, errs)
// 	default:
// 		// look for custom type? or should it be done before this switch, must check out bson.ObjectId because is of typee
// 		// string but requires a specific method to ensure that it's valid
// 		fmt.Println("Currently unknown!")
// 	}
// }
