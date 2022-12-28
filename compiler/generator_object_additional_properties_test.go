package compiler

import (
	"encoding/json"
	"reflect"
	"testing"

	objectwithprops "github.com/sourcegraph/go-jsonschema/compiler/testdata/object-with-props"
)

func TestAdditionalProperties(t *testing.T) {
	o := objectwithprops.ObjectWithProps{
		A:          "1",
		B:          "2",
		Additional: map[string]any{"c": "3", "d": "4"},
	}

	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"a":"1","b":"2","c":"3","d":"4"}`; string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}

	var o2 objectwithprops.ObjectWithProps
	if err := json.Unmarshal(data, &o2); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(o, o2) {
		t.Errorf("got %+v, want %+v", o, o2)
	}
}
