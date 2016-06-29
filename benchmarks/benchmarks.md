## Benchmarks

### go-playground/form
```go
BenchmarkSimpleUserDecodeStruct-8                                 	 5000000	       299 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                         	20000000	       106 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                                 	 2000000	       825 ns/op	     466 B/op	       7 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                         	 5000000	       301 ns/op	     466 B/op	       7 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8               	 2000000	       901 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8       	 5000000	       283 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8               	  300000	      4302 ns/op	    2912 B/op	      32 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8       	 1000000	      1542 ns/op	    2913 B/op	      32 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                       	  100000	     18125 ns/op	    2730 B/op	     135 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8               	  300000	      6178 ns/op	    2734 B/op	     135 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                       	  100000	     17738 ns/op	    7190 B/op	     154 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8               	  300000	      5360 ns/op	    7192 B/op	     154 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                         	   50000	     28476 ns/op	    7567 B/op	     163 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8                 	  200000	      9849 ns/op	    7604 B/op	     163 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                         	  100000	     18424 ns/op	    7097 B/op	     177 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8                 	  300000	      5852 ns/op	    7096 B/op	     177 allocs/op
BenchmarkDecodeNestedStruct-8                                     	  200000	      5199 ns/op	    1040 B/op	      31 allocs/op
BenchmarkDecodeNestedStructParallel-8                             	 1000000	      1713 ns/op	    1040 B/op	      31 allocs/op
BenchmarkEncodeNestedStruct-8                                     	  500000	      3025 ns/op	     848 B/op	      26 allocs/op
BenchmarkEncodeNestedStructParallel-8                             	 1000000	      1096 ns/op	     848 B/op	      26 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-8                          	  500000	      3063 ns/op	     520 B/op	      23 allocs/op
BenchmarkSimpleUserStructGorillaParallel-8                  	 1000000	      1026 ns/op	     520 B/op	      23 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-8        	  200000	     11136 ns/op	    1536 B/op	      84 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-8	  500000	      3928 ns/op	    1536 B/op	      84 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-8                	   50000	     34162 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-8        	  200000	     11937 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexMapStructAllTypesGorilla-8                  	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorilla-8
	gorilla_scheme_test.go:116: Gorilla does not support map parsing at this time
BenchmarkComplexMapStructAllTypesGorillaParallel-8          	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorillaParallel-8
	gorilla_scheme_test.go:121: Gorilla does not support map parsing at this time
BenchmarkArrayMapNestedStructGorilla-8                      	  200000	     10393 ns/op	    2269 B/op	      73 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-8              	  500000	      3484 ns/op	    2269 B/op	      73 allocs/op
No Encoder Support At This Time
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                                 	  500000	      2823 ns/op	     232 B/op	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                         	 2000000	       825 ns/op	     232 B/op	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8               	  200000	      9179 ns/op	    1088 B/op	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8       	  500000	      2949 ns/op	    1088 B/op	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                       	   30000	     41381 ns/op	    5668 B/op	     496 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8               	  100000	     14455 ns/op	    5705 B/op	     498 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                         	panic: reflect: reflect.Value.Set using unaddressable value
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-8                          	  200000	      6104 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-8                  	 1000000	      1748 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-8                          	  300000	      5076 ns/op	    1272 B/op	      29 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-8                  	 1000000	      1440 ns/op	    1272 B/op	      29 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-8        	  100000	     19713 ns/op	    5661 B/op	     143 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-8	  300000	      5618 ns/op	    5663 B/op	     143 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-8        	  100000	     15097 ns/op	    5792 B/op	      82 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-8	  300000	      4340 ns/op	    5792 B/op	      82 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8                	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8
	agj_form_test.go:127:  is not a valid index for type []uint16
BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8        	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8
	agj_form_test.go:140:  is not a valid index for type []int32
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-8                	   20000	     66822 ns/op	   21684 B/op	     400 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-8        	  100000	     20173 ns/op	   21681 B/op	     400 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-8                  	   20000	     86037 ns/op	   22295 B/op	     592 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-8          	   50000	     25491 ns/op	   22297 B/op	     592 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-8                  	   30000	     44735 ns/op	   17959 B/op	     323 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-8          	  100000	     13961 ns/op	   17958 B/op	     323 allocs/op
BenchmarkDecodeNestedStructAGJForm-8                              	--- FAIL: BenchmarkDecodeNestedStructAGJForm-8
	agj_form_test.go:261: NestedArray[1] doesn't exist in benchmarks.NestedStruct
BenchmarkDecodeNestedStructParallelAGJForm-8                      	--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm-8
	agj_form_test.go:275: NestedArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-8                              	  100000	     17904 ns/op	    5704 B/op	     113 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-8                      	  300000	      5502 ns/op	    5704 B/op	     113 allocs/op
```