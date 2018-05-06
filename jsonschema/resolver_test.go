package jsonschema

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestResolve(t *testing.T) {
	var (
		schemaA = &Schema{Title: strptr("a")}
		schemaC = &Schema{ID: strptr("c"), Title: strptr("c")}
	)

	tests := map[string]struct {
		scope     Scope
		ref       string
		wantFound *Schema // tested for pointer equality
		wantErr   error
	}{
		"empty": {},
		"simple": {
			scope: Scope{
				Schemas: []*Schema{
					{Definitions: &map[string]*Schema{"a": schemaA}},
				},
			},
			ref:       "#/definitions/a",
			wantFound: schemaA,
		},
		"base URI": {
			scope: Scope{
				Schemas: []*Schema{
					{Definitions: &map[string]*Schema{"a": schemaA}},
				},
				Base: &url.URL{Path: "b"},
			},
			ref:       "b#/definitions/a",
			wantFound: schemaA,
		},
		"explicit parent schema ID": {
			scope: Scope{
				Schemas: []*Schema{{
					ID: strptr("b2"),
					Definitions: &map[string]*Schema{
						"a": schemaA,
					}},
				},
				Base: &url.URL{Path: "b"},
			},
			ref:       "b2#/definitions/a",
			wantFound: schemaA,
		},
		"explicit schema ID": {
			scope: Scope{
				Schemas: []*Schema{{
					Definitions: &map[string]*Schema{
						"c": schemaC,
					}},
				},
				Base: &url.URL{Path: "b"},
			},
			ref:       "c",
			wantFound: schemaC,
		},
		"not found": {
			scope: Scope{
				Schemas: []*Schema{
					{Definitions: &map[string]*Schema{"a": schemaA}},
				},
			},
			ref:       "#/definitions/b",
			wantFound: nil,
		},
		"invalid URI in ID": {
			scope: Scope{
				Schemas: []*Schema{{ID: strptr("%")}},
			},
			wantErr: &url.Error{Op: "parse", URL: "%", Err: url.EscapeError("%")},
		},
		// TODO(sqs): Many cases are unimplemented.
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			u, err := url.Parse(test.ref)
			if err != nil {
				t.Fatal(err)
			}
			schema, err := Resolve(test.scope, u)
			if !reflect.DeepEqual(err, test.wantErr) {
				if test.wantErr == nil {
					t.Fatal(err)
				}
				t.Fatalf("got error %v, want %v", err, test.wantErr)
			}
			if err != nil {
				return
			}
			if schema != test.wantFound {
				t.Errorf("got schema != want schema\n\n%v", pretty.Diff(schema, test.wantFound))
			}
		})
	}
}
