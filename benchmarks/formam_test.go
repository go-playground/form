package benchmarks

import (
	"net/url"
	"testing"

	"github.com/monoculum/formam"
)

// Simple Benchmarks

func BenchmarkSimpleUserStructFormam(b *testing.B) {

	values := getUserStructValues()
	var err error
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test User
		if err = formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSimpleUserStructFormamParallel(b *testing.B) {

	values := getUserStructValues()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test User
			if err = formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Primitives ALL types

func BenchmarkPrimitivesStructAllPrimitivesFormamTypes(b *testing.B) {
	values := getPrimitivesStructValues()
	var err error

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test PrimitivesStruct
		if err = formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel(b *testing.B) {
	values := getPrimitivesStructValues()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test PrimitivesStruct
			if err = formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Array ALL types

func BenchmarkComplexArrayStructAllTypesFormam(b *testing.B) {
	values := getComplexArrayStructValues()
	var err error

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexArrayStruct
		if err = formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexArrayStructAllTypesFormamParallel(b *testing.B) {
	values := getComplexArrayStructValues()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexArrayStruct
			if err = formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Map ALL types

func getComplexMapStructValuesFormam() url.Values {
	return url.Values{
		"String.key":       []string{"value"},
		"StringPtr.key":    []string{"value"},
		"Int.0":            []string{"1"},
		"IntPtr.0":         []string{"1"},
		"Int8.0":           []string{"1"},
		"Int8Ptr.0":        []string{"1"},
		"Int16.0":          []string{"1"},
		"Int16Ptr.0":       []string{"1"},
		"Int32.0":          []string{"1"},
		"Int32Ptr.0":       []string{"1"},
		"Int64.0":          []string{"1"},
		"Int64Ptr.0":       []string{"1"},
		"Uint.0":           []string{"1"},
		"UintPtr.0":        []string{"1"},
		"Uint8.0":          []string{"1"},
		"Uint8Ptr.0":       []string{"1"},
		"Uint16.0":         []string{"1"},
		"Uint16Ptr.0":      []string{"1"},
		"Uint32.0":         []string{"1"},
		"Uint32Ptr.0":      []string{"1"},
		"Uint64.0":         []string{"1"},
		"Uint64Ptr.0":      []string{"1"},
		"NestedInt.1.2":    []string{"3"},
		"NestedIntPtr.1.2": []string{"3"},
	}
}

func BenchmarkComplexMapStructAllTypesFormam(b *testing.B) {
	// b.Log("Formam only supports map key of string at this time")
	// b.SkipNow()
	values := getComplexMapStructValuesFormam()
	var err error

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexMapStruct
		if err = formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexMapStructAllTypesFormamParallel(b *testing.B) {
	// b.Log("Formam only supports map key of string at this time")
	// b.SkipNow()
	values := getComplexMapStructValuesFormam()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexMapStruct
			if err = formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// NestedStruct Benchmarks

func BenchmarkArrayMapNestedStructFormam(b *testing.B) {

	values := getNestedStructValues()
	var err error
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test NestedStruct
		if err = formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkArrayMapNestedStructFormamParallel(b *testing.B) {

	values := getNestedStructValues()
	var err error

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test NestedStruct
			if err = formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}
