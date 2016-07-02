package form

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const (
	errArraySize           = "Array size of '%d' is larger than the maximum currently set on the decoder of '%d'. To increase this limit please see, SetMaxArraySize(size uint)"
	errMissingStartBracket = "Invalid formatting for key '%s' missing '[' bracket"
	errMissingEndBracket   = "Invalid formatting for key '%s' missing ']' bracket"
)

// TODO: test namespace as []byte, except when passing error to reduce allocations.

// TODO: try non pointer methods on decoder...it is a small struct and may be more efficient to copy than by ref
// possibly even remove and pass arount all values as params
type decoder struct {
	d         *Decoder
	errs      DecodeErrors
	dm        dataMap
	values    url.Values
	maxKeyLen int
}

func (d *decoder) setError(namespace string, err error) {
	if d.errs == nil {
		d.errs = make(DecodeErrors)
	}

	d.errs[namespace] = err
}

func (d *decoder) parseMapData() {

	// already parsed
	if d.dm != nil {
		return
	}

	d.dm = make(dataMap)
	var i int
	var idx int
	var idxP1 int
	var insideBracket bool
	var rd *recursiveData
	var ok bool
	var err error

	for k := range d.values {

		if len(k) > d.maxKeyLen {
			d.maxKeyLen = len(k)
		}

		for i = 0; i < len(k); i++ {

			switch k[i] {
			case '[':
				idx = i
				insideBracket = true
			case ']':

				if !insideBracket {
					log.Panicf(errMissingStartBracket, k)
				}

				if rd, ok = d.dm[k[:idx]]; !ok {
					rd = d.d.keyPool.Get().(*recursiveData)
					rd.keys = rd.keys[0:0]
					d.dm[k[:idx]] = rd
				}

				idxP1 = idx + 1

				// is map + key
				ke := key{
					value:       k[idxP1:i],
					searchValue: k[idx : i+1],
				}

				ke.ivalue, err = strconv.Atoi(k[idxP1:i])
				// only if no error otherwise not an index
				if err == nil {

					// may be slice

					if ke.ivalue > rd.sliceLen {
						rd.sliceLen = ke.ivalue
					}
				} else {
					ke.ivalue = -1
				}

				rd.keys = append(rd.keys, ke)

				insideBracket = false
			}
		}

		// if still inside bracket, that means no ending bracket was ever specified
		if insideBracket {
			log.Panicf(errMissingEndBracket, k)
		}
	}
}

func (d *decoder) traverseStruct(v reflect.Value, namespace string) (set bool) {

	typ := v.Type()
	var nn string // new namespace
	first := len(namespace) == 0

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

			if key = fld.Tag.Get(d.d.tagName); key == ignore {
				continue
			}

			if len(key) == 0 {
				key = fld.Name
			}

			if first {
				nn = key
			} else {
				nn = namespace + namespaceSeparator + key
			}

			if d.setFieldByType(v.Field(i), nn, 0) {
				set = true
			}

		}
	} else {
		s, ok := d.d.structCache.Get(typ)
		if !ok {
			s = d.d.parseStruct(v)
		}

		for _, f := range s.fields {

			if first {
				nn = f.name
			} else {
				nn = namespace + namespaceSeparator + f.name
			}

			if d.setFieldByType(v.Field(f.idx), nn, 0) {
				set = true
			}
		}
	}

	return
}

