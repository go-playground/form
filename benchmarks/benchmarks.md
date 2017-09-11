## Benchmarks

All Benchmarks Last Run Sept 10, 2017

Run on MacBook Pro (15-inch, 2017) using go version go1.9 darwin/amd64
go test -run=NONE -bench=. -benchmem=true

### go-playground/form
```go
BenchmarkSimpleUserDecodeStruct-8                             	 5000000	       243 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                     	20000000	        72.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                             	 2000000	       683 ns/op	     485 B/op	      10 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                     	10000000	       205 ns/op	     485 B/op	      10 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8           	 2000000	       739 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8   	10000000	       214 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8           	  500000	      3608 ns/op	    2977 B/op	      36 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8   	 1000000	      1013 ns/op	    2978 B/op	      36 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                   	  100000	     13358 ns/op	    2249 B/op	     121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8           	  500000	      3656 ns/op	    2249 B/op	     121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                   	  100000	     11574 ns/op	    7112 B/op	     104 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8           	  500000	      3278 ns/op	    7112 B/op	     104 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                     	  100000	     18945 ns/op	    5305 B/op	     130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8             	  300000	      5213 ns/op	    5308 B/op	     130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                     	  100000	     12177 ns/op	    6972 B/op	     129 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8             	  500000	      3409 ns/op	    6971 B/op	     129 allocs/op
BenchmarkDecodeNestedStruct-8                                 	  500000	      2642 ns/op	     384 B/op	      14 allocs/op
BenchmarkDecodeNestedStructParallel-8                         	 2000000	       760 ns/op	     384 B/op	      14 allocs/op
BenchmarkEncodeNestedStruct-8                                 	 1000000	      1686 ns/op	     693 B/op	      16 allocs/op
BenchmarkEncodeNestedStructParallel-8                         	 3000000	       468 ns/op	     693 B/op	      16 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-8                                   	  500000	      2602 ns/op	     568 B/op	      26 allocs/op
BenchmarkSimpleUserStructGorillaParallel-8                           	 2000000	       742 ns/op	     568 B/op	      26 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-8                 	  200000	      9801 ns/op	    1616 B/op	      95 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-8         	  500000	      2674 ns/op	    1616 B/op	      95 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-8                         	   50000	     27079 ns/op	    5528 B/op	     240 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-8                 	  200000	      7382 ns/op	    5528 B/op	     240 allocs/op
BenchmarkArrayMapNestedStructGorilla-8                               	  200000	      8578 ns/op	    2397 B/op	      82 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-8                       	 1000000	      2392 ns/op	    2397 B/op	      82 allocs/op
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                                    	  500000	      2657 ns/op	     232 B/op	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                            	 2000000	       700 ns/op	     232 B/op	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8                  	  200000	      8657 ns/op	    1088 B/op	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8          	 1000000	      2168 ns/op	    1088 B/op	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                          	   50000	     35892 ns/op	    5588 B/op	     484 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8                  	  200000	      9238 ns/op	    5554 B/op	     482 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                            	   30000	     40312 ns/op	   14645 B/op	     534 allocs/op
BenchmarkComplexMapStructAllTypesFormamParallel-8                    	  200000	     10826 ns/op	   14646 B/op	     534 allocs/op
--- FAIL: BenchmarkArrayMapNestedStructFormam
	formam_test.go:137: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
--- FAIL: BenchmarkArrayMapNestedStructFormamParallel
	formam_test.go:151: formam: not supported type for field "Value" in path "NestedPtrArray[1].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-8                             	  300000	      5134 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-8                     	 1000000	      1379 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-8                             	  300000	      4245 ns/op	    1272 B/op	      28 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-8                     	 1000000	      1144 ns/op	    1272 B/op	      28 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-8           	  100000	     19222 ns/op	    5662 B/op	     143 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-8   	  300000	      4966 ns/op	    5662 B/op	     143 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-8           	  100000	     13655 ns/op	    5761 B/op	      72 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-8   	  500000	      3617 ns/op	    5760 B/op	      72 allocs/op
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm
	ajg_form_test.go:127:  is not a valid index for type []uint16
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm
	ajg_form_test.go:140:  is not a valid index for type []int
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-8                   	   30000	     57848 ns/op	   21210 B/op	     314 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-8           	  100000	     15527 ns/op	   21207 B/op	     314 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-8                     	   20000	     77528 ns/op	   22295 B/op	     592 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-8             	  100000	     20717 ns/op	   22297 B/op	     592 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-8                     	   50000	     39453 ns/op	   17936 B/op	     299 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-8             	  200000	     11513 ns/op	   17937 B/op	     299 allocs/op
--- FAIL: BenchmarkDecodeNestedStructAGJForm
	ajg_form_test.go:261: NestedArray[0] doesn't exist in benchmarks.NestedStruct
--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm
	ajg_form_test.go:275: NestedArray[1] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-8                                 	  100000	     15791 ns/op	    5656 B/op	     108 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-8                         	  300000	      4625 ns/op	    5656 B/op	     108 allocs/op
```