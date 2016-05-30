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

	case reflect.Invalid:
		return current, reflect.Invalid

	default:
		return current, current.Kind()
	}
}
