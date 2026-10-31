package form

import (
	"net/url"
	"reflect"
	"testing"
)

func TestDecoderSliceMode(t *testing.T) {
	type input struct {
		Field []string `form:"field"`
	}

	tests := []struct {
		name   string
		mode   SliceMode
		values url.Values
		want   []string
	}{
		{
			name:   "default appends",
			values: url.Values{"field": {"new1", "new2"}},
			want:   []string{"old1", "old2", "new1", "new2"},
		},
		{
			name:   "replace unindexed values",
			mode:   SliceReplace,
			values: url.Values{"field": {"new1", "new2"}},
			want:   []string{"new1", "new2"},
		},
		{
			name: "replace leaves absent field alone",
			mode: SliceReplace,
			want: []string{"old1", "old2"},
		},
		{
			name:   "replace keeps indexed updates positional",
			mode:   SliceReplace,
			values: url.Values{"field[1]": {"new2"}},
			want:   []string{"old1", "new2"},
		},
		{
			name:   "replace applies indexed updates after unindexed values",
			mode:   SliceReplace,
			values: url.Values{"field": {"new1"}, "field[1]": {"new2"}},
			want:   []string{"new1", "new2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := input{Field: []string{"old1", "old2"}}
			decoder := NewDecoder()
			if tt.mode != SliceAppend {
				decoder.SetSliceMode(tt.mode)
			}
			if err := decoder.Decode(&got, tt.values); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got.Field, tt.want) {
				t.Fatalf("Field = %q, want %q", got.Field, tt.want)
			}
		})
	}
}
