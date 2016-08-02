## Benchmarks

Competitor Benchmarks Last Run July 26, 2016

### go-playground/form
```go
PASS
BenchmarkSimpleUserDecodeStruct-8                          	 5000000	       336 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                  	20000000	        99.5 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                          	 2000000	       972 ns/op	     485 B/op	      11 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                  	 5000000	       329 ns/op	     485 B/op	      11 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8        	 1000000	      1014 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8	 5000000	       294 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8        	  300000	      4799 ns/op	    3009 B/op	      46 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8	 1000000	      1581 ns/op	    3010 B/op	      46 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                	  100000	     16326 ns/op	    2257 B/op	     121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8        	  300000	      4710 ns/op	    2257 B/op	     121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                	  100000	     16303 ns/op	    7288 B/op	     146 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8        	  300000	      4979 ns/op	    7290 B/op	     146 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                  	  100000	     21998 ns/op	    5306 B/op	     130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8          	  200000	      6542 ns/op	    5308 B/op	     130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                  	  100000	     17069 ns/op	    7100 B/op	     175 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8          	  300000	      5609 ns/op	    7099 B/op	     175 allocs/op
BenchmarkDecodeNestedStruct-8                              	  500000	      3366 ns/op	     384 B/op	      14 allocs/op
BenchmarkDecodeNestedStructParallel-8                      	 2000000	      1096 ns/op	     384 B/op	      14 allocs/op
BenchmarkEncodeNestedStruct-8                              	 1000000	      2230 ns/op	     704 B/op	      16 allocs/op
BenchmarkEncodeNestedStructParallel-8                      	 2000000	       780 ns/op	     704 B/op	      16 allocs/op
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