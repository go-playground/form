## Benchmarks

All Benchmarks Last Run July 30, 2017

Run on Dell XPS 15 i7-7700HQ 32GB using Go version go1.8.3 linux/amd64
go test -run=NONE -bench=. -benchmem=true

### go-playground/form
```go
BenchmarkSimpleUserDecodeStruct-8                                    	 5000000	       264 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserDecodeStructParallel-8                            	20000000	        80.7 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserEncodeStruct-8                                    	 1000000	      1097 ns/op	     485 B/op	      11 allocs/op
BenchmarkSimpleUserEncodeStructParallel-8                            	 5000000	       239 ns/op	     485 B/op	      11 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypes-8                  	 2000000	       799 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallel-8          	 5000000	       237 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypes-8                  	  300000	      6672 ns/op	    3010 B/op	      46 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallel-8          	 1000000	      1207 ns/op	    3010 B/op	      46 allocs/op
BenchmarkComplexArrayDecodeStructAllTypes-8                          	  100000	     14374 ns/op	    2249 B/op	     121 allocs/op
BenchmarkComplexArrayDecodeStructAllTypesParallel-8                  	  300000	      3867 ns/op	    2249 B/op	     121 allocs/op
BenchmarkComplexArrayEncodeStructAllTypes-8                          	  100000	     20531 ns/op	    7161 B/op	     146 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallel-8                  	  300000	      3839 ns/op	    7162 B/op	     146 allocs/op
BenchmarkComplexMapDecodeStructAllTypes-8                            	  100000	     20745 ns/op	    5306 B/op	     130 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallel-8                    	  300000	      5164 ns/op	    5312 B/op	     130 allocs/op
BenchmarkComplexMapEncodeStructAllTypes-8                            	  100000	     18658 ns/op	    7066 B/op	     175 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallel-8                    	  300000	      4019 ns/op	    7068 B/op	     175 allocs/op
BenchmarkDecodeNestedStruct-8                                        	  500000	      3039 ns/op	     384 B/op	      14 allocs/op
BenchmarkDecodeNestedStructParallel-8                                	 2000000	       832 ns/op	     384 B/op	      14 allocs/op
BenchmarkEncodeNestedStruct-8                                        	 1000000	      2005 ns/op	     693 B/op	      16 allocs/op
BenchmarkEncodeNestedStructParallel-8                                	 3000000	       534 ns/op	     693 B/op	      16 allocs/op
```

### gorilla/schema
```go
BenchmarkSimpleUserStructGorilla-8                                   	  500000	      2968 ns/op	     568 B/op	      27 allocs/op
BenchmarkSimpleUserStructGorillaParallel-8                           	 2000000	       798 ns/op	     568 B/op	      27 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorilla-8                 	  200000	     10666 ns/op	    1616 B/op	      98 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesGorillaParallel-8         	  500000	      2814 ns/op	    1616 B/op	      98 allocs/op
BenchmarkComplexArrayStructAllTypesGorilla-8                         	   50000	     29731 ns/op	    5528 B/op	     240 allocs/op
BenchmarkComplexArrayStructAllTypesGorillaParallel-8                 	  200000	      7657 ns/op	    5528 B/op	     240 allocs/op
BenchmarkArrayMapNestedStructGorilla-8                               	  200000	      9546 ns/op	    2397 B/op	      82 allocs/op
BenchmarkArrayMapNestedStructGorillaParallel-8                       	  500000	      2623 ns/op	    2397 B/op	      82 allocs/op
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                                    	  500000	      2888 ns/op	     232 B/op	      16 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                            	 2000000	       766 ns/op	     232 B/op	      16 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8                  	  200000	      8179 ns/op	    1088 B/op	     121 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8          	 1000000	      2235 ns/op	    1088 B/op	     121 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                          	   50000	     36620 ns/op	    5561 B/op	     482 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8                  	  200000	      9460 ns/op	    5560 B/op	     482 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                            	   30000	     43515 ns/op	   14649 B/op	     534 allocs/op
BenchmarkComplexMapStructAllTypesFormamParallel-8                    	  200000	     10842 ns/op	   14652 B/op	     534 allocs/op
--- FAIL: BenchmarkArrayMapNestedStructFormam
	formam_test.go:137: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
--- FAIL: BenchmarkArrayMapNestedStructFormamParallel
	formam_test.go:151: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value". Maybe you should to include it the UnmarshalText interface or register it using custom type?
```

### ajg/form
```go
BenchmarkSimpleUserDecodeStructAGJForm-8                             	  300000	      5567 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserDecodeStructParallelAGJFrom-8                     	 1000000	      1482 ns/op	    1320 B/op	      34 allocs/op
BenchmarkSimpleUserEncodeStructAGJForm-8                             	  300000	      4699 ns/op	    1272 B/op	      29 allocs/op
BenchmarkSimpleUserEncodeStructParallelAGJForm-8                     	 1000000	      1239 ns/op	    1272 B/op	      29 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesAGJForm-8           	  100000	     20476 ns/op	    5662 B/op	     143 allocs/op
BenchmarkPrimitivesDecodeStructAllPrimitivesTypesParallelAGJForm-8   	  300000	      5039 ns/op	    5662 B/op	     143 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesAGJForm-8           	  100000	     15661 ns/op	    5792 B/op	      82 allocs/op
BenchmarkPrimitivesEncodeStructAllPrimitivesTypesParallelAGJForm-8   	  300000	      3851 ns/op	    5792 B/op	      82 allocs/op
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesAGJForm
	ajg_form_test.go:127:  is not a valid index for type []int32
--- FAIL: BenchmarkComplexArrayDecodeStructAllTypesParallelAGJForm
	ajg_form_test.go:140: NestedInt[0][0] doesn't exist in benchmarks.ComplexArrayStruct
BenchmarkComplexArrayEncodeStructAllTypesAGJForm-8                   	   20000	     67813 ns/op	   21553 B/op	     400 allocs/op
BenchmarkComplexArrayEncodeStructAllTypesParallelAGJForm-8           	  100000	     17175 ns/op	   21552 B/op	     400 allocs/op
BenchmarkComplexMapDecodeStructAllTypesAGJForm-8                     	   20000	     84436 ns/op	   22294 B/op	     592 allocs/op
BenchmarkComplexMapDecodeStructAllTypesParallelAGJForm-8             	  100000	     21599 ns/op	   22298 B/op	     592 allocs/op
BenchmarkComplexMapEncodeStructAllTypesAGJForm-8                     	   30000	     49011 ns/op	   17958 B/op	     323 allocs/op
BenchmarkComplexMapEncodeStructAllTypesParallelAGJForm-8             	  200000	     11763 ns/op	   17958 B/op	     323 allocs/op
--- FAIL: BenchmarkDecodeNestedStructAGJForm
	ajg_form_test.go:261: NestedArray[0] doesn't exist in benchmarks.NestedStruct
--- FAIL: BenchmarkDecodeNestedStructParallelAGJForm
	ajg_form_test.go:275: NestedArray[0] doesn't exist in benchmarks.NestedStruct
BenchmarkEncodeNestedStructAGJForm-8                                 	  100000	     18583 ns/op	    5704 B/op	     113 allocs/op
BenchmarkEncodeNestedStructParallelAGJForm-8                         	  300000	      4683 ns/op	    5704 B/op	     113 allocs/op
```