package benchmarks

import (
	"testing"

	"github.com/monoculum/formam"
)

// Simple Benchmarks

func BenchmarkSimpleUserStructFormam(b *testing.B) {

	values := getUserStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test User
		if err := formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSimpleUserStructFormamParallel(b *testing.B) {

	values := getUserStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test User
			if err := formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Primitives ALL types

func BenchmarkPrimitivesStructAllPrimitivesFormamTypes(b *testing.B) {
	values := getPrimitivesStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test PrimitivesStruct
		if err := formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel(b *testing.B) {
	values := getPrimitivesStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test PrimitivesStruct
			if err := formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Array ALL types

func BenchmarkComplexArrayStructAllTypesFormam(b *testing.B) {
	values := getComplexArrayStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexArrayStruct
		if err := formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexArrayStructAllTypesFormamParallel(b *testing.B) {
	values := getComplexArrayStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexArrayStruct
			if err := formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Map ALL types

func BenchmarkComplexMapStructAllTypesFormam(b *testing.B) {
	// b.Log("Formam only supports map key of string at this time")
	// b.SkipNow()
	values := getComplexMapStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexMapStruct
		if err := formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexMapStructAllTypesFormamParallel(b *testing.B) {
	// b.Log("Formam only supports map key of string at this time")
	// b.SkipNow()
	values := getComplexMapStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexMapStruct
			if err := formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}

// NestedStruct Benchmarks

func BenchmarkArrayMapNestedStructFormam(b *testing.B) {

	values := getNestedStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test NestedStruct
		if err := formam.Decode(values, &test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkArrayMapNestedStructFormamParallel(b *testing.B) {

	values := getNestedStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test NestedStruct
			if err := formam.Decode(values, &test); err != nil {
				b.Error(err)
			}
		}
	})
}
