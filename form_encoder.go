package form

import (
	"bytes"
	"net/url"
	"reflect"
	"strings"
)

// EncodeCustomTypeFunc allows for registering/overriding types to be parsed.
type EncodeCustomTypeFunc func(x interface{}) ([]string, error)

// EncodeErrors is a map of errors encountered during form encoding
type EncodeErrors map[string]error

func (e EncodeErrors) Error() string {
	buff := bytes.NewBufferString(blank)

	for k, err := range e {
		buff.WriteString(fieldNS)
		buff.WriteString(k)
		buff.WriteString(errorText)
		buff.WriteString(err.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// Encoder is the main encode instance
type Encoder struct {
	tagName         string
	structCache     *structCacheMap
	customTypeFuncs map[reflect.Type]EncodeCustomTypeFunc
}

// NewEncoder creates a new encoder instance with sane defaults
func NewEncoder() *Encoder {

	return &Encoder{
		tagName:     "form",
		structCache: newStructCacheMap(),
	}
}

// SetTagName sets the given tag name to be used by the decoder.
// Default is "form"
func (e *Encoder) SetTagName(tagName string) {
	e.tagName = tagName
}

// RegisterCustomTypeFunc registers a CustomTypeFunc against a number of types
// NOTE: this method is not thread-safe it is intended that these all be registered prior to any parsing
func (e *Encoder) RegisterCustomTypeFunc(fn EncodeCustomTypeFunc, types ...interface{}) {

	if e.customTypeFuncs == nil {
		e.customTypeFuncs = map[reflect.Type]EncodeCustomTypeFunc{}
	}

	for _, t := range types {
		e.customTypeFuncs[reflect.TypeOf(t)] = fn
	}
}

// Encode encodes the given values and sets the corresponding struct values
func (e *Encoder) Encode(v interface{}) (url.Values, error) {

	enc := &encoder{
		e:      e,
		values: make(url.Values),
	}

	val, kind := ExtractType(reflect.ValueOf(v))

	if kind != reflect.Struct {
		panic("interface must be a struct, pointer to a struct or interface containing one of the aforementioned")
	}

	enc.traverseStruct(val, make([]byte, 0, 64), -1)

	if len(enc.errs) == 0 {
		return enc.values, nil
	}

	return enc.values, enc.errs
}
