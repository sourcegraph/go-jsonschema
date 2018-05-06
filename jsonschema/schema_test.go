package jsonschema

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/sourcegraph/go-jsonschema/internal/testutil"
)

func TestSample(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/json-schema-draft-07-schema.json")
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

func strptr(s string) *string { return &s }
