package assembler

import (
	"fmt"
	"net/url"
	"testing"
)

func TestInt8(t *testing.T) {

	type intStruct struct {
		Int8Field    int8
		Int8PtrField *int8
	}

	tests := []struct {
		values   url.Values
		expected int8
		isPtr    bool
	}{
		{
			values:   url.Values{"Int8Field": []string{"7"}},
			expected: 7,
		},
		{
			values:   url.Values{"Int8PtrField": []string{"6"}},
			expected: 6,
			isPtr:    true,
		},
	}

	decoder := NewDecoder()

	var test intStruct
	var val int8

	for i, tt := range tests {
		decoder.Decode(&test, tt.values)

		if tt.isPtr {
			if test.Int8PtrField == nil {
				t.Errorf("Idx: %d Expected '%d' Got '%d'", i, tt.expected, test.Int8PtrField)
				continue
			}
			val = *test.Int8PtrField
		} else {
			val = test.Int8Field
		}

		if val != tt.expected {
			t.Errorf("Idx: %d Expected '%d' Got '%d'", i, tt.expected, val)
		}
	}
	// values := url.Values{
	// 	"Int8Field": []string{"5"},
	// }

	// fmt.Println(test.Int8Field)
	// type mm map[string]*intStruct
	// // m := map[int]string{}
	// m := make(mm)
	// fmt.Println(reflect.ValueOf(m).Kind())
}

func TestStraighUpArray(t *testing.T) {

	type Array struct {
		MyArray []int8
	}

	type ArrayPtr struct {
		MyArray []*int8
	}

	values := url.Values{"MyArray": []string{"0", "1", "2"}}

	decoder := NewDecoder()

	var test Array
	var test2 ArrayPtr

	decoder.Decode(&test, values)
	decoder.Decode(&test2, values)

	fmt.Println("Test:", test)
	fmt.Println("Test 2:", test2)
	fmt.Println(*test2.MyArray[0])
	fmt.Println(*test2.MyArray[1])
	fmt.Println(*test2.MyArray[2])

}
