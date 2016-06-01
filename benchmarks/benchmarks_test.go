package benchmarks

import (
	"net/url"
	"testing"

	"github.com/go-playground/form"
)

// Simple Benchmarks

type User struct {
	FirstName string `form:"fname" schema:"fname" formam:"fname"`
	LastName  string `form:"lname" schema:"lname" formam:"lname"`
	Email     string `form:"email" schema:"email" formam:"email"`
	Age       uint8  `form:"age"   schema:"age"   formam:"age"`
}

func getUserStructValues() url.Values {
	return url.Values{
		"fname": []string{"Joey"},
		"lname": []string{"Bloggs"},
		"email": []string{"joeybloggs@gmail.com"},
		"age":   []string{"32"},
	}
}

func BenchmarkSimpleUserStruct(b *testing.B) {

	values := getUserStructValues()
	decoder := form.NewDecoder()
	var err error
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test User
		if err = decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSimpleUserStructParallel(b *testing.B) {

	values := getUserStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test User
			if err = decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

// Primitives ALL types

type PrimitivesStruct struct {
	String  string
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Float32 float32
	Float64 float64
	Bool    bool
}

func getPrimitivesStructValues() url.Values {
	return url.Values{
		"String":  []string{"joeybloggs"},
		"Int":     []string{"1"},
		"Int8":    []string{"2"},
		"Int16":   []string{"3"},
		"Int32":   []string{"4"},
		"Int64":   []string{"5"},
		"Uint":    []string{"1"},
		"Uint8":   []string{"2"},
		"Uint16":  []string{"3"},
		"Uint32":  []string{"4"},
		"Uint64":  []string{"5"},
		"Float32": []string{"1.1"},
		"Float64": []string{"5.0"},
		"Bool":    []string{"true"},
	}
}

func BenchmarkPrimitivesStructAllPrimitivesTypes(b *testing.B) {
	values := getPrimitivesStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test PrimitivesStruct
		if err = decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPrimitivesStructAllPrimitivesTypesParallel(b *testing.B) {
	values := getPrimitivesStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test PrimitivesStruct
			if err = decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Array ALL types

type ComplexArrayStruct struct {
	String       []string
	StringPtr    []*string
	Int          []int
	IntPtr       []*int
	Int8         []int8
	Int8Ptr      []*int8
	Int16        []int16
	Int16Ptr     []*int16
	Int32        []int32
	Int32Ptr     []*int32
	Int64        []int64
	Int64Ptr     []*int64
	Uint         []uint
	UintPtr      []*uint
	Uint8        []uint8
	Uint8Ptr     []*uint8
	Uint16       []uint16
	Uint16Ptr    []*uint16
	Uint32       []uint32
	Uint32Ptr    []*uint32
	Uint64       []uint64
	Uint64Ptr    []*uint64
	NestedInt    [][]int
	NestedIntPtr [][]*int
}

func getComplexArrayStructValues() url.Values {
	return url.Values{
		"String":             []string{"joeybloggs"},
		"StringPtr":          []string{"joeybloggs"},
		"Int":                []string{"1", "2"},
		"IntPtr":             []string{"1", "2"},
		"Int8[0]":            []string{"1"},
		"Int8[1]":            []string{"2"},
		"Int8Ptr[0]":         []string{"1"},
		"Int8Ptr[1]":         []string{"2"},
		"Int16":              []string{"1", "2"},
		"Int16Ptr":           []string{"1", "2"},
		"Int32":              []string{"1", "2"},
		"Int32Ptr":           []string{"1", "2"},
		"Int64":              []string{"1", "2"},
		"Int64Ptr":           []string{"1", "2"},
		"Uint":               []string{"1", "2"},
		"UintPtr":            []string{"1", "2"},
		"Uint8[0]":           []string{"1"},
		"Uint8[1]":           []string{"2"},
		"Uint8Ptr[0]":        []string{"1"},
		"Uint8Ptr[1]":        []string{"2"},
		"Uint16":             []string{"1", "2"},
		"Uint16Ptr":          []string{"1", "2"},
		"Uint32":             []string{"1", "2"},
		"Uint32Ptr":          []string{"1", "2"},
		"Uint64":             []string{"1", "2"},
		"Uint64Ptr":          []string{"1", "2"},
		"NestedInt[0][0]":    []string{"1"},
		"NestedIntPtr[0][1]": []string{"1"},
	}
}

func BenchmarkComplexArrayStructAllTypes(b *testing.B) {
	values := getComplexArrayStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexArrayStruct
		if err = decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexArrayStructAllTypesParallel(b *testing.B) {
	values := getComplexArrayStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexArrayStruct
			if err = decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Map ALL types

type ComplexMapStruct struct {
	String       map[string]string
	StringPtr    map[*string]*string
	Int          map[int]int
	IntPtr       map[*int]*int
	Int8         map[int8]int8
	Int8Ptr      map[*int8]*int8
	Int16        map[int16]int16
	Int16Ptr     map[*int16]*int16
	Int32        map[int32]int32
	Int32Ptr     map[*int32]*int32
	Int64        map[int64]int64
	Int64Ptr     map[*int64]*int64
	Uint         map[uint]uint
	UintPtr      map[*uint]*uint
	Uint8        map[uint8]uint8
	Uint8Ptr     map[*uint8]*uint8
	Uint16       map[uint16]uint16
	Uint16Ptr    map[*uint16]*uint16
	Uint32       map[*uint32]*uint32
	Uint32Ptr    map[*uint32]*uint32
	Uint64       map[*uint64]*uint64
	Uint64Ptr    map[*uint64]*uint64
	NestedInt    map[int]map[int]int
	NestedIntPtr map[*int]map[*int]*int
}

func getComplexMapStructValues() url.Values {
	return url.Values{
		"String[key]":        []string{"value"},
		"StringPtr[key]":     []string{"value"},
		"Int[0]":             []string{"1"},
		"IntPtr[0]":          []string{"1"},
		"Int8[0]":            []string{"1"},
		"Int8Ptr[0]":         []string{"1"},
		"Int16[0]":           []string{"1"},
		"Int16Ptr[0]":        []string{"1"},
		"Int32[0]":           []string{"1"},
		"Int32Ptr[0]":        []string{"1"},
		"Int64[0]":           []string{"1"},
		"Int64Ptr[0]":        []string{"1"},
		"Uint[0]":            []string{"1"},
		"UintPtr[0]":         []string{"1"},
		"Uint8[0]":           []string{"1"},
		"Uint8Ptr[0]":        []string{"1"},
		"Uint16[0]":          []string{"1"},
		"Uint16Ptr[0]":       []string{"1"},
		"Uint32[0]":          []string{"1"},
		"Uint32Ptr[0]":       []string{"1"},
		"Uint64[0]":          []string{"1"},
		"Uint64Ptr[0]":       []string{"1"},
		"NestedInt[1][2]":    []string{"3"},
		"NestedIntPtr[1][2]": []string{"3"},
	}
}

func BenchmarkComplexMapStructAllTypes(b *testing.B) {
	values := getComplexMapStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexMapStruct
		if err = decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexMapStructAllTypesParallel(b *testing.B) {
	values := getComplexMapStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexMapStruct
			if err = decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

// NestedStruct Benchmarks

type Nested2 struct {
	Value   string
	Nested2 *Nested2
}

type Nested struct {
	Value string
}

type NestedStruct struct {
	Nested
	NestedArray    []Nested
	NestedPtrArray []*Nested
	Nested2        Nested2
}

func getNestedStructValues() url.Values {
	return url.Values{
		// Nested Field
		"Value": []string{"value"},
		// Nested Array
		"NestedArray[0].Value": []string{"value"},
		"NestedArray[1].Value": []string{"value"},
		// Nested Array Ptr
		"NestedPtrArray[0].Value": []string{"value"},
		"NestedPtrArray[1].Value": []string{"value"},
		// Nested 2
		"Nested2.Value":         []string{"value"},
		"Nested2.Nested2.Value": []string{"value"},
	}
}

func BenchmarkArrayMapNestedStruct(b *testing.B) {

	values := getNestedStructValues()
	decoder := form.NewDecoder()
	var err error
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test NestedStruct
		if err = decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkArrayMapNestedStructParallel(b *testing.B) {

	values := getNestedStructValues()
	decoder := form.NewDecoder()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test NestedStruct
			if err = decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}
