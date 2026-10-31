package form

import (
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

type optionalMapKey int

func TestEncoderEmptyCustomMapKey(t *testing.T) {
	for _, empty := range [][]string{nil, {}} {
		t.Run(strconv.FormatBool(empty == nil), func(t *testing.T) {
			encoder := NewEncoder()
			encoder.RegisterCustomTypeFunc(func(value interface{}) ([]string, error) {
				key := value.(optionalMapKey)
				if key == 0 {
					return empty, nil
				}
				return []string{strconv.Itoa(int(key))}, nil
			}, optionalMapKey(0))
			value := struct{ Entries map[optionalMapKey]string }{
				Entries: map[optionalMapKey]string{0: "omit", 1: "keep"},
			}
			got, err := encoder.Encode(value)
			if err != nil {
				t.Fatal(err)
			}
			want := url.Values{"Entries[1]": {"keep"}}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Encode = %v, want %v", got, want)
			}
		})
	}
}
