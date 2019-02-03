## Benchmarks

All Benchmarks Last Run Feb 2, 2019

Run on MacBook Pro (15-inch, 2017) using go version go1.11.5 darwin/amd64
go test -run=NONE -bench=. -benchmem=true

### go-playground/form
```go
BenchmarkSimpleUserDecodeStruct-8                                5000000               255 ns/op              64 B/op          1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                       20000000                77.1 ns/op            64 B/op          1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                                2000000               625 ns/op             485 B/op         10 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                       10000000               207 ns/op             485 B/op         10 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8              2000000               768 ns/op              96 B/op          1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8     10000000               221 ns/op              96 B/op          1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8               500000              3192 ns/op            2977 B/op         35 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8      1000000              1011 ns/op            2977 B/op         35 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                       100000             13269 ns/op            2248 B/op        121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8               500000              3758 ns/op            2249 B/op        121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                       200000             10728 ns/op            7112 B/op        104 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8               500000              3305 ns/op            7113 B/op        104 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                         100000             17331 ns/op            5306 B/op        130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8                 300000              4901 ns/op            5309 B/op        130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                         100000             11226 ns/op            6970 B/op        129 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8                 500000              3418 ns/op            6970 B/op        129 allocs/op
BenchmarkDecodeNestedStruct-8                                     500000              2448 ns/op             384 B/op         14 allocs/op
BenchmarkDecodeNestedStructParallel-8                            2000000               736 ns/op             384 B/op         14 allocs/op
BenchmarkEncodeNestedStruct-8                                    1000000              1468 ns/op             693 B/op         16 allocs/op
BenchmarkEncodeNestedStructParallel-8                            3000000               467 ns/op             693 B/op         16 allocs/op
```