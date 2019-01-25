package form_test

import (
	"database/sql"
	"github.com/swaggest/form"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

type TestNullTypes struct {
	I sql.NullInt64   `db:"i"`
	F sql.NullFloat64 `db:"f"`
	S sql.NullString  `db:"s"`
	B sql.NullBool    `db:"b"`
}

func TestRegisterSQLNullTypesEncoders(t *testing.T) {
	e := form.NewEncoder()
	e.SetMode(form.ModeExplicit)
	e.SetTagName("db")
	form.RegisterSQLNullTypesEncoders(e, "NULL")

	testNullTypes := TestNullTypes{}
	v, err := e.Encode(testNullTypes)
	assert.Equal(t, err, nil)
	assert.Equal(t, v["i"], []string{"NULL"})
	assert.Equal(t, v["f"], []string{"NULL"})
	assert.Equal(t, v["s"], []string{"NULL"})
	assert.Equal(t, v["b"], []string{"NULL"})

	testNullTypes = TestNullTypes{
		I: sql.NullInt64{Int64: 123, Valid: true},
		F: sql.NullFloat64{Float64: 123.456, Valid: true},
		S: sql.NullString{String: "abc", Valid: true},
		B: sql.NullBool{Bool: false, Valid: true},
	}

	v, err = e.Encode(testNullTypes)
	assert.Equal(t, err, nil)
	assert.Equal(t, v["i"], []string{"123"})
	assert.Equal(t, v["f"], []string{"123.456"})
	assert.Equal(t, v["s"], []string{"abc"})
	assert.Equal(t, v["b"], []string{"false"})
}

func TestRegisterSQLNullTypesDecoders(t *testing.T) {
	d := form.NewDecoder()
	d.SetMode(form.ModeExplicit)
	d.SetTagName("db")
	form.RegisterSQLNullTypesDecoders(d, "NULL", "null")
	v := map[string][]string{
		"i": {"NULL"},
		"f": {"NULL"},
		"s": {"NULL"},
		"b": {"null"},
	}
	testNullTypes := TestNullTypes{}

	err := d.Decode(&testNullTypes, v)
	assert.Equal(t, err, nil)
	assert.Equal(t, testNullTypes.I.Valid, false)
	assert.Equal(t, testNullTypes.F.Valid, false)
	assert.Equal(t, testNullTypes.S.Valid, false)
	assert.Equal(t, testNullTypes.B.Valid, false)

	v = map[string][]string{
		"i": {"123"},
		"f": {"123.456"},
		"s": {"abc"},
		"b": {"false"},
	}

	err = d.Decode(&testNullTypes, v)
	assert.Equal(t, err, nil)
	assert.Equal(t, testNullTypes.I.Int64, int64(123))
	assert.Equal(t, testNullTypes.F.Float64, 123.456)
	assert.Equal(t, testNullTypes.S.String, "abc")
	assert.Equal(t, testNullTypes.B.Bool, false)
}
