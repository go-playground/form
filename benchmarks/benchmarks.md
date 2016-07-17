## Benchmarks

### go-playground/form
```go
PASS
BenchmarkSimpleUserDecodeStruct-8                          	 5000000	       298 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                  	20000000	        91.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                          	 2000000	       966 ns/op	     549 B/op	      12 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                  	 5000000	       313 ns/op	     549 B/op	      12 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8        	 1000000	      1010 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8	 5000000	       285 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8        	  300000	      4718 ns/op	    3073 B/op	      47 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8	 1000000	      1673 ns/op	    3072 B/op	      47 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                	  100000	     16145 ns/op	    2289 B/op	     122 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8        	  300000	      4943 ns/op	    2291 B/op	     122 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                	  100000	     16020 ns/op	    7351 B/op	     147 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8        	  300000	      5129 ns/op	    7351 B/op	     147 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                  	  100000	     22933 ns/op	    5338 B/op	     131 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8          	  200000	      6366 ns/op	    5341 B/op	     131 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                  	  100000	     16861 ns/op	    7161 B/op	     176 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8          	  300000	      5159 ns/op	    7160 B/op	     176 allocs/op
BenchmarkDecodeNestedStruct-8                              	  500000	      3482 ns/op	     416 B/op	      15 allocs/op
BenchmarkDecodeNestedStructParallel-8                      	 2000000	      1011 ns/op	     416 B/op	      15 allocs/op
BenchmarkEncodeNestedStruct-8                              	 1000000	      2255 ns/op	     768 B/op	      17 allocs/op
BenchmarkEncodeNestedStructParallel-8                      	 2000000	       738 ns/op	     768 B/op	      17 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-8                                	  500000	      2974 ns/op	     520 B/op	      23 allocs/op
BenchmarkSimpleUserStructGorillaParallel-8                        	 2000000	      1003 ns/op	     520 B/op	      23 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-8              	  200000	     10855 ns/op	    1536 B/op	      84 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-8      	  500000	      3285 ns/op	    1536 B/op	      84 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-8                      	   50000	     33196 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-8              	  200000	     11022 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexMapStructAllTypesGorilla-8                        	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorilla-8
	gorilla_scheme_test.go:111: Gorilla does not support map parsing at this time
BenchmarkComplexMapStructAllTypesGorillaParallel-8                	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorillaParallel-8
	gorilla_scheme_test.go:116: Gorilla does not support map parsing at this time
BenchmarkArrayMapNestedStructGorilla-8                            	  200000	     10213 ns/op	    2285 B/op	      75 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-8                    	  500000	      3429 ns/op	    2285 B/op	      75 allocs/op
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                                 	  500000	      2753 ns/op	     232 B/op	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                         	 2000000	       811 ns/op	     232 B/op	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8               	  200000	      8442 ns/op	    1088 B/op	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8       	 1000000	      2613 ns/op	    1088 B/op	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                       	   50000	     37377 ns/op	    5539 B/op	     481 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8               	  200000	     13360 ns/op	    5564 B/op	     483 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                         	   30000	     42590 ns/op	   14654 B/op	     534 allocs/op
BenchmarkComplexMapStructAllTypesFormamParallel-8                 	  100000	     16032 ns/op	   14653 B/op	     534 allocs/op
BenchmarkArrayMapNestedStructFormam-8                             	--- FAIL: BenchmarkArrayMapNestedStructFormam-8
	formam_test.go:137: formam: not supported type for field "Value" in path "NestedPtrArray[1].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
BenchmarkArrayMapNestedStructFormamParallel-8                     	--- FAIL: BenchmarkArrayMapNestedStructFormamParallel-8
	formam_test.go:151: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-8                          	  200000	      7080 ns/op	    1336 B/op	      42 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-8                  	 1000000	      1999 ns/op	    1336 B/op	      42 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-8                          	  300000	      5580 ns/op	    1304 B/op	      37 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-8                  	 1000000	      1666 ns/op	    1304 B/op	      37 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-8        	  100000	     22895 ns/op	    5718 B/op	     171 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-8	  200000	      6736 ns/op	    5719 B/op	     171 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-8        	  100000	     17680 ns/op	    5903 B/op	     110 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-8	  300000	      5059 ns/op	    5904 B/op	     110 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8                	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8
	ajg_form_test.go:127:  is not a valid index for type []*int64
BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8        	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8
	ajg_form_test.go:140:  is not a valid index for type []*int
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-8                	   20000	     79749 ns/op	   22198 B/op	     538 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-8        	  100000	     22357 ns/op	   22199 B/op	     538 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-8                  	   20000	     94743 ns/op	   22497 B/op	     692 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-8          	   50000	     26379 ns/op	   22501 B/op	     692 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-8                  	   30000	     54343 ns/op	   18288 B/op	     419 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-8          	  100000	     16416 ns/op	   18291 B/op	     419 allocs/op
BenchmarkDecodeNestedStructAGJForm-8                              	--- FAIL: BenchmarkDecodeNestedStructAGJForm-8
	ajg_form_test.go:261: NestedArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkDecodeNestedStructParallelAGJForm-8                      	--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm-8
	ajg_form_test.go:275: NestedArray[1] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-8                              	  100000	     20759 ns/op	    5838 B/op	     147 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-8                      	  200000	      6277 ns/op	    5838 B/op	     147 allocs/op
```