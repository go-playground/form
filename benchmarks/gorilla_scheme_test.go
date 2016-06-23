package benchmarks

import (
	"net/url"
	"testing"

	"github.com/gorilla/schema"
)

// Simple Benchmarks

func BenchmarkSimpleUserStructGorilla(b *testing.B) {

	values := getUserStructValues()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test User
		if err := decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSimpleUserStructGorillaParallel(b *testing.B) {

	values := getUserStructValues()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test User
			if err := decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

// Primitives ALL types

func BenchmarkPrimitivesStructAllPrimitivesTypesGorilla(b *testing.B) {
	values := getPrimitivesStructValues()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test PrimitivesStruct
		if err := decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel(b *testing.B) {
	values := getPrimitivesStructValues()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test PrimitivesStruct
			if err := decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Array ALL types

func BenchmarkComplexArrayStructAllTypesGorilla(b *testing.B) {
	values := getComplexArrayStructValues()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexArrayStruct
		if err := decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexArrayStructAllTypesGorillaParallel(b *testing.B) {
	values := getComplexArrayStructValues()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexArrayStruct
			if err := decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Map ALL types

func BenchmarkComplexMapStructAllTypesGorilla(b *testing.B) {
	b.Log("Gorilla does not support map parsing at this time")
	b.SkipNow()
}

func BenchmarkComplexMapStructAllTypesGorillaParallel(b *testing.B) {
	b.Log("Gorilla does not support map parsing at this time")
	b.SkipNow()
}

// NestedStruct Benchmarks

func getNestedStructValuesGorilla() url.Values {
	return url.Values{
		// Nested Field
		"Value": []string{"value"},
		// Nested Array
		"NestedArray.0.Value": []string{"value"},
		"NestedArray.1.Value": []string{"value"},
		// Nested Array Ptr
		"NestedPtrArray.0.Value": []string{"value"},
		"NestedPtrArray.1.Value": []string{"value"},
		// Nested 2
		"Nested2.Value":         []string{"value"},
		"Nested2.Nested2.Value": []string{"value"},
	}
}

func BenchmarkArrayMapNestedStructGorilla(b *testing.B) {

	values := getNestedStructValuesGorilla()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test NestedStruct
		if err := decoder.Decode(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkArrayMapNestedStructGorillaParallel(b *testing.B) {

	values := getNestedStructValuesGorilla()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test NestedStruct
			if err := decoder.Decode(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}
