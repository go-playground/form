package form

import (
	"bytes"
	"net/url"
	"reflect"
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

// Decoder is the main decode instance
type Decoder struct {
	tagName         string
	structCache     structCacheMap
	customTypeFuncs map[reflect.Type]DecodeCustomTypeFunc
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

// RegisterCustomTypeFunc registers a DecodeCustomTypeFunc against a number of types; the function
// will also be used within the map key section.
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
	s := cachedStruct{fields: make([]cachedField, 0, 1)}

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

	dec.traverseStruct(val, "")

	if len(dec.errs) == 0 {
		return nil
	}

	err = dec.errs

	return
}
