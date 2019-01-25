package form_test

import "github.com/swaggest/form"

func ExampleRegisterSQLNullTypesDecoders() {
	d := form.NewDecoder()
	form.RegisterSQLNullTypesDecoders(d, "NULL", "null")
}

func ExampleRegisterSQLNullTypesDecoders_omitNullValues() {
	d := form.NewDecoder()
	// If no null value strings are provided, "NULL" is used
	form.RegisterSQLNullTypesDecoders(d)
}

func ExampleRegisterSQLNullTypesEncoders() {
	e := form.NewEncoder()
	form.RegisterSQLNullTypesEncoders(e, "NULL")
}
