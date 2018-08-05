package p

import (
	"github.com/sourcegraph/go-jsonschema/jsonschema"
)

type MetaSchemaRefs struct {
	A *jsonschema.Schema `json:"a,omitempty"`
	B *jsonschema.Schema `json:"b,omitempty"`
}
