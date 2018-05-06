package compiler

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
	"github.com/sourcegraph/go-jsonschema/jsonschema"
)

func TestParseSchema(t *testing.T) {
	schemaA := &jsonschema.Schema{
		Type: jsonschema.PrimitiveTypeList{jsonschema.ObjectType},
		Properties: &map[string]*jsonschema.Schema{
			"b": {Type: jsonschema.PrimitiveTypeList{jsonschema.StringType}},
		},
	}
	schemaC := &jsonschema.Schema{
		Type: jsonschema.PrimitiveTypeList{jsonschema.ObjectType},
		Properties: &map[string]*jsonschema.Schema{
			"c": {Type: jsonschema.PrimitiveTypeList{jsonschema.StringType}},
		},
	}
	schemaE := &jsonschema.Schema{
		ID:    strptr("e"),
		Type:  jsonschema.PrimitiveTypeList{jsonschema.ArrayType},
		Items: &jsonschema.SchemaOrSchemaList{Schema: schemaC},
	}
	schemaF := &jsonschema.Schema{
		Type: jsonschema.PrimitiveTypeList{jsonschema.ObjectType},
		Properties: &map[string]*jsonschema.Schema{
			"f": {Type: jsonschema.PrimitiveTypeList{jsonschema.StringType}},
		},
	}
	schemaRoot := &jsonschema.Schema{
		Title: strptr("root"),
		Type:  jsonschema.PrimitiveTypeList{jsonschema.ObjectType},
		Properties: &map[string]*jsonschema.Schema{
			"a": schemaA,
			"e": schemaE,
		},
		Definitions: &map[string]*jsonschema.Schema{
			"f": schemaF,
		},
	}

	locations, err := parseSchema(schemaRoot)
	if err != nil {
		t.Fatal(err)
	}
	want := map[*jsonschema.Schema]schemaLocation{
		schemaA: {rel: []jsonschema.ReferenceToken{{Name: "properties", Keyword: true}, {Name: "a"}}},
		schemaC: {
			rel: []jsonschema.ReferenceToken{{Name: "properties", Keyword: true}, {Name: "e"}, {Name: "items", Keyword: true}},
			id:  &jsonschema.ID{Base: &url.URL{Path: "e"}, ReferenceTokens: []jsonschema.ReferenceToken{{Name: "items", Keyword: true}}},
		},
		schemaE: {
			rel: []jsonschema.ReferenceToken{{Name: "properties", Keyword: true}, {Name: "e"}},
			id:  &jsonschema.ID{Base: &url.URL{Path: "e"}},
		},
		schemaF:    {rel: []jsonschema.ReferenceToken{{Name: "definitions", Keyword: true}, {Name: "f"}}},
		schemaRoot: {rel: []jsonschema.ReferenceToken{}},
	}
	if !reflect.DeepEqual(locations, want) {
		// Simplify output.
		labels := map[*jsonschema.Schema]string{
			schemaA:    "schemaA",
			schemaC:    "schemaC",
			schemaE:    "schemaE",
			schemaF:    "schemaF",
			schemaRoot: "schemaRoot",
		}
		simplify := func(locations map[*jsonschema.Schema]schemaLocation) map[string][]string {
			m := make(map[string][]string, len(locations))
			unknown := 0
			for schema, location := range locations {
				label := labels[schema]
				if label == "" {
					label = fmt.Sprintf("unknown%d", unknown)
					unknown++
					b, _ := json.Marshal(schema)
					t.Errorf("no label for schema (using %q): %s", label, b)
				}
				m[label] = []string{jsonschema.EncodeReferenceTokens(location.rel)}
				if location.id != nil {
					m[label] = append(m[label], location.id.String())
				}
			}
			return m
		}
		t.Errorf("got locations != want locations\n%s", strings.Join(pretty.Diff(simplify(locations), simplify(want)), "\n"))
	}
}

func strptr(s string) *string { return &s }
