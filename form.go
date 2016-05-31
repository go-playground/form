package form

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
	ignore             = "-"
	fieldNS            = "Field Namespace:"
	errorText          = " ERROR:"
)

var (
	timeType = reflect.TypeOf(time.Time{})
)

// CustomTypeFunc allows for registering/overriding types to be parsed.
type CustomTypeFunc func([]string) (interface{}, error)

// DecodeErrors is a map of errors encountered during form decoding
type DecodeErrors map[string]error

func (d DecodeErrors) Error() string {
	buff := bytes.NewBufferString(blank)

	for k, err := range d {
		buff.WriteString(fieldNS)
		buff.WriteString(k)
		buff.WriteString(errorText)
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
	sliceLen int
	keys     []key
	indicies []index
}

type dataMap map[string]*recursiveData

type formDecoder struct {
	d                   *Decoder
	errs                DecodeErrors
	dm                  dataMap
	values              url.Values
	maxKeyLen           int
	maxKeyLenCalculated bool
}

func (d *formDecoder) setError(namespace string, err error) {
	if d.errs == nil {
		d.errs = make(DecodeErrors)
	}

	d.errs[namespace] = err
}

// Decoder is the main decode instance
type Decoder struct {
	tagName         string
	structCache     structCacheMap
	customTypeFuncs map[reflect.Type]CustomTypeFunc
}

// NewDecoder creates a new decoder instance with sane defaults
func NewDecoder() *Decoder {
	return &Decoder{
		tagName:     "form",
		structCache: structCacheMap{m: map[reflect.Type]cachedStruct{}},
	}
}

// SetTagName sets the given tag name to be used by the decoder.
// Default is "form"
func (d *Decoder) SetTagName(tagName string) {
	d.tagName = tagName
}

// RegisterCustomTypeFunc registers a CustomTypeFunc against a number of types
// NOTE: this method is not thread-safe it is intended that these all be registered prior to any parsing
func (d *Decoder) RegisterCustomTypeFunc(fn CustomTypeFunc, types ...interface{}) {

	if d.customTypeFuncs == nil {
		d.customTypeFuncs = map[reflect.Type]CustomTypeFunc{}
	}

	for _, t := range types {
		d.customTypeFuncs[reflect.TypeOf(t)] = fn
	}
}

// Decode decodes the given values and sets the corresponding struct values
func (d *Decoder) Decode(v interface{}, values url.Values) (err error) {

	dec := &formDecoder{
		d:      d,
		values: values,
	}

	val := reflect.ValueOf(v)

	kind := val.Kind()

	if kind == reflect.Ptr {
		val = val.Elem()
	}

	if kind != reflect.Ptr || val.Kind() != reflect.Struct {
		panic("interface must be a pointer to a struct")
	}

	dec.traverseStruct(val, "")

	if len(dec.errs) == 0 {
		return nil
	}

	err = dec.errs

	return
}

func (d *formDecoder) parseMaxKeyLen() {

	for k := range d.values {

		if len(k) > d.maxKeyLen {
			d.maxKeyLen = len(k)
		}
	}

	d.maxKeyLenCalculated = true
}

func (d *formDecoder) parseMapData() {

	// already parsed
	if d.dm != nil {
		return
	}

	d.dm = make(dataMap)
	var idx int
	var idx2 int
	var cum int
	var cidx int
	var cidx2 int

	for k := range d.values {

		if len(k) > d.maxKeyLen {
			d.maxKeyLen = len(k)
		}

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

			cidx = cum + idx
			cidx2 = cum + idx2

			if rd, ok = d.dm[k[:cidx]]; !ok {
				rd = &recursiveData{
					keys:     make([]key, 0, 8),   // initializing with initial capacity of 8 to avoid too many reallocations of the underlying array
					indicies: make([]index, 0, 8), // initializing with initial capacity of 8 to avoid too many reallocations of the underlying array
				}
				d.dm[k[:cidx]] = rd
			}

			j, err := strconv.Atoi(k[cidx+1 : cidx2])

			// is map + key
			ke := key{
				value:       k[cidx+1 : cidx2],
				searchValue: k[cidx : cidx2+1],
			}
			rd.keys = append(rd.keys, ke)

			// only if no error otherwise not an index
			if err == nil {

				// is slice + indicies

				if j > rd.sliceLen {
					rd.sliceLen = j
				}

				ind := index{
					value:       j,
					searchValue: k[cidx : cidx2+1],
				}
				rd.indicies = append(rd.indicies, ind)
			}

			cum += idx2 + 1
		}
	}

	d.maxKeyLenCalculated = true
}

