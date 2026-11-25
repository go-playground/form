package form

import (
	"net/url"
	"testing"
)

func TestDecoderEmbeddedShadowingIssue31(t *testing.T) {
	type Embed struct {
		A string
	}
	var data struct {
		A string
		Embed
	}
	dec := NewDecoder()
	err := dec.Decode(&data, url.Values{"A": {"one"}})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data.A=%q data.Embed.A=%q", data.A, data.Embed.A)
	if data.A != "one" {
		t.Fatalf("parent field wrong: %q", data.A)
	}
	if data.Embed.A != "" {
		t.Fatalf("embedded field incorrectly set: %q", data.Embed.A)
	}
}
