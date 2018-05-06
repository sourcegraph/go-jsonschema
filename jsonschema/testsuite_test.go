package jsonschema_test

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/sourcegraph/go-jsonschema/internal/jsonschematestsuite"
	"github.com/sourcegraph/go-jsonschema/internal/testutil"
)

func TestJSONUnmarshalMarshal(t *testing.T) {
	// TODO(sqs): Make these tests work.
	skip := map[string]struct{}{
		"TestJSONUnmarshalMarshal/additionalItems/additionalItems_as_schema":              struct{}{},
		"TestJSONUnmarshalMarshal/additionalItems/additionalItems_are_allowed_by_default": struct{}{},
		"TestJSONUnmarshalMarshal/const/const_with_null":                                  struct{}{},
		"TestJSONUnmarshalMarshal/ref/escaped_pointer_ref":                                struct{}{},
		"TestJSONUnmarshalMarshal/required/required_with_empty_array":                     struct{}{},
	}

	files, err := jsonschematestsuite.Files("../internal")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name, "optional"+string(os.PathSeparator)) {
			continue
		}
		t.Run(f.Name, func(t *testing.T) {
			f.ReadT(t)
			for _, g := range f.Groups {
				t.Run(g.Description, func(t *testing.T) {
					if _, ok := skip[t.Name()]; ok {
						t.Skip()
					}

					marshaled, err := json.Marshal(g.Schema)
					if err != nil {
						t.Fatal(err)
					}

					marshaled = testutil.CanonicalJSON(marshaled)
					data := testutil.CanonicalJSON(g.RawSchema)
					if !bytes.Equal(marshaled, data) {
						t.Errorf("got != want\n\ngot:  %s\nwant: %s", marshaled, data)
					}
				})
			}
		})
	}
}