func (d *formDecoder) traverseStruct(v reflect.Value, namespace string) (set bool) {

	typ := v.Type()
	var nn string // new namespace

	// is anonymous struct, cannot parse or cache as
	// it has no name to index by
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

			if len(namespace) == 0 {
				nn = key
			} else {
				nn = namespace + namespaceSeparator + key
			}

			set = d.setFieldByType(v.Field(i), nn, 0)
		}
	} else {
		s, ok := d.d.structCache.Get(typ)
		if !ok {
			s = d.d.parseStruct(v)
		}

		for _, f := range s.fields {

			if len(namespace) == 0 {
				nn = f.name
			} else {
				nn = namespace + namespaceSeparator + f.name
			}

			set = d.setFieldByType(v.Field(f.idx), nn, 0)
		}
	}

	return
}

func (d *formDecoder) setFieldByType(current reflect.Value, namespace string, idx int) (set bool) {

	var err error

	v, kind := d.d.ExtractType(current)

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

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var u64 uint64

		if u64, err = strconv.ParseUint(arr[idx], 10, 64); err != nil || v.OverflowUint(u64) {
			d.setError(namespace, fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetUint(u64)
		set = true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if !ok || len(arr[idx]) == 0 {
			return
		}

		var i64 int64

		if i64, err = strconv.ParseInt(arr[idx], 10, 64); err != nil || v.OverflowInt(i64) {
			d.setError(namespace, fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetInt(i64)
		set = true

	case reflect.Float32, reflect.Float64:

		if !ok || len(arr[idx]) == 0 {
			return
		}

		var f float64

		if f, err = strconv.ParseFloat(arr[idx], 64); err != nil || v.OverflowFloat(f) {
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

		if b, err = strconv.ParseBool(arr[idx]); err != nil {
			d.setError(namespace, fmt.Errorf("Invalid Boolean Value '%s' Type '%v' Namespace '%s'", arr[idx], v.Type(), namespace))
			return
		}

		v.SetBool(b)
		set = true

	case reflect.Slice, reflect.Array:

		if !ok {

			d.parseMapData()

			// maybe it's an numbered array i.e. Pnone[0].Number
			if rd := d.dm[namespace]; rd != nil {

				var varr reflect.Value

				sl := rd.sliceLen + 1

				if v.IsNil() {
					varr = reflect.MakeSlice(v.Type(), sl, sl)
				} else if v.Len() < sl {
					if v.Cap() <= sl {
						varr = reflect.MakeSlice(v.Type(), sl, sl)
					} else {
						varr = reflect.MakeSlice(v.Type(), sl, v.Cap())
					}
					reflect.Copy(varr, v)
				} else {
					varr = v
				}

				for i := 0; i < len(rd.indicies); i++ {
					newVal := reflect.New(v.Type().Elem()).Elem()

					if d.setFieldByType(newVal, namespace+rd.indicies[i].searchValue, 0) {
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

			if err := d.getMapKey(rd.keys[i].value, kv, namespace); err != nil {
				d.setError(namespace, err)
				continue
			}

			if d.setFieldByType(newVal, namespace+rd.keys[i].searchValue, 0) {
				set = true
				mp.SetMapIndex(kv, newVal)
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
			return
		}

		if !d.maxKeyLenCalculated {
			d.parseMaxKeyLen()
		}

		// we must be recursing infinitly...but that's ok we caught it on the very first overun.
		if len(namespace) > d.maxKeyLen {
			return
		}

		set = d.traverseStruct(v, namespace)
	}

	return
}

func (d *formDecoder) getMapKey(key string, current reflect.Value, namespace string) (err error) {

	v, kind := d.d.ExtractType(current)

	switch kind {
	case reflect.Interface:
		v.Set(reflect.ValueOf(key))
		return
	case reflect.Ptr:

		newVal := reflect.New(v.Type().Elem())
		if err = d.getMapKey(key, newVal.Elem(), namespace); err == nil {
			v.Set(newVal)
		}

	case reflect.String:
		v.SetString(key)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		u64, e := strconv.ParseUint(key, 10, 64)
		if e != nil || v.OverflowUint(u64) {
			err = fmt.Errorf("Invalid Unsigned Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetUint(u64)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		i64, e := strconv.ParseInt(key, 10, 64)
		if e != nil || v.OverflowInt(i64) {
			err = fmt.Errorf("Invalid Integer Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetInt(i64)

	case reflect.Float32, reflect.Float64:

		f, e := strconv.ParseFloat(key, 64)
		if e != nil || v.OverflowFloat(f) {
			err = fmt.Errorf("Invalid Float Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetFloat(f)

	case reflect.Bool:

		b, e := strconv.ParseBool(key)
		if e != nil {
			err = fmt.Errorf("Invalid Boolean Value '%s' Type '%v' Namespace '%s'", key, v.Type(), namespace)
			return
		}

		v.SetBool(b)

	default:
		// look for custom type? or should it be done before this switch, must check out bson.ObjectId because is of typee
		// string but requires a specific method to ensure that it's valid
		err = fmt.Errorf("Unsupported Map Key '%s', Type '%v' Namespace '%s'", key, v.Type(), namespace)
	}

	return
}
