package jsonschema

import (
	"net/url"
	"strconv"
	"testing"
)

func TestID_URI(t *testing.T) {
	tests := map[string]ID{
		"":        {},
		"a":       {Base: &url.URL{Path: "a"}},
		"/a#/b":   {Base: &url.URL{Path: "a"}, ReferenceTokens: []ReferenceToken{{Name: "b"}}},
		"/a#/c/b": {Base: &url.URL{Path: "a", Fragment: "/c"}, ReferenceTokens: []ReferenceToken{{Name: "b"}}},
	}
	i := 0
	for want, id := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if uri := id.String(); uri != want {
				t.Errorf("got %q, want %q", uri, want)
			}
		})
		i++
	}
}
