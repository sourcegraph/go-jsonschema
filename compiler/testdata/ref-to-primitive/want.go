package p

// A
type A struct {
	B string `json:"b,omitempty"`
}

// RefToPrimitive
type RefToPrimitive struct {
	A string `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}
