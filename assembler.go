package assembler

import (
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
	errs := map[string]error{}
	dm := d.parseMapData(values)
	d.traverseStruct(reflect.ValueOf(v), values, "", dm, errs)

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

func (d *Decoder) traverseStruct(v reflect.Value, values url.Values, namespace string, dm dataMap, errs map[string]error) (set bool) {

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

func (d *Decoder) setFieldByType(current reflect.Value, values url.Values, namespace string, idx int, dm dataMap, errs map[string]error) (set bool) {

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
		newVal := reflect.New(v.Type().Elem())
		if set = d.setFieldByType(newVal.Elem(), values, namespace, idx, dm, errs); set {
			v.Set(newVal)
		}

	case reflect.String:
		if arr = values[namespace]; len(arr) == 0 {
			return
		}

		v.SetString(arr[0])
		set = true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var u64 uint64
		if arr, ok = values[namespace]; len(arr) == 0 {
			if !ok {
				return
			}

			v.SetInt(0)
			// v.Set(reflect.ValueOf(int8(0)))
			set = true
			return
		}

		if u64, err := strconv.ParseUint(arr[idx], 10, 64); err != nil || v.OverflowUint(u64) {
			errs[namespace] = fmt.Errorf("Invalid Integer Value '%s', Type '%v'", arr[idx], v.Type())
			return
		}

		v.SetUint(u64)
		// v.Set(reflect.ValueOf(int8(i64)))
		set = true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var i64 int64
		if arr, ok = values[namespace]; len(arr) == 0 {
			if !ok {
				return
			}

			v.SetInt(0)
			// v.Set(reflect.ValueOf(int8(0)))
			set = true
			return
		}

		if i64, err = strconv.ParseInt(arr[idx], 10, 64); err != nil || v.OverflowInt(i64) {
			errs[namespace] = fmt.Errorf("Invalid Integer Value '%s', Type '%v'", arr[idx], v.Type())
			return
		}

		v.SetInt(i64)
		// v.Set(reflect.ValueOf(int8(i64)))
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

		if v.IsNil() {
			varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
		} else if v.Len() < len(arr) {
			varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
			reflect.Copy(varr, v)
		} else {
			varr = v
		}

		for i := 0; i < len(arr); i++ {
			newVal := reflect.New(v.Type().Elem()).Elem()

			if d.setFieldByType(newVal, values, namespace, i, dm, errs) {
				set = true
				varr.Index(i).Set(newVal)
			}
		}

		if !set {
			return
		}

		v.Set(varr)

	// case reflect.Map:
	// 	if arr = values[namespace]; len(arr) == 0 {
	// 		return
	// 	}

	// 	v.Set(reflect.ValueOf(arr[0]))
	// 	set = true
	case reflect.Struct:

		if v.Type() == timeType {
			// Parse time... but how?
			// think maybe this should be left out and must specify a custom type function
			// because everyone may have different requirements...
			return
		}

		// TODO: time is struct
		set = d.traverseStruct(v, values, namespace, dm, errs)
	default:
		// look for custom type? or should it be done before this switch, must check out bson.ObjectId because is of typee
		// string but requires a specific method to ensure that it's valid
		fmt.Println("Currently unknown!")
	}

	return
}

// func int8Converter(kv []KeyValues) (value reflect.Value, err error) {
// 	var v int64

// 	v, err = strconv.ParseInt(kv[0].Values[0], 10, 8)

// 	value = reflect.ValueOf(int8(v))
// 	return
// }
