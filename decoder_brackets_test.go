package form

import (
	"fmt"
	"net/url"
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestDecoderMalformedBracketKeys(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want string
	}{
		{
			name: "missing_start",
			key:  "a]",
			want: "invalid formatting for key 'a]' missing '[' bracket",
		},
		{
			name: "missing_end",
			key:  "a[",
			want: "invalid formatting for key 'a[' missing ']' bracket",
		},
		{
			name: "nested_missing_end",
			key:  "Phone[0.Number",
			want: "invalid formatting for key 'Phone[0.Number' missing ']' bracket",
		},
		{
			name: "nested_missing_start",
			key:  "Phone0].Number",
			want: "invalid formatting for key 'Phone0].Number' missing '[' bracket",
		},
		{
			name: "repeated_missing_end",
			key:  "Phone[[0.Number",
			want: "invalid formatting for key 'Phone[[0.Number' missing ']' bracket",
		},
		{
			name: "repeated_missing_start",
			key:  "Phone0]].Number",
			want: "invalid formatting for key 'Phone0]].Number' missing '[' bracket",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dst struct{ Items []string }
			values := url.Values{tt.key: []string{"bad"}}

			err := NewDecoder().Decode(&dst, values)
			errs, ok := err.(DecodeErrors)
			if !ok {
				t.Fatalf("expected DecodeErrors, got %T: %v", err, err)
			}
			Equal(t, len(errs), 1)
			if keyErr := errs[tt.key]; keyErr == nil {
				t.Fatalf("missing error for key %q: %v", tt.key, errs)
			} else {
				Equal(t, keyErr.Error(), tt.want)
			}
			Equal(t, dst.Items, []string(nil))
			Equal(t, values, url.Values{tt.key: []string{"bad"}})
		})
	}
}

func TestDecoderMalformedBracketKeysSkipPartialIndexes(t *testing.T) {
	var dst struct {
		Items  []string
		Array  [1]string
		Map    map[int]string
		Nested struct{ Items []string }
		Count  int
		Name   string
	}
	dst.Map = map[int]string{2: "keep"}
	values := url.Values{
		"Items[0]":             []string{"ok"},
		"Items[9][":            []string{"bad"},
		"Items[bad]x]":         []string{"bad"},
		"Array[0]":             []string{"array"},
		"Array[bad]x]":         []string{"bad"},
		"Map[1]":               []string{"map"},
		"Map[bad]x]":           []string{"bad"},
		"Nested.Items[0]":      []string{"nested"},
		"Nested.Items[99999][": []string{"bad"},
		"Count":                []string{"bad"},
		"Name":                 []string{"name"},
	}
	original := make(url.Values, len(values))
	for key, value := range values {
		original[key] = append([]string(nil), value...)
	}

	decoder := NewDecoder()
	err := decoder.Decode(&dst, values)
	errs, ok := err.(DecodeErrors)
	if !ok {
		t.Fatalf("expected DecodeErrors, got %T: %v", err, err)
	}
	Equal(t, len(errs), 6)
	for key, bracket := range map[string]string{
		"Items[9][":            "]",
		"Items[bad]x]":         "[",
		"Array[bad]x]":         "[",
		"Map[bad]x]":           "[",
		"Nested.Items[99999][": "]",
	} {
		if keyErr := errs[key]; keyErr == nil {
			t.Errorf("missing error for key %q: %v", key, errs)
		} else {
			Equal(t, keyErr.Error(), fmt.Sprintf("invalid formatting for key '%s' missing '%s' bracket", key, bracket))
		}
	}
	if countErr := errs["Count"]; countErr == nil {
		t.Errorf("missing conversion error: %v", errs)
	} else {
		Equal(t, countErr.Error(), "invalid integer value 'bad' type 'int' namespace 'Count'")
	}
	Equal(t, dst.Items, []string{"ok"})
	Equal(t, dst.Array, [1]string{"array"})
	Equal(t, dst.Map, map[int]string{1: "map", 2: "keep"})
	Equal(t, dst.Nested.Items, []string{"nested"})
	Equal(t, dst.Name, "name")
	Equal(t, values, original)

	var clean struct {
		Items []string
		Map   map[int]string
	}
	err = decoder.Decode(&clean, url.Values{
		"Items[0]": []string{"clean"},
		"Map[3]":   []string{"fresh"},
	})
	Equal(t, err, nil)
	Equal(t, clean.Items, []string{"clean"})
	Equal(t, clean.Map, map[int]string{3: "fresh"})
}

func TestDecoderMalformedBracketKeysRootCollections(t *testing.T) {
	tests := []struct {
		name   string
		dst    interface{}
		want   interface{}
		values url.Values
		key    string
	}{
		{
			name:   "slice",
			dst:    &[]string{},
			want:   &[]string{"ok"},
			values: url.Values{"[0]": []string{"ok"}, "[9][": []string{"bad"}},
			key:    "[9][",
		},
		{
			name:   "array",
			dst:    &[1]string{},
			want:   &[1]string{"ok"},
			values: url.Values{"[0]": []string{"ok"}, "[bad]x]": []string{"bad"}},
			key:    "[bad]x]",
		},
		{
			name:   "map",
			dst:    &map[int]string{2: "keep"},
			want:   &map[int]string{1: "ok", 2: "keep"},
			values: url.Values{"[1]": []string{"ok"}, "[bad]x]": []string{"bad"}},
			key:    "[bad]x]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewDecoder().Decode(tt.dst, tt.values)
			errs, ok := err.(DecodeErrors)
			if !ok {
				t.Fatalf("expected DecodeErrors, got %T: %v", err, err)
			}
			Equal(t, len(errs), 1)
			NotEqual(t, errs[tt.key], nil)
			Equal(t, tt.dst, tt.want)
		})
	}
}

func TestDecoderBracketKeyCompatibility(t *testing.T) {
	var dst struct {
		Literal string            `form:"Literal[[0]"`
		Items   []string          `form:"Items[]"`
		Map     map[string]string `form:"Map[["`
	}
	err := NewDecoder().Decode(&dst, url.Values{
		"Literal[[0]": []string{"literal"},
		"Items[]":     []string{"first", "second"},
		"Map[[[key]":  []string{"map"},
	})
	Equal(t, err, nil)
	Equal(t, dst.Literal, "literal")
	Equal(t, dst.Items, []string{"first", "second"})
	Equal(t, dst.Map, map[string]string{"key": "map"})

	var simple struct{ Name string }
	err = NewDecoder().Decode(&simple, url.Values{
		"Name": []string{"ok"},
		"a[":   []string{"ignored"},
	})
	Equal(t, err, nil)
	Equal(t, simple.Name, "ok")
}
