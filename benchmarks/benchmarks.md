## Benchmarks

### go-playground/form
```go
PASS
BenchmarkSimpleUserDecodeStruct-8                          	 5000000	       308 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                  	20000000	        94.8 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                          	 2000000	       989 ns/op	     549 B/op	      12 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                  	 5000000	       332 ns/op	     549 B/op	      12 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8        	 1000000	      1004 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8	 5000000	       291 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8        	  300000	      4771 ns/op	    3073 B/op	      47 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8	 1000000	      1575 ns/op	    3073 B/op	      47 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                	  100000	     17087 ns/op	    2513 B/op	     123 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8        	  300000	      5020 ns/op	    2518 B/op	     123 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                	  100000	     16219 ns/op	    7350 B/op	     147 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8        	  300000	      4961 ns/op	    7351 B/op	     147 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                  	   50000	     24898 ns/op	    7088 B/op	     135 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8          	  200000	      7771 ns/op	    7121 B/op	     135 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                  	  100000	     16885 ns/op	    7159 B/op	     176 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8          	  300000	      5851 ns/op	    7161 B/op	     176 allocs/op
BenchmarkDecodeNestedStruct-8                              	  300000	      3848 ns/op	     640 B/op	      16 allocs/op
BenchmarkDecodeNestedStructParallel-8                      	 1000000	      1325 ns/op	     640 B/op	      16 allocs/op
BenchmarkEncodeNestedStruct-8                              	  500000	      2319 ns/op	     768 B/op	      17 allocs/op
BenchmarkEncodeNestedStructParallel-8                      	 2000000	       874 ns/op	     768 B/op	      17 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-8                                	  500000	      2997 ns/op	     520 B/op	      23 allocs/op
BenchmarkSimpleUserStructGorillaParallel-8                        	 2000000	       874 ns/op	     520 B/op	      23 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-8              	  200000	     11052 ns/op	    1536 B/op	      84 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-8      	  500000	      3479 ns/op	    1536 B/op	      84 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-8                      	   50000	     33887 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-8              	  200000	     10720 ns/op	    5416 B/op	     223 allocs/op
BenchmarkComplexMapStructAllTypesGorilla-8                        	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorilla-8
	gorilla_scheme_test.go:111: Gorilla does not support map parsing at this time
BenchmarkComplexMapStructAllTypesGorillaParallel-8                	       0	         0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkComplexMapStructAllTypesGorillaParallel-8
	gorilla_scheme_test.go:116: Gorilla does not support map parsing at this time
BenchmarkArrayMapNestedStructGorilla-8                            	  200000	     10156 ns/op	    2269 B/op	      73 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-8                    	  500000	      3306 ns/op	    2269 B/op	      73 allocs/op
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                                 	  500000	      2769 ns/op	     232 B/op	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                         	 2000000	       847 ns/op	     232 B/op	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8               	  200000	      8469 ns/op	    1088 B/op	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8       	  500000	      2452 ns/op	    1088 B/op	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                       	   50000	     37945 ns/op	    5584 B/op	     484 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8               	  200000	     12417 ns/op	    5614 B/op	     485 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                         	   30000	     43519 ns/op	   14644 B/op	     534 allocs/op
BenchmarkComplexMapStructAllTypesFormamParallel-8                 	  100000	     14447 ns/op	   14653 B/op	     534 allocs/op
BenchmarkArrayMapNestedStructFormam-8                             	--- FAIL: BenchmarkArrayMapNestedStructFormam-8
	formam_test.go:137: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
BenchmarkArrayMapNestedStructFormamParallel-8                     	--- FAIL: BenchmarkArrayMapNestedStructFormamParallel-8
	formam_test.go:151: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-8                          	  200000	      5960 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-8                  	 1000000	      1693 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-8                          	  300000	      4679 ns/op	    1272 B/op	      29 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-8                  	 1000000	      1379 ns/op	    1272 B/op	      29 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-8        	  100000	     19211 ns/op	    5662 B/op	     143 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-8	  300000	      5490 ns/op	    5662 B/op	     143 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-8        	  100000	     14807 ns/op	    5791 B/op	      82 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-8	  300000	      4355 ns/op	    5792 B/op	      82 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8                	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm-8
	ajg_form_test.go:127:  is not a valid index for type []*string
BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8        	--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm-8
	ajg_form_test.go:140: NestedInt[0][0] doesn't exist in benchmarks.ComplexArrayStruct
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-8                	   20000	     66264 ns/op	   21683 B/op	     400 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-8        	  100000	     20511 ns/op	   21679 B/op	     400 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-8                  	   20000	     83384 ns/op	   22295 B/op	     592 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-8          	   50000	     24943 ns/op	   22298 B/op	     592 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-8                  	   30000	     43976 ns/op	   17959 B/op	     323 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-8          	  100000	     13729 ns/op	   17959 B/op	     323 allocs/op
BenchmarkDecodeNestedStructAGJForm-8                              	--- FAIL: BenchmarkDecodeNestedStructAGJForm-8
	ajg_form_test.go:261: NestedArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkDecodeNestedStructParallelAGJForm-8                      	--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm-8
	ajg_form_test.go:275: NestedPtrArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-8                              	  100000	     17287 ns/op	    5704 B/op	     113 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-8                      	  300000	      5311 ns/op	    5704 B/op	     113 allocs/op
```