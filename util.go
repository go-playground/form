package form

import "reflect"

// ExtractType gets the actual underlying type of field value.
// it is exposed for use within you Custom Functions
func (d *Decoder) ExtractType(current reflect.Value) (reflect.Value, reflect.Kind) {

	switch current.Kind() {
	case reflect.Ptr:

		if current.IsNil() {
			return current, reflect.Ptr
		}

		return d.ExtractType(current.Elem())

	case reflect.Interface:

		if current.IsNil() {
			return current, reflect.Interface
		}

		return d.ExtractType(current.Elem())

	default:
		return current, current.Kind()
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