func (d *decoder) setFieldByType(current reflect.Value, namespace string, idx int) (set bool) {

	var err error

	v, kind := ExtractType(current)

	arr, ok := d.values[namespace]

	if d.d.customTypeFuncs != nil {

		if ok {

			if cf, ok := d.d.customTypeFuncs[v.Type()]; ok {
				val, err := cf(arr)
				if err != nil {
					d.setError(namespace, err)
					return
				}

				v.Set(reflect.ValueOf(val))
				set = true
				return
			}
		}
	}

	switch kind {
	case reflect.Interface, reflect.Invalid:
		return
	case reflect.Ptr:

		newVal := reflect.New(v.Type().Elem())
		if set = d.setFieldByType(newVal.Elem(), namespace, idx); set {
			v.Set(newVal)
		}

	case reflect.String:

		if !ok {
			return
		}

		v.SetString(arr[idx])
		set = true

	case reflect.Uint, reflect.Uint64:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var u64 uint64

		if u64, err = strconv.ParseUint(arr[idx], 10, 64); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetUint(u64)
		set = true

	case reflect.Uint8:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var u64 uint64

		if u64, err = strconv.ParseUint(arr[idx], 10, 8); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetUint(u64)
		set = true

	case reflect.Uint16:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var u64 uint64

		if u64, err = strconv.ParseUint(arr[idx], 10, 16); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetUint(u64)
		set = true

	case reflect.Uint32:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var u64 uint64

		if u64, err = strconv.ParseUint(arr[idx], 10, 32); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetUint(u64)
		set = true

	case reflect.Int, reflect.Int64:
		if !ok || len(arr[idx]) == 0 {
			return
		}

		var i64 int64

		if i64, err = strconv.ParseInt(arr[idx], 10, 64); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetInt(i64)
		set = true

	case reflect.Int8:
		if !ok || len(arr[idx]) == 0 {
			return
		}

		var i64 int64

		if i64, err = strconv.ParseInt(arr[idx], 10, 8); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetInt(i64)
		set = true

	case reflect.Int16:
		if !ok || len(arr[idx]) == 0 {
			return
		}

		var i64 int64

		if i64, err = strconv.ParseInt(arr[idx], 10, 16); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetInt(i64)
		set = true

	case reflect.Int32:
		if !ok || len(arr[idx]) == 0 {
			return
		}

		var i64 int64

		if i64, err = strconv.ParseInt(arr[idx], 10, 32); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetInt(i64)
		set = true

	case reflect.Float32:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var f float64

		if f, err = strconv.ParseFloat(arr[idx], 32); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Float Value '%s' Type '%v' Namespace '%s'", arr[0], v.Type(), namespace))
			return
		}

		v.SetFloat(f)
		set = true

	case reflect.Float64:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var f float64

		if f, err = strconv.ParseFloat(arr[idx], 64); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Float Value '%s' Type '%v' Namespace '%s'", arr[0], v.Type(), namespace))
			return
		}

		v.SetFloat(f)
		set = true

	case reflect.Bool:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var b bool

		if b, err = parseBool(arr[idx]); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Boolean Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetBool(b)
		set = true

	case reflect.Slice, reflect.Array:

		if !ok {

			d.parseMapData()

			// maybe it's an numbered array i.e. Phone[0].Number
			if rd := d.dm[namespace]; rd != nil {

				var varr reflect.Value
				var kv key

				sl := rd.sliceLen + 1

				// checking below for maxArraySize, but if array exists and already
				// has sufficient capacity allocated then we do not check as the code
				// obviously allows that capacity.

				if v.IsNil() {

					if sl > d.d.maxArraySize {
						d.setError(namespace, fmt.Errorf(errArraySize, sl, d.d.maxArraySize))
						return
					}

					varr = reflect.MakeSlice(v.Type(), sl, sl)

				} else if v.Len() < sl {

					if v.Cap() <= sl {

						if sl > d.d.maxArraySize {
							d.setError(namespace, fmt.Errorf(errArraySize, sl, d.d.maxArraySize))
							return
						}

						varr = reflect.MakeSlice(v.Type(), sl, sl)
					} else {
						varr = reflect.MakeSlice(v.Type(), sl, v.Cap())
					}

					reflect.Copy(varr, v)

				} else {
					varr = v
				}

				for i := 0; i < len(rd.keys); i++ {

					kv = rd.keys[i]
					newVal := reflect.New(varr.Type().Elem()).Elem()

					if kv.ivalue == -1 {
						d.setError(namespace, fmt.Errorf("Invalid Array index '%s'", kv.value))
						continue
					}

					if d.setFieldByType(newVal, namespace+kv.searchValue, 0) {
						set = true
						varr.Index(kv.ivalue).Set(newVal)
					}
				}

				if !set {
					return
				}

				v.Set(varr)
			}

			return
		}

		if len(arr) == 0 {
			return
		}

		var varr reflect.Value
		var existing bool

		if v.IsNil() {
			varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
		} else if v.Len() < len(arr) {
			if v.Cap() <= len(arr) {
				varr = reflect.MakeSlice(v.Type(), len(arr), len(arr))
			} else {
				varr = reflect.MakeSlice(v.Type(), len(arr), v.Cap())
			}
			reflect.Copy(varr, v)
		} else {
			existing = true
			varr = v
		}

		for i := 0; i < len(arr); i++ {
			newVal := reflect.New(v.Type().Elem()).Elem()

			if d.setFieldByType(newVal, namespace, i) {
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

		d.parseMapData()

		// no natural map support so skip directly to dm lookup
		if rd = d.dm[namespace]; rd == nil {
			return
		}

		var existing bool
		var kv key
		var mp reflect.Value
		var mk reflect.Value

		typ := v.Type()

		if v.IsNil() {
			mp = reflect.MakeMap(typ)
		} else {
			existing = true
			mp = v
		}

		for i := 0; i < len(rd.keys); i++ {
			newVal := reflect.New(typ.Elem()).Elem()
			mk = reflect.New(typ.Key()).Elem()
			kv = rd.keys[i]

			if err := d.getMapKey(kv.value, mk, namespace); err != nil {
				d.setError(namespace, err)
				continue
			}

			if d.setFieldByType(newVal, namespace+kv.searchValue, 0) {
				set = true
				mp.SetMapIndex(mk, newVal)
			}
		}

		if !set || existing {
			return
		}

		v.Set(mp)

	case reflect.Struct:

		// if we get here then no custom time function declared so use RFC3339 by default
		if v.Type() == timeType {

			if !ok || len(arr[idx]) == 0 {
				return
			}

			t, err := time.Parse(time.RFC3339, arr[idx])
			if err != nil {
				d.setError(namespace, err)
			}

			v.Set(reflect.ValueOf(t))
			set = true
			return
		}

		d.parseMapData()

		// we must be recursing infinitly...but that's ok we caught it on the very first overun.
		if len(namespace) > d.maxKeyLen {
			return
		}

		set = d.traverseStruct(v, namespace)
	}

	return
}

