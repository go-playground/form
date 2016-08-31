Package form
============
<img align="right" src="https://raw.githubusercontent.com/go-playground/form/master/logo.jpg">
![Project status](https://img.shields.io/badge/version-2.2.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/form/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/form)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/form/badge.svg?branch=master)](https://coveralls.io/github/go-playground/form?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/form)](https://goreportcard.com/report/github.com/go-playground/form)
[![GoDoc](https://godoc.org/github.com/go-playground/form?status.svg)](https://godoc.org/github.com/go-playground/form)
![License](https://img.shields.io/dub/l/vibe-d.svg)
[![Gitter](https://badges.gitter.im/go-playground/form.svg)](https://gitter.im/go-playground/form?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Package form Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values.

It has the following features:

-   Supports map of almost all types.  
-   Supports both Numbered and Normal arrays eg. `"Array[0]"` and just `"Array"` with multiple values passed.
-   Array honours the specified index. eg. if `"Array[2]"` is the only Array value passed down, it will be put at index 2; if array isn't big enough it will be expanded.
-   Only creates objects as necessary eg. if no `array` or `map` values are passed down, the `array` and `map` are left as their default values in the struct.
-   Allows for Custom Type registration.
-   Handles time.Time using RFC3339 time format by default, but can easily be changed by registering a Custom Type, see below.
-   Handles Encoding & Decoding of almost all Go types eg. can Decode into struct, array, map, int... and Encode a struct, array, map, int...

Common Questions

-   Does it support encoding.TextUnmarshaler? No because TextUnmarshaler only accepts []byte but posted values can have multiple values, so is not suitable.

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
* many other types may be supported inherently (eg. `bson.ObjectId` is `type ObjectId string`, which will get populated by the string type

**NOTE**: `map`, `struct` and `slice` nesting are ad infinitum.

Installation
------------

Use go get.

	go get github.com/go-playground/form

Then import the form package into your own code.

	import "github.com/go-playground/form"
    
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

	"github.com/go-playground/form"
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

	"github.com/go-playground/form"
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
###### Run on MacBook Pro (Retina, 15-inch, Late 2013) 2.6 GHz Intel Core i7 16 GB 1600 MHz DDR3 using Go version go1.7 darwin/amd64

NOTE: the 1 allocation and B/op in the first 4 decodes is actually the struct allocating when passing it in, so primitives are actually zero allocation.

```go
go test -bench=. -benchmem=true

PASS
BenchmarkSimpleUserDecodeStruct-8                              	 5000000       	       312 ns/op       	      64 B/op  	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                      	20000000       	        91.7 ns/op     	      64 B/op  	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                              	 2000000       	       902 ns/op       	     485 B/op  	      11 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                      	 5000000       	       301 ns/op       	     485 B/op  	      11 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8            	 2000000       	      1028 ns/op       	      96 B/op  	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8    	 5000000       	       292 ns/op       	      96 B/op  	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8            	  300000       	      4770 ns/op       	    3009 B/op  	      46 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8    	 1000000       	      1569 ns/op       	    3010 B/op  	      46 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                    	  100000       	     15973 ns/op       	    2257 B/op  	     121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8            	  300000       	      4801 ns/op       	    2257 B/op  	     121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                    	  100000       	     15401 ns/op       	    7289 B/op  	     146 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8            	  300000       	      5167 ns/op       	    7289 B/op  	     146 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                      	   50000       	     20683 ns/op       	    5307 B/op  	     130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8              	  300000       	      6880 ns/op       	    5310 B/op  	     130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                      	  100000       	     15567 ns/op       	    7098 B/op  	     175 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8              	  300000       	      5546 ns/op       	    7099 B/op  	     175 allocs/op
BenchmarkDecodeNestedStruct-8                                  	  500000       	      3142 ns/op       	     384 B/op  	      14 allocs/op
BenchmarkDecodeNestedStructParallel-8                          	 1000000       	      1012 ns/op       	     384 B/op  	      14 allocs/op
BenchmarkEncodeNestedStruct-8                                  	 1000000       	      2106 ns/op       	     704 B/op  	      16 allocs/op
BenchmarkEncodeNestedStructParallel-8                          	 2000000       	       772 ns/op       	     704 B/op  	      16 allocs/op
```

Competitor benchmarks can be found [here](https://github.com/go-playground/form/blob/master/benchmarks/benchmarks.md)

Complimentary Software
----------------------

Here is a list of software that compliments using this library post decoding.

* [Validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
* [Conform](https://github.com/leebenson/conform) - Trims, sanitizes & scrubs data based on struct tags.

Package Versioning
----------
I'm jumping on the vendoring bandwagon, you should vendor this package as I will not
be creating different version with gopkg.in like allot of my other libraries.

Why? because my time is spread pretty thin maintaining all of the libraries I have + LIFE,
it is so freeing not to worry about it and will help me keep pouring out bigger and better
things for you the community.

License
------
Distributed under MIT License, please see license file in code for more details.
