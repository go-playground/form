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
	alias    string
	sliceLen int
	keys     []key
}

// type dataMap map[string]*recursiveData
type dataMap []*recursiveData

// Decoder is the main decode instance
type Decoder struct {
	tagName         string
	structCache     *structCacheMap
	customTypeFuncs map[reflect.Type]DecodeCustomTypeFunc
	maxArraySize    int
	keyPool         *sync.Pool
}

// NewDecoder creates a new decoder instance with sane defaults
func NewDecoder() *Decoder {

	return &Decoder{
		tagName:      "form",
		structCache:  newStructCacheMap(),
		maxArraySize: 10000,
		keyPool: &sync.Pool{New: func() interface{} {
			return make(dataMap, 0, 0)
			// return &recursiveData{
			// 	keys: make([]key, 0, 8), // initializing with initial capacity of 8 to avoid too many reallocations of the underlying array
			// }
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

	if len(dec.dm) > 0 {
		d.keyPool.Put(dec.dm)
	}
	// for _, v := range dec.dm {
	// 	d.keyPool.Put(v)
	// }

	if len(dec.errs) == 0 {
		return nil
	}

	err = dec.errs

	return
}