func (d *decoder) getMapKey(key string, current reflect.Value, namespace string) (err error) {

	v, kind := ExtractType(current)

	if d.d.customTypeFuncs != nil {
		if cf, ok := d.d.customTypeFuncs[v.Type()]; ok {
			val, er := cf([]string{key}) // TODO: []string escapes to heap, possible reuseable []string
			if er != nil {
				err = er
				return
			}

			v.Set(reflect.ValueOf(val))
			return
		}
	}

	switch kind {
	case reflect.Interface:
		// If interface would have been set on the struct before decoding,
		// say to a struct value we would not get here but kind would be struct.
		v.Set(reflect.ValueOf(key))
		return
	case reflect.Ptr:
		newVal := reflect.New(v.Type().Elem())
		if err = d.getMapKey(key, newVal.Elem(), namespace); err == nil {
			v.Set(newVal)
		}

	case reflect.String:
		v.SetString(key)

	case reflect.Uint, reflect.Uint64:

		u64, e := strconv.ParseUint(key, 10, 64)
		if e != nil {
			err = fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetUint(u64)

	case reflect.Uint8:

		u64, e := strconv.ParseUint(key, 10, 8)
		if e != nil {
			err = fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetUint(u64)

	case reflect.Uint16:

		u64, e := strconv.ParseUint(key, 10, 16)
		if e != nil {
			err = fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetUint(u64)

	case reflect.Uint32:

		u64, e := strconv.ParseUint(key, 10, 32)
		if e != nil {
			err = fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetUint(u64)

	case reflect.Int, reflect.Int64:

		i64, e := strconv.ParseInt(key, 10, 64)
		if e != nil {
			err = fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetInt(i64)

	case reflect.Int8:

		i64, e := strconv.ParseInt(key, 10, 8)
		if e != nil {
			err = fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetInt(i64)

	case reflect.Int16:

		i64, e := strconv.ParseInt(key, 10, 16)
		if e != nil {
			err = fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetInt(i64)

	case reflect.Int32:

		i64, e := strconv.ParseInt(key, 10, 32)
		if e != nil {
			err = fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetInt(i64)

	case reflect.Float32:

		f, e := strconv.ParseFloat(key, 32)
		if e != nil {
			err = fmt.Errorf("Invalid Float Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetFloat(f)

	case reflect.Float64:

		f, e := strconv.ParseFloat(key, 64)
		if e != nil {
			err = fmt.Errorf("Invalid Float Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetFloat(f)

	case reflect.Bool:

		b, e := parseBool(key)
		if e != nil {
			err = fmt.Errorf("Invalid Boolean Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetBool(b)

	default:
		err = fmt.Errorf("Unsupported Map Key '%s', Type '%v' Namespace '%s'", key, v.Type(), namespace)
	}

	return
}
