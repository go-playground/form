package assembler

import "reflect"

// ExtractType gets the actual underlying type of field value.
// It will dive into pointers, customTypes and return you the
// underlying value and it's kind.
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

		// if d.hasCustomFuncs {
		// 	// fmt.Println("Type", current.Type())
		// 	if fn, ok := d.customTypeFuncs[current.Type()]; ok {

		// 		// fmt.Println("OK")

		// 		return d.ExtractType(reflect.ValueOf(fn(current)))
		// 	}

		// 	// fmt.Println("NOT OK")
		// }

		return current, current.Kind()
	}
}
