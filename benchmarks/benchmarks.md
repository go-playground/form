## Benchmarks

### go-playground/form
```go
BenchmarkSimpleUserStruct-8                                 	 5000000	       299 ns/op	      64 B/op	       1 allocs/op
BenchmarkSimpleUserStructParallel-8                         	20000000	       110 ns/op	      64 B/op	       1 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypes-8               	 2000000	       956 ns/op	      96 B/op	       1 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesParallel-8       	 5000000	       285 ns/op	      96 B/op	       1 allocs/op
BenchmarkComplexArrayStructAllTypes-8                       	  100000	     20706 ns/op	    6776 B/op	     159 allocs/op
BenchmarkComplexArrayStructAllTypesParallel-8               	  200000	      6158 ns/op	    6776 B/op	     159 allocs/op
BenchmarkComplexMapStructAllTypes-8                         	   50000	     35548 ns/op	   20966 B/op	     245 allocs/op
BenchmarkComplexMapStructAllTypesParallel-8                 	  200000	     11984 ns/op	   20966 B/op	     245 allocs/op
BenchmarkArrayMapNestedStruct-8                             	  200000	      5617 ns/op	    2064 B/op	      37 allocs/op
BenchmarkArrayMapNestedStructParallel-8                     	 1000000	      2032 ns/op	    2064 B/op	      37 allocs/op
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
```

### monoculum/formam
```go
BenchmarkSimpleUserStructFormam-8                           	  500000	      3713 ns/op	     264 B/op	      19 allocs/op
BenchmarkSimpleUserStructFormamParallel-8                   	 1000000	      1017 ns/op	     264 B/op	      19 allocs/op
BenchmarkPrimitivesStructAllPrimitivesFormamTypes-8         	  100000	     12197 ns/op	    1280 B/op	     134 allocs/op
BenchmarkPrimitivesStructAllPrimitivesTypesFormamParallel-8 	  500000	      3905 ns/op	    1280 B/op	     134 allocs/op
BenchmarkComplexArrayStructAllTypesFormam-8                 	   30000	     58650 ns/op	    6371 B/op	     522 allocs/op
BenchmarkComplexArrayStructAllTypesFormamParallel-8         	  100000	     19046 ns/op	    6349 B/op	     521 allocs/op
BenchmarkComplexMapStructAllTypesFormam-8                   	--- FAIL: BenchmarkComplexMapStructAllTypesFormam-8
	formam_test.go:142: formam: the key with uint16 type (map[uint16]uint16) in the path Uint16.0 should implements the TextUnmarshaler interface for to can decode it
BenchmarkComplexMapStructAllTypesFormamParallel-8           	--- FAIL: BenchmarkComplexMapStructAllTypesFormamParallel-8
	formam_test.go:158: formam: the key with int type (map[int]int) in the path Int.0 should implements the TextUnmarshaler interface for to can decode it
BenchmarkArrayMapNestedStructFormam-8                       	--- FAIL: BenchmarkArrayMapNestedStructFormam-8
	formam_test.go:174: formam: not supported type for field "Value" in path "NestedPtrArray[1].Value"
BenchmarkArrayMapNestedStructFormamParallel-8               	--- FAIL: BenchmarkArrayMapNestedStructFormamParallel-8
	formam_test.go:189: formam: not supported type for field "Value" in path "NestedPtrArray[0].Value"
```