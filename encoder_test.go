package form

import (
	"testing"

	. "gopkg.in/go-playground/assert.v1"
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

func TestEncoderInt(t *testing.T) {

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
		NoValue          int
		NoPtrValue       *int
	}

	i := int(3)
	i8 := int8(3)
	i16 := int16(3)
	i32 := int32(3)
	i64 := int64(3)

	zero := int(0)
	one := int(1)
	two := int(2)
	three := int(3)

	test := TestInt{
		Int:              i,
		Int8:             i8,
		Int16:            i16,
		Int32:            i32,
		Int64:            i64,
		IntPtr:           &i,
		Int8Ptr:          &i8,
		Int16Ptr:         &i16,
		Int32Ptr:         &i32,
		Int64Ptr:         &i64,
		IntArray:         []int{one, two, three},
		IntPtrArray:      []*int{&one, &two, &three},
		IntArrayArray:    [][]int{{one, zero, three}},
		IntPtrArrayArray: [][]*int{{&one, &zero, &three}},
		IntMap:           map[int]int{one: three, zero: two},
		IntPtrMap:        map[*int]*int{&one: &three, &zero: &two},
	}

	encoder := NewEncoder()
	values, errs := encoder.Encode(test)

	Equal(t, errs, nil)
	Equal(t, len(values), 25)

	val, ok := values["Int8"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int8"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int16"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int32"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int64"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int8"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int8Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int16Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int32Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Int64Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["IntArray"]
	Equal(t, ok, true)
	Equal(t, len(val), 3)
	Equal(t, val[0], "1")
	Equal(t, val[1], "2")
	Equal(t, val[2], "3")

	val, ok = values["IntPtrArray[0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["IntPtrArray[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["IntPtrArray[2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["IntArrayArray[0][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["IntArrayArray[0][1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "0")

	val, ok = values["IntArrayArray[0][2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["IntPtrArrayArray[0][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["IntPtrArrayArray[0][1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "0")

	val, ok = values["IntPtrArrayArray[0][2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["IntMap[0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["IntMap[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["IntPtrMap[0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["IntPtrMap[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["NoValue"]
	Equal(t, ok, true)
	Equal(t, val[0], "0")

	val, ok = values["NoPtrValue"]
	Equal(t, ok, false)
}

func TestEncoderUint(t *testing.T) {

	type TestUint struct {
		Uint              uint
		Uint8             uint8
		Uint16            uint16
		Uint32            uint32
		Uint64            uint64
		UintPtr           *uint
		Uint8Ptr          *uint8
		Uint16Ptr         *uint16
		Uint32Ptr         *uint32
		Uint64Ptr         *uint64
		UintArray         []uint
		UintPtrArray      []*uint
		UintArrayArray    [][]uint
		UintPtrArrayArray [][]*uint
		UintMap           map[uint]uint
		UintPtrMap        map[*uint]*uint
		NoValue           uint
		NoPtrValue        *uint
	}

	i := uint(3)
	i8 := uint8(3)
	i16 := uint16(3)
	i32 := uint32(3)
	i64 := uint64(3)

	zero := uint(0)
	one := uint(1)
	two := uint(2)
	three := uint(3)

	test := TestUint{
		Uint:              i,
		Uint8:             i8,
		Uint16:            i16,
		Uint32:            i32,
		Uint64:            i64,
		UintPtr:           &i,
		Uint8Ptr:          &i8,
		Uint16Ptr:         &i16,
		Uint32Ptr:         &i32,
		Uint64Ptr:         &i64,
		UintArray:         []uint{one, two, three},
		UintPtrArray:      []*uint{&one, &two, &three},
		UintArrayArray:    [][]uint{{one, zero, three}},
		UintPtrArrayArray: [][]*uint{{&one, &zero, &three}},
		UintMap:           map[uint]uint{one: three, zero: two},
		UintPtrMap:        map[*uint]*uint{&one: &three, &zero: &two},
	}

	encoder := NewEncoder()
	values, errs := encoder.Encode(test)

	Equal(t, errs, nil)
	Equal(t, len(values), 25)

	val, ok := values["Uint8"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint8"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint16"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint32"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint64"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint8"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint8Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint16Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint32Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["Uint64Ptr"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["UintArray"]
	Equal(t, ok, true)
	Equal(t, len(val), 3)
	Equal(t, val[0], "1")
	Equal(t, val[1], "2")
	Equal(t, val[2], "3")

	val, ok = values["UintPtrArray[0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["UintPtrArray[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["UintPtrArray[2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["UintArrayArray[0][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["UintArrayArray[0][1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "0")

	val, ok = values["UintArrayArray[0][2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["UintPtrArrayArray[0][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["UintPtrArrayArray[0][1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "0")

	val, ok = values["UintPtrArrayArray[0][2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["UintMap[0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["UintMap[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["UintPtrMap[0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["UintPtrMap[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["NoValue"]
	Equal(t, ok, true)
	Equal(t, val[0], "0")

	val, ok = values["NoPtrValue"]
	Equal(t, ok, false)
}

func TestEncoderString(t *testing.T) {

	type TestString struct {
		String              string
		StringPtr           *string
		StringArray         []string
		StringPtrArray      []*string
		StringArrayArray    [][]string
		StringPtrArrayArray [][]*string
		StringMap           map[string]string
		StringPtrMap        map[*string]*string
		NoValue             string
	}

	one := "1"
	two := "2"
	three := "3"

	test := TestString{
		String:              three,
		StringPtr:           &two,
		StringArray:         []string{one, "", three},
		StringPtrArray:      []*string{&one, nil, &three},
		StringArrayArray:    [][]string{{one, "", three}, nil, {one}},
		StringPtrArrayArray: [][]*string{{&one, nil, &three}, nil, {&one}},
		StringMap:           map[string]string{one: three, three: two},
		StringPtrMap:        map[*string]*string{&one: &three, &three: &two},
	}

	encoder := NewEncoder()
	values, errs := encoder.Encode(test)

	Equal(t, errs, nil)
	Equal(t, len(values), 17)

	val, ok := values["String"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["StringPtr"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["StringArray"]
	Equal(t, ok, true)
	Equal(t, len(val), 3)
	Equal(t, val[0], "1")
	Equal(t, val[1], "")
	Equal(t, val[2], "3")

	val, ok = values["StringPtr"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["StringPtrArray[0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["StringPtrArray[1]"]
	Equal(t, ok, false)

	val, ok = values["StringPtrArray[2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["StringPtrArray[2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["StringArrayArray[0][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["StringArrayArray[0][1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "")

	val, ok = values["StringArrayArray[0][2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["StringArrayArray[1][1]"]
	Equal(t, ok, false)

	val, ok = values["StringArrayArray[2][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["StringPtrArrayArray[0][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["StringPtrArrayArray[0][1]"]
	Equal(t, ok, false)

	val, ok = values["StringPtrArrayArray[0][2]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["StringPtrArrayArray[1][1]"]
	Equal(t, ok, false)

	val, ok = values["StringPtrArrayArray[2][0]"]
	Equal(t, ok, true)
	Equal(t, val[0], "1")

	val, ok = values["StringMap[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["StringMap[3]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["StringPtrMap[1]"]
	Equal(t, ok, true)
	Equal(t, val[0], "3")

	val, ok = values["StringPtrMap[3]"]
	Equal(t, ok, true)
	Equal(t, val[0], "2")

	val, ok = values["NoValue"]
	Equal(t, ok, true)
	Equal(t, val[0], "")
}
