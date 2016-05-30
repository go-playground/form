package form

import (
	"net/url"
	"testing"

	. "github.com/go-playground/assert"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//
//
// go test -cpuprofile cpu.out
// ./validator.test -test.bench=. -test.cpuprofile=cpu.prof
// go tool pprof validator.test cpu.prof
//
//
// go test -memprofile mem.out

func TestInt(t *testing.T) {

	type TestInt struct {
		Int              int
		Int8             int8
		Int16            int16
		Int32            int32
		Int64            int64
		IntPtr           *int
		Int8Ptr          *int8
		Int16Ptr         *int16
		Int32Ptr         *int32
		Int64Ptr         *int64
		IntArray         []int
		IntPtrArray      []*int
		IntArrayArray    [][]int
		IntPtrArrayArray [][]*int
		IntMap           map[int]int
		IntPtrMap        map[*int]*int
	}

	values := url.Values{
		"Int":                    []string{"3"},
		"Int8":                   []string{"3"},
		"Int16":                  []string{"3"},
		"Int32":                  []string{"3"},
		"Int64":                  []string{"3"},
		"IntPtr":                 []string{"3"},
		"Int8Ptr":                []string{"3"},
		"Int16Ptr":               []string{"3"},
		"Int32Ptr":               []string{"3"},
		"Int64Ptr":               []string{"3"},
		"IntArray":               []string{"1", "2", "3"},
		"IntPtrArray[0]":         []string{"1"},
		"IntPtrArray[2]":         []string{"3"},
		"IntArrayArray[0][0]":    []string{"1"},
		"IntArrayArray[0][2]":    []string{"3"},
		"IntArrayArray[2][0]":    []string{"1"},
		"IntPtrArrayArray[0][0]": []string{"1"},
		"IntPtrArrayArray[0][2]": []string{"3"},
		"IntPtrArrayArray[2][0]": []string{"1"},
		"IntMap[1]":              []string{"3"},
		"IntPtrMap[1]":           []string{"3"},
	}

	var test TestInt

	test.IntArray = make([]int, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Int, int(3))
	Equal(t, test.Int8, int8(3))
	Equal(t, test.Int16, int16(3))
	Equal(t, test.Int32, int32(3))
	Equal(t, test.Int64, int64(3))

	Equal(t, *test.IntPtr, int(3))
	Equal(t, *test.Int8Ptr, int8(3))
	Equal(t, *test.Int16Ptr, int16(3))
	Equal(t, *test.Int32Ptr, int32(3))
	Equal(t, *test.Int64Ptr, int64(3))

	Equal(t, len(test.IntArray), 4)
	Equal(t, test.IntArray[0], int(1))
	Equal(t, test.IntArray[1], int(2))
	Equal(t, test.IntArray[2], int(3))
	Equal(t, test.IntArray[3], int(0))

	Equal(t, len(test.IntPtrArray), 3)
	Equal(t, *test.IntPtrArray[0], int(1))
	Equal(t, test.IntPtrArray[1], nil)
	Equal(t, *test.IntPtrArray[2], int(3))

	Equal(t, len(test.IntArrayArray), 3)
	Equal(t, len(test.IntArrayArray[0]), 3)
	Equal(t, len(test.IntArrayArray[1]), 0)
	Equal(t, len(test.IntArrayArray[2]), 1)
	Equal(t, test.IntArrayArray[0][0], int(1))
	Equal(t, test.IntArrayArray[0][1], int(0))
	Equal(t, test.IntArrayArray[0][2], int(3))
	Equal(t, test.IntArrayArray[2][0], int(1))

	Equal(t, len(test.IntPtrArrayArray), 3)
	Equal(t, len(test.IntPtrArrayArray[0]), 3)
	Equal(t, len(test.IntPtrArrayArray[1]), 0)
	Equal(t, len(test.IntPtrArrayArray[2]), 1)
	Equal(t, *test.IntPtrArrayArray[0][0], int(1))
	Equal(t, test.IntPtrArrayArray[0][1], nil)
	Equal(t, *test.IntPtrArrayArray[0][2], int(3))
	Equal(t, *test.IntPtrArrayArray[2][0], int(1))

	Equal(t, len(test.IntMap), 1)
	Equal(t, len(test.IntPtrMap), 1)

	v, ok := test.IntMap[1]
	Equal(t, ok, true)
	Equal(t, v, 3)
}

// func TestInt8(t *testing.T) {

// 	type intStruct struct {
// 		Int8Field    int8
// 		Int8PtrField *int8
// 	}

// 	tests := []struct {
// 		values   url.Values
// 		expected int8
// 		isPtr    bool
// 	}{
// 		{
// 			values:   url.Values{"Int8Field": []string{"7"}},
// 			expected: 7,
// 		},
// 		{
// 			values:   url.Values{"Int8PtrField": []string{"6"}},
// 			expected: 6,
// 			isPtr:    true,
// 		},
// 	}

// 	decoder := NewDecoder()

// 	var test intStruct
// 	var val int8

// 	for i, tt := range tests {
// 		decoder.Decode(&test, tt.values)

// 		if tt.isPtr {
// 			if test.Int8PtrField == nil {
// 				t.Errorf("Idx: %d Expected '%d' Got '%d'", i, tt.expected, test.Int8PtrField)
// 				continue
// 			}
// 			val = *test.Int8PtrField
// 		} else {
// 			val = test.Int8Field
// 		}

// 		if val != tt.expected {
// 			t.Errorf("Idx: %d Expected '%d' Got '%d'", i, tt.expected, val)
// 		}
// 	}
// 	// values := url.Values{
// 	// 	"Int8Field": []string{"5"},
// 	// }

// 	// fmt.Println(test.Int8Field)
// 	// type mm map[string]*intStruct
// 	// // m := map[int]string{}
// 	// m := make(mm)
// 	// fmt.Println(reflect.ValueOf(m).Kind())
// }

// func TestStraighUpArray(t *testing.T) {

// 	type Array struct {
// 		MyArray []int8
// 	}

// 	type ArrayPtr struct {
// 		MyArray []*int8
// 	}

// 	values := url.Values{"MyArray": []string{"0", "1", "2"}}

// 	decoder := NewDecoder()

// 	var test Array
// 	var test2 ArrayPtr

// 	// test2.MyArray = make([]*int8, 1, 1)

// 	decoder.Decode(&test, values)
// 	decoder.Decode(&test2, values)

// 	// fmt.Println("Test:", test)
// 	// fmt.Println("Test 2:", test2)
// 	// fmt.Println(*test2.MyArray[0])
// 	// fmt.Println(*test2.MyArray[1])
// 	// fmt.Println(*test2.MyArray[2])

// }

// func TestArrayNumbered(t *testing.T) {

// 	type Array struct {
// 		MyArray []int8
// 	}

// 	type ArrayPtr struct {
// 		MyArray []*int8
// 	}

// 	values := url.Values{"MyArray[0]": []string{"32"}, "MyArray[2]": []string{"33"}}

// 	decoder := NewDecoder()

// 	var test Array
// 	var test2 ArrayPtr

// 	decoder.Decode(&test, values)
// 	decoder.Decode(&test2, values)

// 	// fmt.Println("Test:", test)
// 	// fmt.Println("Test 2:", test2)
// 	// fmt.Println(*test2.MyArray[0])
// 	// fmt.Println(test2.MyArray[1])
// 	// fmt.Println(*test2.MyArray[2])

// }

// func TestArrayOfArray(t *testing.T) {

// 	type Array struct {
// 		MyArray [][]int8
// 	}

// 	type ArrayPtr struct {
// 		MyArray [][]*int8
// 	}

// 	values := url.Values{"MyArray[0][0]": []string{"32"}, "MyArray[0][1]": []string{"35"}, "MyArray[2][0]": []string{"33"}}

// 	decoder := NewDecoder()

// 	var test Array
// 	var test2 ArrayPtr

// 	decoder.Decode(&test, values)
// 	decoder.Decode(&test2, values)

// 	// fmt.Println("Test:", test)
// 	// fmt.Println("Test 2:", test2)
// }

// func TestString(t *testing.T) {

// 	type stringStruct struct {
// 		StringField    string
// 		StringPtrField *string
// 	}

// 	tests := []struct {
// 		values   url.Values
// 		expected string
// 		isPtr    bool
// 	}{
// 		{
// 			values:   url.Values{"StringField": []string{"7"}},
// 			expected: "7",
// 		},
// 		{
// 			values:   url.Values{"StringPtrField": []string{"6"}},
// 			expected: "6",
// 			isPtr:    true,
// 		},
// 	}

// 	decoder := NewDecoder()

// 	var test stringStruct
// 	var val string

// 	for i, tt := range tests {
// 		decoder.Decode(&test, tt.values)

// 		if tt.isPtr {
// 			if test.StringPtrField == nil {
// 				t.Errorf("Idx: %d Expected '%s' Got '%v'", i, tt.expected, test.StringPtrField)
// 				continue
// 			}
// 			val = *test.StringPtrField
// 		} else {
// 			val = test.StringField
// 		}

// 		if val != tt.expected {
// 			t.Errorf("Idx: %d Expected '%s' Got '%s'", i, tt.expected, val)
// 		}
// 	}
// 	// values := url.Values{
// 	// 	"Int8Field": []string{"5"},
// 	// }

// 	// fmt.Println(test.Int8Field)
// 	// type mm map[string]*intStruct
// 	// // m := map[int]string{}
// 	// m := make(mm)
// 	// fmt.Println(reflect.ValueOf(m).Kind())
// }

// func TestArrayStructString(t *testing.T) {

// 	type Phone struct {
// 		Number string
// 	}

// 	type User struct {
// 		Name         string
// 		PhoneNumbers []Phone
// 		ID           bson.ObjectId
// 	}

// 	values := url.Values{"ID": []string{"bson.ID"}, "Name": []string{"Joey Bloggs"}, "PhoneNumbers[0].Number": []string{"1(111)111-1111"}, "PhoneNumbers[1].Number": []string{"9(999)999-9999"}}

// 	decoder := NewDecoder()

// 	var test User

// 	decoder.Decode(&test, values)
// 	// fmt.Println(test.ID)
// 	// fmt.Println(test.ID.Hex())
// 	// fmt.Println("Test Phone:", test)
// }

// func TestArrayStructStringArrayString(t *testing.T) {
// 	type Home struct {
// 		Address string
// 	}

// 	type Phone struct {
// 		Number string
// 		Homes  []Home
// 	}

// 	type User struct {
// 		Name         string
// 		PhoneNumbers []Phone
// 	}

// 	values := url.Values{"Name": []string{"Joey Bloggs"}, "PhoneNumbers[0].Number": []string{"1(111)111-1111"}, "PhoneNumbers[1].Number": []string{"9(999)999-9999"}, "PhoneNumbers[0].Homes[0].Address": []string{"Beaumont"}}

// 	decoder := NewDecoder()

// 	var test User

// 	decoder.Decode(&test, values)

// 	// fmt.Println("Test Phone2:", test)
// }

// func TestBool(t *testing.T) {

// 	type boolStruct struct {
// 		OK    bool
// 		OKPtr *bool
// 	}

// 	values := url.Values{"OKPtr": []string{"true"}, "OK": []string{"t"}}

// 	decoder := NewDecoder()

// 	var test boolStruct

// 	decoder.Decode(&test, values)

// 	// fmt.Println("Test Bool:", test)
// }

// func TestFloat(t *testing.T) {

// 	type floatStruct struct {
// 		Float    float64
// 		FloatPtr *float64
// 	}

// 	values := url.Values{"Float": []string{"1.3333"}, "FloatPtr": []string{"13.546"}}

// 	decoder := NewDecoder()

// 	var test floatStruct

// 	decoder.Decode(&test, values)

// 	// fmt.Println("Test Float:", test)
// }

// func TestMap(t *testing.T) {

// 	type mapStruct struct {
// 		MapStringInt       map[string]int
// 		MapStringPtrString map[*string]string
// 		MapPrtString       *map[string]string
// 	}

// 	values := url.Values{"MapStringInt[key1]": []string{"3"}, "MapStringInt[key2]": []string{"5"}, "MapStringPtrString[ptrkey]": []string{"13"}, "MapPrtString[mpkey]": []string{"mpvalue"}}

// 	decoder := NewDecoder()

// 	var test mapStruct

// 	decoder.Decode(&test, values)

// 	// fmt.Println("Test Map:", test)

// 	// for k, v := range test.MapStringPtrString {
// 	// 	fmt.Println(*k, v)
// 	// }
// }

// func TestCustomType(t *testing.T) {

// 	type customStruct struct {
// 		Time    time.Time
// 		TimePtr *time.Time
// 	}

// 	values := url.Values{"Time": []string{"2016-01-02"}, "TimePtr": []string{"2017-01-02"}}

// 	decoder := NewDecoder()
// 	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
// 		return time.Parse("2006-01-02", vals[0])
// 	}, time.Time{})

// 	var test customStruct

// 	errs := decoder.Decode(&test, values)
// 	if errs != nil {
// 		t.Error("ERRORS!:", errs)
// 	}

// 	// fmt.Println("Test Custom Type:", test)
// }

// func TestBench(t *testing.T) {

// 	values := url.Values{
// 		"Nest.Children[0].ID":   []string{"joeybloggs_id"},
// 		"Nest.Children[0].Name": []string{"Joeybloggs"},
// 		"String":                []string{"golang is very fun"},
// 		"Slice[0]":              []string{"1"},
// 		"Slice[1]":              []string{"2"},
// 		"Slice[2]":              []string{"3"},
// 		"Slice[3]":              []string{"4"},
// 		"Bool":                  []string{"true"},
// 	}

// 	// values := url.Values{
// 	// 	"Nest.Children[0].ID":   []string{"joeybloggs_id"},
// 	// 	"Nest.Children[0].Name": []string{"Joeybloggs"},
// 	// 	"String":                []string{"golang is very fun"},
// 	// 	"Slice":                 []string{"1", "2", "3", "4"},
// 	// 	"Bool":                  []string{"true"},
// 	// }

// 	// type BenchFormamSchema struct {
// 	// 	Nest struct {
// 	// 		Children []struct {
// 	// 			ID   string
// 	// 			Name string
// 	// 		}
// 	// 	}
// 	// 	String string
// 	// 	Slice  []int
// 	// 	Bool   bool
// 	// }

// 	decoder := NewDecoder()

// 	test := new(BenchFormamSchema)

// 	decoder.Decode(test, values)

// 	// fmt.Printf("Test Bench: %#v", test)

// 	Equal(t, len(test.Nest.Children), 1)
// 	Equal(t, test.Nest.Children[0].ID, "joeybloggs_id")
// 	Equal(t, test.Nest.Children[0].Name, "Joeybloggs")
// 	Equal(t, test.String, "golang is very fun")
// 	Equal(t, test.Slice[0], int(1))
// 	Equal(t, test.Slice[1], int(2))
// 	Equal(t, test.Slice[2], int(3))
// 	Equal(t, test.Slice[3], int(4))
// 	Equal(t, test.Bool, true)
// }

// type BenchFormamSchema struct {
// 	Nest struct {
// 		Children []struct {
// 			ID   string
// 			Name string
// 		}
// 	}
// 	String string
// 	Slice  []int
// 	Bool   bool
// }

// func BenchmarkAssemblerTest2Parallel(b *testing.B) {

// 	values := url.Values{
// 		"Nest.Children[0].ID":   []string{"joeybloggs_id"},
// 		"Nest.Children[0].Name": []string{"Joeybloggs"},
// 		"String":                []string{"golang is very fun"},
// 		"Slice[0]":              []string{"1"},
// 		"Slice[1]":              []string{"2"},
// 		"Slice[2]":              []string{"3"},
// 		"Slice[3]":              []string{"4"},
// 		"Bool":                  []string{"true"},
// 	}

// 	ass := NewDecoder()

// 	b.ReportAllocs()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			test := new(BenchFormamSchema)
// 			if errs := ass.Decode(test, values); errs != nil {
// 				b.Error(errs)
// 			}
// 		}
// 	})
// }
