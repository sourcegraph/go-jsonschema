package jsonschema

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/sourcegraph/go-jsonschema/internal/testutil"
)

func TestSample(t *testing.T) {
	data, err := os.ReadFile("../testdata/json-schema-draft-07-schema.json")
	if err != nil {
		t.Fatal(err)
	}

	var schema Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatal(err)
	}

	marshaled, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}

	marshaled = testutil.CanonicalJSON(marshaled)
	data = testutil.CanonicalJSON(data)
	if !bytes.Equal(marshaled, data) {
		t.Errorf("got %q, want %q", marshaled, data)
	}
}

func TestRaw(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		b, err := json.Marshal(Schema{Comment: strptr("c")})
		if err != nil {
			t.Fatal(err)
		}
		if want := `{"$comment":"c"}`; string(b) != want {
			t.Errorf("got %s, want %s", b, want)
		}
	})
	t.Run("unmarshal", func(t *testing.T) {
		input := []byte(`{"$comment":"c"}`)
		var o Schema
		if err := json.Unmarshal([]byte(input), &o); err != nil {
			t.Fatal(err)
		}
		if want := (Schema{Comment: strptr("c"), Raw: (*json.RawMessage)(&input)}); !reflect.DeepEqual(o, want) {
			t.Errorf("got %+v, want %+v", o, want)
		}
	})
}

func strptr(s string) *string { return &s }
