package benchmarks

import (
	"net/url"
	"testing"

	ajg "github.com/ajg/form"
)

// Simple Benchmarks

func BenchmarkSimpleUserDecodeStructAGJForm(b *testing.B) {

	values := getUserStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test User
		if err := ajg.DecodeValues(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSimpleUserDecodeStructParallelAGJFrom(b *testing.B) {

	values := getUserStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test User
			if err := ajg.DecodeValues(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkSimpleUserEncodeStructAGJForm(b *testing.B) {

	test := getUserStruct()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		if _, err := ajg.EncodeToValues(&test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSimpleUserEncodeStructParallelAGJForm(b *testing.B) {

	test := getUserStruct()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := ajg.EncodeToValues(&test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Primitives ALL types

func BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm(b *testing.B) {
	values := getPrimitivesStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test PrimitivesStruct
		if err := ajg.DecodeValues(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm(b *testing.B) {
	values := getPrimitivesStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test PrimitivesStruct
			if err := ajg.DecodeValues(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm(b *testing.B) {
	test := getPrimitivesStruct()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		if _, err := ajg.EncodeToValues(&test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm(b *testing.B) {
	test := getPrimitivesStruct()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := ajg.EncodeToValues(&test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Array ALL types

func BenchmarkComplexArrayDecodeStructAllTypesAGJForm(b *testing.B) {
	values := getComplexArrayStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexArrayStruct
		if err := ajg.DecodeValues(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm(b *testing.B) {
	values := getComplexArrayStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexArrayStruct
			if err := ajg.DecodeValues(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkComplexArrayEncodeStructAllTypesAGJForm(b *testing.B) {
	test := getComplexArrayStruct()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		if _, err := ajg.EncodeToValues(&test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm(b *testing.B) {
	test := getComplexArrayStruct()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := ajg.EncodeToValues(&test); err != nil {
				b.Error(err)
			}
		}
	})
}

// Complex Map ALL types

func getComplexMapStructValuesAGJForm() url.Values {
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

func BenchmarkComplexMapDecodeStructAllTypesAGJForm(b *testing.B) {
	values := getComplexMapStructValuesAGJForm()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test ComplexMapStruct
		if err := ajg.DecodeValues(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm(b *testing.B) {
	values := getComplexMapStructValuesAGJForm()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test ComplexMapStruct
			if err := ajg.DecodeValues(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkComplexMapEncodeStructAllTypesAGJForm(b *testing.B) {
	test := getComplexMapStructValuesAGJForm()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		if _, err := ajg.EncodeToValues(&test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm(b *testing.B) {
	test := getComplexMapStructValuesAGJForm()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := ajg.EncodeToValues(&test); err != nil {
				b.Error(err)
			}
		}
	})
}

// NestedStruct Benchmarks

func BenchmarkDecodeNestedStructAGJForm(b *testing.B) {

	values := getNestedStructValues()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var test NestedStruct
		if err := ajg.DecodeValues(&test, values); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDecodeNestedStructParallelAGJForm(b *testing.B) {

	values := getNestedStructValues()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var test NestedStruct
			if err := ajg.DecodeValues(&test, values); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkEncodeNestedStructAGJForm(b *testing.B) {

	test := getNestedStruct()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		if _, err := ajg.EncodeToValues(&test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEncodeNestedStructParallelAGJForm(b *testing.B) {

	test := getNestedStruct()

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := ajg.EncodeToValues(&test); err != nil {
				b.Error(err)
			}
		}
	})
}
