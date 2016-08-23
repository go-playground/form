## Benchmarks

Competitor Benchmarks Last Run Aug 23, 2016

### go-playground/form
```go
PASS
BenchmarkSimpleUserDecodeStruct-8                              	 5000000       	       312 ns/op       	      64 B/op  	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                      	20000000       	        91.7 ns/op     	      64 B/op  	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                              	 2000000       	       902 ns/op       	     485 B/op  	      11 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                      	 5000000       	       301 ns/op       	     485 B/op  	      11 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8            	 2000000       	      1028 ns/op       	      96 B/op  	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8    	 5000000       	       292 ns/op       	      96 B/op  	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8            	  300000       	      4770 ns/op       	    3009 B/op  	      46 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8    	 1000000       	      1569 ns/op       	    3010 B/op  	      46 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                    	  100000       	     15973 ns/op       	    2257 B/op  	     121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8            	  300000       	      4801 ns/op       	    2257 B/op  	     121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                    	  100000       	     15401 ns/op       	    7289 B/op  	     146 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8            	  300000       	      5167 ns/op       	    7289 B/op  	     146 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                      	   50000       	     20683 ns/op       	    5307 B/op  	     130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8              	  300000       	      6880 ns/op       	    5310 B/op  	     130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                      	  100000       	     15567 ns/op       	    7098 B/op  	     175 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8              	  300000       	      5546 ns/op       	    7099 B/op  	     175 allocs/op
BenchmarkDecodeNestedStruct-8                                  	  500000       	      3142 ns/op       	     384 B/op  	      14 allocs/op
BenchmarkDecodeNestedStructParallel-8                          	 1000000       	      1012 ns/op       	     384 B/op  	      14 allocs/op
BenchmarkEncodeNestedStruct-8                                  	 1000000       	      2106 ns/op       	     704 B/op  	      16 allocs/op
BenchmarkEncodeNestedStructParallel-8                          	 2000000       	       772 ns/op       	     704 B/op  	      16 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-8                                     	  500000       	      2412 ns/op       	     520 B/op  	      23 allocs/op
BenchmarkSimpleUserStructGorillaParallel-8                             	 2000000       	       900 ns/op       	     520 B/op  	      23 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-8                   	  200000       	      9495 ns/op       	    1536 B/op  	      84 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-8           	  500000       	      3018 ns/op       	    1536 B/op  	      84 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-8                           	   50000       	     29454 ns/op       	    5416 B/op  	     223 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-8                   	  200000       	      9406 ns/op       	    5416 B/op  	     223 allocs/op
BenchmarkArrayMapNestedStructGorilla-8                                 	  200000       	     10212 ns/op       	    2285 B/op  	      75 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-8                         	  500000       	      3305 ns/op       	    2285 B/op  	      75 allocs/op
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                                      	  500000       	      2311 ns/op       	     232 B/op  	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                              	 2000000       	       793 ns/op       	     232 B/op  	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8                    	  200000       	      8539 ns/op       	    1088 B/op  	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8            	 1000000       	      2876 ns/op       	    1088 B/op  	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                            	   50000       	     37409 ns/op       	    5608 B/op  	     485 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8                    	  200000       	     12292 ns/op       	    5556 B/op  	     482 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                              	   30000       	     41450 ns/op       	   14647 B/op  	     534 allocs/op
BenchmarkComplexMapStructAllTypesFormamParallel-8                      	  100000       	     14764 ns/op       	   14647 B/op  	     534 allocs/op
--- FAIL: BenchmarkArrayMapNestedStructFormam
       	formam_test.go:137: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
--- FAIL: BenchmarkArrayMapNestedStructFormamParallel
       	formam_test.go:151: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-8                               	  200000       	      6201 ns/op       	    1336 B/op  	      42 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-8                       	 1000000       	      1728 ns/op       	    1336 B/op  	      42 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-8                               	  300000       	      4969 ns/op       	    1304 B/op  	      37 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-8                       	 1000000       	      1431 ns/op       	    1304 B/op  	      37 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-8             	  100000       	     21879 ns/op       	    5718 B/op  	     171 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-8     	  300000       	      6298 ns/op       	    5718 B/op  	     171 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-8             	  100000       	     17249 ns/op       	    5904 B/op  	     110 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-8     	  300000       	      5032 ns/op       	    5904 B/op  	     110 allocs/op
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm
       	ajg_form_test.go:127:  is not a valid index for type []int16
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm
       	ajg_form_test.go:140:  is not a valid index for type []*string
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-8                     	   20000       	     75231 ns/op       	   22195 B/op  	     538 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-8             	  100000       	     25228 ns/op       	   22199 B/op  	     538 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-8                       	   10000       	    100451 ns/op       	   22496 B/op  	     692 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-8               	   50000       	     29669 ns/op       	   22496 B/op  	     692 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-8                       	   30000       	     50423 ns/op       	   18288 B/op  	     419 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-8               	  100000       	     16833 ns/op       	   18289 B/op  	     419 allocs/op
--- FAIL: BenchmarkDecodeNestedStructAGJForm
       	ajg_form_test.go:261: NestedArray[0] doesn't exist in benchmarks.NestedStruct
--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm
       	ajg_form_test.go:275: NestedPtrArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-8                                   	  100000       	     18447 ns/op       	    5838 B/op  	     147 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-8                           	  300000       	      6898 ns/op       	    5838 B/op  	     147 allocs/op
```