package form

import (
	"bytes"
	"net/url"
	"reflect"
	"strings"
	"sync"
)

// DecodeCustomTypeFunc allows for registering/overriding types to be parsed.
type DecodeCustomTypeFunc func([]string) (interface{}, error)

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
	ivalue      int
	value       string
	searchValue string
}

type recursiveData struct {
	sliceLen int
	keys     []key
}

type dataMap map[string]*recursiveData

// Decoder is the main decode instance
type Decoder struct {
	tagName         string
	structCache     structCacheMap
	customTypeFuncs map[reflect.Type]DecodeCustomTypeFunc
	maxArraySize    int
	keyPool         *sync.Pool
}

// NewDecoder creates a new decoder instance with sane defaults
func NewDecoder() *Decoder {
	return &Decoder{
		tagName:      "form",
		structCache:  structCacheMap{m: map[reflect.Type]cachedStruct{}},
		maxArraySize: 10000,
		keyPool: &sync.Pool{New: func() interface{} {
			return &recursiveData{
				keys: make([]key, 0, 8), // initializing with initial capacity of 8 to avoid too many reallocations of the underlying array
			}
		}},
	}
}

// SetTagName sets the given tag name to be used by the decoder.
// Default is "form"
func (d *Decoder) SetTagName(tagName string) {
	d.tagName = tagName
}

// SetMaxArraySize sets maximum array size that can be created.
// This limit is for the array indexing this library supports to
// avoid potential DOS or man-in-the-middle attacks using an unusually
// high number.
// DEFAULT: 10000
func (d *Decoder) SetMaxArraySize(size uint) {
	d.maxArraySize = int(size)
}

// RegisterCustomTypeFunc registers a CustomTypeFunc against a number of types
// NOTE: this method is not thread-safe it is intended that these all be registered prior to any parsing
func (d *Decoder) RegisterCustomTypeFunc(fn DecodeCustomTypeFunc, types ...interface{}) {

	if d.customTypeFuncs == nil {
		d.customTypeFuncs = map[reflect.Type]DecodeCustomTypeFunc{}
	}

	for _, t := range types {
		d.customTypeFuncs[reflect.TypeOf(t)] = fn
	}
}

func (d *Decoder) parseStruct(current reflect.Value) cachedStruct {

	typ := current.Type()
	s := cachedStruct{fields: make([]cachedField, 0, 4)} // init 4, betting most structs decoding into have at aleast 4 fields.

	numFields := current.NumField()

	var fld reflect.StructField
	var name string

	for i := 0; i < numFields; i++ {

		fld = typ.Field(i)

		if fld.PkgPath != blank && !fld.Anonymous {
			continue
		}

		if name = fld.Tag.Get(d.tagName); name == ignore {
			continue
		}

		if len(name) == 0 {
			name = fld.Name
		}

		s.fields = append(s.fields, cachedField{idx: i, name: name})
	}

	d.structCache.Set(typ, s)

	return s
}

// Decode decodes the given values and sets the corresponding struct values
func (d *Decoder) Decode(v interface{}, values url.Values) (err error) {

	dec := &decoder{
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

	dec.traverseStruct(val, make([]byte, 0, 64))

	for _, v := range dec.dm {
		d.keyPool.Put(v)
	}

	if len(dec.errs) == 0 {
		return nil
	}

	err = dec.errs

	return
}
