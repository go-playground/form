Package form
============
<img align="right" src="https://raw.githubusercontent.com/go-playground/form/master/logo.jpg">
![Project status](https://img.shields.io/badge/version-1.3.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/form/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/form)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/form/badge.svg?branch=master)](https://coveralls.io/github/go-playground/form?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/form)](https://goreportcard.com/report/github.com/go-playground/form)
[![GoDoc](https://godoc.org/github.com/go-playground/form?status.svg)](https://godoc.org/github.com/go-playground/form)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package form parses url.Values and fills a struct with values, creating objects as necessary.

It has the following features:

-   Supports map of almost all types.  
-   Supports both Numbered and Normal arrays eg. `"Array[0]"` and just `"Array"` with multiple values passed.
-   Array honours the specified index. eg. if `"Array[2]"` is the only Array value passed down, it will be put at index 2; if array isn't big enough it will be expanded.
-   Only creates objects as necessary eg. if no `array` or `map` values are passed down, the `array` and `map` are left as their default values in the struct.
-   Allows for Custom Type registration.
-   Handles time.Time using RFC3339 time format by default, but can easily be changed by registering a Custom Type, see below.

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

Example
-------
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

Registering Custom Types
--------------
```go
decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02", vals[0])
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

Benchmarks
------
###### Run on MacBook Pro (Retina, 15-inch, Late 2013) 2.6 GHz Intel Core i7 16 GB 1600 MHz DDR3 using Go version go1.6.2 darwin/amd64

NOTE: the 1 allocation and B/op in the first 4 is actually the struct allocating when passing it in, so primitives are actually zero allocation.

```go
go test -bench=. -benchmem=true

PASS
BenchmarkSimpleUserStruct-8                                 	 5000000	       299 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserStructParallel-8                         	20000000	       110 ns/op	      64 B/op	       1 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypes-8               	 2000000	       956 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesParallel-8       	 5000000	       285 ns/op	      96 B/op	       1 allocs/op
BenchmarkComplexArrayStructAllTypes-8                       	  100000	     20706 ns/op	    6776 B/op	     159 allocs/op
BenchmarkComplexArrayStructAllTypesParallel-8               	  200000	      6158 ns/op	    6776 B/op	     159 allocs/op
BenchmarkComplexMapStructAllTypes-8                         	   50000	     35548 ns/op	   20966 B/op	     245 allocs/op
BenchmarkComplexMapStructAllTypesParallel-8                 	  200000	     11984 ns/op	   20966 B/op	     245 allocs/op
BenchmarkArrayMapNestedStruct-8                             	  200000	      5617 ns/op	    2064 B/op	      37 allocs/op
BenchmarkArrayMapNestedStructParallel-8                     	 1000000	      2032 ns/op	    2064 B/op	      37 allocs/op
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
