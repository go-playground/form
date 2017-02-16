## Benchmarks

Competitor Benchmarks Last Run Feb 15, 2017

Run on i5-7600 16 GB DDR4-2400 using Go version go1.7.5 linux/amd64

### go-playground/form
```go
PASS
BenchmarkSimpleUserDecodeStruct-4                                    	 5000000	       252 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-4                            	20000000	        74.1 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-4                                    	 2000000	       800 ns/op	     485 B/op	      11 allocs/op
BenchmarkSimpleUserEncodeStructParallel-4                            	10000000	       220 ns/op	     485 B/op	      11 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-4                  	 2000000	       773 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-4          	10000000	       225 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-4                  	  300000	      4330 ns/op	    3009 B/op	      46 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-4          	 1000000	      1131 ns/op	    3009 B/op	      46 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-4                          	  100000	     13316 ns/op	    2256 B/op	     121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-4                  	  300000	      3980 ns/op	    2256 B/op	     121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-4                          	  100000	     14038 ns/op	    7287 B/op	     146 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-4                  	  300000	      3646 ns/op	    7288 B/op	     146 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-4                            	  100000	     18042 ns/op	    5305 B/op	     130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-4                    	  300000	      5051 ns/op	    5306 B/op	     130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-4                            	  100000	     14177 ns/op	    7098 B/op	     175 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-4                    	  300000	      3693 ns/op	    7098 B/op	     175 allocs/op
BenchmarkDecodeNestedStruct-4                                        	  500000	      2762 ns/op	     384 B/op	      14 allocs/op
BenchmarkDecodeNestedStructParallel-4                                	 2000000	       785 ns/op	     384 B/op	      14 allocs/op
BenchmarkEncodeNestedStruct-4                                        	 1000000	      1779 ns/op	     704 B/op	      16 allocs/op
BenchmarkEncodeNestedStructParallel-4                                	 3000000	       493 ns/op	     704 B/op	      16 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-4                                   	 1000000	      2077 ns/op	     520 B/op	      23 allocs/op
BenchmarkSimpleUserStructGorillaParallel-4                           	 2000000	       667 ns/op	     520 B/op	      23 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-4                 	  200000	      7997 ns/op	    1536 B/op	      84 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-4         	  500000	      2429 ns/op	    1536 B/op	      84 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-4                         	   50000	     23286 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-4                 	  200000	      6773 ns/op	    5416 B/op	     223 allocs/op
BenchmarkArrayMapNestedStructGorilla-4                               	  200000	      7745 ns/op	    2285 B/op	      75 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-4                       	  500000	      2421 ns/op	    2285 B/op	      75 allocs/op
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-4                                    	 1000000	      1927 ns/op	     232 B/op	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-4                            	 3000000	       514 ns/op	     232 B/op	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-4                  	  200000	      7135 ns/op	    1088 B/op	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-4          	 1000000	      1870 ns/op	    1088 B/op	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-4                          	   50000	     30957 ns/op	    5608 B/op	     485 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-4                  	  200000	      8848 ns/op	    5594 B/op	     484 allocs/op
BenchmarkComplexMapStructAllTypesFormam-4                            	   50000	     37984 ns/op	   14647 B/op	     534 allocs/op
BenchmarkComplexMapStructAllTypesFormamParallel-4                    	  200000	      9795 ns/op	   14658 B/op	     534 allocs/op
--- FAIL: BenchmarkArrayMapNestedStructFormam
	formam_test.go:137: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
--- FAIL: BenchmarkArrayMapNestedStructFormamParallel
	formam_test.go:151: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-4                             	  300000	      5150 ns/op	    1336 B/op	      42 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-4                     	 1000000	      1372 ns/op	    1336 B/op	      42 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-4                             	  300000	      4196 ns/op	    1304 B/op	      37 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-4                     	 1000000	      1129 ns/op	    1304 B/op	      37 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-4           	  100000	     18605 ns/op	    5718 B/op	     171 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-4   	  300000	      4941 ns/op	    5718 B/op	     171 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-4           	  100000	     14417 ns/op	    5903 B/op	     110 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-4   	  500000	      3702 ns/op	    5904 B/op	     110 allocs/op
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm
	ajg_form_test.go:127:  is not a valid index for type []uint32
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm
	ajg_form_test.go:140: Int8Ptr[1] doesn't exist in benchmarks.ComplexArrayStruct
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-4                   	   30000	     61548 ns/op	   22195 B/op	     538 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-4           	  100000	     16386 ns/op	   22193 B/op	     538 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-4                     	   20000	     76838 ns/op	   22498 B/op	     692 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-4             	  100000	     20941 ns/op	   22496 B/op	     692 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-4                     	   50000	     40160 ns/op	   18286 B/op	     419 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-4             	  200000	     10761 ns/op	   18289 B/op	     419 allocs/op
--- FAIL: BenchmarkDecodeNestedStructAGJForm
	ajg_form_test.go:261: NestedArray[0] doesn't exist in benchmarks.NestedStruct
--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm
	ajg_form_test.go:275: NestedArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-4                                 	  100000	     15989 ns/op	    5838 B/op	     147 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-4                         	  300000	      4272 ns/op	    5838 B/op	     147 allocs/op
```