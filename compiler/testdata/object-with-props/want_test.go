package p

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestX(t *testing.T) {
	o := ObjectWithProps{
		A:          "y",
		Additional: map[string]interface{}{"b": "z"},
	}

	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"a":"y","b":"z"}`; string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}

	var o2 ObjectWithProps
	if err := json.Unmarshal(data, &o2); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(o, o2) {
		t.Errorf("got %+v, want %+v", o, o2)
	}
}
