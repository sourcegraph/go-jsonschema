package jsonschema

import (
	"net/url"
)

// Scope is an environment in which to resolve references to JSON Schemas.
type Scope struct {
	Schemas []*Schema
	Base    *url.URL // TODO(sqs): Support for this base URI is not fully implemented.
}

// Resolve returns the JSON Schema referenced by ref in the scope, if it exists.
//
// NOTE(sqs): This does not work well and will probably be removed.
//
// TODO(sqs): Resolve does not yet support loading schemas that are referenced but not present in
// the scope.
func Resolve(scope Scope, ref *url.URL) (found *Schema, err error) {
	if scope.Base != nil {
		ref = scope.Base.ResolveReference(ref)
	}
	v := resolverVisitor{
		base: ID{Base: scope.Base},

		ref:   ref,
		found: &found,
		err:   &err,
	}
	for _, schema := range scope.Schemas {
		Walk(&v, schema)
		if found != nil || err != nil {
			break
		}
	}
	return found, err
}

type resolverVisitor struct {
	base ID

	ref   *url.URL
	found **Schema
	err   *error
}

// Visit implements Visitor.
func (v *resolverVisitor) Visit(schema *Schema, rel []ReferenceToken) (w Visitor) {
	if schema == nil || *v.found != nil || *v.err != nil {
		return nil
	}

	rw := *v // copy
	rw.base = rw.base.ResolveReference(rel)

	// The schema has 2 possible IDs here: (1) the ID we constructed by traversing from the base and
	// appending relative reference tokens, and (2) the ID that it itself sets in its "$id"
	// property. Check both for a match.

	if rw.base.String() == v.ref.String() {
		*v.found = schema
		return nil
	}

	if schema.ID != nil {
		u, err := url.Parse(*schema.ID)
		if err != nil {
			*v.err = err
			return nil
		}
		if v.base.Base != nil {
			u = v.base.Base.ResolveReference(u)
		}
		rw.base = ID{Base: u}
		if rw.base.String() == v.ref.String() {
			*v.found = schema
			return nil
		}
	}

	return &rw
}
