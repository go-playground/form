package form

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

type encoder struct {
	e      *Encoder
	errs   EncodeErrors
	values url.Values
}

func (e *encoder) setError(namespace string, err error) {
	if e.errs == nil {
		e.errs = make(EncodeErrors)
	}

	e.errs[namespace] = err
}

func (e *encoder) traverseStruct(v reflect.Value, namespace string) {

	typ := v.Type()
	var nn string // new namespace

	// is anonymous struct, cannot parse or cache as
	// it has no name to index by and potentially a
	// dynamic value
	if len(typ.Name()) == 0 {

		numFields := v.NumField()
		var fld reflect.StructField
		var key string

		for i := 0; i < numFields; i++ {

			fld = typ.Field(i)

			if fld.PkgPath != blank && !fld.Anonymous {
				continue
			}

			if key = fld.Tag.Get(e.e.tagName); key == ignore {
				continue
			}

			if len(key) == 0 {
				key = fld.Name
			}

			if len(namespace) == 0 {
				nn = key
			} else {
				nn = namespace + namespaceSeparator + key
			}

			e.setFieldByType(v.Field(i), nn)

		}
	} else {
		s, ok := e.e.structCache.Get(typ)
		if !ok {
			s = e.e.parseStruct(v)
		}

		for _, f := range s.fields {

			if len(namespace) == 0 {
				nn = f.name
			} else {
				nn = namespace + namespaceSeparator + f.name
			}

			e.setFieldByType(v.Field(f.idx), nn)
		}
	}

	return
}

func (e *encoder) setFieldByType(current reflect.Value, namespace string) {

	v, kind := ExtractType(current)

	if e.e.customTypeFuncs != nil {

		if cf, ok := e.e.customTypeFuncs[v.Type()]; ok {
			arr, err := cf(v.Interface())
			if err != nil {
				e.setError(namespace, err)
				return
			}

			e.values[namespace] = arr
			return
		}
	}

	switch kind {
	case reflect.Interface, reflect.Invalid:
		return
	case reflect.Ptr:

		e.setFieldByType(v.Elem(), namespace)

	case reflect.String:

		e.values[namespace] = []string{v.String()}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		e.values[namespace] = []string{strconv.FormatUint(v.Uint(), 10)}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		e.values[namespace] = []string{strconv.FormatInt(v.Int(), 10)}

	case reflect.Float32:

		e.values[namespace] = []string{strconv.FormatFloat(v.Float(), 'f', -1, 32)}

	case reflect.Float64:

		e.values[namespace] = []string{strconv.FormatFloat(v.Float(), 'f', -1, 64)}

	case reflect.Bool:

		e.values[namespace] = []string{strconv.FormatBool(v.Bool())}

	case reflect.Slice, reflect.Array:

		for i := 0; i < v.Len(); i++ {
			e.setFieldByType(v.Index(i), namespace+"["+strconv.Itoa(i)+"]")
		}

	case reflect.Map:

		for _, key := range v.MapKeys() {
			e.setFieldByType(current.MapIndex(key), namespace+"["+e.getMapKey(key, namespace)+"]")
		}

	case reflect.Struct:

		// if we get here then no custom time function declared so use RFC3339 by default
		if v.Type() == timeType {
			e.values[namespace] = []string{v.Interface().(time.Time).Format(time.RFC3339)}
			return
		}

		e.traverseStruct(v, namespace)
	}

	return
}

func (e *encoder) getMapKey(key reflect.Value, namespace string) string {

	v, kind := ExtractType(key)

	if e.e.customTypeFuncs != nil {

		if cf, ok := e.e.customTypeFuncs[v.Type()]; ok {
			arr, err := cf(v.Interface())
			if err != nil {
				e.setError(namespace, err)
				return ""
			}

			return arr[0]
		}
	}

	switch kind {
	case reflect.Interface, reflect.Ptr:
		return ""

	case reflect.String:
		return v.String()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)

	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)

	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	default:
		e.setError(namespace, fmt.Errorf("Unsupported Map Key '%v' Namespace '%s'", v.String(), namespace))
		return ""
	}
}
