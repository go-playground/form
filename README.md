Package form
============
<img align="right" src="logo.jpg">[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/go-playground/form)](https://github.com/go-playground/form/releases)
[![Build Status](https://github.com/go-playground/form/actions/workflows/workflow.yml/badge.svg)](https://github.com/go-playground/form/actions/workflows/workflow.yml)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/form/badge.svg?branch=master)](https://coveralls.io/github/go-playground/form?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/form)](https://goreportcard.com/report/github.com/go-playground/form)
[![GoDoc](https://godoc.org/github.com/go-playground/form?status.svg)](https://godoc.org/github.com/go-playground/form)
![License](https://img.shields.io/dub/l/vibe-d.svg)
[![Gitter](https://badges.gitter.im/go-playground/form.svg)](https://gitter.im/go-playground/form?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Package form Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values.

It has the following features:

- Supports map of almost all types.
- Supports both Numbered and Normal arrays eg. `"Array[0]"` and just `"Array"` with multiple values passed.
- Slice honours the specified index. eg. if "Slice[2]" is the only Slice value passed down, it will be put at index 2; if slice isn't big enough it will be expanded.
- Array honours the specified index. eg. if "Array[2]" is the only Array value passed down, it will be put at index 2; if array isn't big enough a warning will be printed and value ignored.
- Only creates objects as necessary eg. if no `array` or `map` values are passed down, the `array` and `map` are left as their default values in the struct.
- Allows for Custom Type registration.
- Handles time.Time using RFC3339 time format by default, but can easily be changed by registering a Custom Type, see below.
- Handles Encoding & Decoding of almost all Go types eg. can Decode into struct, array, map, int... and Encode a struct, array, map, int...

Common Questions

- Does it support encoding.TextUnmarshaler? No because TextUnmarshaler only accepts []byte but posted values can have multiple values, so is not suitable.
- Mixing `array/slice` with `array[idx]/slice[idx]`, in which order are they parsed? `array/slice` then `array[idx]/slice[idx]`

Supported Types ( out of the box )
----------

* `string`
* `bool`
* `int`, `int8`, `int16`, `int32`, `int64`
* `uint`, `uint8`, `uint16`, `uint32`, `uint64`
* `float32`, `float64`
* `struct` and `anonymous struct`
* `interface{}`
* `time.Time` - by default using RFC3339
* a `pointer` to one of the above types
* `slice`, `array`
* `map`
* `custom types` can override any of the above types
* many other types may be supported inherently

**NOTE**: `map`, `struct` and `slice` nesting are ad infinitum.

Installation
------------

Use go get.

	go get github.com/go-playground/form

Then import the form package into your own code.

	import "github.com/go-playground/form/v4"

Usage
-----

- Use symbol `.` for separating fields/structs. (eg. `structfield.field`)
- Use `[index or key]` for access to index of a slice/array or key for map. (eg. `arrayfield[0]`, `mapfield[keyvalue]`)

```html
<form method="POST">
  <input type="text" name="Name" value="joeybloggs"/>
  <input type="text" name="Age" value="3"/>
  <input type="text" name="Gender" value="Male"/>
  <input type="text" name="Address[0].Name" value="26 Here Blvd."/>
  <input type="text" name="Address[0].Phone" value="9(999)999-9999"/>
  <input type="text" name="Address[1].Name" value="26 There Blvd."/>
  <input type="text" name="Address[1].Phone" value="1(111)111-1111"/>
  <input type="text" name="active" value="true"/>
  <input type="text" name="MapExample[key]" value="value"/>
  <input type="text" name="NestedMap[key][key]" value="value"/>
  <input type="text" name="NestedArray[0][0]" value="value"/>
  <input type="submit"/>
</form>
```

Examples
-------

Decoding
```go
package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/go-playground/form/v4"
)

// Address contains address information
type Address struct {
	Name  string
	Phone string
}

// User contains user information
type User struct {
	Name        string
	Age         uint8
	Gender      string
	Address     []Address
	Active      bool `form:"active"`
	MapExample  map[string]string
	NestedMap   map[string]map[string]string
	NestedArray [][]string
}

// use a single instance of Decoder, it caches struct info
var decoder *form.Decoder

func main() {
	decoder = form.NewDecoder()

	// this simulates the results of http.Request's ParseForm() function
	values := parseForm()

	var user User

	// must pass a pointer
	err := decoder.Decode(&user, values)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", user)
}

// this simulates the results of http.Request's ParseForm() function
func parseForm() url.Values {
	return url.Values{
		"Name":                []string{"joeybloggs"},
		"Age":                 []string{"3"},
		"Gender":              []string{"Male"},
		"Address[0].Name":     []string{"26 Here Blvd."},
		"Address[0].Phone":    []string{"9(999)999-9999"},
		"Address[1].Name":     []string{"26 There Blvd."},
		"Address[1].Phone":    []string{"1(111)111-1111"},
		"active":              []string{"true"},
		"MapExample[key]":     []string{"value"},
		"NestedMap[key][key]": []string{"value"},
		"NestedArray[0][0]":   []string{"value"},
	}
}
```

Encoding
```go
package main

import (
	"fmt"
	"log"

	"github.com/go-playground/form/v4"
)

// Address contains address information
type Address struct {
	Name  string
	Phone string
}

// User contains user information
type User struct {
	Name        string
	Age         uint8
	Gender      string
	Address     []Address
	Active      bool `form:"active"`
	MapExample  map[string]string
	NestedMap   map[string]map[string]string
	NestedArray [][]string
}

// use a single instance of Encoder, it caches struct info
var encoder *form.Encoder

func main() {
	encoder = form.NewEncoder()

	user := User{
		Name:   "joeybloggs",
		Age:    3,
		Gender: "Male",
		Address: []Address{
			{Name: "26 Here Blvd.", Phone: "9(999)999-9999"},
			{Name: "26 There Blvd.", Phone: "1(111)111-1111"},
		},
		Active:      true,
		MapExample:  map[string]string{"key": "value"},
		NestedMap:   map[string]map[string]string{"key": {"key": "value"}},
		NestedArray: [][]string{{"value"}},
	}

	// must pass a pointer
	values, err := encoder.Encode(&user)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", values)
}
```

Registering Custom Types
--------------

Decoder
```go
decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
	return time.Parse("2006-01-02", vals[0])
}, time.Time{})
```
ADDITIONAL: if a struct type is registered, the function will only be called if a url.Value exists for
the struct and not just the struct fields eg. url.Values{"User":"Name%3Djoeybloggs"} will call the
custom type function with 'User' as the type, however url.Values{"User.Name":"joeybloggs"} will not.


Encoder
```go
encoder.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
	return []string{x.(time.Time).Format("2006-01-02")}, nil
}, time.Time{})
```

Ignoring Fields
--------------
you can tell form to ignore fields using `-` in the tag
```go
type MyStruct struct {
	Field string `form:"-"`
}
```

Omitempty
--------------
you can tell form to omit empty fields using `,omitempty` or `FieldName,omitempty` in the tag
```go
type MyStruct struct {
	Field  string `form:",omitempty"`
	Field2 string `form:"CustomFieldName,omitempty"`
}
```

Notes
------
To maximize compatibility with other systems the Encoder attempts
to avoid using array indexes in url.Values if at all possible.

eg.
```go
// A struct field of
Field []string{"1", "2", "3"}

// will be output a url.Value as
"Field": []string{"1", "2", "3"}

and not
"Field[0]": []string{"1"}
"Field[1]": []string{"2"}
"Field[2]": []string{"3"}

// however there are times where it is unavoidable, like with pointers
i := int(1)
Field []*string{nil, nil, &i}

// to avoid index 1 and 2 must use index
"Field[2]": []string{"1"}
```

Benchmarks
------
###### Run on M1 MacBook Pro using go version go1.20.6 darwin/amd64

NOTE: the 1 allocation and B/op in the first 4 decodes is actually the struct allocating when passing it in, so primitives are actually zero allocation.

```go
go test -run=NONE -bench=. -benchmem ./...
goos: darwin
goarch: arm64
pkg: github.com/go-playground/form/v4
cpu: Apple M3 Max
BenchmarkNestedArrayDecode100-16     	      75	  15782643 ns/op	18754349 B/op	  360810 allocs/op
BenchmarkNestedArrayDecode1000-16    	       1	2227892458 ns/op	1877558216 B/op	36011385 allocs/op
PASS
ok  	github.com/go-playground/form/v4	4.251s
goos: darwin
goarch: arm64
pkg: github.com/go-playground/form/v4/benchmarks
cpu: Apple M3 Max
BenchmarkSimpleUserDecodeStruct-16                              	12669696	        94.60 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-16                      	46715631	        27.79 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-16                              	 4624094	       256.7 ns/op	     485 B/op	      10 allocs/op
BenchmarkSimpleUserEncodeStructParallel-16                      	 7386290	       166.2 ns/op	     485 B/op	      10 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-16            	 3533421	       332.3 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-16    	20706642	        59.43 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-16            	 1228750	       966.4 ns/op	    1465 B/op	      34 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-16    	 1962678	       607.2 ns/op	    1465 B/op	      34 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-16                    	  213568	      5361 ns/op	    2081 B/op	     121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-16            	  960226	      1314 ns/op	    2087 B/op	     121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-16                    	  271944	      4017 ns/op	    6788 B/op	     107 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-16            	  441998	      2829 ns/op	    6791 B/op	     107 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-16                      	  179220	      6359 ns/op	    5300 B/op	     130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-16              	  412233	      2933 ns/op	    5310 B/op	     130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-16                      	  262464	      4122 ns/op	    4083 B/op	     106 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-16              	  622110	      2084 ns/op	    4084 B/op	     106 allocs/op
BenchmarkDecodeNestedStruct-16                                  	  823956	      1247 ns/op	     344 B/op	      14 allocs/op
BenchmarkDecodeNestedStructParallel-16                          	 4689418	       267.5 ns/op	     344 B/op	      14 allocs/op
BenchmarkEncodeNestedStruct-16                                  	 1844667	       636.0 ns/op	     653 B/op	      16 allocs/op
BenchmarkEncodeNestedStructParallel-16                          	 4302678	       278.8 ns/op	     653 B/op	      16 allocs/op
```

Competitor benchmarks can be found [here](https://github.com/go-playground/form/blob/master/benchmarks/benchmarks.md)


Maintenance and support for SDK major versions
----------------------------------------------

This package is aligned with the [Go release policy](https://go.dev/doc/devel/release) in that support is guaranteed for
the two most recent major versions.

This does not mean the package will not work with older versions of Go, only that we reserve the right to increase the
MSGV(Minimum Supported Go Version) when the need arises to address Security issues/patches, OS issues & support or newly
introduced functionality that would greatly benefit the maintenance and/or usage of this package.

If and when the MSGV is increased it will be done so in a minimum of a `Minor` release bump.

Complimentary Software
----------------------

Here is a list of software that compliments using this library post decoding.

* [Validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
* [mold](https://github.com/go-playground/mold) - Is a general library to help modify or set data within data structures and other objects.

License
------
Distributed under MIT License, please see license file in code for more details.
