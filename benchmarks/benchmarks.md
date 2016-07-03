## Benchmarks

### go-playground/form
```go
BenchmarkSimpleUserDecodeStruct-8                                 	 5000000	       290 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                         	20000000	        99.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                                 	 2000000	       806 ns/op	     466 B/op	       7 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                         	 5000000	       297 ns/op	     466 B/op	       7 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8               	 2000000	       911 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8       	 5000000	       252 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8               	  300000	      4293 ns/op	    2912 B/op	      32 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8       	 1000000	      1451 ns/op	    2913 B/op	      32 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                       	  100000	     17295 ns/op	    2682 B/op	     135 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8               	  300000	      5112 ns/op	    2686 B/op	     135 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                       	  100000	     17128 ns/op	    7191 B/op	     154 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8               	  300000	      5520 ns/op	    7192 B/op	     154 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                         	   50000	     26311 ns/op	    7474 B/op	     161 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8                 	  200000	      8539 ns/op	    7504 B/op	     161 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                         	  100000	     18627 ns/op	    7095 B/op	     177 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8                 	  300000	      5922 ns/op	    7098 B/op	     177 allocs/op
BenchmarkDecodeNestedStruct-8                                     	  300000	      4815 ns/op	    1040 B/op	      31 allocs/op
BenchmarkDecodeNestedStructParallel-8                             	 1000000	      1613 ns/op	    1040 B/op	      31 allocs/op
BenchmarkEncodeNestedStruct-8                                     	  500000	      2995 ns/op	     848 B/op	      26 allocs/op
BenchmarkEncodeNestedStructParallel-8                             	 1000000	      1031 ns/op	     848 B/op	      26 allocs/op
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