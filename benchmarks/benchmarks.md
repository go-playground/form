## Benchmarks

Last Run July 26, 2016

### go-playground/form
```go
PASS
BenchmarkSimpleUserDecodeStruct-8                          	 5000000	       318 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                  	20000000	        95.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                          	 1000000	      1000 ns/op	     549 B/op	      12 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                  	 5000000	       325 ns/op	     549 B/op	      12 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8        	 1000000	      1058 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8	 5000000	       324 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8        	  300000	      4823 ns/op	    3073 B/op	      47 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8	 1000000	      1732 ns/op	    3072 B/op	      47 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                	  100000	     16340 ns/op	    2289 B/op	     122 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8        	  300000	      5105 ns/op	    2291 B/op	     122 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                	  100000	     16343 ns/op	    7351 B/op	     147 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8        	  300000	      5969 ns/op	    7351 B/op	     147 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                  	  100000	     21259 ns/op	    5338 B/op	     131 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8          	  200000	      7493 ns/op	    5342 B/op	     131 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                  	  100000	     17060 ns/op	    7161 B/op	     176 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8          	  300000	      6421 ns/op	    7161 B/op	     176 allocs/op
BenchmarkDecodeNestedStruct-8                              	  300000	      3488 ns/op	     416 B/op	      15 allocs/op
BenchmarkDecodeNestedStructParallel-8                      	 1000000	      1203 ns/op	     416 B/op	      15 allocs/op
BenchmarkEncodeNestedStruct-8                              	 1000000	      2286 ns/op	     768 B/op	      17 allocs/op
BenchmarkEncodeNestedStructParallel-8                      	 2000000	       939 ns/op	     768 B/op	      17 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-8                                	  500000	      3205 ns/op	     520 B/op	      23 allocs/op
BenchmarkSimpleUserStructGorillaParallel-8                        	 2000000	      1072 ns/op	     520 B/op	      23 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-8              	  200000	     11491 ns/op	    1536 B/op	      84 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-8      	  500000	      3853 ns/op	    1536 B/op	      84 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-8                      	   50000	     33257 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-8              	  200000	     10940 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexMapStructAllTypesGorilla-8                        	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorilla-8
	gorilla_scheme_test.go:111: Gorilla does not support map parsing at this time
BenchmarkComplexMapStructAllTypesGorillaParallel-8                	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorillaParallel-8
	gorilla_scheme_test.go:116: Gorilla does not support map parsing at this time
BenchmarkArrayMapNestedStructGorilla-8                            	  200000	     10387 ns/op	    2285 B/op	      75 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-8                    	  500000	      3492 ns/op	    2285 B/op	      75 allocs/op
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                                 	  500000	      2839 ns/op	     232 B/op	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                         	 2000000	       919 ns/op	     232 B/op	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8               	  200000	      8605 ns/op	    1088 B/op	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8       	 1000000	      2906 ns/op	    1088 B/op	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                       	   50000	     38928 ns/op	    5606 B/op	     485 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8               	  100000	     13110 ns/op	    5600 B/op	     485 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                         	   30000	     44037 ns/op	   14648 B/op	     534 allocs/op
BenchmarkComplexMapStructAllTypesFormamParallel-8                 	  100000	     15840 ns/op	   14660 B/op	     534 allocs/op
BenchmarkArrayMapNestedStructFormam-8                             	--- FAIL: BenchmarkArrayMapNestedStructFormam-8
	formam_test.go:137: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
BenchmarkArrayMapNestedStructFormamParallel-8                     	--- FAIL: BenchmarkArrayMapNestedStructFormamParallel-8
	formam_test.go:151: formam: not supported type for field "Value" in path "NestedPtrArray[1].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-8                          	  200000	      7159 ns/op	    1336 B/op	      42 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-8                  	 1000000	      1950 ns/op	    1336 B/op	      42 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-8                          	  300000	      5731 ns/op	    1304 B/op	      37 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-8                  	 1000000	      1711 ns/op	    1304 B/op	      37 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-8        	   50000	     23457 ns/op	    5720 B/op	     171 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-8	  200000	      7120 ns/op	    5718 B/op	     171 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-8        	  100000	     18299 ns/op	    5903 B/op	     110 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-8	  300000	      5973 ns/op	    5904 B/op	     110 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8                	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8
	ajg_form_test.go:127: Int8Ptr[1] doesn't exist in benchmarks.ComplexArrayStruct
BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8        	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8
	ajg_form_test.go:140:  is not a valid index for type []*uint
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-8                	   20000	     81649 ns/op	   22206 B/op	     538 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-8        	  100000	     27927 ns/op	   22197 B/op	     538 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-8                  	   20000	     94935 ns/op	   22500 B/op	     692 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-8          	   50000	     34001 ns/op	   22499 B/op	     692 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-8                  	   30000	     54767 ns/op	   18288 B/op	     419 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-8          	  100000	     21300 ns/op	   18285 B/op	     419 allocs/op
BenchmarkDecodeNestedStructAGJForm-8                              	--- FAIL: BenchmarkDecodeNestedStructAGJForm-8
	ajg_form_test.go:261: NestedArray[1] doesn't exist in benchmarks.NestedStruct
BenchmarkDecodeNestedStructParallelAGJForm-8                      	--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm-8
	ajg_form_test.go:275: NestedArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-8                              	   50000	     23387 ns/op	    5838 B/op	     147 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-8                      	  200000	      7320 ns/op	    5838 B/op	     147 allocs/op
```