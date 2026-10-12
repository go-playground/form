package form

import (
	"net/url"
	"reflect"
	"testing"
)

func TestEncodeUnexportedEmbeddedOmitEmpty(t *testing.T) {
	type embedded struct {
		Value string
	}
	type outer struct {
		embedded `form:",omitempty"`
	}
	for _, value := range []string{"", "present"} {
		t.Run(value, func(t *testing.T) {
			values, err := NewEncoder().Encode(outer{embedded{Value: value}})
			if err != nil {
				t.Fatal(err)
			}
			want := url.Values{}
			if value != "" {
				want["Value"] = []string{value}
			}
			if !reflect.DeepEqual(values, want) {
				t.Errorf("Encode() = %v, want %v", values, want)
			}
		})
	}
}
