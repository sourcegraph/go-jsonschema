/////////// +2build ignore

package compiler

import (
	"encoding/json"
	"reflect"
	"testing"

	testdata_oneof "github.com/sourcegraph/go-jsonschema/compiler/testdata/oneOf"
	"github.com/sourcegraph/go-jsonschema/internal/testutil"
)

// TestOneOf depends on the generated ./testdata/oneOf/want.go file, which you can overwrite with
// the latest generated code by running `go test -test.write-want`.
func TestOneOf(t *testing.T) {
	tests := map[string]struct {
		data string
		want testdata_oneof.OneOf
	}{
		"not set": {
			data: `{}`,
			want: testdata_oneof.OneOf{A: nil},
		},
		"one b": {
			data: `{"a":[{"type":"b","b":"x"}]}`,
			want: testdata_oneof.OneOf{
				A: []testdata_oneof.A{
					{B: &testdata_oneof.B{Type: "b", B: "x"}},
				},
			},
		},
		"one c": {
			data: `{"a":[{"type":"c","c":true}]}`,
			want: testdata_oneof.OneOf{
				A: []testdata_oneof.A{
					{C: &testdata_oneof.C{Type: "c", C: true}},
				},
			},
		},
		"one d": {
			data: `{"a":[{"type":"d","d":1}]}`,
			want: testdata_oneof.OneOf{
				A: []testdata_oneof.A{
					{D: &testdata_oneof.D{Type: "d", D: 1}},
				},
			},
		},
		"multiple": {
			data: `{"a":[{"type":"b","b":"x"},{"type":"c","c":true},{"type":"d","d":1}]}`,
			want: testdata_oneof.OneOf{
				A: []testdata_oneof.A{
					{B: &testdata_oneof.B{Type: "b", B: "x"}},
					{C: &testdata_oneof.C{Type: "c", C: true}},
					{D: &testdata_oneof.D{Type: "d", D: 1}},
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got testdata_oneof.OneOf
			if err := json.Unmarshal([]byte(test.data), &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Unmarshal: got (%+v) != want (%+v)", got, test.want)
			}
			data, err := json.Marshal(got)
			if err != nil {
				t.Fatal(err)
			}
			data = testutil.CanonicalJSON(data)
			test.data = string(testutil.CanonicalJSON([]byte(test.data)))
			if string(data) != test.data {
				t.Errorf("Marshal: got != want\n got %s\nwant %s", data, test.data)
			}
		})
	}
}
