package compiler

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	objectwithprops "github.com/sourcegraph/go-jsonschema/compiler/testdata/object-with-props"
)

func TestAdditionalProperties(t *testing.T) {
	tests := []struct {
		json  string
		value objectwithprops.ObjectWithProps
	}{
		{
			json: `{"a":"A","b":"B","p0":"0","p1":1,"p2":{"a":"A2"}}`,
			value: objectwithprops.ObjectWithProps{
				P0:         "0",
				P1:         1,
				P2:         &objectwithprops.P2{A: "A2"},
				Additional: map[string]any{"a": "A", "b": "B"},
			},
		},
		{
			json: `{"a":"A","b":"B","p1":0}`,
			value: objectwithprops.ObjectWithProps{
				P0:         "",  // omitempty
				P1:         0,   // no omitempty
				P2:         nil, // omitempty
				Additional: map[string]any{"a": "A", "b": "B"},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotJSON, err := json.Marshal(test.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(gotJSON) != test.json {
				t.Errorf("got %s, want %s", gotJSON, test.json)
			}

			var gotValue objectwithprops.ObjectWithProps
			if err := json.Unmarshal(gotJSON, &gotValue); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(gotValue, test.value) {
				t.Errorf("got %+v, want %+v", gotValue, test.value)
			}
		})
	}
}
